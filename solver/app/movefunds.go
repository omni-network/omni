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
	moveFundsTo = zeroAddr

	// moveFundsChains are the chains to move funds on. Add as ready.
	moveFundsChains = []uint64{}
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

	// TODO: Add pricer, start with max $1 transfers

	return moveFundsToOn(ctx, backends, solver, moveFundsTo, moveFundsChains)
}

// moveFundsToOn moves all funds from the solver to the target address on specified chains.
func moveFundsToOn(
	ctx context.Context,
	backends ethbackend.Backends,
	solver common.Address,
	to common.Address,
	chains []uint64,
) error {
	for _, chainID := range chains {
		if err := moveFundsOnChain(ctx, backends, solver, chainID, to); err != nil {
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
		if err := transferToken(ctx, backend, token, solver, to); err != nil {
			return errors.Wrap(err, "transfer erc20", "token", token.Symbol)
		}
	}

	// Transfer native token last
	if hasNative {
		if err := transferNativeMax(ctx, backend, nativeToken, solver, to); err != nil {
			return errors.Wrap(err, "transfer native", "token", nativeToken.Symbol)
		}
	}

	return nil
}

// transferToken transfers the full balance of an ERC20 token to the target address.
func transferToken(
	ctx context.Context,
	backend *ethbackend.Backend,
	token tokens.Token,
	solver common.Address,
	to common.Address,
) error {
	balance, err := tokenutil.BalanceOf(ctx, backend, token, solver)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	if balance.Sign() <= 0 {
		log.Debug(ctx, "Skipping token with zero balance", "token", token.Symbol)
		return nil
	}

	log.Info(ctx, "Transferring ERC20 token",
		"token", token.Symbol,
		"balance", balance.String(),
		"to", to.Hex(),
	)

	receipt, err := tokenutil.Transfer(ctx, backend, token, solver, to, balance)
	if err != nil {
		return errors.Wrap(err, "transfer token")
	}

	log.Info(ctx, "ERC20 transfer complete",
		"token", token.Symbol,
		"tx", receipt.TxHash.Hex(),
	)

	return nil
}

// transferNativeMax transfers the maximum possible native token balance to the target address,
// leaving enough for gas.
func transferNativeMax(
	ctx context.Context,
	backend *ethbackend.Backend,
	token tokens.Token,
	solver common.Address,
	to common.Address,
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
			"balance", balance.String(),
			"gas_cost", gasCost.String(),
		)

		return nil
	}

	log.Info(ctx, "Transferring native token",
		"token", token.Symbol,
		"balance", balance.String(),
		"amount", maxAmount.String(),
		"gas_reserve", gasCost.String(),
		"to", to.Hex(),
	)

	receipt, err := tokenutil.Transfer(ctx, backend, token, solver, to, maxAmount)
	if err != nil {
		return errors.Wrap(err, "transfer native token")
	}

	log.Info(ctx, "Native transfer complete",
		"token", token.Symbol,
		"tx", receipt.TxHash.Hex(),
	)

	return nil
}
