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

// Expense is a solver expense on the destination (matches bindings.SolverNetExpense).
type Expense struct {
	Spender common.Address `json:"spender"`
	Token   common.Address `json:"token"`
	Amount  *big.Int       `json:"amount"`
}

// Call is a call to be made on the destination (matches bindings.SolverNetCall).
type Call struct {
	Target   common.Address `json:"target"`
	Selector [4]byte        `json:"selector"`
	Value    *big.Int       `json:"value"`
	Params   []byte         `json:"params"`
}

// Deposit is a user deposit on the source (matches bindings.SolverNetDeposit).
type Deposit struct {
	Token  common.Address `json:"token"`
	Amount *big.Int       `json:"amount"`
}

// Expenses is a list of expenses.
type Expenses []Expense

// Calls is a list of calls.
type Calls []Call

// ToBindings converts a solvernet.Call to bindings.SolverNetCall.
func (c Call) ToBindings() bindings.SolverNetCall {
	return bindings.SolverNetCall{
		Target:   c.Target,
		Selector: c.Selector,
		Value:    c.Value,
		Params:   c.Params,
	}
}

// ToBindings converts a solvernet.Deposit to bindings.SolverNetDeposit.
func (d Deposit) ToBindings() bindings.SolverNetDeposit {
	return bindings.SolverNetDeposit{
		Token:  d.Token,
		Amount: d.Amount,
	}
}

// ToBindings converts a solvernet.Expense to bindings.SolverNetExpense.
func (e Expense) ToBindings() bindings.SolverNetExpense {
	return bindings.SolverNetExpense{
		Spender: e.Spender,
		Token:   e.Token,
		Amount:  e.Amount,
	}
}

// ToBindings converts a solvernet.Expenses to []bindings.SolverNetExpense.
func (es Expenses) ToBindings() []bindings.SolverNetExpense {
	var out []bindings.SolverNetExpense
	for _, e := range es {
		// native expenses are not submitted on chain, they are derived from call values
		if e.Token == (common.Address{}) {
			continue
		}

		out = append(out, e.ToBindings())
	}

	return out
}

// ToBindings converts a solvernet.Calls to []bindings.SolverNetCall.
func (cs Calls) ToBindings() []bindings.SolverNetCall {
	var out []bindings.SolverNetCall
	for _, c := range cs {
		out = append(out, c.ToBindings())
	}

	return out
}
