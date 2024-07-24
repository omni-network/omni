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
	ABI: "[{\"type\":\"function\",\"name\":\"XCALL_WITHDRAW_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimable\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1BridgeBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setup\",\"inputs\":[{\"name\":\"l1ChainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1Bridge_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalL1Supply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bridge\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"claimant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b506111a8806100206000396000f3fe6080604052600436106100f35760003560e01c80638fdcb4c91161008a578063c4d66de811610059578063c4d66de8146102c2578063d9caed12146102e2578063f2fde38b14610302578063f35ea5571461032257600080fd5b80638fdcb4c914610259578063969b53da146102705780639c5451b014610290578063c3de453d146102af57600080fd5b80633abfe55f116100c65780633abfe55f146101ba578063402914f5146101da578063715018a6146102075780638da5cb5b1461021c57600080fd5b806312622e5b146100f85780631e83409a1461013557806323b051d91461015757806339acf9f11461017b575b600080fd5b34801561010457600080fd5b50600054610118906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b34801561014157600080fd5b50610155610150366004610ef5565b610342565b005b34801561016357600080fd5b5061016d60015481565b60405190815260200161012c565b34801561018757600080fd5b506000546101a290600160401b90046001600160a01b031681565b6040516001600160a01b03909116815260200161012c565b3480156101c657600080fd5b5061016d6101d5366004610f19565b61063b565b3480156101e657600080fd5b5061016d6101f5366004610ef5565b60036020526000908152604090205481565b34801561021357600080fd5b5061015561070b565b34801561022857600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166101a2565b34801561026557600080fd5b50610118620124f881565b34801561027c57600080fd5b506002546101a2906001600160a01b031681565b34801561029c57600080fd5b5061016d6a52b7d2dcc80cd2e400000081565b6101556102bd366004610f19565b61071f565b3480156102ce57600080fd5b506101556102dd366004610ef5565b61072d565b3480156102ee57600080fd5b506101556102fd366004610f45565b61083b565b34801561030e57600080fd5b5061015561031d366004610ef5565b610aa2565b34801561032e57600080fd5b5061015561033d366004610f9b565b610ae0565b60008060089054906101000a90046001600160a01b03166001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa158015610395573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b99190610fe6565b600054909150600160401b90046001600160a01b0316331461041a5760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b60448201526064015b60405180910390fd5b60005481516001600160401b0390811691161461046e5760405162461bcd60e51b81526020600482015260126024820152714f6d6e694272696467653a206e6f74204c3160701b6044820152606401610411565b6001600160a01b0382166104c45760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e694272696467653a206e6f20636c61696d20746f207a65726f000000006044820152606401610411565b6020808201516001600160a01b038116600090815260039092526040909120546105305760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e694272696467653a206e6f7468696e6720746f20636c61696d000000006044820152606401610411565b6001600160a01b038181166000908152600360205260408082208054908390559051909286169083908381818185875af1925050503d8060008114610591576040519150601f19603f3d011682016040523d82523d6000602084013e610596565b606091505b50509050806105e75760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e694272696467653a207472616e73666572206661696c656400000000006044820152606401610411565b846001600160a01b0316836001600160a01b03167ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926838460405161062c91815260200190565b60405180910390a35050505050565b60008054604080516001600160a01b038681166024830152604480830187905283518084039091018152606490920183526020820180516001600160e01b031663f3fef3a360e01b1790529151632376548f60e21b8152600160401b840490921692638dd9523c926106c1926001600160401b039092169190620124f890600401611098565b602060405180830381865afa1580156106de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070291906110ce565b90505b92915050565b610713610b37565b61071d6000610b92565b565b6107298282610c03565b5050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156107725750825b90506000826001600160401b0316600114801561078e5750303b155b90508115801561079c575080155b156107ba5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156107e457845460ff60401b1916600160401b1785555b6107ed86610e7e565b831561083357845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b60008060089054906101000a90046001600160a01b03166001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa15801561088e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b29190610fe6565b600054909150600160401b90046001600160a01b0316331461090e5760405162461bcd60e51b815260206004820152601560248201527413db5b9a509c9a5919d94e881b9bdd081e18d85b1b605a1b6044820152606401610411565b60025460208201516001600160a01b039081169116146109695760405162461bcd60e51b81526020600482015260166024820152754f6d6e694272696467653a206e6f742062726964676560501b6044820152606401610411565b60005481516001600160401b039081169116146109bd5760405162461bcd60e51b81526020600482015260126024820152714f6d6e694272696467653a206e6f74204c3160701b6044820152606401610411565b81600160008282546109cf91906110fd565b90915550506040516000906001600160a01b0385169084908381818185875af1925050503d8060008114610a1f576040519150601f19603f3d011682016040523d82523d6000602084013e610a24565b606091505b5050905080610a5b576001600160a01b03851660009081526003602052604081208054859290610a559084906110fd565b90915550505b6040805184815282151560208201526001600160a01b0380871692908816917f2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f61910161062c565b610aaa610b37565b6001600160a01b038116610ad457604051631e4fbdf760e01b815260006004820152602401610411565b610add81610b92565b50565b610ae8610b37565b600080546001600160401b03949094166001600160e01b031990941693909317600160401b6001600160a01b039384160217909255600280546001600160a01b03191692909116919091179055565b33610b697f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461071d5760405163118cdaa760e01b8152336004820152602401610411565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6001600160a01b038216610c595760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e694272696467653a206e6f2062726964676520746f207a65726f0000006044820152606401610411565b60008111610ca95760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694272696467653a20616d6f756e74206d757374206265203e203000006044820152606401610411565b600154811115610cfb5760405162461bcd60e51b815260206004820152601860248201527f4f6d6e694272696467653a206e6f206c697175696469747900000000000000006044820152606401610411565b6000610d07838361063b565b9050610d1381836110fd565b3414610d615760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e694272696467653a20696e636f72726563742066756e647300000000006044820152606401610411565b8160016000828254610d739190611110565b9091555050600054600254604080516001600160a01b038781166024830152604480830188905283518084039091018152606490920183526020820180516001600160e01b031663f3fef3a360e01b179052915163c21dda4f60e01b8152600160401b850483169463c21dda4f948794610e07946001600160401b03909316936004939190921691620124f8908401611123565b6000604051808303818588803b158015610e2057600080fd5b505af1158015610e34573d6000803e3d6000fd5b50506040518581526001600160a01b03871693503392507f59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422915060200160405180910390a3505050565b610e86610e8f565b610add81610ed8565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661071d57604051631afcd79f60e31b815260040160405180910390fd5b610aaa610e8f565b6001600160a01b0381168114610add57600080fd5b600060208284031215610f0757600080fd5b8135610f1281610ee0565b9392505050565b60008060408385031215610f2c57600080fd5b8235610f3781610ee0565b946020939093013593505050565b600080600060608486031215610f5a57600080fd5b8335610f6581610ee0565b92506020840135610f7581610ee0565b929592945050506040919091013590565b6001600160401b0381168114610add57600080fd5b600080600060608486031215610fb057600080fd5b8335610fbb81610f86565b92506020840135610fcb81610ee0565b91506040840135610fdb81610ee0565b809150509250925092565b600060408284031215610ff857600080fd5b604051604081018181106001600160401b038211171561102857634e487b7160e01b600052604160045260246000fd5b604052825161103681610f86565b8152602083015161104681610ee0565b60208201529392505050565b6000815180845260005b818110156110785760208185018101518683018201520161105c565b506000602082860101526020601f19601f83011685010191505092915050565b60006001600160401b038086168352606060208401526110bb6060840186611052565b9150808416604084015250949350505050565b6000602082840312156110e057600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b80820180821115610705576107056110e7565b81810381811115610705576107056110e7565b60006001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a0606084015261115d60a0840186611052565b9150808416608084015250969550505050505056fea2646970667358221220ed48585e84acdc593ffe0e424f2b6ab5d74d835284c961e49c08a8b8d958c4d364736f6c63430008180033",
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.contract.Transact(opts, "initialize", owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_OmniBridgeNative *OmniBridgeNativeSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Initialize(&_OmniBridgeNative.TransactOpts, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_OmniBridgeNative *OmniBridgeNativeTransactorSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _OmniBridgeNative.Contract.Initialize(&_OmniBridgeNative.TransactOpts, owner_)
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
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniBridgeNative *OmniBridgeNativeFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniBridgeNativeInitializedIterator, error) {

	logs, sub, err := _OmniBridgeNative.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniBridgeNativeInitializedIterator{contract: _OmniBridgeNative.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
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
