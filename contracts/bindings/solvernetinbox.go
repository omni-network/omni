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

// IERC7683GaslessCrossChainOrder is an auto generated low-level Go binding around an user-defined struct.
type IERC7683GaslessCrossChainOrder struct {
	OriginSettler common.Address
	User          common.Address
	Nonce         *big.Int
	OriginChainId *big.Int
	OpenDeadline  uint32
	FillDeadline  uint32
	OrderDataType [32]byte
	OrderData     []byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"weth_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"close\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGaslessCrossChainOrderDigest\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.GaslessCrossChainOrder\",\"components\":[{\"name\":\"originSettler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGaslessUserNonce\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestOrderOffset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint248\",\"internalType\":\"uint248\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextGaslessOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextOnchainOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOnchainUserNonce\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrder\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"resolved\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"state\",\"type\":\"tuple\",\"internalType\":\"structISolverNetInbox.OrderState\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumISolverNetInbox.Status\"},{\"name\":\"rejectReason\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"updatedBy\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"offset\",\"type\":\"uint248\",\"internalType\":\"uint248\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderId\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderSponsor\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPermit2Digest\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.GaslessCrossChainOrder\",\"components\":[{\"name\":\"originSettler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSponsorshipDigest\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.GaslessCrossChainOrder\",\"components\":[{\"name\":\"originSettler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSponsorshipDigest\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incrementGaslessNonce\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"markFilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"open\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"openFor\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.GaslessCrossChainOrder\",\"components\":[{\"name\":\"originSettler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originFillerData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseClose\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseOpen\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reject\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveFor\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.GaslessCrossChainOrder\",\"components\":[{\"name\":\"originSettler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"originFillerData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOutboxes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboxes\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validate\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.OnchainCrossChainOrder\",\"components\":[{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateFor\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIERC7683.GaslessCrossChainOrder\",\"components\":[{\"name\":\"originSettler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderDataType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"orderData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"originFillerData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Closed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillOriginData\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillOriginData\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSolverNet.FillOriginData\",\"components\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structSolverNet.Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"expenses\",\"type\":\"tuple[]\",\"internalType\":\"structSolverNet.TokenExpense[]\",\"components\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Filled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Open\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"resolvedOrder\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC7683.ResolvedCrossChainOrder\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fillDeadline\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxSpent\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minReceived\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.Output[]\",\"components\":[{\"name\":\"token\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fillInstructions\",\"type\":\"tuple[]\",\"internalType\":\"structIERC7683.FillInstruction[]\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destinationSettler\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboxSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Rejected\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCallTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestinationChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenseAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpenseToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFillDeadline\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMissingCalls\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNativeDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOpenDeadline\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrderData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrderTypehash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOriginChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOriginFillerData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOriginSettler\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReason\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSponsor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSponsorship\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderStillValid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PortalPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongFillHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongSourceChain\",\"inputs\":[]}]",
	Bin: "0x61016060405234801562000011575f80fd5b5060405162005dc738038062005dc78339810160408190526200003491620001f9565b306080524660a05260608062000083604080518082018252600e81526d0a6ded8eccae49ccae892dcc4def60931b602080830191909152825180840190935260018352603160f81b9083015291565b815160209283012081519183019190912060c082905260e0819052604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8152938401929092529082015246606082015230608082015260a0902061010052505063ffffffff60643b16156200016d5760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801562000153575060408051601f3d908101601f19168201909252620001509181019062000228565b60015b6200016357436101205262000173565b6101205262000173565b43610120525b6001600160a01b038116610140526200018b62000192565b5062000240565b63409feecd1980546001811615620001b15763f92ee8a95f526004601cfd5b6001600160401b03808260011c14620001f4578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b5f602082840312156200020a575f80fd5b81516001600160a01b038116811462000221575f80fd5b9392505050565b5f6020828403121562000239575f80fd5b5051919050565b60805160a05160c05160e051610100516101205161014051615b20620002a75f395f8181612555015281816125cf01528181613caa0152613ce801525f6107c001525f612da601525f612e6001525f612e3a01525f612dea01525f612dc70152615b205ff3fe608060405260043610610275575f3560e01c806371b233dd1161014a578063d9e8407c116100be578063eae4c19f11610078578063eae4c19f146107af578063f04e283e146107e2578063f2fde38b146107f5578063f904d28514610808578063fee81cf414610827578063ff44641014610858575f80fd5b8063d9e8407c14610701578063db3ea55314610720578063e1bd82f91461073f578063e255543d1461075e578063e7917ea51461077d578063e917a9621461079c575f80fd5b806384b0196e1161010f57806384b0196e146106305780638da5cb5b1461065757806396c144f01461066f5780639f31bc881461068e578063c0c53b8b146106c2578063d7118351146106e1575f80fd5b806371b233dd1461058357806374eeb847146105a2578063792aec5c146105d35780637cac41a6146105f2578063844fac8e14610611575f80fd5b80633563de29116101ec57806354d1f13d116101a657806354d1f13d146104e85780635778472a146104f0578063598d939e1461051e5780635e0aec951461053d5780636fa03e831461055c578063715018a61461057b575f80fd5b80633563de29146103f857806339acf9f11461042c57806339c79e0c1461046257806341b477dd146104815780634a4ee7b1146104a0578063514e62fc146104b3575f80fd5b80631c10893f1161023d5780631c10893f146103325780631cd64df41461034557806322bcd51a1461037457806325692962146103a05780632d622343146103a85780632de94807146103c7575f80fd5b806304a873ab1461027957806310dc77781461029a57806311ebce83146102cc578063183a4f6e1461030057806319b20ad114610313575b5f80fd5b348015610284575f80fd5b50610298610293366004614632565b61087f565b005b3480156102a5575f80fd5b506102b96102b43660046146af565b6109d3565b6040519081526020015b60405180910390f35b3480156102d7575f80fd5b506102b96102e6366004614704565b6001600160a01b03165f908152600c602052604090205490565b61029861030e36600461471f565b6109f7565b34801561031e575f80fd5b506102b961032d366004614736565b610a04565b61029861034036600461476c565b610a10565b348015610350575f80fd5b5061036461035f36600461476c565b610a26565b60405190151581526020016102c3565b34801561037f575f80fd5b5061039361038e3660046147d3565b610a44565b6040516102c391906149fb565b610298610a84565b3480156103b3575f80fd5b506102986103c2366004614a0d565b610ad0565b3480156103d2575f80fd5b506102b96103e1366004614704565b638b78c6d8600c9081525f91909152602090205490565b348015610403575f80fd5b506102b9610412366004614704565b6001600160a01b03165f908152600b602052604090205490565b348015610437575f80fd5b505f5461044a906001600160a01b031681565b6040516001600160a01b0390911681526020016102c3565b34801561046d575f80fd5b5061029861047c36600461471f565b610d6b565b34801561048c575f80fd5b5061039361049b366004614736565b6110e4565b6102986104ae36600461476c565b611106565b3480156104be575f80fd5b506103646104cd36600461476c565b638b78c6d8600c9081525f9290925260209091205416151590565b610298611118565b3480156104fb575f80fd5b5061050f61050a36600461471f565b611151565b6040516102c393929190614a57565b348015610529575f80fd5b506102b9610538366004614704565b611236565b348015610548575f80fd5b506102b961055736600461476c565b611259565b348015610567575f80fd5b506102b9610576366004614704565b611264565b610298611287565b34801561058e575f80fd5b506102b961059d3660046146af565b61129a565b3480156105ad575f80fd5b505f546105c190600160a01b900460ff1681565b60405160ff90911681526020016102c3565b3480156105de575f80fd5b506102986105ed366004614ae2565b611501565b3480156105fd575f80fd5b5061029861060c366004614ae2565b611523565b34801561061c575f80fd5b5061029861062b366004614afd565b611563565b34801561063b575f80fd5b50610644611739565b6040516102c39796959493929190614b8a565b348015610662575f80fd5b50638b78c6d8195461044a565b34801561067a575f80fd5b50610298610689366004614c21565b61179a565b348015610699575f80fd5b5061044a6106a836600461471f565b5f908152600960205260409020546001600160a01b031690565b3480156106cd575f80fd5b506102986106dc366004614c4f565b611909565b3480156106ec575f80fd5b506002546105c190600160f81b900460ff1681565b34801561070c575f80fd5b5061036461071b366004614736565b611998565b34801561072b575f80fd5b5061029861073a366004614c9c565b6119ac565b34801561074a575f80fd5b506102b96107593660046146af565b611b45565b348015610769575f80fd5b506103646107783660046147d3565b611b9b565b348015610788575f80fd5b50610298610797366004614cc6565b611bb3565b6102986107aa366004614736565b611c04565b3480156107ba575f80fd5b506102b97f000000000000000000000000000000000000000000000000000000000000000081565b6102986107f0366004614704565b611d50565b610298610803366004614704565b611d8a565b348015610813575f80fd5b50610298610822366004614ae2565b611db0565b348015610832575f80fd5b506102b9610841366004614704565b63389a75e1600c9081525f91909152602090205490565b348015610863575f80fd5b506002546040516001600160f81b0390911681526020016102c3565b610887611dd2565b8281146108a757604051634ec4810560e11b815260040160405180910390fd5b5f5b838110156109cc578282828181106108c3576108c3614ce7565b90506020020160208101906108d89190614704565b60035f8787858181106108ed576108ed614ce7565b90506020020160208101906109029190614d1a565b6001600160401b0316815260208101919091526040015f2080546001600160a01b0319166001600160a01b039290921691909117905582828281811061094a5761094a614ce7565b905060200201602081019061095f9190614704565b6001600160a01b031685858381811061097a5761097a614ce7565b905060200201602081019061098f9190614d1a565b6001600160401b03167ff730978310b4a2a0e6c673324d737afdb93d0afefed14a3d061b60f66e31f4e360405160405180910390a36001016108a9565b5050505050565b5f80806109e1843683611dec565b915091506109ef8282611eb5565b949350505050565b610a013382611f33565b50565b5f805f6109e184611f3e565b610a18611dd2565b610a228282611feb565b5050565b638b78c6d8600c9081525f8390526020902054811681145b92915050565b610a4c614455565b5f80610a59868686611dec565b9092509050610a788282610a7360a08a0160808b01614d51565b611ff7565b925050505b9392505050565b5f6202a3006001600160401b03164201905063389a75e1600c52335f52806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d5f80a250565b5f5460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa158015610b14573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b389190614e60565b8051600180546020938401516001600160401b039384166001600160e01b031990921691909117600160401b6001600160a01b0392831602179091555f868152600484526040808220815160608101835290549384168152600160a01b840490941684860152600160e01b90920463ffffffff168383015286815260089093528083208151608081019092528054929392829060ff166005811115610bdf57610bdf614a43565b6005811115610bf057610bf0614a43565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b03166060909101529050600181516005811115610c3f57610c3f614a43565b14610c5d57604051635d12a4a360e11b815260040160405180910390fd5b60208201516001546001600160401b03908116911614610c9057604051633687f39960e21b815260040160405180910390fd5b6001546001600160401b0381165f908152600360205260409020546001600160a01b03908116600160401b9092041614610cdc576040516282b42960e81b815260040160405180910390fd5b610ce585612079565b8414610d0457604051631f53eaed60e21b815260040160405180910390fd5b610d118560045f86612321565b610d1c856004612435565b826001600160a01b031684867fa7e64de5f8345186f3a39d8f0664d7d6b534e35ca818dbfb1465bb12f80562fc60405160405180910390a45050600180546001600160e01b0319169055505050565b6002545f80516020615acb83398151915290600160f81b900460ff168015610e2c5760ff81166001148015610dac57505f80516020615a8b83398151915282145b15610dca57604051631309a56360e01b815260040160405180910390fd5b60ff81166002148015610de957505f80516020615acb83398151915282145b15610e0757604051631309a56360e01b815260040160405180910390fd5b60021960ff821601610e2c5760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd212685403610e4a5763ab143c065f526004601cfd5b3068929eee149b4bd21268555f838152600860205260408082208151608081019092528054829060ff166005811115610e8557610e85614a43565b6005811115610e9657610e96614a43565b8152905460ff61010082041660208084019190915263ffffffff62010000830481166040808601919091526001600160a01b03600160301b90940484166060958601525f8a815260048452819020815195860182525493841685526001600160401b03600160a01b85041692850192909252600160e01b90920490911690820152909150600182516005811115610f2f57610f2f614a43565b14610f4d57604051635d12a4a360e11b815260040160405180910390fd5b80516001600160a01b03163314610f76576040516282b42960e81b815260040160405180910390fd5b5f5460208201516040516308c3569160e31b81527ffeccba1cfc4544bf9cd83b76f36ae5c464750b6c43f682e26744ee21ec31fc1e60048201526001600160401b0390911660248201526001600160a01b039091169063461ab48890604401602060405180830381865afa158015610ff0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110149190614e9f565b1561103257604051630c2e605760e11b815260040160405180910390fd5b42615460826040015163ffffffff1661104b9190614ece565b10611069576040516321bb6b2160e11b815260040160405180910390fd5b6110768560035f33612321565b80515f8681526009602052604090205461109c9187916001600160a01b031660016124cc565b6110a7856003612435565b60405185907f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79905f90a250503868929eee149b4bd2126855505050565b6110ec614455565b5f806110f784611f3e565b915091506109ef82825f611ff7565b61110e611dd2565b610a228282611f33565b63389a75e1600c52335f525f6020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c925f80a2565b611159614455565b604080516080810182525f8082526020820181905291810182905260608101829052908061118685612687565b905061119381865f611ff7565b5f868152600860209081526040808320600a9092529182902054825160808101909352815491926001600160f81b03909116918390829060ff1660058111156111de576111de614a43565b60058111156111ef576111ef614a43565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b031660609091015292989297509550909350505050565b6001600160a01b0381165f908152600c6020526040812054610a3e908390612924565b5f610a7d8383612924565b6001600160a01b0381165f908152600b6020526040812054610a3e908390612924565b61128f611dd2565b6112985f612966565b565b5f806112a7833683611dec565b5090505f6112bb6040850160208601614704565b60208084015180519101519192509030906001600160601b031665ffffffffffff5f6112ed60a08a0160808b01614d51565b60405163927da10560e01b81526001600160a01b03888116600483015287811660248301528616604482015263ffffffff9190911691505f906e22d473030f116ddee9f6b43ac78ba39063927da10590606401606060405180830381865afa15801561135b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061137f9190614ef6565b925050505f6e22d473030f116ddee9f6b43ac78ba36001600160a01b0316633644e5156040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113cf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113f39190614f38565b604080517f65626cad6cb96493bf6f5ebea28756c966f023ab9e8a83a7101849d5573b36786020808301919091526001600160a01b039a8b1682840152978a16606082015265ffffffffffff96871660808201529390951660a0808501919091528551808503909101815260c0840186528051908701207ff3841cd1ff0085026a6327b620b67997ce40f282c88a8e905a7a5626e310f3d060e08501526101008401529590961661012082015261014080820192909252825180820390920182526101608101835281519184019190912061190160f01b6101808301526101828201949094526101a280820194909452815180820390940184526101c2019052815191012095945050505050565b600161150c816129a3565b610a225f80516020615acb833981519152836129d4565b600161152e816129a3565b8161154657600280546001600160f81b031690555050565b600280546001600160f81b0316600360f81b17905560035b505050565b6002545f80516020615a8b83398151915290600160f81b900460ff1680156116245760ff811660011480156115a457505f80516020615a8b83398151915282145b156115c257604051631309a56360e01b815260040160405180910390fd5b60ff811660021480156115e157505f80516020615acb83398151915282145b156115ff57604051631309a56360e01b815260040160405180910390fd5b60021960ff8216016116245760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd2126854036116425763ab143c065f526004601cfd5b3068929eee149b4bd21268555f8061165b898787611dec565b909250905061167060408a0160208b01614704565b6001600160a01b0316336001600160a01b03161461169357611693898989612a8d565b60808201516001600160a01b0316156116b0576116b08282612af0565b84156116c6576116c68983602001518888612b44565b6116f082826116db60408d0160208e01614704565b6116eb60a08e0160808f01614d51565b612bca565b61172261170360408b0160208c01614704565b6001600160a01b03165f908152600c6020526040902080546001019055565b50503868929eee149b4bd212685550505050505050565b600f60f81b6060805f808083611788604080518082018252600e81526d0a6ded8eccae49ccae892dcc4def60931b602080830191909152825180840190935260018352603160f81b9083015291565b97989097965046955030945091925090565b3068929eee149b4bd2126854036117b85763ab143c065f526004601cfd5b3068929eee149b4bd21268555f828152600860205260408082208151608081019092528054829060ff1660058111156117f3576117f3614a43565b600581111561180457611804614a43565b81529054610100810460ff16602083015262010000810463ffffffff166040830152600160301b90046001600160a01b0316606090910152905060048151600581111561185357611853614a43565b146118715760405163789bae3560e01b815260040160405180910390fd5b60608101516001600160a01b0316331461189d576040516282b42960e81b815260040160405180910390fd5b6118aa8360055f33612321565b6118b683835f806124cc565b6118c1836005612435565b6040516001600160a01b03831690339085907f8428df912f4f2125b442b488df9c7260cb607246895bcd29f262ecca090b1538905f90a4503868929eee149b4bd21268555050565b63409feecd19805460038255801561193f5760018160011c14303b106119365763f92ee8a95f526004601cfd5b818160ff1b1b91505b5061194984612ca3565b611954836001611feb565b61195d82612cde565b8015611992576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b50505050565b5f6119a282611f3e565b5060019392505050565b60016119b781612d80565b3068929eee149b4bd2126854036119d55763ab143c065f526004601cfd5b3068929eee149b4bd21268555f838152600860205260408082208151608081019092528054829060ff166005811115611a1057611a10614a43565b6005811115611a2157611a21614a43565b81529054610100810460ff908116602084015262010000820463ffffffff166040840152600160301b9091046001600160a01b031660609092019190915290915083165f03611a83576040516337b89b9360e21b815260040160405180910390fd5b600181516005811115611a9857611a98614a43565b14611ab657604051635d12a4a360e11b815260040160405180910390fd5b611ac38460028533612321565b5f84815260046020908152604080832054600990925290912054611af79186916001600160a01b03918216911660016124cc565b611b02846002612435565b60405160ff841690339086907f21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f905f90a4503868929eee149b4bd2126855505050565b5f807fa8b61de3f3ac43b6bb317d5c60a46ff14b09bd3c7364af6dd3f53c3ba889170e83604051602001611b7a929190614fb8565b604051602081830303815290604052805190602001209050610a7d81612da4565b5f611ba7848484611dec565b50600195945050505050565b61271061ffff82161115611bda57604051633ab3447f60e11b815260040160405180910390fd5b335f908152600c60205260408120805461ffff84169290611bfc908490614ece565b909155505050565b6002545f80516020615a8b83398151915290600160f81b900460ff168015611cc55760ff81166001148015611c4557505f80516020615a8b83398151915282145b15611c6357604051631309a56360e01b815260040160405180910390fd5b60ff81166002148015611c8257505f80516020615acb83398151915282145b15611ca057604051631309a56360e01b815260040160405180910390fd5b60021960ff821601611cc55760405163aaae8ef760e01b815260040160405180910390fd5b3068929eee149b4bd212685403611ce35763ab143c065f526004601cfd5b3068929eee149b4bd21268555f80611cfa85611f3e565b608082015191935091506001600160a01b031615611d1c57611d1c8282612af0565b611d288282335f612bca565b335f908152600b602052604090208054600101905550503868929eee149b4bd2126855505050565b611d58611dd2565b63389a75e1600c52805f526020600c208054421115611d7e57636f5e88185f526004601cfd5b5f9055610a0181612966565b611d92611dd2565b8060601b611da757637448fbae5f526004601cfd5b610a0181612966565b6001611dbb816129a3565b610a225f80516020615a8b833981519152836129d4565b638b78c6d819543314611298576382b429005f526004601cfd5b611df46144ad565b5f611e00858585612eba565b5f5f80516020615aab8339815191528660c0013503611e4b57611e3e611e2960e0880188615076565b611e3960c08a0160a08b01614d51565b613118565b8051519093509050611e85565b611e70611e5b60e0880188615076565b611e6b60c08a0160a08b01614d51565b613159565b9250611e826040870160208801614704565b90505b6001600160a01b0381165f908152600c602052604081205490611ea88383612924565b9350505050935093915050565b602080830151604080517f3fa21d7497437af9446a527e21741950cef1b1e2dba772dc6587fe77368b960e8185015290810184905281516001600160a01b031660608201529101516001600160601b031660808201525f90819060a0016040516020818303038152906040528051906020012090506109ef81612da4565b610a2282825f6131fd565b611f466144ad565b5f611f4f6144ad565b611f5884613254565b5f80516020615aab833981519152846020013503611f9457611f8d611f806040860186615076565b611e396020880188614d51565b9050611fb4565b611fb1611fa46040860186615076565b611e6b6020880188614d51565b90505b8051516001600160a01b03165f908152600b6020526040812054825151909190611fde9083612924565b9296929550919350505050565b610a22828260016131fd565b611fff614455565b83515f61200b86613317565b90505f61201787613573565b90505f612023886136cb565b604080516101008101825286516001600160a01b0316815246602082015263ffffffff98891681830152950151909616606085015250608083019590955260a082015260c08101939093525060e0820152919050565b5f818152600460209081526040808320815160608101835290546001600160a01b0381168252600160a01b81046001600160401b031682850152600160e01b900463ffffffff1681830152848452600683528184208054835181860281018601909452808452919385939290849084015b828210156121e6575f848152602090819020604080516080810182526003860290920180546001600160a01b0381168452600160a01b900460e01b6001600160e01b03191693830193909352600183015490820152600282018054919291606084019190612157906150b8565b80601f0160208091040260200160405190810160405280929190818152602001828054612183906150b8565b80156121ce5780601f106121a5576101008083540402835291602001916121ce565b820191905f5260205f20905b8154815290600101906020018083116121b157829003601f168201915b505050505081525050815260200190600101906120ea565b5050505090505f60075f8681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015612280575f848152602090819020604080516060810182526002860290920180546001600160a01b03908116845260019182015490811684860152600160a01b90046001600160601b031691830191909152908352909201910161221d565b5050505090505f6040518060a00160405280466001600160401b0316815260200185602001516001600160401b03168152602001856040015163ffffffff16815260200184815260200183815250905085816040516020016122e29190615144565b60408051601f19818403018152908290526123009291602001615237565b60405160208183030381529060405280519060200120945050505050919050565b5f848152600860205260409081902054815160808101909252610100900460ff16908085600581111561235657612356614a43565b81526020015f8560ff161161236b578261236d565b845b60ff16815263ffffffff42166020808301919091526001600160a01b0385166040928301525f88815260089091522081518154829060ff191660018360058111156123ba576123ba614a43565b02179055506020820151815460408401516060909401516001600160a01b0316600160301b026601000000000000600160d01b031963ffffffff909516620100000265ffffffff00001960ff909416610100029390931665ffffffffff00199092169190911791909117929092169190911790555050505050565b600181600581111561244957612449614a43565b03612452575050565b600481600581111561246657612466614a43565b1461247a575f828152600560205260408120555b600581600581111561248e5761248e614a43565b14610a22575f828152600460209081526040808320839055600690915281206124b69161450c565b5f828152600760205260408120610a229161452a565b5f848152600560209081526040918290208251808401909352546001600160a01b038082168452600160a01b9091046001600160601b031691830191909152831615612617575f61253182602001516001600160601b0316601e60ff166127106137d7565b905082801561253f57505f81115b156126155781516001600160a01b03166125ff577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004015f604051808303818588803b1580156125ac575f80fd5b505af11580156125be573d5f803e3d5ffd5b506125fa9350506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016915086905083613865565b612615565b8151612615906001600160a01b03168583613865565b505b60208101516001600160601b0316156109cc5780516001600160a01b031661266057602081015161265b906001600160a01b038616906001600160601b03166138af565b6109cc565b602081015181516109cc916001600160a01b039091169086906001600160601b0316613865565b61268f6144ad565b5f8281526005602090815260408083208151808301835290546001600160a01b038082168352600160a01b918290046001600160601b031683860152835188875260048652848720610140820186525491821660e082019081529282046001600160401b0316610100820152600160e01b90910463ffffffff16610120820152908152808401829052868552600684528285208054845181870281018701865281815293969295948601949192909184015b8282101561283d575f848152602090819020604080516080810182526003860290920180546001600160a01b0381168452600160a01b900460e01b6001600160e01b031916938301939093526001830154908201526002820180549192916060840191906127ae906150b8565b80601f01602080910402602001604051908101604052809291908181526020018280546127da906150b8565b80156128255780601f106127fc57610100808354040283529160200191612825565b820191905f5260205f20905b81548152906001019060200180831161280857829003601f168201915b50505050508152505081526020019060010190612741565b50505050815260200160075f8681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156128d9575f848152602090819020604080516060810182526002860290920180546001600160a01b03908116845260019182015490811684860152600160a01b90046001600160601b0316918301919091529083529092019101612876565b505050908252505f85815260096020908152604091829020546001600160a01b03169083015201612909836138c8565b81526040805160208181019092525f81529101529392505050565b604080516001600160a01b03841660208201529081018290524660608201525f9060800160405160208183030381529060405280519060200120905092915050565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a355565b638b78c6d819543314610a0157638b78c6d8600c52335f52806020600c205416610a01576382b429005f526004601cfd5b600254600160f81b900460ff166002198101612a035760405163aaae8ef760e01b815260040160405180910390fd5b5f5f80516020615a8b8339815191528414612a1f576002612a22565b60015b905082612a38578060ff168260ff161415612a42565b8060ff168260ff16145b15612a6057604051631309a56360e01b815260040160405180910390fd5b82612a6b575f612a6d565b805b6002601f6101000a81548160ff021916908360ff16021790555050505050565b5f612a9784611b45565b9050612ab4612aac6040860160208701614704565b8285856138e7565b61199257612ad3612acb6040860160208701614704565b8285856139ba565b61199257604051638baa579f60e01b815260040160405180910390fd5b5f612afb8383611eb5565b9050612b108360800151828560c001516139ff565b61155e57612b278360800151828560c00151613adc565b61155e5760405163d46ba7a560e01b815260040160405180910390fd5b82516001600160a01b031615611992575f8080612b638486018661524f565b91945092509050612bc1612b7d6040890160208a01614704565b602088015130906001600160601b0316612b9d60a08c0160808d01614d51565b8a516001600160a01b03169392919063ffffffff90811690899089908990613b2616565b50505050505050565b612be284602001518386608001518760a00151613c52565b5f612bee858584613d9c565b905080608001517f71069e0637faca19b5cceb36f6ee664347f592a954e37edf597b6c0f37afd59f8260e001515f81518110612c2c57612c2c614ce7565b602002602001015160400151806020019051810190612c4b91906154c3565b604051612c589190615144565b60405180910390a280608001517fa576d0af275d0c6207ef43ceee8c498a5d7a26b8157a32d3fdf361e64371628c82604051612c9491906149fb565b60405180910390a25050505050565b6001600160a01b0316638b78c6d819819055805f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b038116612d2d5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b604482015260640160405180910390fd5b5f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52335f52806020600c205416610a01576382b429005f526004601cfd5b7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000030147f0000000000000000000000000000000000000000000000000000000000000000461416612e975750604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81527f000000000000000000000000000000000000000000000000000000000000000060208201527f00000000000000000000000000000000000000000000000000000000000000009181019190915246606082015230608082015260a090205b6719010000000000005f5280601a5281603a52604260182090505f603a52919050565b5f600c81612ece6040870160208801614704565b6001600160a01b031681526020808201929092526040015f205491503090612ef890860186614704565b6001600160a01b031614612f1f5760405163096c735760e01b815260040160405180910390fd5b5f612f306040860160208701614704565b6001600160a01b031603612f575760405163fd684c3b60e01b815260040160405180910390fd5b8084604001351080612f775750612f7061271082614ece565b8460400135115b15612f9557604051633ab3447f60e11b815260040160405180910390fd5b46846060013514612fb95760405163c5ac559960e01b815260040160405180910390fd5b42612fca60a0860160808701614d51565b63ffffffff1610806130065750612fe760c0850160a08601614d51565b63ffffffff16612ffd60a0860160808701614d51565b63ffffffff1610155b15613024576040516331bf2bcb60e21b815260040160405180910390fd5b4261303560c0860160a08701614d51565b63ffffffff16116130595760405163582e388960e01b815260040160405180910390fd5b5f80516020615aab8339815191528460c001351415801561309e57507f77d073e8e3323771e6f057c81e50184ab60b77a5984fe96a82c1467e4cbb64978460c0013514155b156130bc57604051636aea87f360e01b815260040160405180910390fd5b6130c960e0850185615076565b90505f036130ea5760405163a342e7d960e01b815260040160405180910390fd5b81158015906130fa575060608214155b15611992576040516330ce5bc160e11b815260040160405180910390fd5b6131206144ad565b5f61312d8486018661578d565b9050613150815f0151826020015185846040015185606001518660800151613ff4565b95945050505050565b6131616144ad565b5f61316e84860186615840565b60208101519091506001600160a01b031661319c57604051632f212e0d60e11b815260040160405180910390fd5b5f6131be825f015183606001518685608001518660a001518760c00151613ff4565b60208301516001600160a01b03166080808301919091528301519091506131e4906138c8565b60a082015260409091015160c082015290509392505050565b638b78c6d8600c52825f526020600c2080548381178361321e575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe265f80a3505050505050565b426132626020830183614d51565b63ffffffff16116132865760405163582e388960e01b815260040160405180910390fd5b5f80516020615aab8339815191528160200135141580156132cb57507f77d073e8e3323771e6f057c81e50184ab60b77a5984fe96a82c1467e4cbb6497816020013514155b156132e957604051636aea87f360e01b815260040160405180910390fd5b6132f66040820182615076565b90505f03610a015760405163a342e7d960e01b815260040160405180910390fd5b80516040820151606083810151909291905f805b835181101561338d575f84828151811061334757613347614ce7565b60200260200101516040015111156133855783818151811061336b5761336b614ce7565b602002602001015160400151826133829190614ece565b91505b60010161332b565b505f80821161339d5782516133aa565b82516133aa906001614ece565b6001600160401b038111156133c1576133c1614d6c565b60405190808252806020026020018201604052801561341157816020015b604080516080810182525f8082526020808301829052928201819052606082015282525f199092019101816133df5790505b5090505f5b83518110156134f257604051806080016040528061346086848151811061343f5761343f614ce7565b6020026020010151602001516001600160a01b03166001600160a01b031690565b815260200185838151811061347757613477614ce7565b6020908102919091018101516040908101516001600160601b03168352898201516001600160401b03165f9081526003835220549101906001600160a01b0316815260200187602001516001600160401b03168152508282815181106134df576134df614ce7565b6020908102919091010152600101613416565b50811561356957604080516080810182525f8082526020808301869052888101516001600160401b03168252600390528290205490918201906001600160a01b0316815260200186602001516001600160401b03168152508184518151811061355d5761355d614ce7565b60200260200101819052505b9695505050505050565b60208101516080820151606091906001600160a01b0316156135b2578260a00151816020018181516135a59190615933565b6001600160601b03169052505b5f8082602001516001600160601b0316116135cd575f6135d0565b60015b60ff166001600160401b038111156135ea576135ea614d6c565b60405190808252806020026020018201604052801561363a57816020015b604080516080810182525f8082526020808301829052928201819052606082015282525f199092019101816136085790505b5060208301519091506001600160601b031615610a7d576040805160808101909152825181906001600160a01b0316815260200183602001516001600160601b0316815260200161369e86608001516001600160a01b03166001600160a01b031690565b815260200146815250815f815181106136b9576136b9614ce7565b60200260200101819052509392505050565b80516040808301516060808501518351600180825281860190955291949390915f91816020015b60408051606080820183525f8083526020830152918101919091528152602001906001900390816136f25750506040805160608082018352602088810180516001600160401b039081168552815181165f90815260038452869020546001600160a01b031683860152855160a0810187524682168152915116818301528985015163ffffffff168186015291820188905260808201879052835194955091939284019261379f9201615144565b604051602081830303815290604052815250815f815181106137c3576137c3614ce7565b602090810291909101015295945050505050565b8282028183858304148515170261385e575f198385098181108201900382848609835f0384168285116138115763ae47f7025f526004601cfd5b93849004938382119092035f839003839004600101029203041760026003830281188084028203028084028203028084028203028084028203028084028203028084029091030202610a7d565b0492915050565b816014528060345263a9059cbb60601b5f5260205f604460105f875af18060015f5114166138a557803d853b1517106138a5576390b8ec185f526004601cfd5b505f603452505050565b5f385f3884865af1610a225763b12d13eb5f526004601cfd5b5f610a3e82602001516001600160601b0316601e60ff166127106137d7565b5f6001600160a01b038516156109ef57604051853b613977578260408114613917576041811461393e57506139b1565b60208581013560ff81901c601b0190915285356040526001600160ff1b031660605261394f565b60408501355f1a6020526040856040375b50845f526020600160805f60015afa5180871860601b3d119250505f606052806040526139b1565b631626ba7e60e01b80825285600483015260248201604081528460448401528486606485013760208160648701858b5afa90519091141691505b50949350505050565b5f604051631626ba7e60e01b80825285600483015260248201604081528460448401528486606485013760208160648701858b5afa9051909114169695505050505050565b5f6001600160a01b03841615610a7d57604051843b613a9457825160408114613a305760418114613a515750613ad4565b604084015160ff81901c601b016020526001600160ff1b0316606052613a64565b60608401515f1a60205260408401516060525b50835f5260208301516040526020600160805f60015afa5180861860601b3d119250505f60605280604052613ad4565b631626ba7e60e01b808252846004830152602482016040815284516020018060448501828860045afa905060208260443d01868b5afa9151911691141691505b509392505050565b5f604051631626ba7e60e01b808252846004830152602482016040815284516020018060448501828860045afa905060208260443d01868b5afa9151911691141695945050505050565b5f73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2891860601b15613c3257633644e5155f5260205f6004601c8c611388fa60203d145f5115101615613c32576040518760348201528860601b60208201528560748201527fdbb8cf42e1ecb028be3f3dbc922e1d878b963f411dc388ced501601c60f7c6f75f5103613bfd5788601452623f675f60691b5f52602060548201602460108d5afa87151060948201526323f2ebc360621b81528460ff1660b48201528360d48201528260f48201525f38610104601084015f8e5af1915050613c32565b63d505accf60601b81528660548201528460ff1660948201528360b48201528260d48201525f3860e4601084015f8e5af19150505b80613c4757613c478989898989898989614242565b505050505050505050565b83516001600160a01b0316613d4d5783602001516001600160601b03163414613c8e5760405163036f810f60e41b815260040160405180910390fd5b6001600160a01b03821615613d4857613cd26001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016833084614329565b604051632e1a7d4d60e01b8152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d906024015f604051808303815f87803b158015613d31575f80fd5b505af1158015613d43573d5f803e3d5ffd5b505050505b611992565b60208401518451613d76916001600160a01b0390911690859030906001600160601b0316614329565b6001600160a01b03821615611992578351611992906001600160a01b0316833084614329565b613da4614455565b613daf848484611ff7565b84515f8581526004602090815260408083208451815484870151968401516001600160a01b039283166001600160e01b031990921691909117600160a01b6001600160401b039098168802176001600160e01b0316600160e01b63ffffffff9092169190910217909155828a0151600584528285208151918501519183166001600160601b03909216909602179094556080890151600990925290912080546001600160a01b031916919092161790559050613e69614345565b5f848152600a6020526040812080546001600160f81b0319166001600160f81b0393909316929092179091555b846040015151811015613f4b575f848152600660205260409081902090860151805183908110613ec857613ec8614ce7565b6020908102919091018101518254600181810185555f94855293839020825160039092020180549383015160e01c600160a01b026001600160c01b03199094166001600160a01b039092169190911792909217825560408101519282019290925560608201516002820190613f3d908261599e565b505050806001019050613e96565b505f5b846060015151811015613fe6575f8481526007602052604090206060860151805183908110613f7f57613f7f614ce7565b6020908102919091018101518254600181810185555f94855293839020825160029092020180546001600160a01b039283166001600160a01b0319909116178155928201516040909201516001600160601b0316600160a01b029116179082015501613f4e565b50610a7d8360015f33612321565b613ffc6144ad565b6001600160a01b03871661400e573396505b6001600160401b038616158061402c575046866001600160401b0316145b1561404a5760405163090eaaa760e41b815260040160405180910390fd5b604080516060810182526001600160a01b03891681526001600160401b038816602082015263ffffffff87169181019190915283515f0361409e57604051639cc71f7d60e01b815260040160405180910390fd5b8351602010156140c157604051634ec4810560e11b815260040160405180910390fd5b5f5b845181101561411e575f8582815181106140df576140df614ce7565b602090810291909101015180519091506001600160a01b03166141155760405163017ab86160e21b815260040160405180910390fd5b506001016140c3565b5082516020101561414257604051634ec4810560e11b815260040160405180910390fd5b5f5b83518110156141ea575f6001600160a01b031684828151811061416957614169614ce7565b6020026020010151602001516001600160a01b03160361419c5760405163027dcfa160e31b815260040160405180910390fd5b8381815181106141ae576141ae614ce7565b6020026020010151604001516001600160601b03165f036141e25760405163a0ce339960e01b815260040160405180910390fd5b600101614144565b506040518060e001604052808281526020018681526020018581526020018481526020015f6001600160a01b031681526020015f815260200160405180602001604052805f8152508152509150509695505050505050565b60405163927da10581525f1960601c88811660208301528981166040830152878116606083015287811660c0830152508560a01c156e22d473030f116ddee9f6b43ac78ba30260608083016064601c8501845afa605f3d11166142b657676b836e6b8757f0fd5f526004811560021b601801fd5b632b67b570825286606083015265ffffffffffff60808301528560e083015261010080830152604161012083015283610140830152826101608301528460f81b6101808301525f38610184601c85015f855af18a3b0261431d57636b836e6b5f526004601cfd5b50505050505050505050565b61433584848484614388565b61199257611992848484846143d7565b600280545f91908290614360906001600160f81b0316615a5d565b91906101000a8154816001600160f81b0302191690836001600160f81b031602179055905090565b5f60405182606052836040528460601b602c526323b872dd60601b600c5260205f6064601c5f8a5af191508160015f5114166143c857813d873b15171091505b5f606052604052949350505050565b6040518460601b60601c60748201528160548201528260348201528360601b6020820152631b63c28b60611b81526e22d473030f116ddee9f6b43ac78ba360014614806144245750803b15155b80873b15105f386084601087015f875af116612bc157677939f4248757f0fd5f5260048460a01c151560021b601801fd5b6040518061010001604052805f6001600160a01b031681526020015f81526020015f63ffffffff1681526020015f63ffffffff1681526020015f80191681526020016060815260200160608152602001606081525090565b60408051610140810182525f60e0820181815261010083018290526101208301829052825282518084018452818152602080820183905283015260609282018390528282018390526080820181905260a082015260c081019190915290565b5080545f8255600302905f5260205f2090810190610a019190614548565b5080545f8255600202905f5260205f2090810190610a019190614580565b8082111561457c5780546001600160c01b03191681555f6001820181905561457360028301826145a5565b50600301614548565b5090565b5b8082111561457c5780546001600160a01b03191681555f6001820155600201614581565b5080546145b1906150b8565b5f825580601f106145c0575050565b601f0160209004905f5260205f2090810190610a0191905b8082111561457c575f81556001016145d8565b5f8083601f8401126145fb575f80fd5b5081356001600160401b03811115614611575f80fd5b6020830191508360208260051b850101111561462b575f80fd5b9250929050565b5f805f8060408587031215614645575f80fd5b84356001600160401b038082111561465b575f80fd5b614667888389016145eb565b9096509450602087013591508082111561467f575f80fd5b5061468c878288016145eb565b95989497509550505050565b5f61010082840312156146a9575f80fd5b50919050565b5f602082840312156146bf575f80fd5b81356001600160401b038111156146d4575f80fd5b6109ef84828501614698565b6001600160a01b0381168114610a01575f80fd5b80356146ff816146e0565b919050565b5f60208284031215614714575f80fd5b8135610a7d816146e0565b5f6020828403121561472f575f80fd5b5035919050565b5f60208284031215614746575f80fd5b81356001600160401b0381111561475b575f80fd5b820160608185031215610a7d575f80fd5b5f806040838503121561477d575f80fd5b8235614788816146e0565b946020939093013593505050565b5f8083601f8401126147a6575f80fd5b5081356001600160401b038111156147bc575f80fd5b60208301915083602082850101111561462b575f80fd5b5f805f604084860312156147e5575f80fd5b83356001600160401b03808211156147fb575f80fd5b61480787838801614698565b9450602086013591508082111561481c575f80fd5b5061482986828701614796565b9497909650939450505050565b5f815180845260208085019450602084015f5b83811015614885578151805188528381015184890152604080820151908901526060908101519088015260809096019590820190600101614849565b509495945050505050565b5f5b838110156148aa578181015183820152602001614892565b50505f910152565b5f81518084526148c9816020860160208601614890565b601f01601f19169290920160200192915050565b5f82825180855260208086019550808260051b8401018186015f5b8481101561494f57858303601f19018952815180516001600160401b03168452848101518585015260409081015160609185018290529061493b818601836148b2565b9a86019a94505050908301906001016148f8565b5090979650505050505050565b5f61010060018060a01b03835116845260208301516020850152604083015161498d604086018263ffffffff169052565b5060608301516149a5606086018263ffffffff169052565b506080830151608085015260a08301518160a08601526149c782860182614836565b91505060c083015184820360c08601526149e18282614836565b91505060e083015184820360e086015261315082826148dd565b602081525f610a7d602083018461495c565b5f805f60608486031215614a1f575f80fd5b83359250602084013591506040840135614a38816146e0565b809150509250925092565b634e487b7160e01b5f52602160045260245ffd5b60c081525f614a6960c083018661495c565b9050835160068110614a8957634e487b7160e01b5f52602160045260245ffd5b8060208401525060ff602085015116604083015263ffffffff604085015116606083015260018060a01b03606085015116608083015260018060f81b03831660a0830152949350505050565b8015158114610a01575f80fd5b5f60208284031215614af2575f80fd5b8135610a7d81614ad5565b5f805f805f60608688031215614b11575f80fd5b85356001600160401b0380821115614b27575f80fd5b614b3389838a01614698565b96506020880135915080821115614b48575f80fd5b614b5489838a01614796565b90965094506040880135915080821115614b6c575f80fd5b50614b7988828901614796565b969995985093965092949392505050565b60ff60f81b881681525f602060e06020840152614baa60e084018a6148b2565b8381036040850152614bbc818a6148b2565b606085018990526001600160a01b038816608086015260a0850187905284810360c0860152855180825260208088019350909101905f5b81811015614c0f57835183529284019291840191600101614bf3565b50909c9b505050505050505050505050565b5f8060408385031215614c32575f80fd5b823591506020830135614c44816146e0565b809150509250929050565b5f805f60608486031215614c61575f80fd5b8335614c6c816146e0565b92506020840135614c7c816146e0565b91506040840135614a38816146e0565b803560ff811681146146ff575f80fd5b5f8060408385031215614cad575f80fd5b82359150614cbd60208401614c8c565b90509250929050565b5f60208284031215614cd6575f80fd5b813561ffff81168114610a7d575f80fd5b634e487b7160e01b5f52603260045260245ffd5b6001600160401b0381168114610a01575f80fd5b80356146ff81614cfb565b5f60208284031215614d2a575f80fd5b8135610a7d81614cfb565b63ffffffff81168114610a01575f80fd5b80356146ff81614d35565b5f60208284031215614d61575f80fd5b8135610a7d81614d35565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715614da257614da2614d6c565b60405290565b604051608081016001600160401b0381118282101715614da257614da2614d6c565b604051606081016001600160401b0381118282101715614da257614da2614d6c565b60405160a081016001600160401b0381118282101715614da257614da2614d6c565b60405160e081016001600160401b0381118282101715614da257614da2614d6c565b604051601f8201601f191681016001600160401b0381118282101715614e5857614e58614d6c565b604052919050565b5f60408284031215614e70575f80fd5b614e78614d80565b8251614e8381614cfb565b81526020830151614e93816146e0565b60208201529392505050565b5f60208284031215614eaf575f80fd5b8151610a7d81614ad5565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610a3e57610a3e614eba565b805165ffffffffffff811681146146ff575f80fd5b5f805f60608486031215614f08575f80fd5b8351614f13816146e0565b9250614f2160208501614ee1565b9150614f2f60408501614ee1565b90509250925092565b5f60208284031215614f48575f80fd5b5051919050565b5f808335601e19843603018112614f64575f80fd5b83016020810192503590506001600160401b03811115614f82575f80fd5b80360382131561462b575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b828152604060208201525f8235614fce816146e0565b6001600160a01b03166040830152614fe8602084016146f4565b6001600160a01b03811660608401525060408301356080830152606083013560a083015261501860808401614d46565b63ffffffff1660c083015261502f60a08401614d46565b63ffffffff811660e08401525061010060c08401358184015261505560e0850185614f4f565b8261012086015261506b61014086018284614f90565b979650505050505050565b5f808335601e1984360301811261508b575f80fd5b8301803591506001600160401b038211156150a4575f80fd5b60200191503681900382131561462b575f80fd5b600181811c908216806150cc57607f821691505b6020821081036146a957634e487b7160e01b5f52602260045260245ffd5b5f815180845260208085019450602084015f5b8381101561488557815180516001600160a01b0390811689528482015116848901526040908101516001600160601b031690880152606090960195908201906001016150fd565b5f602080835260c083016001600160401b0380865116838601528286015160408282166040880152604088015192506060915063ffffffff8316606088015260608801519250608060a0608089015284845180875260e08a01915060e08160051b8b0101965087860195505f5b81811015615213578a880360df19018352865180516001600160a01b03168952898101516001600160e01b0319168a8a015285810151868a0152860151868901859052615200858a01826148b2565b98505095880195918801916001016151b1565b5050505050505060808501519150601f198482030160a085015261315081836150ea565b828152604060208201525f6109ef60408301846148b2565b5f805f60608486031215615261575f80fd5b61526a84614c8c565b95602085013595506040909401359392505050565b80516146ff81614d35565b5f6001600160401b038211156152a2576152a2614d6c565b5060051b60200190565b6001600160e01b031981168114610a01575f80fd5b5f6001600160401b038211156152d9576152d9614d6c565b50601f01601f191660200190565b5f82601f8301126152f6575f80fd5b8151602061530b6153068361528a565b614e30565b82815260059290921b84018101918181019086841115615329575f80fd5b8286015b848110156154075780516001600160401b038082111561534b575f80fd5b908801906080828b03601f1901811315615363575f80fd5b61536b614da8565b87840151615378816146e0565b8152604084810151615389816152ac565b828a0152606085810151828401529285015192848411156153a8575f80fd5b83860195508d603f8701126153bb575f80fd5b8986015194506153cd615306866152c1565b93508484528d828688010111156153e2575f80fd5b6153f1858b8601848901614890565b820192909252865250505091830191830161532d565b509695505050505050565b6001600160601b0381168114610a01575f80fd5b5f82601f830112615435575f80fd5b815160206154456153068361528a565b82815260609283028501820192828201919087851115615463575f80fd5b8387015b8581101561494f5781818a03121561547d575f80fd5b615485614dca565b8151615490816146e0565b81528186015161549f816146e0565b818701526040828101516154b281615412565b908201528452928401928101615467565b5f602082840312156154d3575f80fd5b81516001600160401b03808211156154e9575f80fd5b9083019060a082860312156154fc575f80fd5b615504614dec565b825161550f81614cfb565b8152602083015161551f81614cfb565b60208201526155306040840161527f565b6040820152606083015182811115615546575f80fd5b615552878286016152e7565b606083015250608083015182811115615569575f80fd5b61557587828601615426565b60808301525095945050505050565b5f60408284031215615594575f80fd5b61559c614d80565b905081356155a9816146e0565b815260208201356155b981615412565b602082015292915050565b5f82601f8301126155d3575f80fd5b81356155e1615306826152c1565b8181528460208386010111156155f5575f80fd5b816020850160208301375f918101602001919091529392505050565b5f82601f830112615620575f80fd5b813560206156306153068361528a565b82815260059290921b8401810191818101908684111561564e575f80fd5b8286015b848110156154075780356001600160401b0380821115615670575f80fd5b908801906080828b03601f1901811315615688575f80fd5b615690614da8565b8784013561569d816146e0565b81526040848101356156ae816152ac565b828a0152606085810135828401529285013592848411156156cd575f80fd5b6156db8e8b868901016155c4565b90830152508652505050918301918301615652565b5f82601f8301126156ff575f80fd5b8135602061570f6153068361528a565b8281526060928302850182019282820191908785111561572d575f80fd5b8387015b8581101561494f5781818a031215615747575f80fd5b61574f614dca565b813561575a816146e0565b815281860135615769816146e0565b8187015260408281013561577c81615412565b908201528452928401928101615731565b5f6020828403121561579d575f80fd5b81356001600160401b03808211156157b3575f80fd5b9083019060c082860312156157c6575f80fd5b6157ce614dec565b82356157d9816146e0565b815260208301356157e981614cfb565b60208201526157fb8660408501615584565b6040820152608083013582811115615811575f80fd5b61581d87828601615611565b60608301525060a083013582811115615834575f80fd5b615575878286016156f0565b5f60208284031215615850575f80fd5b81356001600160401b0380821115615866575f80fd5b90830190610100828603121561587a575f80fd5b615882614e0e565b61588b836146f4565b8152615899602084016146f4565b60208201526040830135828111156158af575f80fd5b6158bb878286016155c4565b6040830152506158cd60608401614d0f565b60608201526158df8660808501615584565b608082015260c0830135828111156158f5575f80fd5b61590187828601615611565b60a08301525060e083013582811115615918575f80fd5b615924878286016156f0565b60c08301525095945050505050565b6001600160601b0381811683821601908082111561595357615953614eba565b5092915050565b601f82111561155e57805f5260205f20601f840160051c8101602085101561597f5750805b601f840160051c820191505b818110156109cc575f815560010161598b565b81516001600160401b038111156159b7576159b7614d6c565b6159cb816159c584546150b8565b8461595a565b602080601f8311600181146159fe575f84156159e75750858301515b5f19600386901b1c1916600185901b178555615a55565b5f85815260208120601f198616915b82811015615a2c57888601518255948401946001909101908401615a0d565b5085821015615a4957878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f6001600160f81b038281166002600160f81b03198101615a8057615a80614eba565b600101939250505056fef76fe33b8a0ebf7aa05740f479d10138c7c15bdc75b10e047cc15be2be15e5b42e7de755ca70cb933dc80103af16cc3303580e5712f1a8927d6461441e99a1e65ffb10051d79c19b9690b0842a292cb621fbf85d15269ed21c4e6a431d892bb5a26469706673582212207e662d764b93793ce86d3294175d1b8758eb5f8d05d8b9faa69d33382123fc4c64736f6c63430008180033",
}

// SolverNetInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SolverNetInboxMetaData.ABI instead.
var SolverNetInboxABI = SolverNetInboxMetaData.ABI

// SolverNetInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolverNetInboxMetaData.Bin instead.
var SolverNetInboxBin = SolverNetInboxMetaData.Bin

// DeploySolverNetInbox deploys a new Ethereum contract, binding an instance of SolverNetInbox to it.
func DeploySolverNetInbox(auth *bind.TransactOpts, backend bind.ContractBackend, weth_ common.Address) (common.Address, *types.Transaction, *SolverNetInbox, error) {
	parsed, err := SolverNetInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolverNetInboxBin), backend, weth_)
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

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_SolverNetInbox *SolverNetInboxCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "eip712Domain")

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
func (_SolverNetInbox *SolverNetInboxSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _SolverNetInbox.Contract.Eip712Domain(&_SolverNetInbox.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_SolverNetInbox *SolverNetInboxCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _SolverNetInbox.Contract.Eip712Domain(&_SolverNetInbox.CallOpts)
}

// GetGaslessCrossChainOrderDigest is a free data retrieval call binding the contract method 0xe1bd82f9.
//
// Solidity: function getGaslessCrossChainOrderDigest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetGaslessCrossChainOrderDigest(opts *bind.CallOpts, order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getGaslessCrossChainOrderDigest", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGaslessCrossChainOrderDigest is a free data retrieval call binding the contract method 0xe1bd82f9.
//
// Solidity: function getGaslessCrossChainOrderDigest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetGaslessCrossChainOrderDigest(order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetGaslessCrossChainOrderDigest(&_SolverNetInbox.CallOpts, order)
}

