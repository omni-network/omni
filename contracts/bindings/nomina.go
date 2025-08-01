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

// NominaMetaData contains all meta data concerning the Nomina contract.
var NominaMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_omni\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_mintAuthority\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONVERSION_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintAuthority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMintAuthority\",\"inputs\":[{\"name\":\"_mintAuthority\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintAuthoritySet\",\"inputs\":[{\"name\":\"mintAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterSet\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowanceOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllowanceUnderflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConversionDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Permit2AllowanceIsFixedAtInfinity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PermitExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSupplyOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051610db2380380610db283398101604081905261002e916100e0565b6001600160a01b038084166080525f80548483166001600160a01b03199182168117835560018054948616949092169390931790556040517f7cd028240c863e3069db38011d9a2a8b46b7af1d8e075414a2539f65069012fe9190a26040516001600160a01b038216907f726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6905f90a2505050610120565b80516001600160a01b03811681146100db575f5ffd5b919050565b5f5f5f606084860312156100f2575f5ffd5b6100fb846100c5565b9250610109602085016100c5565b9150610117604085016100c5565b90509250925092565b608051610c7361013f5f395f81816102aa01526105c90152610c735ff3fe608060405234801561000f575f5ffd5b5060043610610132575f3560e01c806340c10f19116100b45780639340b21e116100795780639340b21e1461034f57806395d89b4114610361578063a9059cbb14610380578063d505accf14610393578063dd62ed3e146103a6578063fca3b5aa146103b9575f5ffd5b806340c10f19146102cc57806342966c68146102df57806367c6e39c146102f257806370a08231146103055780637ecebe001461032a575f5ffd5b806323b872dd116100fa57806323b872dd146101e75780632c8bff31146101fa578063313ce567146102145780633644e5151461021b57806339acf9f1146102a5575f5ffd5b806306fdde0314610136578063075461721461016a578063095ea7b31461019557806318160ddd146101b857806323adc150146101d2575b5f5ffd5b6040805180820190915260068152654e6f6d696e6160d01b60208201525b6040516101619190610a73565b60405180910390f35b60015461017d906001600160a01b031681565b6040516001600160a01b039091168152602001610161565b6101a86101a3366004610ac3565b6103cc565b6040519015158152602001610161565b6805345cdf77eb68f44c545b604051908152602001610161565b6101e56101e0366004610aeb565b61044c565b005b6101a86101f5366004610b0b565b6104bc565b610202604b81565b60405160ff9091168152602001610161565b6012610202565b604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81527fc72733118dabad3698b4044c2dc83c8c688bd907b50ed9d09d93a263878bf51860208201527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc69181019190915246606082015230608082015260a090206101c4565b61017d7f000000000000000000000000000000000000000000000000000000000000000081565b6101e56102da366004610ac3565b610578565b6101e56102ed366004610b45565b6105b0565b6101e5610300366004610ac3565b6105c7565b6101c4610313366004610aeb565b6387a211a2600c9081525f91909152602090205490565b6101c4610338366004610aeb565b6338377508600c9081525f91909152602090205490565b5f5461017d906001600160a01b031681565b6040805180820190915260038152624e4f4d60e81b6020820152610154565b6101a861038e366004610ac3565b610673565b6101e56103a1366004610b5c565b6106d7565b6101c46103b4366004610bc9565b61089a565b6101e56103c7366004610aeb565b6108de565b5f6001600160a01b0383166e22d473030f116ddee9f6b43ac78ba318821915176103fd57633f68539a5f526004601cfd5b82602052637f5e9f20600c52335f52816034600c2055815f52602c5160601c337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560205fa35060015b92915050565b5f546001600160a01b03163314610475576040516282b42960e81b815260040160405180910390fd5b5f80546001600160a01b0319166001600160a01b038316908117825560405190917f7cd028240c863e3069db38011d9a2a8b46b7af1d8e075414a2539f65069012fe91a250565b5f8360601b6e22d473030f116ddee9f6b43ac78ba333146105115733602052637f5e9f208117600c526034600c20805480191561050e5780851115610508576313be252b5f526004601cfd5b84810382555b50505b6387a211a28117600c526020600c208054808511156105375763f4d678b85f526004601cfd5b84810382555050835f526020600c208381540181555082602052600c5160601c8160601c5f516020610c1e5f395f51905f52602080a3505060019392505050565b6001546001600160a01b031633146105a2576040516282b42960e81b815260040160405180910390fd5b6105ac8282610950565b5050565b805f036105ba5750565b6105c433826109b9565b50565b7f00000000000000000000000000000000000000000000000000000000000000005f8290036105f557505050565b6001600160a01b03831661061c5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811661064357604051632efe214b60e01b815260040160405180910390fd5b61065a6001600160a01b0382163361dead85610a1a565b61066e83610669604b85610bfa565b610950565b505050565b5f6387a211a2600c52335f526020600c2080548084111561069b5763f4d678b85f526004601cfd5b83810382555050825f526020600c208281540181555081602052600c5160601c335f516020610c1e5f395f51905f52602080a350600192915050565b6001600160a01b0386166e22d473030f116ddee9f6b43ac78ba3188519151761070757633f68539a5f526004601cfd5b7fc72733118dabad3698b4044c2dc83c8c688bd907b50ed9d09d93a263878bf5187fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64286101561075e57631a15a3cc5f526004601cfd5b6040518960601b60601c99508860601b60601c985065383775081901600e52895f526020600c2080547f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f835284602084015283604084015246606084015230608084015260a08320602e527f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c983528b60208401528a60408401528960608401528060808401528860a084015260c08320604e526042602c205f528760ff16602052866040528560605260208060805f60015afa8c3d51146108465763ddafbaef5f526004601cfd5b0190556303faf4f960a51b89176040526034602c20889055888a7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925602060608501a360405250505f60605250505050505050565b5f6e22d473030f116ddee9f6b43ac78ba2196001600160a01b038316016108c357505f19610446565b50602052637f5e9f20600c9081525f91909152603490205490565b5f546001600160a01b03163314610907576040516282b42960e81b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040517f726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6905f90a250565b6805345cdf77eb68f44c54818101818110156109735763e5cfe9575f526004601cfd5b806805345cdf77eb68f44c5550506387a211a2600c52815f526020600c208181540181555080602052600c5160601c5f5f516020610c1e5f395f51905f52602080a35050565b6387a211a2600c52815f526020600c208054808311156109e05763f4d678b85f526004601cfd5b82900390556805345cdf77eb68f44c805482900390555f8181526001600160a01b0383165f516020610c1e5f395f51905f52602083a35050565b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f511416610a6557803d873b151710610a6557637939f4245f526004601cfd5b505f60605260405250505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610abe575f5ffd5b919050565b5f5f60408385031215610ad4575f5ffd5b610add83610aa8565b946020939093013593505050565b5f60208284031215610afb575f5ffd5b610b0482610aa8565b9392505050565b5f5f5f60608486031215610b1d575f5ffd5b610b2684610aa8565b9250610b3460208501610aa8565b929592945050506040919091013590565b5f60208284031215610b55575f5ffd5b5035919050565b5f5f5f5f5f5f5f60e0888a031215610b72575f5ffd5b610b7b88610aa8565b9650610b8960208901610aa8565b95506040880135945060608801359350608088013560ff81168114610bac575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215610bda575f5ffd5b610be383610aa8565b9150610bf160208401610aa8565b90509250929050565b808202811582820484141761044657634e487b7160e01b5f52601160045260245ffdfeddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212202cbc1373f4ea754c6c3f97c23fbecb5b85267d4cbd14b76fc9e51c2e5a244a1864736f6c634300081e0033",
}

