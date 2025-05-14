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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mailbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PACKAGE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"close\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestOrderOffset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint248\",\"internalType\":\"uint248\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrder\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"resolved\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"state\",\"type\":\"tuple\",\"internalType\":\"structISolverNetInbox.OrderState\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISolverNetInbox.Status\"},{\"name\":\"rejectReason\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedBy\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"offset\",\"type\":\"uint248\",\"internalType\":\"uint248\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutbox\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserNonce\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"handle\",\"inputs\":[{\"name\":\"origin\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV2\",\"inputs\":[{\"name\":\"outbox\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"interchainSecurityModule\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIInterchainSecurityModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localDomain\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMailbox\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markFilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"open\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseClose\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseOpen\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reject\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOutboxes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboxes\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validate\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Closed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillOriginData\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillOriginData\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSolverNet.FillOriginData\",\"components\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structSolverNet.Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"expenses\",\"type\":\"tuple[]\",\"internalType\":\"structSolverNet.TokenExpense[]\",\"components\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Filled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HookSet\",\"inputs\":[{\"name\":\"_hook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IsmSet\",\"inputs\":[{\"name\":\"_ism\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Open\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"resolvedOrder\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboxSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Rejected\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenseAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenseToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFillDeadline\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMissingCalls\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNativeDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrderData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrderTypehash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReason\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderStillValid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PortalPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongFillHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongSourceChain\",\"inputs\":[]}]",
	Bin: "0x61012060405234801562000011575f80fd5b50604051620044523803806200445283398101604081905262000034916200025b565b808263ffffffff60643b1615620000bb5760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015620000a3575060408051601f3d908101601f19168201909252620000a09181019062000298565b60015b620000b25743608052620000c0565b608052620000c0565b436080525b6001600160a01b0390811660a052811660c081905215620001c95760c0516001600160a01b0316636e5f516e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200011a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001409190620002b0565b6001600160a01b031660e0816001600160a01b03168152505060c0516001600160a01b0316638d3638f46040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000198573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001be9190620002d5565b63ffffffff16610100525b50620001d4620001dc565b5050620002fa565b63409feecd1980546001811615620001fb5763f92ee8a95f526004601cfd5b6001600160401b03808260011c146200023e578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b6001600160a01b038116811462000258575f80fd5b50565b5f80604083850312156200026d575f80fd5b82516200027a8162000243565b60208401519092506200028d8162000243565b809150509250929050565b5f60208284031215620002a9575f80fd5b5051919050565b5f60208284031215620002c1575f80fd5b8151620002ce8162000243565b9392505050565b5f60208284031215620002e6575f80fd5b815163ffffffff81168114620002ce575f80fd5b60805160a05160c05160e051610100516140fe620003545f395f61054601525f6106e201525f818161063f015261112201525f818161036701528181610aa901528181610ad80152610ee101525f61072801526140fe5ff3fe608060405260043610610233575f3560e01c8063715018a611610129578063d9e8407c116100a8578063f04e283e1161006d578063f04e283e1461074a578063f2fde38b1461075d578063f904d28514610770578063fee81cf41461078f578063ff446410146107c0575f80fd5b8063d9e8407c14610693578063db3ea553146106b2578063de523cf3146106d1578063e917a96214610704578063eae4c19f14610717575f80fd5b806393c44847116100ee57806393c448471461059557806396c144f0146105d2578063b2b29439146105f1578063d5438eae1461062e578063d711835114610661575f80fd5b8063715018a6146104ef578063792aec5c146104f75780637cac41a6146105165780638d3638f4146105355780638da5cb5b1461057d575f80fd5b806339c79e0c116101b557806354d1f13d1161017a57806354d1f13d1461045357806356d5d4751461045b5780635778472a1461046e5780635e0aec951461049c5780636834e3a8146104bb575f80fd5b806339c79e0c146103a157806341b477dd146103c0578063485cc955146103ec5780634a4ee7b11461040b578063514e62fc1461041e575f80fd5b806329b6eca9116101fb57806329b6eca9146102ba57806329d54233146102d95780632d622343146103065780632de948071461032557806339acf9f114610356575f80fd5b806304a873ab14610237578063183a4f6e146102585780631c10893f1461026b5780631cd64df41461027e57806325692962146102b2575b5f80fd5b348015610242575f80fd5b506102566102513660046130f9565b6107e7565b005b61025661026636600461315f565b61093b565b61025661027936600461318a565b610948565b348015610289575f80fd5b5061029d61029836600461318a565b61095e565b60405190151581526020015b60405180910390f35b61025661097c565b3480156102c5575f80fd5b506102566102d43660046131b4565b6109c8565b3480156102e4575f80fd5b506102f86102f33660046131b4565b610a84565b6040519081526020016102a9565b348015610311575f80fd5b506102566103203660046131cf565b610aa7565b348015610330575f80fd5b506102f861033f3660046131b4565b638b78c6d8600c9081525f91909152602090205490565b348015610361575f80fd5b506103897f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102a9565b3480156103ac575f80fd5b506102566103bb36600461315f565b610c79565b3480156103cb575f80fd5b506103df6103da366004613205565b611008565b6040516102a99190613409565b3480156103f7575f80fd5b5061025661040636600461341b565b611057565b61025661041936600461318a565b6110cc565b348015610429575f80fd5b5061029d61043836600461318a565b638b78c6d8600c9081525f9290925260209091205416151590565b6102566110de565b610256610469366004613463565b611117565b348015610479575f80fd5b5061048d61048836600461315f565b611196565b6040516102a9939291906134f8565b3480156104a7575f80fd5b506102f86104b636600461318a565b61127a565b3480156104c6575f80fd5b506102f86104d53660046131b4565b6001600160a01b03165f908152600a602052604090205490565b61025661128c565b348015610502575f80fd5b50610256610511366004613583565b61129f565b348015610521575f80fd5b50610256610530366004613583565b6112c1565b348015610540575f80fd5b506105687f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102a9565b348015610588575f80fd5b50638b78c6d81954610389565b3480156105a0575f80fd5b506105c5604051806040016040528060058152602001640372e302e360dc1b81525081565b6040516102a9919061359e565b3480156105dd575f80fd5b506102566105ec3660046135b0565b6112fd565b3480156105fc575f80fd5b5061038961060b3660046135e7565b6001600160401b03165f908152600360205260409020546001600160a01b031690565b348015610639575f80fd5b506103897f000000000000000000000000000000000000000000000000000000000000000081565b34801561066c575f80fd5b5060025461068190600160f81b900460ff1681565b60405160ff90911681526020016102a9565b34801561069e575f80fd5b5061029d6106ad366004613205565b61146a565b3480156106bd575f80fd5b506102566106cc366004613602565b61147d565b3480156106dc575f80fd5b506103897f000000000000000000000000000000000000000000000000000000000000000081565b610256610712366004613205565b611604565b348015610722575f80fd5b506102f87f000000000000000000000000000000000000000000000000000000000000000081565b6102566107583660046131b4565b6117d3565b61025661076b3660046131b4565b61180d565b34801561077b575f80fd5b5061025661078a366004613583565b611833565b34801561079a575f80fd5b506102f86107a93660046131b4565b63389a75e1600c9081525f91909152602090205490565b3480156107cb575f80fd5b506002546040516001600160f81b0390911681526020016102a9565b6107ef611855565b82811461080f57604051634ec4810560e11b815260040160405180910390fd5b5f5b838110156109345782828281811061082b5761082b61362a565b905060200201602081019061084091906131b4565b60035f8787858181106108555761085561362a565b905060200201602081019061086a91906135e7565b6001600160401b0316815260208101919091526040015f2080546001600160a01b0319166001600160a01b03929092169190911790558282828181106108b2576108b261362a565b90506020020160208101906108c791906131b4565b6001600160a01b03168585838181106108e2576108e261362a565b90506020020160208101906108f791906135e7565b6001600160401b03167ff730978310b4a2a0e6c673324d737afdb93d0afefed14a3d061b60f66e31f4e360405160405180910390a3600101610811565b5050505050565b610945338261186f565b50565b610950611855565b61095a828261187a565b5050565b638b78c6d8600c9081525f8390526020902054811681145b92915050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f80a250565b63409feecd1980546004919082811060018216106109ed5763f92ee8a95f526004601cfd5b50600182178155466001600160401b03165f8181526003602052604080822080546001600160a01b0319166001600160a01b03881690811790915590519092917ff730978310b4a2a0e6c673324d737afdb93d0afefed14a3d061b60f66e31f4e391a38181558160011c6020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a1505050565b6001600160a01b0381165f908152600a6020526040812054610976908390611886565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615610c08577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa158015610b31573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b559190613710565b8051600180546020909301516001600160a01b0316600160401b026001600160e01b03199093166001600160401b039283161792909217918290555f911615610ba9576001546001600160401b0316610bab565b465b6001549091505f90600160401b90046001600160a01b031615610be057600154600160401b90046001600160a01b0316610be2565b335b9050610bf185858585856118c8565b5050600180546001600160e01b0319169055505050565b6001545f906001600160401b031615610c2c576001546001600160401b0316610c2e565b465b6001549091505f90600160401b90046001600160a01b031615610c6357600154600160401b90046001600160a01b0316610c65565b335b905061093485858585856118c8565b505050565b6002545f805160206140a983398151915290600160f81b900460ff168015610d3a5760ff81166001148015610cba57505f8051602061408983398151915282145b15610cd857604051631309a56360e01b815260040160405180910390fd5b60ff81166002148015610cf757505f805160206140a983398151915282145b15610d1557604051631309a56360e01b815260040160405180910390fd5b60021960ff821601610d3a5760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd212685403610d585763ab143c065f526004601cfd5b3068929eee149b4bd21268555f838152600860205260408082208151608081019092528054829060ff166005811115610d9357610d936134e4565b6005811115610da457610da46134e4565b8152905460ff61010082041660208084019190915263ffffffff62010000830481166040808601919091526001600160a01b03600160301b90940484166060958601525f8a815260048452818120825196870183525494851686526001600160401b03600160a01b860416938601849052600160e01b9094049091169084015292935090914614610e3757615460610e39565b5f5b9050600183516005811115610e5057610e506134e4565b14610e6e57604051635d12a4a360e11b815260040160405180910390fd5b81516001600160a01b03163314610e97576040516282b42960e81b815260040160405180910390fd5b60208201516040516308c3569160e31b81527ffeccba1cfc4544bf9cd83b76f36ae5c464750b6c43f682e26744ee21ec31fc1e60048201526001600160401b0390911660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063461ab48890604401602060405180830381865afa158015610f2e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f52919061374f565b15610f7057604051630c2e605760e11b815260040160405180910390fd5b4281836040015163ffffffff16610f87919061377e565b10610fa5576040516321bb6b2160e11b815260040160405180910390fd5b610fb28660035f33611ab9565b610fbf86835f0151611bcd565b610fca866003611c66565b60405186907f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79905f90a25050503868929eee149b4bd2126855505050565b611010612f30565b5f61101a83611cfd565b8051516001600160a01b0381165f908152600a60205260409020549192509061104f90839061104a908490611886565b611f90565b949350505050565b63409feecd19805460049190828110600182161061107c5763f92ee8a95f526004601cfd5b5081600117815561108c8461200f565b61109783600161187a565b8181558160011c6020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a150505050565b6110d4611855565b61095a828261186f565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f80a2565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461115f576040516282b42960e81b815260040160405180910390fd5b5f808061116e848601866131cf565b92509250925061118d8383838a63ffffffff166111888b90565b6118c8565b50505050505050565b61119e612f30565b604080516080810182525f808252602082018190529181018290526060810182905290806111cb8561204a565b90506111d78186611f90565b5f86815260086020908152604080832060099092529182902054825160808101909352815491926001600160f81b03909116918390829060ff166005811115611222576112226134e4565b6005811115611233576112336134e4565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015292989297509550909350505050565b5f6112858383611886565b9392505050565b611294611855565b61129d5f6122a0565b565b60016112aa816122dd565b61095a5f805160206140a98339815191528361230e565b60016112cc816122dd565b816112e457600280546001600160f81b031690555050565b5050600280546001600160f81b0316600360f81b179055565b3068929eee149b4bd21268540361131b5763ab143c065f526004601cfd5b3068929eee149b4bd21268555f828152600860205260408082208151608081019092528054829060ff166005811115611356576113566134e4565b6005811115611367576113676134e4565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015290506004815160058111156113b6576113b66134e4565b146113d45760405163789bae3560e01b815260040160405180910390fd5b60608101516001600160a01b03163314611400576040516282b42960e81b815260040160405180910390fd5b61140d8360055f33611ab9565b6114178383611bcd565b611422836005611c66565b6040516001600160a01b03831690339085907f8428df912f4f2125b442b488df9c7260cb607246895bcd29f262ecca090b1538905f90a4503868929eee149b4bd21268555050565b5f61147482611cfd565b50600192915050565b6001611488816123c7565b3068929eee149b4bd2126854036114a65763ab143c065f526004601cfd5b3068929eee149b4bd21268555f838152600860205260408082208151608081019092528054829060ff1660058111156114e1576114e16134e4565b60058111156114f2576114f26134e4565b81529054610100810460ff908116602084015262010000820463ffffffff166040840152600160301b9091046001600160a01b031660609092019190915290915083165f03611554576040516337b89b9360e21b815260040160405180910390fd5b600181516005811115611569576115696134e4565b1461158757604051635d12a4a360e11b815260040160405180910390fd5b6115948460028533611ab9565b5f848152600460205260409020546115b69085906001600160a01b0316611bcd565b6115c1846002611c66565b60405160ff841690339086907f21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f905f90a4503868929eee149b4bd2126855505050565b6002545f8051602061408983398151915290600160f81b900460ff1680156116c55760ff8116600114801561164557505f8051602061408983398151915282145b1561166357604051631309a56360e01b815260040160405180910390fd5b60ff8116600214801561168257505f805160206140a983398151915282145b156116a057604051631309a56360e01b815260040160405180910390fd5b60021960ff8216016116c55760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd2126854036116e35763ab143c065f526004601cfd5b3068929eee149b4bd21268555f6116f984611cfd565b905061170881602001516123eb565b5f61171282612450565b905080608001517f71069e0637faca19b5cceb36f6ee664347f592a954e37edf597b6c0f37afd59f8260e001515f815181106117505761175061362a565b60200260200101516040015180602001905181019061176f91906139da565b60405161177c9190613af5565b60405180910390a280608001517fa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c826040516117b89190613409565b60405180910390a250503868929eee149b4bd2126855505050565b6117db611855565b63389a75e1600c52805f526020600c20805442111561180157636f5e88185f526004601cfd5b5f9055610945816122a0565b611815611855565b8060601b61182a57637448fbae5f526004601cfd5b610945816122a0565b600161183e816122dd565b61095a5f805160206140898339815191528361230e565b638b78c6d81954331461129d576382b429005f526004601cfd5b61095a82825f6126c6565b61095a828260016126c6565b604080516001600160a01b03841660208201529081018290524660608201525f9060800160405160208183030381529060405280519060200120905092915050565b5f858152600460209081526040808320815160608101835290546001600160a01b0381168252600160a01b81046001600160401b031682850152600160e01b900463ffffffff168183015288845260089092528083208151608081019092528054929392829060ff166005811115611942576119426134e4565b6005811115611953576119536134e4565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015290506001815160058111156119a2576119a26134e4565b146119c057604051635d12a4a360e11b815260040160405180910390fd5b81602001516001600160401b0316846001600160401b0316146119f657604051633687f39960e21b815260040160405180910390fd5b6001600160401b0384165f908152600360205260409020546001600160a01b03848116911614611a38576040516282b42960e81b815260040160405180910390fd5b611a418761271d565b8614611a6057604051631f53eaed60e21b815260040160405180910390fd5b611a6d8760045f88611ab9565b611a78876004611c66565b846001600160a01b031686887fa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc60405160405180910390a450505050505050565b5f848152600860205260409081902054815160808101909252610100900460ff169080856005811115611aee57611aee6134e4565b81526020015f8560ff1611611b035782611b05565b845b60ff16815263ffffffff42166020808301919091526001600160a01b0385166040928301525f88815260089091522081518154829060ff19166001836005811115611b5257611b526134e4565b02179055506020820151815460408401516060909401516001600160a01b0316600160301b026601000000000000600160d01b031963ffffffff909516620100000265ffffffff00001960ff909416610100029390931665ffffffffff00199092169190911791909117929092169190911790555050505050565b5f828152600560209081526040918290208251808401909352546001600160a01b0381168352600160a01b90046001600160601b031690820181905215610c745780516001600160a01b0316611c3f576020810151610c74906001600160a01b038416906001600160601b03166129c5565b60208101518151610c74916001600160a01b039091169084906001600160601b03166129de565b6001816005811115611c7a57611c7a6134e4565b03611c83575050565b6004816005811115611c9757611c976134e4565b14611cab575f828152600560205260408120555b6005816005811115611cbf57611cbf6134e4565b1461095a575f82815260046020908152604080832083905560069091528120611ce791612f88565b5f82815260076020526040812061095a91612fa6565b611d05612fc4565b42611d136020840184613be8565b63ffffffff1611611d375760405163582e388960e01b815260040160405180910390fd5b7f2e7de755ca70cb933dc80103af16cc3303580e5712f1a8927d6461441e99a1e6826020013514611d7b57604051636aea87f360e01b815260040160405180910390fd5b611d886040830183613c03565b90505f03611da95760405163a342e7d960e01b815260040160405180910390fd5b5f611db76040840184613c03565b810190611dc49190613e3d565b80519091506001600160a01b0316611dda573381525b80602001516001600160401b03165f03611e0757604051633d23e4d160e11b815260040160405180910390fd5b6040805160608101825282516001600160a01b031681526020808401516001600160401b0316818301525f92820190611e4290870187613be8565b63ffffffff16905260608301518051919250905f03611e7457604051639cc71f7d60e01b815260040160405180910390fd5b805160201015611e9757604051634ec4810560e11b815260040160405180910390fd5b6080830151805160201015611ebf57604051634ec4810560e11b815260040160405180910390fd5b5f5b8151811015611f67575f6001600160a01b0316828281518110611ee657611ee661362a565b6020026020010151602001516001600160a01b031603611f195760405163027dcfa160e31b815260040160405180910390fd5b818181518110611f2b57611f2b61362a565b6020026020010151604001516001600160601b03165f03611f5f5760405163a0ce339960e01b815260040160405180910390fd5b600101611ec1565b506040805160808101825293845293840151602084015292820152606081019190915292915050565b611f98612f30565b82515f611fa485612a28565b90505f611fb086612c84565b90505f611fbc87612d88565b604080516101008101825286516001600160a01b031681524660208201525f8183015295015163ffffffff1660608601526080850187905260a08501939093525060c083015260e0820152905092915050565b6001600160a01b0316638b78c6d819819055805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b612052612fc4565b604080515f8481526004602090815283822060e084018552546001600160a01b0380821660808601908152600160a01b8084046001600160401b031660a0880152600160e01b90930463ffffffff1660c087015285528784526005835285842086518088018852905491821681529190046001600160601b031681830152818401528582526006815283822080548551818402810184018752818152949586019493919290919084015b828210156121f8575f848152602090819020604080516080810182526003860290920180546001600160a01b0381168452600160a01b900460e01b6001600160e01b0319169383019390935260018301549082015260028201805491929160608401919061216990613ef0565b80601f016020809104026020016040519081016040528092919081815260200182805461219590613ef0565b80156121e05780601f106121b7576101008083540402835291602001916121e0565b820191905f5260205f20905b8154815290600101906020018083116121c357829003601f168201915b505050505081525050815260200190600101906120fc565b50505050815260200160075f8581526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015612294575f848152602090819020604080516060810182526002860290920180546001600160a01b03908116845260019182015490811684860152600160a01b90046001600160601b0316918301919091529083529092019101612231565b50505091525092915050565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a355565b638b78c6d81954331461094557638b78c6d8600c52335f52806020600c205416610945576382b429005f526004601cfd5b600254600160f81b900460ff16600219810161233d5760405163aaae8ef760e01b815260040160405180910390fd5b5f5f80516020614089833981519152841461235957600261235c565b60015b905082612372578060ff168260ff16141561237c565b8060ff168260ff16145b1561239a57604051631309a56360e01b815260040160405180910390fd5b826123a5575f6123a7565b805b6002601f6101000a81548160ff021916908360ff16021790555050505050565b638b78c6d8600c52335f52806020600c205416610945576382b429005f526004601cfd5b80516001600160a01b03166124275780602001516001600160601b031634146109455760405163036f810f60e41b815260040160405180910390fd5b60208101518151610945916001600160a01b0390911690339030906001600160601b0316612e94565b612458612f30565b8151516001600160a01b0381165f908152600a60205260408120805461248e918491908461248583613f28565b91905055611886565b905061249a8482611f90565b84515f8381526004602090815260408083208451815484870151968401516001600160a01b039283166001600160e01b031990921691909117600160a01b6001600160401b039098168802176001600160e01b0316600160e01b63ffffffff9092169190910217909155828a015160058452919093208151919092015192166001600160601b039092169092021790559250612534612eed565b5f82815260096020526040812080546001600160f81b0319166001600160f81b0393909316929092179091555b846040015151811015612616575f8281526006602052604090819020908601518051839081106125935761259361362a565b6020908102919091018101518254600181810185555f94855293839020825160039092020180549383015160e01c600160a01b026001600160c01b03199094166001600160a01b0390921691909117929092178255604081015192820192909255606082015160028201906126089082613f84565b505050806001019050612561565b505f5b8460600151518110156126b1575f828152600760205260409020606086015180518390811061264a5761264a61362a565b6020908102919091018101518254600181810185555f94855293839020825160029092020180546001600160a01b039283166001600160a01b0319909116178155928201516040909201516001600160601b0316600160a01b029116179082015501612619565b506126bf8160015f33611ab9565b5050919050565b638b78c6d8600c52825f526020600c208054838117836126e7575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe265f80a3505050505050565b5f818152600460209081526040808320815160608101835290546001600160a01b0381168252600160a01b81046001600160401b031682850152600160e01b900463ffffffff1681830152848452600683528184208054835181860281018601909452808452919385939290849084015b8282101561288a575f848152602090819020604080516080810182526003860290920180546001600160a01b0381168452600160a01b900460e01b6001600160e01b031916938301939093526001830154908201526002820180549192916060840191906127fb90613ef0565b80601f016020809104026020016040519081016040528092919081815260200182805461282790613ef0565b80156128725780601f1061284957610100808354040283529160200191612872565b820191905f5260205f20905b81548152906001019060200180831161285557829003601f168201915b5050505050815250508152602001906001019061278e565b5050505090505f60075f8681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015612924575f848152602090819020604080516060810182526002860290920180546001600160a01b03908116845260019182015490811684860152600160a01b90046001600160601b03169183019190915290835290920191016128c1565b5050505090505f6040518060a00160405280466001600160401b0316815260200185602001516001600160401b03168152602001856040015163ffffffff16815260200184815260200183815250905085816040516020016129869190613af5565b60408051601f19818403018152908290526129a49291602001614043565b60405160208183030381529060405280519060200120945050505050919050565b5f385f3884865af161095a5763b12d13eb5f526004601cfd5b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f511416612a1e57803d853b151710612a1e576390b8ec185f526004601cfd5b505f603452505050565b80516040820151606083810151909291905f805b8351811015612a9e575f848281518110612a5857612a5861362a565b6020026020010151604001511115612a9657838181518110612a7c57612a7c61362a565b60200260200101516040015182612a93919061377e565b91505b600101612a3c565b505f808211612aae578251612abb565b8251612abb90600161377e565b6001600160401b03811115612ad257612ad261363e565b604051908082528060200260200182016040528015612b2257816020015b604080516080810182525f8082526020808301829052928201819052606082015282525f19909201910181612af05790505b5090505f5b8351811015612c03576040518060800160405280612b71868481518110612b5057612b5061362a565b6020026020010151602001516001600160a01b03166001600160a01b031690565b8152602001858381518110612b8857612b8861362a565b6020908102919091018101516040908101516001600160601b03168352898201516001600160401b03165f9081526003835220549101906001600160a01b0316815260200187602001516001600160401b0316815250828281518110612bf057612bf061362a565b6020908102919091010152600101612b27565b508115612c7a57604080516080810182525f8082526020808301869052888101516001600160401b03168252600390528290205490918201906001600160a01b0316815260200186602001516001600160401b031681525081845181518110612c6e57612c6e61362a565b60200260200101819052505b9695505050505050565b60605f826020015190505f8082602001516001600160601b031611612ca9575f612cac565b60015b60ff166001600160401b03811115612cc657612cc661363e565b604051908082528060200260200182016040528015612d1657816020015b604080516080810182525f8082526020808301829052928201819052606082015282525f19909201910181612ce45790505b5060208301519091506001600160601b031615611285576040805160808101825283516001600160a01b031681526020808501516001600160601b0316908201525f918101829052466060820152825190918391612d7657612d7661362a565b60200260200101819052509392505050565b80516040808301516060808501518351600180825281860190955291949390915f91816020015b60408051606080820183525f808352602083015291810191909152815260200190600190039081612daf5750506040805160608082018352602088810180516001600160401b039081168552815181165f90815260038452869020546001600160a01b031683860152855160a0810187524682168152915116818301528985015163ffffffff1681860152918201889052608082018790528351949550919392840192612e5c9201613af5565b604051602081830303815290604052815250815f81518110612e8057612e8061362a565b602090810291909101015295945050505050565b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f511416612edf57803d873b151710612edf57637939f4245f526004601cfd5b505f60605260405250505050565b600280545f91908290612f08906001600160f81b031661405b565b91906101000a8154816001600160f81b0302191690836001600160f81b031602179055905090565b6040518061010001604052805f6001600160a01b031681526020015f81526020015f63ffffffff1681526020015f63ffffffff1681526020015f80191681526020016060815260200160608152602001606081525090565b5080545f8255600302905f5260205f2090810190610945919061300f565b5080545f8255600202905f5260205f20908101906109459190613047565b6040805160e0810182525f6080820181815260a0830182905260c083018290528252825180840184528181526020808201929092529082015260609181018290528181019190915290565b808211156130435780546001600160c01b03191681555f6001820181905561303a600283018261306c565b5060030161300f565b5090565b5b808211156130435780546001600160a01b03191681555f6001820155600201613048565b50805461307890613ef0565b5f825580601f10613087575050565b601f0160209004905f5260205f209081019061094591905b80821115613043575f815560010161309f565b5f8083601f8401126130c2575f80fd5b5081356001600160401b038111156130d8575f80fd5b6020830191508360208260051b85010111156130f2575f80fd5b9250929050565b5f805f806040858703121561310c575f80fd5b84356001600160401b0380821115613122575f80fd5b61312e888389016130b2565b90965094506020870135915080821115613146575f80fd5b50613153878288016130b2565b95989497509550505050565b5f6020828403121561316f575f80fd5b5035919050565b6001600160a01b0381168114610945575f80fd5b5f806040838503121561319b575f80fd5b82356131a681613176565b946020939093013593505050565b5f602082840312156131c4575f80fd5b813561128581613176565b5f805f606084860312156131e1575f80fd5b833592506020840135915060408401356131fa81613176565b809150509250925092565b5f60208284031215613215575f80fd5b81356001600160401b0381111561322a575f80fd5b820160608185031215611285575f80fd5b5f815180845260208085019450602084015f5b8381101561328a57815180518852838101518489015260408082015190890152606090810151908801526080909601959082019060010161324e565b509495945050505050565b5f5b838110156132af578181015183820152602001613297565b50505f910152565b5f81518084526132ce816020860160208601613295565b601f01601f19169290920160200192915050565b5f82825180855260208086019550808260051b8401018186015f5b8481101561335457858303601f19018952815180516001600160401b031684528481015185850152604090810151606091850182905290613340818601836132b7565b9a86019a94505050908301906001016132fd565b5090979650505050505050565b5f61010060018060a01b038351168452602083015160208501526040830151613392604086018263ffffffff169052565b5060608301516133aa606086018263ffffffff169052565b506080830151608085015260a08301518160a08601526133cc8286018261323b565b91505060c083015184820360c08601526133e6828261323b565b91505060e083015184820360e086015261340082826132e2565b95945050505050565b602081525f6112856020830184613361565b5f806040838503121561342c575f80fd5b823561343781613176565b9150602083013561344781613176565b809150509250929050565b63ffffffff81168114610945575f80fd5b5f805f8060608587031215613476575f80fd5b843561348181613452565b93506020850135925060408501356001600160401b03808211156134a3575f80fd5b818701915087601f8301126134b6575f80fd5b8135818111156134c4575f80fd5b8860208285010111156134d5575f80fd5b95989497505060200194505050565b634e487b7160e01b5f52602160045260245ffd5b60c081525f61350a60c0830186613361565b905083516006811061352a57634e487b7160e01b5f52602160045260245ffd5b8060208401525060ff602085015116604083015263ffffffff604085015116606083015260018060a01b03606085015116608083015260018060f81b03831660a0830152949350505050565b8015158114610945575f80fd5b5f60208284031215613593575f80fd5b813561128581613576565b602081525f61128560208301846132b7565b5f80604083850312156135c1575f80fd5b82359150602083013561344781613176565b6001600160401b0381168114610945575f80fd5b5f602082840312156135f7575f80fd5b8135611285816135d3565b5f8060408385031215613613575f80fd5b82359150602083013560ff81168114613447575f80fd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156136745761367461363e565b60405290565b604051608081016001600160401b03811182821017156136745761367461363e565b604051606081016001600160401b03811182821017156136745761367461363e565b60405160a081016001600160401b03811182821017156136745761367461363e565b604051601f8201601f191681016001600160401b03811182821017156137085761370861363e565b604052919050565b5f60408284031215613720575f80fd5b613728613652565b8251613733816135d3565b8152602083015161374381613176565b60208201529392505050565b5f6020828403121561375f575f80fd5b815161128581613576565b634e487b7160e01b5f52601160045260245ffd5b808201808211156109765761097661376a565b805161379c81613452565b919050565b5f6001600160401b038211156137b9576137b961363e565b5060051b60200190565b6001600160e01b031981168114610945575f80fd5b5f6001600160401b038211156137f0576137f061363e565b50601f01601f191660200190565b5f82601f83011261380d575f80fd5b8151602061382261381d836137a1565b6136e0565b82815260059290921b84018101918181019086841115613840575f80fd5b8286015b8481101561391e5780516001600160401b0380821115613862575f80fd5b908801906080828b03601f190181131561387a575f80fd5b61388261367a565b8784015161388f81613176565b81526040848101516138a0816137c3565b828a0152606085810151828401529285015192848411156138bf575f80fd5b83860195508d603f8701126138d2575f80fd5b8986015194506138e461381d866137d8565b93508484528d828688010111156138f9575f80fd5b613908858b8601848901613295565b8201929092528652505050918301918301613844565b509695505050505050565b6001600160601b0381168114610945575f80fd5b5f82601f83011261394c575f80fd5b8151602061395c61381d836137a1565b8281526060928302850182019282820191908785111561397a575f80fd5b8387015b858110156133545781818a031215613994575f80fd5b61399c61369c565b81516139a781613176565b8152818601516139b681613176565b818701526040828101516139c981613929565b90820152845292840192810161397e565b5f602082840312156139ea575f80fd5b81516001600160401b0380821115613a00575f80fd5b9083019060a08286031215613a13575f80fd5b613a1b6136be565b8251613a26816135d3565b81526020830151613a36816135d3565b6020820152613a4760408401613791565b6040820152606083015182811115613a5d575f80fd5b613a69878286016137fe565b606083015250608083015182811115613a80575f80fd5b613a8c8782860161393d565b60808301525095945050505050565b5f815180845260208085019450602084015f5b8381101561328a57815180516001600160a01b0390811689528482015116848901526040908101516001600160601b03169088015260609096019590820190600101613aae565b5f602080835260c083016001600160401b0380865116838601528286015160408282166040880152604088015192506060915063ffffffff8316606088015260608801519250608060a0608089015284845180875260e08a01915060e08160051b8b0101965087860195505f5b81811015613bc4578a880360df19018352865180516001600160a01b03168952898101516001600160e01b0319168a8a015285810151868a0152860151868901859052613bb1858a01826132b7565b9850509588019591880191600101613b62565b5050505050505060808501519150601f198482030160a08501526134008183613a9b565b5f60208284031215613bf8575f80fd5b813561128581613452565b5f808335601e19843603018112613c18575f80fd5b8301803591506001600160401b03821115613c31575f80fd5b6020019150368190038213156130f2575f80fd5b5f60408284031215613c55575f80fd5b613c5d613652565b90508135613c6a81613176565b81526020820135613c7a81613929565b602082015292915050565b5f82601f830112613c94575f80fd5b81356020613ca461381d836137a1565b82815260059290921b84018101918181019086841115613cc2575f80fd5b8286015b8481101561391e5780356001600160401b0380821115613ce4575f80fd5b908801906080828b03601f1901811315613cfc575f80fd5b613d0461367a565b87840135613d1181613176565b8152604084810135613d22816137c3565b828a015260608581013582840152928501359284841115613d41575f80fd5b83860195508d603f870112613d54575f80fd5b898601359450613d6661381d866137d8565b93508484528d82868801011115613d7b575f80fd5b848287018b8601375f9484018a01949094525091820152845250918301918301613cc6565b5f82601f830112613daf575f80fd5b81356020613dbf61381d836137a1565b82815260609283028501820192828201919087851115613ddd575f80fd5b8387015b858110156133545781818a031215613df7575f80fd5b613dff61369c565b8135613e0a81613176565b815281860135613e1981613176565b81870152604082810135613e2c81613929565b908201528452928401928101613de1565b5f60208284031215613e4d575f80fd5b81356001600160401b0380821115613e63575f80fd5b9083019060c08286031215613e76575f80fd5b613e7e6136be565b8235613e8981613176565b81526020830135613e99816135d3565b6020820152613eab8660408501613c45565b6040820152608083013582811115613ec1575f80fd5b613ecd87828601613c85565b60608301525060a083013582811115613ee4575f80fd5b613a8c87828601613da0565b600181811c90821680613f0457607f821691505b602082108103613f2257634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60018201613f3957613f3961376a565b5060010190565b601f821115610c7457805f5260205f20601f840160051c81016020851015613f655750805b601f840160051c820191505b81811015610934575f8155600101613f71565b81516001600160401b03811115613f9d57613f9d61363e565b613fb181613fab8454613ef0565b84613f40565b602080601f831160018114613fe4575f8415613fcd5750858301515b5f19600386901b1c1916600185901b17855561403b565b5f85815260208120601f198616915b8281101561401257888601518255948401946001909101908401613ff3565b508582101561402f57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b828152604060208201525f61104f60408301846132b7565b5f6001600160f81b038281166002600160f81b0319810161407e5761407e61376a565b600101939250505056fef76fe33b8a0ebf7aa05740f479d10138c7c15bdc75b10e047cc15be2be15e5b45ffb10051d79c19b9690b0842a292cb621fbf85d15269ed21c4e6a431d892bb5a2646970667358221220be221a7c85837bea3c3cd17e7ff226a1cf0d4fa57743be17026de32e4f3e6b5564736f6c63430008180033",
}

// SolverNetInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SolverNetInboxMetaData.ABI instead.
var SolverNetInboxABI = SolverNetInboxMetaData.ABI

// SolverNetInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolverNetInboxMetaData.Bin instead.
var SolverNetInboxBin = SolverNetInboxMetaData.Bin

// DeploySolverNetInbox deploys a new Ethereum contract, binding an instance of SolverNetInbox to it.
func DeploySolverNetInbox(auth *bind.TransactOpts, backend bind.ContractBackend, omni_ common.Address, mailbox_ common.Address) (common.Address, *types.Transaction, *SolverNetInbox, error) {
	parsed, err := SolverNetInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolverNetInboxBin), backend, omni_, mailbox_)
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

// PACKAGEVERSION is a free data retrieval call binding the contract method 0x93c44847.
//
// Solidity: function PACKAGE_VERSION() view returns(string)
func (_SolverNetInbox *SolverNetInboxCaller) PACKAGEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "PACKAGE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PACKAGEVERSION is a free data retrieval call binding the contract method 0x93c44847.
//
// Solidity: function PACKAGE_VERSION() view returns(string)
func (_SolverNetInbox *SolverNetInboxSession) PACKAGEVERSION() (string, error) {
	return _SolverNetInbox.Contract.PACKAGEVERSION(&_SolverNetInbox.CallOpts)
}

