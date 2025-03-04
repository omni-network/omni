// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SolverNetMiddlemanMetaData contains all meta data concerning the SolverNetMiddleman contract.
var SolverNetMiddlemanMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeAndTransfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b610082565b63409feecd198054600181161561003b5763f92ee8a95f526004601cfd5b6001600160401b03808260011c1461007d578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b61027a8061008f5f395ff3fe60806040526004361061001e575f3560e01c8063febe2c2c1461002757005b3661002557005b005b61002561003536600461019b565b5f836001600160a01b0316348484604051610051929190610235565b5f6040518083038185875af1925050503d805f811461008b576040519150601f19603f3d011682016040523d82523d5f602084013e610090565b606091505b50509050806100b257604051633204506f60e01b815260040160405180910390fd5b6001600160a01b0386166100ce576100c9856100eb565b6100e3565b6100e16001600160a01b03871686610107565b505b505050505050565b5f385f3847855af16101045763b12d13eb5f526004601cfd5b50565b5f6370a082315f5230602052602060346024601c865afa601f3d1116610134576390b8ec185f526004601cfd5b81601452603451905063a9059cbb60601b5f5260205f604460105f875af18060015f51141661017557803d853b151710610175576390b8ec185f526004601cfd5b505f60345292915050565b80356001600160a01b0381168114610196575f80fd5b919050565b5f805f805f608086880312156101af575f80fd5b6101b886610180565b94506101c660208701610180565b93506101d460408701610180565b9250606086013567ffffffffffffffff808211156101f0575f80fd5b818801915088601f830112610203575f80fd5b813581811115610211575f80fd5b896020828501011115610222575f80fd5b9699959850939650602001949392505050565b818382375f910190815291905056fea2646970667358221220d4e81f5c5649d5dd936e7caf3b1bf37a49dabafb950eb70bee6087fd4bfed55164736f6c63430008180033",
}

// SolverNetMiddlemanABI is the input ABI used to generate the binding from.
// Deprecated: Use SolverNetMiddlemanMetaData.ABI instead.
var SolverNetMiddlemanABI = SolverNetMiddlemanMetaData.ABI

// SolverNetMiddlemanBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolverNetMiddlemanMetaData.Bin instead.
var SolverNetMiddlemanBin = SolverNetMiddlemanMetaData.Bin

// DeploySolverNetMiddleman deploys a new Ethereum contract, binding an instance of SolverNetMiddleman to it.
func DeploySolverNetMiddleman(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolverNetMiddleman, error) {
	parsed, err := SolverNetMiddlemanMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolverNetMiddlemanBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolverNetMiddleman{SolverNetMiddlemanCaller: SolverNetMiddlemanCaller{contract: contract}, SolverNetMiddlemanTransactor: SolverNetMiddlemanTransactor{contract: contract}, SolverNetMiddlemanFilterer: SolverNetMiddlemanFilterer{contract: contract}}, nil
}

// SolverNetMiddleman is an auto generated Go binding around an Ethereum contract.
type SolverNetMiddleman struct {
	SolverNetMiddlemanCaller     // Read-only binding to the contract
	SolverNetMiddlemanTransactor // Write-only binding to the contract
	SolverNetMiddlemanFilterer   // Log filterer for contract events
}

// SolverNetMiddlemanCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolverNetMiddlemanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetMiddlemanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolverNetMiddlemanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetMiddlemanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolverNetMiddlemanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetMiddlemanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolverNetMiddlemanSession struct {
	Contract     *SolverNetMiddleman // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SolverNetMiddlemanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolverNetMiddlemanCallerSession struct {
	Contract *SolverNetMiddlemanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SolverNetMiddlemanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolverNetMiddlemanTransactorSession struct {
	Contract     *SolverNetMiddlemanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SolverNetMiddlemanRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolverNetMiddlemanRaw struct {
	Contract *SolverNetMiddleman // Generic contract binding to access the raw methods on
}

// SolverNetMiddlemanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolverNetMiddlemanCallerRaw struct {
	Contract *SolverNetMiddlemanCaller // Generic read-only contract binding to access the raw methods on
}

// SolverNetMiddlemanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolverNetMiddlemanTransactorRaw struct {
	Contract *SolverNetMiddlemanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolverNetMiddleman creates a new instance of SolverNetMiddleman, bound to a specific deployed contract.
