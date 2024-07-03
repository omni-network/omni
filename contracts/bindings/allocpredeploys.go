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

// AllocPredeploysConfig is an auto generated low-level Go binding around an user-defined struct.
type AllocPredeploysConfig struct {
	Admin                  common.Address
	ChainId                *big.Int
	EnableStakingAllowlist bool
	Output                 string
}

// AllocPredeploysMetaData contains all meta data concerning the AllocPredeploys contract.
var AllocPredeploysMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_SCRIPT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"runWithCfg\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structAllocPredeploys.Config\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enableStakingAllowlist\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052600c805462ff00ff19166201000117905534801561002157600080fd5b50611845806100316000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630a9254e4146100465780633858e62314610050578063f8ccbf4714610063575b600080fd5b61004e61008a565b005b61004e61005e366004611316565b6100d5565b600c546100769062010000900460ff1681565b604051901515815260200160405180910390f35b6100b3604051806040016040528060088152602001673232b83637bcb2b960c11b8152506103e3565b601180546001600160a01b0319166001600160a01b0392909216919091179055565b80600d6100e28282611516565b50506011546040516303223eab60e11b81526001600160a01b0390911660048201526000805160206117f0833981519152906306447d5690602401600060405180830381600087803b15801561013757600080fd5b505af115801561014b573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156101ad57600080fd5b505af11580156101c1573d6000803e3d6000fd5b505050506101cd6103f5565b604051631c72346d60e01b81523360048201526000805160206117f083398151915290631c72346d90602401600060405180830381600087803b15801561021357600080fd5b505af1158015610227573d6000803e3d6000fd5b505060405163c88a5e6d60e01b8152336004820152600060248201526000805160206117f0833981519152925063c88a5e6d9150604401600060405180830381600087803b15801561027857600080fd5b505af115801561028c573d6000803e3d6000fd5b5050601154604051631c72346d60e01b81526001600160a01b0390911660048201526000805160206117f08339815191529250631c72346d9150602401600060405180830381600087803b1580156102e357600080fd5b505af11580156102f7573d6000803e3d6000fd5b505060115460405163c88a5e6d60e01b81526001600160a01b039091166004820152600060248201526000805160206117f0833981519152925063c88a5e6d9150604401600060405180830381600087803b15801561035557600080fd5b505af1158015610369573d6000803e3d6000fd5b506000805160206117f0833981519152925063709ecd3f915061039190506060840184611370565b6040518363ffffffff1660e01b81526004016103ae9291906115c4565b600060405180830381600087803b1580156103c857600080fd5b505af11580156103dc573d6000803e3d6000fd5b5050505050565b60006103ee8261042f565b5092915050565b6103fd610539565b61040561057f565b61040d6106fa565b61041561089f565b61041d610af4565b610425610be9565b61042d610d87565b565b600080826040516020016104439190611617565b60408051808303601f190181529082905280516020909101206001625e79b760e01b031982526004820181905291506000805160206117f08339815191529063ffa1864990602401602060405180830381865afa1580156104a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cc9190611633565b6040516318caf8e360e31b81529092506000805160206117f08339815191529063c657c71890610502908590879060040161167c565b600060405180830381600087803b15801561051c57600080fd5b505af1158015610530573d6000803e3d6000fd5b50505050915091565b6000610543610dff565b905060005b815181101561057b57610573828281518110610566576105666116a8565b6020026020010151610e96565b600101610548565b5050565b600d546040516370ca10bb60e01b815273aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa60048201526000602482018190526001600160a01b0390921660448201526000805160206117f0833981519152906370ca10bb90606401600060405180830381600087803b1580156105f557600080fd5b505af1158015610609573d6000803e3d6000fd5b5050604051630fafdced60e21b815260206004820152602960248201527f6f75742f50726f787941646d696e2e736f6c2f50726f787941646d696e2e302e6044820152681c17191a173539b7b760b91b60648201526000805160206117f0833981519152925063b4d6c782915073aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa908390633ebf73b4906084015b600060405180830381865afa1580156106b5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106dd91908101906116be565b6040518363ffffffff1660e01b81526004016103ae92919061167c565b600061070e600162048789608a1b016110b2565b604051630fafdced60e21b815260206004820152602160248201527f506f7274616c52656769737472792e736f6c3a506f7274616c526567697374726044820152607960f81b60648201529091506000805160206117f08339815191529063b4d6c7829083908390633ebf73b490608401600060405180830381865afa15801561079c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107c491908101906116be565b6040518363ffffffff1660e01b81526004016107e192919061167c565b600060405180830381600087803b1580156107fb57600080fd5b505af115801561080f573d6000803e3d6000fd5b50505050806001600160a01b03166331f449006040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561084e57600080fd5b505af1158015610862573d6000803e3d6000fd5b5050600d5460405163189acdbd60e31b81526001600160a01b039091166004820152600162048789608a1b01925063c4d66de891506024016103ae565b60405163c88a5e6d60e01b8152600262048789608a1b01600482015269d3c21bcecceda100000060248201819052906000805160206117f08339815191529063c88a5e6d90604401600060405180830381600087803b15801561090157600080fd5b505af1158015610915573d6000803e3d6000fd5b50505050600061092d600262048789608a1b016110b2565b604051630fafdced60e21b815260206004820152602560248201527f4f6d6e694272696467654e61746976652e736f6c3a4f6d6e694272696467654e604482015264617469766560d81b60648201529091506000805160206117f08339815191529063b4d6c7829083908390633ebf73b490608401600060405180830381865afa1580156109bf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109e791908101906116be565b6040518363ffffffff1660e01b8152600401610a0492919061167c565b600060405180830381600087803b158015610a1e57600080fd5b505af1158015610a32573d6000803e3d6000fd5b50505050806001600160a01b03166331f449006040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610a7157600080fd5b505af1158015610a85573d6000803e3d6000fd5b5050600d5460405163189acdbd60e31b81526001600160a01b039091166004820152600262048789608a1b01925063c4d66de891506024015b600060405180830381600087803b158015610ad857600080fd5b505af1158015610aec573d6000803e3d6000fd5b505050505050565b604051630fafdced60e21b815260206004820152600f60248201526e574f6d6e692e736f6c3a574f6d6e6960881b60448201526000805160206117f08339815191529063b4d6c78290600362048789608a1b01908390633ebf73b490606401600060405180830381865afa158015610b70573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b9891908101906116be565b6040518363ffffffff1660e01b8152600401610bb592919061167c565b600060405180830381600087803b158015610bcf57600080fd5b505af1158015610be3573d6000803e3d6000fd5b50505050565b6000610bfd600162333333608a1b016110b2565b604051630fafdced60e21b81526020600482015260136024820152725374616b696e672e736f6c3a5374616b696e6760681b60448201529091506000805160206117f08339815191529063b4d6c7829083908390633ebf73b490606401600060405180830381865afa158015610c77573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c9f91908101906116be565b6040518363ffffffff1660e01b8152600401610cbc92919061167c565b600060405180830381600087803b158015610cd657600080fd5b505af1158015610cea573d6000803e3d6000fd5b50505050806001600160a01b03166331f449006040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610d2957600080fd5b505af1158015610d3d573d6000803e3d6000fd5b5050600d54600f5460405163400ada7560e01b81526001600160a01b03909216600483015260ff1615156024820152600162333333608a1b01925063400ada7591506044016103ae565b6000610d9b600262333333608a1b016110b2565b604051630fafdced60e21b8152602060048201526015602482015274536c617368696e672e736f6c3a536c617368696e6760581b60448201529091506000805160206117f08339815191529063b4d6c7829083908390633ebf73b490606401610698565b604080516002808252606080830184529260208301908036833701905050905062048789608a1b81600081518110610e3957610e396116a8565b60200260200101906001600160a01b031690816001600160a01b03168152505062333333608a1b81600181518110610e7357610e736116a8565b60200260200101906001600160a01b031690816001600160a01b03168152505090565b63ffffffff811615610ee35760405162461bcd60e51b8152602060048201526011602482015270696e76616c6964206e616d65737061636560781b60448201526064015b60405180910390fd5b604051630fafdced60e21b815260206004820152603b60248201527f5472616e73706172656e745570677261646561626c6550726f78792e736f6c3a60448201527f5472616e73706172656e745570677261646561626c6550726f7879000000000060648201526000906000805160206117f083398151915290633ebf73b490608401600060405180830381865afa158015610f83573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610fab91908101906116be565b905060015b610400816001600160a01b0316116110ad576000610fce8285611781565b9050600362048789608a1b016001600160a01b03821603610fef575061109b565b604051635a6b63c160e11b81526000805160206117f08339815191529063b4d6c78290611022908490879060040161167c565b600060405180830381600087803b15801561103c57600080fd5b505af1158015611050573d6000803e3d6000fd5b505050506110728173aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa611181565b61107b816111e9565b1561109957600061108b826110b2565b90506110978282611275565b505b505b806110a5816117a1565b915050610fb0565b505050565b60006110bd826112dd565b6111095760405162461bcd60e51b815260206004820152601b60248201527f5072656465706c6f79733a206e6f742061207072656465706c6f7900000000006044820152606401610eda565b600362048789608a1b016001600160a01b0383160361116a5760405162461bcd60e51b815260206004820152601760248201527f5072656465706c6f79733a206e6f742070726f786965640000000000000000006044820152606401610eda565b61117b826001600160a01b036117cf565b92915050565b6040516370ca10bb60e01b81526001600160a01b0380841660048301527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024830152821660448201526000805160206117f0833981519152906370ca10bb90606401610abe565b60006001600160a01b038216600162048789608a1b01148061121c57506001600160a01b038216600262048789608a1b01145b8061123857506001600160a01b038216600362048789608a1b01145b8061125457506001600160a01b038216600162333333608a1b01145b8061117b57506001600160a01b038216600262333333608a1b011492915050565b6040516370ca10bb60e01b81526001600160a01b0380841660048301527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6024830152821660448201526000805160206117f0833981519152906370ca10bb90606401610abe565b60006001600160951b03600b83901c1662048789607f1b148061117b57505062333333607f1b600b9190911c6001600160951b03161490565b60006020828403121561132857600080fd5b813567ffffffffffffffff81111561133f57600080fd5b82016080818503121561135157600080fd5b9392505050565b6001600160a01b038116811461136d57600080fd5b50565b6000808335601e1984360301811261138757600080fd5b83018035915067ffffffffffffffff8211156113a257600080fd5b6020019150368190038213156113b757600080fd5b9250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806113e857607f821691505b60208210810361140857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156110ad576000816000526020600020601f850160051c810160208610156114375750805b601f850160051c820191505b81811015610aec57828155600101611443565b67ffffffffffffffff83111561146e5761146e6113be565b6114828361147c83546113d4565b8361140e565b6000601f8411600181146114b6576000851561149e5750838201355b600019600387901b1c1916600186901b1783556103dc565b600083815260209020601f19861690835b828110156114e757868501358255602094850194600190920191016114c7565b50868210156115045760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b813561152181611358565b81546001600160a01b0319166001600160a01b03919091161781556020820135600182015560028101604083013580151580821461155e57600080fd5b60ff19835416915060ff8116821783555050506060820135601e1983360301811261158857600080fd5b8201803567ffffffffffffffff8111156115a157600080fd5b6020820191508036038213156115b657600080fd5b610be3818360038601611456565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60005b8381101561160e5781810151838201526020016115f6565b50506000910152565b600082516116298184602087016115f3565b9190910192915050565b60006020828403121561164557600080fd5b815161135181611358565b600081518084526116688160208601602086016115f3565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190526000906116a090830184611650565b949350505050565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156116d057600080fd5b815167ffffffffffffffff808211156116e857600080fd5b818401915084601f8301126116fc57600080fd5b81518181111561170e5761170e6113be565b604051601f8201601f19908116603f01168101908382118183101715611736576117366113be565b8160405282815287602084870101111561174f57600080fd5b6117608360208301602088016115f3565b979650505050505050565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b038181168382160190808211156103ee576103ee61176b565b60006001600160a01b038281166002600160a01b031981016117c5576117c561176b565b6001019392505050565b6001600160a01b038281168282160390808211156103ee576103ee61176b56fe0000000000000000000000007109709ecfa91a80626ff3989d68f67f5b1dd12da26469706673582212207474ab2dccb3c2641a6ff145b373a772f2a6ad98728f2a2a1985d6a57c72c9e864736f6c63430008180033",
}

