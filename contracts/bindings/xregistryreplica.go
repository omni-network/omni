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

// XRegistryReplicaMetaData contains all meta data concerning the XRegistryReplica contract.
var XRegistryReplicaMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"has\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"set\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161084d38038061084d83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516107ae61009f6000396000818160ce0152818161013a015281816101bd015261034d01526107ae6000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063b3449b7714610046578063d7a18fcd14610076578063fd0b64f71461008b575b600080fd5b6100596100543660046105d5565b6100ae565b6040516001600160a01b0390911681526020015b60405180910390f35b610089610084366004610639565b6100c3565b005b61009e6100993660046105d5565b6103b9565b604051901515815260200161006d565b60006100bb8484846103da565b949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101365760405162461bcd60e51b81526020600482015260136024820152721614995c1b1a58d84e881b9bdd081e18d85b1b606a1b60448201526064015b60405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa158015610195573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101b991906106ae565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610219573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023d919061070d565b67ffffffffffffffff16816000015167ffffffffffffffff16146102a35760405162461bcd60e51b815260206004820152601760248201527f585265706c6963613a206e6f742066726f6d206f6d6e69000000000000000000604482015260640161012d565b60208101516001600160a01b031673121e240000000000000000000000000000000001146103135760405162461bcd60e51b815260206004820152601c60248201527f585265706c6963613a206e6f742066726f6d2058526567697374727900000000604482015260640161012d565b61031f85858585610420565b610329848461047d565b156103b257604051632023c96760e11b815267ffffffffffffffff861660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063404792ce90602401600060405180830381600087803b15801561039957600080fd5b505af11580156103ad573d6000803e3d6000fd5b505050505b5050505050565b6000806103c78585856103da565b6001600160a01b03161415949350505050565b67ffffffffffffffff83166000908152602081905260408120816103fe85856104d1565b81526020810191909152604001600020546001600160a01b0316949350505050565b67ffffffffffffffff84166000908152602081905260408120829161044586866104d1565b815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050565b60006104bf6040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b81525073121e2400000000000000000000000000000000026104d1565b6104c984846104d1565b149392505050565b600082826040516020016104e6929190610731565b60405160208183030381529060405280519060200120905092915050565b67ffffffffffffffff8116811461051a57600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261054457600080fd5b813567ffffffffffffffff8082111561055f5761055f61051d565b604051601f8301601f19908116603f011681019082821181831017156105875761058761051d565b816040528381528660208588010111156105a057600080fd5b836020870160208301376000602085830101528094505050505092915050565b6001600160a01b038116811461051a57600080fd5b6000806000606084860312156105ea57600080fd5b83356105f581610504565b9250602084013567ffffffffffffffff81111561061157600080fd5b61061d86828701610533565b925050604084013561062e816105c0565b809150509250925092565b6000806000806080858703121561064f57600080fd5b843561065a81610504565b9350602085013567ffffffffffffffff81111561067657600080fd5b61068287828801610533565b9350506040850135610693816105c0565b915060608501356106a3816105c0565b939692955090935050565b6000604082840312156106c057600080fd5b6040516040810181811067ffffffffffffffff821117156106e3576106e361051d565b60405282516106f181610504565b81526020830151610701816105c0565b60208201529392505050565b60006020828403121561071f57600080fd5b815161072a81610504565b9392505050565b6000835160005b818110156107525760208187018101518583015201610738565b5060609390931b6bffffffffffffffffffffffff1916919092019081526014019291505056fea2646970667358221220e9e20eed3e986cb94a05e519ff8fff658a79ee1de7759cf9cca281f4be86084064736f6c63430008180033",
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
// Solidity: function get(uint64 chainId, string name, address registrant) view returns(address)
func (_XRegistryReplica *XRegistryReplicaCaller) Get(opts *bind.CallOpts, chainId uint64, name string, registrant common.Address) (common.Address, error) {
	var out []interface{}
	err := _XRegistryReplica.contract.Call(opts, &out, "get", chainId, name, registrant)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns(address)
func (_XRegistryReplica *XRegistryReplicaSession) Get(chainId uint64, name string, registrant common.Address) (common.Address, error) {
	return _XRegistryReplica.Contract.Get(&_XRegistryReplica.CallOpts, chainId, name, registrant)
}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns(address)
func (_XRegistryReplica *XRegistryReplicaCallerSession) Get(chainId uint64, name string, registrant common.Address) (common.Address, error) {
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

// Set is a paid mutator transaction binding the contract method 0xd7a18fcd.
//
// Solidity: function set(uint64 chainId, string name, address registrant, address addr) returns()
func (_XRegistryReplica *XRegistryReplicaTransactor) Set(opts *bind.TransactOpts, chainId uint64, name string, registrant common.Address, addr common.Address) (*types.Transaction, error) {
	return _XRegistryReplica.contract.Transact(opts, "set", chainId, name, registrant, addr)
}

// Set is a paid mutator transaction binding the contract method 0xd7a18fcd.
//
// Solidity: function set(uint64 chainId, string name, address registrant, address addr) returns()
func (_XRegistryReplica *XRegistryReplicaSession) Set(chainId uint64, name string, registrant common.Address, addr common.Address) (*types.Transaction, error) {
	return _XRegistryReplica.Contract.Set(&_XRegistryReplica.TransactOpts, chainId, name, registrant, addr)
}

// Set is a paid mutator transaction binding the contract method 0xd7a18fcd.
//
// Solidity: function set(uint64 chainId, string name, address registrant, address addr) returns()
func (_XRegistryReplica *XRegistryReplicaTransactorSession) Set(chainId uint64, name string, registrant common.Address, addr common.Address) (*types.Transaction, error) {
	return _XRegistryReplica.Contract.Set(&_XRegistryReplica.TransactOpts, chainId, name, registrant, addr)
}
