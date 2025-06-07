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

// IAllowanceTransferAllowanceTransferDetails is an auto generated low-level Go binding around an user-defined struct.
type IAllowanceTransferAllowanceTransferDetails struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Token  common.Address
}

// IAllowanceTransferPermitBatch is an auto generated low-level Go binding around an user-defined struct.
type IAllowanceTransferPermitBatch struct {
	Details     []IAllowanceTransferPermitDetails
	Spender     common.Address
	SigDeadline *big.Int
}

// IAllowanceTransferPermitDetails is an auto generated low-level Go binding around an user-defined struct.
type IAllowanceTransferPermitDetails struct {
	Token      common.Address
	Amount     *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}

// IAllowanceTransferPermitSingle is an auto generated low-level Go binding around an user-defined struct.
type IAllowanceTransferPermitSingle struct {
	Details     IAllowanceTransferPermitDetails
	Spender     common.Address
	SigDeadline *big.Int
}

// IAllowanceTransferTokenSpenderPair is an auto generated low-level Go binding around an user-defined struct.
type IAllowanceTransferTokenSpenderPair struct {
	Token   common.Address
	Spender common.Address
}

// ISignatureTransferPermitBatchTransferFrom is an auto generated low-level Go binding around an user-defined struct.
type ISignatureTransferPermitBatchTransferFrom struct {
	Permitted []ISignatureTransferTokenPermissions
	Nonce     *big.Int
	Deadline  *big.Int
}

// ISignatureTransferPermitTransferFrom is an auto generated low-level Go binding around an user-defined struct.
type ISignatureTransferPermitTransferFrom struct {
	Permitted ISignatureTransferTokenPermissions
	Nonce     *big.Int
	Deadline  *big.Int
}

// ISignatureTransferSignatureTransferDetails is an auto generated low-level Go binding around an user-defined struct.
type ISignatureTransferSignatureTransferDetails struct {
	To              common.Address
	RequestedAmount *big.Int
}

// ISignatureTransferTokenPermissions is an auto generated low-level Go binding around an user-defined struct.
type ISignatureTransferTokenPermissions struct {
	Token  common.Address
	Amount *big.Int
}

// IPermit2MetaData contains all meta data concerning the IPermit2 contract.
var IPermit2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"nonce\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"invalidateNonces\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newNonce\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"invalidateUnorderedNonces\",\"inputs\":[{\"name\":\"wordPos\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockdown\",\"inputs\":[{\"name\":\"approvals\",\"type\":\"tuple[]\",\"internalType\":\"structIAllowanceTransfer.TokenSpenderPair[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonceBitmap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permitBatch\",\"type\":\"tuple\",\"internalType\":\"structIAllowanceTransfer.PermitBatch\",\"components\":[{\"name\":\"details\",\"type\":\"tuple[]\",\"internalType\":\"structIAllowanceTransfer.PermitDetails[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"nonce\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sigDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permitSingle\",\"type\":\"tuple\",\"internalType\":\"structIAllowanceTransfer.PermitSingle\",\"components\":[{\"name\":\"details\",\"type\":\"tuple\",\"internalType\":\"structIAllowanceTransfer.PermitDetails\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"nonce\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sigDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"permitTransferFrom\",\"inputs\":[{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.PermitTransferFrom\",\"components\":[{\"name\":\"permitted\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.TokenPermissions\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"transferDetails\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.SignatureTransferDetails\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"permitTransferFrom\",\"inputs\":[{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.PermitBatchTransferFrom\",\"components\":[{\"name\":\"permitted\",\"type\":\"tuple[]\",\"internalType\":\"structISignatureTransfer.TokenPermissions[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"transferDetails\",\"type\":\"tuple[]\",\"internalType\":\"structISignatureTransfer.SignatureTransferDetails[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"permitWitnessTransferFrom\",\"inputs\":[{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.PermitTransferFrom\",\"components\":[{\"name\":\"permitted\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.TokenPermissions\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"transferDetails\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.SignatureTransferDetails\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"witness\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"witnessTypeString\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"permitWitnessTransferFrom\",\"inputs\":[{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structISignatureTransfer.PermitBatchTransferFrom\",\"components\":[{\"name\":\"permitted\",\"type\":\"tuple[]\",\"internalType\":\"structISignatureTransfer.TokenPermissions[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"transferDetails\",\"type\":\"tuple[]\",\"internalType\":\"structISignatureTransfer.SignatureTransferDetails[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"witness\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"witnessTypeString\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"transferDetails\",\"type\":\"tuple[]\",\"internalType\":\"structIAllowanceTransfer.AllowanceTransferDetails[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Lockdown\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceInvalidation\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newNonce\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"oldNonce\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"expiration\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"nonce\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnorderedNonceInvalidation\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"word\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mask\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowanceExpired\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExcessiveInvalidation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllowance\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[{\"name\":\"maxAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[]}]",
}

// IPermit2ABI is the input ABI used to generate the binding from.
// Deprecated: Use IPermit2MetaData.ABI instead.
var IPermit2ABI = IPermit2MetaData.ABI

