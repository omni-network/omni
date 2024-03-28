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

// IOmniXChainRegistryChain is an auto generated low-level Go binding around an user-defined struct.
type IOmniXChainRegistryChain struct {
	ChainId      uint64
	Name         string
	Portal       common.Address
	DeployHeight *big.Int
}

// OmniXChainRegistryMetaData contains all meta data concerning the OmniXChainRegistry contract.
var OmniXChainRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addChain\",\"inputs\":[{\"name\":\"c\",\"type\":\"tuple\",\"internalType\":\"structIOmniXChainRegistry.Chain\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getChain\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOmniXChainRegistry.Chain\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOmniXChainRegistry.Chain[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initPortal\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"portal\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061177c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638da5cb5b116100665780638da5cb5b146100e6578063a1a6d50814610101578063c2a1402d14610121578063c4d66de814610144578063f2fde38b1461015757600080fd5b806312465973146100985780631d8dce4d146100ad578063331b062a146100c0578063715018a6146100de575b600080fd5b6100ab6100a63660046110d0565b61016a565b005b6100ab6100bb36600461111f565b610448565b6100c86106f1565b6040516100d591906111e1565b60405180910390f35b6100ab610887565b6033546040516001600160a01b0390911681526020016100d5565b61011461010f36600461125a565b61089b565b6040516100d59190611277565b61013461012f36600461125a565b61098c565b60405190151581526020016100d5565b6100ab61015236600461128a565b6109a8565b6100ab61016536600461128a565b610ac3565b610172610b3c565b6065546001600160a01b03166101cf5760405162461bcd60e51b815260206004820152601960248201527f5852656769737472793a20706f7274616c206e6f74207365740000000000000060448201526064015b60405180910390fd5b60006101e1606083016040840161128a565b6001600160a01b0316036102375760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a20706f7274616c206973207a65726f2061646472000060448201526064016101c6565b600081606001351161028b5760405162461bcd60e51b815260206004820152601f60248201527f5852656769737472793a206465706c6f79486569676874206973207a65726f0060448201526064016101c6565b6102a061029b602083018361125a565b610b96565b6065546001600160a01b031663e2947e546102be602084018461125a565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401600060405180830381600087803b1580156102ff57600080fd5b505af1158015610313573d6000803e3d6000fd5b506103389250610329915050602083018361125a565b6103336066610d6b565b610d7f565b610350610348602083018361125a565b606690610e91565b61039c5760405162461bcd60e51b815260206004820152601f60248201527f5852656769737472793a20636861696e20616c7265616479206578697374730060448201526064016101c6565b80606860006103ae602084018461125a565b6001600160401b0316815260208101919091526040016000206103d18282611458565b506103e19050602082018261125a565b6001600160401b03167f1ec5b8f876d132a91a4fe43d93a3c5d59f4ae25983b63bce6aa647b692bd76e661041860208401846112a7565b610428606086016040870161128a565b856060013560405161043d9493929190611511565b60405180910390a250565b610450610b3c565b6001600160a01b0382166104a65760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a20706f7274616c206973207a65726f2061646472000060448201526064016101c6565b6065546001600160a01b0316156104ff5760405162461bcd60e51b815260206004820152601d60248201527f5852656769737472793a20706f7274616c20616c72656164792073657400000060448201526064016101c6565b6000811161054f5760405162461bcd60e51b815260206004820152601f60248201527f5852656769737472793a206465706c6f79486569676874206973207a65726f0060448201526064016101c6565b43811061059e5760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a20667574757265206465706c6f79486569676874000060448201526064016101c6565b606580546001600160a01b0319166001600160a01b03841690811790915560408051608081018252466001600160401b031681528151808301835260048152636f6d6e6960e01b6020808301919091528201529081019190915260608101829052805161060d90606690610e91565b5080516001600160401b0390811660009081526068602090815260409091208351815467ffffffffffffffff191693169290921782558201518291906001820190610658908261155a565b506040828101516002830180546001600160a01b0319166001600160a01b03928316179055606090930151600390920191909155606554835191516338a51f9560e21b81526001600160401b0390921660048301529091169063e2947e5490602401600060405180830381600087803b1580156106d457600080fd5b505af11580156106e8573d6000803e3d6000fd5b50505050505050565b606060006106ff6066610ea6565b6001600160401b03811115610716576107166112f4565b60405190808252806020026020018201604052801561074f57816020015b61073c611096565b8152602001906001900390816107345790505b50905060005b61075f6066610ea6565b8110156108815760686000610775606684610eb1565b6001600160401b03908116825260208083019390935260409182016000208251608081019093528054909116825260018101805492939192918401916107ba9061130a565b80601f01602080910402602001604051908101604052809291908181526020018280546107e69061130a565b80156108335780601f1061080857610100808354040283529160200191610833565b820191906000526020600020905b81548152906001019060200180831161081657829003601f168201915b505050918352505060028201546001600160a01b03166020820152600390910154604090910152825183908390811061086e5761086e611619565b6020908102919091010152600101610755565b50919050565b61088f610b3c565b6108996000610ebd565b565b6108a3611096565b6001600160401b0380831660009081526068602090815260409182902082516080810190935280549093168252600183018054929392918401916108e69061130a565b80601f01602080910402602001604051908101604052809291908181526020018280546109129061130a565b801561095f5780601f106109345761010080835404028352916020019161095f565b820191906000526020600020905b81548152906001019060200180831161094257829003601f168201915b505050918352505060028201546001600160a01b0316602082015260039091015460409091015292915050565b60006109a260666001600160401b038416610f0f565b92915050565b600054610100900460ff16158080156109c85750600054600160ff909116105b806109e25750303b1580156109e2575060005460ff166001145b610a455760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016101c6565b6000805460ff191660011790558015610a68576000805461ff0019166101001790555b610a7182610ebd565b610a79610f1b565b8015610abf576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b610acb610b3c565b6001600160a01b038116610b305760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101c6565b610b3981610ebd565b50565b6033546001600160a01b031633146108995760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101c6565b60005b610ba36066610ea6565b811015610abf576000610bb7606683610eb1565b9050606560009054906101000a90046001600160a01b03166001600160a01b0316639a8a05926040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c30919061162f565b6001600160401b0316816001600160401b031603610c4e5750610d63565b60655460408051632d9aa34960e01b815290516001600160a01b03909216916370e8b56a9184918491632d9aa3499160048083019260209291908290030181865afa158015610ca1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc5919061164c565b6040516001600160401b03881660248201526338a51f9560e21b9060440160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e086901b9092168252610d2f93929162030d4090600401611669565b600060405180830381600087803b158015610d4957600080fd5b505af1158015610d5d573d6000803e3d6000fd5b50505050505b600101610b99565b60606000610d7883610f4a565b9392505050565b60655460408051632d9aa34960e01b815290516001600160a01b03909216916370e8b56a9185918491632d9aa3499160048083019260209291908290030181865afa158015610dd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df6919061164c565b604051627e23a960e31b90610e0f9087906024016116ae565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e086901b9092168252610e5b93929162030d4090600401611669565b600060405180830381600087803b158015610e7557600080fd5b505af1158015610e89573d6000803e3d6000fd5b505050505050565b6000610d78836001600160401b038416610f57565b60006109a282610f63565b6000610d788383610f6d565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610d788383610f79565b600054610100900460ff16610f425760405162461bcd60e51b81526004016101c6906116fb565b610899610f91565b60606000610d7883610fc1565b6000610d78838361101d565b60006109a2825490565b6000610d78838361106c565b60008181526001830160205260408120541515610d78565b600054610100900460ff16610fb85760405162461bcd60e51b81526004016101c6906116fb565b61089933610ebd565b60608160000180548060200260200160405190810160405280929190818152602001828054801561101157602002820191906000526020600020905b815481526020019060010190808311610ffd575b50505050509050919050565b6000818152600183016020526040812054611064575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556109a2565b5060006109a2565b600082600001828154811061108357611083611619565b9060005260206000200154905092915050565b604051806080016040528060006001600160401b031681526020016060815260200160006001600160a01b03168152602001600081525090565b6000602082840312156110e257600080fd5b81356001600160401b038111156110f857600080fd5b820160808185031215610d7857600080fd5b6001600160a01b0381168114610b3957600080fd5b6000806040838503121561113257600080fd5b823561113d8161110a565b946020939093013593505050565b6000815180845260005b8181101561117157602081850181015186830182015201611155565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160401b03815116825260006020820151608060208501526111b9608085018261114b565b6040848101516001600160a01b03169086015260609384015193909401929092525090919050565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b8281101561123857603f19888603018452611226858351611191565b9450928501929085019060010161120a565b5092979650505050505050565b6001600160401b0381168114610b3957600080fd5b60006020828403121561126c57600080fd5b8135610d7881611245565b602081526000610d786020830184611191565b60006020828403121561129c57600080fd5b8135610d788161110a565b6000808335601e198436030181126112be57600080fd5b8301803591506001600160401b038211156112d857600080fd5b6020019150368190038213156112ed57600080fd5b9250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061131e57607f821691505b60208210810361088157634e487b7160e01b600052602260045260246000fd5b601f821115611386576000816000526020600020601f850160051c810160208610156113675750805b601f850160051c820191505b81811015610e8957828155600101611373565b505050565b6001600160401b038311156113a2576113a26112f4565b6113b6836113b0835461130a565b8361133e565b6000601f8411600181146113ea57600085156113d25750838201355b600019600387901b1c1916600186901b178355611444565b600083815260209020601f19861690835b8281101561141b57868501358255602094850194600190920191016113fb565b50868210156114385760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b600081356109a28161110a565b813561146381611245565b815467ffffffffffffffff19166001600160401b0391821617825560208301359036849003601e1901821261149757600080fd5b908301908135818111156114aa57600080fd5b6020830192508036038313156114bf57600080fd5b6114cd81846001870161138b565b5050506115036114df6040840161144b565b6002830180546001600160a01b0319166001600160a01b0392909216919091179055565b606082013560038201555050565b6060815283606082015283856080830137600060808583018101919091526001600160a01b039390931660208201526040810191909152601f909201601f191690910101919050565b81516001600160401b03811115611573576115736112f4565b61158781611581845461130a565b8461133e565b602080601f8311600181146115bc57600084156115a45750858301515b600019600386901b1c1916600185901b178555610e89565b600085815260208120601f198616915b828110156115eb578886015182559484019460019091019084016115cc565b50858210156116095787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561164157600080fd5b8151610d7881611245565b60006020828403121561165e57600080fd5b8151610d788161110a565b60006001600160401b03808716835260018060a01b03861660208401526080604084015261169a608084018661114b565b915080841660608401525095945050505050565b6020808252825182820181905260009190848201906040850190845b818110156116ef5783516001600160401b0316835292840192918401916001016116ca565b50909695505050505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fea26469706673582212209f6a1459713598fcdd4fef8640ba44dec4a37d790499206e46bfd276afc305e964736f6c63430008180033",
}

// OmniXChainRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniXChainRegistryMetaData.ABI instead.
var OmniXChainRegistryABI = OmniXChainRegistryMetaData.ABI

// OmniXChainRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniXChainRegistryMetaData.Bin instead.
var OmniXChainRegistryBin = OmniXChainRegistryMetaData.Bin

// DeployOmniXChainRegistry deploys a new Ethereum contract, binding an instance of OmniXChainRegistry to it.
func DeployOmniXChainRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniXChainRegistry, error) {
	parsed, err := OmniXChainRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniXChainRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniXChainRegistry{OmniXChainRegistryCaller: OmniXChainRegistryCaller{contract: contract}, OmniXChainRegistryTransactor: OmniXChainRegistryTransactor{contract: contract}, OmniXChainRegistryFilterer: OmniXChainRegistryFilterer{contract: contract}}, nil
}

// OmniXChainRegistry is an auto generated Go binding around an Ethereum contract.
type OmniXChainRegistry struct {
	OmniXChainRegistryCaller     // Read-only binding to the contract
	OmniXChainRegistryTransactor // Write-only binding to the contract
	OmniXChainRegistryFilterer   // Log filterer for contract events
}

// OmniXChainRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniXChainRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniXChainRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniXChainRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniXChainRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniXChainRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniXChainRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniXChainRegistrySession struct {
	Contract     *OmniXChainRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OmniXChainRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniXChainRegistryCallerSession struct {
	Contract *OmniXChainRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OmniXChainRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniXChainRegistryTransactorSession struct {
	Contract     *OmniXChainRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OmniXChainRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniXChainRegistryRaw struct {
	Contract *OmniXChainRegistry // Generic contract binding to access the raw methods on
}

// OmniXChainRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniXChainRegistryCallerRaw struct {
	Contract *OmniXChainRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// OmniXChainRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniXChainRegistryTransactorRaw struct {
	Contract *OmniXChainRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniXChainRegistry creates a new instance of OmniXChainRegistry, bound to a specific deployed contract.
func NewOmniXChainRegistry(address common.Address, backend bind.ContractBackend) (*OmniXChainRegistry, error) {
	contract, err := bindOmniXChainRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniXChainRegistry{OmniXChainRegistryCaller: OmniXChainRegistryCaller{contract: contract}, OmniXChainRegistryTransactor: OmniXChainRegistryTransactor{contract: contract}, OmniXChainRegistryFilterer: OmniXChainRegistryFilterer{contract: contract}}, nil
}

// NewOmniXChainRegistryCaller creates a new read-only instance of OmniXChainRegistry, bound to a specific deployed contract.
func NewOmniXChainRegistryCaller(address common.Address, caller bind.ContractCaller) (*OmniXChainRegistryCaller, error) {
	contract, err := bindOmniXChainRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniXChainRegistryCaller{contract: contract}, nil
}

// NewOmniXChainRegistryTransactor creates a new write-only instance of OmniXChainRegistry, bound to a specific deployed contract.
func NewOmniXChainRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniXChainRegistryTransactor, error) {
	contract, err := bindOmniXChainRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniXChainRegistryTransactor{contract: contract}, nil
}

// NewOmniXChainRegistryFilterer creates a new log filterer instance of OmniXChainRegistry, bound to a specific deployed contract.
func NewOmniXChainRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniXChainRegistryFilterer, error) {
	contract, err := bindOmniXChainRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniXChainRegistryFilterer{contract: contract}, nil
}

