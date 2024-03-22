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
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50610342806100206000396000f3fe60806040526004361061001e5760003560e01c806398b1e06a14610023575b600080fd5b610036610031366004610206565b610038565b005b670de0b6b3a76400003410156100955760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e695374616b653a206465706f73697420616d7420746f6f206c6f77000060448201526064015b60405180910390fd5b67ffffffffffffffff34106100ec5760405162461bcd60e51b815260206004820152601f60248201527f4f6d6e695374616b653a206465706f73697420616d7420746f6f206869676800604482015260640161008c565b6100f581610191565b6001600160a01b0316336001600160a01b0316146101555760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e695374616b653a207075626b6579206e6f742073656e64657200000000604482015260640161008c565b7f3884ea6795b4d6e562e2ed70773e296848c9e2d02715686fe7adfc60e4d9abca81346040516101869291906102b7565b60405180910390a150565b600081516040146101e45760405162461bcd60e51b815260206004820181905260248201527f536563703235366b313a20696e76616c6964207075626b6579206c656e677468604482015260640161008c565b50805160209091012090565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561021857600080fd5b813567ffffffffffffffff8082111561023057600080fd5b818401915084601f83011261024457600080fd5b813581811115610256576102566101f0565b604051601f8201601f19908116603f0116810190838211818310171561027e5761027e6101f0565b8160405282815287602084870101111561029757600080fd5b826020860160208301376000928101602001929092525095945050505050565b604081526000835180604084015260005b818110156102e557602081870181015160608684010152016102c8565b506000606082850101526060601f19601f830116840101915050826020830152939250505056fea2646970667358221220e4cb621aa1ab7792d1bdbcff49b2e5a26f0fb7449e0e048407ef3a561e51b12564736f6c63430008180033",
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

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(bytes pubkey) payable returns()
func (_OmniStake *OmniStakeTransactor) Deposit(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _OmniStake.contract.Transact(opts, "deposit", pubkey)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(bytes pubkey) payable returns()
func (_OmniStake *OmniStakeSession) Deposit(pubkey []byte) (*types.Transaction, error) {
	return _OmniStake.Contract.Deposit(&_OmniStake.TransactOpts, pubkey)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(bytes pubkey) payable returns()
func (_OmniStake *OmniStakeTransactorSession) Deposit(pubkey []byte) (*types.Transaction, error) {
	return _OmniStake.Contract.Deposit(&_OmniStake.TransactOpts, pubkey)
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
	Pubkey []byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x3884ea6795b4d6e562e2ed70773e296848c9e2d02715686fe7adfc60e4d9abca.
//
// Solidity: event Deposit(bytes pubkey, uint256 amount)
func (_OmniStake *OmniStakeFilterer) FilterDeposit(opts *bind.FilterOpts) (*OmniStakeDepositIterator, error) {

	logs, sub, err := _OmniStake.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &OmniStakeDepositIterator{contract: _OmniStake.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x3884ea6795b4d6e562e2ed70773e296848c9e2d02715686fe7adfc60e4d9abca.
//
// Solidity: event Deposit(bytes pubkey, uint256 amount)
func (_OmniStake *OmniStakeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *OmniStakeDeposit) (event.Subscription, error) {

	logs, sub, err := _OmniStake.contract.WatchLogs(opts, "Deposit")
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

// ParseDeposit is a log parse operation binding the contract event 0x3884ea6795b4d6e562e2ed70773e296848c9e2d02715686fe7adfc60e4d9abca.
//
// Solidity: event Deposit(bytes pubkey, uint256 amount)
func (_OmniStake *OmniStakeFilterer) ParseDeposit(log types.Log) (*OmniStakeDeposit, error) {
	event := new(OmniStakeDeposit)
	if err := _OmniStake.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
