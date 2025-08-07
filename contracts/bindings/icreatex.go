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

// ICreateXValues is an auto generated low-level Go binding around an user-defined struct.
type ICreateXValues struct {
	ConstructorAmount *big.Int
	InitCallAmount    *big.Int
}

// ICreateXMetaData contains all meta data concerning the ICreateX contract.
var ICreateXMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"computeCreate2Address\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"computedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeCreate2Address\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"computedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"computeCreate3Address\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"computedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"computeCreate3Address\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"computedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeCreateAddress\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"computedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeCreateAddress\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"computedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployCreate\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate2\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate2\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate2AndInit\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"refundAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate2AndInit\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate2AndInit\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"refundAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate2AndInit\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate2Clone\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate2Clone\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate3\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate3\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate3AndInit\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate3AndInit\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate3AndInit\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"refundAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreate3AndInit\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"refundAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreateAndInit\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreateAndInit\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"values\",\"type\":\"tuple\",\"internalType\":\"structICreateX.Values\",\"components\":[{\"name\":\"constructorAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCallAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"refundAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployCreateClone\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ContractCreation\",\"inputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractCreation\",\"inputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Create3ProxyContractCreation\",\"inputs\":[{\"name\":\"newContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FailedContractCreation\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedContractInitialisation\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"FailedEtherTransfer\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidNonceValue\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSalt\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ICreateXABI is the input ABI used to generate the binding from.
// Deprecated: Use ICreateXMetaData.ABI instead.
var ICreateXABI = ICreateXMetaData.ABI

// ICreateX is an auto generated Go binding around an Ethereum contract.
type ICreateX struct {
	ICreateXCaller     // Read-only binding to the contract
	ICreateXTransactor // Write-only binding to the contract
	ICreateXFilterer   // Log filterer for contract events
}

// ICreateXCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICreateXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreateXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICreateXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreateXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICreateXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICreateXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICreateXSession struct {
	Contract     *ICreateX         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICreateXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICreateXCallerSession struct {
	Contract *ICreateXCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ICreateXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICreateXTransactorSession struct {
	Contract     *ICreateXTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICreateXRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICreateXRaw struct {
	Contract *ICreateX // Generic contract binding to access the raw methods on
}

// ICreateXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICreateXCallerRaw struct {
	Contract *ICreateXCaller // Generic read-only contract binding to access the raw methods on
}

// ICreateXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICreateXTransactorRaw struct {
	Contract *ICreateXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICreateX creates a new instance of ICreateX, bound to a specific deployed contract.
func NewICreateX(address common.Address, backend bind.ContractBackend) (*ICreateX, error) {
	contract, err := bindICreateX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICreateX{ICreateXCaller: ICreateXCaller{contract: contract}, ICreateXTransactor: ICreateXTransactor{contract: contract}, ICreateXFilterer: ICreateXFilterer{contract: contract}}, nil
}

// NewICreateXCaller creates a new read-only instance of ICreateX, bound to a specific deployed contract.
func NewICreateXCaller(address common.Address, caller bind.ContractCaller) (*ICreateXCaller, error) {
	contract, err := bindICreateX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICreateXCaller{contract: contract}, nil
}

// NewICreateXTransactor creates a new write-only instance of ICreateX, bound to a specific deployed contract.
func NewICreateXTransactor(address common.Address, transactor bind.ContractTransactor) (*ICreateXTransactor, error) {
	contract, err := bindICreateX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICreateXTransactor{contract: contract}, nil
}

// NewICreateXFilterer creates a new log filterer instance of ICreateX, bound to a specific deployed contract.
func NewICreateXFilterer(address common.Address, filterer bind.ContractFilterer) (*ICreateXFilterer, error) {
	contract, err := bindICreateX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICreateXFilterer{contract: contract}, nil
}

// bindICreateX binds a generic wrapper to an already deployed contract.
func bindICreateX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICreateXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreateX *ICreateXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICreateX.Contract.ICreateXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreateX *ICreateXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreateX.Contract.ICreateXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreateX *ICreateXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreateX.Contract.ICreateXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICreateX *ICreateXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICreateX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICreateX *ICreateXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICreateX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICreateX *ICreateXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICreateX.Contract.contract.Transact(opts, method, params...)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) view returns(address computedAddress)
func (_ICreateX *ICreateXCaller) ComputeCreate2Address(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ICreateX.contract.Call(opts, &out, "computeCreate2Address", salt, initCodeHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) view returns(address computedAddress)
func (_ICreateX *ICreateXSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreate2Address(&_ICreateX.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) view returns(address computedAddress)
func (_ICreateX *ICreateXCallerSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreate2Address(&_ICreateX.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address computedAddress)
func (_ICreateX *ICreateXCaller) ComputeCreate2Address0(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	var out []interface{}
	err := _ICreateX.contract.Call(opts, &out, "computeCreate2Address0", salt, initCodeHash, deployer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address computedAddress)
func (_ICreateX *ICreateXSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreate2Address0(&_ICreateX.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address computedAddress)
func (_ICreateX *ICreateXCallerSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreate2Address0(&_ICreateX.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreate3Address is a free data retrieval call binding the contract method 0x42d654fc.
//
// Solidity: function computeCreate3Address(bytes32 salt, address deployer) pure returns(address computedAddress)
func (_ICreateX *ICreateXCaller) ComputeCreate3Address(opts *bind.CallOpts, salt [32]byte, deployer common.Address) (common.Address, error) {
	var out []interface{}
	err := _ICreateX.contract.Call(opts, &out, "computeCreate3Address", salt, deployer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate3Address is a free data retrieval call binding the contract method 0x42d654fc.
//
// Solidity: function computeCreate3Address(bytes32 salt, address deployer) pure returns(address computedAddress)
func (_ICreateX *ICreateXSession) ComputeCreate3Address(salt [32]byte, deployer common.Address) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreate3Address(&_ICreateX.CallOpts, salt, deployer)
}

// ComputeCreate3Address is a free data retrieval call binding the contract method 0x42d654fc.
//
// Solidity: function computeCreate3Address(bytes32 salt, address deployer) pure returns(address computedAddress)
func (_ICreateX *ICreateXCallerSession) ComputeCreate3Address(salt [32]byte, deployer common.Address) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreate3Address(&_ICreateX.CallOpts, salt, deployer)
}

// ComputeCreate3Address0 is a free data retrieval call binding the contract method 0x6cec2536.
//
// Solidity: function computeCreate3Address(bytes32 salt) view returns(address computedAddress)
func (_ICreateX *ICreateXCaller) ComputeCreate3Address0(opts *bind.CallOpts, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ICreateX.contract.Call(opts, &out, "computeCreate3Address0", salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate3Address0 is a free data retrieval call binding the contract method 0x6cec2536.
//
// Solidity: function computeCreate3Address(bytes32 salt) view returns(address computedAddress)
func (_ICreateX *ICreateXSession) ComputeCreate3Address0(salt [32]byte) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreate3Address0(&_ICreateX.CallOpts, salt)
}

// ComputeCreate3Address0 is a free data retrieval call binding the contract method 0x6cec2536.
//
// Solidity: function computeCreate3Address(bytes32 salt) view returns(address computedAddress)
func (_ICreateX *ICreateXCallerSession) ComputeCreate3Address0(salt [32]byte) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreate3Address0(&_ICreateX.CallOpts, salt)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x28ddd046.
//
// Solidity: function computeCreateAddress(uint256 nonce) view returns(address computedAddress)
func (_ICreateX *ICreateXCaller) ComputeCreateAddress(opts *bind.CallOpts, nonce *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICreateX.contract.Call(opts, &out, "computeCreateAddress", nonce)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x28ddd046.
//
// Solidity: function computeCreateAddress(uint256 nonce) view returns(address computedAddress)
func (_ICreateX *ICreateXSession) ComputeCreateAddress(nonce *big.Int) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreateAddress(&_ICreateX.CallOpts, nonce)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x28ddd046.
//
// Solidity: function computeCreateAddress(uint256 nonce) view returns(address computedAddress)
func (_ICreateX *ICreateXCallerSession) ComputeCreateAddress(nonce *big.Int) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreateAddress(&_ICreateX.CallOpts, nonce)
}

// ComputeCreateAddress0 is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) view returns(address computedAddress)
func (_ICreateX *ICreateXCaller) ComputeCreateAddress0(opts *bind.CallOpts, deployer common.Address, nonce *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICreateX.contract.Call(opts, &out, "computeCreateAddress0", deployer, nonce)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreateAddress0 is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) view returns(address computedAddress)
func (_ICreateX *ICreateXSession) ComputeCreateAddress0(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreateAddress0(&_ICreateX.CallOpts, deployer, nonce)
}

// ComputeCreateAddress0 is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) view returns(address computedAddress)
func (_ICreateX *ICreateXCallerSession) ComputeCreateAddress0(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _ICreateX.Contract.ComputeCreateAddress0(&_ICreateX.CallOpts, deployer, nonce)
}

// DeployCreate is a paid mutator transaction binding the contract method 0x27fe1822.
//
// Solidity: function deployCreate(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate", initCode)
}

// DeployCreate is a paid mutator transaction binding the contract method 0x27fe1822.
//
// Solidity: function deployCreate(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate(initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate(&_ICreateX.TransactOpts, initCode)
}

// DeployCreate is a paid mutator transaction binding the contract method 0x27fe1822.
//
// Solidity: function deployCreate(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate(initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate(&_ICreateX.TransactOpts, initCode)
}

// DeployCreate2 is a paid mutator transaction binding the contract method 0x26307668.
//
// Solidity: function deployCreate2(bytes32 salt, bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate2(opts *bind.TransactOpts, salt [32]byte, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate2", salt, initCode)
}

// DeployCreate2 is a paid mutator transaction binding the contract method 0x26307668.
//
// Solidity: function deployCreate2(bytes32 salt, bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate2(salt [32]byte, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2(&_ICreateX.TransactOpts, salt, initCode)
}

// DeployCreate2 is a paid mutator transaction binding the contract method 0x26307668.
//
// Solidity: function deployCreate2(bytes32 salt, bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate2(salt [32]byte, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2(&_ICreateX.TransactOpts, salt, initCode)
}

// DeployCreate20 is a paid mutator transaction binding the contract method 0x26a32fc7.
//
// Solidity: function deployCreate2(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate20(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate20", initCode)
}

// DeployCreate20 is a paid mutator transaction binding the contract method 0x26a32fc7.
//
// Solidity: function deployCreate2(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate20(initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate20(&_ICreateX.TransactOpts, initCode)
}

