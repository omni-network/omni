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

// UpgradePlan is an auto generated low-level Go binding around an user-defined struct.
type UpgradePlan struct {
	Name   string
	Height uint64
	Info   string
}

// UpgradeMetaData contains all meta data concerning the Upgrade contract.
var UpgradeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cancelUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"planUpgrade\",\"inputs\":[{\"name\":\"plan\",\"type\":\"tuple\",\"internalType\":\"structUpgrade.Plan\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"info\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CancelUpgrade\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlanUpgrade\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"info\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b506105a9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806322a1cc691461006757806355f291661461007c578063715018a6146100845780638da5cb5b1461008c578063c4d66de8146100ca578063f2fde38b146100dd575b600080fd5b61007a610075366004610424565b6100f0565b005b61007a61015c565b61007a61018f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054604080516001600160a01b039092168252519081900360200190f35b61007a6100d8366004610466565b6101a3565b61007a6100eb366004610466565b6102b3565b6100f86102f6565b7fdc944f678a7c5418bb140dca5a53930eee76148b496659ecfafdf9ccd8537138610123828061048f565b61013360408501602086016104dd565b610140604086018661048f565b604051610151959493929190610530565b60405180910390a150565b6101646102f6565b6040517f812c36a273ff85c1871fc7c629fa4c010821a53f3a2492dcc0ea00a396b6a64f90600090a1565b6101976102f6565b6101a16000610351565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156101e95750825b905060008267ffffffffffffffff1660011480156102065750303b155b905081158015610214575080155b156102325760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561025c57845460ff60401b1916600160401b1785555b610265866103c2565b83156102ab57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6102bb6102f6565b6001600160a01b0381166102ea57604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b6102f381610351565b50565b336103287f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146101a15760405163118cdaa760e01b81523360048201526024016102e1565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6103ca6103d3565b6102f38161041c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166101a157604051631afcd79f60e31b815260040160405180910390fd5b6102bb6103d3565b60006020828403121561043657600080fd5b813567ffffffffffffffff81111561044d57600080fd5b82016060818503121561045f57600080fd5b9392505050565b60006020828403121561047857600080fd5b81356001600160a01b038116811461045f57600080fd5b6000808335601e198436030181126104a657600080fd5b83018035915067ffffffffffffffff8211156104c157600080fd5b6020019150368190038213156104d657600080fd5b9250929050565b6000602082840312156104ef57600080fd5b813567ffffffffffffffff8116811461045f57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b606081526000610544606083018789610507565b67ffffffffffffffff861660208401528281036040840152610567818587610507565b9897505050505050505056fea2646970667358221220ef5695f5e690c1adad0d8790e1b3a586ce94a2a873de65cd5a560dee7c3349ba64736f6c63430008180033",
}

// UpgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeMetaData.ABI instead.
var UpgradeABI = UpgradeMetaData.ABI

// UpgradeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UpgradeMetaData.Bin instead.
var UpgradeBin = UpgradeMetaData.Bin

// DeployUpgrade deploys a new Ethereum contract, binding an instance of Upgrade to it.
func DeployUpgrade(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Upgrade, error) {
	parsed, err := UpgradeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UpgradeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Upgrade{UpgradeCaller: UpgradeCaller{contract: contract}, UpgradeTransactor: UpgradeTransactor{contract: contract}, UpgradeFilterer: UpgradeFilterer{contract: contract}}, nil
}

// Upgrade is an auto generated Go binding around an Ethereum contract.
type Upgrade struct {
	UpgradeCaller     // Read-only binding to the contract
	UpgradeTransactor // Write-only binding to the contract
	UpgradeFilterer   // Log filterer for contract events
}

// UpgradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeSession struct {
	Contract     *Upgrade          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeCallerSession struct {
	Contract *UpgradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UpgradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeTransactorSession struct {
	Contract     *UpgradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UpgradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeRaw struct {
	Contract *Upgrade // Generic contract binding to access the raw methods on
}

// UpgradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeCallerRaw struct {
	Contract *UpgradeCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeTransactorRaw struct {
	Contract *UpgradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgrade creates a new instance of Upgrade, bound to a specific deployed contract.
func NewUpgrade(address common.Address, backend bind.ContractBackend) (*Upgrade, error) {
	contract, err := bindUpgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Upgrade{UpgradeCaller: UpgradeCaller{contract: contract}, UpgradeTransactor: UpgradeTransactor{contract: contract}, UpgradeFilterer: UpgradeFilterer{contract: contract}}, nil
}

// NewUpgradeCaller creates a new read-only instance of Upgrade, bound to a specific deployed contract.
func NewUpgradeCaller(address common.Address, caller bind.ContractCaller) (*UpgradeCaller, error) {
	contract, err := bindUpgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeCaller{contract: contract}, nil
}

// NewUpgradeTransactor creates a new write-only instance of Upgrade, bound to a specific deployed contract.
func NewUpgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeTransactor, error) {
	contract, err := bindUpgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeTransactor{contract: contract}, nil
}

// NewUpgradeFilterer creates a new log filterer instance of Upgrade, bound to a specific deployed contract.
func NewUpgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeFilterer, error) {
	contract, err := bindUpgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeFilterer{contract: contract}, nil
}

// bindUpgrade binds a generic wrapper to an already deployed contract.
func bindUpgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Upgrade *UpgradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Upgrade.Contract.UpgradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Upgrade *UpgradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgrade.Contract.UpgradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Upgrade *UpgradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Upgrade.Contract.UpgradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Upgrade *UpgradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Upgrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Upgrade *UpgradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Upgrade *UpgradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Upgrade.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Upgrade *UpgradeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Upgrade.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Upgrade *UpgradeSession) Owner() (common.Address, error) {
	return _Upgrade.Contract.Owner(&_Upgrade.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Upgrade *UpgradeCallerSession) Owner() (common.Address, error) {
	return _Upgrade.Contract.Owner(&_Upgrade.CallOpts)
}

// CancelUpgrade is a paid mutator transaction binding the contract method 0x55f29166.
//
// Solidity: function cancelUpgrade() returns()
func (_Upgrade *UpgradeTransactor) CancelUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgrade.contract.Transact(opts, "cancelUpgrade")
}

// CancelUpgrade is a paid mutator transaction binding the contract method 0x55f29166.
//
// Solidity: function cancelUpgrade() returns()
func (_Upgrade *UpgradeSession) CancelUpgrade() (*types.Transaction, error) {
	return _Upgrade.Contract.CancelUpgrade(&_Upgrade.TransactOpts)
}

