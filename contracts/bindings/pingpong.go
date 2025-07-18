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

// PingPongMetaData contains all meta data concerning the PingPong contract.
var PingPongMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pingpong\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"conf\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"n\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"start\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destChainID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"conf\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"times\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Done\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destChainID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"times\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Ping\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"srcChainID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"n\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610c16380380610c1683398101604081905261002e916101ad565b80600161003a8261004b565b610043816100ef565b5050506101da565b6001600160a01b03811661009b5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b60448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47906020015b60405180910390a150565b6100f881610191565b6101445760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e76616c696420636f6e66206c6576656c00000000000000006044820152606401610092565b5f805460ff60a01b1916600160a01b60ff8416908102919091179091556040519081527f8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483906020016100e4565b5f60ff8216600114806101a7575060ff82166004145b92915050565b5f602082840312156101bd575f80fd5b81516001600160a01b03811681146101d3575f80fd5b9392505050565b610a2f806101e75f395ff3fe60806040526004361061004c575f3560e01c8063091d27881461005757806339acf9f11461008a57806374eeb847146100c0578063a36d4241146100f1578063b81ce37614610112575f80fd5b3661005357005b5f80fd5b348015610062575f80fd5b5061006d62030d4081565b6040516001600160401b0390911681526020015b60405180910390f35b348015610095575f80fd5b505f546100a8906001600160a01b031681565b6040516001600160a01b039091168152602001610081565b3480156100cb575f80fd5b505f546100df90600160a01b900460ff1681565b60405160ff9091168152602001610081565b3480156100fc575f80fd5b5061011061010b366004610681565b610131565b005b34801561011d575f80fd5b5061011061012c366004610707565b6101bb565b5f816001600160401b03161161018e5760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a2074696d6573206d757374206265203e2030000000000060448201526064015b60405180910390fd5b6101b386868686868660016101a4826002610791565b6101ae91906107bc565b6103cc565b505050505050565b5f5460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa1580156101ff573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061022391906107e3565b8051600180546020909301516001600160a01b0316600160401b026001600160e01b03199093166001600160401b0390921691909117919091179055610267610438565b6102b35760405162461bcd60e51b815260206004820152601b60248201527f50696e67506f6e673a206e6f7420616e206f6d6e69207863616c6c00000000006044820152606401610185565b6001546040517f47768d18c1bf0d890e19b6f020bd996b385273187be95e2b6a58feccf3bd75669161030691889188916001600160401b03821691600160401b90046001600160a01b0316908790610873565b60405180910390a1806001600160401b03165f0361037e576001546040517f5335290c57b37808bcd506c1e858915c89a2c9adaa76dc0319044eab5b9d18c29161037191889188916001600160401b03821691600160401b90046001600160a01b0316908890610873565b60405180910390a16103b5565b600180546103b591879187916001600160401b038216918891600160401b9091046001600160a01b03169088906101ae90896107bc565b5050600180546001600160e01b0319169055505050565b61042e85858563b81ce37660e01b8b8b8a89896040516024016103f39594939291906108b9565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915262030d406104c6565b5050505050505050565b5f805f9054906101000a90046001600160a01b03166001600160a01b03166355e2448e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610488573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104ac91906108f6565b80156104c157505f546001600160a01b031633145b905090565b5f8054604051632376548f60e21b815282916001600160a01b031690638dd9523c906104fa908a908890889060040161095f565b602060405180830381865afa158015610515573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105399190610994565b90508047101561058b5760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401610185565b5f5460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f9083906105c4908b908b908b908b908b906004016109ab565b5f604051808303818588803b1580156105db575f80fd5b505af11580156105ed573d5f803e3d5ffd5b50939a9950505050505050505050565b5f8083601f84011261060d575f80fd5b5081356001600160401b03811115610623575f80fd5b60208301915083602082850101111561063a575f80fd5b9250929050565b6001600160401b0381168114610655575f80fd5b50565b803560ff81168114610668575f80fd5b919050565b6001600160a01b0381168114610655575f80fd5b5f805f805f8060a08789031215610696575f80fd5b86356001600160401b038111156106ab575f80fd5b6106b789828a016105fd565b90975095505060208701356106cb81610641565b93506106d960408801610658565b925060608701356106e98161066d565b915060808701356106f981610641565b809150509295509295509295565b5f805f805f6080868803121561071b575f80fd5b85356001600160401b03811115610730575f80fd5b61073c888289016105fd565b909650945061074f905060208701610658565b9250604086013561075f81610641565b9150606086013561076f81610641565b809150509295509295909350565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038181168382160280821691908281146107b4576107b461077d565b505092915050565b6001600160401b038281168282160390808211156107dc576107dc61077d565b5092915050565b5f604082840312156107f3575f80fd5b604051604081018181106001600160401b038211171561082157634e487b7160e01b5f52604160045260245ffd5b604052825161082f81610641565b8152602083015161083f8161066d565b60208201529392505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b608081525f61088660808301878961084b565b6001600160401b0395861660208401526001600160a01b0394909416604083015250921660609092019190915292915050565b608081525f6108cc60808301878961084b565b60ff959095166020830152506001600160401b039283166040820152911660609091015292915050565b5f60208284031215610906575f80fd5b81518015158114610915575f80fd5b9392505050565b5f81518084525f5b8181101561094057602081850181015186830182015201610924565b505f602082860101526020601f19601f83011685010191505092915050565b5f6001600160401b03808616835260606020840152610981606084018661091c565b9150808416604084015250949350505050565b5f602082840312156109a4575f80fd5b5051919050565b5f6001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a060608401526109e460a084018661091c565b9150808416608084015250969550505050505056fea264697066735822122064e44505d8600dd67227f858675403242c4679a21d4b50362e098586e813a52e64736f6c63430008180033",
}

