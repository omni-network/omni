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

// XRegistryBaseDeployment is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type XRegistryBaseDeployment struct {
// 	Addr     common.Address
// 	Metadata []byte
// }

// XRegistryMetaData contains all meta data concerning the XRegistry contract.
var XRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"XSET_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSET_PORTAL_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainIds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"has\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"portal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deployment\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"registrationFee\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dep\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"replicas\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPortal\",\"inputs\":[{\"name\":\"_portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReplica\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"replica\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50611a37806100206000396000f3fe6080604052600436106100e85760003560e01c806388f9380b1161008a578063b3449b7711610059578063b3449b77146102ad578063f2fde38b146102da578063fbe4b7c0146102fa578063fd0b64f71461031c57600080fd5b806388f9380b146101da5780638926f54f146102105780638da5cb5b146102615780638f0d79a71461027f57600080fd5b80636425666b116100c65780636425666b14610156578063715018a61461018e5780637b7c0ddc146101a35780638210a458146101c357600080fd5b80634a0f9973146100ed5780634ff56192146101215780635f3d926014610143575b600080fd5b3480156100f957600080fd5b50610104620249f081565b6040516001600160401b0390911681526020015b60405180910390f35b34801561012d57600080fd5b5061014161013c366004611222565b61033c565b005b610141610151366004611259565b610366565b34801561016257600080fd5b50606654610176906001600160a01b031681565b6040516001600160a01b039091168152602001610118565b34801561019a57600080fd5b50610141610378565b3480156101af57600080fd5b506101416101be366004611307565b61038c565b3480156101cf57600080fd5b506101046203d09081565b3480156101e657600080fd5b506101766101f536600461133c565b6067602052600090815260409020546001600160a01b031681565b34801561021c57600080fd5b5061025161022b36600461133c565b6001600160401b03166000908152606760205260409020546001600160a01b0316151590565b6040519015158152602001610118565b34801561026d57600080fd5b506033546001600160a01b0316610176565b34801561028b57600080fd5b5061029f61029a366004611259565b610491565b604051908152602001610118565b3480156102b957600080fd5b506102cd6102c83660046113f6565b6104a9565b60405161011891906114e0565b3480156102e657600080fd5b506101416102f5366004611222565b6104d1565b34801561030657600080fd5b5061030f61054a565b6040516101189190611538565b34801561032857600080fd5b506102516103373660046113f6565b6105d4565b6103446105f6565b606680546001600160a01b0319166001600160a01b0392909216919091179055565b61037284848484610650565b50505050565b6103806105f6565b61038a60006108b9565b565b6103946105f6565b6001600160401b0382166000908152606760205260409020546001600160a01b0316156104085760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a207265706c69636120616c726561647920736574000060448201526064015b60405180910390fd5b6001600160401b03918216600081815260676020526040812080546001600160a01b0319166001600160a01b0394909416939093179092556068805460018101825592527fa2153420d844928b4421650203c77babc8b33d7f2e7b450e2966db0c220977536004830401805460039093166008026101000a938402199092169202919091179055565b60006104a0858585338661090b565b95945050505050565b6040805180820190915260008152606060208201526104c9848484610a3d565b949350505050565b6104d96105f6565b6001600160a01b03811661053e5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103ff565b610547816108b9565b50565b606060688054806020026020016040519081016040528092919081815260200182805480156105ca57602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116105875790505b5050505050905090565b6000806105e2858585610a3d565b516001600160a01b03161415949350505050565b6033546001600160a01b0316331461038a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103ff565b6001600160401b0384166000908152606760205260409020546001600160a01b03166106be5760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a20636861696e206e6f7420737570706f72746564000060448201526064016103ff565b6106cb848484338561090b565b34101561071a5760405162461bcd60e51b815260206004820152601b60248201527f5852656769737472793a20696e73756666696369656e7420666565000000000060448201526064016103ff565b6107668484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250339250610761915086905061154b565b610b3f565b6107738484843385610bac565b6107b483838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250339250610cbe915050565b1561083c5760006107c860208301836115d0565b8101906107d5919061161d565b606654604051633ba5ccd560e01b81529192506001600160a01b031690633ba5ccd59061080890889085906004016116ce565b600060405180830381600087803b15801561082257600080fd5b505af1158015610836573d6000803e3d6000fd5b50505050505b604051339061084e90859085906116f0565b6040519081900390206001600160401b0386167f5f4d1d58125a7c776395dadc42b44dae862ffe3b0df3964431c8e942405bb43f61088f6020860186611222565b61089c60208701876115d0565b6040516108ab93929190611729565b60405180910390a450505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080610924878088888861091f8961154b565b610d14565b905060005b606854811015610a32576000606882815481106109485761094861174e565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b0316905060006109b7828a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610a3d915050565b80519091506001600160a01b03166109d0575050610a2a565b896001600160401b0316826001600160401b0316036109f0575050610a2a565b6109fe8a838b8b8b86610d14565b610a089085611764565b9350610a1b828b8b8b8b61091f8c61154b565b610a259085611764565b935050505b600101610929565b509695505050505050565b6040805180820190915260008152606060208201526001600160401b038416600090815260656020526040812090610a758585610f43565b81526020808201929092526040908101600020815180830190925280546001600160a01b031682526001810180549293919291840191610ab490611785565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae090611785565b8015610b2d5780601f10610b0257610100808354040283529160200191610b2d565b820191906000526020600020905b815481529060010190602001808311610b1057829003601f168201915b50505050508152505090509392505050565b6001600160401b03841660009081526065602052604081208291610b638686610f43565b8152602080820192909252604001600020825181546001600160a01b0319166001600160a01b03909116178155908201516001820190610ba3908261180c565b50505050505050565b610bc28580868686610bbd8761154b565b610f76565b60005b606854811015610cb657600060688281548110610be457610be461174e565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b031690506000610c538288888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250610a3d915050565b80519091506001600160a01b0316610c6c575050610cae565b876001600160401b0316826001600160401b031603610c8c575050610cae565b610c9a888389898986610f76565b610cab8289898989610bbd8a61154b565b50505b600101610bc5565b505050505050565b6000610d006040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b81525073121e240000000000000000000000000000000002610f43565b610d0a8484610f43565b1490505b92915050565b60665460408051634d4502c960e11b815290516000926001600160a01b0316918291639a8a0592916004808201926020929091908290030181865afa158015610d61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8591906118cb565b6001600160401b0316886001600160401b031603610da7576000915050610f39565b6001600160401b0388166000908152606760205260409020546001600160a01b031680610e115760405162461bcd60e51b81526020600482015260186024820152772c2932b3b4b9ba393c9d103ab735b737bbb71031b430b4b760411b60448201526064016103ff565b600063e4f1c67760e01b8989898989604051602401610e349594939291906118e8565b604051602081830303815290604052906001600160e01b0319166020820180516001600160e01b03838183161783525050505090506000610eac89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b9250610cbe915050565b610eb957620249f0610ebe565b6203d0905b604051632376548f60e21b81529091506001600160a01b03851690638dd9523c90610ef1908e908690869060040161192c565b602060405180830381865afa158015610f0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f329190611962565b9450505050505b9695505050505050565b60008282604051602001610f5892919061197b565b60405160208183030381529060405280519060200120905092915050565b60665460408051634d4502c960e11b815290516001600160a01b03909216918291639a8a05929160048083019260209291908290030181865afa158015610fc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe591906118cb565b6001600160401b0316876001600160401b0316036110035750610cb6565b6001600160401b0387166000908152606760205260409020546001600160a01b03168061106d5760405162461bcd60e51b81526020600482015260186024820152772c2932b3b4b9ba393c9d103ab735b737bbb71031b430b4b760411b60448201526064016103ff565b600063e4f1c67760e01b88888888886040516024016110909594939291906118e8565b604051602081830303815290604052906001600160e01b0319166020820180516001600160e01b0383818316178352505050509050600061110888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250610cbe915050565b61111557620249f061111a565b6203d0905b90506000846001600160a01b0316638dd9523c8c85856040518463ffffffff1660e01b815260040161114e9392919061192c565b602060405180830381865afa15801561116b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118f9190611962565b9050846001600160a01b031663c21dda4f828d60048888886040518763ffffffff1660e01b81526004016111c79594939291906119b2565b6000604051808303818588803b1580156111e057600080fd5b505af11580156111f4573d6000803e3d6000fd5b50505050505050505050505050505050565b80356001600160a01b038116811461121d57600080fd5b919050565b60006020828403121561123457600080fd5b61123d82611206565b9392505050565b6001600160401b038116811461054757600080fd5b6000806000806060858703121561126f57600080fd5b843561127a81611244565b935060208501356001600160401b038082111561129657600080fd5b818701915087601f8301126112aa57600080fd5b8135818111156112b957600080fd5b8860208285010111156112cb57600080fd5b6020830195508094505060408701359150808211156112e957600080fd5b508501604081880312156112fc57600080fd5b939692955090935050565b6000806040838503121561131a57600080fd5b823561132581611244565b915061133360208401611206565b90509250929050565b60006020828403121561134e57600080fd5b813561123d81611244565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561139757611397611359565b604052919050565b60006001600160401b038311156113b8576113b8611359565b6113cb601f8401601f191660200161136f565b90508281528383830111156113df57600080fd5b828260208301376000602084830101529392505050565b60008060006060848603121561140b57600080fd5b833561141681611244565b925060208401356001600160401b0381111561143157600080fd5b8401601f8101861361144257600080fd5b6114518682356020840161139f565b92505061146060408501611206565b90509250925092565b60005b8381101561148457818101518382015260200161146c565b50506000910152565b600081518084526114a5816020860160208601611469565b601f01601f19169290920160200192915050565b60018060a01b03815116825260006020820151604060208501526104c9604085018261148d565b60208152600061123d60208301846114b9565b60008151808452602080850194506020840160005b8381101561152d5781516001600160401b031687529582019590820190600101611508565b509495945050505050565b60208152600061123d60208301846114f3565b60006040823603121561155d57600080fd5b604051604081016001600160401b03828210818311171561158057611580611359565b8160405261158d85611206565b835260208501359150808211156115a357600080fd5b50830136601f8201126115b557600080fd5b6115c43682356020840161139f565b60208301525092915050565b6000808335601e198436030181126115e757600080fd5b8301803591506001600160401b0382111561160157600080fd5b60200191503681900382131561161657600080fd5b9250929050565b6000602080838503121561163057600080fd5b82356001600160401b038082111561164757600080fd5b818501915085601f83011261165b57600080fd5b81358181111561166d5761166d611359565b8060051b915061167e84830161136f565b818152918301840191848101908884111561169857600080fd5b938501935b838510156116c257843592506116b283611244565b828252938501939085019061169d565b98975050505050505050565b6001600160401b03831681526040602082015260006104c960408301846114f3565b8183823760009101908152919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03841681526040602082018190526000906104a09083018486611700565b634e487b7160e01b600052603260045260246000fd5b80820180821115610d0e57634e487b7160e01b600052601160045260246000fd5b600181811c9082168061179957607f821691505b6020821081036117b957634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115611807576000816000526020600020601f850160051c810160208610156117e85750805b601f850160051c820191505b81811015610cb6578281556001016117f4565b505050565b81516001600160401b0381111561182557611825611359565b611839816118338454611785565b846117bf565b602080601f83116001811461186e57600084156118565750858301515b600019600386901b1c1916600185901b178555610cb6565b600085815260208120601f198616915b8281101561189d5788860151825594840194600190910190840161187e565b50858210156118bb5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000602082840312156118dd57600080fd5b815161123d81611244565b6001600160401b038616815260806020820152600061190b608083018688611700565b6001600160a01b038516604084015282810360608401526116c281856114b9565b60006001600160401b0380861683526060602084015261194f606084018661148d565b9150808416604084015250949350505050565b60006020828403121561197457600080fd5b5051919050565b6000835161198d818460208801611469565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b60006001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a060608401526119ec60a084018661148d565b9150808416608084015250969550505050505056fea264697066735822122022e6ed965ee2ea864be68e9afb021bfe59a5cbdf64e5a1b7dec392c770796bba64736f6c63430008180033",
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

