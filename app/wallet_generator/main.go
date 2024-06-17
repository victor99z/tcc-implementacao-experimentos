package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/victor99z/blockchain-app/utils"
)

func main() {
	godotenv.Load()
	arg := os.Args[1:]

	if len(arg) == 0 {
		log.Fatal("No argument provided")
	}

	num_wallet, err := strconv.Atoi(arg[0])
	if err != nil {
		log.Fatal("Invalid argument")
	}

	for i := 0; i < num_wallet; i++ {
		ks := utils.CreateKeyStore(os.Getenv("KEYSTORE_PASSWORD"))
		fmt.Println(ks.Address.Hex())
	}

}
