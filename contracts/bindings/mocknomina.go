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

// MockNominaMetaData contains all meta data concerning the MockNomina contract.
var MockNominaMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_omni\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONVERSION_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OMNI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowanceOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllowanceUnderflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Permit2AllowanceIsFixedAtInfinity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PermitExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSupplyOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610b50380380610b5083398101604081905261002e91610066565b6001600160a01b0381166100555760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0316608052610093565b5f60208284031215610076575f80fd5b81516001600160a01b038116811461008c575f80fd5b9392505050565b608051610a9e6100b25f395f818161010f01526104f10152610a9e5ff3fe608060405234801561000f575f80fd5b5060043610610106575f3560e01c806340c10f191161009e5780637ecebe001161006e5780637ecebe00146102d457806395d89b41146102f9578063a9059cbb14610318578063d505accf1461032b578063dd62ed3e1461033e575f80fd5b806340c10f191461027457806342966c681461028957806367c6e39c1461029c57806370a08231146102af575f80fd5b806323b872dd116100d957806323b872dd146101b65780632c8bff31146101c9578063313ce567146101e35780633644e515146101ea575f80fd5b8063063bdf281461010a57806306fdde031461014e578063095ea7b31461017957806318160ddd1461019c575b5f80fd5b6101317f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6040805180820190915260068152654e6f6d696e6160d01b60208201525b60405161014591906108bd565b61018c610187366004610924565b610351565b6040519015158152602001610145565b6805345cdf77eb68f44c545b604051908152602001610145565b61018c6101c436600461094c565b6103d1565b6101d1604b81565b60405160ff9091168152602001610145565b60126101d1565b604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81527fc72733118dabad3698b4044c2dc83c8c688bd907b50ed9d09d93a263878bf51860208201527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc69181019190915246606082015230608082015260a090206101a8565b610287610282366004610924565b61048d565b005b610287610297366004610985565b61049b565b6102876102aa366004610924565b6104b2565b6101a86102bd36600461099c565b6387a211a2600c9081525f91909152602090205490565b6101a86102e236600461099c565b6338377508600c9081525f91909152602090205490565b6040805180820190915260038152624e4f4d60e81b602082015261016c565b61018c610326366004610924565b61052f565b6102876103393660046109bc565b610593565b6101a861034c366004610a29565b610756565b5f6001600160a01b0383166e22d473030f116ddee9f6b43ac78ba3188219151761038257633f68539a5f526004601cfd5b82602052637f5e9f20600c52335f52816034600c2055815f52602c5160601c337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560205fa35060015b92915050565b5f8360601b6e22d473030f116ddee9f6b43ac78ba333146104265733602052637f5e9f208117600c526034600c208054801915610423578085111561041d576313be252b5f526004601cfd5b84810382555b50505b6387a211a28117600c526020600c2080548085111561044c5763f4d678b85f526004601cfd5b84810382555050835f526020600c208381540181555082602052600c5160601c8160601c5f80516020610a7e833981519152602080a3505060019392505050565b610497828261079a565b5050565b805f036104a55750565b6104af3382610803565b50565b805f036104bd575050565b6001600160a01b0382166104e45760405163d92e233d60e01b815260040160405180910390fd5b61051b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163361dead84610864565b6104978261052a604b84610a5a565b61079a565b5f6387a211a2600c52335f526020600c208054808411156105575763f4d678b85f526004601cfd5b83810382555050825f526020600c208281540181555081602052600c5160601c335f80516020610a7e833981519152602080a350600192915050565b6001600160a01b0386166e22d473030f116ddee9f6b43ac78ba318851915176105c357633f68539a5f526004601cfd5b7fc72733118dabad3698b4044c2dc83c8c688bd907b50ed9d09d93a263878bf5187fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64286101561061a57631a15a3cc5f526004601cfd5b6040518960601b60601c99508860601b60601c985065383775081901600e52895f526020600c2080547f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f835284602084015283604084015246606084015230608084015260a08320602e527f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c983528b60208401528a60408401528960608401528060808401528860a084015260c08320604e526042602c205f528760ff16602052866040528560605260208060805f60015afa8c3d51146107025763ddafbaef5f526004601cfd5b0190556303faf4f960a51b89176040526034602c20889055888a7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925602060608501a360405250505f60605250505050505050565b5f6e22d473030f116ddee9f6b43ac78ba2196001600160a01b0383160161077f57505f196103cb565b50602052637f5e9f20600c9081525f91909152603490205490565b6805345cdf77eb68f44c54818101818110156107bd5763e5cfe9575f526004601cfd5b806805345cdf77eb68f44c5550506387a211a2600c52815f526020600c208181540181555080602052600c5160601c5f5f80516020610a7e833981519152602080a35050565b6387a211a2600c52815f526020600c2080548083111561082a5763f4d678b85f526004601cfd5b82900390556805345cdf77eb68f44c805482900390555f8181526001600160a01b0383165f80516020610a7e833981519152602083a35050565b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f5114166108af57803d873b1517106108af57637939f4245f526004601cfd5b505f60605260405250505050565b5f602080835283518060208501525f5b818110156108e9578581018301518582016040015282016108cd565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461091f575f80fd5b919050565b5f8060408385031215610935575f80fd5b61093e83610909565b946020939093013593505050565b5f805f6060848603121561095e575f80fd5b61096784610909565b925061097560208501610909565b9150604084013590509250925092565b5f60208284031215610995575f80fd5b5035919050565b5f602082840312156109ac575f80fd5b6109b582610909565b9392505050565b5f805f805f805f60e0888a0312156109d2575f80fd5b6109db88610909565b96506109e960208901610909565b95506040880135945060608801359350608088013560ff81168114610a0c575f80fd5b9699959850939692959460a0840135945060c09093013592915050565b5f8060408385031215610a3a575f80fd5b610a4383610909565b9150610a5160208401610909565b90509250929050565b80820281158282048414176103cb57634e487b7160e01b5f52601160045260245ffdfeddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
}

