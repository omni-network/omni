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

// OmniBridgeL1MetaData contains all meta data concerning the OmniBridgeL1 contract.
var OmniBridgeL1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"XCALL_WITHDRAW_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalL1Supply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610f48380380610f4883398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051610eaf6100996000396000818161023301528181610779015261098e0152610eaf6000f3fe6080604052600436106100a75760003560e01c806390fd50b31161006457806390fd50b31461019c5780639c5451b0146101af578063c3de453d146101ce578063f2fde38b146101e1578063f3fef3a314610201578063fc0c546a1461022157600080fd5b806339acf9f1146100ac5780633abfe55f146100e9578063485cc95514610117578063715018a6146101395780638da5cb5b1461014e5780638fdcb4c91461016c575b600080fd5b3480156100b857600080fd5b506065546100cc906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100f557600080fd5b50610109610104366004610be8565b610255565b6040519081526020016100e0565b34801561012357600080fd5b50610137610132366004610c14565b61037f565b005b34801561014557600080fd5b506101376104bb565b34801561015a57600080fd5b506033546001600160a01b03166100cc565b34801561017857600080fd5b50610183620249f081565b60405167ffffffffffffffff90911681526020016100e0565b6101376101aa366004610c4d565b6104cf565b3480156101bb57600080fd5b506101096a52b7d2dcc80cd2e400000081565b6101376101dc366004610be8565b6104dd565b3480156101ed57600080fd5b506101376101fc366004610c66565b6104ec565b34801561020d57600080fd5b5061013761021c366004610be8565b610562565b34801561022d57600080fd5b506100cc7f000000000000000000000000000000000000000000000000000000000000000081565b6065546040805163110ff5f160e01b815290516000926001600160a01b031691638dd9523c91839163110ff5f19160048083019260209291908290030181865afa1580156102a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cb9190610ca0565b604080516001600160a01b0388166024820152604480820188905282518083039091018152606490910182526020810180516001600160e01b031663f3fef3a360e01b179052905160e084901b6001600160e01b0319168152610337929190620249f090600401610d01565b602060405180830381865afa158015610354573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103789190610d38565b9392505050565b600054610100900460ff161580801561039f5750600054600160ff909116105b806103b95750303b1580156103b9575060005460ff166001145b6104215760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610444576000805461ff0019166101001790555b61044c61082f565b610455836104ec565b606580546001600160a01b0319166001600160a01b03841617905580156104b6576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6104c361085e565b6104cd60006108b8565b565b6104da33338361090a565b50565b6104e833838361090a565b5050565b6104f461085e565b6001600160a01b0381166105595760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610418565b6104da816108b8565b60655460408051631799380760e11b815281516000936001600160a01b031692632f32700e92600480820193918290030181865afa1580156105a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105cc9190610d51565b6065549091506001600160a01b031633146106215760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b6044820152606401610418565b60208101516001600160a01b0316600462048789608a1b011461067f5760405162461bcd60e51b81526020600482015260166024820152754f6d6e694272696467653a206e6f742062726964676560501b6044820152606401610418565b606560009054906101000a90046001600160a01b03166001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f69190610ca0565b67ffffffffffffffff16816000015167ffffffffffffffff16146107535760405162461bcd60e51b81526020600482015260146024820152734f6d6e694272696467653a206e6f74206f6d6e6960601b6044820152606401610418565b60405163a9059cbb60e01b81526001600160a01b038481166004830152602482018490527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af11580156107c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e69190610dbc565b50826001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161082291815260200190565b60405180910390a2505050565b600054610100900460ff166108565760405162461bcd60e51b815260040161041890610dde565b6104cd610ba3565b6033546001600160a01b031633146104cd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610418565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6109148282610255565b34146109625760405162461bcd60e51b815260206004820152601760248201527f4f6d6e694272696467653a20696e76616c6964206665650000000000000000006044820152606401610418565b6040516323b872dd60e01b81526001600160a01b038481166004830152306024830152604482018390527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af11580156109d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109fb9190610dbc565b610a475760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e694272696467653a207472616e73666572206661696c656400000000006044820152606401610418565b6065546040805163110ff5f160e01b815290516001600160a01b039092169163c21dda4f91839163110ff5f1916004808201926020929091908290030181865afa158015610a99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abd9190610ca0565b604080516001600160a01b0387166024820152604480820187905282518083039091018152606490910182526020810180516001600160e01b031663f3fef3a360e01b179052905160e084901b6001600160e01b0319168152610b369291600491600462048789608a1b019190620249f0908401610e29565b600060405180830381600087803b158015610b5057600080fd5b505af1158015610b64573d6000803e3d6000fd5b50505050826001600160a01b03167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c8260405161082291815260200190565b600054610100900460ff16610bca5760405162461bcd60e51b815260040161041890610dde565b6104cd336108b8565b6001600160a01b03811681146104da57600080fd5b60008060408385031215610bfb57600080fd5b8235610c0681610bd3565b946020939093013593505050565b60008060408385031215610c2757600080fd5b8235610c3281610bd3565b91506020830135610c4281610bd3565b809150509250929050565b600060208284031215610c5f57600080fd5b5035919050565b600060208284031215610c7857600080fd5b813561037881610bd3565b805167ffffffffffffffff81168114610c9b57600080fd5b919050565b600060208284031215610cb257600080fd5b61037882610c83565b6000815180845260005b81811015610ce157602081850181015186830182015201610cc5565b506000602082860101526020601f19601f83011685010191505092915050565b600067ffffffffffffffff808616835260606020840152610d256060840186610cbb565b9150808416604084015250949350505050565b600060208284031215610d4a57600080fd5b5051919050565b600060408284031215610d6357600080fd5b6040516040810181811067ffffffffffffffff82111715610d9457634e487b7160e01b600052604160045260246000fd5b604052610da083610c83565b81526020830151610db081610bd3565b60208201529392505050565b600060208284031215610dce57600080fd5b8151801515811461037857600080fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600067ffffffffffffffff808816835260ff8716602084015260018060a01b038616604084015260a06060840152610e6460a0840186610cbb565b9150808416608084015250969550505050505056fea2646970667358221220c81342524045b37cc6735dff0a199e03f2d27655120bcaea3c503822a527a13064736f6c63430008180033",
}