// IPermit2 is an auto generated Go binding around an Ethereum contract.
type IPermit2 struct {
	IPermit2Caller     // Read-only binding to the contract
	IPermit2Transactor // Write-only binding to the contract
	IPermit2Filterer   // Log filterer for contract events
}

// IPermit2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IPermit2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPermit2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IPermit2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPermit2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPermit2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPermit2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPermit2Session struct {
	Contract     *IPermit2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPermit2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPermit2CallerSession struct {
	Contract *IPermit2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IPermit2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPermit2TransactorSession struct {
	Contract     *IPermit2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IPermit2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IPermit2Raw struct {
	Contract *IPermit2 // Generic contract binding to access the raw methods on
}

// IPermit2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPermit2CallerRaw struct {
	Contract *IPermit2Caller // Generic read-only contract binding to access the raw methods on
}

// IPermit2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPermit2TransactorRaw struct {
	Contract *IPermit2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIPermit2 creates a new instance of IPermit2, bound to a specific deployed contract.
func NewIPermit2(address common.Address, backend bind.ContractBackend) (*IPermit2, error) {
	contract, err := bindIPermit2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPermit2{IPermit2Caller: IPermit2Caller{contract: contract}, IPermit2Transactor: IPermit2Transactor{contract: contract}, IPermit2Filterer: IPermit2Filterer{contract: contract}}, nil
}

// NewIPermit2Caller creates a new read-only instance of IPermit2, bound to a specific deployed contract.
func NewIPermit2Caller(address common.Address, caller bind.ContractCaller) (*IPermit2Caller, error) {
	contract, err := bindIPermit2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPermit2Caller{contract: contract}, nil
}

// NewIPermit2Transactor creates a new write-only instance of IPermit2, bound to a specific deployed contract.
func NewIPermit2Transactor(address common.Address, transactor bind.ContractTransactor) (*IPermit2Transactor, error) {
	contract, err := bindIPermit2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPermit2Transactor{contract: contract}, nil
}

// NewIPermit2Filterer creates a new log filterer instance of IPermit2, bound to a specific deployed contract.
func NewIPermit2Filterer(address common.Address, filterer bind.ContractFilterer) (*IPermit2Filterer, error) {
	contract, err := bindIPermit2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPermit2Filterer{contract: contract}, nil
}

// bindIPermit2 binds a generic wrapper to an already deployed contract.
func bindIPermit2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPermit2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPermit2 *IPermit2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPermit2.Contract.IPermit2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPermit2 *IPermit2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPermit2.Contract.IPermit2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPermit2 *IPermit2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPermit2.Contract.IPermit2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPermit2 *IPermit2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPermit2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPermit2 *IPermit2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPermit2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPermit2 *IPermit2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPermit2.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPermit2 *IPermit2Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IPermit2.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPermit2 *IPermit2Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _IPermit2.Contract.DOMAINSEPARATOR(&_IPermit2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IPermit2 *IPermit2CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IPermit2.Contract.DOMAINSEPARATOR(&_IPermit2.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0x927da105.
//
// Solidity: function allowance(address user, address token, address spender) view returns(uint160 amount, uint48 expiration, uint48 nonce)
func (_IPermit2 *IPermit2Caller) Allowance(opts *bind.CallOpts, user common.Address, token common.Address, spender common.Address) (struct {
	Amount     *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}, error) {
	var out []interface{}
	err := _IPermit2.contract.Call(opts, &out, "allowance", user, token, spender)

	outstruct := new(struct {
		Amount     *big.Int
		Expiration *big.Int
		Nonce      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Allowance is a free data retrieval call binding the contract method 0x927da105.
//
// Solidity: function allowance(address user, address token, address spender) view returns(uint160 amount, uint48 expiration, uint48 nonce)
func (_IPermit2 *IPermit2Session) Allowance(user common.Address, token common.Address, spender common.Address) (struct {
	Amount     *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}, error) {
	return _IPermit2.Contract.Allowance(&_IPermit2.CallOpts, user, token, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x927da105.
//
// Solidity: function allowance(address user, address token, address spender) view returns(uint160 amount, uint48 expiration, uint48 nonce)
func (_IPermit2 *IPermit2CallerSession) Allowance(user common.Address, token common.Address, spender common.Address) (struct {
	Amount     *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}, error) {
	return _IPermit2.Contract.Allowance(&_IPermit2.CallOpts, user, token, spender)
}

// NonceBitmap is a free data retrieval call binding the contract method 0x4fe02b44.
//
// Solidity: function nonceBitmap(address , uint256 ) view returns(uint256)
func (_IPermit2 *IPermit2Caller) NonceBitmap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPermit2.contract.Call(opts, &out, "nonceBitmap", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceBitmap is a free data retrieval call binding the contract method 0x4fe02b44.
//
// Solidity: function nonceBitmap(address , uint256 ) view returns(uint256)
func (_IPermit2 *IPermit2Session) NonceBitmap(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _IPermit2.Contract.NonceBitmap(&_IPermit2.CallOpts, arg0, arg1)
}

// NonceBitmap is a free data retrieval call binding the contract method 0x4fe02b44.
//
// Solidity: function nonceBitmap(address , uint256 ) view returns(uint256)
func (_IPermit2 *IPermit2CallerSession) NonceBitmap(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _IPermit2.Contract.NonceBitmap(&_IPermit2.CallOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x87517c45.
//
// Solidity: function approve(address token, address spender, uint160 amount, uint48 expiration) returns()
func (_IPermit2 *IPermit2Transactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address, amount *big.Int, expiration *big.Int) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "approve", token, spender, amount, expiration)
}

// Approve is a paid mutator transaction binding the contract method 0x87517c45.
//
// Solidity: function approve(address token, address spender, uint160 amount, uint48 expiration) returns()
func (_IPermit2 *IPermit2Session) Approve(token common.Address, spender common.Address, amount *big.Int, expiration *big.Int) (*types.Transaction, error) {
	return _IPermit2.Contract.Approve(&_IPermit2.TransactOpts, token, spender, amount, expiration)
}

// Approve is a paid mutator transaction binding the contract method 0x87517c45.
//
// Solidity: function approve(address token, address spender, uint160 amount, uint48 expiration) returns()
func (_IPermit2 *IPermit2TransactorSession) Approve(token common.Address, spender common.Address, amount *big.Int, expiration *big.Int) (*types.Transaction, error) {
	return _IPermit2.Contract.Approve(&_IPermit2.TransactOpts, token, spender, amount, expiration)
}

// InvalidateNonces is a paid mutator transaction binding the contract method 0x65d9723c.
//
// Solidity: function invalidateNonces(address token, address spender, uint48 newNonce) returns()
func (_IPermit2 *IPermit2Transactor) InvalidateNonces(opts *bind.TransactOpts, token common.Address, spender common.Address, newNonce *big.Int) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "invalidateNonces", token, spender, newNonce)
}

// InvalidateNonces is a paid mutator transaction binding the contract method 0x65d9723c.
//
// Solidity: function invalidateNonces(address token, address spender, uint48 newNonce) returns()
func (_IPermit2 *IPermit2Session) InvalidateNonces(token common.Address, spender common.Address, newNonce *big.Int) (*types.Transaction, error) {
	return _IPermit2.Contract.InvalidateNonces(&_IPermit2.TransactOpts, token, spender, newNonce)
}

// InvalidateNonces is a paid mutator transaction binding the contract method 0x65d9723c.
//
// Solidity: function invalidateNonces(address token, address spender, uint48 newNonce) returns()
func (_IPermit2 *IPermit2TransactorSession) InvalidateNonces(token common.Address, spender common.Address, newNonce *big.Int) (*types.Transaction, error) {
	return _IPermit2.Contract.InvalidateNonces(&_IPermit2.TransactOpts, token, spender, newNonce)
}

// InvalidateUnorderedNonces is a paid mutator transaction binding the contract method 0x3ff9dcb1.
//
// Solidity: function invalidateUnorderedNonces(uint256 wordPos, uint256 mask) returns()
func (_IPermit2 *IPermit2Transactor) InvalidateUnorderedNonces(opts *bind.TransactOpts, wordPos *big.Int, mask *big.Int) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "invalidateUnorderedNonces", wordPos, mask)
}

// InvalidateUnorderedNonces is a paid mutator transaction binding the contract method 0x3ff9dcb1.
//
// Solidity: function invalidateUnorderedNonces(uint256 wordPos, uint256 mask) returns()
func (_IPermit2 *IPermit2Session) InvalidateUnorderedNonces(wordPos *big.Int, mask *big.Int) (*types.Transaction, error) {
	return _IPermit2.Contract.InvalidateUnorderedNonces(&_IPermit2.TransactOpts, wordPos, mask)
}

// InvalidateUnorderedNonces is a paid mutator transaction binding the contract method 0x3ff9dcb1.
//
// Solidity: function invalidateUnorderedNonces(uint256 wordPos, uint256 mask) returns()
func (_IPermit2 *IPermit2TransactorSession) InvalidateUnorderedNonces(wordPos *big.Int, mask *big.Int) (*types.Transaction, error) {
	return _IPermit2.Contract.InvalidateUnorderedNonces(&_IPermit2.TransactOpts, wordPos, mask)
}

// Lockdown is a paid mutator transaction binding the contract method 0xcc53287f.
//
// Solidity: function lockdown((address,address)[] approvals) returns()
func (_IPermit2 *IPermit2Transactor) Lockdown(opts *bind.TransactOpts, approvals []IAllowanceTransferTokenSpenderPair) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "lockdown", approvals)
}

// Lockdown is a paid mutator transaction binding the contract method 0xcc53287f.
//
// Solidity: function lockdown((address,address)[] approvals) returns()
func (_IPermit2 *IPermit2Session) Lockdown(approvals []IAllowanceTransferTokenSpenderPair) (*types.Transaction, error) {
	return _IPermit2.Contract.Lockdown(&_IPermit2.TransactOpts, approvals)
}

// Lockdown is a paid mutator transaction binding the contract method 0xcc53287f.
//
// Solidity: function lockdown((address,address)[] approvals) returns()
func (_IPermit2 *IPermit2TransactorSession) Lockdown(approvals []IAllowanceTransferTokenSpenderPair) (*types.Transaction, error) {
	return _IPermit2.Contract.Lockdown(&_IPermit2.TransactOpts, approvals)
}

// Permit is a paid mutator transaction binding the contract method 0x2a2d80d1.
//
// Solidity: function permit(address owner, ((address,uint160,uint48,uint48)[],address,uint256) permitBatch, bytes signature) returns()
func (_IPermit2 *IPermit2Transactor) Permit(opts *bind.TransactOpts, owner common.Address, permitBatch IAllowanceTransferPermitBatch, signature []byte) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "permit", owner, permitBatch, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x2a2d80d1.
//
// Solidity: function permit(address owner, ((address,uint160,uint48,uint48)[],address,uint256) permitBatch, bytes signature) returns()
func (_IPermit2 *IPermit2Session) Permit(owner common.Address, permitBatch IAllowanceTransferPermitBatch, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.Permit(&_IPermit2.TransactOpts, owner, permitBatch, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x2a2d80d1.
//
// Solidity: function permit(address owner, ((address,uint160,uint48,uint48)[],address,uint256) permitBatch, bytes signature) returns()
func (_IPermit2 *IPermit2TransactorSession) Permit(owner common.Address, permitBatch IAllowanceTransferPermitBatch, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.Permit(&_IPermit2.TransactOpts, owner, permitBatch, signature)
}

// Permit0 is a paid mutator transaction binding the contract method 0x2b67b570.
//
// Solidity: function permit(address owner, ((address,uint160,uint48,uint48),address,uint256) permitSingle, bytes signature) returns()
func (_IPermit2 *IPermit2Transactor) Permit0(opts *bind.TransactOpts, owner common.Address, permitSingle IAllowanceTransferPermitSingle, signature []byte) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "permit0", owner, permitSingle, signature)
}

// Permit0 is a paid mutator transaction binding the contract method 0x2b67b570.
//
// Solidity: function permit(address owner, ((address,uint160,uint48,uint48),address,uint256) permitSingle, bytes signature) returns()
func (_IPermit2 *IPermit2Session) Permit0(owner common.Address, permitSingle IAllowanceTransferPermitSingle, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.Permit0(&_IPermit2.TransactOpts, owner, permitSingle, signature)
}

// Permit0 is a paid mutator transaction binding the contract method 0x2b67b570.
//
// Solidity: function permit(address owner, ((address,uint160,uint48,uint48),address,uint256) permitSingle, bytes signature) returns()
func (_IPermit2 *IPermit2TransactorSession) Permit0(owner common.Address, permitSingle IAllowanceTransferPermitSingle, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.Permit0(&_IPermit2.TransactOpts, owner, permitSingle, signature)
}

// PermitTransferFrom is a paid mutator transaction binding the contract method 0x30f28b7a.
//
// Solidity: function permitTransferFrom(((address,uint256),uint256,uint256) permit, (address,uint256) transferDetails, address owner, bytes signature) returns()
func (_IPermit2 *IPermit2Transactor) PermitTransferFrom(opts *bind.TransactOpts, permit ISignatureTransferPermitTransferFrom, transferDetails ISignatureTransferSignatureTransferDetails, owner common.Address, signature []byte) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "permitTransferFrom", permit, transferDetails, owner, signature)
}

// PermitTransferFrom is a paid mutator transaction binding the contract method 0x30f28b7a.
//
// Solidity: function permitTransferFrom(((address,uint256),uint256,uint256) permit, (address,uint256) transferDetails, address owner, bytes signature) returns()
func (_IPermit2 *IPermit2Session) PermitTransferFrom(permit ISignatureTransferPermitTransferFrom, transferDetails ISignatureTransferSignatureTransferDetails, owner common.Address, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.PermitTransferFrom(&_IPermit2.TransactOpts, permit, transferDetails, owner, signature)
}

// PermitTransferFrom is a paid mutator transaction binding the contract method 0x30f28b7a.
//
// Solidity: function permitTransferFrom(((address,uint256),uint256,uint256) permit, (address,uint256) transferDetails, address owner, bytes signature) returns()
func (_IPermit2 *IPermit2TransactorSession) PermitTransferFrom(permit ISignatureTransferPermitTransferFrom, transferDetails ISignatureTransferSignatureTransferDetails, owner common.Address, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.PermitTransferFrom(&_IPermit2.TransactOpts, permit, transferDetails, owner, signature)
}

// PermitTransferFrom0 is a paid mutator transaction binding the contract method 0xedd9444b.
//
// Solidity: function permitTransferFrom(((address,uint256)[],uint256,uint256) permit, (address,uint256)[] transferDetails, address owner, bytes signature) returns()
func (_IPermit2 *IPermit2Transactor) PermitTransferFrom0(opts *bind.TransactOpts, permit ISignatureTransferPermitBatchTransferFrom, transferDetails []ISignatureTransferSignatureTransferDetails, owner common.Address, signature []byte) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "permitTransferFrom0", permit, transferDetails, owner, signature)
}

