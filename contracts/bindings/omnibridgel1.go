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

// OmniBridgeL1MetaData contains all meta data concerning the OmniBridgeL1 contract.
var OmniBridgeL1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"XCALL_WITHDRAW_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalL1Supply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bridge\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161108c38038061108c83398101604081905261002f916100fc565b6001600160a01b03811660805261004461004a565b5061012c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561009a5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100f95780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60006020828403121561010e57600080fd5b81516001600160a01b038116811461012557600080fd5b9392505050565b608051610f37610155600039600081816102340152818161073301526109f10152610f376000f3fe60806040526004361061009c5760003560e01c80638fdcb4c9116100645780638fdcb4c9146101805780639c5451b0146101b0578063c3de453d146101cf578063f2fde38b146101e2578063f3fef3a314610202578063fc0c546a1461022257600080fd5b80633794999d146100a157806339acf9f1146100d4578063485cc9551461010c578063715018a61461012e5780638da5cb5b14610143575b600080fd5b3480156100ad57600080fd5b506100c16100bc366004610c8c565b610256565b6040519081526020015b60405180910390f35b3480156100e057600080fd5b506000546100f4906001600160a01b031681565b6040516001600160a01b0390911681526020016100cb565b34801561011857600080fd5b5061012c610127366004610ccd565b61038a565b005b34801561013a57600080fd5b5061012c6104b6565b34801561014f57600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166100f4565b34801561018c57600080fd5b50610197620124f881565b60405167ffffffffffffffff90911681526020016100cb565b3480156101bc57600080fd5b506100c16a52b7d2dcc80cd2e400000081565b61012c6101dd366004610d06565b6104ca565b3480156101ee57600080fd5b5061012c6101fd366004610d32565b6104d9565b34801561020e57600080fd5b5061012c61021d366004610d06565b61051c565b34801561022e57600080fd5b506100f47f000000000000000000000000000000000000000000000000000000000000000081565b600080546040805163110ff5f160e01b815290516001600160a01b0390921691638dd9523c91839163110ff5f1916004808201926020929091908290030181865afa1580156102a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cd9190610d73565b604080516001600160a01b03898116602483015288166044820152606480820188905282518083039091018152608490910182526020810180516001600160e01b0316636ce5768960e11b179052905160e084901b6001600160e01b0319168152610341929190620124f890600401610dd4565b602060405180830381865afa15801561035e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103829190610e0b565b949350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156103d05750825b905060008267ffffffffffffffff1660011480156103ed5750303b155b9050811580156103fb575080155b156104195760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561044357845460ff60401b1916600160401b1785555b61044c876107e9565b600080546001600160a01b0319166001600160a01b03881617905583156104ad57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b6104be6107fa565b6104c86000610855565b565b6104d53383836108c6565b5050565b6104e16107fa565b6001600160a01b03811661051057604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61051981610855565b50565b6000805460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa158015610564573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105889190610e24565b6000549091506001600160a01b031633146105dd5760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b6044820152606401610507565b60208101516001600160a01b0316600262048789608a1b011461063b5760405162461bcd60e51b81526020600482015260166024820152754f6d6e694272696467653a206e6f742062726964676560501b6044820152606401610507565b60008054906101000a90046001600160a01b03166001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa15801561068c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b09190610d73565b67ffffffffffffffff16816000015167ffffffffffffffff161461070d5760405162461bcd60e51b81526020600482015260146024820152734f6d6e694272696467653a206e6f74206f6d6e6960601b6044820152606401610507565b60405163a9059cbb60e01b81526001600160a01b038481166004830152602482018490527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af115801561077c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a09190610e8f565b50826001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364836040516107dc91815260200190565b60405180910390a2505050565b6107f1610c26565b61051981610c6f565b3361082c7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146104c85760405163118cdaa760e01b8152336004820152602401610507565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b600081116109165760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694272696467653a20616d6f756e74206d757374206265203e203000006044820152606401610507565b6001600160a01b03821661096c5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e694272696467653a206e6f2062726964676520746f207a65726f0000006044820152606401610507565b610977838383610256565b34146109c55760405162461bcd60e51b815260206004820152601960248201527f4f6d6e694272696467653a20696e636f727265637420666565000000000000006044820152606401610507565b6040516323b872dd60e01b81526001600160a01b038481166004830152306024830152604482018390527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af1158015610a3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5e9190610e8f565b610aaa5760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e694272696467653a207472616e73666572206661696c656400000000006044820152606401610507565b6000546040805163110ff5f160e01b815290516001600160a01b039092169163c21dda4f913491849163110ff5f19160048083019260209291908290030181865afa158015610afd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b219190610d73565b604080516001600160a01b03898116602483015288166044820152606480820188905282518083039091018152608490910182526020810180516001600160e01b0316636ce5768960e11b179052905160e085901b6001600160e01b0319168152610ba29291600491600262048789608a1b019190620124f8908401610eb1565b6000604051808303818588803b158015610bbb57600080fd5b505af1158015610bcf573d6000803e3d6000fd5b5050505050816001600160a01b0316836001600160a01b03167f59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b42283604051610c1991815260200190565b60405180910390a3505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166104c857604051631afcd79f60e31b815260040160405180910390fd5b6104e1610c26565b6001600160a01b038116811461051957600080fd5b600080600060608486031215610ca157600080fd5b8335610cac81610c77565b92506020840135610cbc81610c77565b929592945050506040919091013590565b60008060408385031215610ce057600080fd5b8235610ceb81610c77565b91506020830135610cfb81610c77565b809150509250929050565b60008060408385031215610d1957600080fd5b8235610d2481610c77565b946020939093013593505050565b600060208284031215610d4457600080fd5b8135610d4f81610c77565b9392505050565b805167ffffffffffffffff81168114610d6e57600080fd5b919050565b600060208284031215610d8557600080fd5b610d4f82610d56565b6000815180845260005b81811015610db457602081850181015186830182015201610d98565b506000602082860101526020601f19601f83011685010191505092915050565b600067ffffffffffffffff808616835260606020840152610df86060840186610d8e565b9150808416604084015250949350505050565b600060208284031215610e1d57600080fd5b5051919050565b600060408284031215610e3657600080fd5b6040516040810181811067ffffffffffffffff82111715610e6757634e487b7160e01b600052604160045260246000fd5b604052610e7383610d56565b81526020830151610e8381610c77565b60208201529392505050565b600060208284031215610ea157600080fd5b81518015158114610d4f57600080fd5b600067ffffffffffffffff808816835260ff8716602084015260018060a01b038616604084015260a06060840152610eec60a0840186610d8e565b9150808416608084015250969550505050505056fea264697066735822122023469baef3c0250b53c151258684228e56cde8d16067f36f21e760a116afe34064736f6c63430008180033",
}