// DeployCreate20 is a paid mutator transaction binding the contract method 0x26a32fc7.
//
// Solidity: function deployCreate2(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate20(initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate20(&_ICreateX.TransactOpts, initCode)
}

// DeployCreate2AndInit is a paid mutator transaction binding the contract method 0xa7db93f2.
//
// Solidity: function deployCreate2AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate2AndInit(opts *bind.TransactOpts, salt [32]byte, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate2AndInit", salt, initCode, data, values, refundAddress)
}

// DeployCreate2AndInit is a paid mutator transaction binding the contract method 0xa7db93f2.
//
// Solidity: function deployCreate2AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate2AndInit(salt [32]byte, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2AndInit(&_ICreateX.TransactOpts, salt, initCode, data, values, refundAddress)
}

// DeployCreate2AndInit is a paid mutator transaction binding the contract method 0xa7db93f2.
//
// Solidity: function deployCreate2AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate2AndInit(salt [32]byte, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2AndInit(&_ICreateX.TransactOpts, salt, initCode, data, values, refundAddress)
}

// DeployCreate2AndInit0 is a paid mutator transaction binding the contract method 0xc3fe107b.
//
// Solidity: function deployCreate2AndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate2AndInit0(opts *bind.TransactOpts, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate2AndInit0", initCode, data, values)
}