// GetGaslessCrossChainOrderDigest is a free data retrieval call binding the contract method 0xe1bd82f9.
//
// Solidity: function getGaslessCrossChainOrderDigest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetGaslessCrossChainOrderDigest(order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetGaslessCrossChainOrderDigest(&_SolverNetInbox.CallOpts, order)
}

// GetGaslessUserNonce is a free data retrieval call binding the contract method 0x11ebce83.
//
// Solidity: function getGaslessUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxCaller) GetGaslessUserNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getGaslessUserNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaslessUserNonce is a free data retrieval call binding the contract method 0x11ebce83.
//
// Solidity: function getGaslessUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxSession) GetGaslessUserNonce(user common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.GetGaslessUserNonce(&_SolverNetInbox.CallOpts, user)
}

// GetGaslessUserNonce is a free data retrieval call binding the contract method 0x11ebce83.
//
// Solidity: function getGaslessUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetGaslessUserNonce(user common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.GetGaslessUserNonce(&_SolverNetInbox.CallOpts, user)
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

// GetNextGaslessOrderId is a free data retrieval call binding the contract method 0x598d939e.
//
// Solidity: function getNextGaslessOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetNextGaslessOrderId(opts *bind.CallOpts, user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getNextGaslessOrderId", user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextGaslessOrderId is a free data retrieval call binding the contract method 0x598d939e.
//
// Solidity: function getNextGaslessOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetNextGaslessOrderId(user common.Address) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetNextGaslessOrderId(&_SolverNetInbox.CallOpts, user)
}

// GetNextGaslessOrderId is a free data retrieval call binding the contract method 0x598d939e.
//
// Solidity: function getNextGaslessOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetNextGaslessOrderId(user common.Address) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetNextGaslessOrderId(&_SolverNetInbox.CallOpts, user)
}

// GetNextOnchainOrderId is a free data retrieval call binding the contract method 0x6fa03e83.
//
// Solidity: function getNextOnchainOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetNextOnchainOrderId(opts *bind.CallOpts, user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getNextOnchainOrderId", user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextOnchainOrderId is a free data retrieval call binding the contract method 0x6fa03e83.
//
// Solidity: function getNextOnchainOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetNextOnchainOrderId(user common.Address) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetNextOnchainOrderId(&_SolverNetInbox.CallOpts, user)
}

// GetNextOnchainOrderId is a free data retrieval call binding the contract method 0x6fa03e83.
//
// Solidity: function getNextOnchainOrderId(address user) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetNextOnchainOrderId(user common.Address) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetNextOnchainOrderId(&_SolverNetInbox.CallOpts, user)
}

// GetOnchainUserNonce is a free data retrieval call binding the contract method 0x3563de29.
//
// Solidity: function getOnchainUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxCaller) GetOnchainUserNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getOnchainUserNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOnchainUserNonce is a free data retrieval call binding the contract method 0x3563de29.
//
// Solidity: function getOnchainUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxSession) GetOnchainUserNonce(user common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.GetOnchainUserNonce(&_SolverNetInbox.CallOpts, user)
}