// CancelUpgrade is a paid mutator transaction binding the contract method 0x55f29166.
//
// Solidity: function cancelUpgrade() returns()
func (_Upgrade *UpgradeTransactorSession) CancelUpgrade() (*types.Transaction, error) {
	return _Upgrade.Contract.CancelUpgrade(&_Upgrade.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Upgrade *UpgradeTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Upgrade.contract.Transact(opts, "initialize", owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Upgrade *UpgradeSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _Upgrade.Contract.Initialize(&_Upgrade.TransactOpts, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Upgrade *UpgradeTransactorSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _Upgrade.Contract.Initialize(&_Upgrade.TransactOpts, owner_)
}

// PlanUpgrade is a paid mutator transaction binding the contract method 0x22a1cc69.
//
// Solidity: function planUpgrade((string,uint64,string) plan) returns()
func (_Upgrade *UpgradeTransactor) PlanUpgrade(opts *bind.TransactOpts, plan UpgradePlan) (*types.Transaction, error) {
	return _Upgrade.contract.Transact(opts, "planUpgrade", plan)
}

// PlanUpgrade is a paid mutator transaction binding the contract method 0x22a1cc69.
//
// Solidity: function planUpgrade((string,uint64,string) plan) returns()
func (_Upgrade *UpgradeSession) PlanUpgrade(plan UpgradePlan) (*types.Transaction, error) {
	return _Upgrade.Contract.PlanUpgrade(&_Upgrade.TransactOpts, plan)
}

// PlanUpgrade is a paid mutator transaction binding the contract method 0x22a1cc69.
//
// Solidity: function planUpgrade((string,uint64,string) plan) returns()
func (_Upgrade *UpgradeTransactorSession) PlanUpgrade(plan UpgradePlan) (*types.Transaction, error) {
	return _Upgrade.Contract.PlanUpgrade(&_Upgrade.TransactOpts, plan)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Upgrade *UpgradeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgrade.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Upgrade *UpgradeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Upgrade.Contract.RenounceOwnership(&_Upgrade.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Upgrade *UpgradeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Upgrade.Contract.RenounceOwnership(&_Upgrade.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Upgrade *UpgradeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Upgrade.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Upgrade *UpgradeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Upgrade.Contract.TransferOwnership(&_Upgrade.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Upgrade *UpgradeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Upgrade.Contract.TransferOwnership(&_Upgrade.TransactOpts, newOwner)
}

// UpgradeCancelUpgradeIterator is returned from FilterCancelUpgrade and is used to iterate over the raw logs and unpacked data for CancelUpgrade events raised by the Upgrade contract.
type UpgradeCancelUpgradeIterator struct {
	Event *UpgradeCancelUpgrade // Event containing the contract specifics and raw log

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
func (it *UpgradeCancelUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeCancelUpgrade)
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
		it.Event = new(UpgradeCancelUpgrade)
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
func (it *UpgradeCancelUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeCancelUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeCancelUpgrade represents a CancelUpgrade event raised by the Upgrade contract.
type UpgradeCancelUpgrade struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelUpgrade is a free log retrieval operation binding the contract event 0x812c36a273ff85c1871fc7c629fa4c010821a53f3a2492dcc0ea00a396b6a64f.
//
// Solidity: event CancelUpgrade()
func (_Upgrade *UpgradeFilterer) FilterCancelUpgrade(opts *bind.FilterOpts) (*UpgradeCancelUpgradeIterator, error) {

	logs, sub, err := _Upgrade.contract.FilterLogs(opts, "CancelUpgrade")
	if err != nil {
		return nil, err
	}
	return &UpgradeCancelUpgradeIterator{contract: _Upgrade.contract, event: "CancelUpgrade", logs: logs, sub: sub}, nil
}

// WatchCancelUpgrade is a free log subscription operation binding the contract event 0x812c36a273ff85c1871fc7c629fa4c010821a53f3a2492dcc0ea00a396b6a64f.
//
// Solidity: event CancelUpgrade()
func (_Upgrade *UpgradeFilterer) WatchCancelUpgrade(opts *bind.WatchOpts, sink chan<- *UpgradeCancelUpgrade) (event.Subscription, error) {

	logs, sub, err := _Upgrade.contract.WatchLogs(opts, "CancelUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeCancelUpgrade)
				if err := _Upgrade.contract.UnpackLog(event, "CancelUpgrade", log); err != nil {
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

// ParseCancelUpgrade is a log parse operation binding the contract event 0x812c36a273ff85c1871fc7c629fa4c010821a53f3a2492dcc0ea00a396b6a64f.
//
// Solidity: event CancelUpgrade()
func (_Upgrade *UpgradeFilterer) ParseCancelUpgrade(log types.Log) (*UpgradeCancelUpgrade, error) {
	event := new(UpgradeCancelUpgrade)
	if err := _Upgrade.contract.UnpackLog(event, "CancelUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Upgrade contract.
type UpgradeInitializedIterator struct {
	Event *UpgradeInitialized // Event containing the contract specifics and raw log

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
func (it *UpgradeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeInitialized)
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
		it.Event = new(UpgradeInitialized)
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
func (it *UpgradeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeInitialized represents a Initialized event raised by the Upgrade contract.
type UpgradeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Upgrade *UpgradeFilterer) FilterInitialized(opts *bind.FilterOpts) (*UpgradeInitializedIterator, error) {

	logs, sub, err := _Upgrade.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UpgradeInitializedIterator{contract: _Upgrade.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Upgrade *UpgradeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UpgradeInitialized) (event.Subscription, error) {

	logs, sub, err := _Upgrade.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeInitialized)
				if err := _Upgrade.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Upgrade *UpgradeFilterer) ParseInitialized(log types.Log) (*UpgradeInitialized, error) {
	event := new(UpgradeInitialized)
	if err := _Upgrade.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Upgrade contract.
type UpgradeOwnershipTransferredIterator struct {
	Event *UpgradeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UpgradeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeOwnershipTransferred)
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
		it.Event = new(UpgradeOwnershipTransferred)
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
func (it *UpgradeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeOwnershipTransferred represents a OwnershipTransferred event raised by the Upgrade contract.
type UpgradeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Upgrade *UpgradeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UpgradeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Upgrade.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeOwnershipTransferredIterator{contract: _Upgrade.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Upgrade *UpgradeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UpgradeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Upgrade.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeOwnershipTransferred)
				if err := _Upgrade.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Upgrade *UpgradeFilterer) ParseOwnershipTransferred(log types.Log) (*UpgradeOwnershipTransferred, error) {
	event := new(UpgradeOwnershipTransferred)
	if err := _Upgrade.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradePlanUpgradeIterator is returned from FilterPlanUpgrade and is used to iterate over the raw logs and unpacked data for PlanUpgrade events raised by the Upgrade contract.
type UpgradePlanUpgradeIterator struct {
	Event *UpgradePlanUpgrade // Event containing the contract specifics and raw log

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
func (it *UpgradePlanUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradePlanUpgrade)
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
		it.Event = new(UpgradePlanUpgrade)
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
func (it *UpgradePlanUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradePlanUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradePlanUpgrade represents a PlanUpgrade event raised by the Upgrade contract.
type UpgradePlanUpgrade struct {
	Name   string
	Height uint64
	Info   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlanUpgrade is a free log retrieval operation binding the contract event 0xdc944f678a7c5418bb140dca5a53930eee76148b496659ecfafdf9ccd8537138.
//
// Solidity: event PlanUpgrade(string name, uint64 height, string info)
func (_Upgrade *UpgradeFilterer) FilterPlanUpgrade(opts *bind.FilterOpts) (*UpgradePlanUpgradeIterator, error) {

	logs, sub, err := _Upgrade.contract.FilterLogs(opts, "PlanUpgrade")
	if err != nil {
		return nil, err
	}
	return &UpgradePlanUpgradeIterator{contract: _Upgrade.contract, event: "PlanUpgrade", logs: logs, sub: sub}, nil
}

// WatchPlanUpgrade is a free log subscription operation binding the contract event 0xdc944f678a7c5418bb140dca5a53930eee76148b496659ecfafdf9ccd8537138.
//
// Solidity: event PlanUpgrade(string name, uint64 height, string info)
func (_Upgrade *UpgradeFilterer) WatchPlanUpgrade(opts *bind.WatchOpts, sink chan<- *UpgradePlanUpgrade) (event.Subscription, error) {

	logs, sub, err := _Upgrade.contract.WatchLogs(opts, "PlanUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradePlanUpgrade)
				if err := _Upgrade.contract.UnpackLog(event, "PlanUpgrade", log); err != nil {
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

// ParsePlanUpgrade is a log parse operation binding the contract event 0xdc944f678a7c5418bb140dca5a53930eee76148b496659ecfafdf9ccd8537138.
//
// Solidity: event PlanUpgrade(string name, uint64 height, string info)
func (_Upgrade *UpgradeFilterer) ParsePlanUpgrade(log types.Log) (*UpgradePlanUpgrade, error) {
	event := new(UpgradePlanUpgrade)
	if err := _Upgrade.contract.UnpackLog(event, "PlanUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