// NominaABI is the input ABI used to generate the binding from.
// Deprecated: Use NominaMetaData.ABI instead.
var NominaABI = NominaMetaData.ABI

// NominaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NominaMetaData.Bin instead.
var NominaBin = NominaMetaData.Bin

// DeployNomina deploys a new Ethereum contract, binding an instance of Nomina to it.
func DeployNomina(auth *bind.TransactOpts, backend bind.ContractBackend, _omni common.Address, _mintAuthority common.Address, _minter common.Address) (common.Address, *types.Transaction, *Nomina, error) {
	parsed, err := NominaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NominaBin), backend, _omni, _mintAuthority, _minter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nomina{NominaCaller: NominaCaller{contract: contract}, NominaTransactor: NominaTransactor{contract: contract}, NominaFilterer: NominaFilterer{contract: contract}}, nil
}

// Nomina is an auto generated Go binding around an Ethereum contract.
type Nomina struct {
	NominaCaller     // Read-only binding to the contract
	NominaTransactor // Write-only binding to the contract
	NominaFilterer   // Log filterer for contract events
}

// NominaCaller is an auto generated read-only Go binding around an Ethereum contract.
type NominaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NominaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NominaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NominaSession struct {
	Contract     *Nomina           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NominaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NominaCallerSession struct {
	Contract *NominaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NominaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NominaTransactorSession struct {
	Contract     *NominaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NominaRaw is an auto generated low-level Go binding around an Ethereum contract.
type NominaRaw struct {
	Contract *Nomina // Generic contract binding to access the raw methods on
}

// NominaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NominaCallerRaw struct {
	Contract *NominaCaller // Generic read-only contract binding to access the raw methods on
}

// NominaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NominaTransactorRaw struct {
	Contract *NominaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNomina creates a new instance of Nomina, bound to a specific deployed contract.
func NewNomina(address common.Address, backend bind.ContractBackend) (*Nomina, error) {
	contract, err := bindNomina(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nomina{NominaCaller: NominaCaller{contract: contract}, NominaTransactor: NominaTransactor{contract: contract}, NominaFilterer: NominaFilterer{contract: contract}}, nil
}

// NewNominaCaller creates a new read-only instance of Nomina, bound to a specific deployed contract.
func NewNominaCaller(address common.Address, caller bind.ContractCaller) (*NominaCaller, error) {
	contract, err := bindNomina(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NominaCaller{contract: contract}, nil
}

// NewNominaTransactor creates a new write-only instance of Nomina, bound to a specific deployed contract.
func NewNominaTransactor(address common.Address, transactor bind.ContractTransactor) (*NominaTransactor, error) {
	contract, err := bindNomina(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NominaTransactor{contract: contract}, nil
}

// NewNominaFilterer creates a new log filterer instance of Nomina, bound to a specific deployed contract.
func NewNominaFilterer(address common.Address, filterer bind.ContractFilterer) (*NominaFilterer, error) {
	contract, err := bindNomina(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NominaFilterer{contract: contract}, nil
}

// bindNomina binds a generic wrapper to an already deployed contract.
func bindNomina(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NominaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nomina *NominaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nomina.Contract.NominaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nomina *NominaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nomina.Contract.NominaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nomina *NominaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nomina.Contract.NominaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nomina *NominaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nomina.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nomina *NominaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nomina.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nomina *NominaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nomina.Contract.contract.Transact(opts, method, params...)
}

// CONVERSIONRATE is a free data retrieval call binding the contract method 0x2c8bff31.
//
// Solidity: function CONVERSION_RATE() view returns(uint8)
func (_Nomina *NominaCaller) CONVERSIONRATE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "CONVERSION_RATE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CONVERSIONRATE is a free data retrieval call binding the contract method 0x2c8bff31.
//
// Solidity: function CONVERSION_RATE() view returns(uint8)
func (_Nomina *NominaSession) CONVERSIONRATE() (uint8, error) {
	return _Nomina.Contract.CONVERSIONRATE(&_Nomina.CallOpts)
}

// CONVERSIONRATE is a free data retrieval call binding the contract method 0x2c8bff31.
//
// Solidity: function CONVERSION_RATE() view returns(uint8)
func (_Nomina *NominaCallerSession) CONVERSIONRATE() (uint8, error) {
	return _Nomina.Contract.CONVERSIONRATE(&_Nomina.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_Nomina *NominaCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_Nomina *NominaSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Nomina.Contract.DOMAINSEPARATOR(&_Nomina.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_Nomina *NominaCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Nomina.Contract.DOMAINSEPARATOR(&_Nomina.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_Nomina *NominaCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_Nomina *NominaSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Nomina.Contract.Allowance(&_Nomina.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_Nomina *NominaCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Nomina.Contract.Allowance(&_Nomina.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_Nomina *NominaCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_Nomina *NominaSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nomina.Contract.BalanceOf(&_Nomina.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_Nomina *NominaCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nomina.Contract.BalanceOf(&_Nomina.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nomina *NominaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nomina *NominaSession) Decimals() (uint8, error) {
	return _Nomina.Contract.Decimals(&_Nomina.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nomina *NominaCallerSession) Decimals() (uint8, error) {
	return _Nomina.Contract.Decimals(&_Nomina.CallOpts)
}

// MintAuthority is a free data retrieval call binding the contract method 0x9340b21e.
//
// Solidity: function mintAuthority() view returns(address)
func (_Nomina *NominaCaller) MintAuthority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "mintAuthority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintAuthority is a free data retrieval call binding the contract method 0x9340b21e.
//
// Solidity: function mintAuthority() view returns(address)
func (_Nomina *NominaSession) MintAuthority() (common.Address, error) {
	return _Nomina.Contract.MintAuthority(&_Nomina.CallOpts)
}

// MintAuthority is a free data retrieval call binding the contract method 0x9340b21e.
//
// Solidity: function mintAuthority() view returns(address)
func (_Nomina *NominaCallerSession) MintAuthority() (common.Address, error) {
	return _Nomina.Contract.MintAuthority(&_Nomina.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Nomina *NominaCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Nomina *NominaSession) Minter() (common.Address, error) {
	return _Nomina.Contract.Minter(&_Nomina.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Nomina *NominaCallerSession) Minter() (common.Address, error) {
	return _Nomina.Contract.Minter(&_Nomina.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Nomina *NominaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Nomina *NominaSession) Name() (string, error) {
	return _Nomina.Contract.Name(&_Nomina.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Nomina *NominaCallerSession) Name() (string, error) {
	return _Nomina.Contract.Name(&_Nomina.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_Nomina *NominaCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_Nomina *NominaSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Nomina.Contract.Nonces(&_Nomina.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_Nomina *NominaCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Nomina.Contract.Nonces(&_Nomina.CallOpts, owner)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_Nomina *NominaCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_Nomina *NominaSession) Omni() (common.Address, error) {
	return _Nomina.Contract.Omni(&_Nomina.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_Nomina *NominaCallerSession) Omni() (common.Address, error) {
	return _Nomina.Contract.Omni(&_Nomina.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Nomina *NominaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Nomina *NominaSession) Symbol() (string, error) {
	return _Nomina.Contract.Symbol(&_Nomina.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Nomina *NominaCallerSession) Symbol() (string, error) {
	return _Nomina.Contract.Symbol(&_Nomina.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Nomina *NominaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nomina.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Nomina *NominaSession) TotalSupply() (*big.Int, error) {
	return _Nomina.Contract.TotalSupply(&_Nomina.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Nomina *NominaCallerSession) TotalSupply() (*big.Int, error) {
	return _Nomina.Contract.TotalSupply(&_Nomina.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nomina *NominaTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nomina *NominaSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Approve(&_Nomina.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nomina *NominaTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Approve(&_Nomina.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Nomina *NominaTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Nomina *NominaSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Burn(&_Nomina.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Nomina *NominaTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Burn(&_Nomina.TransactOpts, amount)
}

// Convert is a paid mutator transaction binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address to, uint256 amount) returns()
func (_Nomina *NominaTransactor) Convert(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "convert", to, amount)
}

// Convert is a paid mutator transaction binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address to, uint256 amount) returns()
func (_Nomina *NominaSession) Convert(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Convert(&_Nomina.TransactOpts, to, amount)
}

// Convert is a paid mutator transaction binding the contract method 0x67c6e39c.
//
// Solidity: function convert(address to, uint256 amount) returns()
func (_Nomina *NominaTransactorSession) Convert(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Convert(&_Nomina.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Nomina *NominaTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Nomina *NominaSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Mint(&_Nomina.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Nomina *NominaTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Mint(&_Nomina.TransactOpts, to, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Nomina *NominaTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Nomina *NominaSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Nomina.Contract.Permit(&_Nomina.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Nomina *NominaTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Nomina.Contract.Permit(&_Nomina.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SetMintAuthority is a paid mutator transaction binding the contract method 0x23adc150.
//
// Solidity: function setMintAuthority(address _mintAuthority) returns()
func (_Nomina *NominaTransactor) SetMintAuthority(opts *bind.TransactOpts, _mintAuthority common.Address) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "setMintAuthority", _mintAuthority)
}

// SetMintAuthority is a paid mutator transaction binding the contract method 0x23adc150.
//
// Solidity: function setMintAuthority(address _mintAuthority) returns()
func (_Nomina *NominaSession) SetMintAuthority(_mintAuthority common.Address) (*types.Transaction, error) {
	return _Nomina.Contract.SetMintAuthority(&_Nomina.TransactOpts, _mintAuthority)
}

// SetMintAuthority is a paid mutator transaction binding the contract method 0x23adc150.
//
// Solidity: function setMintAuthority(address _mintAuthority) returns()
func (_Nomina *NominaTransactorSession) SetMintAuthority(_mintAuthority common.Address) (*types.Transaction, error) {
	return _Nomina.Contract.SetMintAuthority(&_Nomina.TransactOpts, _mintAuthority)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Nomina *NominaTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Nomina *NominaSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Nomina.Contract.SetMinter(&_Nomina.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_Nomina *NominaTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Nomina.Contract.SetMinter(&_Nomina.TransactOpts, _minter)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nomina *NominaTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nomina *NominaSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Transfer(&_Nomina.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nomina *NominaTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.Transfer(&_Nomina.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nomina *NominaTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nomina *NominaSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.TransferFrom(&_Nomina.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nomina *NominaTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nomina.Contract.TransferFrom(&_Nomina.TransactOpts, from, to, amount)
}

// NominaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nomina contract.
type NominaApprovalIterator struct {
	Event *NominaApproval // Event containing the contract specifics and raw log

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
func (it *NominaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaApproval)
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
		it.Event = new(NominaApproval)
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
func (it *NominaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaApproval represents a Approval event raised by the Nomina contract.
type NominaApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Nomina *NominaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NominaApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nomina.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NominaApprovalIterator{contract: _Nomina.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Nomina *NominaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NominaApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nomina.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaApproval)
				if err := _Nomina.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Nomina *NominaFilterer) ParseApproval(log types.Log) (*NominaApproval, error) {
	event := new(NominaApproval)
	if err := _Nomina.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaMintAuthoritySetIterator is returned from FilterMintAuthoritySet and is used to iterate over the raw logs and unpacked data for MintAuthoritySet events raised by the Nomina contract.
type NominaMintAuthoritySetIterator struct {
	Event *NominaMintAuthoritySet // Event containing the contract specifics and raw log

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
func (it *NominaMintAuthoritySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaMintAuthoritySet)
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
		it.Event = new(NominaMintAuthoritySet)
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
func (it *NominaMintAuthoritySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaMintAuthoritySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaMintAuthoritySet represents a MintAuthoritySet event raised by the Nomina contract.
type NominaMintAuthoritySet struct {
	MintAuthority common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMintAuthoritySet is a free log retrieval operation binding the contract event 0x7cd028240c863e3069db38011d9a2a8b46b7af1d8e075414a2539f65069012fe.
//
// Solidity: event MintAuthoritySet(address indexed mintAuthority)
func (_Nomina *NominaFilterer) FilterMintAuthoritySet(opts *bind.FilterOpts, mintAuthority []common.Address) (*NominaMintAuthoritySetIterator, error) {

	var mintAuthorityRule []interface{}
	for _, mintAuthorityItem := range mintAuthority {
		mintAuthorityRule = append(mintAuthorityRule, mintAuthorityItem)
	}

	logs, sub, err := _Nomina.contract.FilterLogs(opts, "MintAuthoritySet", mintAuthorityRule)
	if err != nil {
		return nil, err
	}
	return &NominaMintAuthoritySetIterator{contract: _Nomina.contract, event: "MintAuthoritySet", logs: logs, sub: sub}, nil
}

// WatchMintAuthoritySet is a free log subscription operation binding the contract event 0x7cd028240c863e3069db38011d9a2a8b46b7af1d8e075414a2539f65069012fe.
//
// Solidity: event MintAuthoritySet(address indexed mintAuthority)
func (_Nomina *NominaFilterer) WatchMintAuthoritySet(opts *bind.WatchOpts, sink chan<- *NominaMintAuthoritySet, mintAuthority []common.Address) (event.Subscription, error) {

	var mintAuthorityRule []interface{}
	for _, mintAuthorityItem := range mintAuthority {
		mintAuthorityRule = append(mintAuthorityRule, mintAuthorityItem)
	}

	logs, sub, err := _Nomina.contract.WatchLogs(opts, "MintAuthoritySet", mintAuthorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaMintAuthoritySet)
				if err := _Nomina.contract.UnpackLog(event, "MintAuthoritySet", log); err != nil {
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

// ParseMintAuthoritySet is a log parse operation binding the contract event 0x7cd028240c863e3069db38011d9a2a8b46b7af1d8e075414a2539f65069012fe.
//
// Solidity: event MintAuthoritySet(address indexed mintAuthority)
func (_Nomina *NominaFilterer) ParseMintAuthoritySet(log types.Log) (*NominaMintAuthoritySet, error) {
	event := new(NominaMintAuthoritySet)
	if err := _Nomina.contract.UnpackLog(event, "MintAuthoritySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaMinterSetIterator is returned from FilterMinterSet and is used to iterate over the raw logs and unpacked data for MinterSet events raised by the Nomina contract.
type NominaMinterSetIterator struct {
	Event *NominaMinterSet // Event containing the contract specifics and raw log

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
func (it *NominaMinterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaMinterSet)
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
		it.Event = new(NominaMinterSet)
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
func (it *NominaMinterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaMinterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaMinterSet represents a MinterSet event raised by the Nomina contract.
type NominaMinterSet struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterSet is a free log retrieval operation binding the contract event 0x726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6.
//
// Solidity: event MinterSet(address indexed minter)
func (_Nomina *NominaFilterer) FilterMinterSet(opts *bind.FilterOpts, minter []common.Address) (*NominaMinterSetIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Nomina.contract.FilterLogs(opts, "MinterSet", minterRule)
	if err != nil {
		return nil, err
	}
	return &NominaMinterSetIterator{contract: _Nomina.contract, event: "MinterSet", logs: logs, sub: sub}, nil
}

// WatchMinterSet is a free log subscription operation binding the contract event 0x726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6.
//
// Solidity: event MinterSet(address indexed minter)
func (_Nomina *NominaFilterer) WatchMinterSet(opts *bind.WatchOpts, sink chan<- *NominaMinterSet, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Nomina.contract.WatchLogs(opts, "MinterSet", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaMinterSet)
				if err := _Nomina.contract.UnpackLog(event, "MinterSet", log); err != nil {
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

// ParseMinterSet is a log parse operation binding the contract event 0x726b590ef91a8c76ad05bbe91a57ef84605276528f49cd47d787f558a4e755b6.
//
// Solidity: event MinterSet(address indexed minter)
func (_Nomina *NominaFilterer) ParseMinterSet(log types.Log) (*NominaMinterSet, error) {
	event := new(NominaMinterSet)
	if err := _Nomina.contract.UnpackLog(event, "MinterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nomina contract.
type NominaTransferIterator struct {
	Event *NominaTransfer // Event containing the contract specifics and raw log

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
func (it *NominaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaTransfer)
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
		it.Event = new(NominaTransfer)
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
func (it *NominaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaTransfer represents a Transfer event raised by the Nomina contract.
type NominaTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Nomina *NominaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NominaTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nomina.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NominaTransferIterator{contract: _Nomina.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Nomina *NominaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NominaTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nomina.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaTransfer)
				if err := _Nomina.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Nomina *NominaFilterer) ParseTransfer(log types.Log) (*NominaTransfer, error) {
	event := new(NominaTransfer)
	if err := _Nomina.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