// bindOmniXChainRegistry binds a generic wrapper to an already deployed contract.
func bindOmniXChainRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniXChainRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniXChainRegistry *OmniXChainRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniXChainRegistry.Contract.OmniXChainRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniXChainRegistry *OmniXChainRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.OmniXChainRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniXChainRegistry *OmniXChainRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.OmniXChainRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniXChainRegistry *OmniXChainRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniXChainRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniXChainRegistry *OmniXChainRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniXChainRegistry *OmniXChainRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetChain is a free data retrieval call binding the contract method 0xa1a6d508.
//
// Solidity: function getChain(uint64 chainId) view returns((uint64,string,address,uint256))
func (_OmniXChainRegistry *OmniXChainRegistryCaller) GetChain(opts *bind.CallOpts, chainId uint64) (IOmniXChainRegistryChain, error) {
	var out []interface{}
	err := _OmniXChainRegistry.contract.Call(opts, &out, "getChain", chainId)

	if err != nil {
		return *new(IOmniXChainRegistryChain), err
	}

	out0 := *abi.ConvertType(out[0], new(IOmniXChainRegistryChain)).(*IOmniXChainRegistryChain)

	return out0, err

}

// GetChain is a free data retrieval call binding the contract method 0xa1a6d508.
//
// Solidity: function getChain(uint64 chainId) view returns((uint64,string,address,uint256))
func (_OmniXChainRegistry *OmniXChainRegistrySession) GetChain(chainId uint64) (IOmniXChainRegistryChain, error) {
	return _OmniXChainRegistry.Contract.GetChain(&_OmniXChainRegistry.CallOpts, chainId)
}

