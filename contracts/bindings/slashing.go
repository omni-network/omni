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

// SlashingMetaData contains all meta data concerning the Slashing contract.
var SlashingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"Fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unjail\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Unjail\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061015e806100206000396000f3fe6080604052600436106100295760003560e01c8063bef7a2f01461002e578063f679d3051461005c575b600080fd5b34801561003a57600080fd5b5061004a67016345785d8a000081565b60405190815260200160405180910390f35b610064610066565b005b61006e61009b565b60405133907fc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b90600090a2565b67016345785d8a00003410156100f75760405162461bcd60e51b815260206004820152601a60248201527f536c617368696e673a20696e73756666696369656e7420666565000000000000604482015260640160405180910390fd5b60405161dead903480156108fc02916000818181858888f19350505050158015610125573d6000803e3d6000fd5b5056fea264697066735822122098ff290ceb24f37902f2fb4055b624bcfb5bceef5332235a9893a7004a17fc5d64736f6c63430008180033",
}

// SlashingABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashingMetaData.ABI instead.
var SlashingABI = SlashingMetaData.ABI

// SlashingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashingMetaData.Bin instead.
var SlashingBin = SlashingMetaData.Bin

// DeploySlashing deploys a new Ethereum contract, binding an instance of Slashing to it.
func DeploySlashing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Slashing, error) {
	parsed, err := SlashingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Slashing{SlashingCaller: SlashingCaller{contract: contract}, SlashingTransactor: SlashingTransactor{contract: contract}, SlashingFilterer: SlashingFilterer{contract: contract}}, nil
}

// Slashing is an auto generated Go binding around an Ethereum contract.
type Slashing struct {
	SlashingCaller     // Read-only binding to the contract
	SlashingTransactor // Write-only binding to the contract
	SlashingFilterer   // Log filterer for contract events
}

// SlashingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashingSession struct {
	Contract     *Slashing         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashingCallerSession struct {
	Contract *SlashingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SlashingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashingTransactorSession struct {
	Contract     *SlashingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SlashingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashingRaw struct {
	Contract *Slashing // Generic contract binding to access the raw methods on
}

// SlashingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashingCallerRaw struct {
	Contract *SlashingCaller // Generic read-only contract binding to access the raw methods on
}

// SlashingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashingTransactorRaw struct {
	Contract *SlashingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashing creates a new instance of Slashing, bound to a specific deployed contract.
func NewSlashing(address common.Address, backend bind.ContractBackend) (*Slashing, error) {
	contract, err := bindSlashing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Slashing{SlashingCaller: SlashingCaller{contract: contract}, SlashingTransactor: SlashingTransactor{contract: contract}, SlashingFilterer: SlashingFilterer{contract: contract}}, nil
}

// NewSlashingCaller creates a new read-only instance of Slashing, bound to a specific deployed contract.
func NewSlashingCaller(address common.Address, caller bind.ContractCaller) (*SlashingCaller, error) {
	contract, err := bindSlashing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingCaller{contract: contract}, nil
}

// NewSlashingTransactor creates a new write-only instance of Slashing, bound to a specific deployed contract.
func NewSlashingTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashingTransactor, error) {
	contract, err := bindSlashing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingTransactor{contract: contract}, nil
}

// NewSlashingFilterer creates a new log filterer instance of Slashing, bound to a specific deployed contract.
func NewSlashingFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashingFilterer, error) {
	contract, err := bindSlashing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashingFilterer{contract: contract}, nil
}

// bindSlashing binds a generic wrapper to an already deployed contract.
func bindSlashing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slashing *SlashingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slashing.Contract.SlashingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slashing *SlashingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slashing.Contract.SlashingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slashing *SlashingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slashing.Contract.SlashingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slashing *SlashingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slashing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slashing *SlashingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slashing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slashing *SlashingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slashing.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xbef7a2f0.
//
// Solidity: function Fee() view returns(uint256)
func (_Slashing *SlashingCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Slashing.contract.Call(opts, &out, "Fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xbef7a2f0.
//
// Solidity: function Fee() view returns(uint256)
func (_Slashing *SlashingSession) Fee() (*big.Int, error) {
	return _Slashing.Contract.Fee(&_Slashing.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xbef7a2f0.
//
// Solidity: function Fee() view returns(uint256)
func (_Slashing *SlashingCallerSession) Fee() (*big.Int, error) {
	return _Slashing.Contract.Fee(&_Slashing.CallOpts)
}

// Unjail is a paid mutator transaction binding the contract method 0xf679d305.
//
// Solidity: function unjail() payable returns()
func (_Slashing *SlashingTransactor) Unjail(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slashing.contract.Transact(opts, "unjail")
}

// Unjail is a paid mutator transaction binding the contract method 0xf679d305.
//
// Solidity: function unjail() payable returns()
func (_Slashing *SlashingSession) Unjail() (*types.Transaction, error) {
	return _Slashing.Contract.Unjail(&_Slashing.TransactOpts)
}

// Unjail is a paid mutator transaction binding the contract method 0xf679d305.
//
// Solidity: function unjail() payable returns()
func (_Slashing *SlashingTransactorSession) Unjail() (*types.Transaction, error) {
	return _Slashing.Contract.Unjail(&_Slashing.TransactOpts)
}

// SlashingUnjailIterator is returned from FilterUnjail and is used to iterate over the raw logs and unpacked data for Unjail events raised by the Slashing contract.
type SlashingUnjailIterator struct {
	Event *SlashingUnjail // Event containing the contract specifics and raw log

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
func (it *SlashingUnjailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashingUnjail)
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
		it.Event = new(SlashingUnjail)
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
func (it *SlashingUnjailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashingUnjailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashingUnjail represents a Unjail event raised by the Slashing contract.
type SlashingUnjail struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnjail is a free log retrieval operation binding the contract event 0xc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b.
//
// Solidity: event Unjail(address indexed validator)
func (_Slashing *SlashingFilterer) FilterUnjail(opts *bind.FilterOpts, validator []common.Address) (*SlashingUnjailIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Slashing.contract.FilterLogs(opts, "Unjail", validatorRule)
	if err != nil {
		return nil, err
	}
	return &SlashingUnjailIterator{contract: _Slashing.contract, event: "Unjail", logs: logs, sub: sub}, nil
}

// WatchUnjail is a free log subscription operation binding the contract event 0xc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b.
//
// Solidity: event Unjail(address indexed validator)
func (_Slashing *SlashingFilterer) WatchUnjail(opts *bind.WatchOpts, sink chan<- *SlashingUnjail, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Slashing.contract.WatchLogs(opts, "Unjail", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashingUnjail)
				if err := _Slashing.contract.UnpackLog(event, "Unjail", log); err != nil {
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

// ParseUnjail is a log parse operation binding the contract event 0xc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b.
//
// Solidity: event Unjail(address indexed validator)
func (_Slashing *SlashingFilterer) ParseUnjail(log types.Log) (*SlashingUnjail, error) {
	event := new(SlashingUnjail)
	if err := _Slashing.contract.UnpackLog(event, "Unjail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
