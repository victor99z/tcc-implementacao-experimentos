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

// ExampleNFTMetaData contains all meta data concerning the ExampleNFT contract.
var ExampleNFTMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160005534801561001557600080fd5b506108e9806100256000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630178fe3f14610046578063a9059cbb14610076578063d0def52114610092575b600080fd5b610060600480360381019061005b9190610465565b6100c2565b60405161006d919061052b565b60405180910390f35b610090600480360381019061008b91906105ab565b610167565b005b6100ac60048036038101906100a79190610720565b610283565b6040516100b9919061078b565b60405180910390f35b60606002600083815260200190815260200160002080546100e2906107d5565b80601f016020809104026020016040519081016040528092919081815260200182805461010e906107d5565b801561015b5780601f106101305761010080835404028352916020019161015b565b820191906000526020600020905b81548152906001019060200180831161013e57829003601f168201915b50505050509050919050565b6001600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101d257600080fd5b816001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b6000806000549050836001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600260008381526020019081526020016000209080519060200190610304929190610378565b5060008081548092919061031790610836565b9190505550808473ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688560405160405180910390a3600160005461036f919061087f565b91505092915050565b828054610384906107d5565b90600052602060002090601f0160209004810192826103a657600085556103ed565b82601f106103bf57805160ff19168380011785556103ed565b828001600101855582156103ed579182015b828111156103ec5782518255916020019190600101906103d1565b5b5090506103fa91906103fe565b5090565b5b808211156104175760008160009055506001016103ff565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6104428161042f565b811461044d57600080fd5b50565b60008135905061045f81610439565b92915050565b60006020828403121561047b5761047a610425565b5b600061048984828501610450565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104cc5780820151818401526020810190506104b1565b838111156104db576000848401525b50505050565b6000601f19601f8301169050919050565b60006104fd82610492565b610507818561049d565b93506105178185602086016104ae565b610520816104e1565b840191505092915050565b6000602082019050818103600083015261054581846104f2565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105788261054d565b9050919050565b6105888161056d565b811461059357600080fd5b50565b6000813590506105a58161057f565b92915050565b600080604083850312156105c2576105c1610425565b5b60006105d085828601610596565b92505060206105e185828601610450565b9150509250929050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61062d826104e1565b810181811067ffffffffffffffff8211171561064c5761064b6105f5565b5b80604052505050565b600061065f61041b565b905061066b8282610624565b919050565b600067ffffffffffffffff82111561068b5761068a6105f5565b5b610694826104e1565b9050602081019050919050565b82818337600083830152505050565b60006106c36106be84610670565b610655565b9050828152602081018484840111156106df576106de6105f0565b5b6106ea8482856106a1565b509392505050565b600082601f830112610707576107066105eb565b5b81356107178482602086016106b0565b91505092915050565b6000806040838503121561073757610736610425565b5b600061074585828601610596565b925050602083013567ffffffffffffffff8111156107665761076561042a565b5b610772858286016106f2565b9150509250929050565b6107858161042f565b82525050565b60006020820190506107a0600083018461077c565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806107ed57607f821691505b60208210811415610801576108006107a6565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006108418261042f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561087457610873610807565b5b600182019050919050565b600061088a8261042f565b91506108958361042f565b9250828210156108a8576108a7610807565b5b82820390509291505056fea2646970667358221220e3f7c3ad45f2108157a4b85e9fc3cbbb6508223a4b6a7ec3217b667c7045df5664736f6c634300080a0033",
}

// ExampleNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleNFTMetaData.ABI instead.
var ExampleNFTABI = ExampleNFTMetaData.ABI

// ExampleNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleNFTMetaData.Bin instead.
var ExampleNFTBin = ExampleNFTMetaData.Bin

// DeployExampleNFT deploys a new Ethereum contract, binding an instance of ExampleNFT to it.
func DeployExampleNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExampleNFT, error) {
	parsed, err := ExampleNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleNFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleNFT{ExampleNFTCaller: ExampleNFTCaller{contract: contract}, ExampleNFTTransactor: ExampleNFTTransactor{contract: contract}, ExampleNFTFilterer: ExampleNFTFilterer{contract: contract}}, nil
}

// ExampleNFT is an auto generated Go binding around an Ethereum contract.
type ExampleNFT struct {
	ExampleNFTCaller     // Read-only binding to the contract
	ExampleNFTTransactor // Write-only binding to the contract
	ExampleNFTFilterer   // Log filterer for contract events
}

// ExampleNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleNFTSession struct {
	Contract     *ExampleNFT       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleNFTCallerSession struct {
	Contract *ExampleNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExampleNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleNFTTransactorSession struct {
	Contract     *ExampleNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExampleNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleNFTRaw struct {
	Contract *ExampleNFT // Generic contract binding to access the raw methods on
}

// ExampleNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleNFTCallerRaw struct {
	Contract *ExampleNFTCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleNFTTransactorRaw struct {
	Contract *ExampleNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleNFT creates a new instance of ExampleNFT, bound to a specific deployed contract.
func NewExampleNFT(address common.Address, backend bind.ContractBackend) (*ExampleNFT, error) {
	contract, err := bindExampleNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleNFT{ExampleNFTCaller: ExampleNFTCaller{contract: contract}, ExampleNFTTransactor: ExampleNFTTransactor{contract: contract}, ExampleNFTFilterer: ExampleNFTFilterer{contract: contract}}, nil
}

// NewExampleNFTCaller creates a new read-only instance of ExampleNFT, bound to a specific deployed contract.
func NewExampleNFTCaller(address common.Address, caller bind.ContractCaller) (*ExampleNFTCaller, error) {
	contract, err := bindExampleNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleNFTCaller{contract: contract}, nil
}

// NewExampleNFTTransactor creates a new write-only instance of ExampleNFT, bound to a specific deployed contract.
func NewExampleNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleNFTTransactor, error) {
	contract, err := bindExampleNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleNFTTransactor{contract: contract}, nil
}

// NewExampleNFTFilterer creates a new log filterer instance of ExampleNFT, bound to a specific deployed contract.
func NewExampleNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleNFTFilterer, error) {
	contract, err := bindExampleNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleNFTFilterer{contract: contract}, nil
}

// bindExampleNFT binds a generic wrapper to an already deployed contract.
func bindExampleNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleNFT *ExampleNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleNFT.Contract.ExampleNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleNFT *ExampleNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleNFT.Contract.ExampleNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleNFT *ExampleNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleNFT.Contract.ExampleNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleNFT *ExampleNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleNFT *ExampleNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleNFT *ExampleNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleNFT.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 _tokenId) view returns(string)
func (_ExampleNFT *ExampleNFTCaller) GetData(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ExampleNFT.contract.Call(opts, &out, "getData", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 _tokenId) view returns(string)
func (_ExampleNFT *ExampleNFTSession) GetData(_tokenId *big.Int) (string, error) {
	return _ExampleNFT.Contract.GetData(&_ExampleNFT.CallOpts, _tokenId)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 _tokenId) view returns(string)
func (_ExampleNFT *ExampleNFTCallerSession) GetData(_tokenId *big.Int) (string, error) {
	return _ExampleNFT.Contract.GetData(&_ExampleNFT.CallOpts, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _data) returns(uint256)
func (_ExampleNFT *ExampleNFTTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _data string) (*types.Transaction, error) {
	return _ExampleNFT.contract.Transact(opts, "mint", _to, _data)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _data) returns(uint256)
func (_ExampleNFT *ExampleNFTSession) Mint(_to common.Address, _data string) (*types.Transaction, error) {
	return _ExampleNFT.Contract.Mint(&_ExampleNFT.TransactOpts, _to, _data)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _to, string _data) returns(uint256)
func (_ExampleNFT *ExampleNFTTransactorSession) Mint(_to common.Address, _data string) (*types.Transaction, error) {
	return _ExampleNFT.Contract.Mint(&_ExampleNFT.TransactOpts, _to, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ExampleNFT *ExampleNFTTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleNFT.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ExampleNFT *ExampleNFTSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleNFT.Contract.Transfer(&_ExampleNFT.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ExampleNFT *ExampleNFTTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ExampleNFT.Contract.Transfer(&_ExampleNFT.TransactOpts, _to, _tokenId)
}

// ExampleNFTMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ExampleNFT contract.
type ExampleNFTMintIterator struct {
	Event *ExampleNFTMint // Event containing the contract specifics and raw log

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
func (it *ExampleNFTMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleNFTMint)
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
		it.Event = new(ExampleNFTMint)
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
func (it *ExampleNFTMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleNFTMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleNFTMint represents a Mint event raised by the ExampleNFT contract.
type ExampleNFTMint struct {
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_ExampleNFT *ExampleNFTFilterer) FilterMint(opts *bind.FilterOpts, _to []common.Address, _tokenId []*big.Int) (*ExampleNFTMintIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ExampleNFT.contract.FilterLogs(opts, "Mint", _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ExampleNFTMintIterator{contract: _ExampleNFT.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_ExampleNFT *ExampleNFTFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ExampleNFTMint, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ExampleNFT.contract.WatchLogs(opts, "Mint", _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleNFTMint)
				if err := _ExampleNFT.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_ExampleNFT *ExampleNFTFilterer) ParseMint(log types.Log) (*ExampleNFTMint, error) {
	event := new(ExampleNFTMint)
	if err := _ExampleNFT.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ExampleNFT contract.
type ExampleNFTTransferIterator struct {
	Event *ExampleNFTTransfer // Event containing the contract specifics and raw log

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
func (it *ExampleNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleNFTTransfer)
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
		it.Event = new(ExampleNFTTransfer)
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
func (it *ExampleNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleNFTTransfer represents a Transfer event raised by the ExampleNFT contract.
type ExampleNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ExampleNFT *ExampleNFTFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ExampleNFTTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ExampleNFT.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ExampleNFTTransferIterator{contract: _ExampleNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ExampleNFT *ExampleNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExampleNFTTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ExampleNFT.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleNFTTransfer)
				if err := _ExampleNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_ExampleNFT *ExampleNFTFilterer) ParseTransfer(log types.Log) (*ExampleNFTTransfer, error) {
	event := new(ExampleNFTTransfer)
	if err := _ExampleNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
