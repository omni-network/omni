package app

import (
	"context"
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

// priceFunc returns the unit price of the `base` denominated in `quote`.
// That is, how many units of `quote` one unit of `base` is worth.
//
// E.g.: if base = ETH, quote = USDC, and priceFunc returns 3200, then 1 ETH = 3200 USDC.
//
// Usage:
//
//	quoteAmount = baseAmount * priceFunc(base, quote)
type priceFunc func(ctx context.Context, base, quote tokens.Token) (float64, error)

// unaryPrice is a priceFunc that returns a price for like-for-like 1-to-1 pairs or an error.
// This is the legacy (pre-swaps) behavior.
func unaryPrice(_ context.Context, base, quote tokens.Token) (float64, error) {
	if !areEqualBySymbol(base, quote) {
		return 0, errors.New("deposit token must match expense token")
	}

	if base.ChainClass != quote.ChainClass {
		// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
		return 0, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)")
	}

	return 1, nil
}

// newPriceFunc returns a priceFunc that uses the provided tokenpricer.Pricer to get the price.
func newPriceFunc(pricer tokenpricer.Pricer) priceFunc {
	return func(ctx context.Context, base, quote tokens.Token) (float64, error) {
		if base.ChainClass != quote.ChainClass {
			// we should reject with UnsupportedDestChain before quoting tokens of different chain classes.
			return 0, errors.New("deposit and expense must be of the same chain class (e.g. mainnet, testnet)")
		}

		if areEqualBySymbol(base, quote) {
			return 1, nil
		}

		return pricer.Price(ctx, base.Asset, quote.Asset)
	}
}

type priceHandlerFunc func(ctx context.Context, request *types.PriceRequest) (*types.PriceResponse, error)

func wrapPriceHandlerFunc(priceFunc priceFunc) priceHandlerFunc {
	return func(ctx context.Context, req *types.PriceRequest) (*types.PriceResponse, error) {
		srcToken, ok := tokens.ByAddress(req.SourceChainID, req.DepositToken)
		if !ok {
			return nil, errors.New("deposit token not found", "token", req.DepositToken, "src_chain", evmchain.Name(req.SourceChainID))
		}
		dstToken, ok := tokens.ByAddress(req.DestinationChainID, req.ExpenseToken)
		if !ok {
			return nil, errors.New("expense token not found", "token", req.ExpenseToken, "dst_chain", evmchain.Name(req.DestinationChainID))
		}

		price, err := priceFunc(ctx, srcToken, dstToken)
		if err != nil {
			return nil, errors.Wrap(err, "price")
		}

		return &types.PriceResponse{
			Price: price,
		}, nil
	}
}
