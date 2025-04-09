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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"didFill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fillerData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fillFee\",\"inputs\":[{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executor_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setInboxes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"inboxes\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Filled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filledBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboxSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"inbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Open\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"resolvedOrder\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadFillerData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FillDeadlinePassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongDestChain\",\"inputs\":[]}]",
	Bin: "0x60a060405234801562000010575f80fd5b5063ffffffff60643b1615620000965760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156200007e575060408051601f3d908101601f191682019092526200007b9181019062000112565b60015b6200008d57436080526200009b565b6080526200009b565b436080525b620000a5620000ab565b6200012a565b63409feecd1980546001811615620000ca5763f92ee8a95f526004601cfd5b6001600160401b03808260011c146200010d578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b5f6020828403121562000123575f80fd5b5051919050565b608051611d2b620001435f395f61034d0152611d2b5ff3fe608060405260043610610131575f3560e01c8063715018a6116100a8578063c34c08e51161006d578063c34c08e51461031f578063eae4c19f1461033c578063f04e283e1461036f578063f2fde38b14610382578063f8c8765e14610395578063fee81cf4146103b4575f80fd5b8063715018a61461029c57806374eeb847146102a457806382e2c43f146102d55780638da5cb5b146102e8578063ac7318a714610300575f80fd5b80632773a339116100f95780632773a339146101b85780632de94807146101e557806339acf9f1146102165780634a4ee7b11461024c578063514e62fc1461025f57806354d1f13d14610294575f80fd5b8063183a4f6e146101355780631c10893f1461014a5780631cd64df41461015d578063248689cc1461019157806325692962146101b0575b5f80fd5b610148610143366004611374565b6103e5565b005b61014861015836600461139f565b6103f2565b348015610168575f80fd5b5061017c61017736600461139f565b610408565b60405190151581526020015b60405180910390f35b34801561019c575f80fd5b5061017c6101ab36600461140d565b610426565b610148610484565b3480156101c3575f80fd5b506101d76101d2366004611454565b6104d0565b604051908152602001610188565b3480156101f0575f80fd5b506101d76101ff366004611492565b638b78c6d8600c9081525f91909152602090205490565b348015610221575f80fd5b505f54610234906001600160a01b031681565b6040516001600160a01b039091168152602001610188565b61014861025a36600461139f565b610548565b34801561026a575f80fd5b5061017c61027936600461139f565b638b78c6d8600c9081525f9290925260209091205416151590565b61014861055a565b610148610593565b3480156102af575f80fd5b505f546102c390600160a01b900460ff1681565b60405160ff9091168152602001610188565b6101486102e33660046114b4565b6105a6565b3480156102f3575f80fd5b50638b78c6d81954610234565b34801561030b575f80fd5b5061014861031a366004611567565b6106ba565b34801561032a575f80fd5b506003546001600160a01b0316610234565b348015610347575f80fd5b506101d77f000000000000000000000000000000000000000000000000000000000000000081565b61014861037d366004611492565b61080e565b610148610390366004611492565b610848565b3480156103a0575f80fd5b506101486103af3660046115cd565b61086e565b3480156103bf575f80fd5b506101d76103ce366004611492565b63389a75e1600c9081525f91909152602090205490565b6103ef3382610918565b50565b6103fa610923565b610404828261093d565b5050565b638b78c6d8600c9081525f8390526020902054811681145b92915050565b5f60045f6104698686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061094992505050565b815260208101919091526040015f205460ff16949350505050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f80a250565b5f806104de8385018561191a565b80516040515f196024820181905260448201526001600160a01b0360648201529192506105409160840160408051601f198184030181529190526020810180516001600160e01b0316632d62234360e01b17905261053b8461097b565b610a3e565b949350505050565b610550610923565b6104048282610918565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f80a2565b61059b610923565b6105a45f610ab1565b565b60016105b181610aee565b3068929eee149b4bd2126854036105cf5763ab143c065f526004601cfd5b3068929eee149b4bd21268555f6105e88587018761191a565b602081015190915033906001600160401b0316461461061a5760405163fd24301760e01b815260040160405180910390fd5b42826040015163ffffffff161015610645576040516302857b7560e01b815260040160405180910390fd5b8315801590610655575060208414155b1561067357604051637bca9dfb60e11b815260040160405180910390fd5b602084900361068b5761068884860186611492565b90505b5f61069583610b12565b90506106a389848484610eac565b5050503868929eee149b4bd2126855505050505050565b6106c2610923565b8281146106e257604051634ec4810560e11b815260040160405180910390fd5b5f5b83811015610807578282828181106106fe576106fe6119e1565b90506020020160208101906107139190611492565b60025f878785818110610728576107286119e1565b905060200201602081019061073d91906119f5565b6001600160401b0316815260208101919091526040015f2080546001600160a01b0319166001600160a01b0392909216919091179055828282818110610785576107856119e1565b905060200201602081019061079a9190611492565b6001600160a01b03168585838181106107b5576107b56119e1565b90506020020160208101906107ca91906119f5565b6001600160401b03167f24de2e051f42b7671445d76d7ce15f3805ce949f655541a45ad259358f43c34a60405160405180910390a36001016106e4565b5050505050565b610816610923565b63389a75e1600c52805f526020600c20805442111561083c57636f5e88185f526004601cfd5b5f90556103ef81610ab1565b610850610923565b8060601b61086557637448fbae5f526004601cfd5b6103ef81610ab1565b63409feecd1980546003825580156108a45760018160011c14303b1061089b5763f92ee8a95f526004601cfd5b818160ff1b1b91505b506108ae85611020565b6108b984600161093d565b6108c28361105b565b600380546001600160a01b0319166001600160a01b0384161790558015610807576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15050505050565b61040482825f6110fe565b638b78c6d8195433146105a4576382b429005f526004601cfd5b610404828260016110fe565b5f828260405160200161095d929190611a5b565b60405160208183030381529060405280519060200120905092915050565b5f6109c480825b846060015151811015610a04575f856060015182815181106109a6576109a66119e1565b602002602001015190505f81606001515190506109cc8260600151516020016020611155565b6109c40261138801840193506109e6818262080000611174565b6109f1826020611155565b6003020193909301925050600101610982565b50608084015151611388026109c40180610a1e8385611a87565b610a289190611a87565b610a3590620186a0611a87565b95945050505050565b5f8054604051632376548f60e21b81526001600160a01b0390911690638dd9523c90610a7290879087908790600401611a9a565b602060405180830381865afa158015610a8d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105409190611acf565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a355565b638b78c6d8600c52335f52806020600c2054166103ef576382b429005f526004601cfd5b5f81608001515f5b8151811015610c02575f828281518110610b3657610b366119e1565b602090810291909101810151805191810151604082015160035492945090916001600160601b0390911690610b7a906001600160a01b0380851691339116846111a1565b6001600160a01b03831615610bf35760035460405163e1f21c6760e01b81526001600160a01b0384811660048301528581166024830152604482018490529091169063e1f21c67906064015f604051808303815f87803b158015610bdc575f80fd5b505af1158015610bee573d5f803e3d5ffd5b505050505b50505050806001019050610b1a565b505f805b846060015151811015610cee575f85606001518281518110610c2a57610c2a6119e1565b6020026020010151905060035f9054906101000a90046001600160a01b03166001600160a01b031663b61d27f68260400151835f0151846040015185602001518660600151604051602001610c80929190611ae6565b6040516020818303038152906040526040518563ffffffff1660e01b8152600401610cad93929190611b16565b5f604051808303818588803b158015610cc4575f80fd5b505af1158015610cd6573d5f803e3d5ffd5b50505060409092015193909301925050600101610c06565b5091505f5b8151811015610e31575f828281518110610d0f57610d0f6119e1565b60209081029190910181015190810151600354919250905f90610d3e906001600160a01b0380851691166111fa565b90508015610e235782516001600160a01b03811615610dba57600354604051639d32b56960e01b81526001600160a01b038581166004830152838116602483015290911690639d32b569906044015f604051808303815f87803b158015610da3575f80fd5b505af1158015610db5573d5f803e3d5ffd5b505050505b6003546040516317d5759960e31b81526001600160a01b038581166004830152336024830152604482018590529091169063beabacc8906064015f604051808303815f87803b158015610e0b575f80fd5b505af1158015610e1d573d5f803e3d5ffd5b50505050505b505050806001019050610cf3565b506003546001600160a01b0316318015610ea557600354604051633e97486160e11b8152336004820152602481018390526001600160a01b0390911690637d2e90c2906044015f604051808303815f87803b158015610e8e575f80fd5b505af1158015610ea0573d5f803e3d5ffd5b505050505b5050919050565b5f610ed68585604051602001610ec29190611ba1565b604051602081830303815290604052610949565b5f8181526004602052604090205490915060ff1615610f08576040516341a26a6360e01b815260040160405180910390fd5b5f818152600460208181526040808420805460ff1916600117905587516001600160401b03811685526002909252808420549051602481018a9052604481018690526001600160a01b038881166064830152610f9c9492169060840160408051601f198184030181529190526020810180516001600160e01b0316632d62234360e01b179052610f978a61097b565b611224565b90505f610fa98285611a87565b905080341015610fcb5760405162976f7560e21b815260040160405180910390fd5b5f610fd68234611c94565b90508015610fe857610fe8338261135b565b604051339085908a907fa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc905f90a45050505050505050565b6001600160a01b0316638b78c6d819819055805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b0381166110ab5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b60448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52825f526020600c2080548381178361111f575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe265f80a3505050505050565b5f81611168576365244e4e5f526004601cfd5b50808206151591040190565b828202831584820484141782026111925763ad251c275f526004601cfd5b81810615159190040192915050565b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f5114166111ec57803d873b1517106111ec57637939f4245f526004601cfd5b505f60605260405250505050565b5f816014526370a0823160601b5f5260208060246010865afa601f3d111660205102905092915050565b5f8054604051632376548f60e21b815282916001600160a01b031690638dd9523c90611258908a9088908890600401611a9a565b602060405180830381865afa158015611273573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112979190611acf565b9050804710156112e95760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e6473000000000000000060448201526064016110a2565b5f5460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f908390611322908b908b908b908b908b90600401611ca7565b5f604051808303818588803b158015611339575f80fd5b505af115801561134b573d5f803e3d5ffd5b50939a9950505050505050505050565b5f385f3884865af16104045763b12d13eb5f526004601cfd5b5f60208284031215611384575f80fd5b5035919050565b6001600160a01b03811681146103ef575f80fd5b5f80604083850312156113b0575f80fd5b82356113bb8161138b565b946020939093013593505050565b5f8083601f8401126113d9575f80fd5b5081356001600160401b038111156113ef575f80fd5b602083019150836020828501011115611406575f80fd5b9250929050565b5f805f6040848603121561141f575f80fd5b8335925060208401356001600160401b0381111561143b575f80fd5b611447868287016113c9565b9497909650939450505050565b5f8060208385031215611465575f80fd5b82356001600160401b0381111561147a575f80fd5b611486858286016113c9565b90969095509350505050565b5f602082840312156114a2575f80fd5b81356114ad8161138b565b9392505050565b5f805f805f606086880312156114c8575f80fd5b8535945060208601356001600160401b03808211156114e5575f80fd5b6114f189838a016113c9565b90965094506040880135915080821115611509575f80fd5b50611516888289016113c9565b969995985093965092949392505050565b5f8083601f840112611537575f80fd5b5081356001600160401b0381111561154d575f80fd5b6020830191508360208260051b8501011115611406575f80fd5b5f805f806040858703121561157a575f80fd5b84356001600160401b0380821115611590575f80fd5b61159c88838901611527565b909650945060208701359150808211156115b4575f80fd5b506115c187828801611527565b95989497509550505050565b5f805f80608085870312156115e0575f80fd5b84356115eb8161138b565b935060208501356115fb8161138b565b9250604085013561160b8161138b565b9150606085013561161b8161138b565b939692955090935050565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b038111828210171561165c5761165c611626565b60405290565b604051606081016001600160401b038111828210171561165c5761165c611626565b60405160a081016001600160401b038111828210171561165c5761165c611626565b604051601f8201601f191681016001600160401b03811182821017156116ce576116ce611626565b604052919050565b80356001600160401b03811681146116ec575f80fd5b919050565b5f6001600160401b0382111561170957611709611626565b5060051b60200190565b5f601f83601f840112611724575f80fd5b82356020611739611734836116f1565b6116a6565b82815260059290921b85018101918181019087841115611757575f80fd5b8287015b848110156118595780356001600160401b0380821115611779575f80fd5b908901906080601f19838d038101821315611792575f80fd5b61179a61163a565b888501356117a78161138b565b81526040858101356001600160e01b0319811681146117c4575f80fd5b828b0152606086810135828401529386013593858511156117e3575f80fd5b84870196508f603f8801126117f6575f80fd5b8a87013594508585111561180c5761180c611626565b61181b8b858f880116016116a6565b95508486528f82868901011115611830575f80fd5b848288018c8801375f9486018b019490945250918201929092528552505091830191830161175b565b50979650505050505050565b5f82601f830112611874575f80fd5b81356020611884611734836116f1565b828152606092830285018201928282019190878511156118a2575f80fd5b8387015b8581101561190d5781818a0312156118bc575f80fd5b6118c4611662565b81356118cf8161138b565b8152818601356118de8161138b565b818701526040828101356001600160601b03811681146118fc575f80fd5b9082015284529284019281016118a6565b5090979650505050505050565b5f6020828403121561192a575f80fd5b81356001600160401b0380821115611940575f80fd5b9083019060a08286031215611953575f80fd5b61195b611684565b611964836116d6565b8152611972602084016116d6565b6020820152604083013563ffffffff8116811461198d575f80fd5b60408201526060830135828111156119a3575f80fd5b6119af87828601611713565b6060830152506080830135828111156119c6575f80fd5b6119d287828601611865565b60808301525095945050505050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215611a05575f80fd5b6114ad826116d6565b5f5b83811015611a28578181015183820152602001611a10565b50505f910152565b5f8151808452611a47816020860160208601611a0e565b601f01601f19169290920160200192915050565b828152604060208201525f6105406040830184611a30565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561042057610420611a73565b5f6001600160401b03808616835260606020840152611abc6060840186611a30565b9150808416604084015250949350505050565b5f60208284031215611adf575f80fd5b5051919050565b6001600160e01b03198316815281515f90611b08816004850160208701611a0e565b919091016004019392505050565b60018060a01b0384168152826020820152606060408201525f610a356060830184611a30565b5f815180845260208085019450602084015f5b83811015611b9657815180516001600160a01b0390811689528482015116848901526040908101516001600160601b03169088015260609096019590820190600101611b4f565b509495945050505050565b5f602080835260c083016001600160401b0380865116838601528286015160408282166040880152604088015192506060915063ffffffff8316606088015260608801519250608060a0608089015284845180875260e08a01915060e08160051b8b0101965087860195505f5b81811015611c70578a880360df19018352865180516001600160a01b03168952898101516001600160e01b0319168a8a015285810151868a0152860151868901859052611c5d858a0182611a30565b9850509588019591880191600101611c0e565b5050505050505060808501519150601f198482030160a0850152610a358183611b3c565b8181038181111561042057610420611a73565b5f6001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a06060840152611ce060a0840186611a30565b9150808416608084015250969550505050505056fea2646970667358221220b36ed8257d2285dbce6d2607f31355ac507163ba3fe777046cf313181ea3b3e764736f6c63430008180033",
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

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxSession) Executor() (common.Address, error) {
	return _SolverNetOutbox.Contract.Executor(&_SolverNetOutbox.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) Executor() (common.Address, error) {
	return _SolverNetOutbox.Contract.Executor(&_SolverNetOutbox.CallOpts)
}

// FillFee is a free data retrieval call binding the contract method 0x2773a339.
//
// Solidity: function fillFee(bytes originData) view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxCaller) FillFee(opts *bind.CallOpts, originData []byte) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "fillFee", originData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillFee is a free data retrieval call binding the contract method 0x2773a339.
//
// Solidity: function fillFee(bytes originData) view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxSession) FillFee(originData []byte) (*big.Int, error) {
	return _SolverNetOutbox.Contract.FillFee(&_SolverNetOutbox.CallOpts, originData)
}

