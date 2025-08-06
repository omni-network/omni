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
		returnErr := func(code int, msg string, attrs ...any) (types.QuoteResponse, error) {
			return types.QuoteResponse{}, newAPIError(errors.New(msg, attrs...), code)
		}

		isDepositQuote := bi.IsZero(req.Deposit.Amount)
		isExpenseQuote := bi.IsZero(req.Expense.Amount)

		if isDepositQuote == isExpenseQuote {
			return returnErr(http.StatusBadRequest, "deposit and expense amount cannot be both zero or both non-zero")
		}
		depositTkn, ok := tokens.ByUniAddress(req.SourceChainID, req.Deposit.Token)
		if !ok {
			return returnErr(http.StatusNotFound, "unsupported deposit token", "chain", req.DestinationChainID, "address", req.Expense.Token)
		}

		expenseTkn, ok := tokens.ByUniAddress(req.DestinationChainID, req.Expense.Token)
		if !ok {
			return returnErr(http.StatusNotFound, "unsupported expense token", "chain", req.DestinationChainID, "address", req.Expense.Token)
		}

		// Get the price of the order
		price, err := priceFunc(ctx, depositTkn, expenseTkn)
		if err != nil {
			return types.QuoteResponse{}, newAPIError(err, http.StatusBadRequest)
		}

		// Add solver fee to price
		price = price.WithFeeBips(feeBips(depositTkn.Asset, expenseTkn.Asset))

		if isDepositQuote {
			req.Deposit.Amount = price.ToDeposit(req.Expense.Amount)
		} else {
			req.Expense.Amount = price.ToExpense(req.Deposit.Amount)
		}

		return types.QuoteResponse{
			Deposit: req.Deposit,
			Expense: req.Expense,
		}, maybeMinMaxReject(req.Expense.Amount, expenseTkn)
	}
}

func maybeMinMaxReject(expenseAmt *big.Int, expenseTkn tokens.Token) error {
	bounds, ok := GetSpendBounds(expenseTkn)
	if !ok {
		return nil
	}

	if bi.GT(expenseAmt, bounds.MaxSpend) {
		return newRejection(types.RejectExpenseOverMax,
			errors.New("requested expense exceeds maximum",
				"ask", expenseTkn.FormatAmt(expenseAmt),
				"max", expenseTkn.FormatAmt(bounds.MaxSpend),
			))
	}

	if bi.LT(expenseAmt, bounds.MinSpend) {
		return newRejection(types.RejectExpenseUnderMin,
			errors.New("requested expense is below minimum",
				"ask", expenseTkn.FormatAmt(expenseAmt),
				"min", expenseTkn.FormatAmt(bounds.MinSpend),
			))
	}

	return nil
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
	price, err := priceFunc(ctx, depositTkn, expense.Token)
	if err != nil {
		return TokenAmt{}, newRejection(types.RejectInvalidDeposit, errors.Wrap(err, "", "expense", expense, "deposit", depositTkn))
	}

	// Add fee to price.
	price = price.WithFeeBips(feeBips(depositTkn.Asset, expense.Token.Asset))

	return TokenAmt{
		Token:  depositTkn,
		Amount: price.ToDeposit(expense.Amount),
	}, nil
}

func areEqualBySymbol(a, b tokens.Token) bool {
	return a.Symbol == b.Symbol
}

// feeBips returns the fee in bips for a given pair.
func feeBips(a, b tokens.Asset) int64 {
	// if OMNI<>OMNI, charge no fee
	if a == tokens.OMNI && b == tokens.OMNI {
		return 0
	}

	// if NOM<>NOM, charge no fee
	if a == tokens.NOM && b == tokens.NOM {
		return 0
	}

	// if NOM<>OMNI, charge no fee
	if (a == tokens.NOM && b == tokens.OMNI) || (a == tokens.OMNI && b == tokens.NOM) {
		return 0
	}

	return standardFeeBips
}
