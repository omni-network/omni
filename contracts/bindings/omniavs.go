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

// IOmniAVSStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type IOmniAVSStrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// IOmniAVSValidator is an auto generated low-level Go binding around an user-defined struct.
type IOmniAVSValidator struct {
	Addr      common.Address
	Delegated *big.Int
	Staked    *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// OmniAVSMetaData contains all meta data concerning the OmniAVS contract.
var OmniAVSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"delegationManager_\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"avsDirectory_\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WEIGHTING_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeForSync\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOmniAVS.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegated\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"staked\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"},{\"name\":\"omniChainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumOperatorStake_\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"maxOperatorCount_\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategyParams_\",\"type\":\"tuple[]\",\"internalType\":\"structIOmniAVS.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxOperatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumOperatorStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxOperatorCount\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumOperatorStake\",\"inputs\":[{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOmniChainId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOmniPortal\",\"inputs\":[{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIOmniAVS.StrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXcallGasLimits\",\"inputs\":[{\"name\":\"base\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"perValidator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyParams\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"syncWithOmni\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAdded\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemoved\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60c0604052612710606955620124f8606a553480156200001e57600080fd5b506040516200236138038062002361833981016040819052620000419162000140565b6001600160a01b03808316608052811660a0526200005e62000066565b50506200017f565b600054610100900460ff1615620000d35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000125576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013d57600080fd5b50565b600080604083850312156200015457600080fd5b8251620001618162000127565b6020840151909250620001748162000127565b809150509250929050565b60805160a051612199620001c86000396000818161031a015281816108b601528181610ab10152610c6801526000818161073f01528181611237015261173f01526121996000f3fe6080604052600436106101815760003560e01c8063750521f5116100d1578063b7ab4db51161008a578063e28d490611610064578063e28d4906146104d9578063e481af9d146104f9578063f2fde38b1461050e578063f36b8d361461052e57600080fd5b8063b7ab4db514610470578063b98912b014610492578063c75e3aed146104a757600080fd5b8063750521f5146103b25780638da5cb5b146103d25780639926ee7d146103f0578063a364f4da14610410578063ad6426c114610430578063ae30f16d1461045057600080fd5b806345f59f791161013e5780636b3aa72e116101185780636b3aa72e1461030b578063715018a61461033e5780637182a94414610353578063718fef611461039257600080fd5b806345f59f79146102a15780635c78b0e2146102c15780635e5a6775146102e157600080fd5b80630c41588414610186578063110ff5f1146101a857806312466b68146101ed57806313efbe921461023457806333cfb7b71461023c57806339acf9f114610269575b600080fd5b34801561019257600080fd5b506101a66101a13660046118dd565b61054e565b005b3480156101b457600080fd5b506065546101d09064010000000090046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b3480156101f957600080fd5b5061020d6102083660046118fa565b610578565b604080516001600160a01b0390931683526001600160601b039091166020830152016101e4565b6101a66105b3565b34801561024857600080fd5b5061025c6102573660046118dd565b61069b565b6040516101e49190611913565b34801561027557600080fd5b50606754610289906001600160a01b031681565b6040516001600160a01b0390911681526020016101e4565b3480156102ad57600080fd5b506101a66102bc366004611975565b610800565b3480156102cd57600080fd5b506101a66102dc3660046119ae565b610839565b3480156102ed57600080fd5b506102fd670de0b6b3a764000081565b6040519081526020016101e4565b34801561031757600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610289565b34801561034a57600080fd5b506101a6610870565b34801561035f57600080fd5b5060655461037a90600160601b90046001600160601b031681565b6040516001600160601b0390911681526020016101e4565b34801561039e57600080fd5b506101a66103ad3660046119c9565b610884565b3480156103be57600080fd5b506101a66103cd366004611ab0565b610897565b3480156103de57600080fd5b506033546001600160a01b0316610289565b3480156103fc57600080fd5b506101a661040b366004611b00565b610920565b34801561041c57600080fd5b506101a661042b3660046118dd565b610ba3565b34801561043c57600080fd5b506101a661044b366004611c09565b610d04565b34801561045c57600080fd5b506101a661046b366004611ca5565b610e9a565b34801561047c57600080fd5b50610485610eb0565b6040516101e49190611ce6565b34801561049e57600080fd5b506102fd610ebf565b3480156104b357600080fd5b506065546104c49063ffffffff1681565b60405163ffffffff90911681526020016101e4565b3480156104e557600080fd5b506102896104f43660046118fa565b610fb3565b34801561050557600080fd5b5061025c610fdd565b34801561051a57600080fd5b506101a66105293660046118dd565b6110a5565b34801561053a57600080fd5b506101a6610549366004611d55565b61111e565b610556611142565b606780546001600160a01b0319166001600160a01b0392909216919091179055565b6068818154811061058857600080fd5b6000918252602090912001546001600160a01b0381169150600160a01b90046001600160601b031682565b60006105bd610eb0565b6067546065546040519293506001600160a01b03909116916370e8b56a9134916401000000009091046001600160401b0316906000906333364ffb60e11b9061060a908890602401611ce6565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915286516106479061119c565b6040518663ffffffff1660e01b81526004016106669493929190611dbd565b6000604051808303818588803b15801561067f57600080fd5b505af1158015610693573d6000803e3d6000fd5b505050505050565b6068546060906000906001600160401b038111156106bb576106bb6119eb565b6040519080825280602002602001820160405280156106e4578160200160208202803683370190505b50905060005b6068548110156107f95760006068828154811061070957610709611e02565b600091825260208220015460405163778e55f360e01b81526001600160a01b0388811660048301529182166024820181905293507f00000000000000000000000000000000000000000000000000000000000000009091169063778e55f390604401602060405180830381865afa158015610788573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ac9190611e18565b11156107e657808383815181106107c5576107c5611e02565b60200260200101906001600160a01b031690816001600160a01b0316815250505b50806107f181611e47565b9150506106ea565b5092915050565b610808611142565b606580546001600160601b03909216600160601b026bffffffffffffffffffffffff60601b19909216919091179055565b610841611142565b606580546001600160401b03909216640100000000026bffffffffffffffff0000000019909216919091179055565b610878611142565b61088260006111bf565b565b61088c611142565b606a91909155606955565b61089f611142565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906108eb908490600401611e62565b600060405180830381600087803b15801561090557600080fd5b505af1158015610919573d6000803e3d6000fd5b5050505050565b336001600160a01b038316146109765760405162461bcd60e51b815260206004820152601660248201527527b6b734a0ab299d1037b7363c9037b832b930ba37b960511b60448201526064015b60405180910390fd5b60655460665463ffffffff909116116109d15760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694156533a206d6178206f70657261746f727320726561636865640000604482015260640161096d565b606554600160601b90046001600160601b03166109ed83611211565b6001600160601b03161015610a445760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e694156533a206d696e696d756d207374616b65206e6f74206d65740000604482015260640161096d565b610a4d82611422565b15610a9a5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e694156533a20616c726561647920616e206f70657261746f7200000000604482015260640161096d565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610ae89085908590600401611e75565b600060405180830381600087803b158015610b0257600080fd5b505af1158015610b16573d6000803e3d6000fd5b5050606680546001810182556000919091527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540180546001600160a01b0319166001600160a01b03861617905550610b6b9050565b6040516001600160a01b038316907fac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d90600090a25050565b336001600160a01b03821614610bf45760405162461bcd60e51b815260206004820152601660248201527527b6b734a0ab299d1037b7363c9037b832b930ba37b960511b604482015260640161096d565b610bfd81611422565b610c495760405162461bcd60e51b815260206004820152601860248201527f4f6d6e694156533a206e6f7420616e206f70657261746f720000000000000000604482015260640161096d565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401600060405180830381600087803b158015610cac57600080fd5b505af1158015610cc0573d6000803e3d6000fd5b50505050610ccd8161148c565b6040516001600160a01b038216907f80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d90600090a250565b600054610100900460ff1615808015610d245750600054600160ff909116105b80610d3e5750303b158015610d3e575060005460ff166001145b610da15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161096d565b6000805460ff191660011790558015610dc4576000805461ff0019166101001790555b610dcd886111bf565b606780546001600160a01b0319166001600160a01b03891617905560658054640100000000600160c01b0319166401000000006001600160401b038916026bffffffffffffffffffffffff60601b191617600160601b6001600160601b038816021763ffffffff191663ffffffff8616179055610e4a8383611594565b8015610e90576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050505050565b610ea2611142565b610eac8282611594565b5050565b6060610eba611605565b905090565b600080610eca610eb0565b6067546065546040519293506001600160a01b0390911691638dd9523c9164010000000090046001600160401b0316906333364ffb60e11b90610f11908690602401611ce6565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091528451610f4e9061119c565b6040518463ffffffff1660e01b8152600401610f6c93929190611ec0565b602060405180830381865afa158015610f89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fad9190611e18565b91505090565b60668181548110610fc357600080fd5b6000918252602090912001546001600160a01b0316905081565b6068546060906000906001600160401b03811115610ffd57610ffd6119eb565b604051908082528060200260200182016040528015611026578160200160208202803683370190505b50905060005b60685481101561109f576068818154811061104957611049611e02565b60009182526020909120015482516001600160a01b039091169083908390811061107557611075611e02565b6001600160a01b03909216602092830291909101909101528061109781611e47565b91505061102c565b50919050565b6110ad611142565b6001600160a01b0381166111125760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161096d565b61111b816111bf565b50565b611126611142565b6065805463ffffffff191663ffffffff92909216919091179055565b6033546001600160a01b031633146108825760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161096d565b6000606a54606954836111af9190611ef6565b6111b99190611f15565b92915050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040516367c0439f60e11b81526001600160a01b038281166004830152600091829182917f00000000000000000000000000000000000000000000000000000000000000009091169063cf80873e90602401600060405180830381865afa158015611280573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112a89190810190611fbb565b90925090506000805b83518110156114195760008482815181106112ce576112ce611e02565b6020026020010151905060008483815181106112ec576112ec611e02565b60200260200101519050611325604051806040016040528060006001600160a01b0316815260200160006001600160601b031681525090565b60005b6068548110156113d157836001600160a01b03166068828154811061134f5761134f611e02565b6000918252602090912001546001600160a01b031614156113bf576068818154811061137d5761137d611e02565b6000918252602091829020604080518082019091529101546001600160a01b0381168252600160a01b90046001600160601b03169181019190915291506113d1565b806113c981611e47565b915050611328565b5080516001600160a01b03166113e957505050611407565b6113f7828260200151611867565b611401908661207f565b94505050505b8061141181611e47565b9150506112b1565b50949350505050565b6000805b60665481101561148357826001600160a01b03166066828154811061144d5761144d611e02565b6000918252602090912001546001600160a01b031614156114715750600192915050565b8061147b81611e47565b915050611426565b50600092915050565b60005b606654811015610eac57816001600160a01b0316606682815481106114b6576114b6611e02565b6000918252602090912001546001600160a01b0316141561158257606680546114e1906001906120aa565b815481106114f1576114f1611e02565b600091825260209091200154606680546001600160a01b03909216918390811061151d5761151d611e02565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550606680548061155c5761155c6120c1565b600082815260209020810160001990810180546001600160a01b03191690550190555050565b8061158c81611e47565b91505061148f565b6115a060686000611896565b60005b818110156116005760688383838181106115bf576115bf611e02565b835460018101855560009485526020909420604090910292909201929190910190506115eb82826120d7565b505080806115f890611e47565b9150506115a3565b505050565b6066546060906000906001600160401b03811115611625576116256119eb565b60405190808252806020026020018201604052801561167057816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816116435790505b50905060005b815181101561109f5760006066828154811061169457611694611e02565b6000918252602080832090910154604080518082019091528381529182018390526001600160a01b0316925060005b6068548110156117e057606881815481106116e0576116e0611e02565b6000918252602080832060408051808201825293909101546001600160a01b03808216808652600160a01b9092046001600160601b031693850193909352905163778e55f360e01b8152888316600482015260248101919091529194507f0000000000000000000000000000000000000000000000000000000000000000169063778e55f390604401602060405180830381865afa158015611786573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117aa9190611e18565b905080156117cd576117c0818460200151611867565b6117ca908561207f565b93505b50806117d881611e47565b9150506116c3565b5060006117ec84611211565b905060006117fa8285612119565b90506040518060600160405280866001600160a01b03168152602001826001600160601b03168152602001836001600160601b031681525087878151811061184457611844611e02565b60200260200101819052505050505050808061185f90611e47565b915050611676565b6000670de0b6b3a76400006118856001600160601b03841685611ef6565b61188f9190612141565b9392505050565b508054600082559060005260206000209081019061111b91905b808211156118c457600081556001016118b0565b5090565b6001600160a01b038116811461111b57600080fd5b6000602082840312156118ef57600080fd5b813561188f816118c8565b60006020828403121561190c57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156119545783516001600160a01b03168352928401929184019160010161192f565b50909695505050505050565b6001600160601b038116811461111b57600080fd5b60006020828403121561198757600080fd5b813561188f81611960565b80356001600160401b03811681146119a957600080fd5b919050565b6000602082840312156119c057600080fd5b61188f82611992565b600080604083850312156119dc57600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715611a2357611a236119eb565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611a5157611a516119eb565b604052919050565b60006001600160401b03831115611a7257611a726119eb565b611a85601f8401601f1916602001611a29565b9050828152838383011115611a9957600080fd5b828260208301376000602084830101529392505050565b600060208284031215611ac257600080fd5b81356001600160401b03811115611ad857600080fd5b8201601f81018413611ae957600080fd5b611af884823560208401611a59565b949350505050565b60008060408385031215611b1357600080fd5b8235611b1e816118c8565b915060208301356001600160401b0380821115611b3a57600080fd5b9084019060608287031215611b4e57600080fd5b611b56611a01565b823582811115611b6557600080fd5b83019150601f82018713611b7857600080fd5b611b8787833560208501611a59565b815260208301356020820152604083013560408201528093505050509250929050565b803563ffffffff811681146119a957600080fd5b60008083601f840112611bd057600080fd5b5081356001600160401b03811115611be757600080fd5b6020830191508360208260061b8501011115611c0257600080fd5b9250929050565b600080600080600080600060c0888a031215611c2457600080fd5b8735611c2f816118c8565b96506020880135611c3f816118c8565b9550611c4d60408901611992565b94506060880135611c5d81611960565b9350611c6b60808901611baa565b925060a08801356001600160401b03811115611c8657600080fd5b611c928a828b01611bbe565b989b979a50959850939692959293505050565b60008060208385031215611cb857600080fd5b82356001600160401b03811115611cce57600080fd5b611cda85828601611bbe565b90969095509350505050565b602080825282518282018190526000919060409081850190868401855b82811015611d4857815180516001600160a01b03168552868101516001600160601b039081168887015290860151168585015260609093019290850190600101611d03565b5091979650505050505050565b600060208284031215611d6757600080fd5b61188f82611baa565b6000815180845260005b81811015611d9657602081850181015186830182015201611d7a565b81811115611da8576000602083870101525b50601f01601f19169290920160200192915050565b60006001600160401b03808716835260018060a01b038616602084015260806040840152611dee6080840186611d70565b915080841660608401525095945050505050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611e2a57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b6000600019821415611e5b57611e5b611e31565b5060010190565b60208152600061188f6020830184611d70565b60018060a01b0383168152604060208201526000825160606040840152611e9f60a0840182611d70565b90506020840151606084015260408401516080840152809150509392505050565b60006001600160401b03808616835260606020840152611ee36060840186611d70565b9150808416604084015250949350505050565b6000816000190483118215151615611f1057611f10611e31565b500290565b60008219821115611f2857611f28611e31565b500190565b60006001600160401b03821115611f4657611f466119eb565b5060051b60200190565b600082601f830112611f6157600080fd5b81516020611f76611f7183611f2d565b611a29565b82815260059290921b84018101918181019086841115611f9557600080fd5b8286015b84811015611fb05780518352918301918301611f99565b509695505050505050565b60008060408385031215611fce57600080fd5b82516001600160401b0380821115611fe557600080fd5b818501915085601f830112611ff957600080fd5b81516020612009611f7183611f2d565b82815260059290921b8401810191818101908984111561202857600080fd5b948201945b8386101561204f578551612040816118c8565b8252948201949082019061202d565b9188015191965090935050508082111561206857600080fd5b5061207585828601611f50565b9150509250929050565b60006001600160601b038083168185168083038211156120a1576120a1611e31565b01949350505050565b6000828210156120bc576120bc611e31565b500390565b634e487b7160e01b600052603160045260246000fd5b81356120e2816118c8565b81546001600160a01b03199081166001600160a01b03929092169182178355602084013561210f81611960565b60a01b1617905550565b60006001600160601b038381169083168181101561213957612139611e31565b039392505050565b60008261215e57634e487b7160e01b600052601260045260246000fd5b50049056fea26469706673582212202525e3056faeeea907718ec1f9fc179f39538cee5dd29e27f8baa9fd136e624564736f6c634300080c0033",
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

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_OmniAVS *OmniAVSCaller) WEIGHTINGDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "WEIGHTING_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_OmniAVS *OmniAVSSession) WEIGHTINGDIVISOR() (*big.Int, error) {
	return _OmniAVS.Contract.WEIGHTINGDIVISOR(&_OmniAVS.CallOpts)
}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_OmniAVS *OmniAVSCallerSession) WEIGHTINGDIVISOR() (*big.Int, error) {
	return _OmniAVS.Contract.WEIGHTINGDIVISOR(&_OmniAVS.CallOpts)
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

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint96,uint96)[])
func (_OmniAVS *OmniAVSCaller) GetValidators(opts *bind.CallOpts) ([]IOmniAVSValidator, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]IOmniAVSValidator), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOmniAVSValidator)).(*[]IOmniAVSValidator)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint96,uint96)[])
func (_OmniAVS *OmniAVSSession) GetValidators() ([]IOmniAVSValidator, error) {
	return _OmniAVS.Contract.GetValidators(&_OmniAVS.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint96,uint96)[])
