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

// Create3MetaData contains all meta data concerning the Create3 contract.
var Create3MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deploy\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getDeployed\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506104b18061001d5f395ff3fe608060405260043610610028575f3560e01c806350f1c4641461002c578063cdcb760a14610067575b5f80fd5b348015610037575f80fd5b5061004b610046366004610351565b61007a565b6040516001600160a01b03909116815260200160405180910390f35b61004b61007536600461039a565b6100c5565b6040516001600160601b0319606084901b166020820152603481018290525f906054016040516020818303038152906040528051906020012091506100be8261010a565b9392505050565b6040516001600160601b03193360601b166020820152603481018390525f906054016040516020818303038152906040528051906020012092506100be83833461011b565b5f610115823061026f565b92915050565b5f806040518060400160405280601081526020016f67363d3d37363d34f03d5260086018f360801b81525090505f858251602084015ff590506001600160a01b0381166101a35760405162461bcd60e51b81526020600482015260116024820152701111541313d65351539517d19052531151607a1b60448201526064015b60405180910390fd5b6101ac8661010a565b92505f816001600160a01b031685876040516101c8919061044f565b5f6040518083038185875af1925050503d805f8114610202576040519150601f19603f3d011682016040523d82523d5f602084013e610207565b606091505b5050905080801561022157506001600160a01b0384163b15155b6102655760405162461bcd60e51b815260206004820152601560248201527412539255125053125690551253d397d19052531151605a1b604482015260640161019a565b5050509392505050565b604080518082018252601081526f67363d3d37363d34f03d5260086018f360801b60209182015290516001600160f81b0319918101919091526001600160601b0319606083901b166021820152603581018390527f21c35dbe1b344a2488cf3321d6ce542f8e9f305544ff09e4993a62319a497c1f60558201525f90819061030e906075015b6040516020818303038152906040528051906020012090565b6040516135a560f21b60208201526001600160601b0319606083901b166022820152600160f81b6036820152909150610349906037016102f5565b949350505050565b5f8060408385031215610362575f80fd5b82356001600160a01b0381168114610378575f80fd5b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f80604083850312156103ab575f80fd5b82359150602083013567ffffffffffffffff808211156103c9575f80fd5b818501915085601f8301126103dc575f80fd5b8135818111156103ee576103ee610386565b604051601f8201601f19908116603f0116810190838211818310171561041657610416610386565b8160405282815288602084870101111561042e575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f82515f5b8181101561046e5760208186018101518583015201610454565b505f92019182525091905056fea2646970667358221220a35268b33df8d9d434830057faa2da513798dc47d56f109b1622fcd2ecda96bf64736f6c63430008180033",
}

// Create3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Create3MetaData.ABI instead.
var Create3ABI = Create3MetaData.ABI

// Create3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Create3MetaData.Bin instead.
var Create3Bin = Create3MetaData.Bin

// DeployCreate3 deploys a new Ethereum contract, binding an instance of Create3 to it.
func DeployCreate3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Create3, error) {
	parsed, err := Create3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Create3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Create3{Create3Caller: Create3Caller{contract: contract}, Create3Transactor: Create3Transactor{contract: contract}, Create3Filterer: Create3Filterer{contract: contract}}, nil
}

// Create3 is an auto generated Go binding around an Ethereum contract.
type Create3 struct {
	Create3Caller     // Read-only binding to the contract
	Create3Transactor // Write-only binding to the contract
	Create3Filterer   // Log filterer for contract events
}