// FillFee is a free data retrieval call binding the contract method 0x2773a339.
//
// Solidity: function fillFee(bytes originData) view returns(uint256)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) FillFee(originData []byte) (*big.Int, error) {
	return _SolverNetOutbox.Contract.FillFee(&_SolverNetOutbox.CallOpts, originData)
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
// Solidity: function fill(bytes32 orderId, bytes originData, bytes fillerData) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) Fill(opts *bind.TransactOpts, orderId [32]byte, originData []byte, fillerData []byte) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "fill", orderId, originData, fillerData)
}

// Fill is a paid mutator transaction binding the contract method 0x82e2c43f.
//
// Solidity: function fill(bytes32 orderId, bytes originData, bytes fillerData) payable returns()
func (_SolverNetOutbox *SolverNetOutboxSession) Fill(orderId [32]byte, originData []byte, fillerData []byte) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Fill(&_SolverNetOutbox.TransactOpts, orderId, originData, fillerData)
}

// Fill is a paid mutator transaction binding the contract method 0x82e2c43f.
//
// Solidity: function fill(bytes32 orderId, bytes originData, bytes fillerData) payable returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) Fill(orderId [32]byte, originData []byte, fillerData []byte) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Fill(&_SolverNetOutbox.TransactOpts, orderId, originData, fillerData)
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
// Solidity: function initialize(address owner_, address solver_, address omni_, address executor_) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, solver_ common.Address, omni_ common.Address, executor_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "initialize", owner_, solver_, omni_, executor_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address executor_) returns()
func (_SolverNetOutbox *SolverNetOutboxSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address, executor_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Initialize(&_SolverNetOutbox.TransactOpts, owner_, solver_, omni_, executor_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address executor_) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address, executor_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Initialize(&_SolverNetOutbox.TransactOpts, owner_, solver_, omni_, executor_)
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

// SetInboxes is a paid mutator transaction binding the contract method 0xac7318a7.
//
// Solidity: function setInboxes(uint64[] chainIds, address[] inboxes) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) SetInboxes(opts *bind.TransactOpts, chainIds []uint64, inboxes []common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "setInboxes", chainIds, inboxes)
}

