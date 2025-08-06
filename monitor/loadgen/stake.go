//nolint:unused // Temporary
package loadgen

import (
	"context"
	"math/rand/v2"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

const loadgenJitter = 0.2 // 20% jitter

func delegateForever(ctx context.Context, contract *bindings.Staking, backend *ethbackend.Backend, delegator, validator common.Address, period time.Duration) {
	log.Info(ctx, "Starting periodic delegation", "delegator", delegator.Hex(), "validator", validator.Hex(), "period", period)

	nextPeriod := func() time.Duration {
		jitter := time.Duration(float64(period) * rand.Float64() * loadgenJitter) //nolint:gosec // Weak random ok for load tests.
		return period + jitter
	}

	timer := time.NewTimer(nextPeriod())
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := delegateOnce(ctx, contract, backend, delegator, validator); err != nil {
				log.Warn(ctx, "Failed to delegate (will retry)", err)
			}
			timer.Reset(nextPeriod())
		}
	}
}

func delegateOnce(ctx context.Context, contract *bindings.Staking, backend *ethbackend.Backend, delegator, validator common.Address) error {
	backoff := expbackoff.New(ctx)
	for {
		ethBalance, err := backend.EtherBalanceAt(ctx, delegator)
		if err != nil {
			return err
		} else if ethBalance < 1 {
			log.Info(ctx, "Waiting for delegator to be funded", "balance", ethBalance, "delegator", delegator.Hex())
			backoff()

			continue
		}

		break // Continue funding below
	}

	txOpts, err := backend.BindOpts(ctx, delegator)
	if err != nil {
		return err
	}
	txOpts.Value = bi.Ether(1) // 1 ETH (in wei)

	tx, err := contract.Delegate(txOpts, validator)
	if err != nil {
		return errors.Wrap(err, "deposit")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	log.Info(ctx, "Deposited delegation",
		"height", rec.BlockNumber,
		"delegator", delegator,
		"validator", validator,
	)

	return nil
}
