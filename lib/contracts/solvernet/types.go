package solvernet

import (
	"encoding/binary"
	"math/big"
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

// Call is a bindings.SolverNetCall with Selector and Params joined into Data.
type Call struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

type (
	Expense = bindings.SolverNetTokenExpense
	Deposit = bindings.SolverNetDeposit

	Expenses []Expense
	Calls    []Call
)

func (cs Calls) ToBindings() []bindings.SolverNetCall {
	var out []bindings.SolverNetCall
	for _, c := range cs {
		var selector [4]byte
		if len(c.Data) >= 4 {
			copy(selector[:], c.Data[:4])
		}

		var params []byte
		if len(c.Data) > 4 {
			params = make([]byte, len(c.Data)-4)
			copy(params, c.Data[4:])
		}

		out = append(out, bindings.SolverNetCall{
			Target:   c.Target,
			Value:    c.Value,
			Selector: selector,
			Params:   params,
		})
	}

	return out
}

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
