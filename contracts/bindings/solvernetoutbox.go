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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"didFill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fillerData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fillFee\",\"inputs\":[{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setInboxes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"inboxes\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Filled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filledBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboxSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"inbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Open\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"resolvedOrder\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadFillerData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FillDeadlinePassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongDestChain\",\"inputs\":[]}]",
	Bin: "0x60a06040523480156200001157600080fd5b5063ffffffff60643b1615620000975760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156200007f575060408051601f3d908101601f191682019092526200007c9181019062000114565b60015b6200008e57436080526200009c565b6080526200009c565b436080525b620000a6620000ac565b6200012e565b63409feecd1980546001811615620000cc5763f92ee8a96000526004601cfd5b6001600160401b03808260011c146200010f578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b6000602082840312156200012757600080fd5b5051919050565b60805161237b6200014a6000396000610381015261237b6000f3fe6080604052600436106101355760003560e01c8063715018a6116100ab578063c0c53b8b1161006f578063c0c53b8b14610331578063c34c08e514610351578063eae4c19f1461036f578063f04e283e146103a3578063f2fde38b146103b6578063fee81cf4146103c957600080fd5b8063715018a6146102aa57806374eeb847146102b257806382e2c43f146102e55780638da5cb5b146102f8578063ac7318a71461031157600080fd5b80632773a339116100fd5780632773a339146101bf5780632de94807146101ed57806339acf9f1146102205780634a4ee7b114610258578063514e62fc1461026b57806354d1f13d146102a257600080fd5b8063183a4f6e1461013a5780631c10893f1461014f5780631cd64df414610162578063248689cc1461019757806325692962146101b7575b600080fd5b61014d6101483660046113ca565b6103fc565b005b61014d61015d3660046113f8565b610409565b34801561016e57600080fd5b5061018261017d3660046113f8565b61041f565b60405190151581526020015b60405180910390f35b3480156101a357600080fd5b506101826101b236600461146c565b61043e565b61014d6104a0565b3480156101cb57600080fd5b506101df6101da3660046114b7565b6104ef565b60405190815260200161018e565b3480156101f957600080fd5b506101df6102083660046114f8565b638b78c6d8600c908152600091909152602090205490565b34801561022c57600080fd5b50600054610240906001600160a01b031681565b6040516001600160a01b03909116815260200161018e565b61014d6102663660046113f8565b610569565b34801561027757600080fd5b506101826102863660046113f8565b638b78c6d8600c90815260009290925260209091205416151590565b61014d61057b565b61014d6105b7565b3480156102be57600080fd5b506000546102d390600160a01b900460ff1681565b60405160ff909116815260200161018e565b61014d6102f336600461151c565b6105cb565b34801561030457600080fd5b50638b78c6d81954610240565b34801561031d57600080fd5b5061014d61032c3660046115d9565b6106e2565b34801561033d57600080fd5b5061014d61034c366004611644565b610819565b34801561035d57600080fd5b506003546001600160a01b0316610240565b34801561037b57600080fd5b506101df7f000000000000000000000000000000000000000000000000000000000000000081565b61014d6103b13660046114f8565b610903565b61014d6103c43660046114f8565b610940565b3480156103d557600080fd5b506101df6103e43660046114f8565b63389a75e1600c908152600091909152602090205490565b6104063382610967565b50565b610411610973565b61041b828261098e565b5050565b638b78c6d8600c90815260008390526020902054811681145b92915050565b6000600460006104848686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061099a92505050565b815260208101919091526040016000205460ff16949350505050565b60006202a3006001600160401b03164201905063389a75e1600c5233600052806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d600080a250565b6000806104fe838501856119a2565b80516040516000196024820181905260448201526001600160a01b0360648201529192506105619160840160408051601f198184030181529190526020810180516001600160e01b0316632d62234360e01b17905261055c846109cd565b610a67565b949350505050565b610571610973565b61041b8282610967565b63389a75e1600c523360005260006020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92600080a2565b6105bf610973565b6105c96000610add565b565b60016105d681610b1b565b3068929eee149b4bd2126854036105f55763ab143c066000526004601cfd5b3068929eee149b4bd2126855600061060f858701876119a2565b602081015190915033906001600160401b031646146106415760405163fd24301760e01b815260040160405180910390fd5b42826040015163ffffffff16101561066c576040516302857b7560e01b815260040160405180910390fd5b831580159061067c575060208414155b1561069a57604051637bca9dfb60e11b815260040160405180910390fd5b60208490036106b2576106af848601866114f8565b90505b60006106bd83610b41565b90506106cb89848484610f04565b5050503868929eee149b4bd2126855505050505050565b6106ea610973565b60005b838110156108125782828281811061070757610707611a70565b905060200201602081019061071c91906114f8565b6002600087878581811061073257610732611a70565b90506020020160208101906107479190611a86565b6001600160401b03168152602081019190915260400160002080546001600160a01b0319166001600160a01b039290921691909117905582828281811061079057610790611a70565b90506020020160208101906107a591906114f8565b6001600160a01b03168585838181106107c0576107c0611a70565b90506020020160208101906107d59190611a86565b6001600160401b03167f24de2e051f42b7671445d76d7ce15f3805ce949f655541a45ad259358f43c34a60405160405180910390a36001016106ed565b5050505050565b63409feecd1980546003825580156108505760018160011c14303b106108475763f92ee8a96000526004601cfd5b818160ff1b1b91505b5061085a8461107e565b61086583600161098e565b61086e826110ba565b3060405161087b906113bd565b6001600160a01b039091168152602001604051809103906000f0801580156108a7573d6000803e3d6000fd5b50600380546001600160a01b0319166001600160a01b039290921691909117905580156108fd576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b50505050565b61090b610973565b63389a75e1600c52806000526020600c20805442111561093357636f5e88186000526004601cfd5b6000905561040681610add565b610948610973565b8060601b61095e57637448fbae6000526004601cfd5b61040681610add565b61041b8282600061115e565b638b78c6d8195433146105c9576382b429006000526004601cfd5b61041b8282600161115e565b600082826040516020016109af929190611af1565b60405160208183030381529060405280519060200120905092915050565b60006109c480825b846060015151811015610a2d576000856060015182815181106109fa576109fa611a70565b60200260200101519050610a1781606001515160200160206111b7565b6109c402929092016113880191506001016109d5565b50608084015151611388026109c40180610a478385611b20565b610a519190611b20565b610a5e90620186a0611b20565b95945050505050565b60008054604051632376548f60e21b81526001600160a01b0390911690638dd9523c90610a9c90879087908790600401611b33565b602060405180830381865afa158015610ab9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105619190611b69565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a355565b638b78c6d8600c5233600052806020600c205416610406576382b429006000526004601cfd5b6000816080015160005b8151811015610c39576000828281518110610b6857610b68611a70565b602090810291909101810151805191810151604082015160035492945090916001600160601b0390911690610bac906001600160a01b0380851691339116846111d8565b6001600160a01b03831615610c2a5760035460405163e1f21c6760e01b81526001600160a01b0384811660048301528581166024830152604482018490529091169063e1f21c6790606401600060405180830381600087803b158015610c1157600080fd5b505af1158015610c25573d6000803e3d6000fd5b505050505b50505050806001019050610b4b565b506000805b846060015151811015610d2d57600085606001518281518110610c6357610c63611a70565b60200260200101519050600360009054906101000a90046001600160a01b03166001600160a01b031663b61d27f682604001518360000151846040015185602001518660600151604051602001610cbb929190611b82565b6040516020818303038152906040526040518563ffffffff1660e01b8152600401610ce893929190611bb3565b6000604051808303818588803b158015610d0157600080fd5b505af1158015610d15573d6000803e3d6000fd5b50505060409092015193909301925050600101610c3e565b50915060005b8151811015610e84576000828281518110610d5057610d50611a70565b6020908102919091018101519081015160035491925090600090610d80906001600160a01b038085169116611236565b90508015610e765782516001600160a01b03811615610e085760035460405163e1f21c6760e01b81526001600160a01b0385811660048301528381166024830152600060448301529091169063e1f21c6790606401600060405180830381600087803b158015610def57600080fd5b505af1158015610e03573d6000803e3d6000fd5b505050505b6003546040516317d5759960e31b81526001600160a01b038581166004830152336024830152604482018590529091169063beabacc890606401600060405180830381600087803b158015610e5c57600080fd5b505af1158015610e70573d6000803e3d6000fd5b50505050505b505050806001019050610d33565b506003546001600160a01b0316318015610efd57600354604051633e97486160e11b8152336004820152602481018390526001600160a01b0390911690637d2e90c290604401600060405180830381600087803b158015610ee457600080fd5b505af1158015610ef8573d6000803e3d6000fd5b505050505b5050919050565b6000610f2f8585604051602001610f1b9190611c41565b60405160208183030381529060405261099a565b60008181526004602052604090205490915060ff1615610f62576040516341a26a6360e01b815260040160405180910390fd5b6000818152600460208181526040808420805460ff1916600117905587516001600160401b03811685526002909252808420549051602481018a9052604481018690526001600160a01b038881166064830152610ff79492169060840160408051601f198184030181529190526020810180516001600160e01b0316632d62234360e01b179052610ff28a6109cd565b611262565b905060006110058285611b20565b9050803410156110275760405162976f7560e21b815260040160405180910390fd5b60006110338234611d36565b905080156110455761104533826113a1565b604051339085908a907fa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc90600090a45050505050505050565b6001600160a01b0316638b78c6d8198190558060007f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b03811661110a5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b60448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52826000526020600c20805483811783611180575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26600080a3505050505050565b6000816111cc576365244e4e6000526004601cfd5b50808206151591040190565b60405181606052826040528360601b602c526323b872dd60601b600c52602060006064601c6000895af1806001600051141661122757803d873b15171061122757637939f4246000526004601cfd5b50600060605260405250505050565b6000816014526370a0823160601b60005260208060246010865afa601f3d111660205102905092915050565b60008054604051632376548f60e21b815282916001600160a01b031690638dd9523c90611297908a9088908890600401611b33565b602060405180830381865afa1580156112b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d89190611b69565b90508047101561132a5760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401611101565b60005460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f908390611364908b908b908b908b908b90600401611d49565b6000604051808303818588803b15801561137d57600080fd5b505af1158015611391573d6000803e3d6000fd5b50939a9950505050505050505050565b60003860003884865af161041b5763b12d13eb6000526004601cfd5b6105ad80611d9983390190565b6000602082840312156113dc57600080fd5b5035919050565b6001600160a01b038116811461040657600080fd5b6000806040838503121561140b57600080fd5b8235611416816113e3565b946020939093013593505050565b60008083601f84011261143657600080fd5b5081356001600160401b0381111561144d57600080fd5b60208301915083602082850101111561146557600080fd5b9250929050565b60008060006040848603121561148157600080fd5b8335925060208401356001600160401b0381111561149e57600080fd5b6114aa86828701611424565b9497909650939450505050565b600080602083850312156114ca57600080fd5b82356001600160401b038111156114e057600080fd5b6114ec85828601611424565b90969095509350505050565b60006020828403121561150a57600080fd5b8135611515816113e3565b9392505050565b60008060008060006060868803121561153457600080fd5b8535945060208601356001600160401b038082111561155257600080fd5b61155e89838a01611424565b9096509450604088013591508082111561157757600080fd5b5061158488828901611424565b969995985093965092949392505050565b60008083601f8401126115a757600080fd5b5081356001600160401b038111156115be57600080fd5b6020830191508360208260051b850101111561146557600080fd5b600080600080604085870312156115ef57600080fd5b84356001600160401b038082111561160657600080fd5b61161288838901611595565b9096509450602087013591508082111561162b57600080fd5b5061163887828801611595565b95989497509550505050565b60008060006060848603121561165957600080fd5b8335611664816113e3565b92506020840135611674816113e3565b91506040840135611684816113e3565b809150509250925092565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156116c7576116c761168f565b60405290565b604051606081016001600160401b03811182821017156116c7576116c761168f565b60405160a081016001600160401b03811182821017156116c7576116c761168f565b604051601f8201601f191681016001600160401b03811182821017156117395761173961168f565b604052919050565b80356001600160401b038116811461175857600080fd5b919050565b60006001600160401b038211156117765761177661168f565b5060051b60200190565b6000601f83601f84011261179357600080fd5b823560206117a86117a38361175d565b611711565b82815260059290921b850181019181810190878411156117c757600080fd5b8287015b848110156118da5780356001600160401b03808211156117eb5760008081fd5b908901906080601f19838d0381018213156118065760008081fd5b61180e6116a5565b8885013561181b816113e3565b81526040858101356001600160e01b03198116811461183a5760008081fd5b828b01526060868101358284015293860135938585111561185b5760008081fd5b84870196508f603f88011261187257600094508485fd5b8a8701359450858511156118885761188861168f565b6118978b858f88011601611711565b95508486528f828689010111156118b057600093508384fd5b848288018c88013760009486018b01949094525091820192909252855250509183019183016117cb565b50979650505050505050565b600082601f8301126118f757600080fd5b813560206119076117a38361175d565b8281526060928302850182019282820191908785111561192657600080fd5b8387015b858110156119955781818a0312156119425760008081fd5b61194a6116cd565b8135611955816113e3565b815281860135611964816113e3565b818701526040828101356001600160601b03811681146119845760008081fd5b90820152845292840192810161192a565b5090979650505050505050565b6000602082840312156119b457600080fd5b81356001600160401b03808211156119cb57600080fd5b9083019060a082860312156119df57600080fd5b6119e76116ef565b6119f083611741565b81526119fe60208401611741565b6020820152604083013563ffffffff81168114611a1a57600080fd5b6040820152606083013582811115611a3157600080fd5b611a3d87828601611780565b606083015250608083013582811115611a5557600080fd5b611a61878286016118e6565b60808301525095945050505050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611a9857600080fd5b61151582611741565b60005b83811015611abc578181015183820152602001611aa4565b50506000910152565b60008151808452611add816020860160208601611aa1565b601f01601f19169290920160200192915050565b8281526040602082015260006105616040830184611ac5565b634e487b7160e01b600052601160045260246000fd5b8082018082111561043857610438611b0a565b60006001600160401b03808616835260606020840152611b566060840186611ac5565b9150808416604084015250949350505050565b600060208284031215611b7b57600080fd5b5051919050565b6001600160e01b0319831681528151600090611ba5816004850160208701611aa1565b919091016004019392505050565b60018060a01b0384168152826020820152606060408201526000610a5e6060830184611ac5565b60008151808452602080850194506020840160005b83811015611c3657815180516001600160a01b0390811689528482015116848901526040908101516001600160601b03169088015260609096019590820190600101611bef565b509495945050505050565b6000602080835260c083016001600160401b0380865116838601528286015160408282166040880152604088015192506060915063ffffffff8316606088015260608801519250608060a0608089015284845180875260e08a01915060e08160051b8b01019650878601955060005b81811015611d12578a880360df19018352865180516001600160a01b03168952898101516001600160e01b0319168a8a015285810151868a0152860151868901859052611cff858a0182611ac5565b9850509588019591880191600101611cb0565b5050505050505060808501519150601f198482030160a0850152610a5e8183611bda565b8181038181111561043857610438611b0a565b60006001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a06060840152611d8360a0840186611ac5565b9150808416608084015250969550505050505056fe60a060405234801561001057600080fd5b506040516105ad3803806105ad83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516105076100a66000396000818160b901528181610122015281816101820152818161025301526102b501526105076000f3fe60806040526004361061004b5760003560e01c80637d2e90c214610054578063b61d27f614610074578063beabacc814610087578063ce11e6ab146100a7578063e1f21c67146100f757005b3661005257005b005b34801561006057600080fd5b5061005261006f3660046103d4565b610117565b6100526100823660046103fe565b610177565b34801561009357600080fd5b506100526100a2366004610485565b610248565b3480156100b357600080fd5b506100db7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b34801561010357600080fd5b50610052610112366004610485565b6102aa565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101605760405163bda8fc9560e01b815260040160405180910390fd5b6101736001600160a01b03831682610307565b5050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101c05760405163bda8fc9560e01b815260040160405180910390fd5b6000846001600160a01b03168484846040516101dd9291906104c1565b60006040518083038185875af1925050503d806000811461021a576040519150601f19603f3d011682016040523d82523d6000602084013e61021f565b606091505b505090508061024157604051633204506f60e01b815260040160405180910390fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146102915760405163bda8fc9560e01b815260040160405180910390fd5b6102a56001600160a01b0384168383610323565b505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146102f35760405163bda8fc9560e01b815260040160405180910390fd5b6102a56001600160a01b0384168383610373565b60003860003884865af16101735763b12d13eb6000526004601cfd5b816014528060345263a9059cbb60601b60005260206000604460106000875af1806001600051141661036857803d853b151710610368576390b8ec186000526004601cfd5b506000603452505050565b816014528060345263095ea7b360601b60005260206000604460106000875af1806001600051141661036857803d853b15171061036857633e3f8f736000526004601cfd5b80356001600160a01b03811681146103cf57600080fd5b919050565b600080604083850312156103e757600080fd5b6103f0836103b8565b946020939093013593505050565b6000806000806060858703121561041457600080fd5b61041d856103b8565b935060208501359250604085013567ffffffffffffffff8082111561044157600080fd5b818701915087601f83011261045557600080fd5b81358181111561046457600080fd5b88602082850101111561047657600080fd5b95989497505060200194505050565b60008060006060848603121561049a57600080fd5b6104a3846103b8565b92506104b1602085016103b8565b9150604084013590509250925092565b818382376000910190815291905056fea26469706673582212204e2cefec039437b1ebc715912a486f808020f056b422d619fba0f36429219cbe64736f6c63430008180033a2646970667358221220da3edfee4e44453b4a4dfa6f9e01b56866d25f916d6643b16cca3d0815af7c1664736f6c63430008180033",
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

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner_, address solver_, address omni_) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, solver_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "initialize", owner_, solver_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner_, address solver_, address omni_) returns()
func (_SolverNetOutbox *SolverNetOutboxSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Initialize(&_SolverNetOutbox.TransactOpts, owner_, solver_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner_, address solver_, address omni_) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.Initialize(&_SolverNetOutbox.TransactOpts, owner_, solver_, omni_)
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
