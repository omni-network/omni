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

// WOmniMetaData contains all meta data concerning the WOmni contract.
var WOmniMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"guy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"guy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b506106e58061001d5f395ff3fe60806040526004361061009f575f3560e01c8063313ce56711610063578063313ce5671461018557806370a08231146101ab57806395d89b41146101d6578063a9059cbb14610203578063d0e30db0146100ae578063dd62ed3e14610222576100ae565b806306fdde03146100b6578063095ea7b3146100fc57806318160ddd1461012b57806323b872dd146101475780632e1a7d4d14610166576100ae565b366100ae576100ac610258565b005b6100ac610258565b3480156100c1575f80fd5b5060408051808201909152600c81526b57726170706564204f6d6e6960a01b60208201525b6040516100f3919061054c565b60405180910390f35b348015610107575f80fd5b5061011b6101163660046105b3565b6102b2565b60405190151581526020016100f3565b348015610136575f80fd5b50475b6040519081526020016100f3565b348015610152575f80fd5b5061011b6101613660046105db565b61031e565b348015610171575f80fd5b506100ac610180366004610614565b610499565b348015610190575f80fd5b50610199601281565b60405160ff90911681526020016100f3565b3480156101b6575f80fd5b506101396101c536600461062b565b5f6020819052908152604090205481565b3480156101e1575f80fd5b50604080518082019091526005815264574f4d4e4960d81b60208201526100e6565b34801561020e575f80fd5b5061011b61021d3660046105b3565b610539565b34801561022d575f80fd5b5061013961023c366004610644565b600160209081525f928352604080842090915290825290205481565b335f9081526020819052604081208054349290610276908490610689565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b335f8181526001602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259061030c9086815260200190565b60405180910390a35060015b92915050565b6001600160a01b0383165f90815260208190526040812054821115610341575f80fd5b6001600160a01b038416331480159061037d57506001600160a01b0384165f9081526001602090815260408083203384529091529020545f1914155b156103e8576001600160a01b0384165f9081526001602090815260408083203384529091529020548211156103b0575f80fd5b6001600160a01b0384165f908152600160209081526040808320338452909152812080548492906103e290849061069c565b90915550505b6001600160a01b0384165f908152602081905260408120805484929061040f90849061069c565b90915550506001600160a01b0383165f908152602081905260408120805484929061043b908490610689565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161048791815260200190565b60405180910390a35060019392505050565b335f908152602081905260409020548111156104b3575f80fd5b335f90815260208190526040812080548392906104d190849061069c565b9091555050604051339082156108fc029083905f818181858888f19350505050158015610500573d5f803e3d5ffd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b5f61054533848461031e565b9392505050565b5f602080835283518060208501525f5b818110156105785785810183015185820160400152820161055c565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146105ae575f80fd5b919050565b5f80604083850312156105c4575f80fd5b6105cd83610598565b946020939093013593505050565b5f805f606084860312156105ed575f80fd5b6105f684610598565b925061060460208501610598565b9150604084013590509250925092565b5f60208284031215610624575f80fd5b5035919050565b5f6020828403121561063b575f80fd5b61054582610598565b5f8060408385031215610655575f80fd5b61065e83610598565b915061066c60208401610598565b90509250929050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561031857610318610675565b818103818111156103185761031861067556fea2646970667358221220c519ca043c31f4d44aba7fd72f0f11c408c46fd1b691dcdad2d62c51e30d775f64736f6c63430008180033",
}

// WOmniABI is the input ABI used to generate the binding from.
// Deprecated: Use WOmniMetaData.ABI instead.
var WOmniABI = WOmniMetaData.ABI

// WOmniBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WOmniMetaData.Bin instead.
var WOmniBin = WOmniMetaData.Bin

// DeployWOmni deploys a new Ethereum contract, binding an instance of WOmni to it.
func DeployWOmni(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WOmni, error) {
	parsed, err := WOmniMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WOmniBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WOmni{WOmniCaller: WOmniCaller{contract: contract}, WOmniTransactor: WOmniTransactor{contract: contract}, WOmniFilterer: WOmniFilterer{contract: contract}}, nil
}

// WOmni is an auto generated Go binding around an Ethereum contract.
type WOmni struct {
	WOmniCaller     // Read-only binding to the contract
	WOmniTransactor // Write-only binding to the contract
	WOmniFilterer   // Log filterer for contract events
}

