// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package models

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StorageOnChainRecord is an auto generated low-level Go binding around an user-defined struct.
type StorageOnChainRecord struct {
	Doctor      common.Address
	Recipient   common.Address
	Institution common.Address
	Data        string
	CreateAt    *big.Int
}

// StorageOnChainMetaData contains all meta data concerning the StorageOnChain contract.
var StorageOnChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_doctor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_institution\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"create_at\",\"type\":\"uint256\"}],\"name\":\"Save\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"getAllRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"doctor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"institution\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"create_at\",\"type\":\"uint256\"}],\"internalType\":\"structStorageOnChain.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"doctor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"institution\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"create_at\",\"type\":\"uint256\"}],\"internalType\":\"structStorageOnChain.Record\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"getRecordCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_doctor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_institution\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"save\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506111f1806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806302e2df1c146100515780633f84975b1461008157806378b531051461009d578063cdb8acf0146100cd575b600080fd5b61006b60048036038101906100669190610926565b6100fd565b6040516100789190610b43565b60405180910390f35b61009b60048036038101906100969190610c9a565b610338565b005b6100b760048036038101906100b29190610926565b6105e3565b6040516100c49190610d2c565b60405180910390f35b6100e760048036038101906100e29190610d73565b61062c565b6040516100f49190610e29565b60405180910390f35b60606000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561032d57838290600052602060002090600502016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805461029290610e7a565b80601f01602080910402602001604051908101604052809291908181526020018280546102be90610e7a565b801561030b5780601f106102e05761010080835404028352916020019161030b565b820191906000526020600020905b8154815290600101906020018083116102ee57829003601f168201915b505050505081526020016004820154815250508152602001906001019061015d565b505050509050919050565b600042905060006040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381525090506000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030190816104fc9190611057565b5060808201518160040155505060018060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105589190611158565b925050819055508473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8e9fcc0f71060a700903d14598fd64a54c8965f36c53310ee8f9c601bb0e8605856040516105d39190610d2c565b60405180910390a4505050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610634610843565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082815481106106845761068361118c565b5b90600052602060002090600502016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820180546107af90610e7a565b80601f01602080910402602001604051908101604052809291908181526020018280546107db90610e7a565b80156108285780601f106107fd57610100808354040283529160200191610828565b820191906000526020600020905b81548152906001019060200180831161080b57829003601f168201915b50505050508152602001600482015481525050905092915050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108f3826108c8565b9050919050565b610903816108e8565b811461090e57600080fd5b50565b600081359050610920816108fa565b92915050565b60006020828403121561093c5761093b6108be565b5b600061094a84828501610911565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610988816108e8565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109c85780820151818401526020810190506109ad565b60008484015250505050565b6000601f19601f8301169050919050565b60006109f08261098e565b6109fa8185610999565b9350610a0a8185602086016109aa565b610a13816109d4565b840191505092915050565b6000819050919050565b610a3181610a1e565b82525050565b600060a083016000830151610a4f600086018261097f565b506020830151610a62602086018261097f565b506040830151610a75604086018261097f565b5060608301518482036060860152610a8d82826109e5565b9150506080830151610aa26080860182610a28565b508091505092915050565b6000610ab98383610a37565b905092915050565b6000602082019050919050565b6000610ad982610953565b610ae3818561095e565b935083602082028501610af58561096f565b8060005b85811015610b315784840389528151610b128582610aad565b9450610b1d83610ac1565b925060208a01995050600181019050610af9565b50829750879550505050505092915050565b60006020820190508181036000830152610b5d8184610ace565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610ba7826109d4565b810181811067ffffffffffffffff82111715610bc657610bc5610b6f565b5b80604052505050565b6000610bd96108b4565b9050610be58282610b9e565b919050565b600067ffffffffffffffff821115610c0557610c04610b6f565b5b610c0e826109d4565b9050602081019050919050565b82818337600083830152505050565b6000610c3d610c3884610bea565b610bcf565b905082815260208101848484011115610c5957610c58610b6a565b5b610c64848285610c1b565b509392505050565b600082601f830112610c8157610c80610b65565b5b8135610c91848260208601610c2a565b91505092915050565b60008060008060808587031215610cb457610cb36108be565b5b6000610cc287828801610911565b9450506020610cd387828801610911565b9350506040610ce487828801610911565b925050606085013567ffffffffffffffff811115610d0557610d046108c3565b5b610d1187828801610c6c565b91505092959194509250565b610d2681610a1e565b82525050565b6000602082019050610d416000830184610d1d565b92915050565b610d5081610a1e565b8114610d5b57600080fd5b50565b600081359050610d6d81610d47565b92915050565b60008060408385031215610d8a57610d896108be565b5b6000610d9885828601610911565b9250506020610da985828601610d5e565b9150509250929050565b600060a083016000830151610dcb600086018261097f565b506020830151610dde602086018261097f565b506040830151610df1604086018261097f565b5060608301518482036060860152610e0982826109e5565b9150506080830151610e1e6080860182610a28565b508091505092915050565b60006020820190508181036000830152610e438184610db3565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e9257607f821691505b602082108103610ea557610ea4610e4b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610f0d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610ed0565b610f178683610ed0565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610f54610f4f610f4a84610a1e565b610f2f565b610a1e565b9050919050565b6000819050919050565b610f6e83610f39565b610f82610f7a82610f5b565b848454610edd565b825550505050565b600090565b610f97610f8a565b610fa2818484610f65565b505050565b5b81811015610fc657610fbb600082610f8f565b600181019050610fa8565b5050565b601f82111561100b57610fdc81610eab565b610fe584610ec0565b81016020851015610ff4578190505b61100861100085610ec0565b830182610fa7565b50505b505050565b600082821c905092915050565b600061102e60001984600802611010565b1980831691505092915050565b6000611047838361101d565b9150826002028217905092915050565b6110608261098e565b67ffffffffffffffff81111561107957611078610b6f565b5b6110838254610e7a565b61108e828285610fca565b600060209050601f8311600181146110c157600084156110af578287015190505b6110b9858261103b565b865550611121565b601f1984166110cf86610eab565b60005b828110156110f7578489015182556001820191506020850194506020810190506110d2565b868310156111145784890151611110601f89168261101d565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061116382610a1e565b915061116e83610a1e565b925082820190508082111561118657611185611129565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea26469706673582212202e9e2062e62df9be16c74c8c017a4d15fecec89fe57b4847b781871b358ebc1d64736f6c63430008130033",
}

// StorageOnChainABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageOnChainMetaData.ABI instead.
var StorageOnChainABI = StorageOnChainMetaData.ABI

// StorageOnChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageOnChainMetaData.Bin instead.
var StorageOnChainBin = StorageOnChainMetaData.Bin

// DeployStorageOnChain deploys a new Ethereum contract, binding an instance of StorageOnChain to it.
func DeployStorageOnChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageOnChain, error) {
	parsed, err := StorageOnChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageOnChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageOnChain{StorageOnChainCaller: StorageOnChainCaller{contract: contract}, StorageOnChainTransactor: StorageOnChainTransactor{contract: contract}, StorageOnChainFilterer: StorageOnChainFilterer{contract: contract}}, nil
}

// StorageOnChain is an auto generated Go binding around an Ethereum contract.
type StorageOnChain struct {
	StorageOnChainCaller     // Read-only binding to the contract
	StorageOnChainTransactor // Write-only binding to the contract
	StorageOnChainFilterer   // Log filterer for contract events
}

// StorageOnChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageOnChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageOnChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageOnChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageOnChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageOnChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageOnChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageOnChainSession struct {
	Contract     *StorageOnChain   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageOnChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageOnChainCallerSession struct {
	Contract *StorageOnChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StorageOnChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageOnChainTransactorSession struct {
	Contract     *StorageOnChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StorageOnChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageOnChainRaw struct {
	Contract *StorageOnChain // Generic contract binding to access the raw methods on
}

// StorageOnChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageOnChainCallerRaw struct {
	Contract *StorageOnChainCaller // Generic read-only contract binding to access the raw methods on
}

// StorageOnChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageOnChainTransactorRaw struct {
	Contract *StorageOnChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageOnChain creates a new instance of StorageOnChain, bound to a specific deployed contract.
func NewStorageOnChain(address common.Address, backend bind.ContractBackend) (*StorageOnChain, error) {
	contract, err := bindStorageOnChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageOnChain{StorageOnChainCaller: StorageOnChainCaller{contract: contract}, StorageOnChainTransactor: StorageOnChainTransactor{contract: contract}, StorageOnChainFilterer: StorageOnChainFilterer{contract: contract}}, nil
}

// NewStorageOnChainCaller creates a new read-only instance of StorageOnChain, bound to a specific deployed contract.
func NewStorageOnChainCaller(address common.Address, caller bind.ContractCaller) (*StorageOnChainCaller, error) {
	contract, err := bindStorageOnChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageOnChainCaller{contract: contract}, nil
}

// NewStorageOnChainTransactor creates a new write-only instance of StorageOnChain, bound to a specific deployed contract.
func NewStorageOnChainTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageOnChainTransactor, error) {
	contract, err := bindStorageOnChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageOnChainTransactor{contract: contract}, nil
}

// NewStorageOnChainFilterer creates a new log filterer instance of StorageOnChain, bound to a specific deployed contract.
func NewStorageOnChainFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageOnChainFilterer, error) {
	contract, err := bindStorageOnChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageOnChainFilterer{contract: contract}, nil
}

