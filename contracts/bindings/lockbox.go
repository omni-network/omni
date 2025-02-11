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

// LockboxMetaData contains all meta data concerning the Lockbox contract.
var LockboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositTo\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauser_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wrapped_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTo\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wrapped\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b610c84806100d96000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80638456cb59116100a2578063d547741f11610071578063d547741f14610233578063e63ab1e914610246578063f8c8765e1461025b578063fc0c546a1461026e578063ffaad6a51461028157600080fd5b80638456cb59146101fd57806391d1485414610205578063a217fddf14610218578063b6b55f251461022057600080fd5b80632f2ff15d116100e95780632f2ff15d1461018c57806336568abe1461019f5780633f4ba83a146101b257806350e70d48146101ba5780635c975abb146101e557600080fd5b806301ffc9a71461011b578063205c287814610143578063248a9ca3146101585780632e1a7d4d14610179575b600080fd5b61012e610129366004610ade565b610294565b60405190151581526020015b60405180910390f35b610156610151366004610b2b565b6102cb565b005b61016b610166366004610b55565b6102e2565b60405190815260200161013a565b610156610187366004610b55565b610304565b61015661019a366004610b6e565b61031a565b6101566101ad366004610b6e565b61033c565b610156610374565b6001546101cd906001600160a01b031681565b6040516001600160a01b03909116815260200161013a565b600080516020610c2f8339815191525460ff1661012e565b610156610394565b61012e610213366004610b6e565b6103b4565b61016b600081565b61015661022e366004610b55565b6103ec565b610156610241366004610b6e565b6103ff565b61016b600080516020610bef83398151915281565b610156610269366004610b9a565b61041b565b6000546101cd906001600160a01b031681565b61015661028f366004610b2b565b610626565b60006001600160e01b03198216637965db0b60e01b14806102c557506301ffc9a760e01b6001600160e01b03198316145b92915050565b6102d3610639565b6102de33838361066c565b5050565b6000908152600080516020610c0f833981519152602052604090206001015490565b61030c610639565b61031733338361066c565b50565b610323826102e2565b61032c816106ea565b61033683836106f4565b50505050565b6001600160a01b03811633146103655760405163334bd91960e11b815260040160405180910390fd5b61036f8282610799565b505050565b600080516020610bef83398151915261038c816106ea565b610317610815565b600080516020610bef8339815191526103ac816106ea565b610317610875565b6000918252600080516020610c0f833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6103f4610639565b6103173333836108be565b610408826102e2565b610411816106ea565b6103368383610799565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156104615750825b905060008267ffffffffffffffff16600114801561047e5750303b155b90508115801561048c575080155b156104aa5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156104d457845460ff60401b1916600160401b1785555b6001600160a01b0389166104fb5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0388166105225760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0387166105495760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0386166105705760405163d92e233d60e01b815260040160405180910390fd5b610578610941565b610580610949565b61058b60008a6106f4565b506105a4600080516020610bef833981519152896106f4565b50600080546001600160a01b03808a166001600160a01b0319928316179092556001805492891692909116919091179055831561061b57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b61062e610639565b6102de3383836108be565b600080516020610c2f8339815191525460ff161561066a5760405163d93c066560e01b815260040160405180910390fd5b565b6001546040516379ef98bf60e11b81526001600160a01b038581166004830152602482018490529091169063f3df317e90604401600060405180830381600087803b1580156106ba57600080fd5b505af11580156106ce573d6000803e3d6000fd5b505060005461036f92506001600160a01b031690508383610959565b61031781336109a9565b6000600080516020610c0f83398151915261070f84846103b4565b61078f576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556107453390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506102c5565b60009150506102c5565b6000600080516020610c0f8339815191526107b484846103b4565b1561078f576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506102c5565b61081d6109e6565b600080516020610c2f833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b61087d610639565b600080516020610c2f833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610857565b6000546108d6906001600160a01b0316843084610a16565b6001546040516340c10f1960e01b81526001600160a01b03848116600483015260248201849052909116906340c10f1990604401600060405180830381600087803b15801561092457600080fd5b505af1158015610938573d6000803e3d6000fd5b50505050505050565b61066a610a74565b610951610a74565b61066a610abd565b816014528060345263a9059cbb60601b60005260206000604460106000875af1806001600051141661099e57803d853b15171061099e576390b8ec186000526004601cfd5b506000603452505050565b6109b382826103b4565b6102de5760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440160405180910390fd5b600080516020610c2f8339815191525460ff1661066a57604051638dfc202b60e01b815260040160405180910390fd5b60405181606052826040528360601b602c526323b872dd60601b600c52602060006064601c6000895af18060016000511416610a6557803d873b151710610a6557637939f4246000526004601cfd5b50600060605260405250505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661066a57604051631afcd79f60e31b815260040160405180910390fd5b610ac5610a74565b600080516020610c2f833981519152805460ff19169055565b600060208284031215610af057600080fd5b81356001600160e01b031981168114610b0857600080fd5b9392505050565b80356001600160a01b0381168114610b2657600080fd5b919050565b60008060408385031215610b3e57600080fd5b610b4783610b0f565b946020939093013593505050565b600060208284031215610b6757600080fd5b5035919050565b60008060408385031215610b8157600080fd5b82359150610b9160208401610b0f565b90509250929050565b60008060008060808587031215610bb057600080fd5b610bb985610b0f565b9350610bc760208601610b0f565b9250610bd560408601610b0f565b9150610be360608601610b0f565b90509295919450925056fe539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a264697066735822122003f0f824d70119f2c32f09f500a1cec53630d1e5efd0d9d87326b02aefec4d1e64736f6c634300081a0033",
}

