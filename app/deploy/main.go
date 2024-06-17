package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/victor99z/blockchain-app/models"
	"github.com/victor99z/blockchain-app/utils"
)

func contractBuilder(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction) {
	godotenv.Load()
	sc_flag := os.Getenv("SMART_CONTRACT_VERSION") // change

	if sc_flag == "onchain" {
		addr, tx, _, err := models.DeployStorageOnChain(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		return addr, tx
	} else if sc_flag == "sidechain" {
		addr, tx, _, err := models.DeployStorageOnSidechain(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		return addr, tx
	}
	log.Fatal("SMART_CONTRACT_VERSION not found")
	return common.Address{}, nil
}

const brute_key = `{"address":"881466d0000b87cc0ebe032adeec41d1cc52102f","crypto":{"cipher":"aes-128-ctr","ciphertext":"15673b83b28621afe90c4f3bda47526a3090f72996aa0d6b45868214c5203f2e","cipherparams":{"iv":"df18cc787a39aaffb84439c83dc7774c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"7f392369e14002780615f22dee33fedb8bdde72ec637883d9c5a77a433de7224"},"mac":"310a9a2db226e8b1b3c411e3af84e2de2f5031f38d4e0508080536902c94e3cf"},"id":"7de64149-e8fc-447e-a226-5d791d5670f0","version":3}`

func main() {
	godotenv.Load()
	key, err := keystore.DecryptKey([]byte(brute_key), "password")

	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial("http://172.16.239.10:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, utils.GetChainID())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	addr, tx := contractBuilder(auth, client)

	r := struct {
		Address string `json:"address"`
		TxHash  string `json:"txHash"`
	}{
		Address: addr.Hex(),
		TxHash:  tx.Hash().Hex(),
	}

	parser, _ := json.Marshal(r)
	ioutil.WriteFile("./deployed-contract.json", parser, os.ModePerm)

	fmt.Println("-----------------------------------")
	fmt.Println(addr.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-----------------------------------")
}
