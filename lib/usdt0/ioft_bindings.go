// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdt0

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

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// OFTFeeDetail is an auto generated low-level Go binding around an user-defined struct.
type OFTFeeDetail struct {
	FeeAmountLD *big.Int
	Description string
}

// OFTLimit is an auto generated low-level Go binding around an user-defined struct.
type OFTLimit struct {
	MinAmountLD *big.Int
	MaxAmountLD *big.Int
}

// OFTReceipt is an auto generated low-level Go binding around an user-defined struct.
type OFTReceipt struct {
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
}

// SendParam is an auto generated low-level Go binding around an user-defined struct.
type SendParam struct {
	DstEid       uint32
	To           [32]byte
	AmountLD     *big.Int
	MinAmountLD  *big.Int
	ExtraOptions []byte
	ComposeMsg   []byte
	OftCmd       []byte
}

// IOFTMetaData contains all meta data concerning the IOFT contract.
var IOFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"approvalRequired\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oftVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteOFT\",\"inputs\":[{\"name\":\"_sendParam\",\"type\":\"tuple\",\"internalType\":\"structSendParam\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"to\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraOptions\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"composeMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"oftCmd\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOFTLimit\",\"components\":[{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"oftFeeDetails\",\"type\":\"tuple[]\",\"internalType\":\"structOFTFeeDetail[]\",\"components\":[{\"name\":\"feeAmountLD\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOFTReceipt\",\"components\":[{\"name\":\"amountSentLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountReceivedLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteSend\",\"inputs\":[{\"name\":\"_sendParam\",\"type\":\"tuple\",\"internalType\":\"structSendParam\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"to\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraOptions\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"composeMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"oftCmd\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_payInLzToken\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMessagingFee\",\"components\":[{\"name\":\"nativeFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lzTokenFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"_sendParam\",\"type\":\"tuple\",\"internalType\":\"structSendParam\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"to\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraOptions\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"composeMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"oftCmd\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_fee\",\"type\":\"tuple\",\"internalType\":\"structMessagingFee\",\"components\":[{\"name\":\"nativeFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lzTokenFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_refundAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMessagingReceipt\",\"components\":[{\"name\":\"guid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structMessagingFee\",\"components\":[{\"name\":\"nativeFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lzTokenFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOFTReceipt\",\"components\":[{\"name\":\"amountSentLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountReceivedLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sharedDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OFTReceived\",\"inputs\":[{\"name\":\"guid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcEid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"toAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountReceivedLD\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OFTSent\",\"inputs\":[{\"name\":\"guid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"dstEid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"fromAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountSentLD\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountReceivedLD\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidLocalDecimals\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlippageExceeded\",\"inputs\":[{\"name\":\"amountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// IOFTABI is the input ABI used to generate the binding from.
// Deprecated: Use IOFTMetaData.ABI instead.
var IOFTABI = IOFTMetaData.ABI

// IOFT is an auto generated Go binding around an Ethereum contract.
type IOFT struct {
	IOFTCaller     // Read-only binding to the contract
	IOFTTransactor // Write-only binding to the contract
	IOFTFilterer   // Log filterer for contract events
}

// IOFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOFTSession struct {
	Contract     *IOFT             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOFTCallerSession struct {
	Contract *IOFTCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IOFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOFTTransactorSession struct {
	Contract     *IOFTTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOFTRaw struct {
	Contract *IOFT // Generic contract binding to access the raw methods on
}

// IOFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOFTCallerRaw struct {
	Contract *IOFTCaller // Generic read-only contract binding to access the raw methods on
}

// IOFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOFTTransactorRaw struct {
	Contract *IOFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOFT creates a new instance of IOFT, bound to a specific deployed contract.
func NewIOFT(address common.Address, backend bind.ContractBackend) (*IOFT, error) {
	contract, err := bindIOFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOFT{IOFTCaller: IOFTCaller{contract: contract}, IOFTTransactor: IOFTTransactor{contract: contract}, IOFTFilterer: IOFTFilterer{contract: contract}}, nil
}

// NewIOFTCaller creates a new read-only instance of IOFT, bound to a specific deployed contract.
func NewIOFTCaller(address common.Address, caller bind.ContractCaller) (*IOFTCaller, error) {
	contract, err := bindIOFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOFTCaller{contract: contract}, nil
}

// NewIOFTTransactor creates a new write-only instance of IOFT, bound to a specific deployed contract.
func NewIOFTTransactor(address common.Address, transactor bind.ContractTransactor) (*IOFTTransactor, error) {
	contract, err := bindIOFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOFTTransactor{contract: contract}, nil
}

// NewIOFTFilterer creates a new log filterer instance of IOFT, bound to a specific deployed contract.
func NewIOFTFilterer(address common.Address, filterer bind.ContractFilterer) (*IOFTFilterer, error) {
	contract, err := bindIOFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOFTFilterer{contract: contract}, nil
}

// bindIOFT binds a generic wrapper to an already deployed contract.
func bindIOFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOFT *IOFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOFT.Contract.IOFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOFT *IOFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOFT.Contract.IOFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOFT *IOFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOFT.Contract.IOFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOFT *IOFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOFT *IOFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOFT *IOFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOFT.Contract.contract.Transact(opts, method, params...)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() view returns(bool)
func (_IOFT *IOFTCaller) ApprovalRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IOFT.contract.Call(opts, &out, "approvalRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() view returns(bool)
func (_IOFT *IOFTSession) ApprovalRequired() (bool, error) {
	return _IOFT.Contract.ApprovalRequired(&_IOFT.CallOpts)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() view returns(bool)
func (_IOFT *IOFTCallerSession) ApprovalRequired() (bool, error) {
	return _IOFT.Contract.ApprovalRequired(&_IOFT.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() view returns(bytes4 interfaceId, uint64 version)
func (_IOFT *IOFTCaller) OftVersion(opts *bind.CallOpts) (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	var out []interface{}
	err := _IOFT.contract.Call(opts, &out, "oftVersion")

	outstruct := new(struct {
		InterfaceId [4]byte
		Version     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InterfaceId = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.Version = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() view returns(bytes4 interfaceId, uint64 version)
func (_IOFT *IOFTSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _IOFT.Contract.OftVersion(&_IOFT.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() view returns(bytes4 interfaceId, uint64 version)
func (_IOFT *IOFTCallerSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _IOFT.Contract.OftVersion(&_IOFT.CallOpts)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256), (int256,string)[] oftFeeDetails, (uint256,uint256))
func (_IOFT *IOFTCaller) QuoteOFT(opts *bind.CallOpts, _sendParam SendParam) (OFTLimit, []OFTFeeDetail, OFTReceipt, error) {
	var out []interface{}
	err := _IOFT.contract.Call(opts, &out, "quoteOFT", _sendParam)

	if err != nil {
		return *new(OFTLimit), *new([]OFTFeeDetail), *new(OFTReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(OFTLimit)).(*OFTLimit)
	out1 := *abi.ConvertType(out[1], new([]OFTFeeDetail)).(*[]OFTFeeDetail)
	out2 := *abi.ConvertType(out[2], new(OFTReceipt)).(*OFTReceipt)

	return out0, out1, out2, err

}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256), (int256,string)[] oftFeeDetails, (uint256,uint256))
func (_IOFT *IOFTSession) QuoteOFT(_sendParam SendParam) (OFTLimit, []OFTFeeDetail, OFTReceipt, error) {
	return _IOFT.Contract.QuoteOFT(&_IOFT.CallOpts, _sendParam)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256), (int256,string)[] oftFeeDetails, (uint256,uint256))
func (_IOFT *IOFTCallerSession) QuoteOFT(_sendParam SendParam) (OFTLimit, []OFTFeeDetail, OFTReceipt, error) {
	return _IOFT.Contract.QuoteOFT(&_IOFT.CallOpts, _sendParam)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256))
func (_IOFT *IOFTCaller) QuoteSend(opts *bind.CallOpts, _sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _IOFT.contract.Call(opts, &out, "quoteSend", _sendParam, _payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256))
func (_IOFT *IOFTSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _IOFT.Contract.QuoteSend(&_IOFT.CallOpts, _sendParam, _payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256))
func (_IOFT *IOFTCallerSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _IOFT.Contract.QuoteSend(&_IOFT.CallOpts, _sendParam, _payInLzToken)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_IOFT *IOFTCaller) SharedDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IOFT.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_IOFT *IOFTSession) SharedDecimals() (uint8, error) {
	return _IOFT.Contract.SharedDecimals(&_IOFT.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_IOFT *IOFTCallerSession) SharedDecimals() (uint8, error) {
	return _IOFT.Contract.SharedDecimals(&_IOFT.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IOFT *IOFTCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOFT.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IOFT *IOFTSession) Token() (common.Address, error) {
	return _IOFT.Contract.Token(&_IOFT.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IOFT *IOFTCallerSession) Token() (common.Address, error) {
	return _IOFT.Contract.Token(&_IOFT.CallOpts)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)), (uint256,uint256))
func (_IOFT *IOFTTransactor) Send(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _IOFT.contract.Transact(opts, "send", _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)), (uint256,uint256))
func (_IOFT *IOFTSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _IOFT.Contract.Send(&_IOFT.TransactOpts, _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)), (uint256,uint256))
func (_IOFT *IOFTTransactorSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _IOFT.Contract.Send(&_IOFT.TransactOpts, _sendParam, _fee, _refundAddress)
}

// IOFTOFTReceivedIterator is returned from FilterOFTReceived and is used to iterate over the raw logs and unpacked data for OFTReceived events raised by the IOFT contract.
type IOFTOFTReceivedIterator struct {
	Event *IOFTOFTReceived // Event containing the contract specifics and raw log

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
func (it *IOFTOFTReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOFTOFTReceived)
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
		it.Event = new(IOFTOFTReceived)
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
func (it *IOFTOFTReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOFTOFTReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOFTOFTReceived represents a OFTReceived event raised by the IOFT contract.
type IOFTOFTReceived struct {
	Guid             [32]byte
	SrcEid           uint32
	ToAddress        common.Address
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTReceived is a free log retrieval operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_IOFT *IOFTFilterer) FilterOFTReceived(opts *bind.FilterOpts, guid [][32]byte, toAddress []common.Address) (*IOFTOFTReceivedIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _IOFT.contract.FilterLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &IOFTOFTReceivedIterator{contract: _IOFT.contract, event: "OFTReceived", logs: logs, sub: sub}, nil
}

// WatchOFTReceived is a free log subscription operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_IOFT *IOFTFilterer) WatchOFTReceived(opts *bind.WatchOpts, sink chan<- *IOFTOFTReceived, guid [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _IOFT.contract.WatchLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOFTOFTReceived)
				if err := _IOFT.contract.UnpackLog(event, "OFTReceived", log); err != nil {
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

// ParseOFTReceived is a log parse operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_IOFT *IOFTFilterer) ParseOFTReceived(log types.Log) (*IOFTOFTReceived, error) {
	event := new(IOFTOFTReceived)
	if err := _IOFT.contract.UnpackLog(event, "OFTReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOFTOFTSentIterator is returned from FilterOFTSent and is used to iterate over the raw logs and unpacked data for OFTSent events raised by the IOFT contract.
type IOFTOFTSentIterator struct {
	Event *IOFTOFTSent // Event containing the contract specifics and raw log

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
func (it *IOFTOFTSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOFTOFTSent)
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
		it.Event = new(IOFTOFTSent)
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
func (it *IOFTOFTSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOFTOFTSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOFTOFTSent represents a OFTSent event raised by the IOFT contract.
type IOFTOFTSent struct {
	Guid             [32]byte
	DstEid           uint32
	FromAddress      common.Address
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTSent is a free log retrieval operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_IOFT *IOFTFilterer) FilterOFTSent(opts *bind.FilterOpts, guid [][32]byte, fromAddress []common.Address) (*IOFTOFTSentIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _IOFT.contract.FilterLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &IOFTOFTSentIterator{contract: _IOFT.contract, event: "OFTSent", logs: logs, sub: sub}, nil
}

// WatchOFTSent is a free log subscription operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_IOFT *IOFTFilterer) WatchOFTSent(opts *bind.WatchOpts, sink chan<- *IOFTOFTSent, guid [][32]byte, fromAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _IOFT.contract.WatchLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOFTOFTSent)
				if err := _IOFT.contract.UnpackLog(event, "OFTSent", log); err != nil {
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

// ParseOFTSent is a log parse operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_IOFT *IOFTFilterer) ParseOFTSent(log types.Log) (*IOFTOFTSent, error) {
	event := new(IOFTOFTSent)
	if err := _IOFT.contract.UnpackLog(event, "OFTSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
