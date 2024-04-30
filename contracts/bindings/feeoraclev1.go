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

// IFeeOracleV1ChainFeeParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV1ChainFeeParams struct {
	ChainId      uint64
	GasPrice     *big.Int
	ToNativeRate *big.Int
}

// FeeOracleV1MetaData contains all meta data concerning the FeeOracleV1 contract.
var FeeOracleV1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONVERSION_RATE_DENOM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bulkSetFeeParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV1.ChainFeeParams[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toNativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasPriceOn\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"baseGasLimit_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"protocolFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV1.ChainFeeParams[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toNativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setToNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toNativeRate\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BaseGasLimitSet\",\"inputs\":[{\"name\":\"baseGasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasPriceSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolFeeSet\",\"inputs\":[{\"name\":\"protocolFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ToNativeRateSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100dd565b600054610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100db576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610c7a806100ec6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638f9d6ace11610097578063b0e21e8a11610066578063b0e21e8a146101f9578063d070415714610202578063ee590a5314610222578063f2fde38b1461023557600080fd5b80638f9d6ace146101b657806390b97cf0146101c057806398563b03146101d3578063a34e7abb146101e657600080fd5b80638b7bfd70116100d35780638b7bfd701461012a5780638da5cb5b1461015d5780638dd9523c146101785780638df66e341461018b57600080fd5b80633fde7da9146100fa578063715018a61461010f578063787dce3d14610117575b600080fd5b61010d610108366004610963565b610248565b005b61010d61025e565b61010d6101253660046109a4565b610272565b61014a6101383660046109d9565b60686020526000908152604090205481565b6040519081526020015b60405180910390f35b6033546040516001600160a01b039091168152602001610154565b61014a6101863660046109fb565b610286565b60655461019e906001600160401b031681565b6040516001600160401b039091168152602001610154565b61014a620f424081565b61010d6101ce366004610aa5565b61039e565b61010d6101e1366004610b13565b6104d1565b61010d6101f4366004610b13565b6104e3565b61014a60665481565b61014a6102103660046109d9565b60676020526000908152604090205481565b61010d6102303660046109d9565b6104f5565b61010d610243366004610b3d565b610506565b61025061057c565b61025a82826105d6565b5050565b61026661057c565b610270600061063f565b565b61027a61057c565b61028381610691565b50565b6001600160401b038416600090815260676020526040812054158015906102c457506001600160401b03851660009081526068602052604090205415155b6103155760405162461bcd60e51b815260206004820152601a60248201527f4665654f7261636c6556313a206e6f2066656520706172616d7300000000000060448201526064015b60405180910390fd5b6001600160401b0385166000908152606860209081526040808320546067909252822054620f42409161034791610b6e565b6103519190610b8b565b9050610366816001600160401b038516610b6e565b60655461037d9083906001600160401b0316610b6e565b60665461038a9190610bad565b6103949190610bad565b9695505050505050565b600054610100900460ff16158080156103be5750600054600160ff909116105b806103d85750303b1580156103d8575060005460ff166001145b61043b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161030c565b6000805460ff19166001179055801561045e576000805461ff0019166101001790555b6104678661063f565b610470856106cd565b61047984610691565b61048383836105d6565b80156104c9576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b6104d961057c565b61025a828261071c565b6104eb61057c565b61025a828261081e565b6104fd61057c565b610283816106cd565b61050e61057c565b6001600160a01b0381166105735760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161030c565b6102838161063f565b6033546001600160a01b031633146102705760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161030c565b60005b8181101561063a5760008383838181106105f5576105f5610bc0565b90506060020180360381019061060b9190610bd6565b905061061f8160000151826020015161071c565b6106318160000151826040015161081e565b506001016105d9565b505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60668190556040518181527fdb5aafdb29539329e37d4e3ee869bc4031941fd55a5dfc92824fbe34b204e30d906020015b60405180910390a150565b6065805467ffffffffffffffff19166001600160401b0383169081179091556040519081527f6185fbe062d94552cf644f5cb643f583db7b2e7e66fdc4b4c75ff8876a257ba6906020016106c2565b6000811161076c5760405162461bcd60e51b815260206004820152601e60248201527f4665654f7261636c6556313a206e6f207a65726f206761732070726963650000604482015260640161030c565b816001600160401b03166000036107c55760405162461bcd60e51b815260206004820152601d60248201527f4665654f7261636c6556313a206e6f207a65726f20636861696e206964000000604482015260640161030c565b6001600160401b038216600081815260676020908152604091829020849055815192835282018390527f3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff91015b60405180910390a15050565b6000811161086e5760405162461bcd60e51b815260206004820152601960248201527f4665654f7261636c6556313a206e6f207a65726f207261746500000000000000604482015260640161030c565b816001600160401b03166000036108c75760405162461bcd60e51b815260206004820152601d60248201527f4665654f7261636c6556313a206e6f207a65726f20636861696e206964000000604482015260640161030c565b6001600160401b038216600081815260686020908152604091829020849055815192835282018390527f4b4594c9f06af25bc504eead96f7f0eaa3f1577f8d9b075b236520ec712e13089101610812565b60008083601f84011261092a57600080fd5b5081356001600160401b0381111561094157600080fd5b60208301915083602060608302850101111561095c57600080fd5b9250929050565b6000806020838503121561097657600080fd5b82356001600160401b0381111561098c57600080fd5b61099885828601610918565b90969095509350505050565b6000602082840312156109b657600080fd5b5035919050565b80356001600160401b03811681146109d457600080fd5b919050565b6000602082840312156109eb57600080fd5b6109f4826109bd565b9392505050565b60008060008060608587031215610a1157600080fd5b610a1a856109bd565b935060208501356001600160401b0380821115610a3657600080fd5b818701915087601f830112610a4a57600080fd5b813581811115610a5957600080fd5b886020828501011115610a6b57600080fd5b602083019550809450505050610a83604086016109bd565b905092959194509250565b80356001600160a01b03811681146109d457600080fd5b600080600080600060808688031215610abd57600080fd5b610ac686610a8e565b9450610ad4602087016109bd565b93506040860135925060608601356001600160401b03811115610af657600080fd5b610b0288828901610918565b969995985093965092949392505050565b60008060408385031215610b2657600080fd5b610b2f836109bd565b946020939093013593505050565b600060208284031215610b4f57600080fd5b6109f482610a8e565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610b8557610b85610b58565b92915050565b600082610ba857634e487b7160e01b600052601260045260246000fd5b500490565b80820180821115610b8557610b85610b58565b634e487b7160e01b600052603260045260246000fd5b600060608284031215610be857600080fd5b604051606081018181106001600160401b0382111715610c1857634e487b7160e01b600052604160045260246000fd5b604052610c24836109bd565b81526020830135602082015260408301356040820152809150509291505056fea2646970667358221220cf6018675176cdb7da93afc7ea3a51ca62d965d53f2c129f0b18d2d8b0c274f664736f6c63430008180033",
}

// FeeOracleV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeOracleV1MetaData.ABI instead.
var FeeOracleV1ABI = FeeOracleV1MetaData.ABI

// FeeOracleV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeOracleV1MetaData.Bin instead.
var FeeOracleV1Bin = FeeOracleV1MetaData.Bin

// DeployFeeOracleV1 deploys a new Ethereum contract, binding an instance of FeeOracleV1 to it.
func DeployFeeOracleV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FeeOracleV1, error) {
	parsed, err := FeeOracleV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeOracleV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeOracleV1{FeeOracleV1Caller: FeeOracleV1Caller{contract: contract}, FeeOracleV1Transactor: FeeOracleV1Transactor{contract: contract}, FeeOracleV1Filterer: FeeOracleV1Filterer{contract: contract}}, nil
}

