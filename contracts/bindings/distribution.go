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
	ABI: "[{\"type\":\"function\",\"name\":\"Fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b506104e98061001d5f395ff3fe608060405260043610610054575f3560e01c806351cff8d914610058578063715018a61461006d5780638da5cb5b14610081578063bef7a2f0146100cc578063c4d66de8146100f5578063f2fde38b14610114575b5f80fd5b61006b610066366004610486565b610133565b005b348015610078575f80fd5b5061006b610173565b34801561008c575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b0390911681526020015b60405180910390f35b3480156100d7575f80fd5b506100e767016345785d8a000081565b6040519081526020016100c3565b348015610100575f80fd5b5061006b61010f366004610486565b610186565b34801561011f575f80fd5b5061006b61012e366004610486565b610294565b61013b6102d6565b6040516001600160a01b0382169033907f34d58c18c6c1df2c698ccac556acea92941ca7b99d2fccf9e3ac1852d0dec36f905f90a350565b61017b610359565b6101845f6103b4565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156101cb5750825b90505f8267ffffffffffffffff1660011480156101e75750303b155b9050811580156101f5575080155b156102135760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561023d57845460ff60401b1916600160401b1785555b61024686610424565b831561028c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b61029c610359565b6001600160a01b0381166102ca57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6102d3816103b4565b50565b67016345785d8a000034101561032e5760405162461bcd60e51b815260206004820152601e60248201527f446973747269627574696f6e3a20696e73756666696369656e7420666565000060448201526064016102c1565b60405161dead903480156108fc02915f818181858888f193505050501580156102d3573d5f803e3d5ffd5b3361038b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146101845760405163118cdaa760e01b81523360048201526024016102c1565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61042c610435565b6102d38161047e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661018457604051631afcd79f60e31b815260040160405180910390fd5b61029c610435565b5f60208284031215610496575f80fd5b81356001600160a01b03811681146104ac575f80fd5b939250505056fea264697066735822122079cff0339412c5440a6a4cecc42a2991970802f3f70ef91ea3aac0207721017964736f6c63430008180033",
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distribution *DistributionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Distribution.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distribution *DistributionSession) Owner() (common.Address, error) {
	return _Distribution.Contract.Owner(&_Distribution.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distribution *DistributionCallerSession) Owner() (common.Address, error) {
	return _Distribution.Contract.Owner(&_Distribution.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Distribution *DistributionTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Distribution.contract.Transact(opts, "initialize", owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Distribution *DistributionSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.Initialize(&_Distribution.TransactOpts, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Distribution *DistributionTransactorSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.Initialize(&_Distribution.TransactOpts, owner_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distribution *DistributionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distribution.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distribution *DistributionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Distribution.Contract.RenounceOwnership(&_Distribution.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distribution *DistributionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Distribution.Contract.RenounceOwnership(&_Distribution.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distribution *DistributionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Distribution.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distribution *DistributionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.TransferOwnership(&_Distribution.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distribution *DistributionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Distribution.Contract.TransferOwnership(&_Distribution.TransactOpts, newOwner)
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

// DistributionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Distribution contract.
type DistributionInitializedIterator struct {
	Event *DistributionInitialized // Event containing the contract specifics and raw log

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
func (it *DistributionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributionInitialized)
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
		it.Event = new(DistributionInitialized)
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
func (it *DistributionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributionInitialized represents a Initialized event raised by the Distribution contract.
type DistributionInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Distribution *DistributionFilterer) FilterInitialized(opts *bind.FilterOpts) (*DistributionInitializedIterator, error) {

	logs, sub, err := _Distribution.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DistributionInitializedIterator{contract: _Distribution.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Distribution *DistributionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DistributionInitialized) (event.Subscription, error) {

	logs, sub, err := _Distribution.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributionInitialized)
				if err := _Distribution.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Distribution *DistributionFilterer) ParseInitialized(log types.Log) (*DistributionInitialized, error) {
	event := new(DistributionInitialized)
	if err := _Distribution.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Distribution contract.
type DistributionOwnershipTransferredIterator struct {
	Event *DistributionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DistributionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributionOwnershipTransferred)
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
		it.Event = new(DistributionOwnershipTransferred)
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
func (it *DistributionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributionOwnershipTransferred represents a OwnershipTransferred event raised by the Distribution contract.
type DistributionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Distribution *DistributionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DistributionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Distribution.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DistributionOwnershipTransferredIterator{contract: _Distribution.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Distribution *DistributionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DistributionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Distribution.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributionOwnershipTransferred)
				if err := _Distribution.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Distribution *DistributionFilterer) ParseOwnershipTransferred(log types.Log) (*DistributionOwnershipTransferred, error) {
	event := new(DistributionOwnershipTransferred)
	if err := _Distribution.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
