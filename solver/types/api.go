package types

import (
	"encoding/json"
	"math/big"
	"strings"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// JSONErrorResponse is a json response for http errors (e.g 4xx, 5xx), not used for rejections.
type JSONErrorResponse struct {
	Error JSONError `json:"error"`
}

type JSONError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// CheckRequest is the expected request body for the /api/v1/check endpoint.
//
// NOTE: Check request / response types mirror SolvertNet.OrderData, built
// specifically for EVM -> EVM orders via SolverNetInbox / Outbox contracts,
// with ERC7683 type hash matching SolverNetInbox.ORDERDATA_TYPEHASH.
//
// To support multiple order types with this api (e.g. EVM -> Solana, Solana -> EVM)
// we'd need a more generic request / response format that discriminates on
// order type hash.
type CheckRequest struct {
	SourceChainID      uint64    `json:"sourceChainId"`
	DestinationChainID uint64    `json:"destChainId"`
	FillDeadline       uint32    `json:"fillDeadline"`
	Calls              []Call    `json:"calls"`
	Expenses           []Expense `json:"expenses"`
	Deposit            AddrAmt   `json:"deposit"`
}

// CheckResponse is the response json for the /check endpoint.
type CheckResponse struct {
	Accepted          bool         `json:"accepted"`
	Rejected          bool         `json:"rejected"`
	RejectCode        RejectReason `json:"rejectCode"`
	RejectReason      string       `json:"rejectReason"`
	RejectDescription string       `json:"rejectDescription"`
}

// QuoteRequest is the expected request body for the /api/v1/quote endpoint.
// If deposit amount is omitted, the response will include the required deposit amount.
// If expense amount is omitted, the response will include the required expense amount.
type QuoteRequest struct {
	SourceChainID      uint64  `json:"sourceChainId"`
	DestinationChainID uint64  `json:"destChainId"`
	Deposit            AddrAmt `json:"deposit"`
	Expense            AddrAmt `json:"expense"`
}

type PriceRequest struct {
	SourceChainID      uint64         `json:"sourceChainId"`
	DestinationChainID uint64         `json:"destChainId"`
	DepositToken       common.Address `json:"depositToken"`
	ExpenseToken       common.Address `json:"expenseToken"`
}

type TokensResponse struct {
	Tokens []TokenResponse `json:"tokens"`
}

type TokenResponse struct {
	Enabled    bool
	Name       string
	Symbol     string
	ChainID    uint64
	Address    common.Address
	Decimals   uint
	ExpenseMin *big.Int
	ExpenseMax *big.Int
}

type tokenResponseJSON struct {
	Enabled    bool           `json:"enabled"`
	Name       string         `json:"name"`
	Symbol     string         `json:"symbol"`
	ChainID    uint64         `json:"chainId"`
	Address    common.Address `json:"address"`
	Decimals   uint           `json:"decimals"`
	ExpenseMin *bigIntJSON    `json:"expenseMin"`
	ExpenseMax *bigIntJSON    `json:"expenseMax"`
}

func (t TokenResponse) MarshalJSON() ([]byte, error) {
	return marshal(tokenResponseJSON{
		Enabled:    t.Enabled,
		Name:       t.Name,
		Symbol:     t.Symbol,
		ChainID:    t.ChainID,
		Address:    t.Address,
		Decimals:   t.Decimals,
		ExpenseMin: (*bigIntJSON)(t.ExpenseMin),
		ExpenseMax: (*bigIntJSON)(t.ExpenseMax),
	})
}

func (t *TokenResponse) UnmarshalJSON(bz []byte) error {
	v := new(tokenResponseJSON)
	if err := unmarshal(bz, v); err != nil {
		return err
	}

	t.Enabled = v.Enabled
	t.Name = v.Name
	t.Symbol = v.Symbol
	t.ChainID = v.ChainID
	t.Address = v.Address
	t.Decimals = v.Decimals
	t.ExpenseMin = v.ExpenseMin.ToIntOrZero()
	t.ExpenseMax = v.ExpenseMax.ToIntOrZero()

	return nil
}

// bigIntJSON is a wrapper around big.Int that can unmarshal both 0xhex and decimal string numbers.
// Note that similar to *big.Int, it must always be a pointer.
type bigIntJSON big.Int

// ToIntOrZero returns the value of the bigIntJSON as a *big.Int or zero if nil.
func (b *bigIntJSON) ToIntOrZero() *big.Int {
	if b == nil {
		return bi.Zero()
	}

	return (*big.Int)(b)
}

func (b *bigIntJSON) MarshalJSON() ([]byte, error) {
	return marshal((*hexutil.Big)(b))
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *bigIntJSON) UnmarshalJSON(input []byte) error {
	// First try hexutil.Big "0x1234".
	h := new(hexutil.Big)
	err := h.UnmarshalJSON(input)
	if jerr := new(json.UnmarshalTypeError); errors.As(err, &jerr) && jerr.Value == hexutil.ErrMissingPrefix.Error() { //nolint:revive // Explicit empty block
		// Swallow ErrMissingPrefix, number isn't 0xhex, try below as decimal.
	} else if err != nil {
		// Note that this also includes "not a string" errors.
		return errors.Wrap(err, "unmarshal hex number")
	} else /* err == nil */ {
		i := h.ToInt()
		*b = (bigIntJSON)(*i)

		return nil
	}

	// Trim quotes from the input string.
	decimal := strings.Trim(string(input), "\"")

	// If ErrMissingPrefix, try to unmarshal as a decimal string.
	i, ok := new(big.Int).SetString(decimal, 10)
	if !ok {
		return errors.New("invalid decimal number", "input", string(input))
	}

	*b = (bigIntJSON)(*i)

	return nil
}

type addrAmtJSON struct {
	Token  common.Address `json:"token"`
	Amount *bigIntJSON    `json:"amount,omitempty"`
}

// AddrAmt represents a token address and amount pair, with the amount being optional.
// If amount is nil or zero, quote response should inform the amount.
type AddrAmt struct {
	Token  common.Address
	Amount *big.Int
}

func (u AddrAmt) MarshalJSON() ([]byte, error) {
	return marshal(addrAmtJSON{
		Token:  u.Token,
		Amount: (*bigIntJSON)(u.Amount),
	})
}

func (u *AddrAmt) UnmarshalJSON(bz []byte) error {
	v := new(addrAmtJSON)
	if err := unmarshal(bz, v); err != nil {
		return err
	}

	u.Token = v.Token
	u.Amount = v.Amount.ToIntOrZero()

	return nil
}

// QuoteResponse is the response json for the /api/v1/quote endpoint.
type QuoteResponse struct {
	Deposit           AddrAmt      `json:"deposit"`
	Expense           AddrAmt      `json:"expense"`
	Rejected          bool         `json:"rejected"`
	RejectCode        RejectReason `json:"rejectCode"`
	RejectReason      string       `json:"rejectReason"`
	RejectDescription string       `json:"rejectDescription"`
}

// ContractsResponse is the response json for the /api/vi/contracts endpoint.
type ContractsResponse struct {
	Portal    common.Address `json:"portal"`
	Inbox     common.Address `json:"inbox"`
	Outbox    common.Address `json:"outbox"`
	Middleman common.Address `json:"middleman"`
	Executor  common.Address `json:"executor"`
}

// expenseJSON is a json marshal-able solvernt.Expense.
type expenseJSON struct {
	Spender common.Address `json:"spender"`
	Token   common.Address `json:"token"`
	Amount  *bigIntJSON    `json:"amount"`
}

// Expense wraps solvernet.Expense to provide custom json marshaling.
type Expense solvernet.Expense

func (e *Expense) UnmarshalJSON(bz []byte) error {
	v := new(expenseJSON)
	if err := unmarshal(bz, v); err != nil {
		return err
	}

	e.Spender = v.Spender
	e.Token = v.Token
	e.Amount = v.Amount.ToIntOrZero()

	return nil
}

func (e Expense) MarshalJSON() ([]byte, error) {
	return marshal(expenseJSON{
		Spender: e.Spender,
		Token:   e.Token,
		Amount:  (*bigIntJSON)(e.Amount),
	})
}

// callJSON is a json marshal-able solvernet.Call.
type callJSON struct {
	Target common.Address `json:"target"`
	Data   hexutil.Bytes  `json:"data"`
	Value  *bigIntJSON    `json:"value"`
}

// Call wraps solvernet.Call to provide custom json marshaling.
type Call solvernet.Call

func (c *Call) UnmarshalJSON(bz []byte) error {
	v := new(callJSON)
	if err := unmarshal(bz, v); err != nil {
		return err
	}

	c.Target = v.Target
	c.Value = v.Value.ToIntOrZero()
	c.Data = v.Data

	return nil
}

func (c Call) MarshalJSON() ([]byte, error) {
	return marshal(callJSON{
		Target: c.Target,
		Value:  (*bigIntJSON)(c.Value),
		Data:   c.Data,
	})
}

func marshal(v any) ([]byte, error) {
	bz, err := json.Marshal(v)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return bz, nil
}

func unmarshal(bz []byte, v any) error {
	if err := json.Unmarshal(bz, v); err != nil {
		return errors.Wrap(err, "unmarshal")
	}

	return nil
}

func CallsToBindings(calls []Call) []bindings.SolverNetCall {
	var resp []bindings.SolverNetCall
	for _, c := range calls {
		resp = append(resp, solvernet.Call(c).ToBinding())
	}

	return resp
}

func CallsFromBindings(calls []bindings.SolverNetCall) []Call {
	var resp []Call
	for _, c := range calls {
		resp = append(resp, Call(solvernet.CallFromBinding(c)))
	}

	return resp
}

func ExpensesToBindings(expenses []Expense) []solvernet.Expense {
	var resp []solvernet.Expense
	for _, e := range expenses {
		resp = append(resp, solvernet.Expense(e))
	}

	return resp
}

func ExpensesFromBindings(expenses []solvernet.Expense) []Expense {
	var resp []Expense
	for _, e := range expenses {
		resp = append(resp, Expense(e))
	}

	return resp
}

func CheckRequestFromOrderData(srcChainID uint64, data bindings.SolverNetOrderData) (CheckRequest, error) {
	deadline, err := umath.ToUint32(time.Now().Add(time.Hour).Unix())
	if err != nil {
		return CheckRequest{}, err
	}

	expenses := ExpensesFromBindings(data.Expenses)

	// Add native calls as expenses.
	// Note this is inconsistent with OpenOrder where native calls MUST NOT be included as expenses.
	for _, call := range data.Calls {
		if call.Value == nil || bi.IsZero(call.Value) {
			continue
		}
		expenses = append(expenses, Expense{
			// TODO(corver): What should spender be?
			Amount: bi.Clone(call.Value),
		})
	}

	return CheckRequest{
		SourceChainID:      srcChainID,
		DestinationChainID: data.DestChainId,
		FillDeadline:       deadline,
		Deposit:            AddrAmt(data.Deposit),
		Expenses:           expenses,
		Calls:              CallsFromBindings(data.Calls),
	}, nil
}