// AllocPredeploysABI is the input ABI used to generate the binding from.
// Deprecated: Use AllocPredeploysMetaData.ABI instead.
var AllocPredeploysABI = AllocPredeploysMetaData.ABI

// AllocPredeploysBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AllocPredeploysMetaData.Bin instead.
var AllocPredeploysBin = AllocPredeploysMetaData.Bin

// DeployAllocPredeploys deploys a new Ethereum contract, binding an instance of AllocPredeploys to it.
func DeployAllocPredeploys(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AllocPredeploys, error) {
	parsed, err := AllocPredeploysMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AllocPredeploysBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AllocPredeploys{AllocPredeploysCaller: AllocPredeploysCaller{contract: contract}, AllocPredeploysTransactor: AllocPredeploysTransactor{contract: contract}, AllocPredeploysFilterer: AllocPredeploysFilterer{contract: contract}}, nil
}

// AllocPredeploys is an auto generated Go binding around an Ethereum contract.
type AllocPredeploys struct {
	AllocPredeploysCaller     // Read-only binding to the contract
	AllocPredeploysTransactor // Write-only binding to the contract
	AllocPredeploysFilterer   // Log filterer for contract events
}

// AllocPredeploysCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllocPredeploysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocPredeploysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllocPredeploysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocPredeploysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllocPredeploysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocPredeploysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllocPredeploysSession struct {
	Contract     *AllocPredeploys  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AllocPredeploysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllocPredeploysCallerSession struct {
	Contract *AllocPredeploysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AllocPredeploysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllocPredeploysTransactorSession struct {
	Contract     *AllocPredeploysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AllocPredeploysRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllocPredeploysRaw struct {
	Contract *AllocPredeploys // Generic contract binding to access the raw methods on
}

// AllocPredeploysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllocPredeploysCallerRaw struct {
	Contract *AllocPredeploysCaller // Generic read-only contract binding to access the raw methods on
}

// AllocPredeploysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllocPredeploysTransactorRaw struct {
	Contract *AllocPredeploysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllocPredeploys creates a new instance of AllocPredeploys, bound to a specific deployed contract.
func NewAllocPredeploys(address common.Address, backend bind.ContractBackend) (*AllocPredeploys, error) {
	contract, err := bindAllocPredeploys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllocPredeploys{AllocPredeploysCaller: AllocPredeploysCaller{contract: contract}, AllocPredeploysTransactor: AllocPredeploysTransactor{contract: contract}, AllocPredeploysFilterer: AllocPredeploysFilterer{contract: contract}}, nil
}

// NewAllocPredeploysCaller creates a new read-only instance of AllocPredeploys, bound to a specific deployed contract.
func NewAllocPredeploysCaller(address common.Address, caller bind.ContractCaller) (*AllocPredeploysCaller, error) {
	contract, err := bindAllocPredeploys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllocPredeploysCaller{contract: contract}, nil
}

// NewAllocPredeploysTransactor creates a new write-only instance of AllocPredeploys, bound to a specific deployed contract.
func NewAllocPredeploysTransactor(address common.Address, transactor bind.ContractTransactor) (*AllocPredeploysTransactor, error) {
	contract, err := bindAllocPredeploys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllocPredeploysTransactor{contract: contract}, nil
}

// NewAllocPredeploysFilterer creates a new log filterer instance of AllocPredeploys, bound to a specific deployed contract.
func NewAllocPredeploysFilterer(address common.Address, filterer bind.ContractFilterer) (*AllocPredeploysFilterer, error) {
	contract, err := bindAllocPredeploys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllocPredeploysFilterer{contract: contract}, nil
}

// bindAllocPredeploys binds a generic wrapper to an already deployed contract.
func bindAllocPredeploys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AllocPredeploysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocPredeploys *AllocPredeploysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocPredeploys.Contract.AllocPredeploysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocPredeploys *AllocPredeploysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.AllocPredeploysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocPredeploys *AllocPredeploysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.AllocPredeploysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocPredeploys *AllocPredeploysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocPredeploys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocPredeploys *AllocPredeploysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocPredeploys *AllocPredeploysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.contract.Transact(opts, method, params...)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_AllocPredeploys *AllocPredeploysCaller) ISSCRIPT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AllocPredeploys.contract.Call(opts, &out, "IS_SCRIPT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_AllocPredeploys *AllocPredeploysSession) ISSCRIPT() (bool, error) {
	return _AllocPredeploys.Contract.ISSCRIPT(&_AllocPredeploys.CallOpts)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_AllocPredeploys *AllocPredeploysCallerSession) ISSCRIPT() (bool, error) {
	return _AllocPredeploys.Contract.ISSCRIPT(&_AllocPredeploys.CallOpts)
}