// FeeOracleV1 is an auto generated Go binding around an Ethereum contract.
type FeeOracleV1 struct {
	FeeOracleV1Caller     // Read-only binding to the contract
	FeeOracleV1Transactor // Write-only binding to the contract
	FeeOracleV1Filterer   // Log filterer for contract events
}

// FeeOracleV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeeOracleV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeOracleV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeOracleV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeOracleV1Session struct {
	Contract     *FeeOracleV1      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeOracleV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeOracleV1CallerSession struct {
	Contract *FeeOracleV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FeeOracleV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeOracleV1TransactorSession struct {
	Contract     *FeeOracleV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FeeOracleV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeeOracleV1Raw struct {
	Contract *FeeOracleV1 // Generic contract binding to access the raw methods on
}

// FeeOracleV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeOracleV1CallerRaw struct {
	Contract *FeeOracleV1Caller // Generic read-only contract binding to access the raw methods on
}

// FeeOracleV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeOracleV1TransactorRaw struct {
	Contract *FeeOracleV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeOracleV1 creates a new instance of FeeOracleV1, bound to a specific deployed contract.
func NewFeeOracleV1(address common.Address, backend bind.ContractBackend) (*FeeOracleV1, error) {
	contract, err := bindFeeOracleV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1{FeeOracleV1Caller: FeeOracleV1Caller{contract: contract}, FeeOracleV1Transactor: FeeOracleV1Transactor{contract: contract}, FeeOracleV1Filterer: FeeOracleV1Filterer{contract: contract}}, nil
}

// NewFeeOracleV1Caller creates a new read-only instance of FeeOracleV1, bound to a specific deployed contract.
func NewFeeOracleV1Caller(address common.Address, caller bind.ContractCaller) (*FeeOracleV1Caller, error) {
	contract, err := bindFeeOracleV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1Caller{contract: contract}, nil
}

// NewFeeOracleV1Transactor creates a new write-only instance of FeeOracleV1, bound to a specific deployed contract.
func NewFeeOracleV1Transactor(address common.Address, transactor bind.ContractTransactor) (*FeeOracleV1Transactor, error) {
	contract, err := bindFeeOracleV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1Transactor{contract: contract}, nil
}

// NewFeeOracleV1Filterer creates a new log filterer instance of FeeOracleV1, bound to a specific deployed contract.
func NewFeeOracleV1Filterer(address common.Address, filterer bind.ContractFilterer) (*FeeOracleV1Filterer, error) {
	contract, err := bindFeeOracleV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1Filterer{contract: contract}, nil
}

// bindFeeOracleV1 binds a generic wrapper to an already deployed contract.
func bindFeeOracleV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeOracleV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeOracleV1 *FeeOracleV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeOracleV1.Contract.FeeOracleV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeOracleV1 *FeeOracleV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.FeeOracleV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeOracleV1 *FeeOracleV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.FeeOracleV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeOracleV1 *FeeOracleV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeOracleV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeOracleV1 *FeeOracleV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeOracleV1 *FeeOracleV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.contract.Transact(opts, method, params...)
}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) CONVERSIONRATEDENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "CONVERSION_RATE_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) CONVERSIONRATEDENOM() (*big.Int, error) {
	return _FeeOracleV1.Contract.CONVERSIONRATEDENOM(&_FeeOracleV1.CallOpts)
}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) CONVERSIONRATEDENOM() (*big.Int, error) {
	return _FeeOracleV1.Contract.CONVERSIONRATEDENOM(&_FeeOracleV1.CallOpts)
}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1Caller) BaseGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "baseGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1Session) BaseGasLimit() (uint64, error) {
	return _FeeOracleV1.Contract.BaseGasLimit(&_FeeOracleV1.CallOpts)
}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1CallerSession) BaseGasLimit() (uint64, error) {
	return _FeeOracleV1.Contract.BaseGasLimit(&_FeeOracleV1.CallOpts)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes , uint64 gasLimit) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) FeeFor(opts *bind.CallOpts, destChainId uint64, arg1 []byte, gasLimit uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "feeFor", destChainId, arg1, gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes , uint64 gasLimit) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) FeeFor(destChainId uint64, arg1 []byte, gasLimit uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.FeeFor(&_FeeOracleV1.CallOpts, destChainId, arg1, gasLimit)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes , uint64 gasLimit) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) FeeFor(destChainId uint64, arg1 []byte, gasLimit uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.FeeFor(&_FeeOracleV1.CallOpts, destChainId, arg1, gasLimit)
}

