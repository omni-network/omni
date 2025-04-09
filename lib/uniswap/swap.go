//nolint:unparam // skeleton code
package uniswap

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/tokens"
)

// SwapToUSDC swaps the given token to USDC.
func SwapToUSDC(
	ctx context.Context,
	backend *ethbackend.Backend,
	token tokens.Token,
	amount *big.Int,
) (*big.Int, error) {
	if token.Is(tokens.USDC) { // Already USDC
		return amount, nil
	}

	swaps, err := routeToUSDC(token)
	if err != nil {
		return nil, errors.Wrap(err, "route to USDC")
	}

	if len(swaps) == 1 {
		return singleSwap(ctx, backend, swaps[0], amount)
	}

	return mulitSwap(ctx, backend, swaps, amount)
}

type Swap struct {
	TokenIn  tokens.Token
	TokenOut tokens.Token
}

// routeToUSDC returns a list of swap to get USDC.
// Routes defined statically, dynamic route optimization not (yet) worth the complexity.
func routeToUSDC(tkn tokens.Token) ([]Swap, error) {
	usdc, ok := tokens.ByAsset(tkn.ChainID, tokens.USDC)
	if !ok {
		return nil, errors.New("no USDC", "chain", tkn.ChainID)
	}

	// WETH required native ETH -> USDC. SwapRouter02 wraps ETH to WETH.
	weth, ok := tokens.ByAsset(tkn.ChainID, tokens.WETH)
	if !ok {
		return nil, errors.New("no WETH", "chain", tkn.ChainID)
	}

	if tkn.Is(tokens.ETH) { // Swap ETH to USDC directly.
		return []Swap{
			{TokenIn: weth, TokenOut: usdc},
		}, nil
	}

	if tkn.Is(tokens.WSTETH) { // Swap WSTETH to USDC via WETH.
		return []Swap{
			{TokenIn: tkn, TokenOut: weth},
			{TokenIn: weth, TokenOut: usdc},
		}, nil
	}

	return nil, errors.New("no route to USDC", "token", tkn.Asset, "chain", tkn.ChainID)
}

// singleSwap executes a single swap.
func singleSwap(
	ctx context.Context,
	backend *ethbackend.Backend,
	swap Swap,
	amountIn *big.Int,
) (*big.Int, error) {
	chainID := swap.TokenIn.ChainID

	router, err := newRouter(chainID, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new router")
	}

	quoter, err := newQuoter(chainID, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new quoter")
	}

	_ = ctx
	_ = amountIn
	_ = swap
	_ = quoter
	_ = router

	// TODO

	return bi.Zero(), nil
}

// mulitSwap executes a multi hop swap.
func mulitSwap(
	ctx context.Context,
	backend *ethbackend.Backend,
	swaps []Swap,
	amountIn *big.Int,
) (*big.Int, error) {
	if len(swaps) < 2 {
		return nil, errors.New("no multi hop", "swaps", len(swaps))
	}

	chainID := swaps[0].TokenIn.ChainID

	router, err := newRouter(chainID, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new router")
	}

	quoter, err := newQuoter(chainID, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new quoter")
	}

	_ = ctx
	_ = amountIn
	_ = swaps
	_ = quoter
	_ = router

	// TODO

	return bi.Zero(), nil
}

func newRouter(chainID uint64, backend *ethbackend.Backend) (*UniSwapRouter02, error) {
	addr, ok := routers[chainID]
	if !ok {
		return nil, errors.New("no router", "chain", chainID)
	}

	return NewUniSwapRouter02(addr, backend)
}

func newQuoter(chainID uint64, backend *ethbackend.Backend) (*UniQuoterV2, error) {
	addr, ok := quoters[chainID]
	if !ok {
		return nil, errors.New("no quoter", "chain", chainID)
	}

	return NewUniQuoterV2(addr, backend)
}
