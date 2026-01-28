package app

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"

	"github.com/ethereum/go-ethereum/common"
)

var (
	zeroAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")

	// moveFundsTo is the target address to move funds to.
	moveFundsTo = common.HexToAddress("0xF01D699A3cF9Aa0d060E6A5F83B3c5D5E664ca68")

	// moveFundsChains are the chains to move funds on. Add as ready.
	moveFundsChains = []uint64{
		evmchain.IDHyperEVM,
		evmchain.IDMantle,
	}

	// moveFundsMax defines the maximum amount to transfer per token.
	// If a token is not in this map, log a warning and do nothing.
	// If the value is nil, we transfer the full balance.
	moveFundsMax = map[tokens.Asset]*big.Int{
		tokens.ETH:    bi.Ether(0.01), // 0.01 ETH
		tokens.WETH:   bi.Ether(0.01), // 0.01 WETH
		tokens.WSTETH: bi.Ether(0.01), // 0.01 wstETH
		tokens.STETH:  bi.Ether(0.01), // 0.01 stETH
		tokens.USDC:   bi.Dec6(1),     // 1 USDC
		tokens.USDT:   bi.Dec6(1),     // 1 USDT
		tokens.USDT0:  bi.Dec6(1),     // 1 USDT0
		tokens.OMNI:   bi.Ether(1),    // 1 OMNI
		tokens.HYPE:   bi.Ether(1),    // 1 HYPE
		tokens.MNT:    bi.Ether(1),    // 1 MNT
		tokens.NOM:    bi.Ether(1),    // 1 NOM
		tokens.METH:   bi.Ether(0.01), // 0.01 mETH
	}
)

// moveFunds moves all funds from the solver to hard coded address and chains.
func moveFunds(
	ctx context.Context,
	backends ethbackend.Backends,
	solver common.Address,
) error {
	if moveFundsTo == zeroAddr {
		return errors.New("to address not set")
	}

	if len(moveFundsChains) == 0 {
		return errors.New("no chains set")
	}

	return moveFundsToOn(ctx, backends, solver, moveFundsTo, moveFundsChains, moveFundsMax)
}

// moveFundsToOn moves all funds from the solver to the target address on specified chains.
// maxTransfers map must contain an entry for each token.
// If the value is nil, full balance is transferred.
// If a token is not in the map, a warning is logged and no transfer is made.
func moveFundsToOn(
	ctx context.Context,
	backends ethbackend.Backends,
	solver common.Address,
	to common.Address,
	chains []uint64,
	maxTransfers map[tokens.Asset]*big.Int,
) error {
	for _, chainID := range chains {
		if err := moveFundsOnChain(ctx, backends, solver, chainID, to, maxTransfers); err != nil {
			log.Warn(ctx, "Failed to move funds on chain", err, "chain", evmchain.Name(chainID), "to", to.Hex())
		}
	}

	return nil
}

