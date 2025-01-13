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

// IERC7683FillInstruction is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type IERC7683FillInstruction struct {
// 	DestinationChainId uint64
// 	DestinationSettler [32]byte
// 	OriginData         []byte
// }

// IERC7683Output is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type IERC7683Output struct {
// 	Token     [32]byte
// 	Amount    *big.Int
// 	Recipient [32]byte
// 	ChainId   *big.Int
// }

// IERC7683ResolvedCrossChainOrder is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type IERC7683ResolvedCrossChainOrder struct {
// 	User             common.Address
// 	OriginChainId    *big.Int
// 	OpenDeadline     uint32
// 	FillDeadline     uint32
// 	OrderId          [32]byte
// 	MaxSpent         []IERC7683Output
// 	MinReceived      []IERC7683Output
// 	FillInstructions []IERC7683FillInstruction
// }

// SolverNetOutboxMetaData contains all meta data concerning the SolverNetOutbox contract.
var SolverNetOutboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"didFill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fillFee\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Filled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filledBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Open\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"resolvedOrder\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenses\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongDestChain\",\"inputs\":[]}]",
	Bin: "0x60a06040523480156200001157600080fd5b5063ffffffff60643b1615620000975760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156200007f575060408051601f3d908101601f191682019092526200007c9181019062000110565b60015b6200008e57436080526200009c565b6080526200009c565b436080525b620000a6620000ac565b6200012a565b63409feecd1980546001811615620000cc5763f92ee8a96000526004601cfd5b8160c01c808260011c146200010b578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b6000602082840312156200012357600080fd5b5051919050565b608051611f0362000146600039600061030d0152611f036000f3fe60806040526004361061011f5760003560e01c806354d1f13d116100a0578063eae4c19f11610064578063eae4c19f146102fb578063f04e283e1461032f578063f2fde38b14610342578063f8c8765e14610355578063fee81cf41461037557600080fd5b806354d1f13d1461028c578063715018a61461029457806374eeb8471461029c57806382e2c43f146102cf5780638da5cb5b146102e257600080fd5b806325692962116100e757806325692962146101cf5780632de94807146101d757806339acf9f11461020a5780634a4ee7b114610242578063514e62fc1461025557600080fd5b8063183a4f6e146101245780631c10893f146101395780631cd64df41461014c57806320dcd41614610181578063248689cc146101af575b600080fd5b610137610132366004611073565b6103a8565b005b6101376101473660046110a8565b6103b5565b34801561015857600080fd5b5061016c6101673660046110a8565b6103cb565b60405190151581526020015b60405180910390f35b34801561018d57600080fd5b506101a161019c3660046110e9565b6103ea565b604051908152602001610178565b3480156101bb57600080fd5b5061016c6101ca366004611153565b61043b565b61013761049d565b3480156101e357600080fd5b506101a16101f236600461119e565b638b78c6d8600c908152600091909152602090205490565b34801561021657600080fd5b5060005461022a906001600160a01b031681565b6040516001600160a01b039091168152602001610178565b6101376102503660046110a8565b6104ec565b34801561026157600080fd5b5061016c6102703660046110a8565b638b78c6d8600c90815260009290925260209091205416151590565b6101376104fe565b61013761053a565b3480156102a857600080fd5b506000546102bd90600160a01b900460ff1681565b60405160ff9091168152602001610178565b6101376102dd3660046111b9565b61054e565b3480156102ee57600080fd5b50638b78c6d8195461022a565b34801561030757600080fd5b506101a17f000000000000000000000000000000000000000000000000000000000000000081565b61013761033d36600461119e565b610608565b61013761035036600461119e565b610645565b34801561036157600080fd5b50610137610370366004611232565b61066c565b34801561038157600080fd5b506101a161039036600461119e565b63389a75e1600c908152600091909152602090205490565b6103b23382610773565b50565b6103bd61077f565b6103c7828261079a565b5050565b638b78c6d8600c90815260008390526020902054811681145b92915050565b6040516000196024820181905260448201526000906103e490839060640160408051601f198184030181529190526020810180516001600160e01b0316637b0f383b60e01b179052620186a06107a6565b6000600460006104818686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061082492505050565b815260208101919091526040016000205460ff16949350505050565b60006202a3006001600160401b03164201905063389a75e1600c5233600052806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d600080a250565b6104f461077f565b6103c78282610773565b63389a75e1600c523360005260006020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92600080a2565b61054261077f565b61054c6000610857565b565b600161055981610895565b3068929eee149b4bd2126854036105785763ab143c066000526004601cfd5b3068929eee149b4bd21268556000610592858701876113e5565b60208101519091506105a3816108bb565b6105f2888360000151836105ed8c8c8c8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061082492505050565b610c08565b50503868929eee149b4bd2126855505050505050565b61061061077f565b63389a75e1600c52806000526020600c20805442111561063857636f5e88186000526004601cfd5b600090556103b281610857565b61064d61077f565b8060601b61066357637448fbae6000526004601cfd5b6103b281610857565b63409feecd1980546003825580156106a35760018160011c14303b1061069a5763f92ee8a96000526004601cfd5b818160ff1b1b91505b506106ad85610d48565b6106b884600161079a565b6106c183610d84565b600280546001600160a01b0319166001600160a01b03841617905560405130906106ea90611066565b6001600160a01b039091168152602001604051809103906000f080158015610716573d6000803e3d6000fd5b50600380546001600160a01b0319166001600160a01b0392909216919091179055801561076c576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b5050505050565b6103c782826000610e28565b638b78c6d81954331461054c576382b429006000526004601cfd5b6103c782826001610e28565b60008054604051632376548f60e21b81526001600160a01b0390911690638dd9523c906107db90879087908790600401611570565b602060405180830381865afa1580156107f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081c91906115a6565b949350505050565b600082826040516020016108399291906115bf565b60405160208183030381529060405280519060200120905092915050565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a355565b638b78c6d8600c5233600052806020600c2054166103b2576382b429006000526004601cfd5b806080015160005b81518110156109ac5760008282815181106108e0576108e06115d8565b6020026020010151905060006108f7826000015190565b90506000610906836020015190565b600354604085015191925061092b916001600160a01b03858116923392911690610e81565b600354604084810151905163e1f21c6760e01b81526001600160a01b0385811660048301528481166024830152604482019290925291169063e1f21c6790606401600060405180830381600087803b15801561098657600080fd5b505af115801561099a573d6000803e3d6000fd5b505050505050508060010190506108c3565b5081516001600160401b031646146109d75760405163fd24301760e01b815260040160405180910390fd5b6003546040808401519051635b8eb34960e11b81526001600160a01b039092169163b71d66929190610a0d9086906004016115ee565b6000604051808303818588803b158015610a2657600080fd5b505af1158015610a3a573d6000803e3d6000fd5b505050505060005b8151811015610b8a576000828281518110610a5f57610a5f6115d8565b602002602001015190506000610a76826000015190565b600354909150600090610a95906001600160a01b038085169116610edf565b90508015610b7c57600354602084015160405163e1f21c6760e01b81526001600160a01b03858116600483015291821660248201526000604482015291169063e1f21c6790606401600060405180830381600087803b158015610af757600080fd5b505af1158015610b0b573d6000803e3d6000fd5b50506003546040516317d5759960e31b81526001600160a01b03868116600483015233602483015260448201869052909116925063beabacc89150606401600060405180830381600087803b158015610b6357600080fd5b505af1158015610b77573d6000803e3d6000fd5b505050505b505050806001019050610a42565b506003546001600160a01b0316318015610c0357600354604051633e97486160e11b8152336004820152602481018390526001600160a01b0390911690637d2e90c290604401600060405180830381600087803b158015610bea57600080fd5b505af1158015610bfe573d6000803e3d6000fd5b505050505b505050565b60008181526004602052604090205460ff1615610c38576040516341a26a6360e01b815260040160405180910390fd5b6000818152600460208190526040808320805460ff1916600117905560025490516024810188905260448101859052610cb192879290916001600160a01b039091169060640160408051601f198184030181529190526020810180516001600160e01b0316637b0f383b60e01b179052620186a0610f0b565b905080836040015134610cc4919061169a565b1015610ce25760405162976f7560e21b815260040160405180910390fd5b600081846040015134610cf5919061169a565b610cff919061169a565b90508015610d1157610d11338261104a565b6040513390849088907fa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc90600090a4505050505050565b6001600160a01b0316638b78c6d8198190558060007f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b038116610dd45760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b60448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52826000526020600c20805483811783610e4a575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26600080a3505050505050565b60405181606052826040528360601b602c526323b872dd60601b600c52602060006064601c6000895af18060016000511416610ed057803d873b151710610ed057637939f4246000526004601cfd5b50600060605260405250505050565b6000816014526370a0823160601b60005260208060246010865afa601f3d111660205102905092915050565b60008054604051632376548f60e21b815282916001600160a01b031690638dd9523c90610f40908a9088908890600401611570565b602060405180830381865afa158015610f5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8191906115a6565b905080471015610fd35760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401610dcb565b60005460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f90839061100d908b908b908b908b908b906004016116bb565b6000604051808303818588803b15801561102657600080fd5b505af115801561103a573d6000803e3d6000fd5b50939a9950505050505050505050565b60003860003884865af16103c75763b12d13eb6000526004601cfd5b6107c38061170b83390190565b60006020828403121561108557600080fd5b5035919050565b80356001600160a01b03811681146110a357600080fd5b919050565b600080604083850312156110bb57600080fd5b6110c48361108c565b946020939093013593505050565b80356001600160401b03811681146110a357600080fd5b6000602082840312156110fb57600080fd5b611104826110d2565b9392505050565b60008083601f84011261111d57600080fd5b5081356001600160401b0381111561113457600080fd5b60208301915083602082850101111561114c57600080fd5b9250929050565b60008060006040848603121561116857600080fd5b8335925060208401356001600160401b0381111561118557600080fd5b6111918682870161110b565b9497909650939450505050565b6000602082840312156111b057600080fd5b6111048261108c565b6000806000806000606086880312156111d157600080fd5b8535945060208601356001600160401b03808211156111ef57600080fd5b6111fb89838a0161110b565b9096509450604088013591508082111561121457600080fd5b506112218882890161110b565b969995985093965092949392505050565b6000806000806080858703121561124857600080fd5b6112518561108c565b935061125f6020860161108c565b925061126d6040860161108c565b915061127b6060860161108c565b905092959194509250565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156112be576112be611286565b60405290565b604080519081016001600160401b03811182821017156112be576112be611286565b60405160a081016001600160401b03811182821017156112be576112be611286565b604051601f8201601f191681016001600160401b038111828210171561133057611330611286565b604052919050565b600082601f83011261134957600080fd5b813560206001600160401b0382111561136457611364611286565b611372818360051b01611308565b8281526060928302850182019282820191908785111561139157600080fd5b8387015b858110156113d85781818a0312156113ad5760008081fd5b6113b561129c565b813581528582013586820152604080830135908201528452928401928101611395565b5090979650505050505050565b600060208083850312156113f857600080fd5b82356001600160401b038082111561140f57600080fd5b908401906040828703121561142357600080fd5b61142b6112c4565b611434836110d2565b8152838301358281111561144757600080fd5b929092019160a0838803121561145c57600080fd5b6114646112e6565b61146d846110d2565b815284840135858201526040840135604082015260608401358381111561149357600080fd5b8401601f810189136114a457600080fd5b8035848111156114b6576114b6611286565b6114c8601f8201601f19168801611308565b8181528a888385010111156114dc57600080fd5b8188840189830137600088838301015280606085015250505060808401358381111561150757600080fd5b61151389828701611338565b608083015250938101939093525090949350505050565b6000815180845260005b8181101561155057602081850181015186830182015201611534565b506000602082860101526020601f19601f83011685010191505092915050565b60006001600160401b03808616835260606020840152611593606084018661152a565b9150808416604084015250949350505050565b6000602082840312156115b857600080fd5b5051919050565b82815260406020820152600061081c604083018461152a565b634e487b7160e01b600052603260045260246000fd5b600060208083526001600160401b03845116818401528084015160408160408601526040860151915060608260608701526060870151925060a0608087015261163a60c087018461152a565b6080880151878203601f190160a0890152805180835290860194506000918601905b8083101561168d5785518051835287810151888401528501518583015294860194600192909201919083019061165c565b5098975050505050505050565b818103818111156103e457634e487b7160e01b600052601160045260246000fd5b60006001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a060608401526116f560a084018661152a565b9150808416608084015250969550505050505056fe60a060405234801561001057600080fd5b506040516107c33803806107c383398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b60805161071d6100a66000396000818160b901528181610122015281816101820152818161026601526102c3015261071d6000f3fe60806040526004361061004b5760003560e01c80637d2e90c214610054578063b71d669214610074578063beabacc814610087578063ce11e6ab146100a7578063e1f21c67146100f757005b3661005257005b005b34801561006057600080fd5b5061005261006f3660046103e2565b610117565b6100526100823660046105bd565b610177565b34801561009357600080fd5b506100526100a236600461067c565b61025b565b3480156100b357600080fd5b506100db7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b34801561010357600080fd5b5061005261011236600461067c565b6102b8565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101605760405163bda8fc9560e01b815260040160405180910390fd5b6101736001600160a01b03831682610315565b5050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101c05760405163bda8fc9560e01b815260040160405180910390fd5b60006101cd826020015190565b90506000816001600160a01b0316836040015184606001516040516101f291906106b8565b60006040518083038185875af1925050503d806000811461022f576040519150601f19603f3d011682016040523d82523d6000602084013e610234565b606091505b505090508061025657604051633204506f60e01b815260040160405180910390fd5b505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146102a45760405163bda8fc9560e01b815260040160405180910390fd5b6102566001600160a01b0384168383610331565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103015760405163bda8fc9560e01b815260040160405180910390fd5b6102566001600160a01b0384168383610381565b60003860003884865af16101735763b12d13eb6000526004601cfd5b816014528060345263a9059cbb60601b60005260206000604460106000875af1806001600051141661037657803d853b151710610376576390b8ec186000526004601cfd5b506000603452505050565b816014528060345263095ea7b360601b60005260206000604460106000875af1806001600051141661037657803d853b15171061037657633e3f8f736000526004601cfd5b80356001600160a01b03811681146103dd57600080fd5b919050565b600080604083850312156103f557600080fd5b6103fe836103c6565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156104455761044561040c565b60405290565b60405160a0810167ffffffffffffffff811182821017156104455761044561040c565b604051601f8201601f1916810167ffffffffffffffff811182821017156104975761049761040c565b604052919050565b600082601f8301126104b057600080fd5b813567ffffffffffffffff8111156104ca576104ca61040c565b6104dd601f8201601f191660200161046e565b8181528460208386010111156104f257600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261052057600080fd5b8135602067ffffffffffffffff82111561053c5761053c61040c565b61054a818360051b0161046e565b8281526060928302850182019282820191908785111561056957600080fd5b8387015b858110156105b05781818a0312156105855760008081fd5b61058d610422565b81358152858201358682015260408083013590820152845292840192810161056d565b5090979650505050505050565b6000602082840312156105cf57600080fd5b813567ffffffffffffffff808211156105e757600080fd5b9083019060a082860312156105fb57600080fd5b61060361044b565b8235828116811461061357600080fd5b80825250602083013560208201526040830135604082015260608301358281111561063d57600080fd5b6106498782860161049f565b60608301525060808301358281111561066157600080fd5b61066d8782860161050f565b60808301525095945050505050565b60008060006060848603121561069157600080fd5b61069a846103c6565b92506106a8602085016103c6565b9150604084013590509250925092565b6000825160005b818110156106d957602081860181015185830152016106bf565b50600092019182525091905056fea264697066735822122039845f33fb91200ca5d4e2630ebbf90440a5eddc4956e363513a6ec69679c08564736f6c63430008180033a264697066735822122009087387d017510769b8afc67372773f7e759a8e7493030f6e0c80c0df83335564736f6c63430008180033",
}

// SolverNetOutboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SolverNetOutboxMetaData.ABI instead.
var SolverNetOutboxABI = SolverNetOutboxMetaData.ABI

// SolverNetOutboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolverNetOutboxMetaData.Bin instead.
var SolverNetOutboxBin = SolverNetOutboxMetaData.Bin

// DeploySolverNetOutbox deploys a new Ethereum contract, binding an instance of SolverNetOutbox to it.
func DeploySolverNetOutbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolverNetOutbox, error) {
	parsed, err := SolverNetOutboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolverNetOutboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolverNetOutbox{SolverNetOutboxCaller: SolverNetOutboxCaller{contract: contract}, SolverNetOutboxTransactor: SolverNetOutboxTransactor{contract: contract}, SolverNetOutboxFilterer: SolverNetOutboxFilterer{contract: contract}}, nil
}

// SolverNetOutbox is an auto generated Go binding around an Ethereum contract.
type SolverNetOutbox struct {
	SolverNetOutboxCaller     // Read-only binding to the contract
	SolverNetOutboxTransactor // Write-only binding to the contract
	SolverNetOutboxFilterer   // Log filterer for contract events
}

// SolverNetOutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolverNetOutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetOutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolverNetOutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetOutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolverNetOutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetOutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolverNetOutboxSession struct {
	Contract     *SolverNetOutbox  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolverNetOutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolverNetOutboxCallerSession struct {
	Contract *SolverNetOutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SolverNetOutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolverNetOutboxTransactorSession struct {
	Contract     *SolverNetOutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SolverNetOutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolverNetOutboxRaw struct {
	Contract *SolverNetOutbox // Generic contract binding to access the raw methods on
}

// SolverNetOutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolverNetOutboxCallerRaw struct {
	Contract *SolverNetOutboxCaller // Generic read-only contract binding to access the raw methods on
}

// SolverNetOutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolverNetOutboxTransactorRaw struct {
	Contract *SolverNetOutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolverNetOutbox creates a new instance of SolverNetOutbox, bound to a specific deployed contract.
func NewSolverNetOutbox(address common.Address, backend bind.ContractBackend) (*SolverNetOutbox, error) {
	contract, err := bindSolverNetOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutbox{SolverNetOutboxCaller: SolverNetOutboxCaller{contract: contract}, SolverNetOutboxTransactor: SolverNetOutboxTransactor{contract: contract}, SolverNetOutboxFilterer: SolverNetOutboxFilterer{contract: contract}}, nil
}

// NewSolverNetOutboxCaller creates a new read-only instance of SolverNetOutbox, bound to a specific deployed contract.
func NewSolverNetOutboxCaller(address common.Address, caller bind.ContractCaller) (*SolverNetOutboxCaller, error) {
	contract, err := bindSolverNetOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxCaller{contract: contract}, nil
}

// NewSolverNetOutboxTransactor creates a new write-only instance of SolverNetOutbox, bound to a specific deployed contract.
func NewSolverNetOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SolverNetOutboxTransactor, error) {
	contract, err := bindSolverNetOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxTransactor{contract: contract}, nil
}

