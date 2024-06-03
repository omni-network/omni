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

// XRegistryMetaData contains all meta data concerning the XRegistry contract.
var XRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"XSET_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainIds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"has\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"registrationFee\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"replicas\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPortal\",\"inputs\":[{\"name\":\"_omni\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReplica\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"replica\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6114108061007e6000396000f3fe6080604052600436106100dd5760003560e01c80638926f54f1161007f578063b3449b7711610059578063b3449b771461028b578063f2fde38b146102ab578063fbe4b7c0146102cb578063fd0b64f7146102ed57600080fd5b80638926f54f146101ee5780638da5cb5b1461023f578063a4861b421461025d57600080fd5b80634ff56192116100bb5780634ff5619214610163578063715018a6146101835780637b7c0ddc1461019857806388f9380b146101b857600080fd5b806314a28fcb146100e257806339acf9f1146100f75780634a0f997314610134575b600080fd5b6100f56100f0366004610fc5565b61030d565b005b34801561010357600080fd5b50600254610117906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561014057600080fd5b5061014b620186a081565b6040516001600160401b03909116815260200161012b565b34801561016f57600080fd5b506100f561017e36600461105a565b610540565b34801561018f57600080fd5b506100f561056a565b3480156101a457600080fd5b506100f56101b336600461107c565b61057e565b3480156101c457600080fd5b506101176101d33660046110b1565b6003602052600090815260409020546001600160a01b031681565b3480156101fa57600080fd5b5061022f6102093660046110b1565b6001600160401b03166000908152600360205260409020546001600160a01b0316151590565b604051901515815260200161012b565b34801561024b57600080fd5b506000546001600160a01b0316610117565b34801561026957600080fd5b5061027d610278366004610fc5565b61067e565b60405190815260200161012b565b34801561029757600080fd5b506101176102a63660046110e4565b610696565b3480156102b757600080fd5b506100f56102c636600461105a565b6106ab565b3480156102d757600080fd5b506102e0610724565b60405161012b91906111b7565b3480156102f957600080fd5b5061022f6103083660046110e4565b6107ae565b6001600160401b0384166000908152600360205260409020546001600160a01b03166103805760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a20636861696e206e6f7420737570706f72746564000060448201526064015b60405180910390fd5b61038d84848433856107cf565b3410156103dc5760405162461bcd60e51b815260206004820152601b60248201527f5852656769737472793a20696e73756666696369656e742066656500000000006044820152606401610377565b6104208484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152503392508691506108d29050565b61042d848484338561092e565b61046e83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250339250610a13915050565b156104da57600254604051632023c96760e11b81526001600160401b03861660048201526001600160a01b039091169063404792ce90602401600060405180830381600087803b1580156104c157600080fd5b505af11580156104d5573d6000803e3d6000fd5b505050505b60405133906104ec9085908590611204565b6040519081900381206001600160a01b0384168252906001600160401b038716907fd06596d338531cdc7b5a36893f1bae3902aa239c822f38d70301ca871f855fac9060200160405180910390a450505050565b610548610a69565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b610572610a69565b61057c6000610ac3565b565b610586610a69565b6001600160401b0382166000908152600360205260409020546001600160a01b0316156105f55760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a207265706c69636120616c72656164792073657400006044820152606401610377565b6001600160401b039182166000818152600360208190526040822080546001600160a01b0319166001600160a01b03959095169490941790935560048054600181018255918190527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b908204018054919093166008026101000a93840219169202919091179055565b600061068d85858533866107cf565b95945050505050565b60006106a3848484610b13565b949350505050565b6106b3610a69565b6001600160a01b0381166107185760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610377565b61072181610ac3565b50565b606060048054806020026020016040519081016040528092919081815260200182805480156107a457602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116107615790505b5050505050905090565b6000806107bc858585610b13565b6001600160a01b03161415949350505050565b60008060005b6004548110156108c7576000600482815481106107f4576107f4611214565b6000918252602090912060048204015460039091166008026101000a90046001600160401b0390811691508916810361082d57506108bf565b6000610871828a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610b13915050565b90506001600160a01b0381166108885750506108bf565b6108968a838b8b8b86610b58565b6108a0908561122a565b93506108b0828b8b8b8b8b610b58565b6108ba908561122a565b935050505b6001016107d5565b509695505050505050565b6001600160401b038416600090815260016020526040812082916108f68686610d2e565b815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050565b60005b600454811015610a0b5760006004828154811061095057610950611214565b6000918252602090912060048204015460039091166008026101000a90046001600160401b039081169150871681036109895750610a03565b60006109cd8288888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250610b13915050565b90506001600160a01b0381166109e4575050610a03565b6109f2888389898986610d61565b610a00828989898989610d61565b50505b600101610931565b505050505050565b6000610a556040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b81525073121e240000000000000000000000000000000002610d2e565b610a5f8484610d2e565b1490505b92915050565b6000546001600160a01b0316331461057c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610377565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160401b038316600090815260016020526040812081610b368585610d2e565b81526020810191909152604001600020546001600160a01b0316949350505050565b60025460408051634d4502c960e11b815290516000926001600160a01b031691639a8a05929160048083019260209291908290030181865afa158015610ba2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc6919061124b565b6001600160401b0316876001600160401b031603610be657506000610d24565b6001600160401b0387166000908152600360205260409020546001600160a01b031680610c505760405162461bcd60e51b81526020600482015260186024820152772c2932b3b4b9ba393c9d103ab735b737bbb71031b430b4b760411b6044820152606401610377565b600063d7a18fcd60e01b8888888888604051602401610c73959493929190611268565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092526002549151632376548f60e21b81529092506001600160a01b0390911690638dd9523c90610cde908c908590620186a09060040161130f565b602060405180830381865afa158015610cfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1f9190611345565b925050505b9695505050505050565b60008282604051602001610d4392919061135e565b60405160208183030381529060405280519060200120905092915050565b600260009054906101000a90046001600160a01b03166001600160a01b0316639a8a05926040518163ffffffff1660e01b8152600401602060405180830381865afa158015610db4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd8919061124b565b6001600160401b0316866001600160401b03160315610a0b576001600160401b0386166000908152600360205260409020546001600160a01b031680610e5b5760405162461bcd60e51b81526020600482015260186024820152772c2932b3b4b9ba393c9d103ab735b737bbb71031b430b4b760411b6044820152606401610377565b600063d7a18fcd60e01b8787878787604051602401610e7e959493929190611268565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092526002549151632376548f60e21b81529092506001600160a01b03909116906370e8b56a908290638dd9523c90610ef1908d908790620186a09060040161130f565b602060405180830381865afa158015610f0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f329190611345565b8a8585620186a06040518663ffffffff1660e01b8152600401610f589493929190611395565b6000604051808303818588803b158015610f7157600080fd5b505af1158015610f85573d6000803e3d6000fd5b50505050505050505050505050565b6001600160401b038116811461072157600080fd5b80356001600160a01b0381168114610fc057600080fd5b919050565b60008060008060608587031215610fdb57600080fd5b8435610fe681610f94565b935060208501356001600160401b038082111561100257600080fd5b818701915087601f83011261101657600080fd5b81358181111561102557600080fd5b88602082850101111561103757600080fd5b60208301955080945050505061104f60408601610fa9565b905092959194509250565b60006020828403121561106c57600080fd5b61107582610fa9565b9392505050565b6000806040838503121561108f57600080fd5b823561109a81610f94565b91506110a860208401610fa9565b90509250929050565b6000602082840312156110c357600080fd5b813561107581610f94565b634e487b7160e01b600052604160045260246000fd5b6000806000606084860312156110f957600080fd5b833561110481610f94565b925060208401356001600160401b038082111561112057600080fd5b818601915086601f83011261113457600080fd5b813581811115611146576111466110ce565b604051601f8201601f19908116603f0116810190838211818310171561116e5761116e6110ce565b8160405282815289602084870101111561118757600080fd5b8260208601602083013760006020848301015280965050505050506111ae60408501610fa9565b90509250925092565b6020808252825182820181905260009190848201906040850190845b818110156111f85783516001600160401b0316835292840192918401916001016111d3565b50909695505050505050565b8183823760009101908152919050565b634e487b7160e01b600052603260045260246000fd5b80820180821115610a6357634e487b7160e01b600052601160045260246000fd5b60006020828403121561125d57600080fd5b815161107581610f94565b6001600160401b038616815260806020820152836080820152838560a0830137600060a08583018101919091526001600160a01b039384166040830152919092166060830152601f909201601f1916010192915050565b60005b838110156112da5781810151838201526020016112c2565b50506000910152565b600081518084526112fb8160208601602086016112bf565b601f01601f19169290920160200192915050565b60006001600160401b0380861683526060602084015261133260608401866112e3565b9150808416604084015250949350505050565b60006020828403121561135757600080fd5b5051919050565b600083516113708184602088016112bf565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b60006001600160401b03808716835260018060a01b0386166020840152608060408401526113c660808401866112e3565b91508084166060840152509594505050505056fea264697066735822122033738c9135e91475725d915f4559e6f707965107068b32b9fcadad6ae464e15964736f6c63430008180033",
}

// XRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use XRegistryMetaData.ABI instead.
var XRegistryABI = XRegistryMetaData.ABI

// XRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use XRegistryMetaData.Bin instead.
var XRegistryBin = XRegistryMetaData.Bin

// DeployXRegistry deploys a new Ethereum contract, binding an instance of XRegistry to it.
func DeployXRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XRegistry, error) {
	parsed, err := XRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(XRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XRegistry{XRegistryCaller: XRegistryCaller{contract: contract}, XRegistryTransactor: XRegistryTransactor{contract: contract}, XRegistryFilterer: XRegistryFilterer{contract: contract}}, nil
}

// XRegistry is an auto generated Go binding around an Ethereum contract.
type XRegistry struct {
	XRegistryCaller     // Read-only binding to the contract
	XRegistryTransactor // Write-only binding to the contract
	XRegistryFilterer   // Log filterer for contract events
}

// XRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type XRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XRegistrySession struct {
	Contract     *XRegistry        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XRegistryCallerSession struct {
	Contract *XRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// XRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XRegistryTransactorSession struct {
	Contract     *XRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// XRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type XRegistryRaw struct {
	Contract *XRegistry // Generic contract binding to access the raw methods on
}

// XRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XRegistryCallerRaw struct {
	Contract *XRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// XRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XRegistryTransactorRaw struct {
	Contract *XRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXRegistry creates a new instance of XRegistry, bound to a specific deployed contract.
func NewXRegistry(address common.Address, backend bind.ContractBackend) (*XRegistry, error) {
	contract, err := bindXRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XRegistry{XRegistryCaller: XRegistryCaller{contract: contract}, XRegistryTransactor: XRegistryTransactor{contract: contract}, XRegistryFilterer: XRegistryFilterer{contract: contract}}, nil
}

// NewXRegistryCaller creates a new read-only instance of XRegistry, bound to a specific deployed contract.
func NewXRegistryCaller(address common.Address, caller bind.ContractCaller) (*XRegistryCaller, error) {
	contract, err := bindXRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XRegistryCaller{contract: contract}, nil
}

// NewXRegistryTransactor creates a new write-only instance of XRegistry, bound to a specific deployed contract.
func NewXRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*XRegistryTransactor, error) {
	contract, err := bindXRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XRegistryTransactor{contract: contract}, nil
}