func (_OmniAVS *OmniAVSCallerSession) GetValidators() ([]IOmniAVSValidator, error) {
	return _OmniAVS.Contract.GetValidators(&_OmniAVS.CallOpts)
}

// MaxOperatorCount is a free data retrieval call binding the contract method 0xc75e3aed.
//
// Solidity: function maxOperatorCount() view returns(uint32)
func (_OmniAVS *OmniAVSCaller) MaxOperatorCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "maxOperatorCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorCount is a free data retrieval call binding the contract method 0xc75e3aed.
//
// Solidity: function maxOperatorCount() view returns(uint32)
func (_OmniAVS *OmniAVSSession) MaxOperatorCount() (uint32, error) {
	return _OmniAVS.Contract.MaxOperatorCount(&_OmniAVS.CallOpts)
}

// MaxOperatorCount is a free data retrieval call binding the contract method 0xc75e3aed.
//
// Solidity: function maxOperatorCount() view returns(uint32)
func (_OmniAVS *OmniAVSCallerSession) MaxOperatorCount() (uint32, error) {
	return _OmniAVS.Contract.MaxOperatorCount(&_OmniAVS.CallOpts)
}

// MinimumOperatorStake is a free data retrieval call binding the contract method 0x7182a944.
//
// Solidity: function minimumOperatorStake() view returns(uint96)
func (_OmniAVS *OmniAVSCaller) MinimumOperatorStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "minimumOperatorStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumOperatorStake is a free data retrieval call binding the contract method 0x7182a944.
//
// Solidity: function minimumOperatorStake() view returns(uint96)
func (_OmniAVS *OmniAVSSession) MinimumOperatorStake() (*big.Int, error) {
	return _OmniAVS.Contract.MinimumOperatorStake(&_OmniAVS.CallOpts)
}

