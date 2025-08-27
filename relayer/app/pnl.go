package relayer

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/pnl"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokenpricer/coingecko"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/xchain"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	priceCacheEvictInterval = time.Minute
)

// newTokenPricer creates a new cached pricer with priceCacheEvictInterval.
func newTokenPricer(ctx context.Context, cgAPIKey string) tokenpricer.Pricer {
	pricer := tokenpricer.NewCached(coingecko.New(coingecko.WithAPIKey(cgAPIKey)))

	// use cached pricer avoid spamming coingecko public api
	go pricer.ClearCacheForever(ctx, priceCacheEvictInterval)

	return pricer
}

type pnlLogger struct {
	network netconf.ID
	pricer  tokenpricer.Pricer
}

// newPnlLogger creates a new pnl logger.
func newPnlLogger(network netconf.ID, pricer tokenpricer.Pricer) pnlLogger {
	return pnlLogger{network: network, pricer: pricer}
}

// log logs the pnl for an xsubmit transaction, warning on error.
func (l pnlLogger) log(ctx context.Context, tx *ethtypes.Transaction, receipt *ethclient.Receipt, sub xchain.Submission) {
	if err := l.logE(ctx, tx, receipt, sub); err != nil {
		log.Warn(ctx, "Failed to log pnl", err)
	}
}

// logE logs the pnl for an xsubmit transaction, returning any errors.
func (l pnlLogger) logE(ctx context.Context, tx *ethtypes.Transaction, receipt *ethclient.Receipt, sub xchain.Submission) error {
	srcChainID := sub.BlockHeader.ChainID
	dstChainID := sub.DestChainID

	dest, ok := evmchain.MetadataByID(dstChainID)
	if !ok {
		return errors.New("unknown chain ID")
	}

	spendGwei := totalSpendGwei(tx, receipt)
	spendTotal.WithLabelValues(dest.Name, dest.NativeToken.Symbol).Add(spendGwei)

	prices, err := l.pricer.USDPrices(ctx, tokens.NOM, tokens.ETH)
	if err != nil {
		return errors.Wrap(err, "get prices")
	}

	log.Debug(ctx, "Using token prices", "nom", prices[tokens.NOM], "eth", prices[tokens.ETH])

	spend, err := spendByDenom(dest, spendGwei, prices)
	if err != nil {
		return errors.Wrap(err, "get spend")
	}

	md := map[string]any{
		"tx":       tx.Hash().Hex(),
		"gas_used": receipt.GasUsed,
		"status":   receipt.Status,
		"num_msgs": len(sub.Msgs),
	}

	id := tx.Hash().Hex()

	// log expenses
	pnl.Log(ctx,
		pnl.LogP{
			Type: pnl.Expense, AmountGwei: spend.nUSD, Currency: pnl.USD,
			Category: "gas", Subcategory: "xsubmit",
			Chain: dest.Name, ID: id, Metadata: md,
		},
		pnl.LogP{
			Type: pnl.Expense, AmountGwei: spend.nETH, Currency: pnl.ETH,
			Category: "gas", Subcategory: "xsubmit",
			Chain: dest.Name, ID: id, Metadata: md,
		},
		pnl.LogP{
			Type: pnl.Expense, AmountGwei: spend.nNOM, Currency: pnl.NOM,
			Category: "gas", Subcategory: "xsubmit",
			Chain: dest.Name, ID: id, Metadata: md,
		},
	)

	// do not log income if:
	//	 - the submission failed (avoid double counting)
	// 	 - source chain is omni consensus chain (no fees collected)
	if netconf.IsOmniConsensus(l.network, srcChainID) || receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return nil
	}

	src, ok := evmchain.MetadataByID(srcChainID)
	if !ok {
		return errors.New("unknown source chain ID")
	}

	fees, err := feeByDenom(src, sub, prices)
	if err != nil {
		return errors.Wrap(err, "get fees")
	}

	// log income
	pnl.Log(ctx,
		pnl.LogP{
			Type: pnl.Income, AmountGwei: fees.nUSD, Currency: pnl.USD,
			Category: "fees", Subcategory: "xcall",
			Chain: src.Name, ID: id, Metadata: md,
		},
		pnl.LogP{
			Type: pnl.Income, AmountGwei: fees.nETH, Currency: pnl.ETH,
			Category: "fees", Subcategory: "xcall",
			Chain: src.Name, ID: id, Metadata: md,
		},
		pnl.LogP{
			Type: pnl.Income, AmountGwei: fees.nNOM, Currency: pnl.NOM,
			Category: "fees", Subcategory: "xcall",
			Chain: src.Name, ID: id, Metadata: md,
		},
	)

	return nil
}

type amtByDenom struct {
	nUSD float64 // "nano" USD (gwei)
	nNOM float64 // "nano" NOM (gwei)
	nETH float64 // "nano" ETH (gwei)
}

// feeByDenom returns the amount fees collected from a receipt in nom, eth, and usd.
func feeByDenom(
	src evmchain.Metadata,
	sub xchain.Submission,
	prices map[tokens.Asset]float64,
) (amtByDenom, error) {
	var fees amtByDenom

	for _, msg := range sub.Msgs {
		if msg.SourceChainID != src.ChainID {
			return amtByDenom{}, errors.New("source chain ID mismatch [BUG]", "expected", src.ChainID, "got", msg.SourceChainID)
		}

		feesGwei := bi.ToGweiF64(msg.Fees)

		switch src.NativeToken {
		case tokens.NOM:
			fees.nNOM += feesGwei
			fees.nUSD += feesGwei * prices[tokens.NOM]
		case tokens.ETH:
			fees.nETH += feesGwei
			fees.nUSD += feesGwei * prices[tokens.ETH]
		default:
			return amtByDenom{}, errors.New("unknown native token", "token", src.NativeToken)
		}
	}

	return fees, nil
}

// spendByDenom returns the amount spent on a transaction in nom, eth, and usd.
func spendByDenom(
	dest evmchain.Metadata,
	spendGwei float64,
	prices map[tokens.Asset]float64,
) (amtByDenom, error) {
	var spend amtByDenom

	switch dest.NativeToken {
	case tokens.NOM:
		spend.nNOM = spendGwei
		spend.nUSD = spendGwei * prices[tokens.NOM]
	case tokens.ETH:
		spend.nETH = spendGwei
		spend.nUSD = spendGwei * prices[tokens.ETH]
	default:
		return amtByDenom{}, errors.New("unknown native token", "token", dest.NativeToken)
	}

	return spend, nil
}

// totalSpendGwei returns the total amount spent on a transaction in gwei.
func totalSpendGwei(tx *ethtypes.Transaction, rec *ethclient.Receipt) float64 {
	spend := bi.MulRaw(rec.EffectiveGasPrice, rec.GasUsed)

	// add op l1 fee, if any
	if rec.OPL1Fee != nil {
		spend = bi.Add(spend, rec.OPL1Fee)
	}

	// add tx value
	spend = bi.Add(spend, tx.Value())

	return bi.ToGweiF64(spend)
}
