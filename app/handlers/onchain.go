package handlers

import (
	"encoding/json"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/victor99z/blockchain-app/models"
	"github.com/victor99z/blockchain-app/utils"
)

type OnchainHandler struct {
	Conn    *ethclient.Client
	ChainId *big.Int
	Storage *models.StorageOnChain
	Auth    *bind.TransactOpts
}

const doctor_keystore = `{"address":"823630d906c13965478da03ae24004ad5ff7c691","crypto":{"cipher":"aes-128-ctr","ciphertext":"2ce99513d1c6a0cbad8077e2dcf8a6fa44007c56db2d222c13d0290f0227565f","cipherparams":{"iv":"650dc162d1fa5372ab2a6f3ea5cbbb41"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"19b2c52938abb5db1a3f74e4a9f7ce244e493c4847d765ec37e15688eb998065"},"mac":"0f992004c53db2820d785778ccd37c84c9d9d042f6b0156f5c70db9e2ad3a42b"},"id":"dd92eda9-a06c-4ddb-889a-670698d04d40","version":3}`
const institution_keystore = `{"address":"688104018a9fe62a555f2835d614df7e23b2312d","crypto":{"cipher":"aes-128-ctr","ciphertext":"eaab0254b6815316eb2e8228f8dbcde1f3954bcfe324c1a9487525f1e7f87514","cipherparams":{"iv":"b6978ac4fa2ea4a5688fd84beafd3ba2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"c9106a2e1850b960fdb0b9b0e00e5d7acc138690c18878a035dee29b5f5bb14a"},"mac":"88bfc7a1d93da184b79dc21ff623523c161cfeca2d4b1a3b96fcf9b5040afeb2"},"id":"a63a7de7-c9c3-4026-a4c7-c3da826a37ff","version":3}`
const pacient_keystore = `{"address":"881466d0000b87cc0ebe032adeec41d1cc52102f","crypto":{"cipher":"aes-128-ctr","ciphertext":"15673b83b28621afe90c4f3bda47526a3090f72996aa0d6b45868214c5203f2e","cipherparams":{"iv":"df18cc787a39aaffb84439c83dc7774c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7f392369e14002780615f22dee33fedb8bdde72ec637883d9c5a77a433de7224"},"mac":"310a9a2db226e8b1b3c411e3af84e2de2f5031f38d4e0508080536902c94e3cf"},"id":"7de64149-e8fc-447e-a226-5d791d5670f0","version":3}`

func NewOnchainHandler() *OnchainHandler {
	godotenv.Load()

	conn, err := ethclient.Dial(os.Getenv("ETH_HOST"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	chainId := utils.GetChainID()

	alreadyDeployedContract := utils.ReadDeployedContract()

	storage, _ := models.NewStorageOnChain(common.HexToAddress(alreadyDeployedContract.Address), conn)

	doctor, _ := keystore.DecryptKey([]byte(doctor_keystore), os.Getenv("KEYSTORE_PASSWORD"))
	auth, _ := bind.NewKeyedTransactorWithChainID(doctor.PrivateKey, chainId)

	auth.GasLimit = 300000000
	auth.GasPrice = big.NewInt(0)
	auth.Value = big.NewInt(0)

	return &OnchainHandler{
		conn,
		chainId,
		storage,
		auth,
	}
}

func (s *OnchainHandler) GetRecords(c *fiber.Ctx) error {

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

func (s *OnchainHandler) GetRecordCounterByAddr(c *fiber.Ctx) error {
	pacient_addr := c.Params("id")
	counter, _ := s.Storage.GetRecordCounter(nil, common.HexToAddress(pacient_addr))

	return c.JSON(fiber.Map{
		"address":         pacient_addr,
		"records_counter": counter,
	})
}

func (s *OnchainHandler) GetRecord(c *fiber.Ctx) error {
	pacient_addr := c.Params("id")
	idx := c.Params("index")

	index, _ := strconv.ParseInt(idx, 10, 64)
	tx, err := s.Storage.GetRecord(nil, common.HexToAddress(pacient_addr), big.NewInt(index))
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	_ = os.WriteFile(`./tmp/`+time.Now().Format(time.RFC3339), []byte(tx.Data), 0644)

	return c.JSON(fiber.Map{
		"pacient":     tx.Recipient.Hex(),
		"doctor":      tx.Doctor.Hex(),
		"institution": tx.Institution.Hex(),
		"create_at":   tx.CreateAt,
	})
}

// Quem realiza inserção de registros é o médico
// Ou seja, apenas para o médico é necessario um KeydTransactor
func (s *OnchainHandler) SaveRecord(c *fiber.Ctx) error {

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

	// readFile, err := os.ReadFile(payload.File)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
	// 		"message": "Failed to read file",
	// 	})
	// }

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

	// Read file into byte slice
	fileBytes, err := io.ReadAll(fileHeader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot read file",
		})
	}

	// Convert byte slice to a byte string
	fileByteString := string(fileBytes)

	tx, err := s.Storage.Save(s.Auth,
		common.HexToAddress(payload.Doctor),
		common.HexToAddress(payload.Pacient),
		common.HexToAddress(payload.Institution),
		fileByteString,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(fiber.Map{
		"txHash": tx.Hash().Hex(),
	})
}
