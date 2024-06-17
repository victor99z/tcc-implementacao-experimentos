package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/victor99z/blockchain-app/models"
	"github.com/victor99z/blockchain-app/utils"
)

func main() {
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

	chainId, _ := conn.NetworkID(context.Background())

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
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
	// fmt.Println("Addr: ", auth.From.Hex())

	// tx, err := storage.GetAllRecords(nil, common.HexToAddress("0xe235c40230580f7EF0B5F237fd3F7A8D989Cc42d")) // Pass data as argument
	// if err != nil {
	// 	log.Fatalf("Failed to send a transaction: %v", err)
	// }

	// fmt.Println(len(tx))

	bigs, _ := storage.GetRecordCounter(nil, common.HexToAddress("0xe235c40230580f7EF0B5F237fd3F7A8D989Cc42d"))
	fmt.Println(bigs)

	fodase, _ := storage.GetAllRecords(nil, common.HexToAddress("0xe235c40230580f7EF0B5F237fd3F7A8D989Cc42d"))
	marshal, _ := json.MarshalIndent(fodase, "", "  ")
	fmt.Println(string(marshal))

}
