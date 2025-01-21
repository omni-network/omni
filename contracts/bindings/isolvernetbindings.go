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

// ISolverNetCall is an auto generated low-level Go binding around an user-defined struct.
type ISolverNetCall struct {
	ChainId  uint64
	Target   [32]byte
	Value    *big.Int
	Data     []byte
	Expenses []ISolverNetTokenExpense
}

// ISolverNetDeposit is an auto generated low-level Go binding around an user-defined struct.
type ISolverNetDeposit struct {
	Token  [32]byte
	Amount *big.Int
}

// ISolverNetFillOriginData is an auto generated low-level Go binding around an user-defined struct.
type ISolverNetFillOriginData struct {
	SrcChainId   uint64
	FillDeadline *big.Int
	Call         ISolverNetCall
}

// ISolverNetOrderData is an auto generated low-level Go binding around an user-defined struct.
type ISolverNetOrderData struct {
	Call     ISolverNetCall
	Deposits []ISolverNetDeposit
}

// ISolverNetTokenExpense is an auto generated low-level Go binding around an user-defined struct.
type ISolverNetTokenExpense struct {
	Token   [32]byte
	Spender [32]byte
	Amount  *big.Int
}

// ISolverNetBindingsMetaData contains all meta data concerning the ISolverNetBindings contract.
var ISolverNetBindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISolverNet.Call\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expenses\",\"type\":\"tuple[]\",\"internalType\":\"structISolverNet.TokenExpense[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"spender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISolverNet.Deposit\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fillOriginData\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISolverNet.FillOriginData\",\"components\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fillDeadline\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structISolverNet.Call\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expenses\",\"type\":\"tuple[]\",\"internalType\":\"structISolverNet.TokenExpense[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"spender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderData\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISolverNet.OrderData\",\"components\":[{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structISolverNet.Call\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expenses\",\"type\":\"tuple[]\",\"internalType\":\"structISolverNet.TokenExpense[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"spender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structISolverNet.Deposit[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenExpense\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISolverNet.TokenExpense\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"spender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"}]",
}

// ISolverNetBindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use ISolverNetBindingsMetaData.ABI instead.
var ISolverNetBindingsABI = ISolverNetBindingsMetaData.ABI

// ISolverNetBindings is an auto generated Go binding around an Ethereum contract.
type ISolverNetBindings struct {
	ISolverNetBindingsCaller     // Read-only binding to the contract
	ISolverNetBindingsTransactor // Write-only binding to the contract
	ISolverNetBindingsFilterer   // Log filterer for contract events
}

// ISolverNetBindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISolverNetBindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISolverNetBindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISolverNetBindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISolverNetBindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISolverNetBindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISolverNetBindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISolverNetBindingsSession struct {
	Contract     *ISolverNetBindings // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISolverNetBindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISolverNetBindingsCallerSession struct {
	Contract *ISolverNetBindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ISolverNetBindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISolverNetBindingsTransactorSession struct {
	Contract     *ISolverNetBindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ISolverNetBindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISolverNetBindingsRaw struct {
	Contract *ISolverNetBindings // Generic contract binding to access the raw methods on
}

// ISolverNetBindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISolverNetBindingsCallerRaw struct {
	Contract *ISolverNetBindingsCaller // Generic read-only contract binding to access the raw methods on
}

// ISolverNetBindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISolverNetBindingsTransactorRaw struct {
	Contract *ISolverNetBindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISolverNetBindings creates a new instance of ISolverNetBindings, bound to a specific deployed contract.
func NewISolverNetBindings(address common.Address, backend bind.ContractBackend) (*ISolverNetBindings, error) {
	contract, err := bindISolverNetBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISolverNetBindings{ISolverNetBindingsCaller: ISolverNetBindingsCaller{contract: contract}, ISolverNetBindingsTransactor: ISolverNetBindingsTransactor{contract: contract}, ISolverNetBindingsFilterer: ISolverNetBindingsFilterer{contract: contract}}, nil
}

// NewISolverNetBindingsCaller creates a new read-only instance of ISolverNetBindings, bound to a specific deployed contract.
func NewISolverNetBindingsCaller(address common.Address, caller bind.ContractCaller) (*ISolverNetBindingsCaller, error) {
	contract, err := bindISolverNetBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISolverNetBindingsCaller{contract: contract}, nil
}

// NewISolverNetBindingsTransactor creates a new write-only instance of ISolverNetBindings, bound to a specific deployed contract.
func NewISolverNetBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*ISolverNetBindingsTransactor, error) {
	contract, err := bindISolverNetBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISolverNetBindingsTransactor{contract: contract}, nil
}

// NewISolverNetBindingsFilterer creates a new log filterer instance of ISolverNetBindings, bound to a specific deployed contract.
func NewISolverNetBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*ISolverNetBindingsFilterer, error) {
	contract, err := bindISolverNetBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISolverNetBindingsFilterer{contract: contract}, nil
}

// bindISolverNetBindings binds a generic wrapper to an already deployed contract.
func bindISolverNetBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISolverNetBindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISolverNetBindings *ISolverNetBindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISolverNetBindings.Contract.ISolverNetBindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISolverNetBindings *ISolverNetBindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISolverNetBindings.Contract.ISolverNetBindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISolverNetBindings *ISolverNetBindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISolverNetBindings.Contract.ISolverNetBindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISolverNetBindings *ISolverNetBindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISolverNetBindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISolverNetBindings *ISolverNetBindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISolverNetBindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISolverNetBindings *ISolverNetBindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISolverNetBindings.Contract.contract.Transact(opts, method, params...)
}

