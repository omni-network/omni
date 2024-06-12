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

// OmniBridgeNativeMetaData contains all meta data concerning the OmniBridgeNative contract.
var OmniBridgeNativeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"XCALL_WITHDRAW_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimable\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1Bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1BridgeBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setup\",\"inputs\":[{\"name\":\"l1ChainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1Bridge_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalL1Supply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bridge\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"claimant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50611009806100206000396000f3fe6080604052600436106100e85760003560e01c80638da5cb5b1161008a578063c3de453d11610059578063c3de453d14610287578063d9caed121461029a578063f2fde38b146102ba578063f35ea557146102da57600080fd5b80638da5cb5b146102135780638fdcb4c914610231578063969b53da146102485780639c5451b01461026857600080fd5b806339acf9f1116100c657806339acf9f1146101725780633abfe55f146101b1578063402914f5146101d1578063715018a6146101fe57600080fd5b806312622e5b146100ed5780631e83409a1461012c57806323b051d91461014e575b600080fd5b3480156100f957600080fd5b5060655461010e9067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b34801561013857600080fd5b5061014c610147366004610d52565b6102fa565b005b34801561015a57600080fd5b5061016460665481565b604051908152602001610123565b34801561017e57600080fd5b5060655461019990600160401b90046001600160a01b031681565b6040516001600160a01b039091168152602001610123565b3480156101bd57600080fd5b506101646101cc366004610d76565b6105ee565b3480156101dd57600080fd5b506101646101ec366004610d52565b60686020526000908152604090205481565b34801561020a57600080fd5b5061014c6106c0565b34801561021f57600080fd5b506033546001600160a01b0316610199565b34801561023d57600080fd5b5061010e620124f881565b34801561025457600080fd5b50606754610199906001600160a01b031681565b34801561027457600080fd5b506101646a52b7d2dcc80cd2e400000081565b61014c610295366004610d76565b6106d4565b3480156102a657600080fd5b5061014c6102b5366004610da2565b6106e2565b3480156102c657600080fd5b5061014c6102d5366004610d52565b610944565b3480156102e657600080fd5b5061014c6102f5366004610df9565b6109bd565b60655460408051631799380760e11b81528151600093600160401b90046001600160a01b031692632f32700e92600480820193918290030181865afa158015610347573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036b9190610e44565b606554909150600160401b90046001600160a01b031633146103cc5760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b60448201526064015b60405180910390fd5b606554815167ffffffffffffffff9081169116146104215760405162461bcd60e51b81526020600482015260126024820152714f6d6e694272696467653a206e6f74204c3160701b60448201526064016103c3565b6001600160a01b0382166104775760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e694272696467653a206e6f20636c61696d20746f207a65726f0000000060448201526064016103c3565b6020808201516001600160a01b038116600090815260689092526040909120546104e35760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e694272696467653a206e6f7468696e6720746f20636c61696d0000000060448201526064016103c3565b6001600160a01b038181166000908152606860205260408082208054908390559051909286169083908381818185875af1925050503d8060008114610544576040519150601f19603f3d011682016040523d82523d6000602084013e610549565b606091505b505090508061059a5760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e694272696467653a207472616e73666572206661696c6564000000000060448201526064016103c3565b846001600160a01b0316836001600160a01b03167ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683846040516105df91815260200190565b60405180910390a35050505050565b606554604080516001600160a01b038581166024830152604480830186905283518084039091018152606490920183526020820180516001600160e01b031663f3fef3a360e01b1790529151632376548f60e21b8152600093600160401b810490931692638dd9523c926106769267ffffffffffffffff90921691620124f890600401610ef7565b602060405180830381865afa158015610693573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b79190610f2e565b90505b92915050565b6106c8610a15565b6106d26000610a6f565b565b6106de8282610ac1565b5050565b60655460408051631799380760e11b81528151600093600160401b90046001600160a01b031692632f32700e92600480820193918290030181865afa15801561072f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107539190610e44565b606554909150600160401b90046001600160a01b031633146107af5760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b60448201526064016103c3565b60675460208201516001600160a01b0390811691161461080a5760405162461bcd60e51b81526020600482015260166024820152754f6d6e694272696467653a206e6f742062726964676560501b60448201526064016103c3565b606554815167ffffffffffffffff90811691161461085f5760405162461bcd60e51b81526020600482015260126024820152714f6d6e694272696467653a206e6f74204c3160701b60448201526064016103c3565b81606660008282546108719190610f5d565b90915550506040516000906001600160a01b0385169084908381818185875af1925050503d80600081146108c1576040519150601f19603f3d011682016040523d82523d6000602084013e6108c6565b606091505b50509050806108fd576001600160a01b038516600090815260686020526040812080548592906108f7908490610f5d565b90915550505b6040805184815282151560208201526001600160a01b0380871692908816917f2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f6191016105df565b61094c610a15565b6001600160a01b0381166109b15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103c3565b6109ba81610a6f565b50565b6109c5610a15565b6065805467ffffffffffffffff949094166001600160e01b031990941693909317600160401b6001600160a01b039384160217909255606780546001600160a01b03191692909116919091179055565b6033546001600160a01b031633146106d25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103c3565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038216610b175760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e694272696467653a206e6f2062726964676520746f207a65726f00000060448201526064016103c3565b60008111610b675760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694272696467653a20616d6f756e74206d757374206265203e2030000060448201526064016103c3565b606654811115610bb95760405162461bcd60e51b815260206004820152601860248201527f4f6d6e694272696467653a206e6f206c6971756964697479000000000000000060448201526064016103c3565b6000610bc583836105ee565b9050610bd18183610f5d565b3414610c1f5760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e694272696467653a20696e636f72726563742066756e6473000000000060448201526064016103c3565b8160666000828254610c319190610f70565b9091555050606554606754604080516001600160a01b038781166024830152604480830188905283518084039091018152606490920183526020820180516001600160e01b031663f3fef3a360e01b179052915163c21dda4f60e01b8152600160401b850483169463c21dda4f948794610cc69467ffffffffffffffff909316936004939190921691620124f8908401610f83565b6000604051808303818588803b158015610cdf57600080fd5b505af1158015610cf3573d6000803e3d6000fd5b50506040518581526001600160a01b03871693503392507f59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422915060200160405180910390a3505050565b6001600160a01b03811681146109ba57600080fd5b600060208284031215610d6457600080fd5b8135610d6f81610d3d565b9392505050565b60008060408385031215610d8957600080fd5b8235610d9481610d3d565b946020939093013593505050565b600080600060608486031215610db757600080fd5b8335610dc281610d3d565b92506020840135610dd281610d3d565b929592945050506040919091013590565b67ffffffffffffffff811681146109ba57600080fd5b600080600060608486031215610e0e57600080fd5b8335610e1981610de3565b92506020840135610e2981610d3d565b91506040840135610e3981610d3d565b809150509250925092565b600060408284031215610e5657600080fd5b6040516040810181811067ffffffffffffffff82111715610e8757634e487b7160e01b600052604160045260246000fd5b6040528251610e9581610de3565b81526020830151610ea581610d3d565b60208201529392505050565b6000815180845260005b81811015610ed757602081850181015186830182015201610ebb565b506000602082860101526020601f19601f83011685010191505092915050565b600067ffffffffffffffff808616835260606020840152610f1b6060840186610eb1565b9150808416604084015250949350505050565b600060208284031215610f4057600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b808201808211156106ba576106ba610f47565b818103818111156106ba576106ba610f47565b600067ffffffffffffffff808816835260ff8716602084015260018060a01b038616604084015260a06060840152610fbe60a0840186610eb1565b9150808416608084015250969550505050505056fea26469706673582212208387388a7e9923d5a88231afb39cf25012876b5e5a9cd6c50524f387447172e764736f6c63430008180033",
}

// OmniBridgeNativeABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniBridgeNativeMetaData.ABI instead.
var OmniBridgeNativeABI = OmniBridgeNativeMetaData.ABI

// OmniBridgeNativeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniBridgeNativeMetaData.Bin instead.
var OmniBridgeNativeBin = OmniBridgeNativeMetaData.Bin

// DeployOmniBridgeNative deploys a new Ethereum contract, binding an instance of OmniBridgeNative to it.
func DeployOmniBridgeNative(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniBridgeNative, error) {
	parsed, err := OmniBridgeNativeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniBridgeNativeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniBridgeNative{OmniBridgeNativeCaller: OmniBridgeNativeCaller{contract: contract}, OmniBridgeNativeTransactor: OmniBridgeNativeTransactor{contract: contract}, OmniBridgeNativeFilterer: OmniBridgeNativeFilterer{contract: contract}}, nil
}

// OmniBridgeNative is an auto generated Go binding around an Ethereum contract.
type OmniBridgeNative struct {
	OmniBridgeNativeCaller     // Read-only binding to the contract
	OmniBridgeNativeTransactor // Write-only binding to the contract
	OmniBridgeNativeFilterer   // Log filterer for contract events
}

// OmniBridgeNativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniBridgeNativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeNativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniBridgeNativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeNativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniBridgeNativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniBridgeNativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniBridgeNativeSession struct {
	Contract     *OmniBridgeNative // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniBridgeNativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniBridgeNativeCallerSession struct {
	Contract *OmniBridgeNativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OmniBridgeNativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniBridgeNativeTransactorSession struct {
	Contract     *OmniBridgeNativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OmniBridgeNativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniBridgeNativeRaw struct {
	Contract *OmniBridgeNative // Generic contract binding to access the raw methods on
}

// OmniBridgeNativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniBridgeNativeCallerRaw struct {
	Contract *OmniBridgeNativeCaller // Generic read-only contract binding to access the raw methods on
}

// OmniBridgeNativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniBridgeNativeTransactorRaw struct {
	Contract *OmniBridgeNativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniBridgeNative creates a new instance of OmniBridgeNative, bound to a specific deployed contract.
func NewOmniBridgeNative(address common.Address, backend bind.ContractBackend) (*OmniBridgeNative, error) {
	contract, err := bindOmniBridgeNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNative{OmniBridgeNativeCaller: OmniBridgeNativeCaller{contract: contract}, OmniBridgeNativeTransactor: OmniBridgeNativeTransactor{contract: contract}, OmniBridgeNativeFilterer: OmniBridgeNativeFilterer{contract: contract}}, nil
}

// NewOmniBridgeNativeCaller creates a new read-only instance of OmniBridgeNative, bound to a specific deployed contract.
func NewOmniBridgeNativeCaller(address common.Address, caller bind.ContractCaller) (*OmniBridgeNativeCaller, error) {
	contract, err := bindOmniBridgeNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeCaller{contract: contract}, nil
}

// NewOmniBridgeNativeTransactor creates a new write-only instance of OmniBridgeNative, bound to a specific deployed contract.
func NewOmniBridgeNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniBridgeNativeTransactor, error) {
	contract, err := bindOmniBridgeNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeTransactor{contract: contract}, nil
}

// NewOmniBridgeNativeFilterer creates a new log filterer instance of OmniBridgeNative, bound to a specific deployed contract.
func NewOmniBridgeNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniBridgeNativeFilterer, error) {
	contract, err := bindOmniBridgeNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeFilterer{contract: contract}, nil
}

// bindOmniBridgeNative binds a generic wrapper to an already deployed contract.
func bindOmniBridgeNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniBridgeNativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniBridgeNative *OmniBridgeNativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniBridgeNative.Contract.OmniBridgeNativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniBridgeNative *OmniBridgeNativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.OmniBridgeNativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniBridgeNative *OmniBridgeNativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.OmniBridgeNativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniBridgeNative *OmniBridgeNativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniBridgeNative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniBridgeNative *OmniBridgeNativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniBridgeNative *OmniBridgeNativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.contract.Transact(opts, method, params...)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeNative *OmniBridgeNativeCaller) XCALLWITHDRAWGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "XCALL_WITHDRAW_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeNative *OmniBridgeNativeSession) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _OmniBridgeNative.Contract.XCALLWITHDRAWGASLIMIT(&_OmniBridgeNative.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _OmniBridgeNative.Contract.XCALLWITHDRAWGASLIMIT(&_OmniBridgeNative.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeCaller) BridgeFee(opts *bind.CallOpts, to common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "bridgeFee", to, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeSession) BridgeFee(to common.Address, amount *big.Int) (*big.Int, error) {
	return _OmniBridgeNative.Contract.BridgeFee(&_OmniBridgeNative.CallOpts, to, amount)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) BridgeFee(to common.Address, amount *big.Int) (*big.Int, error) {
	return _OmniBridgeNative.Contract.BridgeFee(&_OmniBridgeNative.CallOpts, to, amount)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeCaller) Claimable(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "claimable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _OmniBridgeNative.Contract.Claimable(&_OmniBridgeNative.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _OmniBridgeNative.Contract.Claimable(&_OmniBridgeNative.CallOpts, arg0)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeCaller) L1Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "l1Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeSession) L1Bridge() (common.Address, error) {
	return _OmniBridgeNative.Contract.L1Bridge(&_OmniBridgeNative.CallOpts)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) L1Bridge() (common.Address, error) {
	return _OmniBridgeNative.Contract.L1Bridge(&_OmniBridgeNative.CallOpts)
}