func NewSolverNetMiddleman(address common.Address, backend bind.ContractBackend) (*SolverNetMiddleman, error) {
	contract, err := bindSolverNetMiddleman(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddleman{SolverNetMiddlemanCaller: SolverNetMiddlemanCaller{contract: contract}, SolverNetMiddlemanTransactor: SolverNetMiddlemanTransactor{contract: contract}, SolverNetMiddlemanFilterer: SolverNetMiddlemanFilterer{contract: contract}}, nil
}

// NewSolverNetMiddlemanCaller creates a new read-only instance of SolverNetMiddleman, bound to a specific deployed contract.
func NewSolverNetMiddlemanCaller(address common.Address, caller bind.ContractCaller) (*SolverNetMiddlemanCaller, error) {
	contract, err := bindSolverNetMiddleman(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddlemanCaller{contract: contract}, nil
}

// NewSolverNetMiddlemanTransactor creates a new write-only instance of SolverNetMiddleman, bound to a specific deployed contract.
func NewSolverNetMiddlemanTransactor(address common.Address, transactor bind.ContractTransactor) (*SolverNetMiddlemanTransactor, error) {
	contract, err := bindSolverNetMiddleman(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddlemanTransactor{contract: contract}, nil
}

// NewSolverNetMiddlemanFilterer creates a new log filterer instance of SolverNetMiddleman, bound to a specific deployed contract.
func NewSolverNetMiddlemanFilterer(address common.Address, filterer bind.ContractFilterer) (*SolverNetMiddlemanFilterer, error) {
	contract, err := bindSolverNetMiddleman(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddlemanFilterer{contract: contract}, nil
}

// bindSolverNetMiddleman binds a generic wrapper to an already deployed contract.
func bindSolverNetMiddleman(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolverNetMiddlemanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetMiddleman *SolverNetMiddlemanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetMiddleman.Contract.SolverNetMiddlemanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetMiddleman *SolverNetMiddlemanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.SolverNetMiddlemanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetMiddleman *SolverNetMiddlemanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.SolverNetMiddlemanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetMiddleman *SolverNetMiddlemanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetMiddleman.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.contract.Transact(opts, method, params...)
}

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactor) ExecuteAndTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.contract.Transact(opts, "executeAndTransfer", token, to, target, data)
}

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanSession) ExecuteAndTransfer(token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.ExecuteAndTransfer(&_SolverNetMiddleman.TransactOpts, token, to, target, data)
}

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorSession) ExecuteAndTransfer(token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.ExecuteAndTransfer(&_SolverNetMiddleman.TransactOpts, token, to, target, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.Fallback(&_SolverNetMiddleman.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.Fallback(&_SolverNetMiddleman.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetMiddleman.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanSession) Receive() (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.Receive(&_SolverNetMiddleman.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorSession) Receive() (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.Receive(&_SolverNetMiddleman.TransactOpts)
}

// SolverNetMiddlemanInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SolverNetMiddleman contract.
type SolverNetMiddlemanInitializedIterator struct {
	Event *SolverNetMiddlemanInitialized // Event containing the contract specifics and raw log

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
func (it *SolverNetMiddlemanInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetMiddlemanInitialized)
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
		it.Event = new(SolverNetMiddlemanInitialized)
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
func (it *SolverNetMiddlemanInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetMiddlemanInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetMiddlemanInitialized represents a Initialized event raised by the SolverNetMiddleman contract.
type SolverNetMiddlemanInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolverNetMiddleman *SolverNetMiddlemanFilterer) FilterInitialized(opts *bind.FilterOpts) (*SolverNetMiddlemanInitializedIterator, error) {

	logs, sub, err := _SolverNetMiddleman.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddlemanInitializedIterator{contract: _SolverNetMiddleman.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolverNetMiddleman *SolverNetMiddlemanFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SolverNetMiddlemanInitialized) (event.Subscription, error) {

	logs, sub, err := _SolverNetMiddleman.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetMiddlemanInitialized)
				if err := _SolverNetMiddleman.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolverNetMiddleman *SolverNetMiddlemanFilterer) ParseInitialized(log types.Log) (*SolverNetMiddlemanInitialized, error) {
	event := new(SolverNetMiddlemanInitialized)
	if err := _SolverNetMiddleman.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