// bindStorageOnChain binds a generic wrapper to an already deployed contract.
func bindStorageOnChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageOnChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageOnChain *StorageOnChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageOnChain.Contract.StorageOnChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageOnChain *StorageOnChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageOnChain.Contract.StorageOnChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageOnChain *StorageOnChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageOnChain.Contract.StorageOnChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageOnChain *StorageOnChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageOnChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageOnChain *StorageOnChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageOnChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageOnChain *StorageOnChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageOnChain.Contract.contract.Transact(opts, method, params...)
}

// GetAllRecords is a free data retrieval call binding the contract method 0x02e2df1c.
//
// Solidity: function getAllRecords(address _recipient) view returns((address,address,address,string,uint256)[])
func (_StorageOnChain *StorageOnChainCaller) GetAllRecords(opts *bind.CallOpts, _recipient common.Address) ([]StorageOnChainRecord, error) {
	var out []interface{}
	err := _StorageOnChain.contract.Call(opts, &out, "getAllRecords", _recipient)

	if err != nil {
		return *new([]StorageOnChainRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]StorageOnChainRecord)).(*[]StorageOnChainRecord)

	return out0, err

}

// GetAllRecords is a free data retrieval call binding the contract method 0x02e2df1c.
//
// Solidity: function getAllRecords(address _recipient) view returns((address,address,address,string,uint256)[])
func (_StorageOnChain *StorageOnChainSession) GetAllRecords(_recipient common.Address) ([]StorageOnChainRecord, error) {
	return _StorageOnChain.Contract.GetAllRecords(&_StorageOnChain.CallOpts, _recipient)
}

// GetAllRecords is a free data retrieval call binding the contract method 0x02e2df1c.
//
// Solidity: function getAllRecords(address _recipient) view returns((address,address,address,string,uint256)[])
func (_StorageOnChain *StorageOnChainCallerSession) GetAllRecords(_recipient common.Address) ([]StorageOnChainRecord, error) {
	return _StorageOnChain.Contract.GetAllRecords(&_StorageOnChain.CallOpts, _recipient)
}

// GetRecord is a free data retrieval call binding the contract method 0xcdb8acf0.
//
// Solidity: function getRecord(address _recipient, uint256 _index) view returns((address,address,address,string,uint256))
func (_StorageOnChain *StorageOnChainCaller) GetRecord(opts *bind.CallOpts, _recipient common.Address, _index *big.Int) (StorageOnChainRecord, error) {
	var out []interface{}
	err := _StorageOnChain.contract.Call(opts, &out, "getRecord", _recipient, _index)

	if err != nil {
		return *new(StorageOnChainRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageOnChainRecord)).(*StorageOnChainRecord)

	return out0, err

}

// GetRecord is a free data retrieval call binding the contract method 0xcdb8acf0.
//
// Solidity: function getRecord(address _recipient, uint256 _index) view returns((address,address,address,string,uint256))
func (_StorageOnChain *StorageOnChainSession) GetRecord(_recipient common.Address, _index *big.Int) (StorageOnChainRecord, error) {
	return _StorageOnChain.Contract.GetRecord(&_StorageOnChain.CallOpts, _recipient, _index)
}

// GetRecord is a free data retrieval call binding the contract method 0xcdb8acf0.
//
// Solidity: function getRecord(address _recipient, uint256 _index) view returns((address,address,address,string,uint256))
func (_StorageOnChain *StorageOnChainCallerSession) GetRecord(_recipient common.Address, _index *big.Int) (StorageOnChainRecord, error) {
	return _StorageOnChain.Contract.GetRecord(&_StorageOnChain.CallOpts, _recipient, _index)
}

