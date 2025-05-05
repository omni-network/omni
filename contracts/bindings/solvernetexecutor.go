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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_outbox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeAndTransfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeAndTransfer721\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"outbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tryRevokeApproval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOutbox\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelf\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610aec380380610aec83398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f80fd5b81516001600160a01b0381168114610065575f80fd5b9392505050565b608051610a456100a75f395f818161011501528181610190015281816101f0015281816103dd015281816104aa01526105070152610a455ff3fe608060405260043610610078575f3560e01c8063beabacc81161004a578063beabacc8146100e5578063ce11e6ab14610104578063e1f21c6714610153578063febe2c2c1461017257005b80637d2e90c2146100815780639d32b569146100a0578063a5524cb5146100bf578063b61d27f6146100d257005b3661007f57005b005b34801561008c575f80fd5b5061007f61009b3660046107c2565b610185565b3480156100ab575f80fd5b5061007f6100ba3660046107ea565b6101e5565b61007f6100cd366004610860565b6102a2565b61007f6100e03660046108da565b6103d2565b3480156100f0575f80fd5b5061007f6100ff366004610930565b61049f565b34801561010f575f80fd5b506101377f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b34801561015e575f80fd5b5061007f61016d366004610930565b6104fc565b61007f610180366004610969565b610559565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101ce5760405163bda8fc9560e01b815260040160405180910390fd5b6101e16001600160a01b0383168261062f565b5050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461022e5760405163bda8fc9560e01b815260040160405180910390fd5b60405163095ea7b360e01b81526001600160a01b0382811660048301525f602483015283169063095ea7b3906044016020604051808303815f875af1925050508015610297575060408051601f3d908101601f19168201909252610294918101906109da565b60015b156101e1575b505050565b3330146102c2576040516314e1dbf760e11b815260040160405180910390fd5b5f836001600160a01b03163484846040516102de929190610a00565b5f6040518083038185875af1925050503d805f8114610318576040519150601f19603f3d011682016040523d82523d5f602084013e61031d565b606091505b505090508061033f57604051633204506f60e01b815260040160405180910390fd5b6001600160a01b0387166103665760405163c1ab6dc160e01b815260040160405180910390fd5b6040516323b872dd60e01b81523060048201526001600160a01b038681166024830152604482018890528816906323b872dd906064015f604051808303815f87803b1580156103b3575f80fd5b505af11580156103c5573d5f803e3d5ffd5b5050505050505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461041b5760405163bda8fc9560e01b815260040160405180910390fd5b5f846001600160a01b0316848484604051610437929190610a00565b5f6040518083038185875af1925050503d805f8114610471576040519150601f19603f3d011682016040523d82523d5f602084013e610476565b606091505b505090508061049857604051633204506f60e01b815260040160405180910390fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104e85760405163bda8fc9560e01b815260040160405180910390fd5b61029d6001600160a01b0384168383610648565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105455760405163bda8fc9560e01b815260040160405180910390fd5b61029d6001600160a01b0384168383610692565b333014610579576040516314e1dbf760e11b815260040160405180910390fd5b5f836001600160a01b0316348484604051610595929190610a00565b5f6040518083038185875af1925050503d805f81146105cf576040519150601f19603f3d011682016040523d82523d5f602084013e6105d4565b606091505b50509050806105f657604051633204506f60e01b815260040160405180910390fd5b6001600160a01b0386166106125761060d85610712565b610627565b6106256001600160a01b0387168661072e565b505b505050505050565b5f385f3884865af16101e15763b12d13eb5f526004601cfd5b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f51141661068857803d853b151710610688576390b8ec185f526004601cfd5b505f603452505050565b816014528060345263095ea7b360601b5f5260205f604460105f875af18060015f51141661068857803d853b151710610688575f60345263095ea7b360601b5f525f38604460105f885af1508160345260205f604460105f885af190508060015f51141661068857803d853b15171061068857633e3f8f735f526004601cfd5b5f385f3847855af161072b5763b12d13eb5f526004601cfd5b50565b5f6370a082315f5230602052602060346024601c865afa601f3d111661075b576390b8ec185f526004601cfd5b81601452603451905063a9059cbb60601b5f5260205f604460105f875af18060015f51141661079c57803d853b15171061079c576390b8ec185f526004601cfd5b505f60345292915050565b80356001600160a01b03811681146107bd575f80fd5b919050565b5f80604083850312156107d3575f80fd5b6107dc836107a7565b946020939093013593505050565b5f80604083850312156107fb575f80fd5b610804836107a7565b9150610812602084016107a7565b90509250929050565b5f8083601f84011261082b575f80fd5b50813567ffffffffffffffff811115610842575f80fd5b602083019150836020828501011115610859575f80fd5b9250929050565b5f805f805f8060a08789031215610875575f80fd5b61087e876107a7565b955060208701359450610893604088016107a7565b93506108a1606088016107a7565b9250608087013567ffffffffffffffff8111156108bc575f80fd5b6108c889828a0161081b565b979a9699509497509295939492505050565b5f805f80606085870312156108ed575f80fd5b6108f6856107a7565b935060208501359250604085013567ffffffffffffffff811115610918575f80fd5b6109248782880161081b565b95989497509550505050565b5f805f60608486031215610942575f80fd5b61094b846107a7565b9250610959602085016107a7565b9150604084013590509250925092565b5f805f805f6080868803121561097d575f80fd5b610986866107a7565b9450610994602087016107a7565b93506109a2604087016107a7565b9250606086013567ffffffffffffffff8111156109bd575f80fd5b6109c98882890161081b565b969995985093965092949392505050565b5f602082840312156109ea575f80fd5b815180151581146109f9575f80fd5b9392505050565b818382375f910190815291905056fea264697066735822122099ae9da10b33a88c9e5263b8b625c83087605a713af9d331fd29b967a36efe2364736f6c63430008180033",
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