// GetOnchainUserNonce is a free data retrieval call binding the contract method 0x3563de29.
//
// Solidity: function getOnchainUserNonce(address user) view returns(uint256)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetOnchainUserNonce(user common.Address) (*big.Int, error) {
	return _SolverNetInbox.Contract.GetOnchainUserNonce(&_SolverNetInbox.CallOpts, user)
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

// GetOrderSponsor is a free data retrieval call binding the contract method 0x9f31bc88.
//
// Solidity: function getOrderSponsor(bytes32 id) view returns(address)
func (_SolverNetInbox *SolverNetInboxCaller) GetOrderSponsor(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getOrderSponsor", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOrderSponsor is a free data retrieval call binding the contract method 0x9f31bc88.
//
// Solidity: function getOrderSponsor(bytes32 id) view returns(address)
func (_SolverNetInbox *SolverNetInboxSession) GetOrderSponsor(id [32]byte) (common.Address, error) {
	return _SolverNetInbox.Contract.GetOrderSponsor(&_SolverNetInbox.CallOpts, id)
}

// GetOrderSponsor is a free data retrieval call binding the contract method 0x9f31bc88.
//
// Solidity: function getOrderSponsor(bytes32 id) view returns(address)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetOrderSponsor(id [32]byte) (common.Address, error) {
	return _SolverNetInbox.Contract.GetOrderSponsor(&_SolverNetInbox.CallOpts, id)
}