// PermitTransferFrom0 is a paid mutator transaction binding the contract method 0xedd9444b.
//
// Solidity: function permitTransferFrom(((address,uint256)[],uint256,uint256) permit, (address,uint256)[] transferDetails, address owner, bytes signature) returns()
func (_IPermit2 *IPermit2Session) PermitTransferFrom0(permit ISignatureTransferPermitBatchTransferFrom, transferDetails []ISignatureTransferSignatureTransferDetails, owner common.Address, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.PermitTransferFrom0(&_IPermit2.TransactOpts, permit, transferDetails, owner, signature)
}

// PermitTransferFrom0 is a paid mutator transaction binding the contract method 0xedd9444b.
//
// Solidity: function permitTransferFrom(((address,uint256)[],uint256,uint256) permit, (address,uint256)[] transferDetails, address owner, bytes signature) returns()
func (_IPermit2 *IPermit2TransactorSession) PermitTransferFrom0(permit ISignatureTransferPermitBatchTransferFrom, transferDetails []ISignatureTransferSignatureTransferDetails, owner common.Address, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.PermitTransferFrom0(&_IPermit2.TransactOpts, permit, transferDetails, owner, signature)
}

// PermitWitnessTransferFrom is a paid mutator transaction binding the contract method 0x137c29fe.
//
// Solidity: function permitWitnessTransferFrom(((address,uint256),uint256,uint256) permit, (address,uint256) transferDetails, address owner, bytes32 witness, string witnessTypeString, bytes signature) returns()
func (_IPermit2 *IPermit2Transactor) PermitWitnessTransferFrom(opts *bind.TransactOpts, permit ISignatureTransferPermitTransferFrom, transferDetails ISignatureTransferSignatureTransferDetails, owner common.Address, witness [32]byte, witnessTypeString string, signature []byte) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "permitWitnessTransferFrom", permit, transferDetails, owner, witness, witnessTypeString, signature)
}