// WOmniCaller is an auto generated read-only Go binding around an Ethereum contract.
type WOmniCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WOmniTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WOmniTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WOmniFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WOmniFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WOmniSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WOmniSession struct {
	Contract     *WOmni            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WOmniCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WOmniCallerSession struct {
	Contract *WOmniCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WOmniTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WOmniTransactorSession struct {
	Contract     *WOmniTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WOmniRaw is an auto generated low-level Go binding around an Ethereum contract.
type WOmniRaw struct {
	Contract *WOmni // Generic contract binding to access the raw methods on
}

// WOmniCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WOmniCallerRaw struct {
	Contract *WOmniCaller // Generic read-only contract binding to access the raw methods on
}

// WOmniTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WOmniTransactorRaw struct {
	Contract *WOmniTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWOmni creates a new instance of WOmni, bound to a specific deployed contract.
func NewWOmni(address common.Address, backend bind.ContractBackend) (*WOmni, error) {
	contract, err := bindWOmni(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WOmni{WOmniCaller: WOmniCaller{contract: contract}, WOmniTransactor: WOmniTransactor{contract: contract}, WOmniFilterer: WOmniFilterer{contract: contract}}, nil
}

// NewWOmniCaller creates a new read-only instance of WOmni, bound to a specific deployed contract.
func NewWOmniCaller(address common.Address, caller bind.ContractCaller) (*WOmniCaller, error) {
	contract, err := bindWOmni(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WOmniCaller{contract: contract}, nil
}

// NewWOmniTransactor creates a new write-only instance of WOmni, bound to a specific deployed contract.
func NewWOmniTransactor(address common.Address, transactor bind.ContractTransactor) (*WOmniTransactor, error) {
	contract, err := bindWOmni(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WOmniTransactor{contract: contract}, nil
}

// NewWOmniFilterer creates a new log filterer instance of WOmni, bound to a specific deployed contract.
func NewWOmniFilterer(address common.Address, filterer bind.ContractFilterer) (*WOmniFilterer, error) {
	contract, err := bindWOmni(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WOmniFilterer{contract: contract}, nil
}

// bindWOmni binds a generic wrapper to an already deployed contract.
func bindWOmni(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WOmniMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WOmni *WOmniRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WOmni.Contract.WOmniCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WOmni *WOmniRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WOmni.Contract.WOmniTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WOmni *WOmniRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WOmni.Contract.WOmniTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WOmni *WOmniCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WOmni.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WOmni *WOmniTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WOmni.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WOmni *WOmniTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WOmni.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WOmni *WOmniCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WOmni.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WOmni *WOmniSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WOmni.Contract.Allowance(&_WOmni.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WOmni *WOmniCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WOmni.Contract.Allowance(&_WOmni.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WOmni *WOmniCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WOmni.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WOmni *WOmniSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WOmni.Contract.BalanceOf(&_WOmni.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WOmni *WOmniCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WOmni.Contract.BalanceOf(&_WOmni.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WOmni *WOmniCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WOmni.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WOmni *WOmniSession) Decimals() (uint8, error) {
	return _WOmni.Contract.Decimals(&_WOmni.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WOmni *WOmniCallerSession) Decimals() (uint8, error) {
	return _WOmni.Contract.Decimals(&_WOmni.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WOmni *WOmniCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WOmni.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WOmni *WOmniSession) Name() (string, error) {
	return _WOmni.Contract.Name(&_WOmni.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WOmni *WOmniCallerSession) Name() (string, error) {
	return _WOmni.Contract.Name(&_WOmni.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WOmni *WOmniCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WOmni.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WOmni *WOmniSession) Symbol() (string, error) {
	return _WOmni.Contract.Symbol(&_WOmni.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WOmni *WOmniCallerSession) Symbol() (string, error) {
	return _WOmni.Contract.Symbol(&_WOmni.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WOmni *WOmniCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WOmni.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WOmni *WOmniSession) TotalSupply() (*big.Int, error) {
	return _WOmni.Contract.TotalSupply(&_WOmni.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WOmni *WOmniCallerSession) TotalSupply() (*big.Int, error) {
	return _WOmni.Contract.TotalSupply(&_WOmni.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WOmni *WOmniTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WOmni *WOmniSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.Contract.Approve(&_WOmni.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WOmni *WOmniTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.Contract.Approve(&_WOmni.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WOmni *WOmniTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WOmni.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WOmni *WOmniSession) Deposit() (*types.Transaction, error) {
	return _WOmni.Contract.Deposit(&_WOmni.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WOmni *WOmniTransactorSession) Deposit() (*types.Transaction, error) {
	return _WOmni.Contract.Deposit(&_WOmni.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WOmni *WOmniTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WOmni *WOmniSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.Contract.Transfer(&_WOmni.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WOmni *WOmniTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.Contract.Transfer(&_WOmni.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WOmni *WOmniTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WOmni *WOmniSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.Contract.TransferFrom(&_WOmni.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WOmni *WOmniTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.Contract.TransferFrom(&_WOmni.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WOmni *WOmniTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WOmni.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WOmni *WOmniSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WOmni.Contract.Withdraw(&_WOmni.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WOmni *WOmniTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WOmni.Contract.Withdraw(&_WOmni.TransactOpts, wad)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WOmni *WOmniTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WOmni.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WOmni *WOmniSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WOmni.Contract.Fallback(&_WOmni.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WOmni *WOmniTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WOmni.Contract.Fallback(&_WOmni.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WOmni *WOmniTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WOmni.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WOmni *WOmniSession) Receive() (*types.Transaction, error) {
	return _WOmni.Contract.Receive(&_WOmni.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WOmni *WOmniTransactorSession) Receive() (*types.Transaction, error) {
	return _WOmni.Contract.Receive(&_WOmni.TransactOpts)
}

// WOmniApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WOmni contract.
type WOmniApprovalIterator struct {
	Event *WOmniApproval // Event containing the contract specifics and raw log

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
func (it *WOmniApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WOmniApproval)
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
		it.Event = new(WOmniApproval)
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
func (it *WOmniApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WOmniApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WOmniApproval represents a Approval event raised by the WOmni contract.
type WOmniApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WOmni *WOmniFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WOmniApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WOmni.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &WOmniApprovalIterator{contract: _WOmni.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WOmni *WOmniFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WOmniApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WOmni.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WOmniApproval)
				if err := _WOmni.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WOmni *WOmniFilterer) ParseApproval(log types.Log) (*WOmniApproval, error) {
	event := new(WOmniApproval)
	if err := _WOmni.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WOmniDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WOmni contract.
type WOmniDepositIterator struct {
	Event *WOmniDeposit // Event containing the contract specifics and raw log

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
func (it *WOmniDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WOmniDeposit)
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
		it.Event = new(WOmniDeposit)
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
func (it *WOmniDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WOmniDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WOmniDeposit represents a Deposit event raised by the WOmni contract.
type WOmniDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WOmni *WOmniFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WOmniDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WOmni.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WOmniDepositIterator{contract: _WOmni.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WOmni *WOmniFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WOmniDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WOmni.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WOmniDeposit)
				if err := _WOmni.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WOmni *WOmniFilterer) ParseDeposit(log types.Log) (*WOmniDeposit, error) {
	event := new(WOmniDeposit)
	if err := _WOmni.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WOmniTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WOmni contract.
type WOmniTransferIterator struct {
	Event *WOmniTransfer // Event containing the contract specifics and raw log

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
func (it *WOmniTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WOmniTransfer)
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
		it.Event = new(WOmniTransfer)
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
func (it *WOmniTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WOmniTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WOmniTransfer represents a Transfer event raised by the WOmni contract.
type WOmniTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WOmni *WOmniFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WOmniTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WOmni.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &WOmniTransferIterator{contract: _WOmni.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WOmni *WOmniFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WOmniTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WOmni.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WOmniTransfer)
				if err := _WOmni.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WOmni *WOmniFilterer) ParseTransfer(log types.Log) (*WOmniTransfer, error) {
	event := new(WOmniTransfer)
	if err := _WOmni.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WOmniWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WOmni contract.
type WOmniWithdrawalIterator struct {
	Event *WOmniWithdrawal // Event containing the contract specifics and raw log

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
func (it *WOmniWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WOmniWithdrawal)
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
		it.Event = new(WOmniWithdrawal)
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
func (it *WOmniWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WOmniWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WOmniWithdrawal represents a Withdrawal event raised by the WOmni contract.
type WOmniWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WOmni *WOmniFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WOmniWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WOmni.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WOmniWithdrawalIterator{contract: _WOmni.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WOmni *WOmniFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WOmniWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WOmni.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WOmniWithdrawal)
				if err := _WOmni.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WOmni *WOmniFilterer) ParseWithdrawal(log types.Log) (*WOmniWithdrawal, error) {
	event := new(WOmniWithdrawal)
	if err := _WOmni.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
