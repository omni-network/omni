//nolint:unused // WIP
package rebalance

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/uniswap"
	"github.com/omni-network/omni/lib/usdt0"
	"github.com/omni-network/omni/solver/fundthresh"

	"github.com/ethereum/go-ethereum/common"
)

var (
	hypUSDT0 = mustToken(evmchain.IDHyperEVM, tokens.USDT0)
	ethUSDT  = mustToken(evmchain.IDEthereum, tokens.USDT)
	ethUSDC  = mustToken(evmchain.IDEthereum, tokens.USDC)
)

// rebalanceHyperEVMForever starts rebalancing HyperVM forever.
// HyperVM rebalancing is restricted to USDT0 refills from Ethereum -> HyperVM.
func rebalanceHyperEVMForever(
	ctx context.Context,
	interval time.Duration,
	backends ethbackend.Backends,
	solver common.Address,
) {
	_, err := backends.Backend(evmchain.IDHyperEVM)
	if err != nil {
		log.Warn(ctx, "HyperVM backend not available, skipping rebalancing", err)
		return
	}

	ctx = log.WithCtx(ctx, "step", "hypervm")

	for {
		start := time.Now()
		elapsed := time.Since(start)

		err := rebalanceHyperEVMOnce(ctx, backends, solver)
		if err != nil {
			log.Warn(ctx, "Failed to rebalance HyperVM (will retry)", err)
		}

		// Sleep for the remaining time in the interval, if any.
		if elapsed < interval {
			time.Sleep(interval - elapsed)
		}
	}
}

// rebalanceHyperEVMOnce refills HyperEVM USDT0 from Ethereum USDT, swapping from USDC if needed.
func rebalanceHyperEVMOnce(
	ctx context.Context,
	backends ethbackend.Backends,
	solver common.Address,
) error {
	ethBackend, err := backends.Backend(evmchain.IDEthereum)
	if err != nil {
		return errors.New("ethereum backend")
	}

	hypBackend, err := backends.Backend(evmchain.IDHyperEVM)
	if err != nil {
		return errors.New("hypervm backend")
	}

	//  Check USDT0 deficit on HyperEVM
	deficitUSDT0, err := GetDeficit(ctx, hypBackend, hypUSDT0, solver)
	if err != nil {
		return errors.Wrap(err, "get deficit usdt0")
	}

	// No deficit, return
	if bi.LTE(deficitUSDT0, bi.Zero()) {
		return nil
	}

	// Check if we have enough surplus USDT on Ethereum
	ethUSDTThresh := fundthresh.Get(ethUSDT)
	ethUSDTBalance, err := tokenutil.BalanceOf(ctx, ethBackend, ethUSDT, solver)
	if err != nil {
		return errors.Wrap(err, "get usdt balance")
	}

	surplusUSDT := bi.Sub(ethUSDTBalance, ethUSDTThresh.Surplus())

	// If we have enough, send USDT right to HyperEVM
	if bi.GTE(surplusUSDT, deficitUSDT0) {
		return sendUSDTToHyperEVM(ctx, ethBackend, solver, deficitUSDT0)
	}

	// If we don't, check if we have USDC surplus to swap
	neededUSDT := bi.Sub(deficitUSDT0, surplusUSDT)

	surplusUSDC, err := GetSurplus(ctx, ethBackend, ethUSDC, solver)
	if err != nil {
		return errors.Wrap(err, "get surplus usdc")
	}

	// Cap to available
	toSwap := bi.Sub(surplusUSDC, neededUSDT)
	if bi.LT(surplusUSDC, toSwap) {
		toSwap = surplusUSDC
	}

	// Limit to min swap
	if bi.LT(toSwap, fundthresh.Get(ethUSDC).MinSwap()) {
		return nil
	}

	// Swap to USDT, send
	_, err = uniswap.SwapFromUSDC(ctx, ethBackend, solver, ethUSDT, neededUSDT)
	if err != nil {
		return errors.Wrap(err, "swap usdc to usdt")
	}

	return sendUSDTToHyperEVM(ctx, ethBackend, solver, deficitUSDT0)
}

// sendUSDTToHyperEVM sends USDT from Ethereum to HyperEVM USDT0.
func sendUSDTToHyperEVM(
	ctx context.Context,
	ethBackend *ethbackend.Backend,
	solver common.Address,
	amount *big.Int,
) error {
	const maxSend = 10000 // 10k USDT
	toSend := amount
	if bi.GT(amount, bi.Dec6(maxSend)) {
		toSend = bi.Dec6(maxSend)
	}

	const minSend = 100 // 100 USDT
	if bi.LT(toSend, bi.Dec6(minSend)) {
		return nil
	}

	_, err := usdt0.Send(
		ctx,
		ethBackend,
		solver,
		evmchain.IDEthereum,
		evmchain.IDHyperEVM,
		toSend,
		nil, // TODO: add db
	)
	if err != nil {
		return errors.Wrap(err, "deposit usdt0")
	}

	return nil
}