// PACKAGEVERSION is a free data retrieval call binding the contract method 0x93c44847.
//
// Solidity: function PACKAGE_VERSION() view returns(string)
func (_SolverNetInbox *SolverNetInboxCallerSession) PACKAGEVERSION() (string, error) {
	return _SolverNetInbox.Contract.PACKAGEVERSION(&_SolverNetInbox.CallOpts)
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

// GetOutbox is a free data retrieval call binding the contract method 0xb2b29439.
//
// Solidity: function getOutbox(uint64 chainId) view returns(address)
func (_SolverNetInbox *SolverNetInboxCaller) GetOutbox(opts *bind.CallOpts, chainId uint64) (common.Address, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getOutbox", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOutbox is a free data retrieval call binding the contract method 0xb2b29439.
//
// Solidity: function getOutbox(uint64 chainId) view returns(address)
func (_SolverNetInbox *SolverNetInboxSession) GetOutbox(chainId uint64) (common.Address, error) {
	return _SolverNetInbox.Contract.GetOutbox(&_SolverNetInbox.CallOpts, chainId)
}

// GetOutbox is a free data retrieval call binding the contract method 0xb2b29439.
//
// Solidity: function getOutbox(uint64 chainId) view returns(address)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetOutbox(chainId uint64) (common.Address, error) {
	return _SolverNetInbox.Contract.GetOutbox(&_SolverNetInbox.CallOpts, chainId)
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

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_SolverNetInbox *SolverNetInboxCaller) InterchainSecurityModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "interchainSecurityModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_SolverNetInbox *SolverNetInboxSession) InterchainSecurityModule() (common.Address, error) {
	return _SolverNetInbox.Contract.InterchainSecurityModule(&_SolverNetInbox.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_SolverNetInbox *SolverNetInboxCallerSession) InterchainSecurityModule() (common.Address, error) {
	return _SolverNetInbox.Contract.InterchainSecurityModule(&_SolverNetInbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SolverNetInbox *SolverNetInboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SolverNetInbox *SolverNetInboxSession) LocalDomain() (uint32, error) {
	return _SolverNetInbox.Contract.LocalDomain(&_SolverNetInbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SolverNetInbox *SolverNetInboxCallerSession) LocalDomain() (uint32, error) {
	return _SolverNetInbox.Contract.LocalDomain(&_SolverNetInbox.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_SolverNetInbox *SolverNetInboxCaller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_SolverNetInbox *SolverNetInboxSession) Mailbox() (common.Address, error) {
	return _SolverNetInbox.Contract.Mailbox(&_SolverNetInbox.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_SolverNetInbox *SolverNetInboxCallerSession) Mailbox() (common.Address, error) {
	return _SolverNetInbox.Contract.Mailbox(&_SolverNetInbox.CallOpts)
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

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 origin, bytes32 sender, bytes message) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactor) Handle(opts *bind.TransactOpts, origin uint32, sender [32]byte, message []byte) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "handle", origin, sender, message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 origin, bytes32 sender, bytes message) payable returns()
func (_SolverNetInbox *SolverNetInboxSession) Handle(origin uint32, sender [32]byte, message []byte) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Handle(&_SolverNetInbox.TransactOpts, origin, sender, message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 origin, bytes32 sender, bytes message) payable returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) Handle(origin uint32, sender [32]byte, message []byte) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Handle(&_SolverNetInbox.TransactOpts, origin, sender, message)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address solver_) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, solver_ common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "initialize", owner_, solver_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address solver_) returns()
func (_SolverNetInbox *SolverNetInboxSession) Initialize(owner_ common.Address, solver_ common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Initialize(&_SolverNetInbox.TransactOpts, owner_, solver_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address solver_) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) Initialize(owner_ common.Address, solver_ common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.Initialize(&_SolverNetInbox.TransactOpts, owner_, solver_)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address outbox) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) InitializeV2(opts *bind.TransactOpts, outbox common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "initializeV2", outbox)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address outbox) returns()
func (_SolverNetInbox *SolverNetInboxSession) InitializeV2(outbox common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.InitializeV2(&_SolverNetInbox.TransactOpts, outbox)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address outbox) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) InitializeV2(outbox common.Address) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.InitializeV2(&_SolverNetInbox.TransactOpts, outbox)
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