// OmniBridgeL1ABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniBridgeL1MetaData.ABI instead.
var OmniBridgeL1ABI = OmniBridgeL1MetaData.ABI

// OmniBridgeL1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniBridgeL1MetaData.Bin instead.
var OmniBridgeL1Bin = OmniBridgeL1MetaData.Bin

// DeployOmniBridgeL1 deploys a new Ethereum contract, binding an instance of OmniBridgeL1 to it.
func DeployOmniBridgeL1(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address) (common.Address, *types.Transaction, *OmniBridgeL1, error) {
	parsed, err := OmniBridgeL1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniBridgeL1Bin), backend, token_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniBridgeL1{OmniBridgeL1Caller: OmniBridgeL1Caller{contract: contract}, OmniBridgeL1Transactor: OmniBridgeL1Transactor{contract: contract}, OmniBridgeL1Filterer: OmniBridgeL1Filterer{contract: contract}}, nil
}

// OmniBridgeL1 is an auto generated Go binding around an Ethereum contract.
type OmniBridgeL1 struct {
	OmniBridgeL1Caller     // Read-only binding to the contract
	OmniBridgeL1Transactor // Write-only binding to the contract
	OmniBridgeL1Filterer   // Log filterer for contract events
}

// OmniBridgeL1Caller is an auto generated read-only Go binding around an Ethereum contract.
type OmniBridgeL1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniBridgeL1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniBridgeL1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniBridgeL1Session struct {
	Contract     *OmniBridgeL1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniBridgeL1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniBridgeL1CallerSession struct {
	Contract *OmniBridgeL1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OmniBridgeL1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniBridgeL1TransactorSession struct {
	Contract     *OmniBridgeL1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OmniBridgeL1Raw is an auto generated low-level Go binding around an Ethereum contract.
type OmniBridgeL1Raw struct {
	Contract *OmniBridgeL1 // Generic contract binding to access the raw methods on
}

// OmniBridgeL1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniBridgeL1CallerRaw struct {
	Contract *OmniBridgeL1Caller // Generic read-only contract binding to access the raw methods on
}

// OmniBridgeL1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniBridgeL1TransactorRaw struct {
	Contract *OmniBridgeL1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniBridgeL1 creates a new instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1(address common.Address, backend bind.ContractBackend) (*OmniBridgeL1, error) {
	contract, err := bindOmniBridgeL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1{OmniBridgeL1Caller: OmniBridgeL1Caller{contract: contract}, OmniBridgeL1Transactor: OmniBridgeL1Transactor{contract: contract}, OmniBridgeL1Filterer: OmniBridgeL1Filterer{contract: contract}}, nil
}

// NewOmniBridgeL1Caller creates a new read-only instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Caller(address common.Address, caller bind.ContractCaller) (*OmniBridgeL1Caller, error) {
	contract, err := bindOmniBridgeL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Caller{contract: contract}, nil
}

// NewOmniBridgeL1Transactor creates a new write-only instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Transactor(address common.Address, transactor bind.ContractTransactor) (*OmniBridgeL1Transactor, error) {
	contract, err := bindOmniBridgeL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Transactor{contract: contract}, nil
}

