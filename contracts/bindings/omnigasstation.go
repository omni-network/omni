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

// OmniGasStationGasPump is an auto generated low-level Go binding around an user-defined struct.
type OmniGasStationGasPump struct {
	ChainID uint64
	Addr    common.Address
}

// OmniGasStationMetaData contains all meta data concerning the OmniGasStation contract.
var OmniGasStationMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fueled\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pumps_\",\"type\":\"tuple[]\",\"internalType\":\"structOmniGasStation.GasPump[]\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPump\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pumps\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPump\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settleUp\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasPumpAdded\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SettledUp\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"owed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fueled\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61104b806100df6000396000f3fe6080604052600436106100e15760003560e01c806374eeb8471161007f578063a04feb8611610059578063a04feb8614610282578063bac34036146102a2578063ed116d08146102e8578063f2fde38b1461030857600080fd5b806374eeb847146101fd5780638456cb59146102305780638da5cb5b1461024557600080fd5b80634afa17dd116100bb5780634afa17dd146101615780634e0dc4f0146101815780635c975abb146101b7578063715018a6146101e857600080fd5b806339acf9f1146100ed5780633bd9b9f61461012a5780633f4ba83a1461014c57600080fd5b366100e857005b600080fd5b3480156100f957600080fd5b5060005461010d906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561013657600080fd5b5061014a610145366004610db3565b610328565b005b34801561015857600080fd5b5061014a6105f5565b34801561016d57600080fd5b5061014a61017c366004610df4565b610607565b34801561018d57600080fd5b5061010d61019c366004610e2d565b6032602052600090815260409020546001600160a01b031681565b3480156101c357600080fd5b50600080516020610ff68339815191525460ff165b6040519015158152602001610121565b3480156101f457600080fd5b5061014a61061d565b34801561020957600080fd5b5060005461021e90600160a01b900460ff1681565b60405160ff9091168152602001610121565b34801561023c57600080fd5b5061014a61062f565b34801561025157600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661010d565b34801561028e57600080fd5b5061014a61029d366004610e51565b61063f565b3480156102ae57600080fd5b506102da6102bd366004610ee5565b603360209081526000928352604080842090915290825290205481565b604051908152602001610121565b3480156102f457600080fd5b506101d8610303366004610df4565b6107c9565b34801561031457600080fd5b5061014a610323366004610f13565b61080f565b60005460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa15801561036f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103939190610f30565b8051600180546020909301516001600160a01b0316600160401b026001600160e01b03199093166001600160401b03909216919091179190911790556103d761084d565b6103df61087e565b801561040d575060015461040d906001600160401b03811690600160401b90046001600160a01b03166107c9565b61045e5760405162461bcd60e51b815260206004820152601860248201527f47617353746174696f6e3a20756e617574686f72697a6564000000000000000060448201526064015b60405180910390fd5b6001600160a01b03821660009081526033602090815260408083206001546001600160401b031684529091529020548082116104dc5760405162461bcd60e51b815260206004820152601a60248201527f47617353746174696f6e3a20616c72656164792066756e6465640000000000006044820152606401610455565b60006001600160a01b0384166104f28385610f9c565b604051600081818185875af1925050503d806000811461052e576040519150601f19603f3d011682016040523d82523d6000602084013e610533565b606091505b50509050801561056f576001600160a01b03841660009081526033602090815260408083206001546001600160401b0316845290915290208390555b6001546001600160a01b03851660008181526033602090815260408083206001600160401b0390951680845294825291829020548251888152918201528415158183015290517f4264b2d9471008d8513ddd06a5da387491ccaf43988f604f5eca833d30551c9d9181900360600190a35050600180546001600160e01b03191690555050565b6105fd610911565b61060561096c565b565b61060f610911565b61061982826109cd565b5050565b610625610911565b6106056000610ae0565b610637610911565b610605610b51565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156106845750825b90506000826001600160401b031660011480156106a05750303b155b9050811580156106ae575080155b156106cc5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156106f657845460ff60401b1916600160401b1785555b610701896004610b9a565b61070a88610bb4565b60005b868110156107775761076f88888381811061072a5761072a610fbd565b6107409260206040909202019081019150610e2d565b89898481811061075257610752610fbd565b905060400201602001602081019061076a9190610f13565b6109cd565b60010161070d565b5083156107be57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b60006001600160a01b0382161580159061080657506001600160401b0383166000908152603260205260409020546001600160a01b038381169116145b90505b92915050565b610817610911565b6001600160a01b03811661084157604051631e4fbdf760e01b815260006004820152602401610455565b61084a81610ae0565b50565b600080516020610ff68339815191525460ff16156106055760405163d93c066560e01b815260040160405180910390fd5b60008060009054906101000a90046001600160a01b03166001600160a01b03166355e2448e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f69190610fd3565b801561090c57506000546001600160a01b031633145b905090565b336109437f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146106055760405163118cdaa760e01b8152336004820152602401610455565b610974610bc5565b600080516020610ff6833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020015b60405180910390a150565b6001600160a01b038116610a1b5760405162461bcd60e51b815260206004820152601560248201527423b0b9a9ba30ba34b7b71d103d32b9379030b2323960591b6044820152606401610455565b816001600160401b0316600003610a745760405162461bcd60e51b815260206004820152601860248201527f47617353746174696f6e3a207a65726f20636861696e496400000000000000006044820152606401610455565b6001600160401b03821660008181526032602090815260409182902080546001600160a01b0319166001600160a01b03861690811790915591519182527f193d7693e9edcedc2357451dd62125041be2c877aa64852b126906cd438f08f5910160405180910390a25050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b610b5961084d565b600080516020610ff6833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336109ae565b610ba2610bf5565b610bab82610c3e565b61061981610cd7565b610bbc610bf5565b61084a81610d7a565b600080516020610ff68339815191525460ff1661060557604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661060557604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b038116610c895760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b6044820152606401610455565b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47906020016109c2565b610ce081610d82565b610d2c5760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e76616c696420636f6e66206c6576656c00000000000000006044820152606401610455565b6000805460ff60a01b1916600160a01b60ff8416908102919091179091556040519081527f8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483906020016109c2565b610817610bf5565b600060ff821660011480610809575060ff821660041492915050565b6001600160a01b038116811461084a57600080fd5b60008060408385031215610dc657600080fd5b8235610dd181610d9e565b946020939093013593505050565b6001600160401b038116811461084a57600080fd5b60008060408385031215610e0757600080fd5b8235610e1281610ddf565b91506020830135610e2281610d9e565b809150509250929050565b600060208284031215610e3f57600080fd5b8135610e4a81610ddf565b9392505050565b60008060008060608587031215610e6757600080fd5b8435610e7281610d9e565b93506020850135610e8281610d9e565b925060408501356001600160401b0380821115610e9e57600080fd5b818701915087601f830112610eb257600080fd5b813581811115610ec157600080fd5b8860208260061b8501011115610ed657600080fd5b95989497505060200194505050565b60008060408385031215610ef857600080fd5b8235610f0381610d9e565b91506020830135610e2281610ddf565b600060208284031215610f2557600080fd5b8135610e4a81610d9e565b600060408284031215610f4257600080fd5b604051604081018181106001600160401b0382111715610f7257634e487b7160e01b600052604160045260246000fd5b6040528251610f8081610ddf565b81526020830151610f9081610d9e565b60208201529392505050565b8181038181111561080957634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060208284031215610fe557600080fd5b81518015158114610e4a57600080fdfecd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220d0bd338e753dc81ce87f0d7d39806756dbb7db8f06c18e4871173f7b78a70a6564736f6c63430008180033",
}

// OmniGasStationABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniGasStationMetaData.ABI instead.
var OmniGasStationABI = OmniGasStationMetaData.ABI

// OmniGasStationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniGasStationMetaData.Bin instead.
var OmniGasStationBin = OmniGasStationMetaData.Bin

// DeployOmniGasStation deploys a new Ethereum contract, binding an instance of OmniGasStation to it.
func DeployOmniGasStation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniGasStation, error) {
	parsed, err := OmniGasStationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniGasStationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniGasStation{OmniGasStationCaller: OmniGasStationCaller{contract: contract}, OmniGasStationTransactor: OmniGasStationTransactor{contract: contract}, OmniGasStationFilterer: OmniGasStationFilterer{contract: contract}}, nil
}

// OmniGasStation is an auto generated Go binding around an Ethereum contract.
type OmniGasStation struct {
	OmniGasStationCaller     // Read-only binding to the contract
	OmniGasStationTransactor // Write-only binding to the contract
	OmniGasStationFilterer   // Log filterer for contract events
}

// OmniGasStationCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniGasStationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniGasStationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniGasStationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniGasStationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniGasStationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniGasStationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniGasStationSession struct {
	Contract     *OmniGasStation   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniGasStationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniGasStationCallerSession struct {
	Contract *OmniGasStationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OmniGasStationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniGasStationTransactorSession struct {
	Contract     *OmniGasStationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OmniGasStationRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniGasStationRaw struct {
	Contract *OmniGasStation // Generic contract binding to access the raw methods on
}

// OmniGasStationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniGasStationCallerRaw struct {
	Contract *OmniGasStationCaller // Generic read-only contract binding to access the raw methods on
}

// OmniGasStationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniGasStationTransactorRaw struct {
	Contract *OmniGasStationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniGasStation creates a new instance of OmniGasStation, bound to a specific deployed contract.
func NewOmniGasStation(address common.Address, backend bind.ContractBackend) (*OmniGasStation, error) {
	contract, err := bindOmniGasStation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniGasStation{OmniGasStationCaller: OmniGasStationCaller{contract: contract}, OmniGasStationTransactor: OmniGasStationTransactor{contract: contract}, OmniGasStationFilterer: OmniGasStationFilterer{contract: contract}}, nil
}

// NewOmniGasStationCaller creates a new read-only instance of OmniGasStation, bound to a specific deployed contract.
func NewOmniGasStationCaller(address common.Address, caller bind.ContractCaller) (*OmniGasStationCaller, error) {
	contract, err := bindOmniGasStation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniGasStationCaller{contract: contract}, nil
}

// NewOmniGasStationTransactor creates a new write-only instance of OmniGasStation, bound to a specific deployed contract.
func NewOmniGasStationTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniGasStationTransactor, error) {
	contract, err := bindOmniGasStation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniGasStationTransactor{contract: contract}, nil
}

