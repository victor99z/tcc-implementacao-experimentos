# Application (interaction level)


### TODO

- [x] Create new accounts
- [x] Deploy smart contracts
- [x] Interact with smart contracts

Smart contracts:
- [ ] Save files in the blockchain using only blockchain
- [ ] Save files in IPFS and save the CID in the blockchain 


## Instruções para geração de código;


### 1. Gera o abi do contrato
```bash
$ solc --bin SimpleStorage.sol --abi SimpleStorage.sol -o build --overwrite
```
### 2. gera a struct em golang junto com o abi no arquivo

```bash
$ abigen --abi Storage.abi --pkg main --type Storage --out Storage.go
$ abigen --abi ./build/Storage.abi --pkg contract --type Storage --out Storage.go --bin ./build/Storage.bin
```

### Para executar a API

```bash

go run deploy/main.go
go run api/main.go

```

## Links

- https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings