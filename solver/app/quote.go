package app

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NOTE: Quote request / response types mirror SolvertNet.OrderData, built
// specifically for EVM -> EVM orders via SolverNetInbox / Outbox contracts,
// with ERC7683 type hash matching SolverNetInbox.ORDERDATA_TYPEHASH.
//
// To support multiple order types with this api (e.g. EVM -> Solana, Solana -> EVM)
// we'd need a more generic request / response format that discriminates on
// order type hash.

// QuoteRequest is the expected request body for the /api/v1/quote endpoint.
type QuoteRequest struct {
	SourceChainID      uint64         `json:"sourceChainId"`
	DestinationChainID uint64         `json:"destChainId"`
	FillDeadline       uint32         `json:"fillDeadline"`
	Calls              []Call         `json:"calls"`
	Expenses           []Expense      `json:"expenses"`
	DepositToken       common.Address `json:"depositToken"`
}

// QuoteResponse is the response json for the /quote endpoint.
type QuoteResponse struct {
	Rejected          bool               `json:"rejected,omitempty"`
	RejectReason      string             `json:"rejectReason,omitempty"`
	RejectDescription string             `json:"rejectDescription,omitempty"`
	Deposit           *Deposit           `json:"deposit,omitempty"`
	Error             *JSONErrorResponse `json:"error,omitempty"`
}

var _ JSONResponse = (*QuoteResponse)(nil)

func (r QuoteResponse) StatusCode() int {
	if r.Error != nil {
		return r.Error.Code
	}

	return http.StatusOK
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

type quoteFunc func(context.Context, QuoteRequest) (Deposit, error)

// newQuoter returns a quoteFunc that can be used to quote deposits for expenses.
// It is the logic behind the /quote endpoint.
func newQuoter(backends ethbackend.Backends, solverAddr, inboxAddr, outboxAddr common.Address) quoteFunc {
	return func(ctx context.Context, req QuoteRequest) (Deposit, error) {
		if req.SourceChainID == req.DestinationChainID {
			return Deposit{}, newRejection(rejectSameChain, errors.New("source and destination chain are the same"))
		}

		srcBackend, err := backends.Backend(req.SourceChainID)
		if err != nil {
			return Deposit{}, newRejection(rejectUnsupportedSrcChain, errors.New("unsupported source chain", "chain_id", req.SourceChainID))
		}

		dstBackend, err := backends.Backend(req.DestinationChainID)
		if err != nil {
			return Deposit{}, newRejection(rejectUnsupportedDestChain, errors.New("unsupported destination chain", "chain_id", req.DestinationChainID))
		}

		depositTkn, ok := tokens.Find(req.SourceChainID, req.DepositToken)
		if !ok {
			return Deposit{}, newRejection(rejectUnsupportedDeposit, errors.New("unsupported deposit token", "addr", req.DepositToken))
		}

		expenses, err := parseQuoteExpenses(req)
		if err != nil {
			return Deposit{}, err
		}

		quote, err := getQuote([]Token{depositTkn}, expenses)
		if err != nil {
			return Deposit{}, err
		}

		if err := checkLiquidity(ctx, expenses, dstBackend, solverAddr); err != nil {
			return Deposit{}, err
		}

		if err := checkApprovals(ctx, expenses, dstBackend, solverAddr, outboxAddr); err != nil {
			return Deposit{}, err
		}

		orderID, err := getNextOrderID(ctx, srcBackend, inboxAddr)
		if err != nil {
			return Deposit{}, err
		}

		fillOriginData, err := getFillOriginData(req)
		if err != nil {
			return Deposit{}, err
		}

		if err := checkFill(ctx, dstBackend, orderID, fillOriginData, nativeAmt(expenses), solverAddr, outboxAddr); err != nil {
			return Deposit{}, err
		}

		return Deposit{
			Token:  quote[0].Token.Address,
			Amount: quote[0].Amount,
		}, nil
	}
}

// getNextOrderID returns the next order ID for the given inbox.
func getNextOrderID(ctx context.Context, client ethclient.Client, inboxAddr common.Address) (OrderID, error) {
	inbox, err := bindings.NewSolverNetInbox(inboxAddr, client)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "new inbox")
	}

	orderID, err := inbox.GetNextId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return OrderID{}, errors.Wrap(err, "get next order id")
	}

	return orderID, nil
}

