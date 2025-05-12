// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rebalance

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

// WSTETHMetaData contains all meta data concerning the WSTETH contract.
var WSTETHMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"_stETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wstETHAmount\",\"type\":\"uint256\"}],\"name\":\"getStETHByWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stETHAmount\",\"type\":\"uint256\"}],\"name\":\"getWstETHByStETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stEthPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerStEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wstETHAmount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stETHAmount\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WSTETHABI is the input ABI used to generate the binding from.
// Deprecated: Use WSTETHMetaData.ABI instead.
var WSTETHABI = WSTETHMetaData.ABI

// WSTETH is an auto generated Go binding around an Ethereum contract.
type WSTETH struct {
	WSTETHCaller     // Read-only binding to the contract
	WSTETHTransactor // Write-only binding to the contract
	WSTETHFilterer   // Log filterer for contract events
}

// WSTETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type WSTETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WSTETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WSTETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WSTETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WSTETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WSTETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WSTETHSession struct {
	Contract     *WSTETH           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WSTETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WSTETHCallerSession struct {
	Contract *WSTETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WSTETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WSTETHTransactorSession struct {
	Contract     *WSTETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WSTETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type WSTETHRaw struct {
	Contract *WSTETH // Generic contract binding to access the raw methods on
}

// WSTETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WSTETHCallerRaw struct {
	Contract *WSTETHCaller // Generic read-only contract binding to access the raw methods on
}

// WSTETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WSTETHTransactorRaw struct {
	Contract *WSTETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWSTETH creates a new instance of WSTETH, bound to a specific deployed contract.
func NewWSTETH(address common.Address, backend bind.ContractBackend) (*WSTETH, error) {
	contract, err := bindWSTETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WSTETH{WSTETHCaller: WSTETHCaller{contract: contract}, WSTETHTransactor: WSTETHTransactor{contract: contract}, WSTETHFilterer: WSTETHFilterer{contract: contract}}, nil
}

// NewWSTETHCaller creates a new read-only instance of WSTETH, bound to a specific deployed contract.
func NewWSTETHCaller(address common.Address, caller bind.ContractCaller) (*WSTETHCaller, error) {
	contract, err := bindWSTETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WSTETHCaller{contract: contract}, nil
}

// NewWSTETHTransactor creates a new write-only instance of WSTETH, bound to a specific deployed contract.
func NewWSTETHTransactor(address common.Address, transactor bind.ContractTransactor) (*WSTETHTransactor, error) {
	contract, err := bindWSTETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WSTETHTransactor{contract: contract}, nil
}

// NewWSTETHFilterer creates a new log filterer instance of WSTETH, bound to a specific deployed contract.
func NewWSTETHFilterer(address common.Address, filterer bind.ContractFilterer) (*WSTETHFilterer, error) {
	contract, err := bindWSTETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WSTETHFilterer{contract: contract}, nil
}

// bindWSTETH binds a generic wrapper to an already deployed contract.
func bindWSTETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WSTETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WSTETH *WSTETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WSTETH.Contract.WSTETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WSTETH *WSTETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WSTETH.Contract.WSTETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WSTETH *WSTETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WSTETH.Contract.WSTETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WSTETH *WSTETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WSTETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WSTETH *WSTETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WSTETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WSTETH *WSTETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WSTETH.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WSTETH *WSTETHCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WSTETH *WSTETHSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WSTETH.Contract.DOMAINSEPARATOR(&_WSTETH.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WSTETH *WSTETHCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WSTETH.Contract.DOMAINSEPARATOR(&_WSTETH.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WSTETH *WSTETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WSTETH *WSTETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WSTETH.Contract.Allowance(&_WSTETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WSTETH *WSTETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WSTETH.Contract.Allowance(&_WSTETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WSTETH *WSTETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WSTETH *WSTETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WSTETH.Contract.BalanceOf(&_WSTETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WSTETH *WSTETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WSTETH.Contract.BalanceOf(&_WSTETH.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WSTETH *WSTETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WSTETH *WSTETHSession) Decimals() (uint8, error) {
	return _WSTETH.Contract.Decimals(&_WSTETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WSTETH *WSTETHCallerSession) Decimals() (uint8, error) {
	return _WSTETH.Contract.Decimals(&_WSTETH.CallOpts)
}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_WSTETH *WSTETHCaller) GetStETHByWstETH(opts *bind.CallOpts, _wstETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "getStETHByWstETH", _wstETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_WSTETH *WSTETHSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*big.Int, error) {
	return _WSTETH.Contract.GetStETHByWstETH(&_WSTETH.CallOpts, _wstETHAmount)
}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_WSTETH *WSTETHCallerSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*big.Int, error) {
	return _WSTETH.Contract.GetStETHByWstETH(&_WSTETH.CallOpts, _wstETHAmount)
}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_WSTETH *WSTETHCaller) GetWstETHByStETH(opts *bind.CallOpts, _stETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "getWstETHByStETH", _stETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_WSTETH *WSTETHSession) GetWstETHByStETH(_stETHAmount *big.Int) (*big.Int, error) {
	return _WSTETH.Contract.GetWstETHByStETH(&_WSTETH.CallOpts, _stETHAmount)
}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_WSTETH *WSTETHCallerSession) GetWstETHByStETH(_stETHAmount *big.Int) (*big.Int, error) {
	return _WSTETH.Contract.GetWstETHByStETH(&_WSTETH.CallOpts, _stETHAmount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WSTETH *WSTETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WSTETH *WSTETHSession) Name() (string, error) {
	return _WSTETH.Contract.Name(&_WSTETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WSTETH *WSTETHCallerSession) Name() (string, error) {
	return _WSTETH.Contract.Name(&_WSTETH.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WSTETH *WSTETHCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WSTETH *WSTETHSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WSTETH.Contract.Nonces(&_WSTETH.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WSTETH *WSTETHCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WSTETH.Contract.Nonces(&_WSTETH.CallOpts, owner)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WSTETH *WSTETHCaller) StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "stETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WSTETH *WSTETHSession) StETH() (common.Address, error) {
	return _WSTETH.Contract.StETH(&_WSTETH.CallOpts)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WSTETH *WSTETHCallerSession) StETH() (common.Address, error) {
	return _WSTETH.Contract.StETH(&_WSTETH.CallOpts)
}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_WSTETH *WSTETHCaller) StEthPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "stEthPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_WSTETH *WSTETHSession) StEthPerToken() (*big.Int, error) {
	return _WSTETH.Contract.StEthPerToken(&_WSTETH.CallOpts)
}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_WSTETH *WSTETHCallerSession) StEthPerToken() (*big.Int, error) {
	return _WSTETH.Contract.StEthPerToken(&_WSTETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WSTETH *WSTETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WSTETH *WSTETHSession) Symbol() (string, error) {
	return _WSTETH.Contract.Symbol(&_WSTETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WSTETH *WSTETHCallerSession) Symbol() (string, error) {
	return _WSTETH.Contract.Symbol(&_WSTETH.CallOpts)
}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_WSTETH *WSTETHCaller) TokensPerStEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "tokensPerStEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_WSTETH *WSTETHSession) TokensPerStEth() (*big.Int, error) {
	return _WSTETH.Contract.TokensPerStEth(&_WSTETH.CallOpts)
}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_WSTETH *WSTETHCallerSession) TokensPerStEth() (*big.Int, error) {
	return _WSTETH.Contract.TokensPerStEth(&_WSTETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WSTETH *WSTETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WSTETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WSTETH *WSTETHSession) TotalSupply() (*big.Int, error) {
	return _WSTETH.Contract.TotalSupply(&_WSTETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WSTETH *WSTETHCallerSession) TotalSupply() (*big.Int, error) {
	return _WSTETH.Contract.TotalSupply(&_WSTETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WSTETH *WSTETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WSTETH *WSTETHSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.Approve(&_WSTETH.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WSTETH *WSTETHTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.Approve(&_WSTETH.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WSTETH *WSTETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WSTETH.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WSTETH *WSTETHSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.DecreaseAllowance(&_WSTETH.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WSTETH *WSTETHTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.DecreaseAllowance(&_WSTETH.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WSTETH *WSTETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WSTETH.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WSTETH *WSTETHSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.IncreaseAllowance(&_WSTETH.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WSTETH *WSTETHTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.IncreaseAllowance(&_WSTETH.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WSTETH *WSTETHTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WSTETH.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WSTETH *WSTETHSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WSTETH.Contract.Permit(&_WSTETH.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WSTETH *WSTETHTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WSTETH.Contract.Permit(&_WSTETH.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WSTETH *WSTETHTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WSTETH *WSTETHSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.Transfer(&_WSTETH.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WSTETH *WSTETHTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.Transfer(&_WSTETH.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WSTETH *WSTETHTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WSTETH *WSTETHSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.TransferFrom(&_WSTETH.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WSTETH *WSTETHTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.TransferFrom(&_WSTETH.TransactOpts, sender, recipient, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_WSTETH *WSTETHTransactor) Unwrap(opts *bind.TransactOpts, _wstETHAmount *big.Int) (*types.Transaction, error) {
	return _WSTETH.contract.Transact(opts, "unwrap", _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_WSTETH *WSTETHSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.Unwrap(&_WSTETH.TransactOpts, _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_WSTETH *WSTETHTransactorSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.Unwrap(&_WSTETH.TransactOpts, _wstETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_WSTETH *WSTETHTransactor) Wrap(opts *bind.TransactOpts, _stETHAmount *big.Int) (*types.Transaction, error) {
	return _WSTETH.contract.Transact(opts, "wrap", _stETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_WSTETH *WSTETHSession) Wrap(_stETHAmount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.Wrap(&_WSTETH.TransactOpts, _stETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_WSTETH *WSTETHTransactorSession) Wrap(_stETHAmount *big.Int) (*types.Transaction, error) {
	return _WSTETH.Contract.Wrap(&_WSTETH.TransactOpts, _stETHAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WSTETH *WSTETHTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WSTETH.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WSTETH *WSTETHSession) Receive() (*types.Transaction, error) {
	return _WSTETH.Contract.Receive(&_WSTETH.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WSTETH *WSTETHTransactorSession) Receive() (*types.Transaction, error) {
	return _WSTETH.Contract.Receive(&_WSTETH.TransactOpts)
}

// WSTETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WSTETH contract.
type WSTETHApprovalIterator struct {
	Event *WSTETHApproval // Event containing the contract specifics and raw log

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
func (it *WSTETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WSTETHApproval)
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
		it.Event = new(WSTETHApproval)
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
func (it *WSTETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WSTETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WSTETHApproval represents a Approval event raised by the WSTETH contract.
type WSTETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WSTETH *WSTETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WSTETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WSTETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WSTETHApprovalIterator{contract: _WSTETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WSTETH *WSTETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WSTETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WSTETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WSTETHApproval)
				if err := _WSTETH.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WSTETH *WSTETHFilterer) ParseApproval(log types.Log) (*WSTETHApproval, error) {
	event := new(WSTETHApproval)
	if err := _WSTETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WSTETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WSTETH contract.
type WSTETHTransferIterator struct {
	Event *WSTETHTransfer // Event containing the contract specifics and raw log

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
func (it *WSTETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WSTETHTransfer)
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
		it.Event = new(WSTETHTransfer)
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
func (it *WSTETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WSTETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WSTETHTransfer represents a Transfer event raised by the WSTETH contract.
type WSTETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WSTETH *WSTETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WSTETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WSTETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WSTETHTransferIterator{contract: _WSTETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WSTETH *WSTETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WSTETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WSTETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WSTETHTransfer)
				if err := _WSTETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WSTETH *WSTETHFilterer) ParseTransfer(log types.Log) (*WSTETHTransfer, error) {
	event := new(WSTETHTransfer)
	if err := _WSTETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
