package loadgen

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/connect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	dead = "0x000000000000000000000000000000000000dead"
)

type XCallerConfig struct {
	Enabled  bool
	ChainIDs map[string]string
}

func xCallForever(ctx context.Context, xCallerAddr common.Address, period time.Duration, chainIDs map[string]string, backends ethbackend.Backends, connector connect.Connector) {
	log.Info(ctx, "Starting periodic xcalls", "period", period)

	nextPeriod := func() time.Duration {
		jitter := time.Duration(float64(period) * rand.Float64() * loadgenJitter) //nolint:gosec // Weak random ok for load tests.
		return period + jitter
	}

	timer := time.NewTimer(nextPeriod())
	defer timer.Stop()

	// tick immediately
	if err := xCallOnce(ctx, xCallerAddr, chainIDs, backends, connector); err != nil {
		log.Warn(ctx, "Failed to xcall (will retry)", err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := xCallOnce(ctx, xCallerAddr, chainIDs, backends, connector); err != nil {
				log.Warn(ctx, "Failed to xcall (will retry)", err)
			}
			timer.Reset(nextPeriod())
		}
	}
}

func xCallOnce(ctx context.Context, xCallerAddr common.Address, chainIDs map[string]string, backends ethbackend.Backends, connector connect.Connector) error {
	for from, to := range chainIDs {
		fromChain, ok := evmchain.MetadataByName(from)
		if !ok {
			return errors.New("unknown source chain name", from)
		}

		dstChain, ok := evmchain.MetadataByName(to)
		if !ok {
			return errors.New("unknown destination chain name", from)
		}

		backend, err := backends.Backend(fromChain.ChainID)
		if err != nil {
			return err
		}

		fromPortal, err := connector.GetPortal(fromChain.ChainID, backend)
		if err != nil {
			return err
		}

		var data []byte
		to := common.HexToAddress(dead)
		gasLimit := uint64(100_000)
		fee, err := fromPortal.FeeFor(&bind.CallOpts{}, dstChain.ChainID, data, gasLimit)
		if err != nil {
			return errors.Wrap(err, "feeFor",
				"src_chain", fromChain.ChainID,
				"dst_chain_id", dstChain.ChainID,
			)
		}

		txOpts, err := backend.BindOpts(ctx, xCallerAddr)
		if err != nil {
			return errors.Wrap(err, "bindOpts")
		}

		txOpts.Value = fee
		tx, err := fromPortal.Xcall(txOpts, dstChain.ChainID, uint8(xchain.ConfFinalized), to, data, gasLimit)
		if err != nil {
			return errors.Wrap(err, "xcall",
				"src_chain", fromChain.ChainID,
				"dst_chain_id", dstChain.ChainID,
			)
		}

		log.Debug(ctx, fmt.Sprintf("xcall made %d -> %d %s", fromChain.ChainID, dstChain.ChainID, tx.To()))
	}

	return nil
}
