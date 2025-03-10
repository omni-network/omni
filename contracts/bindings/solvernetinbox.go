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
type IERC7683FillInstruction struct {
	DestinationChainId uint64
	DestinationSettler [32]byte
	OriginData         []byte
}

// IERC7683OnchainCrossChainOrder is an auto generated low-level Go binding around an user-defined struct.
type IERC7683OnchainCrossChainOrder struct {
	FillDeadline  uint32
	OrderDataType [32]byte
	OrderData     []byte
}

// IERC7683Output is an auto generated low-level Go binding around an user-defined struct.
type IERC7683Output struct {
	Token     [32]byte
	Amount    *big.Int
	Recipient [32]byte
	ChainId   *big.Int
}

// IERC7683ResolvedCrossChainOrder is an auto generated low-level Go binding around an user-defined struct.
type IERC7683ResolvedCrossChainOrder struct {
	User             common.Address
	OriginChainId    *big.Int
	OpenDeadline     uint32
	FillDeadline     uint32
	OrderId          [32]byte
	MaxSpent         []IERC7683Output
	MinReceived      []IERC7683Output
	FillInstructions []IERC7683FillInstruction
}

// ISolverNetInboxOrderState is an auto generated low-level Go binding around an user-defined struct.
type ISolverNetInboxOrderState struct {
	Status       uint8
	RejectReason uint8
	Timestamp    uint32
	UpdatedBy    common.Address
}

