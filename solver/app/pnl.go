package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/pnl"
	tokenslib "github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"
)

// pnlExpenses logs the solver expense PnL for the order.
func pnlExpenses(ctx context.Context, pricer tokenslib.Pricer, order Order, outboxAddr common.Address, dstChainName string) error {
	pendingData, err := order.PendingData()
	if err != nil {
		return errors.Wrap(err, "get pending data [BUG]")
	}

	maxSpent, err := parseMaxSpent(pendingData, outboxAddr)
	if err != nil {
		return errors.Wrap(err, "parse max spent [BUG]") // This should never fail here.
	}

	for _, tknAmt := range maxSpent {
		p := pnl.LogP{
			Type:        pnl.Expense,
			AmountGwei:  toGweiF64(tknAmt.Amount),
			Currency:    pnl.Currency(tknAmt.Token.Symbol),
			Category:    "solver",
			Subcategory: "spend",
			Chain:       dstChainName,
			ID:          order.ID.String(),
		}
		pnl.Log(ctx, p)
		usdPnL(ctx, pricer, tknAmt.Token.Token, p)
	}

	return nil
}

// pnlIncome logs the solver income PnL for the order.
func pnlIncome(ctx context.Context, pricer tokenslib.Pricer, order Order, srcChainName string) error {
	minReceived, err := parseMinReceived(order)
	if err != nil {
		return errors.Wrap(err, "parse min received [BUG]") // This should never fail here.
	}

	for _, tknAmt := range minReceived {
		p := pnl.LogP{
			Type:        pnl.Income,
			AmountGwei:  toGweiF64(tknAmt.Amount),
			Currency:    pnl.Currency(tknAmt.Token.Symbol),
			Category:    "solver",
			Subcategory: "receive",
			Chain:       srcChainName,
			ID:          order.ID.String(),
		}

		pnl.Log(ctx, p)
		usdPnL(ctx, pricer, tknAmt.Token.Token, p)
	}

	return nil
}

// usdPnL logs the USD equivalent PnL.
// This is best effort.
func usdPnL(ctx context.Context, pricer tokenslib.Pricer, token tokenslib.Token, p pnl.LogP) {
	usdPrice, err := pricer.Price(ctx, token)
	if err != nil {
		log.Warn(ctx, "Failed to get token USD price (will retry)", err, "token", token.Name)
		return
	}

	p.Currency = pnl.USD
	p.AmountGwei *= usdPrice // USDAmount = TokenAmount * USDPricePerToken
	pnl.Log(ctx, p)
}
