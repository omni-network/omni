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

// ISolverNetOutboxInboxConfig is an auto generated low-level Go binding around an user-defined struct.
type ISolverNetOutboxInboxConfig struct {
	Inbox    common.Address
	Provider uint8
}

// SolverNetOutboxMetaData contains all meta data concerning the SolverNetOutbox contract.
var SolverNetOutboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_mailbox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PACKAGE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"didFill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fill\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fillerData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fillFee\",\"inputs\":[{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executor_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"interchainSecurityModule\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIInterchainSecurityModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localDomain\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMailbox\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setInboxes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structISolverNetOutbox.InboxConfig[]\",\"components\":[{\"name\":\"inbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"uint8\",\"internalType\":\"enumISolverNetOutbox.Provider\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Filled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filledBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HookSet\",\"inputs\":[{\"name\":\"_hook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboxSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"inbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumISolverNetOutbox.Provider\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IsmSet\",\"inputs\":[{\"name\":\"_ism\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Open\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"resolvedOrder\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadFillerData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FillDeadlinePassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongDestChain\",\"inputs\":[]}]",
	Bin: "0x61010060405234801562000011575f80fd5b50604051620026c9380380620026c9833981016040819052620000349162000252565b8063ffffffff60643b1615620000ba5760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015620000a2575060408051601f3d908101601f191682019092526200009f9181019062000277565b60015b620000b15743608052620000bf565b608052620000bf565b436080525b6001600160a01b03811660a081905215620001c15760a0516001600160a01b0316636e5f516e6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000113573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000139919062000252565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b0316638d3638f46040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000191573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001b791906200028f565b63ffffffff1660e0525b50620001cc620001d3565b50620002b4565b63409feecd1980546001811615620001f25763f92ee8a95f526004601cfd5b6001600160401b03808260011c1462000235578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b6001600160a01b03811681146200024f575f80fd5b50565b5f6020828403121562000263575f80fd5b815162000270816200023a565b9392505050565b5f6020828403121562000288575f80fd5b5051919050565b5f60208284031215620002a0575f80fd5b815163ffffffff8116811462000270575f80fd5b60805160a05160c05160e0516123d3620002f65f395f61036201525f61044f01525f818161041c01528181610d23015261180801525f61048201526123d35ff3fe60806040526004361061017b575f3560e01c806382e2c43f116100cd578063d5438eae11610087578063f04e283e11610062578063f04e283e146104a4578063f2fde38b146104b7578063f8c8765e146104ca578063fee81cf4146104e9575f80fd5b8063d5438eae1461040b578063de523cf31461043e578063eae4c19f14610471575f80fd5b806382e2c43f1461031f57806382e51365146103325780638d3638f4146103515780638da5cb5b1461039957806393c44847146103b1578063c34c08e5146103ee575f80fd5b80632de9480711610138578063514e62fc11610113578063514e62fc146102a957806354d1f13d146102de578063715018a6146102e657806374eeb847146102ee575f80fd5b80632de948071461022f57806339acf9f1146102605780634a4ee7b114610296575f80fd5b8063183a4f6e1461017f5780631c10893f146101945780631cd64df4146101a7578063248689cc146101db57806325692962146101fa5780632773a33914610202575b5f80fd5b61019261018d3660046118bd565b61051a565b005b6101926101a23660046118e8565b610527565b3480156101b2575f80fd5b506101c66101c13660046118e8565b61053d565b60405190151581526020015b60405180910390f35b3480156101e6575f80fd5b506101c66101f5366004611956565b61055b565b6101926105b9565b34801561020d575f80fd5b5061022161021c36600461199d565b610605565b6040519081526020016101d2565b34801561023a575f80fd5b506102216102493660046119db565b638b78c6d8600c9081525f91909152602090205490565b34801561026b575f80fd5b505f5461027e906001600160a01b031681565b6040516001600160a01b0390911681526020016101d2565b6101926102a43660046118e8565b610787565b3480156102b4575f80fd5b506101c66102c33660046118e8565b638b78c6d8600c9081525f9290925260209091205416151590565b610192610799565b6101926107d2565b3480156102f9575f80fd5b505f5461030d90600160a01b900460ff1681565b60405160ff90911681526020016101d2565b61019261032d3660046119fd565b6107e5565b34801561033d575f80fd5b5061019261034c366004611a70565b6108f9565b34801561035c575f80fd5b506103847f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101d2565b3480156103a4575f80fd5b50638b78c6d8195461027e565b3480156103bc575f80fd5b506103e1604051806040016040528060058152602001640372e302e360dc1b81525081565b6040516101d29190611b79565b3480156103f9575f80fd5b506033546001600160a01b031661027e565b348015610416575f80fd5b5061027e7f000000000000000000000000000000000000000000000000000000000000000081565b348015610449575f80fd5b5061027e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561047c575f80fd5b506102217f000000000000000000000000000000000000000000000000000000000000000081565b6101926104b23660046119db565b610a69565b6101926104c53660046119db565b610aa3565b3480156104d5575f80fd5b506101926104e4366004611b8b565b610ac9565b3480156104f4575f80fd5b506102216105033660046119db565b63389a75e1600c9081525f91909152602090205490565b6105243382610b73565b50565b61052f610b7e565b6105398282610b98565b5050565b638b78c6d8600c9081525f8390526020902054811681145b92915050565b5f60345f61059e8686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610ba492505050565b815260208101919091526040015f205460ff16949350505050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f80a250565b5f8061061383850185611ed8565b80516001600160401b03165f908152603260209081526040808320815180830190925280546001600160a01b03811683529495509293909291830190600160a01b900460ff16600281111561066a5761066a611f9f565b600281111561067b5761067b611f9f565b905250905060018160200151600281111561069857610698611f9f565b036107065781516040515f196024820181905260448201526001600160a01b0360648201526106fd919060840160408051601f198184030181529190526020810180516001600160e01b0316632d62234360e01b1790526106f885610bd6565b610c99565b92505050610555565b60028160200151600281111561071e5761071e611f9f565b0361076e5781518151604080515f1960208201819052918101919091526001600160a01b0360608201526106fd92919060800160405160208183030381529060405261076986610bd6565b610d14565b6040516306b7c75960e31b815260040160405180910390fd5b61078f610b7e565b6105398282610b73565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f80a2565b6107da610b7e565b6107e35f610dd4565b565b60016107f081610e11565b3068929eee149b4bd21268540361080e5763ab143c065f526004601cfd5b3068929eee149b4bd21268555f61082785870187611ed8565b602081015190915033906001600160401b031646146108595760405163fd24301760e01b815260040160405180910390fd5b42826040015163ffffffff161015610884576040516302857b7560e01b815260040160405180910390fd5b8315801590610894575060208414155b156108b257604051637bca9dfb60e11b815260040160405180910390fd5b60208490036108ca576108c7848601866119db565b90505b5f6108d483610e35565b90506108e2898484846111cf565b5050503868929eee149b4bd2126855505050505050565b610901610b7e565b82811461092157604051634ec4810560e11b815260040160405180910390fd5b5f5b83811015610a625782828281811061093d5761093d611fb3565b90506040020160325f87878581811061095857610958611fb3565b905060200201602081019061096d9190611fc7565b6001600160401b0316815260208101919091526040015f2061098f8282611fec565b9050508282828181106109a4576109a4611fb3565b90506040020160200160208101906109bc9190612061565b60028111156109cd576109cd611f9f565b8383838181106109df576109df611fb3565b6109f592602060409092020190810191506119db565b6001600160a01b0316868684818110610a1057610a10611fb3565b9050602002016020810190610a259190611fc7565b6001600160401b03167f41d1e58dad7679eeaefbd227553cdede676d33bf46808e7b54b2550a4e44057e60405160405180910390a4600101610923565b5050505050565b610a71610b7e565b63389a75e1600c52805f526020600c208054421115610a9757636f5e88185f526004601cfd5b5f905561052481610dd4565b610aab610b7e565b8060601b610ac057637448fbae5f526004601cfd5b61052481610dd4565b63409feecd198054600382558015610aff5760018160011c14303b10610af65763f92ee8a95f526004601cfd5b818160ff1b1b91505b50610b09856112d3565b610b14846001610b98565b610b1d8361130e565b603380546001600160a01b0319166001600160a01b0384161790558015610a62576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15050505050565b61053982825f6113b1565b638b78c6d8195433146107e3576382b429005f526004601cfd5b610539828260016113b1565b5f8282604051602001610bb892919061207c565b60405160208183030381529060405280519060200120905092915050565b5f6109c480825b846060015151811015610c5f575f85606001518281518110610c0157610c01611fb3565b602002602001015190505f8160600151519050610c278260600151516020016020611408565b6109c4026113880184019350610c41818262080000611427565b610c4c826020611408565b6003020193909301925050600101610bdd565b50608084015151611388026109c40180610c7983856120a8565b610c8391906120a8565b610c9090620186a06120a8565b95945050505050565b5f8054604051632376548f60e21b81526001600160a01b0390911690638dd9523c90610ccd908790879087906004016120bb565b602060405180830381865afa158015610ce8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d0c91906120f0565b949350505050565b5f80610d1f83611454565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f7ccd32187610d6a886001600160a01b03166001600160a01b031690565b87856040518563ffffffff1660e01b8152600401610d8b9493929190612107565b602060405180830381865afa158015610da6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dca91906120f0565b9695505050505050565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a355565b638b78c6d8600c52335f52806020600c205416610524576382b429005f526004601cfd5b5f81608001515f5b8151811015610f25575f828281518110610e5957610e59611fb3565b602090810291909101810151805191810151604082015160335492945090916001600160601b0390911690610e9d906001600160a01b038085169133911684611470565b6001600160a01b03831615610f165760335460405163e1f21c6760e01b81526001600160a01b0384811660048301528581166024830152604482018490529091169063e1f21c67906064015f604051808303815f87803b158015610eff575f80fd5b505af1158015610f11573d5f803e3d5ffd5b505050505b50505050806001019050610e3d565b505f805b846060015151811015611011575f85606001518281518110610f4d57610f4d611fb3565b6020026020010151905060335f9054906101000a90046001600160a01b03166001600160a01b031663b61d27f68260400151835f0151846040015185602001518660600151604051602001610fa392919061213d565b6040516020818303038152906040526040518563ffffffff1660e01b8152600401610fd09392919061216d565b5f604051808303818588803b158015610fe7575f80fd5b505af1158015610ff9573d5f803e3d5ffd5b50505060409092015193909301925050600101610f29565b5091505f5b8151811015611154575f82828151811061103257611032611fb3565b60209081029190910181015190810151603354919250905f90611061906001600160a01b0380851691166114c9565b905080156111465782516001600160a01b038116156110dd57603354604051639d32b56960e01b81526001600160a01b038581166004830152838116602483015290911690639d32b569906044015f604051808303815f87803b1580156110c6575f80fd5b505af11580156110d8573d5f803e3d5ffd5b505050505b6033546040516317d5759960e31b81526001600160a01b038581166004830152336024830152604482018590529091169063beabacc8906064015f604051808303815f87803b15801561112e575f80fd5b505af1158015611140573d5f803e3d5ffd5b50505050505b505050806001019050611016565b506033546001600160a01b03163180156111c857603354604051633e97486160e11b8152336004820152602481018390526001600160a01b0390911690637d2e90c2906044015f604051808303815f87803b1580156111b1575f80fd5b505af11580156111c3573d5f803e3d5ffd5b505050505b5050919050565b5f6111f985856040516020016111e591906121f8565b604051602081830303815290604052610ba4565b5f8181526034602052604090205490915060ff161561122b576040516341a26a6360e01b815260040160405180910390fd5b5f818152603460205260408120805460ff1916600117905561124f868386886114f3565b90505f61125c82856120a8565b90508034101561127e5760405162976f7560e21b815260040160405180910390fd5b5f61128982346122eb565b9050801561129b5761129b3382611674565b604051339085908a907fa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc905f90a45050505050505050565b6001600160a01b0316638b78c6d819819055805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b03811661135e5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b60448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52825f526020600c208054838117836113d2575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe265f80a3505050505050565b5f8161141b576365244e4e5f526004601cfd5b50808206151591040190565b828202831584820484141782026114455763ad251c275f526004601cfd5b81810615159190040192915050565b60606105555f833360405180602001604052805f81525061168d565b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f5114166114bb57803d873b1517106114bb57637939f4245f526004601cfd5b505f60605260405250505050565b5f816014526370a0823160601b5f5260208060246010865afa601f3d111660205102905092915050565b80516001600160401b03165f908152603260209081526040808320815180830190925280546001600160a01b03811683528493830190600160a01b900460ff16600281111561154457611544611f9f565b600281111561155557611555611f9f565b90525090505f60018260200151600281111561157357611573611f9f565b036115e65783518251604051602481018a9052604481018990526001600160a01b03881660648201526115df929160049160840160408051601f198184030181529190526020810180516001600160e01b0316632d62234360e01b1790526115da89610bd6565b6116c2565b9050610dca565b6002826020015160028111156115fe576115fe611f9f565b0361076e5760408051602081018990529081018790526001600160a01b03861660608201525f9060800160408051601f198184030181529190528551845191925061164d918361076989610bd6565b855184519193506116689184846116638a610bd6565b6117f9565b50509695505050505050565b5f385f3884865af16105395763b12d13eb5f526004601cfd5b60606001858585856040516020016116a99594939291906122fe565b6040516020818303038152906040529050949350505050565b5f8054604051632376548f60e21b815282916001600160a01b031690638dd9523c906116f6908a90889088906004016120bb565b602060405180830381865afa158015611711573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061173591906120f0565b9050804710156117875760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401611355565b5f5460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f9083906117c0908b908b908b908b908b9060040161234f565b5f604051808303818588803b1580156117d7575f80fd5b505af11580156117e9573d5f803e3d5ffd5b50939a9950505050505050505050565b5f8061180483611454565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166348aee8d486896118508a6001600160a01b03166001600160a01b031690565b88866040518663ffffffff1660e01b81526004016118719493929190612107565b60206040518083038185885af115801561188d573d5f803e3d5ffd5b50505050506040513d601f19601f820116820180604052508101906118b291906120f0565b979650505050505050565b5f602082840312156118cd575f80fd5b5035919050565b6001600160a01b0381168114610524575f80fd5b5f80604083850312156118f9575f80fd5b8235611904816118d4565b946020939093013593505050565b5f8083601f840112611922575f80fd5b5081356001600160401b03811115611938575f80fd5b60208301915083602082850101111561194f575f80fd5b9250929050565b5f805f60408486031215611968575f80fd5b8335925060208401356001600160401b03811115611984575f80fd5b61199086828701611912565b9497909650939450505050565b5f80602083850312156119ae575f80fd5b82356001600160401b038111156119c3575f80fd5b6119cf85828601611912565b90969095509350505050565b5f602082840312156119eb575f80fd5b81356119f6816118d4565b9392505050565b5f805f805f60608688031215611a11575f80fd5b8535945060208601356001600160401b0380821115611a2e575f80fd5b611a3a89838a01611912565b90965094506040880135915080821115611a52575f80fd5b50611a5f88828901611912565b969995985093965092949392505050565b5f805f8060408587031215611a83575f80fd5b84356001600160401b0380821115611a99575f80fd5b818701915087601f830112611aac575f80fd5b813581811115611aba575f80fd5b8860208260051b8501011115611ace575f80fd5b602092830196509450908601359080821115611ae8575f80fd5b818701915087601f830112611afb575f80fd5b813581811115611b09575f80fd5b8860208260061b8501011115611b1d575f80fd5b95989497505060200194505050565b5f5b83811015611b46578181015183820152602001611b2e565b50505f910152565b5f8151808452611b65816020860160208601611b2c565b601f01601f19169290920160200192915050565b602081525f6119f66020830184611b4e565b5f805f8060808587031215611b9e575f80fd5b8435611ba9816118d4565b93506020850135611bb9816118d4565b92506040850135611bc9816118d4565b91506060850135611bd9816118d4565b939692955090935050565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b0381118282101715611c1a57611c1a611be4565b60405290565b604051606081016001600160401b0381118282101715611c1a57611c1a611be4565b60405160a081016001600160401b0381118282101715611c1a57611c1a611be4565b604051601f8201601f191681016001600160401b0381118282101715611c8c57611c8c611be4565b604052919050565b80356001600160401b0381168114611caa575f80fd5b919050565b5f6001600160401b03821115611cc757611cc7611be4565b5060051b60200190565b5f601f83601f840112611ce2575f80fd5b82356020611cf7611cf283611caf565b611c64565b82815260059290921b85018101918181019087841115611d15575f80fd5b8287015b84811015611e175780356001600160401b0380821115611d37575f80fd5b908901906080601f19838d038101821315611d50575f80fd5b611d58611bf8565b88850135611d65816118d4565b81526040858101356001600160e01b031981168114611d82575f80fd5b828b015260608681013582840152938601359385851115611da1575f80fd5b84870196508f603f880112611db4575f80fd5b8a870135945085851115611dca57611dca611be4565b611dd98b858f88011601611c64565b95508486528f82868901011115611dee575f80fd5b848288018c8801375f9486018b0194909452509182019290925285525050918301918301611d19565b50979650505050505050565b5f82601f830112611e32575f80fd5b81356020611e42611cf283611caf565b82815260609283028501820192828201919087851115611e60575f80fd5b8387015b85811015611ecb5781818a031215611e7a575f80fd5b611e82611c20565b8135611e8d816118d4565b815281860135611e9c816118d4565b818701526040828101356001600160601b0381168114611eba575f80fd5b908201528452928401928101611e64565b5090979650505050505050565b5f60208284031215611ee8575f80fd5b81356001600160401b0380821115611efe575f80fd5b9083019060a08286031215611f11575f80fd5b611f19611c42565b611f2283611c94565b8152611f3060208401611c94565b6020820152604083013563ffffffff81168114611f4b575f80fd5b6040820152606083013582811115611f61575f80fd5b611f6d87828601611cd1565b606083015250608083013582811115611f84575f80fd5b611f9087828601611e23565b60808301525095945050505050565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215611fd7575f80fd5b6119f682611c94565b60038110610524575f80fd5b8135611ff7816118d4565b81546001600160a01b031981166001600160a01b03929092169182178355602084013561202381611fe0565b6003811061203f57634e487b7160e01b5f52602160045260245ffd5b6001600160a81b03199190911690911760a09190911b60ff60a01b1617905550565b5f60208284031215612071575f80fd5b81356119f681611fe0565b828152604060208201525f610d0c6040830184611b4e565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561055557610555612094565b5f6001600160401b038086168352606060208401526120dd6060840186611b4e565b9150808416604084015250949350505050565b5f60208284031215612100575f80fd5b5051919050565b63ffffffff85168152836020820152608060408201525f61212b6080830185611b4e565b82810360608401526118b28185611b4e565b6001600160e01b03198316815281515f9061215f816004850160208701611b2c565b919091016004019392505050565b60018060a01b0384168152826020820152606060408201525f610c906060830184611b4e565b5f815180845260208085019450602084015f5b838110156121ed57815180516001600160a01b0390811689528482015116848901526040908101516001600160601b031690880152606090960195908201906001016121a6565b509495945050505050565b5f602080835260c083016001600160401b0380865116838601528286015160408282166040880152604088015192506060915063ffffffff8316606088015260608801519250608060a0608089015284845180875260e08a01915060e08160051b8b0101965087860195505f5b818110156122c7578a880360df19018352865180516001600160a01b03168952898101516001600160e01b0319168a8a015285810151868a01528601518689018590526122b4858a0182611b4e565b9850509588019591880191600101612265565b5050505050505060808501519150601f198482030160a0850152610c908183612193565b8181038181111561055557610555612094565b61ffff60f01b8660f01b1681528460028201528360228201526001600160601b03198360601b1660428201525f825161233e816056850160208701611b2c565b919091016056019695505050505050565b5f6001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a0606084015261238860a0840186611b4e565b9150808416608084015250969550505050505056fea2646970667358221220625865c93815b49b526d958ed729c6f8a18d131f5e1e9b5a719630de3a52f9e164736f6c63430008180033",
}

// SolverNetOutboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SolverNetOutboxMetaData.ABI instead.
var SolverNetOutboxABI = SolverNetOutboxMetaData.ABI

// SolverNetOutboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolverNetOutboxMetaData.Bin instead.
var SolverNetOutboxBin = SolverNetOutboxMetaData.Bin

// DeploySolverNetOutbox deploys a new Ethereum contract, binding an instance of SolverNetOutbox to it.
func DeploySolverNetOutbox(auth *bind.TransactOpts, backend bind.ContractBackend, _mailbox common.Address) (common.Address, *types.Transaction, *SolverNetOutbox, error) {
	parsed, err := SolverNetOutboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolverNetOutboxBin), backend, _mailbox)
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

// PACKAGEVERSION is a free data retrieval call binding the contract method 0x93c44847.
//
// Solidity: function PACKAGE_VERSION() view returns(string)
func (_SolverNetOutbox *SolverNetOutboxCaller) PACKAGEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "PACKAGE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PACKAGEVERSION is a free data retrieval call binding the contract method 0x93c44847.
//
// Solidity: function PACKAGE_VERSION() view returns(string)
func (_SolverNetOutbox *SolverNetOutboxSession) PACKAGEVERSION() (string, error) {
	return _SolverNetOutbox.Contract.PACKAGEVERSION(&_SolverNetOutbox.CallOpts)
}

