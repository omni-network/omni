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

// MockERC20MetaData contains all meta data concerning the MockERC20 contract.
var MockERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000c3d38038062000c3d833981016040819052620000349162000126565b8181600362000044838262000221565b50600462000053828262000221565b5050505050620002ed565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200008657600080fd5b81516001600160401b0380821115620000a357620000a36200005e565b604051601f8301601f19908116603f01168101908282118183101715620000ce57620000ce6200005e565b8160405283815260209250866020858801011115620000ec57600080fd5b600091505b83821015620001105785820183015181830184015290820190620000f1565b6000602085830101528094505050505092915050565b600080604083850312156200013a57600080fd5b82516001600160401b03808211156200015257600080fd5b620001608683870162000074565b935060208501519150808211156200017757600080fd5b50620001868582860162000074565b9150509250929050565b600181811c90821680620001a557607f821691505b602082108103620001c657634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200021c576000816000526020600020601f850160051c81016020861015620001f75750805b601f850160051c820191505b81811015620002185782815560010162000203565b5050505b505050565b81516001600160401b038111156200023d576200023d6200005e565b62000255816200024e845462000190565b84620001cc565b602080601f8311600181146200028d5760008415620002745750858301515b600019600386901b1c1916600185901b17855562000218565b600085815260208120601f198616915b82811015620002be578886015182559484019460019091019084016200029d565b5085821015620002dd5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61094080620002fd6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806340c10f191161007157806340c10f191461014157806370a082311461015657806395d89b411461017f578063a457c2d714610187578063a9059cbb1461019a578063dd62ed3e146101ad57600080fd5b806306fdde03146100b9578063095ea7b3146100d757806318160ddd146100fa57806323b872dd1461010c578063313ce5671461011f578063395093511461012e575b600080fd5b6100c16101c0565b6040516100ce9190610789565b60405180910390f35b6100ea6100e53660046107f4565b610252565b60405190151581526020016100ce565b6002545b6040519081526020016100ce565b6100ea61011a36600461081e565b61026c565b604051601281526020016100ce565b6100ea61013c3660046107f4565b610290565b61015461014f3660046107f4565b6102b2565b005b6100fe61016436600461085a565b6001600160a01b031660009081526020819052604090205490565b6100c16102c0565b6100ea6101953660046107f4565b6102cf565b6100ea6101a83660046107f4565b61034f565b6100fe6101bb36600461087c565b61035d565b6060600380546101cf906108af565b80601f01602080910402602001604051908101604052809291908181526020018280546101fb906108af565b80156102485780601f1061021d57610100808354040283529160200191610248565b820191906000526020600020905b81548152906001019060200180831161022b57829003601f168201915b5050505050905090565b600033610260818585610388565b60019150505b92915050565b60003361027a8582856104ac565b610285858585610526565b506001949350505050565b6000336102608185856102a3838361035d565b6102ad91906108e9565b610388565b6102bc82826106ca565b5050565b6060600480546101cf906108af565b600033816102dd828661035d565b9050838110156103425760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102858286868403610388565b600033610260818585610526565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103ea5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610339565b6001600160a01b03821661044b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610339565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60006104b8848461035d565b9050600019811461052057818110156105135760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610339565b6105208484848403610388565b50505050565b6001600160a01b03831661058a5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610339565b6001600160a01b0382166105ec5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610339565b6001600160a01b038316600090815260208190526040902054818110156106645760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610339565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610520565b6001600160a01b0382166107205760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610339565b806002600082825461073291906108e9565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b60006020808352835180602085015260005b818110156107b75785810183015185820160400152820161079b565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146107ef57600080fd5b919050565b6000806040838503121561080757600080fd5b610810836107d8565b946020939093013593505050565b60008060006060848603121561083357600080fd5b61083c846107d8565b925061084a602085016107d8565b9150604084013590509250925092565b60006020828403121561086c57600080fd5b610875826107d8565b9392505050565b6000806040838503121561088f57600080fd5b610898836107d8565b91506108a6602084016107d8565b90509250929050565b600181811c908216806108c357607f821691505b6020821081036108e357634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561026657634e487b7160e01b600052601160045260246000fdfea26469706673582212200e5ff8af1ddb3a64655b3cac018e983177503fd5a8a56a4a55979999bdb48dff64736f6c63430008180033",
}

// MockERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockERC20MetaData.ABI instead.
var MockERC20ABI = MockERC20MetaData.ABI

// MockERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockERC20MetaData.Bin instead.
var MockERC20Bin = MockERC20MetaData.Bin

// DeployMockERC20 deploys a new Ethereum contract, binding an instance of MockERC20 to it.
func DeployMockERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *MockERC20, error) {
	parsed, err := MockERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockERC20Bin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockERC20{MockERC20Caller: MockERC20Caller{contract: contract}, MockERC20Transactor: MockERC20Transactor{contract: contract}, MockERC20Filterer: MockERC20Filterer{contract: contract}}, nil
}

// MockERC20 is an auto generated Go binding around an Ethereum contract.
type MockERC20 struct {
	MockERC20Caller     // Read-only binding to the contract
	MockERC20Transactor // Write-only binding to the contract
	MockERC20Filterer   // Log filterer for contract events
}

// MockERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockERC20Session struct {
	Contract     *MockERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockERC20CallerSession struct {
	Contract *MockERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MockERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockERC20TransactorSession struct {
	Contract     *MockERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MockERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockERC20Raw struct {
	Contract *MockERC20 // Generic contract binding to access the raw methods on
}

// MockERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockERC20CallerRaw struct {
	Contract *MockERC20Caller // Generic read-only contract binding to access the raw methods on
}

// MockERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockERC20TransactorRaw struct {
	Contract *MockERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockERC20 creates a new instance of MockERC20, bound to a specific deployed contract.
func NewMockERC20(address common.Address, backend bind.ContractBackend) (*MockERC20, error) {
	contract, err := bindMockERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockERC20{MockERC20Caller: MockERC20Caller{contract: contract}, MockERC20Transactor: MockERC20Transactor{contract: contract}, MockERC20Filterer: MockERC20Filterer{contract: contract}}, nil
}

// NewMockERC20Caller creates a new read-only instance of MockERC20, bound to a specific deployed contract.
func NewMockERC20Caller(address common.Address, caller bind.ContractCaller) (*MockERC20Caller, error) {
	contract, err := bindMockERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC20Caller{contract: contract}, nil
}

// NewMockERC20Transactor creates a new write-only instance of MockERC20, bound to a specific deployed contract.
func NewMockERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*MockERC20Transactor, error) {
	contract, err := bindMockERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC20Transactor{contract: contract}, nil
}

// NewMockERC20Filterer creates a new log filterer instance of MockERC20, bound to a specific deployed contract.
func NewMockERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*MockERC20Filterer, error) {
	contract, err := bindMockERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockERC20Filterer{contract: contract}, nil
}

// bindMockERC20 binds a generic wrapper to an already deployed contract.
func bindMockERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC20 *MockERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC20.Contract.MockERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC20 *MockERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC20.Contract.MockERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC20 *MockERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC20.Contract.MockERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC20 *MockERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC20 *MockERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC20 *MockERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockERC20 *MockERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockERC20 *MockERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockERC20.Contract.Allowance(&_MockERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockERC20 *MockERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockERC20.Contract.Allowance(&_MockERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockERC20 *MockERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockERC20 *MockERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockERC20.Contract.BalanceOf(&_MockERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockERC20 *MockERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockERC20.Contract.BalanceOf(&_MockERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockERC20 *MockERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockERC20 *MockERC20Session) Decimals() (uint8, error) {
	return _MockERC20.Contract.Decimals(&_MockERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockERC20 *MockERC20CallerSession) Decimals() (uint8, error) {
	return _MockERC20.Contract.Decimals(&_MockERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC20 *MockERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC20 *MockERC20Session) Name() (string, error) {
	return _MockERC20.Contract.Name(&_MockERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC20 *MockERC20CallerSession) Name() (string, error) {
	return _MockERC20.Contract.Name(&_MockERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC20 *MockERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC20 *MockERC20Session) Symbol() (string, error) {
	return _MockERC20.Contract.Symbol(&_MockERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC20 *MockERC20CallerSession) Symbol() (string, error) {
	return _MockERC20.Contract.Symbol(&_MockERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockERC20 *MockERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockERC20 *MockERC20Session) TotalSupply() (*big.Int, error) {
	return _MockERC20.Contract.TotalSupply(&_MockERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockERC20 *MockERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _MockERC20.Contract.TotalSupply(&_MockERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Approve(&_MockERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Approve(&_MockERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockERC20 *MockERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockERC20 *MockERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.DecreaseAllowance(&_MockERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockERC20 *MockERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.DecreaseAllowance(&_MockERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockERC20 *MockERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockERC20 *MockERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.IncreaseAllowance(&_MockERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockERC20 *MockERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.IncreaseAllowance(&_MockERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockERC20 *MockERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockERC20 *MockERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Mint(&_MockERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockERC20 *MockERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Mint(&_MockERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Transfer(&_MockERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Transfer(&_MockERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.TransferFrom(&_MockERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockERC20 *MockERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.TransferFrom(&_MockERC20.TransactOpts, from, to, amount)
}

// MockERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockERC20 contract.
type MockERC20ApprovalIterator struct {
	Event *MockERC20Approval // Event containing the contract specifics and raw log

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
func (it *MockERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC20Approval)
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
		it.Event = new(MockERC20Approval)
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
func (it *MockERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC20Approval represents a Approval event raised by the MockERC20 contract.
type MockERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockERC20 *MockERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MockERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MockERC20ApprovalIterator{contract: _MockERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockERC20 *MockERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC20Approval)
				if err := _MockERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MockERC20 *MockERC20Filterer) ParseApproval(log types.Log) (*MockERC20Approval, error) {
	event := new(MockERC20Approval)
	if err := _MockERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockERC20 contract.
type MockERC20TransferIterator struct {
	Event *MockERC20Transfer // Event containing the contract specifics and raw log

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
func (it *MockERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC20Transfer)
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
		it.Event = new(MockERC20Transfer)
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
func (it *MockERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC20Transfer represents a Transfer event raised by the MockERC20 contract.
type MockERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockERC20 *MockERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockERC20TransferIterator{contract: _MockERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockERC20 *MockERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC20Transfer)
				if err := _MockERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MockERC20 *MockERC20Filterer) ParseTransfer(log types.Log) (*MockERC20Transfer, error) {
	event := new(MockERC20Transfer)
	if err := _MockERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
