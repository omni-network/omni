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

// AdminMetaData contains all meta data concerning the Admin contract.
var AdminMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_SCRIPT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pausePortal\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpausePortal\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052600c805462ff00ff19166201000117905534801561002157600080fd5b506102ee806100316000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063690e19b014610046578063d31e1cea1461005b578063f8ccbf471461006e575b600080fd5b610059610054366004610288565b610095565b005b610059610069366004610288565b6101d7565b600c546100819062010000900460ff1681565b604051901515815260200160405180910390f35b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b0316637fb5297f6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156100f357600080fd5b505af1158015610107573d6000803e3d6000fd5b50505050806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561014657600080fd5b505af115801561015a573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166376eadd366040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156101bc57600080fd5b505af11580156101d0573d6000803e3d6000fd5b5050505050565b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b0316637fb5297f6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561023557600080fd5b505af1158015610249573d6000803e3d6000fd5b50505050806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561014657600080fd5b60006020828403121561029a57600080fd5b81356001600160a01b03811681146102b157600080fd5b939250505056fea2646970667358221220d52a3d322bc2cc80adbd8721d357be971a333cdbb1810a65cbfa01cb04c5e47564736f6c63430008180033",
}

// AdminABI is the input ABI used to generate the binding from.
// Deprecated: Use AdminMetaData.ABI instead.
var AdminABI = AdminMetaData.ABI

// AdminBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdminMetaData.Bin instead.
var AdminBin = AdminMetaData.Bin

// DeployAdmin deploys a new Ethereum contract, binding an instance of Admin to it.
func DeployAdmin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Admin, error) {
	parsed, err := AdminMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdminBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// Admin is an auto generated Go binding around an Ethereum contract.
type Admin struct {
	AdminCaller     // Read-only binding to the contract
	AdminTransactor // Write-only binding to the contract
	AdminFilterer   // Log filterer for contract events
}

// AdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminSession struct {
	Contract     *Admin            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminCallerSession struct {
	Contract *AdminCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminTransactorSession struct {
	Contract     *AdminTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminRaw struct {
	Contract *Admin // Generic contract binding to access the raw methods on
}

// AdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminCallerRaw struct {
	Contract *AdminCaller // Generic read-only contract binding to access the raw methods on
}

// AdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminTransactorRaw struct {
	Contract *AdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdmin creates a new instance of Admin, bound to a specific deployed contract.
func NewAdmin(address common.Address, backend bind.ContractBackend) (*Admin, error) {
	contract, err := bindAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// NewAdminCaller creates a new read-only instance of Admin, bound to a specific deployed contract.
func NewAdminCaller(address common.Address, caller bind.ContractCaller) (*AdminCaller, error) {
	contract, err := bindAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminCaller{contract: contract}, nil
}

// NewAdminTransactor creates a new write-only instance of Admin, bound to a specific deployed contract.
func NewAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminTransactor, error) {
	contract, err := bindAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminTransactor{contract: contract}, nil
}

// NewAdminFilterer creates a new log filterer instance of Admin, bound to a specific deployed contract.
func NewAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminFilterer, error) {
	contract, err := bindAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminFilterer{contract: contract}, nil
}

// bindAdmin binds a generic wrapper to an already deployed contract.
func bindAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.AdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transact(opts, method, params...)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_Admin *AdminCaller) ISSCRIPT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "IS_SCRIPT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_Admin *AdminSession) ISSCRIPT() (bool, error) {
	return _Admin.Contract.ISSCRIPT(&_Admin.CallOpts)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_Admin *AdminCallerSession) ISSCRIPT() (bool, error) {
	return _Admin.Contract.ISSCRIPT(&_Admin.CallOpts)
}

// PausePortal is a paid mutator transaction binding the contract method 0x690e19b0.
//
// Solidity: function pausePortal(address portal) returns()
func (_Admin *AdminTransactor) PausePortal(opts *bind.TransactOpts, portal common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "pausePortal", portal)
}

// PausePortal is a paid mutator transaction binding the contract method 0x690e19b0.
//
// Solidity: function pausePortal(address portal) returns()
func (_Admin *AdminSession) PausePortal(portal common.Address) (*types.Transaction, error) {
	return _Admin.Contract.PausePortal(&_Admin.TransactOpts, portal)
}

// PausePortal is a paid mutator transaction binding the contract method 0x690e19b0.
//
// Solidity: function pausePortal(address portal) returns()
func (_Admin *AdminTransactorSession) PausePortal(portal common.Address) (*types.Transaction, error) {
	return _Admin.Contract.PausePortal(&_Admin.TransactOpts, portal)
}

// UnpausePortal is a paid mutator transaction binding the contract method 0xd31e1cea.
//
// Solidity: function unpausePortal(address portal) returns()
func (_Admin *AdminTransactor) UnpausePortal(opts *bind.TransactOpts, portal common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "unpausePortal", portal)
}

// UnpausePortal is a paid mutator transaction binding the contract method 0xd31e1cea.
//
// Solidity: function unpausePortal(address portal) returns()
func (_Admin *AdminSession) UnpausePortal(portal common.Address) (*types.Transaction, error) {
	return _Admin.Contract.UnpausePortal(&_Admin.TransactOpts, portal)
}

// UnpausePortal is a paid mutator transaction binding the contract method 0xd31e1cea.
//
// Solidity: function unpausePortal(address portal) returns()
func (_Admin *AdminTransactorSession) UnpausePortal(portal common.Address) (*types.Transaction, error) {
	return _Admin.Contract.UnpausePortal(&_Admin.TransactOpts, portal)
}
