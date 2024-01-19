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
)

// XTypesBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type XTypesBlockHeader struct {
	SourceChainId uint64
	BlockHeight   uint64
	BlockHash     [32]byte
}

// XTypesMsg is an auto generated low-level Go binding around an user-defined struct.
type XTypesMsg struct {
	SourceChainId uint64
	DestChainId   uint64
	StreamOffset  uint64
	Sender        common.Address
	To            common.Address
	Data          []byte
	GasLimit      uint64
}

// XTypesSigTuple is an auto generated low-level Go binding around an user-defined struct.
type XTypesSigTuple struct {
	ValidatorPubKey []byte
	Signature       []byte
}

// XTypesSubmission is an auto generated low-level Go binding around an user-defined struct.
type XTypesSubmission struct {
	AttestationRoot [32]byte
	BlockHeader     XTypesBlockHeader
	Msgs            []XTypesMsg
	Proof           [][32]byte
	ProofFlags      []bool
	Signatures      []XTypesSigTuple
}

// OmniPortalMetaData contains all meta data concerning the OmniPortal contract.
var OmniPortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"XMSG_DEFAULT_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XMSG_MAX_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XMSG_MIN_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XRECEIPT_MAX_RETURN_DATA_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXStreamOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outXStreamOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xsubmit\",\"inputs\":[{\"name\":\"xsub\",\"type\":\"tuple\",\"internalType\":\"structXTypes.Submission\",\"components\":[{\"name\":\"attestationRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockHeader\",\"type\":\"tuple\",\"internalType\":\"structXTypes.BlockHeader\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Msg[]\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlags\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.SigTuple[]\",\"components\":[{\"name\":\"validatorPubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"XMsg\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceipt\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b506001600160401b034616608052608051610e6b6100406000396000818161015801526103420152610e6b6000f3fe6080604052600436106100915760003560e01c80639a8a0592116100595780639a8a0592146101465780639dad9aae1461017a578063a2cc111b14610191578063b58e964f146101a8578063fa590d14146101de57600080fd5b8063016655691461009657806306f9f174146100c857806350e646dd146100ea57806370e8b56a146100fd57806390ab417c14610110575b600080fd5b3480156100a257600080fd5b506100ac61010081565b6040516001600160401b03909116815260200160405180910390f35b3480156100d457600080fd5b506100e86100e3366004610931565b6101f4565b005b6100e86100f83660046109ee565b610250565b6100e861010b366004610a4e565b610267565b34801561011c57600080fd5b506100ac61012b366004610ac3565b6000602081905290815260409020546001600160401b031681565b34801561015257600080fd5b506100ac7f000000000000000000000000000000000000000000000000000000000000000081565b34801561018657600080fd5b506100ac62030d4081565b34801561019d57600080fd5b506100ac624c4b4081565b3480156101b457600080fd5b506100ac6101c3366004610ac3565b6001602052600090815260409020546001600160401b031681565b3480156101ea57600080fd5b506100ac61520881565b60005b6102046080830183610ade565b905081101561024c5761024461021d6080840184610ade565b8381811061022d5761022d610b27565b905060200281019061023f9190610b3d565b610338565b6001016101f7565b5050565b610261843385858562030d40610596565b50505050565b624c4b406001600160401b03821611156102c85760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f206869676800000060448201526064015b60405180910390fd5b6152086001600160401b03821610156103235760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f206c6f770000000060448201526064016102bf565b610331853386868686610596565b5050505050565b6001600160401b037f0000000000000000000000000000000000000000000000000000000000000000166103726040830160208401610ac3565b6001600160401b0316146103c85760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a2077726f6e672064657374436861696e496400000060448201526064016102bf565b600160006103d96020840184610ac3565b6001600160401b03908116825260208201929092526040908101600020549091169061040b9060608401908401610ac3565b6001600160401b0316146104615760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a2077726f6e672073747265616d4f6666736574000060448201526064016102bf565b60018060006104736020850185610ac3565b6001600160401b03908116825260208201929092526040016000908120805490926104a091859116610b73565b92506101000a8154816001600160401b0302191690836001600160401b031602179055506000610527826104d390610c64565b3360405180604001604052806101006001600160401b03168152602001604051806040016040528060168152602001754f6d6e69506f7274616c3a206f7574206f662067617360501b81525081525061064d565b905080604001516001600160401b031681600001516001600160401b03167fc2f03986562c9333b7872c7694832b7583b6a4dbde9868c9d43a2a3896c64343836060015184608001518560a001518660c0015160405161058a9493929190610d5f565b60405180910390a35050565b6001600160401b03808716600081815260208190526040908190205490519216917fac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283906105ec9089908990899089908990610d9b565b60405180910390a36001600160401b038087166000908152602081905260408120805460019391929161062191859116610b73565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550505050505050565b6040805160e08101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820152905a905060008086608001516001600160a01b03168760c001516001600160401b03168860a001516040516106b99190610df7565b60006040518083038160008787f1925050503d80600081146106f7576040519150601f19603f3d011682016040523d82523d6000602084013e6106fc565b606091505b50915091505a61070c9084610e09565b925084600001516001600160401b0316815111156107405761073d81600087600001516001600160401b03166107ef565b90505b8660c001516001600160401b031683106107955784602001516040516024016107699190610e22565b60408051601f198184030181529190526020810180516001600160e01b031662461bcd60e51b17905290505b6040805160e08101825288516001600160401b0390811682526020808b0151821690830152988201519098169088015260608701929092526001600160a01b03909416608086015292151560a0850152505060c082015290565b60608182601f0110156108355760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016102bf565b8282840110156108785760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016102bf565b818301845110156108bf5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b60448201526064016102bf565b6060821580156108de5760405191506000825260208201604052610928565b6040519150601f8416801560200281840101858101878315602002848b0101015b818310156109175780518352602092830192016108ff565b5050858452601f01601f1916604052505b50949350505050565b60006020828403121561094357600080fd5b81356001600160401b0381111561095957600080fd5b8201610100818503121561096c57600080fd5b9392505050565b80356001600160401b038116811461098a57600080fd5b919050565b80356001600160a01b038116811461098a57600080fd5b60008083601f8401126109b857600080fd5b5081356001600160401b038111156109cf57600080fd5b6020830191508360208285010111156109e757600080fd5b9250929050565b60008060008060608587031215610a0457600080fd5b610a0d85610973565b9350610a1b6020860161098f565b925060408501356001600160401b03811115610a3657600080fd5b610a42878288016109a6565b95989497509550505050565b600080600080600060808688031215610a6657600080fd5b610a6f86610973565b9450610a7d6020870161098f565b935060408601356001600160401b03811115610a9857600080fd5b610aa4888289016109a6565b9094509250610ab7905060608701610973565b90509295509295909350565b600060208284031215610ad557600080fd5b61096c82610973565b6000808335601e19843603018112610af557600080fd5b8301803591506001600160401b03821115610b0f57600080fd5b6020019150600581901b36038213156109e757600080fd5b634e487b7160e01b600052603260045260246000fd5b6000823560de19833603018112610b5357600080fd5b9190910192915050565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03818116838216019080821115610b9357610b93610b5d565b5092915050565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715610bd257610bd2610b9a565b60405290565b600082601f830112610be957600080fd5b81356001600160401b0380821115610c0357610c03610b9a565b604051601f8301601f19908116603f01168101908282118183101715610c2b57610c2b610b9a565b81604052838152866020858801011115610c4457600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060e08236031215610c7657600080fd5b610c7e610bb0565b610c8783610973565b8152610c9560208401610973565b6020820152610ca660408401610973565b6040820152610cb76060840161098f565b6060820152610cc86080840161098f565b608082015260a08301356001600160401b03811115610ce657600080fd5b610cf236828601610bd8565b60a083015250610d0460c08401610973565b60c082015292915050565b60005b83811015610d2a578181015183820152602001610d12565b50506000910152565b60008151808452610d4b816020860160208601610d0f565b601f01601f19169290920160200192915050565b8481526001600160a01b03841660208201528215156040820152608060608201819052600090610d9190830184610d33565b9695505050505050565b6001600160a01b0386811682528516602082015260806040820181905281018390526000838560a0840137600060a0858401015260a0601f19601f86011683010190506001600160401b03831660608301529695505050505050565b60008251610b53818460208701610d0f565b81810381811115610e1c57610e1c610b5d565b92915050565b60208152600061096c6020830184610d3356fea26469706673582212205c348a5a884ce192bc444dbba7605690244b82b64c093a91a917a1010254a34664736f6c63430008170033",
}

// OmniPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniPortalMetaData.ABI instead.
var OmniPortalABI = OmniPortalMetaData.ABI

// OmniPortalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniPortalMetaData.Bin instead.
var OmniPortalBin = OmniPortalMetaData.Bin

// DeployOmniPortal deploys a new Ethereum contract, binding an instance of OmniPortal to it.
func DeployOmniPortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniPortal, error) {
	parsed, err := OmniPortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniPortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// OmniPortal is an auto generated Go binding around an Ethereum contract.
type OmniPortal struct {
	OmniPortalCaller     // Read-only binding to the contract
	OmniPortalTransactor // Write-only binding to the contract
	OmniPortalFilterer   // Log filterer for contract events
}

// OmniPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniPortalSession struct {
	Contract     *OmniPortal       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniPortalCallerSession struct {
	Contract *OmniPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OmniPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniPortalTransactorSession struct {
	Contract     *OmniPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OmniPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniPortalRaw struct {
	Contract *OmniPortal // Generic contract binding to access the raw methods on
}

// OmniPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniPortalCallerRaw struct {
	Contract *OmniPortalCaller // Generic read-only contract binding to access the raw methods on
}

// OmniPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniPortalTransactorRaw struct {
	Contract *OmniPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniPortal creates a new instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortal(address common.Address, backend bind.ContractBackend) (*OmniPortal, error) {
	contract, err := bindOmniPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// NewOmniPortalCaller creates a new read-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalCaller(address common.Address, caller bind.ContractCaller) (*OmniPortalCaller, error) {
	contract, err := bindOmniPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalCaller{contract: contract}, nil
}

// NewOmniPortalTransactor creates a new write-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniPortalTransactor, error) {
	contract, err := bindOmniPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalTransactor{contract: contract}, nil
}

// NewOmniPortalFilterer creates a new log filterer instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniPortalFilterer, error) {
	contract, err := bindOmniPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFilterer{contract: contract}, nil
}

