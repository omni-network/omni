// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// IQuoterV2QuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IQuoterV2QuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// UniQuoterV2MetaData contains all meta data concerning the UniQuoterV2 contract.
var UniQuoterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniQuoterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniQuoterV2MetaData.ABI instead.
var UniQuoterV2ABI = UniQuoterV2MetaData.ABI

// UniQuoterV2 is an auto generated Go binding around an Ethereum contract.
type UniQuoterV2 struct {
	UniQuoterV2Caller     // Read-only binding to the contract
	UniQuoterV2Transactor // Write-only binding to the contract
	UniQuoterV2Filterer   // Log filterer for contract events
}

// UniQuoterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniQuoterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniQuoterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniQuoterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniQuoterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniQuoterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniQuoterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniQuoterV2Session struct {
	Contract     *UniQuoterV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniQuoterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniQuoterV2CallerSession struct {
	Contract *UniQuoterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UniQuoterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniQuoterV2TransactorSession struct {
	Contract     *UniQuoterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UniQuoterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniQuoterV2Raw struct {
	Contract *UniQuoterV2 // Generic contract binding to access the raw methods on
}

// UniQuoterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniQuoterV2CallerRaw struct {
	Contract *UniQuoterV2Caller // Generic read-only contract binding to access the raw methods on
}

// UniQuoterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniQuoterV2TransactorRaw struct {
	Contract *UniQuoterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniQuoterV2 creates a new instance of UniQuoterV2, bound to a specific deployed contract.
func NewUniQuoterV2(address common.Address, backend bind.ContractBackend) (*UniQuoterV2, error) {
	contract, err := bindUniQuoterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniQuoterV2{UniQuoterV2Caller: UniQuoterV2Caller{contract: contract}, UniQuoterV2Transactor: UniQuoterV2Transactor{contract: contract}, UniQuoterV2Filterer: UniQuoterV2Filterer{contract: contract}}, nil
}

// NewUniQuoterV2Caller creates a new read-only instance of UniQuoterV2, bound to a specific deployed contract.
func NewUniQuoterV2Caller(address common.Address, caller bind.ContractCaller) (*UniQuoterV2Caller, error) {
	contract, err := bindUniQuoterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniQuoterV2Caller{contract: contract}, nil
}

// NewUniQuoterV2Transactor creates a new write-only instance of UniQuoterV2, bound to a specific deployed contract.
func NewUniQuoterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*UniQuoterV2Transactor, error) {
	contract, err := bindUniQuoterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniQuoterV2Transactor{contract: contract}, nil
}

// NewUniQuoterV2Filterer creates a new log filterer instance of UniQuoterV2, bound to a specific deployed contract.
func NewUniQuoterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*UniQuoterV2Filterer, error) {
	contract, err := bindUniQuoterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniQuoterV2Filterer{contract: contract}, nil
}

// bindUniQuoterV2 binds a generic wrapper to an already deployed contract.
func bindUniQuoterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniQuoterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniQuoterV2 *UniQuoterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniQuoterV2.Contract.UniQuoterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniQuoterV2 *UniQuoterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.UniQuoterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniQuoterV2 *UniQuoterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.UniQuoterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniQuoterV2 *UniQuoterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniQuoterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniQuoterV2 *UniQuoterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniQuoterV2 *UniQuoterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniQuoterV2 *UniQuoterV2Caller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniQuoterV2.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniQuoterV2 *UniQuoterV2Session) WETH9() (common.Address, error) {
	return _UniQuoterV2.Contract.WETH9(&_UniQuoterV2.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniQuoterV2 *UniQuoterV2CallerSession) WETH9() (common.Address, error) {
	return _UniQuoterV2.Contract.WETH9(&_UniQuoterV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniQuoterV2 *UniQuoterV2Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniQuoterV2.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniQuoterV2 *UniQuoterV2Session) Factory() (common.Address, error) {
	return _UniQuoterV2.Contract.Factory(&_UniQuoterV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniQuoterV2 *UniQuoterV2CallerSession) Factory() (common.Address, error) {
	return _UniQuoterV2.Contract.Factory(&_UniQuoterV2.CallOpts)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_UniQuoterV2 *UniQuoterV2Caller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _UniQuoterV2.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_UniQuoterV2 *UniQuoterV2Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _UniQuoterV2.Contract.UniswapV3SwapCallback(&_UniQuoterV2.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_UniQuoterV2 *UniQuoterV2CallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _UniQuoterV2.Contract.UniswapV3SwapCallback(&_UniQuoterV2.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2Transactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _UniQuoterV2.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2Session) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.QuoteExactInput(&_UniQuoterV2.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2TransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.QuoteExactInput(&_UniQuoterV2.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2Transactor) QuoteExactInputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _UniQuoterV2.contract.Transact(opts, "quoteExactInputSingle", params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2Session) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.QuoteExactInputSingle(&_UniQuoterV2.TransactOpts, params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2TransactorSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.QuoteExactInputSingle(&_UniQuoterV2.TransactOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2Transactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _UniQuoterV2.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2Session) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.QuoteExactOutput(&_UniQuoterV2.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2TransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.QuoteExactOutput(&_UniQuoterV2.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2Transactor) QuoteExactOutputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _UniQuoterV2.contract.Transact(opts, "quoteExactOutputSingle", params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2Session) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.QuoteExactOutputSingle(&_UniQuoterV2.TransactOpts, params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_UniQuoterV2 *UniQuoterV2TransactorSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _UniQuoterV2.Contract.QuoteExactOutputSingle(&_UniQuoterV2.TransactOpts, params)
}
