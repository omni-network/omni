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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakingContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegateN\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"n\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"proxy\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structStakingProxy.Call[]\",\"components\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumStakingProxy.Method\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"stakingContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegateN\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"n\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161069d38038061069d83398101604081905261002e91610052565b5f80546001600160a01b0319166001600160a01b039290921691909117905561007f565b5f60208284031215610062575f80fd5b81516001600160a01b0381168114610078575f80fd5b9392505050565b6106118061008c5f395ff3fe608060405260043610610041575f3560e01c806325db2ab51461004c578063557707fa14610061578063b1b69ebc14610074578063ee99205c14610087575f80fd5b3661004857005b5f80fd5b61005f61005a3660046104a0565b6100c1565b005b61005f61006f3660046104d0565b610143565b61005f610082366004610506565b6101cd565b348015610092575f80fd5b505f546100a5906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b5f80546001600160a01b0316905b8281101561013c576040516317066a5760e21b81526001600160a01b038681166004830152831690635c19a95c9086906024015f604051808303818588803b158015610119575f80fd5b505af115801561012b573d5f803e3d5ffd5b5050600190930192506100cf915050565b5050505050565b5f80546001600160a01b0316905b828110156101c5576040516326ccee8b60e11b81526001600160a01b03878116600483015260248201869052831690634d99dd169087906044015f604051808303818588803b1580156101a2575f80fd5b505af11580156101b4573d5f803e3d5ffd5b505060019093019250610151915050565b505050505050565b5f80546001600160a01b0316905b828110156103ea575f8484838181106101f6576101f6610589565b61020c926020608090920201908101915061059d565b600181111561021d5761021d610575565b036102d157816001600160a01b0316635c19a95c85858481811061024357610243610589565b9050608002016020013586868581811061025f5761025f610589565b905060800201604001602081019061027791906105c2565b6040516001600160e01b031960e085901b1681526001600160a01b0390911660048201526024015f604051808303818588803b1580156102b5575f80fd5b505af11580156102c7573d5f803e3d5ffd5b50505050506103e2565b60018484838181106102e5576102e5610589565b6102fb926020608090920201908101915061059d565b600181111561030c5761030c610575565b036103e257816001600160a01b0316634d99dd1685858481811061033257610332610589565b9050608002016020013586868581811061034e5761034e610589565b905060800201604001602081019061036691906105c2565b87878681811061037857610378610589565b905060800201606001356040518463ffffffff1660e01b81526004016103b39291906001600160a01b03929092168252602082015260400190565b5f604051808303818588803b1580156103ca575f80fd5b505af11580156103dc573d5f803e3d5ffd5b50505050505b6001016101db565b5047801561047f576040515f90339083908381818185875af1925050503d805f8114610431576040519150601f19603f3d011682016040523d82523d5f602084013e610436565b606091505b505090508061013c5760405162461bcd60e51b8152602060048201526011602482015270115512081c99599d5b990819985a5b1959607a1b604482015260640160405180910390fd5b50505050565b80356001600160a01b038116811461049b575f80fd5b919050565b5f805f606084860312156104b2575f80fd5b6104bb84610485565b95602085013595506040909401359392505050565b5f805f80608085870312156104e3575f80fd5b6104ec85610485565b966020860135965060408601359560600135945092505050565b5f8060208385031215610517575f80fd5b823567ffffffffffffffff8082111561052e575f80fd5b818501915085601f830112610541575f80fd5b81358181111561054f575f80fd5b8660208260071b8501011115610563575f80fd5b60209290920196919550909350505050565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156105ad575f80fd5b8135600281106105bb575f80fd5b9392505050565b5f602082840312156105d2575f80fd5b6105bb8261048556fea264697066735822122069bf7e7df8968030abfdfa4da9d70d02852b8d5a15f32085e71d0fbfe14134f364736f6c63430008180033",
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

// DelegateN is a paid mutator transaction binding the contract method 0x25db2ab5.
//
// Solidity: function delegateN(address validator, uint256 value, uint256 n) payable returns()
func (_StakingProxy *StakingProxyTransactor) DelegateN(opts *bind.TransactOpts, validator common.Address, value *big.Int, n *big.Int) (*types.Transaction, error) {
	return _StakingProxy.contract.Transact(opts, "delegateN", validator, value, n)
}

// DelegateN is a paid mutator transaction binding the contract method 0x25db2ab5.
//
// Solidity: function delegateN(address validator, uint256 value, uint256 n) payable returns()
func (_StakingProxy *StakingProxySession) DelegateN(validator common.Address, value *big.Int, n *big.Int) (*types.Transaction, error) {
	return _StakingProxy.Contract.DelegateN(&_StakingProxy.TransactOpts, validator, value, n)
}

// DelegateN is a paid mutator transaction binding the contract method 0x25db2ab5.
//
// Solidity: function delegateN(address validator, uint256 value, uint256 n) payable returns()
func (_StakingProxy *StakingProxyTransactorSession) DelegateN(validator common.Address, value *big.Int, n *big.Int) (*types.Transaction, error) {
	return _StakingProxy.Contract.DelegateN(&_StakingProxy.TransactOpts, validator, value, n)
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

// UndelegateN is a paid mutator transaction binding the contract method 0x557707fa.
//
// Solidity: function undelegateN(address validator, uint256 value, uint256 amount, uint256 n) payable returns()
func (_StakingProxy *StakingProxyTransactor) UndelegateN(opts *bind.TransactOpts, validator common.Address, value *big.Int, amount *big.Int, n *big.Int) (*types.Transaction, error) {
	return _StakingProxy.contract.Transact(opts, "undelegateN", validator, value, amount, n)
}

// UndelegateN is a paid mutator transaction binding the contract method 0x557707fa.
//
// Solidity: function undelegateN(address validator, uint256 value, uint256 amount, uint256 n) payable returns()
func (_StakingProxy *StakingProxySession) UndelegateN(validator common.Address, value *big.Int, amount *big.Int, n *big.Int) (*types.Transaction, error) {
	return _StakingProxy.Contract.UndelegateN(&_StakingProxy.TransactOpts, validator, value, amount, n)
}

// UndelegateN is a paid mutator transaction binding the contract method 0x557707fa.
//
// Solidity: function undelegateN(address validator, uint256 value, uint256 amount, uint256 n) payable returns()
func (_StakingProxy *StakingProxyTransactorSession) UndelegateN(validator common.Address, value *big.Int, amount *big.Int, n *big.Int) (*types.Transaction, error) {
	return _StakingProxy.Contract.UndelegateN(&_StakingProxy.TransactOpts, validator, value, amount, n)
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