// SolverNetInboxMetaData contains all meta data concerning the SolverNetInbox contract.
var SolverNetInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"close\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestOrderIdByStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISolverNetInbox.Status\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOffset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrder\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"resolved\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"state\",\"type\":\"tuple\",\"internalType\":\"structISolverNetInbox.OrderState\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISolverNetInbox.Status\"},{\"name\":\"rejectReason\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedBy\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"offset\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserNonce\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"markFilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"open\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseClose\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseOpen\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reject\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOutboxes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboxes\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validate\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Closed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Filled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Open\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"resolvedOrder\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboxSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Rejected\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCallTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenseAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenseToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFillDeadline\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMissingCalls\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNativeDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrderData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrderTypehash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReason\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderStillValid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PortalPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongFillHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongSourceChain\",\"inputs\":[]}]",
	Bin: "0x60a060405234801562000010575f80fd5b5063ffffffff60643b1615620000965760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156200007e575060408051601f3d908101601f191682019092526200007b9181019062000112565b60015b6200008d57436080526200009b565b6080526200009b565b436080525b620000a5620000ab565b6200012a565b63409feecd1980546001811615620000ca5763f92ee8a95f526004601cfd5b6001600160401b03808260011c146200010d578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b5f6020828403121562000123575f80fd5b5051919050565b608051613b1a620001435f395f6105c00152613b1a5ff3fe6080604052600436106101fc575f3560e01c8063715018a611610113578063d9e8407c1161009d578063eb6ad24e1161006d578063eb6ad24e146105e2578063f04e283e146105ff578063f2fde38b14610612578063f904d28514610625578063fee81cf414610644575f80fd5b8063d9e8407c1461055e578063db3ea5531461057d578063e917a9621461059c578063eae4c19f146105af575f80fd5b80638da5cb5b116100e35780638da5cb5b146104c957806396c144f0146104e1578063af844c9b14610500578063c0c53b8b1461051f578063d71183511461053e575f80fd5b8063715018a61461045257806374eeb8471461045a578063792aec5c1461048b5780637cac41a6146104aa575f80fd5b806339acf9f111610194578063514e62fc11610164578063514e62fc1461039457806354d1f13d146103c95780635778472a146103d15780635e0aec95146103ff5780636834e3a81461041e575f80fd5b806339acf9f11461030057806339c79e0c1461033657806341b477dd146103555780634a4ee7b114610381575f80fd5b806325692962116101cf578063256929621461027b57806329d54233146102835780632d622343146102b05780632de94807146102cf575f80fd5b806304a873ab14610200578063183a4f6e146102215780631c10893f146102345780631cd64df414610247575b5f80fd5b34801561020b575f80fd5b5061021f61021a366004612e1e565b610675565b005b61021f61022f366004612e84565b6107c9565b61021f610242366004612eaf565b6107d6565b348015610252575f80fd5b50610266610261366004612eaf565b6107ec565b60405190151581526020015b60405180910390f35b61021f61080a565b34801561028e575f80fd5b506102a261029d366004612ed9565b610856565b604051908152602001610272565b3480156102bb575f80fd5b5061021f6102ca366004612ef4565b610879565b3480156102da575f80fd5b506102a26102e9366004612ed9565b638b78c6d8600c9081525f91909152602090205490565b34801561030b575f80fd5b505f5461031e906001600160a01b031681565b6040516001600160a01b039091168152602001610272565b348015610341575f80fd5b5061021f610350366004612e84565b610b14565b348015610360575f80fd5b5061037461036f366004612f2a565b610e74565b6040516102729190613124565b61021f61038f366004612eaf565b610ec3565b34801561039f575f80fd5b506102666103ae366004612eaf565b638b78c6d8600c9081525f9290925260209091205416151590565b61021f610ed5565b3480156103dc575f80fd5b506103f06103eb366004612e84565b610f0e565b6040516102729392919061314a565b34801561040a575f80fd5b506102a2610419366004612eaf565b610fe8565b348015610429575f80fd5b506102a2610438366004612ed9565b6001600160a01b03165f908152600a602052604090205490565b61021f610ffa565b348015610465575f80fd5b505f5461047990600160a01b900460ff1681565b60405160ff9091168152602001610272565b348015610496575f80fd5b5061021f6104a53660046131ce565b61100d565b3480156104b5575f80fd5b5061021f6104c43660046131ce565b61102f565b3480156104d4575f80fd5b50638b78c6d8195461031e565b3480156104ec575f80fd5b5061021f6104fb3660046131e9565b61106f565b34801561050b575f80fd5b506102a261051a366004613217565b6111dc565b34801561052a575f80fd5b5061021f610539366004613235565b611218565b348015610549575f80fd5b5060025461047990600160f81b900460ff1681565b348015610569575f80fd5b50610266610578366004612f2a565b6112a7565b348015610588575f80fd5b5061021f610597366004613272565b6112ba565b61021f6105aa366004612f2a565b611441565b3480156105ba575f80fd5b506102a27f000000000000000000000000000000000000000000000000000000000000000081565b3480156105ed575f80fd5b506002546001600160f81b03166102a2565b61021f61060d366004612ed9565b6115a0565b61021f610620366004612ed9565b6115da565b348015610630575f80fd5b5061021f61063f3660046131ce565b611600565b34801561064f575f80fd5b506102a261065e366004612ed9565b63389a75e1600c9081525f91909152602090205490565b61067d611622565b82811461069d57604051634ec4810560e11b815260040160405180910390fd5b5f5b838110156107c2578282828181106106b9576106b961329a565b90506020020160208101906106ce9190612ed9565b60035f8787858181106106e3576106e361329a565b90506020020160208101906106f891906132c2565b6001600160401b0316815260208101919091526040015f2080546001600160a01b0319166001600160a01b03929092169190911790558282828181106107405761074061329a565b90506020020160208101906107559190612ed9565b6001600160a01b03168585838181106107705761077061329a565b905060200201602081019061078591906132c2565b6001600160401b03167ff730978310b4a2a0e6c673324d737afdb93d0afefed14a3d061b60f66e31f4e360405160405180910390a360010161069f565b5050505050565b6107d3338261163c565b50565b6107de611622565b6107e88282611647565b5050565b638b78c6d8600c9081525f8390526020902054811681145b92915050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f80a250565b6001600160a01b0381165f908152600a6020526040812054610804908390611653565b5f5460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa1580156108bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e191906133af565b8051600180546020938401516001600160401b039384166001600160e01b031990921691909117600160401b6001600160a01b0392831602179091555f868152600484526040808220815160608101835290549384168152600160a01b840490941684860152600160e01b90920463ffffffff168383015286815260089093528083208151608081019092528054929392829060ff16600581111561098857610988613136565b600581111561099957610999613136565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015290506001815160058111156109e8576109e8613136565b14610a0657604051635d12a4a360e11b815260040160405180910390fd5b60208201516001546001600160401b03908116911614610a3957604051633687f39960e21b815260040160405180910390fd5b6001546001600160401b0381165f908152600360205260409020546001600160a01b03908116600160401b9092041614610a85576040516282b42960e81b815260040160405180910390fd5b610a8e85611695565b8414610aad57604051631f53eaed60e21b815260040160405180910390fd5b610aba8560045f8661193d565b610ac5856004611ab7565b826001600160a01b031684867fa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc60405160405180910390a45050600180546001600160e01b0319169055505050565b6002545f80516020613ac583398151915290600160f81b900460ff168015610bd55760ff81166001148015610b5557505f80516020613aa583398151915282145b15610b7357604051631309a56360e01b815260040160405180910390fd5b60ff81166002148015610b9257505f80516020613ac583398151915282145b15610bb057604051631309a56360e01b815260040160405180910390fd5b60021960ff821601610bd55760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd212685403610bf35763ab143c065f526004601cfd5b3068929eee149b4bd21268555f838152600860205260408082208151608081019092528054829060ff166005811115610c2e57610c2e613136565b6005811115610c3f57610c3f613136565b8152905460ff61010082041660208084019190915263ffffffff62010000830481166040808601919091526001600160a01b03600160301b90940484166060958601525f8a815260048452819020815195860182525493841685526001600160401b03600160a01b85041692850192909252600160e01b90920490911690820152909150600182516005811115610cd857610cd8613136565b14610cf657604051635d12a4a360e11b815260040160405180910390fd5b80516001600160a01b03163314610d1f576040516282b42960e81b815260040160405180910390fd5b5f5460208201516040516308c3569160e31b81527ffeccba1cfc4544bf9cd83b76f36ae5c464750b6c43f682e26744ee21ec31fc1e60048201526001600160401b0390911660248201526001600160a01b039091169063461ab48890604401602060405180830381865afa158015610d99573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dbd91906133ee565b15610ddb57604051630c2e605760e11b815260040160405180910390fd5b42615460826040015163ffffffff16610df4919061341d565b10610e12576040516321bb6b2160e11b815260040160405180910390fd5b610e1f8560035f3361193d565b610e2c85825f0151611b4e565b610e37856003611ab7565b60405185907f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79905f90a250503868929eee149b4bd2126855505050565b610e7c612c55565b5f610e8683611be7565b8051516001600160a01b0381165f908152600a602052604090205491925090610ebb908390610eb6908490611653565b611eef565b949350505050565b610ecb611622565b6107e8828261163c565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f80a2565b610f16612c55565b604080516080810182525f80825260208201819052918101829052606081018290529080610f4385611f6e565b9050610f4f8186611eef565b5f868152600860209081526040808320600990925291829020548251608081019093528154919290918390829060ff166005811115610f9057610f90613136565b6005811115610fa157610fa1613136565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015292989297509550909350505050565b5f610ff38383611653565b9392505050565b611002611622565b61100b5f6121c4565b565b600161101881612201565b6107e85f80516020613ac583398151915283612232565b600161103a81612201565b8161105257600280546001600160f81b031690555050565b600280546001600160f81b0316600360f81b17905560035b505050565b3068929eee149b4bd21268540361108d5763ab143c065f526004601cfd5b3068929eee149b4bd21268555f828152600860205260408082208151608081019092528054829060ff1660058111156110c8576110c8613136565b60058111156110d9576110d9613136565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b0316606090910152905060048151600581111561112857611128613136565b146111465760405163789bae3560e01b815260040160405180910390fd5b60608101516001600160a01b03163314611172576040516282b42960e81b815260040160405180910390fd5b61117f8360055f3361193d565b6111898383611b4e565b611194836005611ab7565b6040516001600160a01b03831690339085907f8428df912f4f2125b442b488df9c7260cb607246895bcd29f262ecca090b1538905f90a4503868929eee149b4bd21268555050565b5f600b5f8360058111156111f2576111f2613136565b600581111561120357611203613136565b81526020019081526020015f20549050919050565b63409feecd19805460038255801561124e5760018160011c14303b106112455763f92ee8a95f526004601cfd5b818160ff1b1b91505b50611258846122eb565b611263836001611647565b61126c82612326565b80156112a1576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b50505050565b5f6112b182611be7565b50600192915050565b60016112c5816123c8565b3068929eee149b4bd2126854036112e35763ab143c065f526004601cfd5b3068929eee149b4bd21268555f838152600860205260408082208151608081019092528054829060ff16600581111561131e5761131e613136565b600581111561132f5761132f613136565b81529054610100810460ff908116602084015262010000820463ffffffff166040840152600160301b9091046001600160a01b031660609092019190915290915083165f03611391576040516337b89b9360e21b815260040160405180910390fd5b6001815160058111156113a6576113a6613136565b146113c457604051635d12a4a360e11b815260040160405180910390fd5b6113d1846002853361193d565b5f848152600460205260409020546113f39085906001600160a01b0316611b4e565b6113fe846002611ab7565b60405160ff841690339086907f21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f905f90a4503868929eee149b4bd2126855505050565b6002545f80516020613aa583398151915290600160f81b900460ff1680156115025760ff8116600114801561148257505f80516020613aa583398151915282145b156114a057604051631309a56360e01b815260040160405180910390fd5b60ff811660021480156114bf57505f80516020613ac583398151915282145b156114dd57604051631309a56360e01b815260040160405180910390fd5b60021960ff8216016115025760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd2126854036115205763ab143c065f526004601cfd5b3068929eee149b4bd21268555f61153684611be7565b905061154581602001516123ec565b5f61154f82612451565b905080608001517fa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c826040516115859190613124565b60405180910390a250503868929eee149b4bd2126855505050565b6115a8611622565b63389a75e1600c52805f526020600c2080544211156115ce57636f5e88185f526004601cfd5b5f90556107d3816121c4565b6115e2611622565b8060601b6115f757637448fbae5f526004601cfd5b6107d3816121c4565b600161160b81612201565b6107e85f80516020613aa583398151915283612232565b638b78c6d81954331461100b576382b429005f526004601cfd5b6107e882825f612697565b6107e882826001612697565b604080516001600160a01b03841660208201529081018290524660608201525f9060800160405160208183030381529060405280519060200120905092915050565b5f818152600460209081526040808320815160608101835290546001600160a01b0381168252600160a01b81046001600160401b031682850152600160e01b900463ffffffff1681830152848452600683528184208054835181860281018601909452808452919385939290849084015b82821015611802575f848152602090819020604080516080810182526003860290920180546001600160a01b0381168452600160a01b900460e01b6001600160e01b0319169383019390935260018301549082015260028201805491929160608401919061177390613430565b80601f016020809104026020016040519081016040528092919081815260200182805461179f90613430565b80156117ea5780601f106117c1576101008083540402835291602001916117ea565b820191905f5260205f20905b8154815290600101906020018083116117cd57829003601f168201915b50505050508152505081526020019060010190611706565b5050505090505f60075f8681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b8282101561189c575f848152602090819020604080516060810182526002860290920180546001600160a01b03908116845260019182015490811684860152600160a01b90046001600160601b0316918301919091529083529092019101611839565b5050505090505f6040518060a00160405280466001600160401b0316815260200185602001516001600160401b03168152602001856040015163ffffffff16815260200184815260200183815250905085816040516020016118fe91906134c2565b60408051601f198184030181529082905261191c92916020016135b5565b60405160208183030381529060405280519060200120945050505050919050565b5f848152600860205260409081902054815160808101909252610100900460ff16908085600581111561197257611972613136565b81526020015f8560ff16116119875782611989565b845b60ff16815263ffffffff42166020808301919091526001600160a01b0385166040928301525f88815260089091522081518154829060ff191660018360058111156119d6576119d6613136565b02179055506020820151815460408401516060909401516001600160a01b0316600160301b026601000000000000600160d01b031963ffffffff909516620100000265ffffffff00001960ff909416610100029390931665ffffffffff001990921691909117919091179290921691909117905584600b5f866005811115611a6057611a60613136565b6005811115611a7157611a71613136565b815260208101919091526040015f20556001846005811115611a9557611a95613136565b036107c257611aa26126ee565b5f868152600960205260409020555050505050565b6001816005811115611acb57611acb613136565b03611ad4575050565b6004816005811115611ae857611ae8613136565b14611afc575f828152600560205260408120555b6005816005811115611b1057611b10613136565b146107e8575f82815260046020908152604080832083905560069091528120611b3891612cad565b5f8281526007602052604081206107e891612ccb565b5f828152600560209081526040918290208251808401909352546001600160a01b0381168352600160a01b90046001600160601b03169082018190521561106a5780516001600160a01b0316611bc057602081015161106a906001600160a01b038416906001600160601b031661272d565b6020810151815161106a916001600160a01b039091169084906001600160601b0316612746565b611bef612ce9565b42611bfd60208401846135cd565b63ffffffff1611611c215760405163582e388960e01b815260040160405180910390fd5b7f2e7de755ca70cb933dc80103af16cc3303580e5712f1a8927d6461441e99a1e6826020013514611c6557604051636aea87f360e01b815260040160405180910390fd5b611c7260408301836135f0565b90505f03611c935760405163a342e7d960e01b815260040160405180910390fd5b5f611ca160408401846135f0565b810190611cae919061389a565b80519091506001600160a01b0316611cc4573381525b60208101516001600160401b03161580611cea57504681602001516001600160401b0316145b15611d0857604051633d23e4d160e11b815260040160405180910390fd5b6040805160608101825282516001600160a01b031681526020808401516001600160401b0316818301525f92820190611d43908701876135cd565b63ffffffff16905260608301518051919250905f03611d7557604051639cc71f7d60e01b815260040160405180910390fd5b805160201015611d9857604051634ec4810560e11b815260040160405180910390fd5b5f5b8151811015611df5575f828281518110611db657611db661329a565b602090810291909101015180519091506001600160a01b0316611dec5760405163017ab86160e21b815260040160405180910390fd5b50600101611d9a565b506080830151805160201015611e1e57604051634ec4810560e11b815260040160405180910390fd5b5f5b8151811015611ec6575f6001600160a01b0316828281518110611e4557611e4561329a565b6020026020010151602001516001600160a01b031603611e785760405163027dcfa160e31b815260040160405180910390fd5b818181518110611e8a57611e8a61329a565b6020026020010151604001516001600160601b03165f03611ebe5760405163a0ce339960e01b815260040160405180910390fd5b600101611e20565b506040805160808101825293845293840151602084015292820152606081019190915292915050565b611ef7612c55565b82515f611f0385612790565b90505f611f0f866129ec565b90505f611f1b87612af0565b604080516101008101825286516001600160a01b031681524660208201525f8183015295015163ffffffff1660608601526080850187905260a08501939093525060c083015260e0820152905092915050565b611f76612ce9565b604080515f8481526004602090815283822060e084018552546001600160a01b0380821660808601908152600160a01b8084046001600160401b031660a0880152600160e01b90930463ffffffff1660c087015285528784526005835285842086518088018852905491821681529190046001600160601b031681830152818401528582526006815283822080548551818402810184018752818152949586019493919290919084015b8282101561211c575f848152602090819020604080516080810182526003860290920180546001600160a01b0381168452600160a01b900460e01b6001600160e01b0319169383019390935260018301549082015260028201805491929160608401919061208d90613430565b80601f01602080910402602001604051908101604052809291908181526020018280546120b990613430565b80156121045780601f106120db57610100808354040283529160200191612104565b820191905f5260205f20905b8154815290600101906020018083116120e757829003601f168201915b50505050508152505081526020019060010190612020565b50505050815260200160075f8581526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156121b8575f848152602090819020604080516060810182526002860290920180546001600160a01b03908116845260019182015490811684860152600160a01b90046001600160601b0316918301919091529083529092019101612155565b50505091525092915050565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a355565b638b78c6d8195433146107d357638b78c6d8600c52335f52806020600c2054166107d3576382b429005f526004601cfd5b600254600160f81b900460ff1660021981016122615760405163aaae8ef760e01b815260040160405180910390fd5b5f5f80516020613aa5833981519152841461227d576002612280565b60015b905082612296578060ff168260ff1614156122a0565b8060ff168260ff16145b156122be57604051631309a56360e01b815260040160405180910390fd5b826122c9575f6122cb565b805b6002601f6101000a81548160ff021916908360ff16021790555050505050565b6001600160a01b0316638b78c6d819819055805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b0381166123755760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b604482015260640160405180910390fd5b5f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52335f52806020600c2054166107d3576382b429005f526004601cfd5b80516001600160a01b03166124285780602001516001600160601b031634146107d35760405163036f810f60e41b815260040160405180910390fd5b602081015181516107d3916001600160a01b0390911690339030906001600160601b0316612bfc565b612459612c55565b8151516001600160a01b0381165f908152600a60205260408120805461248f91849190846124868361395c565b91905055611653565b905061249b8482611eef565b84515f83815260046020908152604080832084518154848701519684015163ffffffff16600160e01b026001600160e01b036001600160401b03909816600160a01b9081026001600160e01b03199093166001600160a01b0394851617929092179790971696909617909155828a015160058452918420825192909301516001600160601b03169094029316929092179091559093505b8460400151518110156125e7575f8281526006602052604090819020908601518051839081106125645761256461329a565b6020908102919091018101518254600181810185555f94855293839020825160039092020180549383015160e01c600160a01b026001600160c01b03199094166001600160a01b0390921691909117929092178255604081015192820192909255606082015160028201906125d990826139b8565b505050806001019050612532565b505f5b846060015151811015612682575f828152600760205260409020606086015180518390811061261b5761261b61329a565b6020908102919091018101518254600181810185555f94855293839020825160029092020180546001600160a01b039283166001600160a01b0319909116178155928201516040909201516001600160601b0316600160a01b0291161790820155016125ea565b506126908160015f3361193d565b5050919050565b638b78c6d8600c52825f526020600c208054838117836126b8575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe265f80a3505050505050565b600280545f91908290612709906001600160f81b0316613a77565b82546001600160f81b039182166101009390930a8381029202191617909155919050565b5f385f3884865af16107e85763b12d13eb5f526004601cfd5b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f51141661278657803d853b151710612786576390b8ec185f526004601cfd5b505f603452505050565b80516040820151606083810151909291905f805b8351811015612806575f8482815181106127c0576127c061329a565b60200260200101516040015111156127fe578381815181106127e4576127e461329a565b602002602001015160400151826127fb919061341d565b91505b6001016127a4565b505f808211612816578251612823565b825161282390600161341d565b6001600160401b0381111561283a5761283a6132dd565b60405190808252806020026020018201604052801561288a57816020015b604080516080810182525f8082526020808301829052928201819052606082015282525f199092019101816128585790505b5090505f5b835181101561296b5760405180608001604052806128d98684815181106128b8576128b861329a565b6020026020010151602001516001600160a01b03166001600160a01b031690565b81526020018583815181106128f0576128f061329a565b6020908102919091018101516040908101516001600160601b03168352898201516001600160401b03165f9081526003835220549101906001600160a01b0316815260200187602001516001600160401b03168152508282815181106129585761295861329a565b602090810291909101015260010161288f565b5081156129e257604080516080810182525f8082526020808301869052888101516001600160401b03168252600390528290205490918201906001600160a01b0316815260200186602001516001600160401b0316815250818451815181106129d6576129d661329a565b60200260200101819052505b9695505050505050565b60605f826020015190505f8082602001516001600160601b031611612a11575f612a14565b60015b60ff166001600160401b03811115612a2e57612a2e6132dd565b604051908082528060200260200182016040528015612a7e57816020015b604080516080810182525f8082526020808301829052928201819052606082015282525f19909201910181612a4c5790505b5060208301519091506001600160601b031615610ff3576040805160808101825283516001600160a01b031681526020808501516001600160601b0316908201525f918101829052466060820152825190918391612ade57612ade61329a565b60200260200101819052509392505050565b80516040808301516060808501518351600180825281860190955291949390915f91816020015b60408051606080820183525f808352602083015291810191909152815260200190600190039081612b175750506040805160608082018352602088810180516001600160401b039081168552815181165f90815260038452869020546001600160a01b031683860152855160a0810187524682168152915116818301528985015163ffffffff1681860152918201889052608082018790528351949550919392840192612bc492016134c2565b604051602081830303815290604052815250815f81518110612be857612be861329a565b602090810291909101015295945050505050565b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f511416612c4757803d873b151710612c4757637939f4245f526004601cfd5b505f60605260405250505050565b6040518061010001604052805f6001600160a01b031681526020015f81526020015f63ffffffff1681526020015f63ffffffff1681526020015f80191681526020016060815260200160608152602001606081525090565b5080545f8255600302905f5260205f20908101906107d39190612d34565b5080545f8255600202905f5260205f20908101906107d39190612d6c565b6040805160e0810182525f6080820181815260a0830182905260c083018290528252825180840184528181526020808201929092529082015260609181018290528181019190915290565b80821115612d685780546001600160c01b03191681555f60018201819055612d5f6002830182612d91565b50600301612d34565b5090565b5b80821115612d685780546001600160a01b03191681555f6001820155600201612d6d565b508054612d9d90613430565b5f825580601f10612dac575050565b601f0160209004905f5260205f20908101906107d391905b80821115612d68575f8155600101612dc4565b5f8083601f840112612de7575f80fd5b5081356001600160401b03811115612dfd575f80fd5b6020830191508360208260051b8501011115612e17575f80fd5b9250929050565b5f805f8060408587031215612e31575f80fd5b84356001600160401b0380821115612e47575f80fd5b612e5388838901612dd7565b90965094506020870135915080821115612e6b575f80fd5b50612e7887828801612dd7565b95989497509550505050565b5f60208284031215612e94575f80fd5b5035919050565b6001600160a01b03811681146107d3575f80fd5b5f8060408385031215612ec0575f80fd5b8235612ecb81612e9b565b946020939093013593505050565b5f60208284031215612ee9575f80fd5b8135610ff381612e9b565b5f805f60608486031215612f06575f80fd5b83359250602084013591506040840135612f1f81612e9b565b809150509250925092565b5f60208284031215612f3a575f80fd5b81356001600160401b03811115612f4f575f80fd5b820160608185031215610ff3575f80fd5b5f815180845260208085019450602084015f5b83811015612faf578151805188528381015184890152604080820151908901526060908101519088015260809096019590820190600101612f73565b509495945050505050565b5f81518084525f5b81811015612fde57602081850181015186830182015201612fc2565b505f602082860101526020601f19601f83011685010191505092915050565b5f82825180855260208086019550808260051b8401018186015f5b8481101561306f57858303601f19018952815180516001600160401b03168452848101518585015260409081015160609185018290529061305b81860183612fba565b9a86019a9450505090830190600101613018565b5090979650505050505050565b5f61010060018060a01b0383511684526020830151602085015260408301516130ad604086018263ffffffff169052565b5060608301516130c5606086018263ffffffff169052565b506080830151608085015260a08301518160a08601526130e782860182612f60565b91505060c083015184820360c08601526131018282612f60565b91505060e083015184820360e086015261311b8282612ffd565b95945050505050565b602081525f610ff3602083018461307c565b634e487b7160e01b5f52602160045260245ffd5b60c081525f61315c60c083018661307c565b905083516006811061317c57634e487b7160e01b5f52602160045260245ffd5b60208381019190915284015160ff1660408084019190915284015163ffffffff16606080840191909152909301516001600160a01b0316608082015260a00152919050565b80151581146107d3575f80fd5b5f602082840312156131de575f80fd5b8135610ff3816131c1565b5f80604083850312156131fa575f80fd5b82359150602083013561320c81612e9b565b809150509250929050565b5f60208284031215613227575f80fd5b813560068110610ff3575f80fd5b5f805f60608486031215613247575f80fd5b833561325281612e9b565b9250602084013561326281612e9b565b91506040840135612f1f81612e9b565b5f8060408385031215613283575f80fd5b82359150602083013560ff8116811461320c575f80fd5b634e487b7160e01b5f52603260045260245ffd5b6001600160401b03811681146107d3575f80fd5b5f602082840312156132d2575f80fd5b8135610ff3816132ae565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715613313576133136132dd565b60405290565b604051608081016001600160401b0381118282101715613313576133136132dd565b604051606081016001600160401b0381118282101715613313576133136132dd565b60405160a081016001600160401b0381118282101715613313576133136132dd565b604051601f8201601f191681016001600160401b03811182821017156133a7576133a76132dd565b604052919050565b5f604082840312156133bf575f80fd5b6133c76132f1565b82516133d2816132ae565b815260208301516133e281612e9b565b60208201529392505050565b5f602082840312156133fe575f80fd5b8151610ff3816131c1565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561080457610804613409565b600181811c9082168061344457607f821691505b60208210810361346257634e487b7160e01b5f52602260045260245ffd5b50919050565b5f815180845260208085019450602084015f5b83811015612faf57815180516001600160a01b0390811689528482015116848901526040908101516001600160601b0316908801526060909601959082019060010161347b565b5f602080835260c083016001600160401b0380865116838601528286015160408282166040880152604088015192506060915063ffffffff8316606088015260608801519250608060a0608089015284845180875260e08a01915060e08160051b8b0101965087860195505f5b81811015613591578a880360df19018352865180516001600160a01b03168952898101516001600160e01b0319168a8a015285810151868a015286015186890185905261357e858a0182612fba565b985050958801959188019160010161352f565b5050505050505060808501519150601f198482030160a085015261311b8183613468565b828152604060208201525f610ebb6040830184612fba565b5f602082840312156135dd575f80fd5b813563ffffffff81168114610ff3575f80fd5b5f808335601e19843603018112613605575f80fd5b8301803591506001600160401b0382111561361e575f80fd5b602001915036819003821315612e17575f80fd5b80356001600160601b0381168114613648575f80fd5b919050565b5f6040828403121561365d575f80fd5b6136656132f1565b9050813561367281612e9b565b815261368060208301613632565b602082015292915050565b5f6001600160401b038211156136a3576136a36132dd565b5060051b60200190565b5f601f83601f8401126136be575f80fd5b823560206136d36136ce8361368b565b61337f565b82815260059290921b850181019181810190878411156136f1575f80fd5b8287015b848110156137f35780356001600160401b0380821115613713575f80fd5b908901906080601f19838d03810182131561372c575f80fd5b613734613319565b8885013561374181612e9b565b81526040858101356001600160e01b03198116811461375e575f80fd5b828b01526060868101358284015293860135938585111561377d575f80fd5b84870196508f603f880112613790575f80fd5b8a8701359450858511156137a6576137a66132dd565b6137b58b858f8801160161337f565b95508486528f828689010111156137ca575f80fd5b848288018c8801375f9486018b01949094525091820192909252855250509183019183016136f5565b50979650505050505050565b5f82601f83011261380e575f80fd5b8135602061381e6136ce8361368b565b8281526060928302850182019282820191908785111561383c575f80fd5b8387015b8581101561306f5781818a031215613856575f80fd5b61385e61333b565b813561386981612e9b565b81528186013561387881612e9b565b818701526040613889838201613632565b908201528452928401928101613840565b5f602082840312156138aa575f80fd5b81356001600160401b03808211156138c0575f80fd5b9083019060c082860312156138d3575f80fd5b6138db61335d565b82356138e681612e9b565b815260208301356138f6816132ae565b6020820152613908866040850161364d565b604082015260808301358281111561391e575f80fd5b61392a878286016136ad565b60608301525060a083013582811115613941575f80fd5b61394d878286016137ff565b60808301525095945050505050565b5f6001820161396d5761396d613409565b5060010190565b601f82111561106a57805f5260205f20601f840160051c810160208510156139995750805b601f840160051c820191505b818110156107c2575f81556001016139a5565b81516001600160401b038111156139d1576139d16132dd565b6139e5816139df8454613430565b84613974565b602080601f831160018114613a18575f8415613a015750858301515b5f19600386901b1c1916600185901b178555613a6f565b5f85815260208120601f198616915b82811015613a4657888601518255948401946001909101908401613a27565b5085821015613a6357878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f6001600160f81b038281166002600160f81b03198101613a9a57613a9a613409565b600101939250505056fef76fe33b8a0ebf7aa05740f479d10138c7c15bdc75b10e047cc15be2be15e5b45ffb10051d79c19b9690b0842a292cb621fbf85d15269ed21c4e6a431d892bb5a26469706673582212209bcacf54e51bcc01ad7c07df4c5e1467ed4c7d55f843f45f435f9f995ea19dc664736f6c63430008180033",
}

// SolverNetInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SolverNetInboxMetaData.ABI instead.
var SolverNetInboxABI = SolverNetInboxMetaData.ABI

// SolverNetInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolverNetInboxMetaData.Bin instead.
var SolverNetInboxBin = SolverNetInboxMetaData.Bin

// DeploySolverNetInbox deploys a new Ethereum contract, binding an instance of SolverNetInbox to it.
func DeploySolverNetInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolverNetInbox, error) {
	parsed, err := SolverNetInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolverNetInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolverNetInbox{SolverNetInboxCaller: SolverNetInboxCaller{contract: contract}, SolverNetInboxTransactor: SolverNetInboxTransactor{contract: contract}, SolverNetInboxFilterer: SolverNetInboxFilterer{contract: contract}}, nil
}

// SolverNetInbox is an auto generated Go binding around an Ethereum contract.
type SolverNetInbox struct {
	SolverNetInboxCaller     // Read-only binding to the contract
	SolverNetInboxTransactor // Write-only binding to the contract
	SolverNetInboxFilterer   // Log filterer for contract events
}

// SolverNetInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolverNetInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolverNetInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolverNetInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolverNetInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolverNetInboxSession struct {
	Contract     *SolverNetInbox   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolverNetInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolverNetInboxCallerSession struct {
	Contract *SolverNetInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SolverNetInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolverNetInboxTransactorSession struct {
	Contract     *SolverNetInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SolverNetInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolverNetInboxRaw struct {
	Contract *SolverNetInbox // Generic contract binding to access the raw methods on
}

// SolverNetInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolverNetInboxCallerRaw struct {
	Contract *SolverNetInboxCaller // Generic read-only contract binding to access the raw methods on
}

// SolverNetInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolverNetInboxTransactorRaw struct {
	Contract *SolverNetInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolverNetInbox creates a new instance of SolverNetInbox, bound to a specific deployed contract.
func NewSolverNetInbox(address common.Address, backend bind.ContractBackend) (*SolverNetInbox, error) {
	contract, err := bindSolverNetInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolverNetInbox{SolverNetInboxCaller: SolverNetInboxCaller{contract: contract}, SolverNetInboxTransactor: SolverNetInboxTransactor{contract: contract}, SolverNetInboxFilterer: SolverNetInboxFilterer{contract: contract}}, nil
}

// NewSolverNetInboxCaller creates a new read-only instance of SolverNetInbox, bound to a specific deployed contract.
func NewSolverNetInboxCaller(address common.Address, caller bind.ContractCaller) (*SolverNetInboxCaller, error) {
	contract, err := bindSolverNetInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxCaller{contract: contract}, nil
}

// NewSolverNetInboxTransactor creates a new write-only instance of SolverNetInbox, bound to a specific deployed contract.
func NewSolverNetInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SolverNetInboxTransactor, error) {
	contract, err := bindSolverNetInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxTransactor{contract: contract}, nil
}