// GetPermit2Digest is a free data retrieval call binding the contract method 0x71b233dd.
//
// Solidity: function getPermit2Digest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetPermit2Digest(opts *bind.CallOpts, order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getPermit2Digest", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPermit2Digest is a free data retrieval call binding the contract method 0x71b233dd.
//
// Solidity: function getPermit2Digest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetPermit2Digest(order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetPermit2Digest(&_SolverNetInbox.CallOpts, order)
}

// GetPermit2Digest is a free data retrieval call binding the contract method 0x71b233dd.
//
// Solidity: function getPermit2Digest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetPermit2Digest(order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetPermit2Digest(&_SolverNetInbox.CallOpts, order)
}

// GetSponsorshipDigest is a free data retrieval call binding the contract method 0x10dc7778.
//
// Solidity: function getSponsorshipDigest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetSponsorshipDigest(opts *bind.CallOpts, order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getSponsorshipDigest", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSponsorshipDigest is a free data retrieval call binding the contract method 0x10dc7778.
//
// Solidity: function getSponsorshipDigest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetSponsorshipDigest(order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetSponsorshipDigest(&_SolverNetInbox.CallOpts, order)
}

// GetSponsorshipDigest is a free data retrieval call binding the contract method 0x10dc7778.
//
// Solidity: function getSponsorshipDigest((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetSponsorshipDigest(order IERC7683GaslessCrossChainOrder) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetSponsorshipDigest(&_SolverNetInbox.CallOpts, order)
}

