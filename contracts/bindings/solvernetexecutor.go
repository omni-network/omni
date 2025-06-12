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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_outbox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeAndTransfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeAndTransfer721\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"outbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tryRevokeApproval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FnSelectorNotRecognized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOutbox\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelf\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610bcd380380610bcd83398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f80fd5b81516001600160a01b0381168114610065575f80fd5b9392505050565b608051610b266100a75f395f8181610172015281816101ed0152818161024d015281816104ac0152818161058b01526105e80152610b265ff3fe60806040526004361061007e575f3560e01c8063beabacc81161004d578063beabacc814610142578063ce11e6ab14610161578063e1f21c67146101b0578063febe2c2c146101cf576100a5565b80637d2e90c2146100de5780639d32b569146100fd578063a5524cb51461011c578063b61d27f61461012f576100a5565b366100a5573033036100a357604051633204506f60e01b815260040160405180910390fd5b005b5f3560e01c63bc197c81811463f23a6e6182141763150b7a02821417156100d057806020526020603cf35b50633c10b94e5f526004601cfd5b3480156100e9575f80fd5b506100a36100f83660046108a3565b6101e2565b348015610108575f80fd5b506100a36101173660046108cb565b610242565b6100a361012a366004610941565b610371565b6100a361013d3660046109bb565b6104a1565b34801561014d575f80fd5b506100a361015c366004610a11565b610580565b34801561016c575f80fd5b506101947f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b3480156101bb575f80fd5b506100a36101ca366004610a11565b6105dd565b6100a36101dd366004610a4a565b61063a565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461022b5760405163bda8fc9560e01b815260040160405180910390fd5b61023e6001600160a01b03831682610710565b5050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461028b5760405163bda8fc9560e01b815260040160405180910390fd5b60405163095ea7b360e01b81526001600160a01b0382811660048301525f602483015283169063095ea7b3906044016020604051808303815f875af19250505080156102f4575060408051601f3d908101601f191682019092526102f191810190610abb565b60015b61036c5760405163095ea7b360e01b81526001600160a01b0382811660048301526001602483015283169063095ea7b3906044016020604051808303815f875af1925050508015610362575060408051601f3d908101601f1916820190925261035f91810190610abb565b60015b1561023e57505050565b505050565b333014610391576040516314e1dbf760e11b815260040160405180910390fd5b5f836001600160a01b03163484846040516103ad929190610ae1565b5f6040518083038185875af1925050503d805f81146103e7576040519150601f19603f3d011682016040523d82523d5f602084013e6103ec565b606091505b505090508061040e57604051633204506f60e01b815260040160405180910390fd5b6001600160a01b0387166104355760405163c1ab6dc160e01b815260040160405180910390fd5b6040516323b872dd60e01b81523060048201526001600160a01b038681166024830152604482018890528816906323b872dd906064015f604051808303815f87803b158015610482575f80fd5b505af1158015610494573d5f803e3d5ffd5b5050505050505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104ea5760405163bda8fc9560e01b815260040160405180910390fd5b6001600160a01b0384166104fc573093505b5f846001600160a01b0316848484604051610518929190610ae1565b5f6040518083038185875af1925050503d805f8114610552576040519150601f19603f3d011682016040523d82523d5f602084013e610557565b606091505b505090508061057957604051633204506f60e01b815260040160405180910390fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105c95760405163bda8fc9560e01b815260040160405180910390fd5b61036c6001600160a01b0384168383610729565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106265760405163bda8fc9560e01b815260040160405180910390fd5b61036c6001600160a01b0384168383610773565b33301461065a576040516314e1dbf760e11b815260040160405180910390fd5b5f836001600160a01b0316348484604051610676929190610ae1565b5f6040518083038185875af1925050503d805f81146106b0576040519150601f19603f3d011682016040523d82523d5f602084013e6106b5565b606091505b50509050806106d757604051633204506f60e01b815260040160405180910390fd5b6001600160a01b0386166106f3576106ee856107f3565b610708565b6107066001600160a01b0387168661080f565b505b505050505050565b5f385f3884865af161023e5763b12d13eb5f526004601cfd5b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f51141661076957803d853b151710610769576390b8ec185f526004601cfd5b505f603452505050565b816014528060345263095ea7b360601b5f5260205f604460105f875af18060015f51141661076957803d853b151710610769575f60345263095ea7b360601b5f525f38604460105f885af1508160345260205f604460105f885af190508060015f51141661076957803d853b15171061076957633e3f8f735f526004601cfd5b5f385f3847855af161080c5763b12d13eb5f526004601cfd5b50565b5f6370a082315f5230602052602060346024601c865afa601f3d111661083c576390b8ec185f526004601cfd5b81601452603451905063a9059cbb60601b5f5260205f604460105f875af18060015f51141661087d57803d853b15171061087d576390b8ec185f526004601cfd5b505f60345292915050565b80356001600160a01b038116811461089e575f80fd5b919050565b5f80604083850312156108b4575f80fd5b6108bd83610888565b946020939093013593505050565b5f80604083850312156108dc575f80fd5b6108e583610888565b91506108f360208401610888565b90509250929050565b5f8083601f84011261090c575f80fd5b50813567ffffffffffffffff811115610923575f80fd5b60208301915083602082850101111561093a575f80fd5b9250929050565b5f805f805f8060a08789031215610956575f80fd5b61095f87610888565b95506020870135945061097460408801610888565b935061098260608801610888565b9250608087013567ffffffffffffffff81111561099d575f80fd5b6109a989828a016108fc565b979a9699509497509295939492505050565b5f805f80606085870312156109ce575f80fd5b6109d785610888565b935060208501359250604085013567ffffffffffffffff8111156109f9575f80fd5b610a05878288016108fc565b95989497509550505050565b5f805f60608486031215610a23575f80fd5b610a2c84610888565b9250610a3a60208501610888565b9150604084013590509250925092565b5f805f805f60808688031215610a5e575f80fd5b610a6786610888565b9450610a7560208701610888565b9350610a8360408701610888565b9250606086013567ffffffffffffffff811115610a9e575f80fd5b610aaa888289016108fc565b969995985093965092949392505050565b5f60208284031215610acb575f80fd5b81518015158114610ada575f80fd5b9392505050565b818382375f910190815291905056fea2646970667358221220999c7b4f53981506a709664b8bf9ab9dd3b6b073da3bd27a9508d6328f6a662964736f6c63430008180033",
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

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) ExecuteAndTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.Transact(opts, "executeAndTransfer", token, to, target, data)
}

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorSession) ExecuteAndTransfer(token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.ExecuteAndTransfer(&_SolverNetExecutor.TransactOpts, token, to, target, data)
}

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) ExecuteAndTransfer(token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.ExecuteAndTransfer(&_SolverNetExecutor.TransactOpts, token, to, target, data)
}

// ExecuteAndTransfer721 is a paid mutator transaction binding the contract method 0xa5524cb5.
//
// Solidity: function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactor) ExecuteAndTransfer721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.contract.Transact(opts, "executeAndTransfer721", token, tokenId, to, target, data)
}

// ExecuteAndTransfer721 is a paid mutator transaction binding the contract method 0xa5524cb5.
//
// Solidity: function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorSession) ExecuteAndTransfer721(token common.Address, tokenId *big.Int, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.ExecuteAndTransfer721(&_SolverNetExecutor.TransactOpts, token, tokenId, to, target, data)
}

// ExecuteAndTransfer721 is a paid mutator transaction binding the contract method 0xa5524cb5.
//
// Solidity: function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes data) payable returns()
func (_SolverNetExecutor *SolverNetExecutorTransactorSession) ExecuteAndTransfer721(token common.Address, tokenId *big.Int, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetExecutor.Contract.ExecuteAndTransfer721(&_SolverNetExecutor.TransactOpts, token, tokenId, to, target, data)
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