// NewSolverNetOutboxFilterer creates a new log filterer instance of SolverNetOutbox, bound to a specific deployed contract.
func NewSolverNetOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SolverNetOutboxFilterer, error) {
	contract, err := bindSolverNetOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxFilterer{contract: contract}, nil
}

// bindSolverNetOutbox binds a generic wrapper to an already deployed contract.
func bindSolverNetOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolverNetOutboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetOutbox *SolverNetOutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetOutbox.Contract.SolverNetOutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetOutbox *SolverNetOutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.SolverNetOutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetOutbox *SolverNetOutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.SolverNetOutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetOutbox *SolverNetOutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetOutbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetOutbox *SolverNetOutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetOutbox *SolverNetOutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.contract.Transact(opts, method, params...)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolverNetOutbox *SolverNetOutboxCaller) DefaultConfLevel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "defaultConfLevel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolverNetOutbox *SolverNetOutboxSession) DefaultConfLevel() (uint8, error) {
	return _SolverNetOutbox.Contract.DefaultConfLevel(&_SolverNetOutbox.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) DefaultConfLevel() (uint8, error) {
	return _SolverNetOutbox.Contract.DefaultConfLevel(&_SolverNetOutbox.CallOpts)
}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxCaller) DeployedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "deployedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxSession) DeployedAt() (*big.Int, error) {
	return _SolverNetOutbox.Contract.DeployedAt(&_SolverNetOutbox.CallOpts)
}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) DeployedAt() (*big.Int, error) {
	return _SolverNetOutbox.Contract.DeployedAt(&_SolverNetOutbox.CallOpts)
}