// OmniBridgeL1ABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniBridgeL1MetaData.ABI instead.
var OmniBridgeL1ABI = OmniBridgeL1MetaData.ABI

// OmniBridgeL1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniBridgeL1MetaData.Bin instead.
var OmniBridgeL1Bin = OmniBridgeL1MetaData.Bin

// DeployOmniBridgeL1 deploys a new Ethereum contract, binding an instance of OmniBridgeL1 to it.
func DeployOmniBridgeL1(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address) (common.Address, *types.Transaction, *OmniBridgeL1, error) {
	parsed, err := OmniBridgeL1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniBridgeL1Bin), backend, token_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniBridgeL1{OmniBridgeL1Caller: OmniBridgeL1Caller{contract: contract}, OmniBridgeL1Transactor: OmniBridgeL1Transactor{contract: contract}, OmniBridgeL1Filterer: OmniBridgeL1Filterer{contract: contract}}, nil
}

// OmniBridgeL1 is an auto generated Go binding around an Ethereum contract.
type OmniBridgeL1 struct {
	OmniBridgeL1Caller     // Read-only binding to the contract
	OmniBridgeL1Transactor // Write-only binding to the contract
	OmniBridgeL1Filterer   // Log filterer for contract events
}

// OmniBridgeL1Caller is an auto generated read-only Go binding around an Ethereum contract.
type OmniBridgeL1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniBridgeL1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniBridgeL1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeL1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniBridgeL1Session struct {
	Contract     *OmniBridgeL1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniBridgeL1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniBridgeL1CallerSession struct {
	Contract *OmniBridgeL1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OmniBridgeL1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniBridgeL1TransactorSession struct {
	Contract     *OmniBridgeL1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OmniBridgeL1Raw is an auto generated low-level Go binding around an Ethereum contract.
type OmniBridgeL1Raw struct {
	Contract *OmniBridgeL1 // Generic contract binding to access the raw methods on
}

// OmniBridgeL1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniBridgeL1CallerRaw struct {
	Contract *OmniBridgeL1Caller // Generic read-only contract binding to access the raw methods on
}

// OmniBridgeL1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniBridgeL1TransactorRaw struct {
	Contract *OmniBridgeL1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniBridgeL1 creates a new instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1(address common.Address, backend bind.ContractBackend) (*OmniBridgeL1, error) {
	contract, err := bindOmniBridgeL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1{OmniBridgeL1Caller: OmniBridgeL1Caller{contract: contract}, OmniBridgeL1Transactor: OmniBridgeL1Transactor{contract: contract}, OmniBridgeL1Filterer: OmniBridgeL1Filterer{contract: contract}}, nil
}

// NewOmniBridgeL1Caller creates a new read-only instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Caller(address common.Address, caller bind.ContractCaller) (*OmniBridgeL1Caller, error) {
	contract, err := bindOmniBridgeL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Caller{contract: contract}, nil
}

// NewOmniBridgeL1Transactor creates a new write-only instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Transactor(address common.Address, transactor bind.ContractTransactor) (*OmniBridgeL1Transactor, error) {
	contract, err := bindOmniBridgeL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Transactor{contract: contract}, nil
}

