package utils

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func CreateKeyStore(password string) accounts.Account {

	ks := keystore.NewKeyStore("./wallets/", keystore.StandardScryptN, keystore.StandardScryptP)

	acc, err := ks.NewAccount(password)

	if err != nil {
		log.Println(err)
	}

	return acc
}

func ListKeyStore() []common.Address {
	ks := keystore.NewKeyStore("./wallets/", keystore.StandardScryptN, keystore.StandardScryptP)
	accounts := ks.Accounts()

	vec := make([]common.Address, 0)
	for _, account := range accounts {
		vec = append(vec, account.Address)
	}
	return vec
}
