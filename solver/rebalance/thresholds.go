// This file mirrors e2e/app/eoa/solver.go and extends
// The two should be merged in the future, or reconciled in tests.
//

package rebalance

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

// GetFundThreshold returns the fund thesholds for `token`.
func GetFundThreshold(token tokens.Token) FundThreshold {
	t, ok := thresholds[token]
	if !ok {
		// If threshold not explicitly set, return 0 target w/ inf surplus.
		// So that the token is never considered in deficit / surplus.
		return FundThreshold{
			token:   token,
			surplus: inf,
		}
	}

	return FundThreshold{
		token:   token,
		min:     t.min,
		target:  t.target,
		surplus: t.surplus,
		maxSwap: t.maxSwap,
		minSwap: t.minSwap,
	}
}

var (
	thresholds = map[tokens.Token]FundThreshold{
		mustToken(evmchain.IDEthereum, tokens.WSTETH): {
			min:     10,
			target:  50,
			surplus: inf,
		},
		mustToken(evmchain.IDEthereum, tokens.USDC): {
			min:     50_000,
			target:  100_000,
			surplus: 120_000,
			minSwap: 1000,
			maxSwap: 5000,
		},
		mustToken(evmchain.IDBase, tokens.WSTETH): {
			minSwap: 1,
			maxSwap: 3,
		},
		mustToken(evmchain.IDBase, tokens.USDC): {
			min:     20_000,
			target:  40_000,
			surplus: 50_000,
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
