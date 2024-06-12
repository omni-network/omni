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

// XRegistryBaseDeployment is an auto generated low-level Go binding around an user-defined struct.
type XRegistryBaseDeployment struct {
	Addr     common.Address
	Metadata []byte
}

// XRegistryReplicaMetaData contains all meta data concerning the XRegistryReplica contract.
var XRegistryReplicaMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"portal_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"has\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"set\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dep\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610d95380380610d9583398101604081905261002f9161009a565b6001600160a01b0381166100895760405162461bcd60e51b815260206004820181905260248201527f5852656769737472795265706c6963613a206e6f207a65726f20706f7274616c604482015260640160405180910390fd5b6001600160a01b03166080526100ca565b6000602082840312156100ac57600080fd5b81516001600160a01b03811681146100c357600080fd5b9392505050565b608051610ca36100f26000396000818160d10152818161030501526103d50152610ca36000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063b3449b7714610046578063e4f1c6771461006f578063fd0b64f714610084575b600080fd5b610059610054366004610777565b6100a7565b60405161006691906107fe565b60405180910390f35b61008261007d36600461084c565b6100cf565b005b610097610092366004610777565b610446565b6040519015158152602001610066565b6040805180820190915260008152606060208201526100c7848484610468565b949350505050565b7f0000000000000000000000000000000000000000000000000000000000000000336001600160a01b038216146101435760405162461bcd60e51b81526020600482015260136024820152721614995c1b1a58d84e881b9bdd081e18d85b1b606a1b60448201526064015b60405180910390fd5b6000816001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa158015610182573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101a691906108db565b9050816001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020a919061091c565b6001600160401b031681600001516001600160401b03161461026e5760405162461bcd60e51b815260206004820152601760248201527f585265706c6963613a206e6f742066726f6d206f6d6e69000000000000000000604482015260640161013a565b60208101516001600160a01b031673121e240000000000000000000000000000000001146102de5760405162461bcd60e51b815260206004820152601c60248201527f585265706c6963613a206e6f742066726f6d2058526567697374727900000000604482015260640161013a565b6102f28686866102ed87610939565b61056a565b6102fc85856105d7565b801561039a57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639a8a05926040518163ffffffff1660e01b8152600401602060405180830381865afa158015610361573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610385919061091c565b6001600160401b0316866001600160401b0316145b1561043e5760006103ae60208501856109a5565b8101906103bb91906109f2565b60405163c0f25e8b60e01b81529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c0f25e8b9061040a908490600401610aa3565b600060405180830381600087803b15801561042457600080fd5b505af1158015610438573d6000803e3d6000fd5b50505050505b505050505050565b600080610454858585610468565b516001600160a01b03161415949350505050565b6040805180820190915260008152606060208201526001600160401b0384166000908152602081905260408120906104a0858561062b565b81526020808201929092526040908101600020815180830190925280546001600160a01b0316825260018101805492939192918401916104df90610af0565b80601f016020809104026020016040519081016040528092919081815260200182805461050b90610af0565b80156105585780601f1061052d57610100808354040283529160200191610558565b820191906000526020600020905b81548152906001019060200180831161053b57829003601f168201915b50505050508152505090509392505050565b6001600160401b0384166000908152602081905260408120829161058e868661062b565b8152602080820192909252604001600020825181546001600160a01b0319166001600160a01b039091161781559082015160018201906105ce9082610b77565b50505050505050565b60006106196040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b81525073121e24000000000000000000000000000000000261062b565b610623848461062b565b149392505050565b60008282604051602001610640929190610c36565b60405160208183030381529060405280519060200120905092915050565b6001600160401b038116811461067357600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156106ae576106ae610676565b60405290565b604051601f8201601f191681016001600160401b03811182821017156106dc576106dc610676565b604052919050565b60006001600160401b038311156106fd576106fd610676565b610710601f8401601f19166020016106b4565b905082815283838301111561072457600080fd5b828260208301376000602084830101529392505050565b600082601f83011261074c57600080fd5b61075b838335602085016106e4565b9392505050565b6001600160a01b038116811461067357600080fd5b60008060006060848603121561078c57600080fd5b83356107978161065e565b925060208401356001600160401b038111156107b257600080fd5b6107be8682870161073b565b92505060408401356107cf81610762565b809150509250925092565b60005b838110156107f55781810151838201526020016107dd565b50506000910152565b6020815260018060a01b0382511660208201526000602083015160408084015280518060608501526108378160808601602085016107da565b601f01601f1916929092016080019392505050565b6000806000806080858703121561086257600080fd5b843561086d8161065e565b935060208501356001600160401b038082111561088957600080fd5b6108958883890161073b565b9450604087013591506108a782610762565b909250606086013590808211156108bd57600080fd5b508501604081880312156108d057600080fd5b939692955090935050565b6000604082840312156108ed57600080fd5b6108f561068c565b82516109008161065e565b8152602083015161091081610762565b60208201529392505050565b60006020828403121561092e57600080fd5b815161075b8161065e565b60006040823603121561094b57600080fd5b61095361068c565b823561095e81610762565b815260208301356001600160401b0381111561097957600080fd5b830136601f82011261098a57600080fd5b610999368235602084016106e4565b60208301525092915050565b6000808335601e198436030181126109bc57600080fd5b8301803591506001600160401b038211156109d657600080fd5b6020019150368190038213156109eb57600080fd5b9250929050565b60006020808385031215610a0557600080fd5b82356001600160401b0380821115610a1c57600080fd5b818501915085601f830112610a3057600080fd5b813581811115610a4257610a42610676565b8060051b9150610a538483016106b4565b8181529183018401918481019088841115610a6d57600080fd5b938501935b83851015610a975784359250610a878361065e565b8282529385019390850190610a72565b98975050505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610ae45783516001600160401b031683529284019291840191600101610abf565b50909695505050505050565b600181811c90821680610b0457607f821691505b602082108103610b2457634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610b72576000816000526020600020601f850160051c81016020861015610b535750805b601f850160051c820191505b8181101561043e57828155600101610b5f565b505050565b81516001600160401b03811115610b9057610b90610676565b610ba481610b9e8454610af0565b84610b2a565b602080601f831160018114610bd95760008415610bc15750858301515b600019600386901b1c1916600185901b17855561043e565b600085815260208120601f198616915b82811015610c0857888601518255948401946001909101908401610be9565b5085821015610c265787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008351610c488184602088016107da565b60609390931b6bffffffffffffffffffffffff1916919092019081526014019291505056fea264697066735822122096ad20825c6695ad2b5b0c554d7dae731a4a44a62b1d146c950a53791753fa0b64736f6c63430008180033",
}