// bindOmniPortal binds a generic wrapper to an already deployed contract.
func bindOmniPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OmniPortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.OmniPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transact(opts, method, params...)
}

// XMSGDEFAULTGASLIMIT is a free data retrieval call binding the contract method 0x9dad9aae.
//
// Solidity: function XMSG_DEFAULT_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XMSGDEFAULTGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XMSG_DEFAULT_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XMSGDEFAULTGASLIMIT is a free data retrieval call binding the contract method 0x9dad9aae.
//
// Solidity: function XMSG_DEFAULT_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XMSGDEFAULTGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGDEFAULTGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGDEFAULTGASLIMIT is a free data retrieval call binding the contract method 0x9dad9aae.
//
// Solidity: function XMSG_DEFAULT_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XMSGDEFAULTGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGDEFAULTGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGMAXGASLIMIT is a free data retrieval call binding the contract method 0xa2cc111b.
//
// Solidity: function XMSG_MAX_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XMSGMAXGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XMSG_MAX_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XMSGMAXGASLIMIT is a free data retrieval call binding the contract method 0xa2cc111b.
//
// Solidity: function XMSG_MAX_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XMSGMAXGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGMAXGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGMAXGASLIMIT is a free data retrieval call binding the contract method 0xa2cc111b.
//
// Solidity: function XMSG_MAX_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XMSGMAXGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGMAXGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGMINGASLIMIT is a free data retrieval call binding the contract method 0xfa590d14.
//
// Solidity: function XMSG_MIN_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XMSGMINGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XMSG_MIN_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XMSGMINGASLIMIT is a free data retrieval call binding the contract method 0xfa590d14.
//
// Solidity: function XMSG_MIN_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XMSGMINGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGMINGASLIMIT(&_OmniPortal.CallOpts)
}

// XMSGMINGASLIMIT is a free data retrieval call binding the contract method 0xfa590d14.
//
// Solidity: function XMSG_MIN_GAS_LIMIT() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XMSGMINGASLIMIT() (uint64, error) {
	return _OmniPortal.Contract.XMSGMINGASLIMIT(&_OmniPortal.CallOpts)
}

