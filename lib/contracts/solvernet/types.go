package solvernet

import (
	"encoding/binary"
	"strconv"

	"github.com/omni-network/omni/contracts/bindings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	StatusInvalid  OrderStatus = 0
	StatusPending  OrderStatus = 1
	StatusRejected OrderStatus = 2
	StatusClosed   OrderStatus = 3
	StatusFilled   OrderStatus = 4
	StatusClaimed  OrderStatus = 5
)

type (
	OrderID        [32]byte
	OrderStatus    uint8
	OrderResolved  = bindings.IERC7683ResolvedCrossChainOrder
	OrderState     = bindings.ISolverNetInboxOrderState
	FillOriginData = bindings.SolverNetFillOriginData
)

// Uint64 returns the order ID as a BigEndian uint64 (monotonically incrementing number).
func (id OrderID) Uint64() uint64 {
	return binary.BigEndian.Uint64(id[32-8:])
}

// String returns the Uint64 representation of the order ID as a string.
func (id OrderID) String() string {
	return strconv.FormatUint(id.Uint64(), 10)
}

func (s OrderStatus) String() string {
	switch s {
	case StatusInvalid:
		return "invalid"
	case StatusPending:
		return "pending"
	case StatusRejected:
		return "rejected"
	case StatusClosed:
		return "closed"
	case StatusFilled:
		return "filled"
	case StatusClaimed:
		return "claimed"
	default:
		return "unknown"
	}
}

func (s OrderStatus) Uint8() uint8 {
	return uint8(s)
}

type (
	Expense  = bindings.SolverNetTokenExpense
	Deposit  = bindings.SolverNetDeposit
	Call     = bindings.SolverNetCall
	Expenses []Expense
	Calls    []Call
)

func (es Expenses) NoNative() Expenses {
	var out Expenses
	for _, e := range es {
		if !isNative(e) {
			out = append(out, e)
		}
	}

	return out
}

func isNative(e Expense) bool { return e.Token == (common.Address{}) }