// PermitWitnessTransferFrom is a paid mutator transaction binding the contract method 0x137c29fe.
//
// Solidity: function permitWitnessTransferFrom(((address,uint256),uint256,uint256) permit, (address,uint256) transferDetails, address owner, bytes32 witness, string witnessTypeString, bytes signature) returns()
func (_IPermit2 *IPermit2Session) PermitWitnessTransferFrom(permit ISignatureTransferPermitTransferFrom, transferDetails ISignatureTransferSignatureTransferDetails, owner common.Address, witness [32]byte, witnessTypeString string, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.PermitWitnessTransferFrom(&_IPermit2.TransactOpts, permit, transferDetails, owner, witness, witnessTypeString, signature)
}

// PermitWitnessTransferFrom is a paid mutator transaction binding the contract method 0x137c29fe.
//
// Solidity: function permitWitnessTransferFrom(((address,uint256),uint256,uint256) permit, (address,uint256) transferDetails, address owner, bytes32 witness, string witnessTypeString, bytes signature) returns()
func (_IPermit2 *IPermit2TransactorSession) PermitWitnessTransferFrom(permit ISignatureTransferPermitTransferFrom, transferDetails ISignatureTransferSignatureTransferDetails, owner common.Address, witness [32]byte, witnessTypeString string, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.PermitWitnessTransferFrom(&_IPermit2.TransactOpts, permit, transferDetails, owner, witness, witnessTypeString, signature)
}

