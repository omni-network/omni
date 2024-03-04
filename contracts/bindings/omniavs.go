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

// IOmniAVSOperator is an auto generated low-level Go binding around an user-defined struct.
type IOmniAVSOperator struct {
	Addr      common.Address
	Delegated *big.Int
	Staked    *big.Int
}

// IOmniAVSStrategyParam is an auto generated low-level Go binding around an user-defined struct.
type IOmniAVSStrategyParam struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// OmniAVSMetaData contains all meta data concerning the OmniAVS contract.
var OmniAVSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"delegationManager_\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"avsDirectory_\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ethStakeInbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeForSync\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"},{\"name\":\"omniChainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ethStakeInbox_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategyParams_\",\"type\":\"tuple[]\",\"internalType\":\"structIOmniAVS.StrategyParam[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isInAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOmniAVS.Operator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegated\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"staked\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFromAllowlist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEthStakeInbox\",\"inputs\":[{\"name\":\"inbox\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOmniChainId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOmniPortal\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIOmniAVS.StrategyParam[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXcallGasLimits\",\"inputs\":[{\"name\":\"base\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"perOperator\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyParams\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOmniAVS.StrategyParam[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"syncWithOmni\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"xcallBaseGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xcallGasLimitPerOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAdded\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAllowed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDisallowed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemoved\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200274938038062002749833981016040819052620000349162000133565b6001600160a01b03808316608052811660a0526200005162000059565b505062000172565b600054610100900460ff1615620000c65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000118576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013057600080fd5b50565b600080604083850312156200014757600080fd5b825162000154816200011a565b602084015190925062000167816200011a565b809150509250929050565b60805160a051612595620001b4600039600081816103d1015281816108b901528181610aa30152610c360152600081816119700152611a2f01526125956000f3fe6080604052600436106101cd5760003560e01c8063750521f5116100f7578063b98912b011610095578063e673df8a11610064578063e673df8a14610555578063f2fde38b14610577578063f57f33d014610597578063f8e86ece146105b957600080fd5b8063b98912b0146104dd578063c4d179cf14610500578063d17efb3614610520578063e481af9d1461054057600080fd5b80638da5cb5b116100d15780638da5cb5b1461045f5780639926ee7d1461047d578063a364f4da1461049d578063ae30f16d146104bd57600080fd5b8063750521f51461040a5780637815873d1461042a5780638456cb591461044a57600080fd5b80633f4ba83a1161016f5780635c975abb1161013e5780635c975abb1461038a5780635da93d7e146103a25780636b3aa72e146103c2578063715018a6146103f557600080fd5b80633f4ba83a1461030e5780634a8fa4591461032357806354c74ed3146103435780635c78b0e21461036a57600080fd5b8063243d51c7116101ab578063243d51c71461023957806329d0fdc01461026057806333cfb7b7146102a957806339acf9f1146102d657600080fd5b80630c415884146101d2578063110ff5f1146101f457806313efbe9214610231575b600080fd5b3480156101de57600080fd5b506101f26101ed366004611cbd565b6105d9565b005b34801561020057600080fd5b50609a54610214906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b6101f2610603565b34801561024557600080fd5b50609a5461021490600160801b90046001600160401b031681565b34801561026c57600080fd5b5061029961027b366004611cbd565b6001600160a01b031660009081526099602052604090205460ff1690565b6040519015158152602001610228565b3480156102b557600080fd5b506102c96102c4366004611cbd565b6106ec565b6040516102289190611cda565b3480156102e257600080fd5b50609c546102f6906001600160a01b031681565b6040516001600160a01b039091168152602001610228565b34801561031a57600080fd5b506101f261071d565b34801561032f57600080fd5b506101f261033e366004611d43565b61072f565b34801561034f57600080fd5b50609a5461021490600160401b90046001600160401b031681565b34801561037657600080fd5b506101f2610385366004611d76565b610790565b34801561039657600080fd5b5060655460ff16610299565b3480156103ae57600080fd5b506101f26103bd366004611cbd565b6107bb565b3480156103ce57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102f6565b34801561040157600080fd5b506101f2610888565b34801561041657600080fd5b506101f2610425366004611e56565b61089a565b34801561043657600080fd5b50609b546102f6906001600160a01b031681565b34801561045657600080fd5b506101f2610923565b34801561046b57600080fd5b506033546001600160a01b03166102f6565b34801561048957600080fd5b506101f2610498366004611ea6565b610933565b3480156104a957600080fd5b506101f26104b8366004611cbd565b610b44565b3480156104c957600080fd5b506101f26104d8366004611f9b565b610cc9565b3480156104e957600080fd5b506104f2610cdf565b604051908152602001610228565b34801561050c57600080fd5b506101f261051b366004611fdc565b610dcb565b34801561052c57600080fd5b506101f261053b366004611cbd565b610f66565b34801561054c57600080fd5b506102c9610fde565b34801561056157600080fd5b5061056a610fed565b6040516102289190612067565b34801561058357600080fd5b506101f2610592366004611cbd565b610ff7565b3480156105a357600080fd5b506105ac611070565b60405161022891906120d6565b3480156105c557600080fd5b506101f26105d4366004611cbd565b6110ec565b6105e16111ff565b609c80546001600160a01b0319166001600160a01b0392909216919091179055565b61060b611259565b600061061561129f565b609c54609a54609b546040519394506001600160a01b03928316936370e8b56a9334936001600160401b03169216906333364ffb60e11b9061065b908890602401612067565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091528651610698906113f4565b6040518663ffffffff1660e01b81526004016106b79493929190612177565b6000604051808303818588803b1580156106d057600080fd5b505af11580156106e4573d6000803e3d6000fd5b505050505050565b60606106f782611429565b61070f57505060408051600081526020810190915290565b610717611489565b92915050565b6107256111ff565b61072d611541565b565b6107376111ff565b609a805477ffffffffffffffffffffffffffffffff00000000000000001916600160801b6001600160401b03948516026fffffffffffffffff0000000000000000191617600160401b9290931691909102919091179055565b6107986111ff565b609a805467ffffffffffffffff19166001600160401b0392909216919091179055565b6107c36111ff565b6001600160a01b03811660009081526099602052604090205460ff166108305760405162461bcd60e51b815260206004820152601960248201527f4f6d6e694156533a206e6f7420696e20616c6c6f776c6973740000000000000060448201526064015b60405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916905590519182527f8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb91015b60405180910390a150565b6108906111ff565b61072d6000611593565b6108a26111ff565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906108ee9084906004016121bc565b600060405180830381600087803b15801561090857600080fd5b505af115801561091c573d6000803e3d6000fd5b5050505050565b61092b6111ff565b61072d6115e5565b61093b611259565b336001600160a01b0383161461098c5760405162461bcd60e51b815260206004820152601660248201527527b6b734a0ab299d1037b7363c9037b832b930ba37b960511b6044820152606401610827565b6001600160a01b03821660009081526099602052604090205460ff166109eb5760405162461bcd60e51b815260206004820152601460248201527313db5b9a505594ce881b9bdd08185b1b1bddd95960621b6044820152606401610827565b6109f482611429565b15610a415760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e694156533a20616c726561647920616e206f70657261746f72000000006044820152606401610827565b609880546001810182556000919091527f2237a976fa961f5921fd19f2b03c925c725d77b20ce8f790c19709c03de4d8140180546001600160a01b0319166001600160a01b038416179055604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610ada90859085906004016121cf565b600060405180830381600087803b158015610af457600080fd5b505af1158015610b08573d6000803e3d6000fd5b50506040516001600160a01b03851692507fac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d9150600090a25050565b610b4c611259565b336001600160a01b0382161480610b6d57506033546001600160a01b031633145b610bb95760405162461bcd60e51b815260206004820152601f60248201527f4f6d6e694156533a206f6e6c79206f70657261746f72206f72206f776e6572006044820152606401610827565b610bc281611429565b610c0e5760405162461bcd60e51b815260206004820152601860248201527f4f6d6e694156533a206e6f7420616e206f70657261746f7200000000000000006044820152606401610827565b610c1781611622565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401600060405180830381600087803b158015610c7a57600080fd5b505af1158015610c8e573d6000803e3d6000fd5b50506040516001600160a01b03841692507f80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d9150600090a250565b610cd16111ff565b610cdb8282611720565b5050565b600080610cea61129f565b609c54609a546040519293506001600160a01b0390911691638dd9523c916001600160401b0316906333364ffb60e11b90610d29908690602401612067565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091528451610d66906113f4565b6040518463ffffffff1660e01b8152600401610d849392919061221a565b602060405180830381865afa158015610da1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc59190612250565b91505090565b600054610100900460ff1615808015610deb5750600054600160ff909116105b80610e055750303b158015610e05575060005460ff166001145b610e685760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610827565b6000805460ff191660011790558015610e8b576000805461ff0019166101001790555b609c80546001600160a01b03199081166001600160a01b0389811691909117909255609a80546001600160401b0389166fffffffffffffffffffffffffffffffff199091161769c35000000000000000001767ffffffffffffffff60801b191661249f60831b179055609b8054909116918616919091179055610f0d87611593565b610f178383611720565b8015610f5d576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b610f6e6111ff565b6001600160a01b038116610fbc5760405162461bcd60e51b81526020600482015260156024820152744f6d6e694156533a207a65726f206164647265737360581b6044820152606401610827565b609b80546001600160a01b0319166001600160a01b0392909216919091179055565b6060610fe8611489565b905090565b6060610fe861129f565b610fff6111ff565b6001600160a01b0381166110645760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610827565b61106d81611593565b50565b60606097805480602002602001604051908101604052809291908181526020016000905b828210156110e357600084815260209081902060408051808201909152908401546001600160a01b0381168252600160a01b90046001600160601b031681830152825260019092019101611094565b50505050905090565b6110f46111ff565b6001600160a01b0381166111425760405162461bcd60e51b81526020600482015260156024820152744f6d6e694156533a207a65726f206164647265737360581b6044820152606401610827565b6001600160a01b03811660009081526099602052604090205460ff16156111ab5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e694156533a20616c726561647920696e20616c6c6f776c6973740000006044820152606401610827565b6001600160a01b038116600081815260996020908152604091829020805460ff1916600117905590519182527fdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75910161087d565b6033546001600160a01b0316331461072d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610827565b60655460ff161561072d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610827565b6098546060906000906001600160401b038111156112bf576112bf611d91565b60405190808252806020026020018201604052801561130a57816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816112dd5790505b50905060005b81518110156113ee5760006098828154811061132e5761132e612269565b60009182526020822001546001600160a01b0316915061134d826118d8565b9050600061135a83611a09565b90506000816001600160601b0316836001600160601b03161161137e576000611388565b6113888284612295565b90506040518060600160405280856001600160a01b03168152602001826001600160601b03168152602001836001600160601b03168152508686815181106113d2576113d2612269565b6020026020010181905250848060010195505050505050611310565b50919050565b609a546000906001600160401b03600160801b820481169161141f91600160401b90910416846122bd565b61071791906122ec565b6000805b60985481101561148057826001600160a01b03166098828154811061145457611454612269565b6000918252602090912001546001600160a01b031614156114785750600192915050565b60010161142d565b50600092915050565b6097546060906000906001600160401b038111156114a9576114a9611d91565b6040519080825280602002602001820160405280156114d2578160200160208202803683370190505b50905060005b6097548110156113ee57609781815481106114f5576114f5612269565b60009182526020909120015482516001600160a01b039091169083908390811061152157611521612269565b6001600160a01b03909216602092830291909101909101526001016114d8565b611549611bfe565b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6115ed611259565b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586115763390565b60005b609854811015610cdb57816001600160a01b03166098828154811061164c5761164c612269565b6000918252602090912001546001600160a01b03161415611718576098805461167790600190612317565b8154811061168757611687612269565b600091825260209091200154609880546001600160a01b0390921691839081106116b3576116b3612269565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060988054806116f2576116f261232e565b600082815260209020810160001990810180546001600160a01b03191690550190555050565b600101611625565b61172c60976000611c76565b60005b818110156118d357600083838381811061174b5761174b612269565b6117619260206040909202019081019150611cbd565b6001600160a01b031614156117b85760405162461bcd60e51b815260206004820152601960248201527f4f6d6e694156533a206e6f207a65726f207374726174656779000000000000006044820152606401610827565b60006117c5826001612344565b90505b82811015611888578383828181106117e2576117e2612269565b6117f89260206040909202019081019150611cbd565b6001600160a01b031684848481811061181357611813612269565b6118299260206040909202019081019150611cbd565b6001600160a01b031614156118805760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694156533a206e6f206475706c696361746520737472617465677900006044820152606401610827565b6001016117c8565b50609783838381811061189d5761189d612269565b835460018101855560009485526020909420604090910292909201929190910190506118c9828261235c565b505060010161172f565b505050565b6040805180820190915260008082526020820181905290819060005b609754811015611a00576097818154811061191157611911612269565b6000918252602080832060408051808201825293909101546001600160a01b03808216808652600160a01b9092046001600160601b031693850193909352905163778e55f360e01b8152898316600482015260248101919091529194507f0000000000000000000000000000000000000000000000000000000000000000169063778e55f390604401602060405180830381865afa1580156119b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119db9190612250565b90506119eb818460200151611c47565b6119f590856123aa565b9350506001016118f4565b50909392505050565b6040516367c0439f60e11b81526001600160a01b038281166004830152600091829182917f00000000000000000000000000000000000000000000000000000000000000009091169063cf80873e90602401600060405180830381865afa158015611a78573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611aa0919081019061245a565b90925090506000805b8351811015611bf5576000848281518110611ac657611ac6612269565b60200260200101519050611aff604051806040016040528060006001600160a01b0316815260200160006001600160601b031681525090565b60005b609754811015611ba157826001600160a01b031660978281548110611b2957611b29612269565b6000918252602090912001546001600160a01b03161415611b995760978181548110611b5757611b57612269565b6000918252602091829020604080518082019091529101546001600160a01b0381168252600160a01b90046001600160601b0316918101919091529150611ba1565b600101611b02565b5080516001600160a01b0316611bb8575050611aa9565b611bdf858481518110611bcd57611bcd612269565b60200260200101518260200151611c47565b611be990856123aa565b93505050600101611aa9565b50949350505050565b60655460ff1661072d5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610827565b6000670de0b6b3a7640000611c656001600160601b0384168561251e565b611c6f919061253d565b9392505050565b508054600082559060005260206000209081019061106d91905b80821115611ca45760008155600101611c90565b5090565b6001600160a01b038116811461106d57600080fd5b600060208284031215611ccf57600080fd5b8135611c6f81611ca8565b6020808252825182820181905260009190848201906040850190845b81811015611d1b5783516001600160a01b031683529284019291840191600101611cf6565b50909695505050505050565b80356001600160401b0381168114611d3e57600080fd5b919050565b60008060408385031215611d5657600080fd5b611d5f83611d27565b9150611d6d60208401611d27565b90509250929050565b600060208284031215611d8857600080fd5b611c6f82611d27565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715611dc957611dc9611d91565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611df757611df7611d91565b604052919050565b60006001600160401b03831115611e1857611e18611d91565b611e2b601f8401601f1916602001611dcf565b9050828152838383011115611e3f57600080fd5b828260208301376000602084830101529392505050565b600060208284031215611e6857600080fd5b81356001600160401b03811115611e7e57600080fd5b8201601f81018413611e8f57600080fd5b611e9e84823560208401611dff565b949350505050565b60008060408385031215611eb957600080fd5b8235611ec481611ca8565b915060208301356001600160401b0380821115611ee057600080fd5b9084019060608287031215611ef457600080fd5b611efc611da7565b823582811115611f0b57600080fd5b83019150601f82018713611f1e57600080fd5b611f2d87833560208501611dff565b815260208301356020820152604083013560408201528093505050509250929050565b60008083601f840112611f6257600080fd5b5081356001600160401b03811115611f7957600080fd5b6020830191508360208260061b8501011115611f9457600080fd5b9250929050565b60008060208385031215611fae57600080fd5b82356001600160401b03811115611fc457600080fd5b611fd085828601611f50565b90969095509350505050565b60008060008060008060a08789031215611ff557600080fd5b863561200081611ca8565b9550602087013561201081611ca8565b945061201e60408801611d27565b9350606087013561202e81611ca8565b925060808701356001600160401b0381111561204957600080fd5b61205589828a01611f50565b979a9699509497509295939492505050565b602080825282518282018190526000919060409081850190868401855b828110156120c957815180516001600160a01b03168552868101516001600160601b039081168887015290860151168585015260609093019290850190600101612084565b5091979650505050505050565b602080825282518282018190526000919060409081850190868401855b828110156120c957815180516001600160a01b031685528601516001600160601b03168685015292840192908501906001016120f3565b6000815180845260005b8181101561215057602081850181015186830182015201612134565b81811115612162576000602083870101525b50601f01601f19169290920160200192915050565b60006001600160401b03808716835260018060a01b0386166020840152608060408401526121a8608084018661212a565b915080841660608401525095945050505050565b602081526000611c6f602083018461212a565b60018060a01b03831681526040602082015260008251606060408401526121f960a084018261212a565b90506020840151606084015260408401516080840152809150509392505050565b60006001600160401b0380861683526060602084015261223d606084018661212a565b9150808416604084015250949350505050565b60006020828403121561226257600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001600160601b03838116908316818110156122b5576122b561227f565b039392505050565b60006001600160401b03808316818516818304811182151516156122e3576122e361227f565b02949350505050565b60006001600160401b0380831681851680830382111561230e5761230e61227f565b01949350505050565b6000828210156123295761232961227f565b500390565b634e487b7160e01b600052603160045260246000fd5b600082198211156123575761235761227f565b500190565b813561236781611ca8565b81546001600160a01b03199081166001600160a01b0392909216918217835560208401356001600160601b03811681146123a057600080fd5b60a01b1617905550565b60006001600160601b0380831681851680830382111561230e5761230e61227f565b60006001600160401b038211156123e5576123e5611d91565b5060051b60200190565b600082601f83011261240057600080fd5b81516020612415612410836123cc565b611dcf565b82815260059290921b8401810191818101908684111561243457600080fd5b8286015b8481101561244f5780518352918301918301612438565b509695505050505050565b6000806040838503121561246d57600080fd5b82516001600160401b038082111561248457600080fd5b818501915085601f83011261249857600080fd5b815160206124a8612410836123cc565b82815260059290921b840181019181810190898411156124c757600080fd5b948201945b838610156124ee5785516124df81611ca8565b825294820194908201906124cc565b9188015191965090935050508082111561250757600080fd5b50612514858286016123ef565b9150509250929050565b60008160001904831182151516156125385761253861227f565b500290565b60008261255a57634e487b7160e01b600052601260045260246000fd5b50049056fea2646970667358221220fd4618ae6a6976b4a693de7d01bfef7433733f7c4dee4a642fc73ba43527360e64736f6c634300080c0033",
}

// OmniAVSABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniAVSMetaData.ABI instead.
var OmniAVSABI = OmniAVSMetaData.ABI

// OmniAVSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniAVSMetaData.Bin instead.
var OmniAVSBin = OmniAVSMetaData.Bin

// DeployOmniAVS deploys a new Ethereum contract, binding an instance of OmniAVS to it.
func DeployOmniAVS(auth *bind.TransactOpts, backend bind.ContractBackend, delegationManager_ common.Address, avsDirectory_ common.Address) (common.Address, *types.Transaction, *OmniAVS, error) {
	parsed, err := OmniAVSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniAVSBin), backend, delegationManager_, avsDirectory_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniAVS{OmniAVSCaller: OmniAVSCaller{contract: contract}, OmniAVSTransactor: OmniAVSTransactor{contract: contract}, OmniAVSFilterer: OmniAVSFilterer{contract: contract}}, nil
}

// OmniAVS is an auto generated Go binding around an Ethereum contract.
type OmniAVS struct {
	OmniAVSCaller     // Read-only binding to the contract
	OmniAVSTransactor // Write-only binding to the contract
	OmniAVSFilterer   // Log filterer for contract events
}

// OmniAVSCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniAVSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniAVSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniAVSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniAVSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniAVSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniAVSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniAVSSession struct {
	Contract     *OmniAVS          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniAVSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniAVSCallerSession struct {
	Contract *OmniAVSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OmniAVSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniAVSTransactorSession struct {
	Contract     *OmniAVSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OmniAVSRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniAVSRaw struct {
	Contract *OmniAVS // Generic contract binding to access the raw methods on
}

// OmniAVSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniAVSCallerRaw struct {
	Contract *OmniAVSCaller // Generic read-only contract binding to access the raw methods on
}

// OmniAVSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniAVSTransactorRaw struct {
	Contract *OmniAVSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniAVS creates a new instance of OmniAVS, bound to a specific deployed contract.
func NewOmniAVS(address common.Address, backend bind.ContractBackend) (*OmniAVS, error) {
	contract, err := bindOmniAVS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniAVS{OmniAVSCaller: OmniAVSCaller{contract: contract}, OmniAVSTransactor: OmniAVSTransactor{contract: contract}, OmniAVSFilterer: OmniAVSFilterer{contract: contract}}, nil
}

// NewOmniAVSCaller creates a new read-only instance of OmniAVS, bound to a specific deployed contract.
func NewOmniAVSCaller(address common.Address, caller bind.ContractCaller) (*OmniAVSCaller, error) {
	contract, err := bindOmniAVS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniAVSCaller{contract: contract}, nil
}

// NewOmniAVSTransactor creates a new write-only instance of OmniAVS, bound to a specific deployed contract.
func NewOmniAVSTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniAVSTransactor, error) {
	contract, err := bindOmniAVS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniAVSTransactor{contract: contract}, nil
}

