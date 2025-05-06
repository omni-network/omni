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

// SolverNetMiddlemanMetaData contains all meta data concerning the SolverNetMiddleman contract.
var SolverNetMiddlemanMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeAndTransfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeAndTransfer721\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b506104af8061001d5f395ff3fe608060405260043610610029575f3560e01c8063a5524cb514610032578063febe2c2c1461004557005b3661003057005b005b61003061004036600461037f565b610058565b6100306100533660046103f9565b61019e565b3068929eee149b4bd2126854036100765763ab143c065f526004601cfd5b3068929eee149b4bd21268555f836001600160a01b031634848460405161009e92919061046a565b5f6040518083038185875af1925050503d805f81146100d8576040519150601f19603f3d011682016040523d82523d5f602084013e6100dd565b606091505b50509050806100ff57604051633204506f60e01b815260040160405180910390fd5b6001600160a01b0387166101265760405163c1ab6dc160e01b815260040160405180910390fd5b6040516323b872dd60e01b81523060048201526001600160a01b038681166024830152604482018890528816906323b872dd906064015f604051808303815f87803b158015610173575f80fd5b505af1158015610185573d5f803e3d5ffd5b50505050503868929eee149b4bd2126855505050505050565b3068929eee149b4bd2126854036101bc5763ab143c065f526004601cfd5b3068929eee149b4bd21268555f836001600160a01b03163484846040516101e492919061046a565b5f6040518083038185875af1925050503d805f811461021e576040519150601f19603f3d011682016040523d82523d5f602084013e610223565b606091505b505090508061024557604051633204506f60e01b815260040160405180910390fd5b6001600160a01b0386166102615761025c8561028a565b610276565b6102746001600160a01b038716866102a6565b505b503868929eee149b4bd21268555050505050565b5f385f3847855af16102a35763b12d13eb5f526004601cfd5b50565b5f6370a082315f5230602052602060346024601c865afa601f3d11166102d3576390b8ec185f526004601cfd5b81601452603451905063a9059cbb60601b5f5260205f604460105f875af18060015f51141661031457803d853b151710610314576390b8ec185f526004601cfd5b505f60345292915050565b80356001600160a01b0381168114610335575f80fd5b919050565b5f8083601f84011261034a575f80fd5b50813567ffffffffffffffff811115610361575f80fd5b602083019150836020828501011115610378575f80fd5b9250929050565b5f805f805f8060a08789031215610394575f80fd5b61039d8761031f565b9550602087013594506103b26040880161031f565b93506103c06060880161031f565b9250608087013567ffffffffffffffff8111156103db575f80fd5b6103e789828a0161033a565b979a9699509497509295939492505050565b5f805f805f6080868803121561040d575f80fd5b6104168661031f565b94506104246020870161031f565b93506104326040870161031f565b9250606086013567ffffffffffffffff81111561044d575f80fd5b6104598882890161033a565b969995985093965092949392505050565b818382375f910190815291905056fea2646970667358221220dae8d0a3bef1232b3ccb539348490b2de8f0e15004e1895f92312a2ca62e344a64736f6c63430008180033",
}

// SolverNetMiddlemanABI is the input ABI used to generate the binding from.
// Deprecated: Use SolverNetMiddlemanMetaData.ABI instead.
var SolverNetMiddlemanABI = SolverNetMiddlemanMetaData.ABI

// SolverNetMiddlemanBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolverNetMiddlemanMetaData.Bin instead.
var SolverNetMiddlemanBin = SolverNetMiddlemanMetaData.Bin

// DeploySolverNetMiddleman deploys a new Ethereum contract, binding an instance of SolverNetMiddleman to it.
func DeploySolverNetMiddleman(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolverNetMiddleman, error) {
	parsed, err := SolverNetMiddlemanMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolverNetMiddlemanBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolverNetMiddleman{SolverNetMiddlemanCaller: SolverNetMiddlemanCaller{contract: contract}, SolverNetMiddlemanTransactor: SolverNetMiddlemanTransactor{contract: contract}, SolverNetMiddlemanFilterer: SolverNetMiddlemanFilterer{contract: contract}}, nil
}

// SolverNetMiddleman is an auto generated Go binding around an Ethereum contract.
type SolverNetMiddleman struct {
	SolverNetMiddlemanCaller     // Read-only binding to the contract
	SolverNetMiddlemanTransactor // Write-only binding to the contract
	SolverNetMiddlemanFilterer   // Log filterer for contract events
}

// SolverNetMiddlemanCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolverNetMiddlemanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetMiddlemanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolverNetMiddlemanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetMiddlemanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolverNetMiddlemanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetMiddlemanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolverNetMiddlemanSession struct {
	Contract     *SolverNetMiddleman // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SolverNetMiddlemanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolverNetMiddlemanCallerSession struct {
	Contract *SolverNetMiddlemanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SolverNetMiddlemanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolverNetMiddlemanTransactorSession struct {
	Contract     *SolverNetMiddlemanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SolverNetMiddlemanRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolverNetMiddlemanRaw struct {
	Contract *SolverNetMiddleman // Generic contract binding to access the raw methods on
}

// SolverNetMiddlemanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolverNetMiddlemanCallerRaw struct {
	Contract *SolverNetMiddlemanCaller // Generic read-only contract binding to access the raw methods on
}

// SolverNetMiddlemanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolverNetMiddlemanTransactorRaw struct {
	Contract *SolverNetMiddlemanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolverNetMiddleman creates a new instance of SolverNetMiddleman, bound to a specific deployed contract.
func NewSolverNetMiddleman(address common.Address, backend bind.ContractBackend) (*SolverNetMiddleman, error) {
	contract, err := bindSolverNetMiddleman(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddleman{SolverNetMiddlemanCaller: SolverNetMiddlemanCaller{contract: contract}, SolverNetMiddlemanTransactor: SolverNetMiddlemanTransactor{contract: contract}, SolverNetMiddlemanFilterer: SolverNetMiddlemanFilterer{contract: contract}}, nil
}

// NewSolverNetMiddlemanCaller creates a new read-only instance of SolverNetMiddleman, bound to a specific deployed contract.
func NewSolverNetMiddlemanCaller(address common.Address, caller bind.ContractCaller) (*SolverNetMiddlemanCaller, error) {
	contract, err := bindSolverNetMiddleman(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddlemanCaller{contract: contract}, nil
}

// NewSolverNetMiddlemanTransactor creates a new write-only instance of SolverNetMiddleman, bound to a specific deployed contract.
func NewSolverNetMiddlemanTransactor(address common.Address, transactor bind.ContractTransactor) (*SolverNetMiddlemanTransactor, error) {
	contract, err := bindSolverNetMiddleman(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddlemanTransactor{contract: contract}, nil
}

// NewSolverNetMiddlemanFilterer creates a new log filterer instance of SolverNetMiddleman, bound to a specific deployed contract.
func NewSolverNetMiddlemanFilterer(address common.Address, filterer bind.ContractFilterer) (*SolverNetMiddlemanFilterer, error) {
	contract, err := bindSolverNetMiddleman(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolverNetMiddlemanFilterer{contract: contract}, nil
}

// bindSolverNetMiddleman binds a generic wrapper to an already deployed contract.
func bindSolverNetMiddleman(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolverNetMiddlemanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetMiddleman *SolverNetMiddlemanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetMiddleman.Contract.SolverNetMiddlemanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetMiddleman *SolverNetMiddlemanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.SolverNetMiddlemanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetMiddleman *SolverNetMiddlemanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.SolverNetMiddlemanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetMiddleman *SolverNetMiddlemanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetMiddleman.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.contract.Transact(opts, method, params...)
}

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactor) ExecuteAndTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.contract.Transact(opts, "executeAndTransfer", token, to, target, data)
}

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanSession) ExecuteAndTransfer(token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.ExecuteAndTransfer(&_SolverNetMiddleman.TransactOpts, token, to, target, data)
}

// ExecuteAndTransfer is a paid mutator transaction binding the contract method 0xfebe2c2c.
//
// Solidity: function executeAndTransfer(address token, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorSession) ExecuteAndTransfer(token common.Address, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.ExecuteAndTransfer(&_SolverNetMiddleman.TransactOpts, token, to, target, data)
}

// ExecuteAndTransfer721 is a paid mutator transaction binding the contract method 0xa5524cb5.
//
// Solidity: function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactor) ExecuteAndTransfer721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.contract.Transact(opts, "executeAndTransfer721", token, tokenId, to, target, data)
}

// ExecuteAndTransfer721 is a paid mutator transaction binding the contract method 0xa5524cb5.
//
// Solidity: function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanSession) ExecuteAndTransfer721(token common.Address, tokenId *big.Int, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.ExecuteAndTransfer721(&_SolverNetMiddleman.TransactOpts, token, tokenId, to, target, data)
}

// ExecuteAndTransfer721 is a paid mutator transaction binding the contract method 0xa5524cb5.
//
// Solidity: function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes data) payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorSession) ExecuteAndTransfer721(token common.Address, tokenId *big.Int, to common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.ExecuteAndTransfer721(&_SolverNetMiddleman.TransactOpts, token, tokenId, to, target, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.Fallback(&_SolverNetMiddleman.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.Fallback(&_SolverNetMiddleman.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetMiddleman.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanSession) Receive() (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.Receive(&_SolverNetMiddleman.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SolverNetMiddleman *SolverNetMiddlemanTransactorSession) Receive() (*types.Transaction, error) {
	return _SolverNetMiddleman.Contract.Receive(&_SolverNetMiddleman.TransactOpts)
}