// DeployCreate2AndInit0 is a paid mutator transaction binding the contract method 0xc3fe107b.
//
// Solidity: function deployCreate2AndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate2AndInit0(initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2AndInit0(&_ICreateX.TransactOpts, initCode, data, values)
}

// DeployCreate2AndInit0 is a paid mutator transaction binding the contract method 0xc3fe107b.
//
// Solidity: function deployCreate2AndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate2AndInit0(initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2AndInit0(&_ICreateX.TransactOpts, initCode, data, values)
}

// DeployCreate2AndInit1 is a paid mutator transaction binding the contract method 0xe437252a.
//
// Solidity: function deployCreate2AndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate2AndInit1(opts *bind.TransactOpts, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate2AndInit1", initCode, data, values, refundAddress)
}

// DeployCreate2AndInit1 is a paid mutator transaction binding the contract method 0xe437252a.
//
// Solidity: function deployCreate2AndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate2AndInit1(initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2AndInit1(&_ICreateX.TransactOpts, initCode, data, values, refundAddress)
}

// DeployCreate2AndInit1 is a paid mutator transaction binding the contract method 0xe437252a.
//
// Solidity: function deployCreate2AndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate2AndInit1(initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2AndInit1(&_ICreateX.TransactOpts, initCode, data, values, refundAddress)
}

// DeployCreate2AndInit2 is a paid mutator transaction binding the contract method 0xe96deee4.
//
// Solidity: function deployCreate2AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate2AndInit2(opts *bind.TransactOpts, salt [32]byte, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate2AndInit2", salt, initCode, data, values)
}

// DeployCreate2AndInit2 is a paid mutator transaction binding the contract method 0xe96deee4.
//
// Solidity: function deployCreate2AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate2AndInit2(salt [32]byte, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2AndInit2(&_ICreateX.TransactOpts, salt, initCode, data, values)
}

// DeployCreate2AndInit2 is a paid mutator transaction binding the contract method 0xe96deee4.
//
// Solidity: function deployCreate2AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate2AndInit2(salt [32]byte, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2AndInit2(&_ICreateX.TransactOpts, salt, initCode, data, values)
}

// DeployCreate2Clone is a paid mutator transaction binding the contract method 0x2852527a.
//
// Solidity: function deployCreate2Clone(bytes32 salt, address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXTransactor) DeployCreate2Clone(opts *bind.TransactOpts, salt [32]byte, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate2Clone", salt, implementation, data)
}

// DeployCreate2Clone is a paid mutator transaction binding the contract method 0x2852527a.
//
// Solidity: function deployCreate2Clone(bytes32 salt, address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXSession) DeployCreate2Clone(salt [32]byte, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2Clone(&_ICreateX.TransactOpts, salt, implementation, data)
}

// DeployCreate2Clone is a paid mutator transaction binding the contract method 0x2852527a.
//
// Solidity: function deployCreate2Clone(bytes32 salt, address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXTransactorSession) DeployCreate2Clone(salt [32]byte, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2Clone(&_ICreateX.TransactOpts, salt, implementation, data)
}

// DeployCreate2Clone0 is a paid mutator transaction binding the contract method 0x81503da1.
//
// Solidity: function deployCreate2Clone(address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXTransactor) DeployCreate2Clone0(opts *bind.TransactOpts, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate2Clone0", implementation, data)
}

// DeployCreate2Clone0 is a paid mutator transaction binding the contract method 0x81503da1.
//
// Solidity: function deployCreate2Clone(address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXSession) DeployCreate2Clone0(implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2Clone0(&_ICreateX.TransactOpts, implementation, data)
}

// DeployCreate2Clone0 is a paid mutator transaction binding the contract method 0x81503da1.
//
// Solidity: function deployCreate2Clone(address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXTransactorSession) DeployCreate2Clone0(implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate2Clone0(&_ICreateX.TransactOpts, implementation, data)
}

