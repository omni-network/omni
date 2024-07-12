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

// UnpausePortalMetaData contains all meta data concerning the UnpausePortal contract.
var UnpausePortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_SCRIPT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"run\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052600c805462ff00ff19166201000117905534801561002157600080fd5b5061021f806100316000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063522bb7041461003b578063f8ccbf4714610050575b600080fd5b61004e6100493660046101b9565b610077565b005b600c546100639062010000900460ff1681565b604051901515815260200160405180910390f35b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b0316637fb5297f6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156100d557600080fd5b505af11580156100e9573d6000803e3d6000fd5b50505050806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561012857600080fd5b505af115801561013c573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166376eadd366040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561019e57600080fd5b505af11580156101b2573d6000803e3d6000fd5b5050505050565b6000602082840312156101cb57600080fd5b81356001600160a01b03811681146101e257600080fd5b939250505056fea264697066735822122002b6accc43378ebb855731d76a3635c092aa09aa671d5ab5c4a2508252d52ea664736f6c63430008180033",
}

// UnpausePortalABI is the input ABI used to generate the binding from.
// Deprecated: Use UnpausePortalMetaData.ABI instead.
var UnpausePortalABI = UnpausePortalMetaData.ABI

// UnpausePortalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UnpausePortalMetaData.Bin instead.
var UnpausePortalBin = UnpausePortalMetaData.Bin

// DeployUnpausePortal deploys a new Ethereum contract, binding an instance of UnpausePortal to it.
func DeployUnpausePortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnpausePortal, error) {
	parsed, err := UnpausePortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UnpausePortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnpausePortal{UnpausePortalCaller: UnpausePortalCaller{contract: contract}, UnpausePortalTransactor: UnpausePortalTransactor{contract: contract}, UnpausePortalFilterer: UnpausePortalFilterer{contract: contract}}, nil
}

// UnpausePortal is an auto generated Go binding around an Ethereum contract.
type UnpausePortal struct {
	UnpausePortalCaller     // Read-only binding to the contract
	UnpausePortalTransactor // Write-only binding to the contract
	UnpausePortalFilterer   // Log filterer for contract events
}

// UnpausePortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnpausePortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnpausePortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnpausePortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnpausePortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnpausePortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnpausePortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnpausePortalSession struct {
	Contract     *UnpausePortal    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnpausePortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnpausePortalCallerSession struct {
	Contract *UnpausePortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UnpausePortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnpausePortalTransactorSession struct {
	Contract     *UnpausePortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UnpausePortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnpausePortalRaw struct {
	Contract *UnpausePortal // Generic contract binding to access the raw methods on
}

// UnpausePortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnpausePortalCallerRaw struct {
	Contract *UnpausePortalCaller // Generic read-only contract binding to access the raw methods on
}

// UnpausePortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnpausePortalTransactorRaw struct {
	Contract *UnpausePortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnpausePortal creates a new instance of UnpausePortal, bound to a specific deployed contract.
func NewUnpausePortal(address common.Address, backend bind.ContractBackend) (*UnpausePortal, error) {
	contract, err := bindUnpausePortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnpausePortal{UnpausePortalCaller: UnpausePortalCaller{contract: contract}, UnpausePortalTransactor: UnpausePortalTransactor{contract: contract}, UnpausePortalFilterer: UnpausePortalFilterer{contract: contract}}, nil
}

// NewUnpausePortalCaller creates a new read-only instance of UnpausePortal, bound to a specific deployed contract.
func NewUnpausePortalCaller(address common.Address, caller bind.ContractCaller) (*UnpausePortalCaller, error) {
	contract, err := bindUnpausePortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnpausePortalCaller{contract: contract}, nil
}

// NewUnpausePortalTransactor creates a new write-only instance of UnpausePortal, bound to a specific deployed contract.
func NewUnpausePortalTransactor(address common.Address, transactor bind.ContractTransactor) (*UnpausePortalTransactor, error) {
	contract, err := bindUnpausePortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnpausePortalTransactor{contract: contract}, nil
}

// NewUnpausePortalFilterer creates a new log filterer instance of UnpausePortal, bound to a specific deployed contract.
func NewUnpausePortalFilterer(address common.Address, filterer bind.ContractFilterer) (*UnpausePortalFilterer, error) {
	contract, err := bindUnpausePortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnpausePortalFilterer{contract: contract}, nil
}

// bindUnpausePortal binds a generic wrapper to an already deployed contract.
func bindUnpausePortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UnpausePortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnpausePortal *UnpausePortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnpausePortal.Contract.UnpausePortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnpausePortal *UnpausePortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnpausePortal.Contract.UnpausePortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnpausePortal *UnpausePortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnpausePortal.Contract.UnpausePortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnpausePortal *UnpausePortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnpausePortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnpausePortal *UnpausePortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnpausePortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnpausePortal *UnpausePortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnpausePortal.Contract.contract.Transact(opts, method, params...)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_UnpausePortal *UnpausePortalCaller) ISSCRIPT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UnpausePortal.contract.Call(opts, &out, "IS_SCRIPT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_UnpausePortal *UnpausePortalSession) ISSCRIPT() (bool, error) {
	return _UnpausePortal.Contract.ISSCRIPT(&_UnpausePortal.CallOpts)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_UnpausePortal *UnpausePortalCallerSession) ISSCRIPT() (bool, error) {
	return _UnpausePortal.Contract.ISSCRIPT(&_UnpausePortal.CallOpts)
}

// Run is a paid mutator transaction binding the contract method 0x522bb704.
//
// Solidity: function run(address portal) returns()
func (_UnpausePortal *UnpausePortalTransactor) Run(opts *bind.TransactOpts, portal common.Address) (*types.Transaction, error) {
	return _UnpausePortal.contract.Transact(opts, "run", portal)
}

// Run is a paid mutator transaction binding the contract method 0x522bb704.
//
// Solidity: function run(address portal) returns()
func (_UnpausePortal *UnpausePortalSession) Run(portal common.Address) (*types.Transaction, error) {
	return _UnpausePortal.Contract.Run(&_UnpausePortal.TransactOpts, portal)
}

// Run is a paid mutator transaction binding the contract method 0x522bb704.
//
// Solidity: function run(address portal) returns()
func (_UnpausePortal *UnpausePortalTransactorSession) Run(portal common.Address) (*types.Transaction, error) {
	return _UnpausePortal.Contract.Run(&_UnpausePortal.TransactOpts, portal)
}
