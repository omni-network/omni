package app

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
)

type checkFunc func(context.Context, types.CheckRequest) error

// newChecker returns a checkFunc that can be used to see if an order would be accepted or rejected.
// It is the logic behind the /check endpoint.
func newChecker(backends ethbackend.Backends, isAllowedCall callAllowFunc, priceFunc priceFunc, solverAddr, outboxAddr common.Address) checkFunc {
	return func(ctx context.Context, req types.CheckRequest) error {
		if req.SourceChainID == req.DestinationChainID {
			return newRejection(types.RejectSameChain, errors.New("source and destination chain are the same"))
		}

		if _, err := backends.Backend(req.SourceChainID); err != nil {
			return newRejection(types.RejectUnsupportedSrcChain, errors.New("unsupported source chain", "chain_id", req.SourceChainID))
		}

		dstBackend, err := backends.Backend(req.DestinationChainID)
		if err != nil {
			return newRejection(types.RejectUnsupportedDestChain, errors.New("unsupported destination chain", "chain_id", req.DestinationChainID))
		}

		deposit, err := parseTokenAmt(req.SourceChainID, req.Deposit)
		if err != nil {
			return err
		}

		expenses, err := parseExpenses(req.DestinationChainID, req.Expenses, req.Calls)
		if err != nil {
			return err
		}

		if err := checkQuote(ctx, priceFunc, []TokenAmt{deposit}, expenses); err != nil {
			return err
		}

		if err := checkLiquidity(ctx, expenses, dstBackend, solverAddr); err != nil {
			return err
		}

		if err := checkCalls(req.DestinationChainID, req.Calls, isAllowedCall); err != nil {
			return err
		}

		if err := checkApprovals(ctx, expenses, dstBackend, solverAddr, outboxAddr); err != nil {
			return err
		}

		fillOriginData, err := getFillOriginData(req)
		if err != nil {
			return err
		}

		// Random orderID (since unfilled).
		var orderID OrderID
		_, _ = rand.Read(orderID[:])

		return checkFill(ctx, dstBackend, orderID, fillOriginData, nativeAmt(expenses), solverAddr, outboxAddr)
	}
}

// getFillOriginData returns packed fill origin data for a check request.
func getFillOriginData(req types.CheckRequest) ([]byte, error) {
	fillOriginData := bindings.SolverNetFillOriginData{
		FillDeadline: req.FillDeadline,
		SrcChainId:   req.SourceChainID,
		DestChainId:  req.DestinationChainID,
		Expenses:     solvernet.FilterNativeExpenses(types.ExpensesToBindings(req.Expenses)),
		Calls:        types.CallsToBindings(req.Calls),
	}

	fillOriginDataBz, err := solvernet.PackFillOriginData(fillOriginData)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill origin data")
	}

	return fillOriginDataBz, nil
}

// coversQuote checks if `deposits` match or exceed a `quote` for expenses.
func coversQuote(deposits, quote []TokenAmt) error {
	byTkn := func(ps []TokenAmt) map[tokens.Token]*big.Int {
		res := make(map[tokens.Token]*big.Int)
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
			return newRejection(types.RejectInsufficientDeposit, errors.New("missing deposit", "token", tkn))
		}

		if bi.LT(d, q) {
			return newRejection(types.RejectInsufficientDeposit, errors.New("insufficient deposit", "deposit", tkn.FormatAmt(d), "min", tkn.FormatAmt(q)))
		}
	}

	return nil
}

func parseExpenses(destChainID uint64, expenses []types.Expense, calls []types.Call) ([]TokenAmt, error) {
	var ps []TokenAmt

	// Sum of call value must be represented in expenses
	callValues := bi.Zero()
	for _, c := range calls {
		if c.Value == nil {
			continue
		}

		callValues.Add(callValues, c.Value)
	}

	nativeExpense := bi.Zero()
	for _, e := range expenses {
		if e.Amount.Sign() <= 0 {
			return nil, newRejection(types.RejectInvalidExpense, errors.New("expense amount not positive"))
		}

		tkn, ok := tokens.ByAddress(destChainID, e.Token)
		if !ok || !IsSupportedToken(tkn) {
			return nil, newRejection(types.RejectUnsupportedExpense, errors.New("unsupported expense token", "addr", e.Token))
		}

		if isNative(e) {
			// sum of call values must be represented by a single expense
			if nativeExpense.Sign() > 0 {
				return nil, newRejection(types.RejectInvalidExpense, errors.New("only one native expense supported"))
			}

			nativeExpense = bi.Add(nativeExpense, e.Amount)
		}

		bounds, ok := GetSpendBounds(tkn)
		if ok && bi.GT(e.Amount, bounds.MaxSpend) {
			return nil, newRejection(types.RejectExpenseOverMax,
				errors.New("requested expense exceeds maximum",
					"ask", tkn.FormatAmt(e.Amount),
					"max", tkn.FormatAmt(bounds.MaxSpend),
				))
		}

		if ok && bi.LT(e.Amount, bounds.MinSpend) {
			return nil, newRejection(types.RejectExpenseUnderMin,
				errors.New("requested expense is below minimum",
					"ask", tkn.FormatAmt(e.Amount),
					"min", tkn.FormatAmt(bounds.MinSpend),
				))
		}

		ps = append(ps, TokenAmt{
			Token:  tkn,
			Amount: e.Amount,
		})
	}

	native, ok := tokens.Native(destChainID)
	if !ok {
		return nil, errors.New("invalid destination chain ID [BUG]") // This shouldn't happen here.
	}

	// native expense must match sum of call values
	if bi.NEQ(nativeExpense, callValues) {
		return nil, newRejection(types.RejectInvalidExpense,
			errors.New("native expense must match native value",
				"expense", native.FormatAmt(nativeExpense),
				"values", native.FormatAmt(callValues),
			))
	}

	return ps, nil
}

func parseTokenAmt(srcChainID uint64, dep types.AddrAmt) (TokenAmt, error) {
	tkn, ok := tokens.ByAddress(srcChainID, dep.Token)
	if !ok {
		return TokenAmt{}, newRejection(types.RejectUnsupportedDeposit, errors.New("unsupported source chain deposit token", "addr", dep.Token, "src_chain", srcChainID))
	}

	return TokenAmt{
		Token:  tkn,
		Amount: dep.Amount,
	}, nil
}

func isNative(e types.Expense) bool { return e.Token == tokens.NativeAddr }
