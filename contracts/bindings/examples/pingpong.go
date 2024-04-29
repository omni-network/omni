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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pingpong\",\"inputs\":[{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"n\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pingpong_norecv\",\"inputs\":[{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"n\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"start\",\"inputs\":[{\"name\":\"destChainID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Done\",\"inputs\":[{\"name\":\"destChainID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"times\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610aa1380380610aa183398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516109f36100ae6000396000818161016b01528181610388015281816104fb0152818161058d015281816105b901526106bc01526109f36000f3fe6080604052600436106100435760003560e01c8063091d27881461004f5780634d4fc29314610082578063645953db146100a457806365aef3f4146100c457600080fd5b3661004a57005b600080fd5b34801561005b57600080fd5b50610066620186a081565b6040516001600160401b03909116815260200160405180910390f35b34801561008e57600080fd5b506100a261009d366004610762565b6100e4565b005b3480156100b057600080fd5b506100a26100bf3660046107ad565b610169565b3480156100d057600080fd5b506100a26100df3660046107ad565b610330565b6000816001600160401b0316116101425760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a2074696d6573206d757374206265203e2030000000000060448201526064015b60405180910390fd5b61016483838360016101558260026107fc565b61015f9190610827565b610494565b505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa1580156101c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101ea919061084e565b8051600080546020909301516001600160a01b0316600160401b026001600160e01b03199093166001600160401b039092169190911791909117905561022e6104f7565b61027a5760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a206e6f7420616e206f6d6e69207863616c6c00000000006044820152606401610139565b806001600160401b03166000036102ee57600054604080516001600160401b038084168252600160401b9093046001600160a01b03166020820152918416908201527fa73c132355e7f2c5ced079328c310d267f1e830c197c5579b6177218a777313e9060600160405180910390a161031c565b60005461031c906001600160401b03811690600160401b90046001600160a01b03168461015f600186610827565b5050600080546001600160e01b0319169055565b6103386104f7565b6103845760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a206e6f7420616e206f6d6e69207863616c6c00000000006044820152606401610139565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa1580156103e3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610407919061084e565b9050816001600160401b031660000361047c578051602080830151604080516001600160401b0394851681526001600160a01b039092169282019290925291851682820152517fa73c132355e7f2c5ced079328c310d267f1e830c197c5579b6177218a777313e9181900360600190a1505050565b8051602082015161016491908561015f600187610827565b604080516001600160401b03848116602483015283166044808301919091528251808303909101815260649091019091526020810180516001600160e01b031663645953db60e01b1790526104f09085908590620186a06105b4565b5050505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166355e2448e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610557573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057b91906108ba565b80156105af5750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b905090565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638dd9523c8786866040518463ffffffff1660e01b815260040161060793929190610929565b602060405180830381865afa158015610624573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610648919061095f565b905080471015806106595750803410155b6106a55760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401610139565b6040516338745ab560e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370e8b56a9083906106f9908a908a908a908a90600401610978565b6000604051808303818588803b15801561071257600080fd5b505af1158015610726573d6000803e3d6000fd5b50939998505050505050505050565b6001600160401b038116811461074a57600080fd5b50565b6001600160a01b038116811461074a57600080fd5b60008060006060848603121561077757600080fd5b833561078281610735565b925060208401356107928161074d565b915060408401356107a281610735565b809150509250925092565b600080604083850312156107c057600080fd5b82356107cb81610735565b915060208301356107db81610735565b809150509250929050565b634e487b7160e01b600052601160045260246000fd5b6001600160401b0381811683821602808216919082811461081f5761081f6107e6565b505092915050565b6001600160401b03828116828216039080821115610847576108476107e6565b5092915050565b60006040828403121561086057600080fd5b604051604081018181106001600160401b038211171561089057634e487b7160e01b600052604160045260246000fd5b604052825161089e81610735565b815260208301516108ae8161074d565b60208201529392505050565b6000602082840312156108cc57600080fd5b815180151581146108dc57600080fd5b9392505050565b6000815180845260005b81811015610909576020818501810151868301820152016108ed565b506000602082860101526020601f19601f83011685010191505092915050565b60006001600160401b0380861683526060602084015261094c60608401866108e3565b9150808416604084015250949350505050565b60006020828403121561097157600080fd5b5051919050565b60006001600160401b03808716835260018060a01b0386166020840152608060408401526109a960808401866108e3565b91508084166060840152509594505050505056fea2646970667358221220e2366274dcae076daf044225ae4a85386ebedcdbd319b608ecf58f06cfca72fe64736f6c63430008180033",
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

// PingpongNorecv is a paid mutator transaction binding the contract method 0x65aef3f4.
//
// Solidity: function pingpong_norecv(uint64 times, uint64 n) returns()
func (_PingPong *PingPongTransactor) PingpongNorecv(opts *bind.TransactOpts, times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.contract.Transact(opts, "pingpong_norecv", times, n)
}

// PingpongNorecv is a paid mutator transaction binding the contract method 0x65aef3f4.
//
// Solidity: function pingpong_norecv(uint64 times, uint64 n) returns()
func (_PingPong *PingPongSession) PingpongNorecv(times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.Contract.PingpongNorecv(&_PingPong.TransactOpts, times, n)
}

// PingpongNorecv is a paid mutator transaction binding the contract method 0x65aef3f4.
//
// Solidity: function pingpong_norecv(uint64 times, uint64 n) returns()
func (_PingPong *PingPongTransactorSession) PingpongNorecv(times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.Contract.PingpongNorecv(&_PingPong.TransactOpts, times, n)
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