// DidFill is a free data retrieval call binding the contract method 0x248689cc.
//
// Solidity: function didFill(bytes32 orderId, bytes originData) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxCaller) DidFill(opts *bind.CallOpts, orderId [32]byte, originData []byte) (bool, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "didFill", orderId, originData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidFill is a free data retrieval call binding the contract method 0x248689cc.
//
// Solidity: function didFill(bytes32 orderId, bytes originData) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxSession) DidFill(orderId [32]byte, originData []byte) (bool, error) {
	return _SolverNetOutbox.Contract.DidFill(&_SolverNetOutbox.CallOpts, orderId, originData)
}

// DidFill is a free data retrieval call binding the contract method 0x248689cc.
//
// Solidity: function didFill(bytes32 orderId, bytes originData) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) DidFill(orderId [32]byte, originData []byte) (bool, error) {
	return _SolverNetOutbox.Contract.DidFill(&_SolverNetOutbox.CallOpts, orderId, originData)
}

// FillFee is a free data retrieval call binding the contract method 0x20dcd416.
//
// Solidity: function fillFee(uint64 srcChainId) view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxCaller) FillFee(opts *bind.CallOpts, srcChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "fillFee", srcChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillFee is a free data retrieval call binding the contract method 0x20dcd416.
//
// Solidity: function fillFee(uint64 srcChainId) view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxSession) FillFee(srcChainId uint64) (*big.Int, error) {
	return _SolverNetOutbox.Contract.FillFee(&_SolverNetOutbox.CallOpts, srcChainId)
}