// NewXRegistryFilterer creates a new log filterer instance of XRegistry, bound to a specific deployed contract.
func NewXRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*XRegistryFilterer, error) {
	contract, err := bindXRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XRegistryFilterer{contract: contract}, nil
}

// bindXRegistry binds a generic wrapper to an already deployed contract.
func bindXRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := XRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XRegistry *XRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XRegistry.Contract.XRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XRegistry *XRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XRegistry.Contract.XRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XRegistry *XRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XRegistry.Contract.XRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XRegistry *XRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XRegistry *XRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XRegistry *XRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XRegistry.Contract.contract.Transact(opts, method, params...)
}

// XSETGASLIMIT is a free data retrieval call binding the contract method 0x4a0f9973.
//
// Solidity: function XSET_GAS_LIMIT() view returns(uint64)
func (_XRegistry *XRegistryCaller) XSETGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "XSET_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XSETGASLIMIT is a free data retrieval call binding the contract method 0x4a0f9973.
//
// Solidity: function XSET_GAS_LIMIT() view returns(uint64)
func (_XRegistry *XRegistrySession) XSETGASLIMIT() (uint64, error) {
	return _XRegistry.Contract.XSETGASLIMIT(&_XRegistry.CallOpts)
}

// XSETGASLIMIT is a free data retrieval call binding the contract method 0x4a0f9973.
//
// Solidity: function XSET_GAS_LIMIT() view returns(uint64)
func (_XRegistry *XRegistryCallerSession) XSETGASLIMIT() (uint64, error) {
	return _XRegistry.Contract.XSETGASLIMIT(&_XRegistry.CallOpts)
}