// XRECEIPTMAXRETURNDATASIZE is a free data retrieval call binding the contract method 0x01665569.
//
// Solidity: function XRECEIPT_MAX_RETURN_DATA_SIZE() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XRECEIPTMAXRETURNDATASIZE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XRECEIPT_MAX_RETURN_DATA_SIZE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XRECEIPTMAXRETURNDATASIZE is a free data retrieval call binding the contract method 0x01665569.
//
// Solidity: function XRECEIPT_MAX_RETURN_DATA_SIZE() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XRECEIPTMAXRETURNDATASIZE() (uint64, error) {
	return _OmniPortal.Contract.XRECEIPTMAXRETURNDATASIZE(&_OmniPortal.CallOpts)
}

// XRECEIPTMAXRETURNDATASIZE is a free data retrieval call binding the contract method 0x01665569.
//
// Solidity: function XRECEIPT_MAX_RETURN_DATA_SIZE() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XRECEIPTMAXRETURNDATASIZE() (uint64, error) {
	return _OmniPortal.Contract.XRECEIPTMAXRETURNDATASIZE(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// InXStreamOffset is a free data retrieval call binding the contract method 0xb58e964f.
//
// Solidity: function inXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXStreamOffset(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXStreamOffset", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXStreamOffset is a free data retrieval call binding the contract method 0xb58e964f.
//
// Solidity: function inXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXStreamOffset(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXStreamOffset(&_OmniPortal.CallOpts, arg0)
}

// InXStreamOffset is a free data retrieval call binding the contract method 0xb58e964f.
//
// Solidity: function inXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXStreamOffset(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXStreamOffset(&_OmniPortal.CallOpts, arg0)
}

// OutXStreamOffset is a free data retrieval call binding the contract method 0x90ab417c.
//
// Solidity: function outXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OutXStreamOffset(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "outXStreamOffset", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutXStreamOffset is a free data retrieval call binding the contract method 0x90ab417c.
//
// Solidity: function outXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) OutXStreamOffset(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXStreamOffset(&_OmniPortal.CallOpts, arg0)
}

// OutXStreamOffset is a free data retrieval call binding the contract method 0x90ab417c.
//
// Solidity: function outXStreamOffset(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OutXStreamOffset(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXStreamOffset(&_OmniPortal.CallOpts, arg0)
}

// Xcall is a paid mutator transaction binding the contract method 0x50e646dd.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data) payable returns()
func (_OmniPortal *OmniPortalTransactor) Xcall(opts *bind.TransactOpts, destChainId uint64, to common.Address, data []byte) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xcall", destChainId, to, data)
}

// Xcall is a paid mutator transaction binding the contract method 0x50e646dd.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data) payable returns()
func (_OmniPortal *OmniPortalSession) Xcall(destChainId uint64, to common.Address, data []byte) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, to, data)
}

// Xcall is a paid mutator transaction binding the contract method 0x50e646dd.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data) payable returns()
func (_OmniPortal *OmniPortalTransactorSession) Xcall(destChainId uint64, to common.Address, data []byte) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, to, data)
}

// Xcall0 is a paid mutator transaction binding the contract method 0x70e8b56a.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactor) Xcall0(opts *bind.TransactOpts, destChainId uint64, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xcall0", destChainId, to, data, gasLimit)
}

// Xcall0 is a paid mutator transaction binding the contract method 0x70e8b56a.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalSession) Xcall0(destChainId uint64, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall0(&_OmniPortal.TransactOpts, destChainId, to, data, gasLimit)
}