// NewSolverNetInboxFilterer creates a new log filterer instance of SolverNetInbox, bound to a specific deployed contract.
func NewSolverNetInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SolverNetInboxFilterer, error) {
	contract, err := bindSolverNetInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxFilterer{contract: contract}, nil
}

// bindSolverNetInbox binds a generic wrapper to an already deployed contract.
func bindSolverNetInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolverNetInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetInbox *SolverNetInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetInbox.Contract.SolverNetInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetInbox *SolverNetInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.SolverNetInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetInbox *SolverNetInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.SolverNetInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolverNetInbox *SolverNetInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolverNetInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolverNetInbox *SolverNetInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolverNetInbox *SolverNetInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.contract.Transact(opts, method, params...)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolverNetInbox *SolverNetInboxCaller) DefaultConfLevel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "defaultConfLevel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolverNetInbox *SolverNetInboxSession) DefaultConfLevel() (uint8, error) {
	return _SolverNetInbox.Contract.DefaultConfLevel(&_SolverNetInbox.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolverNetInbox *SolverNetInboxCallerSession) DefaultConfLevel() (uint8, error) {
	return _SolverNetInbox.Contract.DefaultConfLevel(&_SolverNetInbox.CallOpts)
}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolverNetInbox *SolverNetInboxCaller) DeployedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "deployedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolverNetInbox *SolverNetInboxSession) DeployedAt() (*big.Int, error) {
	return _SolverNetInbox.Contract.DeployedAt(&_SolverNetInbox.CallOpts)
}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolverNetInbox *SolverNetInboxCallerSession) DeployedAt() (*big.Int, error) {
	return _SolverNetInbox.Contract.DeployedAt(&_SolverNetInbox.CallOpts)
}

