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

// SolverNetCall is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type SolverNetCall struct {
// 	Target   common.Address
// 	Selector [4]byte
// 	Value    *big.Int
// 	Params   []byte
// }

// SolverNetFillOriginData is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type SolverNetFillOriginData struct {
// 	SrcChainId   uint64
// 	DestChainId  uint64
// 	FillDeadline uint32
// 	Calls        []SolverNetCall
// 	Expenses     []SolverNetTokenExpense
// }

// SolverNetTokenExpense is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type SolverNetTokenExpense struct {
// 	Spender common.Address
// 	Token   common.Address
// 	Amount  *big.Int
// }

// SolverNetInboxMetaData contains all meta data concerning the SolverNetInbox contract.
var SolverNetInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"close\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestOrderOffset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint248\",\"internalType\":\"uint248\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrder\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"resolved\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"state\",\"type\":\"tuple\",\"internalType\":\"structISolverNetInbox.OrderState\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISolverNetInbox.Status\"},{\"name\":\"rejectReason\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedBy\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"offset\",\"type\":\"uint248\",\"internalType\":\"uint248\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserNonce\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"markFilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"open\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseClose\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseOpen\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reject\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOutboxes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboxes\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validate\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Closed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillOriginData\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillOriginData\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSolverNet.FillOriginData\",\"components\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structSolverNet.Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"expenses\",\"type\":\"tuple[]\",\"internalType\":\"structSolverNet.TokenExpense[]\",\"components\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Filled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Open\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"resolvedOrder\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboxSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Rejected\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCallTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenseAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenseToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFillDeadline\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMissingCalls\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNativeDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrderData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrderTypehash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReason\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderStillValid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PortalPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongFillHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongSourceChain\",\"inputs\":[]}]",
	Bin: "0x60a060405234801562000010575f80fd5b5063ffffffff60643b1615620000965760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156200007e575060408051601f3d908101601f191682019092526200007b9181019062000112565b60015b6200008d57436080526200009b565b6080526200009b565b436080525b620000a5620000ab565b6200012a565b63409feecd1980546001811615620000ca5763f92ee8a95f526004601cfd5b6001600160401b03808260011c146200010d578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b5f6020828403121562000123575f80fd5b5051919050565b608051613d8d620001435f395f6105960152613d8d5ff3fe6080604052600436106101f1575f3560e01c8063715018a611610108578063d9e8407c1161009d578063f04e283e1161006d578063f04e283e146105b8578063f2fde38b146105cb578063f904d285146105de578063fee81cf4146105fd578063ff4464101461062e575f80fd5b8063d9e8407c14610534578063db3ea55314610553578063e917a96214610572578063eae4c19f14610585575f80fd5b80638da5cb5b116100d85780638da5cb5b146104be57806396c144f0146104d6578063c0c53b8b146104f5578063d711835114610514575f80fd5b8063715018a61461044757806374eeb8471461044f578063792aec5c146104805780637cac41a61461049f575f80fd5b806339acf9f111610189578063514e62fc11610159578063514e62fc1461038957806354d1f13d146103be5780635778472a146103c65780635e0aec95146103f45780636834e3a814610413575f80fd5b806339acf9f1146102f557806339c79e0c1461032b57806341b477dd1461034a5780634a4ee7b114610376575f80fd5b806325692962116101c4578063256929621461027057806329d54233146102785780632d622343146102a55780632de94807146102c4575f80fd5b806304a873ab146101f5578063183a4f6e146102165780631c10893f146102295780631cd64df41461023c575b5f80fd5b348015610200575f80fd5b5061021461020f366004612e0a565b610655565b005b610214610224366004612e70565b6107a9565b610214610237366004612e9b565b6107b6565b348015610247575f80fd5b5061025b610256366004612e9b565b6107cc565b60405190151581526020015b60405180910390f35b6102146107ea565b348015610283575f80fd5b50610297610292366004612ec5565b610836565b604051908152602001610267565b3480156102b0575f80fd5b506102146102bf366004612ee0565b610859565b3480156102cf575f80fd5b506102976102de366004612ec5565b638b78c6d8600c9081525f91909152602090205490565b348015610300575f80fd5b505f54610313906001600160a01b031681565b6040516001600160a01b039091168152602001610267565b348015610336575f80fd5b50610214610345366004612e70565b610af4565b348015610355575f80fd5b50610369610364366004612f16565b610e54565b604051610267919061311a565b610214610384366004612e9b565b610ea3565b348015610394575f80fd5b5061025b6103a3366004612e9b565b638b78c6d8600c9081525f9290925260209091205416151590565b610214610eb5565b3480156103d1575f80fd5b506103e56103e0366004612e70565b610eee565b60405161026793929190613140565b3480156103ff575f80fd5b5061029761040e366004612e9b565b610fd2565b34801561041e575f80fd5b5061029761042d366004612ec5565b6001600160a01b03165f908152600a602052604090205490565b610214610fe4565b34801561045a575f80fd5b505f5461046e90600160a01b900460ff1681565b60405160ff9091168152602001610267565b34801561048b575f80fd5b5061021461049a3660046131cb565b610ff7565b3480156104aa575f80fd5b506102146104b93660046131cb565b611019565b3480156104c9575f80fd5b50638b78c6d81954610313565b3480156104e1575f80fd5b506102146104f03660046131e6565b611059565b348015610500575f80fd5b5061021461050f366004613214565b6111c6565b34801561051f575f80fd5b5060025461046e90600160f81b900460ff1681565b34801561053f575f80fd5b5061025b61054e366004612f16565b611255565b34801561055e575f80fd5b5061021461056d366004613251565b611268565b610214610580366004612f16565b6113ef565b348015610590575f80fd5b506102977f000000000000000000000000000000000000000000000000000000000000000081565b6102146105c6366004612ec5565b6115be565b6102146105d9366004612ec5565b6115f8565b3480156105e9575f80fd5b506102146105f83660046131cb565b61161e565b348015610608575f80fd5b50610297610617366004612ec5565b63389a75e1600c9081525f91909152602090205490565b348015610639575f80fd5b506002546040516001600160f81b039091168152602001610267565b61065d611640565b82811461067d57604051634ec4810560e11b815260040160405180910390fd5b5f5b838110156107a25782828281811061069957610699613279565b90506020020160208101906106ae9190612ec5565b60035f8787858181106106c3576106c3613279565b90506020020160208101906106d891906132a1565b6001600160401b0316815260208101919091526040015f2080546001600160a01b0319166001600160a01b039290921691909117905582828281811061072057610720613279565b90506020020160208101906107359190612ec5565b6001600160a01b031685858381811061075057610750613279565b905060200201602081019061076591906132a1565b6001600160401b03167ff730978310b4a2a0e6c673324d737afdb93d0afefed14a3d061b60f66e31f4e360405160405180910390a360010161067f565b5050505050565b6107b3338261165a565b50565b6107be611640565b6107c88282611665565b5050565b638b78c6d8600c9081525f8390526020902054811681145b92915050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f80a250565b6001600160a01b0381165f908152600a60205260408120546107e4908390611671565b5f5460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa15801561089d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108c1919061338e565b8051600180546020938401516001600160401b039384166001600160e01b031990921691909117600160401b6001600160a01b0392831602179091555f868152600484526040808220815160608101835290549384168152600160a01b840490941684860152600160e01b90920463ffffffff168383015286815260089093528083208151608081019092528054929392829060ff1660058111156109685761096861312c565b60058111156109795761097961312c565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015290506001815160058111156109c8576109c861312c565b146109e657604051635d12a4a360e11b815260040160405180910390fd5b60208201516001546001600160401b03908116911614610a1957604051633687f39960e21b815260040160405180910390fd5b6001546001600160401b0381165f908152600360205260409020546001600160a01b03908116600160401b9092041614610a65576040516282b42960e81b815260040160405180910390fd5b610a6e856116b3565b8414610a8d57604051631f53eaed60e21b815260040160405180910390fd5b610a9a8560045f8661195b565b610aa5856004611a6f565b826001600160a01b031684867fa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc60405160405180910390a45050600180546001600160e01b0319169055505050565b6002545f80516020613d3883398151915290600160f81b900460ff168015610bb55760ff81166001148015610b3557505f80516020613d1883398151915282145b15610b5357604051631309a56360e01b815260040160405180910390fd5b60ff81166002148015610b7257505f80516020613d3883398151915282145b15610b9057604051631309a56360e01b815260040160405180910390fd5b60021960ff821601610bb55760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd212685403610bd35763ab143c065f526004601cfd5b3068929eee149b4bd21268555f838152600860205260408082208151608081019092528054829060ff166005811115610c0e57610c0e61312c565b6005811115610c1f57610c1f61312c565b8152905460ff61010082041660208084019190915263ffffffff62010000830481166040808601919091526001600160a01b03600160301b90940484166060958601525f8a815260048452819020815195860182525493841685526001600160401b03600160a01b85041692850192909252600160e01b90920490911690820152909150600182516005811115610cb857610cb861312c565b14610cd657604051635d12a4a360e11b815260040160405180910390fd5b80516001600160a01b03163314610cff576040516282b42960e81b815260040160405180910390fd5b5f5460208201516040516308c3569160e31b81527ffeccba1cfc4544bf9cd83b76f36ae5c464750b6c43f682e26744ee21ec31fc1e60048201526001600160401b0390911660248201526001600160a01b039091169063461ab48890604401602060405180830381865afa158015610d79573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d9d91906133cd565b15610dbb57604051630c2e605760e11b815260040160405180910390fd5b42615460826040015163ffffffff16610dd491906133fc565b10610df2576040516321bb6b2160e11b815260040160405180910390fd5b610dff8560035f3361195b565b610e0c85825f0151611b06565b610e17856003611a6f565b60405185907f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79905f90a250503868929eee149b4bd2126855505050565b610e5c612c41565b5f610e6683611b9f565b8051516001600160a01b0381165f908152600a602052604090205491925090610e9b908390610e96908490611671565b611ea7565b949350505050565b610eab611640565b6107c8828261165a565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f80a2565b610ef6612c41565b604080516080810182525f80825260208201819052918101829052606081018290529080610f2385611f26565b9050610f2f8186611ea7565b5f86815260086020908152604080832060099092529182902054825160808101909352815491926001600160f81b03909116918390829060ff166005811115610f7a57610f7a61312c565b6005811115610f8b57610f8b61312c565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015292989297509550909350505050565b5f610fdd8383611671565b9392505050565b610fec611640565b610ff55f61217c565b565b6001611002816121b9565b6107c85f80516020613d38833981519152836121ea565b6001611024816121b9565b8161103c57600280546001600160f81b031690555050565b600280546001600160f81b0316600360f81b17905560035b505050565b3068929eee149b4bd2126854036110775763ab143c065f526004601cfd5b3068929eee149b4bd21268555f828152600860205260408082208151608081019092528054829060ff1660058111156110b2576110b261312c565b60058111156110c3576110c361312c565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015290506004815160058111156111125761111261312c565b146111305760405163789bae3560e01b815260040160405180910390fd5b60608101516001600160a01b0316331461115c576040516282b42960e81b815260040160405180910390fd5b6111698360055f3361195b565b6111738383611b06565b61117e836005611a6f565b6040516001600160a01b03831690339085907f8428df912f4f2125b442b488df9c7260cb607246895bcd29f262ecca090b1538905f90a4503868929eee149b4bd21268555050565b63409feecd1980546003825580156111fc5760018160011c14303b106111f35763f92ee8a95f526004601cfd5b818160ff1b1b91505b50611206846122a3565b611211836001611665565b61121a826122de565b801561124f576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b50505050565b5f61125f82611b9f565b50600192915050565b600161127381612380565b3068929eee149b4bd2126854036112915763ab143c065f526004601cfd5b3068929eee149b4bd21268555f838152600860205260408082208151608081019092528054829060ff1660058111156112cc576112cc61312c565b60058111156112dd576112dd61312c565b81529054610100810460ff908116602084015262010000820463ffffffff166040840152600160301b9091046001600160a01b031660609092019190915290915083165f0361133f576040516337b89b9360e21b815260040160405180910390fd5b6001815160058111156113545761135461312c565b1461137257604051635d12a4a360e11b815260040160405180910390fd5b61137f846002853361195b565b5f848152600460205260409020546113a19085906001600160a01b0316611b06565b6113ac846002611a6f565b60405160ff841690339086907f21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f905f90a4503868929eee149b4bd2126855505050565b6002545f80516020613d1883398151915290600160f81b900460ff1680156114b05760ff8116600114801561143057505f80516020613d1883398151915282145b1561144e57604051631309a56360e01b815260040160405180910390fd5b60ff8116600214801561146d57505f80516020613d3883398151915282145b1561148b57604051631309a56360e01b815260040160405180910390fd5b60021960ff8216016114b05760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd2126854036114ce5763ab143c065f526004601cfd5b3068929eee149b4bd21268555f6114e484611b9f565b90506114f381602001516123a4565b5f6114fd82612409565b905080608001517f71069e0637faca19b5cceb36f6ee664347f592a954e37edf597b6c0f37afd59f8260e001515f8151811061153b5761153b613279565b60200260200101516040015180602001905181019061155a9190613669565b6040516115679190613784565b60405180910390a280608001517fa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c826040516115a3919061311a565b60405180910390a250503868929eee149b4bd2126855505050565b6115c6611640565b63389a75e1600c52805f526020600c2080544211156115ec57636f5e88185f526004601cfd5b5f90556107b38161217c565b611600611640565b8060601b61161557637448fbae5f526004601cfd5b6107b38161217c565b6001611629816121b9565b6107c85f80516020613d18833981519152836121ea565b638b78c6d819543314610ff5576382b429005f526004601cfd5b6107c882825f61267f565b6107c88282600161267f565b604080516001600160a01b03841660208201529081018290524660608201525f9060800160405160208183030381529060405280519060200120905092915050565b5f818152600460209081526040808320815160608101835290546001600160a01b0381168252600160a01b81046001600160401b031682850152600160e01b900463ffffffff1681830152848452600683528184208054835181860281018601909452808452919385939290849084015b82821015611820575f848152602090819020604080516080810182526003860290920180546001600160a01b0381168452600160a01b900460e01b6001600160e01b0319169383019390935260018301549082015260028201805491929160608401919061179190613877565b80601f01602080910402602001604051908101604052809291908181526020018280546117bd90613877565b80156118085780601f106117df57610100808354040283529160200191611808565b820191905f5260205f20905b8154815290600101906020018083116117eb57829003601f168201915b50505050508152505081526020019060010190611724565b5050505090505f60075f8681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156118ba575f848152602090819020604080516060810182526002860290920180546001600160a01b03908116845260019182015490811684860152600160a01b90046001600160601b0316918301919091529083529092019101611857565b5050505090505f6040518060a00160405280466001600160401b0316815260200185602001516001600160401b03168152602001856040015163ffffffff168152602001848152602001838152509050858160405160200161191c9190613784565b60408051601f198184030181529082905261193a92916020016138af565b60405160208183030381529060405280519060200120945050505050919050565b5f848152600860205260409081902054815160808101909252610100900460ff1690808560058111156119905761199061312c565b81526020015f8560ff16116119a557826119a7565b845b60ff16815263ffffffff42166020808301919091526001600160a01b0385166040928301525f88815260089091522081518154829060ff191660018360058111156119f4576119f461312c565b02179055506020820151815460408401516060909401516001600160a01b0316600160301b026601000000000000600160d01b031963ffffffff909516620100000265ffffffff00001960ff909416610100029390931665ffffffffff00199092169190911791909117929092169190911790555050505050565b6001816005811115611a8357611a8361312c565b03611a8c575050565b6004816005811115611aa057611aa061312c565b14611ab4575f828152600560205260408120555b6005816005811115611ac857611ac861312c565b146107c8575f82815260046020908152604080832083905560069091528120611af091612c99565b5f8281526007602052604081206107c891612cb7565b5f828152600560209081526040918290208251808401909352546001600160a01b0381168352600160a01b90046001600160601b0316908201819052156110545780516001600160a01b0316611b78576020810151611054906001600160a01b038416906001600160601b03166126d6565b60208101518151611054916001600160a01b039091169084906001600160601b03166126ef565b611ba7612cd5565b42611bb560208401846138c7565b63ffffffff1611611bd95760405163582e388960e01b815260040160405180910390fd5b7f2e7de755ca70cb933dc80103af16cc3303580e5712f1a8927d6461441e99a1e6826020013514611c1d57604051636aea87f360e01b815260040160405180910390fd5b611c2a60408301836138e2565b90505f03611c4b5760405163a342e7d960e01b815260040160405180910390fd5b5f611c5960408401846138e2565b810190611c669190613b1c565b80519091506001600160a01b0316611c7c573381525b60208101516001600160401b03161580611ca257504681602001516001600160401b0316145b15611cc057604051633d23e4d160e11b815260040160405180910390fd5b6040805160608101825282516001600160a01b031681526020808401516001600160401b0316818301525f92820190611cfb908701876138c7565b63ffffffff16905260608301518051919250905f03611d2d57604051639cc71f7d60e01b815260040160405180910390fd5b805160201015611d5057604051634ec4810560e11b815260040160405180910390fd5b5f5b8151811015611dad575f828281518110611d6e57611d6e613279565b602090810291909101015180519091506001600160a01b0316611da45760405163017ab86160e21b815260040160405180910390fd5b50600101611d52565b506080830151805160201015611dd657604051634ec4810560e11b815260040160405180910390fd5b5f5b8151811015611e7e575f6001600160a01b0316828281518110611dfd57611dfd613279565b6020026020010151602001516001600160a01b031603611e305760405163027dcfa160e31b815260040160405180910390fd5b818181518110611e4257611e42613279565b6020026020010151604001516001600160601b03165f03611e765760405163a0ce339960e01b815260040160405180910390fd5b600101611dd8565b506040805160808101825293845293840151602084015292820152606081019190915292915050565b611eaf612c41565b82515f611ebb85612739565b90505f611ec786612995565b90505f611ed387612a99565b604080516101008101825286516001600160a01b031681524660208201525f8183015295015163ffffffff1660608601526080850187905260a08501939093525060c083015260e0820152905092915050565b611f2e612cd5565b604080515f8481526004602090815283822060e084018552546001600160a01b0380821660808601908152600160a01b8084046001600160401b031660a0880152600160e01b90930463ffffffff1660c087015285528784526005835285842086518088018852905491821681529190046001600160601b031681830152818401528582526006815283822080548551818402810184018752818152949586019493919290919084015b828210156120d4575f848152602090819020604080516080810182526003860290920180546001600160a01b0381168452600160a01b900460e01b6001600160e01b0319169383019390935260018301549082015260028201805491929160608401919061204590613877565b80601f016020809104026020016040519081016040528092919081815260200182805461207190613877565b80156120bc5780601f10612093576101008083540402835291602001916120bc565b820191905f5260205f20905b81548152906001019060200180831161209f57829003601f168201915b50505050508152505081526020019060010190611fd8565b50505050815260200160075f8581526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015612170575f848152602090819020604080516060810182526002860290920180546001600160a01b03908116845260019182015490811684860152600160a01b90046001600160601b031691830191909152908352909201910161210d565b50505091525092915050565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a355565b638b78c6d8195433146107b357638b78c6d8600c52335f52806020600c2054166107b3576382b429005f526004601cfd5b600254600160f81b900460ff1660021981016122195760405163aaae8ef760e01b815260040160405180910390fd5b5f5f80516020613d188339815191528414612235576002612238565b60015b90508261224e578060ff168260ff161415612258565b8060ff168260ff16145b1561227657604051631309a56360e01b815260040160405180910390fd5b82612281575f612283565b805b6002601f6101000a81548160ff021916908360ff16021790555050505050565b6001600160a01b0316638b78c6d819819055805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b03811661232d5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b604482015260640160405180910390fd5b5f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52335f52806020600c2054166107b3576382b429005f526004601cfd5b80516001600160a01b03166123e05780602001516001600160601b031634146107b35760405163036f810f60e41b815260040160405180910390fd5b602081015181516107b3916001600160a01b0390911690339030906001600160601b0316612ba5565b612411612c41565b8151516001600160a01b0381165f908152600a602052604081208054612447918491908461243e83613bcf565b91905055611671565b90506124538482611ea7565b84515f8381526004602090815260408083208451815484870151968401516001600160a01b039283166001600160e01b031990921691909117600160a01b6001600160401b039098168802176001600160e01b0316600160e01b63ffffffff9092169190910217909155828a015160058452919093208151919092015192166001600160601b0390921690920217905592506124ed612bfe565b5f82815260096020526040812080546001600160f81b0319166001600160f81b0393909316929092179091555b8460400151518110156125cf575f82815260066020526040908190209086015180518390811061254c5761254c613279565b6020908102919091018101518254600181810185555f94855293839020825160039092020180549383015160e01c600160a01b026001600160c01b03199094166001600160a01b0390921691909117929092178255604081015192820192909255606082015160028201906125c19082613c2b565b50505080600101905061251a565b505f5b84606001515181101561266a575f828152600760205260409020606086015180518390811061260357612603613279565b6020908102919091018101518254600181810185555f94855293839020825160029092020180546001600160a01b039283166001600160a01b0319909116178155928201516040909201516001600160601b0316600160a01b0291161790820155016125d2565b506126788160015f3361195b565b5050919050565b638b78c6d8600c52825f526020600c208054838117836126a0575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe265f80a3505050505050565b5f385f3884865af16107c85763b12d13eb5f526004601cfd5b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f51141661272f57803d853b15171061272f576390b8ec185f526004601cfd5b505f603452505050565b80516040820151606083810151909291905f805b83518110156127af575f84828151811061276957612769613279565b60200260200101516040015111156127a75783818151811061278d5761278d613279565b602002602001015160400151826127a491906133fc565b91505b60010161274d565b505f8082116127bf5782516127cc565b82516127cc9060016133fc565b6001600160401b038111156127e3576127e36132bc565b60405190808252806020026020018201604052801561283357816020015b604080516080810182525f8082526020808301829052928201819052606082015282525f199092019101816128015790505b5090505f5b835181101561291457604051806080016040528061288286848151811061286157612861613279565b6020026020010151602001516001600160a01b03166001600160a01b031690565b815260200185838151811061289957612899613279565b6020908102919091018101516040908101516001600160601b03168352898201516001600160401b03165f9081526003835220549101906001600160a01b0316815260200187602001516001600160401b031681525082828151811061290157612901613279565b6020908102919091010152600101612838565b50811561298b57604080516080810182525f8082526020808301869052888101516001600160401b03168252600390528290205490918201906001600160a01b0316815260200186602001516001600160401b03168152508184518151811061297f5761297f613279565b60200260200101819052505b9695505050505050565b60605f826020015190505f8082602001516001600160601b0316116129ba575f6129bd565b60015b60ff166001600160401b038111156129d7576129d76132bc565b604051908082528060200260200182016040528015612a2757816020015b604080516080810182525f8082526020808301829052928201819052606082015282525f199092019101816129f55790505b5060208301519091506001600160601b031615610fdd576040805160808101825283516001600160a01b031681526020808501516001600160601b0316908201525f918101829052466060820152825190918391612a8757612a87613279565b60200260200101819052509392505050565b80516040808301516060808501518351600180825281860190955291949390915f91816020015b60408051606080820183525f808352602083015291810191909152815260200190600190039081612ac05750506040805160608082018352602088810180516001600160401b039081168552815181165f90815260038452869020546001600160a01b031683860152855160a0810187524682168152915116818301528985015163ffffffff1681860152918201889052608082018790528351949550919392840192612b6d9201613784565b604051602081830303815290604052815250815f81518110612b9157612b91613279565b602090810291909101015295945050505050565b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f511416612bf057803d873b151710612bf057637939f4245f526004601cfd5b505f60605260405250505050565b600280545f91908290612c19906001600160f81b0316613cea565b91906101000a8154816001600160f81b0302191690836001600160f81b031602179055905090565b6040518061010001604052805f6001600160a01b031681526020015f81526020015f63ffffffff1681526020015f63ffffffff1681526020015f80191681526020016060815260200160608152602001606081525090565b5080545f8255600302905f5260205f20908101906107b39190612d20565b5080545f8255600202905f5260205f20908101906107b39190612d58565b6040805160e0810182525f6080820181815260a0830182905260c083018290528252825180840184528181526020808201929092529082015260609181018290528181019190915290565b80821115612d545780546001600160c01b03191681555f60018201819055612d4b6002830182612d7d565b50600301612d20565b5090565b5b80821115612d545780546001600160a01b03191681555f6001820155600201612d59565b508054612d8990613877565b5f825580601f10612d98575050565b601f0160209004905f5260205f20908101906107b391905b80821115612d54575f8155600101612db0565b5f8083601f840112612dd3575f80fd5b5081356001600160401b03811115612de9575f80fd5b6020830191508360208260051b8501011115612e03575f80fd5b9250929050565b5f805f8060408587031215612e1d575f80fd5b84356001600160401b0380821115612e33575f80fd5b612e3f88838901612dc3565b90965094506020870135915080821115612e57575f80fd5b50612e6487828801612dc3565b95989497509550505050565b5f60208284031215612e80575f80fd5b5035919050565b6001600160a01b03811681146107b3575f80fd5b5f8060408385031215612eac575f80fd5b8235612eb781612e87565b946020939093013593505050565b5f60208284031215612ed5575f80fd5b8135610fdd81612e87565b5f805f60608486031215612ef2575f80fd5b83359250602084013591506040840135612f0b81612e87565b809150509250925092565b5f60208284031215612f26575f80fd5b81356001600160401b03811115612f3b575f80fd5b820160608185031215610fdd575f80fd5b5f815180845260208085019450602084015f5b83811015612f9b578151805188528381015184890152604080820151908901526060908101519088015260809096019590820190600101612f5f565b509495945050505050565b5f5b83811015612fc0578181015183820152602001612fa8565b50505f910152565b5f8151808452612fdf816020860160208601612fa6565b601f01601f19169290920160200192915050565b5f82825180855260208086019550808260051b8401018186015f5b8481101561306557858303601f19018952815180516001600160401b03168452848101518585015260409081015160609185018290529061305181860183612fc8565b9a86019a945050509083019060010161300e565b5090979650505050505050565b5f61010060018060a01b0383511684526020830151602085015260408301516130a3604086018263ffffffff169052565b5060608301516130bb606086018263ffffffff169052565b506080830151608085015260a08301518160a08601526130dd82860182612f4c565b91505060c083015184820360c08601526130f78282612f4c565b91505060e083015184820360e08601526131118282612ff3565b95945050505050565b602081525f610fdd6020830184613072565b634e487b7160e01b5f52602160045260245ffd5b60c081525f61315260c0830186613072565b905083516006811061317257634e487b7160e01b5f52602160045260245ffd5b8060208401525060ff602085015116604083015263ffffffff604085015116606083015260018060a01b03606085015116608083015260018060f81b03831660a0830152949350505050565b80151581146107b3575f80fd5b5f602082840312156131db575f80fd5b8135610fdd816131be565b5f80604083850312156131f7575f80fd5b82359150602083013561320981612e87565b809150509250929050565b5f805f60608486031215613226575f80fd5b833561323181612e87565b9250602084013561324181612e87565b91506040840135612f0b81612e87565b5f8060408385031215613262575f80fd5b82359150602083013560ff81168114613209575f80fd5b634e487b7160e01b5f52603260045260245ffd5b6001600160401b03811681146107b3575f80fd5b5f602082840312156132b1575f80fd5b8135610fdd8161328d565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156132f2576132f26132bc565b60405290565b604051608081016001600160401b03811182821017156132f2576132f26132bc565b604051606081016001600160401b03811182821017156132f2576132f26132bc565b60405160a081016001600160401b03811182821017156132f2576132f26132bc565b604051601f8201601f191681016001600160401b0381118282101715613386576133866132bc565b604052919050565b5f6040828403121561339e575f80fd5b6133a66132d0565b82516133b18161328d565b815260208301516133c181612e87565b60208201529392505050565b5f602082840312156133dd575f80fd5b8151610fdd816131be565b634e487b7160e01b5f52601160045260245ffd5b808201808211156107e4576107e46133e8565b63ffffffff811681146107b3575f80fd5b805161342b8161340f565b919050565b5f6001600160401b03821115613448576134486132bc565b5060051b60200190565b6001600160e01b0319811681146107b3575f80fd5b5f6001600160401b0382111561347f5761347f6132bc565b50601f01601f191660200190565b5f82601f83011261349c575f80fd5b815160206134b16134ac83613430565b61335e565b82815260059290921b840181019181810190868411156134cf575f80fd5b8286015b848110156135ad5780516001600160401b03808211156134f1575f80fd5b908801906080828b03601f1901811315613509575f80fd5b6135116132f8565b8784015161351e81612e87565b815260408481015161352f81613452565b828a01526060858101518284015292850151928484111561354e575f80fd5b83860195508d603f870112613561575f80fd5b8986015194506135736134ac86613467565b93508484528d82868801011115613588575f80fd5b613597858b8601848901612fa6565b82019290925286525050509183019183016134d3565b509695505050505050565b6001600160601b03811681146107b3575f80fd5b5f82601f8301126135db575f80fd5b815160206135eb6134ac83613430565b82815260609283028501820192828201919087851115613609575f80fd5b8387015b858110156130655781818a031215613623575f80fd5b61362b61331a565b815161363681612e87565b81528186015161364581612e87565b81870152604082810151613658816135b8565b90820152845292840192810161360d565b5f60208284031215613679575f80fd5b81516001600160401b038082111561368f575f80fd5b9083019060a082860312156136a2575f80fd5b6136aa61333c565b82516136b58161328d565b815260208301516136c58161328d565b60208201526136d660408401613420565b60408201526060830151828111156136ec575f80fd5b6136f88782860161348d565b60608301525060808301518281111561370f575f80fd5b61371b878286016135cc565b60808301525095945050505050565b5f815180845260208085019450602084015f5b83811015612f9b57815180516001600160a01b0390811689528482015116848901526040908101516001600160601b0316908801526060909601959082019060010161373d565b5f602080835260c083016001600160401b0380865116838601528286015160408282166040880152604088015192506060915063ffffffff8316606088015260608801519250608060a0608089015284845180875260e08a01915060e08160051b8b0101965087860195505f5b81811015613853578a880360df19018352865180516001600160a01b03168952898101516001600160e01b0319168a8a015285810151868a0152860151868901859052613840858a0182612fc8565b98505095880195918801916001016137f1565b5050505050505060808501519150601f198482030160a0850152613111818361372a565b600181811c9082168061388b57607f821691505b6020821081036138a957634e487b7160e01b5f52602260045260245ffd5b50919050565b828152604060208201525f610e9b6040830184612fc8565b5f602082840312156138d7575f80fd5b8135610fdd8161340f565b5f808335601e198436030181126138f7575f80fd5b8301803591506001600160401b03821115613910575f80fd5b602001915036819003821315612e03575f80fd5b5f60408284031215613934575f80fd5b61393c6132d0565b9050813561394981612e87565b81526020820135613959816135b8565b602082015292915050565b5f82601f830112613973575f80fd5b813560206139836134ac83613430565b82815260059290921b840181019181810190868411156139a1575f80fd5b8286015b848110156135ad5780356001600160401b03808211156139c3575f80fd5b908801906080828b03601f19018113156139db575f80fd5b6139e36132f8565b878401356139f081612e87565b8152604084810135613a0181613452565b828a015260608581013582840152928501359284841115613a20575f80fd5b83860195508d603f870112613a33575f80fd5b898601359450613a456134ac86613467565b93508484528d82868801011115613a5a575f80fd5b848287018b8601375f9484018a019490945250918201528452509183019183016139a5565b5f82601f830112613a8e575f80fd5b81356020613a9e6134ac83613430565b82815260609283028501820192828201919087851115613abc575f80fd5b8387015b858110156130655781818a031215613ad6575f80fd5b613ade61331a565b8135613ae981612e87565b815281860135613af881612e87565b81870152604082810135613b0b816135b8565b908201528452928401928101613ac0565b5f60208284031215613b2c575f80fd5b81356001600160401b0380821115613b42575f80fd5b9083019060c08286031215613b55575f80fd5b613b5d61333c565b8235613b6881612e87565b81526020830135613b788161328d565b6020820152613b8a8660408501613924565b6040820152608083013582811115613ba0575f80fd5b613bac87828601613964565b60608301525060a083013582811115613bc3575f80fd5b61371b87828601613a7f565b5f60018201613be057613be06133e8565b5060010190565b601f82111561105457805f5260205f20601f840160051c81016020851015613c0c5750805b601f840160051c820191505b818110156107a2575f8155600101613c18565b81516001600160401b03811115613c4457613c446132bc565b613c5881613c528454613877565b84613be7565b602080601f831160018114613c8b575f8415613c745750858301515b5f19600386901b1c1916600185901b178555613ce2565b5f85815260208120601f198616915b82811015613cb957888601518255948401946001909101908401613c9a565b5085821015613cd657878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f6001600160f81b038281166002600160f81b03198101613d0d57613d0d6133e8565b600101939250505056fef76fe33b8a0ebf7aa05740f479d10138c7c15bdc75b10e047cc15be2be15e5b45ffb10051d79c19b9690b0842a292cb621fbf85d15269ed21c4e6a431d892bb5a26469706673582212201faff0d51727ee6e05f6f5364906d942798c3cec1a99eba8891d2f78b7e5444464736f6c63430008180033",
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