// GasPriceOn is a free data retrieval call binding the contract method 0xd0704157.
//
// Solidity: function gasPriceOn(uint64 ) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) GasPriceOn(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "gasPriceOn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPriceOn is a free data retrieval call binding the contract method 0xd0704157.
//
// Solidity: function gasPriceOn(uint64 ) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) GasPriceOn(arg0 uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.GasPriceOn(&_FeeOracleV1.CallOpts, arg0)
}

// GasPriceOn is a free data retrieval call binding the contract method 0xd0704157.
//
// Solidity: function gasPriceOn(uint64 ) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) GasPriceOn(arg0 uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.GasPriceOn(&_FeeOracleV1.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV1 *FeeOracleV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV1 *FeeOracleV1Session) Owner() (common.Address, error) {
	return _FeeOracleV1.Contract.Owner(&_FeeOracleV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV1 *FeeOracleV1CallerSession) Owner() (common.Address, error) {
	return _FeeOracleV1.Contract.Owner(&_FeeOracleV1.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) ProtocolFee() (*big.Int, error) {
	return _FeeOracleV1.Contract.ProtocolFee(&_FeeOracleV1.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) ProtocolFee() (*big.Int, error) {
	return _FeeOracleV1.Contract.ProtocolFee(&_FeeOracleV1.CallOpts)
}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 ) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) ToNativeRate(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "toNativeRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 ) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) ToNativeRate(arg0 uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.ToNativeRate(&_FeeOracleV1.CallOpts, arg0)
}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 ) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) ToNativeRate(arg0 uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.ToNativeRate(&_FeeOracleV1.CallOpts, arg0)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0x3fde7da9.
//
// Solidity: function bulkSetFeeParams((uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) BulkSetFeeParams(opts *bind.TransactOpts, params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "bulkSetFeeParams", params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0x3fde7da9.
//
// Solidity: function bulkSetFeeParams((uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1Session) BulkSetFeeParams(params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.BulkSetFeeParams(&_FeeOracleV1.TransactOpts, params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0x3fde7da9.
//
// Solidity: function bulkSetFeeParams((uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) BulkSetFeeParams(params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.BulkSetFeeParams(&_FeeOracleV1.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x90b97cf0.
//
// Solidity: function initialize(address owner_, uint64 baseGasLimit_, uint256 protocolFee_, (uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, baseGasLimit_ uint64, protocolFee_ *big.Int, params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "initialize", owner_, baseGasLimit_, protocolFee_, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x90b97cf0.
//
// Solidity: function initialize(address owner_, uint64 baseGasLimit_, uint256 protocolFee_, (uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1Session) Initialize(owner_ common.Address, baseGasLimit_ uint64, protocolFee_ *big.Int, params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.Initialize(&_FeeOracleV1.TransactOpts, owner_, baseGasLimit_, protocolFee_, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x90b97cf0.
//
// Solidity: function initialize(address owner_, uint64 baseGasLimit_, uint256 protocolFee_, (uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) Initialize(owner_ common.Address, baseGasLimit_ uint64, protocolFee_ *big.Int, params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.Initialize(&_FeeOracleV1.TransactOpts, owner_, baseGasLimit_, protocolFee_, params)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV1 *FeeOracleV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _FeeOracleV1.Contract.RenounceOwnership(&_FeeOracleV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeOracleV1.Contract.RenounceOwnership(&_FeeOracleV1.TransactOpts)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xee590a53.
//
// Solidity: function setBaseGasLimit(uint64 gasLimit) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetBaseGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setBaseGasLimit", gasLimit)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xee590a53.
//
// Solidity: function setBaseGasLimit(uint64 gasLimit) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetBaseGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetBaseGasLimit(&_FeeOracleV1.TransactOpts, gasLimit)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xee590a53.
//
// Solidity: function setBaseGasLimit(uint64 gasLimit) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetBaseGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetBaseGasLimit(&_FeeOracleV1.TransactOpts, gasLimit)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0x98563b03.
//
// Solidity: function setGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetGasPrice(opts *bind.TransactOpts, chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setGasPrice", chainId, gasPrice)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0x98563b03.
//
// Solidity: function setGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetGasPrice(chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetGasPrice(&_FeeOracleV1.TransactOpts, chainId, gasPrice)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0x98563b03.
//
// Solidity: function setGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetGasPrice(chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetGasPrice(&_FeeOracleV1.TransactOpts, chainId, gasPrice)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 fee) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetProtocolFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setProtocolFee", fee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 fee) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetProtocolFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetProtocolFee(&_FeeOracleV1.TransactOpts, fee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 fee) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetProtocolFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetProtocolFee(&_FeeOracleV1.TransactOpts, fee)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa34e7abb.
//
// Solidity: function setToNativeRate(uint64 chainId, uint256 rate) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetToNativeRate(opts *bind.TransactOpts, chainId uint64, rate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setToNativeRate", chainId, rate)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa34e7abb.
//
// Solidity: function setToNativeRate(uint64 chainId, uint256 rate) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetToNativeRate(chainId uint64, rate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetToNativeRate(&_FeeOracleV1.TransactOpts, chainId, rate)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa34e7abb.
//
// Solidity: function setToNativeRate(uint64 chainId, uint256 rate) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetToNativeRate(chainId uint64, rate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetToNativeRate(&_FeeOracleV1.TransactOpts, chainId, rate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV1 *FeeOracleV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.TransferOwnership(&_FeeOracleV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.TransferOwnership(&_FeeOracleV1.TransactOpts, newOwner)
}