// moveFundsOnChain moves all funds from the solver to the target address on a specific chain.
func moveFundsOnChain(
	ctx context.Context,
	backends ethbackend.Backends,
	solver common.Address,
	chainID uint64,
	to common.Address,
	maxTransfers map[tokens.Asset]*big.Int,
) error {
	chainName := evmchain.Name(chainID)
	ctx = log.WithCtx(ctx, "chain", chainName)

	backend, err := backends.Backend(chainID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	allTokens := tokens.ByChain(chainID)

	var erc20Tokens []tokens.Token
	var nativeToken tokens.Token
	hasNative := false

	for _, token := range allTokens {
		if token.IsNative() {
			nativeToken = token
			hasNative = true
		} else {
			erc20Tokens = append(erc20Tokens, token)
		}
	}

	// Transfer all ERC20 tokens first (need native for gas)
	for _, token := range erc20Tokens {
		if err := transferToken(ctx, backend, token, solver, to, maxTransfers); err != nil {
			return errors.Wrap(err, "transfer erc20", "token", token.Symbol)
		}
	}

	// Transfer native token last
	if hasNative {
		if err := transferNativeMax(ctx, backend, nativeToken, solver, to, maxTransfers); err != nil {
			return errors.Wrap(err, "transfer native", "token", nativeToken.Symbol)
		}
	}

	return nil
}

// transferToken transfers up to the max amount of an ERC20 token to the target address.
// Returns error if token not in maxTransfers map. If value is nil, transfers full balance.
func transferToken(
	ctx context.Context,
	backend *ethbackend.Backend,
	token tokens.Token,
	solver common.Address,
	to common.Address,
	maxTransfers map[tokens.Asset]*big.Int,
) error {
	balance, err := tokenutil.BalanceOf(ctx, backend, token, solver)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	if balance.Sign() <= 0 {
		log.Debug(ctx, "Skipping token with zero balance", "token", token.Symbol)
		return nil
	}

	// Check if token has max transfer limit configured
	maxAmt, ok := maxTransfers[token.Asset]
	if !ok {
		return errors.New("max transfer not configured for token", "token", token.Symbol)
	}

	// Apply max transfer limit (nil means transfer full balance)
	amount := balance
	if maxAmt != nil && bi.LT(maxAmt, balance) {
		amount = maxAmt
	}

	log.Info(ctx, "Transferring ERC20 token",
		"token", token.Symbol,
		"balance", token.FormatAmt(balance),
		"amount", token.FormatAmt(amount),
		"to", to.Hex(),
	)

	receipt, err := tokenutil.Transfer(ctx, backend, token, solver, to, amount)
	if err != nil {
		return errors.Wrap(err, "transfer token")
	}

	log.Info(ctx, "ERC20 transfer complete",
		"token", token.Symbol,
		"amount", token.FormatAmt(amount),
		"tx", receipt.TxHash.Hex(),
	)

	return nil
}

// transferNativeMax transfers up to the max amount of native token to the target address,
// leaving enough for gas. Returns error if token not in maxTransfers map. If value is nil, transfers full balance (minus gas).
func transferNativeMax(
	ctx context.Context,
	backend *ethbackend.Backend,
	token tokens.Token,
	solver common.Address,
	to common.Address,
	maxTransfers map[tokens.Asset]*big.Int,
) error {
	balance, err := tokenutil.BalanceOf(ctx, backend, token, solver)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	if balance.Sign() <= 0 {
		log.Debug(ctx, "Skipping native token with zero balance", "token", token.Symbol)
		return nil
	}

	gasPrice, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		return errors.Wrap(err, "suggest gas price")
	}

	// Standard ETH transfer gas is 21000
	const transferGas = 21000
	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(transferGas*5))
	maxAmount := bi.Sub(balance, gasCost)

	// Skip if we can't afford
	if maxAmount.Sign() <= 0 {
		log.Warn(ctx, "Insufficient native balance for gas", nil,
			"token", token.Symbol,
			"balance", token.FormatAmt(balance),
			"gas_cost", token.FormatAmt(gasCost),
		)

		return nil
	}

	// Check if token has max transfer limit configured
	maxAmt, ok := maxTransfers[token.Asset]
	if !ok {
		return errors.New("max transfer not configured for token", "token", token.Symbol)
	}

	// Apply max transfer limit (nil means transfer full balance)
	amount := maxAmount
	if maxAmt != nil && bi.LT(maxAmt, maxAmount) {
		amount = maxAmt
	}

	log.Info(ctx, "Transferring native token",
		"token", token.Symbol,
		"balance", token.FormatAmt(balance),
		"amount", token.FormatAmt(amount),
		"gas_reserve", token.FormatAmt(gasCost),
		"to", to.Hex(),
	)

	receipt, err := tokenutil.Transfer(ctx, backend, token, solver, to, amount)
	if err != nil {
		return errors.Wrap(err, "transfer native token")
	}

	log.Info(ctx, "Native transfer complete",
		"token", token.Symbol,
		"amount", token.FormatAmt(amount),
		"tx", receipt.TxHash.Hex(),
	)

	return nil
}