// L1BridgeBalance is a free data retrieval call binding the contract method 0x23b051d9.
//
// Solidity: function l1BridgeBalance() view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeCaller) L1BridgeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "l1BridgeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1BridgeBalance is a free data retrieval call binding the contract method 0x23b051d9.
//
// Solidity: function l1BridgeBalance() view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeSession) L1BridgeBalance() (*big.Int, error) {
	return _OmniBridgeNative.Contract.L1BridgeBalance(&_OmniBridgeNative.CallOpts)
}

// L1BridgeBalance is a free data retrieval call binding the contract method 0x23b051d9.
//
// Solidity: function l1BridgeBalance() view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) L1BridgeBalance() (*big.Int, error) {
	return _OmniBridgeNative.Contract.L1BridgeBalance(&_OmniBridgeNative.CallOpts)
}

// L1ChainId is a free data retrieval call binding the contract method 0x12622e5b.
//
// Solidity: function l1ChainId() view returns(uint64)
func (_OmniBridgeNative *OmniBridgeNativeCaller) L1ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "l1ChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// L1ChainId is a free data retrieval call binding the contract method 0x12622e5b.
//
// Solidity: function l1ChainId() view returns(uint64)
func (_OmniBridgeNative *OmniBridgeNativeSession) L1ChainId() (uint64, error) {
	return _OmniBridgeNative.Contract.L1ChainId(&_OmniBridgeNative.CallOpts)
}