// NewOmniBridgeL1Filterer creates a new log filterer instance of OmniBridgeL1, bound to a specific deployed contract.
func NewOmniBridgeL1Filterer(address common.Address, filterer bind.ContractFilterer) (*OmniBridgeL1Filterer, error) {
	contract, err := bindOmniBridgeL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1Filterer{contract: contract}, nil
}

// bindOmniBridgeL1 binds a generic wrapper to an already deployed contract.
func bindOmniBridgeL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniBridgeL1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniBridgeL1.Contract.OmniBridgeL1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.OmniBridgeL1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniBridgeL1 *OmniBridgeL1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.OmniBridgeL1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniBridgeL1 *OmniBridgeL1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniBridgeL1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniBridgeL1 *OmniBridgeL1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniBridgeL1 *OmniBridgeL1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.contract.Transact(opts, method, params...)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1Caller) XCALLWITHDRAWGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "XCALL_WITHDRAW_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1Session) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _OmniBridgeL1.Contract.XCALLWITHDRAWGASLIMIT(&_OmniBridgeL1.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _OmniBridgeL1.Contract.XCALLWITHDRAWGASLIMIT(&_OmniBridgeL1.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Caller) BridgeFee(opts *bind.CallOpts, to common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "bridgeFee", to, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Session) BridgeFee(to common.Address, amount *big.Int) (*big.Int, error) {
	return _OmniBridgeL1.Contract.BridgeFee(&_OmniBridgeL1.CallOpts, to, amount)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) BridgeFee(to common.Address, amount *big.Int) (*big.Int, error) {
	return _OmniBridgeL1.Contract.BridgeFee(&_OmniBridgeL1.CallOpts, to, amount)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Omni() (common.Address, error) {
	return _OmniBridgeL1.Contract.Omni(&_OmniBridgeL1.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Omni() (common.Address, error) {
	return _OmniBridgeL1.Contract.Omni(&_OmniBridgeL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Owner() (common.Address, error) {
	return _OmniBridgeL1.Contract.Owner(&_OmniBridgeL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Owner() (common.Address, error) {
	return _OmniBridgeL1.Contract.Owner(&_OmniBridgeL1.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1Session) Token() (common.Address, error) {
	return _OmniBridgeL1.Contract.Token(&_OmniBridgeL1.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) Token() (common.Address, error) {
	return _OmniBridgeL1.Contract.Token(&_OmniBridgeL1.CallOpts)
}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Caller) TotalL1Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeL1.contract.Call(opts, &out, "totalL1Supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1Session) TotalL1Supply() (*big.Int, error) {
	return _OmniBridgeL1.Contract.TotalL1Supply(&_OmniBridgeL1.CallOpts)
}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeL1 *OmniBridgeL1CallerSession) TotalL1Supply() (*big.Int, error) {
	return _OmniBridgeL1.Contract.TotalL1Supply(&_OmniBridgeL1.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0x90fd50b3.
//
// Solidity: function bridge(uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Bridge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "bridge", amount)
}

// Bridge is a paid mutator transaction binding the contract method 0x90fd50b3.
//
// Solidity: function bridge(uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Bridge(amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Bridge(&_OmniBridgeL1.TransactOpts, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0x90fd50b3.
//
// Solidity: function bridge(uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Bridge(amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Bridge(&_OmniBridgeL1.TransactOpts, amount)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Bridge0(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "bridge0", to, amount)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Bridge0(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Bridge0(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Bridge0(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Bridge0(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "initialize", owner_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Initialize(owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Initialize(&_OmniBridgeL1.TransactOpts, owner_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address omni_) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Initialize(owner_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Initialize(&_OmniBridgeL1.TransactOpts, owner_, omni_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) RenounceOwnership() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.RenounceOwnership(&_OmniBridgeL1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.RenounceOwnership(&_OmniBridgeL1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.TransferOwnership(&_OmniBridgeL1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.TransferOwnership(&_OmniBridgeL1.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1Transactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1Session) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Withdraw(&_OmniBridgeL1.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_OmniBridgeL1 *OmniBridgeL1TransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeL1.Contract.Withdraw(&_OmniBridgeL1.TransactOpts, to, amount)
}

// OmniBridgeL1DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the OmniBridgeL1 contract.
type OmniBridgeL1DepositIterator struct {
	Event *OmniBridgeL1Deposit // Event containing the contract specifics and raw log

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
func (it *OmniBridgeL1DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Deposit)
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
		it.Event = new(OmniBridgeL1Deposit)
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
func (it *OmniBridgeL1DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Deposit represents a Deposit event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Deposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address) (*OmniBridgeL1DepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1DepositIterator{contract: _OmniBridgeL1.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Deposit, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Deposit)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseDeposit(log types.Log) (*OmniBridgeL1Deposit, error) {
	event := new(OmniBridgeL1Deposit)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniBridgeL1 contract.
type OmniBridgeL1InitializedIterator struct {
	Event *OmniBridgeL1Initialized // Event containing the contract specifics and raw log

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
func (it *OmniBridgeL1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Initialized)
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
		it.Event = new(OmniBridgeL1Initialized)
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
func (it *OmniBridgeL1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Initialized represents a Initialized event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterInitialized(opts *bind.FilterOpts) (*OmniBridgeL1InitializedIterator, error) {

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1InitializedIterator{contract: _OmniBridgeL1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Initialized) (event.Subscription, error) {

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Initialized)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseInitialized(log types.Log) (*OmniBridgeL1Initialized, error) {
	event := new(OmniBridgeL1Initialized)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniBridgeL1 contract.
type OmniBridgeL1OwnershipTransferredIterator struct {
	Event *OmniBridgeL1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniBridgeL1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1OwnershipTransferred)
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
		it.Event = new(OmniBridgeL1OwnershipTransferred)
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
func (it *OmniBridgeL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1OwnershipTransferred represents a OwnershipTransferred event raised by the OmniBridgeL1 contract.
type OmniBridgeL1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniBridgeL1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1OwnershipTransferredIterator{contract: _OmniBridgeL1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1OwnershipTransferred)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseOwnershipTransferred(log types.Log) (*OmniBridgeL1OwnershipTransferred, error) {
	event := new(OmniBridgeL1OwnershipTransferred)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeL1WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the OmniBridgeL1 contract.
type OmniBridgeL1WithdrawIterator struct {
	Event *OmniBridgeL1Withdraw // Event containing the contract specifics and raw log

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
func (it *OmniBridgeL1WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeL1Withdraw)
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
		it.Event = new(OmniBridgeL1Withdraw)
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
func (it *OmniBridgeL1WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeL1WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeL1Withdraw represents a Withdraw event raised by the OmniBridgeL1 contract.
type OmniBridgeL1Withdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*OmniBridgeL1WithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeL1WithdrawIterator{contract: _OmniBridgeL1.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *OmniBridgeL1Withdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeL1.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeL1Withdraw)
				if err := _OmniBridgeL1.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_OmniBridgeL1 *OmniBridgeL1Filterer) ParseWithdraw(log types.Log) (*OmniBridgeL1Withdraw, error) {
	event := new(OmniBridgeL1Withdraw)
	if err := _OmniBridgeL1.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