// FeeOracleV1BaseGasLimitSetIterator is returned from FilterBaseGasLimitSet and is used to iterate over the raw logs and unpacked data for BaseGasLimitSet events raised by the FeeOracleV1 contract.
type FeeOracleV1BaseGasLimitSetIterator struct {
	Event *FeeOracleV1BaseGasLimitSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV1BaseGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1BaseGasLimitSet)
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
		it.Event = new(FeeOracleV1BaseGasLimitSet)
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
func (it *FeeOracleV1BaseGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1BaseGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1BaseGasLimitSet represents a BaseGasLimitSet event raised by the FeeOracleV1 contract.
type FeeOracleV1BaseGasLimitSet struct {
	BaseGasLimit uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBaseGasLimitSet is a free log retrieval operation binding the contract event 0x6185fbe062d94552cf644f5cb643f583db7b2e7e66fdc4b4c75ff8876a257ba6.
//
// Solidity: event BaseGasLimitSet(uint64 baseGasLimit)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterBaseGasLimitSet(opts *bind.FilterOpts) (*FeeOracleV1BaseGasLimitSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "BaseGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1BaseGasLimitSetIterator{contract: _FeeOracleV1.contract, event: "BaseGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchBaseGasLimitSet is a free log subscription operation binding the contract event 0x6185fbe062d94552cf644f5cb643f583db7b2e7e66fdc4b4c75ff8876a257ba6.
//
// Solidity: event BaseGasLimitSet(uint64 baseGasLimit)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchBaseGasLimitSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1BaseGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "BaseGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1BaseGasLimitSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "BaseGasLimitSet", log); err != nil {
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

// ParseBaseGasLimitSet is a log parse operation binding the contract event 0x6185fbe062d94552cf644f5cb643f583db7b2e7e66fdc4b4c75ff8876a257ba6.
//
// Solidity: event BaseGasLimitSet(uint64 baseGasLimit)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseBaseGasLimitSet(log types.Log) (*FeeOracleV1BaseGasLimitSet, error) {
	event := new(FeeOracleV1BaseGasLimitSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "BaseGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1GasPriceSetIterator is returned from FilterGasPriceSet and is used to iterate over the raw logs and unpacked data for GasPriceSet events raised by the FeeOracleV1 contract.
type FeeOracleV1GasPriceSetIterator struct {
	Event *FeeOracleV1GasPriceSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV1GasPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1GasPriceSet)
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
		it.Event = new(FeeOracleV1GasPriceSet)
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
func (it *FeeOracleV1GasPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1GasPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1GasPriceSet represents a GasPriceSet event raised by the FeeOracleV1 contract.
type FeeOracleV1GasPriceSet struct {
	ChainId  uint64
	GasPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasPriceSet is a free log retrieval operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 chainId, uint256 gasPrice)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterGasPriceSet(opts *bind.FilterOpts) (*FeeOracleV1GasPriceSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "GasPriceSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1GasPriceSetIterator{contract: _FeeOracleV1.contract, event: "GasPriceSet", logs: logs, sub: sub}, nil
}

// WatchGasPriceSet is a free log subscription operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 chainId, uint256 gasPrice)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchGasPriceSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1GasPriceSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "GasPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1GasPriceSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "GasPriceSet", log); err != nil {
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

// ParseGasPriceSet is a log parse operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 chainId, uint256 gasPrice)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseGasPriceSet(log types.Log) (*FeeOracleV1GasPriceSet, error) {
	event := new(FeeOracleV1GasPriceSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "GasPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FeeOracleV1 contract.
type FeeOracleV1InitializedIterator struct {
	Event *FeeOracleV1Initialized // Event containing the contract specifics and raw log

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
func (it *FeeOracleV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1Initialized)
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
		it.Event = new(FeeOracleV1Initialized)
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
func (it *FeeOracleV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1Initialized represents a Initialized event raised by the FeeOracleV1 contract.
type FeeOracleV1Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*FeeOracleV1InitializedIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1InitializedIterator{contract: _FeeOracleV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FeeOracleV1Initialized) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1Initialized)
				if err := _FeeOracleV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseInitialized(log types.Log) (*FeeOracleV1Initialized, error) {
	event := new(FeeOracleV1Initialized)
	if err := _FeeOracleV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeOracleV1 contract.
type FeeOracleV1OwnershipTransferredIterator struct {
	Event *FeeOracleV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeOracleV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1OwnershipTransferred)
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
		it.Event = new(FeeOracleV1OwnershipTransferred)
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
func (it *FeeOracleV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1OwnershipTransferred represents a OwnershipTransferred event raised by the FeeOracleV1 contract.
type FeeOracleV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeOracleV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1OwnershipTransferredIterator{contract: _FeeOracleV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeOracleV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1OwnershipTransferred)
				if err := _FeeOracleV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseOwnershipTransferred(log types.Log) (*FeeOracleV1OwnershipTransferred, error) {
	event := new(FeeOracleV1OwnershipTransferred)
	if err := _FeeOracleV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1ProtocolFeeSetIterator is returned from FilterProtocolFeeSet and is used to iterate over the raw logs and unpacked data for ProtocolFeeSet events raised by the FeeOracleV1 contract.
type FeeOracleV1ProtocolFeeSetIterator struct {
	Event *FeeOracleV1ProtocolFeeSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV1ProtocolFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1ProtocolFeeSet)
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
		it.Event = new(FeeOracleV1ProtocolFeeSet)
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
func (it *FeeOracleV1ProtocolFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1ProtocolFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1ProtocolFeeSet represents a ProtocolFeeSet event raised by the FeeOracleV1 contract.
type FeeOracleV1ProtocolFeeSet struct {
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeSet is a free log retrieval operation binding the contract event 0xdb5aafdb29539329e37d4e3ee869bc4031941fd55a5dfc92824fbe34b204e30d.
//
// Solidity: event ProtocolFeeSet(uint256 protocolFee)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterProtocolFeeSet(opts *bind.FilterOpts) (*FeeOracleV1ProtocolFeeSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "ProtocolFeeSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1ProtocolFeeSetIterator{contract: _FeeOracleV1.contract, event: "ProtocolFeeSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeSet is a free log subscription operation binding the contract event 0xdb5aafdb29539329e37d4e3ee869bc4031941fd55a5dfc92824fbe34b204e30d.
//
// Solidity: event ProtocolFeeSet(uint256 protocolFee)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchProtocolFeeSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1ProtocolFeeSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "ProtocolFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1ProtocolFeeSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "ProtocolFeeSet", log); err != nil {
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

// ParseProtocolFeeSet is a log parse operation binding the contract event 0xdb5aafdb29539329e37d4e3ee869bc4031941fd55a5dfc92824fbe34b204e30d.
//
// Solidity: event ProtocolFeeSet(uint256 protocolFee)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseProtocolFeeSet(log types.Log) (*FeeOracleV1ProtocolFeeSet, error) {
	event := new(FeeOracleV1ProtocolFeeSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "ProtocolFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1ToNativeRateSetIterator is returned from FilterToNativeRateSet and is used to iterate over the raw logs and unpacked data for ToNativeRateSet events raised by the FeeOracleV1 contract.
type FeeOracleV1ToNativeRateSetIterator struct {
	Event *FeeOracleV1ToNativeRateSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV1ToNativeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1ToNativeRateSet)
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
		it.Event = new(FeeOracleV1ToNativeRateSet)
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
func (it *FeeOracleV1ToNativeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1ToNativeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1ToNativeRateSet represents a ToNativeRateSet event raised by the FeeOracleV1 contract.
type FeeOracleV1ToNativeRateSet struct {
	ChainId      uint64
	ToNativeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterToNativeRateSet is a free log retrieval operation binding the contract event 0x4b4594c9f06af25bc504eead96f7f0eaa3f1577f8d9b075b236520ec712e1308.
//
// Solidity: event ToNativeRateSet(uint64 chainId, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterToNativeRateSet(opts *bind.FilterOpts) (*FeeOracleV1ToNativeRateSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "ToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1ToNativeRateSetIterator{contract: _FeeOracleV1.contract, event: "ToNativeRateSet", logs: logs, sub: sub}, nil
}

// WatchToNativeRateSet is a free log subscription operation binding the contract event 0x4b4594c9f06af25bc504eead96f7f0eaa3f1577f8d9b075b236520ec712e1308.
//
// Solidity: event ToNativeRateSet(uint64 chainId, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchToNativeRateSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1ToNativeRateSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "ToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1ToNativeRateSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "ToNativeRateSet", log); err != nil {
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

// ParseToNativeRateSet is a log parse operation binding the contract event 0x4b4594c9f06af25bc504eead96f7f0eaa3f1577f8d9b075b236520ec712e1308.
//
// Solidity: event ToNativeRateSet(uint64 chainId, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseToNativeRateSet(log types.Log) (*FeeOracleV1ToNativeRateSet, error) {
	event := new(FeeOracleV1ToNativeRateSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "ToNativeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