// SolverNetInboxHookSetIterator is returned from FilterHookSet and is used to iterate over the raw logs and unpacked data for HookSet events raised by the SolverNetInbox contract.
type SolverNetInboxHookSetIterator struct {
	Event *SolverNetInboxHookSet // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxHookSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxHookSet)
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
		it.Event = new(SolverNetInboxHookSet)
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
func (it *SolverNetInboxHookSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxHookSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxHookSet represents a HookSet event raised by the SolverNetInbox contract.
type SolverNetInboxHookSet struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookSet is a free log retrieval operation binding the contract event 0x4eab7b127c764308788622363ad3e9532de3dfba7845bd4f84c125a22544255a.
//
// Solidity: event HookSet(address _hook)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterHookSet(opts *bind.FilterOpts) (*SolverNetInboxHookSetIterator, error) {

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "HookSet")
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxHookSetIterator{contract: _SolverNetInbox.contract, event: "HookSet", logs: logs, sub: sub}, nil
}

// WatchHookSet is a free log subscription operation binding the contract event 0x4eab7b127c764308788622363ad3e9532de3dfba7845bd4f84c125a22544255a.
//
// Solidity: event HookSet(address _hook)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchHookSet(opts *bind.WatchOpts, sink chan<- *SolverNetInboxHookSet) (event.Subscription, error) {

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "HookSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxHookSet)
				if err := _SolverNetInbox.contract.UnpackLog(event, "HookSet", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseHookSet(log types.Log) (*SolverNetInboxHookSet, error) {
	event := new(SolverNetInboxHookSet)
	if err := _SolverNetInbox.contract.UnpackLog(event, "HookSet", log); err != nil {
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

// SolverNetInboxIsmSetIterator is returned from FilterIsmSet and is used to iterate over the raw logs and unpacked data for IsmSet events raised by the SolverNetInbox contract.
type SolverNetInboxIsmSetIterator struct {
	Event *SolverNetInboxIsmSet // Event containing the contract specifics and raw log

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
func (it *SolverNetInboxIsmSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolverNetInboxIsmSet)
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
		it.Event = new(SolverNetInboxIsmSet)
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
func (it *SolverNetInboxIsmSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolverNetInboxIsmSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolverNetInboxIsmSet represents a IsmSet event raised by the SolverNetInbox contract.
type SolverNetInboxIsmSet struct {
	Ism common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIsmSet is a free log retrieval operation binding the contract event 0xc47cbcc588c67679e52261c45cc315e56562f8d0ccaba16facb9093ff9498799.
//
// Solidity: event IsmSet(address _ism)
func (_SolverNetInbox *SolverNetInboxFilterer) FilterIsmSet(opts *bind.FilterOpts) (*SolverNetInboxIsmSetIterator, error) {

	logs, sub, err := _SolverNetInbox.contract.FilterLogs(opts, "IsmSet")
	if err != nil {
		return nil, err
	}
	return &SolverNetInboxIsmSetIterator{contract: _SolverNetInbox.contract, event: "IsmSet", logs: logs, sub: sub}, nil
}

// WatchIsmSet is a free log subscription operation binding the contract event 0xc47cbcc588c67679e52261c45cc315e56562f8d0ccaba16facb9093ff9498799.
//
// Solidity: event IsmSet(address _ism)
func (_SolverNetInbox *SolverNetInboxFilterer) WatchIsmSet(opts *bind.WatchOpts, sink chan<- *SolverNetInboxIsmSet) (event.Subscription, error) {

	logs, sub, err := _SolverNetInbox.contract.WatchLogs(opts, "IsmSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolverNetInboxIsmSet)
				if err := _SolverNetInbox.contract.UnpackLog(event, "IsmSet", log); err != nil {
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
func (_SolverNetInbox *SolverNetInboxFilterer) ParseIsmSet(log types.Log) (*SolverNetInboxIsmSet, error) {
	event := new(SolverNetInboxIsmSet)
	if err := _SolverNetInbox.contract.UnpackLog(event, "IsmSet", log); err != nil {
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
