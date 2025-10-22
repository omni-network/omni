package fundthresh

import (
	"math"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"
)

var inf = math.Inf(1)

type FundThreshold struct {
	token   tokens.Token
	min     float64 // alert if below
	target  float64 // fund target, below which, consider deficit
	surplus float64 // above, consider surplus
	minSwap float64 // min amount needed to swap
	maxSwap float64 // max amount allowed in single swap
}

// Min returns the minimum balance, below which we should alert.
func (t FundThreshold) Min() *big.Int {
	return t.balance(t.min)
}

// Target returns the target balance, below which we should fund.
func (t FundThreshold) Target() *big.Int {
	return t.balance(t.target)
}

// Surplus returns the surplus treshold, if any, above which we can swap/send elsewhere.
// If surplus threshold set to inf, it returns the max uint256.
func (t FundThreshold) Surplus() *big.Int {
	if t.surplus == inf {
		return bi.Clone(umath.MaxUint256)
	}

	return t.balance(t.surplus)
}

// NeverSurplus returns true if the surplus threshold is set to inf.
func (t FundThreshold) NeverSurplus() bool {
	return t.surplus == inf
}

// MinSwap returns the minimum amount needed to swap.
func (t FundThreshold) MinSwap() *big.Int {
	if t.minSwap == 0 {
		return bi.Zero()
	}

	return t.balance(t.minSwap)
}

// MaxSwap returns the maximum amount allowed in single swap.
func (t FundThreshold) MaxSwap() *big.Int {
	if t.maxSwap == 0 {
		return bi.Zero()
	}

	return t.balance(t.maxSwap)
}

// balance returns the float balance as a big.Int, normalized to the token's decimals.
func (t FundThreshold) balance(f float64) *big.Int {
	if t.token.Decimals == 6 {
		return bi.Dec6(f)
	}

	return bi.Ether(f)
}

// Get returns the fund thesholds for `token`.
func Get(token tokens.Token) FundThreshold {
	t, ok := thresholds[token]
	if !ok {
		// If threshold not explicitly set, return 0 target w/ inf surplus.
		// So that the token is never considered in deficit / surplus.
		return FundThreshold{
			token:   token,
			surplus: inf,
		}
	}

	surplus := t.surplus
	if surplus == 0 && t.target > 0 {
		// Surplus threshold should always be >= target.
		surplus = t.target
	}

	return FundThreshold{
		token:   token,
		min:     t.min,
		target:  t.target,
		surplus: surplus,
		maxSwap: t.maxSwap,
		minSwap: t.minSwap,
	}
}

var (
	thresholds = map[tokens.Token]FundThreshold{
		// ETH
		mustToken(evmchain.IDEthereum, tokens.ETH): {
			min:     1,
			target:  3,
			surplus: 5,
			minSwap: 1,
			maxSwap: 5,
		},
		mustToken(evmchain.IDBase, tokens.ETH): {
			min:     1,
			target:  3,
			surplus: 5,
			minSwap: 1,
			maxSwap: 5,
		},
		mustToken(evmchain.IDArbitrumOne, tokens.ETH): {
			min:     1,
			target:  3,
			surplus: 5,
			minSwap: 1,
			maxSwap: 5,
		},
		mustToken(evmchain.IDOptimism, tokens.ETH): {
			min:     1,
			target:  3,
			surplus: 5,
			minSwap: 1,
			maxSwap: 5,
		},

		// USDC
		mustToken(evmchain.IDEthereum, tokens.USDC): {
			min:     1000,
			target:  5000,
			surplus: 10000,
			minSwap: 1000,
			maxSwap: 10000,
		},
		mustToken(evmchain.IDBase, tokens.USDC): {
			min:     1000,
			target:  5000,
			surplus: 10000,
			minSwap: 1000,
			maxSwap: 10000,
		},
		mustToken(evmchain.IDArbitrumOne, tokens.USDC): {
			min:     1000,
			target:  5000,
			surplus: 10000,
			minSwap: 1000,
			maxSwap: 10000,
		},
		mustToken(evmchain.IDOptimism, tokens.USDC): {
			min:     1000,
			target:  5000,
			surplus: 10000,
			minSwap: 1000,
			maxSwap: 10000,
		},
		mustToken(evmchain.IDMantle, tokens.USDC): {
			target: 2_000,
		},

		// USDT
		mustToken(evmchain.IDEthereum, tokens.USDT): {
			min:     1000,
			target:  5000,
			surplus: 10000,
			minSwap: 1000,
			maxSwap: 10000,
		},
		mustToken(evmchain.IDOptimism, tokens.USDT): {
			min:     1000,
			target:  5000,
			surplus: 10000,
			minSwap: 1000,
			maxSwap: 10000,
		},
		mustToken(evmchain.IDArbitrumOne, tokens.USDT): {
			min:     1000,
			target:  5000,
			surplus: 10000,
			minSwap: 1000,
			maxSwap: 10000,
		},

		// WSTETH
		mustToken(evmchain.IDEthereum, tokens.WSTETH): {
			min:     1,
			target:  3,
			surplus: 5,
			minSwap: 1,
			maxSwap: 5,
		},
		mustToken(evmchain.IDBase, tokens.WSTETH): {
			minSwap: 1,
			maxSwap: 5,
		},

		// MNT
		mustToken(evmchain.IDMantle, tokens.MNT): {
			min:    1000, // Alert on 15k (rebalance is manual)
			target: 3000,
		},

		// mETH
		mustToken(evmchain.IDMantle, tokens.METH): {
			min:    1, // Alert on 1 mETH (rebalance is manual)
			target: 3,
		},

		// WETH
		mustToken(evmchain.IDMantle, tokens.WETH): {
			min:    1, // Alert on 1 WETH (rebalance is manual)
			target: 3,
		},
	}
)

func mustToken(chainID uint64, asset tokens.Asset) tokens.Token {
	tkn, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic("token not found")
	}

	return tkn
}
