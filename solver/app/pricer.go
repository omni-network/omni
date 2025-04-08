package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"
)

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