// getFillOriginData returns packed fill origin data for a quote request.
func getFillOriginData(req QuoteRequest) ([]byte, error) {
	calls := make([]bindings.SolverNetCall, len(req.Calls))
	for i, c := range req.Calls {
		calls[i] = bindings.SolverNetCall{
			Target:   c.Target,
			Selector: c.Selector,
			Value:    c.Value,
			Params:   c.Params,
		}
	}

	expenses := make([]bindings.SolverNetExpense, len(req.Expenses))
	for i, e := range req.Expenses {
		expenses[i] = bindings.SolverNetExpense{
			Spender: e.Spender,
			Token:   e.Token,
			Amount:  e.Amount,
		}
	}

	fillOriginData := bindings.SolverNetFillOriginData{
		FillDeadline: req.FillDeadline,
		SrcChainId:   req.SourceChainID,
		DestChainId:  req.DestinationChainID,
		Expenses:     expenses,
		Calls:        calls,
	}

	fillOriginDataBz, err := solvernet.PackFillOriginData(fillOriginData)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill origin data")
	}

	return fillOriginDataBz, nil
}

// newQuoteHandler returns a handler for the /quote endpoint.
// It is responsible to http request / response handling, and delegates
// logic to a quoteFunc.
func newQuoteHandler(quoteFunc quoteFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rr *http.Request) {
		ctx := rr.Context()

		w.Header().Set("Content-Type", "application/json")

		writeError := func(statusCode int, err error) {
			log.DebugErr(ctx, "Error handling /quote request", err)

			writeJSON(ctx, w, QuoteResponse{
				Error: &JSONErrorResponse{
					Code:    statusCode,
					Status:  http.StatusText(statusCode),
					Message: removeBUG(err.Error()),
				},
			})
		}

		var req QuoteRequest
		if err := json.NewDecoder(rr.Body).Decode(&req); err != nil {
			writeError(http.StatusBadRequest, errors.Wrap(err, "decode request"))
			return
		}

		deposit, err := quoteFunc(ctx, req)
		if r := new(RejectionError); errors.As(err, &r) { // RejectionError
			writeJSON(ctx, w, QuoteResponse{
				Rejected:          true,
				RejectReason:      r.Reason.String(),
				RejectDescription: r.Err.Error(),
			})
		} else if err != nil { // Error
			writeError(http.StatusInternalServerError, err)
		} else {
			writeJSON(ctx, w, QuoteResponse{Deposit: &deposit}) // Success
		}
	})
}

// getQuote returns payment in `depositTkns` required to pay for `expenses`.
//
// For now, this is a simple quote that requires a single expense, paid
// for by an equal amount of an equivalent deposit token. Token equivalence is
// determined by symbol (ex arbitrum "ETH" is equivalent to optimism "ETH").
func getQuote(depositTkns []Token, expenses []Payment) ([]Payment, error) {
	if len(depositTkns) != 1 {
		return nil, newRejection(rejectInvalidDeposit, errors.New("only single deposit token supported"))
	}

	if len(expenses) != 1 {
		return nil, newRejection(rejectInvalidExpense, errors.New("only single expense supported"))
	}

	expense := expenses[0]
	depositTkn := depositTkns[0]

	if expense.Token.Symbol != depositTkn.Symbol {
		return nil, newRejection(rejectInvalidDeposit, errors.New("deposit token must match expense token"))
	}

	// make sure chain class (e.g. mainnet, testnet) matches
	// we should reject with UnsupportedDestChain before this. the solver is
	// initialized by network, which only includes chains of the same class
	if expense.Token.ChainClass != depositTkn.ChainClass {
		return nil, newRejection(rejectInvalidDeposit, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"))
	}

	return []Payment{
		{
			Token:  depositTkn,
			Amount: expense.Amount,
		},
	}, nil
}

// coversQuote checks if `deposits` match or exceed a `quote` for expenses.
func coversQuote(deposits, quote []Payment) error {
	byTkn := func(ps []Payment) map[Token]*big.Int {
		res := make(map[Token]*big.Int)
		for _, p := range ps {
			res[p.Token] = p.Amount
		}

		return res
	}

	quoteByTkn := byTkn(quote)
	depositsByTkn := byTkn(deposits)

	for tkn, q := range quoteByTkn {
		d, ok := depositsByTkn[tkn]
		if !ok {
			return newRejection(rejectInsufficientDeposit, errors.New("missing deposit", "token", tkn))
		}

		if d.Cmp(q) < 0 {
			return newRejection(rejectInsufficientDeposit, errors.New("insufficient deposit", "token", tkn, "deposit", d, "quote", q))
		}
	}

	return nil
}

func parseQuoteExpenses(req QuoteRequest) ([]Payment, error) {
	var expenses []Payment

	hasNative := false
	for _, e := range req.Expenses {
		tkn, ok := tokens.Find(req.DestinationChainID, e.Token)
		if !ok {
			return nil, newRejection(rejectUnsupportedExpense, errors.New("unsupported expense token", "addr", e.Token))
		}

		if tkn.IsNative() {
			if hasNative {
				return nil, newRejection(rejectInvalidExpense, errors.New("multiple native expenses not supported"))
			}

			hasNative = true
		}

		expenses = append(expenses, Payment{
			Token:  tkn,
			Amount: e.Amount,
		})
	}

	return expenses, nil
}
