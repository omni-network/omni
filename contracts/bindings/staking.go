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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MinDelegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MinDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowValidators\",\"inputs\":[{\"name\":\"validators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createValidator\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"disableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowValidators\",\"inputs\":[{\"name\":\"validators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowlistEnabled_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAllowedValidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowlistDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowlistEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"deposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delegate\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAllowed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDisallowed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611097806100df6000396000f3fe6080604052600436106100dd5760003560e01c806384768b7a1161007f578063c6a2aac811610059578063c6a2aac814610252578063cf8e629a14610267578063d146fd1b1461027c578063f2fde38b1461029657600080fd5b806384768b7a146101b85780638da5cb5b146101f8578063a5a470ad1461023f57600080fd5b8063400ada75116100bb578063400ada751461015457806359bcddde146101745780635c19a95c14610190578063715018a6146101a357600080fd5b8063117407e3146100e257806311bcd830146101045780633f0b1edf14610134575b600080fd5b3480156100ee57600080fd5b506101026100fd366004610e50565b6102b6565b005b34801561011057600080fd5b5061012168056bc75e2d6310000081565b6040519081526020015b60405180910390f35b34801561014057600080fd5b5061010261014f366004610e50565b610386565b34801561016057600080fd5b5061010261016f366004610ee1565b610452565b34801561018057600080fd5b50610121670de0b6b3a764000081565b61010261019e366004610f1d565b610571565b3480156101af57600080fd5b506101026106df565b3480156101c457600080fd5b506101e86101d3366004610f1d565b60016020526000908152604090205460ff1681565b604051901515815260200161012b565b34801561020457600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b03909116815260200161012b565b61010261024d366004610f38565b6106f3565b34801561025e57600080fd5b5061010261084f565b34801561027357600080fd5b5061010261088d565b34801561028857600080fd5b506000546101e89060ff1681565b3480156102a257600080fd5b506101026102b1366004610f1d565b6108c8565b6102be610906565b60005b818110156103815760018060008585858181106102e0576102e0610f98565b90506020020160208101906102f59190610f1d565b6001600160a01b031681526020810191909152604001600020805460ff191691151591909117905582828281811061032f5761032f610f98565b90506020020160208101906103449190610f1d565b6001600160a01b03167fc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de9860405160405180910390a26001016102c1565b505050565b61038e610906565b60005b81811015610381576000600160008585858181106103b1576103b1610f98565b90506020020160208101906103c69190610f1d565b6001600160a01b031681526020810191909152604001600020805460ff191691151591909117905582828281811061040057610400610f98565b90506020020160208101906104159190610f1d565b6001600160a01b03167f3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f60405160405180910390a2600101610391565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156104985750825b905060008267ffffffffffffffff1660011480156104b55750303b155b9050811580156104c3575080155b156104e15760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561050b57845460ff60401b1916600160401b1785555b61051487610961565b6000805460ff1916871515179055831561056857845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b60005460ff16158061059b57506001600160a01b03811660009081526001602052604090205460ff165b6105ec5760405162461bcd60e51b815260206004820152601860248201527f5374616b696e673a206e6f7420616c6c6f7765642076616c000000000000000060448201526064015b60405180910390fd5b670de0b6b3a76400003410156106445760405162461bcd60e51b815260206004820152601d60248201527f5374616b696e673a20696e73756666696369656e74206465706f73697400000060448201526064016105e3565b336001600160a01b0382161461069c5760405162461bcd60e51b815260206004820152601d60248201527f5374616b696e673a206f6e6c792073656c662064656c65676174696f6e00000060448201526064016105e3565b6040513481526001600160a01b0382169033907f510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc9060200160405180910390a350565b6106e7610906565b6106f16000610972565b565b60005460ff16158061071457503360009081526001602052604090205460ff165b6107575760405162461bcd60e51b815260206004820152601460248201527314dd185ada5b99ce881b9bdd08185b1b1bddd95960621b60448201526064016105e3565b68056bc75e2d631000003410156107b05760405162461bcd60e51b815260206004820152601d60248201527f5374616b696e673a20696e73756666696369656e74206465706f73697400000060448201526064016105e3565b6107ba82826109e3565b6108065760405162461bcd60e51b815260206004820152601760248201527f5374616b696e673a20696e76616c6964207075626b657900000000000000000060448201526064016105e3565b336001600160a01b03167fc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a45383833460405161084393929190610fae565b60405180910390a25050565b610857610906565b6000805460ff191660011781556040517f8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec39190a1565b610895610906565b6000805460ff191681556040517f2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb079190a1565b6108d0610906565b6001600160a01b0381166108fa57604051631e4fbdf760e01b8152600060048201526024016105e3565b61090381610972565b50565b336109387f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146106f15760405163118cdaa760e01b81523360048201526024016105e3565b610969610b39565b61090381610b82565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b600060218214610a355760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a20696e76616c6964207075626b6579206c656e677468000060448201526064016105e3565b82826000818110610a4857610a48610f98565b9050013560f81c60f81b6001600160f81b031916600260f81b1480610a96575082826000818110610a7b57610a7b610f98565b9050013560f81c60f81b6001600160f81b031916600360f81b145b610ae25760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a20696e76616c6964207075626b657920707265666978000060448201526064016105e3565b60018301356000610b1785858381610afc57610afc610f98565b919091013560f81c905083600060076401000003d019610b8a565b9050610b2e8282600060076401000003d019610cbe565b925050505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166106f157604051631afcd79f60e31b815260040160405180910390fd5b6108d0610b39565b60008560ff1660021480610ba157508560ff166003145b610c075760405162461bcd60e51b815260206004820152603160248201527f456c6c697074696343757276653a696e6e76616c696420636f6d7072657373656044820152700c8408a8640e0ded2dce840e0e4caccd2f607b1b60648201526084016105e3565b60008280610c1757610c17610fe7565b8380610c2557610c25610fe7565b858580610c3457610c34610fe7565b888a09088480610c4657610c46610fe7565b8580610c5457610c54610fe7565b898a098909089050610c7d816004610c6d866001611013565b610c779190611026565b85610d77565b905060006002610c9060ff8a1684611013565b610c9a919061103a565b15610cae57610ca9828561104e565b610cb0565b815b925050505b95945050505050565b6000851580610ccd5750818610155b80610cd6575084155b80610ce15750818510155b15610cee57506000610cb5565b60008280610cfe57610cfe610fe7565b868709905060008380610d1357610d13610fe7565b888580610d2257610d22610fe7565b8a8b090990508515610d52578380610d3c57610d3c610fe7565b8480610d4a57610d4a610fe7565b878a09820890505b8415610d6c578380610d6657610d66610fe7565b85820890505b149695505050505050565b600081600003610dc95760405162461bcd60e51b815260206004820152601e60248201527f456c6c697074696343757276653a206d6f64756c7573206973207a65726f000060448201526064016105e3565b83600003610dd957506000610e49565b82600003610de957506001610e49565b6001600160ff1b5b8015610e4557838186161515870a85848509099150836002820486161515870a85848509099150836004820486161515870a85848509099150836008820486161515870a8584850909915060109004610df1565b5090505b9392505050565b60008060208385031215610e6357600080fd5b823567ffffffffffffffff80821115610e7b57600080fd5b818501915085601f830112610e8f57600080fd5b813581811115610e9e57600080fd5b8660208260051b8501011115610eb357600080fd5b60209290920196919550909350505050565b80356001600160a01b0381168114610edc57600080fd5b919050565b60008060408385031215610ef457600080fd5b610efd83610ec5565b915060208301358015158114610f1257600080fd5b809150509250929050565b600060208284031215610f2f57600080fd5b610e4982610ec5565b60008060208385031215610f4b57600080fd5b823567ffffffffffffffff80821115610f6357600080fd5b818501915085601f830112610f7757600080fd5b813581811115610f8657600080fd5b866020828501011115610eb357600080fd5b634e487b7160e01b600052603260045260246000fd5b604081528260408201528284606083013760006060848301015260006060601f19601f8601168301019050826020830152949350505050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610b3357610b33610ffd565b60008261103557611035610fe7565b500490565b60008261104957611049610fe7565b500690565b81810381811115610b3357610b33610ffd56fea2646970667358221220e1023227918cc5ebf28c67929cd41fa7c6a328a4894fc94c415c1d9ba8b8fa7b64736f6c63430008180033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// MinDelegation is a free data retrieval call binding the contract method 0x59bcddde.
//
// Solidity: function MinDelegation() view returns(uint256)
func (_Staking *StakingCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MinDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegation is a free data retrieval call binding the contract method 0x59bcddde.
//
// Solidity: function MinDelegation() view returns(uint256)
func (_Staking *StakingSession) MinDelegation() (*big.Int, error) {
	return _Staking.Contract.MinDelegation(&_Staking.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x59bcddde.
//
// Solidity: function MinDelegation() view returns(uint256)
func (_Staking *StakingCallerSession) MinDelegation() (*big.Int, error) {
	return _Staking.Contract.MinDelegation(&_Staking.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x11bcd830.
//
// Solidity: function MinDeposit() view returns(uint256)
func (_Staking *StakingCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MinDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x11bcd830.
//
// Solidity: function MinDeposit() view returns(uint256)
func (_Staking *StakingSession) MinDeposit() (*big.Int, error) {
	return _Staking.Contract.MinDeposit(&_Staking.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x11bcd830.
//
// Solidity: function MinDeposit() view returns(uint256)
func (_Staking *StakingCallerSession) MinDeposit() (*big.Int, error) {
	return _Staking.Contract.MinDeposit(&_Staking.CallOpts)
}

// IsAllowedValidator is a free data retrieval call binding the contract method 0x84768b7a.
//
// Solidity: function isAllowedValidator(address ) view returns(bool)
func (_Staking *StakingCaller) IsAllowedValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isAllowedValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedValidator is a free data retrieval call binding the contract method 0x84768b7a.
//
// Solidity: function isAllowedValidator(address ) view returns(bool)
func (_Staking *StakingSession) IsAllowedValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.IsAllowedValidator(&_Staking.CallOpts, arg0)
}

// IsAllowedValidator is a free data retrieval call binding the contract method 0x84768b7a.
//
// Solidity: function isAllowedValidator(address ) view returns(bool)
func (_Staking *StakingCallerSession) IsAllowedValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.IsAllowedValidator(&_Staking.CallOpts, arg0)
}

// IsAllowlistEnabled is a free data retrieval call binding the contract method 0xd146fd1b.
//
// Solidity: function isAllowlistEnabled() view returns(bool)
func (_Staking *StakingCaller) IsAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowlistEnabled is a free data retrieval call binding the contract method 0xd146fd1b.
//
// Solidity: function isAllowlistEnabled() view returns(bool)
func (_Staking *StakingSession) IsAllowlistEnabled() (bool, error) {
	return _Staking.Contract.IsAllowlistEnabled(&_Staking.CallOpts)
}

// IsAllowlistEnabled is a free data retrieval call binding the contract method 0xd146fd1b.
//
// Solidity: function isAllowlistEnabled() view returns(bool)
func (_Staking *StakingCallerSession) IsAllowlistEnabled() (bool, error) {
	return _Staking.Contract.IsAllowlistEnabled(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// AllowValidators is a paid mutator transaction binding the contract method 0x117407e3.
//
// Solidity: function allowValidators(address[] validators) returns()
func (_Staking *StakingTransactor) AllowValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "allowValidators", validators)
}

// AllowValidators is a paid mutator transaction binding the contract method 0x117407e3.
//
// Solidity: function allowValidators(address[] validators) returns()
func (_Staking *StakingSession) AllowValidators(validators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AllowValidators(&_Staking.TransactOpts, validators)
}

// AllowValidators is a paid mutator transaction binding the contract method 0x117407e3.
//
// Solidity: function allowValidators(address[] validators) returns()
func (_Staking *StakingTransactorSession) AllowValidators(validators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AllowValidators(&_Staking.TransactOpts, validators)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Staking *StakingTransactor) CreateValidator(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "createValidator", pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Staking *StakingSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator(&_Staking.TransactOpts, pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Staking *StakingTransactorSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator(&_Staking.TransactOpts, pubkey)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_Staking *StakingSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_Staking *StakingTransactorSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, validator)
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_Staking *StakingTransactor) DisableAllowlist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "disableAllowlist")
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_Staking *StakingSession) DisableAllowlist() (*types.Transaction, error) {
	return _Staking.Contract.DisableAllowlist(&_Staking.TransactOpts)
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_Staking *StakingTransactorSession) DisableAllowlist() (*types.Transaction, error) {
	return _Staking.Contract.DisableAllowlist(&_Staking.TransactOpts)
}

// DisallowValidators is a paid mutator transaction binding the contract method 0x3f0b1edf.
//
// Solidity: function disallowValidators(address[] validators) returns()
func (_Staking *StakingTransactor) DisallowValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "disallowValidators", validators)
}

// DisallowValidators is a paid mutator transaction binding the contract method 0x3f0b1edf.
//
// Solidity: function disallowValidators(address[] validators) returns()
func (_Staking *StakingSession) DisallowValidators(validators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DisallowValidators(&_Staking.TransactOpts, validators)
}

// DisallowValidators is a paid mutator transaction binding the contract method 0x3f0b1edf.
//
// Solidity: function disallowValidators(address[] validators) returns()
func (_Staking *StakingTransactorSession) DisallowValidators(validators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DisallowValidators(&_Staking.TransactOpts, validators)
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_Staking *StakingTransactor) EnableAllowlist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "enableAllowlist")
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_Staking *StakingSession) EnableAllowlist() (*types.Transaction, error) {
	return _Staking.Contract.EnableAllowlist(&_Staking.TransactOpts)
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_Staking *StakingTransactorSession) EnableAllowlist() (*types.Transaction, error) {
	return _Staking.Contract.EnableAllowlist(&_Staking.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", owner_, isAllowlistEnabled_)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingSession) Initialize(owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, owner_, isAllowlistEnabled_)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingTransactorSession) Initialize(owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, owner_, isAllowlistEnabled_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// StakingAllowlistDisabledIterator is returned from FilterAllowlistDisabled and is used to iterate over the raw logs and unpacked data for AllowlistDisabled events raised by the Staking contract.
type StakingAllowlistDisabledIterator struct {
	Event *StakingAllowlistDisabled // Event containing the contract specifics and raw log

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
func (it *StakingAllowlistDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAllowlistDisabled)
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
		it.Event = new(StakingAllowlistDisabled)
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
func (it *StakingAllowlistDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAllowlistDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAllowlistDisabled represents a AllowlistDisabled event raised by the Staking contract.
type StakingAllowlistDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllowlistDisabled is a free log retrieval operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_Staking *StakingFilterer) FilterAllowlistDisabled(opts *bind.FilterOpts) (*StakingAllowlistDisabledIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AllowlistDisabled")
	if err != nil {
		return nil, err
	}
	return &StakingAllowlistDisabledIterator{contract: _Staking.contract, event: "AllowlistDisabled", logs: logs, sub: sub}, nil
}

// WatchAllowlistDisabled is a free log subscription operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_Staking *StakingFilterer) WatchAllowlistDisabled(opts *bind.WatchOpts, sink chan<- *StakingAllowlistDisabled) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AllowlistDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAllowlistDisabled)
				if err := _Staking.contract.UnpackLog(event, "AllowlistDisabled", log); err != nil {
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

// ParseAllowlistDisabled is a log parse operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_Staking *StakingFilterer) ParseAllowlistDisabled(log types.Log) (*StakingAllowlistDisabled, error) {
	event := new(StakingAllowlistDisabled)
	if err := _Staking.contract.UnpackLog(event, "AllowlistDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingAllowlistEnabledIterator is returned from FilterAllowlistEnabled and is used to iterate over the raw logs and unpacked data for AllowlistEnabled events raised by the Staking contract.
type StakingAllowlistEnabledIterator struct {
	Event *StakingAllowlistEnabled // Event containing the contract specifics and raw log

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
func (it *StakingAllowlistEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAllowlistEnabled)
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
		it.Event = new(StakingAllowlistEnabled)
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
func (it *StakingAllowlistEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAllowlistEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAllowlistEnabled represents a AllowlistEnabled event raised by the Staking contract.
type StakingAllowlistEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllowlistEnabled is a free log retrieval operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_Staking *StakingFilterer) FilterAllowlistEnabled(opts *bind.FilterOpts) (*StakingAllowlistEnabledIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AllowlistEnabled")
	if err != nil {
		return nil, err
	}
	return &StakingAllowlistEnabledIterator{contract: _Staking.contract, event: "AllowlistEnabled", logs: logs, sub: sub}, nil
}

// WatchAllowlistEnabled is a free log subscription operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_Staking *StakingFilterer) WatchAllowlistEnabled(opts *bind.WatchOpts, sink chan<- *StakingAllowlistEnabled) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AllowlistEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAllowlistEnabled)
				if err := _Staking.contract.UnpackLog(event, "AllowlistEnabled", log); err != nil {
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

// ParseAllowlistEnabled is a log parse operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_Staking *StakingFilterer) ParseAllowlistEnabled(log types.Log) (*StakingAllowlistEnabled, error) {
	event := new(StakingAllowlistEnabled)
	if err := _Staking.contract.UnpackLog(event, "AllowlistEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCreateValidatorIterator is returned from FilterCreateValidator and is used to iterate over the raw logs and unpacked data for CreateValidator events raised by the Staking contract.
type StakingCreateValidatorIterator struct {
	Event *StakingCreateValidator // Event containing the contract specifics and raw log

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
func (it *StakingCreateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCreateValidator)
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
		it.Event = new(StakingCreateValidator)
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
func (it *StakingCreateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCreateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCreateValidator represents a CreateValidator event raised by the Staking contract.
type StakingCreateValidator struct {
	Validator common.Address
	Pubkey    []byte
	Deposit   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateValidator is a free log retrieval operation binding the contract event 0xc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a453.
//
// Solidity: event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit)
func (_Staking *StakingFilterer) FilterCreateValidator(opts *bind.FilterOpts, validator []common.Address) (*StakingCreateValidatorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "CreateValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingCreateValidatorIterator{contract: _Staking.contract, event: "CreateValidator", logs: logs, sub: sub}, nil
}

// WatchCreateValidator is a free log subscription operation binding the contract event 0xc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a453.
//
// Solidity: event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit)
func (_Staking *StakingFilterer) WatchCreateValidator(opts *bind.WatchOpts, sink chan<- *StakingCreateValidator, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "CreateValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCreateValidator)
				if err := _Staking.contract.UnpackLog(event, "CreateValidator", log); err != nil {
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

// ParseCreateValidator is a log parse operation binding the contract event 0xc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a453.
//
// Solidity: event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit)
func (_Staking *StakingFilterer) ParseCreateValidator(log types.Log) (*StakingCreateValidator, error) {
	event := new(StakingCreateValidator)
	if err := _Staking.contract.UnpackLog(event, "CreateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the Staking contract.
type StakingDelegateIterator struct {
	Event *StakingDelegate // Event containing the contract specifics and raw log

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
func (it *StakingDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegate)
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
		it.Event = new(StakingDelegate)
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
func (it *StakingDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegate represents a Delegate event raised by the Staking contract.
type StakingDelegate struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegate(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*StakingDelegateIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Delegate", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegateIterator{contract: _Staking.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakingDelegate, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Delegate", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegate)
				if err := _Staking.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegate(log types.Log) (*StakingDelegate, error) {
	event := new(StakingDelegate)
	if err := _Staking.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorAllowedIterator is returned from FilterValidatorAllowed and is used to iterate over the raw logs and unpacked data for ValidatorAllowed events raised by the Staking contract.
type StakingValidatorAllowedIterator struct {
	Event *StakingValidatorAllowed // Event containing the contract specifics and raw log

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
func (it *StakingValidatorAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorAllowed)
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
		it.Event = new(StakingValidatorAllowed)
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
func (it *StakingValidatorAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorAllowed represents a ValidatorAllowed event raised by the Staking contract.
type StakingValidatorAllowed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAllowed is a free log retrieval operation binding the contract event 0xc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de98.
//
// Solidity: event ValidatorAllowed(address indexed validator)
func (_Staking *StakingFilterer) FilterValidatorAllowed(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorAllowedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorAllowed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorAllowedIterator{contract: _Staking.contract, event: "ValidatorAllowed", logs: logs, sub: sub}, nil
}

// WatchValidatorAllowed is a free log subscription operation binding the contract event 0xc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de98.
//
// Solidity: event ValidatorAllowed(address indexed validator)
func (_Staking *StakingFilterer) WatchValidatorAllowed(opts *bind.WatchOpts, sink chan<- *StakingValidatorAllowed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorAllowed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorAllowed)
				if err := _Staking.contract.UnpackLog(event, "ValidatorAllowed", log); err != nil {
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

// ParseValidatorAllowed is a log parse operation binding the contract event 0xc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de98.
//
// Solidity: event ValidatorAllowed(address indexed validator)
func (_Staking *StakingFilterer) ParseValidatorAllowed(log types.Log) (*StakingValidatorAllowed, error) {
	event := new(StakingValidatorAllowed)
	if err := _Staking.contract.UnpackLog(event, "ValidatorAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorDisallowedIterator is returned from FilterValidatorDisallowed and is used to iterate over the raw logs and unpacked data for ValidatorDisallowed events raised by the Staking contract.
type StakingValidatorDisallowedIterator struct {
	Event *StakingValidatorDisallowed // Event containing the contract specifics and raw log

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
func (it *StakingValidatorDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorDisallowed)
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
		it.Event = new(StakingValidatorDisallowed)
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
func (it *StakingValidatorDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorDisallowed represents a ValidatorDisallowed event raised by the Staking contract.
type StakingValidatorDisallowed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDisallowed is a free log retrieval operation binding the contract event 0x3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f.
//
// Solidity: event ValidatorDisallowed(address indexed validator)
func (_Staking *StakingFilterer) FilterValidatorDisallowed(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorDisallowedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorDisallowed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorDisallowedIterator{contract: _Staking.contract, event: "ValidatorDisallowed", logs: logs, sub: sub}, nil
}

// WatchValidatorDisallowed is a free log subscription operation binding the contract event 0x3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f.
//
// Solidity: event ValidatorDisallowed(address indexed validator)
func (_Staking *StakingFilterer) WatchValidatorDisallowed(opts *bind.WatchOpts, sink chan<- *StakingValidatorDisallowed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorDisallowed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorDisallowed)
				if err := _Staking.contract.UnpackLog(event, "ValidatorDisallowed", log); err != nil {
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

// ParseValidatorDisallowed is a log parse operation binding the contract event 0x3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f.
//
// Solidity: event ValidatorDisallowed(address indexed validator)
func (_Staking *StakingFilterer) ParseValidatorDisallowed(log types.Log) (*StakingValidatorDisallowed, error) {
	event := new(StakingValidatorDisallowed)
	if err := _Staking.contract.UnpackLog(event, "ValidatorDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
