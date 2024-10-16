package relayer

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/pnl"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/coingecko"
	"github.com/omni-network/omni/lib/xchain"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	priceCacheEvictInterval = time.Minute
)

// newTokenPricer creates a new cached pricer with priceCacheEvictInterval.
func newTokenPricer(ctx context.Context) *tokens.CachedPricer {
	pricer := tokens.NewCachedPricer(coingecko.New())

	// use cached pricer avoid spamming coingecko public api
	// TODO: use api key
	go pricer.ClearCacheForever(ctx, priceCacheEvictInterval)

	return pricer
}

type pnlLogger struct {
	network netconf.ID
	pricer  tokens.Pricer
}

// newPnlLogger creates a new pnl logger.
func newPnlLogger(network netconf.ID, pricer tokens.Pricer) pnlLogger {
	return pnlLogger{network: network, pricer: pricer}
}

// log logs the pnl for an xsubmit transaction, warning on error.
func (l pnlLogger) log(ctx context.Context, tx *ethtypes.Transaction, receipt *ethtypes.Receipt, sub xchain.Submission) {
	if err := l.logE(ctx, tx, receipt, sub); err != nil {
		log.Warn(ctx, "Failed to log pnl", err)
	}
}

// logE logs the pnl for an xsubmit transaction, returning any errors.
func (l pnlLogger) logE(ctx context.Context, tx *ethtypes.Transaction, receipt *ethtypes.Receipt, sub xchain.Submission) error {
	srcChainID := sub.BlockHeader.ChainID
	dstChainID := sub.DestChainID

	dest, ok := evmchain.MetadataByID(dstChainID)
	if !ok {
		return errors.New("unknown chain ID")
	}

	prices, err := l.pricer.Price(ctx, tokens.OMNI, tokens.ETH)
	if err != nil {
		return errors.Wrap(err, "get prices")
	}

	log.Debug(ctx, "Using token prices", "omni", prices[tokens.OMNI], "eth", prices[tokens.ETH])

	spend, err := getSpend(dest, tx, receipt, prices)
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
			Type: pnl.Expense, AmountGwei: spend.nOMNI, Currency: pnl.OMNI,
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

	fees, err := getFees(src, sub, prices)
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
			Type: pnl.Income, AmountGwei: fees.nOMNI, Currency: pnl.OMNI,
			Category: "fees", Subcategory: "xcall",
			Chain: src.Name, ID: id, Metadata: md,
		},
	)

	return nil
}

type amounts struct {
	nUSD  float64 // "nano" USD (gwei)
	nOMNI float64 // "nano" OMNI (gwei)
	nETH  float64 // "nano" ETH (gwei)
}

// getFees returns the amount fees collected from a receipt in omni, eth, and usd.
func getFees(
	src evmchain.Metadata,
	sub xchain.Submission,
	prices map[tokens.Token]float64,
) (amounts, error) {
	fees := amounts{}

	for _, msg := range sub.Msgs {
		if msg.SourceChainID != src.ChainID {
			return amounts{}, errors.New("source chain ID mismatch [BUG]", "expected", src.ChainID, "got", msg.SourceChainID)
		}

		feesGwei := toGwei(msg.Fees)

		switch src.NativeToken {
		case tokens.OMNI:
			fees.nOMNI += feesGwei
			fees.nUSD += feesGwei * prices[tokens.OMNI]
		case tokens.ETH:
			fees.nETH += feesGwei
			fees.nUSD += feesGwei * prices[tokens.ETH]
		default:
			return amounts{}, errors.New("unknown native token", "token", src.NativeToken)
		}
	}

	return fees, nil
}

// getSpend returns the amount spent on a transaction in omni, eth, and usd.
func getSpend(
	dest evmchain.Metadata,
	tx *ethtypes.Transaction,
	receipt *ethtypes.Receipt,
	prices map[tokens.Token]float64,
) (amounts, error) {
	spendGwei := totalSpendGwei(tx, receipt)

	spend := amounts{}
	switch dest.NativeToken {
	case tokens.OMNI:
		spend.nOMNI = spendGwei
		spend.nUSD = spendGwei * prices[tokens.OMNI]
	case tokens.ETH:
		spend.nETH = spendGwei
		spend.nUSD = spendGwei * prices[tokens.ETH]
	default:
		return amounts{}, errors.New("unknown native token", "token", dest.NativeToken)
	}

	return spend, nil
}
