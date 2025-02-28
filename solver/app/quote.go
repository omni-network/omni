package app

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/solver/types"
)

// standardFeeBips is the standard fee charged by the solver (0.3%).
const standardFeeBips = 30

type (
	QuoteRequest  = types.QuoteRequest
	QuoteResponse = types.QuoteResponse
	QuoteUnit     = types.QuoteUnit

	quoteFunc func(QuoteRequest) QuoteResponse
)

// quoter is quoteFunc that can be used to quoter an expense or deposit.
// It is the logic behind the /quoter endpoint.
func quoter(req QuoteRequest) QuoteResponse {
	deposit := req.Deposit.Parse()
	expense := req.Expense.Parse()

	isDepositQuote := deposit.Amount == nil || deposit.Amount.Sign() == 0
	isExpenseQuote := expense.Amount == nil || expense.Amount.Sign() == 0

	returnErr := func(code int, msg string) QuoteResponse {
		return QuoteResponse{
			Error: &JSONErrorResponse{
				Code:    code,
				Status:  http.StatusText(code),
				Message: msg,
			},
		}
	}

	if isDepositQuote == isExpenseQuote {
		return returnErr(http.StatusBadRequest, "deposit and expense amount cannot be both zero or both non-zero")
	}

	depositTkn, ok := tokens.Find(req.SourceChainID, req.Deposit.Token)
	if !ok {
		return returnErr(http.StatusNotFound, "unsupported deposit token")
	}

	expenseTkn, ok := tokens.Find(req.DestinationChainID, req.Expense.Token)
	if !ok {
		return returnErr(http.StatusNotFound, "unsupported expense token")
	}

	returnQuote := func(depositAmt, expenseAmt *big.Int) QuoteResponse {
		return QuoteResponse{
			Deposit: QuoteUnit{
				Token:  deposit.Token,
				Amount: depositAmt,
			}.ToJSON(),
			Expense: QuoteUnit{
				Token:  expense.Token,
				Amount: expenseAmt,
			}.ToJSON(),
		}
	}

	if isDepositQuote {
		quoted, err := quoteDeposit(depositTkn, Payment{Token: expenseTkn, Amount: expense.Amount})
		if err != nil {
			return returnErr(http.StatusBadRequest, err.Error())
		}

		return returnQuote(quoted.Amount, expense.Amount)
	}

	quoted, err := QuoteExpense(expenseTkn, Payment{Token: depositTkn, Amount: deposit.Amount})
	if err != nil {
		return returnErr(http.StatusBadRequest, err.Error())
	}

	return returnQuote(deposit.Amount, quoted.Amount)
}

// newQuoteHandler returns a handler for the /quote endpoint.
// It is responsible to http request / response handling, and delegates
// logic to a quoteFunc.
func newQuoteHandler(quoteFunc quoteFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rr *http.Request) {
		ctx := rr.Context()

		w.Header().Set("Content-Type", "application/json")

		var req QuoteRequest
		if err := json.NewDecoder(rr.Body).Decode(&req); err != nil {
			writeJSON(ctx, w, QuoteResponse{
				Error: &JSONErrorResponse{
					Code:    http.StatusBadRequest,
					Status:  http.StatusText(http.StatusBadRequest),
					Message: err.Error(),
				},
			})

			return
		}

		res := quoteFunc(req)
		writeJSON(ctx, w, res)
	})
}

// getQuote returns payment in `depositTkns` required to pay for `expenses`.
func getQuote(depositTkns []Token, expenses []Payment) ([]Payment, error) {
	if len(depositTkns) != 1 {
		return nil, newRejection(rejectInvalidDeposit, errors.New("only single deposit token supported"))
	}

	if len(expenses) != 1 {
		return nil, newRejection(rejectInvalidExpense, errors.New("only single expense supported"))
	}

	expense := expenses[0]
	depositTkn := depositTkns[0]

	deposit, err := quoteDeposit(depositTkn, expense)
	if err != nil {
		return nil, err
	}

	return []Payment{deposit}, nil
}

// quoteDeposit returns the deposit required to cover `expense`.
func quoteDeposit(tkn Token, expense Payment) (Payment, error) {
	if expense.Token.Symbol != tkn.Symbol {
		return Payment{}, newRejection(rejectInvalidDeposit, errors.New("deposit token must match expense token"))
	}

	if expense.Token.ChainClass != tkn.ChainClass {
		// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
		return Payment{}, newRejection(rejectInvalidDeposit, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"))
	}

	return Payment{
		Token:  tkn,
		Amount: depositFor(expense.Amount, feeBipsFor(tkn)),
	}, nil
}

// QuoteExpense returns the expense allowed for `deposit`.
func QuoteExpense(tkn Token, deposit Payment) (Payment, error) {
	if deposit.Token.Symbol != tkn.Symbol {
		return Payment{}, newRejection(rejectInvalidDeposit, errors.New("deposit token must match expense token"))
	}

	if deposit.Token.ChainClass != tkn.ChainClass {
		// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
		return Payment{}, newRejection(rejectInvalidDeposit, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"))
	}

	return Payment{
		Token:  tkn,
		Amount: expenseFor(deposit.Amount, feeBipsFor(tkn)),
	}, nil
}

// feeBipsFor returns the fee in bips for a given token.
func feeBipsFor(tkn Token) int64 {
	// if OMNI, charge no fee
	if tkn.IsOMNI() {
		return 0
	}

	return standardFeeBips
}

// depositFor returns the deposit required to cover `expense` with a fee in bips.
func depositFor(expense *big.Int, bips int64) *big.Int {
	// deposit = expense + (expense * bips / 10_000)

	quo := big.NewInt(10_000)
	num := new(big.Int).Mul(expense, big.NewInt(bips))
	fee := new(big.Int).Div(num, quo)

	return new(big.Int).Add(expense, fee)
}

// expenseFor returns the expense allowed for `deposit` with a fee in bips.
func expenseFor(deposit *big.Int, bips int64) *big.Int {
	// expense = 10_000 * d / (10_000 + bips)

	quo := big.NewInt(10_000 + bips)
	num := new(big.Int).Mul(deposit, big.NewInt(10_000))

	return new(big.Int).Div(num, quo)
}
