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
)

// PingPongMetaData contains all meta data concerning the PingPong contract.
var PingPongMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_omni\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pingpong\",\"inputs\":[{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"n\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"start\",\"inputs\":[{\"name\":\"destChainID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Done\",\"inputs\":[{\"name\":\"destChainID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"times\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506040516108cb3803806108cb83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610838806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806339acf9f1146100465780634d4fc29314610075578063645953db1461008a575b600080fd5b600054610059906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b610088610083366004610465565b61009d565b005b6100886100983660046104b0565b610123565b60008167ffffffffffffffff16116100fc5760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a2074696d6573206d757374206265203e2030000000000060448201526064015b60405180910390fd5b61011e838383600161010f8260026104ff565b610119919061052b565b610304565b505050565b60008054906101000a90046001600160a01b03166001600160a01b03166355e2448e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610174573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101989190610553565b80156101ae57506000546001600160a01b031633145b6101fa5760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a206e6f7420616e206f6d6e69207863616c6c000000000060448201526064016100f3565b60008060009054906101000a90046001600160a01b03166001600160a01b0316632f32700e6040518163ffffffff1660e01b8152600401600060405180830381865afa15801561024e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102769190810190610682565b90508167ffffffffffffffff166000036102ec5780516060808301516040805167ffffffffffffffff94851681526001600160a01b0390921660208301529286168184015291517fa73c132355e7f2c5ced079328c310d267f1e830c197c5579b6177218a777313e9281900390910190a1505050565b8051606082015161011e91908561011960018761052b565b6040805167ffffffffffffffff8481166024830152831660448083019190915282518083039091018152606490910182526020810180516001600160e01b031663645953db60e01b179052600080549251634115ab7960e01b8152919290916001600160a01b0390911690634115ab79906103859089908690600401610781565b602060405180830381865afa1580156103a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c691906107ac565b6000546040516350e646dd60e01b81529192506001600160a01b0316906350e646dd9083906103fd908a908a9088906004016107c5565b6000604051808303818588803b15801561041657600080fd5b505af115801561042a573d6000803e3d6000fd5b5050505050505050505050565b67ffffffffffffffff8116811461044d57600080fd5b50565b6001600160a01b038116811461044d57600080fd5b60008060006060848603121561047a57600080fd5b833561048581610437565b9250602084013561049581610450565b915060408401356104a581610437565b809150509250925092565b600080604083850312156104c357600080fd5b82356104ce81610437565b915060208301356104de81610437565b809150509250929050565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff818116838216028082169190828114610523576105236104e9565b505092915050565b67ffffffffffffffff82811682821603908082111561054c5761054c6104e9565b5092915050565b60006020828403121561056557600080fd5b8151801515811461057557600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b60405160e0810167ffffffffffffffff811182821017156105b5576105b561057c565b60405290565b80516105c681610437565b919050565b80516105c681610450565b60005b838110156105f15781810151838201526020016105d9565b50506000910152565b600082601f83011261060b57600080fd5b815167ffffffffffffffff808211156106265761062661057c565b604051601f8301601f19908116603f0116810190828211818310171561064e5761064e61057c565b8160405283815286602085880101111561066757600080fd5b6106788460208301602089016105d6565b9695505050505050565b60006020828403121561069457600080fd5b815167ffffffffffffffff808211156106ac57600080fd5b9083019060e082860312156106c057600080fd5b6106c8610592565b6106d1836105bb565b81526106df602084016105bb565b60208201526106f0604084016105bb565b6040820152610701606084016105cb565b6060820152610712608084016105cb565b608082015260a08301518281111561072957600080fd5b610735878286016105fa565b60a08301525061074760c084016105bb565b60c082015295945050505050565b6000815180845261076d8160208601602086016105d6565b601f01601f19169290920160200192915050565b67ffffffffffffffff831681526040602082015260006107a46040830184610755565b949350505050565b6000602082840312156107be57600080fd5b5051919050565b67ffffffffffffffff841681526001600160a01b03831660208201526060604082018190526000906107f990830184610755565b9594505050505056fea26469706673582212209c06b5c6fdd8051ac0a7e69a1247d82605c7c14955ac0fd19cd494b1ac4b94e864736f6c63430008170033",
}

