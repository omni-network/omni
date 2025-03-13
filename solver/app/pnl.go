package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/pnl"
	tokenslib "github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"
)

type pnlFunc func(ctx context.Context, order Order) error
type targetFunc func(PendingData) string

// newPnlFunc returns a pnlFunc that logs the PnL for an order.
// This is done immediately after successful fill.
//
// Technically, income should only be logged after successful claim, but
// target is used as subcategory and this is only available in pending data.
//
// So technically, the claim can fail and the PnL income will still be logged.
func newPnlFunc(
	pricer tokenslib.Pricer,
	targetName targetFunc,
	namer func(uint64) string,
	outbox common.Address,
) pnlFunc {
	return func(ctx context.Context, order Order) error {
		pendingData, err := order.PendingData()
		if err != nil {
			return errors.Wrap(err, "get pending data [BUG]")
		}

		target := targetName(pendingData)
		srcChainName := namer(order.SourceChainID)
		dstChainName := namer(pendingData.DestinationChainID)

		maxSpent, err := parseMaxSpent(pendingData, outbox)
		if err != nil {
			return errors.Wrap(err, "parse max spent [BUG]") // This should never fail here.
		}

		minReceived, err := parseMinReceived(order)
		if err != nil {
			return errors.Wrap(err, "parse min received [BUG]") // This should never fail here.
		}

		const pnlCategory = "solver"

		for _, tknAmt := range maxSpent {
			p := pnl.LogP{
				Type:        pnl.Expense,
				AmountGwei:  toGweiF64(tknAmt.Amount),
				Currency:    pnl.Currency(tknAmt.Token.Symbol),
				Category:    pnlCategory,
				Subcategory: target,
				Chain:       dstChainName,
				ID:          order.ID.String(),
			}
			pnl.Log(ctx, p)
			usdPnL(ctx, pricer, tknAmt.Token.Token, p)
		}

		for _, tknAmt := range minReceived {
			p := pnl.LogP{
				Type:        pnl.Income,
				AmountGwei:  toGweiF64(tknAmt.Amount),
				Currency:    pnl.Currency(tknAmt.Token.Symbol),
				Category:    pnlCategory,
				Subcategory: target,
				Chain:       srcChainName,
				ID:          order.ID.String(),
			}

			pnl.Log(ctx, p)
			usdPnL(ctx, pricer, tknAmt.Token.Token, p)
		}

		return nil
	}
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