// FillFee is a free data retrieval call binding the contract method 0x20dcd416.
//
// Solidity: function fillFee(uint64 srcChainId) view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) FillFee(srcChainId uint64) (*big.Int, error) {
	return _SolverNetOutbox.Contract.FillFee(&_SolverNetOutbox.CallOpts, srcChainId)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _SolverNetOutbox.Contract.HasAllRoles(&_SolverNetOutbox.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _SolverNetOutbox.Contract.HasAllRoles(&_SolverNetOutbox.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _SolverNetOutbox.Contract.HasAnyRole(&_SolverNetOutbox.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _SolverNetOutbox.Contract.HasAnyRole(&_SolverNetOutbox.CallOpts, user, roles)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxSession) Omni() (common.Address, error) {
	return _SolverNetOutbox.Contract.Omni(&_SolverNetOutbox.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) Omni() (common.Address, error) {
	return _SolverNetOutbox.Contract.Omni(&_SolverNetOutbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolverNetOutbox *SolverNetOutboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolverNetOutbox *SolverNetOutboxSession) Owner() (common.Address, error) {
	return _SolverNetOutbox.Contract.Owner(&_SolverNetOutbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) Owner() (common.Address, error) {
	return _SolverNetOutbox.Contract.Owner(&_SolverNetOutbox.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolverNetOutbox *SolverNetOutboxCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolverNetOutbox *SolverNetOutboxSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _SolverNetOutbox.Contract.OwnershipHandoverExpiresAt(&_SolverNetOutbox.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _SolverNetOutbox.Contract.OwnershipHandoverExpiresAt(&_SolverNetOutbox.CallOpts, pendingOwner)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolverNetOutbox *SolverNetOutboxCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolverNetOutbox *SolverNetOutboxSession) RolesOf(user common.Address) (*big.Int, error) {
	return _SolverNetOutbox.Contract.RolesOf(&_SolverNetOutbox.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _SolverNetOutbox.Contract.RolesOf(&_SolverNetOutbox.CallOpts, user)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.CancelOwnershipHandover(&_SolverNetOutbox.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.CancelOwnershipHandover(&_SolverNetOutbox.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.CompleteOwnershipHandover(&_SolverNetOutbox.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.CompleteOwnershipHandover(&_SolverNetOutbox.TransactOpts, pendingOwner)
}

// Fill is a paid mutator transaction binding the contract method 0x82e2c43f.
//
// Solidity: function fill(bytes32 orderId, bytes originData, bytes ) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) Fill(opts *bind.TransactOpts, orderId [32]byte, originData []byte, arg2 []byte) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "fill", orderId, originData, arg2)
}

// Fill is a paid mutator transaction binding the contract method 0x82e2c43f.
//
// Solidity: function fill(bytes32 orderId, bytes originData, bytes ) payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) Fill(orderId [32]byte, originData []byte, arg2 []byte) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Fill(&_SolverNetOutbox.TransactOpts, orderId, originData, arg2)
}

// Fill is a paid mutator transaction binding the contract method 0x82e2c43f.
//
// Solidity: function fill(bytes32 orderId, bytes originData, bytes ) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) Fill(orderId [32]byte, originData []byte, arg2 []byte) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Fill(&_SolverNetOutbox.TransactOpts, orderId, originData, arg2)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.GrantRoles(&_SolverNetOutbox.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.GrantRoles(&_SolverNetOutbox.TransactOpts, user, roles)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address inbox_) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, solver_ common.Address, omni_ common.Address, inbox_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "initialize", owner_, solver_, omni_, inbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address inbox_) returns()
func (_SolverNetOutbox *SolverNetOutboxSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address, inbox_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Initialize(&_SolverNetOutbox.TransactOpts, owner_, solver_, omni_, inbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address inbox_) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address, inbox_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Initialize(&_SolverNetOutbox.TransactOpts, owner_, solver_, omni_, inbox_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.RenounceOwnership(&_SolverNetOutbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.RenounceOwnership(&_SolverNetOutbox.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.RenounceRoles(&_SolverNetOutbox.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.RenounceRoles(&_SolverNetOutbox.TransactOpts, roles)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.RequestOwnershipHandover(&_SolverNetOutbox.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.RequestOwnershipHandover(&_SolverNetOutbox.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.RevokeRoles(&_SolverNetOutbox.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.RevokeRoles(&_SolverNetOutbox.TransactOpts, user, roles)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.TransferOwnership(&_SolverNetOutbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.TransferOwnership(&_SolverNetOutbox.TransactOpts, newOwner)
}

// SolverNetOutboxDefaultConfLevelSetIterator is returned from FilterDefaultConfLevelSet and is used to iterate over the raw logs and unpacked data for DefaultConfLevelSet events raised by the SolverNetOutbox contract.
type SolverNetOutboxDefaultConfLevelSetIterator struct {
	Event *SolverNetOutboxDefaultConfLevelSet // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxDefaultConfLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxDefaultConfLevelSet)
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
		it.Event = new(SolverNetOutboxDefaultConfLevelSet)
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
func (it *SolverNetOutboxDefaultConfLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxDefaultConfLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxDefaultConfLevelSet represents a DefaultConfLevelSet event raised by the SolverNetOutbox contract.
type SolverNetOutboxDefaultConfLevelSet struct {
	Conf uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultConfLevelSet is a free log retrieval operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterDefaultConfLevelSet(opts *bind.FilterOpts) (*SolverNetOutboxDefaultConfLevelSetIterator, error) {

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxDefaultConfLevelSetIterator{contract: _SolverNetOutbox.contract, event: "DefaultConfLevelSet", logs: logs, sub: sub}, nil
}

// WatchDefaultConfLevelSet is a free log subscription operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchDefaultConfLevelSet(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxDefaultConfLevelSet) (event.Subscription, error) {

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxDefaultConfLevelSet)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
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
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseDefaultConfLevelSet(log types.Log) (*SolverNetOutboxDefaultConfLevelSet, error) {
	event := new(SolverNetOutboxDefaultConfLevelSet)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetOutboxFilledIterator is returned from FilterFilled and is used to iterate over the raw logs and unpacked data for Filled events raised by the SolverNetOutbox contract.
type SolverNetOutboxFilledIterator struct {
	Event *SolverNetOutboxFilled // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxFilled)
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
		it.Event = new(SolverNetOutboxFilled)
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
func (it *SolverNetOutboxFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxFilled represents a Filled event raised by the SolverNetOutbox contract.
type SolverNetOutboxFilled struct {
	OrderId  [32]byte
	FillHash [32]byte
	FilledBy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFilled is a free log retrieval operation binding the contract event 0xa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc.
//
// Solidity: event Filled(bytes32 indexed orderId, bytes32 indexed fillHash, address indexed filledBy)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterFilled(opts *bind.FilterOpts, orderId [][32]byte, fillHash [][32]byte, filledBy []common.Address) (*SolverNetOutboxFilledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var fillHashRule []interface{}
	for _, fillHashItem := range fillHash {
		fillHashRule = append(fillHashRule, fillHashItem)
	}
	var filledByRule []interface{}
	for _, filledByItem := range filledBy {
		filledByRule = append(filledByRule, filledByItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "Filled", orderIdRule, fillHashRule, filledByRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxFilledIterator{contract: _SolverNetOutbox.contract, event: "Filled", logs: logs, sub: sub}, nil
}

// WatchFilled is a free log subscription operation binding the contract event 0xa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc.
//
// Solidity: event Filled(bytes32 indexed orderId, bytes32 indexed fillHash, address indexed filledBy)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchFilled(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxFilled, orderId [][32]byte, fillHash [][32]byte, filledBy []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var fillHashRule []interface{}
	for _, fillHashItem := range fillHash {
		fillHashRule = append(fillHashRule, fillHashItem)
	}
	var filledByRule []interface{}
	for _, filledByItem := range filledBy {
		filledByRule = append(filledByRule, filledByItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "Filled", orderIdRule, fillHashRule, filledByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxFilled)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "Filled", log); err != nil {
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

// ParseFilled is a log parse operation binding the contract event 0xa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc.
//
// Solidity: event Filled(bytes32 indexed orderId, bytes32 indexed fillHash, address indexed filledBy)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseFilled(log types.Log) (*SolverNetOutboxFilled, error) {
	event := new(SolverNetOutboxFilled)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "Filled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetOutboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SolverNetOutbox contract.
type SolverNetOutboxInitializedIterator struct {
	Event *SolverNetOutboxInitialized // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxInitialized)
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
		it.Event = new(SolverNetOutboxInitialized)
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
func (it *SolverNetOutboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxInitialized represents a Initialized event raised by the SolverNetOutbox contract.
type SolverNetOutboxInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*SolverNetOutboxInitializedIterator, error) {

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxInitializedIterator{contract: _SolverNetOutbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxInitialized) (event.Subscription, error) {

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxInitialized)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseInitialized(log types.Log) (*SolverNetOutboxInitialized, error) {
	event := new(SolverNetOutboxInitialized)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetOutboxOmniPortalSetIterator is returned from FilterOmniPortalSet and is used to iterate over the raw logs and unpacked data for OmniPortalSet events raised by the SolverNetOutbox contract.
type SolverNetOutboxOmniPortalSetIterator struct {
	Event *SolverNetOutboxOmniPortalSet // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxOmniPortalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxOmniPortalSet)
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
		it.Event = new(SolverNetOutboxOmniPortalSet)
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
func (it *SolverNetOutboxOmniPortalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxOmniPortalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxOmniPortalSet represents a OmniPortalSet event raised by the SolverNetOutbox contract.
type SolverNetOutboxOmniPortalSet struct {
	Omni common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOmniPortalSet is a free log retrieval operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterOmniPortalSet(opts *bind.FilterOpts) (*SolverNetOutboxOmniPortalSetIterator, error) {

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxOmniPortalSetIterator{contract: _SolverNetOutbox.contract, event: "OmniPortalSet", logs: logs, sub: sub}, nil
}

// WatchOmniPortalSet is a free log subscription operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchOmniPortalSet(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxOmniPortalSet) (event.Subscription, error) {

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxOmniPortalSet)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
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
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseOmniPortalSet(log types.Log) (*SolverNetOutboxOmniPortalSet, error) {
	event := new(SolverNetOutboxOmniPortalSet)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetOutboxOpenIterator is returned from FilterOpen and is used to iterate over the raw logs and unpacked data for Open events raised by the SolverNetOutbox contract.
type SolverNetOutboxOpenIterator struct {
	Event *SolverNetOutboxOpen // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxOpen)
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
		it.Event = new(SolverNetOutboxOpen)
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
func (it *SolverNetOutboxOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxOpen represents a Open event raised by the SolverNetOutbox contract.
type SolverNetOutboxOpen struct {
	OrderId       [32]byte
	ResolvedOrder IERC7683ResolvedCrossChainOrder
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpen is a free log retrieval operation binding the contract event 0xa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c.
//
// Solidity: event Open(bytes32 indexed orderId, (address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolvedOrder)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterOpen(opts *bind.FilterOpts, orderId [][32]byte) (*SolverNetOutboxOpenIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "Open", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxOpenIterator{contract: _SolverNetOutbox.contract, event: "Open", logs: logs, sub: sub}, nil
}

// WatchOpen is a free log subscription operation binding the contract event 0xa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c.
//
// Solidity: event Open(bytes32 indexed orderId, (address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolvedOrder)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchOpen(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxOpen, orderId [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "Open", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxOpen)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "Open", log); err != nil {
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

// ParseOpen is a log parse operation binding the contract event 0xa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c.
//
// Solidity: event Open(bytes32 indexed orderId, (address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolvedOrder)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseOpen(log types.Log) (*SolverNetOutboxOpen, error) {
	event := new(SolverNetOutboxOpen)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "Open", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetOutboxOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the SolverNetOutbox contract.
type SolverNetOutboxOwnershipHandoverCanceledIterator struct {
	Event *SolverNetOutboxOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxOwnershipHandoverCanceled)
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
		it.Event = new(SolverNetOutboxOwnershipHandoverCanceled)
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
func (it *SolverNetOutboxOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the SolverNetOutbox contract.
type SolverNetOutboxOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*SolverNetOutboxOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxOwnershipHandoverCanceledIterator{contract: _SolverNetOutbox.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxOwnershipHandoverCanceled)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*SolverNetOutboxOwnershipHandoverCanceled, error) {
	event := new(SolverNetOutboxOwnershipHandoverCanceled)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetOutboxOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the SolverNetOutbox contract.
type SolverNetOutboxOwnershipHandoverRequestedIterator struct {
	Event *SolverNetOutboxOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxOwnershipHandoverRequested)
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
		it.Event = new(SolverNetOutboxOwnershipHandoverRequested)
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
func (it *SolverNetOutboxOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the SolverNetOutbox contract.
type SolverNetOutboxOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*SolverNetOutboxOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxOwnershipHandoverRequestedIterator{contract: _SolverNetOutbox.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxOwnershipHandoverRequested)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseOwnershipHandoverRequested(log types.Log) (*SolverNetOutboxOwnershipHandoverRequested, error) {
	event := new(SolverNetOutboxOwnershipHandoverRequested)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetOutboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SolverNetOutbox contract.
type SolverNetOutboxOwnershipTransferredIterator struct {
	Event *SolverNetOutboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxOwnershipTransferred)
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
		it.Event = new(SolverNetOutboxOwnershipTransferred)
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
func (it *SolverNetOutboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxOwnershipTransferred represents a OwnershipTransferred event raised by the SolverNetOutbox contract.
type SolverNetOutboxOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*SolverNetOutboxOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxOwnershipTransferredIterator{contract: _SolverNetOutbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxOwnershipTransferred)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseOwnershipTransferred(log types.Log) (*SolverNetOutboxOwnershipTransferred, error) {
	event := new(SolverNetOutboxOwnershipTransferred)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetOutboxRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the SolverNetOutbox contract.
type SolverNetOutboxRolesUpdatedIterator struct {
	Event *SolverNetOutboxRolesUpdated // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxRolesUpdated)
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
		it.Event = new(SolverNetOutboxRolesUpdated)
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
func (it *SolverNetOutboxRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxRolesUpdated represents a RolesUpdated event raised by the SolverNetOutbox contract.
type SolverNetOutboxRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*SolverNetOutboxRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxRolesUpdatedIterator{contract: _SolverNetOutbox.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxRolesUpdated)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
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

// ParseRolesUpdated is a log parse operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseRolesUpdated(log types.Log) (*SolverNetOutboxRolesUpdated, error) {
	event := new(SolverNetOutboxRolesUpdated)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