// GetSponsorshipDigest0 is a free data retrieval call binding the contract method 0x19b20ad1.
//
// Solidity: function getSponsorshipDigest((uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCaller) GetSponsorshipDigest0(opts *bind.CallOpts, order IERC7683OnchainCrossChainOrder) ([32]byte, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "getSponsorshipDigest0", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSponsorshipDigest0 is a free data retrieval call binding the contract method 0x19b20ad1.
//
// Solidity: function getSponsorshipDigest((uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxSession) GetSponsorshipDigest0(order IERC7683OnchainCrossChainOrder) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetSponsorshipDigest0(&_SolverNetInbox.CallOpts, order)
}

// GetSponsorshipDigest0 is a free data retrieval call binding the contract method 0x19b20ad1.
//
// Solidity: function getSponsorshipDigest((uint32,bytes32,bytes) order) view returns(bytes32)
func (_SolverNetInbox *SolverNetInboxCallerSession) GetSponsorshipDigest0(order IERC7683OnchainCrossChainOrder) ([32]byte, error) {
	return _SolverNetInbox.Contract.GetSponsorshipDigest0(&_SolverNetInbox.CallOpts, order)
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

// ResolveFor is a free data retrieval call binding the contract method 0x22bcd51a.
//
// Solidity: function resolveFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes originFillerData) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]))
func (_SolverNetInbox *SolverNetInboxCaller) ResolveFor(opts *bind.CallOpts, order IERC7683GaslessCrossChainOrder, originFillerData []byte) (IERC7683ResolvedCrossChainOrder, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "resolveFor", order, originFillerData)

	if err != nil {
		return *new(IERC7683ResolvedCrossChainOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC7683ResolvedCrossChainOrder)).(*IERC7683ResolvedCrossChainOrder)

	return out0, err

}