// PingPongABI is the input ABI used to generate the binding from.
// Deprecated: Use PingPongMetaData.ABI instead.
var PingPongABI = PingPongMetaData.ABI

// PingPongBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PingPongMetaData.Bin instead.
var PingPongBin = PingPongMetaData.Bin

// DeployPingPong deploys a new Ethereum contract, binding an instance of PingPong to it.
func DeployPingPong(auth *bind.TransactOpts, backend bind.ContractBackend, _omni common.Address) (common.Address, *types.Transaction, *PingPong, error) {
	parsed, err := PingPongMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PingPongBin), backend, _omni)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PingPong{PingPongCaller: PingPongCaller{contract: contract}, PingPongTransactor: PingPongTransactor{contract: contract}, PingPongFilterer: PingPongFilterer{contract: contract}}, nil
}

// PingPong is an auto generated Go binding around an Ethereum contract.
type PingPong struct {
	PingPongCaller     // Read-only binding to the contract
	PingPongTransactor // Write-only binding to the contract
	PingPongFilterer   // Log filterer for contract events
}

// PingPongCaller is an auto generated read-only Go binding around an Ethereum contract.
type PingPongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PingPongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PingPongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PingPongSession struct {
	Contract     *PingPong         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingPongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PingPongCallerSession struct {
	Contract *PingPongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PingPongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PingPongTransactorSession struct {
	Contract     *PingPongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PingPongRaw is an auto generated low-level Go binding around an Ethereum contract.
type PingPongRaw struct {
	Contract *PingPong // Generic contract binding to access the raw methods on
}

// PingPongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PingPongCallerRaw struct {
	Contract *PingPongCaller // Generic read-only contract binding to access the raw methods on
}

// PingPongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PingPongTransactorRaw struct {
	Contract *PingPongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPingPong creates a new instance of PingPong, bound to a specific deployed contract.
func NewPingPong(address common.Address, backend bind.ContractBackend) (*PingPong, error) {
	contract, err := bindPingPong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PingPong{PingPongCaller: PingPongCaller{contract: contract}, PingPongTransactor: PingPongTransactor{contract: contract}, PingPongFilterer: PingPongFilterer{contract: contract}}, nil
}

// NewPingPongCaller creates a new read-only instance of PingPong, bound to a specific deployed contract.
func NewPingPongCaller(address common.Address, caller bind.ContractCaller) (*PingPongCaller, error) {
	contract, err := bindPingPong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongCaller{contract: contract}, nil
}

// NewPingPongTransactor creates a new write-only instance of PingPong, bound to a specific deployed contract.
func NewPingPongTransactor(address common.Address, transactor bind.ContractTransactor) (*PingPongTransactor, error) {
	contract, err := bindPingPong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongTransactor{contract: contract}, nil
}

// NewPingPongFilterer creates a new log filterer instance of PingPong, bound to a specific deployed contract.
func NewPingPongFilterer(address common.Address, filterer bind.ContractFilterer) (*PingPongFilterer, error) {
	contract, err := bindPingPong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingPongFilterer{contract: contract}, nil
}

// bindPingPong binds a generic wrapper to an already deployed contract.
func bindPingPong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PingPongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PingPong *PingPongRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPong.Contract.PingPongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PingPong *PingPongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.Contract.PingPongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PingPong *PingPongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPong.Contract.PingPongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PingPong *PingPongCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PingPong *PingPongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PingPong *PingPongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPong.Contract.contract.Transact(opts, method, params...)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_PingPong *PingPongCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPong.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_PingPong *PingPongSession) Omni() (common.Address, error) {
	return _PingPong.Contract.Omni(&_PingPong.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_PingPong *PingPongCallerSession) Omni() (common.Address, error) {
	return _PingPong.Contract.Omni(&_PingPong.CallOpts)
}

// Pingpong is a paid mutator transaction binding the contract method 0x645953db.
//
// Solidity: function pingpong(uint64 times, uint64 n) returns()
func (_PingPong *PingPongTransactor) Pingpong(opts *bind.TransactOpts, times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.contract.Transact(opts, "pingpong", times, n)
}

// Pingpong is a paid mutator transaction binding the contract method 0x645953db.
//
// Solidity: function pingpong(uint64 times, uint64 n) returns()
func (_PingPong *PingPongSession) Pingpong(times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Pingpong(&_PingPong.TransactOpts, times, n)
}

// Pingpong is a paid mutator transaction binding the contract method 0x645953db.
//
// Solidity: function pingpong(uint64 times, uint64 n) returns()
func (_PingPong *PingPongTransactorSession) Pingpong(times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Pingpong(&_PingPong.TransactOpts, times, n)
}

// Start is a paid mutator transaction binding the contract method 0x4d4fc293.
//
// Solidity: function start(uint64 destChainID, address to, uint64 times) returns()
func (_PingPong *PingPongTransactor) Start(opts *bind.TransactOpts, destChainID uint64, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.contract.Transact(opts, "start", destChainID, to, times)
}

// Start is a paid mutator transaction binding the contract method 0x4d4fc293.
//
// Solidity: function start(uint64 destChainID, address to, uint64 times) returns()
func (_PingPong *PingPongSession) Start(destChainID uint64, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Start(&_PingPong.TransactOpts, destChainID, to, times)
}

// Start is a paid mutator transaction binding the contract method 0x4d4fc293.
//
// Solidity: function start(uint64 destChainID, address to, uint64 times) returns()
func (_PingPong *PingPongTransactorSession) Start(destChainID uint64, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Start(&_PingPong.TransactOpts, destChainID, to, times)
}

// PingPongDoneIterator is returned from FilterDone and is used to iterate over the raw logs and unpacked data for Done events raised by the PingPong contract.
type PingPongDoneIterator struct {
	Event *PingPongDone // Event containing the contract specifics and raw log

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
func (it *PingPongDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDone)
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
		it.Event = new(PingPongDone)
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
func (it *PingPongDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongDone represents a Done event raised by the PingPong contract.
type PingPongDone struct {
	DestChainID uint64
	To          common.Address
	Times       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDone is a free log retrieval operation binding the contract event 0xa73c132355e7f2c5ced079328c310d267f1e830c197c5579b6177218a777313e.
//
// Solidity: event Done(uint64 destChainID, address to, uint64 times)
func (_PingPong *PingPongFilterer) FilterDone(opts *bind.FilterOpts) (*PingPongDoneIterator, error) {

	logs, sub, err := _PingPong.contract.FilterLogs(opts, "Done")
	if err != nil {
		return nil, err
	}
	return &PingPongDoneIterator{contract: _PingPong.contract, event: "Done", logs: logs, sub: sub}, nil
}

// WatchDone is a free log subscription operation binding the contract event 0xa73c132355e7f2c5ced079328c310d267f1e830c197c5579b6177218a777313e.
//
// Solidity: event Done(uint64 destChainID, address to, uint64 times)
func (_PingPong *PingPongFilterer) WatchDone(opts *bind.WatchOpts, sink chan<- *PingPongDone) (event.Subscription, error) {

	logs, sub, err := _PingPong.contract.WatchLogs(opts, "Done")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongDone)
				if err := _PingPong.contract.UnpackLog(event, "Done", log); err != nil {
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

// ParseDone is a log parse operation binding the contract event 0xa73c132355e7f2c5ced079328c310d267f1e830c197c5579b6177218a777313e.
//
// Solidity: event Done(uint64 destChainID, address to, uint64 times)
func (_PingPong *PingPongFilterer) ParseDone(log types.Log) (*PingPongDone, error) {
	event := new(PingPongDone)
	if err := _PingPong.contract.UnpackLog(event, "Done", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