// NewOmniGasStationFilterer creates a new log filterer instance of OmniGasStation, bound to a specific deployed contract.
func NewOmniGasStationFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniGasStationFilterer, error) {
	contract, err := bindOmniGasStation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniGasStationFilterer{contract: contract}, nil
}

// bindOmniGasStation binds a generic wrapper to an already deployed contract.
func bindOmniGasStation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniGasStationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniGasStation *OmniGasStationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniGasStation.Contract.OmniGasStationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniGasStation *OmniGasStationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasStation.Contract.OmniGasStationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniGasStation *OmniGasStationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniGasStation.Contract.OmniGasStationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniGasStation *OmniGasStationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniGasStation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniGasStation *OmniGasStationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasStation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniGasStation *OmniGasStationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniGasStation.Contract.contract.Transact(opts, method, params...)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_OmniGasStation *OmniGasStationCaller) DefaultConfLevel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniGasStation.contract.Call(opts, &out, "defaultConfLevel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_OmniGasStation *OmniGasStationSession) DefaultConfLevel() (uint8, error) {
	return _OmniGasStation.Contract.DefaultConfLevel(&_OmniGasStation.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_OmniGasStation *OmniGasStationCallerSession) DefaultConfLevel() (uint8, error) {
	return _OmniGasStation.Contract.DefaultConfLevel(&_OmniGasStation.CallOpts)
}

// Fueled is a free data retrieval call binding the contract method 0xbac34036.
//
// Solidity: function fueled(address , uint64 ) view returns(uint256)
func (_OmniGasStation *OmniGasStationCaller) Fueled(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) (*big.Int, error) {
	var out []interface{}
	err := _OmniGasStation.contract.Call(opts, &out, "fueled", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fueled is a free data retrieval call binding the contract method 0xbac34036.
//
// Solidity: function fueled(address , uint64 ) view returns(uint256)
func (_OmniGasStation *OmniGasStationSession) Fueled(arg0 common.Address, arg1 uint64) (*big.Int, error) {
	return _OmniGasStation.Contract.Fueled(&_OmniGasStation.CallOpts, arg0, arg1)
}

// Fueled is a free data retrieval call binding the contract method 0xbac34036.
//
// Solidity: function fueled(address , uint64 ) view returns(uint256)
func (_OmniGasStation *OmniGasStationCallerSession) Fueled(arg0 common.Address, arg1 uint64) (*big.Int, error) {
	return _OmniGasStation.Contract.Fueled(&_OmniGasStation.CallOpts, arg0, arg1)
}

// IsPump is a free data retrieval call binding the contract method 0xed116d08.
//
// Solidity: function isPump(uint64 chainID, address addr) view returns(bool)
func (_OmniGasStation *OmniGasStationCaller) IsPump(opts *bind.CallOpts, chainID uint64, addr common.Address) (bool, error) {
	var out []interface{}
	err := _OmniGasStation.contract.Call(opts, &out, "isPump", chainID, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPump is a free data retrieval call binding the contract method 0xed116d08.
//
// Solidity: function isPump(uint64 chainID, address addr) view returns(bool)
func (_OmniGasStation *OmniGasStationSession) IsPump(chainID uint64, addr common.Address) (bool, error) {
	return _OmniGasStation.Contract.IsPump(&_OmniGasStation.CallOpts, chainID, addr)
}

// IsPump is a free data retrieval call binding the contract method 0xed116d08.
//
// Solidity: function isPump(uint64 chainID, address addr) view returns(bool)
func (_OmniGasStation *OmniGasStationCallerSession) IsPump(chainID uint64, addr common.Address) (bool, error) {
	return _OmniGasStation.Contract.IsPump(&_OmniGasStation.CallOpts, chainID, addr)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniGasStation *OmniGasStationCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniGasStation.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniGasStation *OmniGasStationSession) Omni() (common.Address, error) {
	return _OmniGasStation.Contract.Omni(&_OmniGasStation.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniGasStation *OmniGasStationCallerSession) Omni() (common.Address, error) {
	return _OmniGasStation.Contract.Omni(&_OmniGasStation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniGasStation *OmniGasStationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniGasStation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniGasStation *OmniGasStationSession) Owner() (common.Address, error) {
	return _OmniGasStation.Contract.Owner(&_OmniGasStation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniGasStation *OmniGasStationCallerSession) Owner() (common.Address, error) {
	return _OmniGasStation.Contract.Owner(&_OmniGasStation.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniGasStation *OmniGasStationCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniGasStation.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniGasStation *OmniGasStationSession) Paused() (bool, error) {
	return _OmniGasStation.Contract.Paused(&_OmniGasStation.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniGasStation *OmniGasStationCallerSession) Paused() (bool, error) {
	return _OmniGasStation.Contract.Paused(&_OmniGasStation.CallOpts)
}

// Pumps is a free data retrieval call binding the contract method 0x4e0dc4f0.
//
// Solidity: function pumps(uint64 ) view returns(address)
func (_OmniGasStation *OmniGasStationCaller) Pumps(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _OmniGasStation.contract.Call(opts, &out, "pumps", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pumps is a free data retrieval call binding the contract method 0x4e0dc4f0.
//
// Solidity: function pumps(uint64 ) view returns(address)
func (_OmniGasStation *OmniGasStationSession) Pumps(arg0 uint64) (common.Address, error) {
	return _OmniGasStation.Contract.Pumps(&_OmniGasStation.CallOpts, arg0)
}

// Pumps is a free data retrieval call binding the contract method 0x4e0dc4f0.
//
// Solidity: function pumps(uint64 ) view returns(address)
func (_OmniGasStation *OmniGasStationCallerSession) Pumps(arg0 uint64) (common.Address, error) {
	return _OmniGasStation.Contract.Pumps(&_OmniGasStation.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xa04feb86.
//
// Solidity: function initialize(address portal, address owner, (uint64,address)[] pumps_) returns()
func (_OmniGasStation *OmniGasStationTransactor) Initialize(opts *bind.TransactOpts, portal common.Address, owner common.Address, pumps_ []OmniGasStationGasPump) (*types.Transaction, error) {
	return _OmniGasStation.contract.Transact(opts, "initialize", portal, owner, pumps_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa04feb86.
//
// Solidity: function initialize(address portal, address owner, (uint64,address)[] pumps_) returns()
func (_OmniGasStation *OmniGasStationSession) Initialize(portal common.Address, owner common.Address, pumps_ []OmniGasStationGasPump) (*types.Transaction, error) {
	return _OmniGasStation.Contract.Initialize(&_OmniGasStation.TransactOpts, portal, owner, pumps_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa04feb86.
//
// Solidity: function initialize(address portal, address owner, (uint64,address)[] pumps_) returns()
func (_OmniGasStation *OmniGasStationTransactorSession) Initialize(portal common.Address, owner common.Address, pumps_ []OmniGasStationGasPump) (*types.Transaction, error) {
	return _OmniGasStation.Contract.Initialize(&_OmniGasStation.TransactOpts, portal, owner, pumps_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniGasStation *OmniGasStationTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasStation.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniGasStation *OmniGasStationSession) Pause() (*types.Transaction, error) {
	return _OmniGasStation.Contract.Pause(&_OmniGasStation.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniGasStation *OmniGasStationTransactorSession) Pause() (*types.Transaction, error) {
	return _OmniGasStation.Contract.Pause(&_OmniGasStation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniGasStation *OmniGasStationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasStation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniGasStation *OmniGasStationSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniGasStation.Contract.RenounceOwnership(&_OmniGasStation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniGasStation *OmniGasStationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniGasStation.Contract.RenounceOwnership(&_OmniGasStation.TransactOpts)
}

// SetPump is a paid mutator transaction binding the contract method 0x4afa17dd.
//
// Solidity: function setPump(uint64 chainId, address addr) returns()
func (_OmniGasStation *OmniGasStationTransactor) SetPump(opts *bind.TransactOpts, chainId uint64, addr common.Address) (*types.Transaction, error) {
	return _OmniGasStation.contract.Transact(opts, "setPump", chainId, addr)
}

// SetPump is a paid mutator transaction binding the contract method 0x4afa17dd.
//
// Solidity: function setPump(uint64 chainId, address addr) returns()
func (_OmniGasStation *OmniGasStationSession) SetPump(chainId uint64, addr common.Address) (*types.Transaction, error) {
	return _OmniGasStation.Contract.SetPump(&_OmniGasStation.TransactOpts, chainId, addr)
}

// SetPump is a paid mutator transaction binding the contract method 0x4afa17dd.
//
// Solidity: function setPump(uint64 chainId, address addr) returns()
func (_OmniGasStation *OmniGasStationTransactorSession) SetPump(chainId uint64, addr common.Address) (*types.Transaction, error) {
	return _OmniGasStation.Contract.SetPump(&_OmniGasStation.TransactOpts, chainId, addr)
}

// SettleUp is a paid mutator transaction binding the contract method 0x3bd9b9f6.
//
// Solidity: function settleUp(address recipient, uint256 owed) returns()
func (_OmniGasStation *OmniGasStationTransactor) SettleUp(opts *bind.TransactOpts, recipient common.Address, owed *big.Int) (*types.Transaction, error) {
	return _OmniGasStation.contract.Transact(opts, "settleUp", recipient, owed)
}

// SettleUp is a paid mutator transaction binding the contract method 0x3bd9b9f6.
//
// Solidity: function settleUp(address recipient, uint256 owed) returns()
func (_OmniGasStation *OmniGasStationSession) SettleUp(recipient common.Address, owed *big.Int) (*types.Transaction, error) {
	return _OmniGasStation.Contract.SettleUp(&_OmniGasStation.TransactOpts, recipient, owed)
}

// SettleUp is a paid mutator transaction binding the contract method 0x3bd9b9f6.
//
// Solidity: function settleUp(address recipient, uint256 owed) returns()
func (_OmniGasStation *OmniGasStationTransactorSession) SettleUp(recipient common.Address, owed *big.Int) (*types.Transaction, error) {
	return _OmniGasStation.Contract.SettleUp(&_OmniGasStation.TransactOpts, recipient, owed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniGasStation *OmniGasStationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniGasStation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniGasStation *OmniGasStationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniGasStation.Contract.TransferOwnership(&_OmniGasStation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniGasStation *OmniGasStationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniGasStation.Contract.TransferOwnership(&_OmniGasStation.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniGasStation *OmniGasStationTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasStation.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniGasStation *OmniGasStationSession) Unpause() (*types.Transaction, error) {
	return _OmniGasStation.Contract.Unpause(&_OmniGasStation.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniGasStation *OmniGasStationTransactorSession) Unpause() (*types.Transaction, error) {
	return _OmniGasStation.Contract.Unpause(&_OmniGasStation.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OmniGasStation *OmniGasStationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasStation.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OmniGasStation *OmniGasStationSession) Receive() (*types.Transaction, error) {
	return _OmniGasStation.Contract.Receive(&_OmniGasStation.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OmniGasStation *OmniGasStationTransactorSession) Receive() (*types.Transaction, error) {
	return _OmniGasStation.Contract.Receive(&_OmniGasStation.TransactOpts)
}

// OmniGasStationDefaultConfLevelSetIterator is returned from FilterDefaultConfLevelSet and is used to iterate over the raw logs and unpacked data for DefaultConfLevelSet events raised by the OmniGasStation contract.
type OmniGasStationDefaultConfLevelSetIterator struct {
	Event *OmniGasStationDefaultConfLevelSet // Event containing the contract specifics and raw log

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
func (it *OmniGasStationDefaultConfLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasStationDefaultConfLevelSet)
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
		it.Event = new(OmniGasStationDefaultConfLevelSet)
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
func (it *OmniGasStationDefaultConfLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasStationDefaultConfLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasStationDefaultConfLevelSet represents a DefaultConfLevelSet event raised by the OmniGasStation contract.
type OmniGasStationDefaultConfLevelSet struct {
	Conf uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultConfLevelSet is a free log retrieval operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_OmniGasStation *OmniGasStationFilterer) FilterDefaultConfLevelSet(opts *bind.FilterOpts) (*OmniGasStationDefaultConfLevelSetIterator, error) {

	logs, sub, err := _OmniGasStation.contract.FilterLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return &OmniGasStationDefaultConfLevelSetIterator{contract: _OmniGasStation.contract, event: "DefaultConfLevelSet", logs: logs, sub: sub}, nil
}

// WatchDefaultConfLevelSet is a free log subscription operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_OmniGasStation *OmniGasStationFilterer) WatchDefaultConfLevelSet(opts *bind.WatchOpts, sink chan<- *OmniGasStationDefaultConfLevelSet) (event.Subscription, error) {

	logs, sub, err := _OmniGasStation.contract.WatchLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasStationDefaultConfLevelSet)
				if err := _OmniGasStation.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
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

// ParseDefaultConfLevelSet is a log parse operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_OmniGasStation *OmniGasStationFilterer) ParseDefaultConfLevelSet(log types.Log) (*OmniGasStationDefaultConfLevelSet, error) {
	event := new(OmniGasStationDefaultConfLevelSet)
	if err := _OmniGasStation.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasStationGasPumpAddedIterator is returned from FilterGasPumpAdded and is used to iterate over the raw logs and unpacked data for GasPumpAdded events raised by the OmniGasStation contract.
type OmniGasStationGasPumpAddedIterator struct {
	Event *OmniGasStationGasPumpAdded // Event containing the contract specifics and raw log

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
func (it *OmniGasStationGasPumpAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasStationGasPumpAdded)
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
		it.Event = new(OmniGasStationGasPumpAdded)
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
func (it *OmniGasStationGasPumpAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasStationGasPumpAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasStationGasPumpAdded represents a GasPumpAdded event raised by the OmniGasStation contract.
type OmniGasStationGasPumpAdded struct {
	ChainID uint64
	Addr    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGasPumpAdded is a free log retrieval operation binding the contract event 0x193d7693e9edcedc2357451dd62125041be2c877aa64852b126906cd438f08f5.
//
// Solidity: event GasPumpAdded(uint64 indexed chainID, address addr)
func (_OmniGasStation *OmniGasStationFilterer) FilterGasPumpAdded(opts *bind.FilterOpts, chainID []uint64) (*OmniGasStationGasPumpAddedIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _OmniGasStation.contract.FilterLogs(opts, "GasPumpAdded", chainIDRule)
	if err != nil {
		return nil, err
	}
	return &OmniGasStationGasPumpAddedIterator{contract: _OmniGasStation.contract, event: "GasPumpAdded", logs: logs, sub: sub}, nil
}

// WatchGasPumpAdded is a free log subscription operation binding the contract event 0x193d7693e9edcedc2357451dd62125041be2c877aa64852b126906cd438f08f5.
//
// Solidity: event GasPumpAdded(uint64 indexed chainID, address addr)
func (_OmniGasStation *OmniGasStationFilterer) WatchGasPumpAdded(opts *bind.WatchOpts, sink chan<- *OmniGasStationGasPumpAdded, chainID []uint64) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _OmniGasStation.contract.WatchLogs(opts, "GasPumpAdded", chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasStationGasPumpAdded)
				if err := _OmniGasStation.contract.UnpackLog(event, "GasPumpAdded", log); err != nil {
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

// ParseGasPumpAdded is a log parse operation binding the contract event 0x193d7693e9edcedc2357451dd62125041be2c877aa64852b126906cd438f08f5.
//
// Solidity: event GasPumpAdded(uint64 indexed chainID, address addr)
func (_OmniGasStation *OmniGasStationFilterer) ParseGasPumpAdded(log types.Log) (*OmniGasStationGasPumpAdded, error) {
	event := new(OmniGasStationGasPumpAdded)
	if err := _OmniGasStation.contract.UnpackLog(event, "GasPumpAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasStationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniGasStation contract.
type OmniGasStationInitializedIterator struct {
	Event *OmniGasStationInitialized // Event containing the contract specifics and raw log

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
func (it *OmniGasStationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasStationInitialized)
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
		it.Event = new(OmniGasStationInitialized)
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
func (it *OmniGasStationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasStationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasStationInitialized represents a Initialized event raised by the OmniGasStation contract.
type OmniGasStationInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniGasStation *OmniGasStationFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniGasStationInitializedIterator, error) {

	logs, sub, err := _OmniGasStation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniGasStationInitializedIterator{contract: _OmniGasStation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniGasStation *OmniGasStationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniGasStationInitialized) (event.Subscription, error) {

	logs, sub, err := _OmniGasStation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasStationInitialized)
				if err := _OmniGasStation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OmniGasStation *OmniGasStationFilterer) ParseInitialized(log types.Log) (*OmniGasStationInitialized, error) {
	event := new(OmniGasStationInitialized)
	if err := _OmniGasStation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasStationOmniPortalSetIterator is returned from FilterOmniPortalSet and is used to iterate over the raw logs and unpacked data for OmniPortalSet events raised by the OmniGasStation contract.
type OmniGasStationOmniPortalSetIterator struct {
	Event *OmniGasStationOmniPortalSet // Event containing the contract specifics and raw log

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
func (it *OmniGasStationOmniPortalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasStationOmniPortalSet)
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
		it.Event = new(OmniGasStationOmniPortalSet)
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
func (it *OmniGasStationOmniPortalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasStationOmniPortalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasStationOmniPortalSet represents a OmniPortalSet event raised by the OmniGasStation contract.
type OmniGasStationOmniPortalSet struct {
	Omni common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOmniPortalSet is a free log retrieval operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_OmniGasStation *OmniGasStationFilterer) FilterOmniPortalSet(opts *bind.FilterOpts) (*OmniGasStationOmniPortalSetIterator, error) {

	logs, sub, err := _OmniGasStation.contract.FilterLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return &OmniGasStationOmniPortalSetIterator{contract: _OmniGasStation.contract, event: "OmniPortalSet", logs: logs, sub: sub}, nil
}

// WatchOmniPortalSet is a free log subscription operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_OmniGasStation *OmniGasStationFilterer) WatchOmniPortalSet(opts *bind.WatchOpts, sink chan<- *OmniGasStationOmniPortalSet) (event.Subscription, error) {

	logs, sub, err := _OmniGasStation.contract.WatchLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasStationOmniPortalSet)
				if err := _OmniGasStation.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
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

// ParseOmniPortalSet is a log parse operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_OmniGasStation *OmniGasStationFilterer) ParseOmniPortalSet(log types.Log) (*OmniGasStationOmniPortalSet, error) {
	event := new(OmniGasStationOmniPortalSet)
	if err := _OmniGasStation.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasStationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniGasStation contract.
type OmniGasStationOwnershipTransferredIterator struct {
	Event *OmniGasStationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniGasStationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasStationOwnershipTransferred)
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
		it.Event = new(OmniGasStationOwnershipTransferred)
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
func (it *OmniGasStationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasStationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasStationOwnershipTransferred represents a OwnershipTransferred event raised by the OmniGasStation contract.
type OmniGasStationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniGasStation *OmniGasStationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniGasStationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniGasStation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniGasStationOwnershipTransferredIterator{contract: _OmniGasStation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniGasStation *OmniGasStationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniGasStationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniGasStation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasStationOwnershipTransferred)
				if err := _OmniGasStation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OmniGasStation *OmniGasStationFilterer) ParseOwnershipTransferred(log types.Log) (*OmniGasStationOwnershipTransferred, error) {
	event := new(OmniGasStationOwnershipTransferred)
	if err := _OmniGasStation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasStationPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OmniGasStation contract.
type OmniGasStationPausedIterator struct {
	Event *OmniGasStationPaused // Event containing the contract specifics and raw log

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
func (it *OmniGasStationPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasStationPaused)
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
		it.Event = new(OmniGasStationPaused)
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
func (it *OmniGasStationPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasStationPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasStationPaused represents a Paused event raised by the OmniGasStation contract.
type OmniGasStationPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OmniGasStation *OmniGasStationFilterer) FilterPaused(opts *bind.FilterOpts) (*OmniGasStationPausedIterator, error) {

	logs, sub, err := _OmniGasStation.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OmniGasStationPausedIterator{contract: _OmniGasStation.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OmniGasStation *OmniGasStationFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OmniGasStationPaused) (event.Subscription, error) {

	logs, sub, err := _OmniGasStation.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasStationPaused)
				if err := _OmniGasStation.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OmniGasStation *OmniGasStationFilterer) ParsePaused(log types.Log) (*OmniGasStationPaused, error) {
	event := new(OmniGasStationPaused)
	if err := _OmniGasStation.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasStationSettledUpIterator is returned from FilterSettledUp and is used to iterate over the raw logs and unpacked data for SettledUp events raised by the OmniGasStation contract.
type OmniGasStationSettledUpIterator struct {
	Event *OmniGasStationSettledUp // Event containing the contract specifics and raw log

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
func (it *OmniGasStationSettledUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasStationSettledUp)
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
		it.Event = new(OmniGasStationSettledUp)
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
func (it *OmniGasStationSettledUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasStationSettledUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasStationSettledUp represents a SettledUp event raised by the OmniGasStation contract.
type OmniGasStationSettledUp struct {
	Recipient common.Address
	ChainID   uint64
	Owed      *big.Int
	Fueled    *big.Int
	Success   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSettledUp is a free log retrieval operation binding the contract event 0x4264b2d9471008d8513ddd06a5da387491ccaf43988f604f5eca833d30551c9d.
//
// Solidity: event SettledUp(address indexed recipient, uint64 indexed chainID, uint256 owed, uint256 fueled, bool success)
func (_OmniGasStation *OmniGasStationFilterer) FilterSettledUp(opts *bind.FilterOpts, recipient []common.Address, chainID []uint64) (*OmniGasStationSettledUpIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _OmniGasStation.contract.FilterLogs(opts, "SettledUp", recipientRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return &OmniGasStationSettledUpIterator{contract: _OmniGasStation.contract, event: "SettledUp", logs: logs, sub: sub}, nil
}

// WatchSettledUp is a free log subscription operation binding the contract event 0x4264b2d9471008d8513ddd06a5da387491ccaf43988f604f5eca833d30551c9d.
//
// Solidity: event SettledUp(address indexed recipient, uint64 indexed chainID, uint256 owed, uint256 fueled, bool success)
func (_OmniGasStation *OmniGasStationFilterer) WatchSettledUp(opts *bind.WatchOpts, sink chan<- *OmniGasStationSettledUp, recipient []common.Address, chainID []uint64) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _OmniGasStation.contract.WatchLogs(opts, "SettledUp", recipientRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasStationSettledUp)
				if err := _OmniGasStation.contract.UnpackLog(event, "SettledUp", log); err != nil {
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

// ParseSettledUp is a log parse operation binding the contract event 0x4264b2d9471008d8513ddd06a5da387491ccaf43988f604f5eca833d30551c9d.
//
// Solidity: event SettledUp(address indexed recipient, uint64 indexed chainID, uint256 owed, uint256 fueled, bool success)
func (_OmniGasStation *OmniGasStationFilterer) ParseSettledUp(log types.Log) (*OmniGasStationSettledUp, error) {
	event := new(OmniGasStationSettledUp)
	if err := _OmniGasStation.contract.UnpackLog(event, "SettledUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasStationUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OmniGasStation contract.
type OmniGasStationUnpausedIterator struct {
	Event *OmniGasStationUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniGasStationUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasStationUnpaused)
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
		it.Event = new(OmniGasStationUnpaused)
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
func (it *OmniGasStationUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasStationUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasStationUnpaused represents a Unpaused event raised by the OmniGasStation contract.
type OmniGasStationUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OmniGasStation *OmniGasStationFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OmniGasStationUnpausedIterator, error) {

	logs, sub, err := _OmniGasStation.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OmniGasStationUnpausedIterator{contract: _OmniGasStation.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OmniGasStation *OmniGasStationFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OmniGasStationUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniGasStation.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasStationUnpaused)
				if err := _OmniGasStation.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OmniGasStation *OmniGasStationFilterer) ParseUnpaused(log types.Log) (*OmniGasStationUnpaused, error) {
	event := new(OmniGasStationUnpaused)
	if err := _OmniGasStation.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
