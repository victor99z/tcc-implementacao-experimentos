package main

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/victor99z/blockchain-app/handlers"
	"github.com/victor99z/blockchain-app/models"
	"github.com/victor99z/blockchain-app/utils"
)

type Connection struct {
	Conn    *ethclient.Client
	Auth    *bind.TransactOpts
	ChainId *big.Int
}

func getRecords(addr_pacient string) []models.StorageOnChainRecord {
	godotenv.Load()

	b, _ := os.ReadFile("./wallets/UTC--2024-05-02T02-48-49.339988302Z--823630d906c13965478da03ae24004ad5ff7c691")
	key, _ := keystore.DecryptKey([]byte(b), os.Getenv("KEYSTORE_PASSWORD"))

	conn, err := ethclient.Dial(os.Getenv("ETH_HOST"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	alreadyDeployedContract := utils.ReadDeployedContract()

	storage, _ := models.NewStorageOnChain(
		common.HexToAddress(alreadyDeployedContract.Address), // Smart contract address
		conn,
	)

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, utils.GetChainID())
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = 300000000
	auth.GasPrice = big.NewInt(0)
	auth.Value = big.NewInt(0)

	// file, _ := ioutil.ReadFile("./file_test")
	// _ = file

	// ks := keystore.NewKeyStore("./wallets/", keystore.StandardScryptN, keystore.StandardScryptP)

	// pacient, _ := ks.NewAccount(os.Getenv("KEYSTORE_PASSWORD"))
	// doctor, _ := ks.NewAccount(os.Getenv("KEYSTORE_PASSWORD"))
	// institution, _ := ks.NewAccount(os.Getenv("KEYSTORE_PASSWORD"))

	// fmt.Println("Pacient: ", pacient.Address.Hex())
	// fmt.Println("Doctor: ", doctor.Address.Hex())
	// fmt.Println("Institution: ", institution.Address.Hex())

	// tx, err := storage.Save(auth, auth.From, auth.From, auth.From, "YEAHHHHHHHHHHHHHHHHHH BABY")
	// if err != nil {
	// 	log.Fatalf("Failed to send a transaction: %v", err)
	// }
	// fmt.Println("Transaction hash: ", tx.Hash().Hex())

	tx, err := storage.GetAllRecords(nil, common.HexToAddress(addr_pacient)) // Pass data as argument
	if err != nil {
		log.Fatalf("Failed to send a transaction: %v", err)
	}
	return tx
}

func main() {
	app := fiber.New()

	api := app.Group("/api")

	onchain_handler := handlers.NewOnchainHandler()
	// sidechain_handler := handlers.NewSidechainHandler()

	v1 := api.Group("/v1")
	v1.Get("/record/:id", onchain_handler.GetRecords)
	v1.Get("/record/:id/:counter", onchain_handler.GetRecord)
	v1.Post("/record", onchain_handler.SaveRecord)

	// v1.Get("/record/:id", func(c *fiber.Ctx) error {
	// 	pacient_addr := c.Params("id")
	// 	plain_records := getRecords(pacient_addr)

	// 	return c.JSON(fiber.Map{
	// 		"records": plain_records,
	// 		"count":   len(plain_records),
	// 	})
	// })

	// v1.Get("/record/:id/:counter", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")
	// 	counter := c.Params("counter")
	// 	return c.JSON(fiber.Map{
	// 		"id":      id,
	// 		"counter": counter,
	// 	})
	// })

	// app.Post("/record", func(c *fiber.Ctx) error {

	// 	if err != nil {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"message": "Invalid file",
	// 		})
	// 	}

	// 	payload := struct {
	// 		Doctor      string
	// 		Pacient     string
	// 		Institution string
	// 	}{}

	// 	if err := c.BodyParser(&payload); err != nil {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"message": "Invalid JSON",
	// 		})
	// 	}
	// 	return c.JSON(payload)
	// })

	// v2 := api.Group("/v2")
	// v2.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Hello, World 👋!",
	// 	})
	// })

	log.Fatal(app.Listen(":5001"))
}
