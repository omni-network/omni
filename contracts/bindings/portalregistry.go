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

// PortalRegistryDeployment is an auto generated low-level Go binding around an user-defined struct.
type PortalRegistryDeployment struct {
	ChainId           uint64
	Addr              common.Address
	DeployHeight      uint64
	FinalizationStrat string
}

// PortalRegistryMetaData contains all meta data concerning the PortalRegistry contract.
var PortalRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deployments\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizationStrat\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizationStrat\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"list\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structPortalRegistry.Deployment[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizationStrat\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"deployment\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizationStrat\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"registrationFee\",\"inputs\":[{\"name\":\"deployment\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"finalizationStrat\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"xregistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractXRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PortalRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"finalizationStrat\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61132a8061007e6000396000f3fe6080604052600436106100915760003560e01c806386ac16841161005957806386ac1684146101675780638da5cb5b1461017a578063ada8679814610198578063c2a1402d146101c5578063f2fde38b1461021d57600080fd5b80630470c5fc146100965780630f560cd7146100c9578063473d0452146100eb578063715018a61461011b578063738b07a614610132575b600080fd5b3480156100a257600080fd5b506100b66100b1366004610cbb565b61023d565b6040519081526020015b60405180910390f35b3480156100d557600080fd5b506100de6102f6565b6040516100c09190610d92565b3480156100f757600080fd5b5061010b610106366004610e0b565b610611565b6040516100c09493929190610e28565b34801561012757600080fd5b506101306106dc565b005b34801561013e57600080fd5b5061014f600462048789608a1b0181565b6040516001600160a01b0390911681526020016100c0565b610130610175366004610cbb565b6106f0565b34801561018657600080fd5b506000546001600160a01b031661014f565b3480156101a457600080fd5b506101b86101b3366004610e0b565b610a86565b6040516100c09190610e6d565b3480156101d157600080fd5b5061020d6101e0366004610e0b565b6001600160401b0316600090815260016020526040902054600160401b90046001600160a01b0316151590565b60405190151581526020016100c0565b34801561022957600080fd5b50610130610238366004610e95565b610b98565b6000600462048789608a1b0163a4861b4261025b6020850185610e0b565b6040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b8152508560200160208101906102919190610e95565b6040518463ffffffff1660e01b81526004016102af93929190610eb2565b602060405180830381865afa1580156102cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f09190610eec565b92915050565b60606000600462048789608a1b016001600160a01b031663fbe4b7c06040518163ffffffff1660e01b8152600401600060405180830381865afa158015610341573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103699190810190610f2b565b90506000805b8251816001600160401b031610156103f9576103d483826001600160401b03168151811061039f5761039f610fef565b60200260200101516001600160401b0316600090815260016020526040902054600160401b90046001600160a01b0316151590565b156103e757816103e381611005565b9250505b806103f181611005565b91505061036f565b506000816001600160401b03166001600160401b0381111561041d5761041d610f05565b60405190808252806020026020018201604052801561046e57816020015b604080516080810182526000808252602080830182905292820152606080820152825260001990920191018161043b5790505b5090506000805b8451816001600160401b03161015610607576104a585826001600160401b03168151811061039f5761039f610fef565b156105f5576001600086836001600160401b0316815181106104c9576104c9610fef565b6020908102919091018101516001600160401b039081168352828201939093526040918201600020825160808101845281548086168252600160401b90046001600160a01b03169281019290925260018101549093169181019190915260028201805491929160608401919061053e90611039565b80601f016020809104026020016040519081016040528092919081815260200182805461056a90611039565b80156105b75780601f1061058c576101008083540402835291602001916105b7565b820191906000526020600020905b81548152906001019060200180831161059a57829003601f168201915b50505050508152505083836001600160401b0316815181106105db576105db610fef565b602002602001018190525081806105f190611005565b9250505b806105ff81611005565b915050610475565b5090949350505050565b60016020819052600091825260409091208054918101546002820180546001600160401b0380861695600160401b90046001600160a01b031694931692919061065990611039565b80601f016020809104026020016040519081016040528092919081815260200182805461068590611039565b80156106d25780601f106106a7576101008083540402835291602001916106d2565b820191906000526020600020905b8154815290600101906020018083116106b557829003601f168201915b5050505050905084565b6106e4610c11565b6106ee6000610c6b565b565b6106f8610c11565b6107086101e06020830183610e0b565b1561075a5760405162461bcd60e51b815260206004820152601b60248201527f506f7274616c52656769737472793a20616c726561647920736574000000000060448201526064015b60405180910390fd5b600061076c6040830160208401610e95565b6001600160a01b0316036107c25760405162461bcd60e51b815260206004820152601c60248201527f506f7274616c52656769737472793a207a65726f2061646472657373000000006044820152606401610751565b6107cf6060820182611073565b90506000036108205760405162461bcd60e51b815260206004820152601b60248201527f506f7274616c52656769737472793a20656d70747920737472617400000000006044820152606401610751565b600462048789608a1b0163a4861b4261083c6020840184610e0b565b6040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b8152508460200160208101906108729190610e95565b6040518463ffffffff1660e01b815260040161089093929190610eb2565b602060405180830381865afa1580156108ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d19190610eec565b3410156109205760405162461bcd60e51b815260206004820181905260248201527f506f7274616c52656769737472793a20696e73756666696369656e74206665656044820152606401610751565b600462048789608a1b016314a28fcb3461093d6020850185610e0b565b6040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b8152508560200160208101906109739190610e95565b6040518563ffffffff1660e01b815260040161099193929190610eb2565b6000604051808303818588803b1580156109aa57600080fd5b505af11580156109be573d6000803e3d6000fd5b5084935060019250600091506109d990506020840184610e0b565b6001600160401b0316815260208101919091526040016000206109fc82826111d1565b50610a0f90506040820160208301610e95565b6001600160a01b0316610a256020830183610e0b565b6001600160401b03167f20c588e9e1e07556ad236d2666ba0a232806d504da63e50d5db8429d15ac77c8610a5f6060850160408601610e0b565b610a6c6060860186611073565b604051610a7b939291906112b5565b60405180910390a350565b60408051608081018252600080825260208201819052918101919091526060808201526001600160401b03828116600090815260016020818152604092839020835160808101855281548087168252600160401b90046001600160a01b0316928101929092529182015490931691830191909152600281018054606084019190610b0f90611039565b80601f0160208091040260200160405190810160405280929190818152602001828054610b3b90611039565b8015610b885780601f10610b5d57610100808354040283529160200191610b88565b820191906000526020600020905b815481529060010190602001808311610b6b57829003601f168201915b5050505050815250509050919050565b610ba0610c11565b6001600160a01b038116610c055760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610751565b610c0e81610c6b565b50565b6000546001600160a01b031633146106ee5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610751565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215610ccd57600080fd5b81356001600160401b03811115610ce357600080fd5b820160808185031215610cf557600080fd5b9392505050565b6000815180845260005b81811015610d2257602081850181015186830182015201610d06565b506000602082860101526020601f19601f83011685010191505092915050565b60006001600160401b0380835116845260018060a01b03602084015116602085015280604084015116604085015250606082015160806060850152610d8a6080850182610cfc565b949350505050565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b82811015610de957603f19888603018452610dd7858351610d42565b94509285019290850190600101610dbb565b5092979650505050505050565b6001600160401b0381168114610c0e57600080fd5b600060208284031215610e1d57600080fd5b8135610cf581610df6565b6001600160401b0385811682526001600160a01b038516602083015283166040820152608060608201819052600090610e6390830184610cfc565b9695505050505050565b602081526000610cf56020830184610d42565b6001600160a01b0381168114610c0e57600080fd5b600060208284031215610ea757600080fd5b8135610cf581610e80565b6001600160401b0384168152606060208201526000610ed46060830185610cfc565b905060018060a01b0383166040830152949350505050565b600060208284031215610efe57600080fd5b5051919050565b634e487b7160e01b600052604160045260246000fd5b8051610f2681610df6565b919050565b60006020808385031215610f3e57600080fd5b82516001600160401b0380821115610f5557600080fd5b818501915085601f830112610f6957600080fd5b815181811115610f7b57610f7b610f05565b8060051b604051601f19603f83011681018181108582111715610fa057610fa0610f05565b604052918252848201925083810185019188831115610fbe57600080fd5b938501935b82851015610fe357610fd485610f1b565b84529385019392850192610fc3565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b60006001600160401b0380831681810361102f57634e487b7160e01b600052601160045260246000fd5b6001019392505050565b600181811c9082168061104d57607f821691505b60208210810361106d57634e487b7160e01b600052602260045260246000fd5b50919050565b6000808335601e1984360301811261108a57600080fd5b8301803591506001600160401b038211156110a457600080fd5b6020019150368190038213156110b957600080fd5b9250929050565b601f82111561110c576000816000526020600020601f850160051c810160208610156110e95750805b601f850160051c820191505b81811015611108578281556001016110f5565b5050505b505050565b6001600160401b0383111561112857611128610f05565b61113c836111368354611039565b836110c0565b6000601f84116001811461117057600085156111585750838201355b600019600387901b1c1916600186901b1783556111ca565b600083815260209020601f19861690835b828110156111a15786850135825560209485019460019092019101611181565b50868210156111be5760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b81356111dc81610df6565b815467ffffffffffffffff19166001600160401b03821617825550602082013561120581610e80565b815468010000000000000000600160e01b031916604091821b68010000000000000000600160e01b031617825582013561123e81610df6565b60018201805467ffffffffffffffff19166001600160401b038316179055506060820135601e1983360301811261127457600080fd5b820180356001600160401b0381111561128c57600080fd5b6020820191508036038213156112a157600080fd5b6112af818360028601611111565b50505050565b6001600160401b038416815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f191601019291505056fea2646970667358221220cfa65a02784574c4ba96b1086123732b52a1f2fa3270c517b0e92e44df4aa26864736f6c63430008180033",
}

// PortalRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PortalRegistryMetaData.ABI instead.
var PortalRegistryABI = PortalRegistryMetaData.ABI

// PortalRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PortalRegistryMetaData.Bin instead.
var PortalRegistryBin = PortalRegistryMetaData.Bin

// DeployPortalRegistry deploys a new Ethereum contract, binding an instance of PortalRegistry to it.
func DeployPortalRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PortalRegistry, error) {
	parsed, err := PortalRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PortalRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PortalRegistry{PortalRegistryCaller: PortalRegistryCaller{contract: contract}, PortalRegistryTransactor: PortalRegistryTransactor{contract: contract}, PortalRegistryFilterer: PortalRegistryFilterer{contract: contract}}, nil
}

// PortalRegistry is an auto generated Go binding around an Ethereum contract.
type PortalRegistry struct {
	PortalRegistryCaller     // Read-only binding to the contract
	PortalRegistryTransactor // Write-only binding to the contract
	PortalRegistryFilterer   // Log filterer for contract events
}

// PortalRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PortalRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PortalRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PortalRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PortalRegistrySession struct {
	Contract     *PortalRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortalRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PortalRegistryCallerSession struct {
	Contract *PortalRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PortalRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PortalRegistryTransactorSession struct {
	Contract     *PortalRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PortalRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PortalRegistryRaw struct {
	Contract *PortalRegistry // Generic contract binding to access the raw methods on
}

// PortalRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PortalRegistryCallerRaw struct {
	Contract *PortalRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PortalRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PortalRegistryTransactorRaw struct {
	Contract *PortalRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPortalRegistry creates a new instance of PortalRegistry, bound to a specific deployed contract.
func NewPortalRegistry(address common.Address, backend bind.ContractBackend) (*PortalRegistry, error) {
	contract, err := bindPortalRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PortalRegistry{PortalRegistryCaller: PortalRegistryCaller{contract: contract}, PortalRegistryTransactor: PortalRegistryTransactor{contract: contract}, PortalRegistryFilterer: PortalRegistryFilterer{contract: contract}}, nil
}

// NewPortalRegistryCaller creates a new read-only instance of PortalRegistry, bound to a specific deployed contract.
func NewPortalRegistryCaller(address common.Address, caller bind.ContractCaller) (*PortalRegistryCaller, error) {
	contract, err := bindPortalRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryCaller{contract: contract}, nil
}

// NewPortalRegistryTransactor creates a new write-only instance of PortalRegistry, bound to a specific deployed contract.
func NewPortalRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PortalRegistryTransactor, error) {
	contract, err := bindPortalRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryTransactor{contract: contract}, nil
}

// NewPortalRegistryFilterer creates a new log filterer instance of PortalRegistry, bound to a specific deployed contract.
func NewPortalRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PortalRegistryFilterer, error) {
	contract, err := bindPortalRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryFilterer{contract: contract}, nil
}

// bindPortalRegistry binds a generic wrapper to an already deployed contract.
func bindPortalRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PortalRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PortalRegistry *PortalRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PortalRegistry.Contract.PortalRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PortalRegistry *PortalRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PortalRegistry.Contract.PortalRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PortalRegistry *PortalRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PortalRegistry.Contract.PortalRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PortalRegistry *PortalRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PortalRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PortalRegistry *PortalRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PortalRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PortalRegistry *PortalRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PortalRegistry.Contract.contract.Transact(opts, method, params...)
}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(uint64 chainId, address addr, uint64 deployHeight, string finalizationStrat)
func (_PortalRegistry *PortalRegistryCaller) Deployments(opts *bind.CallOpts, arg0 uint64) (struct {
	ChainId           uint64
	Addr              common.Address
	DeployHeight      uint64
	FinalizationStrat string
}, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "deployments", arg0)

	outstruct := new(struct {
		ChainId           uint64
		Addr              common.Address
		DeployHeight      uint64
		FinalizationStrat string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Addr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DeployHeight = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.FinalizationStrat = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(uint64 chainId, address addr, uint64 deployHeight, string finalizationStrat)
func (_PortalRegistry *PortalRegistrySession) Deployments(arg0 uint64) (struct {
	ChainId           uint64
	Addr              common.Address
	DeployHeight      uint64
	FinalizationStrat string
}, error) {
	return _PortalRegistry.Contract.Deployments(&_PortalRegistry.CallOpts, arg0)
}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(uint64 chainId, address addr, uint64 deployHeight, string finalizationStrat)
func (_PortalRegistry *PortalRegistryCallerSession) Deployments(arg0 uint64) (struct {
	ChainId           uint64
	Addr              common.Address
	DeployHeight      uint64
	FinalizationStrat string
}, error) {
	return _PortalRegistry.Contract.Deployments(&_PortalRegistry.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 chainId) view returns((uint64,address,uint64,string))
func (_PortalRegistry *PortalRegistryCaller) Get(opts *bind.CallOpts, chainId uint64) (PortalRegistryDeployment, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "get", chainId)

	if err != nil {
		return *new(PortalRegistryDeployment), err
	}

	out0 := *abi.ConvertType(out[0], new(PortalRegistryDeployment)).(*PortalRegistryDeployment)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 chainId) view returns((uint64,address,uint64,string))
func (_PortalRegistry *PortalRegistrySession) Get(chainId uint64) (PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.Get(&_PortalRegistry.CallOpts, chainId)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 chainId) view returns((uint64,address,uint64,string))
func (_PortalRegistry *PortalRegistryCallerSession) Get(chainId uint64) (PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.Get(&_PortalRegistry.CallOpts, chainId)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc2a1402d.
//
// Solidity: function isRegistered(uint64 chainId) view returns(bool)
func (_PortalRegistry *PortalRegistryCaller) IsRegistered(opts *bind.CallOpts, chainId uint64) (bool, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "isRegistered", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc2a1402d.
//
// Solidity: function isRegistered(uint64 chainId) view returns(bool)
func (_PortalRegistry *PortalRegistrySession) IsRegistered(chainId uint64) (bool, error) {
	return _PortalRegistry.Contract.IsRegistered(&_PortalRegistry.CallOpts, chainId)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc2a1402d.
//
// Solidity: function isRegistered(uint64 chainId) view returns(bool)
func (_PortalRegistry *PortalRegistryCallerSession) IsRegistered(chainId uint64) (bool, error) {
	return _PortalRegistry.Contract.IsRegistered(&_PortalRegistry.CallOpts, chainId)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((uint64,address,uint64,string)[])
func (_PortalRegistry *PortalRegistryCaller) List(opts *bind.CallOpts) ([]PortalRegistryDeployment, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]PortalRegistryDeployment), err
	}

	out0 := *abi.ConvertType(out[0], new([]PortalRegistryDeployment)).(*[]PortalRegistryDeployment)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((uint64,address,uint64,string)[])
func (_PortalRegistry *PortalRegistrySession) List() ([]PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.List(&_PortalRegistry.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((uint64,address,uint64,string)[])
func (_PortalRegistry *PortalRegistryCallerSession) List() ([]PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.List(&_PortalRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PortalRegistry *PortalRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PortalRegistry *PortalRegistrySession) Owner() (common.Address, error) {
	return _PortalRegistry.Contract.Owner(&_PortalRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PortalRegistry *PortalRegistryCallerSession) Owner() (common.Address, error) {
	return _PortalRegistry.Contract.Owner(&_PortalRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x0470c5fc.
//
// Solidity: function registrationFee((uint64,address,uint64,string) deployment) view returns(uint256)
func (_PortalRegistry *PortalRegistryCaller) RegistrationFee(opts *bind.CallOpts, deployment PortalRegistryDeployment) (*big.Int, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "registrationFee", deployment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationFee is a free data retrieval call binding the contract method 0x0470c5fc.
//
// Solidity: function registrationFee((uint64,address,uint64,string) deployment) view returns(uint256)
func (_PortalRegistry *PortalRegistrySession) RegistrationFee(deployment PortalRegistryDeployment) (*big.Int, error) {
	return _PortalRegistry.Contract.RegistrationFee(&_PortalRegistry.CallOpts, deployment)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x0470c5fc.
//
// Solidity: function registrationFee((uint64,address,uint64,string) deployment) view returns(uint256)
func (_PortalRegistry *PortalRegistryCallerSession) RegistrationFee(deployment PortalRegistryDeployment) (*big.Int, error) {
	return _PortalRegistry.Contract.RegistrationFee(&_PortalRegistry.CallOpts, deployment)
}

// Xregistry is a free data retrieval call binding the contract method 0x738b07a6.
//
// Solidity: function xregistry() view returns(address)
func (_PortalRegistry *PortalRegistryCaller) Xregistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "xregistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xregistry is a free data retrieval call binding the contract method 0x738b07a6.
//
// Solidity: function xregistry() view returns(address)
func (_PortalRegistry *PortalRegistrySession) Xregistry() (common.Address, error) {
	return _PortalRegistry.Contract.Xregistry(&_PortalRegistry.CallOpts)
}

// Xregistry is a free data retrieval call binding the contract method 0x738b07a6.
//
// Solidity: function xregistry() view returns(address)
func (_PortalRegistry *PortalRegistryCallerSession) Xregistry() (common.Address, error) {
	return _PortalRegistry.Contract.Xregistry(&_PortalRegistry.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0x86ac1684.
//
// Solidity: function register((uint64,address,uint64,string) deployment) payable returns()
func (_PortalRegistry *PortalRegistryTransactor) Register(opts *bind.TransactOpts, deployment PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "register", deployment)
}

// Register is a paid mutator transaction binding the contract method 0x86ac1684.
//
// Solidity: function register((uint64,address,uint64,string) deployment) payable returns()
func (_PortalRegistry *PortalRegistrySession) Register(deployment PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.Register(&_PortalRegistry.TransactOpts, deployment)
}

// Register is a paid mutator transaction binding the contract method 0x86ac1684.
//
// Solidity: function register((uint64,address,uint64,string) deployment) payable returns()
func (_PortalRegistry *PortalRegistryTransactorSession) Register(deployment PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.Register(&_PortalRegistry.TransactOpts, deployment)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PortalRegistry *PortalRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PortalRegistry *PortalRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _PortalRegistry.Contract.RenounceOwnership(&_PortalRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PortalRegistry *PortalRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PortalRegistry.Contract.RenounceOwnership(&_PortalRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PortalRegistry *PortalRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PortalRegistry *PortalRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PortalRegistry.Contract.TransferOwnership(&_PortalRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PortalRegistry *PortalRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PortalRegistry.Contract.TransferOwnership(&_PortalRegistry.TransactOpts, newOwner)
}

// PortalRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PortalRegistry contract.
type PortalRegistryOwnershipTransferredIterator struct {
	Event *PortalRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PortalRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRegistryOwnershipTransferred)
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
		it.Event = new(PortalRegistryOwnershipTransferred)
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
func (it *PortalRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the PortalRegistry contract.
type PortalRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PortalRegistry *PortalRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PortalRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PortalRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryOwnershipTransferredIterator{contract: _PortalRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PortalRegistry *PortalRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PortalRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PortalRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRegistryOwnershipTransferred)
				if err := _PortalRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PortalRegistry *PortalRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*PortalRegistryOwnershipTransferred, error) {
	event := new(PortalRegistryOwnershipTransferred)
	if err := _PortalRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalRegistryPortalRegisteredIterator is returned from FilterPortalRegistered and is used to iterate over the raw logs and unpacked data for PortalRegistered events raised by the PortalRegistry contract.
type PortalRegistryPortalRegisteredIterator struct {
	Event *PortalRegistryPortalRegistered // Event containing the contract specifics and raw log

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
func (it *PortalRegistryPortalRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRegistryPortalRegistered)
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
		it.Event = new(PortalRegistryPortalRegistered)
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
func (it *PortalRegistryPortalRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRegistryPortalRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRegistryPortalRegistered represents a PortalRegistered event raised by the PortalRegistry contract.
type PortalRegistryPortalRegistered struct {
	ChainId           uint64
	Addr              common.Address
	DeployHeight      uint64
	FinalizationStrat string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPortalRegistered is a free log retrieval operation binding the contract event 0x20c588e9e1e07556ad236d2666ba0a232806d504da63e50d5db8429d15ac77c8.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, string finalizationStrat)
func (_PortalRegistry *PortalRegistryFilterer) FilterPortalRegistered(opts *bind.FilterOpts, chainId []uint64, addr []common.Address) (*PortalRegistryPortalRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _PortalRegistry.contract.FilterLogs(opts, "PortalRegistered", chainIdRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryPortalRegisteredIterator{contract: _PortalRegistry.contract, event: "PortalRegistered", logs: logs, sub: sub}, nil
}

// WatchPortalRegistered is a free log subscription operation binding the contract event 0x20c588e9e1e07556ad236d2666ba0a232806d504da63e50d5db8429d15ac77c8.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, string finalizationStrat)
func (_PortalRegistry *PortalRegistryFilterer) WatchPortalRegistered(opts *bind.WatchOpts, sink chan<- *PortalRegistryPortalRegistered, chainId []uint64, addr []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _PortalRegistry.contract.WatchLogs(opts, "PortalRegistered", chainIdRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRegistryPortalRegistered)
				if err := _PortalRegistry.contract.UnpackLog(event, "PortalRegistered", log); err != nil {
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

// ParsePortalRegistered is a log parse operation binding the contract event 0x20c588e9e1e07556ad236d2666ba0a232806d504da63e50d5db8429d15ac77c8.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, string finalizationStrat)
func (_PortalRegistry *PortalRegistryFilterer) ParsePortalRegistered(log types.Log) (*PortalRegistryPortalRegistered, error) {
	event := new(PortalRegistryPortalRegistered)
	if err := _PortalRegistry.contract.UnpackLog(event, "PortalRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
