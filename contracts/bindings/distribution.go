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

// DistributionMetaData contains all meta data concerning the Distribution contract.
var DistributionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"Fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b506101eb8061001d5f395ff3fe608060405260043610610028575f3560e01c806351cff8d91461002c578063bef7a2f014610041575b5f80fd5b61003f61003a366004610163565b61006e565b005b34801561004c575f80fd5b5061005c67016345785d8a000081565b60405190815260200160405180910390f35b6100766100ae565b6040516001600160a01b0382169033907f34d58c18c6c1df2c698ccac556acea92941ca7b99d2fccf9e3ac1852d0dec36f905f90a350565b67016345785d8a000034101561010a5760405162461bcd60e51b815260206004820152601e60248201527f446973747269627574696f6e3a20696e73756666696369656e74206665650000604482015260640160405180910390fd5b61011e61dead67016345785d8a0000610146565b5f61013167016345785d8a000034610190565b90508015610143576101433382610146565b50565b5f385f3884865af161015f5763b12d13eb5f526004601cfd5b5050565b5f60208284031215610173575f80fd5b81356001600160a01b0381168114610189575f80fd5b9392505050565b818103818111156101af57634e487b7160e01b5f52601160045260245ffd5b9291505056fea2646970667358221220958c13fa5c3c6da85228eb17f0dca1968144741e2edcbd9fb5ec98dea14d82d464736f6c63430008180033",
}

// DistributionABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributionMetaData.ABI instead.
var DistributionABI = DistributionMetaData.ABI

// DistributionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DistributionMetaData.Bin instead.
var DistributionBin = DistributionMetaData.Bin

// DeployDistribution deploys a new Ethereum contract, binding an instance of Distribution to it.
func DeployDistribution(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Distribution, error) {
	parsed, err := DistributionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DistributionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Distribution{DistributionCaller: DistributionCaller{contract: contract}, DistributionTransactor: DistributionTransactor{contract: contract}, DistributionFilterer: DistributionFilterer{contract: contract}}, nil
}

// Distribution is an auto generated Go binding around an Ethereum contract.
type Distribution struct {
	DistributionCaller     // Read-only binding to the contract
	DistributionTransactor // Write-only binding to the contract
	DistributionFilterer   // Log filterer for contract events
}

// DistributionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributionSession struct {
	Contract     *Distribution     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributionCallerSession struct {
	Contract *DistributionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DistributionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributionTransactorSession struct {
	Contract     *DistributionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DistributionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributionRaw struct {
	Contract *Distribution // Generic contract binding to access the raw methods on
}

// DistributionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributionCallerRaw struct {
	Contract *DistributionCaller // Generic read-only contract binding to access the raw methods on
}

// DistributionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributionTransactorRaw struct {
	Contract *DistributionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistribution creates a new instance of Distribution, bound to a specific deployed contract.
func NewDistribution(address common.Address, backend bind.ContractBackend) (*Distribution, error) {
	contract, err := bindDistribution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Distribution{DistributionCaller: DistributionCaller{contract: contract}, DistributionTransactor: DistributionTransactor{contract: contract}, DistributionFilterer: DistributionFilterer{contract: contract}}, nil
}

// NewDistributionCaller creates a new read-only instance of Distribution, bound to a specific deployed contract.
func NewDistributionCaller(address common.Address, caller bind.ContractCaller) (*DistributionCaller, error) {
	contract, err := bindDistribution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionCaller{contract: contract}, nil
}

// NewDistributionTransactor creates a new write-only instance of Distribution, bound to a specific deployed contract.
func NewDistributionTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributionTransactor, error) {
	contract, err := bindDistribution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionTransactor{contract: contract}, nil
}

// NewDistributionFilterer creates a new log filterer instance of Distribution, bound to a specific deployed contract.
func NewDistributionFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributionFilterer, error) {
	contract, err := bindDistribution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributionFilterer{contract: contract}, nil
}

// bindDistribution binds a generic wrapper to an already deployed contract.
func bindDistribution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DistributionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distribution *DistributionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distribution.Contract.DistributionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distribution *DistributionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribution.Contract.DistributionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distribution *DistributionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distribution.Contract.DistributionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distribution *DistributionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distribution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distribution *DistributionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distribution *DistributionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distribution.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xbef7a2f0.
//
// Solidity: function Fee() view returns(uint256)
func (_Distribution *DistributionCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "Fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xbef7a2f0.
//
// Solidity: function Fee() view returns(uint256)
func (_Distribution *DistributionSession) Fee() (*big.Int, error) {
	return _Distribution.Contract.Fee(&_Distribution.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xbef7a2f0.
//
// Solidity: function Fee() view returns(uint256)
func (_Distribution *DistributionCallerSession) Fee() (*big.Int, error) {
	return _Distribution.Contract.Fee(&_Distribution.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address validator) payable returns()
func (_Distribution *DistributionTransactor) Withdraw(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Distribution.contract.Transact(opts, "withdraw", validator)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address validator) payable returns()
func (_Distribution *DistributionSession) Withdraw(validator common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.Withdraw(&_Distribution.TransactOpts, validator)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address validator) payable returns()
func (_Distribution *DistributionTransactorSession) Withdraw(validator common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.Withdraw(&_Distribution.TransactOpts, validator)
}

// DistributionWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Distribution contract.
type DistributionWithdrawIterator struct {
	Event *DistributionWithdraw // Event containing the contract specifics and raw log

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
func (it *DistributionWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributionWithdraw)
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
		it.Event = new(DistributionWithdraw)
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
func (it *DistributionWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributionWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributionWithdraw represents a Withdraw event raised by the Distribution contract.
type DistributionWithdraw struct {
	Delegator common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x34d58c18c6c1df2c698ccac556acea92941ca7b99d2fccf9e3ac1852d0dec36f.
//
// Solidity: event Withdraw(address indexed delegator, address indexed validator)
func (_Distribution *DistributionFilterer) FilterWithdraw(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DistributionWithdrawIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Distribution.contract.FilterLogs(opts, "Withdraw", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DistributionWithdrawIterator{contract: _Distribution.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x34d58c18c6c1df2c698ccac556acea92941ca7b99d2fccf9e3ac1852d0dec36f.
//
// Solidity: event Withdraw(address indexed delegator, address indexed validator)
func (_Distribution *DistributionFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DistributionWithdraw, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Distribution.contract.WatchLogs(opts, "Withdraw", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributionWithdraw)
				if err := _Distribution.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x34d58c18c6c1df2c698ccac556acea92941ca7b99d2fccf9e3ac1852d0dec36f.
//
// Solidity: event Withdraw(address indexed delegator, address indexed validator)
func (_Distribution *DistributionFilterer) ParseWithdraw(log types.Log) (*DistributionWithdraw, error) {
	event := new(DistributionWithdraw)
	if err := _Distribution.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
