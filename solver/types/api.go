package types

import (
	"math/big"
	"net/http"

	"github.com/omni-network/omni/lib/contracts/solvernet"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type JSONResponse interface {
	StatusCode() int
}

// JSONErrorResponse is a json response for http errors (e.g 4xx, 5xx), not used for rejections.
type JSONErrorResponse struct {
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
	SourceChainID      uint64       `json:"sourceChainId"`
	DestinationChainID uint64       `json:"destChainId"`
	FillDeadline       uint32       `json:"fillDeadline"`
	Calls              JSONCalls    `json:"calls"`
	Expenses           JSONExpenses `json:"expenses"`
	Deposit            JSONDeposit  `json:"deposit"`
}

// CheckResponse is the response json for the /check endpoint.
type CheckResponse struct {
	Accepted          bool               `json:"accepted,omitempty"`
	Rejected          bool               `json:"rejected,omitempty"`
	RejectReason      string             `json:"rejectReason,omitempty"`
	RejectDescription string             `json:"rejectDescription,omitempty"`
	Error             *JSONErrorResponse `json:"error,omitempty"`
}

var _ JSONResponse = (*CheckResponse)(nil)

func (r CheckResponse) StatusCode() int {
	if r.Error != nil {
		return r.Error.Code
	}

	return http.StatusOK
}

// QuoteRequest is the expected request body for the /api/v1/quote endpoint.
// If deposit amount is omitted, the response will include the required deposit amount.
// If expense amount is omitted, the response will include the required expense amount.
type QuoteRequest struct {
	SourceChainID      uint64        `json:"sourceChainId"`
	DestinationChainID uint64        `json:"destChainId"`
	Deposit            JSONQuoteUnit `json:"deposit"`
	Expense            JSONQuoteUnit `json:"expense"`
}

// QuoteUnit represents a token and amount pair, with the amount being optional.
// If amount is nil or zero, quote response should inform the amount.
type QuoteUnit struct {
	Token  common.Address
	Amount *big.Int
}

// JSONQuoteUnit is a json marshal-able QuoteUnit.
type JSONQuoteUnit struct {
	Token  common.Address `json:"token"`
	Amount *hexutil.Big   `json:"amount,omitempty"`
}

func (qu QuoteUnit) ToJSON() JSONQuoteUnit {
	return JSONQuoteUnit{
		Token:  qu.Token,
		Amount: (*hexutil.Big)(qu.Amount),
	}
}
func (qu JSONQuoteUnit) Parse() QuoteUnit {
	return QuoteUnit{
		Token:  qu.Token,
		Amount: intOrZero(qu.Amount),
	}
}

// QuoteResponse is the response json for the /api/v1/quote endpoint.
type QuoteResponse struct {
	Deposit JSONQuoteUnit      `json:"deposit"`
	Expense JSONQuoteUnit      `json:"expense"`
	Error   *JSONErrorResponse `json:"error,omitempty"`
}

var _ JSONResponse = (*QuoteResponse)(nil)

func (r QuoteResponse) StatusCode() int {
	if r.Error != nil {
		return r.Error.Code
	}

	return http.StatusOK
}

// ContractsResponse is the response json for the /api/vi/contracts endpoint.
type ContractsResponse struct {
	Portal    string             `json:"portal,omitempty"`
	Inbox     string             `json:"inbox,omitempty"`
	Outbox    string             `json:"outbox,omitempty"`
	Middleman string             `json:"middleman,omitempty"`
	Error     *JSONErrorResponse `json:"error,omitempty"`
}

var _ JSONResponse = (*ContractsResponse)(nil)

func (r ContractsResponse) StatusCode() int {
	if r.Error != nil {
		return r.Error.Code
	}

	return http.StatusOK
}

// JSONExpense is a json marshal-able solvernt.Expense.
type JSONExpense struct {
	Spender common.Address `json:"spender"`
	Token   common.Address `json:"token"`
	Amount  *hexutil.Big   `json:"amount"`
}

// JSONCall is a json marshal-able solvernet.Call.
type JSONCall struct {
	Target common.Address `json:"target"`
	Data   *hexutil.Bytes `json:"data"`
	Value  *hexutil.Big   `json:"value"`
}

// JSONDeposit is a json marshal-able solvernet.Deposit.
type JSONDeposit struct {
	Token  common.Address `json:"token"`
	Amount *hexutil.Big   `json:"amount"`
}

type (
	JSONCalls    []JSONCall
	JSONExpenses []JSONExpense
)

func ToJSONCalls(calls []solvernet.Call) JSONCalls {
	var out JSONCalls
	for _, c := range calls {
		data := c.Data

		out = append(out, JSONCall{
			Target: c.Target,
			Value:  (*hexutil.Big)(c.Value),
			Data:   (*hexutil.Bytes)(&data),
		})
	}

	return out
}

func ToJSONExpenses(expenses []solvernet.Expense) JSONExpenses {
	var out JSONExpenses
	for _, e := range expenses {
		out = append(out, JSONExpense{
			Spender: e.Spender,
			Token:   e.Token,
			Amount:  (*hexutil.Big)(e.Amount),
		})
	}

	return out
}

func ToJSONDeposit(deposit solvernet.Deposit) JSONDeposit {
	return JSONDeposit{
		Token:  deposit.Token,
		Amount: (*hexutil.Big)(deposit.Amount),
	}
}

func (cs JSONCalls) Parse() solvernet.Calls {
	var out []solvernet.Call
	for _, c := range cs {
		out = append(out, solvernet.Call{
			Target: c.Target,
			Value:  intOrZero(c.Value),
			Data:   bzOrNil(c.Data),
		})
	}

	return out
}

func (es JSONExpenses) Parse() solvernet.Expenses {
	var out []solvernet.Expense
	for _, e := range es {
		out = append(out, solvernet.Expense{
			Spender: e.Spender,
			Token:   e.Token,
			Amount:  intOrZero(e.Amount),
		})
	}

	return out
}

func (d JSONDeposit) Parse() solvernet.Deposit {
	return solvernet.Deposit{
		Token:  d.Token,
		Amount: intOrZero(d.Amount),
	}
}

func intOrZero(i *hexutil.Big) *big.Int {
	if i == nil {
		return big.NewInt(0)
	}

	return i.ToInt()
}

func bzOrNil(b *hexutil.Bytes) []byte {
	if b == nil {
		return nil
	}

	return *b
}