// DeployCreate3 is a paid mutator transaction binding the contract method 0x7f565360.
//
// Solidity: function deployCreate3(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate3(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate3", initCode)
}

// DeployCreate3 is a paid mutator transaction binding the contract method 0x7f565360.
//
// Solidity: function deployCreate3(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate3(initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3(&_ICreateX.TransactOpts, initCode)
}

// DeployCreate3 is a paid mutator transaction binding the contract method 0x7f565360.
//
// Solidity: function deployCreate3(bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate3(initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3(&_ICreateX.TransactOpts, initCode)
}

// DeployCreate30 is a paid mutator transaction binding the contract method 0x9c36a286.
//
// Solidity: function deployCreate3(bytes32 salt, bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate30(opts *bind.TransactOpts, salt [32]byte, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate30", salt, initCode)
}

// DeployCreate30 is a paid mutator transaction binding the contract method 0x9c36a286.
//
// Solidity: function deployCreate3(bytes32 salt, bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate30(salt [32]byte, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate30(&_ICreateX.TransactOpts, salt, initCode)
}

// DeployCreate30 is a paid mutator transaction binding the contract method 0x9c36a286.
//
// Solidity: function deployCreate3(bytes32 salt, bytes initCode) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate30(salt [32]byte, initCode []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate30(&_ICreateX.TransactOpts, salt, initCode)
}

// DeployCreate3AndInit is a paid mutator transaction binding the contract method 0x00d84acb.
//
// Solidity: function deployCreate3AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate3AndInit(opts *bind.TransactOpts, salt [32]byte, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate3AndInit", salt, initCode, data, values)
}

// DeployCreate3AndInit is a paid mutator transaction binding the contract method 0x00d84acb.
//
// Solidity: function deployCreate3AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate3AndInit(salt [32]byte, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3AndInit(&_ICreateX.TransactOpts, salt, initCode, data, values)
}

// DeployCreate3AndInit is a paid mutator transaction binding the contract method 0x00d84acb.
//
// Solidity: function deployCreate3AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate3AndInit(salt [32]byte, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3AndInit(&_ICreateX.TransactOpts, salt, initCode, data, values)
}

// DeployCreate3AndInit0 is a paid mutator transaction binding the contract method 0x2f990e3f.
//
// Solidity: function deployCreate3AndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate3AndInit0(opts *bind.TransactOpts, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate3AndInit0", initCode, data, values)
}

// DeployCreate3AndInit0 is a paid mutator transaction binding the contract method 0x2f990e3f.
//
// Solidity: function deployCreate3AndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate3AndInit0(initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3AndInit0(&_ICreateX.TransactOpts, initCode, data, values)
}

// DeployCreate3AndInit0 is a paid mutator transaction binding the contract method 0x2f990e3f.
//
// Solidity: function deployCreate3AndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate3AndInit0(initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3AndInit0(&_ICreateX.TransactOpts, initCode, data, values)
}

// DeployCreate3AndInit1 is a paid mutator transaction binding the contract method 0xddda0acb.
//
// Solidity: function deployCreate3AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate3AndInit1(opts *bind.TransactOpts, salt [32]byte, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate3AndInit1", salt, initCode, data, values, refundAddress)
}

// DeployCreate3AndInit1 is a paid mutator transaction binding the contract method 0xddda0acb.
//
// Solidity: function deployCreate3AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate3AndInit1(salt [32]byte, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3AndInit1(&_ICreateX.TransactOpts, salt, initCode, data, values, refundAddress)
}

// DeployCreate3AndInit1 is a paid mutator transaction binding the contract method 0xddda0acb.
//
// Solidity: function deployCreate3AndInit(bytes32 salt, bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate3AndInit1(salt [32]byte, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3AndInit1(&_ICreateX.TransactOpts, salt, initCode, data, values, refundAddress)
}

