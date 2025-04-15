package app

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/tokens"
)

var (
	supportedTokens = map[tokens.Asset]bool{
		tokens.ETH:    true,
		tokens.OMNI:   true,
		tokens.WSTETH: true,
		tokens.STETH:  true,
		tokens.USDC:   true,
	}
)

func IsSupportedToken(token tokens.Token) bool {
	return supportedTokens[token.Asset]
}

type SpendBounds struct {
	MinSpend *big.Int // minimum spend amount
	MaxSpend *big.Int // maximum spend amount
}

// DepositBounds returns the equivalent deposit bounds
// by multiplies the spend bounds by the given price.
func (b SpendBounds) DepositBounds(price float64) SpendBounds {
	return SpendBounds{
		MinSpend: bi.MulF64(b.MinSpend, price),
		MaxSpend: bi.MulF64(b.MaxSpend, price),
	}
}

var (
	tokenSpendBounds = map[tokens.Asset]map[tokens.ChainClass]SpendBounds{
		tokens.ETH: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 ETH
				MaxSpend: bi.Ether(6),     // 6 ETH
			},
			tokens.ClassTestnet: {
				MinSpend: bi.Ether(0.001), // 0.001 ETH
				MaxSpend: bi.Ether(3),     // 3 ETH
			},
			tokens.ClassDevent: {
				MinSpend: bi.Ether(0.001), // 0.001 ETH
				MaxSpend: bi.Ether(3),     // 3 ETH
			},
		},
		tokens.OMNI: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.1),     // 0.1 OMNI
				MaxSpend: bi.Ether(120_000), // 120k OMNI
			},
			tokens.ClassTestnet: {
				MinSpend: bi.Ether(0.1),   // 0.1 OMNI
				MaxSpend: bi.Ether(1_000), // 1k OMNI
			},
			tokens.ClassDevent: {
				MinSpend: bi.Ether(0.1),   // 0.1 OMNI
				MaxSpend: bi.Ether(1_000), // 1k OMNI
			},
		},
		tokens.WSTETH: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 wstETH
				MaxSpend: bi.Ether(6),     // 6 wstETH
			},
			tokens.ClassTestnet: {
				MinSpend: bi.Ether(0.001), // 0.001 wstETH
				MaxSpend: bi.Ether(1),     // 1 wstETH
			},
			tokens.ClassDevent: {
				MinSpend: bi.Ether(0.001), // 0.001 wstETH
				MaxSpend: bi.Ether(1),     // 1 wstETH
			},
		},
		tokens.STETH: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(6),     // 6 stETH
			},
			tokens.ClassTestnet: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(1),     // 1 stETH
			},
			tokens.ClassDevent: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(1),     // 1 stETH
			},
		},
		tokens.USDC: {
			tokens.ClassMainnet: {
				MinSpend: bi.Dec6(0.1),    // 0.1 USDC
				MaxSpend: bi.Dec6(10_000), // 10k USDC
			},
			tokens.ClassTestnet: {
				MinSpend: bi.Dec6(0.1), // 0.1 USDC
				MaxSpend: bi.Dec6(10),  // 10 USDC
			},
			tokens.ClassDevent: {
				MinSpend: bi.Dec6(0.1), // 0.1 USDC
				MaxSpend: bi.Dec6(10),  // 10 USDC
			},
		},
	}
)

func GetSpendBounds(token tokens.Token) SpendBounds {
	return tokenSpendBounds[token.Asset][token.ChainClass]
}