// Call is a free data retrieval call binding the contract method 0xf8b6c506.
//
// Solidity: function call((uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[]) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCaller) Call(opts *bind.CallOpts, arg0 ISolverNetCall) error {
	var out []interface{}
	err := _ISolverNetBindings.contract.Call(opts, &out, "call", arg0)

	if err != nil {
		return err
	}

	return err

}

// Call is a free data retrieval call binding the contract method 0xf8b6c506.
//
// Solidity: function call((uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[]) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsSession) Call(arg0 ISolverNetCall) error {
	return _ISolverNetBindings.Contract.Call(&_ISolverNetBindings.CallOpts, arg0)
}

// Call is a free data retrieval call binding the contract method 0xf8b6c506.
//
// Solidity: function call((uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[]) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCallerSession) Call(arg0 ISolverNetCall) error {
	return _ISolverNetBindings.Contract.Call(&_ISolverNetBindings.CallOpts, arg0)
}

// Deposit is a free data retrieval call binding the contract method 0x32defc2c.
//
// Solidity: function deposit((bytes32,uint256) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCaller) Deposit(opts *bind.CallOpts, arg0 ISolverNetDeposit) error {
	var out []interface{}
	err := _ISolverNetBindings.contract.Call(opts, &out, "deposit", arg0)

	if err != nil {
		return err
	}

	return err

}

// Deposit is a free data retrieval call binding the contract method 0x32defc2c.
//
// Solidity: function deposit((bytes32,uint256) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsSession) Deposit(arg0 ISolverNetDeposit) error {
	return _ISolverNetBindings.Contract.Deposit(&_ISolverNetBindings.CallOpts, arg0)
}

// Deposit is a free data retrieval call binding the contract method 0x32defc2c.
//
// Solidity: function deposit((bytes32,uint256) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCallerSession) Deposit(arg0 ISolverNetDeposit) error {
	return _ISolverNetBindings.Contract.Deposit(&_ISolverNetBindings.CallOpts, arg0)
}

// FillOriginData is a free data retrieval call binding the contract method 0x81f90fa5.
//
// Solidity: function fillOriginData((uint64,uint40,(uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[])) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCaller) FillOriginData(opts *bind.CallOpts, arg0 ISolverNetFillOriginData) error {
	var out []interface{}
	err := _ISolverNetBindings.contract.Call(opts, &out, "fillOriginData", arg0)

	if err != nil {
		return err
	}

	return err

}

// FillOriginData is a free data retrieval call binding the contract method 0x81f90fa5.
//
// Solidity: function fillOriginData((uint64,uint40,(uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[])) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsSession) FillOriginData(arg0 ISolverNetFillOriginData) error {
	return _ISolverNetBindings.Contract.FillOriginData(&_ISolverNetBindings.CallOpts, arg0)
}

// FillOriginData is a free data retrieval call binding the contract method 0x81f90fa5.
//
// Solidity: function fillOriginData((uint64,uint40,(uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[])) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCallerSession) FillOriginData(arg0 ISolverNetFillOriginData) error {
	return _ISolverNetBindings.Contract.FillOriginData(&_ISolverNetBindings.CallOpts, arg0)
}

// OrderData is a free data retrieval call binding the contract method 0x0f5c779a.
//
// Solidity: function orderData(((uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[]),(bytes32,uint256)[]) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCaller) OrderData(opts *bind.CallOpts, arg0 ISolverNetOrderData) error {
	var out []interface{}
	err := _ISolverNetBindings.contract.Call(opts, &out, "orderData", arg0)

	if err != nil {
		return err
	}

	return err

}

// OrderData is a free data retrieval call binding the contract method 0x0f5c779a.
//
// Solidity: function orderData(((uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[]),(bytes32,uint256)[]) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsSession) OrderData(arg0 ISolverNetOrderData) error {
	return _ISolverNetBindings.Contract.OrderData(&_ISolverNetBindings.CallOpts, arg0)
}

// OrderData is a free data retrieval call binding the contract method 0x0f5c779a.
//
// Solidity: function orderData(((uint64,bytes32,uint256,bytes,(bytes32,bytes32,uint256)[]),(bytes32,uint256)[]) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCallerSession) OrderData(arg0 ISolverNetOrderData) error {
	return _ISolverNetBindings.Contract.OrderData(&_ISolverNetBindings.CallOpts, arg0)
}

// TokenExpense is a free data retrieval call binding the contract method 0xde5e1c62.
//
// Solidity: function tokenExpense((bytes32,bytes32,uint256) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCaller) TokenExpense(opts *bind.CallOpts, arg0 ISolverNetTokenExpense) error {
	var out []interface{}
	err := _ISolverNetBindings.contract.Call(opts, &out, "tokenExpense", arg0)

	if err != nil {
		return err
	}

	return err

}

// TokenExpense is a free data retrieval call binding the contract method 0xde5e1c62.
//
// Solidity: function tokenExpense((bytes32,bytes32,uint256) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsSession) TokenExpense(arg0 ISolverNetTokenExpense) error {
	return _ISolverNetBindings.Contract.TokenExpense(&_ISolverNetBindings.CallOpts, arg0)
}

// TokenExpense is a free data retrieval call binding the contract method 0xde5e1c62.
//
// Solidity: function tokenExpense((bytes32,bytes32,uint256) ) view returns()
func (_ISolverNetBindings *ISolverNetBindingsCallerSession) TokenExpense(arg0 ISolverNetTokenExpense) error {
	return _ISolverNetBindings.Contract.TokenExpense(&_ISolverNetBindings.CallOpts, arg0)
}
