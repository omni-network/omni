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

// RedenomAccountRange is an auto generated low-level Go binding around an user-defined struct.
type RedenomAccountRange struct {
	Addresses []common.Address
	Accounts  [][]byte
	Proof     [][]byte
}

// RedenomMetaData contains all meta data concerning the Redenom contract.
var RedenomMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submit\",\"inputs\":[{\"name\":\"range\",\"type\":\"tuple\",\"internalType\":\"structRedenom.AccountRange\",\"components\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"accounts\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Submitted\",\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"accounts\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"proof\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b506105fa8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c8063715018a6146100595780638da5cb5b14610063578063c4d66de8146100a1578063f154157c146100b4578063f2fde38b146100c7575b5f80fd5b6100616100da565b005b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054604080516001600160a01b039092168252519081900360200190f35b6100616100af3660046103ef565b6100ed565b6100616100c236600461040f565b6101fb565b6100616100d53660046103ef565b610265565b6100e26102a7565b6100eb5f610302565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156101325750825b90505f8267ffffffffffffffff16600114801561014e5750303b155b90508115801561015c575080155b1561017a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156101a457845460ff60401b1916600160401b1785555b6101ad86610372565b83156101f357845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6102036102a7565b7f62f3dd1c4db5f3313296af6a52e5897512a583f9cb18a863c7c126e308b2c98061022e8280610446565b61023b6020850185610446565b6102486040870187610446565b60405161025a9695949392919061054d565b60405180910390a150565b61026d6102a7565b6001600160a01b03811661029b57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6102a481610302565b50565b336102d97f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146100eb5760405163118cdaa760e01b8152336004820152602401610292565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61037a610383565b6102a4816103cc565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166100eb57604051631afcd79f60e31b815260040160405180910390fd5b61026d610383565b80356001600160a01b03811681146103ea575f80fd5b919050565b5f602082840312156103ff575f80fd5b610408826103d4565b9392505050565b5f6020828403121561041f575f80fd5b813567ffffffffffffffff811115610435575f80fd5b820160608185031215610408575f80fd5b5f808335601e1984360301811261045b575f80fd5b83018035915067ffffffffffffffff821115610475575f80fd5b6020019150600581901b360382131561048c575f80fd5b9250929050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f838385526020808601955060208560051b830101845f5b8781101561054057848303601f19018952813536889003601e190181126104f8575f80fd5b8701848101903567ffffffffffffffff811115610513575f80fd5b803603821315610521575f80fd5b61052c858284610493565b9a86019a94505050908301906001016104d3565b5090979650505050505050565b606080825281018690525f8760808301825b8981101561058d576001600160a01b03610578846103d4565b1682526020928301929091019060010161055f565b5083810360208501526105a181888a6104bb565b91505082810360408401526105b78185876104bb565b999850505050505050505056fea264697066735822122069c4bbe9fb5cfffbecc3279486bb350b6eaed526d7ec4605b438b423c93c5a2764736f6c63430008180033",
}

// RedenomABI is the input ABI used to generate the binding from.
// Deprecated: Use RedenomMetaData.ABI instead.
var RedenomABI = RedenomMetaData.ABI

// RedenomBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RedenomMetaData.Bin instead.
var RedenomBin = RedenomMetaData.Bin

// DeployRedenom deploys a new Ethereum contract, binding an instance of Redenom to it.
func DeployRedenom(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Redenom, error) {
	parsed, err := RedenomMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RedenomBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Redenom{RedenomCaller: RedenomCaller{contract: contract}, RedenomTransactor: RedenomTransactor{contract: contract}, RedenomFilterer: RedenomFilterer{contract: contract}}, nil
}

// Redenom is an auto generated Go binding around an Ethereum contract.
type Redenom struct {
	RedenomCaller     // Read-only binding to the contract
	RedenomTransactor // Write-only binding to the contract
	RedenomFilterer   // Log filterer for contract events
}

// RedenomCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedenomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedenomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedenomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedenomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedenomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedenomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedenomSession struct {
	Contract     *Redenom          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedenomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedenomCallerSession struct {
	Contract *RedenomCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RedenomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedenomTransactorSession struct {
	Contract     *RedenomTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RedenomRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedenomRaw struct {
	Contract *Redenom // Generic contract binding to access the raw methods on
}

// RedenomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedenomCallerRaw struct {
	Contract *RedenomCaller // Generic read-only contract binding to access the raw methods on
}

// RedenomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedenomTransactorRaw struct {
	Contract *RedenomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedenom creates a new instance of Redenom, bound to a specific deployed contract.
func NewRedenom(address common.Address, backend bind.ContractBackend) (*Redenom, error) {
	contract, err := bindRedenom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Redenom{RedenomCaller: RedenomCaller{contract: contract}, RedenomTransactor: RedenomTransactor{contract: contract}, RedenomFilterer: RedenomFilterer{contract: contract}}, nil
}

// NewRedenomCaller creates a new read-only instance of Redenom, bound to a specific deployed contract.
func NewRedenomCaller(address common.Address, caller bind.ContractCaller) (*RedenomCaller, error) {
	contract, err := bindRedenom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedenomCaller{contract: contract}, nil
}

// NewRedenomTransactor creates a new write-only instance of Redenom, bound to a specific deployed contract.
func NewRedenomTransactor(address common.Address, transactor bind.ContractTransactor) (*RedenomTransactor, error) {
	contract, err := bindRedenom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedenomTransactor{contract: contract}, nil
}

// NewRedenomFilterer creates a new log filterer instance of Redenom, bound to a specific deployed contract.
func NewRedenomFilterer(address common.Address, filterer bind.ContractFilterer) (*RedenomFilterer, error) {
	contract, err := bindRedenom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedenomFilterer{contract: contract}, nil
}

// bindRedenom binds a generic wrapper to an already deployed contract.
func bindRedenom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RedenomMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redenom *RedenomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redenom.Contract.RedenomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redenom *RedenomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redenom.Contract.RedenomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redenom *RedenomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redenom.Contract.RedenomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redenom *RedenomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redenom.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redenom *RedenomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redenom.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redenom *RedenomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redenom.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Redenom *RedenomCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redenom.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Redenom *RedenomSession) Owner() (common.Address, error) {
	return _Redenom.Contract.Owner(&_Redenom.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Redenom *RedenomCallerSession) Owner() (common.Address, error) {
	return _Redenom.Contract.Owner(&_Redenom.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Redenom *RedenomTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Redenom.contract.Transact(opts, "initialize", owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Redenom *RedenomSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _Redenom.Contract.Initialize(&_Redenom.TransactOpts, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_Redenom *RedenomTransactorSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _Redenom.Contract.Initialize(&_Redenom.TransactOpts, owner_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Redenom *RedenomTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redenom.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Redenom *RedenomSession) RenounceOwnership() (*types.Transaction, error) {
	return _Redenom.Contract.RenounceOwnership(&_Redenom.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Redenom *RedenomTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Redenom.Contract.RenounceOwnership(&_Redenom.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xf154157c.
//
// Solidity: function submit((address[],bytes[],bytes[]) range) returns()
func (_Redenom *RedenomTransactor) Submit(opts *bind.TransactOpts, arg0 RedenomAccountRange) (*types.Transaction, error) {
	return _Redenom.contract.Transact(opts, "submit", arg0)
}

// Submit is a paid mutator transaction binding the contract method 0xf154157c.
//
// Solidity: function submit((address[],bytes[],bytes[]) range) returns()
func (_Redenom *RedenomSession) Submit(arg0 RedenomAccountRange) (*types.Transaction, error) {
	return _Redenom.Contract.Submit(&_Redenom.TransactOpts, arg0)
}

// Submit is a paid mutator transaction binding the contract method 0xf154157c.
//
// Solidity: function submit((address[],bytes[],bytes[]) range) returns()
func (_Redenom *RedenomTransactorSession) Submit(arg0 RedenomAccountRange) (*types.Transaction, error) {
	return _Redenom.Contract.Submit(&_Redenom.TransactOpts, arg0)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Redenom *RedenomTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Redenom.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Redenom *RedenomSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Redenom.Contract.TransferOwnership(&_Redenom.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Redenom *RedenomTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Redenom.Contract.TransferOwnership(&_Redenom.TransactOpts, newOwner)
}

// RedenomInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Redenom contract.
type RedenomInitializedIterator struct {
	Event *RedenomInitialized // Event containing the contract specifics and raw log

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
func (it *RedenomInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedenomInitialized)
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
		it.Event = new(RedenomInitialized)
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
func (it *RedenomInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedenomInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedenomInitialized represents a Initialized event raised by the Redenom contract.
type RedenomInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Redenom *RedenomFilterer) FilterInitialized(opts *bind.FilterOpts) (*RedenomInitializedIterator, error) {

	logs, sub, err := _Redenom.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RedenomInitializedIterator{contract: _Redenom.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Redenom *RedenomFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RedenomInitialized) (event.Subscription, error) {

	logs, sub, err := _Redenom.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedenomInitialized)
				if err := _Redenom.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Redenom *RedenomFilterer) ParseInitialized(log types.Log) (*RedenomInitialized, error) {
	event := new(RedenomInitialized)
	if err := _Redenom.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedenomOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Redenom contract.
type RedenomOwnershipTransferredIterator struct {
	Event *RedenomOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RedenomOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedenomOwnershipTransferred)
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
		it.Event = new(RedenomOwnershipTransferred)
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
func (it *RedenomOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedenomOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedenomOwnershipTransferred represents a OwnershipTransferred event raised by the Redenom contract.
type RedenomOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Redenom *RedenomFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RedenomOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Redenom.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RedenomOwnershipTransferredIterator{contract: _Redenom.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Redenom *RedenomFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RedenomOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Redenom.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedenomOwnershipTransferred)
				if err := _Redenom.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Redenom *RedenomFilterer) ParseOwnershipTransferred(log types.Log) (*RedenomOwnershipTransferred, error) {
	event := new(RedenomOwnershipTransferred)
	if err := _Redenom.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedenomSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the Redenom contract.
type RedenomSubmittedIterator struct {
	Event *RedenomSubmitted // Event containing the contract specifics and raw log

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
func (it *RedenomSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedenomSubmitted)
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
		it.Event = new(RedenomSubmitted)
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
func (it *RedenomSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedenomSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedenomSubmitted represents a Submitted event raised by the Redenom contract.
type RedenomSubmitted struct {
	Addresses []common.Address
	Accounts  [][]byte
	Proof     [][]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x62f3dd1c4db5f3313296af6a52e5897512a583f9cb18a863c7c126e308b2c980.
//
// Solidity: event Submitted(address[] addresses, bytes[] accounts, bytes[] proof)
func (_Redenom *RedenomFilterer) FilterSubmitted(opts *bind.FilterOpts) (*RedenomSubmittedIterator, error) {

	logs, sub, err := _Redenom.contract.FilterLogs(opts, "Submitted")
	if err != nil {
		return nil, err
	}
	return &RedenomSubmittedIterator{contract: _Redenom.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x62f3dd1c4db5f3313296af6a52e5897512a583f9cb18a863c7c126e308b2c980.
//
// Solidity: event Submitted(address[] addresses, bytes[] accounts, bytes[] proof)
func (_Redenom *RedenomFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *RedenomSubmitted) (event.Subscription, error) {

	logs, sub, err := _Redenom.contract.WatchLogs(opts, "Submitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedenomSubmitted)
				if err := _Redenom.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x62f3dd1c4db5f3313296af6a52e5897512a583f9cb18a863c7c126e308b2c980.
//
// Solidity: event Submitted(address[] addresses, bytes[] accounts, bytes[] proof)
func (_Redenom *RedenomFilterer) ParseSubmitted(log types.Log) (*RedenomSubmitted, error) {
	event := new(RedenomSubmitted)
	if err := _Redenom.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