// SetInboxes is a paid mutator transaction binding the contract method 0xac7318a7.
//
// Solidity: function setInboxes(uint64[] chainIds, address[] inboxes) returns()
func (_SolverNetOutbox *SolverNetOutboxSession) SetInboxes(chainIds []uint64, inboxes []common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.SetInboxes(&_SolverNetOutbox.TransactOpts, chainIds, inboxes)
}

// SetInboxes is a paid mutator transaction binding the contract method 0xac7318a7.
//
// Solidity: function setInboxes(uint64[] chainIds, address[] inboxes) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) SetInboxes(chainIds []uint64, inboxes []common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.SetInboxes(&_SolverNetOutbox.TransactOpts, chainIds, inboxes)
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

// SolverNetOutboxInboxSetIterator is returned from FilterInboxSet and is used to iterate over the raw logs and unpacked data for InboxSet events raised by the SolverNetOutbox contract.
type SolverNetOutboxInboxSetIterator struct {
	Event *SolverNetOutboxInboxSet // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxInboxSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxInboxSet)
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
		it.Event = new(SolverNetOutboxInboxSet)
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
func (it *SolverNetOutboxInboxSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxInboxSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxInboxSet represents a InboxSet event raised by the SolverNetOutbox contract.
type SolverNetOutboxInboxSet struct {
	ChainId uint64
	Inbox   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxSet is a free log retrieval operation binding the contract event 0x24de2e051f42b7671445d76d7ce15f3805ce949f655541a45ad259358f43c34a.
//
// Solidity: event InboxSet(uint64 indexed chainId, address indexed inbox)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterInboxSet(opts *bind.FilterOpts, chainId []uint64, inbox []common.Address) (*SolverNetOutboxInboxSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "InboxSet", chainIdRule, inboxRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxInboxSetIterator{contract: _SolverNetOutbox.contract, event: "InboxSet", logs: logs, sub: sub}, nil
}

// WatchInboxSet is a free log subscription operation binding the contract event 0x24de2e051f42b7671445d76d7ce15f3805ce949f655541a45ad259358f43c34a.
//
// Solidity: event InboxSet(uint64 indexed chainId, address indexed inbox)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchInboxSet(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxInboxSet, chainId []uint64, inbox []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "InboxSet", chainIdRule, inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxInboxSet)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "InboxSet", log); err != nil {
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

// ParseInboxSet is a log parse operation binding the contract event 0x24de2e051f42b7671445d76d7ce15f3805ce949f655541a45ad259358f43c34a.
//
// Solidity: event InboxSet(uint64 indexed chainId, address indexed inbox)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseInboxSet(log types.Log) (*SolverNetOutboxInboxSet, error) {
	event := new(SolverNetOutboxInboxSet)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "InboxSet", log); err != nil {
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
