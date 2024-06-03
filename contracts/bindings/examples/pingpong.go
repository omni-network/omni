// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examples

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

// PingPongMetaData contains all meta data concerning the PingPong contract.
var PingPongMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pingpong\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"n\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"start\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destChainID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Done\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destChainID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"times\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Ping\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"srcChainID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"n\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610acc380380610acc83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610a266100a66000396000818160bb0152818161036b015281816103fd0152818161049301526105960152610a266000f3fe6080604052600436106100385760003560e01c8063091d278814610044578063a04444fa14610077578063d32f5cda1461009957600080fd5b3661003f57005b600080fd5b34801561005057600080fd5b5061005b62030d4081565b6040516001600160401b03909116815260200160405180910390f35b34801561008357600080fd5b50610097610092366004610672565b6100b9565b005b3480156100a557600080fd5b506100976100b43660046106ee565b6102e8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa158015610116573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061013a919061076a565b8051600080546020909301516001600160a01b0316600160401b026001600160e01b03199093166001600160401b039092169190911791909117905561017e610367565b6101cf5760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a206e6f7420616e206f6d6e69207863616c6c000000000060448201526064015b60405180910390fd5b6000546040517f47768d18c1bf0d890e19b6f020bd996b385273187be95e2b6a58feccf3bd75669161022291879187916001600160401b03821691600160401b90046001600160a01b03169087906107ff565b60405180910390a1806001600160401b031660000361029b576000546040517f5335290c57b37808bcd506c1e858915c89a2c9adaa76dc0319044eab5b9d18c29161028e91879187916001600160401b03821691600160401b90046001600160a01b03169088906107ff565b60405180910390a16102d2565b6000546102d290859085906001600160401b03811690600160401b90046001600160a01b0316866102cd60018861085c565b610424565b5050600080546001600160e01b03191690555050565b6000816001600160401b0316116103415760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a2074696d6573206d757374206265203e2030000000000060448201526064016101c6565b61036085858585856001610356826002610883565b6102cd919061085c565b5050505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166355e2448e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103eb91906108ae565b801561041f5750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b905090565b6104858460018563a04444fa60e01b8a8a888860405160240161044a94939291906108d7565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915262030d4061048e565b50505050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638dd9523c8886866040518463ffffffff1660e01b81526004016104e193929190610952565b602060405180830381865afa1580156104fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105229190610988565b905080471015806105335750803410155b61057f5760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e6473000000000000000060448201526064016101c6565b60405163c21dda4f60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c21dda4f9083906105d5908b908b908b908b908b906004016109a1565b6000604051808303818588803b1580156105ee57600080fd5b505af1158015610602573d6000803e3d6000fd5b50939a9950505050505050505050565b60008083601f84011261062457600080fd5b5081356001600160401b0381111561063b57600080fd5b60208301915083602082850101111561065357600080fd5b9250929050565b6001600160401b038116811461066f57600080fd5b50565b6000806000806060858703121561068857600080fd5b84356001600160401b0381111561069e57600080fd5b6106aa87828801610612565b90955093505060208501356106be8161065a565b915060408501356106ce8161065a565b939692955090935050565b6001600160a01b038116811461066f57600080fd5b60008060008060006080868803121561070657600080fd5b85356001600160401b0381111561071c57600080fd5b61072888828901610612565b909650945050602086013561073c8161065a565b9250604086013561074c816106d9565b9150606086013561075c8161065a565b809150509295509295909350565b60006040828403121561077c57600080fd5b604051604081018181106001600160401b03821117156107ac57634e487b7160e01b600052604160045260246000fd5b60405282516107ba8161065a565b815260208301516107ca816106d9565b60208201529392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6080815260006108136080830187896107d6565b6001600160401b0395861660208401526001600160a01b0394909416604083015250921660609092019190915292915050565b634e487b7160e01b600052601160045260246000fd5b6001600160401b0382811682821603908082111561087c5761087c610846565b5092915050565b6001600160401b038181168382160280821691908281146108a6576108a6610846565b505092915050565b6000602082840312156108c057600080fd5b815180151581146108d057600080fd5b9392505050565b6060815260006108eb6060830186886107d6565b6001600160401b039485166020840152929093166040909101529392505050565b6000815180845260005b8181101561093257602081850181015186830182015201610916565b506000602082860101526020601f19601f83011685010191505092915050565b60006001600160401b03808616835260606020840152610975606084018661090c565b9150808416604084015250949350505050565b60006020828403121561099a57600080fd5b5051919050565b60006001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a060608401526109db60a084018661090c565b9150808416608084015250969550505050505056fea2646970667358221220abd3a1f14296976e105313063e58b971afb3568781f20e003f50eb3580ed10b664736f6c63430008180033",
}

