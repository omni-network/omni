package uniswap

import (
	"context"
	"encoding/binary"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Swap represents a single hop in a Uniswap V3 swap path.
type Swap struct {
	TokenIn  tokens.Token
	TokenOut tokens.Token
	PoolFee  uint32
}

// SwapToUSDC swaps a token for USDC using Uniswap V3.
func SwapToUSDC(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	token tokens.Token,
	amount *big.Int,
) (*big.Int, error) {
	if token.Is(tokens.USDC) {
		return amount, nil
	}

	swaps, err := routeToUSDC(token)
	if err != nil {
		return nil, errors.Wrap(err, "route to USDC")
	}

	if err := maybeApproveRouter(ctx, backend, user, token, amount); err != nil {
		return nil, errors.Wrap(err, "approve token")
	}

	amountOut, receipt, err := executeSwaps(ctx, backend, user, token, swaps, amount)
	if err != nil {
		return nil, errors.Wrap(err, "execute swaps")
	}

	log.Debug(ctx, "Swapped for USDC",
		"in", token.FormatAmt(amount),
		"out", swaps[len(swaps)-1].TokenOut.FormatAmt(amountOut),
		"tx", receipt.TxHash,
		"gas_used", receipt.GasUsed,
	)

	return amountOut, nil
}

// routeToUSDC returns the optimal swap path to convert a token to USDC.
func routeToUSDC(tkn tokens.Token) ([]Swap, error) {
	usdc, ok := tokens.ByAsset(tkn.ChainID, tokens.USDC)
	if !ok {
		return nil, errors.New("no USDC", "chain", tkn.ChainID)
	}

	weth, ok := tokens.ByAsset(tkn.ChainID, tokens.WETH)
	if !ok {
		return nil, errors.New("no WETH", "chain", tkn.ChainID)
	}

	if tkn.Is(tokens.ETH) { // Swap WETH to USDC, direct
		return []Swap{wethToUSDC(weth, usdc)}, nil
	}

	if tkn.Is(tokens.USDT) { // Swap USDT to USDC, direct
		return []Swap{usdtToUSDC(tkn, usdc)}, nil
	}

	if tkn.Is(tokens.WSTETH) { // Swap WSTETH to USDC, via WETH
		return []Swap{wstethToWETH(tkn, weth), wethToUSDC(weth, usdc)}, nil
	}

	return nil, errors.New("no route to USDC", "token", tkn.Asset, "chain", tkn.ChainID)
}

// usdtToUSDC returns a swap from USDT to USDC, 0.01% fee tier.
func usdtToUSDC(usdt, usdc tokens.Token) Swap {
	return Swap{TokenIn: usdt, TokenOut: usdc, PoolFee: FeeBips1}
}

// wethToUSDC returns a swap from WETH to USDC, 0.05% fee tier.
func wethToUSDC(weth, usdc tokens.Token) Swap {
	return Swap{TokenIn: weth, TokenOut: usdc, PoolFee: FeeBips5}
}

// wstethToWETH returns a swap from WSTETH to WETH, 0.01% fee tier.
func wstethToWETH(wsteth, weth tokens.Token) Swap {
	return Swap{TokenIn: wsteth, TokenOut: weth, PoolFee: FeeBips1}
}

// executeSwaps executes a series of Uniswap V3 swaps and returns the amount received.
func executeSwaps(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	tokenIn tokens.Token,
	swaps []Swap,
	amountIn *big.Int,
) (*big.Int, *ethtypes.Receipt, error) {
	chainID := swaps[0].TokenIn.ChainID

	router, err := newRouter(chainID, backend)
	if err != nil {
		return nil, nil, errors.Wrap(err, "new router")
	}

	quoter, err := newQuoter(chainID, backend)
	if err != nil {
		return nil, nil, errors.Wrap(err, "new quoter")
	}

	path, err := encodePath(swaps)
	if err != nil {
		return nil, nil, errors.Wrap(err, "encode path")
	}

	amountOut, err := quoter.CallQuoteExactInput(ctx, path, amountIn)
	if err != nil {
		return nil, nil, errors.Wrap(err, "quote exact input")
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return nil, nil, errors.Wrap(err, "bind opts")
	}

	if tokenIn.IsNative() {
		txOpts.Value = amountIn
	}

	params := IV3SwapRouterExactInputParams{
		Path:             path,
		Recipient:        user,
		AmountIn:         amountIn,
		AmountOutMinimum: amountOut,
	}

	tx, err := router.ExactInput(txOpts, params)
	if err != nil {
		return nil, nil, errors.Wrap(err, "exact input")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "wait mined")
	}

	return amountOut, receipt, nil
}

// maybeApproveRouter checks if the router needs approval for the token and approves if necessary.
func maybeApproveRouter(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	token tokens.Token,
	amount *big.Int,
) error {
	if token.IsNative() {
		return nil
	}

	addr, ok := routers[token.ChainID]
	if !ok {
		return errors.New("no router", "chain", token.ChainID)
	}

	erc20, err := bindings.NewIERC20(token.Address, backend)
	if err != nil {
		return errors.Wrap(err, "new token")
	}

	allowance, err := erc20.Allowance(&bind.CallOpts{Context: ctx}, user, addr)
	if err != nil {
		return errors.Wrap(err, "get allowance")
	}

	if bi.GTE(allowance, amount) {
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := erc20.Approve(txOpts, addr, umath.MaxUint256)
	if err != nil {
		return errors.Wrap(err, "approve")
	}

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Approved token spend",
		"token", token.Symbol,
		"router", addr.Hex(),
		"tx", tx.Hash().Hex(),
	)

	return nil
}

// newRouter creates a new SwapRouter02 contract instance for the given chain.
func newRouter(chainID uint64, backend *ethbackend.Backend) (*UniSwapRouter02, error) {
	addr, ok := routers[chainID]
	if !ok {
		return nil, errors.New("no router", "chain", chainID)
	}

	return NewUniSwapRouter02(addr, backend)
}

// newQuoter creates a new QuoterV2 contract instance for the given chain.
func newQuoter(chainID uint64, backend *ethbackend.Backend) (*UniQuoterV2, error) {
	addr, ok := quoters[chainID]
	if !ok {
		return nil, errors.New("no quoter", "chain", chainID)
	}

	return NewUniQuoterV2(addr, backend)
}

// encodePath converts swaps to Uniswap V3 path format.
// Each hop is tokenIn (20b) + fee (3b) + tokenOut (20b).
// Multi-hop swaps use intermediate tokens.
func encodePath(swaps []Swap) ([]byte, error) {
	if len(swaps) == 0 {
		return nil, errors.New("empty swaps")
	}

	var path []byte

	for i, swap := range swaps {
		path = append(path, swap.TokenIn.Address.Bytes()...)

		// Write fee as 3 bytes (uint24)
		feeBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(feeBytes, swap.PoolFee)
		path = append(path, feeBytes[1:]...) // Take last 3 bytes

		// If last, append tokenOut address
		if i == len(swaps)-1 {
			path = append(path, swap.TokenOut.Address.Bytes()...)
			continue
		}

		// If not last, asset swap.TokenOut == next.TokenIn
		if swap.TokenOut != swaps[i+1].TokenIn {
			return nil, errors.New("invalid swap path")
		}
	}

	return path, nil
}