// XRegistryReplicaABI is the input ABI used to generate the binding from.
// Deprecated: Use XRegistryReplicaMetaData.ABI instead.
var XRegistryReplicaABI = XRegistryReplicaMetaData.ABI

// XRegistryReplicaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use XRegistryReplicaMetaData.Bin instead.
var XRegistryReplicaBin = XRegistryReplicaMetaData.Bin

// DeployXRegistryReplica deploys a new Ethereum contract, binding an instance of XRegistryReplica to it.
func DeployXRegistryReplica(auth *bind.TransactOpts, backend bind.ContractBackend, portal_ common.Address) (common.Address, *types.Transaction, *XRegistryReplica, error) {
	parsed, err := XRegistryReplicaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(XRegistryReplicaBin), backend, portal_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XRegistryReplica{XRegistryReplicaCaller: XRegistryReplicaCaller{contract: contract}, XRegistryReplicaTransactor: XRegistryReplicaTransactor{contract: contract}, XRegistryReplicaFilterer: XRegistryReplicaFilterer{contract: contract}}, nil
}

// XRegistryReplica is an auto generated Go binding around an Ethereum contract.
type XRegistryReplica struct {
	XRegistryReplicaCaller     // Read-only binding to the contract
	XRegistryReplicaTransactor // Write-only binding to the contract
	XRegistryReplicaFilterer   // Log filterer for contract events
}

// XRegistryReplicaCaller is an auto generated read-only Go binding around an Ethereum contract.
type XRegistryReplicaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRegistryReplicaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XRegistryReplicaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRegistryReplicaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XRegistryReplicaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRegistryReplicaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XRegistryReplicaSession struct {
	Contract     *XRegistryReplica // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XRegistryReplicaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XRegistryReplicaCallerSession struct {
	Contract *XRegistryReplicaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// XRegistryReplicaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XRegistryReplicaTransactorSession struct {
	Contract     *XRegistryReplicaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// XRegistryReplicaRaw is an auto generated low-level Go binding around an Ethereum contract.
type XRegistryReplicaRaw struct {
	Contract *XRegistryReplica // Generic contract binding to access the raw methods on
}

// XRegistryReplicaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XRegistryReplicaCallerRaw struct {
	Contract *XRegistryReplicaCaller // Generic read-only contract binding to access the raw methods on
}

// XRegistryReplicaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XRegistryReplicaTransactorRaw struct {
	Contract *XRegistryReplicaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXRegistryReplica creates a new instance of XRegistryReplica, bound to a specific deployed contract.