// PingPongABI is the input ABI used to generate the binding from.
// Deprecated: Use PingPongMetaData.ABI instead.
var PingPongABI = PingPongMetaData.ABI

// PingPongBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PingPongMetaData.Bin instead.
var PingPongBin = PingPongMetaData.Bin

// DeployPingPong deploys a new Ethereum contract, binding an instance of PingPong to it.
func DeployPingPong(auth *bind.TransactOpts, backend bind.ContractBackend, portal common.Address) (common.Address, *types.Transaction, *PingPong, error) {
	parsed, err := PingPongMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PingPongBin), backend, portal)
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
	parsed, err := PingPongMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint64)
func (_PingPong *PingPongCaller) GASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PingPong.contract.Call(opts, &out, "GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint64)
func (_PingPong *PingPongSession) GASLIMIT() (uint64, error) {
	return _PingPong.Contract.GASLIMIT(&_PingPong.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint64)
func (_PingPong *PingPongCallerSession) GASLIMIT() (uint64, error) {
	return _PingPong.Contract.GASLIMIT(&_PingPong.CallOpts)
}

// Pingpong is a paid mutator transaction binding the contract method 0xa04444fa.
//
// Solidity: function pingpong(string id, uint64 times, uint64 n) returns()
func (_PingPong *PingPongTransactor) Pingpong(opts *bind.TransactOpts, id string, times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.contract.Transact(opts, "pingpong", id, times, n)
}

// Pingpong is a paid mutator transaction binding the contract method 0xa04444fa.
//
// Solidity: function pingpong(string id, uint64 times, uint64 n) returns()
func (_PingPong *PingPongSession) Pingpong(id string, times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Pingpong(&_PingPong.TransactOpts, id, times, n)
}

// Pingpong is a paid mutator transaction binding the contract method 0xa04444fa.
//
// Solidity: function pingpong(string id, uint64 times, uint64 n) returns()
func (_PingPong *PingPongTransactorSession) Pingpong(id string, times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Pingpong(&_PingPong.TransactOpts, id, times, n)
}

// Start is a paid mutator transaction binding the contract method 0xd32f5cda.
//
// Solidity: function start(string id, uint64 destChainID, address to, uint64 times) returns()
func (_PingPong *PingPongTransactor) Start(opts *bind.TransactOpts, id string, destChainID uint64, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.contract.Transact(opts, "start", id, destChainID, to, times)
}

// Start is a paid mutator transaction binding the contract method 0xd32f5cda.
//
// Solidity: function start(string id, uint64 destChainID, address to, uint64 times) returns()
func (_PingPong *PingPongSession) Start(id string, destChainID uint64, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Start(&_PingPong.TransactOpts, id, destChainID, to, times)
}

// Start is a paid mutator transaction binding the contract method 0xd32f5cda.
//
// Solidity: function start(string id, uint64 destChainID, address to, uint64 times) returns()
func (_PingPong *PingPongTransactorSession) Start(id string, destChainID uint64, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Start(&_PingPong.TransactOpts, id, destChainID, to, times)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PingPong *PingPongTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PingPong *PingPongSession) Receive() (*types.Transaction, error) {
	return _PingPong.Contract.Receive(&_PingPong.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PingPong *PingPongTransactorSession) Receive() (*types.Transaction, error) {
	return _PingPong.Contract.Receive(&_PingPong.TransactOpts)
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
	Id          string
	DestChainID uint64
	To          common.Address
	Times       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDone is a free log retrieval operation binding the contract event 0x5335290c57b37808bcd506c1e858915c89a2c9adaa76dc0319044eab5b9d18c2.
//
// Solidity: event Done(string id, uint64 destChainID, address to, uint64 times)
func (_PingPong *PingPongFilterer) FilterDone(opts *bind.FilterOpts) (*PingPongDoneIterator, error) {

	logs, sub, err := _PingPong.contract.FilterLogs(opts, "Done")
	if err != nil {
		return nil, err
	}
	return &PingPongDoneIterator{contract: _PingPong.contract, event: "Done", logs: logs, sub: sub}, nil
}

// WatchDone is a free log subscription operation binding the contract event 0x5335290c57b37808bcd506c1e858915c89a2c9adaa76dc0319044eab5b9d18c2.
//
// Solidity: event Done(string id, uint64 destChainID, address to, uint64 times)
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

// ParseDone is a log parse operation binding the contract event 0x5335290c57b37808bcd506c1e858915c89a2c9adaa76dc0319044eab5b9d18c2.
//
// Solidity: event Done(string id, uint64 destChainID, address to, uint64 times)
func (_PingPong *PingPongFilterer) ParseDone(log types.Log) (*PingPongDone, error) {
	event := new(PingPongDone)
	if err := _PingPong.contract.UnpackLog(event, "Done", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingPongPingIterator is returned from FilterPing and is used to iterate over the raw logs and unpacked data for Ping events raised by the PingPong contract.
type PingPongPingIterator struct {
	Event *PingPongPing // Event containing the contract specifics and raw log

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
func (it *PingPongPingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongPing)
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
		it.Event = new(PingPongPing)
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
func (it *PingPongPingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongPingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongPing represents a Ping event raised by the PingPong contract.
type PingPongPing struct {
	Id         string
	SrcChainID uint64
	From       common.Address
	N          uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPing is a free log retrieval operation binding the contract event 0x47768d18c1bf0d890e19b6f020bd996b385273187be95e2b6a58feccf3bd7566.
//
// Solidity: event Ping(string id, uint64 srcChainID, address from, uint64 n)
func (_PingPong *PingPongFilterer) FilterPing(opts *bind.FilterOpts) (*PingPongPingIterator, error) {

	logs, sub, err := _PingPong.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &PingPongPingIterator{contract: _PingPong.contract, event: "Ping", logs: logs, sub: sub}, nil
}

// WatchPing is a free log subscription operation binding the contract event 0x47768d18c1bf0d890e19b6f020bd996b385273187be95e2b6a58feccf3bd7566.
//
// Solidity: event Ping(string id, uint64 srcChainID, address from, uint64 n)
func (_PingPong *PingPongFilterer) WatchPing(opts *bind.WatchOpts, sink chan<- *PingPongPing) (event.Subscription, error) {

	logs, sub, err := _PingPong.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongPing)
				if err := _PingPong.contract.UnpackLog(event, "Ping", log); err != nil {
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

// ParsePing is a log parse operation binding the contract event 0x47768d18c1bf0d890e19b6f020bd996b385273187be95e2b6a58feccf3bd7566.
//
// Solidity: event Ping(string id, uint64 srcChainID, address from, uint64 n)
func (_PingPong *PingPongFilterer) ParsePing(log types.Log) (*PingPongPing, error) {
	event := new(PingPongPing)
	if err := _PingPong.contract.UnpackLog(event, "Ping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
