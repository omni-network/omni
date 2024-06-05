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
	Bin: "0x60a060405234801561001057600080fd5b50604051610d02380380610d0283398101604081905261002f9161009a565b6001600160a01b0381166100895760405162461bcd60e51b815260206004820181905260248201527f5852656769737472795265706c6963613a206e6f207a65726f20706f7274616c604482015260640160405180910390fd5b6001600160a01b03166080526100ca565b6000602082840312156100ac57600080fd5b81516001600160a01b03811681146100c357600080fd5b9392505050565b608051610c176100eb6000396000818160d101526103370152610c176000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063b3449b7714610046578063e4f1c6771461006f578063fd0b64f714610084575b600080fd5b6100596100543660046106db565b6100a7565b6040516100669190610762565b60405180910390f35b61008261007d3660046107b0565b6100cf565b005b6100976100923660046106db565b6103aa565b6040519015158152602001610066565b6040805180820190915260008152606060208201526100c78484846103cc565b949350505050565b7f0000000000000000000000000000000000000000000000000000000000000000336001600160a01b038216146101435760405162461bcd60e51b81526020600482015260136024820152721614995c1b1a58d84e881b9bdd081e18d85b1b606a1b60448201526064015b60405180910390fd5b6000816001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa158015610182573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101a6919061083f565b9050816001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020a9190610880565b6001600160401b031681600001516001600160401b03161461026e5760405162461bcd60e51b815260206004820152601760248201527f585265706c6963613a206e6f742066726f6d206f6d6e69000000000000000000604482015260640161013a565b60208101516001600160a01b031673121e240000000000000000000000000000000001146102de5760405162461bcd60e51b815260206004820152601c60248201527f585265706c6963613a206e6f742066726f6d2058526567697374727900000000604482015260640161013a565b6102f28686866102ed8761089d565b6104ce565b6102fc858561053b565b156103a25760006103106020850185610909565b81019061031d9190610956565b604051633ba5ccd560e01b81529091506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ba5ccd59061036e908a908590600401610a07565b600060405180830381600087803b15801561038857600080fd5b505af115801561039c573d6000803e3d6000fd5b50505050505b505050505050565b6000806103b88585856103cc565b516001600160a01b03161415949350505050565b6040805180820190915260008152606060208201526001600160401b038416600090815260208190526040812090610404858561058f565b81526020808201929092526040908101600020815180830190925280546001600160a01b03168252600181018054929391929184019161044390610a64565b80601f016020809104026020016040519081016040528092919081815260200182805461046f90610a64565b80156104bc5780601f10610491576101008083540402835291602001916104bc565b820191906000526020600020905b81548152906001019060200180831161049f57829003601f168201915b50505050508152505090509392505050565b6001600160401b038416600090815260208190526040812082916104f2868661058f565b8152602080820192909252604001600020825181546001600160a01b0319166001600160a01b039091161781559082015160018201906105329082610aeb565b50505050505050565b600061057d6040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b81525073121e24000000000000000000000000000000000261058f565b610587848461058f565b149392505050565b600082826040516020016105a4929190610baa565b60405160208183030381529060405280519060200120905092915050565b6001600160401b03811681146105d757600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715610612576106126105da565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610640576106406105da565b604052919050565b60006001600160401b03831115610661576106616105da565b610674601f8401601f1916602001610618565b905082815283838301111561068857600080fd5b828260208301376000602084830101529392505050565b600082601f8301126106b057600080fd5b6106bf83833560208501610648565b9392505050565b6001600160a01b03811681146105d757600080fd5b6000806000606084860312156106f057600080fd5b83356106fb816105c2565b925060208401356001600160401b0381111561071657600080fd5b6107228682870161069f565b9250506040840135610733816106c6565b809150509250925092565b60005b83811015610759578181015183820152602001610741565b50506000910152565b6020815260018060a01b03825116602082015260006020830151604080840152805180606085015261079b81608086016020850161073e565b601f01601f1916929092016080019392505050565b600080600080608085870312156107c657600080fd5b84356107d1816105c2565b935060208501356001600160401b03808211156107ed57600080fd5b6107f98883890161069f565b94506040870135915061080b826106c6565b9092506060860135908082111561082157600080fd5b5085016040818803121561083457600080fd5b939692955090935050565b60006040828403121561085157600080fd5b6108596105f0565b8251610864816105c2565b81526020830151610874816106c6565b60208201529392505050565b60006020828403121561089257600080fd5b81516106bf816105c2565b6000604082360312156108af57600080fd5b6108b76105f0565b82356108c2816106c6565b815260208301356001600160401b038111156108dd57600080fd5b830136601f8201126108ee57600080fd5b6108fd36823560208401610648565b60208301525092915050565b6000808335601e1984360301811261092057600080fd5b8301803591506001600160401b0382111561093a57600080fd5b60200191503681900382131561094f57600080fd5b9250929050565b6000602080838503121561096957600080fd5b82356001600160401b038082111561098057600080fd5b818501915085601f83011261099457600080fd5b8135818111156109a6576109a66105da565b8060051b91506109b7848301610618565b81815291830184019184810190888411156109d157600080fd5b938501935b838510156109fb57843592506109eb836105c2565b82825293850193908501906109d6565b98975050505050505050565b6000604082016001600160401b03808616845260206040602086015282865180855260608701915060208801945060005b81811015610a56578551851683529483019491830191600101610a38565b509098975050505050505050565b600181811c90821680610a7857607f821691505b602082108103610a9857634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610ae6576000816000526020600020601f850160051c81016020861015610ac75750805b601f850160051c820191505b818110156103a257828155600101610ad3565b505050565b81516001600160401b03811115610b0457610b046105da565b610b1881610b128454610a64565b84610a9e565b602080601f831160018114610b4d5760008415610b355750858301515b600019600386901b1c1916600185901b1785556103a2565b600085815260208120601f198616915b82811015610b7c57888601518255948401946001909101908401610b5d565b5085821015610b9a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008351610bbc81846020880161073e565b60609390931b6bffffffffffffffffffffffff1916919092019081526014019291505056fea2646970667358221220d41b26f549a53da6077155258aff0872e62f80d86fd7080b2655e1fc425cbc5064736f6c63430008180033",
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
