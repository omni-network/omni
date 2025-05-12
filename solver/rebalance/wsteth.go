package rebalance

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate abigen --abi wsteth-abi.json --type WSTETH --pkg rebalance --out wsteth_bindings.go

// wrapSTETH takes the solvers stETH balance and wraps it to wstETH.
func wrapSTETH(
	ctx context.Context,
	backends ethbackend.Backends,
	solver common.Address,
) error {
	backend, err := backends.Backend(evmchain.IDEthereum)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	steth, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.STETH)
	if !ok {
		return errors.New("steth token not found")
	}

	stethBalance, err := tokenutil.BalanceOf(ctx, backend, steth, solver)
	if err != nil {
		return errors.Wrap(err, "steth balance")
	}

	if bi.IsZero(stethBalance) { // Nothing to wrap.
		return nil
	}

	wsteth, ok := tokens.ByAsset(evmchain.IDEthereum, tokens.WSTETH)
	if !ok {
		return errors.New("wsteth token not found")
	}

	wstethContract, err := NewWSTETH(wsteth.Address, backend)
	if err != nil {
		return errors.Wrap(err, "wsteth contract")
	}

	stethContract, err := bindings.NewIERC20(steth.Address, backend)
	if err != nil {
		return errors.Wrap(err, "steth contract")
	}

	txOpts, err := backend.BindOpts(ctx, solver)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	// Approve wsteth to wrap
	tx, err := stethContract.Approve(txOpts, wsteth.Address, stethBalance)
	if err != nil {
		return errors.Wrap(err, "approve wsteth")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	// Wrap stETH to wstETH.
	tx, err = wstethContract.Wrap(txOpts, stethBalance)
	if err != nil {
		return errors.Wrap(err, "wrap steth")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Wrapped stETH to wstETH", "tx", receipt.TxHash, "amount", steth.FormatAmt(stethBalance))

	return nil
}