// PermitWitnessTransferFrom0 is a paid mutator transaction binding the contract method 0xfe8ec1a7.
//
// Solidity: function permitWitnessTransferFrom(((address,uint256)[],uint256,uint256) permit, (address,uint256)[] transferDetails, address owner, bytes32 witness, string witnessTypeString, bytes signature) returns()
func (_IPermit2 *IPermit2Transactor) PermitWitnessTransferFrom0(opts *bind.TransactOpts, permit ISignatureTransferPermitBatchTransferFrom, transferDetails []ISignatureTransferSignatureTransferDetails, owner common.Address, witness [32]byte, witnessTypeString string, signature []byte) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "permitWitnessTransferFrom0", permit, transferDetails, owner, witness, witnessTypeString, signature)
}

// PermitWitnessTransferFrom0 is a paid mutator transaction binding the contract method 0xfe8ec1a7.
//
// Solidity: function permitWitnessTransferFrom(((address,uint256)[],uint256,uint256) permit, (address,uint256)[] transferDetails, address owner, bytes32 witness, string witnessTypeString, bytes signature) returns()
func (_IPermit2 *IPermit2Session) PermitWitnessTransferFrom0(permit ISignatureTransferPermitBatchTransferFrom, transferDetails []ISignatureTransferSignatureTransferDetails, owner common.Address, witness [32]byte, witnessTypeString string, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.PermitWitnessTransferFrom0(&_IPermit2.TransactOpts, permit, transferDetails, owner, witness, witnessTypeString, signature)
}

