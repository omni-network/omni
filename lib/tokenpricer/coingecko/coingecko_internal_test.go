package coingecko

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/tokens"
)

// GetAPIKeyHeader returns the header key for the CoinGecko API key for testing purposes.
func GetAPIKeyHeader() string {
	return apikeyHeader
}

// GetPrice exports the getPrice method for testing purposes.
func (c Client) GetPrice(ctx context.Context, currency string, bases ...tokens.Asset) (map[tokens.Asset]*big.Rat, error) {
	return c.getPrice(ctx, currency, bases...)
}