// GetLatestOrderOffset is a free data retrieval call binding the contract method 0xff446410.
//
// Solidity: function getLatestOrderOffset() view returns(uint248)
func (_SolverNetInbox *SolverNetInboxCaller) GetLatestOrderOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getLatestOrderOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestOrderOffset is a free data retrieval call binding the contract method 0xff446410.
//
// Solidity: function getLatestOrderOffset() view returns(uint248)
func (_SolverNetInbox *SolverNetInboxSession) GetLatestOrderOffset() (*big.Int, error) {
	return _SolverNetInbox.Contract.GetLatestOrderOffset(&_SolverNetInbox.CallOpts)
}

// GetLatestOrderOffset is a free data retrieval call binding the contract method 0xff446410.
//
// Solidity: function getLatestOrderOffset() view returns(uint248)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetLatestOrderOffset() (*big.Int, error) {
	return _SolverNetInbox.Contract.GetLatestOrderOffset(&_SolverNetInbox.CallOpts)
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

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 id) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolved, (uint8,uint8,uint32,address) state, uint248 offset)
func (_SolverNetInbox *SolverNetInboxCaller) GetOrder(opts *bind.CallOpts, id [32]byte) (struct {
	Resolved IERC7683ResolvedCrossChainOrder
	State    ISolverNetInboxOrderState
	Offset   *big.Int
}, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getOrder", id)

	outstruct := new(struct {
		Resolved IERC7683ResolvedCrossChainOrder
		State    ISolverNetInboxOrderState
		Offset   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolved = *abi.ConvertType(out[0], new(IERC7683ResolvedCrossChainOrder)).(*IERC7683ResolvedCrossChainOrder)
	outstruct.State = *abi.ConvertType(out[1], new(ISolverNetInboxOrderState)).(*ISolverNetInboxOrderState)
	outstruct.Offset = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 id) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolved, (uint8,uint8,uint32,address) state, uint248 offset)
func (_SolverNetInbox *SolverNetInboxSession) GetOrder(id [32]byte) (struct {
	Resolved IERC7683ResolvedCrossChainOrder
	State    ISolverNetInboxOrderState
	Offset   *big.Int
}, error) {
	return _SolverNetInbox.Contract.GetOrder(&_SolverNetInbox.CallOpts, id)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 id) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]) resolved, (uint8,uint8,uint32,address) state, uint248 offset)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetOrder(id [32]byte) (struct {
	Resolved IERC7683ResolvedCrossChainOrder
	State    ISolverNetInboxOrderState
	Offset   *big.Int
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

// SolverNetInboxFillOriginDataIterator is returned from FilterFillOriginData and is used to iterate over the raw logs and unpacked data for FillOriginData events raised by the SolverNetInbox contract.
type SolverNetInboxFillOriginDataIterator struct {
	Event *SolverNetInboxFillOriginData // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxFillOriginDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxFillOriginData)
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
		it.Event = new(SolverNetInboxFillOriginData)
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
func (it *SolverNetInboxFillOriginDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxFillOriginDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxFillOriginData represents a FillOriginData event raised by the SolverNetInbox contract.
type SolverNetInboxFillOriginData struct {
	Id             [32]byte
	FillOriginData SolverNetFillOriginData
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFillOriginData is a free log retrieval operation binding the contract event 0x71069e0637faca19b5cceb36f6ee664347f592a954e37edf597b6c0f37afd59f.
//
// Solidity: event FillOriginData(bytes32 indexed id, (uint64,uint64,uint32,(address,bytes4,uint256,bytes)[],(address,address,uint96)[]) fillOriginData)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterFillOriginData(opts *bind.FilterOpts, id [][32]byte) (*SolverNetInboxFillOriginDataIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "FillOriginData", idRule)
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxFillOriginDataIterator{contract: _SolverNetInbox.contract, event: "FillOriginData", logs: logs, sub: sub}, nil
}

// WatchFillOriginData is a free log subscription operation binding the contract event 0x71069e0637faca19b5cceb36f6ee664347f592a954e37edf597b6c0f37afd59f.
//
// Solidity: event FillOriginData(bytes32 indexed id, (uint64,uint64,uint32,(address,bytes4,uint256,bytes)[],(address,address,uint96)[]) fillOriginData)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchFillOriginData(opts *bind.WatchOpts, sink chan<- *SolverNetInboxFillOriginData, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "FillOriginData", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxFillOriginData)
				if err := _SolverNetInbox.contract.UnpackLog(event, "FillOriginData", log); err != nil {
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

// ParseFillOriginData is a log parse operation binding the contract event 0x71069e0637faca19b5cceb36f6ee664347f592a954e37edf597b6c0f37afd59f.
//
// Solidity: event FillOriginData(bytes32 indexed id, (uint64,uint64,uint32,(address,bytes4,uint256,bytes)[],(address,address,uint96)[]) fillOriginData)
func (_SolverNetInbox *SolverNetInboxFilterer) ParseFillOriginData(log types.Log) (*SolverNetInboxFillOriginData, error) {
	event := new(SolverNetInboxFillOriginData)
	if err := _SolverNetInbox.contract.UnpackLog(event, "FillOriginData", log); err != nil {
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