// RunWithCfg is a paid mutator transaction binding the contract method 0x3858e623.
//
// Solidity: function runWithCfg((address,uint256,bool,string) config) returns()
func (_AllocPredeploys *AllocPredeploysTransactor) RunWithCfg(opts *bind.TransactOpts, config AllocPredeploysConfig) (*types.Transaction, error) {
	return _AllocPredeploys.contract.Transact(opts, "runWithCfg", config)
}

// RunWithCfg is a paid mutator transaction binding the contract method 0x3858e623.
//
// Solidity: function runWithCfg((address,uint256,bool,string) config) returns()
func (_AllocPredeploys *AllocPredeploysSession) RunWithCfg(config AllocPredeploysConfig) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.RunWithCfg(&_AllocPredeploys.TransactOpts, config)
}

// RunWithCfg is a paid mutator transaction binding the contract method 0x3858e623.
//
// Solidity: function runWithCfg((address,uint256,bool,string) config) returns()
func (_AllocPredeploys *AllocPredeploysTransactorSession) RunWithCfg(config AllocPredeploysConfig) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.RunWithCfg(&_AllocPredeploys.TransactOpts, config)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_AllocPredeploys *AllocPredeploysTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocPredeploys.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_AllocPredeploys *AllocPredeploysSession) SetUp() (*types.Transaction, error) {
	return _AllocPredeploys.Contract.SetUp(&_AllocPredeploys.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_AllocPredeploys *AllocPredeploysTransactorSession) SetUp() (*types.Transaction, error) {
	return _AllocPredeploys.Contract.SetUp(&_AllocPredeploys.TransactOpts)
}
