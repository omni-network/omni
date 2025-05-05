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

// SwapToUSDC swaps `amountIn` of `token` to USDC, returning the minimum USDC received.
func SwapToUSDC(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	token tokens.Token,
	amountIn *big.Int,
) (*big.Int, error) {
	if token.Is(tokens.USDC) {
		return amountIn, nil
	}

	swaps, err := routeToUSDC(token)
	if err != nil {
		return nil, errors.Wrap(err, "route to USDC")
	}

	if err := maybeApproveRouter(ctx, backend, user, token, amountIn); err != nil {
		return nil, errors.Wrap(err, "approve router")
	}

	amountOutMin, receipt, err := swapExactInput(ctx, backend, user, token, swaps, amountIn)
	if err != nil {
		return nil, errors.Wrap(err, "execute swaps")
	}

	log.Debug(ctx, "Swapped to USDC",
		"in", token.FormatAmt(amountIn),
		"out", swaps[len(swaps)-1].TokenOut.FormatAmt(amountOutMin),
		"tx", receipt.TxHash,
		"gas_used", receipt.GasUsed,
	)

	return amountOutMin, nil
}

// SwapUSDCTo swaps USDC for `amountOut` of `token`, returning the max USDC spent.
func SwapUSDCTo(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	token tokens.Token,
	amountOut *big.Int,
) (*big.Int, error) {
	if token.Is(tokens.USDC) {
		return amountOut, nil
	}

	swaps, err := routeFromUSDC(token)
	if err != nil {
		return nil, errors.Wrap(err, "route to USDC")
	}

	// Need to approve quoted max in
	if err := maybeApproveRouter(ctx, backend, user, swaps[0].TokenIn, umath.MaxUint256); err != nil {
		return nil, errors.Wrap(err, "approve router")
	}

	amountInMax, receipt, err := swapExactOutput(ctx, backend, user, swaps, amountOut)
	if err != nil {
		return nil, errors.Wrap(err, "execute swaps")
	}

	// If output token is ETH, swap will return WETH. We need to unwrap
	// TODO(kevin): unwrap a single tx using SwapRouter02.multicall
	if token.Is(tokens.ETH) {
		receipt, err := unwrapWETH(ctx, token.ChainID, backend, user, amountOut)
		if err != nil {
			return nil, errors.Wrap(err, "unwrap WETH")
		}

		log.Debug(ctx, "Unwrapped WETH",
			"amount", tokens.WETH.FormatAmt(amountOut),
			"tx", receipt.TxHash,
			"gas_used", receipt.GasUsed)
	}

	log.Debug(ctx, "Swapped USDC to",
		"in", swaps[0].TokenIn.FormatAmt(amountInMax),
		"out", token.FormatAmt(amountOut),
		"tx", receipt.TxHash,
		"gas_used", receipt.GasUsed)

	return amountInMax, nil
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

// routeFromUSDC returns the optimal swap path to convert USDC to a token.
func routeFromUSDC(tkn tokens.Token) ([]Swap, error) {
	to, err := routeToUSDC(tkn)
	if err != nil {
		return nil, errors.Wrap(err, "route to USDC")
	}

	return reverse(to), nil
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

// reverse reverses a swap path, swapping order and token in <> out for each swap.
func reverse(swaps []Swap) []Swap {
	reversed := make([]Swap, len(swaps))
	for i, swap := range swaps {
		reversed[len(swaps)-1-i] = Swap{
			TokenIn:  swap.TokenOut,
			TokenOut: swap.TokenIn,
			PoolFee:  swap.PoolFee,
		}
	}

	return reversed
}

// swapExactInput executes a series of Uniswap V3 swaps with exact input, and returns the minimum output amount.
func swapExactInput(
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

	amountOutMin, err := quoter.CallQuoteExactInput(ctx, path, amountIn)
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
		AmountOutMinimum: amountOutMin,
	}

	tx, err := router.ExactInput(txOpts, params)
	if err != nil {
		return nil, nil, errors.Wrap(err, "exact input")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "wait mined")
	}

	return amountOutMin, receipt, nil
}

// swapExactOutput executes a series of Uniswap V3 swaps with exact output, and returns the maximum input amount.
func swapExactOutput(
	ctx context.Context,
	backend *ethbackend.Backend,
	user common.Address,
	swaps []Swap,
	amountOut *big.Int,
) (*big.Int, *ethtypes.Receipt, error) {
	tokenIn := swaps[0].TokenIn
	tokenOut := swaps[len(swaps)-1].TokenOut
	chainID := tokenIn.ChainID

	router, err := newRouter(chainID, backend)
	if err != nil {
		return nil, nil, errors.Wrap(err, "new router")
	}

	quoter, err := newQuoter(chainID, backend)
	if err != nil {
		return nil, nil, errors.Wrap(err, "new quoter")
	}

	// QuoterV2.quoteExactOutput expects the path to be reversed
	path, err := encodePath(reverse(swaps))
	if err != nil {
		return nil, nil, errors.Wrap(err, "encode path")
	}

	amountInMax, err := quoter.CallQuoteExactOutput(ctx, path, amountOut)
	if err != nil {
		return nil, nil, errors.Wrap(err, "quote exact output")
	}


	log.Debug(ctx, "Quoted exact output",
		"token_in", tokenIn.Asset,
		"token_out", tokenOut.Asset,
		"amount_in_max", tokenIn.FormatAmt(amountInMax),
		"amount_out", tokenOut.FormatAmt(amountOut))

	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return nil, nil, errors.Wrap(err, "bind opts")
	}

	if tokenIn.IsNative() {
		log.Debug(ctx, "Adding value to tx opts")
		txOpts.Value = amountInMax // Overpayment is refunded
	}

	params := IV3SwapRouterExactOutputParams{
		Path:            path,
		Recipient:       user,
		AmountOut:       amountOut,
		AmountInMaximum: amountInMax,
	}

	tx, err := router.ExactOutput(txOpts, params)
	if err != nil {
		return nil, nil, errors.Wrap(err, "exact output")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "wait mined")
	}

	return amountInMax, receipt, nil
}

func unwrapWETH(
	ctx context.Context,
	chainID uint64,
	backend *ethbackend.Backend,
	user common.Address,
	amount *big.Int,
) (*ethtypes.Receipt, error) {
	txOpts, err := backend.BindOpts(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "bind opts")
	}

	weth, ok := tokens.ByAsset(chainID, tokens.WETH)
	if !ok {
		return nil, errors.New("no WETH", "chain", chainID)
	}

	contract, err := NewWETH9(weth.Address, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new WETH9")
	}

	tx, err := contract.Withdraw(txOpts, amount)
	if err != nil {
		return nil, errors.Wrap(err, "unwrap WETH")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return nil, errors.Wrap(err, "wait mined")
	}

	return receipt, nil
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
