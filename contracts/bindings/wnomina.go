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

// WNominaMetaData contains all meta data concerning the WNomina contract.
var WNominaMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositTo\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTo\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowanceOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllowanceUnderflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NOMTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Permit2AllowanceIsFixedAtInfinity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PermitExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSupplyOverflow\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610a238061001c5f395ff3fe6080604052600436106100f6575f3560e01c806370a0823111610089578063b760faf911610058578063b760faf914610365578063d0e30db014610105578063d505accf14610378578063dd62ed3e1461039757610105565b806370a08231146102b85780637ecebe00146102e957806395d89b411461031a578063a9059cbb1461034657610105565b806323b872dd116100c557806323b872dd146101c95780632e1a7d4d146101e8578063313ce567146102075780633644e5151461022257610105565b806306fdde031461010d578063095ea7b31461015557806318160ddd14610184578063205c2878146101aa57610105565b36610105576101036103b6565b005b6101036103b6565b348015610118575f5ffd5b5060408051808201909152600e81526d57726170706564204e6f6d696e6160901b60208201525b60405161014c919061087b565b60405180910390f35b348015610160575f5ffd5b5061017461016f3660046108cb565b6103c2565b604051901515815260200161014c565b34801561018f575f5ffd5b506805345cdf77eb68f44c545b60405190815260200161014c565b3480156101b5575f5ffd5b506101036101c43660046108cb565b610442565b3480156101d4575f5ffd5b506101746101e33660046108f3565b610450565b3480156101f3575f5ffd5b5061010361020236600461092d565b61050c565b348015610212575f5ffd5b506040516012815260200161014c565b34801561022d575f5ffd5b50604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81527f84fe092f00820265d8abbb39159a80f1dc751d5463fcd172fce598140f7b360c60208201527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc69181019190915246606082015230608082015260a0902061019c565b3480156102c3575f5ffd5b5061019c6102d2366004610944565b6387a211a2600c9081525f91909152602090205490565b3480156102f4575f5ffd5b5061019c610303366004610944565b6338377508600c9081525f91909152602090205490565b348015610325575f5ffd5b50604080518082019091526004815263574e4f4d60e01b602082015261013f565b348015610351575f5ffd5b506101746103603660046108cb565b610519565b610103610373366004610944565b61057d565b348015610383575f5ffd5b50610103610392366004610964565b610587565b3480156103a2575f5ffd5b5061019c6103b13660046109d1565b61074a565b6103c0333461078e565b565b5f6001600160a01b0383166e22d473030f116ddee9f6b43ac78ba318821915176103f357633f68539a5f526004601cfd5b82602052637f5e9f20600c52335f52816034600c2055815f52602c5160601c337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560205fa35060015b92915050565b61044c82826107f7565b5050565b5f8360601b6e22d473030f116ddee9f6b43ac78ba333146104a55733602052637f5e9f208117600c526034600c2080548019156104a2578085111561049c576313be252b5f526004601cfd5b84810382555b50505b6387a211a28117600c526020600c208054808511156104cb5763f4d678b85f526004601cfd5b84810382555050835f526020600c208381540181555082602052600c5160601c8160601c5f516020610a035f395f51905f52602080a3505060019392505050565b61051633826107f7565b50565b5f6387a211a2600c52335f526020600c208054808411156105415763f4d678b85f526004601cfd5b83810382555050825f526020600c208281540181555081602052600c5160601c335f516020610a035f395f51905f52602080a350600192915050565b610516813461078e565b6001600160a01b0386166e22d473030f116ddee9f6b43ac78ba318851915176105b757633f68539a5f526004601cfd5b7f84fe092f00820265d8abbb39159a80f1dc751d5463fcd172fce598140f7b360c7fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64286101561060e57631a15a3cc5f526004601cfd5b6040518960601b60601c99508860601b60601c985065383775081901600e52895f526020600c2080547f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f835284602084015283604084015246606084015230608084015260a08320602e527f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c983528b60208401528a60408401528960608401528060808401528860a084015260c08320604e526042602c205f528760ff16602052866040528560605260208060805f60015afa8c3d51146106f65763ddafbaef5f526004601cfd5b0190556303faf4f960a51b89176040526034602c20889055888a7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925602060608501a360405250505f60605250505050505050565b5f6e22d473030f116ddee9f6b43ac78ba2196001600160a01b0383160161077357505f1961043c565b50602052637f5e9f20600c9081525f91909152603490205490565b6805345cdf77eb68f44c54818101818110156107b15763e5cfe9575f526004601cfd5b806805345cdf77eb68f44c5550506387a211a2600c52815f526020600c208181540181555080602052600c5160601c5f5f516020610a035f395f51905f52602080a35050565b610801338261081a565b5f385f3884865af161044c5763252db5955f526004601cfd5b6387a211a2600c52815f526020600c208054808311156108415763f4d678b85f526004601cfd5b82900390556805345cdf77eb68f44c805482900390555f8181526001600160a01b0383165f516020610a035f395f51905f52602083a35050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b03811681146108c6575f5ffd5b919050565b5f5f604083850312156108dc575f5ffd5b6108e5836108b0565b946020939093013593505050565b5f5f5f60608486031215610905575f5ffd5b61090e846108b0565b925061091c602085016108b0565b929592945050506040919091013590565b5f6020828403121561093d575f5ffd5b5035919050565b5f60208284031215610954575f5ffd5b61095d826108b0565b9392505050565b5f5f5f5f5f5f5f60e0888a03121561097a575f5ffd5b610983886108b0565b9650610991602089016108b0565b95506040880135945060608801359350608088013560ff811681146109b4575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156109e2575f5ffd5b6109eb836108b0565b91506109f9602084016108b0565b9050925092905056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
}

