package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/victor99z/blockchain-app/utils"
)

func main() {
	godotenv.Load()
	conn, err := ethclient.Dial(os.Getenv("ETH_HOST_WS"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	alreadyDeployedContract := utils.ReadDeployedContract()

	logs := make(chan types.Log)

	fmt.Println(alreadyDeployedContract.Address)

	sub, err := conn.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(alreadyDeployedContract.Address)},
	}, logs)

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("Received event:", vLog)
			// Parse event data from vLog
			// Do something with the event data
		}
	}

}