// PingPongABI is the input ABI used to generate the binding from.
// Deprecated: Use PingPongMetaData.ABI instead.
var PingPongABI = PingPongMetaData.ABI

// PingPongBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PingPongMetaData.Bin instead.
var PingPongBin = PingPongMetaData.Bin

// DeployPingPong deploys a new Ethereum contract, binding an instance of PingPong to it.
func DeployPingPong(auth *bind.TransactOpts, backend bind.ContractBackend, portal common.Address) (common.Address, *types.Transaction, *PingPong, error) {
	parsed, err := PingPongMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PingPongBin), backend, portal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PingPong{PingPongCaller: PingPongCaller{contract: contract}, PingPongTransactor: PingPongTransactor{contract: contract}, PingPongFilterer: PingPongFilterer{contract: contract}}, nil
}

// PingPong is an auto generated Go binding around an Ethereum contract.
type PingPong struct {
	PingPongCaller     // Read-only binding to the contract
	PingPongTransactor // Write-only binding to the contract
	PingPongFilterer   // Log filterer for contract events
}

// PingPongCaller is an auto generated read-only Go binding around an Ethereum contract.
type PingPongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PingPongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PingPongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PingPongSession struct {
	Contract     *PingPong         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingPongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PingPongCallerSession struct {
	Contract *PingPongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PingPongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PingPongTransactorSession struct {
	Contract     *PingPongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PingPongRaw is an auto generated low-level Go binding around an Ethereum contract.
type PingPongRaw struct {
	Contract *PingPong // Generic contract binding to access the raw methods on
}

// PingPongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PingPongCallerRaw struct {
	Contract *PingPongCaller // Generic read-only contract binding to access the raw methods on
}

// PingPongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PingPongTransactorRaw struct {
	Contract *PingPongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPingPong creates a new instance of PingPong, bound to a specific deployed contract.
func NewPingPong(address common.Address, backend bind.ContractBackend) (*PingPong, error) {
	contract, err := bindPingPong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PingPong{PingPongCaller: PingPongCaller{contract: contract}, PingPongTransactor: PingPongTransactor{contract: contract}, PingPongFilterer: PingPongFilterer{contract: contract}}, nil
}

// NewPingPongCaller creates a new read-only instance of PingPong, bound to a specific deployed contract.
func NewPingPongCaller(address common.Address, caller bind.ContractCaller) (*PingPongCaller, error) {
	contract, err := bindPingPong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongCaller{contract: contract}, nil
}

// NewPingPongTransactor creates a new write-only instance of PingPong, bound to a specific deployed contract.
func NewPingPongTransactor(address common.Address, transactor bind.ContractTransactor) (*PingPongTransactor, error) {
	contract, err := bindPingPong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongTransactor{contract: contract}, nil
}

// NewPingPongFilterer creates a new log filterer instance of PingPong, bound to a specific deployed contract.
func NewPingPongFilterer(address common.Address, filterer bind.ContractFilterer) (*PingPongFilterer, error) {
	contract, err := bindPingPong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingPongFilterer{contract: contract}, nil
}

// bindPingPong binds a generic wrapper to an already deployed contract.
func bindPingPong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PingPongMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PingPong *PingPongRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPong.Contract.PingPongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PingPong *PingPongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.Contract.PingPongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PingPong *PingPongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPong.Contract.PingPongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PingPong *PingPongCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PingPong *PingPongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PingPong *PingPongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPong.Contract.contract.Transact(opts, method, params...)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint64)
func (_PingPong *PingPongCaller) GASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PingPong.contract.Call(opts, &out, "GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint64)
func (_PingPong *PingPongSession) GASLIMIT() (uint64, error) {
	return _PingPong.Contract.GASLIMIT(&_PingPong.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint64)
func (_PingPong *PingPongCallerSession) GASLIMIT() (uint64, error) {
	return _PingPong.Contract.GASLIMIT(&_PingPong.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_PingPong *PingPongCaller) DefaultConfLevel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PingPong.contract.Call(opts, &out, "defaultConfLevel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_PingPong *PingPongSession) DefaultConfLevel() (uint8, error) {
	return _PingPong.Contract.DefaultConfLevel(&_PingPong.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_PingPong *PingPongCallerSession) DefaultConfLevel() (uint8, error) {
	return _PingPong.Contract.DefaultConfLevel(&_PingPong.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_PingPong *PingPongCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPong.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_PingPong *PingPongSession) Omni() (common.Address, error) {
	return _PingPong.Contract.Omni(&_PingPong.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_PingPong *PingPongCallerSession) Omni() (common.Address, error) {
	return _PingPong.Contract.Omni(&_PingPong.CallOpts)
}

// Pingpong is a paid mutator transaction binding the contract method 0xb81ce376.
//
// Solidity: function pingpong(string id, uint8 conf, uint64 times, uint64 n) returns()
func (_PingPong *PingPongTransactor) Pingpong(opts *bind.TransactOpts, id string, conf uint8, times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.contract.Transact(opts, "pingpong", id, conf, times, n)
}

// Pingpong is a paid mutator transaction binding the contract method 0xb81ce376.
//
// Solidity: function pingpong(string id, uint8 conf, uint64 times, uint64 n) returns()
func (_PingPong *PingPongSession) Pingpong(id string, conf uint8, times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Pingpong(&_PingPong.TransactOpts, id, conf, times, n)
}

// Pingpong is a paid mutator transaction binding the contract method 0xb81ce376.
//
// Solidity: function pingpong(string id, uint8 conf, uint64 times, uint64 n) returns()
func (_PingPong *PingPongTransactorSession) Pingpong(id string, conf uint8, times uint64, n uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Pingpong(&_PingPong.TransactOpts, id, conf, times, n)
}

// Start is a paid mutator transaction binding the contract method 0xa36d4241.
//
// Solidity: function start(string id, uint64 destChainID, uint8 conf, address to, uint64 times) returns()
func (_PingPong *PingPongTransactor) Start(opts *bind.TransactOpts, id string, destChainID uint64, conf uint8, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.contract.Transact(opts, "start", id, destChainID, conf, to, times)
}

// Start is a paid mutator transaction binding the contract method 0xa36d4241.
//
// Solidity: function start(string id, uint64 destChainID, uint8 conf, address to, uint64 times) returns()
func (_PingPong *PingPongSession) Start(id string, destChainID uint64, conf uint8, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Start(&_PingPong.TransactOpts, id, destChainID, conf, to, times)
}

// Start is a paid mutator transaction binding the contract method 0xa36d4241.
//
// Solidity: function start(string id, uint64 destChainID, uint8 conf, address to, uint64 times) returns()
func (_PingPong *PingPongTransactorSession) Start(id string, destChainID uint64, conf uint8, to common.Address, times uint64) (*types.Transaction, error) {
	return _PingPong.Contract.Start(&_PingPong.TransactOpts, id, destChainID, conf, to, times)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PingPong *PingPongTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PingPong *PingPongSession) Receive() (*types.Transaction, error) {
	return _PingPong.Contract.Receive(&_PingPong.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PingPong *PingPongTransactorSession) Receive() (*types.Transaction, error) {
	return _PingPong.Contract.Receive(&_PingPong.TransactOpts)
}

// PingPongDefaultConfLevelSetIterator is returned from FilterDefaultConfLevelSet and is used to iterate over the raw logs and unpacked data for DefaultConfLevelSet events raised by the PingPong contract.
type PingPongDefaultConfLevelSetIterator struct {
	Event *PingPongDefaultConfLevelSet // Event containing the contract specifics and raw log

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
func (it *PingPongDefaultConfLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDefaultConfLevelSet)
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
		it.Event = new(PingPongDefaultConfLevelSet)
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
func (it *PingPongDefaultConfLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongDefaultConfLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongDefaultConfLevelSet represents a DefaultConfLevelSet event raised by the PingPong contract.
type PingPongDefaultConfLevelSet struct {
	Conf uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultConfLevelSet is a free log retrieval operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_PingPong *PingPongFilterer) FilterDefaultConfLevelSet(opts *bind.FilterOpts) (*PingPongDefaultConfLevelSetIterator, error) {

	logs, sub, err := _PingPong.contract.FilterLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return &PingPongDefaultConfLevelSetIterator{contract: _PingPong.contract, event: "DefaultConfLevelSet", logs: logs, sub: sub}, nil
}

// WatchDefaultConfLevelSet is a free log subscription operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_PingPong *PingPongFilterer) WatchDefaultConfLevelSet(opts *bind.WatchOpts, sink chan<- *PingPongDefaultConfLevelSet) (event.Subscription, error) {

	logs, sub, err := _PingPong.contract.WatchLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongDefaultConfLevelSet)
				if err := _PingPong.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
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

// ParseDefaultConfLevelSet is a log parse operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_PingPong *PingPongFilterer) ParseDefaultConfLevelSet(log types.Log) (*PingPongDefaultConfLevelSet, error) {
	event := new(PingPongDefaultConfLevelSet)
	if err := _PingPong.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingPongDoneIterator is returned from FilterDone and is used to iterate over the raw logs and unpacked data for Done events raised by the PingPong contract.
type PingPongDoneIterator struct {
	Event *PingPongDone // Event containing the contract specifics and raw log

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
func (it *PingPongDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongDone)
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
		it.Event = new(PingPongDone)
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
func (it *PingPongDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongDone represents a Done event raised by the PingPong contract.
type PingPongDone struct {
	Id          string
	DestChainID uint64
	To          common.Address
	Times       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDone is a free log retrieval operation binding the contract event 0x5335290c57b37808bcd506c1e858915c89a2c9adaa76dc0319044eab5b9d18c2.
//
// Solidity: event Done(string id, uint64 destChainID, address to, uint64 times)
func (_PingPong *PingPongFilterer) FilterDone(opts *bind.FilterOpts) (*PingPongDoneIterator, error) {

	logs, sub, err := _PingPong.contract.FilterLogs(opts, "Done")
	if err != nil {
		return nil, err
	}
	return &PingPongDoneIterator{contract: _PingPong.contract, event: "Done", logs: logs, sub: sub}, nil
}

// WatchDone is a free log subscription operation binding the contract event 0x5335290c57b37808bcd506c1e858915c89a2c9adaa76dc0319044eab5b9d18c2.
//
// Solidity: event Done(string id, uint64 destChainID, address to, uint64 times)
func (_PingPong *PingPongFilterer) WatchDone(opts *bind.WatchOpts, sink chan<- *PingPongDone) (event.Subscription, error) {

	logs, sub, err := _PingPong.contract.WatchLogs(opts, "Done")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongDone)
				if err := _PingPong.contract.UnpackLog(event, "Done", log); err != nil {
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

// ParseDone is a log parse operation binding the contract event 0x5335290c57b37808bcd506c1e858915c89a2c9adaa76dc0319044eab5b9d18c2.
//
// Solidity: event Done(string id, uint64 destChainID, address to, uint64 times)
func (_PingPong *PingPongFilterer) ParseDone(log types.Log) (*PingPongDone, error) {
	event := new(PingPongDone)
	if err := _PingPong.contract.UnpackLog(event, "Done", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingPongOmniPortalSetIterator is returned from FilterOmniPortalSet and is used to iterate over the raw logs and unpacked data for OmniPortalSet events raised by the PingPong contract.
type PingPongOmniPortalSetIterator struct {
	Event *PingPongOmniPortalSet // Event containing the contract specifics and raw log

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
func (it *PingPongOmniPortalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongOmniPortalSet)
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
		it.Event = new(PingPongOmniPortalSet)
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
func (it *PingPongOmniPortalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongOmniPortalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongOmniPortalSet represents a OmniPortalSet event raised by the PingPong contract.
type PingPongOmniPortalSet struct {
	Omni common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOmniPortalSet is a free log retrieval operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_PingPong *PingPongFilterer) FilterOmniPortalSet(opts *bind.FilterOpts) (*PingPongOmniPortalSetIterator, error) {

	logs, sub, err := _PingPong.contract.FilterLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return &PingPongOmniPortalSetIterator{contract: _PingPong.contract, event: "OmniPortalSet", logs: logs, sub: sub}, nil
}

// WatchOmniPortalSet is a free log subscription operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_PingPong *PingPongFilterer) WatchOmniPortalSet(opts *bind.WatchOpts, sink chan<- *PingPongOmniPortalSet) (event.Subscription, error) {

	logs, sub, err := _PingPong.contract.WatchLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongOmniPortalSet)
				if err := _PingPong.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
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

// ParseOmniPortalSet is a log parse operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_PingPong *PingPongFilterer) ParseOmniPortalSet(log types.Log) (*PingPongOmniPortalSet, error) {
	event := new(PingPongOmniPortalSet)
	if err := _PingPong.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingPongPingIterator is returned from FilterPing and is used to iterate over the raw logs and unpacked data for Ping events raised by the PingPong contract.
type PingPongPingIterator struct {
	Event *PingPongPing // Event containing the contract specifics and raw log

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
func (it *PingPongPingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongPing)
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
		it.Event = new(PingPongPing)
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
func (it *PingPongPingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongPingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongPing represents a Ping event raised by the PingPong contract.
type PingPongPing struct {
	Id         string
	SrcChainID uint64
	From       common.Address
	N          uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPing is a free log retrieval operation binding the contract event 0x47768d18c1bf0d890e19b6f020bd996b385273187be95e2b6a58feccf3bd7566.
//
// Solidity: event Ping(string id, uint64 srcChainID, address from, uint64 n)
func (_PingPong *PingPongFilterer) FilterPing(opts *bind.FilterOpts) (*PingPongPingIterator, error) {

	logs, sub, err := _PingPong.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &PingPongPingIterator{contract: _PingPong.contract, event: "Ping", logs: logs, sub: sub}, nil
}

// WatchPing is a free log subscription operation binding the contract event 0x47768d18c1bf0d890e19b6f020bd996b385273187be95e2b6a58feccf3bd7566.
//
// Solidity: event Ping(string id, uint64 srcChainID, address from, uint64 n)
func (_PingPong *PingPongFilterer) WatchPing(opts *bind.WatchOpts, sink chan<- *PingPongPing) (event.Subscription, error) {

	logs, sub, err := _PingPong.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongPing)
				if err := _PingPong.contract.UnpackLog(event, "Ping", log); err != nil {
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

// ParsePing is a log parse operation binding the contract event 0x47768d18c1bf0d890e19b6f020bd996b385273187be95e2b6a58feccf3bd7566.
//
// Solidity: event Ping(string id, uint64 srcChainID, address from, uint64 n)
func (_PingPong *PingPongFilterer) ParsePing(log types.Log) (*PingPongPing, error) {
	event := new(PingPongPing)
	if err := _PingPong.contract.UnpackLog(event, "Ping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