// LockboxABI is the input ABI used to generate the binding from.
// Deprecated: Use LockboxMetaData.ABI instead.
var LockboxABI = LockboxMetaData.ABI

// LockboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LockboxMetaData.Bin instead.
var LockboxBin = LockboxMetaData.Bin

// DeployLockbox deploys a new Ethereum contract, binding an instance of Lockbox to it.
func DeployLockbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lockbox, error) {
	parsed, err := LockboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LockboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lockbox{LockboxCaller: LockboxCaller{contract: contract}, LockboxTransactor: LockboxTransactor{contract: contract}, LockboxFilterer: LockboxFilterer{contract: contract}}, nil
}

// Lockbox is an auto generated Go binding around an Ethereum contract.
type Lockbox struct {
	LockboxCaller     // Read-only binding to the contract
	LockboxTransactor // Write-only binding to the contract
	LockboxFilterer   // Log filterer for contract events
}

// LockboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockboxSession struct {
	Contract     *Lockbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockboxCallerSession struct {
	Contract *LockboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LockboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockboxTransactorSession struct {
	Contract     *LockboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LockboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockboxRaw struct {
	Contract *Lockbox // Generic contract binding to access the raw methods on
}

// LockboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockboxCallerRaw struct {
	Contract *LockboxCaller // Generic read-only contract binding to access the raw methods on
}

// LockboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockboxTransactorRaw struct {
	Contract *LockboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockbox creates a new instance of Lockbox, bound to a specific deployed contract.
func NewLockbox(address common.Address, backend bind.ContractBackend) (*Lockbox, error) {
	contract, err := bindLockbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lockbox{LockboxCaller: LockboxCaller{contract: contract}, LockboxTransactor: LockboxTransactor{contract: contract}, LockboxFilterer: LockboxFilterer{contract: contract}}, nil
}

// NewLockboxCaller creates a new read-only instance of Lockbox, bound to a specific deployed contract.
func NewLockboxCaller(address common.Address, caller bind.ContractCaller) (*LockboxCaller, error) {
	contract, err := bindLockbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockboxCaller{contract: contract}, nil
}

// NewLockboxTransactor creates a new write-only instance of Lockbox, bound to a specific deployed contract.
func NewLockboxTransactor(address common.Address, transactor bind.ContractTransactor) (*LockboxTransactor, error) {
	contract, err := bindLockbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockboxTransactor{contract: contract}, nil
}

// NewLockboxFilterer creates a new log filterer instance of Lockbox, bound to a specific deployed contract.
func NewLockboxFilterer(address common.Address, filterer bind.ContractFilterer) (*LockboxFilterer, error) {
	contract, err := bindLockbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockboxFilterer{contract: contract}, nil
}

// bindLockbox binds a generic wrapper to an already deployed contract.
func bindLockbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LockboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lockbox *LockboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lockbox.Contract.LockboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lockbox *LockboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockbox.Contract.LockboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lockbox *LockboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lockbox.Contract.LockboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lockbox *LockboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lockbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lockbox *LockboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lockbox *LockboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lockbox.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lockbox *LockboxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lockbox.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lockbox *LockboxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Lockbox.Contract.DEFAULTADMINROLE(&_Lockbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lockbox *LockboxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Lockbox.Contract.DEFAULTADMINROLE(&_Lockbox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Lockbox *LockboxCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lockbox.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Lockbox *LockboxSession) PAUSERROLE() ([32]byte, error) {
	return _Lockbox.Contract.PAUSERROLE(&_Lockbox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Lockbox *LockboxCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Lockbox.Contract.PAUSERROLE(&_Lockbox.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lockbox *LockboxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Lockbox.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lockbox *LockboxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Lockbox.Contract.GetRoleAdmin(&_Lockbox.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lockbox *LockboxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Lockbox.Contract.GetRoleAdmin(&_Lockbox.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lockbox *LockboxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Lockbox.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lockbox *LockboxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Lockbox.Contract.HasRole(&_Lockbox.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lockbox *LockboxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Lockbox.Contract.HasRole(&_Lockbox.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lockbox *LockboxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lockbox.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lockbox *LockboxSession) Paused() (bool, error) {
	return _Lockbox.Contract.Paused(&_Lockbox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lockbox *LockboxCallerSession) Paused() (bool, error) {
	return _Lockbox.Contract.Paused(&_Lockbox.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lockbox *LockboxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Lockbox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lockbox *LockboxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lockbox.Contract.SupportsInterface(&_Lockbox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lockbox *LockboxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lockbox.Contract.SupportsInterface(&_Lockbox.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Lockbox *LockboxCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockbox.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Lockbox *LockboxSession) Token() (common.Address, error) {
	return _Lockbox.Contract.Token(&_Lockbox.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Lockbox *LockboxCallerSession) Token() (common.Address, error) {
	return _Lockbox.Contract.Token(&_Lockbox.CallOpts)
}

// Wrapped is a free data retrieval call binding the contract method 0x50e70d48.
//
// Solidity: function wrapped() view returns(address)
func (_Lockbox *LockboxCaller) Wrapped(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lockbox.contract.Call(opts, &out, "wrapped")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wrapped is a free data retrieval call binding the contract method 0x50e70d48.
//
// Solidity: function wrapped() view returns(address)
func (_Lockbox *LockboxSession) Wrapped() (common.Address, error) {
	return _Lockbox.Contract.Wrapped(&_Lockbox.CallOpts)
}

// Wrapped is a free data retrieval call binding the contract method 0x50e70d48.
//
// Solidity: function wrapped() view returns(address)
func (_Lockbox *LockboxCallerSession) Wrapped() (common.Address, error) {
	return _Lockbox.Contract.Wrapped(&_Lockbox.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Lockbox *LockboxTransactor) Deposit(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "deposit", value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Lockbox *LockboxSession) Deposit(value *big.Int) (*types.Transaction, error) {
	return _Lockbox.Contract.Deposit(&_Lockbox.TransactOpts, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 value) returns()
func (_Lockbox *LockboxTransactorSession) Deposit(value *big.Int) (*types.Transaction, error) {
	return _Lockbox.Contract.Deposit(&_Lockbox.TransactOpts, value)
}

// DepositTo is a paid mutator transaction binding the contract method 0xffaad6a5.
//
// Solidity: function depositTo(address to, uint256 value) returns()
func (_Lockbox *LockboxTransactor) DepositTo(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "depositTo", to, value)
}

// DepositTo is a paid mutator transaction binding the contract method 0xffaad6a5.
//
// Solidity: function depositTo(address to, uint256 value) returns()
func (_Lockbox *LockboxSession) DepositTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lockbox.Contract.DepositTo(&_Lockbox.TransactOpts, to, value)
}

// DepositTo is a paid mutator transaction binding the contract method 0xffaad6a5.
//
// Solidity: function depositTo(address to, uint256 value) returns()
func (_Lockbox *LockboxTransactorSession) DepositTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lockbox.Contract.DepositTo(&_Lockbox.TransactOpts, to, value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lockbox *LockboxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lockbox *LockboxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lockbox.Contract.GrantRole(&_Lockbox.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lockbox *LockboxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lockbox.Contract.GrantRole(&_Lockbox.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address pauser_, address token_, address wrapped_) returns()
func (_Lockbox *LockboxTransactor) Initialize(opts *bind.TransactOpts, admin_ common.Address, pauser_ common.Address, token_ common.Address, wrapped_ common.Address) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "initialize", admin_, pauser_, token_, wrapped_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address pauser_, address token_, address wrapped_) returns()
func (_Lockbox *LockboxSession) Initialize(admin_ common.Address, pauser_ common.Address, token_ common.Address, wrapped_ common.Address) (*types.Transaction, error) {
	return _Lockbox.Contract.Initialize(&_Lockbox.TransactOpts, admin_, pauser_, token_, wrapped_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address pauser_, address token_, address wrapped_) returns()
func (_Lockbox *LockboxTransactorSession) Initialize(admin_ common.Address, pauser_ common.Address, token_ common.Address, wrapped_ common.Address) (*types.Transaction, error) {
	return _Lockbox.Contract.Initialize(&_Lockbox.TransactOpts, admin_, pauser_, token_, wrapped_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lockbox *LockboxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lockbox *LockboxSession) Pause() (*types.Transaction, error) {
	return _Lockbox.Contract.Pause(&_Lockbox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lockbox *LockboxTransactorSession) Pause() (*types.Transaction, error) {
	return _Lockbox.Contract.Pause(&_Lockbox.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lockbox *LockboxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lockbox *LockboxSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lockbox.Contract.RenounceRole(&_Lockbox.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lockbox *LockboxTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lockbox.Contract.RenounceRole(&_Lockbox.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lockbox *LockboxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lockbox *LockboxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lockbox.Contract.RevokeRole(&_Lockbox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lockbox *LockboxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lockbox.Contract.RevokeRole(&_Lockbox.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lockbox *LockboxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lockbox *LockboxSession) Unpause() (*types.Transaction, error) {
	return _Lockbox.Contract.Unpause(&_Lockbox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lockbox *LockboxTransactorSession) Unpause() (*types.Transaction, error) {
	return _Lockbox.Contract.Unpause(&_Lockbox.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Lockbox *LockboxTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Lockbox *LockboxSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Lockbox.Contract.Withdraw(&_Lockbox.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Lockbox *LockboxTransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Lockbox.Contract.Withdraw(&_Lockbox.TransactOpts, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_Lockbox *LockboxTransactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lockbox.contract.Transact(opts, "withdrawTo", to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_Lockbox *LockboxSession) WithdrawTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lockbox.Contract.WithdrawTo(&_Lockbox.TransactOpts, to, value)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 value) returns()
func (_Lockbox *LockboxTransactorSession) WithdrawTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lockbox.Contract.WithdrawTo(&_Lockbox.TransactOpts, to, value)
}

// LockboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lockbox contract.
type LockboxInitializedIterator struct {
	Event *LockboxInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockboxInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockboxInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockboxInitialized represents a Initialized event raised by the Lockbox contract.
type LockboxInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lockbox *LockboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*LockboxInitializedIterator, error) {

	logs, sub, err := _Lockbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LockboxInitializedIterator{contract: _Lockbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lockbox *LockboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LockboxInitialized) (event.Subscription, error) {

	logs, sub, err := _Lockbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockboxInitialized)
				if err := _Lockbox.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Lockbox *LockboxFilterer) ParseInitialized(log types.Log) (*LockboxInitialized, error) {
	event := new(LockboxInitialized)
	if err := _Lockbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockboxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Lockbox contract.
type LockboxPausedIterator struct {
	Event *LockboxPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockboxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockboxPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockboxPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockboxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockboxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockboxPaused represents a Paused event raised by the Lockbox contract.
type LockboxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lockbox *LockboxFilterer) FilterPaused(opts *bind.FilterOpts) (*LockboxPausedIterator, error) {

	logs, sub, err := _Lockbox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LockboxPausedIterator{contract: _Lockbox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lockbox *LockboxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LockboxPaused) (event.Subscription, error) {

	logs, sub, err := _Lockbox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockboxPaused)
				if err := _Lockbox.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lockbox *LockboxFilterer) ParsePaused(log types.Log) (*LockboxPaused, error) {
	event := new(LockboxPaused)
	if err := _Lockbox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockboxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Lockbox contract.
type LockboxRoleAdminChangedIterator struct {
	Event *LockboxRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockboxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockboxRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockboxRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockboxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockboxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockboxRoleAdminChanged represents a RoleAdminChanged event raised by the Lockbox contract.
type LockboxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Lockbox *LockboxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LockboxRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Lockbox.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LockboxRoleAdminChangedIterator{contract: _Lockbox.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Lockbox *LockboxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LockboxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Lockbox.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockboxRoleAdminChanged)
				if err := _Lockbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Lockbox *LockboxFilterer) ParseRoleAdminChanged(log types.Log) (*LockboxRoleAdminChanged, error) {
	event := new(LockboxRoleAdminChanged)
	if err := _Lockbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockboxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Lockbox contract.
type LockboxRoleGrantedIterator struct {
	Event *LockboxRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockboxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockboxRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockboxRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockboxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockboxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockboxRoleGranted represents a RoleGranted event raised by the Lockbox contract.
type LockboxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lockbox *LockboxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LockboxRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lockbox.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LockboxRoleGrantedIterator{contract: _Lockbox.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lockbox *LockboxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LockboxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lockbox.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockboxRoleGranted)
				if err := _Lockbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lockbox *LockboxFilterer) ParseRoleGranted(log types.Log) (*LockboxRoleGranted, error) {
	event := new(LockboxRoleGranted)
	if err := _Lockbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockboxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Lockbox contract.
type LockboxRoleRevokedIterator struct {
	Event *LockboxRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockboxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockboxRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockboxRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockboxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockboxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockboxRoleRevoked represents a RoleRevoked event raised by the Lockbox contract.
type LockboxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lockbox *LockboxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LockboxRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lockbox.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LockboxRoleRevokedIterator{contract: _Lockbox.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lockbox *LockboxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LockboxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lockbox.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockboxRoleRevoked)
				if err := _Lockbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lockbox *LockboxFilterer) ParseRoleRevoked(log types.Log) (*LockboxRoleRevoked, error) {
	event := new(LockboxRoleRevoked)
	if err := _Lockbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockboxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Lockbox contract.
type LockboxUnpausedIterator struct {
	Event *LockboxUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockboxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockboxUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockboxUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockboxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockboxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockboxUnpaused represents a Unpaused event raised by the Lockbox contract.
type LockboxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lockbox *LockboxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LockboxUnpausedIterator, error) {

	logs, sub, err := _Lockbox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LockboxUnpausedIterator{contract: _Lockbox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lockbox *LockboxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LockboxUnpaused) (event.Subscription, error) {

	logs, sub, err := _Lockbox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockboxUnpaused)
				if err := _Lockbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lockbox *LockboxFilterer) ParseUnpaused(log types.Log) (*LockboxUnpaused, error) {
	event := new(LockboxUnpaused)
	if err := _Lockbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
