package app

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/pnl"
	tokenslib "github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type orderPnLFunc func(ctx context.Context, order Order, rec *types.Receipt) error
type simpleGasPnLFunc func(ctx context.Context, chainID uint64, rec *types.Receipt, subCat string) error

type targetFunc func(PendingData) string

// newFilledPnlFunc returns a orderPnLFunc that logs the PnL for successfully filled orders.
//
// Technically, income should only be logged after successful claim, but
// target is used as subcategory and this is only available in pending data.
//
// So technically, the claim can fail and the PnL income will still be logged.
func newFilledPnlFunc(
	pricer tokenslib.Pricer,
	targetName targetFunc,
	namer func(uint64) string,
	outbox common.Address,
	destFilledAge destFilledAge,
) orderPnLFunc {
	return func(ctx context.Context, order Order, rec *types.Receipt) error {
		pendingData, err := order.PendingData()
		if err != nil {
			return errors.Wrap(err, "get pending data [BUG]")
		}

		target := targetName(pendingData)
		srcChainName := namer(order.SourceChainID)
		dstChainName := namer(pendingData.DestinationChainID)
		age := destFilledAge(ctx, pendingData.DestinationChainID, rec.BlockNumber.Uint64(), order)

		maxSpent, err := parseMaxSpent(pendingData, outbox)
		if err != nil {
			return errors.Wrap(err, "parse max spent [BUG]") // This should never fail here.
		}

		minReceived, err := parseMinReceived(order)
		if err != nil {
			return errors.Wrap(err, "parse min received [BUG]") // This should never fail here.
		}

		log.Info(ctx, "Order filled", age, "src_chain", srcChainName, "dst_chain", dstChainName, "target", target)
		filledOrders.WithLabelValues(srcChainName, dstChainName, target).Inc()

		if err := gasPnL(ctx, pricer, pendingData.DestinationChainID, dstChainName, rec, target, order.ID.String()); err != nil {
			return err
		}

		// Log expenses and deposits
		for _, tknAmt := range maxSpent {
			p := pnl.LogP{
				Type:        pnl.Expense,
				AmountGwei:  umath.WeiToGweiF64(tknAmt.Amount),
				Currency:    pnl.Currency(tknAmt.Token.Symbol),
				Category:    "solver_expense",
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
				AmountGwei:  umath.WeiToGweiF64(tknAmt.Amount),
				Currency:    pnl.Currency(tknAmt.Token.Symbol),
				Category:    "solver_deposit",
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

// newOrderGasPnLFunc returns a orderPnLFunc that logs the gas expense PnL for updating order status, except for filled orders.
func newOrderGasPnLFunc(pricer tokenslib.Pricer, namer func(uint64) string) orderPnLFunc {
	return func(ctx context.Context, order Order, rec *types.Receipt) error {
		srcChainName := namer(order.SourceChainID)
		return gasPnL(ctx, pricer, order.SourceChainID, srcChainName, rec, order.Status.String(), order.ID.String())
	}
}

// newSimpleGasPnLFunc returns a simpleGasPnLFunc that logs simple gas PnL, not related to orders.
func newSimpleGasPnLFunc(pricer tokenslib.Pricer, namer func(uint64) string) simpleGasPnLFunc {
	return func(ctx context.Context, chainID uint64, rec *types.Receipt, subCat string) error {
		chainName := namer(chainID)
		return gasPnL(ctx, pricer, chainID, chainName, rec, subCat, rec.TxHash.Hex())
	}
}

func gasPnL(
	ctx context.Context,
	pricer tokenslib.Pricer,
	chainID uint64,
	chainName string,
	rec *types.Receipt,
	subCat string,
	id string,
) error {
	amount := new(big.Int).Mul(rec.EffectiveGasPrice, umath.NewBigInt(rec.GasUsed))

	// Add any xcall fees included in tx
	if fee, ok := maybeParseXCallFee(rec); ok {
		amount = new(big.Int).Add(amount, fee)
	}

	nativeToken, ok := tokens.Find(chainID, NativeAddr)
	if !ok {
		return errors.New("native token not found [BUG]")
	}

	// Log native gas as expense
	p := pnl.LogP{
		Type:        pnl.Expense,
		AmountGwei:  umath.WeiToGweiF64(amount),
		Currency:    pnl.Currency(nativeToken.Symbol),
		Category:    "gas",
		Subcategory: subCat,
		Chain:       chainName,
		ID:          id,
	}
	pnl.Log(ctx, p)
	usdPnL(ctx, pricer, nativeToken.Token, p)

	return nil
}

// maybeParseXCallFee returns the xcall fee from the receipt if present or false.
func maybeParseXCallFee(rec *types.Receipt) (*big.Int, bool) {
	portal, _ := bindings.NewOmniPortal(common.Address{}, nil) // Safe to pass in zeros since we only parse events.
	for _, event := range rec.Logs {
		if event == nil {
			continue
		}
		xmsg, err := portal.OmniPortalFilterer.ParseXMsg(*event)
		if err != nil {
			continue
		}

		return xmsg.Fees, true
	}

	return nil, false
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
