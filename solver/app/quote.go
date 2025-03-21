package app

import (
	"context"
	"math/big"
	"net/http"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	ltokens "github.com/omni-network/omni/lib/tokens"
	stokens "github.com/omni-network/omni/solver/tokens"
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

	depositTkn, ok := stokens.ByAddress(req.SourceChainID, req.Deposit.Token)
	if !ok {
		return returnErr(http.StatusNotFound, "unsupported deposit token")
	}

	expenseTkn, ok := stokens.ByAddress(req.DestinationChainID, req.Expense.Token)
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
func getQuote(depositTkns []stokens.Token, expenses []TokenAmt) ([]TokenAmt, error) {
	if len(depositTkns) != 1 {
		return nil, newRejection(types.RejectInvalidDeposit, errors.New("only single deposit token supported"))
	}

	if len(expenses) != 1 {
		return nil, newRejection(types.RejectInvalidExpense, errors.New("only single expense supported"))
	}

	expense := expenses[0]
	depositTkn := depositTkns[0]

	deposit, err := QuoteDeposit(depositTkn, expense)
	if err != nil {
		return nil, err
	}

	return []TokenAmt{deposit}, nil
}

// QuoteDeposit returns the source chain deposit required to cover `expense`.
func QuoteDeposit(depositTkn stokens.Token, expense TokenAmt) (TokenAmt, error) {
	if !areEqualBySymbol(depositTkn, expense.Token) {
		return TokenAmt{}, newRejection(types.RejectInvalidDeposit, errors.New("deposit token must match expense token"))
	}

	if expense.Token.ChainClass != depositTkn.ChainClass {
		// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
		return TokenAmt{}, newRejection(types.RejectInvalidDeposit, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"))
	}

	return TokenAmt{
		Token:  depositTkn,
		Amount: depositFor(expense.Amount, feeBipsFor(depositTkn)),
	}, nil
}

// QuoteExpense returns the destination chain expense allowed for `deposit`.
func quoteExpense(expenseTkn stokens.Token, deposit TokenAmt) (TokenAmt, error) {
	if !areEqualBySymbol(deposit.Token, expenseTkn) {
		return TokenAmt{}, newRejection(types.RejectInvalidDeposit, errors.New("deposit token must match expense token"))
	}

	if deposit.Token.ChainClass != expenseTkn.ChainClass {
		// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
		return TokenAmt{}, newRejection(types.RejectInvalidDeposit, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)"))
	}

	return TokenAmt{
		Token:  expenseTkn,
		Amount: expenseFor(deposit.Amount, feeBipsFor(expenseTkn)),
	}, nil
}

func areEqualBySymbol(a, b stokens.Token) bool {
	if a.Symbol == b.Symbol {
		return true
	}

	equivalents := map[string]string{}
	makeEq := func(a, b string) {
		equivalents[a] = b
		equivalents[b] = a
	}

	// consider stETH and ETH as equivalent
	makeEq(ltokens.STETH.Symbol, ltokens.ETH.Symbol)

	return equivalents[a.Symbol] == b.Symbol
}

// feeBipsFor returns the fee in bips for a given token.
func feeBipsFor(tkn stokens.Token) int64 {
	// if OMNI, charge no fee
	if tkn.IsOMNI() {
		return 0
	}

	return standardFeeBips
}

// depositFor returns the deposit required to cover `expense` with a fee in bips.
func depositFor(expense *big.Int, bips int64) *big.Int {
	// deposit = expense + ceil(expense * bips / 10_000)

	feeDividend := bi.Mul(expense, bi.N(bips))
	feeDividend = bi.Add(feeDividend, bi.N(9_999)) // Add 9_999 to dividend to round up.
	feeDivisor := bi.N(10_000)

	fee := bi.Div(feeDividend, feeDivisor)

	return bi.Add(expense, fee)
}

// expenseFor returns the expense allowed for `deposit` with a fee in bips.
func expenseFor(deposit *big.Int, bips int64) *big.Int {
	// expense = floor(d * 10_000 / (10_000 + bips))

	return bi.DivRaw(
		bi.MulRaw(deposit, 10_000),
		10_000+bips,
	)
}
