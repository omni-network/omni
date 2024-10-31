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

// MockSymbioticVaultMetaData contains all meta data concerning the MockSymbioticVault contract.
var MockSymbioticVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"newCollateral\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516105af3803806105af83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b60805161051861009760003960008181609e0152818160f8015261018301526105186000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806327e235e31461005157806347e7ef2414610084578063d8dfeb4514610099578063f3fef3a3146100d8575b600080fd5b61007161005f366004610410565b60006020819052908152604090205481565b6040519081526020015b60405180910390f35b61009761009236600461042b565b6100eb565b005b6100c07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161007b565b6100976100e636600461042b565b610151565b6101206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163330846101ae565b6001600160a01b0382166000908152602081905260408120805483929061014890849061046b565b90915550505050565b336000908152602081905260408120805483929061017090849061047e565b909155506101aa90506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016838361021b565b5050565b6040516001600160a01b0384811660248301528381166044830152606482018390526102159186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050610251565b50505050565b6040516001600160a01b0383811660248301526044820183905261024c91859182169063a9059cbb906064016101e3565b505050565b60006102666001600160a01b038416836102b9565b9050805160001415801561028b5750808060200190518101906102899190610491565b155b1561024c57604051635274afe760e01b81526001600160a01b03841660048201526024015b60405180910390fd5b60606102c7838360006102d0565b90505b92915050565b6060814710156102f55760405163cd78605960e01b81523060048201526024016102b0565b600080856001600160a01b0316848660405161031191906104b3565b60006040518083038185875af1925050503d806000811461034e576040519150601f19603f3d011682016040523d82523d6000602084013e610353565b606091505b509150915061036386838361036f565b925050505b9392505050565b6060826103845761037f826103cb565b610368565b815115801561039b57506001600160a01b0384163b155b156103c457604051639996b31560e01b81526001600160a01b03851660048201526024016102b0565b5080610368565b8051156103db5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80356001600160a01b038116811461040b57600080fd5b919050565b60006020828403121561042257600080fd5b6102c7826103f4565b6000806040838503121561043e57600080fd5b610447836103f4565b946020939093013593505050565b634e487b7160e01b600052601160045260246000fd5b808201808211156102ca576102ca610455565b818103818111156102ca576102ca610455565b6000602082840312156104a357600080fd5b8151801515811461036857600080fd5b6000825160005b818110156104d457602081860181015185830152016104ba565b50600092019182525091905056fea2646970667358221220cf906c76df6feca017ffa507594c347ed9e47e71be5d82cff000c6c553a2c87c64736f6c63430008180033",
}

// MockSymbioticVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use MockSymbioticVaultMetaData.ABI instead.
var MockSymbioticVaultABI = MockSymbioticVaultMetaData.ABI

// MockSymbioticVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockSymbioticVaultMetaData.Bin instead.
var MockSymbioticVaultBin = MockSymbioticVaultMetaData.Bin

// DeployMockSymbioticVault deploys a new Ethereum contract, binding an instance of MockSymbioticVault to it.
func DeployMockSymbioticVault(auth *bind.TransactOpts, backend bind.ContractBackend, newCollateral common.Address) (common.Address, *types.Transaction, *MockSymbioticVault, error) {
	parsed, err := MockSymbioticVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockSymbioticVaultBin), backend, newCollateral)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockSymbioticVault{MockSymbioticVaultCaller: MockSymbioticVaultCaller{contract: contract}, MockSymbioticVaultTransactor: MockSymbioticVaultTransactor{contract: contract}, MockSymbioticVaultFilterer: MockSymbioticVaultFilterer{contract: contract}}, nil
}

// MockSymbioticVault is an auto generated Go binding around an Ethereum contract.
type MockSymbioticVault struct {
	MockSymbioticVaultCaller     // Read-only binding to the contract
	MockSymbioticVaultTransactor // Write-only binding to the contract
	MockSymbioticVaultFilterer   // Log filterer for contract events
}

// MockSymbioticVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockSymbioticVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockSymbioticVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockSymbioticVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockSymbioticVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockSymbioticVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockSymbioticVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockSymbioticVaultSession struct {
	Contract     *MockSymbioticVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockSymbioticVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockSymbioticVaultCallerSession struct {
	Contract *MockSymbioticVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MockSymbioticVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockSymbioticVaultTransactorSession struct {
	Contract     *MockSymbioticVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockSymbioticVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockSymbioticVaultRaw struct {
	Contract *MockSymbioticVault // Generic contract binding to access the raw methods on
}

// MockSymbioticVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockSymbioticVaultCallerRaw struct {
	Contract *MockSymbioticVaultCaller // Generic read-only contract binding to access the raw methods on
}

// MockSymbioticVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockSymbioticVaultTransactorRaw struct {
	Contract *MockSymbioticVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockSymbioticVault creates a new instance of MockSymbioticVault, bound to a specific deployed contract.
func NewMockSymbioticVault(address common.Address, backend bind.ContractBackend) (*MockSymbioticVault, error) {
	contract, err := bindMockSymbioticVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockSymbioticVault{MockSymbioticVaultCaller: MockSymbioticVaultCaller{contract: contract}, MockSymbioticVaultTransactor: MockSymbioticVaultTransactor{contract: contract}, MockSymbioticVaultFilterer: MockSymbioticVaultFilterer{contract: contract}}, nil
}

// NewMockSymbioticVaultCaller creates a new read-only instance of MockSymbioticVault, bound to a specific deployed contract.
func NewMockSymbioticVaultCaller(address common.Address, caller bind.ContractCaller) (*MockSymbioticVaultCaller, error) {
	contract, err := bindMockSymbioticVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockSymbioticVaultCaller{contract: contract}, nil
}

// NewMockSymbioticVaultTransactor creates a new write-only instance of MockSymbioticVault, bound to a specific deployed contract.
func NewMockSymbioticVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*MockSymbioticVaultTransactor, error) {
	contract, err := bindMockSymbioticVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockSymbioticVaultTransactor{contract: contract}, nil
}

// NewMockSymbioticVaultFilterer creates a new log filterer instance of MockSymbioticVault, bound to a specific deployed contract.
func NewMockSymbioticVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*MockSymbioticVaultFilterer, error) {
	contract, err := bindMockSymbioticVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockSymbioticVaultFilterer{contract: contract}, nil
}

// bindMockSymbioticVault binds a generic wrapper to an already deployed contract.
func bindMockSymbioticVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockSymbioticVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockSymbioticVault *MockSymbioticVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockSymbioticVault.Contract.MockSymbioticVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockSymbioticVault *MockSymbioticVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockSymbioticVault.Contract.MockSymbioticVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockSymbioticVault *MockSymbioticVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockSymbioticVault.Contract.MockSymbioticVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockSymbioticVault *MockSymbioticVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockSymbioticVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockSymbioticVault *MockSymbioticVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockSymbioticVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockSymbioticVault *MockSymbioticVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockSymbioticVault.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address depositor) view returns(uint256 balance)
func (_MockSymbioticVault *MockSymbioticVaultCaller) Balances(opts *bind.CallOpts, depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockSymbioticVault.contract.Call(opts, &out, "balances", depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address depositor) view returns(uint256 balance)
func (_MockSymbioticVault *MockSymbioticVaultSession) Balances(depositor common.Address) (*big.Int, error) {
	return _MockSymbioticVault.Contract.Balances(&_MockSymbioticVault.CallOpts, depositor)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address depositor) view returns(uint256 balance)
func (_MockSymbioticVault *MockSymbioticVaultCallerSession) Balances(depositor common.Address) (*big.Int, error) {
	return _MockSymbioticVault.Contract.Balances(&_MockSymbioticVault.CallOpts, depositor)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_MockSymbioticVault *MockSymbioticVaultCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockSymbioticVault.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_MockSymbioticVault *MockSymbioticVaultSession) Collateral() (common.Address, error) {
	return _MockSymbioticVault.Contract.Collateral(&_MockSymbioticVault.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_MockSymbioticVault *MockSymbioticVaultCallerSession) Collateral() (common.Address, error) {
	return _MockSymbioticVault.Contract.Collateral(&_MockSymbioticVault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) returns()
func (_MockSymbioticVault *MockSymbioticVaultTransactor) Deposit(opts *bind.TransactOpts, onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockSymbioticVault.contract.Transact(opts, "deposit", onBehalfOf, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) returns()
func (_MockSymbioticVault *MockSymbioticVaultSession) Deposit(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockSymbioticVault.Contract.Deposit(&_MockSymbioticVault.TransactOpts, onBehalfOf, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) returns()
func (_MockSymbioticVault *MockSymbioticVaultTransactorSession) Deposit(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockSymbioticVault.Contract.Deposit(&_MockSymbioticVault.TransactOpts, onBehalfOf, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_MockSymbioticVault *MockSymbioticVaultTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockSymbioticVault.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_MockSymbioticVault *MockSymbioticVaultSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockSymbioticVault.Contract.Withdraw(&_MockSymbioticVault.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_MockSymbioticVault *MockSymbioticVaultTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockSymbioticVault.Contract.Withdraw(&_MockSymbioticVault.TransactOpts, to, amount)
}
