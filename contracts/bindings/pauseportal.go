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

// PausePortalMetaData contains all meta data concerning the PausePortal contract.
var PausePortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_SCRIPT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"run\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052600c805462ff00ff19166201000117905534801561002157600080fd5b5061021f806100316000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063522bb7041461003b578063f8ccbf4714610050575b600080fd5b61004e6100493660046101b9565b610077565b005b600c546100639062010000900460ff1681565b604051901515815260200160405180910390f35b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b0316637fb5297f6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156100d557600080fd5b505af11580156100e9573d6000803e3d6000fd5b50505050806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561012857600080fd5b505af115801561013c573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166376eadd366040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561019e57600080fd5b505af11580156101b2573d6000803e3d6000fd5b5050505050565b6000602082840312156101cb57600080fd5b81356001600160a01b03811681146101e257600080fd5b939250505056fea26469706673582212205c6ff86a97d468038b138b795d987116497b07bf48e0d98a38dc12a9c015a6a664736f6c63430008180033",
}

// PausePortalABI is the input ABI used to generate the binding from.
// Deprecated: Use PausePortalMetaData.ABI instead.
var PausePortalABI = PausePortalMetaData.ABI

// PausePortalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PausePortalMetaData.Bin instead.
var PausePortalBin = PausePortalMetaData.Bin

// DeployPausePortal deploys a new Ethereum contract, binding an instance of PausePortal to it.
func DeployPausePortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PausePortal, error) {
	parsed, err := PausePortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PausePortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PausePortal{PausePortalCaller: PausePortalCaller{contract: contract}, PausePortalTransactor: PausePortalTransactor{contract: contract}, PausePortalFilterer: PausePortalFilterer{contract: contract}}, nil
}

// PausePortal is an auto generated Go binding around an Ethereum contract.
type PausePortal struct {
	PausePortalCaller     // Read-only binding to the contract
	PausePortalTransactor // Write-only binding to the contract
	PausePortalFilterer   // Log filterer for contract events
}

// PausePortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausePortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausePortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausePortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausePortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausePortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausePortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausePortalSession struct {
	Contract     *PausePortal      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausePortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausePortalCallerSession struct {
	Contract *PausePortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PausePortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausePortalTransactorSession struct {
	Contract     *PausePortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PausePortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausePortalRaw struct {
	Contract *PausePortal // Generic contract binding to access the raw methods on
}

// PausePortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausePortalCallerRaw struct {
	Contract *PausePortalCaller // Generic read-only contract binding to access the raw methods on
}

// PausePortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausePortalTransactorRaw struct {
	Contract *PausePortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausePortal creates a new instance of PausePortal, bound to a specific deployed contract.
func NewPausePortal(address common.Address, backend bind.ContractBackend) (*PausePortal, error) {
	contract, err := bindPausePortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausePortal{PausePortalCaller: PausePortalCaller{contract: contract}, PausePortalTransactor: PausePortalTransactor{contract: contract}, PausePortalFilterer: PausePortalFilterer{contract: contract}}, nil
}

// NewPausePortalCaller creates a new read-only instance of PausePortal, bound to a specific deployed contract.
func NewPausePortalCaller(address common.Address, caller bind.ContractCaller) (*PausePortalCaller, error) {
	contract, err := bindPausePortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausePortalCaller{contract: contract}, nil
}

// NewPausePortalTransactor creates a new write-only instance of PausePortal, bound to a specific deployed contract.
func NewPausePortalTransactor(address common.Address, transactor bind.ContractTransactor) (*PausePortalTransactor, error) {
	contract, err := bindPausePortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausePortalTransactor{contract: contract}, nil
}

// NewPausePortalFilterer creates a new log filterer instance of PausePortal, bound to a specific deployed contract.
func NewPausePortalFilterer(address common.Address, filterer bind.ContractFilterer) (*PausePortalFilterer, error) {
	contract, err := bindPausePortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausePortalFilterer{contract: contract}, nil
}

// bindPausePortal binds a generic wrapper to an already deployed contract.
func bindPausePortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PausePortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausePortal *PausePortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausePortal.Contract.PausePortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausePortal *PausePortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausePortal.Contract.PausePortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausePortal *PausePortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausePortal.Contract.PausePortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausePortal *PausePortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausePortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausePortal *PausePortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausePortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausePortal *PausePortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausePortal.Contract.contract.Transact(opts, method, params...)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_PausePortal *PausePortalCaller) ISSCRIPT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PausePortal.contract.Call(opts, &out, "IS_SCRIPT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_PausePortal *PausePortalSession) ISSCRIPT() (bool, error) {
	return _PausePortal.Contract.ISSCRIPT(&_PausePortal.CallOpts)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_PausePortal *PausePortalCallerSession) ISSCRIPT() (bool, error) {
	return _PausePortal.Contract.ISSCRIPT(&_PausePortal.CallOpts)
}

// Run is a paid mutator transaction binding the contract method 0x522bb704.
//
// Solidity: function run(address portal) returns()
func (_PausePortal *PausePortalTransactor) Run(opts *bind.TransactOpts, portal common.Address) (*types.Transaction, error) {
	return _PausePortal.contract.Transact(opts, "run", portal)
}

// Run is a paid mutator transaction binding the contract method 0x522bb704.
//
// Solidity: function run(address portal) returns()
func (_PausePortal *PausePortalSession) Run(portal common.Address) (*types.Transaction, error) {
	return _PausePortal.Contract.Run(&_PausePortal.TransactOpts, portal)
}

// Run is a paid mutator transaction binding the contract method 0x522bb704.
//
// Solidity: function run(address portal) returns()
func (_PausePortal *PausePortalTransactorSession) Run(portal common.Address) (*types.Transaction, error) {
	return _PausePortal.Contract.Run(&_PausePortal.TransactOpts, portal)
}