// GetLatestOrderIdByStatus is a free data retrieval call binding the contract method 0xaf844c9b.
//
// Solidity: function getLatestOrderIdByStatus(uint8 status) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetLatestOrderIdByStatus(opts *bind.CallOpts, status uint8) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getLatestOrderIdByStatus", status)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestOrderIdByStatus is a free data retrieval call binding the contract method 0xaf844c9b.
//
// Solidity: function getLatestOrderIdByStatus(uint8 status) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetLatestOrderIdByStatus(status uint8) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetLatestOrderIdByStatus(&_SolverNetInbox.CallOpts, status)
}

// GetLatestOrderIdByStatus is a free data retrieval call binding the contract method 0xaf844c9b.
//
// Solidity: function getLatestOrderIdByStatus(uint8 status) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetLatestOrderIdByStatus(status uint8) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetLatestOrderIdByStatus(&_SolverNetInbox.CallOpts, status)
}

// GetNextOrderId is a free data retrieval call binding the contract method 0x29d54233.
//
// Solidity: function getNextOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetNextOrderId(opts *bind.CallOpts, user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getNextOrderId", user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextOrderId is a free data retrieval call binding the contract method 0x29d54233.
//
// Solidity: function getNextOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetNextOrderId(user common.Address) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetNextOrderId(&_SolverNetInbox.CallOpts, user)
}

// GetNextOrderId is a free data retrieval call binding the contract method 0x29d54233.
//
// Solidity: function getNextOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetNextOrderId(user common.Address) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetNextOrderId(&_SolverNetInbox.CallOpts, user)
}

// GetOffset is a free data retrieval call binding the contract method 0xeb6ad24e.
//
// Solidity: function getOffset() view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetOffset(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getOffset")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOffset is a free data retrieval call binding the contract method 0xeb6ad24e.
//
// Solidity: function getOffset() view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetOffset() ([32]byte, error) {
	return _SolverNetInbox.Contract.GetOffset(&_SolverNetInbox.CallOpts)
}

// GetOffset is a free data retrieval call binding the contract method 0xeb6ad24e.
//
// Solidity: function getOffset() view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetOffset() ([32]byte, error) {
	return _SolverNetInbox.Contract.GetOffset(&_SolverNetInbox.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 id) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolved, (uint8,uint8,uint32,address) state, bytes32 offset)
func (_SolverNetInbox *SolverNetInboxCaller) GetOrder(opts *bind.CallOpts, id [32]byte) (struct {
	Resolved IERC7683ResolvedCrossChainOrder
	State    ISolverNetInboxOrderState
	Offset   [32]byte
}, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getOrder", id)

	outstruct := new(struct {
		Resolved IERC7683ResolvedCrossChainOrder
		State    ISolverNetInboxOrderState
		Offset   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolved = *abi.ConvertType(out[0], new(IERC7683ResolvedCrossChainOrder)).(*IERC7683ResolvedCrossChainOrder)
	outstruct.State = *abi.ConvertType(out[1], new(ISolverNetInboxOrderState)).(*ISolverNetInboxOrderState)
	outstruct.Offset = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 id) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolved, (uint8,uint8,uint32,address) state, bytes32 offset)
func (_SolverNetInbox *SolverNetInboxSession) GetOrder(id [32]byte) (struct {
	Resolved IERC7683ResolvedCrossChainOrder
	State    ISolverNetInboxOrderState
	Offset   [32]byte
}, error) {
	return _SolverNetInbox.Contract.GetOrder(&_SolverNetInbox.CallOpts, id)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 id) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolved, (uint8,uint8,uint32,address) state, bytes32 offset)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetOrder(id [32]byte) (struct {
	Resolved IERC7683ResolvedCrossChainOrder
	State    ISolverNetInboxOrderState
	Offset   [32]byte
}, error) {
	return _SolverNetInbox.Contract.GetOrder(&_SolverNetInbox.CallOpts, id)
}

// GetOrderId is a free data retrieval call binding the contract method 0x5e0aec95.
//
// Solidity: function getOrderId(address user, uint256 nonce) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetOrderId(opts *bind.CallOpts, user common.Address, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getOrderId", user, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderId is a free data retrieval call binding the contract method 0x5e0aec95.
//
// Solidity: function getOrderId(address user, uint256 nonce) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetOrderId(user common.Address, nonce *big.Int) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetOrderId(&_SolverNetInbox.CallOpts, user, nonce)
}

// GetOrderId is a free data retrieval call binding the contract method 0x5e0aec95.
//
// Solidity: function getOrderId(address user, uint256 nonce) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetOrderId(user common.Address, nonce *big.Int) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetOrderId(&_SolverNetInbox.CallOpts, user, nonce)
}

// GetUserNonce is a free data retrieval call binding the contract method 0x6834e3a8.
//
// Solidity: function getUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxCaller) GetUserNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getUserNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNonce is a free data retrieval call binding the contract method 0x6834e3a8.
//
// Solidity: function getUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxSession) GetUserNonce(user common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.GetUserNonce(&_SolverNetInbox.CallOpts, user)
}

