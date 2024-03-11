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

// OmniStakeMetaData contains all meta data concerning the OmniStake contract.
var OmniStakeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600f57600080fd5b5060968061001e6000396000f3fe608060405260043610601c5760003560e01c8063d0e30db0146021575b600080fd5b60276029565b005b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a256fea2646970667358221220a691d6a4925781aa228829d6819ea739aaeccd931fb7f409c754af6bfe9bfb5d64736f6c63430008170033",
}

// OmniStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniStakeMetaData.ABI instead.
var OmniStakeABI = OmniStakeMetaData.ABI

// OmniStakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniStakeMetaData.Bin instead.
var OmniStakeBin = OmniStakeMetaData.Bin

// DeployOmniStake deploys a new Ethereum contract, binding an instance of OmniStake to it.
func DeployOmniStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniStake, error) {
	parsed, err := OmniStakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniStakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniStake{OmniStakeCaller: OmniStakeCaller{contract: contract}, OmniStakeTransactor: OmniStakeTransactor{contract: contract}, OmniStakeFilterer: OmniStakeFilterer{contract: contract}}, nil
}

// OmniStake is an auto generated Go binding around an Ethereum contract.
type OmniStake struct {
	OmniStakeCaller     // Read-only binding to the contract
	OmniStakeTransactor // Write-only binding to the contract
	OmniStakeFilterer   // Log filterer for contract events
}

// OmniStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniStakeSession struct {
	Contract     *OmniStake        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniStakeCallerSession struct {
	Contract *OmniStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OmniStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniStakeTransactorSession struct {
	Contract     *OmniStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OmniStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniStakeRaw struct {
	Contract *OmniStake // Generic contract binding to access the raw methods on
}

// OmniStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniStakeCallerRaw struct {
	Contract *OmniStakeCaller // Generic read-only contract binding to access the raw methods on
}

// OmniStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniStakeTransactorRaw struct {
	Contract *OmniStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniStake creates a new instance of OmniStake, bound to a specific deployed contract.
func NewOmniStake(address common.Address, backend bind.ContractBackend) (*OmniStake, error) {
	contract, err := bindOmniStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniStake{OmniStakeCaller: OmniStakeCaller{contract: contract}, OmniStakeTransactor: OmniStakeTransactor{contract: contract}, OmniStakeFilterer: OmniStakeFilterer{contract: contract}}, nil
}

// NewOmniStakeCaller creates a new read-only instance of OmniStake, bound to a specific deployed contract.
func NewOmniStakeCaller(address common.Address, caller bind.ContractCaller) (*OmniStakeCaller, error) {
	contract, err := bindOmniStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniStakeCaller{contract: contract}, nil
}

// NewOmniStakeTransactor creates a new write-only instance of OmniStake, bound to a specific deployed contract.
func NewOmniStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniStakeTransactor, error) {
	contract, err := bindOmniStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniStakeTransactor{contract: contract}, nil
}

// NewOmniStakeFilterer creates a new log filterer instance of OmniStake, bound to a specific deployed contract.
func NewOmniStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniStakeFilterer, error) {
	contract, err := bindOmniStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniStakeFilterer{contract: contract}, nil
}

// bindOmniStake binds a generic wrapper to an already deployed contract.
func bindOmniStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniStakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniStake *OmniStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniStake.Contract.OmniStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniStake *OmniStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniStake.Contract.OmniStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniStake *OmniStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniStake.Contract.OmniStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniStake *OmniStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniStake *OmniStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniStake *OmniStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniStake.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_OmniStake *OmniStakeTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniStake.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_OmniStake *OmniStakeSession) Deposit() (*types.Transaction, error) {
	return _OmniStake.Contract.Deposit(&_OmniStake.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_OmniStake *OmniStakeTransactorSession) Deposit() (*types.Transaction, error) {
	return _OmniStake.Contract.Deposit(&_OmniStake.TransactOpts)
}

// OmniStakeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the OmniStake contract.
type OmniStakeDepositIterator struct {
	Event *OmniStakeDeposit // Event containing the contract specifics and raw log

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
func (it *OmniStakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniStakeDeposit)
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
		it.Event = new(OmniStakeDeposit)
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
func (it *OmniStakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniStakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniStakeDeposit represents a Deposit event raised by the OmniStake contract.
type OmniStakeDeposit struct {
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed depositor, uint256 amount)
func (_OmniStake *OmniStakeFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address) (*OmniStakeDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _OmniStake.contract.FilterLogs(opts, "Deposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return &OmniStakeDepositIterator{contract: _OmniStake.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed depositor, uint256 amount)
func (_OmniStake *OmniStakeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *OmniStakeDeposit, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _OmniStake.contract.WatchLogs(opts, "Deposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniStakeDeposit)
				if err := _OmniStake.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed depositor, uint256 amount)
func (_OmniStake *OmniStakeFilterer) ParseDeposit(log types.Log) (*OmniStakeDeposit, error) {
	event := new(OmniStakeDeposit)
	if err := _OmniStake.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