// NewOmniBridgeL1Filterer creates a new log filterer instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Filterer(address common.Address, filterer bind.ContractFilterer) (*OmniBridgeL1Filterer, error) {
	contract, err := bindOmniBridgeL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Filterer{contract: contract}, nil
}

// bindOmniBridgeL1 binds a generic wrapper to an already deployed contract.
func bindOmniBridgeL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniBridgeL1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniBridgeL1.Contract.OmniBridgeL1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.OmniBridgeL1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.OmniBridgeL1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniBridgeL1 *OmniBridgeL1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniBridgeL1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniBridgeL1 *OmniBridgeL1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniBridgeL1 *OmniBridgeL1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.contract.Transact(opts, method, params...)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1Caller) XCALLWITHDRAWGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "XCALL_WITHDRAW_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1Session) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _OmniBridgeL1.Contract.XCALLWITHDRAWGASLIMIT(&_OmniBridgeL1.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _OmniBridgeL1.Contract.XCALLWITHDRAWGASLIMIT(&_OmniBridgeL1.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Caller) BridgeFee(opts *bind.CallOpts, payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "bridgeFee", payor, to, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Session) BridgeFee(payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	return _OmniBridgeL1.Contract.BridgeFee(&_OmniBridgeL1.CallOpts, payor, to, amount)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) BridgeFee(payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	return _OmniBridgeL1.Contract.BridgeFee(&_OmniBridgeL1.CallOpts, payor, to, amount)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Omni() (common.Address, error) {
	return _OmniBridgeL1.Contract.Omni(&_OmniBridgeL1.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Omni() (common.Address, error) {
	return _OmniBridgeL1.Contract.Omni(&_OmniBridgeL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Owner() (common.Address, error) {
	return _OmniBridgeL1.Contract.Owner(&_OmniBridgeL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Owner() (common.Address, error) {
	return _OmniBridgeL1.Contract.Owner(&_OmniBridgeL1.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Token() (common.Address, error) {
	return _OmniBridgeL1.Contract.Token(&_OmniBridgeL1.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Token() (common.Address, error) {
	return _OmniBridgeL1.Contract.Token(&_OmniBridgeL1.CallOpts)
}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Caller) TotalL1Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "totalL1Supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Session) TotalL1Supply() (*big.Int, error) {
	return _OmniBridgeL1.Contract.TotalL1Supply(&_OmniBridgeL1.CallOpts)
}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) TotalL1Supply() (*big.Int, error) {
	return _OmniBridgeL1.Contract.TotalL1Supply(&_OmniBridgeL1.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Bridge(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "bridge", to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Bridge(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Bridge(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "initialize", owner_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Initialize(owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Initialize(&_OmniBridgeL1.TransactOpts, owner_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Initialize(owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Initialize(&_OmniBridgeL1.TransactOpts, owner_, omni_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) RenounceOwnership() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.RenounceOwnership(&_OmniBridgeL1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.RenounceOwnership(&_OmniBridgeL1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.TransferOwnership(&_OmniBridgeL1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.TransferOwnership(&_OmniBridgeL1.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Withdraw(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Withdraw(&_OmniBridgeL1.TransactOpts, to, amount)
}

// OmniBridgeL1BridgeIterator is returned from FilterBridge and is used to iterate over the raw logs and unpacked data for Bridge events raised by the OmniBridgeL1 contract.
type OmniBridgeL1BridgeIterator struct {
	Event *OmniBridgeL1Bridge // Event containing the contract specifics and raw log

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
func (it *OmniBridgeL1BridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Bridge)
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
		it.Event = new(OmniBridgeL1Bridge)
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
func (it *OmniBridgeL1BridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1BridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Bridge represents a Bridge event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Bridge struct {
	Payor  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridge is a free log retrieval operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterBridge(opts *bind.FilterOpts, payor []common.Address, to []common.Address) (*OmniBridgeL1BridgeIterator, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1BridgeIterator{contract: _OmniBridgeL1.contract, event: "Bridge", logs: logs, sub: sub}, nil
}

// WatchBridge is a free log subscription operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchBridge(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Bridge, payor []common.Address, to []common.Address) (event.Subscription, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Bridge)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Bridge", log); err != nil {
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

// ParseBridge is a log parse operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseBridge(log types.Log) (*OmniBridgeL1Bridge, error) {
	event := new(OmniBridgeL1Bridge)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Bridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniBridgeL1 contract.
type OmniBridgeL1InitializedIterator struct {
	Event *OmniBridgeL1Initialized // Event containing the contract specifics and raw log

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
func (it *OmniBridgeL1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Initialized)
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
		it.Event = new(OmniBridgeL1Initialized)
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
func (it *OmniBridgeL1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Initialized represents a Initialized event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterInitialized(opts *bind.FilterOpts) (*OmniBridgeL1InitializedIterator, error) {

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1InitializedIterator{contract: _OmniBridgeL1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Initialized) (event.Subscription, error) {

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Initialized)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseInitialized(log types.Log) (*OmniBridgeL1Initialized, error) {
	event := new(OmniBridgeL1Initialized)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniBridgeL1 contract.
type OmniBridgeL1OwnershipTransferredIterator struct {
	Event *OmniBridgeL1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniBridgeL1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1OwnershipTransferred)
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
		it.Event = new(OmniBridgeL1OwnershipTransferred)
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
func (it *OmniBridgeL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1OwnershipTransferred represents a OwnershipTransferred event raised by the OmniBridgeL1 contract.
type OmniBridgeL1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniBridgeL1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1OwnershipTransferredIterator{contract: _OmniBridgeL1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1OwnershipTransferred)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseOwnershipTransferred(log types.Log) (*OmniBridgeL1OwnershipTransferred, error) {
	event := new(OmniBridgeL1OwnershipTransferred)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the OmniBridgeL1 contract.
type OmniBridgeL1WithdrawIterator struct {
	Event *OmniBridgeL1Withdraw // Event containing the contract specifics and raw log

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
func (it *OmniBridgeL1WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Withdraw)
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
		it.Event = new(OmniBridgeL1Withdraw)
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
func (it *OmniBridgeL1WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Withdraw represents a Withdraw event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Withdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*OmniBridgeL1WithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1WithdrawIterator{contract: _OmniBridgeL1.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Withdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Withdraw)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseWithdraw(log types.Log) (*OmniBridgeL1Withdraw, error) {
	event := new(OmniBridgeL1Withdraw)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
