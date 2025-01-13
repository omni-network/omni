package loadgen

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/portal"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	deadAddr = "0x000000000000000000000000000000000000dead"
)

type xCallConfig struct {
	NetworkID   netconf.ID
	XCallerAddr common.Address
	Period      time.Duration
	Backends    ethbackend.Backends
	Chains      []netconf.Chain
}

func xCallForever(ctx context.Context, cfg xCallConfig) {
	log.Info(ctx, "Starting periodic xcalls", "period", cfg.Period)

	nextPeriod := func() time.Duration {
		jitter := time.Duration(float64(cfg.Period) * rand.Float64() * loadgenJitter) //nolint:gosec // Weak random ok for load tests.
		return cfg.Period + jitter
	}

	// timer will tick immediately
	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := xCall(ctx, cfg); err != nil {
				log.Warn(ctx, "Failed to xcall (will retry)", err)
			}
			timer.Reset(nextPeriod())
		}
	}
}

func xCall(ctx context.Context, cfg xCallConfig) error {
	fromChain, dstChain := getChainPair(cfg.Chains)
	backend, err := cfg.Backends.Backend(fromChain.ID)
	if err != nil {
		return err
	}

	fromPortal, err := getPortal(ctx, cfg.NetworkID, fromChain.ID, cfg.Backends)
	if err != nil {
		return err
	}

	var nilData []byte
	to := common.HexToAddress(deadAddr)
	fee, err := fromPortal.FeeFor(&bind.CallOpts{Context: ctx}, dstChain.ID, nilData, portal.XMsgMinGasLimit)
	if err != nil {
		return errors.Wrap(err, "feeFor",
			"src_chain", fromChain.ID,
			"dst_chain_id", dstChain.ID,
		)
	}

	txOpts, err := backend.BindOpts(ctx, cfg.XCallerAddr)
	if err != nil {
		return errors.Wrap(err, "bindOpts")
	}

	txOpts.Value = fee
	tx, err := fromPortal.Xcall(txOpts, dstChain.ID, uint8(xchain.ConfLatest), to, nilData, portal.XMsgMinGasLimit)
	if err != nil {
		return errors.Wrap(err, "xcall",
			"src_chain", fromChain.ID,
			"dst_chain_id", dstChain.ID,
		)
	}
	log.Debug(ctx, fmt.Sprintf("xcall made %s -> %s %s", fromChain.Name, dstChain.Name, tx.Hash()))

	return nil
}

func getPortal(ctx context.Context, networkID netconf.ID, chainID uint64, backends ethbackend.Backends) (*bindings.OmniPortal, error) {
	backend, err := backends.Backend(chainID)
	if err != nil {
		return nil, err
	}

	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return nil, err
	}

	contract, err := bindings.NewOmniPortal(addrs.Portal, backend)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func getChainPair(chains []netconf.Chain) (netconf.Chain, netconf.Chain) {
	i := rand.IntN(len(chains)) //nolint:gosec // Weak random ok for load tests.
	j := rand.IntN(len(chains)) //nolint:gosec // Weak random ok for load tests.
	for i == j {
		j = rand.IntN(len(chains)) //nolint:gosec // Weak random ok for load tests.
	}

	return chains[i], chains[j]
}