// Create3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Create3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Create3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Create3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Create3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Create3Session struct {
	Contract     *Create3          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Create3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Create3CallerSession struct {
	Contract *Create3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Create3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Create3TransactorSession struct {
	Contract     *Create3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Create3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Create3Raw struct {
	Contract *Create3 // Generic contract binding to access the raw methods on
}

// Create3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Create3CallerRaw struct {
	Contract *Create3Caller // Generic read-only contract binding to access the raw methods on
}

// Create3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Create3TransactorRaw struct {
	Contract *Create3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreate3 creates a new instance of Create3, bound to a specific deployed contract.
func NewCreate3(address common.Address, backend bind.ContractBackend) (*Create3, error) {
	contract, err := bindCreate3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Create3{Create3Caller: Create3Caller{contract: contract}, Create3Transactor: Create3Transactor{contract: contract}, Create3Filterer: Create3Filterer{contract: contract}}, nil
}

// NewCreate3Caller creates a new read-only instance of Create3, bound to a specific deployed contract.
func NewCreate3Caller(address common.Address, caller bind.ContractCaller) (*Create3Caller, error) {
	contract, err := bindCreate3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Create3Caller{contract: contract}, nil
}

// NewCreate3Transactor creates a new write-only instance of Create3, bound to a specific deployed contract.
func NewCreate3Transactor(address common.Address, transactor bind.ContractTransactor) (*Create3Transactor, error) {
	contract, err := bindCreate3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Create3Transactor{contract: contract}, nil
}

// NewCreate3Filterer creates a new log filterer instance of Create3, bound to a specific deployed contract.
func NewCreate3Filterer(address common.Address, filterer bind.ContractFilterer) (*Create3Filterer, error) {
	contract, err := bindCreate3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Create3Filterer{contract: contract}, nil
}

// bindCreate3 binds a generic wrapper to an already deployed contract.
func bindCreate3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Create3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Create3 *Create3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Create3.Contract.Create3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Create3 *Create3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Create3.Contract.Create3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Create3 *Create3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Create3.Contract.Create3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Create3 *Create3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Create3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Create3 *Create3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Create3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Create3 *Create3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Create3.Contract.contract.Transact(opts, method, params...)
}

// GetDeployed is a free data retrieval call binding the contract method 0x50f1c464.
//
// Solidity: function getDeployed(address deployer, bytes32 salt) view returns(address deployed)
func (_Create3 *Create3Caller) GetDeployed(opts *bind.CallOpts, deployer common.Address, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Create3.contract.Call(opts, &out, "getDeployed", deployer, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDeployed is a free data retrieval call binding the contract method 0x50f1c464.
//
// Solidity: function getDeployed(address deployer, bytes32 salt) view returns(address deployed)
func (_Create3 *Create3Session) GetDeployed(deployer common.Address, salt [32]byte) (common.Address, error) {
	return _Create3.Contract.GetDeployed(&_Create3.CallOpts, deployer, salt)
}

// GetDeployed is a free data retrieval call binding the contract method 0x50f1c464.
//
// Solidity: function getDeployed(address deployer, bytes32 salt) view returns(address deployed)
func (_Create3 *Create3CallerSession) GetDeployed(deployer common.Address, salt [32]byte) (common.Address, error) {
	return _Create3.Contract.GetDeployed(&_Create3.CallOpts, deployer, salt)
}

// Deploy is a paid mutator transaction binding the contract method 0xcdcb760a.
//
// Solidity: function deploy(bytes32 salt, bytes creationCode) payable returns(address deployed)
func (_Create3 *Create3Transactor) Deploy(opts *bind.TransactOpts, salt [32]byte, creationCode []byte) (*types.Transaction, error) {
	return _Create3.contract.Transact(opts, "deploy", salt, creationCode)
}

// Deploy is a paid mutator transaction binding the contract method 0xcdcb760a.
//
// Solidity: function deploy(bytes32 salt, bytes creationCode) payable returns(address deployed)
func (_Create3 *Create3Session) Deploy(salt [32]byte, creationCode []byte) (*types.Transaction, error) {
	return _Create3.Contract.Deploy(&_Create3.TransactOpts, salt, creationCode)
}

// Deploy is a paid mutator transaction binding the contract method 0xcdcb760a.
//
// Solidity: function deploy(bytes32 salt, bytes creationCode) payable returns(address deployed)
func (_Create3 *Create3TransactorSession) Deploy(salt [32]byte, creationCode []byte) (*types.Transaction, error) {
	return _Create3.Contract.Deploy(&_Create3.TransactOpts, salt, creationCode)
}
