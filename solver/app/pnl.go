package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/pnl"

	"github.com/ethereum/go-ethereum/common"
)

// pnlExpenses logs the solver expense PnL for the order.
func pnlExpenses(ctx context.Context, order Order, outboxAddr common.Address, dstChainName string) error {
	maxSpent, err := parseMaxSpent(order, outboxAddr)
	if err != nil {
		return errors.Wrap(err, "parse max spent [BUG]") // This should never fail here.
	}

	for _, tknAmt := range maxSpent {
		pnl.Log(ctx, pnl.LogP{
			Type:        pnl.Expense,
			AmountGwei:  toGweiF64(tknAmt.Amount),
			Currency:    pnl.Currency(tknAmt.Token.Symbol),
			Category:    "solver",
			Subcategory: "spend",
			Chain:       dstChainName,
			ID:          order.ID.String(),
		})
	}

	return nil
}

// pnlIncome logs the solver income PnL for the order.
func pnlIncome(ctx context.Context, order Order, srcChainName string) error {
	minReceived, err := parseMinReceived(order)
	if err != nil {
		return errors.Wrap(err, "parse min received [BUG]") // This should never fail here.
	}

	for _, tknAmt := range minReceived {
		pnl.Log(ctx, pnl.LogP{
			Type:        pnl.Income,
			AmountGwei:  toGweiF64(tknAmt.Amount),
			Currency:    pnl.Currency(tknAmt.Token.Symbol),
			Category:    "solver",
			Subcategory: "receive",
			Chain:       srcChainName,
			ID:          order.ID.String(),
		})
	}

	return nil
}