// XSETPORTALGASLIMIT is a free data retrieval call binding the contract method 0x8210a458.
//
// Solidity: function XSET_PORTAL_GAS_LIMIT() view returns(uint64)
func (_XRegistry *XRegistryCaller) XSETPORTALGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "XSET_PORTAL_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XSETPORTALGASLIMIT is a free data retrieval call binding the contract method 0x8210a458.
//
// Solidity: function XSET_PORTAL_GAS_LIMIT() view returns(uint64)
func (_XRegistry *XRegistrySession) XSETPORTALGASLIMIT() (uint64, error) {
	return _XRegistry.Contract.XSETPORTALGASLIMIT(&_XRegistry.CallOpts)
}

// XSETPORTALGASLIMIT is a free data retrieval call binding the contract method 0x8210a458.
//
// Solidity: function XSET_PORTAL_GAS_LIMIT() view returns(uint64)
func (_XRegistry *XRegistryCallerSession) XSETPORTALGASLIMIT() (uint64, error) {
	return _XRegistry.Contract.XSETPORTALGASLIMIT(&_XRegistry.CallOpts)
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
// Solidity: function get(uint64 chainId, string name, address registrant) view returns((address,bytes))
func (_XRegistry *XRegistryCaller) Get(opts *bind.CallOpts, chainId uint64, name string, registrant common.Address) (XRegistryBaseDeployment, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "get", chainId, name, registrant)

	if err != nil {
		return *new(XRegistryBaseDeployment), err
	}

	out0 := *abi.ConvertType(out[0], new(XRegistryBaseDeployment)).(*XRegistryBaseDeployment)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns((address,bytes))
func (_XRegistry *XRegistrySession) Get(chainId uint64, name string, registrant common.Address) (XRegistryBaseDeployment, error) {
	return _XRegistry.Contract.Get(&_XRegistry.CallOpts, chainId, name, registrant)
}

// Get is a free data retrieval call binding the contract method 0xb3449b77.
//
// Solidity: function get(uint64 chainId, string name, address registrant) view returns((address,bytes))
func (_XRegistry *XRegistryCallerSession) Get(chainId uint64, name string, registrant common.Address) (XRegistryBaseDeployment, error) {
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

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_XRegistry *XRegistryCaller) Portal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "portal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_XRegistry *XRegistrySession) Portal() (common.Address, error) {
	return _XRegistry.Contract.Portal(&_XRegistry.CallOpts)
}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_XRegistry *XRegistryCallerSession) Portal() (common.Address, error) {
	return _XRegistry.Contract.Portal(&_XRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x8f0d79a7.
//
// Solidity: function registrationFee(uint64 chainId, string name, (address,bytes) dep) view returns(uint256)
func (_XRegistry *XRegistryCaller) RegistrationFee(opts *bind.CallOpts, chainId uint64, name string, dep XRegistryBaseDeployment) (*big.Int, error) {
	var out []interface{}
	err := _XRegistry.contract.Call(opts, &out, "registrationFee", chainId, name, dep)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationFee is a free data retrieval call binding the contract method 0x8f0d79a7.
//
// Solidity: function registrationFee(uint64 chainId, string name, (address,bytes) dep) view returns(uint256)
func (_XRegistry *XRegistrySession) RegistrationFee(chainId uint64, name string, dep XRegistryBaseDeployment) (*big.Int, error) {
	return _XRegistry.Contract.RegistrationFee(&_XRegistry.CallOpts, chainId, name, dep)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x8f0d79a7.
//
// Solidity: function registrationFee(uint64 chainId, string name, (address,bytes) dep) view returns(uint256)
func (_XRegistry *XRegistryCallerSession) RegistrationFee(chainId uint64, name string, dep XRegistryBaseDeployment) (*big.Int, error) {
	return _XRegistry.Contract.RegistrationFee(&_XRegistry.CallOpts, chainId, name, dep)
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

// Register is a paid mutator transaction binding the contract method 0x5f3d9260.
//
// Solidity: function register(uint64 chainId, string name, (address,bytes) deployment) payable returns()
func (_XRegistry *XRegistryTransactor) Register(opts *bind.TransactOpts, chainId uint64, name string, deployment XRegistryBaseDeployment) (*types.Transaction, error) {
	return _XRegistry.contract.Transact(opts, "register", chainId, name, deployment)
}

// Register is a paid mutator transaction binding the contract method 0x5f3d9260.
//
// Solidity: function register(uint64 chainId, string name, (address,bytes) deployment) payable returns()
func (_XRegistry *XRegistrySession) Register(chainId uint64, name string, deployment XRegistryBaseDeployment) (*types.Transaction, error) {
	return _XRegistry.Contract.Register(&_XRegistry.TransactOpts, chainId, name, deployment)
}

// Register is a paid mutator transaction binding the contract method 0x5f3d9260.
//
// Solidity: function register(uint64 chainId, string name, (address,bytes) deployment) payable returns()
func (_XRegistry *XRegistryTransactorSession) Register(chainId uint64, name string, deployment XRegistryBaseDeployment) (*types.Transaction, error) {
	return _XRegistry.Contract.Register(&_XRegistry.TransactOpts, chainId, name, deployment)
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
// Solidity: function setPortal(address _portal) returns()
func (_XRegistry *XRegistryTransactor) SetPortal(opts *bind.TransactOpts, _portal common.Address) (*types.Transaction, error) {
	return _XRegistry.contract.Transact(opts, "setPortal", _portal)
}

// SetPortal is a paid mutator transaction binding the contract method 0x4ff56192.
//
// Solidity: function setPortal(address _portal) returns()
func (_XRegistry *XRegistrySession) SetPortal(_portal common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.SetPortal(&_XRegistry.TransactOpts, _portal)
}

// SetPortal is a paid mutator transaction binding the contract method 0x4ff56192.
//
// Solidity: function setPortal(address _portal) returns()
func (_XRegistry *XRegistryTransactorSession) SetPortal(_portal common.Address) (*types.Transaction, error) {
	return _XRegistry.Contract.SetPortal(&_XRegistry.TransactOpts, _portal)
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
	Metadata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterContractRegistered is a free log retrieval operation binding the contract event 0x5f4d1d58125a7c776395dadc42b44dae862ffe3b0df3964431c8e942405bb43f.
//
// Solidity: event ContractRegistered(uint64 indexed chainId, string indexed name, address indexed registrant, address addr, bytes metadata)
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

// WatchContractRegistered is a free log subscription operation binding the contract event 0x5f4d1d58125a7c776395dadc42b44dae862ffe3b0df3964431c8e942405bb43f.
//
// Solidity: event ContractRegistered(uint64 indexed chainId, string indexed name, address indexed registrant, address addr, bytes metadata)
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

// ParseContractRegistered is a log parse operation binding the contract event 0x5f4d1d58125a7c776395dadc42b44dae862ffe3b0df3964431c8e942405bb43f.
//
// Solidity: event ContractRegistered(uint64 indexed chainId, string indexed name, address indexed registrant, address addr, bytes metadata)
func (_XRegistry *XRegistryFilterer) ParseContractRegistered(log types.Log) (*XRegistryContractRegistered, error) {
	event := new(XRegistryContractRegistered)
	if err := _XRegistry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the XRegistry contract.
type XRegistryInitializedIterator struct {
	Event *XRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *XRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XRegistryInitialized)
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
		it.Event = new(XRegistryInitialized)
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
func (it *XRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XRegistryInitialized represents a Initialized event raised by the XRegistry contract.
type XRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_XRegistry *XRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*XRegistryInitializedIterator, error) {

	logs, sub, err := _XRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &XRegistryInitializedIterator{contract: _XRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_XRegistry *XRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *XRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _XRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XRegistryInitialized)
				if err := _XRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_XRegistry *XRegistryFilterer) ParseInitialized(log types.Log) (*XRegistryInitialized, error) {
	event := new(XRegistryInitialized)
	if err := _XRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