// MockNominaABI is the input ABI used to generate the binding from.
// Deprecated: Use MockNominaMetaData.ABI instead.
var MockNominaABI = MockNominaMetaData.ABI

// MockNominaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockNominaMetaData.Bin instead.
var MockNominaBin = MockNominaMetaData.Bin

// DeployMockNomina deploys a new Ethereum contract, binding an instance of MockNomina to it.
func DeployMockNomina(auth *bind.TransactOpts, backend bind.ContractBackend, _omni common.Address) (common.Address, *types.Transaction, *MockNomina, error) {
	parsed, err := MockNominaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockNominaBin), backend, _omni)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockNomina{MockNominaCaller: MockNominaCaller{contract: contract}, MockNominaTransactor: MockNominaTransactor{contract: contract}, MockNominaFilterer: MockNominaFilterer{contract: contract}}, nil
}

// MockNomina is an auto generated Go binding around an Ethereum contract.
type MockNomina struct {
	MockNominaCaller     // Read-only binding to the contract
	MockNominaTransactor // Write-only binding to the contract
	MockNominaFilterer   // Log filterer for contract events
}

// MockNominaCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockNominaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNominaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockNominaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNominaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockNominaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNominaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockNominaSession struct {
	Contract     *MockNomina       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockNominaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockNominaCallerSession struct {
	Contract *MockNominaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MockNominaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockNominaTransactorSession struct {
	Contract     *MockNominaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MockNominaRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockNominaRaw struct {
	Contract *MockNomina // Generic contract binding to access the raw methods on
}

// MockNominaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockNominaCallerRaw struct {
	Contract *MockNominaCaller // Generic read-only contract binding to access the raw methods on
}

// MockNominaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockNominaTransactorRaw struct {
	Contract *MockNominaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockNomina creates a new instance of MockNomina, bound to a specific deployed contract.
func NewMockNomina(address common.Address, backend bind.ContractBackend) (*MockNomina, error) {
	contract, err := bindMockNomina(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockNomina{MockNominaCaller: MockNominaCaller{contract: contract}, MockNominaTransactor: MockNominaTransactor{contract: contract}, MockNominaFilterer: MockNominaFilterer{contract: contract}}, nil
}

// NewMockNominaCaller creates a new read-only instance of MockNomina, bound to a specific deployed contract.
func NewMockNominaCaller(address common.Address, caller bind.ContractCaller) (*MockNominaCaller, error) {
	contract, err := bindMockNomina(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockNominaCaller{contract: contract}, nil
}

// NewMockNominaTransactor creates a new write-only instance of MockNomina, bound to a specific deployed contract.
func NewMockNominaTransactor(address common.Address, transactor bind.ContractTransactor) (*MockNominaTransactor, error) {
	contract, err := bindMockNomina(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockNominaTransactor{contract: contract}, nil
}

// NewMockNominaFilterer creates a new log filterer instance of MockNomina, bound to a specific deployed contract.
func NewMockNominaFilterer(address common.Address, filterer bind.ContractFilterer) (*MockNominaFilterer, error) {
	contract, err := bindMockNomina(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockNominaFilterer{contract: contract}, nil
}

// bindMockNomina binds a generic wrapper to an already deployed contract.
func bindMockNomina(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockNominaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockNomina *MockNominaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockNomina.Contract.MockNominaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockNomina *MockNominaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockNomina.Contract.MockNominaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockNomina *MockNominaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockNomina.Contract.MockNominaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockNomina *MockNominaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockNomina.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockNomina *MockNominaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockNomina.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockNomina *MockNominaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockNomina.Contract.contract.Transact(opts, method, params...)
}

// CONVERSIONRATE is a free data retrieval call binding the contract method 0x2c8bff31.
//
// Solidity: function CONVERSION_RATE() view returns(uint8)
func (_MockNomina *MockNominaCaller) CONVERSIONRATE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "CONVERSION_RATE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CONVERSIONRATE is a free data retrieval call binding the contract method 0x2c8bff31.
//
// Solidity: function CONVERSION_RATE() view returns(uint8)
func (_MockNomina *MockNominaSession) CONVERSIONRATE() (uint8, error) {
	return _MockNomina.Contract.CONVERSIONRATE(&_MockNomina.CallOpts)
}

// CONVERSIONRATE is a free data retrieval call binding the contract method 0x2c8bff31.
//
// Solidity: function CONVERSION_RATE() view returns(uint8)
func (_MockNomina *MockNominaCallerSession) CONVERSIONRATE() (uint8, error) {
	return _MockNomina.Contract.CONVERSIONRATE(&_MockNomina.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_MockNomina *MockNominaCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_MockNomina *MockNominaSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MockNomina.Contract.DOMAINSEPARATOR(&_MockNomina.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_MockNomina *MockNominaCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MockNomina.Contract.DOMAINSEPARATOR(&_MockNomina.CallOpts)
}

// OMNI is a free data retrieval call binding the contract method 0x063bdf28.
//
// Solidity: function OMNI() view returns(address)
func (_MockNomina *MockNominaCaller) OMNI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "OMNI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OMNI is a free data retrieval call binding the contract method 0x063bdf28.
//
// Solidity: function OMNI() view returns(address)
func (_MockNomina *MockNominaSession) OMNI() (common.Address, error) {
	return _MockNomina.Contract.OMNI(&_MockNomina.CallOpts)
}

// OMNI is a free data retrieval call binding the contract method 0x063bdf28.
//
// Solidity: function OMNI() view returns(address)
func (_MockNomina *MockNominaCallerSession) OMNI() (common.Address, error) {
	return _MockNomina.Contract.OMNI(&_MockNomina.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_MockNomina *MockNominaCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_MockNomina *MockNominaSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockNomina.Contract.Allowance(&_MockNomina.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_MockNomina *MockNominaCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockNomina.Contract.Allowance(&_MockNomina.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_MockNomina *MockNominaCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_MockNomina *MockNominaSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MockNomina.Contract.BalanceOf(&_MockNomina.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_MockNomina *MockNominaCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MockNomina.Contract.BalanceOf(&_MockNomina.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockNomina *MockNominaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockNomina *MockNominaSession) Decimals() (uint8, error) {
	return _MockNomina.Contract.Decimals(&_MockNomina.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockNomina *MockNominaCallerSession) Decimals() (uint8, error) {
	return _MockNomina.Contract.Decimals(&_MockNomina.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_MockNomina *MockNominaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_MockNomina *MockNominaSession) Name() (string, error) {
	return _MockNomina.Contract.Name(&_MockNomina.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_MockNomina *MockNominaCallerSession) Name() (string, error) {
	return _MockNomina.Contract.Name(&_MockNomina.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_MockNomina *MockNominaCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_MockNomina *MockNominaSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MockNomina.Contract.Nonces(&_MockNomina.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_MockNomina *MockNominaCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MockNomina.Contract.Nonces(&_MockNomina.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_MockNomina *MockNominaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_MockNomina *MockNominaSession) Symbol() (string, error) {
	return _MockNomina.Contract.Symbol(&_MockNomina.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_MockNomina *MockNominaCallerSession) Symbol() (string, error) {
	return _MockNomina.Contract.Symbol(&_MockNomina.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_MockNomina *MockNominaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockNomina.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_MockNomina *MockNominaSession) TotalSupply() (*big.Int, error) {
	return _MockNomina.Contract.TotalSupply(&_MockNomina.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_MockNomina *MockNominaCallerSession) TotalSupply() (*big.Int, error) {
	return _MockNomina.Contract.TotalSupply(&_MockNomina.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockNomina *MockNominaTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockNomina *MockNominaSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Approve(&_MockNomina.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockNomina *MockNominaTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Approve(&_MockNomina.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MockNomina *MockNominaTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MockNomina *MockNominaSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Burn(&_MockNomina.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MockNomina *MockNominaTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Burn(&_MockNomina.TransactOpts, amount)
}

// Convert is a paid mutator transaction binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address to, uint256 amount) returns()
func (_MockNomina *MockNominaTransactor) Convert(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.contract.Transact(opts, "convert", to, amount)
}

// Convert is a paid mutator transaction binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address to, uint256 amount) returns()
func (_MockNomina *MockNominaSession) Convert(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Convert(&_MockNomina.TransactOpts, to, amount)
}

// Convert is a paid mutator transaction binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address to, uint256 amount) returns()
func (_MockNomina *MockNominaTransactorSession) Convert(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Convert(&_MockNomina.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockNomina *MockNominaTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockNomina *MockNominaSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Mint(&_MockNomina.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockNomina *MockNominaTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Mint(&_MockNomina.TransactOpts, to, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MockNomina *MockNominaTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MockNomina.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MockNomina *MockNominaSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MockNomina.Contract.Permit(&_MockNomina.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MockNomina *MockNominaTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MockNomina.Contract.Permit(&_MockNomina.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockNomina *MockNominaTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockNomina *MockNominaSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Transfer(&_MockNomina.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockNomina *MockNominaTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.Transfer(&_MockNomina.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockNomina *MockNominaTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockNomina *MockNominaSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.TransferFrom(&_MockNomina.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockNomina *MockNominaTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockNomina.Contract.TransferFrom(&_MockNomina.TransactOpts, from, to, amount)
}

// MockNominaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockNomina contract.
type MockNominaApprovalIterator struct {
	Event *MockNominaApproval // Event containing the contract specifics and raw log

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
func (it *MockNominaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNominaApproval)
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
		it.Event = new(MockNominaApproval)
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
func (it *MockNominaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNominaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNominaApproval represents a Approval event raised by the MockNomina contract.
type MockNominaApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_MockNomina *MockNominaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MockNominaApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockNomina.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MockNominaApprovalIterator{contract: _MockNomina.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_MockNomina *MockNominaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockNominaApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockNomina.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNominaApproval)
				if err := _MockNomina.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_MockNomina *MockNominaFilterer) ParseApproval(log types.Log) (*MockNominaApproval, error) {
	event := new(MockNominaApproval)
	if err := _MockNomina.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockNominaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockNomina contract.
type MockNominaTransferIterator struct {
	Event *MockNominaTransfer // Event containing the contract specifics and raw log

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
func (it *MockNominaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNominaTransfer)
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
		it.Event = new(MockNominaTransfer)
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
func (it *MockNominaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNominaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNominaTransfer represents a Transfer event raised by the MockNomina contract.
type MockNominaTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_MockNomina *MockNominaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockNominaTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockNomina.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockNominaTransferIterator{contract: _MockNomina.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_MockNomina *MockNominaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockNominaTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockNomina.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNominaTransfer)
				if err := _MockNomina.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_MockNomina *MockNominaFilterer) ParseTransfer(log types.Log) (*MockNominaTransfer, error) {
	event := new(MockNominaTransfer)
	if err := _MockNomina.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
