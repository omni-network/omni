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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"has\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"set\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dep\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610cf8380380610cf883398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610c5961009f6000396000818160da01528181610146015281816101c901526103760152610c596000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063b3449b7714610046578063e4f1c6771461006f578063fd0b64f714610084575b600080fd5b610059610054366004610719565b6100a7565b60405161006691906107a0565b60405180910390f35b61008261007d3660046107ee565b6100cf565b005b610097610092366004610719565b6103e8565b6040519015158152602001610066565b6040805180820190915260008152606060208201526100c784848461040a565b949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101425760405162461bcd60e51b81526020600482015260136024820152721614995c1b1a58d84e881b9bdd081e18d85b1b606a1b60448201526064015b60405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa1580156101a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101c5919061087d565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610225573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024991906108be565b6001600160401b031681600001516001600160401b0316146102ad5760405162461bcd60e51b815260206004820152601760248201527f585265706c6963613a206e6f742066726f6d206f6d6e690000000000000000006044820152606401610139565b60208101516001600160a01b031673121e2400000000000000000000000000000000011461031d5760405162461bcd60e51b815260206004820152601c60248201527f585265706c6963613a206e6f742066726f6d20585265676973747279000000006044820152606401610139565b61033185858561032c866108db565b61050c565b61033b8484610579565b156103e157600061034f6020840184610947565b81019061035c9190610994565b604051633ba5ccd560e01b81529091506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ba5ccd5906103ad9089908590600401610a45565b600060405180830381600087803b1580156103c757600080fd5b505af11580156103db573d6000803e3d6000fd5b50505050505b5050505050565b6000806103f685858561040a565b516001600160a01b03161415949350505050565b6040805180820190915260008152606060208201526001600160401b03841660009081526020819052604081209061044285856105cd565b81526020808201929092526040908101600020815180830190925280546001600160a01b03168252600181018054929391929184019161048190610aa2565b80601f01602080910402602001604051908101604052809291908181526020018280546104ad90610aa2565b80156104fa5780601f106104cf576101008083540402835291602001916104fa565b820191906000526020600020905b8154815290600101906020018083116104dd57829003601f168201915b50505050508152505090509392505050565b6001600160401b0384166000908152602081905260408120829161053086866105cd565b8152602080820192909252604001600020825181546001600160a01b0319166001600160a01b039091161781559082015160018201906105709082610b2d565b50505050505050565b60006105bb6040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b81525073121e2400000000000000000000000000000000026105cd565b6105c584846105cd565b149392505050565b600082826040516020016105e2929190610bec565b60405160208183030381529060405280519060200120905092915050565b6001600160401b038116811461061557600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561065057610650610618565b60405290565b604051601f8201601f191681016001600160401b038111828210171561067e5761067e610618565b604052919050565b60006001600160401b0383111561069f5761069f610618565b6106b2601f8401601f1916602001610656565b90508281528383830111156106c657600080fd5b828260208301376000602084830101529392505050565b600082601f8301126106ee57600080fd5b6106fd83833560208501610686565b9392505050565b6001600160a01b038116811461061557600080fd5b60008060006060848603121561072e57600080fd5b833561073981610600565b925060208401356001600160401b0381111561075457600080fd5b610760868287016106dd565b925050604084013561077181610704565b809150509250925092565b60005b8381101561079757818101518382015260200161077f565b50506000910152565b6020815260018060a01b0382511660208201526000602083015160408084015280518060608501526107d981608086016020850161077c565b601f01601f1916929092016080019392505050565b6000806000806080858703121561080457600080fd5b843561080f81610600565b935060208501356001600160401b038082111561082b57600080fd5b610837888389016106dd565b94506040870135915061084982610704565b9092506060860135908082111561085f57600080fd5b5085016040818803121561087257600080fd5b939692955090935050565b60006040828403121561088f57600080fd5b61089761062e565b82516108a281610600565b815260208301516108b281610704565b60208201529392505050565b6000602082840312156108d057600080fd5b81516106fd81610600565b6000604082360312156108ed57600080fd5b6108f561062e565b823561090081610704565b815260208301356001600160401b0381111561091b57600080fd5b830136601f82011261092c57600080fd5b61093b36823560208401610686565b60208301525092915050565b6000808335601e1984360301811261095e57600080fd5b8301803591506001600160401b0382111561097857600080fd5b60200191503681900382131561098d57600080fd5b9250929050565b600060208083850312156109a757600080fd5b82356001600160401b03808211156109be57600080fd5b818501915085601f8301126109d257600080fd5b8135818111156109e4576109e4610618565b8060051b91506109f5848301610656565b8181529183018401918481019088841115610a0f57600080fd5b938501935b83851015610a395784359250610a2983610600565b8282529385019390850190610a14565b98975050505050505050565b6000604082016001600160401b03808616845260206040602086015282865180855260608701915060208801945060005b81811015610a94578551851683529483019491830191600101610a76565b509098975050505050505050565b600181811c90821680610ab657607f821691505b602082108103610ad657634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610b28576000816000526020600020601f850160051c81016020861015610b055750805b601f850160051c820191505b81811015610b2457828155600101610b11565b5050505b505050565b81516001600160401b03811115610b4657610b46610618565b610b5a81610b548454610aa2565b84610adc565b602080601f831160018114610b8f5760008415610b775750858301515b600019600386901b1c1916600185901b178555610b24565b600085815260208120601f198616915b82811015610bbe57888601518255948401946001909101908401610b9f565b5085821015610bdc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008351610bfe81846020880161077c565b60609390931b6bffffffffffffffffffffffff1916919092019081526014019291505056fea2646970667358221220ea8138abf6576bd9a467e7775645e73e50b9d2a54b73fb453e84bffdf22d06a464736f6c63430008180033",
}

// XRegistryReplicaABI is the input ABI used to generate the binding from.
// Deprecated: Use XRegistryReplicaMetaData.ABI instead.
var XRegistryReplicaABI = XRegistryReplicaMetaData.ABI

// XRegistryReplicaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use XRegistryReplicaMetaData.Bin instead.
var XRegistryReplicaBin = XRegistryReplicaMetaData.Bin

// DeployXRegistryReplica deploys a new Ethereum contract, binding an instance of XRegistryReplica to it.
func DeployXRegistryReplica(auth *bind.TransactOpts, backend bind.ContractBackend, omni_ common.Address) (common.Address, *types.Transaction, *XRegistryReplica, error) {
	parsed, err := XRegistryReplicaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(XRegistryReplicaBin), backend, omni_)
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
