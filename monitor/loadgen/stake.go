package loadgen

import (
	"context"
	"math/big"
	"math/rand/v2"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

const selfDelegateJitter = 0.2 // 20% jitter

func selfDelegateForever(ctx context.Context, contract *bindings.Staking, backend *ethbackend.Backend, validator common.Address, period time.Duration) {
	log.Info(ctx, "Starting periodic self-delegation", "validator", validator.Hex(), "period", period)

	nextPeriod := func() time.Duration {
		jitter := time.Duration(float64(period) * rand.Float64() * selfDelegateJitter) //nolint:gosec // Weak random ok for load tests.
		return period + jitter
	}

	timer := time.NewTimer(nextPeriod())
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := selfDelegateOnce(ctx, contract, backend, validator); err != nil {
				log.Warn(ctx, "Failed to self-delegate (will retry)", err)
			}
			timer.Reset(nextPeriod())
		}
	}
}

func selfDelegateOnce(ctx context.Context, contract *bindings.Staking, backend *ethbackend.Backend, validator common.Address) error {
	backoff := expbackoff.New(ctx)
	for {
		ethBalance, err := backend.EtherBalanceAt(ctx, validator)
		if err != nil {
			return err
		} else if ethBalance < 1 {
			log.Info(ctx, "Waiting for validator to be funded", "balance", ethBalance, "validator", validator.Hex())
			backoff()

			continue
		}

		break // Continue funding below
	}

	txOpts, err := backend.BindOpts(ctx, validator)
	if err != nil {
		return err
	}
	txOpts.Value = big.NewInt(params.Ether) // 1 ETH (in wei)

	tx, err := contract.Delegate(txOpts, validator, validator)
	if err != nil {
		return errors.Wrap(err, "deposit")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	log.Info(ctx, "Deposited validator self-delegation",
		"height", rec.BlockNumber,
		"validator", validator,
	)

	return nil
}