// DeployCreate3AndInit2 is a paid mutator transaction binding the contract method 0xf5745aba.
//
// Solidity: function deployCreate3AndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreate3AndInit2(opts *bind.TransactOpts, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreate3AndInit2", initCode, data, values, refundAddress)
}

// DeployCreate3AndInit2 is a paid mutator transaction binding the contract method 0xf5745aba.
//
// Solidity: function deployCreate3AndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreate3AndInit2(initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3AndInit2(&_ICreateX.TransactOpts, initCode, data, values, refundAddress)
}

// DeployCreate3AndInit2 is a paid mutator transaction binding the contract method 0xf5745aba.
//
// Solidity: function deployCreate3AndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreate3AndInit2(initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreate3AndInit2(&_ICreateX.TransactOpts, initCode, data, values, refundAddress)
}

// DeployCreateAndInit is a paid mutator transaction binding the contract method 0x31a7c8c8.
//
// Solidity: function deployCreateAndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreateAndInit(opts *bind.TransactOpts, initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreateAndInit", initCode, data, values)
}

// DeployCreateAndInit is a paid mutator transaction binding the contract method 0x31a7c8c8.
//
// Solidity: function deployCreateAndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreateAndInit(initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreateAndInit(&_ICreateX.TransactOpts, initCode, data, values)
}

// DeployCreateAndInit is a paid mutator transaction binding the contract method 0x31a7c8c8.
//
// Solidity: function deployCreateAndInit(bytes initCode, bytes data, (uint256,uint256) values) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreateAndInit(initCode []byte, data []byte, values ICreateXValues) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreateAndInit(&_ICreateX.TransactOpts, initCode, data, values)
}

// DeployCreateAndInit0 is a paid mutator transaction binding the contract method 0x98e81077.
//
// Solidity: function deployCreateAndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactor) DeployCreateAndInit0(opts *bind.TransactOpts, initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreateAndInit0", initCode, data, values, refundAddress)
}

// DeployCreateAndInit0 is a paid mutator transaction binding the contract method 0x98e81077.
//
// Solidity: function deployCreateAndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXSession) DeployCreateAndInit0(initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreateAndInit0(&_ICreateX.TransactOpts, initCode, data, values, refundAddress)
}

// DeployCreateAndInit0 is a paid mutator transaction binding the contract method 0x98e81077.
//
// Solidity: function deployCreateAndInit(bytes initCode, bytes data, (uint256,uint256) values, address refundAddress) payable returns(address newContract)
func (_ICreateX *ICreateXTransactorSession) DeployCreateAndInit0(initCode []byte, data []byte, values ICreateXValues, refundAddress common.Address) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreateAndInit0(&_ICreateX.TransactOpts, initCode, data, values, refundAddress)
}

// DeployCreateClone is a paid mutator transaction binding the contract method 0xf9664498.
//
// Solidity: function deployCreateClone(address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXTransactor) DeployCreateClone(opts *bind.TransactOpts, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.contract.Transact(opts, "deployCreateClone", implementation, data)
}

// DeployCreateClone is a paid mutator transaction binding the contract method 0xf9664498.
//
// Solidity: function deployCreateClone(address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXSession) DeployCreateClone(implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreateClone(&_ICreateX.TransactOpts, implementation, data)
}

// DeployCreateClone is a paid mutator transaction binding the contract method 0xf9664498.
//
// Solidity: function deployCreateClone(address implementation, bytes data) payable returns(address proxy)
func (_ICreateX *ICreateXTransactorSession) DeployCreateClone(implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ICreateX.Contract.DeployCreateClone(&_ICreateX.TransactOpts, implementation, data)
}

