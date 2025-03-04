package app

import (
	"context"
	"math/big"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/solver/types"
)

// standardFeeBips is the standard fee charged by the solver (0.3%).
const standardFeeBips = 30

type quoteFunc func(context.Context, types.QuoteRequest) (types.QuoteResponse, error)

// quoter is quoteFunc that can be used to quoter an expense or deposit.
// It is the logic behind the /quoter endpoint.
func quoter(_ context.Context, req types.QuoteRequest) (types.QuoteResponse, error) {
	deposit := req.Deposit
	expense := req.Expense

	isDepositQuote := deposit.Amount == nil || deposit.Amount.Sign() == 0
	isExpenseQuote := expense.Amount == nil || expense.Amount.Sign() == 0

	returnErr := func(code int, msg string) (types.QuoteResponse, error) {
		return types.QuoteResponse{}, newAPIError(errors.New(msg), code)
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

	returnQuote := func(depositAmt, expenseAmt *big.Int) types.QuoteResponse {
		return types.QuoteResponse{
			Deposit: types.AddrAmt{
				Token:  deposit.Token,
				Amount: depositAmt,
			},
			Expense: types.AddrAmt{
				Token:  expense.Token,
				Amount: expenseAmt,
			},
		}
	}

	if isDepositQuote {
		quoted, err := QuoteDeposit(depositTkn, TokenAmt{Token: expenseTkn, Amount: expense.Amount})
		if err != nil {
			return types.QuoteResponse{}, newAPIError(err, http.StatusBadRequest)
		}

		return returnQuote(quoted.Amount, expense.Amount), nil
	}

	quoted, err := quoteExpense(expenseTkn, TokenAmt{Token: depositTkn, Amount: deposit.Amount})
	if err != nil {
		return types.QuoteResponse{}, newAPIError(err, http.StatusBadRequest)
	}

	return returnQuote(deposit.Amount, quoted.Amount), nil
}

// getQuote returns payment in `depositTkns` required to pay for `expenses`.
func getQuote(depositTkns []Token, expenses []TokenAmt) ([]TokenAmt, error) {
	if len(depositTkns) != 1 {
		return nil, newRejection(rejectInvalidDeposit, errors.New("only single deposit token supported"))
	}

	if len(expenses) != 1 {
		return nil, newRejection(rejectInvalidExpense, errors.New("only single expense supported"))
	}

	expense := expenses[0]
	depositTkn := depositTkns[0]

	deposit, err := QuoteDeposit(depositTkn, expense)
	if err != nil {
		return nil, err
	}

	return []TokenAmt{deposit}, nil
}

// QuoteDeposit returns the deposit required to cover `expense`.
func QuoteDeposit(tkn Token, expense TokenAmt) (TokenAmt, error) {
	if expense.Token.Symbol != tkn.Symbol {
		return TokenAmt{}, newRejection(rejectInvalidDeposit, errors.New("deposit token must match expense token"))
	}

	if expense.Token.ChainClass != tkn.ChainClass {
		// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
		return TokenAmt{}, newRejection(rejectInvalidDeposit, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"))
	}

	return TokenAmt{
		Token:  tkn,
		Amount: depositFor(expense.Amount, feeBipsFor(tkn)),
	}, nil
}

// QuoteExpense returns the expense allowed for `deposit`.
func quoteExpense(tkn Token, deposit TokenAmt) (TokenAmt, error) {
	if deposit.Token.Symbol != tkn.Symbol {
		return TokenAmt{}, newRejection(rejectInvalidDeposit, errors.New("deposit token must match expense token"))
	}

	if deposit.Token.ChainClass != tkn.ChainClass {
		// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
		return TokenAmt{}, newRejection(rejectInvalidDeposit, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"))
	}

	return TokenAmt{
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
