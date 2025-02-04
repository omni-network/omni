package app

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"

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
	FillDeadline       uint64         `json:"fillDeadline"`
	Calls              []Call         `json:"calls"`
	Expenses           []Expense      `json:"expenses"`
	DepositToken       common.Address `json:"depositToken"`
}

// QuoteResponse is the response json for the /quote endpoint.
type QuoteResponse struct {
	Rejected          bool    `json:"rejected"`
	RejectReason      string  `json:"rejectReason"`
	RejectDescription string  `json:"rejectDescription"`
	Deposit           Deposit `json:"deposit"`
	Error             Error   `json:"error"`
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

// Error is a json response for http errors (e.g 4xx, 5xx), not used for rejections.
type Error struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// newQuoteHandler returns a handler for the /quote endpoint.
func newQuoteHandler(backends ethbackend.Backends, solverAddr common.Address) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		// TODO: better request / response logging

		write := func(res QuoteResponse) {
			if err := json.NewEncoder(w).Encode(res); err != nil {
				log.Error(ctx, "[BUG] error writing /quote response", err)
				w.WriteHeader(http.StatusInternalServerError)

				return
			}

			status := http.StatusOK
			if res.Error.Code != 0 {
				status = res.Error.Code
			}

			w.WriteHeader(status)
		}

		writeError := func(statusCode int, err error) {
			log.DebugErr(ctx, "error handling /quote request", err, "code", statusCode)

			res := QuoteResponse{
				Error: Error{
					Code:    statusCode,
					Status:  http.StatusText(statusCode),
					Message: err.Error(),
				},
			}

			write(res)
		}

		writeRejectOrErr := func(reject RejectOrErr) {
			if reject.ShouldError() {
				writeError(http.StatusInternalServerError, reject.Err)
				return
			}

			res := QuoteResponse{
				Rejected:          true,
				RejectReason:      reject.Reason.String(),
				RejectDescription: reject.Err.Error(),
			}

			write(res)
		}

		writeReject := func(reason rejectReason, err error) {
			writeRejectOrErr(RejectOrErr{Reason: reason, Err: err})
		}

		var req QuoteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(http.StatusBadRequest, errors.Wrap(err, "decode request"))
			return
		}

		backend, err := backends.Backend(req.DestinationChainID)
		if err != nil {
			writeReject(rejectUnsupportedDestChain, err)
			return
		}

		depositTkn, ok := tokens.find(req.SourceChainID, req.DepositToken)
		if !ok {
			writeReject(rejectUnsupportedDeposit, errors.New("unsupported deposit token", "addr", req.DepositToken))
			return
		}

		var expenses []Payment
		for _, e := range req.Expenses {
			tkn, ok := tokens.find(req.DestinationChainID, e.Token)
			if !ok {
				writeReject(rejectUnsupportedExpense, errors.New("unsupported expense token", "addr", e.Token))
				return
			}
			expenses = append(expenses, Payment{
				Token:  tkn,
				Amount: e.Amount,
			})
		}

		quote, rejectOrErr := getQuote([]Token{depositTkn}, expenses)
		if rejectOrErr.ShouldReturn() {
			writeRejectOrErr(rejectOrErr)
			return
		}

		rejectOrErr = checkLiquidity(ctx, expenses, backend, solverAddr)
		if rejectOrErr.ShouldReturn() {
			writeRejectOrErr(rejectOrErr)
			return
		}

		w.WriteHeader(http.StatusOK)

		res := QuoteResponse{
			Deposit: Deposit{
				Token:  quote[0].Token.Address,
				Amount: quote[0].Amount,
			},
		}

		write(res)
	})
}

// getQuote returns payment in `depositTkns` required to pay for `expenses`.
//
// For now, this is a simple quote that requires a single expense, paid
// for by an equal amount of an equivalent deposit token. Token equivalence is
// determined by symbol (ex arbitrum "ETH" is equivalent to optimism "ETH").
func getQuote(depositTkns []Token, expenses []Payment) ([]Payment, RejectOrErr) {
	if len(depositTkns) != 1 {
		return nil, RejectOrErr{
			Reason: rejectInvalidDeposit,
			Err:    errors.New("only single deposit token supported"),
		}
	}

	if len(expenses) != 1 {
		return nil, RejectOrErr{
			Reason: rejectInvalidExpense,
			Err:    errors.New("only single expense supported"),
		}
	}

	expense := expenses[0]
	depositTkn := depositTkns[0]

	if expense.Token.Symbol != depositTkn.Symbol {
		return nil, RejectOrErr{
			Reason: rejectInvalidDeposit,
			Err:    errors.New("deposit token must match expense token"),
		}
	}

	// make sure chain class (e.g. mainnet, testnet) matches
	// we should reject with UnsupportedDestChain before this. the solver is
	// initialized by network, which only includes chains of the same class
	if expense.Token.ChainClass != depositTkn.ChainClass {
		return nil, RejectOrErr{
			Reason: rejectInvalidDeposit,
			Err:    errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"),
		}
	}

	return []Payment{
		{
			Token:  depositTkn,
			Amount: expense.Amount,
		},
	}, RejectOrErr{}
}

// coversQuote checks if `deposits` match or exceed a `quote` for expenses.
func coversQuote(deposits, quote []Payment) RejectOrErr {
	if len(quote) != len(deposits) {
		return RejectOrErr{}
	}

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
			return RejectOrErr{
				Reason: rejectInsufficientDeposit,
				Err:    errors.New("missing deposit", "token", tkn),
			}
		}

		if d.Cmp(q) < 0 {
			return RejectOrErr{
				Reason: rejectInsufficientDeposit,
				Err:    errors.New("insufficient deposit", "token", tkn, "deposit", d, "quote", q),
			}
		}
	}

	return RejectOrErr{}
}
