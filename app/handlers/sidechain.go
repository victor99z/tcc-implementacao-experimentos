package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/joho/godotenv"
	"github.com/victor99z/blockchain-app/models"
	"github.com/victor99z/blockchain-app/utils"
)

type SidechainHandler struct {
	Conn    *ethclient.Client
	ChainId *big.Int
	Storage *models.StorageOnSidechain
	Auth    *bind.TransactOpts
	Ipfs    *shell.Shell
}

func NewSidechainHandler() *SidechainHandler {
	godotenv.Load()

	sh := shell.NewShell(os.Getenv("IPFS_HOST"))
	conn, err := ethclient.Dial(os.Getenv("ETH_HOST"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	chainId := utils.GetChainID()

	alreadyDeployedContract := utils.ReadDeployedContract()

	storage, _ := models.NewStorageOnSidechain(common.HexToAddress(alreadyDeployedContract.Address), conn)

	doctor, _ := keystore.DecryptKey([]byte(doctor_keystore), os.Getenv("KEYSTORE_PASSWORD"))
	auth, _ := bind.NewKeyedTransactorWithChainID(doctor.PrivateKey, chainId)

	auth.GasLimit = 300000000
	auth.GasPrice = big.NewInt(0)
	auth.Value = big.NewInt(0)

	return &SidechainHandler{
		conn,
		chainId,
		storage,
		auth,
		sh,
	}
}

func (s *SidechainHandler) GetRecords(c *fiber.Ctx) error {

	id := c.Params("id")

	records, err := s.Storage.GetAllRecords(
		nil,
		common.HexToAddress(id),
	)

	if err != nil {
		return c.Status(500).JSON(err)
	}

	if len(records) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No records found",
		})
	}

	return c.JSON(records)
}

func (s *SidechainHandler) GetRecordCounterByAddr(c *fiber.Ctx) error {
	pacient_addr := c.Params("id")
	counter, _ := s.Storage.GetRecordCounter(nil, common.HexToAddress(pacient_addr))

	return c.JSON(fiber.Map{
		"address":         pacient_addr,
		"records_counter": counter,
	})
}

func (s *SidechainHandler) GetRecord(c *fiber.Ctx) error {
	pacient_addr := c.Params("id")
	idx := c.Params("index")

	index, _ := strconv.ParseInt(idx, 10, 64)
	tx, err := s.Storage.GetRecord(nil, common.HexToAddress(pacient_addr), big.NewInt(index))
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	readerFileIpfs, err := s.Ipfs.Cat(tx.Cid)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
	defer readerFileIpfs.Close()

	bytes, err := ioutil.ReadAll(readerFileIpfs)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
	_ = os.WriteFile(`./tmp/`+tx.Cid, bytes, 0644)

	return c.JSON(fiber.Map{
		"pacient":     tx.Recipient.Hex(),
		"doctor":      tx.Doctor.Hex(),
		"institution": tx.Institution.Hex(),
		"cid":         tx.Cid,
		"create_at":   tx.CreateAt,
	})
}

func PinFileCluster(cid string) {
	url := os.Getenv("IPFS_CLUSTER_API") + "/pins/" + cid
	agent := fiber.Post(url)
	statusCode, _, errs := agent.Bytes()
	if errs != nil {
		log.Default().Print("Error to pin file ", cid)
	}
	if statusCode != fiber.StatusOK {
		log.Default().Print("Error to pin file [status code] ", cid)

	}
}

func (s *SidechainHandler) SaveRecord(c *fiber.Ctx) error {
	payload := struct {
		Pacient     string
		Doctor      string
		Institution string
		File        string
	}{}

	jsonData := c.FormValue("json")

	if err := json.Unmarshal([]byte(jsonData), &payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Parse file part
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse file",
		})
	}
	// Open the uploaded file
	fileHeader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot open file",
		})
	}
	defer fileHeader.Close()

	// Convert byte slice to a byte string
	cid, err := s.Ipfs.Add(fileHeader)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err,
		})
	}

	tx, err := s.Storage.Save(s.Auth,
		common.HexToAddress(payload.Doctor),
		common.HexToAddress(payload.Pacient),
		common.HexToAddress(payload.Institution),
		cid,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	PinFileCluster(cid)

	return c.JSON(fiber.Map{
		"txHash": tx.Hash().Hex(),
		"cid":    cid,
	})
}