// GetChain is a free data retrieval call binding the contract method 0xa1a6d508.
//
// Solidity: function getChain(uint64 chainId) view returns((uint64,string,address,uint256))
func (_OmniXChainRegistry *OmniXChainRegistryCallerSession) GetChain(chainId uint64) (IOmniXChainRegistryChain, error) {
	return _OmniXChainRegistry.Contract.GetChain(&_OmniXChainRegistry.CallOpts, chainId)
}

// GetChains is a free data retrieval call binding the contract method 0x331b062a.
//
// Solidity: function getChains() view returns((uint64,string,address,uint256)[])
func (_OmniXChainRegistry *OmniXChainRegistryCaller) GetChains(opts *bind.CallOpts) ([]IOmniXChainRegistryChain, error) {
	var out []interface{}
	err := _OmniXChainRegistry.contract.Call(opts, &out, "getChains")

	if err != nil {
		return *new([]IOmniXChainRegistryChain), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOmniXChainRegistryChain)).(*[]IOmniXChainRegistryChain)

	return out0, err

}

// GetChains is a free data retrieval call binding the contract method 0x331b062a.
//
// Solidity: function getChains() view returns((uint64,string,address,uint256)[])
func (_OmniXChainRegistry *OmniXChainRegistrySession) GetChains() ([]IOmniXChainRegistryChain, error) {
	return _OmniXChainRegistry.Contract.GetChains(&_OmniXChainRegistry.CallOpts)
}