func NewXRegistryReplica(address common.Address, backend bind.ContractBackend) (*XRegistryReplica, error) {
	contract, err := bindXRegistryReplica(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XRegistryReplica{XRegistryReplicaCaller: XRegistryReplicaCaller{contract: contract}, XRegistryReplicaTransactor: XRegistryReplicaTransactor{contract: contract}, XRegistryReplicaFilterer: XRegistryReplicaFilterer{contract: contract}}, nil
}

// NewXRegistryReplicaCaller creates a new read-only instance of XRegistryReplica, bound to a specific deployed contract.
func NewXRegistryReplicaCaller(address common.Address, caller bind.ContractCaller) (*XRegistryReplicaCaller, error) {
	contract, err := bindXRegistryReplica(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XRegistryReplicaCaller{contract: contract}, nil
}

// NewXRegistryReplicaTransactor creates a new write-only instance of XRegistryReplica, bound to a specific deployed contract.
func NewXRegistryReplicaTransactor(address common.Address, transactor bind.ContractTransactor) (*XRegistryReplicaTransactor, error) {
	contract, err := bindXRegistryReplica(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XRegistryReplicaTransactor{contract: contract}, nil
}

// NewXRegistryReplicaFilterer creates a new log filterer instance of XRegistryReplica, bound to a specific deployed contract.
func NewXRegistryReplicaFilterer(address common.Address, filterer bind.ContractFilterer) (*XRegistryReplicaFilterer, error) {
	contract, err := bindXRegistryReplica(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XRegistryReplicaFilterer{contract: contract}, nil
}

// bindXRegistryReplica binds a generic wrapper to an already deployed contract.
func bindXRegistryReplica(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := XRegistryReplicaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XRegistryReplica *XRegistryReplicaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XRegistryReplica.Contract.XRegistryReplicaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XRegistryReplica *XRegistryReplicaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XRegistryReplica.Contract.XRegistryReplicaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XRegistryReplica *XRegistryReplicaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XRegistryReplica.Contract.XRegistryReplicaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XRegistryReplica *XRegistryReplicaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XRegistryReplica.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XRegistryReplica *XRegistryReplicaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XRegistryReplica.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XRegistryReplica *XRegistryReplicaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XRegistryReplica.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns((address,bytes))
func (_XRegistryReplica *XRegistryReplicaCaller) Get(opts *bind.CallOpts, chainId uint64, name string, registrant common.Address) (XRegistryBaseDeployment, error) {
	var out []interface{}
	err := _XRegistryReplica.contract.Call(opts, &out, "get", chainId, name, registrant)

	if err != nil {
		return *new(XRegistryBaseDeployment), err
	}

	out0 := *abi.ConvertType(out[0], new(XRegistryBaseDeployment)).(*XRegistryBaseDeployment)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns((address,bytes))
func (_XRegistryReplica *XRegistryReplicaSession) Get(chainId uint64, name string, registrant common.Address) (XRegistryBaseDeployment, error) {
	return _XRegistryReplica.Contract.Get(&_XRegistryReplica.CallOpts, chainId, name, registrant)
}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns((address,bytes))
func (_XRegistryReplica *XRegistryReplicaCallerSession) Get(chainId uint64, name string, registrant common.Address) (XRegistryBaseDeployment, error) {
	return _XRegistryReplica.Contract.Get(&_XRegistryReplica.CallOpts, chainId, name, registrant)
}

// Has is a free data retrieval call binding the contract method 0xfd0b64f7.
//
// Solidity: function has(uint64 chainId, string name, address registrant) view returns(bool)
func (_XRegistryReplica *XRegistryReplicaCaller) Has(opts *bind.CallOpts, chainId uint64, name string, registrant common.Address) (bool, error) {
	var out []interface{}
	err := _XRegistryReplica.contract.Call(opts, &out, "has", chainId, name, registrant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Has is a free data retrieval call binding the contract method 0xfd0b64f7.
//
// Solidity: function has(uint64 chainId, string name, address registrant) view returns(bool)
func (_XRegistryReplica *XRegistryReplicaSession) Has(chainId uint64, name string, registrant common.Address) (bool, error) {
	return _XRegistryReplica.Contract.Has(&_XRegistryReplica.CallOpts, chainId, name, registrant)
}

// Has is a free data retrieval call binding the contract method 0xfd0b64f7.
//
// Solidity: function has(uint64 chainId, string name, address registrant) view returns(bool)
func (_XRegistryReplica *XRegistryReplicaCallerSession) Has(chainId uint64, name string, registrant common.Address) (bool, error) {
	return _XRegistryReplica.Contract.Has(&_XRegistryReplica.CallOpts, chainId, name, registrant)
}

// Set is a paid mutator transaction binding the contract method 0xe4f1c677.
//
// Solidity: function set(uint64 chainId, string name, address registrant, (address,bytes) dep) returns()
func (_XRegistryReplica *XRegistryReplicaTransactor) Set(opts *bind.TransactOpts, chainId uint64, name string, registrant common.Address, dep XRegistryBaseDeployment) (*types.Transaction, error) {
	return _XRegistryReplica.contract.Transact(opts, "set", chainId, name, registrant, dep)
}

// Set is a paid mutator transaction binding the contract method 0xe4f1c677.
//
// Solidity: function set(uint64 chainId, string name, address registrant, (address,bytes) dep) returns()
func (_XRegistryReplica *XRegistryReplicaSession) Set(chainId uint64, name string, registrant common.Address, dep XRegistryBaseDeployment) (*types.Transaction, error) {
	return _XRegistryReplica.Contract.Set(&_XRegistryReplica.TransactOpts, chainId, name, registrant, dep)
}

// Set is a paid mutator transaction binding the contract method 0xe4f1c677.
//
// Solidity: function set(uint64 chainId, string name, address registrant, (address,bytes) dep) returns()
func (_XRegistryReplica *XRegistryReplicaTransactorSession) Set(chainId uint64, name string, registrant common.Address, dep XRegistryBaseDeployment) (*types.Transaction, error) {
	return _XRegistryReplica.Contract.Set(&_XRegistryReplica.TransactOpts, chainId, name, registrant, dep)
}
