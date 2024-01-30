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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"XMSG_DEFAULT_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XMSG_MAX_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XMSG_MIN_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXStreamOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outXStreamOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xsubmit\",\"inputs\":[{\"name\":\"xsub\",\"type\":\"tuple\",\"internalType\":\"structXTypes.Submission\",\"components\":[{\"name\":\"attestationRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockHeader\",\"type\":\"tuple\",\"internalType\":\"structXTypes.BlockHeader\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Msg[]\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlags\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.SigTuple[]\",\"components\":[{\"name\":\"validatorPubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"XMsg\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceipt\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"streamOffset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MerkleProofInvalidMultiproof\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561001057600080fd5b506001600160401b0346166080526080516111096100476000396000818161013701528181610375015261070e01526111096000f3fe6080604052600436106100865760003560e01c80639a8a0592116100595780639a8a0592146101255780639dad9aae14610159578063a2cc111b14610170578063b58e964f14610187578063fa590d14146101bd57600080fd5b806306f9f1741461008b57806350e646dd146100ad57806370e8b56a146100c057806390ab417c146100d3575b600080fd5b34801561009757600080fd5b506100ab6100a6366004610c11565b6101d3565b005b6100ab6100bb366004610cc7565b6102b5565b6100ab6100ce366004610d27565b6102cc565b3480156100df57600080fd5b506101096100ee366004610d9c565b6000602081905290815260409020546001600160401b031681565b6040516001600160401b03909116815260200160405180910390f35b34801561013157600080fd5b506101097f000000000000000000000000000000000000000000000000000000000000000081565b34801561016557600080fd5b5061010962030d4081565b34801561017c57600080fd5b50610109624c4b4081565b34801561019357600080fd5b506101096101a2366004610d9c565b6001602052600090815260409020546001600160401b031681565b3480156101c957600080fd5b5061010961520881565b6102088135602083016101e96080850185610db7565b6101f660a0870187610db7565b61020360c0890189610db7565b6102e1565b6102595760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a20696e76616c69642070726f6f660000000000000060448201526064015b60405180910390fd5b60005b6102696080830183610db7565b90508110156102b1576102a96102826080840184610db7565b8381811061029257610292610e00565b90506020028101906102a49190610e16565b61036b565b60010161025c565b5050565b6102c6843385858562030d40610655565b50505050565b6102da853386868686610655565b5050505050565b600061035e858580806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160208089028281018201909352888252909350889250879182918501908490808284376000920191909152508e925061035991508d90508c8c610844565b610959565b9998505050505050505050565b6001600160401b037f0000000000000000000000000000000000000000000000000000000000000000166103a56040830160208401610d9c565b6001600160401b0316146103fb5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a2077726f6e672064657374436861696e49640000006044820152606401610250565b6001600061040c6020840184610d9c565b6001600160401b03908116825260208201929092526040908101600020549091169061043e9060608401908401610d9c565b6001600160401b0316146104945760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a2077726f6e672073747265616d4f666673657400006044820152606401610250565b60018060006104a66020850185610d9c565b6001600160401b03908116825260208201929092526040016000908120805490926104d391859116610e4c565b92506101000a8154816001600160401b0302191690836001600160401b031602179055506000624c4b406001600160401b03168260c00160208101906105199190610d9c565b6001600160401b03161161053c5761053760e0830160c08401610d9c565b610541565b624c4b405b6001600160401b0316905060005a9050600061056360a0850160808601610e73565b6001600160a01b03168361057a60a0870187610e8e565b604051610588929190610ed4565b60006040518083038160008787f1925050503d80600081146105c6576040519150601f19603f3d011682016040523d82523d6000602084013e6105cb565b606091505b505090505a6105da9083610ee4565b91506105ec6060850160408601610d9c565b6001600160401b03166106026020860186610d9c565b604080518581523360208201528415158183015290516001600160401b0392909216917f34515b4105a7bb34f3af3cd490137ab292bb2ff14efb800df5c7d59e28944f259181900360600190a350505050565b624c4b406001600160401b03821611156106b15760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f20686967680000006044820152606401610250565b6152086001600160401b038216101561070c5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f206c6f77000000006044820152606401610250565b7f00000000000000000000000000000000000000000000000000000000000000006001600160401b0316866001600160401b03160361078d5760405162461bcd60e51b815260206004820152601f60248201527f4f6d6e69506f7274616c3a206e6f2073616d652d636861696e207863616c6c006044820152606401610250565b6001600160401b03808716600081815260208190526040908190205490519216917fac3afbbff5be7c4af1610721cf4793840bd167251fd6f184ee708f752a731283906107e39089908990899089908990610f20565b60405180910390a36001600160401b038087166000908152602081905260408120805460019391929161081891859116610e4c565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550505050505050565b60606000610853836001610f68565b6001600160401b0381111561086a5761086a610f7b565b604051908082528060200260200182016040528015610893578160200160208202803683370190505b5090506108be856040516020016108aa9190610f91565b604051602081830303815290604052610971565b816000815181106108d1576108d1610e00565b60200260200101818152505060005b8381101561094e5761091f8585838181106108fd576108fd610e00565b905060200281019061090f9190610e16565b6040516020016108aa9190610fce565b8261092b836001610f68565b8151811061093b5761093b610e00565b60209081029190910101526001016108e0565b5090505b9392505050565b6000826109678686856109aa565b1495945050505050565b6000818051906020012060405160200161098d91815260200190565b604051602081830303815290604052805190602001209050919050565b80518351835160009291906109c0816001610f68565b6109ca8385610f68565b146109e857604051631a8a024960e11b815260040160405180910390fd5b6000816001600160401b03811115610a0257610a02610f7b565b604051908082528060200260200182016040528015610a2b578160200160208202803683370190505b5090506000806000805b85811015610b5f576000888510610a70578584610a51816110ba565b955081518110610a6357610a63610e00565b6020026020010151610a96565b8a85610a7b816110ba565b965081518110610a8d57610a8d610e00565b60200260200101515b905060008c8381518110610aac57610aac610e00565b6020026020010151610ae2578d84610ac3816110ba565b955081518110610ad557610ad5610e00565b6020026020010151610b2c565b898610610b06578685610af4816110ba565b965081518110610ad557610ad5610e00565b8b86610b11816110ba565b975081518110610b2357610b23610e00565b60200260200101515b9050610b388282610bdd565b878481518110610b4a57610b4a610e00565b60209081029190910101525050600101610a35565b508415610bb157858114610b8657604051631a8a024960e11b815260040160405180910390fd5b836001860381518110610b9b57610b9b610e00565b6020026020010151975050505050505050610952565b8615610bca5788600081518110610b9b57610b9b610e00565b8a600081518110610b9b57610b9b610e00565b6000818310610bf9576000828152602084905260409020610c08565b60008381526020839052604090205b90505b92915050565b600060208284031215610c2357600080fd5b81356001600160401b03811115610c3957600080fd5b8201610100818503121561095257600080fd5b80356001600160401b0381168114610c6357600080fd5b919050565b80356001600160a01b0381168114610c6357600080fd5b60008083601f840112610c9157600080fd5b5081356001600160401b03811115610ca857600080fd5b602083019150836020828501011115610cc057600080fd5b9250929050565b60008060008060608587031215610cdd57600080fd5b610ce685610c4c565b9350610cf460208601610c68565b925060408501356001600160401b03811115610d0f57600080fd5b610d1b87828801610c7f565b95989497509550505050565b600080600080600060808688031215610d3f57600080fd5b610d4886610c4c565b9450610d5660208701610c68565b935060408601356001600160401b03811115610d7157600080fd5b610d7d88828901610c7f565b9094509250610d90905060608701610c4c565b90509295509295909350565b600060208284031215610dae57600080fd5b610c0882610c4c565b6000808335601e19843603018112610dce57600080fd5b8301803591506001600160401b03821115610de857600080fd5b6020019150600581901b3603821315610cc057600080fd5b634e487b7160e01b600052603260045260246000fd5b6000823560de19833603018112610e2c57600080fd5b9190910192915050565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03818116838216019080821115610e6c57610e6c610e36565b5092915050565b600060208284031215610e8557600080fd5b610c0882610c68565b6000808335601e19843603018112610ea557600080fd5b8301803591506001600160401b03821115610ebf57600080fd5b602001915036819003821315610cc057600080fd5b8183823760009101908152919050565b81810381811115610c0b57610c0b610e36565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03868116825285166020820152608060408201819052600090610f4d9083018587610ef7565b90506001600160401b03831660608301529695505050505050565b80820180821115610c0b57610c0b610e36565b634e487b7160e01b600052604160045260246000fd5b606081016001600160401b0380610fa785610c4c565b16835280610fb760208601610c4c565b166020840152506040830135604083015292915050565b6020815260006001600160401b0380610fe685610c4c565b16602084015280610ff960208601610c4c565b1660408401528061100c60408601610c4c565b16606084015261101e60608501610c68565b60018060a01b0380821660808601528061103a60808801610c68565b1660a0860152505060a0840135601e1985360301811261105957600080fd5b84016020810190358281111561106e57600080fd5b80360382131561107d57600080fd5b60e060c086015261109361010086018284610ef7565b925050506110a360c08501610c4c565b6001600160401b03811660e0850152509392505050565b6000600182016110cc576110cc610e36565b506001019056fea26469706673582212208442d9d19fd697700accc727f15cce00d9506a91ddbd3297ff528cfb6d610b7864736f6c63430008170033",
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
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXReceipt is a free log retrieval operation binding the contract event 0x34515b4105a7bb34f3af3cd490137ab292bb2ff14efb800df5c7d59e28944f25.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success)
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

// WatchXReceipt is a free log subscription operation binding the contract event 0x34515b4105a7bb34f3af3cd490137ab292bb2ff14efb800df5c7d59e28944f25.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success)
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

// ParseXReceipt is a log parse operation binding the contract event 0x34515b4105a7bb34f3af3cd490137ab292bb2ff14efb800df5c7d59e28944f25.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success)
func (_OmniPortal *OmniPortalFilterer) ParseXReceipt(log types.Log) (*OmniPortalXReceipt, error) {
	event := new(OmniPortalXReceipt)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
