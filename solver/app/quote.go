package app

import (
	"context"
	"math/big"
	"net/http"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/solver/types"
)

// standardFeeBips is the standard fee charged by the solver (0.3%).
const standardFeeBips = 30

type quoteFunc func(context.Context, types.QuoteRequest) (types.QuoteResponse, error)

// quoter is quoteFunc that can be used to quoter an expense or deposit.
// It is the logic behind the /quoter endpoint.
func newQuoter(priceFunc priceFunc) quoteFunc {
	return func(ctx context.Context, req types.QuoteRequest) (types.QuoteResponse, error) {
		deposit := req.Deposit
		expense := req.Expense

		isDepositQuote := deposit.Amount == nil || bi.IsZero(deposit.Amount)
		isExpenseQuote := expense.Amount == nil || bi.IsZero(expense.Amount)

		returnErr := func(code int, msg string) (types.QuoteResponse, error) {
			return types.QuoteResponse{}, newAPIError(errors.New(msg), code)
		}

		if isDepositQuote == isExpenseQuote {
			return returnErr(http.StatusBadRequest, "deposit and expense amount cannot be both zero or both non-zero")
		}

		depositTkn, ok := tokens.ByAddress(req.SourceChainID, req.Deposit.Token)
		if !ok {
			return returnErr(http.StatusNotFound, "unsupported deposit token")
		}

		expenseTkn, ok := tokens.ByAddress(req.DestinationChainID, req.Expense.Token)
		if !ok {
			return returnErr(http.StatusNotFound, "unsupported expense token")
		}

		maybeMinMaxReject := func(expenseAmt *big.Int) error {
			bounds := GetSpendBounds(expenseTkn)
			overMax := bounds.MaxSpend != nil && bi.GT(expenseAmt, bounds.MaxSpend)
			underMin := bounds.MinSpend != nil && bi.LT(expenseAmt, bounds.MinSpend)

			if overMax {
				return newRejection(types.RejectExpenseOverMax,
					errors.New("requested expense exceeds maximum",
						"ask", expenseTkn.FormatAmt(expenseAmt),
						"max", expenseTkn.FormatAmt(bounds.MaxSpend),
					))
			}

			if underMin {
				return newRejection(types.RejectExpenseUnderMin,
					errors.New("requested expense is below minimum",
						"ask", expenseTkn.FormatAmt(expenseAmt),
						"min", expenseTkn.FormatAmt(bounds.MinSpend),
					))
			}

			return nil
		}

		returnQuote := func(depositAmt, expenseAmt *big.Int) (types.QuoteResponse, error) {
			return types.QuoteResponse{
				Deposit: types.AddrAmt{
					Token:  deposit.Token,
					Amount: depositAmt,
				},
				Expense: types.AddrAmt{
					Token:  expense.Token,
					Amount: expenseAmt,
				},
			}, maybeMinMaxReject(expenseAmt)
		}

		if isDepositQuote {
			quoted, err := quoteDeposit(ctx, priceFunc, depositTkn, TokenAmt{Token: expenseTkn, Amount: expense.Amount})
			if err != nil {
				return types.QuoteResponse{}, newAPIError(err, http.StatusBadRequest)
			}

			return returnQuote(quoted.Amount, expense.Amount)
		}

		quoted, err := quoteExpense(ctx, priceFunc, expenseTkn, TokenAmt{Token: depositTkn, Amount: deposit.Amount})
		if err != nil {
			return types.QuoteResponse{}, newAPIError(err, http.StatusBadRequest)
		}

		return returnQuote(deposit.Amount, quoted.Amount)
	}
}

// getQuote returns payment in `depositTkns` required to pay for `expenses`.
func getQuote(ctx context.Context, priceFunc priceFunc, depositTkns []tokens.Token, expenses []TokenAmt) ([]TokenAmt, error) {
	if len(depositTkns) != 1 {
		return nil, newRejection(types.RejectInvalidDeposit, errors.New("only single deposit token supported"))
	}

	if len(expenses) != 1 {
		return nil, newRejection(types.RejectInvalidExpense, errors.New("only single expense supported"))
	}

	expense := expenses[0]
	depositTkn := depositTkns[0]

	deposit, err := quoteDeposit(ctx, priceFunc, depositTkn, expense)
	if err != nil {
		return nil, err
	}

	return []TokenAmt{deposit}, nil
}

// quoteDeposit returns the source chain deposit required to cover `expense`.
func quoteDeposit(ctx context.Context, priceFunc priceFunc, depositTkn tokens.Token, expense TokenAmt) (TokenAmt, error) {
	price, err := priceFunc(ctx, expense.Token, depositTkn)
	if err != nil {
		return TokenAmt{}, newRejection(types.RejectInvalidDeposit, errors.Wrap(err, "", "expense", expense, "deposit", depositTkn))
	}

	depositAmount := bi.MulF64(expense.Amount, price)
	depositAmount = depositFor(depositAmount, feeBips(depositTkn.Asset, expense.Token.Asset))

	return TokenAmt{
		Token:  depositTkn,
		Amount: depositAmount,
	}, nil
}

// QuoteExpense returns the destination chain expense allowed for `deposit`.
func quoteExpense(ctx context.Context, priceFunc priceFunc, expenseTkn tokens.Token, deposit TokenAmt) (TokenAmt, error) {
	price, err := priceFunc(ctx, deposit.Token, expenseTkn)
	if err != nil {
		return TokenAmt{}, newRejection(types.RejectInvalidDeposit, errors.Wrap(err, "", "deposit", deposit, "expense", expenseTkn))
	}

	expenseAmount := bi.MulF64(deposit.Amount, price)
	expenseAmount = expenseFor(expenseAmount, feeBips(expenseTkn.Asset, deposit.Token.Asset))

	return TokenAmt{
		Token:  expenseTkn,
		Amount: expenseAmount,
	}, nil
}

func areEqualBySymbol(a, b tokens.Token) bool {
	if a.Symbol == b.Symbol {
		return true
	}

	equivalents := map[string]string{}
	makeEq := func(a, b string) {
		equivalents[a] = b
		equivalents[b] = a
	}

	// consider stETH and ETH as equivalent
	makeEq(tokens.STETH.Symbol, tokens.ETH.Symbol)

	return equivalents[a.Symbol] == b.Symbol
}

// feeBips returns the fee in bips for a given pair.
func feeBips(a, b tokens.Asset) int64 {
	// if OMNI<>OMNI, charge no fee
	if a == tokens.OMNI && b == tokens.OMNI {
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