// L1ChainId is a free data retrieval call binding the contract method 0x12622e5b.
//
// Solidity: function l1ChainId() view returns(uint64)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) L1ChainId() (uint64, error) {
	return _OmniBridgeNative.Contract.L1ChainId(&_OmniBridgeNative.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeSession) Omni() (common.Address, error) {
	return _OmniBridgeNative.Contract.Omni(&_OmniBridgeNative.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) Omni() (common.Address, error) {
	return _OmniBridgeNative.Contract.Omni(&_OmniBridgeNative.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeSession) Owner() (common.Address, error) {
	return _OmniBridgeNative.Contract.Owner(&_OmniBridgeNative.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) Owner() (common.Address, error) {
	return _OmniBridgeNative.Contract.Owner(&_OmniBridgeNative.CallOpts)
}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeCaller) TotalL1Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniBridgeNative.contract.Call(opts, &out, "totalL1Supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeSession) TotalL1Supply() (*big.Int, error) {
	return _OmniBridgeNative.Contract.TotalL1Supply(&_OmniBridgeNative.CallOpts)
}

// TotalL1Supply is a free data retrieval call binding the contract method 0x9c5451b0.
//
// Solidity: function totalL1Supply() view returns(uint256)
func (_OmniBridgeNative *OmniBridgeNativeCallerSession) TotalL1Supply() (*big.Int, error) {
	return _OmniBridgeNative.Contract.TotalL1Supply(&_OmniBridgeNative.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactor) Bridge(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeNative.contract.Transact(opts, "bridge", to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeNative *OmniBridgeNativeSession) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Bridge(&_OmniBridgeNative.TransactOpts, to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactorSession) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Bridge(&_OmniBridgeNative.TransactOpts, to, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address to) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactor) Claim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.contract.Transact(opts, "claim", to)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address to) returns()
func (_OmniBridgeNative *OmniBridgeNativeSession) Claim(to common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Claim(&_OmniBridgeNative.TransactOpts, to)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address to) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactorSession) Claim(to common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Claim(&_OmniBridgeNative.TransactOpts, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniBridgeNative.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeNative *OmniBridgeNativeSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.RenounceOwnership(&_OmniBridgeNative.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.RenounceOwnership(&_OmniBridgeNative.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0xf35ea557.
//
// Solidity: function setup(uint64 l1ChainId_, address omni_, address l1Bridge_) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactor) Setup(opts *bind.TransactOpts, l1ChainId_ uint64, omni_ common.Address, l1Bridge_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.contract.Transact(opts, "setup", l1ChainId_, omni_, l1Bridge_)
}

// Setup is a paid mutator transaction binding the contract method 0xf35ea557.
//
// Solidity: function setup(uint64 l1ChainId_, address omni_, address l1Bridge_) returns()
func (_OmniBridgeNative *OmniBridgeNativeSession) Setup(l1ChainId_ uint64, omni_ common.Address, l1Bridge_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Setup(&_OmniBridgeNative.TransactOpts, l1ChainId_, omni_, l1Bridge_)
}

// Setup is a paid mutator transaction binding the contract method 0xf35ea557.
//
// Solidity: function setup(uint64 l1ChainId_, address omni_, address l1Bridge_) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactorSession) Setup(l1ChainId_ uint64, omni_ common.Address, l1Bridge_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Setup(&_OmniBridgeNative.TransactOpts, l1ChainId_, omni_, l1Bridge_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeNative *OmniBridgeNativeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.TransferOwnership(&_OmniBridgeNative.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.TransferOwnership(&_OmniBridgeNative.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address payor, address to, uint256 amount) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactor) Withdraw(opts *bind.TransactOpts, payor common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeNative.contract.Transact(opts, "withdraw", payor, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address payor, address to, uint256 amount) returns()
func (_OmniBridgeNative *OmniBridgeNativeSession) Withdraw(payor common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Withdraw(&_OmniBridgeNative.TransactOpts, payor, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address payor, address to, uint256 amount) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactorSession) Withdraw(payor common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Withdraw(&_OmniBridgeNative.TransactOpts, payor, to, amount)
}

// OmniBridgeNativeBridgeIterator is returned from FilterBridge and is used to iterate over the raw logs and unpacked data for Bridge events raised by the OmniBridgeNative contract.
type OmniBridgeNativeBridgeIterator struct {
	Event *OmniBridgeNativeBridge // Event containing the contract specifics and raw log

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
func (it *OmniBridgeNativeBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeNativeBridge)
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
		it.Event = new(OmniBridgeNativeBridge)
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
func (it *OmniBridgeNativeBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeNativeBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeNativeBridge represents a Bridge event raised by the OmniBridgeNative contract.
type OmniBridgeNativeBridge struct {
	Payor  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridge is a free log retrieval operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) FilterBridge(opts *bind.FilterOpts, payor []common.Address, to []common.Address) (*OmniBridgeNativeBridgeIterator, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeNative.contract.FilterLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeBridgeIterator{contract: _OmniBridgeNative.contract, event: "Bridge", logs: logs, sub: sub}, nil
}

// WatchBridge is a free log subscription operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) WatchBridge(opts *bind.WatchOpts, sink chan<- *OmniBridgeNativeBridge, payor []common.Address, to []common.Address) (event.Subscription, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeNative.contract.WatchLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeNativeBridge)
				if err := _OmniBridgeNative.contract.UnpackLog(event, "Bridge", log); err != nil {
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

// ParseBridge is a log parse operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) ParseBridge(log types.Log) (*OmniBridgeNativeBridge, error) {
	event := new(OmniBridgeNativeBridge)
	if err := _OmniBridgeNative.contract.UnpackLog(event, "Bridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeNativeClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the OmniBridgeNative contract.
type OmniBridgeNativeClaimedIterator struct {
	Event *OmniBridgeNativeClaimed // Event containing the contract specifics and raw log

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
func (it *OmniBridgeNativeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeNativeClaimed)
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
		it.Event = new(OmniBridgeNativeClaimed)
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
func (it *OmniBridgeNativeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeNativeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeNativeClaimed represents a Claimed event raised by the OmniBridgeNative contract.
type OmniBridgeNativeClaimed struct {
	Claimant common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimant, address indexed to, uint256 amount)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) FilterClaimed(opts *bind.FilterOpts, claimant []common.Address, to []common.Address) (*OmniBridgeNativeClaimedIterator, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeNative.contract.FilterLogs(opts, "Claimed", claimantRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeClaimedIterator{contract: _OmniBridgeNative.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimant, address indexed to, uint256 amount)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *OmniBridgeNativeClaimed, claimant []common.Address, to []common.Address) (event.Subscription, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeNative.contract.WatchLogs(opts, "Claimed", claimantRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeNativeClaimed)
				if err := _OmniBridgeNative.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimant, address indexed to, uint256 amount)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) ParseClaimed(log types.Log) (*OmniBridgeNativeClaimed, error) {
	event := new(OmniBridgeNativeClaimed)
	if err := _OmniBridgeNative.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeNativeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniBridgeNative contract.
type OmniBridgeNativeInitializedIterator struct {
	Event *OmniBridgeNativeInitialized // Event containing the contract specifics and raw log

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
func (it *OmniBridgeNativeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeNativeInitialized)
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
		it.Event = new(OmniBridgeNativeInitialized)
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
func (it *OmniBridgeNativeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeNativeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeNativeInitialized represents a Initialized event raised by the OmniBridgeNative contract.
type OmniBridgeNativeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniBridgeNativeInitializedIterator, error) {

	logs, sub, err := _OmniBridgeNative.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeInitializedIterator{contract: _OmniBridgeNative.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniBridgeNativeInitialized) (event.Subscription, error) {

	logs, sub, err := _OmniBridgeNative.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeNativeInitialized)
				if err := _OmniBridgeNative.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OmniBridgeNative *OmniBridgeNativeFilterer) ParseInitialized(log types.Log) (*OmniBridgeNativeInitialized, error) {
	event := new(OmniBridgeNativeInitialized)
	if err := _OmniBridgeNative.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeNativeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniBridgeNative contract.
type OmniBridgeNativeOwnershipTransferredIterator struct {
	Event *OmniBridgeNativeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniBridgeNativeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeNativeOwnershipTransferred)
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
		it.Event = new(OmniBridgeNativeOwnershipTransferred)
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
func (it *OmniBridgeNativeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeNativeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeNativeOwnershipTransferred represents a OwnershipTransferred event raised by the OmniBridgeNative contract.
type OmniBridgeNativeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniBridgeNativeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniBridgeNative.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeOwnershipTransferredIterator{contract: _OmniBridgeNative.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniBridgeNativeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniBridgeNative.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeNativeOwnershipTransferred)
				if err := _OmniBridgeNative.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OmniBridgeNative *OmniBridgeNativeFilterer) ParseOwnershipTransferred(log types.Log) (*OmniBridgeNativeOwnershipTransferred, error) {
	event := new(OmniBridgeNativeOwnershipTransferred)
	if err := _OmniBridgeNative.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniBridgeNativeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the OmniBridgeNative contract.
type OmniBridgeNativeWithdrawIterator struct {
	Event *OmniBridgeNativeWithdraw // Event containing the contract specifics and raw log

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
func (it *OmniBridgeNativeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniBridgeNativeWithdraw)
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
		it.Event = new(OmniBridgeNativeWithdraw)
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
func (it *OmniBridgeNativeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniBridgeNativeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniBridgeNativeWithdraw represents a Withdraw event raised by the OmniBridgeNative contract.
type OmniBridgeNativeWithdraw struct {
	Payor   common.Address
	To      common.Address
	Amount  *big.Int
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f61.
//
// Solidity: event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) FilterWithdraw(opts *bind.FilterOpts, payor []common.Address, to []common.Address) (*OmniBridgeNativeWithdrawIterator, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeNative.contract.FilterLogs(opts, "Withdraw", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeWithdrawIterator{contract: _OmniBridgeNative.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f61.
//
// Solidity: event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *OmniBridgeNativeWithdraw, payor []common.Address, to []common.Address) (event.Subscription, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniBridgeNative.contract.WatchLogs(opts, "Withdraw", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniBridgeNativeWithdraw)
				if err := _OmniBridgeNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f61.
//
// Solidity: event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) ParseWithdraw(log types.Log) (*OmniBridgeNativeWithdraw, error) {
	event := new(OmniBridgeNativeWithdraw)
	if err := _OmniBridgeNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
