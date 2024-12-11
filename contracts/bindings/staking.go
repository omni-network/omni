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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MinDelegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MinDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowValidators\",\"inputs\":[{\"name\":\"validators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createValidator\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createValidator\",\"inputs\":[{\"name\":\"x\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"disableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowValidators\",\"inputs\":[{\"name\":\"validators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValidatorPubkeyDigest\",\"inputs\":[{\"name\":\"x\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowlistEnabled_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV1\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowlistEnabled_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAllowedValidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowlistDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowlistEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"deposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delegate\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAllowed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDisallowed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611e22806100df6000396000f3fe6080604052600436106101145760003560e01c806384b0196e116100a0578063c6a2aac811610064578063c6a2aac814610306578063cf8e629a1461031b578063d146fd1b14610330578063eb4bd8441461034a578063f2fde38b1461035d57600080fd5b806384b0196e146102515780638ab65686146102795780638da5cb5b1461028c5780638f38fae8146102d3578063a5a470ad146102f357600080fd5b8063400ada75116100e7578063400ada75146101ab57806359bcddde146101cb5780635cd8a76b146101e7578063715018a6146101fc57806384768b7a1461021157600080fd5b8063117407e31461011957806311bcd8301461013b578063296192f41461016b5780633f0b1edf1461018b575b600080fd5b34801561012557600080fd5b50610139610134366004611817565b61037d565b005b34801561014757600080fd5b5061015868056bc75e2d6310000081565b6040519081526020015b60405180910390f35b34801561017757600080fd5b5061015861018636600461188c565b61044d565b34801561019757600080fd5b506101396101a6366004611817565b6104b3565b3480156101b757600080fd5b506101396101c63660046118ca565b61057f565b3480156101d757600080fd5b50610158670de0b6b3a764000081565b3480156101f357600080fd5b506101396106ce565b34801561020857600080fd5b506101396107d1565b34801561021d57600080fd5b5061024161022c366004611906565b60016020526000908152604090205460ff1681565b6040519015158152602001610162565b34801561025d57600080fd5b506102666107e5565b6040516101629796959493929190611967565b610139610287366004611a00565b610896565b34801561029857600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b039091168152602001610162565b3480156102df57600080fd5b506101396102ee3660046118ca565b610985565b610139610301366004611a7c565b610a35565b34801561031257600080fd5b50610139610b5b565b34801561032757600080fd5b50610139610b99565b34801561033c57600080fd5b506000546102419060ff1681565b610139610358366004611abe565b610bd4565b34801561036957600080fd5b50610139610378366004611906565b610d61565b610385610d9f565b60005b818110156104485760018060008585858181106103a7576103a7611b11565b90506020020160208101906103bc9190611906565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790558282828181106103f6576103f6611b11565b905060200201602081019061040b9190611906565b6001600160a01b03167fc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de9860405160405180910390a2600101610388565b505050565b604080517fc9a51567e61a6d1a243a60e57bf4560e7e543694b79349ce2cba3a14fe21b0426020820152908101839052606081018290526000906104aa906080015b60405160208183030381529060405280519060200120610dfa565b90505b92915050565b6104bb610d9f565b60005b81811015610448576000600160008585858181106104de576104de611b11565b90506020020160208101906104f39190611906565b6001600160a01b031681526020810191909152604001600020805460ff191691151591909117905582828281811061052d5761052d611b11565b90506020020160208101906105429190611906565b6001600160a01b03167f3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f60405160405180910390a26001016104be565b600080516020611dcd8339815191528054600160401b810460ff16159067ffffffffffffffff166000811580156105b35750825b905060008267ffffffffffffffff1660011480156105d05750303b155b9050811580156105de575080155b156105fc5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561062657845460ff60401b1916600160401b1785555b61062f87610e27565b610671604051806040016040528060078152602001665374616b696e6760c81b815250604051806040016040528060018152602001603160f81b815250610e38565b6000805460ff191687151517905583156106c557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b600080516020611dcd833981519152805460029190600160401b900460ff16806107065750805467ffffffffffffffff808416911610155b156107245760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff191667ffffffffffffffff831617600160401b17815560408051808201825260078152665374616b696e6760c81b602080830191909152825180840190935260018352603160f81b9083015261078591610e38565b805460ff60401b1916815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050565b6107d9610d9f565b6107e36000610e4e565b565b60006060808280808381600080516020611dad833981519152805490915015801561081257506001810154155b61085b5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b610863610ebf565b61086b610f82565b60408051600080825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b60005460ff1615806108c057506001600160a01b03811660009081526001602052604090205460ff165b61090c5760405162461bcd60e51b815260206004820152601860248201527f5374616b696e673a206e6f7420616c6c6f7765642076616c00000000000000006044820152606401610852565b670de0b6b3a76400003410156109345760405162461bcd60e51b815260040161085290611b3d565b806001600160a01b0316826001600160a01b03167f510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc3460405161097991815260200190565b60405180910390a35050565b600080516020611dcd8339815191528054600160401b810460ff16159067ffffffffffffffff166000811580156109b95750825b905060008267ffffffffffffffff1660011480156109d65750303b155b9050811580156109e4575080155b15610a025760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610a2c57845460ff60401b1916600160401b1785555b61067187610e27565b60005460ff161580610a5657503360009081526001602052604090205460ff165b610a995760405162461bcd60e51b815260206004820152601460248201527314dd185ada5b99ce881b9bdd08185b1b1bddd95960621b6044820152606401610852565b68056bc75e2d63100000341015610ac25760405162461bcd60e51b815260040161085290611b3d565b610acc8282610fc1565b610b125760405162461bcd60e51b81526020600482015260176024820152765374616b696e673a20696e76616c6964207075626b657960481b6044820152606401610852565b336001600160a01b03167fc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a453838334604051610b4f93929190611b74565b60405180910390a25050565b610b63610d9f565b6000805460ff191660011781556040517f8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec39190a1565b610ba1610d9f565b6000805460ff191681556040517f2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb079190a1565b60005460ff161580610bf557503360009081526001602052604090205460ff165b610c385760405162461bcd60e51b815260206004820152601460248201527314dd185ada5b99ce881b9bdd08185b1b1bddd95960621b6044820152606401610852565b68056bc75e2d63100000341015610c615760405162461bcd60e51b815260040161085290611b3d565b610c6b8484611115565b610cb15760405162461bcd60e51b81526020600482015260176024820152765374616b696e673a20696e76616c6964207075626b657960481b6044820152606401610852565b610cbd8484848461112b565b610d095760405162461bcd60e51b815260206004820152601a60248201527f5374616b696e673a20696e76616c6964207369676e61747572650000000000006044820152606401610852565b6000610d1585856111e3565b9050336001600160a01b03167fc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a4538234604051610d52929190611bad565b60405180910390a25050505050565b610d69610d9f565b6001600160a01b038116610d9357604051631e4fbdf760e01b815260006004820152602401610852565b610d9c81610e4e565b50565b33610dd17f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107e35760405163118cdaa760e01b8152336004820152602401610852565b60006104ad610e07611230565b8360405161190160f01b8152600281019290925260228201526042902090565b610e2f61123f565b610d9c81611276565b610e4061123f565b610e4a828261127e565b5050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1028054606091600080516020611dad83398151915291610efe90611bcf565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2a90611bcf565b8015610f775780601f10610f4c57610100808354040283529160200191610f77565b820191906000526020600020905b815481529060010190602001808311610f5a57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1038054606091600080516020611dad83398151915291610efe90611bcf565b6000602182146110135760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a20696e76616c6964207075626b6579206c656e67746800006044820152606401610852565b8282600081811061102657611026611b11565b9050013560f81c60f81b6001600160f81b031916600260f81b148061107457508282600081811061105957611059611b11565b9050013560f81c60f81b6001600160f81b031916600360f81b145b6110c05760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a20696e76616c6964207075626b65792070726566697800006044820152606401610852565b600183013560006110f5858583816110da576110da611b11565b919091013560f81c905083600060076401000003d0196112df565b905061110c8282600060076401000003d019611411565b95945050505050565b60006104aa83838360076401000003d019611411565b604080517fc9a51567e61a6d1a243a60e57bf4560e7e543694b79349ce2cba3a14fe21b04260208201529081018590526060810184905260009081906111739060800161048f565b905060006111b78286868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114ca92505050565b5050905060006111c78888611517565b6001600160a01b03928316921691909114979650505050505050565b606060006111f5600184166002611c1f565b6040805160f89290921b6001600160f81b03191660208301526021808301969096528051808303909601865260419091019052509192915050565b600061123a61154d565b905090565b600080516020611dcd83398151915254600160401b900460ff166107e357604051631afcd79f60e31b815260040160405180910390fd5b610d6961123f565b61128661123f565b600080516020611dad8339815191527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026112c08482611c88565b50600381016112cf8382611c88565b5060008082556001909101555050565b60008560ff16600214806112f657508560ff166003145b61135c5760405162461bcd60e51b815260206004820152603160248201527f456c6c697074696343757276653a696e6e76616c696420636f6d7072657373656044820152700c8408a8640e0ded2dce840e0e4caccd2f607b1b6064820152608401610852565b6000828061136c5761136c611d48565b838061137a5761137a611d48565b85858061138957611389611d48565b888a0908848061139b5761139b611d48565b85806113a9576113a9611d48565b898a0989090890506113d28160046113c2866001611d5e565b6113cc9190611d71565b856115c1565b9050600060026113e560ff8a1684611d5e565b6113ef9190611d85565b15611403576113fe8285611d99565b611405565b815b98975050505050505050565b60008515806114205750818610155b80611429575084155b806114345750818510155b156114415750600061110c565b6000828061145157611451611d48565b86870990506000838061146657611466611d48565b88858061147557611475611d48565b8a8b0909905085156114a557838061148f5761148f611d48565b848061149d5761149d611d48565b878a09820890505b84156114bf5783806114b9576114b9611d48565b85820890505b149695505050505050565b600080600083516041036115045760208401516040850151606086015160001a6114f68882858561169a565b955095509550505050611510565b50508151600091506002905b9250925092565b60408051818152606081018252600091829190602082018180368337505050602081019485526040810193909352505051902090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611578611769565b6115806117d3565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6000816000036116135760405162461bcd60e51b815260206004820152601e60248201527f456c6c697074696343757276653a206d6f64756c7573206973207a65726f00006044820152606401610852565b8360000361162357506000611693565b8260000361163357506001611693565b6001600160ff1b5b801561168f57838186161515870a85848509099150836002820486161515870a85848509099150836004820486161515870a85848509099150836008820486161515870a858485090991506010900461163b565b5090505b9392505050565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156116d5575060009150600390508261175f565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611729573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166117555750600092506001915082905061175f565b9250600091508190505b9450945094915050565b6000600080516020611dad83398151915281611783610ebf565b80519091501561179b57805160209091012092915050565b815480156117aa579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b6000600080516020611dad833981519152816117ed610f82565b80519091501561180557805160209091012092915050565b600182015480156117aa579392505050565b6000806020838503121561182a57600080fd5b823567ffffffffffffffff8082111561184257600080fd5b818501915085601f83011261185657600080fd5b81358181111561186557600080fd5b8660208260051b850101111561187a57600080fd5b60209290920196919550909350505050565b6000806040838503121561189f57600080fd5b50508035926020909101359150565b80356001600160a01b03811681146118c557600080fd5b919050565b600080604083850312156118dd57600080fd5b6118e6836118ae565b9150602083013580151581146118fb57600080fd5b809150509250929050565b60006020828403121561191857600080fd5b6104aa826118ae565b6000815180845260005b818110156119475760208185018101518683018201520161192b565b506000602082860101526020601f19601f83011685010191505092915050565b60ff60f81b881681526000602060e0602084015261198860e084018a611921565b838103604085015261199a818a611921565b606085018990526001600160a01b038816608086015260a0850187905284810360c08601528551808252602080880193509091019060005b818110156119ee578351835292840192918401916001016119d2565b50909c9b505050505050505050505050565b60008060408385031215611a1357600080fd5b611a1c836118ae565b9150611a2a602084016118ae565b90509250929050565b60008083601f840112611a4557600080fd5b50813567ffffffffffffffff811115611a5d57600080fd5b602083019150836020828501011115611a7557600080fd5b9250929050565b60008060208385031215611a8f57600080fd5b823567ffffffffffffffff811115611aa657600080fd5b611ab285828601611a33565b90969095509350505050565b60008060008060608587031215611ad457600080fd5b8435935060208501359250604085013567ffffffffffffffff811115611af957600080fd5b611b0587828801611a33565b95989497509550505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6020808252601d908201527f5374616b696e673a20696e73756666696369656e74206465706f736974000000604082015260600190565b604081528260408201528284606083013760006060848301015260006060601f19601f8601168301019050826020830152949350505050565b604081526000611bc06040830185611921565b90508260208301529392505050565b600181811c90821680611be357607f821691505b602082108103611c0357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60ff81811683821601908111156104ad576104ad611c09565b601f821115610448576000816000526020600020601f850160051c81016020861015611c615750805b601f850160051c820191505b81811015611c8057828155600101611c6d565b505050505050565b815167ffffffffffffffff811115611ca257611ca2611b27565b611cb681611cb08454611bcf565b84611c38565b602080601f831160018114611ceb5760008415611cd35750858301515b600019600386901b1c1916600185901b178555611c80565b600085815260208120601f198616915b82811015611d1a57888601518255948401946001909101908401611cfb565b5085821015611d385787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601260045260246000fd5b808201808211156104ad576104ad611c09565b600082611d8057611d80611d48565b500490565b600082611d9457611d94611d48565b500690565b818103818111156104ad576104ad611c0956fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212209ec1de4b1ad9f96937cdc63dd5d0c5e8f62af7eef29a7e85ce59896560b0f83264736f6c63430008180033",
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

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Staking *StakingCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Staking *StakingSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Staking.Contract.Eip712Domain(&_Staking.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Staking *StakingCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Staking.Contract.Eip712Domain(&_Staking.CallOpts)
}