// PermitWitnessTransferFrom0 is a paid mutator transaction binding the contract method 0xfe8ec1a7.
//
// Solidity: function permitWitnessTransferFrom(((address,uint256)[],uint256,uint256) permit, (address,uint256)[] transferDetails, address owner, bytes32 witness, string witnessTypeString, bytes signature) returns()
func (_IPermit2 *IPermit2TransactorSession) PermitWitnessTransferFrom0(permit ISignatureTransferPermitBatchTransferFrom, transferDetails []ISignatureTransferSignatureTransferDetails, owner common.Address, witness [32]byte, witnessTypeString string, signature []byte) (*types.Transaction, error) {
	return _IPermit2.Contract.PermitWitnessTransferFrom0(&_IPermit2.TransactOpts, permit, transferDetails, owner, witness, witnessTypeString, signature)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0d58b1db.
//
// Solidity: function transferFrom((address,address,uint160,address)[] transferDetails) returns()
func (_IPermit2 *IPermit2Transactor) TransferFrom(opts *bind.TransactOpts, transferDetails []IAllowanceTransferAllowanceTransferDetails) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "transferFrom", transferDetails)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0d58b1db.
//
// Solidity: function transferFrom((address,address,uint160,address)[] transferDetails) returns()
func (_IPermit2 *IPermit2Session) TransferFrom(transferDetails []IAllowanceTransferAllowanceTransferDetails) (*types.Transaction, error) {
	return _IPermit2.Contract.TransferFrom(&_IPermit2.TransactOpts, transferDetails)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x0d58b1db.
//
// Solidity: function transferFrom((address,address,uint160,address)[] transferDetails) returns()
func (_IPermit2 *IPermit2TransactorSession) TransferFrom(transferDetails []IAllowanceTransferAllowanceTransferDetails) (*types.Transaction, error) {
	return _IPermit2.Contract.TransferFrom(&_IPermit2.TransactOpts, transferDetails)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x36c78516.
//
// Solidity: function transferFrom(address from, address to, uint160 amount, address token) returns()
func (_IPermit2 *IPermit2Transactor) TransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _IPermit2.contract.Transact(opts, "transferFrom0", from, to, amount, token)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x36c78516.
//
// Solidity: function transferFrom(address from, address to, uint160 amount, address token) returns()
func (_IPermit2 *IPermit2Session) TransferFrom0(from common.Address, to common.Address, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _IPermit2.Contract.TransferFrom0(&_IPermit2.TransactOpts, from, to, amount, token)
}

// TransferFrom0 is a paid mutator transaction binding the contract method 0x36c78516.
//
// Solidity: function transferFrom(address from, address to, uint160 amount, address token) returns()
func (_IPermit2 *IPermit2TransactorSession) TransferFrom0(from common.Address, to common.Address, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _IPermit2.Contract.TransferFrom0(&_IPermit2.TransactOpts, from, to, amount, token)
}

// IPermit2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IPermit2 contract.
type IPermit2ApprovalIterator struct {
	Event *IPermit2Approval // Event containing the contract specifics and raw log

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
func (it *IPermit2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPermit2Approval)
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
		it.Event = new(IPermit2Approval)
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
func (it *IPermit2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPermit2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPermit2Approval represents a Approval event raised by the IPermit2 contract.
type IPermit2Approval struct {
	Owner      common.Address
	Token      common.Address
	Spender    common.Address
	Amount     *big.Int
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xda9fa7c1b00402c17d0161b249b1ab8bbec047c5a52207b9c112deffd817036b.
//
// Solidity: event Approval(address indexed owner, address indexed token, address indexed spender, uint160 amount, uint48 expiration)
func (_IPermit2 *IPermit2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, token []common.Address, spender []common.Address) (*IPermit2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPermit2.contract.FilterLogs(opts, "Approval", ownerRule, tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IPermit2ApprovalIterator{contract: _IPermit2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xda9fa7c1b00402c17d0161b249b1ab8bbec047c5a52207b9c112deffd817036b.
//
// Solidity: event Approval(address indexed owner, address indexed token, address indexed spender, uint160 amount, uint48 expiration)
func (_IPermit2 *IPermit2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IPermit2Approval, owner []common.Address, token []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPermit2.contract.WatchLogs(opts, "Approval", ownerRule, tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPermit2Approval)
				if err := _IPermit2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xda9fa7c1b00402c17d0161b249b1ab8bbec047c5a52207b9c112deffd817036b.
//
// Solidity: event Approval(address indexed owner, address indexed token, address indexed spender, uint160 amount, uint48 expiration)
func (_IPermit2 *IPermit2Filterer) ParseApproval(log types.Log) (*IPermit2Approval, error) {
	event := new(IPermit2Approval)
	if err := _IPermit2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPermit2LockdownIterator is returned from FilterLockdown and is used to iterate over the raw logs and unpacked data for Lockdown events raised by the IPermit2 contract.
type IPermit2LockdownIterator struct {
	Event *IPermit2Lockdown // Event containing the contract specifics and raw log

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
func (it *IPermit2LockdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPermit2Lockdown)
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
		it.Event = new(IPermit2Lockdown)
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
func (it *IPermit2LockdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPermit2LockdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPermit2Lockdown represents a Lockdown event raised by the IPermit2 contract.
type IPermit2Lockdown struct {
	Owner   common.Address
	Token   common.Address
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockdown is a free log retrieval operation binding the contract event 0x89b1add15eff56b3dfe299ad94e01f2b52fbcb80ae1a3baea6ae8c04cb2b98a4.
//
// Solidity: event Lockdown(address indexed owner, address token, address spender)
func (_IPermit2 *IPermit2Filterer) FilterLockdown(opts *bind.FilterOpts, owner []common.Address) (*IPermit2LockdownIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IPermit2.contract.FilterLogs(opts, "Lockdown", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IPermit2LockdownIterator{contract: _IPermit2.contract, event: "Lockdown", logs: logs, sub: sub}, nil
}

// WatchLockdown is a free log subscription operation binding the contract event 0x89b1add15eff56b3dfe299ad94e01f2b52fbcb80ae1a3baea6ae8c04cb2b98a4.
//
// Solidity: event Lockdown(address indexed owner, address token, address spender)
func (_IPermit2 *IPermit2Filterer) WatchLockdown(opts *bind.WatchOpts, sink chan<- *IPermit2Lockdown, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IPermit2.contract.WatchLogs(opts, "Lockdown", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPermit2Lockdown)
				if err := _IPermit2.contract.UnpackLog(event, "Lockdown", log); err != nil {
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

// ParseLockdown is a log parse operation binding the contract event 0x89b1add15eff56b3dfe299ad94e01f2b52fbcb80ae1a3baea6ae8c04cb2b98a4.
//
// Solidity: event Lockdown(address indexed owner, address token, address spender)
func (_IPermit2 *IPermit2Filterer) ParseLockdown(log types.Log) (*IPermit2Lockdown, error) {
	event := new(IPermit2Lockdown)
	if err := _IPermit2.contract.UnpackLog(event, "Lockdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPermit2NonceInvalidationIterator is returned from FilterNonceInvalidation and is used to iterate over the raw logs and unpacked data for NonceInvalidation events raised by the IPermit2 contract.
type IPermit2NonceInvalidationIterator struct {
	Event *IPermit2NonceInvalidation // Event containing the contract specifics and raw log

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
func (it *IPermit2NonceInvalidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPermit2NonceInvalidation)
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
		it.Event = new(IPermit2NonceInvalidation)
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
func (it *IPermit2NonceInvalidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPermit2NonceInvalidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPermit2NonceInvalidation represents a NonceInvalidation event raised by the IPermit2 contract.
type IPermit2NonceInvalidation struct {
	Owner    common.Address
	Token    common.Address
	Spender  common.Address
	NewNonce *big.Int
	OldNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceInvalidation is a free log retrieval operation binding the contract event 0x55eb90d810e1700b35a8e7e25395ff7f2b2259abd7415ca2284dfb1c246418f3.
//
// Solidity: event NonceInvalidation(address indexed owner, address indexed token, address indexed spender, uint48 newNonce, uint48 oldNonce)
func (_IPermit2 *IPermit2Filterer) FilterNonceInvalidation(opts *bind.FilterOpts, owner []common.Address, token []common.Address, spender []common.Address) (*IPermit2NonceInvalidationIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPermit2.contract.FilterLogs(opts, "NonceInvalidation", ownerRule, tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IPermit2NonceInvalidationIterator{contract: _IPermit2.contract, event: "NonceInvalidation", logs: logs, sub: sub}, nil
}

// WatchNonceInvalidation is a free log subscription operation binding the contract event 0x55eb90d810e1700b35a8e7e25395ff7f2b2259abd7415ca2284dfb1c246418f3.
//
// Solidity: event NonceInvalidation(address indexed owner, address indexed token, address indexed spender, uint48 newNonce, uint48 oldNonce)
func (_IPermit2 *IPermit2Filterer) WatchNonceInvalidation(opts *bind.WatchOpts, sink chan<- *IPermit2NonceInvalidation, owner []common.Address, token []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPermit2.contract.WatchLogs(opts, "NonceInvalidation", ownerRule, tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPermit2NonceInvalidation)
				if err := _IPermit2.contract.UnpackLog(event, "NonceInvalidation", log); err != nil {
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

// ParseNonceInvalidation is a log parse operation binding the contract event 0x55eb90d810e1700b35a8e7e25395ff7f2b2259abd7415ca2284dfb1c246418f3.
//
// Solidity: event NonceInvalidation(address indexed owner, address indexed token, address indexed spender, uint48 newNonce, uint48 oldNonce)
func (_IPermit2 *IPermit2Filterer) ParseNonceInvalidation(log types.Log) (*IPermit2NonceInvalidation, error) {
	event := new(IPermit2NonceInvalidation)
	if err := _IPermit2.contract.UnpackLog(event, "NonceInvalidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPermit2PermitIterator is returned from FilterPermit and is used to iterate over the raw logs and unpacked data for Permit events raised by the IPermit2 contract.
type IPermit2PermitIterator struct {
	Event *IPermit2Permit // Event containing the contract specifics and raw log

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
func (it *IPermit2PermitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPermit2Permit)
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
		it.Event = new(IPermit2Permit)
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
func (it *IPermit2PermitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPermit2PermitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPermit2Permit represents a Permit event raised by the IPermit2 contract.
type IPermit2Permit struct {
	Owner      common.Address
	Token      common.Address
	Spender    common.Address
	Amount     *big.Int
	Expiration *big.Int
	Nonce      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermit is a free log retrieval operation binding the contract event 0xc6a377bfc4eb120024a8ac08eef205be16b817020812c73223e81d1bdb9708ec.
//
// Solidity: event Permit(address indexed owner, address indexed token, address indexed spender, uint160 amount, uint48 expiration, uint48 nonce)
func (_IPermit2 *IPermit2Filterer) FilterPermit(opts *bind.FilterOpts, owner []common.Address, token []common.Address, spender []common.Address) (*IPermit2PermitIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPermit2.contract.FilterLogs(opts, "Permit", ownerRule, tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IPermit2PermitIterator{contract: _IPermit2.contract, event: "Permit", logs: logs, sub: sub}, nil
}

// WatchPermit is a free log subscription operation binding the contract event 0xc6a377bfc4eb120024a8ac08eef205be16b817020812c73223e81d1bdb9708ec.
//
// Solidity: event Permit(address indexed owner, address indexed token, address indexed spender, uint160 amount, uint48 expiration, uint48 nonce)
func (_IPermit2 *IPermit2Filterer) WatchPermit(opts *bind.WatchOpts, sink chan<- *IPermit2Permit, owner []common.Address, token []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IPermit2.contract.WatchLogs(opts, "Permit", ownerRule, tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPermit2Permit)
				if err := _IPermit2.contract.UnpackLog(event, "Permit", log); err != nil {
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

// ParsePermit is a log parse operation binding the contract event 0xc6a377bfc4eb120024a8ac08eef205be16b817020812c73223e81d1bdb9708ec.
//
// Solidity: event Permit(address indexed owner, address indexed token, address indexed spender, uint160 amount, uint48 expiration, uint48 nonce)
func (_IPermit2 *IPermit2Filterer) ParsePermit(log types.Log) (*IPermit2Permit, error) {
	event := new(IPermit2Permit)
	if err := _IPermit2.contract.UnpackLog(event, "Permit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPermit2UnorderedNonceInvalidationIterator is returned from FilterUnorderedNonceInvalidation and is used to iterate over the raw logs and unpacked data for UnorderedNonceInvalidation events raised by the IPermit2 contract.
type IPermit2UnorderedNonceInvalidationIterator struct {
	Event *IPermit2UnorderedNonceInvalidation // Event containing the contract specifics and raw log

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
func (it *IPermit2UnorderedNonceInvalidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPermit2UnorderedNonceInvalidation)
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
		it.Event = new(IPermit2UnorderedNonceInvalidation)
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
func (it *IPermit2UnorderedNonceInvalidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPermit2UnorderedNonceInvalidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPermit2UnorderedNonceInvalidation represents a UnorderedNonceInvalidation event raised by the IPermit2 contract.
type IPermit2UnorderedNonceInvalidation struct {
	Owner common.Address
	Word  *big.Int
	Mask  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnorderedNonceInvalidation is a free log retrieval operation binding the contract event 0x3704902f963766a4e561bbaab6e6cdc1b1dd12f6e9e99648da8843b3f46b918d.
//
// Solidity: event UnorderedNonceInvalidation(address indexed owner, uint256 word, uint256 mask)
func (_IPermit2 *IPermit2Filterer) FilterUnorderedNonceInvalidation(opts *bind.FilterOpts, owner []common.Address) (*IPermit2UnorderedNonceInvalidationIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IPermit2.contract.FilterLogs(opts, "UnorderedNonceInvalidation", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IPermit2UnorderedNonceInvalidationIterator{contract: _IPermit2.contract, event: "UnorderedNonceInvalidation", logs: logs, sub: sub}, nil
}

// WatchUnorderedNonceInvalidation is a free log subscription operation binding the contract event 0x3704902f963766a4e561bbaab6e6cdc1b1dd12f6e9e99648da8843b3f46b918d.
//
// Solidity: event UnorderedNonceInvalidation(address indexed owner, uint256 word, uint256 mask)
func (_IPermit2 *IPermit2Filterer) WatchUnorderedNonceInvalidation(opts *bind.WatchOpts, sink chan<- *IPermit2UnorderedNonceInvalidation, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IPermit2.contract.WatchLogs(opts, "UnorderedNonceInvalidation", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPermit2UnorderedNonceInvalidation)
				if err := _IPermit2.contract.UnpackLog(event, "UnorderedNonceInvalidation", log); err != nil {
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

// ParseUnorderedNonceInvalidation is a log parse operation binding the contract event 0x3704902f963766a4e561bbaab6e6cdc1b1dd12f6e9e99648da8843b3f46b918d.
//
// Solidity: event UnorderedNonceInvalidation(address indexed owner, uint256 word, uint256 mask)
func (_IPermit2 *IPermit2Filterer) ParseUnorderedNonceInvalidation(log types.Log) (*IPermit2UnorderedNonceInvalidation, error) {
	event := new(IPermit2UnorderedNonceInvalidation)
	if err := _IPermit2.contract.UnpackLog(event, "UnorderedNonceInvalidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