// ICreateXContractCreationIterator is returned from FilterContractCreation and is used to iterate over the raw logs and unpacked data for ContractCreation events raised by the ICreateX contract.
type ICreateXContractCreationIterator struct {
	Event *ICreateXContractCreation // Event containing the contract specifics and raw log

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
func (it *ICreateXContractCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreateXContractCreation)
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
		it.Event = new(ICreateXContractCreation)
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
func (it *ICreateXContractCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreateXContractCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreateXContractCreation represents a ContractCreation event raised by the ICreateX contract.
type ICreateXContractCreation struct {
	NewContract common.Address
	Salt        [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractCreation is a free log retrieval operation binding the contract event 0xb8fda7e00c6b06a2b54e58521bc5894fee35f1090e5a3bb6390bfe2b98b497f7.
//
// Solidity: event ContractCreation(address indexed newContract, bytes32 indexed salt)
func (_ICreateX *ICreateXFilterer) FilterContractCreation(opts *bind.FilterOpts, newContract []common.Address, salt [][32]byte) (*ICreateXContractCreationIterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}
	var saltRule []interface{}
	for _, saltItem := range salt {
		saltRule = append(saltRule, saltItem)
	}

	logs, sub, err := _ICreateX.contract.FilterLogs(opts, "ContractCreation", newContractRule, saltRule)
	if err != nil {
		return nil, err
	}
	return &ICreateXContractCreationIterator{contract: _ICreateX.contract, event: "ContractCreation", logs: logs, sub: sub}, nil
}

// WatchContractCreation is a free log subscription operation binding the contract event 0xb8fda7e00c6b06a2b54e58521bc5894fee35f1090e5a3bb6390bfe2b98b497f7.
//
// Solidity: event ContractCreation(address indexed newContract, bytes32 indexed salt)
func (_ICreateX *ICreateXFilterer) WatchContractCreation(opts *bind.WatchOpts, sink chan<- *ICreateXContractCreation, newContract []common.Address, salt [][32]byte) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}
	var saltRule []interface{}
	for _, saltItem := range salt {
		saltRule = append(saltRule, saltItem)
	}

	logs, sub, err := _ICreateX.contract.WatchLogs(opts, "ContractCreation", newContractRule, saltRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreateXContractCreation)
				if err := _ICreateX.contract.UnpackLog(event, "ContractCreation", log); err != nil {
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

// ParseContractCreation is a log parse operation binding the contract event 0xb8fda7e00c6b06a2b54e58521bc5894fee35f1090e5a3bb6390bfe2b98b497f7.
//
// Solidity: event ContractCreation(address indexed newContract, bytes32 indexed salt)
func (_ICreateX *ICreateXFilterer) ParseContractCreation(log types.Log) (*ICreateXContractCreation, error) {
	event := new(ICreateXContractCreation)
	if err := _ICreateX.contract.UnpackLog(event, "ContractCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreateXContractCreation0Iterator is returned from FilterContractCreation0 and is used to iterate over the raw logs and unpacked data for ContractCreation0 events raised by the ICreateX contract.
type ICreateXContractCreation0Iterator struct {
	Event *ICreateXContractCreation0 // Event containing the contract specifics and raw log

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
func (it *ICreateXContractCreation0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreateXContractCreation0)
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
		it.Event = new(ICreateXContractCreation0)
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
func (it *ICreateXContractCreation0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreateXContractCreation0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreateXContractCreation0 represents a ContractCreation0 event raised by the ICreateX contract.
type ICreateXContractCreation0 struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractCreation0 is a free log retrieval operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: event ContractCreation(address indexed newContract)
func (_ICreateX *ICreateXFilterer) FilterContractCreation0(opts *bind.FilterOpts, newContract []common.Address) (*ICreateXContractCreation0Iterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _ICreateX.contract.FilterLogs(opts, "ContractCreation0", newContractRule)
	if err != nil {
		return nil, err
	}
	return &ICreateXContractCreation0Iterator{contract: _ICreateX.contract, event: "ContractCreation0", logs: logs, sub: sub}, nil
}

// WatchContractCreation0 is a free log subscription operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: event ContractCreation(address indexed newContract)
func (_ICreateX *ICreateXFilterer) WatchContractCreation0(opts *bind.WatchOpts, sink chan<- *ICreateXContractCreation0, newContract []common.Address) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _ICreateX.contract.WatchLogs(opts, "ContractCreation0", newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreateXContractCreation0)
				if err := _ICreateX.contract.UnpackLog(event, "ContractCreation0", log); err != nil {
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

// ParseContractCreation0 is a log parse operation binding the contract event 0x4db17dd5e4732fb6da34a148104a592783ca119a1e7bb8829eba6cbadef0b511.
//
// Solidity: event ContractCreation(address indexed newContract)
func (_ICreateX *ICreateXFilterer) ParseContractCreation0(log types.Log) (*ICreateXContractCreation0, error) {
	event := new(ICreateXContractCreation0)
	if err := _ICreateX.contract.UnpackLog(event, "ContractCreation0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICreateXCreate3ProxyContractCreationIterator is returned from FilterCreate3ProxyContractCreation and is used to iterate over the raw logs and unpacked data for Create3ProxyContractCreation events raised by the ICreateX contract.
type ICreateXCreate3ProxyContractCreationIterator struct {
	Event *ICreateXCreate3ProxyContractCreation // Event containing the contract specifics and raw log

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
func (it *ICreateXCreate3ProxyContractCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICreateXCreate3ProxyContractCreation)
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
		it.Event = new(ICreateXCreate3ProxyContractCreation)
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
func (it *ICreateXCreate3ProxyContractCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICreateXCreate3ProxyContractCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICreateXCreate3ProxyContractCreation represents a Create3ProxyContractCreation event raised by the ICreateX contract.
type ICreateXCreate3ProxyContractCreation struct {
	NewContract common.Address
	Salt        [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreate3ProxyContractCreation is a free log retrieval operation binding the contract event 0x2feea65dd4e9f9cbd86b74b7734210c59a1b2981b5b137bd0ee3e208200c9067.
//
// Solidity: event Create3ProxyContractCreation(address indexed newContract, bytes32 indexed salt)
func (_ICreateX *ICreateXFilterer) FilterCreate3ProxyContractCreation(opts *bind.FilterOpts, newContract []common.Address, salt [][32]byte) (*ICreateXCreate3ProxyContractCreationIterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}
	var saltRule []interface{}
	for _, saltItem := range salt {
		saltRule = append(saltRule, saltItem)
	}

	logs, sub, err := _ICreateX.contract.FilterLogs(opts, "Create3ProxyContractCreation", newContractRule, saltRule)
	if err != nil {
		return nil, err
	}
	return &ICreateXCreate3ProxyContractCreationIterator{contract: _ICreateX.contract, event: "Create3ProxyContractCreation", logs: logs, sub: sub}, nil
}

// WatchCreate3ProxyContractCreation is a free log subscription operation binding the contract event 0x2feea65dd4e9f9cbd86b74b7734210c59a1b2981b5b137bd0ee3e208200c9067.
//
// Solidity: event Create3ProxyContractCreation(address indexed newContract, bytes32 indexed salt)
func (_ICreateX *ICreateXFilterer) WatchCreate3ProxyContractCreation(opts *bind.WatchOpts, sink chan<- *ICreateXCreate3ProxyContractCreation, newContract []common.Address, salt [][32]byte) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}
	var saltRule []interface{}
	for _, saltItem := range salt {
		saltRule = append(saltRule, saltItem)
	}

	logs, sub, err := _ICreateX.contract.WatchLogs(opts, "Create3ProxyContractCreation", newContractRule, saltRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICreateXCreate3ProxyContractCreation)
				if err := _ICreateX.contract.UnpackLog(event, "Create3ProxyContractCreation", log); err != nil {
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

// ParseCreate3ProxyContractCreation is a log parse operation binding the contract event 0x2feea65dd4e9f9cbd86b74b7734210c59a1b2981b5b137bd0ee3e208200c9067.
//
// Solidity: event Create3ProxyContractCreation(address indexed newContract, bytes32 indexed salt)
func (_ICreateX *ICreateXFilterer) ParseCreate3ProxyContractCreation(log types.Log) (*ICreateXCreate3ProxyContractCreation, error) {
	event := new(ICreateXCreate3ProxyContractCreation)
	if err := _ICreateX.contract.UnpackLog(event, "Create3ProxyContractCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
