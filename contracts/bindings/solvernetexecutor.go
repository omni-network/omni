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

// SolverNetExecutorMetaData contains all meta data concerning the SolverNetExecutor contract.
var SolverNetExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_outbox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"outbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tryRevokeApproval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOutbox\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060405161070138038061070183398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f80fd5b81516001600160a01b0381168114610065575f80fd5b9392505050565b60805161065b6100a65f395f818160df01528181610147015281816101a70152818161026401528181610331015261038e015261065b5ff3fe608060405260043610610055575f3560e01c80637d2e90c21461005e5780639d32b5691461007d578063b61d27f61461009c578063beabacc8146100af578063ce11e6ab146100ce578063e1f21c671461011d57005b3661005c57005b005b348015610069575f80fd5b5061005c6100783660046104de565b61013c565b348015610088575f80fd5b5061005c610097366004610506565b61019c565b61005c6100aa366004610537565b610259565b3480156100ba575f80fd5b5061005c6100c93660046105b7565b610326565b3480156100d9575f80fd5b506101017f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b348015610128575f80fd5b5061005c6101373660046105b7565b610383565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101855760405163bda8fc9560e01b815260040160405180910390fd5b6101986001600160a01b038316826103e0565b5050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101e55760405163bda8fc9560e01b815260040160405180910390fd5b60405163095ea7b360e01b81526001600160a01b0382811660048301525f602483015283169063095ea7b3906044016020604051808303815f875af192505050801561024e575060408051601f3d908101601f1916820190925261024b918101906105f0565b60015b15610198575b505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146102a25760405163bda8fc9560e01b815260040160405180910390fd5b5f846001600160a01b03168484846040516102be929190610616565b5f6040518083038185875af1925050503d805f81146102f8576040519150601f19603f3d011682016040523d82523d5f602084013e6102fd565b606091505b505090508061031f57604051633204506f60e01b815260040160405180910390fd5b5050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461036f5760405163bda8fc9560e01b815260040160405180910390fd5b6102546001600160a01b03841683836103f9565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103cc5760405163bda8fc9560e01b815260040160405180910390fd5b6102546001600160a01b0384168383610443565b5f385f3884865af16101985763b12d13eb5f526004601cfd5b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f51141661043957803d853b151710610439576390b8ec185f526004601cfd5b505f603452505050565b816014528060345263095ea7b360601b5f5260205f604460105f875af18060015f51141661043957803d853b151710610439575f60345263095ea7b360601b5f525f38604460105f885af1508160345260205f604460105f885af190508060015f51141661043957803d853b15171061043957633e3f8f735f526004601cfd5b80356001600160a01b03811681146104d9575f80fd5b919050565b5f80604083850312156104ef575f80fd5b6104f8836104c3565b946020939093013593505050565b5f8060408385031215610517575f80fd5b610520836104c3565b915061052e602084016104c3565b90509250929050565b5f805f806060858703121561054a575f80fd5b610553856104c3565b935060208501359250604085013567ffffffffffffffff80821115610576575f80fd5b818701915087601f830112610589575f80fd5b813581811115610597575f80fd5b8860208285010111156105a8575f80fd5b95989497505060200194505050565b5f805f606084860312156105c9575f80fd5b6105d2846104c3565b92506105e0602085016104c3565b9150604084013590509250925092565b5f60208284031215610600575f80fd5b8151801515811461060f575f80fd5b9392505050565b818382375f910190815291905056fea2646970667358221220a55adfd1d2b47b92a3e5b5c8accc2420164d9929c9eb0873789ce323ac7d9d6864736f6c63430008180033",
}

// SolverNetExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use SolverNetExecutorMetaData.ABI instead.
var SolverNetExecutorABI = SolverNetExecutorMetaData.ABI

// SolverNetExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolverNetExecutorMetaData.Bin instead.
var SolverNetExecutorBin = SolverNetExecutorMetaData.Bin

// DeploySolverNetExecutor deploys a new Ethereum contract, binding an instance of SolverNetExecutor to it.
func DeploySolverNetExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, _outbox common.Address) (common.Address, *types.Transaction, *SolverNetExecutor, error) {
	parsed, err := SolverNetExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolverNetExecutorBin), backend, _outbox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolverNetExecutor{SolverNetExecutorCaller: SolverNetExecutorCaller{contract: contract}, SolverNetExecutorTransactor: SolverNetExecutorTransactor{contract: contract}, SolverNetExecutorFilterer: SolverNetExecutorFilterer{contract: contract}}, nil
}

// SolverNetExecutor is an auto generated Go binding around an Ethereum contract.
type SolverNetExecutor struct {
	SolverNetExecutorCaller     // Read-only binding to the contract
	SolverNetExecutorTransactor // Write-only binding to the contract
	SolverNetExecutorFilterer   // Log filterer for contract events
}

// SolverNetExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolverNetExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolverNetExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolverNetExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolverNetExecutorSession struct {
	Contract     *SolverNetExecutor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SolverNetExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolverNetExecutorCallerSession struct {
	Contract *SolverNetExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SolverNetExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolverNetExecutorTransactorSession struct {
	Contract     *SolverNetExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SolverNetExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolverNetExecutorRaw struct {
	Contract *SolverNetExecutor // Generic contract binding to access the raw methods on
}

// SolverNetExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolverNetExecutorCallerRaw struct {
	Contract *SolverNetExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// SolverNetExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolverNetExecutorTransactorRaw struct {
	Contract *SolverNetExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolverNetExecutor creates a new instance of SolverNetExecutor, bound to a specific deployed contract.
func NewSolverNetExecutor(address common.Address, backend bind.ContractBackend) (*SolverNetExecutor, error) {
	contract, err := bindSolverNetExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolverNetExecutor{SolverNetExecutorCaller: SolverNetExecutorCaller{contract: contract}, SolverNetExecutorTransactor: SolverNetExecutorTransactor{contract: contract}, SolverNetExecutorFilterer: SolverNetExecutorFilterer{contract: contract}}, nil
}

// NewSolverNetExecutorCaller creates a new read-only instance of SolverNetExecutor, bound to a specific deployed contract.
func NewSolverNetExecutorCaller(address common.Address, caller bind.ContractCaller) (*SolverNetExecutorCaller, error) {
	contract, err := bindSolverNetExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetExecutorCaller{contract: contract}, nil
}

// NewSolverNetExecutorTransactor creates a new write-only instance of SolverNetExecutor, bound to a specific deployed contract.
func NewSolverNetExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*SolverNetExecutorTransactor, error) {
	contract, err := bindSolverNetExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetExecutorTransactor{contract: contract}, nil
}

// NewSolverNetExecutorFilterer creates a new log filterer instance of SolverNetExecutor, bound to a specific deployed contract.
func NewSolverNetExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*SolverNetExecutorFilterer, error) {
	contract, err := bindSolverNetExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolverNetExecutorFilterer{contract: contract}, nil
}

// bindSolverNetExecutor binds a generic wrapper to an already deployed contract.
func bindSolverNetExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolverNetExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetExecutor *SolverNetExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetExecutor.Contract.SolverNetExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetExecutor *SolverNetExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.SolverNetExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetExecutor *SolverNetExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.SolverNetExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetExecutor *SolverNetExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetExecutor *SolverNetExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetExecutor *SolverNetExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.contract.Transact(opts, method, params...)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_SolverNetExecutor *SolverNetExecutorCaller) Outbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetExecutor.contract.Call(opts, &out, "outbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_SolverNetExecutor *SolverNetExecutorSession) Outbox() (common.Address, error) {
	return _SolverNetExecutor.Contract.Outbox(&_SolverNetExecutor.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_SolverNetExecutor *SolverNetExecutorCallerSession) Outbox() (common.Address, error) {
	return _SolverNetExecutor.Contract.Outbox(&_SolverNetExecutor.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.Transact(opts, "approve", token, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorSession) Approve(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Approve(&_SolverNetExecutor.TransactOpts, token, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) Approve(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Approve(&_SolverNetExecutor.TransactOpts, token, spender, amount)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.Transact(opts, "execute", target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Execute(&_SolverNetExecutor.TransactOpts, target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Execute(&_SolverNetExecutor.TransactOpts, target, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) Transfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.Transact(opts, "transfer", token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Transfer(&_SolverNetExecutor.TransactOpts, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Transfer(&_SolverNetExecutor.TransactOpts, token, to, amount)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address to, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) TransferNative(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.Transact(opts, "transferNative", to, amount)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address to, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorSession) TransferNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.TransferNative(&_SolverNetExecutor.TransactOpts, to, amount)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address to, uint256 amount) returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) TransferNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.TransferNative(&_SolverNetExecutor.TransactOpts, to, amount)
}

// TryRevokeApproval is a paid mutator transaction binding the contract method 0x9d32b569.
//
// Solidity: function tryRevokeApproval(address token, address spender) returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) TryRevokeApproval(opts *bind.TransactOpts, token common.Address, spender common.Address) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.Transact(opts, "tryRevokeApproval", token, spender)
}

// TryRevokeApproval is a paid mutator transaction binding the contract method 0x9d32b569.
//
// Solidity: function tryRevokeApproval(address token, address spender) returns()
func (_SolverNetExecutor *SolverNetExecutorSession) TryRevokeApproval(token common.Address, spender common.Address) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.TryRevokeApproval(&_SolverNetExecutor.TransactOpts, token, spender)
}

// TryRevokeApproval is a paid mutator transaction binding the contract method 0x9d32b569.
//
// Solidity: function tryRevokeApproval(address token, address spender) returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) TryRevokeApproval(token common.Address, spender common.Address) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.TryRevokeApproval(&_SolverNetExecutor.TransactOpts, token, spender)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetExecutor *SolverNetExecutorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Fallback(&_SolverNetExecutor.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Fallback(&_SolverNetExecutor.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetExecutor *SolverNetExecutorSession) Receive() (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Receive(&_SolverNetExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) Receive() (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.Receive(&_SolverNetExecutor.TransactOpts)
}