// ResolveFor is a free data retrieval call binding the contract method 0x22bcd51a.
//
// Solidity: function resolveFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes originFillerData) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]))
func (_SolverNetInbox *SolverNetInboxSession) ResolveFor(order IERC7683GaslessCrossChainOrder, originFillerData []byte) (IERC7683ResolvedCrossChainOrder, error) {
	return _SolverNetInbox.Contract.ResolveFor(&_SolverNetInbox.CallOpts, order, originFillerData)
}

// ResolveFor is a free data retrieval call binding the contract method 0x22bcd51a.
//
// Solidity: function resolveFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes originFillerData) view returns((address,uint256,uint32,uint32,bytes32,(bytes32,uint256,bytes32,uint256)[],(bytes32,uint256,bytes32,uint256)[],(uint64,bytes32,bytes)[]))
func (_SolverNetInbox *SolverNetInboxCallerSession) ResolveFor(order IERC7683GaslessCrossChainOrder, originFillerData []byte) (IERC7683ResolvedCrossChainOrder, error) {
	return _SolverNetInbox.Contract.ResolveFor(&_SolverNetInbox.CallOpts, order, originFillerData)
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

// ValidateFor is a free data retrieval call binding the contract method 0xe255543d.
//
// Solidity: function validateFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes originFillerData) view returns(bool)
func (_SolverNetInbox *SolverNetInboxCaller) ValidateFor(opts *bind.CallOpts, order IERC7683GaslessCrossChainOrder, originFillerData []byte) (bool, error) {
	var out []interface{}
	err := _SolverNetInbox.contract.Call(opts, &out, "validateFor", order, originFillerData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateFor is a free data retrieval call binding the contract method 0xe255543d.
//
// Solidity: function validateFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes originFillerData) view returns(bool)
func (_SolverNetInbox *SolverNetInboxSession) ValidateFor(order IERC7683GaslessCrossChainOrder, originFillerData []byte) (bool, error) {
	return _SolverNetInbox.Contract.ValidateFor(&_SolverNetInbox.CallOpts, order, originFillerData)
}

// ValidateFor is a free data retrieval call binding the contract method 0xe255543d.
//
// Solidity: function validateFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes originFillerData) view returns(bool)
func (_SolverNetInbox *SolverNetInboxCallerSession) ValidateFor(order IERC7683GaslessCrossChainOrder, originFillerData []byte) (bool, error) {
	return _SolverNetInbox.Contract.ValidateFor(&_SolverNetInbox.CallOpts, order, originFillerData)
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

// IncrementGaslessNonce is a paid mutator transaction binding the contract method 0xe7917ea5.
//
// Solidity: function incrementGaslessNonce(uint16 amount) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) IncrementGaslessNonce(opts *bind.TransactOpts, amount uint16) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "incrementGaslessNonce", amount)
}

// IncrementGaslessNonce is a paid mutator transaction binding the contract method 0xe7917ea5.
//
// Solidity: function incrementGaslessNonce(uint16 amount) returns()
func (_SolverNetInbox *SolverNetInboxSession) IncrementGaslessNonce(amount uint16) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.IncrementGaslessNonce(&_SolverNetInbox.TransactOpts, amount)
}

// IncrementGaslessNonce is a paid mutator transaction binding the contract method 0xe7917ea5.
//
// Solidity: function incrementGaslessNonce(uint16 amount) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) IncrementGaslessNonce(amount uint16) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.IncrementGaslessNonce(&_SolverNetInbox.TransactOpts, amount)
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

// OpenFor is a paid mutator transaction binding the contract method 0x844fac8e.
//
// Solidity: function openFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes signature, bytes originFillerData) returns()
func (_SolverNetInbox *SolverNetInboxTransactor) OpenFor(opts *bind.TransactOpts, order IERC7683GaslessCrossChainOrder, signature []byte, originFillerData []byte) (*types.Transaction, error) {
	return _SolverNetInbox.contract.Transact(opts, "openFor", order, signature, originFillerData)
}

// OpenFor is a paid mutator transaction binding the contract method 0x844fac8e.
//
// Solidity: function openFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes signature, bytes originFillerData) returns()
func (_SolverNetInbox *SolverNetInboxSession) OpenFor(order IERC7683GaslessCrossChainOrder, signature []byte, originFillerData []byte) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.OpenFor(&_SolverNetInbox.TransactOpts, order, signature, originFillerData)
}

// OpenFor is a paid mutator transaction binding the contract method 0x844fac8e.
//
// Solidity: function openFor((address,address,uint256,uint256,uint32,uint32,bytes32,bytes) order, bytes signature, bytes originFillerData) returns()
func (_SolverNetInbox *SolverNetInboxTransactorSession) OpenFor(order IERC7683GaslessCrossChainOrder, signature []byte, originFillerData []byte) (*types.Transaction, error) {
	return _SolverNetInbox.Contract.OpenFor(&_SolverNetInbox.TransactOpts, order, signature, originFillerData)
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
