//nolint:revive,staticcheck,unparam // WIP
package rebalance

import (
	"context"
	"time"

	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// rebalanceForever starts rebalancing loops for each chain in the network.
func rebalanceForever(
	ctx context.Context,
	cfg Config,
	db *cctpdb.DB,
	network netconf.Network,
	backends ethbackend.Backends,
	solver common.Address,
) {
	for _, chain := range network.EVMChains() {
		ctx := log.WithCtx(ctx, "chain", evmchain.Name(chain.ID))

		go swapSurplusForever(ctx, cfg, db, network, backends, solver, chain.ID)
		go sendSurplusForever(ctx, cfg, db, network, backends, solver, chain.ID)
		go fillDeficitForever(ctx, cfg, db, network, backends, solver, chain.ID)
	}
}

// swapSurplusForever swaps surplus tokens to USDC on `chainID` forever.
func swapSurplusForever(
	ctx context.Context,
	cfg Config,
	db *cctpdb.DB,
	network netconf.Network,
	backends ethbackend.Backends,
	solver common.Address,
	chainID uint64,
) {
	ticker := time.NewTicker(0)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Reset(cfg.Interval)

			do := func() error {
				return swapSurplusOnce(ctx, db, network.ID, backends, chainID, solver)
			}

			if err := expbackoff.Retry(ctx, do); err != nil {
				log.Error(ctx, "Swap surplus failed", err)
			}
		}
	}
}

// sendSurplusForever sends surplus USDC on `chainID` to chains in deficit forever.
func sendSurplusForever(
	ctx context.Context,
	cfg Config,
	db *cctpdb.DB,
	network netconf.Network,
	backends ethbackend.Backends,
	solver common.Address,
	chainID uint64,
) {
	ticker := time.NewTicker(0)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Reset(cfg.Interval)

			do := func() error {
				return sendSurplusOnce(ctx, db, network.ID, backends, chainID, solver)
			}

			if err := expbackoff.Retry(ctx, do); err != nil {
				log.Error(ctx, "Send surplus failed", err)
			}
		}
	}
}

// fillDeficitForever fills token deficits from surplus USDC on `chainID` forever.
func fillDeficitForever(
	ctx context.Context,
	cfg Config,
	db *cctpdb.DB,
	network netconf.Network,
	backends ethbackend.Backends,
	solver common.Address,
	chainID uint64,
) {
	ticker := time.NewTicker(0)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Reset(cfg.Interval)

			do := func() error {
				return fillDeficitOnce(ctx, db, network.ID, backends, chainID, solver)
			}

			if err := expbackoff.Retry(ctx, do); err != nil {
				log.Error(ctx, "Fill deficit failed", err)
			}
		}
	}
}

// sendSurplusOnce sends surplus USDC on `chainID` to chains in deficit.
func sendSurplusOnce(
	ctx context.Context,
	db *cctpdb.DB,
	networkID netconf.ID,
	backends ethbackend.Backends,
	chainID uint64,
	solver common.Address,
) error {
	if chainID != evmchain.IDBase {
		// Only sending surplus wstETH on Base.
	}

	return nil
}

// swapSurplusOnce swaps surplus tokens to USDC on `chainID`.
func swapSurplusOnce(
	ctx context.Context,
	db *cctpdb.DB,
	networkID netconf.ID,
	backends ethbackend.Backends,
	chainID uint64,
	solver common.Address,
) error {
	if chainID != evmchain.IDBase {
		// Only swapping surplus wstETH on Base.
	}

	// TODO
	return nil
}

// fillDeficitOnce fills token deficits from surplus USDC on `chainID`.
func fillDeficitOnce(
	ctx context.Context,
	db *cctpdb.DB,
	networkID netconf.ID,
	backends ethbackend.Backends,
	chainID uint64,
	solver common.Address,
) error {
	if chainID != evmchain.IDEthereum {
		// Only filling deficit wstETH on Ethereum.
		return nil
	}

	// TODO

	return nil
}
