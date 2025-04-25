package app

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
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

	// use cached pricer avoid spamming coingecko public api
	const priceCacheEvictInterval = time.Minute * 10
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
