package app

import (
	"math/big"
	"sort"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	supportedAssets = map[tokens.Asset]bool{
		tokens.ETH:    true,
		tokens.OMNI:   true,
		tokens.WSTETH: true,
		tokens.STETH:  true,
		tokens.USDC:   true,
		tokens.USDT:   true,
	}

	// minSafeETH is the minimum amount of ETH the solver can leave itself with post-fill,
	// to ensure it can pay for gas of other orders.
	minSafeETH = bi.Ether(0.05)
)

func IsSupportedToken(token tokens.Token) bool {
	return supportedAssets[token.Asset]
}

type SpendBounds struct {
	MinSpend *big.Int // minimum spend amount
	MaxSpend *big.Int // maximum spend amount
}

// DepositBounds returns the equivalent deposit bounds
// by dividing the spend bounds (expense) by the given price.
func (b SpendBounds) DepositBounds(price types.Price) SpendBounds {
	return SpendBounds{
		MinSpend: price.ToDeposit(b.MinSpend),
		MaxSpend: price.ToDeposit(b.MaxSpend),
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
				MaxSpend: bi.Ether(0.01),  // 0.1 wstETH
			},
			tokens.ClassDevent: {
				MinSpend: bi.Ether(0.001), // 0.001 wstETH
				MaxSpend: bi.Ether(0.01),  // 0.1 wstETH
			},
		},
		tokens.STETH: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(6),     // 6 stETH
			},
			tokens.ClassTestnet: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(0.1),   // 0.1 stETH
			},
			tokens.ClassDevent: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(0.1),   // 0.1 stETH
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
		tokens.USDT: {
			tokens.ClassMainnet: {
				MinSpend: bi.Dec6(0.1),    // 0.1 USDT
				MaxSpend: bi.Dec6(10_000), // 10 USDT
			},
		},
	}

	// chainOverrides map token -> chain id -> override spend bounds.
	chainOverrides = map[tokens.Asset]map[uint64]SpendBounds{
		tokens.ETH: {
			evmchain.IDEthereum: {
				MinSpend: bi.Ether(0.001), // 0.001 ETH
				MaxSpend: bi.Ether(65),    // 65 ETH
			},
		},
		tokens.USDC: {
			evmchain.IDEthereum: {
				MinSpend: bi.Dec6(0.1),     // 0.1 USDC
				MaxSpend: bi.Dec6(100_000), // 100k USDC
			},
		},
	}
)

func GetSpendBounds(token tokens.Token) (SpendBounds, bool) {
	override, ok := chainOverrides[token.Asset][token.ChainID]
	if ok {
		return override, true
	}

	resp, ok := tokenSpendBounds[token.Asset][token.ChainClass]

	return resp, ok
}

func tokensResponse(chains []uint64) (types.TokensResponse, error) {
	var resp []types.TokenResponse
	for _, chain := range chains {
		for asset := range supportedAssets {
			token, ok := tokens.ByAsset(chain, asset)
			if !ok {
				continue
			}

			bounds, ok := GetSpendBounds(token)
			if !ok {
				return types.TokensResponse{}, errors.New("invalid token")
			}

			resp = append(resp, types.TokenResponse{
				Enabled:    true, // Disable not supported yet.
				Name:       token.Name,
				Symbol:     token.Symbol,
				ChainID:    token.ChainID,
				Address:    token.Address,
				Decimals:   token.Decimals,
				ExpenseMin: (*hexutil.Big)(bounds.MinSpend),
				ExpenseMax: (*hexutil.Big)(bounds.MaxSpend),
			})
		}
	}

	sort.Slice(resp, func(i, j int) bool {
		if resp[i].ChainID != resp[j].ChainID {
			return resp[i].ChainID < resp[j].ChainID
		}

		return resp[i].Symbol < resp[j].Symbol
	})

	return types.TokensResponse{
		Tokens: resp,
	}, nil
}
