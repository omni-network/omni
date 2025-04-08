package app

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/tokenmeta"
	"github.com/omni-network/omni/lib/tokens"
)

var (
	supportedTokens = map[tokenmeta.Meta]bool{
		tokenmeta.ETH:    true,
		tokenmeta.OMNI:   true,
		tokenmeta.WSTETH: true,
		tokenmeta.STETH:  true,
		tokenmeta.USDC:   true,
	}
)

func IsSupportedToken(token tokens.Token) bool {
	return supportedTokens[token.Meta]
}

type SpendBounds struct {
	MinSpend *big.Int // minimum spend amount
	MaxSpend *big.Int // maximum spend amount
}

var (
	tokenSpendBounds = map[tokenmeta.Meta]map[tokens.ChainClass]SpendBounds{
		tokenmeta.ETH: {
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
		tokenmeta.OMNI: {
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
		tokenmeta.WSTETH: {
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
		tokenmeta.STETH: {
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
		tokenmeta.USDC: {
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
	return tokenSpendBounds[token.Meta][token.ChainClass]
}