// GetValidatorPubkeyDigest is a free data retrieval call binding the contract method 0x296192f4.
//
// Solidity: function getValidatorPubkeyDigest(bytes32 x, bytes32 y) view returns(bytes32)
func (_Staking *StakingCaller) GetValidatorPubkeyDigest(opts *bind.CallOpts, x [32]byte, y [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorPubkeyDigest", x, y)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetValidatorPubkeyDigest is a free data retrieval call binding the contract method 0x296192f4.
//
// Solidity: function getValidatorPubkeyDigest(bytes32 x, bytes32 y) view returns(bytes32)
func (_Staking *StakingSession) GetValidatorPubkeyDigest(x [32]byte, y [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetValidatorPubkeyDigest(&_Staking.CallOpts, x, y)
}

// GetValidatorPubkeyDigest is a free data retrieval call binding the contract method 0x296192f4.
//
// Solidity: function getValidatorPubkeyDigest(bytes32 x, bytes32 y) view returns(bytes32)
func (_Staking *StakingCallerSession) GetValidatorPubkeyDigest(x [32]byte, y [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetValidatorPubkeyDigest(&_Staking.CallOpts, x, y)
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

// CreateValidator0 is a paid mutator transaction binding the contract method 0xeb4bd844.
//
// Solidity: function createValidator(bytes32 x, bytes32 y, bytes signature) payable returns()
func (_Staking *StakingTransactor) CreateValidator0(opts *bind.TransactOpts, x [32]byte, y [32]byte, signature []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "createValidator0", x, y, signature)
}

// CreateValidator0 is a paid mutator transaction binding the contract method 0xeb4bd844.
//
// Solidity: function createValidator(bytes32 x, bytes32 y, bytes signature) payable returns()
func (_Staking *StakingSession) CreateValidator0(x [32]byte, y [32]byte, signature []byte) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator0(&_Staking.TransactOpts, x, y, signature)
}

// CreateValidator0 is a paid mutator transaction binding the contract method 0xeb4bd844.
//
// Solidity: function createValidator(bytes32 x, bytes32 y, bytes signature) payable returns()
func (_Staking *StakingTransactorSession) CreateValidator0(x [32]byte, y [32]byte, signature []byte) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator0(&_Staking.TransactOpts, x, y, signature)
}

// Delegate is a paid mutator transaction binding the contract method 0x8ab65686.
//
// Solidity: function delegate(address delegator, address validator) payable returns()
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", delegator, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x8ab65686.
//
// Solidity: function delegate(address delegator, address validator) payable returns()
func (_Staking *StakingSession) Delegate(delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, delegator, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x8ab65686.
//
// Solidity: function delegate(address delegator, address validator) payable returns()
func (_Staking *StakingTransactorSession) Delegate(delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, delegator, validator)
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

// InitializeV1 is a paid mutator transaction binding the contract method 0x8f38fae8.
//
// Solidity: function initializeV1(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingTransactor) InitializeV1(opts *bind.TransactOpts, owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initializeV1", owner_, isAllowlistEnabled_)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x8f38fae8.
//
// Solidity: function initializeV1(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingSession) InitializeV1(owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.Contract.InitializeV1(&_Staking.TransactOpts, owner_, isAllowlistEnabled_)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x8f38fae8.
//
// Solidity: function initializeV1(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingTransactorSession) InitializeV1(owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.Contract.InitializeV1(&_Staking.TransactOpts, owner_, isAllowlistEnabled_)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Staking *StakingTransactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Staking *StakingSession) InitializeV2() (*types.Transaction, error) {
	return _Staking.Contract.InitializeV2(&_Staking.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Staking *StakingTransactorSession) InitializeV2() (*types.Transaction, error) {
	return _Staking.Contract.InitializeV2(&_Staking.TransactOpts)
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

// StakingEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Staking contract.
type StakingEIP712DomainChangedIterator struct {
	Event *StakingEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *StakingEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingEIP712DomainChanged)
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
		it.Event = new(StakingEIP712DomainChanged)
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
func (it *StakingEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingEIP712DomainChanged represents a EIP712DomainChanged event raised by the Staking contract.
type StakingEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Staking *StakingFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*StakingEIP712DomainChangedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &StakingEIP712DomainChangedIterator{contract: _Staking.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Staking *StakingFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *StakingEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingEIP712DomainChanged)
				if err := _Staking.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Staking *StakingFilterer) ParseEIP712DomainChanged(log types.Log) (*StakingEIP712DomainChanged, error) {
	event := new(StakingEIP712DomainChanged)
	if err := _Staking.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
