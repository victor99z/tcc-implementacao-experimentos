install:
	go mod tidy

run:
	go run ./main.go

sc-deploy: 
	go run ./deploy/main.go

abigen:
	abigen --bin ./contracts/build/StorageOnChain.bin --abi ./contracts/build/StorageOnChain.abi --pkg models --type StorageOnChain --out ./models/storage_chain.go
	abigen --bin ./contracts/build/StorageOnSidechain.bin --abi ./contracts/build/StorageOnSidechain.abi --pkg models --type StorageOnSidechain --out ./models/storage_sidechain.go