// GetRecordCounter is a free data retrieval call binding the contract method 0x78b53105.
//
// Solidity: function getRecordCounter(address _recipient) view returns(uint256)
func (_StorageOnChain *StorageOnChainCaller) GetRecordCounter(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StorageOnChain.contract.Call(opts, &out, "getRecordCounter", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecordCounter is a free data retrieval call binding the contract method 0x78b53105.
//
// Solidity: function getRecordCounter(address _recipient) view returns(uint256)
func (_StorageOnChain *StorageOnChainSession) GetRecordCounter(_recipient common.Address) (*big.Int, error) {
	return _StorageOnChain.Contract.GetRecordCounter(&_StorageOnChain.CallOpts, _recipient)
}

// GetRecordCounter is a free data retrieval call binding the contract method 0x78b53105.
//
// Solidity: function getRecordCounter(address _recipient) view returns(uint256)
func (_StorageOnChain *StorageOnChainCallerSession) GetRecordCounter(_recipient common.Address) (*big.Int, error) {
	return _StorageOnChain.Contract.GetRecordCounter(&_StorageOnChain.CallOpts, _recipient)
}

// Save is a paid mutator transaction binding the contract method 0x3f84975b.
//
// Solidity: function save(address _doctor, address _recipient, address _institution, string _data) returns()
func (_StorageOnChain *StorageOnChainTransactor) Save(opts *bind.TransactOpts, _doctor common.Address, _recipient common.Address, _institution common.Address, _data string) (*types.Transaction, error) {
	return _StorageOnChain.contract.Transact(opts, "save", _doctor, _recipient, _institution, _data)
}

// Save is a paid mutator transaction binding the contract method 0x3f84975b.
//
// Solidity: function save(address _doctor, address _recipient, address _institution, string _data) returns()
func (_StorageOnChain *StorageOnChainSession) Save(_doctor common.Address, _recipient common.Address, _institution common.Address, _data string) (*types.Transaction, error) {
	return _StorageOnChain.Contract.Save(&_StorageOnChain.TransactOpts, _doctor, _recipient, _institution, _data)
}

// Save is a paid mutator transaction binding the contract method 0x3f84975b.
//
// Solidity: function save(address _doctor, address _recipient, address _institution, string _data) returns()
func (_StorageOnChain *StorageOnChainTransactorSession) Save(_doctor common.Address, _recipient common.Address, _institution common.Address, _data string) (*types.Transaction, error) {
	return _StorageOnChain.Contract.Save(&_StorageOnChain.TransactOpts, _doctor, _recipient, _institution, _data)
}

// StorageOnChainSaveIterator is returned from FilterSave and is used to iterate over the raw logs and unpacked data for Save events raised by the StorageOnChain contract.
type StorageOnChainSaveIterator struct {
	Event *StorageOnChainSave // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageOnChainSaveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOnChainSave)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageOnChainSave)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageOnChainSaveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOnChainSaveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOnChainSave represents a Save event raised by the StorageOnChain contract.
type StorageOnChainSave struct {
	Doctor      common.Address
	Institution common.Address
	Recipient   common.Address
	CreateAt    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSave is a free log retrieval operation binding the contract event 0x8e9fcc0f71060a700903d14598fd64a54c8965f36c53310ee8f9c601bb0e8605.
//
// Solidity: event Save(address indexed _doctor, address indexed _institution, address indexed _recipient, uint256 create_at)
func (_StorageOnChain *StorageOnChainFilterer) FilterSave(opts *bind.FilterOpts, _doctor []common.Address, _institution []common.Address, _recipient []common.Address) (*StorageOnChainSaveIterator, error) {

	var _doctorRule []interface{}
	for _, _doctorItem := range _doctor {
		_doctorRule = append(_doctorRule, _doctorItem)
	}
	var _institutionRule []interface{}
	for _, _institutionItem := range _institution {
		_institutionRule = append(_institutionRule, _institutionItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _StorageOnChain.contract.FilterLogs(opts, "Save", _doctorRule, _institutionRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return &StorageOnChainSaveIterator{contract: _StorageOnChain.contract, event: "Save", logs: logs, sub: sub}, nil
}

// WatchSave is a free log subscription operation binding the contract event 0x8e9fcc0f71060a700903d14598fd64a54c8965f36c53310ee8f9c601bb0e8605.
//
// Solidity: event Save(address indexed _doctor, address indexed _institution, address indexed _recipient, uint256 create_at)
func (_StorageOnChain *StorageOnChainFilterer) WatchSave(opts *bind.WatchOpts, sink chan<- *StorageOnChainSave, _doctor []common.Address, _institution []common.Address, _recipient []common.Address) (event.Subscription, error) {

	var _doctorRule []interface{}
	for _, _doctorItem := range _doctor {
		_doctorRule = append(_doctorRule, _doctorItem)
	}
	var _institutionRule []interface{}
	for _, _institutionItem := range _institution {
		_institutionRule = append(_institutionRule, _institutionItem)
	}
	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}

	logs, sub, err := _StorageOnChain.contract.WatchLogs(opts, "Save", _doctorRule, _institutionRule, _recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOnChainSave)
				if err := _StorageOnChain.contract.UnpackLog(event, "Save", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSave is a log parse operation binding the contract event 0x8e9fcc0f71060a700903d14598fd64a54c8965f36c53310ee8f9c601bb0e8605.
//
// Solidity: event Save(address indexed _doctor, address indexed _institution, address indexed _recipient, uint256 create_at)
func (_StorageOnChain *StorageOnChainFilterer) ParseSave(log types.Log) (*StorageOnChainSave, error) {
	event := new(StorageOnChainSave)
	if err := _StorageOnChain.contract.UnpackLog(event, "Save", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
