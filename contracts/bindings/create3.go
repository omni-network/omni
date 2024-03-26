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
	ABI: "[{\"type\":\"function\",\"name\":\"deploy\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getDeployedAddr\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployedHeight\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610551806100206000396000f3fe6080604052600436106100345760003560e01c80631b5e6d381461003957806373840bff14610076578063cdcb760a146100a4575b600080fd5b34801561004557600080fd5b506100596100543660046103ba565b6100b7565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561008257600080fd5b506100966100913660046103ba565b6100f4565b60405190815260200161006d565b6100596100b2366004610408565b610134565b600082826040516020016100cc9291906104c3565b6040516020818303038152906040528051906020012091506100ed82610185565b9392505050565b600082826040516020016101099291906104c3565b60408051808303601f1901815291815281516020928301206000908152918290529020549392505050565b600033836040516020016101499291906104c3565b60405160208183030381529060405280519060200120925061016c83833461025f565b6000848152602081905260409020439055905092915050565b604080518082018252601081526f67363d3d37363d34f03d5260086018f360801b60209182015290516001600160f81b0319918101919091526001600160601b03193060601b166021820152603581018290527f21c35dbe1b344a2488cf3321d6ce542f8e9f305544ff09e4993a62319a497c1f60558201526000908190610224906075015b6040516020818303038152906040528051906020012090565b6040516135a560f21b60208201526001600160601b0319606083901b166022820152600160f81b60368201529091506100ed9060370161020b565b6000806040518060400160405280601081526020016f67363d3d37363d34f03d5260086018f360801b81525090506000858251602084016000f590506001600160a01b0381166102ea5760405162461bcd60e51b81526020600482015260116024820152701111541313d65351539517d19052531151607a1b60448201526064015b60405180910390fd5b6102f386610185565b92506000816001600160a01b0316858760405161031091906104e0565b60006040518083038185875af1925050503d806000811461034d576040519150601f19603f3d011682016040523d82523d6000602084013e610352565b606091505b5050905080801561036c57506001600160a01b0384163b15155b6103b05760405162461bcd60e51b815260206004820152601560248201527412539255125053125690551253d397d19052531151605a1b60448201526064016102e1565b5050509392505050565b600080604083850312156103cd57600080fd5b82356001600160a01b03811681146103e457600080fd5b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561041b57600080fd5b82359150602083013567ffffffffffffffff8082111561043a57600080fd5b818501915085601f83011261044e57600080fd5b813581811115610460576104606103f2565b604051601f8201601f19908116603f01168101908382118183101715610488576104886103f2565b816040528281528860208487010111156104a157600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60609290921b6001600160601b0319168252601482015260340190565b6000825160005b8181101561050157602081860181015185830152016104e7565b81811115610510576000828501525b50919091019291505056fea2646970667358221220960f3f9115e8e4a72ea67c63cbf8b5fc196508a54561a05ff85e18780b4c151364736f6c634300080c0033",
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

// GetDeployedAddr is a free data retrieval call binding the contract method 0x1b5e6d38.
//
// Solidity: function getDeployedAddr(address deployer, bytes32 salt) view returns(address deployed)
func (_Create3 *Create3Caller) GetDeployedAddr(opts *bind.CallOpts, deployer common.Address, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Create3.contract.Call(opts, &out, "getDeployedAddr", deployer, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDeployedAddr is a free data retrieval call binding the contract method 0x1b5e6d38.
//
// Solidity: function getDeployedAddr(address deployer, bytes32 salt) view returns(address deployed)
func (_Create3 *Create3Session) GetDeployedAddr(deployer common.Address, salt [32]byte) (common.Address, error) {
	return _Create3.Contract.GetDeployedAddr(&_Create3.CallOpts, deployer, salt)
}

// GetDeployedAddr is a free data retrieval call binding the contract method 0x1b5e6d38.
//
// Solidity: function getDeployedAddr(address deployer, bytes32 salt) view returns(address deployed)
func (_Create3 *Create3CallerSession) GetDeployedAddr(deployer common.Address, salt [32]byte) (common.Address, error) {
	return _Create3.Contract.GetDeployedAddr(&_Create3.CallOpts, deployer, salt)
}

// GetDeployedHeight is a free data retrieval call binding the contract method 0x73840bff.
//
// Solidity: function getDeployedHeight(address deployer, bytes32 salt) view returns(uint256)
func (_Create3 *Create3Caller) GetDeployedHeight(opts *bind.CallOpts, deployer common.Address, salt [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Create3.contract.Call(opts, &out, "getDeployedHeight", deployer, salt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeployedHeight is a free data retrieval call binding the contract method 0x73840bff.
//
// Solidity: function getDeployedHeight(address deployer, bytes32 salt) view returns(uint256)
func (_Create3 *Create3Session) GetDeployedHeight(deployer common.Address, salt [32]byte) (*big.Int, error) {
	return _Create3.Contract.GetDeployedHeight(&_Create3.CallOpts, deployer, salt)
}

// GetDeployedHeight is a free data retrieval call binding the contract method 0x73840bff.
//
// Solidity: function getDeployedHeight(address deployer, bytes32 salt) view returns(uint256)
func (_Create3 *Create3CallerSession) GetDeployedHeight(deployer common.Address, salt [32]byte) (*big.Int, error) {
	return _Create3.Contract.GetDeployedHeight(&_Create3.CallOpts, deployer, salt)
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
