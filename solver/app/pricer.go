package app

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokenpricer/coingecko"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/solver/types"
)

func newPricer(ctx context.Context, network netconf.ID, apiKey string) tokenpricer.Pricer {
	if network == netconf.Devnet {
		return tokenpricer.NewDevnetMock()
	}

	pricer := tokenpricer.NewCached(coingecko.New(coingecko.WithAPIKey(apiKey)))

	const priceCacheEvictInterval = time.Minute
	go pricer.ClearCacheForever(ctx, priceCacheEvictInterval)

	return pricer
}

// priceFunc returns the unit price of the `deposit` denominated in `expense`.
// That is, how many units of `quote` one unit of `base` is worth.
//
// E.g.: if deposit = ETH, expense = USDC, and priceFunc returns 3200 USDC/ETH, then 1 ETH = 3200 USDC.
//
// Usage:
//
//	expenseAmount = depositAmount * priceFunc(deposit, expense)
//	depositAmount = expenseAmount / priceFunc(deposit, expense)
type priceFunc func(ctx context.Context, deposit, expense tokens.Token) (types.Price, error)

// newPriceFunc returns a priceFunc that uses the provided tokenpricer.Pricer to get the price.
func newPriceFunc(pricer tokenpricer.Pricer) priceFunc {
	return func(ctx context.Context, deposit, expense tokens.Token) (types.Price, error) {
		if deposit.ChainClass != expense.ChainClass {
			// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
			return types.Price{}, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)", "deposit", deposit.ChainClass, "expense", expense.ChainClass)
		}

		price := big.NewRat(1, 1)
		if !areEqualBySymbol(deposit, expense) {
			var err error
			price, err = pricer.Price(ctx, deposit.Asset, expense.Asset)
			if err != nil {
				return types.Price{}, err
			}
		}

		return types.Price{
			Price:   price,
			Deposit: deposit.Asset,
			Expense: expense.Asset,
		}, nil
	}
}

type priceHandlerFunc func(ctx context.Context, request types.PriceRequest) (types.Price, error)

func wrapPriceHandlerFunc(priceFunc priceFunc) priceHandlerFunc {
	return func(ctx context.Context, req types.PriceRequest) (types.Price, error) {
		srcToken, ok := tokens.ByAddress(req.SourceChainID, req.DepositToken)
		if !ok {
			return types.Price{}, errors.New("deposit token not found", "token", req.DepositToken, "src_chain", evmchain.Name(req.SourceChainID))
		}
		dstToken, ok := tokens.ByAddress(req.DestinationChainID, req.ExpenseToken)
		if !ok {
			return types.Price{}, errors.New("expense token not found", "token", req.ExpenseToken, "dst_chain", evmchain.Name(req.DestinationChainID))
		}

		price, err := priceFunc(ctx, srcToken, dstToken)
		if err != nil {
			return types.Price{}, errors.Wrap(err, "price")
		}

		return price.WithFeeBips(feeBips(srcToken.Asset, dstToken.Asset)), nil // Add fee to the price.
	}
}

// debugOrderPrice log order amounts and prices if applicable.
func debugOrderPrice(ctx context.Context, priceFunc priceFunc, order Order) {
	pendingData, err := order.PendingData()
	if err != nil {
		return
	}

	if len(pendingData.MinReceived) == 0 || len(pendingData.MaxSpent) == 0 {
		return
	}

	depositAmt := pendingData.MinReceived[0].Amount
	depositTkn, ok := tokenByAddr32(order.SourceChainID, pendingData.MinReceived[0].Token)
	if !ok {
		return
	}

	expenseAmt := pendingData.MaxSpent[0].Amount
	expenseTkn, ok := tokenByAddr32(pendingData.DestinationChainID, pendingData.MaxSpent[0].Token)
	if !ok {
		return
	}

	orderPrice := types.Price{
		Price:   new(big.Rat).SetFrac(expenseAmt, depositAmt),
		Deposit: depositTkn.Asset,
		Expense: expenseTkn.Asset,
	}

	currentPrice, err := priceFunc(ctx, depositTkn, expenseTkn)
	if err != nil {
		return
	}

	currentF64, _ := currentPrice.Price.Float64()
	orderF64, _ := orderPrice.Price.Float64()
	profitBips := math.Round((currentF64 - orderF64) / orderF64 * 10_000)
	slippageBips := int64(profitBips) - feeBips(depositTkn.Asset, expenseTkn.Asset)

	log.Debug(ctx, "Pending order amounts",
		"deposit", depositTkn.FormatAmt(depositAmt),
		"expense", expenseTkn.FormatAmt(expenseAmt),
		"slippage_bips", slippageBips,
		"price_current", currentPrice.FormatF64(),
		"price_pair", fmt.Sprintf("%s/%s", orderPrice.Expense, orderPrice.Deposit),
	)
}

func tokenByAddr32(chainID uint64, addr32 [32]byte) (tokens.Token, bool) {
	addr, err := toEthAddr(addr32)
	if err != nil {
		return tokens.Token{}, false
	}

	return tokens.ByAddress(chainID, addr)
}

// monitorPricesForever blocks and instruments all supported asset prices (in USD) periodically.
func monitorPricesForever(ctx context.Context, priceFunc priceFunc) {
	timer := time.NewTimer(0) // Start immediately
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			timer.Reset(time.Minute)

			for asset := range supportedAssets {
				err := monitorPriceOnce(ctx, asset, tokens.USDC, priceFunc)
				if err != nil {
					log.Warn(ctx, "Failed monitoring price (will retry)", err, "asset", asset)
					break
				}
			}
		}
	}
}

func monitorPriceOnce(ctx context.Context, deposit, expense tokens.Asset, priceFunc priceFunc) error {
	if deposit == expense {
		return nil
	}

	expenseTkn, ok := tokens.ByAsset(evmchain.IDEthereum, expense)
	if !ok {
		return nil
	}

	depositTkn, ok := tokens.ByAsset(evmchain.IDEthereum, deposit)
	if !ok {
		return nil
	}

	price, err := priceFunc(ctx, depositTkn, expenseTkn)
	if err != nil {
		return err
	}

	priceGauge.WithLabelValues(price.FormatPair()).Set(price.F64())

	return nil
}