// WNominaABI is the input ABI used to generate the binding from.
// Deprecated: Use WNominaMetaData.ABI instead.
var WNominaABI = WNominaMetaData.ABI

// WNominaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WNominaMetaData.Bin instead.
var WNominaBin = WNominaMetaData.Bin

// DeployWNomina deploys a new Ethereum contract, binding an instance of WNomina to it.
func DeployWNomina(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WNomina, error) {
	parsed, err := WNominaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WNominaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WNomina{WNominaCaller: WNominaCaller{contract: contract}, WNominaTransactor: WNominaTransactor{contract: contract}, WNominaFilterer: WNominaFilterer{contract: contract}}, nil
}

// WNomina is an auto generated Go binding around an Ethereum contract.
type WNomina struct {
	WNominaCaller     // Read-only binding to the contract
	WNominaTransactor // Write-only binding to the contract
	WNominaFilterer   // Log filterer for contract events
}

// WNominaCaller is an auto generated read-only Go binding around an Ethereum contract.
type WNominaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WNominaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WNominaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WNominaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WNominaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WNominaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WNominaSession struct {
	Contract     *WNomina          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WNominaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WNominaCallerSession struct {
	Contract *WNominaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WNominaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WNominaTransactorSession struct {
	Contract     *WNominaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WNominaRaw is an auto generated low-level Go binding around an Ethereum contract.
type WNominaRaw struct {
	Contract *WNomina // Generic contract binding to access the raw methods on
}

// WNominaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WNominaCallerRaw struct {
	Contract *WNominaCaller // Generic read-only contract binding to access the raw methods on
}

// WNominaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WNominaTransactorRaw struct {
	Contract *WNominaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWNomina creates a new instance of WNomina, bound to a specific deployed contract.
func NewWNomina(address common.Address, backend bind.ContractBackend) (*WNomina, error) {
	contract, err := bindWNomina(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WNomina{WNominaCaller: WNominaCaller{contract: contract}, WNominaTransactor: WNominaTransactor{contract: contract}, WNominaFilterer: WNominaFilterer{contract: contract}}, nil
}

// NewWNominaCaller creates a new read-only instance of WNomina, bound to a specific deployed contract.
func NewWNominaCaller(address common.Address, caller bind.ContractCaller) (*WNominaCaller, error) {
	contract, err := bindWNomina(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WNominaCaller{contract: contract}, nil
}

// NewWNominaTransactor creates a new write-only instance of WNomina, bound to a specific deployed contract.
func NewWNominaTransactor(address common.Address, transactor bind.ContractTransactor) (*WNominaTransactor, error) {
	contract, err := bindWNomina(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WNominaTransactor{contract: contract}, nil
}

// NewWNominaFilterer creates a new log filterer instance of WNomina, bound to a specific deployed contract.
func NewWNominaFilterer(address common.Address, filterer bind.ContractFilterer) (*WNominaFilterer, error) {
	contract, err := bindWNomina(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WNominaFilterer{contract: contract}, nil
}

// bindWNomina binds a generic wrapper to an already deployed contract.
func bindWNomina(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WNominaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WNomina *WNominaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WNomina.Contract.WNominaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WNomina *WNominaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNomina.Contract.WNominaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WNomina *WNominaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WNomina.Contract.WNominaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WNomina *WNominaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WNomina.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WNomina *WNominaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNomina.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WNomina *WNominaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WNomina.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_WNomina *WNominaCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WNomina.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_WNomina *WNominaSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WNomina.Contract.DOMAINSEPARATOR(&_WNomina.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_WNomina *WNominaCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WNomina.Contract.DOMAINSEPARATOR(&_WNomina.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_WNomina *WNominaCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNomina.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_WNomina *WNominaSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WNomina.Contract.Allowance(&_WNomina.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_WNomina *WNominaCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WNomina.Contract.Allowance(&_WNomina.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_WNomina *WNominaCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNomina.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_WNomina *WNominaSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WNomina.Contract.BalanceOf(&_WNomina.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_WNomina *WNominaCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WNomina.Contract.BalanceOf(&_WNomina.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WNomina *WNominaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WNomina.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WNomina *WNominaSession) Decimals() (uint8, error) {
	return _WNomina.Contract.Decimals(&_WNomina.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WNomina *WNominaCallerSession) Decimals() (uint8, error) {
	return _WNomina.Contract.Decimals(&_WNomina.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_WNomina *WNominaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WNomina.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_WNomina *WNominaSession) Name() (string, error) {
	return _WNomina.Contract.Name(&_WNomina.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_WNomina *WNominaCallerSession) Name() (string, error) {
	return _WNomina.Contract.Name(&_WNomina.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_WNomina *WNominaCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNomina.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_WNomina *WNominaSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WNomina.Contract.Nonces(&_WNomina.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_WNomina *WNominaCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WNomina.Contract.Nonces(&_WNomina.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_WNomina *WNominaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WNomina.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_WNomina *WNominaSession) Symbol() (string, error) {
	return _WNomina.Contract.Symbol(&_WNomina.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_WNomina *WNominaCallerSession) Symbol() (string, error) {
	return _WNomina.Contract.Symbol(&_WNomina.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_WNomina *WNominaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WNomina.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_WNomina *WNominaSession) TotalSupply() (*big.Int, error) {
	return _WNomina.Contract.TotalSupply(&_WNomina.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_WNomina *WNominaCallerSession) TotalSupply() (*big.Int, error) {
	return _WNomina.Contract.TotalSupply(&_WNomina.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WNomina *WNominaTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WNomina *WNominaSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.Approve(&_WNomina.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WNomina *WNominaTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.Approve(&_WNomina.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WNomina *WNominaTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNomina.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WNomina *WNominaSession) Deposit() (*types.Transaction, error) {
	return _WNomina.Contract.Deposit(&_WNomina.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WNomina *WNominaTransactorSession) Deposit() (*types.Transaction, error) {
	return _WNomina.Contract.Deposit(&_WNomina.TransactOpts)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_WNomina *WNominaTransactor) DepositTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WNomina.contract.Transact(opts, "depositTo", to)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_WNomina *WNominaSession) DepositTo(to common.Address) (*types.Transaction, error) {
	return _WNomina.Contract.DepositTo(&_WNomina.TransactOpts, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address to) payable returns()
func (_WNomina *WNominaTransactorSession) DepositTo(to common.Address) (*types.Transaction, error) {
	return _WNomina.Contract.DepositTo(&_WNomina.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WNomina *WNominaTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WNomina.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WNomina *WNominaSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WNomina.Contract.Permit(&_WNomina.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WNomina *WNominaTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WNomina.Contract.Permit(&_WNomina.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WNomina *WNominaTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WNomina *WNominaSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.Transfer(&_WNomina.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WNomina *WNominaTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.Transfer(&_WNomina.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WNomina *WNominaTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WNomina *WNominaSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.TransferFrom(&_WNomina.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WNomina *WNominaTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.TransferFrom(&_WNomina.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WNomina *WNominaTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WNomina *WNominaSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.Withdraw(&_WNomina.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WNomina *WNominaTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.Withdraw(&_WNomina.TransactOpts, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_WNomina *WNominaTransactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.contract.Transact(opts, "withdrawTo", to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_WNomina *WNominaSession) WithdrawTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.WithdrawTo(&_WNomina.TransactOpts, to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_WNomina *WNominaTransactorSession) WithdrawTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNomina.Contract.WithdrawTo(&_WNomina.TransactOpts, to, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WNomina *WNominaTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WNomina.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WNomina *WNominaSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WNomina.Contract.Fallback(&_WNomina.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WNomina *WNominaTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WNomina.Contract.Fallback(&_WNomina.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WNomina *WNominaTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNomina.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WNomina *WNominaSession) Receive() (*types.Transaction, error) {
	return _WNomina.Contract.Receive(&_WNomina.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WNomina *WNominaTransactorSession) Receive() (*types.Transaction, error) {
	return _WNomina.Contract.Receive(&_WNomina.TransactOpts)
}

// WNominaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WNomina contract.
type WNominaApprovalIterator struct {
	Event *WNominaApproval // Event containing the contract specifics and raw log

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
func (it *WNominaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WNominaApproval)
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
		it.Event = new(WNominaApproval)
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
func (it *WNominaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WNominaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WNominaApproval represents a Approval event raised by the WNomina contract.
type WNominaApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_WNomina *WNominaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WNominaApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WNomina.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WNominaApprovalIterator{contract: _WNomina.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_WNomina *WNominaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WNominaApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WNomina.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WNominaApproval)
				if err := _WNomina.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WNomina *WNominaFilterer) ParseApproval(log types.Log) (*WNominaApproval, error) {
	event := new(WNominaApproval)
	if err := _WNomina.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WNominaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WNomina contract.
type WNominaTransferIterator struct {
	Event *WNominaTransfer // Event containing the contract specifics and raw log

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
func (it *WNominaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WNominaTransfer)
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
		it.Event = new(WNominaTransfer)
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
func (it *WNominaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WNominaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WNominaTransfer represents a Transfer event raised by the WNomina contract.
type WNominaTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_WNomina *WNominaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WNominaTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WNomina.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WNominaTransferIterator{contract: _WNomina.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_WNomina *WNominaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WNominaTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WNomina.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WNominaTransfer)
				if err := _WNomina.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WNomina *WNominaFilterer) ParseTransfer(log types.Log) (*WNominaTransfer, error) {
	event := new(WNominaTransfer)
	if err := _WNomina.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
