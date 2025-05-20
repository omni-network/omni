package rebalance

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/mantle"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/solver/fundthresh"

	"github.com/ethereum/go-ethereum/common"
)

var (
	mantleUSDC   = mustToken(evmchain.IDMantle, tokens.USDC)
	ethereumUSDC = mustToken(evmchain.IDEthereum, tokens.USDC)
)

// rebalanceForever starts rebalancing mantle forever.
// Mantle rebalancing is restricted to USDC refills from L1 -> L2.
func rebalanceMantleForever(
	ctx context.Context,
	interval time.Duration,
	backends ethbackend.Backends,
	solver common.Address,
) {
	_, err := backends.Backend(evmchain.IDMantle)
	if err != nil {
		log.Warn(ctx, "Mantle backend not available, skipping rebalancing", err)
		return
	}

	ctx = log.WithCtx(ctx, "step", "mantle")

	for {
		start := time.Now()
		elapsed := time.Since(start)

		err := rebalanceMantleOnce(ctx, backends, solver)
		if err != nil {
			log.Warn(ctx, "Failed to rebalance mantle (will retry)", err)
		}

		// Sleep for the remaining time in the interval, if any.
		if elapsed < interval {
			time.Sleep(interval - elapsed)
		}
	}
}

// rebalanceMantleOnce rebalances mantle from L1 USDC balance.
// It does not require a surpluse on L1. Instead, Mantle USDC requirements are
// baked into L1 target balance.
func rebalanceMantleOnce(
	ctx context.Context,
	backends ethbackend.Backends,
	solver common.Address,
) error {
	l1USDC := ethereumUSDC
	l2USDC := mantleUSDC

	l1Backend, err := backends.Backend(evmchain.IDEthereum)
	if err != nil {
		return errors.New("ethereum backend")
	}

	l2Backend, err := backends.Backend(evmchain.IDMantle)
	if err != nil {
		return errors.New("mantle backend")
	}

	l2Thresh := fundthresh.Get(l2USDC)

	l1Balance, err := tokenutil.BalanceOf(ctx, l1Backend, l1USDC, solver)
	if err != nil {
		return errors.Wrap(err, "l1 balance")
	}

	l2Balance, err := tokenutil.BalanceOf(ctx, l2Backend, l2USDC, solver)
	if err != nil {
		return errors.Wrap(err, "l2 balance")
	}

	// L2 balance > target, do nothing.
	if bi.LT(l2Thresh.Target(), l2Balance) {
		return nil
	}

	deficit := bi.Sub(l2Thresh.Target(), l2Balance)

	// L1 balance < deficit, error and warn.
	if bi.LT(l1Balance, deficit) {
		return errors.New("deficit > l1 balance")
	}

	// Protect against over-spending L1 USDC - warn if defict more than half L1 balance.
	if bi.LT(bi.Div(l1Balance, bi.N(2)), deficit) {
		return errors.New("deficit > half of l1 balance")
	}

	_, err = mantle.DepositUSDC(
		ctx,
		l1Backend,
		solver,
		deficit,
	)
	if err != nil {
		return errors.Wrap(err, "deposit usdc")
	}

	return nil
}
