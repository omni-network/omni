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

// MockVaultMetaData contains all meta data concerning the MockVault contract.
var MockVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"newCollateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b506040516104b13803806104b183398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f80fd5b81516001600160a01b0381168114610065575f80fd5b9392505050565b60805161041261009f5f395f818160a6015281816101010152818161013b01528181610205015261023f01526104125ff3fe60806040526004361061003e575f3560e01c806327e235e31461004257806347e7ef2414610080578063d8dfeb4514610095578063f3fef3a3146100e0575b5f80fd5b34801561004d575f80fd5b5061006d61005c366004610354565b5f6020819052908152604090205481565b6040519081526020015b60405180910390f35b61009361008e366004610374565b6100ff565b005b3480156100a0575f80fd5b506100c87f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610077565b3480156100eb575f80fd5b506100936100fa366004610374565b6101e0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615610168576101636001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001633308461027d565b6101b0565b8034146101b05760405162461bcd60e51b8152602060048201526012602482015271125b9d985b1a590818dbdb1b185d195c985b60721b604482015260640160405180910390fd5b6001600160a01b0382165f90815260208190526040812080548392906101d79084906103b0565b90915550505050565b335f90815260208190526040812080548392906101fe9084906103c9565b90915550507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161561026a576102666001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001683836102d6565b5050565b6102666001600160a01b03831682610320565b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f5114166102c857803d873b1517106102c857637939f4245f526004601cfd5b505f60605260405250505050565b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f51141661031657803d853b151710610316576390b8ec185f526004601cfd5b505f603452505050565b5f385f3884865af16102665763b12d13eb5f526004601cfd5b80356001600160a01b038116811461034f575f80fd5b919050565b5f60208284031215610364575f80fd5b61036d82610339565b9392505050565b5f8060408385031215610385575f80fd5b61038e83610339565b946020939093013593505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156103c3576103c361039c565b92915050565b818103818111156103c3576103c361039c56fea26469706673582212209c1ed4e7fd38cfb207e3cc64e8aeeec3b6f44846ba65dbb438f476cccf690d4464736f6c63430008180033",
}

// MockVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use MockVaultMetaData.ABI instead.
var MockVaultABI = MockVaultMetaData.ABI

// MockVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockVaultMetaData.Bin instead.
var MockVaultBin = MockVaultMetaData.Bin

// DeployMockVault deploys a new Ethereum contract, binding an instance of MockVault to it.
func DeployMockVault(auth *bind.TransactOpts, backend bind.ContractBackend, newCollateral common.Address) (common.Address, *types.Transaction, *MockVault, error) {
	parsed, err := MockVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockVaultBin), backend, newCollateral)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockVault{MockVaultCaller: MockVaultCaller{contract: contract}, MockVaultTransactor: MockVaultTransactor{contract: contract}, MockVaultFilterer: MockVaultFilterer{contract: contract}}, nil
}

// MockVault is an auto generated Go binding around an Ethereum contract.
type MockVault struct {
	MockVaultCaller     // Read-only binding to the contract
	MockVaultTransactor // Write-only binding to the contract
	MockVaultFilterer   // Log filterer for contract events
}

// MockVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockVaultSession struct {
	Contract     *MockVault        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockVaultCallerSession struct {
	Contract *MockVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MockVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockVaultTransactorSession struct {
	Contract     *MockVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MockVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockVaultRaw struct {
	Contract *MockVault // Generic contract binding to access the raw methods on
}

// MockVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockVaultCallerRaw struct {
	Contract *MockVaultCaller // Generic read-only contract binding to access the raw methods on
}

// MockVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockVaultTransactorRaw struct {
	Contract *MockVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockVault creates a new instance of MockVault, bound to a specific deployed contract.
func NewMockVault(address common.Address, backend bind.ContractBackend) (*MockVault, error) {
	contract, err := bindMockVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockVault{MockVaultCaller: MockVaultCaller{contract: contract}, MockVaultTransactor: MockVaultTransactor{contract: contract}, MockVaultFilterer: MockVaultFilterer{contract: contract}}, nil
}

// NewMockVaultCaller creates a new read-only instance of MockVault, bound to a specific deployed contract.
func NewMockVaultCaller(address common.Address, caller bind.ContractCaller) (*MockVaultCaller, error) {
	contract, err := bindMockVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockVaultCaller{contract: contract}, nil
}

// NewMockVaultTransactor creates a new write-only instance of MockVault, bound to a specific deployed contract.
func NewMockVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*MockVaultTransactor, error) {
	contract, err := bindMockVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockVaultTransactor{contract: contract}, nil
}

// NewMockVaultFilterer creates a new log filterer instance of MockVault, bound to a specific deployed contract.
func NewMockVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*MockVaultFilterer, error) {
	contract, err := bindMockVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockVaultFilterer{contract: contract}, nil
}

// bindMockVault binds a generic wrapper to an already deployed contract.
func bindMockVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockVault *MockVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockVault.Contract.MockVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockVault *MockVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockVault.Contract.MockVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockVault *MockVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockVault.Contract.MockVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockVault *MockVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockVault *MockVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockVault *MockVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockVault.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address depositor) view returns(uint256 balance)
func (_MockVault *MockVaultCaller) Balances(opts *bind.CallOpts, depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockVault.contract.Call(opts, &out, "balances", depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address depositor) view returns(uint256 balance)
func (_MockVault *MockVaultSession) Balances(depositor common.Address) (*big.Int, error) {
	return _MockVault.Contract.Balances(&_MockVault.CallOpts, depositor)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address depositor) view returns(uint256 balance)
func (_MockVault *MockVaultCallerSession) Balances(depositor common.Address) (*big.Int, error) {
	return _MockVault.Contract.Balances(&_MockVault.CallOpts, depositor)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_MockVault *MockVaultCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockVault.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_MockVault *MockVaultSession) Collateral() (common.Address, error) {
	return _MockVault.Contract.Collateral(&_MockVault.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_MockVault *MockVaultCallerSession) Collateral() (common.Address, error) {
	return _MockVault.Contract.Collateral(&_MockVault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) payable returns()
func (_MockVault *MockVaultTransactor) Deposit(opts *bind.TransactOpts, onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockVault.contract.Transact(opts, "deposit", onBehalfOf, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) payable returns()
func (_MockVault *MockVaultSession) Deposit(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockVault.Contract.Deposit(&_MockVault.TransactOpts, onBehalfOf, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) payable returns()
func (_MockVault *MockVaultTransactorSession) Deposit(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockVault.Contract.Deposit(&_MockVault.TransactOpts, onBehalfOf, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_MockVault *MockVaultTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockVault.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_MockVault *MockVaultSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockVault.Contract.Withdraw(&_MockVault.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_MockVault *MockVaultTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockVault.Contract.Withdraw(&_MockVault.TransactOpts, to, amount)
}