// NewOmniAVSFilterer creates a new log filterer instance of OmniAVS, bound to a specific deployed contract.
func NewOmniAVSFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniAVSFilterer, error) {
	contract, err := bindOmniAVS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniAVSFilterer{contract: contract}, nil
}

// bindOmniAVS binds a generic wrapper to an already deployed contract.
func bindOmniAVS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniAVSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniAVS *OmniAVSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniAVS.Contract.OmniAVSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniAVS *OmniAVSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniAVS.Contract.OmniAVSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniAVS *OmniAVSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniAVS.Contract.OmniAVSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniAVS *OmniAVSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniAVS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniAVS *OmniAVSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniAVS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniAVS *OmniAVSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniAVS.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_OmniAVS *OmniAVSCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_OmniAVS *OmniAVSSession) AvsDirectory() (common.Address, error) {
	return _OmniAVS.Contract.AvsDirectory(&_OmniAVS.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_OmniAVS *OmniAVSCallerSession) AvsDirectory() (common.Address, error) {
	return _OmniAVS.Contract.AvsDirectory(&_OmniAVS.CallOpts)
}

// EthStakeInbox is a free data retrieval call binding the contract method 0x7815873d.
//
// Solidity: function ethStakeInbox() view returns(address)
func (_OmniAVS *OmniAVSCaller) EthStakeInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "ethStakeInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthStakeInbox is a free data retrieval call binding the contract method 0x7815873d.
//
// Solidity: function ethStakeInbox() view returns(address)
func (_OmniAVS *OmniAVSSession) EthStakeInbox() (common.Address, error) {
	return _OmniAVS.Contract.EthStakeInbox(&_OmniAVS.CallOpts)
}

// EthStakeInbox is a free data retrieval call binding the contract method 0x7815873d.
//
// Solidity: function ethStakeInbox() view returns(address)
func (_OmniAVS *OmniAVSCallerSession) EthStakeInbox() (common.Address, error) {
	return _OmniAVS.Contract.EthStakeInbox(&_OmniAVS.CallOpts)
}

// FeeForSync is a free data retrieval call binding the contract method 0xb98912b0.
//
// Solidity: function feeForSync() view returns(uint256)
func (_OmniAVS *OmniAVSCaller) FeeForSync(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "feeForSync")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeForSync is a free data retrieval call binding the contract method 0xb98912b0.
//
// Solidity: function feeForSync() view returns(uint256)
func (_OmniAVS *OmniAVSSession) FeeForSync() (*big.Int, error) {
	return _OmniAVS.Contract.FeeForSync(&_OmniAVS.CallOpts)
}

// FeeForSync is a free data retrieval call binding the contract method 0xb98912b0.
//
// Solidity: function feeForSync() view returns(uint256)
func (_OmniAVS *OmniAVSCallerSession) FeeForSync() (*big.Int, error) {
	return _OmniAVS.Contract.FeeForSync(&_OmniAVS.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_OmniAVS *OmniAVSCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_OmniAVS *OmniAVSSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _OmniAVS.Contract.GetOperatorRestakedStrategies(&_OmniAVS.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_OmniAVS *OmniAVSCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _OmniAVS.Contract.GetOperatorRestakedStrategies(&_OmniAVS.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_OmniAVS *OmniAVSCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_OmniAVS *OmniAVSSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _OmniAVS.Contract.GetRestakeableStrategies(&_OmniAVS.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_OmniAVS *OmniAVSCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _OmniAVS.Contract.GetRestakeableStrategies(&_OmniAVS.CallOpts)
}

// IsInAllowlist is a free data retrieval call binding the contract method 0x29d0fdc0.
//
// Solidity: function isInAllowlist(address operator) view returns(bool)
func (_OmniAVS *OmniAVSCaller) IsInAllowlist(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "isInAllowlist", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInAllowlist is a free data retrieval call binding the contract method 0x29d0fdc0.
//
// Solidity: function isInAllowlist(address operator) view returns(bool)
func (_OmniAVS *OmniAVSSession) IsInAllowlist(operator common.Address) (bool, error) {
	return _OmniAVS.Contract.IsInAllowlist(&_OmniAVS.CallOpts, operator)
}

// IsInAllowlist is a free data retrieval call binding the contract method 0x29d0fdc0.
//
// Solidity: function isInAllowlist(address operator) view returns(bool)
func (_OmniAVS *OmniAVSCallerSession) IsInAllowlist(operator common.Address) (bool, error) {
	return _OmniAVS.Contract.IsInAllowlist(&_OmniAVS.CallOpts, operator)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniAVS *OmniAVSCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniAVS *OmniAVSSession) Omni() (common.Address, error) {
	return _OmniAVS.Contract.Omni(&_OmniAVS.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniAVS *OmniAVSCallerSession) Omni() (common.Address, error) {
	return _OmniAVS.Contract.Omni(&_OmniAVS.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniAVS *OmniAVSCaller) OmniChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "omniChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniAVS *OmniAVSSession) OmniChainId() (uint64, error) {
	return _OmniAVS.Contract.OmniChainId(&_OmniAVS.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniAVS *OmniAVSCallerSession) OmniChainId() (uint64, error) {
	return _OmniAVS.Contract.OmniChainId(&_OmniAVS.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns((address,uint96,uint96)[])
func (_OmniAVS *OmniAVSCaller) Operators(opts *bind.CallOpts) ([]IOmniAVSOperator, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "operators")

	if err != nil {
		return *new([]IOmniAVSOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOmniAVSOperator)).(*[]IOmniAVSOperator)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns((address,uint96,uint96)[])
func (_OmniAVS *OmniAVSSession) Operators() ([]IOmniAVSOperator, error) {
	return _OmniAVS.Contract.Operators(&_OmniAVS.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0xe673df8a.
//
// Solidity: function operators() view returns((address,uint96,uint96)[])
func (_OmniAVS *OmniAVSCallerSession) Operators() ([]IOmniAVSOperator, error) {
	return _OmniAVS.Contract.Operators(&_OmniAVS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniAVS *OmniAVSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniAVS *OmniAVSSession) Owner() (common.Address, error) {
	return _OmniAVS.Contract.Owner(&_OmniAVS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniAVS *OmniAVSCallerSession) Owner() (common.Address, error) {
	return _OmniAVS.Contract.Owner(&_OmniAVS.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniAVS *OmniAVSCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniAVS *OmniAVSSession) Paused() (bool, error) {
	return _OmniAVS.Contract.Paused(&_OmniAVS.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniAVS *OmniAVSCallerSession) Paused() (bool, error) {
	return _OmniAVS.Contract.Paused(&_OmniAVS.CallOpts)
}

// StrategyParams is a free data retrieval call binding the contract method 0xf57f33d0.
//
// Solidity: function strategyParams() view returns((address,uint96)[])
func (_OmniAVS *OmniAVSCaller) StrategyParams(opts *bind.CallOpts) ([]IOmniAVSStrategyParam, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "strategyParams")

	if err != nil {
		return *new([]IOmniAVSStrategyParam), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOmniAVSStrategyParam)).(*[]IOmniAVSStrategyParam)

	return out0, err

}

// StrategyParams is a free data retrieval call binding the contract method 0xf57f33d0.
//
// Solidity: function strategyParams() view returns((address,uint96)[])
func (_OmniAVS *OmniAVSSession) StrategyParams() ([]IOmniAVSStrategyParam, error) {
	return _OmniAVS.Contract.StrategyParams(&_OmniAVS.CallOpts)
}

// StrategyParams is a free data retrieval call binding the contract method 0xf57f33d0.
//
// Solidity: function strategyParams() view returns((address,uint96)[])
func (_OmniAVS *OmniAVSCallerSession) StrategyParams() ([]IOmniAVSStrategyParam, error) {
	return _OmniAVS.Contract.StrategyParams(&_OmniAVS.CallOpts)
}

// XcallBaseGasLimit is a free data retrieval call binding the contract method 0x243d51c7.
//
// Solidity: function xcallBaseGasLimit() view returns(uint64)
func (_OmniAVS *OmniAVSCaller) XcallBaseGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "xcallBaseGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XcallBaseGasLimit is a free data retrieval call binding the contract method 0x243d51c7.
//
// Solidity: function xcallBaseGasLimit() view returns(uint64)
func (_OmniAVS *OmniAVSSession) XcallBaseGasLimit() (uint64, error) {
	return _OmniAVS.Contract.XcallBaseGasLimit(&_OmniAVS.CallOpts)
}

// XcallBaseGasLimit is a free data retrieval call binding the contract method 0x243d51c7.
//
// Solidity: function xcallBaseGasLimit() view returns(uint64)
func (_OmniAVS *OmniAVSCallerSession) XcallBaseGasLimit() (uint64, error) {
	return _OmniAVS.Contract.XcallBaseGasLimit(&_OmniAVS.CallOpts)
}

// XcallGasLimitPerOperator is a free data retrieval call binding the contract method 0x54c74ed3.
//
// Solidity: function xcallGasLimitPerOperator() view returns(uint64)
func (_OmniAVS *OmniAVSCaller) XcallGasLimitPerOperator(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "xcallGasLimitPerOperator")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XcallGasLimitPerOperator is a free data retrieval call binding the contract method 0x54c74ed3.
//
// Solidity: function xcallGasLimitPerOperator() view returns(uint64)
func (_OmniAVS *OmniAVSSession) XcallGasLimitPerOperator() (uint64, error) {
	return _OmniAVS.Contract.XcallGasLimitPerOperator(&_OmniAVS.CallOpts)
}

// XcallGasLimitPerOperator is a free data retrieval call binding the contract method 0x54c74ed3.
//
// Solidity: function xcallGasLimitPerOperator() view returns(uint64)
func (_OmniAVS *OmniAVSCallerSession) XcallGasLimitPerOperator() (uint64, error) {
	return _OmniAVS.Contract.XcallGasLimitPerOperator(&_OmniAVS.CallOpts)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_OmniAVS *OmniAVSTransactor) AddToAllowlist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "addToAllowlist", operator)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_OmniAVS *OmniAVSSession) AddToAllowlist(operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.AddToAllowlist(&_OmniAVS.TransactOpts, operator)
}

// AddToAllowlist is a paid mutator transaction binding the contract method 0xf8e86ece.
//
// Solidity: function addToAllowlist(address operator) returns()
func (_OmniAVS *OmniAVSTransactorSession) AddToAllowlist(operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.AddToAllowlist(&_OmniAVS.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_OmniAVS *OmniAVSTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_OmniAVS *OmniAVSSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.DeregisterOperatorFromAVS(&_OmniAVS.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_OmniAVS *OmniAVSTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.DeregisterOperatorFromAVS(&_OmniAVS.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d179cf.
//
// Solidity: function initialize(address owner_, address omni_, uint64 omniChainId_, address ethStakeInbox_, (address,uint96)[] strategyParams_) returns()
func (_OmniAVS *OmniAVSTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, omni_ common.Address, omniChainId_ uint64, ethStakeInbox_ common.Address, strategyParams_ []IOmniAVSStrategyParam) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "initialize", owner_, omni_, omniChainId_, ethStakeInbox_, strategyParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d179cf.
//
// Solidity: function initialize(address owner_, address omni_, uint64 omniChainId_, address ethStakeInbox_, (address,uint96)[] strategyParams_) returns()
func (_OmniAVS *OmniAVSSession) Initialize(owner_ common.Address, omni_ common.Address, omniChainId_ uint64, ethStakeInbox_ common.Address, strategyParams_ []IOmniAVSStrategyParam) (*types.Transaction, error) {
	return _OmniAVS.Contract.Initialize(&_OmniAVS.TransactOpts, owner_, omni_, omniChainId_, ethStakeInbox_, strategyParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d179cf.
//
// Solidity: function initialize(address owner_, address omni_, uint64 omniChainId_, address ethStakeInbox_, (address,uint96)[] strategyParams_) returns()
func (_OmniAVS *OmniAVSTransactorSession) Initialize(owner_ common.Address, omni_ common.Address, omniChainId_ uint64, ethStakeInbox_ common.Address, strategyParams_ []IOmniAVSStrategyParam) (*types.Transaction, error) {
	return _OmniAVS.Contract.Initialize(&_OmniAVS.TransactOpts, owner_, omni_, omniChainId_, ethStakeInbox_, strategyParams_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniAVS *OmniAVSTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniAVS *OmniAVSSession) Pause() (*types.Transaction, error) {
	return _OmniAVS.Contract.Pause(&_OmniAVS.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniAVS *OmniAVSTransactorSession) Pause() (*types.Transaction, error) {
	return _OmniAVS.Contract.Pause(&_OmniAVS.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_OmniAVS *OmniAVSTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_OmniAVS *OmniAVSSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _OmniAVS.Contract.RegisterOperatorToAVS(&_OmniAVS.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_OmniAVS *OmniAVSTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _OmniAVS.Contract.RegisterOperatorToAVS(&_OmniAVS.TransactOpts, operator, operatorSignature)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_OmniAVS *OmniAVSTransactor) RemoveFromAllowlist(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "removeFromAllowlist", operator)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_OmniAVS *OmniAVSSession) RemoveFromAllowlist(operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.RemoveFromAllowlist(&_OmniAVS.TransactOpts, operator)
}

// RemoveFromAllowlist is a paid mutator transaction binding the contract method 0x5da93d7e.
//
// Solidity: function removeFromAllowlist(address operator) returns()
func (_OmniAVS *OmniAVSTransactorSession) RemoveFromAllowlist(operator common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.RemoveFromAllowlist(&_OmniAVS.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniAVS *OmniAVSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniAVS *OmniAVSSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniAVS.Contract.RenounceOwnership(&_OmniAVS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniAVS *OmniAVSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniAVS.Contract.RenounceOwnership(&_OmniAVS.TransactOpts)
}

// SetEthStakeInbox is a paid mutator transaction binding the contract method 0xd17efb36.
//
// Solidity: function setEthStakeInbox(address inbox) returns()
func (_OmniAVS *OmniAVSTransactor) SetEthStakeInbox(opts *bind.TransactOpts, inbox common.Address) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setEthStakeInbox", inbox)
}

// SetEthStakeInbox is a paid mutator transaction binding the contract method 0xd17efb36.
//
// Solidity: function setEthStakeInbox(address inbox) returns()
func (_OmniAVS *OmniAVSSession) SetEthStakeInbox(inbox common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetEthStakeInbox(&_OmniAVS.TransactOpts, inbox)
}

// SetEthStakeInbox is a paid mutator transaction binding the contract method 0xd17efb36.
//
// Solidity: function setEthStakeInbox(address inbox) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetEthStakeInbox(inbox common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetEthStakeInbox(&_OmniAVS.TransactOpts, inbox)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string metadataURI) returns()
func (_OmniAVS *OmniAVSTransactor) SetMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setMetadataURI", metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string metadataURI) returns()
func (_OmniAVS *OmniAVSSession) SetMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetMetadataURI(&_OmniAVS.TransactOpts, metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string metadataURI) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetMetadataURI(&_OmniAVS.TransactOpts, metadataURI)
}

// SetOmniChainId is a paid mutator transaction binding the contract method 0x5c78b0e2.
//
// Solidity: function setOmniChainId(uint64 chainId) returns()
func (_OmniAVS *OmniAVSTransactor) SetOmniChainId(opts *bind.TransactOpts, chainId uint64) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setOmniChainId", chainId)
}

// SetOmniChainId is a paid mutator transaction binding the contract method 0x5c78b0e2.
//
// Solidity: function setOmniChainId(uint64 chainId) returns()
func (_OmniAVS *OmniAVSSession) SetOmniChainId(chainId uint64) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetOmniChainId(&_OmniAVS.TransactOpts, chainId)
}

// SetOmniChainId is a paid mutator transaction binding the contract method 0x5c78b0e2.
//
// Solidity: function setOmniChainId(uint64 chainId) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetOmniChainId(chainId uint64) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetOmniChainId(&_OmniAVS.TransactOpts, chainId)
}

// SetOmniPortal is a paid mutator transaction binding the contract method 0x0c415884.
//
// Solidity: function setOmniPortal(address portal) returns()
func (_OmniAVS *OmniAVSTransactor) SetOmniPortal(opts *bind.TransactOpts, portal common.Address) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setOmniPortal", portal)
}

// SetOmniPortal is a paid mutator transaction binding the contract method 0x0c415884.
//
// Solidity: function setOmniPortal(address portal) returns()
func (_OmniAVS *OmniAVSSession) SetOmniPortal(portal common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetOmniPortal(&_OmniAVS.TransactOpts, portal)
}

// SetOmniPortal is a paid mutator transaction binding the contract method 0x0c415884.
//
// Solidity: function setOmniPortal(address portal) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetOmniPortal(portal common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetOmniPortal(&_OmniAVS.TransactOpts, portal)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_OmniAVS *OmniAVSTransactor) SetStrategyParams(opts *bind.TransactOpts, params []IOmniAVSStrategyParam) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setStrategyParams", params)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_OmniAVS *OmniAVSSession) SetStrategyParams(params []IOmniAVSStrategyParam) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetStrategyParams(&_OmniAVS.TransactOpts, params)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetStrategyParams(params []IOmniAVSStrategyParam) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetStrategyParams(&_OmniAVS.TransactOpts, params)
}

// SetXcallGasLimits is a paid mutator transaction binding the contract method 0x4a8fa459.
//
// Solidity: function setXcallGasLimits(uint64 base, uint64 perOperator) returns()
func (_OmniAVS *OmniAVSTransactor) SetXcallGasLimits(opts *bind.TransactOpts, base uint64, perOperator uint64) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setXcallGasLimits", base, perOperator)
}

// SetXcallGasLimits is a paid mutator transaction binding the contract method 0x4a8fa459.
//
// Solidity: function setXcallGasLimits(uint64 base, uint64 perOperator) returns()
func (_OmniAVS *OmniAVSSession) SetXcallGasLimits(base uint64, perOperator uint64) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetXcallGasLimits(&_OmniAVS.TransactOpts, base, perOperator)
}

// SetXcallGasLimits is a paid mutator transaction binding the contract method 0x4a8fa459.
//
// Solidity: function setXcallGasLimits(uint64 base, uint64 perOperator) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetXcallGasLimits(base uint64, perOperator uint64) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetXcallGasLimits(&_OmniAVS.TransactOpts, base, perOperator)
}

// SyncWithOmni is a paid mutator transaction binding the contract method 0x13efbe92.
//
// Solidity: function syncWithOmni() payable returns()
func (_OmniAVS *OmniAVSTransactor) SyncWithOmni(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "syncWithOmni")
}

// SyncWithOmni is a paid mutator transaction binding the contract method 0x13efbe92.
//
// Solidity: function syncWithOmni() payable returns()
func (_OmniAVS *OmniAVSSession) SyncWithOmni() (*types.Transaction, error) {
	return _OmniAVS.Contract.SyncWithOmni(&_OmniAVS.TransactOpts)
}

// SyncWithOmni is a paid mutator transaction binding the contract method 0x13efbe92.
//
// Solidity: function syncWithOmni() payable returns()
func (_OmniAVS *OmniAVSTransactorSession) SyncWithOmni() (*types.Transaction, error) {
	return _OmniAVS.Contract.SyncWithOmni(&_OmniAVS.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniAVS *OmniAVSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniAVS *OmniAVSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.TransferOwnership(&_OmniAVS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniAVS *OmniAVSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniAVS.Contract.TransferOwnership(&_OmniAVS.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniAVS *OmniAVSTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniAVS *OmniAVSSession) Unpause() (*types.Transaction, error) {
	return _OmniAVS.Contract.Unpause(&_OmniAVS.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniAVS *OmniAVSTransactorSession) Unpause() (*types.Transaction, error) {
	return _OmniAVS.Contract.Unpause(&_OmniAVS.TransactOpts)
}

// OmniAVSInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniAVS contract.
type OmniAVSInitializedIterator struct {
	Event *OmniAVSInitialized // Event containing the contract specifics and raw log

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
func (it *OmniAVSInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniAVSInitialized)
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
		it.Event = new(OmniAVSInitialized)
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
func (it *OmniAVSInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniAVSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniAVSInitialized represents a Initialized event raised by the OmniAVS contract.
type OmniAVSInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniAVS *OmniAVSFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniAVSInitializedIterator, error) {

	logs, sub, err := _OmniAVS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniAVSInitializedIterator{contract: _OmniAVS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OmniAVS *OmniAVSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniAVSInitialized) (event.Subscription, error) {

	logs, sub, err := _OmniAVS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniAVSInitialized)
				if err := _OmniAVS.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OmniAVS *OmniAVSFilterer) ParseInitialized(log types.Log) (*OmniAVSInitialized, error) {
	event := new(OmniAVSInitialized)
	if err := _OmniAVS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniAVSOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the OmniAVS contract.
type OmniAVSOperatorAddedIterator struct {
	Event *OmniAVSOperatorAdded // Event containing the contract specifics and raw log

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
func (it *OmniAVSOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniAVSOperatorAdded)
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
		it.Event = new(OmniAVSOperatorAdded)
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
func (it *OmniAVSOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniAVSOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniAVSOperatorAdded represents a OperatorAdded event raised by the OmniAVS contract.
type OmniAVSOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_OmniAVS *OmniAVSFilterer) FilterOperatorAdded(opts *bind.FilterOpts, operator []common.Address) (*OmniAVSOperatorAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OmniAVS.contract.FilterLogs(opts, "OperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &OmniAVSOperatorAddedIterator{contract: _OmniAVS.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_OmniAVS *OmniAVSFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *OmniAVSOperatorAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OmniAVS.contract.WatchLogs(opts, "OperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniAVSOperatorAdded)
				if err := _OmniAVS.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed operator)
func (_OmniAVS *OmniAVSFilterer) ParseOperatorAdded(log types.Log) (*OmniAVSOperatorAdded, error) {
	event := new(OmniAVSOperatorAdded)
	if err := _OmniAVS.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniAVSOperatorAllowedIterator is returned from FilterOperatorAllowed and is used to iterate over the raw logs and unpacked data for OperatorAllowed events raised by the OmniAVS contract.
type OmniAVSOperatorAllowedIterator struct {
	Event *OmniAVSOperatorAllowed // Event containing the contract specifics and raw log

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
func (it *OmniAVSOperatorAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniAVSOperatorAllowed)
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
		it.Event = new(OmniAVSOperatorAllowed)
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
func (it *OmniAVSOperatorAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniAVSOperatorAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniAVSOperatorAllowed represents a OperatorAllowed event raised by the OmniAVS contract.
type OmniAVSOperatorAllowed struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAllowed is a free log retrieval operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address operator)
func (_OmniAVS *OmniAVSFilterer) FilterOperatorAllowed(opts *bind.FilterOpts) (*OmniAVSOperatorAllowedIterator, error) {

	logs, sub, err := _OmniAVS.contract.FilterLogs(opts, "OperatorAllowed")
	if err != nil {
		return nil, err
	}
	return &OmniAVSOperatorAllowedIterator{contract: _OmniAVS.contract, event: "OperatorAllowed", logs: logs, sub: sub}, nil
}

// WatchOperatorAllowed is a free log subscription operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address operator)
func (_OmniAVS *OmniAVSFilterer) WatchOperatorAllowed(opts *bind.WatchOpts, sink chan<- *OmniAVSOperatorAllowed) (event.Subscription, error) {

	logs, sub, err := _OmniAVS.contract.WatchLogs(opts, "OperatorAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniAVSOperatorAllowed)
				if err := _OmniAVS.contract.UnpackLog(event, "OperatorAllowed", log); err != nil {
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

// ParseOperatorAllowed is a log parse operation binding the contract event 0xdde65206cdee4ea27ef1b170724ba50b41ad09a3bf2dda12935fc40c4dbf6e75.
//
// Solidity: event OperatorAllowed(address operator)
func (_OmniAVS *OmniAVSFilterer) ParseOperatorAllowed(log types.Log) (*OmniAVSOperatorAllowed, error) {
	event := new(OmniAVSOperatorAllowed)
	if err := _OmniAVS.contract.UnpackLog(event, "OperatorAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniAVSOperatorDisallowedIterator is returned from FilterOperatorDisallowed and is used to iterate over the raw logs and unpacked data for OperatorDisallowed events raised by the OmniAVS contract.
type OmniAVSOperatorDisallowedIterator struct {
	Event *OmniAVSOperatorDisallowed // Event containing the contract specifics and raw log

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
func (it *OmniAVSOperatorDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniAVSOperatorDisallowed)
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
		it.Event = new(OmniAVSOperatorDisallowed)
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
func (it *OmniAVSOperatorDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniAVSOperatorDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniAVSOperatorDisallowed represents a OperatorDisallowed event raised by the OmniAVS contract.
type OmniAVSOperatorDisallowed struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDisallowed is a free log retrieval operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address operator)
func (_OmniAVS *OmniAVSFilterer) FilterOperatorDisallowed(opts *bind.FilterOpts) (*OmniAVSOperatorDisallowedIterator, error) {

	logs, sub, err := _OmniAVS.contract.FilterLogs(opts, "OperatorDisallowed")
	if err != nil {
		return nil, err
	}
	return &OmniAVSOperatorDisallowedIterator{contract: _OmniAVS.contract, event: "OperatorDisallowed", logs: logs, sub: sub}, nil
}

// WatchOperatorDisallowed is a free log subscription operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address operator)
func (_OmniAVS *OmniAVSFilterer) WatchOperatorDisallowed(opts *bind.WatchOpts, sink chan<- *OmniAVSOperatorDisallowed) (event.Subscription, error) {

	logs, sub, err := _OmniAVS.contract.WatchLogs(opts, "OperatorDisallowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniAVSOperatorDisallowed)
				if err := _OmniAVS.contract.UnpackLog(event, "OperatorDisallowed", log); err != nil {
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

// ParseOperatorDisallowed is a log parse operation binding the contract event 0x8560daa191dd8e6fba276b053006b3990c46c94b842f85490f52c49b15cfe5cb.
//
// Solidity: event OperatorDisallowed(address operator)
func (_OmniAVS *OmniAVSFilterer) ParseOperatorDisallowed(log types.Log) (*OmniAVSOperatorDisallowed, error) {
	event := new(OmniAVSOperatorDisallowed)
	if err := _OmniAVS.contract.UnpackLog(event, "OperatorDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniAVSOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the OmniAVS contract.
type OmniAVSOperatorRemovedIterator struct {
	Event *OmniAVSOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *OmniAVSOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniAVSOperatorRemoved)
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
		it.Event = new(OmniAVSOperatorRemoved)
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
func (it *OmniAVSOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniAVSOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniAVSOperatorRemoved represents a OperatorRemoved event raised by the OmniAVS contract.
type OmniAVSOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_OmniAVS *OmniAVSFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operator []common.Address) (*OmniAVSOperatorRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OmniAVS.contract.FilterLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &OmniAVSOperatorRemovedIterator{contract: _OmniAVS.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_OmniAVS *OmniAVSFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *OmniAVSOperatorRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _OmniAVS.contract.WatchLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniAVSOperatorRemoved)
				if err := _OmniAVS.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_OmniAVS *OmniAVSFilterer) ParseOperatorRemoved(log types.Log) (*OmniAVSOperatorRemoved, error) {
	event := new(OmniAVSOperatorRemoved)
	if err := _OmniAVS.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniAVSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniAVS contract.
type OmniAVSOwnershipTransferredIterator struct {
	Event *OmniAVSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniAVSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniAVSOwnershipTransferred)
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
		it.Event = new(OmniAVSOwnershipTransferred)
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
func (it *OmniAVSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniAVSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniAVSOwnershipTransferred represents a OwnershipTransferred event raised by the OmniAVS contract.
type OmniAVSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniAVS *OmniAVSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniAVSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniAVS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniAVSOwnershipTransferredIterator{contract: _OmniAVS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniAVS *OmniAVSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniAVSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniAVS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniAVSOwnershipTransferred)
				if err := _OmniAVS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OmniAVS *OmniAVSFilterer) ParseOwnershipTransferred(log types.Log) (*OmniAVSOwnershipTransferred, error) {
	event := new(OmniAVSOwnershipTransferred)
	if err := _OmniAVS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniAVSPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OmniAVS contract.
type OmniAVSPausedIterator struct {
	Event *OmniAVSPaused // Event containing the contract specifics and raw log

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
func (it *OmniAVSPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniAVSPaused)
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
		it.Event = new(OmniAVSPaused)
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
func (it *OmniAVSPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniAVSPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniAVSPaused represents a Paused event raised by the OmniAVS contract.
type OmniAVSPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OmniAVS *OmniAVSFilterer) FilterPaused(opts *bind.FilterOpts) (*OmniAVSPausedIterator, error) {

	logs, sub, err := _OmniAVS.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OmniAVSPausedIterator{contract: _OmniAVS.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OmniAVS *OmniAVSFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OmniAVSPaused) (event.Subscription, error) {

	logs, sub, err := _OmniAVS.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniAVSPaused)
				if err := _OmniAVS.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_OmniAVS *OmniAVSFilterer) ParsePaused(log types.Log) (*OmniAVSPaused, error) {
	event := new(OmniAVSPaused)
	if err := _OmniAVS.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniAVSUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OmniAVS contract.
type OmniAVSUnpausedIterator struct {
	Event *OmniAVSUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniAVSUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniAVSUnpaused)
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
		it.Event = new(OmniAVSUnpaused)
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
func (it *OmniAVSUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniAVSUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniAVSUnpaused represents a Unpaused event raised by the OmniAVS contract.
type OmniAVSUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OmniAVS *OmniAVSFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OmniAVSUnpausedIterator, error) {

	logs, sub, err := _OmniAVS.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OmniAVSUnpausedIterator{contract: _OmniAVS.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OmniAVS *OmniAVSFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OmniAVSUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniAVS.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniAVSUnpaused)
				if err := _OmniAVS.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_OmniAVS *OmniAVSFilterer) ParseUnpaused(log types.Log) (*OmniAVSUnpaused, error) {
	event := new(OmniAVSUnpaused)
	if err := _OmniAVS.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