// GetUserNonce is a free data retrieval call binding the contract method 0x6834e3a8.
//
// Solidity: function getUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetUserNonce(user common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.GetUserNonce(&_SolverNetInbox.CallOpts, user)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolverNetInbox *SolverNetInboxCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolverNetInbox *SolverNetInboxSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _SolverNetInbox.Contract.HasAllRoles(&_SolverNetInbox.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolverNetInbox *SolverNetInboxCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _SolverNetInbox.Contract.HasAllRoles(&_SolverNetInbox.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolverNetInbox *SolverNetInboxCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolverNetInbox *SolverNetInboxSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _SolverNetInbox.Contract.HasAnyRole(&_SolverNetInbox.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolverNetInbox *SolverNetInboxCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _SolverNetInbox.Contract.HasAnyRole(&_SolverNetInbox.CallOpts, user, roles)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolverNetInbox *SolverNetInboxCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolverNetInbox *SolverNetInboxSession) Omni() (common.Address, error) {
	return _SolverNetInbox.Contract.Omni(&_SolverNetInbox.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolverNetInbox *SolverNetInboxCallerSession) Omni() (common.Address, error) {
	return _SolverNetInbox.Contract.Omni(&_SolverNetInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolverNetInbox *SolverNetInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolverNetInbox *SolverNetInboxSession) Owner() (common.Address, error) {
	return _SolverNetInbox.Contract.Owner(&_SolverNetInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolverNetInbox *SolverNetInboxCallerSession) Owner() (common.Address, error) {
	return _SolverNetInbox.Contract.Owner(&_SolverNetInbox.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolverNetInbox *SolverNetInboxCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolverNetInbox *SolverNetInboxSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.OwnershipHandoverExpiresAt(&_SolverNetInbox.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolverNetInbox *SolverNetInboxCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.OwnershipHandoverExpiresAt(&_SolverNetInbox.CallOpts, pendingOwner)
}

// PauseState is a free data retrieval call binding the contract method 0xd7118351.
//
// Solidity: function pauseState() view returns(uint8)
func (_SolverNetInbox *SolverNetInboxCaller) PauseState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "pauseState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PauseState is a free data retrieval call binding the contract method 0xd7118351.
//
// Solidity: function pauseState() view returns(uint8)
func (_SolverNetInbox *SolverNetInboxSession) PauseState() (uint8, error) {
	return _SolverNetInbox.Contract.PauseState(&_SolverNetInbox.CallOpts)
}

// PauseState is a free data retrieval call binding the contract method 0xd7118351.
//
// Solidity: function pauseState() view returns(uint8)
func (_SolverNetInbox *SolverNetInboxCallerSession) PauseState() (uint8, error) {
	return _SolverNetInbox.Contract.PauseState(&_SolverNetInbox.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x41b477dd.
//
// Solidity: function resolve((uint32,bytes32,bytes) order) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]))
func (_SolverNetInbox *SolverNetInboxCaller) Resolve(opts *bind.CallOpts, order IERC7683OnchainCrossChainOrder) (IERC7683ResolvedCrossChainOrder, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "resolve", order)

	if err != nil {
		return *new(IERC7683ResolvedCrossChainOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC7683ResolvedCrossChainOrder)).(*IERC7683ResolvedCrossChainOrder)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x41b477dd.
//
// Solidity: function resolve((uint32,bytes32,bytes) order) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]))
func (_SolverNetInbox *SolverNetInboxSession) Resolve(order IERC7683OnchainCrossChainOrder) (IERC7683ResolvedCrossChainOrder, error) {
	return _SolverNetInbox.Contract.Resolve(&_SolverNetInbox.CallOpts, order)
}

// Resolve is a free data retrieval call binding the contract method 0x41b477dd.
//
// Solidity: function resolve((uint32,bytes32,bytes) order) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]))
func (_SolverNetInbox *SolverNetInboxCallerSession) Resolve(order IERC7683OnchainCrossChainOrder) (IERC7683ResolvedCrossChainOrder, error) {
	return _SolverNetInbox.Contract.Resolve(&_SolverNetInbox.CallOpts, order)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolverNetInbox *SolverNetInboxCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolverNetInbox *SolverNetInboxSession) RolesOf(user common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.RolesOf(&_SolverNetInbox.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolverNetInbox *SolverNetInboxCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.RolesOf(&_SolverNetInbox.CallOpts, user)
}

// Validate is a free data retrieval call binding the contract method 0xd9e8407c.
//
// Solidity: function validate((uint32,bytes32,bytes) order) view returns(bool)
func (_SolverNetInbox *SolverNetInboxCaller) Validate(opts *bind.CallOpts, order IERC7683OnchainCrossChainOrder) (bool, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "validate", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Validate is a free data retrieval call binding the contract method 0xd9e8407c.
//
// Solidity: function validate((uint32,bytes32,bytes) order) view returns(bool)
func (_SolverNetInbox *SolverNetInboxSession) Validate(order IERC7683OnchainCrossChainOrder) (bool, error) {
	return _SolverNetInbox.Contract.Validate(&_SolverNetInbox.CallOpts, order)
}

// Validate is a free data retrieval call binding the contract method 0xd9e8407c.
//
// Solidity: function validate((uint32,bytes32,bytes) order) view returns(bool)
func (_SolverNetInbox *SolverNetInboxCallerSession) Validate(order IERC7683OnchainCrossChainOrder) (bool, error) {
	return _SolverNetInbox.Contract.Validate(&_SolverNetInbox.CallOpts, order)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolverNetInbox *SolverNetInboxSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _SolverNetInbox.Contract.CancelOwnershipHandover(&_SolverNetInbox.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _SolverNetInbox.Contract.CancelOwnershipHandover(&_SolverNetInbox.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x96c144f0.
//
// Solidity: function claim(bytes32 id, address to) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) Claim(opts *bind.TransactOpts, id [32]byte, to common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "claim", id, to)
}

// Claim is a paid mutator transaction binding the contract method 0x96c144f0.
//
// Solidity: function claim(bytes32 id, address to) returns()
func (_SolverNetInbox *SolverNetInboxSession) Claim(id [32]byte, to common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Claim(&_SolverNetInbox.TransactOpts, id, to)
}

// Claim is a paid mutator transaction binding the contract method 0x96c144f0.
//
// Solidity: function claim(bytes32 id, address to) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) Claim(id [32]byte, to common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Claim(&_SolverNetInbox.TransactOpts, id, to)
}

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(bytes32 id) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) Close(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "close", id)
}

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(bytes32 id) returns()
func (_SolverNetInbox *SolverNetInboxSession) Close(id [32]byte) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Close(&_SolverNetInbox.TransactOpts, id)
}

// Close is a paid mutator transaction binding the contract method 0x39c79e0c.
//
// Solidity: function close(bytes32 id) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) Close(id [32]byte) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Close(&_SolverNetInbox.TransactOpts, id)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolverNetInbox *SolverNetInboxSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.CompleteOwnershipHandover(&_SolverNetInbox.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.CompleteOwnershipHandover(&_SolverNetInbox.TransactOpts, pendingOwner)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.GrantRoles(&_SolverNetInbox.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.GrantRoles(&_SolverNetInbox.TransactOpts, user, roles)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner_, address solver_, address omni_) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, solver_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "initialize", owner_, solver_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner_, address solver_, address omni_) returns()
func (_SolverNetInbox *SolverNetInboxSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Initialize(&_SolverNetInbox.TransactOpts, owner_, solver_, omni_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address owner_, address solver_, address omni_) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Initialize(&_SolverNetInbox.TransactOpts, owner_, solver_, omni_)
}

// MarkFilled is a paid mutator transaction binding the contract method 0x2d622343.
//
// Solidity: function markFilled(bytes32 id, bytes32 fillHash, address creditedTo) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) MarkFilled(opts *bind.TransactOpts, id [32]byte, fillHash [32]byte, creditedTo common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "markFilled", id, fillHash, creditedTo)
}

// MarkFilled is a paid mutator transaction binding the contract method 0x2d622343.
//
// Solidity: function markFilled(bytes32 id, bytes32 fillHash, address creditedTo) returns()
func (_SolverNetInbox *SolverNetInboxSession) MarkFilled(id [32]byte, fillHash [32]byte, creditedTo common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.MarkFilled(&_SolverNetInbox.TransactOpts, id, fillHash, creditedTo)
}

// MarkFilled is a paid mutator transaction binding the contract method 0x2d622343.
//
// Solidity: function markFilled(bytes32 id, bytes32 fillHash, address creditedTo) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) MarkFilled(id [32]byte, fillHash [32]byte, creditedTo common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.MarkFilled(&_SolverNetInbox.TransactOpts, id, fillHash, creditedTo)
}

// Open is a paid mutator transaction binding the contract method 0xe917a962.
//
// Solidity: function open((uint32,bytes32,bytes) order) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) Open(opts *bind.TransactOpts, order IERC7683OnchainCrossChainOrder) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "open", order)
}

// Open is a paid mutator transaction binding the contract method 0xe917a962.
//
// Solidity: function open((uint32,bytes32,bytes) order) payable returns()
func (_SolverNetInbox *SolverNetInboxSession) Open(order IERC7683OnchainCrossChainOrder) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Open(&_SolverNetInbox.TransactOpts, order)
}

// Open is a paid mutator transaction binding the contract method 0xe917a962.
//
// Solidity: function open((uint32,bytes32,bytes) order) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) Open(order IERC7683OnchainCrossChainOrder) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Open(&_SolverNetInbox.TransactOpts, order)
}

// PauseAll is a paid mutator transaction binding the contract method 0x7cac41a6.
//
// Solidity: function pauseAll(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) PauseAll(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "pauseAll", pause)
}

// PauseAll is a paid mutator transaction binding the contract method 0x7cac41a6.
//
// Solidity: function pauseAll(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxSession) PauseAll(pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.PauseAll(&_SolverNetInbox.TransactOpts, pause)
}

// PauseAll is a paid mutator transaction binding the contract method 0x7cac41a6.
//
// Solidity: function pauseAll(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) PauseAll(pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.PauseAll(&_SolverNetInbox.TransactOpts, pause)
}

// PauseClose is a paid mutator transaction binding the contract method 0x792aec5c.
//
// Solidity: function pauseClose(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) PauseClose(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "pauseClose", pause)
}

// PauseClose is a paid mutator transaction binding the contract method 0x792aec5c.
//
// Solidity: function pauseClose(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxSession) PauseClose(pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.PauseClose(&_SolverNetInbox.TransactOpts, pause)
}

// PauseClose is a paid mutator transaction binding the contract method 0x792aec5c.
//
// Solidity: function pauseClose(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) PauseClose(pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.PauseClose(&_SolverNetInbox.TransactOpts, pause)
}

// PauseOpen is a paid mutator transaction binding the contract method 0xf904d285.
//
// Solidity: function pauseOpen(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) PauseOpen(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "pauseOpen", pause)
}

// PauseOpen is a paid mutator transaction binding the contract method 0xf904d285.
//
// Solidity: function pauseOpen(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxSession) PauseOpen(pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.PauseOpen(&_SolverNetInbox.TransactOpts, pause)
}

// PauseOpen is a paid mutator transaction binding the contract method 0xf904d285.
//
// Solidity: function pauseOpen(bool pause) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) PauseOpen(pause bool) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.PauseOpen(&_SolverNetInbox.TransactOpts, pause)
}

// Reject is a paid mutator transaction binding the contract method 0xdb3ea553.
//
// Solidity: function reject(bytes32 id, uint8 reason) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) Reject(opts *bind.TransactOpts, id [32]byte, reason uint8) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "reject", id, reason)
}

// Reject is a paid mutator transaction binding the contract method 0xdb3ea553.
//
// Solidity: function reject(bytes32 id, uint8 reason) returns()
func (_SolverNetInbox *SolverNetInboxSession) Reject(id [32]byte, reason uint8) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Reject(&_SolverNetInbox.TransactOpts, id, reason)
}

// Reject is a paid mutator transaction binding the contract method 0xdb3ea553.
//
// Solidity: function reject(bytes32 id, uint8 reason) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) Reject(id [32]byte, reason uint8) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Reject(&_SolverNetInbox.TransactOpts, id, reason)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolverNetInbox *SolverNetInboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolverNetInbox.Contract.RenounceOwnership(&_SolverNetInbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolverNetInbox.Contract.RenounceOwnership(&_SolverNetInbox.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.RenounceRoles(&_SolverNetInbox.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.RenounceRoles(&_SolverNetInbox.TransactOpts, roles)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolverNetInbox *SolverNetInboxSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _SolverNetInbox.Contract.RequestOwnershipHandover(&_SolverNetInbox.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _SolverNetInbox.Contract.RequestOwnershipHandover(&_SolverNetInbox.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.RevokeRoles(&_SolverNetInbox.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.RevokeRoles(&_SolverNetInbox.TransactOpts, user, roles)
}

// SetOutboxes is a paid mutator transaction binding the contract method 0x04a873ab.
//
// Solidity: function setOutboxes(uint64[] chainIds, address[] outboxes) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) SetOutboxes(opts *bind.TransactOpts, chainIds []uint64, outboxes []common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "setOutboxes", chainIds, outboxes)
}

