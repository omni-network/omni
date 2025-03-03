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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162000b2e38038062000b2e833981016040819052620000339162000166565b81604051602001620000469190620001cc565b60405160208183030381529060405281604051602001620000689190620001cc565b60408051601f19818403018152919052600362000086838262000282565b50600462000095828262000282565b50505050506200034e565b634e487b7160e01b5f52604160045260245ffd5b5f5b83811015620000d0578181015183820152602001620000b6565b50505f910152565b5f82601f830112620000e8575f80fd5b81516001600160401b0380821115620001055762000105620000a0565b604051601f8301601f19908116603f01168101908282118183101715620001305762000130620000a0565b8160405283815286602085880101111562000149575f80fd5b6200015c846020830160208901620000b4565b9695505050505050565b5f806040838503121562000178575f80fd5b82516001600160401b03808211156200018f575f80fd5b6200019d86838701620000d8565b93506020850151915080821115620001b3575f80fd5b50620001c285828601620000d8565b9150509250929050565b5f8251620001df818460208701620000b4565b6620284d6f636b2960c81b920191825250600701919050565b600181811c908216806200020d57607f821691505b6020821081036200022c57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200027d57805f5260205f20601f840160051c81016020851015620002595750805b601f840160051c820191505b818110156200027a575f815560010162000265565b50505b505050565b81516001600160401b038111156200029e576200029e620000a0565b620002b681620002af8454620001f8565b8462000232565b602080601f831160018114620002ec575f8415620002d45750858301515b5f19600386901b1c1916600185901b17855562000346565b5f85815260208120601f198616915b828110156200031c57888601518255948401946001909101908401620002fb565b50858210156200033a57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b6107d2806200035c5f395ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c806340c10f191161006e57806340c10f191461011f57806370a082311461013457806395d89b411461015c5780639dc29fac14610164578063a9059cbb14610177578063dd62ed3e1461018a575f80fd5b806306fdde03146100aa578063095ea7b3146100c857806318160ddd146100eb57806323b872dd146100fd578063313ce56714610110575b5f80fd5b6100b26101c2565b6040516100bf919061062c565b60405180910390f35b6100db6100d6366004610693565b610252565b60405190151581526020016100bf565b6002545b6040519081526020016100bf565b6100db61010b3660046106bb565b61026b565b604051601281526020016100bf565b61013261012d366004610693565b61028e565b005b6100ef6101423660046106f4565b6001600160a01b03165f9081526020819052604090205490565b6100b261029c565b610132610172366004610693565b6102ab565b6100db610185366004610693565b6102d0565b6100ef610198366004610714565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101d190610745565b80601f01602080910402602001604051908101604052809291908181526020018280546101fd90610745565b80156102485780601f1061021f57610100808354040283529160200191610248565b820191905f5260205f20905b81548152906001019060200180831161022b57829003601f168201915b5050505050905090565b5f3361025f8185856102dd565b60019150505b92915050565b5f336102788582856102ef565b61028385858561036f565b506001949350505050565b61029882826103cc565b5050565b6060600480546101d190610745565b6001600160a01b03821633146102c6576102c68233836102ef565b6102988282610400565b5f3361025f81858561036f565b6102ea8383836001610434565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f198114610369578181101561035b57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b61036984848484035f610434565b50505050565b6001600160a01b03831661039857604051634b637e8f60e11b81525f6004820152602401610352565b6001600160a01b0382166103c15760405163ec442f0560e01b81525f6004820152602401610352565b6102ea838383610506565b6001600160a01b0382166103f55760405163ec442f0560e01b81525f6004820152602401610352565b6102985f8383610506565b6001600160a01b03821661042957604051634b637e8f60e11b81525f6004820152602401610352565b610298825f83610506565b6001600160a01b03841661045d5760405163e602df0560e01b81525f6004820152602401610352565b6001600160a01b03831661048657604051634a1406b160e11b81525f6004820152602401610352565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561036957826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104f891815260200190565b60405180910390a350505050565b6001600160a01b038316610530578060025f828254610525919061077d565b909155506105a09050565b6001600160a01b0383165f90815260208190526040902054818110156105825760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610352565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b0382166105bc576002805482900390556105da565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161061f91815260200190565b60405180910390a3505050565b5f602080835283518060208501525f5b818110156106585785810183015185820160400152820161063c565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461068e575f80fd5b919050565b5f80604083850312156106a4575f80fd5b6106ad83610678565b946020939093013593505050565b5f805f606084860312156106cd575f80fd5b6106d684610678565b92506106e460208501610678565b9150604084013590509250925092565b5f60208284031215610704575f80fd5b61070d82610678565b9392505050565b5f8060408385031215610725575f80fd5b61072e83610678565b915061073c60208401610678565b90509250929050565b600181811c9082168061075957607f821691505b60208210810361077757634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561026557634e487b7160e01b5f52601160045260245ffdfea2646970667358221220a5e067909b323fc6b3e5429e78358bf5bb5abc6b3f86d6d7d0c54518ac09d3d364736f6c63430008180033",
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
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockERC20 *MockERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockERC20 *MockERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Approve(&_MockERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockERC20 *MockERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Approve(&_MockERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_MockERC20 *MockERC20Transactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_MockERC20 *MockERC20Session) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Burn(&_MockERC20.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_MockERC20 *MockERC20TransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Burn(&_MockERC20.TransactOpts, from, amount)
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
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockERC20 *MockERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockERC20 *MockERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Transfer(&_MockERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockERC20 *MockERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.Transfer(&_MockERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockERC20 *MockERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockERC20 *MockERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.TransferFrom(&_MockERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockERC20 *MockERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockERC20.Contract.TransferFrom(&_MockERC20.TransactOpts, from, to, value)
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
