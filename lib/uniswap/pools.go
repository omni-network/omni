//nolint:unused // skeleton code
package uniswap

import (
	"github.com/omni-network/omni/lib/tokens"
)

const (
	// FeeBips30 is the 0.3% fee tier.
	FeeBips30 uint = 3000

	// FeeBips5 is the 0.05% fee tier.
	FeeBips5 uint = 500

	// FeeBips1 is the 0.01% fee tier.
	FeeBips1 uint = 100
)

// PoolMeta defines an asset pair and fee bips for a uniswap v3 pool.
// Asset0 and Asset1 are unordered, unlike pool Token0 and Token1 (ordered by address).
type PoolMeta struct {
	Asset0 tokens.Asset
	Asset1 tokens.Asset
	Fee    uint
}

// IsPair checks if the given tokenIn / tokenOut pair is in the pool.
func (p PoolMeta) IsPair(tokenIn, tokenOut tokens.Token) bool {
	return tokenIn.Is(p.Asset0) && tokenOut.Is(p.Asset1) || tokenIn.Is(p.Asset1) && tokenOut.Is(p.Asset0)
}

// pools is a static list of uniswap v3 pool meta data (reference: https://app.uniswap.org/explore/pools)
var pools = []PoolMeta{
	newPool(tokens.WETH, tokens.USDC, FeeBips5),   // WETH / UDSC, 0.05%
	newPool(tokens.WSTETH, tokens.WETH, FeeBips1), // WSTETH / WETH, 0.01%
}

func newPool(asset0, asset1 tokens.Asset, fee uint) PoolMeta {
	return PoolMeta{Asset0: asset0, Asset1: asset1, Fee: fee}
}

// poolFee returns the pool fee for the given tokenIn / tokenOut pair.
func poolFee(tokenIn tokens.Token, tokenOut tokens.Token) (uint, bool) {
	for _, pool := range pools {
		if pool.IsPair(tokenIn, tokenOut) {
			return pool.Fee, true
		}
	}

	return 0, false
}