// GetChains is a free data retrieval call binding the contract method 0x331b062a.
//
// Solidity: function getChains() view returns((uint64,string,address,uint256)[])
func (_OmniXChainRegistry *OmniXChainRegistryCallerSession) GetChains() ([]IOmniXChainRegistryChain, error) {
	return _OmniXChainRegistry.Contract.GetChains(&_OmniXChainRegistry.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc2a1402d.
//
// Solidity: function isRegistered(uint64 chainId) view returns(bool)
func (_OmniXChainRegistry *OmniXChainRegistryCaller) IsRegistered(opts *bind.CallOpts, chainId uint64) (bool, error) {
	var out []interface{}
	err := _OmniXChainRegistry.contract.Call(opts, &out, "isRegistered", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc2a1402d.
//
// Solidity: function isRegistered(uint64 chainId) view returns(bool)
func (_OmniXChainRegistry *OmniXChainRegistrySession) IsRegistered(chainId uint64) (bool, error) {
	return _OmniXChainRegistry.Contract.IsRegistered(&_OmniXChainRegistry.CallOpts, chainId)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc2a1402d.
//
// Solidity: function isRegistered(uint64 chainId) view returns(bool)
func (_OmniXChainRegistry *OmniXChainRegistryCallerSession) IsRegistered(chainId uint64) (bool, error) {
	return _OmniXChainRegistry.Contract.IsRegistered(&_OmniXChainRegistry.CallOpts, chainId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniXChainRegistry *OmniXChainRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniXChainRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniXChainRegistry *OmniXChainRegistrySession) Owner() (common.Address, error) {
	return _OmniXChainRegistry.Contract.Owner(&_OmniXChainRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniXChainRegistry *OmniXChainRegistryCallerSession) Owner() (common.Address, error) {
	return _OmniXChainRegistry.Contract.Owner(&_OmniXChainRegistry.CallOpts)
}

// AddChain is a paid mutator transaction binding the contract method 0x12465973.
//
// Solidity: function addChain((uint64,string,address,uint256) c) returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactor) AddChain(opts *bind.TransactOpts, c IOmniXChainRegistryChain) (*types.Transaction, error) {
	return _OmniXChainRegistry.contract.Transact(opts, "addChain", c)
}

// AddChain is a paid mutator transaction binding the contract method 0x12465973.
//
// Solidity: function addChain((uint64,string,address,uint256) c) returns()
func (_OmniXChainRegistry *OmniXChainRegistrySession) AddChain(c IOmniXChainRegistryChain) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.AddChain(&_OmniXChainRegistry.TransactOpts, c)
}

// AddChain is a paid mutator transaction binding the contract method 0x12465973.
//
// Solidity: function addChain((uint64,string,address,uint256) c) returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactorSession) AddChain(c IOmniXChainRegistryChain) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.AddChain(&_OmniXChainRegistry.TransactOpts, c)
}

// InitPortal is a paid mutator transaction binding the contract method 0x1d8dce4d.
//
// Solidity: function initPortal(address portal, uint256 deployHeight) returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactor) InitPortal(opts *bind.TransactOpts, portal common.Address, deployHeight *big.Int) (*types.Transaction, error) {
	return _OmniXChainRegistry.contract.Transact(opts, "initPortal", portal, deployHeight)
}

// InitPortal is a paid mutator transaction binding the contract method 0x1d8dce4d.
//
// Solidity: function initPortal(address portal, uint256 deployHeight) returns()
func (_OmniXChainRegistry *OmniXChainRegistrySession) InitPortal(portal common.Address, deployHeight *big.Int) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.InitPortal(&_OmniXChainRegistry.TransactOpts, portal, deployHeight)
}

// InitPortal is a paid mutator transaction binding the contract method 0x1d8dce4d.
//
// Solidity: function initPortal(address portal, uint256 deployHeight) returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactorSession) InitPortal(portal common.Address, deployHeight *big.Int) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.InitPortal(&_OmniXChainRegistry.TransactOpts, portal, deployHeight)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _OmniXChainRegistry.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_OmniXChainRegistry *OmniXChainRegistrySession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.Initialize(&_OmniXChainRegistry.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.Initialize(&_OmniXChainRegistry.TransactOpts, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniXChainRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniXChainRegistry *OmniXChainRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.RenounceOwnership(&_OmniXChainRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.RenounceOwnership(&_OmniXChainRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniXChainRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniXChainRegistry *OmniXChainRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.TransferOwnership(&_OmniXChainRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniXChainRegistry *OmniXChainRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniXChainRegistry.Contract.TransferOwnership(&_OmniXChainRegistry.TransactOpts, newOwner)
}

// OmniXChainRegistryChainAddedIterator is returned from FilterChainAdded and is used to iterate over the raw logs and unpacked data for ChainAdded events raised by the OmniXChainRegistry contract.
type OmniXChainRegistryChainAddedIterator struct {
	Event *OmniXChainRegistryChainAdded // Event containing the contract specifics and raw log

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
func (it *OmniXChainRegistryChainAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniXChainRegistryChainAdded)
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
		it.Event = new(OmniXChainRegistryChainAdded)
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
func (it *OmniXChainRegistryChainAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniXChainRegistryChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniXChainRegistryChainAdded represents a ChainAdded event raised by the OmniXChainRegistry contract.
type OmniXChainRegistryChainAdded struct {
	ChainId      uint64
	Name         string
	Portal       common.Address
	DeployHeight *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChainAdded is a free log retrieval operation binding the contract event 0x1ec5b8f876d132a91a4fe43d93a3c5d59f4ae25983b63bce6aa647b692bd76e6.
//
// Solidity: event ChainAdded(uint64 indexed chainId, string name, address portal, uint256 deployHeight)
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) FilterChainAdded(opts *bind.FilterOpts, chainId []uint64) (*OmniXChainRegistryChainAddedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniXChainRegistry.contract.FilterLogs(opts, "ChainAdded", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniXChainRegistryChainAddedIterator{contract: _OmniXChainRegistry.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

// WatchChainAdded is a free log subscription operation binding the contract event 0x1ec5b8f876d132a91a4fe43d93a3c5d59f4ae25983b63bce6aa647b692bd76e6.
//
// Solidity: event ChainAdded(uint64 indexed chainId, string name, address portal, uint256 deployHeight)
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *OmniXChainRegistryChainAdded, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniXChainRegistry.contract.WatchLogs(opts, "ChainAdded", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniXChainRegistryChainAdded)
				if err := _OmniXChainRegistry.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

// ParseChainAdded is a log parse operation binding the contract event 0x1ec5b8f876d132a91a4fe43d93a3c5d59f4ae25983b63bce6aa647b692bd76e6.
//
// Solidity: event ChainAdded(uint64 indexed chainId, string name, address portal, uint256 deployHeight)
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) ParseChainAdded(log types.Log) (*OmniXChainRegistryChainAdded, error) {
	event := new(OmniXChainRegistryChainAdded)
	if err := _OmniXChainRegistry.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniXChainRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniXChainRegistry contract.
type OmniXChainRegistryInitializedIterator struct {
	Event *OmniXChainRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *OmniXChainRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniXChainRegistryInitialized)
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
		it.Event = new(OmniXChainRegistryInitialized)
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
func (it *OmniXChainRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniXChainRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniXChainRegistryInitialized represents a Initialized event raised by the OmniXChainRegistry contract.
type OmniXChainRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniXChainRegistryInitializedIterator, error) {

	logs, sub, err := _OmniXChainRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniXChainRegistryInitializedIterator{contract: _OmniXChainRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniXChainRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _OmniXChainRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniXChainRegistryInitialized)
				if err := _OmniXChainRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) ParseInitialized(log types.Log) (*OmniXChainRegistryInitialized, error) {
	event := new(OmniXChainRegistryInitialized)
	if err := _OmniXChainRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniXChainRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniXChainRegistry contract.
type OmniXChainRegistryOwnershipTransferredIterator struct {
	Event *OmniXChainRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniXChainRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniXChainRegistryOwnershipTransferred)
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
		it.Event = new(OmniXChainRegistryOwnershipTransferred)
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
func (it *OmniXChainRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniXChainRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniXChainRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the OmniXChainRegistry contract.
type OmniXChainRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniXChainRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniXChainRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniXChainRegistryOwnershipTransferredIterator{contract: _OmniXChainRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniXChainRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniXChainRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniXChainRegistryOwnershipTransferred)
				if err := _OmniXChainRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OmniXChainRegistry *OmniXChainRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*OmniXChainRegistryOwnershipTransferred, error) {
	event := new(OmniXChainRegistryOwnershipTransferred)
	if err := _OmniXChainRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