// PACKAGEVERSION is a free data retrieval call binding the contract method 0x93c44847.
//
// Solidity: function PACKAGE_VERSION() view returns(string)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) PACKAGEVERSION() (string, error) {
	return _SolverNetOutbox.Contract.PACKAGEVERSION(&_SolverNetOutbox.CallOpts)
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

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxCaller) InterchainSecurityModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "interchainSecurityModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxSession) InterchainSecurityModule() (common.Address, error) {
	return _SolverNetOutbox.Contract.InterchainSecurityModule(&_SolverNetOutbox.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) InterchainSecurityModule() (common.Address, error) {
	return _SolverNetOutbox.Contract.InterchainSecurityModule(&_SolverNetOutbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SolverNetOutbox *SolverNetOutboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SolverNetOutbox *SolverNetOutboxSession) LocalDomain() (uint32, error) {
	return _SolverNetOutbox.Contract.LocalDomain(&_SolverNetOutbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) LocalDomain() (uint32, error) {
	return _SolverNetOutbox.Contract.LocalDomain(&_SolverNetOutbox.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxCaller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetOutbox.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxSession) Mailbox() (common.Address, error) {
	return _SolverNetOutbox.Contract.Mailbox(&_SolverNetOutbox.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_SolverNetOutbox *SolverNetOutboxCallerSession) Mailbox() (common.Address, error) {
	return _SolverNetOutbox.Contract.Mailbox(&_SolverNetOutbox.CallOpts)
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

// SetInboxes is a paid mutator transaction binding the contract method 0x82e51365.
//
// Solidity: function setInboxes(uint64[] chainIds, (address,uint8)[] configs) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactor) SetInboxes(opts *bind.TransactOpts, chainIds []uint64, configs []ISolverNetOutboxInboxConfig) (*types.Transaction, error) {
	return _SolverNetOutbox.contract.Transact(opts, "setInboxes", chainIds, configs)
}

// SetInboxes is a paid mutator transaction binding the contract method 0x82e51365.
//
// Solidity: function setInboxes(uint64[] chainIds, (address,uint8)[] configs) returns()
func (_SolverNetOutbox *SolverNetOutboxSession) SetInboxes(chainIds []uint64, configs []ISolverNetOutboxInboxConfig) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.SetInboxes(&_SolverNetOutbox.TransactOpts, chainIds, configs)
}

// SetInboxes is a paid mutator transaction binding the contract method 0x82e51365.
//
// Solidity: function setInboxes(uint64[] chainIds, (address,uint8)[] configs) returns()
func (_SolverNetOutbox *SolverNetOutboxTransactorSession) SetInboxes(chainIds []uint64, configs []ISolverNetOutboxInboxConfig) (*types.Transaction, error) {
	return _SolverNetOutbox.Contract.SetInboxes(&_SolverNetOutbox.TransactOpts, chainIds, configs)
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

// SolverNetOutboxHookSetIterator is returned from FilterHookSet and is used to iterate over the raw logs and unpacked data for HookSet events raised by the SolverNetOutbox contract.
type SolverNetOutboxHookSetIterator struct {
	Event *SolverNetOutboxHookSet // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxHookSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxHookSet)
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
		it.Event = new(SolverNetOutboxHookSet)
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
func (it *SolverNetOutboxHookSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxHookSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxHookSet represents a HookSet event raised by the SolverNetOutbox contract.
type SolverNetOutboxHookSet struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookSet is a free log retrieval operation binding the contract event 0x4eab7b127c764308788622363ad3e9532de3dfba7845bd4f84c125a22544255a.
//
// Solidity: event HookSet(address _hook)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterHookSet(opts *bind.FilterOpts) (*SolverNetOutboxHookSetIterator, error) {

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "HookSet")
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxHookSetIterator{contract: _SolverNetOutbox.contract, event: "HookSet", logs: logs, sub: sub}, nil
}

// WatchHookSet is a free log subscription operation binding the contract event 0x4eab7b127c764308788622363ad3e9532de3dfba7845bd4f84c125a22544255a.
//
// Solidity: event HookSet(address _hook)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchHookSet(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxHookSet) (event.Subscription, error) {

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "HookSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxHookSet)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "HookSet", log); err != nil {
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

// ParseHookSet is a log parse operation binding the contract event 0x4eab7b127c764308788622363ad3e9532de3dfba7845bd4f84c125a22544255a.
//
// Solidity: event HookSet(address _hook)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseHookSet(log types.Log) (*SolverNetOutboxHookSet, error) {
	event := new(SolverNetOutboxHookSet)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "HookSet", log); err != nil {
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
	ChainId  uint64
	Inbox    common.Address
	Provider uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInboxSet is a free log retrieval operation binding the contract event 0x41d1e58dad7679eeaefbd227553cdede676d33bf46808e7b54b2550a4e44057e.
//
// Solidity: event InboxSet(uint64 indexed chainId, address indexed inbox, uint8 indexed provider)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterInboxSet(opts *bind.FilterOpts, chainId []uint64, inbox []common.Address, provider []uint8) (*SolverNetOutboxInboxSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "InboxSet", chainIdRule, inboxRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxInboxSetIterator{contract: _SolverNetOutbox.contract, event: "InboxSet", logs: logs, sub: sub}, nil
}

// WatchInboxSet is a free log subscription operation binding the contract event 0x41d1e58dad7679eeaefbd227553cdede676d33bf46808e7b54b2550a4e44057e.
//
// Solidity: event InboxSet(uint64 indexed chainId, address indexed inbox, uint8 indexed provider)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchInboxSet(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxInboxSet, chainId []uint64, inbox []common.Address, provider []uint8) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "InboxSet", chainIdRule, inboxRule, providerRule)
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

// ParseInboxSet is a log parse operation binding the contract event 0x41d1e58dad7679eeaefbd227553cdede676d33bf46808e7b54b2550a4e44057e.
//
// Solidity: event InboxSet(uint64 indexed chainId, address indexed inbox, uint8 indexed provider)
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

// SolverNetOutboxIsmSetIterator is returned from FilterIsmSet and is used to iterate over the raw logs and unpacked data for IsmSet events raised by the SolverNetOutbox contract.
type SolverNetOutboxIsmSetIterator struct {
	Event *SolverNetOutboxIsmSet // Event containing the contract specifics and raw log

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
func (it *SolverNetOutboxIsmSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetOutboxIsmSet)
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
		it.Event = new(SolverNetOutboxIsmSet)
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
func (it *SolverNetOutboxIsmSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetOutboxIsmSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetOutboxIsmSet represents a IsmSet event raised by the SolverNetOutbox contract.
type SolverNetOutboxIsmSet struct {
	Ism common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIsmSet is a free log retrieval operation binding the contract event 0xc47cbcc588c67679e52261c45cc315e56562f8d0ccaba16facb9093ff9498799.
//
// Solidity: event IsmSet(address _ism)
func (_SolverNetOutbox *SolverNetOutboxFilterer) FilterIsmSet(opts *bind.FilterOpts) (*SolverNetOutboxIsmSetIterator, error) {

	logs, sub, err := _SolverNetOutbox.contract.FilterLogs(opts, "IsmSet")
	if err != nil {
		return nil, err
	}
	return &SolverNetOutboxIsmSetIterator{contract: _SolverNetOutbox.contract, event: "IsmSet", logs: logs, sub: sub}, nil
}

// WatchIsmSet is a free log subscription operation binding the contract event 0xc47cbcc588c67679e52261c45cc315e56562f8d0ccaba16facb9093ff9498799.
//
// Solidity: event IsmSet(address _ism)
func (_SolverNetOutbox *SolverNetOutboxFilterer) WatchIsmSet(opts *bind.WatchOpts, sink chan<- *SolverNetOutboxIsmSet) (event.Subscription, error) {

	logs, sub, err := _SolverNetOutbox.contract.WatchLogs(opts, "IsmSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetOutboxIsmSet)
				if err := _SolverNetOutbox.contract.UnpackLog(event, "IsmSet", log); err != nil {
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

// ParseIsmSet is a log parse operation binding the contract event 0xc47cbcc588c67679e52261c45cc315e56562f8d0ccaba16facb9093ff9498799.
//
// Solidity: event IsmSet(address _ism)
func (_SolverNetOutbox *SolverNetOutboxFilterer) ParseIsmSet(log types.Log) (*SolverNetOutboxIsmSet, error) {
	event := new(SolverNetOutboxIsmSet)
	if err := _SolverNetOutbox.contract.UnpackLog(event, "IsmSet", log); err != nil {
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