// ChainIds is a free data retrieval call binding the contract method 0xfbe4b7c0.
//
// Solidity: function chainIds() view returns(uint64[])
func (_XRegistry *XRegistryCaller) ChainIds(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "chainIds")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// ChainIds is a free data retrieval call binding the contract method 0xfbe4b7c0.
//
// Solidity: function chainIds() view returns(uint64[])
func (_XRegistry *XRegistrySession) ChainIds() ([]uint64, error) {
	return _XRegistry.Contract.ChainIds(&_XRegistry.CallOpts)
}

// ChainIds is a free data retrieval call binding the contract method 0xfbe4b7c0.
//
// Solidity: function chainIds() view returns(uint64[])
func (_XRegistry *XRegistryCallerSession) ChainIds() ([]uint64, error) {
	return _XRegistry.Contract.ChainIds(&_XRegistry.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns(address)
func (_XRegistry *XRegistryCaller) Get(opts *bind.CallOpts, chainId uint64, name string, registrant common.Address) (common.Address, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "get", chainId, name, registrant)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns(address)
func (_XRegistry *XRegistrySession) Get(chainId uint64, name string, registrant common.Address) (common.Address, error) {
	return _XRegistry.Contract.Get(&_XRegistry.CallOpts, chainId, name, registrant)
}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns(address)
func (_XRegistry *XRegistryCallerSession) Get(chainId uint64, name string, registrant common.Address) (common.Address, error) {
	return _XRegistry.Contract.Get(&_XRegistry.CallOpts, chainId, name, registrant)
}

// Has is a free data retrieval call binding the contract method 0xfd0b64f7.
//
// Solidity: function has(uint64 chainId, string name, address registrant) view returns(bool)
func (_XRegistry *XRegistryCaller) Has(opts *bind.CallOpts, chainId uint64, name string, registrant common.Address) (bool, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "has", chainId, name, registrant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Has is a free data retrieval call binding the contract method 0xfd0b64f7.
//
// Solidity: function has(uint64 chainId, string name, address registrant) view returns(bool)
func (_XRegistry *XRegistrySession) Has(chainId uint64, name string, registrant common.Address) (bool, error) {
	return _XRegistry.Contract.Has(&_XRegistry.CallOpts, chainId, name, registrant)
}

// Has is a free data retrieval call binding the contract method 0xfd0b64f7.
//
// Solidity: function has(uint64 chainId, string name, address registrant) view returns(bool)
func (_XRegistry *XRegistryCallerSession) Has(chainId uint64, name string, registrant common.Address) (bool, error) {
	return _XRegistry.Contract.Has(&_XRegistry.CallOpts, chainId, name, registrant)
}

// IsSupportedChain is a free data retrieval call binding the contract method 0x8926f54f.
//
// Solidity: function isSupportedChain(uint64 chainId) view returns(bool)
func (_XRegistry *XRegistryCaller) IsSupportedChain(opts *bind.CallOpts, chainId uint64) (bool, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "isSupportedChain", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedChain is a free data retrieval call binding the contract method 0x8926f54f.
//
// Solidity: function isSupportedChain(uint64 chainId) view returns(bool)
func (_XRegistry *XRegistrySession) IsSupportedChain(chainId uint64) (bool, error) {
	return _XRegistry.Contract.IsSupportedChain(&_XRegistry.CallOpts, chainId)
}

// IsSupportedChain is a free data retrieval call binding the contract method 0x8926f54f.
//
// Solidity: function isSupportedChain(uint64 chainId) view returns(bool)
func (_XRegistry *XRegistryCallerSession) IsSupportedChain(chainId uint64) (bool, error) {
	return _XRegistry.Contract.IsSupportedChain(&_XRegistry.CallOpts, chainId)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_XRegistry *XRegistryCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_XRegistry *XRegistrySession) Omni() (common.Address, error) {
	return _XRegistry.Contract.Omni(&_XRegistry.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_XRegistry *XRegistryCallerSession) Omni() (common.Address, error) {
	return _XRegistry.Contract.Omni(&_XRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XRegistry *XRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XRegistry *XRegistrySession) Owner() (common.Address, error) {
	return _XRegistry.Contract.Owner(&_XRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XRegistry *XRegistryCallerSession) Owner() (common.Address, error) {
	return _XRegistry.Contract.Owner(&_XRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0xa4861b42.
//
// Solidity: function registrationFee(uint64 chainId, string name, address addr) view returns(uint256)
func (_XRegistry *XRegistryCaller) RegistrationFee(opts *bind.CallOpts, chainId uint64, name string, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "registrationFee", chainId, name, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationFee is a free data retrieval call binding the contract method 0xa4861b42.
//
// Solidity: function registrationFee(uint64 chainId, string name, address addr) view returns(uint256)
func (_XRegistry *XRegistrySession) RegistrationFee(chainId uint64, name string, addr common.Address) (*big.Int, error) {
	return _XRegistry.Contract.RegistrationFee(&_XRegistry.CallOpts, chainId, name, addr)
}

// RegistrationFee is a free data retrieval call binding the contract method 0xa4861b42.
//
// Solidity: function registrationFee(uint64 chainId, string name, address addr) view returns(uint256)
func (_XRegistry *XRegistryCallerSession) RegistrationFee(chainId uint64, name string, addr common.Address) (*big.Int, error) {
	return _XRegistry.Contract.RegistrationFee(&_XRegistry.CallOpts, chainId, name, addr)
}

// Replicas is a free data retrieval call binding the contract method 0x88f9380b.
//
// Solidity: function replicas(uint64 ) view returns(address)
func (_XRegistry *XRegistryCaller) Replicas(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "replicas", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Replicas is a free data retrieval call binding the contract method 0x88f9380b.
//
// Solidity: function replicas(uint64 ) view returns(address)
func (_XRegistry *XRegistrySession) Replicas(arg0 uint64) (common.Address, error) {
	return _XRegistry.Contract.Replicas(&_XRegistry.CallOpts, arg0)
}

// Replicas is a free data retrieval call binding the contract method 0x88f9380b.
//
// Solidity: function replicas(uint64 ) view returns(address)
func (_XRegistry *XRegistryCallerSession) Replicas(arg0 uint64) (common.Address, error) {
	return _XRegistry.Contract.Replicas(&_XRegistry.CallOpts, arg0)
}

// Register is a paid mutator transaction binding the contract method 0x14a28fcb.
//
// Solidity: function register(uint64 chainId, string name, address addr) payable returns()
func (_XRegistry *XRegistryTransactor) Register(opts *bind.TransactOpts, chainId uint64, name string, addr common.Address) (*types.Transaction, error) {
	return _XRegistry.contract.Transact(opts, "register", chainId, name, addr)
}

// Register is a paid mutator transaction binding the contract method 0x14a28fcb.
//
// Solidity: function register(uint64 chainId, string name, address addr) payable returns()
func (_XRegistry *XRegistrySession) Register(chainId uint64, name string, addr common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.Register(&_XRegistry.TransactOpts, chainId, name, addr)
}

// Register is a paid mutator transaction binding the contract method 0x14a28fcb.
//
// Solidity: function register(uint64 chainId, string name, address addr) payable returns()
func (_XRegistry *XRegistryTransactorSession) Register(chainId uint64, name string, addr common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.Register(&_XRegistry.TransactOpts, chainId, name, addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XRegistry *XRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XRegistry *XRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _XRegistry.Contract.RenounceOwnership(&_XRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XRegistry *XRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _XRegistry.Contract.RenounceOwnership(&_XRegistry.TransactOpts)
}

// SetPortal is a paid mutator transaction binding the contract method 0x4ff56192.
//
// Solidity: function setPortal(address _omni) returns()
func (_XRegistry *XRegistryTransactor) SetPortal(opts *bind.TransactOpts, _omni common.Address) (*types.Transaction, error) {
	return _XRegistry.contract.Transact(opts, "setPortal", _omni)
}

// SetPortal is a paid mutator transaction binding the contract method 0x4ff56192.
//
// Solidity: function setPortal(address _omni) returns()
func (_XRegistry *XRegistrySession) SetPortal(_omni common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.SetPortal(&_XRegistry.TransactOpts, _omni)
}

// SetPortal is a paid mutator transaction binding the contract method 0x4ff56192.
//
// Solidity: function setPortal(address _omni) returns()
func (_XRegistry *XRegistryTransactorSession) SetPortal(_omni common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.SetPortal(&_XRegistry.TransactOpts, _omni)
}

// SetReplica is a paid mutator transaction binding the contract method 0x7b7c0ddc.
//
// Solidity: function setReplica(uint64 chainId, address replica) returns()
func (_XRegistry *XRegistryTransactor) SetReplica(opts *bind.TransactOpts, chainId uint64, replica common.Address) (*types.Transaction, error) {
	return _XRegistry.contract.Transact(opts, "setReplica", chainId, replica)
}

// SetReplica is a paid mutator transaction binding the contract method 0x7b7c0ddc.
//
// Solidity: function setReplica(uint64 chainId, address replica) returns()
func (_XRegistry *XRegistrySession) SetReplica(chainId uint64, replica common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.SetReplica(&_XRegistry.TransactOpts, chainId, replica)
}

// SetReplica is a paid mutator transaction binding the contract method 0x7b7c0ddc.
//
// Solidity: function setReplica(uint64 chainId, address replica) returns()
func (_XRegistry *XRegistryTransactorSession) SetReplica(chainId uint64, replica common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.SetReplica(&_XRegistry.TransactOpts, chainId, replica)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XRegistry *XRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _XRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XRegistry *XRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.TransferOwnership(&_XRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XRegistry *XRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.TransferOwnership(&_XRegistry.TransactOpts, newOwner)
}

// XRegistryContractRegisteredIterator is returned from FilterContractRegistered and is used to iterate over the raw logs and unpacked data for ContractRegistered events raised by the XRegistry contract.
type XRegistryContractRegisteredIterator struct {
	Event *XRegistryContractRegistered // Event containing the contract specifics and raw log

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
func (it *XRegistryContractRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XRegistryContractRegistered)
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
		it.Event = new(XRegistryContractRegistered)
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
func (it *XRegistryContractRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XRegistryContractRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XRegistryContractRegistered represents a ContractRegistered event raised by the XRegistry contract.
type XRegistryContractRegistered struct {
	ChainId    uint64
	Name       common.Hash
	Registrant common.Address
	Addr       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterContractRegistered is a free log retrieval operation binding the contract event 0xd06596d338531cdc7b5a36893f1bae3902aa239c822f38d70301ca871f855fac.
//
// Solidity: event ContractRegistered(uint64 indexed chainId, string indexed name, address indexed registrant, address addr)
func (_XRegistry *XRegistryFilterer) FilterContractRegistered(opts *bind.FilterOpts, chainId []uint64, name []string, registrant []common.Address) (*XRegistryContractRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _XRegistry.contract.FilterLogs(opts, "ContractRegistered", chainIdRule, nameRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return &XRegistryContractRegisteredIterator{contract: _XRegistry.contract, event: "ContractRegistered", logs: logs, sub: sub}, nil
}

// WatchContractRegistered is a free log subscription operation binding the contract event 0xd06596d338531cdc7b5a36893f1bae3902aa239c822f38d70301ca871f855fac.
//
// Solidity: event ContractRegistered(uint64 indexed chainId, string indexed name, address indexed registrant, address addr)
func (_XRegistry *XRegistryFilterer) WatchContractRegistered(opts *bind.WatchOpts, sink chan<- *XRegistryContractRegistered, chainId []uint64, name []string, registrant []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _XRegistry.contract.WatchLogs(opts, "ContractRegistered", chainIdRule, nameRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XRegistryContractRegistered)
				if err := _XRegistry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
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

// ParseContractRegistered is a log parse operation binding the contract event 0xd06596d338531cdc7b5a36893f1bae3902aa239c822f38d70301ca871f855fac.
//
// Solidity: event ContractRegistered(uint64 indexed chainId, string indexed name, address indexed registrant, address addr)
func (_XRegistry *XRegistryFilterer) ParseContractRegistered(log types.Log) (*XRegistryContractRegistered, error) {
	event := new(XRegistryContractRegistered)
	if err := _XRegistry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the XRegistry contract.
type XRegistryOwnershipTransferredIterator struct {
	Event *XRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *XRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XRegistryOwnershipTransferred)
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
		it.Event = new(XRegistryOwnershipTransferred)
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
func (it *XRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the XRegistry contract.
type XRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XRegistry *XRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*XRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _XRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &XRegistryOwnershipTransferredIterator{contract: _XRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XRegistry *XRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *XRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _XRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XRegistryOwnershipTransferred)
				if err := _XRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_XRegistry *XRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*XRegistryOwnershipTransferred, error) {
	event := new(XRegistryOwnershipTransferred)
	if err := _XRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
