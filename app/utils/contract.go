package utils

import (
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type deployedContract struct {
	Address string `json:"address"`
	TxHash  string `json:"txHash"`
}

func ReadDeployedContract() deployedContract {

	godotenv.Load()

	for os.Getenv("CONTRACT_ADDR") == "" {
		log.Println("Waiting for contract address")
		time.Sleep(5 * time.Second)
		godotenv.Load()
	}

	tmp := deployedContract{
		Address: os.Getenv("CONTRACT_ADDR"),
	}
	// b, err := os.ReadFile("./deployed-contract.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = json.Unmarshal(b, &tmp)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return tmp
}

func GetChainID() *big.Int {
	num, _ := strconv.Atoi(os.Getenv("CHAIN_ID"))
	res := big.NewInt(int64(num))
	return res
}
