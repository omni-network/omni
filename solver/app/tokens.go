package app

import (
	"context"
	"math/big"
	"sort"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/unibackend"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	// supportedAssets maps asset to true, if the asset is supported on all chains.
	supportedAssets = map[tokens.Asset]bool{
		tokens.ETH:    true,
		tokens.NOM:    true,
		tokens.WSTETH: true,
		tokens.USDC:   true,
		tokens.USDT:   true,
		tokens.WETH:   true,
		tokens.METH:   true,
		tokens.MNT:    true,
		tokens.HYPE:   true,
		tokens.USDT0:  true,
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
		tokens.NOM: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(1),          // 1 NOM
				MaxSpend: bi.Ether(35_000_000), // 35M NOM
			},
			tokens.ClassTestnet: {
				MinSpend: bi.Ether(1),       // 1 NOM
				MaxSpend: bi.Ether(100_000), // 100k NOM
			},
			tokens.ClassDevent: {
				MinSpend: bi.Ether(1),       // 1 NOM
				MaxSpend: bi.Ether(100_000), // 100k NOM
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
				MaxSpend: bi.Dec6(10_000), // 10k USDT
			},
		},
		tokens.WETH: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 WETH
				MaxSpend: bi.Ether(3),     // 3 WETH
			},
		},
		tokens.METH: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 METH
				MaxSpend: bi.Ether(3),     // 3 METH
			},
		},
		tokens.MNT: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 MNT
				MaxSpend: bi.Ether(1000),  // 1000 MNT
			},
		},
		tokens.HYPE: {
			tokens.ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 HYPE
				MaxSpend: bi.Ether(100),   // 100 HYPE
			},
		},
		tokens.USDT0: {
			tokens.ClassMainnet: {
				MinSpend: bi.Dec6(0.1),       // 0.1 USDT0
				MaxSpend: bi.Dec6(2_000_000), // 2M USDT0
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

func tokensResponse(ctx context.Context, backends unibackend.Backends, solverAddr common.Address) (types.TokensResponse, error) {
	// Get all tokens we support
	var tkns []tokens.Token
	for chainID := range backends {
		for asset := range supportedAssets {
			token, ok := tokens.ByAsset(chainID, asset)
			if !ok || token.IsSVM() {
				continue
			}

			tkns = append(tkns, token)
		}
	}

	// Define a forkjoin work function to process each token concurrently.
	workFn := func(ctx context.Context, token tokens.Token) (types.TokenResponse, error) {
		backend, err := backends.Backend(token.ChainID)
		if err != nil {
			return types.TokenResponse{}, err
		}

		bounds, ok := GetSpendBounds(token)
		if !ok {
			return types.TokenResponse{}, errors.New("invalid token spend bounds", "token", token)
		}

		var expenseEnabled bool
		inventory := bi.Zero()
		if backend.IsEVM() {
			bal, err := tokenutil.BalanceOf(ctx, backend.EVMBackend(), token, solverAddr)
			if err == nil { // Disable expense if query fails (should be temporary).
				expenseEnabled = bi.GT(bal, bounds.MinSpend)
				inventory = bal
			}
		}

		return types.TokenResponse{
			Enabled:          expenseEnabled,
			ExpenseEnabled:   expenseEnabled,
			DepositEnabled:   true, // Always enabled for deposits for now.
			Name:             token.Name,
			Symbol:           token.Symbol,
			ChainID:          token.ChainID,
			Address:          token.UniAddress(),
			Decimals:         token.Decimals,
			ExpenseMin:       (*hexutil.Big)(bounds.MinSpend),
			ExpenseMax:       (*hexutil.Big)(bounds.MaxSpend),
			ExpenseInventory: (*hexutil.Big)(inventory),
		}, nil
	}

	// Do forkjoin
	result, cancel := forkjoin.NewWithInputs(ctx, workFn, tkns)
	defer cancel()
	resp, err := result.Flatten()
	if err != nil {
		return types.TokensResponse{}, errors.Wrap(err, "fork join tokens response")
	}

	// Sort deterministically.
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
