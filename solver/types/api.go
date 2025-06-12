package types

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/uni"

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
	Debug              bool      `json:"debug"`
}

// CheckResponse is the response json for the /check endpoint.
type CheckResponse struct {
	Accepted          bool           `json:"accepted"`
	Rejected          bool           `json:"rejected"`
	RejectCode        RejectReason   `json:"rejectCode"`
	RejectReason      string         `json:"rejectReason"`
	RejectDescription string         `json:"rejectDescription"`
	Trace             map[string]any `json:"trace"` // If debug is true, result of debug_traceCall
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
	SourceChainID      uint64      `json:"sourceChainId"`
	DestinationChainID uint64      `json:"destChainId"`
	DepositToken       uni.Address `json:"depositToken"`
	ExpenseToken       uni.Address `json:"expenseToken"`
	IncludeFees        bool        `json:"includeFees"` // If true, include fees in the price calculation
}

type TokensResponse struct {
	Tokens []TokenResponse `json:"tokens"`
}

type TokenResponse struct {
	Enabled          bool         `json:"enabled"` // Deprecated, use ExpenseEnabled instead
	ExpenseEnabled   bool         `json:"expenseEnabled"`
	DepositEnabled   bool         `json:"depositEnabled"`
	Name             string       `json:"name"`
	Symbol           string       `json:"symbol"`
	ChainID          uint64       `json:"chainId"`
	Address          uni.Address  `json:"address"`
	Decimals         uint         `json:"decimals"`
	ExpenseMin       *hexutil.Big `json:"expenseMin"`
	ExpenseMax       *hexutil.Big `json:"expenseMax"`
	ExpenseInventory *hexutil.Big `json:"expenseInventory"`
}

type addrAmtJSON struct {
	Token  uni.Address  `json:"token"`
	Amount *hexutil.Big `json:"amount,omitempty"`
}

// AddrAmt represents a token address and amount pair, with the amount being optional.
// If amount is nil or zero, quote response should inform the amount.
type AddrAmt struct {
	Token  uni.Address
	Amount *big.Int
}

func (u AddrAmt) MarshalJSON() ([]byte, error) {
	return marshal(addrAmtJSON{
		Token:  u.Token,
		Amount: (*hexutil.Big)(u.Amount),
	})
}

func (u *AddrAmt) UnmarshalJSON(bz []byte) error {
	v := new(addrAmtJSON)
	if err := unmarshal(bz, v); err != nil {
		return err
	}

	u.Token = v.Token
	u.Amount = intOrZero(v.Amount)

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
	Portal    uni.Address `json:"portal"`
	Inbox     uni.Address `json:"inbox"`
	Outbox    uni.Address `json:"outbox"`
	Middleman uni.Address `json:"middleman"`
	Executor  uni.Address `json:"executor"`
}

// expenseJSON is a json marshal-able solvernt.Expense.
type expenseJSON struct {
	Spender uni.Address  `json:"spender"`
	Token   uni.Address  `json:"token"`
	Amount  *hexutil.Big `json:"amount"`
}

// Expense wraps solvernet.Expense to provide custom json marshaling.
type Expense solvernet.Expense

func (e *Expense) UnmarshalJSON(bz []byte) error {
	v := new(expenseJSON)
	if err := unmarshal(bz, v); err != nil {
		return err
	}

	if !v.Spender.IsEVM() || !v.Token.IsEVM() {
		return errors.New("expenses must be EVM addresses")
	}

	e.Spender = v.Spender.EVM()
	e.Token = v.Token.EVM()
	e.Amount = intOrZero(v.Amount)

	return nil
}

func (e Expense) MarshalJSON() ([]byte, error) {
	return marshal(expenseJSON{
		Spender: uni.EVMAddress(e.Spender),
		Token:   uni.EVMAddress(e.Token),
		Amount:  (*hexutil.Big)(e.Amount),
	})
}

// callJSON is a json marshal-able solvernet.Call.
type callJSON struct {
	Target common.Address `json:"target"`
	Data   hexutil.Bytes  `json:"data"`
	Value  *hexutil.Big   `json:"value"`
}

// Call wraps solvernet.Call to provide custom json marshaling.
type Call solvernet.Call

func (c *Call) UnmarshalJSON(bz []byte) error {
	v := new(callJSON)
	if err := unmarshal(bz, v); err != nil {
		return err
	}

	c.Target = v.Target
	c.Value = intOrZero(v.Value)
	c.Data = v.Data

	return nil
}

func (c Call) MarshalJSON() ([]byte, error) {
	return marshal(callJSON{
		Target: c.Target,
		Value:  (*hexutil.Big)(c.Value),
		Data:   c.Data,
	})
}

func intOrZero(i *hexutil.Big) *big.Int {
	if i == nil {
		return bi.Zero()
	}

	return i.ToInt()
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
		Deposit: AddrAmt{
			Token:  uni.EVMAddress(data.Deposit.Token),
			Amount: data.Deposit.Amount,
		},
		Expenses: expenses,
		Calls:    CallsFromBindings(data.Calls),
	}, nil
}

// RelayRequest is the expected request body for the /api/v1/relay endpoint.
// This endpoint accepts gasless orders with user signatures and submits them on behalf of users.
type RelayRequest struct {
	// The gasless cross-chain order to be submitted
	Order bindings.IERC7683GaslessCrossChainOrder `json:"order"`
	// User's signature authorizing the order
	Signature hexutil.Bytes `json:"signature"`
	// Optional filler-specific data (currently unused but part of ERC7683 spec)
	OriginFillerData hexutil.Bytes `json:"originFillerData,omitempty"`
}

// RelayResponse is the response json for the /relay endpoint.
type RelayResponse struct {
	// Whether the order was successfully submitted
	Success bool `json:"success"`
	// Transaction hash of the submitted openFor transaction
	TxHash common.Hash `json:"txHash,omitempty"`
	// Order ID that was created
	OrderID common.Hash `json:"orderId,omitempty"`
	// Error details if submission failed
	Error *RelayError `json:"error,omitempty"`
}

// RelayError represents an error in relay submission.
type RelayError struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}
