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

// StakingProxyCall is an auto generated low-level Go binding around an user-defined struct.
type StakingProxyCall struct {
	Method    uint8
	Value     *big.Int
	Validator common.Address
	Amount    *big.Int
}

// StakingProxyMetaData contains all meta data concerning the StakingProxy contract.
var StakingProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakingContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"proxy\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structStakingProxy.Call[]\",\"components\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumStakingProxy.Method\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"stakingContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516104e33803806104e383398101604081905261002e91610052565b5f80546001600160a01b0319166001600160a01b039290921691909117905561007f565b5f60208284031215610062575f80fd5b81516001600160a01b0381168114610078575f80fd5b9392505050565b6104578061008c5f395ff3fe60806040526004361061002b575f3560e01c8063b1b69ebc14610036578063ee99205c1461004b575f80fd5b3661003257005b5f80fd5b61004961004436600461033f565b610085565b005b348015610056575f80fd5b505f54610069906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b5f80546001600160a01b0316905b828110156102a2575f8484838181106100ae576100ae6103c2565b6100c492602060809092020190810191506103d6565b60018111156100d5576100d56103ae565b0361018957816001600160a01b0316635c19a95c8585848181106100fb576100fb6103c2565b90506080020160200135868685818110610117576101176103c2565b905060800201604001602081019061012f91906103fb565b6040516001600160e01b031960e085901b1681526001600160a01b0390911660048201526024015f604051808303818588803b15801561016d575f80fd5b505af115801561017f573d5f803e3d5ffd5b505050505061029a565b600184848381811061019d5761019d6103c2565b6101b392602060809092020190810191506103d6565b60018111156101c4576101c46103ae565b0361029a57816001600160a01b0316634d99dd168585848181106101ea576101ea6103c2565b90506080020160200135868685818110610206576102066103c2565b905060800201604001602081019061021e91906103fb565b878786818110610230576102306103c2565b905060800201606001356040518463ffffffff1660e01b815260040161026b9291906001600160a01b03929092168252602082015260400190565b5f604051808303818588803b158015610282575f80fd5b505af1158015610294573d5f803e3d5ffd5b50505050505b600101610093565b50478015610339576040515f90339083908381818185875af1925050503d805f81146102e9576040519150601f19603f3d011682016040523d82523d5f602084013e6102ee565b606091505b50509050806103375760405162461bcd60e51b8152602060048201526011602482015270115512081c99599d5b990819985a5b1959607a1b604482015260640160405180910390fd5b505b50505050565b5f8060208385031215610350575f80fd5b823567ffffffffffffffff80821115610367575f80fd5b818501915085601f83011261037a575f80fd5b813581811115610388575f80fd5b8660208260071b850101111561039c575f80fd5b60209290920196919550909350505050565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156103e6575f80fd5b8135600281106103f4575f80fd5b9392505050565b5f6020828403121561040b575f80fd5b81356001600160a01b03811681146103f4575f80fdfea2646970667358221220ec9bed4eb10164e91643d5a8b127b7d482b23d4a572f44e24c4b5c6c88e19c2d64736f6c63430008180033",
}

// StakingProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingProxyMetaData.ABI instead.
var StakingProxyABI = StakingProxyMetaData.ABI

// StakingProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingProxyMetaData.Bin instead.
var StakingProxyBin = StakingProxyMetaData.Bin

// DeployStakingProxy deploys a new Ethereum contract, binding an instance of StakingProxy to it.
func DeployStakingProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _stakingContract common.Address) (common.Address, *types.Transaction, *StakingProxy, error) {
	parsed, err := StakingProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingProxyBin), backend, _stakingContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingProxy{StakingProxyCaller: StakingProxyCaller{contract: contract}, StakingProxyTransactor: StakingProxyTransactor{contract: contract}, StakingProxyFilterer: StakingProxyFilterer{contract: contract}}, nil
}

// StakingProxy is an auto generated Go binding around an Ethereum contract.
type StakingProxy struct {
	StakingProxyCaller     // Read-only binding to the contract
	StakingProxyTransactor // Write-only binding to the contract
	StakingProxyFilterer   // Log filterer for contract events
}

// StakingProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingProxySession struct {
	Contract     *StakingProxy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingProxyCallerSession struct {
	Contract *StakingProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakingProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingProxyTransactorSession struct {
	Contract     *StakingProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakingProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingProxyRaw struct {
	Contract *StakingProxy // Generic contract binding to access the raw methods on
}

// StakingProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingProxyCallerRaw struct {
	Contract *StakingProxyCaller // Generic read-only contract binding to access the raw methods on
}

// StakingProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingProxyTransactorRaw struct {
	Contract *StakingProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingProxy creates a new instance of StakingProxy, bound to a specific deployed contract.
func NewStakingProxy(address common.Address, backend bind.ContractBackend) (*StakingProxy, error) {
	contract, err := bindStakingProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingProxy{StakingProxyCaller: StakingProxyCaller{contract: contract}, StakingProxyTransactor: StakingProxyTransactor{contract: contract}, StakingProxyFilterer: StakingProxyFilterer{contract: contract}}, nil
}

// NewStakingProxyCaller creates a new read-only instance of StakingProxy, bound to a specific deployed contract.
func NewStakingProxyCaller(address common.Address, caller bind.ContractCaller) (*StakingProxyCaller, error) {
	contract, err := bindStakingProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingProxyCaller{contract: contract}, nil
}

// NewStakingProxyTransactor creates a new write-only instance of StakingProxy, bound to a specific deployed contract.
func NewStakingProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingProxyTransactor, error) {
	contract, err := bindStakingProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingProxyTransactor{contract: contract}, nil
}

// NewStakingProxyFilterer creates a new log filterer instance of StakingProxy, bound to a specific deployed contract.
func NewStakingProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingProxyFilterer, error) {
	contract, err := bindStakingProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingProxyFilterer{contract: contract}, nil
}

// bindStakingProxy binds a generic wrapper to an already deployed contract.
func bindStakingProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingProxy *StakingProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingProxy.Contract.StakingProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingProxy *StakingProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingProxy.Contract.StakingProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingProxy *StakingProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingProxy.Contract.StakingProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingProxy *StakingProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingProxy *StakingProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingProxy *StakingProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingProxy.Contract.contract.Transact(opts, method, params...)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_StakingProxy *StakingProxyCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingProxy.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_StakingProxy *StakingProxySession) StakingContract() (common.Address, error) {
	return _StakingProxy.Contract.StakingContract(&_StakingProxy.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_StakingProxy *StakingProxyCallerSession) StakingContract() (common.Address, error) {
	return _StakingProxy.Contract.StakingContract(&_StakingProxy.CallOpts)
}

// Proxy is a paid mutator transaction binding the contract method 0xb1b69ebc.
//
// Solidity: function proxy((uint8,uint256,address,uint256)[] calls) payable returns()
func (_StakingProxy *StakingProxyTransactor) Proxy(opts *bind.TransactOpts, calls []StakingProxyCall) (*types.Transaction, error) {
	return _StakingProxy.contract.Transact(opts, "proxy", calls)
}

// Proxy is a paid mutator transaction binding the contract method 0xb1b69ebc.
//
// Solidity: function proxy((uint8,uint256,address,uint256)[] calls) payable returns()
func (_StakingProxy *StakingProxySession) Proxy(calls []StakingProxyCall) (*types.Transaction, error) {
	return _StakingProxy.Contract.Proxy(&_StakingProxy.TransactOpts, calls)
}

// Proxy is a paid mutator transaction binding the contract method 0xb1b69ebc.
//
// Solidity: function proxy((uint8,uint256,address,uint256)[] calls) payable returns()
func (_StakingProxy *StakingProxyTransactorSession) Proxy(calls []StakingProxyCall) (*types.Transaction, error) {
	return _StakingProxy.Contract.Proxy(&_StakingProxy.TransactOpts, calls)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingProxy *StakingProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingProxy *StakingProxySession) Receive() (*types.Transaction, error) {
	return _StakingProxy.Contract.Receive(&_StakingProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingProxy *StakingProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingProxy.Contract.Receive(&_StakingProxy.TransactOpts)
}
