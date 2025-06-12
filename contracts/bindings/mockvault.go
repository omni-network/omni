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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"newCollateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositWithoutBehalfOf\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060405161050c38038061050c83398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f80fd5b81516001600160a01b0381168114610065575f80fd5b9392505050565b60805161046d61009f5f395f818160b9015281816101640152818161019e01528181610264015261029e015261046d5ff3fe608060405260043610610049575f3560e01c80631b190b5f1461004d57806327e235e31461005757806347e7ef2414610095578063d8dfeb45146100a8578063f3fef3a3146100f3575b5f80fd5b610055610112565b005b348015610062575f80fd5b506100826100713660046103af565b5f6020819052908152604090205481565b6040519081526020015b60405180910390f35b6100556100a33660046103cf565b610162565b3480156100b3575f80fd5b506100db7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161008c565b3480156100fe575f80fd5b5061005561010d3660046103cf565b61023f565b5f34116101565760405162461bcd60e51b815260206004820152600d60248201526c139bc8115d1a195c881cd95b9d609a1b60448201526064015b60405180910390fd5b61016033346102d8565b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156101cb576101c66001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163330846102f1565b61020f565b80341461020f5760405162461bcd60e51b8152602060048201526012602482015271125b9d985b1a590818dbdb1b185d195c985b60721b604482015260640161014d565b6001600160a01b0382165f908152602081905260408120805483929061023690849061040b565b90915550505050565b335f908152602081905260408120805483929061025d908490610424565b90915550507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156102c9576102c56001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016838361034a565b5050565b6102c56001600160a01b038316825b5f385f3884865af16102c55763b12d13eb5f526004601cfd5b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f51141661033c57803d873b15171061033c57637939f4245f526004601cfd5b505f60605260405250505050565b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f51141661038a57803d853b15171061038a576390b8ec185f526004601cfd5b505f603452505050565b80356001600160a01b03811681146103aa575f80fd5b919050565b5f602082840312156103bf575f80fd5b6103c882610394565b9392505050565b5f80604083850312156103e0575f80fd5b6103e983610394565b946020939093013593505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561041e5761041e6103f7565b92915050565b8181038181111561041e5761041e6103f756fea2646970667358221220836499affd527b41617767cd44932dfe509466f30be64f1315d0576ee3b9796564736f6c63430008180033",
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

// DepositWithoutBehalfOf is a paid mutator transaction binding the contract method 0x1b190b5f.
//
// Solidity: function depositWithoutBehalfOf() payable returns()
func (_MockVault *MockVaultTransactor) DepositWithoutBehalfOf(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockVault.contract.Transact(opts, "depositWithoutBehalfOf")
}

// DepositWithoutBehalfOf is a paid mutator transaction binding the contract method 0x1b190b5f.
//
// Solidity: function depositWithoutBehalfOf() payable returns()
func (_MockVault *MockVaultSession) DepositWithoutBehalfOf() (*types.Transaction, error) {
	return _MockVault.Contract.DepositWithoutBehalfOf(&_MockVault.TransactOpts)
}

// DepositWithoutBehalfOf is a paid mutator transaction binding the contract method 0x1b190b5f.
//
// Solidity: function depositWithoutBehalfOf() payable returns()
func (_MockVault *MockVaultTransactorSession) DepositWithoutBehalfOf() (*types.Transaction, error) {
	return _MockVault.Contract.DepositWithoutBehalfOf(&_MockVault.TransactOpts)
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
