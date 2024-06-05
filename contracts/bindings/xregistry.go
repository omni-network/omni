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
	ABI: "[{\"type\":\"function\",\"name\":\"XSET_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSET_PORTAL_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainIds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"has\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"portal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deployment\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"registrationFee\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dep\",\"type\":\"tuple\",\"internalType\":\"structXRegistryBase.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"replicas\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPortal\",\"inputs\":[{\"name\":\"_portal\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReplica\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"replica\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"registrant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611a358061007e6000396000f3fe6080604052600436106100e85760003560e01c806388f9380b1161008a578063b3449b7711610059578063b3449b77146102ad578063f2fde38b146102da578063fbe4b7c0146102fa578063fd0b64f71461031c57600080fd5b806388f9380b146101da5780638926f54f146102105780638da5cb5b146102615780638f0d79a71461027f57600080fd5b80636425666b116100c65780636425666b14610156578063715018a61461018e5780637b7c0ddc146101a35780638210a458146101c357600080fd5b80634a0f9973146100ed5780634ff56192146101215780635f3d926014610143575b600080fd5b3480156100f957600080fd5b50610104620249f081565b6040516001600160401b0390911681526020015b60405180910390f35b34801561012d57600080fd5b5061014161013c366004611220565b61033c565b005b610141610151366004611257565b610366565b34801561016257600080fd5b50600254610176906001600160a01b031681565b6040516001600160a01b039091168152602001610118565b34801561019a57600080fd5b50610141610378565b3480156101af57600080fd5b506101416101be366004611305565b61038c565b3480156101cf57600080fd5b506101046203d09081565b3480156101e657600080fd5b506101766101f536600461133a565b6003602052600090815260409020546001600160a01b031681565b34801561021c57600080fd5b5061025161022b36600461133a565b6001600160401b03166000908152600360205260409020546001600160a01b0316151590565b6040519015158152602001610118565b34801561026d57600080fd5b506000546001600160a01b0316610176565b34801561028b57600080fd5b5061029f61029a366004611257565b610491565b604051908152602001610118565b3480156102b957600080fd5b506102cd6102c83660046113f4565b6104a9565b60405161011891906114de565b3480156102e657600080fd5b506101416102f5366004611220565b6104d1565b34801561030657600080fd5b5061030f61054a565b6040516101189190611536565b34801561032857600080fd5b506102516103373660046113f4565b6105d4565b6103446105f6565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b61037284848484610650565b50505050565b6103806105f6565b61038a60006108b9565b565b6103946105f6565b6001600160401b0382166000908152600360205260409020546001600160a01b0316156104085760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a207265706c69636120616c726561647920736574000060448201526064015b60405180910390fd5b6001600160401b039182166000818152600360208190526040822080546001600160a01b0319166001600160a01b03959095169490941790935560048054600181018255918190527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b908204018054919093166008026101000a93840219169202919091179055565b60006104a08585853386610909565b95945050505050565b6040805180820190915260008152606060208201526104c9848484610a3b565b949350505050565b6104d96105f6565b6001600160a01b03811661053e5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103ff565b610547816108b9565b50565b606060048054806020026020016040519081016040528092919081815260200182805480156105ca57602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116105875790505b5050505050905090565b6000806105e2858585610a3b565b516001600160a01b03161415949350505050565b6000546001600160a01b0316331461038a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103ff565b6001600160401b0384166000908152600360205260409020546001600160a01b03166106be5760405162461bcd60e51b815260206004820152601e60248201527f5852656769737472793a20636861696e206e6f7420737570706f72746564000060448201526064016103ff565b6106cb8484843385610909565b34101561071a5760405162461bcd60e51b815260206004820152601b60248201527f5852656769737472793a20696e73756666696369656e7420666565000000000060448201526064016103ff565b6107668484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152503392506107619150869050611549565b610b3d565b6107738484843385610baa565b6107b483838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250339250610cbc915050565b1561083c5760006107c860208301836115ce565b8101906107d5919061161b565b600254604051633ba5ccd560e01b81529192506001600160a01b031690633ba5ccd59061080890889085906004016116cc565b600060405180830381600087803b15801561082257600080fd5b505af1158015610836573d6000803e3d6000fd5b50505050505b604051339061084e90859085906116ee565b6040519081900390206001600160401b0386167f5f4d1d58125a7c776395dadc42b44dae862ffe3b0df3964431c8e942405bb43f61088f6020860186611220565b61089c60208701876115ce565b6040516108ab93929190611727565b60405180910390a450505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600080610922878088888861091d89611549565b610d12565b905060005b600454811015610a30576000600482815481106109465761094661174c565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b0316905060006109b5828a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c9250610a3b915050565b80519091506001600160a01b03166109ce575050610a28565b896001600160401b0316826001600160401b0316036109ee575050610a28565b6109fc8a838b8b8b86610d12565b610a069085611762565b9350610a19828b8b8b8b61091d8c611549565b610a239085611762565b935050505b600101610927565b509695505050505050565b6040805180820190915260008152606060208201526001600160401b038416600090815260016020526040812090610a738585610f41565b81526020808201929092526040908101600020815180830190925280546001600160a01b031682526001810180549293919291840191610ab290611783565b80601f0160208091040260200160405190810160405280929190818152602001828054610ade90611783565b8015610b2b5780601f10610b0057610100808354040283529160200191610b2b565b820191906000526020600020905b815481529060010190602001808311610b0e57829003601f168201915b50505050508152505090509392505050565b6001600160401b03841660009081526001602052604081208291610b618686610f41565b8152602080820192909252604001600020825181546001600160a01b0319166001600160a01b03909116178155908201516001820190610ba1908261180a565b50505050505050565b610bc08580868686610bbb87611549565b610f74565b60005b600454811015610cb457600060048281548110610be257610be261174c565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b031690506000610c518288888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250610a3b915050565b80519091506001600160a01b0316610c6a575050610cac565b876001600160401b0316826001600160401b031603610c8a575050610cac565b610c98888389898986610f74565b610ca98289898989610bbb8a611549565b50505b600101610bc3565b505050505050565b6000610cfe6040518060400160405280600a81526020016913db5b9a541bdc9d185b60b21b81525073121e240000000000000000000000000000000002610f41565b610d088484610f41565b1490505b92915050565b60025460408051634d4502c960e11b815290516000926001600160a01b0316918291639a8a0592916004808201926020929091908290030181865afa158015610d5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8391906118c9565b6001600160401b0316886001600160401b031603610da5576000915050610f37565b6001600160401b0388166000908152600360205260409020546001600160a01b031680610e0f5760405162461bcd60e51b81526020600482015260186024820152772c2932b3b4b9ba393c9d103ab735b737bbb71031b430b4b760411b60448201526064016103ff565b600063e4f1c67760e01b8989898989604051602401610e329594939291906118e6565b604051602081830303815290604052906001600160e01b0319166020820180516001600160e01b03838183161783525050505090506000610eaa89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b9250610cbc915050565b610eb757620249f0610ebc565b6203d0905b604051632376548f60e21b81529091506001600160a01b03851690638dd9523c90610eef908e908690869060040161192a565b602060405180830381865afa158015610f0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f309190611960565b9450505050505b9695505050505050565b60008282604051602001610f56929190611979565b60405160208183030381529060405280519060200120905092915050565b60025460408051634d4502c960e11b815290516001600160a01b03909216918291639a8a05929160048083019260209291908290030181865afa158015610fbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe391906118c9565b6001600160401b0316876001600160401b0316036110015750610cb4565b6001600160401b0387166000908152600360205260409020546001600160a01b03168061106b5760405162461bcd60e51b81526020600482015260186024820152772c2932b3b4b9ba393c9d103ab735b737bbb71031b430b4b760411b60448201526064016103ff565b600063e4f1c67760e01b888888888860405160240161108e9594939291906118e6565b604051602081830303815290604052906001600160e01b0319166020820180516001600160e01b0383818316178352505050509050600061110688888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250610cbc915050565b61111357620249f0611118565b6203d0905b90506000846001600160a01b0316638dd9523c8c85856040518463ffffffff1660e01b815260040161114c9392919061192a565b602060405180830381865afa158015611169573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118d9190611960565b9050846001600160a01b031663c21dda4f828d60048888886040518763ffffffff1660e01b81526004016111c59594939291906119b0565b6000604051808303818588803b1580156111de57600080fd5b505af11580156111f2573d6000803e3d6000fd5b50505050505050505050505050505050565b80356001600160a01b038116811461121b57600080fd5b919050565b60006020828403121561123257600080fd5b61123b82611204565b9392505050565b6001600160401b038116811461054757600080fd5b6000806000806060858703121561126d57600080fd5b843561127881611242565b935060208501356001600160401b038082111561129457600080fd5b818701915087601f8301126112a857600080fd5b8135818111156112b757600080fd5b8860208285010111156112c957600080fd5b6020830195508094505060408701359150808211156112e757600080fd5b508501604081880312156112fa57600080fd5b939692955090935050565b6000806040838503121561131857600080fd5b823561132381611242565b915061133160208401611204565b90509250929050565b60006020828403121561134c57600080fd5b813561123b81611242565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561139557611395611357565b604052919050565b60006001600160401b038311156113b6576113b6611357565b6113c9601f8401601f191660200161136d565b90508281528383830111156113dd57600080fd5b828260208301376000602084830101529392505050565b60008060006060848603121561140957600080fd5b833561141481611242565b925060208401356001600160401b0381111561142f57600080fd5b8401601f8101861361144057600080fd5b61144f8682356020840161139d565b92505061145e60408501611204565b90509250925092565b60005b8381101561148257818101518382015260200161146a565b50506000910152565b600081518084526114a3816020860160208601611467565b601f01601f19169290920160200192915050565b60018060a01b03815116825260006020820151604060208501526104c9604085018261148b565b60208152600061123b60208301846114b7565b60008151808452602080850194506020840160005b8381101561152b5781516001600160401b031687529582019590820190600101611506565b509495945050505050565b60208152600061123b60208301846114f1565b60006040823603121561155b57600080fd5b604051604081016001600160401b03828210818311171561157e5761157e611357565b8160405261158b85611204565b835260208501359150808211156115a157600080fd5b50830136601f8201126115b357600080fd5b6115c23682356020840161139d565b60208301525092915050565b6000808335601e198436030181126115e557600080fd5b8301803591506001600160401b038211156115ff57600080fd5b60200191503681900382131561161457600080fd5b9250929050565b6000602080838503121561162e57600080fd5b82356001600160401b038082111561164557600080fd5b818501915085601f83011261165957600080fd5b81358181111561166b5761166b611357565b8060051b915061167c84830161136d565b818152918301840191848101908884111561169657600080fd5b938501935b838510156116c057843592506116b083611242565b828252938501939085019061169b565b98975050505050505050565b6001600160401b03831681526040602082015260006104c960408301846114f1565b8183823760009101908152919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03841681526040602082018190526000906104a090830184866116fe565b634e487b7160e01b600052603260045260246000fd5b80820180821115610d0c57634e487b7160e01b600052601160045260246000fd5b600181811c9082168061179757607f821691505b6020821081036117b757634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115611805576000816000526020600020601f850160051c810160208610156117e65750805b601f850160051c820191505b81811015610cb4578281556001016117f2565b505050565b81516001600160401b0381111561182357611823611357565b611837816118318454611783565b846117bd565b602080601f83116001811461186c57600084156118545750858301515b600019600386901b1c1916600185901b178555610cb4565b600085815260208120601f198616915b8281101561189b5788860151825594840194600190910190840161187c565b50858210156118b95787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000602082840312156118db57600080fd5b815161123b81611242565b6001600160401b03861681526080602082015260006119096080830186886116fe565b6001600160a01b038516604084015282810360608401526116c081856114b7565b60006001600160401b0380861683526060602084015261194d606084018661148b565b9150808416604084015250949350505050565b60006020828403121561197257600080fd5b5051919050565b6000835161198b818460208801611467565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b60006001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a060608401526119ea60a084018661148b565b9150808416608084015250969550505050505056fea2646970667358221220773b19fcf13e12759e92cd185935480b502d1ae53a315114cf66efe0446b13c264736f6c63430008180033",
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