// SetOutboxes is a paid mutator transaction binding the contract method 0x04a873ab.
//
// Solidity: function setOutboxes(uint64[] chainIds, address[] outboxes) returns()
func (_SolverNetInbox *SolverNetInboxSession) SetOutboxes(chainIds []uint64, outboxes []common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.SetOutboxes(&_SolverNetInbox.TransactOpts, chainIds, outboxes)
}

// SetOutboxes is a paid mutator transaction binding the contract method 0x04a873ab.
//
// Solidity: function setOutboxes(uint64[] chainIds, address[] outboxes) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) SetOutboxes(chainIds []uint64, outboxes []common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.SetOutboxes(&_SolverNetInbox.TransactOpts, chainIds, outboxes)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolverNetInbox *SolverNetInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.TransferOwnership(&_SolverNetInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.TransferOwnership(&_SolverNetInbox.TransactOpts, newOwner)
}

// SolverNetInboxClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the SolverNetInbox contract.
type SolverNetInboxClaimedIterator struct {
	Event *SolverNetInboxClaimed // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxClaimed)
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
		it.Event = new(SolverNetInboxClaimed)
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
func (it *SolverNetInboxClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxClaimed represents a Claimed event raised by the SolverNetInbox contract.
type SolverNetInboxClaimed struct {
	Id  [32]byte
	By  common.Address
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x8428df912f4f2125b442b488df9c7260cb607246895bcd29f262ecca090b1538.
//
// Solidity: event Claimed(bytes32 indexed id, address indexed by, address indexed to)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterClaimed(opts *bind.FilterOpts, id [][32]byte, by []common.Address, to []common.Address) (*SolverNetInboxClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "Claimed", idRule, byRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxClaimedIterator{contract: _SolverNetInbox.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x8428df912f4f2125b442b488df9c7260cb607246895bcd29f262ecca090b1538.
//
// Solidity: event Claimed(bytes32 indexed id, address indexed by, address indexed to)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *SolverNetInboxClaimed, id [][32]byte, by []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "Claimed", idRule, byRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxClaimed)
				if err := _SolverNetInbox.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x8428df912f4f2125b442b488df9c7260cb607246895bcd29f262ecca090b1538.
//
// Solidity: event Claimed(bytes32 indexed id, address indexed by, address indexed to)
func (_SolverNetInbox *SolverNetInboxFilterer) ParseClaimed(log types.Log) (*SolverNetInboxClaimed, error) {
	event := new(SolverNetInboxClaimed)
	if err := _SolverNetInbox.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the SolverNetInbox contract.
type SolverNetInboxClosedIterator struct {
	Event *SolverNetInboxClosed // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxClosed)
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
		it.Event = new(SolverNetInboxClosed)
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
func (it *SolverNetInboxClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxClosed represents a Closed event raised by the SolverNetInbox contract.
type SolverNetInboxClosed struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79.
//
// Solidity: event Closed(bytes32 indexed id)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterClosed(opts *bind.FilterOpts, id [][32]byte) (*SolverNetInboxClosedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "Closed", idRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxClosedIterator{contract: _SolverNetInbox.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79.
//
// Solidity: event Closed(bytes32 indexed id)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *SolverNetInboxClosed, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "Closed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxClosed)
				if err := _SolverNetInbox.contract.UnpackLog(event, "Closed", log); err != nil {
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

// ParseClosed is a log parse operation binding the contract event 0x7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79.
//
// Solidity: event Closed(bytes32 indexed id)
func (_SolverNetInbox *SolverNetInboxFilterer) ParseClosed(log types.Log) (*SolverNetInboxClosed, error) {
	event := new(SolverNetInboxClosed)
	if err := _SolverNetInbox.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxDefaultConfLevelSetIterator is returned from FilterDefaultConfLevelSet and is used to iterate over the raw logs and unpacked data for DefaultConfLevelSet events raised by the SolverNetInbox contract.
type SolverNetInboxDefaultConfLevelSetIterator struct {
	Event *SolverNetInboxDefaultConfLevelSet // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxDefaultConfLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxDefaultConfLevelSet)
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
		it.Event = new(SolverNetInboxDefaultConfLevelSet)
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
func (it *SolverNetInboxDefaultConfLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxDefaultConfLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxDefaultConfLevelSet represents a DefaultConfLevelSet event raised by the SolverNetInbox contract.
type SolverNetInboxDefaultConfLevelSet struct {
	Conf uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultConfLevelSet is a free log retrieval operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterDefaultConfLevelSet(opts *bind.FilterOpts) (*SolverNetInboxDefaultConfLevelSetIterator, error) {

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxDefaultConfLevelSetIterator{contract: _SolverNetInbox.contract, event: "DefaultConfLevelSet", logs: logs, sub: sub}, nil
}

// WatchDefaultConfLevelSet is a free log subscription operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchDefaultConfLevelSet(opts *bind.WatchOpts, sink chan<- *SolverNetInboxDefaultConfLevelSet) (event.Subscription, error) {

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxDefaultConfLevelSet)
				if err := _SolverNetInbox.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseDefaultConfLevelSet(log types.Log) (*SolverNetInboxDefaultConfLevelSet, error) {
	event := new(SolverNetInboxDefaultConfLevelSet)
	if err := _SolverNetInbox.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxFilledIterator is returned from FilterFilled and is used to iterate over the raw logs and unpacked data for Filled events raised by the SolverNetInbox contract.
type SolverNetInboxFilledIterator struct {
	Event *SolverNetInboxFilled // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxFilled)
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
		it.Event = new(SolverNetInboxFilled)
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
func (it *SolverNetInboxFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxFilled represents a Filled event raised by the SolverNetInbox contract.
type SolverNetInboxFilled struct {
	Id         [32]byte
	FillHash   [32]byte
	CreditedTo common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFilled is a free log retrieval operation binding the contract event 0xa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc.
//
// Solidity: event Filled(bytes32 indexed id, bytes32 indexed fillHash, address indexed creditedTo)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterFilled(opts *bind.FilterOpts, id [][32]byte, fillHash [][32]byte, creditedTo []common.Address) (*SolverNetInboxFilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fillHashRule []interface{}
	for _, fillHashItem := range fillHash {
		fillHashRule = append(fillHashRule, fillHashItem)
	}
	var creditedToRule []interface{}
	for _, creditedToItem := range creditedTo {
		creditedToRule = append(creditedToRule, creditedToItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "Filled", idRule, fillHashRule, creditedToRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxFilledIterator{contract: _SolverNetInbox.contract, event: "Filled", logs: logs, sub: sub}, nil
}

// WatchFilled is a free log subscription operation binding the contract event 0xa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc.
//
// Solidity: event Filled(bytes32 indexed id, bytes32 indexed fillHash, address indexed creditedTo)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchFilled(opts *bind.WatchOpts, sink chan<- *SolverNetInboxFilled, id [][32]byte, fillHash [][32]byte, creditedTo []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fillHashRule []interface{}
	for _, fillHashItem := range fillHash {
		fillHashRule = append(fillHashRule, fillHashItem)
	}
	var creditedToRule []interface{}
	for _, creditedToItem := range creditedTo {
		creditedToRule = append(creditedToRule, creditedToItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "Filled", idRule, fillHashRule, creditedToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxFilled)
				if err := _SolverNetInbox.contract.UnpackLog(event, "Filled", log); err != nil {
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
// Solidity: event Filled(bytes32 indexed id, bytes32 indexed fillHash, address indexed creditedTo)
func (_SolverNetInbox *SolverNetInboxFilterer) ParseFilled(log types.Log) (*SolverNetInboxFilled, error) {
	event := new(SolverNetInboxFilled)
	if err := _SolverNetInbox.contract.UnpackLog(event, "Filled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SolverNetInbox contract.
type SolverNetInboxInitializedIterator struct {
	Event *SolverNetInboxInitialized // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxInitialized)
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
		it.Event = new(SolverNetInboxInitialized)
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
func (it *SolverNetInboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxInitialized represents a Initialized event raised by the SolverNetInbox contract.
type SolverNetInboxInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*SolverNetInboxInitializedIterator, error) {

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxInitializedIterator{contract: _SolverNetInbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SolverNetInboxInitialized) (event.Subscription, error) {

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxInitialized)
				if err := _SolverNetInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseInitialized(log types.Log) (*SolverNetInboxInitialized, error) {
	event := new(SolverNetInboxInitialized)
	if err := _SolverNetInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxOmniPortalSetIterator is returned from FilterOmniPortalSet and is used to iterate over the raw logs and unpacked data for OmniPortalSet events raised by the SolverNetInbox contract.
type SolverNetInboxOmniPortalSetIterator struct {
	Event *SolverNetInboxOmniPortalSet // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxOmniPortalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxOmniPortalSet)
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
		it.Event = new(SolverNetInboxOmniPortalSet)
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
func (it *SolverNetInboxOmniPortalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxOmniPortalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxOmniPortalSet represents a OmniPortalSet event raised by the SolverNetInbox contract.
type SolverNetInboxOmniPortalSet struct {
	Omni common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOmniPortalSet is a free log retrieval operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterOmniPortalSet(opts *bind.FilterOpts) (*SolverNetInboxOmniPortalSetIterator, error) {

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxOmniPortalSetIterator{contract: _SolverNetInbox.contract, event: "OmniPortalSet", logs: logs, sub: sub}, nil
}

// WatchOmniPortalSet is a free log subscription operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchOmniPortalSet(opts *bind.WatchOpts, sink chan<- *SolverNetInboxOmniPortalSet) (event.Subscription, error) {

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxOmniPortalSet)
				if err := _SolverNetInbox.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseOmniPortalSet(log types.Log) (*SolverNetInboxOmniPortalSet, error) {
	event := new(SolverNetInboxOmniPortalSet)
	if err := _SolverNetInbox.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxOpenIterator is returned from FilterOpen and is used to iterate over the raw logs and unpacked data for Open events raised by the SolverNetInbox contract.
type SolverNetInboxOpenIterator struct {
	Event *SolverNetInboxOpen // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxOpen)
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
		it.Event = new(SolverNetInboxOpen)
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
func (it *SolverNetInboxOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxOpen represents a Open event raised by the SolverNetInbox contract.
type SolverNetInboxOpen struct {
	OrderId       [32]byte
	ResolvedOrder IERC7683ResolvedCrossChainOrder
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpen is a free log retrieval operation binding the contract event 0xa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c.
//
// Solidity: event Open(bytes32 indexed orderId, (address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolvedOrder)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterOpen(opts *bind.FilterOpts, orderId [][32]byte) (*SolverNetInboxOpenIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "Open", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxOpenIterator{contract: _SolverNetInbox.contract, event: "Open", logs: logs, sub: sub}, nil
}

// WatchOpen is a free log subscription operation binding the contract event 0xa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c.
//
// Solidity: event Open(bytes32 indexed orderId, (address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolvedOrder)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchOpen(opts *bind.WatchOpts, sink chan<- *SolverNetInboxOpen, orderId [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "Open", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxOpen)
				if err := _SolverNetInbox.contract.UnpackLog(event, "Open", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseOpen(log types.Log) (*SolverNetInboxOpen, error) {
	event := new(SolverNetInboxOpen)
	if err := _SolverNetInbox.contract.UnpackLog(event, "Open", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxOutboxSetIterator is returned from FilterOutboxSet and is used to iterate over the raw logs and unpacked data for OutboxSet events raised by the SolverNetInbox contract.
type SolverNetInboxOutboxSetIterator struct {
	Event *SolverNetInboxOutboxSet // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxOutboxSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxOutboxSet)
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
		it.Event = new(SolverNetInboxOutboxSet)
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
func (it *SolverNetInboxOutboxSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxOutboxSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxOutboxSet represents a OutboxSet event raised by the SolverNetInbox contract.
type SolverNetInboxOutboxSet struct {
	ChainId uint64
	Outbox  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxSet is a free log retrieval operation binding the contract event 0xf730978310b4a2a0e6c673324d737afdb93d0afefed14a3d061b60f66e31f4e3.
//
// Solidity: event OutboxSet(uint64 indexed chainId, address indexed outbox)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterOutboxSet(opts *bind.FilterOpts, chainId []uint64, outbox []common.Address) (*SolverNetInboxOutboxSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "OutboxSet", chainIdRule, outboxRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxOutboxSetIterator{contract: _SolverNetInbox.contract, event: "OutboxSet", logs: logs, sub: sub}, nil
}

// WatchOutboxSet is a free log subscription operation binding the contract event 0xf730978310b4a2a0e6c673324d737afdb93d0afefed14a3d061b60f66e31f4e3.
//
// Solidity: event OutboxSet(uint64 indexed chainId, address indexed outbox)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchOutboxSet(opts *bind.WatchOpts, sink chan<- *SolverNetInboxOutboxSet, chainId []uint64, outbox []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "OutboxSet", chainIdRule, outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxOutboxSet)
				if err := _SolverNetInbox.contract.UnpackLog(event, "OutboxSet", log); err != nil {
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

// ParseOutboxSet is a log parse operation binding the contract event 0xf730978310b4a2a0e6c673324d737afdb93d0afefed14a3d061b60f66e31f4e3.
//
// Solidity: event OutboxSet(uint64 indexed chainId, address indexed outbox)
func (_SolverNetInbox *SolverNetInboxFilterer) ParseOutboxSet(log types.Log) (*SolverNetInboxOutboxSet, error) {
	event := new(SolverNetInboxOutboxSet)
	if err := _SolverNetInbox.contract.UnpackLog(event, "OutboxSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the SolverNetInbox contract.
type SolverNetInboxOwnershipHandoverCanceledIterator struct {
	Event *SolverNetInboxOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxOwnershipHandoverCanceled)
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
		it.Event = new(SolverNetInboxOwnershipHandoverCanceled)
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
func (it *SolverNetInboxOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the SolverNetInbox contract.
type SolverNetInboxOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*SolverNetInboxOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxOwnershipHandoverCanceledIterator{contract: _SolverNetInbox.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *SolverNetInboxOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxOwnershipHandoverCanceled)
				if err := _SolverNetInbox.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*SolverNetInboxOwnershipHandoverCanceled, error) {
	event := new(SolverNetInboxOwnershipHandoverCanceled)
	if err := _SolverNetInbox.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the SolverNetInbox contract.
type SolverNetInboxOwnershipHandoverRequestedIterator struct {
	Event *SolverNetInboxOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxOwnershipHandoverRequested)
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
		it.Event = new(SolverNetInboxOwnershipHandoverRequested)
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
func (it *SolverNetInboxOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the SolverNetInbox contract.
type SolverNetInboxOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*SolverNetInboxOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxOwnershipHandoverRequestedIterator{contract: _SolverNetInbox.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *SolverNetInboxOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxOwnershipHandoverRequested)
				if err := _SolverNetInbox.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseOwnershipHandoverRequested(log types.Log) (*SolverNetInboxOwnershipHandoverRequested, error) {
	event := new(SolverNetInboxOwnershipHandoverRequested)
	if err := _SolverNetInbox.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SolverNetInbox contract.
type SolverNetInboxOwnershipTransferredIterator struct {
	Event *SolverNetInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxOwnershipTransferred)
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
		it.Event = new(SolverNetInboxOwnershipTransferred)
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
func (it *SolverNetInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxOwnershipTransferred represents a OwnershipTransferred event raised by the SolverNetInbox contract.
type SolverNetInboxOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*SolverNetInboxOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxOwnershipTransferredIterator{contract: _SolverNetInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SolverNetInboxOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxOwnershipTransferred)
				if err := _SolverNetInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseOwnershipTransferred(log types.Log) (*SolverNetInboxOwnershipTransferred, error) {
	event := new(SolverNetInboxOwnershipTransferred)
	if err := _SolverNetInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxRejectedIterator is returned from FilterRejected and is used to iterate over the raw logs and unpacked data for Rejected events raised by the SolverNetInbox contract.
type SolverNetInboxRejectedIterator struct {
	Event *SolverNetInboxRejected // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxRejected)
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
		it.Event = new(SolverNetInboxRejected)
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
func (it *SolverNetInboxRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxRejected represents a Rejected event raised by the SolverNetInbox contract.
type SolverNetInboxRejected struct {
	Id     [32]byte
	By     common.Address
	Reason uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRejected is a free log retrieval operation binding the contract event 0x21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f.
//
// Solidity: event Rejected(bytes32 indexed id, address indexed by, uint8 indexed reason)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterRejected(opts *bind.FilterOpts, id [][32]byte, by []common.Address, reason []uint8) (*SolverNetInboxRejectedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "Rejected", idRule, byRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxRejectedIterator{contract: _SolverNetInbox.contract, event: "Rejected", logs: logs, sub: sub}, nil
}

// WatchRejected is a free log subscription operation binding the contract event 0x21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f.
//
// Solidity: event Rejected(bytes32 indexed id, address indexed by, uint8 indexed reason)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchRejected(opts *bind.WatchOpts, sink chan<- *SolverNetInboxRejected, id [][32]byte, by []common.Address, reason []uint8) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "Rejected", idRule, byRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxRejected)
				if err := _SolverNetInbox.contract.UnpackLog(event, "Rejected", log); err != nil {
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

// ParseRejected is a log parse operation binding the contract event 0x21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f.
//
// Solidity: event Rejected(bytes32 indexed id, address indexed by, uint8 indexed reason)
func (_SolverNetInbox *SolverNetInboxFilterer) ParseRejected(log types.Log) (*SolverNetInboxRejected, error) {
	event := new(SolverNetInboxRejected)
	if err := _SolverNetInbox.contract.UnpackLog(event, "Rejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolverNetInboxRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the SolverNetInbox contract.
type SolverNetInboxRolesUpdatedIterator struct {
	Event *SolverNetInboxRolesUpdated // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxRolesUpdated)
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
		it.Event = new(SolverNetInboxRolesUpdated)
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
func (it *SolverNetInboxRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxRolesUpdated represents a RolesUpdated event raised by the SolverNetInbox contract.
type SolverNetInboxRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*SolverNetInboxRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxRolesUpdatedIterator{contract: _SolverNetInbox.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *SolverNetInboxRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxRolesUpdated)
				if err := _SolverNetInbox.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseRolesUpdated(log types.Log) (*SolverNetInboxRolesUpdated, error) {
	event := new(SolverNetInboxRolesUpdated)
	if err := _SolverNetInbox.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