// MinimumOperatorStake is a free data retrieval call binding the contract method 0x7182a944.
//
// Solidity: function minimumOperatorStake() view returns(uint96)
func (_OmniAVS *OmniAVSCallerSession) MinimumOperatorStake() (*big.Int, error) {
	return _OmniAVS.Contract.MinimumOperatorStake(&_OmniAVS.CallOpts)
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

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(address)
func (_OmniAVS *OmniAVSCaller) Operators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(address)
func (_OmniAVS *OmniAVSSession) Operators(arg0 *big.Int) (common.Address, error) {
	return _OmniAVS.Contract.Operators(&_OmniAVS.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(address)
func (_OmniAVS *OmniAVSCallerSession) Operators(arg0 *big.Int) (common.Address, error) {
	return _OmniAVS.Contract.Operators(&_OmniAVS.CallOpts, arg0)
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

// StrategyParams is a free data retrieval call binding the contract method 0x12466b68.
//
// Solidity: function strategyParams(uint256 ) view returns(address strategy, uint96 multiplier)
func (_OmniAVS *OmniAVSCaller) StrategyParams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	var out []interface{}
	err := _OmniAVS.contract.Call(opts, &out, "strategyParams", arg0)

	outstruct := new(struct {
		Strategy   common.Address
		Multiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Strategy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Multiplier = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StrategyParams is a free data retrieval call binding the contract method 0x12466b68.
//
// Solidity: function strategyParams(uint256 ) view returns(address strategy, uint96 multiplier)
func (_OmniAVS *OmniAVSSession) StrategyParams(arg0 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	return _OmniAVS.Contract.StrategyParams(&_OmniAVS.CallOpts, arg0)
}

// StrategyParams is a free data retrieval call binding the contract method 0x12466b68.
//
// Solidity: function strategyParams(uint256 ) view returns(address strategy, uint96 multiplier)
func (_OmniAVS *OmniAVSCallerSession) StrategyParams(arg0 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	return _OmniAVS.Contract.StrategyParams(&_OmniAVS.CallOpts, arg0)
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

// Initialize is a paid mutator transaction binding the contract method 0xad6426c1.
//
// Solidity: function initialize(address owner_, address omni_, uint64 omniChainId_, uint96 minimumOperatorStake_, uint32 maxOperatorCount_, (address,uint96)[] strategyParams_) returns()
func (_OmniAVS *OmniAVSTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, omni_ common.Address, omniChainId_ uint64, minimumOperatorStake_ *big.Int, maxOperatorCount_ uint32, strategyParams_ []IOmniAVSStrategyParams) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "initialize", owner_, omni_, omniChainId_, minimumOperatorStake_, maxOperatorCount_, strategyParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0xad6426c1.
//
// Solidity: function initialize(address owner_, address omni_, uint64 omniChainId_, uint96 minimumOperatorStake_, uint32 maxOperatorCount_, (address,uint96)[] strategyParams_) returns()
func (_OmniAVS *OmniAVSSession) Initialize(owner_ common.Address, omni_ common.Address, omniChainId_ uint64, minimumOperatorStake_ *big.Int, maxOperatorCount_ uint32, strategyParams_ []IOmniAVSStrategyParams) (*types.Transaction, error) {
	return _OmniAVS.Contract.Initialize(&_OmniAVS.TransactOpts, owner_, omni_, omniChainId_, minimumOperatorStake_, maxOperatorCount_, strategyParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0xad6426c1.
//
// Solidity: function initialize(address owner_, address omni_, uint64 omniChainId_, uint96 minimumOperatorStake_, uint32 maxOperatorCount_, (address,uint96)[] strategyParams_) returns()
func (_OmniAVS *OmniAVSTransactorSession) Initialize(owner_ common.Address, omni_ common.Address, omniChainId_ uint64, minimumOperatorStake_ *big.Int, maxOperatorCount_ uint32, strategyParams_ []IOmniAVSStrategyParams) (*types.Transaction, error) {
	return _OmniAVS.Contract.Initialize(&_OmniAVS.TransactOpts, owner_, omni_, omniChainId_, minimumOperatorStake_, maxOperatorCount_, strategyParams_)
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

// SetMaxOperatorCount is a paid mutator transaction binding the contract method 0xf36b8d36.
//
// Solidity: function setMaxOperatorCount(uint32 count) returns()
func (_OmniAVS *OmniAVSTransactor) SetMaxOperatorCount(opts *bind.TransactOpts, count uint32) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setMaxOperatorCount", count)
}

// SetMaxOperatorCount is a paid mutator transaction binding the contract method 0xf36b8d36.
//
// Solidity: function setMaxOperatorCount(uint32 count) returns()
func (_OmniAVS *OmniAVSSession) SetMaxOperatorCount(count uint32) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetMaxOperatorCount(&_OmniAVS.TransactOpts, count)
}

// SetMaxOperatorCount is a paid mutator transaction binding the contract method 0xf36b8d36.
//
// Solidity: function setMaxOperatorCount(uint32 count) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetMaxOperatorCount(count uint32) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetMaxOperatorCount(&_OmniAVS.TransactOpts, count)
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

// SetMinimumOperatorStake is a paid mutator transaction binding the contract method 0x45f59f79.
//
// Solidity: function setMinimumOperatorStake(uint96 stake) returns()
func (_OmniAVS *OmniAVSTransactor) SetMinimumOperatorStake(opts *bind.TransactOpts, stake *big.Int) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setMinimumOperatorStake", stake)
}

// SetMinimumOperatorStake is a paid mutator transaction binding the contract method 0x45f59f79.
//
// Solidity: function setMinimumOperatorStake(uint96 stake) returns()
func (_OmniAVS *OmniAVSSession) SetMinimumOperatorStake(stake *big.Int) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetMinimumOperatorStake(&_OmniAVS.TransactOpts, stake)
}

// SetMinimumOperatorStake is a paid mutator transaction binding the contract method 0x45f59f79.
//
// Solidity: function setMinimumOperatorStake(uint96 stake) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetMinimumOperatorStake(stake *big.Int) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetMinimumOperatorStake(&_OmniAVS.TransactOpts, stake)
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
func (_OmniAVS *OmniAVSTransactor) SetStrategyParams(opts *bind.TransactOpts, params []IOmniAVSStrategyParams) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setStrategyParams", params)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_OmniAVS *OmniAVSSession) SetStrategyParams(params []IOmniAVSStrategyParams) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetStrategyParams(&_OmniAVS.TransactOpts, params)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetStrategyParams(params []IOmniAVSStrategyParams) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetStrategyParams(&_OmniAVS.TransactOpts, params)
}

// SetXcallGasLimits is a paid mutator transaction binding the contract method 0x718fef61.
//
// Solidity: function setXcallGasLimits(uint256 base, uint256 perValidator) returns()
func (_OmniAVS *OmniAVSTransactor) SetXcallGasLimits(opts *bind.TransactOpts, base *big.Int, perValidator *big.Int) (*types.Transaction, error) {
	return _OmniAVS.contract.Transact(opts, "setXcallGasLimits", base, perValidator)
}

// SetXcallGasLimits is a paid mutator transaction binding the contract method 0x718fef61.
//
// Solidity: function setXcallGasLimits(uint256 base, uint256 perValidator) returns()
func (_OmniAVS *OmniAVSSession) SetXcallGasLimits(base *big.Int, perValidator *big.Int) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetXcallGasLimits(&_OmniAVS.TransactOpts, base, perValidator)
}

// SetXcallGasLimits is a paid mutator transaction binding the contract method 0x718fef61.
//
// Solidity: function setXcallGasLimits(uint256 base, uint256 perValidator) returns()
func (_OmniAVS *OmniAVSTransactorSession) SetXcallGasLimits(base *big.Int, perValidator *big.Int) (*types.Transaction, error) {
	return _OmniAVS.Contract.SetXcallGasLimits(&_OmniAVS.TransactOpts, base, perValidator)
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