// Xcall0 is a paid mutator transaction binding the contract method 0x70e8b56a.
//
// Solidity: function xcall(uint64 destChainId, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactorSession) Xcall0(destChainId uint64, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall0(&_OmniPortal.TransactOpts, destChainId, to, data, gasLimit)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x06f9f174.
//
// Solidity: function xsubmit((bytes32,(uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(bytes,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactor) Xsubmit(opts *bind.TransactOpts, xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xsubmit", xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x06f9f174.
//
// Solidity: function xsubmit((bytes32,(uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(bytes,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x06f9f174.
//
// Solidity: function xsubmit((bytes32,(uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(bytes,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactorSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// OmniPortalXMsgIterator is returned from FilterXMsg and is used to iterate over the raw logs and unpacked data for XMsg events raised by the OmniPortal contract.
type OmniPortalXMsgIterator struct {
	Event *OmniPortalXMsg // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsg)
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
		it.Event = new(OmniPortalXMsg)
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
func (it *OmniPortalXMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsg represents a XMsg event raised by the OmniPortal contract.
type OmniPortalXMsg struct {
	DestChainId  uint64
	StreamOffset uint64
	Sender       common.Address
	To           common.Address
	Data         []byte
	GasLimit     uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterXMsg is a free log retrieval operation binding the contract event 0xac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) FilterXMsg(opts *bind.FilterOpts, destChainId []uint64, streamOffset []uint64) (*OmniPortalXMsgIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var streamOffsetRule []interface{}
	for _, streamOffsetItem := range streamOffset {
		streamOffsetRule = append(streamOffsetRule, streamOffsetItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsg", destChainIdRule, streamOffsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgIterator{contract: _OmniPortal.contract, event: "XMsg", logs: logs, sub: sub}, nil
}

// WatchXMsg is a free log subscription operation binding the contract event 0xac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) WatchXMsg(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsg, destChainId []uint64, streamOffset []uint64) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var streamOffsetRule []interface{}
	for _, streamOffsetItem := range streamOffset {
		streamOffsetRule = append(streamOffsetRule, streamOffsetItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsg", destChainIdRule, streamOffsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsg)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
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

// ParseXMsg is a log parse operation binding the contract event 0xac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) ParseXMsg(log types.Log) (*OmniPortalXMsg, error) {
	event := new(OmniPortalXMsg)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXReceiptIterator is returned from FilterXReceipt and is used to iterate over the raw logs and unpacked data for XReceipt events raised by the OmniPortal contract.
type OmniPortalXReceiptIterator struct {
	Event *OmniPortalXReceipt // Event containing the contract specifics and raw log

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
func (it *OmniPortalXReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXReceipt)
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
		it.Event = new(OmniPortalXReceipt)
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
func (it *OmniPortalXReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXReceipt represents a XReceipt event raised by the OmniPortal contract.
type OmniPortalXReceipt struct {
	SourceChainId uint64
	StreamOffset  uint64
	GasUsed       *big.Int
	Relayer       common.Address
	Success       bool
	ReturnData    []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXReceipt is a free log retrieval operation binding the contract event 0xc2f03986562c9333b7872c7694832b7583b6a4dbde9868c9d43a2a3896c64343.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success, bytes returnData)
func (_OmniPortal *OmniPortalFilterer) FilterXReceipt(opts *bind.FilterOpts, sourceChainId []uint64, streamOffset []uint64) (*OmniPortalXReceiptIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var streamOffsetRule []interface{}
	for _, streamOffsetItem := range streamOffset {
		streamOffsetRule = append(streamOffsetRule, streamOffsetItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XReceipt", sourceChainIdRule, streamOffsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXReceiptIterator{contract: _OmniPortal.contract, event: "XReceipt", logs: logs, sub: sub}, nil
}

// WatchXReceipt is a free log subscription operation binding the contract event 0xc2f03986562c9333b7872c7694832b7583b6a4dbde9868c9d43a2a3896c64343.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success, bytes returnData)
func (_OmniPortal *OmniPortalFilterer) WatchXReceipt(opts *bind.WatchOpts, sink chan<- *OmniPortalXReceipt, sourceChainId []uint64, streamOffset []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var streamOffsetRule []interface{}
	for _, streamOffsetItem := range streamOffset {
		streamOffsetRule = append(streamOffsetRule, streamOffsetItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XReceipt", sourceChainIdRule, streamOffsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXReceipt)
				if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
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

// ParseXReceipt is a log parse operation binding the contract event 0xc2f03986562c9333b7872c7694832b7583b6a4dbde9868c9d43a2a3896c64343.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success, bytes returnData)
func (_OmniPortal *OmniPortalFilterer) ParseXReceipt(log types.Log) (*OmniPortalXReceipt, error) {
	event := new(OmniPortalXReceipt)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
