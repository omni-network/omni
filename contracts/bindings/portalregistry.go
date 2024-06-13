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
	ChainId      uint64
	Addr         common.Address
	DeployHeight uint64
	Shards       []uint64
}

// PortalRegistryMetaData contains all meta data concerning the PortalRegistry contract.
var PortalRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"bulkRegister\",\"inputs\":[{\"name\":\"deps\",\"type\":\"tuple[]\",\"internalType\":\"structPortalRegistry.Deployment[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"chainIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployments\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"list\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structPortalRegistry.Deployment[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"dep\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PortalRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506110a2806100206000396000f3fe6080604052600436106100865760003560e01c8063510644e811610059578063510644e814610194578063715018a6146101a75780638da5cb5b146101bc578063ada86798146101e4578063f2fde38b1461021157600080fd5b80630f560cd71461008b57806321d93090146100b6578063278679c5146100ee578063473d045214610110575b600080fd5b34801561009757600080fd5b506100a0610231565b6040516100ad9190610b4a565b60405180910390f35b3480156100c257600080fd5b506100d66100d1366004610bae565b61040d565b6040516001600160401b0390911681526020016100ad565b3480156100fa57600080fd5b5061010e610109366004610bc7565b61044a565b005b34801561011c57600080fd5b5061016261012b366004610c1d565b606660205260009081526040902080546001909101546001600160401b0380831692600160401b90046001600160a01b0316911683565b604080516001600160401b0394851681526001600160a01b0390931660208401529216918101919091526060016100ad565b61010e6101a2366004610c3a565b61045e565b3480156101b357600080fd5b5061010e6104c6565b3480156101c857600080fd5b506033546040516001600160a01b0390911681526020016100ad565b3480156101f057600080fd5b506102046101ff366004610c1d565b6104da565b6040516100ad9190610cae565b34801561021d57600080fd5b5061010e61022c366004610cd6565b6105d9565b6065546060906000906001600160401b0381111561025157610251610cf3565b6040519080825280602002602001820160405280156102a257816020015b604080516080810182526000808252602080830182905292820152606080820152825260001990920191018161026f5790505b50905060005b6065546001600160401b038216101561040757606660006065836001600160401b0316815481106102db576102db610d09565b6000918252602080832060048304015460039092166008026101000a9091046001600160401b039081168452838201949094526040928301909120825160808101845281548086168252600160401b90046001600160a01b031681840152600182015490941684840152600281018054845181850281018501909552808552919360608601939092908301828280156103c557602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116103825790505b50505050508152505082826001600160401b0316815181106103e9576103e9610d09565b602002602001018190525080806103ff90610d1f565b9150506102a8565b50919050565b6065818154811061041d57600080fd5b9060005260206000209060049182820401919006600802915054906101000a90046001600160401b031681565b610452610654565b61045b816106ae565b50565b610466610654565b60005b6001600160401b0381168211156104c1576104af8383836001600160401b031681811061049857610498610d09565b90506020028101906104aa9190610d53565b6106ae565b806104b981610d1f565b915050610469565b505050565b6104ce610654565b6104d86000610a53565b565b60408051608081018252600080825260208201819052918101919091526060808201526001600160401b03828116600090815260666020908152604091829020825160808101845281548086168252600160401b90046001600160a01b031681840152600182015490941684840152600281018054845181850281018501909552808552919360608601939092908301828280156105c957602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116105865790505b5050505050815250509050919050565b6105e1610654565b6001600160a01b03811661064b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61045b81610a53565b6033546001600160a01b031633146104d85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610642565b60006106c06040830160208401610cd6565b6001600160a01b0316036107165760405162461bcd60e51b815260206004820152601c60248201527f506f7274616c52656769737472793a206e6f207a65726f2061646472000000006044820152606401610642565b60006107256020830183610c1d565b6001600160401b03161161077b5760405162461bcd60e51b815260206004820181905260248201527f506f7274616c52656769737472793a206e6f207a65726f20636861696e2049446044820152606401610642565b600061078a6060830183610d73565b9050116107d95760405162461bcd60e51b815260206004820152601960248201527f506f7274616c52656769737472793a206e6f20736861726473000000000000006044820152606401610642565b60006066816107eb6020850185610c1d565b6001600160401b03168152602081019190915260400160002054600160401b90046001600160a01b0316146108625760405162461bcd60e51b815260206004820152601b60248201527f506f7274616c52656769737472793a20616c72656164792073657400000000006044820152606401610642565b60005b6108726060830183610d73565b9050816001600160401b031610156109415760006108936060840184610d73565b836001600160401b03168181106108ac576108ac610d09565b90506020020160208101906108c19190610c1d565b90508060ff16816001600160401b03161480156108e257506108e281610aa5565b61092e5760405162461bcd60e51b815260206004820152601d60248201527f506f7274616c52656769737472793a20696e76616c69642073686172640000006044820152606401610642565b508061093981610d1f565b915050610865565b5080606660006109546020840184610c1d565b6001600160401b0316815260208101919091526040016000206109778282610f29565b50606590506109896020830183610c1d565b815460018101835560009283526020928390206004820401805460039092166008026101000a6001600160401b038181021990931693909216919091029190911790556109dc9060408301908301610cd6565b6001600160a01b03166109f26020830183610c1d565b6001600160401b03167f99fa7076dae8857571c277f28e10ca59528e679baa7981e731a3cd4f877b4f75610a2c6060850160408601610c1d565b610a396060860186610d73565b604051610a4893929190611010565b60405180910390a350565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000600160ff831610801590610abf5750600460ff831611155b92915050565b6000608083016001600160401b03808451168552602060018060a01b03602086015116602087015281604086015116604087015260608501516080606088015283815180865260a089019150602083019550600092505b80831015610b3e57855185168252948301946001929092019190830190610b1c565b50979650505050505050565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b82811015610ba157603f19888603018452610b8f858351610ac5565b94509285019290850190600101610b73565b5092979650505050505050565b600060208284031215610bc057600080fd5b5035919050565b600060208284031215610bd957600080fd5b81356001600160401b03811115610bef57600080fd5b820160808185031215610c0157600080fd5b9392505050565b6001600160401b038116811461045b57600080fd5b600060208284031215610c2f57600080fd5b8135610c0181610c08565b60008060208385031215610c4d57600080fd5b82356001600160401b0380821115610c6457600080fd5b818501915085601f830112610c7857600080fd5b813581811115610c8757600080fd5b8660208260051b8501011115610c9c57600080fd5b60209290920196919550909350505050565b602081526000610c016020830184610ac5565b6001600160a01b038116811461045b57600080fd5b600060208284031215610ce857600080fd5b8135610c0181610cc1565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60006001600160401b03808316818103610d4957634e487b7160e01b600052601160045260246000fd5b6001019392505050565b60008235607e19833603018112610d6957600080fd5b9190910192915050565b6000808335601e19843603018112610d8a57600080fd5b8301803591506001600160401b03821115610da457600080fd5b6020019150600581901b3603821315610dbc57600080fd5b9250929050565b60008135610abf81610c08565b600160401b821115610de457610de4610cf3565b8054828255808310156104c15760008260005260206000206003850160021c81016003840160021c8201915060188660031b168015610e34576000198083018054828460200360031b1c16815550505b505b81811015610e4957828155600101610e36565b505050505050565b6001600160401b03831115610e6857610e68610cf3565b610e728382610dd0565b60008181526020902082908460021c60005b81811015610ee0576000805b6004811015610ed357610ec2610ea587610dc3565b6001600160401b03908116600684901b90811b91901b1984161790565b602096909601959150600101610e90565b5083820155600101610e84565b506003198616808703818814610f1f576000805b82811015610f1957610f08610ea588610dc3565b602097909701969150600101610ef4565b50848401555b5050505050505050565b8135610f3481610c08565b815467ffffffffffffffff19166001600160401b038216178255506020820135610f5d81610cc1565b815468010000000000000000600160e01b031916604091821b68010000000000000000600160e01b0316178255820135610f9681610c08565b60018201805467ffffffffffffffff19166001600160401b038316179055506060820135601e19833603018112610fcc57600080fd5b820180356001600160401b03811115610fe457600080fd5b6020820191508060051b3603821315610ffc57600080fd5b61100a818360028601610e51565b50505050565b6001600160401b03848116825260406020808401829052908301849052600091859160608501845b8781101561105f57843561104b81610c08565b841682529382019390820190600101611038565b509897505050505050505056fea264697066735822122098c7b63839660f1f7c9b1d6f5b25ee4f00454e24be05a4b35468a8e9ace7a6e764736f6c63430008180033",
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

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(uint64)
func (_PortalRegistry *PortalRegistryCaller) ChainIds(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "chainIds", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(uint64)
func (_PortalRegistry *PortalRegistrySession) ChainIds(arg0 *big.Int) (uint64, error) {
	return _PortalRegistry.Contract.ChainIds(&_PortalRegistry.CallOpts, arg0)
}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(uint64)
func (_PortalRegistry *PortalRegistryCallerSession) ChainIds(arg0 *big.Int) (uint64, error) {
	return _PortalRegistry.Contract.ChainIds(&_PortalRegistry.CallOpts, arg0)
}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(uint64 chainId, address addr, uint64 deployHeight)
func (_PortalRegistry *PortalRegistryCaller) Deployments(opts *bind.CallOpts, arg0 uint64) (struct {
	ChainId      uint64
	Addr         common.Address
	DeployHeight uint64
}, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "deployments", arg0)

	outstruct := new(struct {
		ChainId      uint64
		Addr         common.Address
		DeployHeight uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Addr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DeployHeight = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(uint64 chainId, address addr, uint64 deployHeight)
func (_PortalRegistry *PortalRegistrySession) Deployments(arg0 uint64) (struct {
	ChainId      uint64
	Addr         common.Address
	DeployHeight uint64
}, error) {
	return _PortalRegistry.Contract.Deployments(&_PortalRegistry.CallOpts, arg0)
}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(uint64 chainId, address addr, uint64 deployHeight)
func (_PortalRegistry *PortalRegistryCallerSession) Deployments(arg0 uint64) (struct {
	ChainId      uint64
	Addr         common.Address
	DeployHeight uint64
}, error) {
	return _PortalRegistry.Contract.Deployments(&_PortalRegistry.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 chainId) view returns((uint64,address,uint64,uint64[]))
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
// Solidity: function get(uint64 chainId) view returns((uint64,address,uint64,uint64[]))
func (_PortalRegistry *PortalRegistrySession) Get(chainId uint64) (PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.Get(&_PortalRegistry.CallOpts, chainId)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 chainId) view returns((uint64,address,uint64,uint64[]))
func (_PortalRegistry *PortalRegistryCallerSession) Get(chainId uint64) (PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.Get(&_PortalRegistry.CallOpts, chainId)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((uint64,address,uint64,uint64[])[])
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
// Solidity: function list() view returns((uint64,address,uint64,uint64[])[])
func (_PortalRegistry *PortalRegistrySession) List() ([]PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.List(&_PortalRegistry.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((uint64,address,uint64,uint64[])[])
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

// BulkRegister is a paid mutator transaction binding the contract method 0x510644e8.
//
// Solidity: function bulkRegister((uint64,address,uint64,uint64[])[] deps) payable returns()
func (_PortalRegistry *PortalRegistryTransactor) BulkRegister(opts *bind.TransactOpts, deps []PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "bulkRegister", deps)
}

// BulkRegister is a paid mutator transaction binding the contract method 0x510644e8.
//
// Solidity: function bulkRegister((uint64,address,uint64,uint64[])[] deps) payable returns()
func (_PortalRegistry *PortalRegistrySession) BulkRegister(deps []PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.BulkRegister(&_PortalRegistry.TransactOpts, deps)
}

// BulkRegister is a paid mutator transaction binding the contract method 0x510644e8.
//
// Solidity: function bulkRegister((uint64,address,uint64,uint64[])[] deps) payable returns()
func (_PortalRegistry *PortalRegistryTransactorSession) BulkRegister(deps []PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.BulkRegister(&_PortalRegistry.TransactOpts, deps)
}

// Register is a paid mutator transaction binding the contract method 0x278679c5.
//
// Solidity: function register((uint64,address,uint64,uint64[]) dep) returns()
func (_PortalRegistry *PortalRegistryTransactor) Register(opts *bind.TransactOpts, dep PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "register", dep)
}

// Register is a paid mutator transaction binding the contract method 0x278679c5.
//
// Solidity: function register((uint64,address,uint64,uint64[]) dep) returns()
func (_PortalRegistry *PortalRegistrySession) Register(dep PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.Register(&_PortalRegistry.TransactOpts, dep)
}

// Register is a paid mutator transaction binding the contract method 0x278679c5.
//
// Solidity: function register((uint64,address,uint64,uint64[]) dep) returns()
func (_PortalRegistry *PortalRegistryTransactorSession) Register(dep PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.Register(&_PortalRegistry.TransactOpts, dep)
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

// PortalRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PortalRegistry contract.
type PortalRegistryInitializedIterator struct {
	Event *PortalRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *PortalRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRegistryInitialized)
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
		it.Event = new(PortalRegistryInitialized)
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
func (it *PortalRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRegistryInitialized represents a Initialized event raised by the PortalRegistry contract.
type PortalRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PortalRegistry *PortalRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*PortalRegistryInitializedIterator, error) {

	logs, sub, err := _PortalRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PortalRegistryInitializedIterator{contract: _PortalRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PortalRegistry *PortalRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PortalRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _PortalRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRegistryInitialized)
				if err := _PortalRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PortalRegistry *PortalRegistryFilterer) ParseInitialized(log types.Log) (*PortalRegistryInitialized, error) {
	event := new(PortalRegistryInitialized)
	if err := _PortalRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	ChainId      uint64
	Addr         common.Address
	DeployHeight uint64
	Shards       []uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPortalRegistered is a free log retrieval operation binding the contract event 0x99fa7076dae8857571c277f28e10ca59528e679baa7981e731a3cd4f877b4f75.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64[] shards)
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

// WatchPortalRegistered is a free log subscription operation binding the contract event 0x99fa7076dae8857571c277f28e10ca59528e679baa7981e731a3cd4f877b4f75.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64[] shards)
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

// ParsePortalRegistered is a log parse operation binding the contract event 0x99fa7076dae8857571c277f28e10ca59528e679baa7981e731a3cd4f877b4f75.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64[] shards)
func (_PortalRegistry *PortalRegistryFilterer) ParsePortalRegistered(log types.Log) (*PortalRegistryPortalRegistered, error) {
	event := new(PortalRegistryPortalRegistered)
	if err := _PortalRegistry.contract.UnpackLog(event, "PortalRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
