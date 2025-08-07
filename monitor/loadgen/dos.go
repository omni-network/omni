package loadgen

import (
	"context"
	"math/big"
	"math/rand/v2"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
)

func maybeDosForever(ctx context.Context, backend *ethbackend.Backend, delegator, validator common.Address, period time.Duration) {
	log.Info(ctx, "Starting periodic staking DoS", "delegator", delegator.Hex(), "validator", validator.Hex(), "period", period)

	nextPeriod := func() time.Duration {
		jitter := time.Duration(float64(period) * rand.Float64() * loadgenJitter) //nolint:gosec // Weak random ok for load tests.
		return period + jitter
	}

	timer := time.NewTimer(nextPeriod())
	defer timer.Stop()

	count := 100
	maxCount := 1500 // Gas limit of 30M exceeded if more.

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := dosOnce(ctx, backend, delegator, validator, count); err != nil {
				log.Warn(ctx, "Failed to submit DoS (will retry)", err, "count", count)
			} else if count < maxCount {
				count += 100
			}
			timer.Reset(nextPeriod())
		}
	}
}

func dosOnce(ctx context.Context, backend *ethbackend.Backend, delegator, validator common.Address, count int) error {
	delegateAmount := bi.Ether(evmredenom.Factor) // 75 $NOM (minimum delegation amount)

	txOpts, err := backend.BindOpts(ctx, delegator)
	if err != nil {
		return err
	}
	txOpts.Value = bi.MulRaw(delegateAmount, count)

	if err := awaitBalance(ctx, backend, delegator, txOpts.Value); err != nil {
		return err
	}

	proxyAddr, err := deployStakingProxy(ctx, backend, delegator)
	if err != nil {
		return errors.Wrap(err, "deploying staking proxy")
	}

	proxy, err := bindings.NewStakingProxy(proxyAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new staking proxy")
	}

	tx, err := proxy.DelegateN(txOpts, validator, delegateAmount, bi.N(count))
	if err != nil {
		return errors.Wrap(err, "deposit")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	log.Info(ctx, "DoS staking delegation submitted",
		"height", rec.BlockNumber,
		"delegator", delegator,
		"validator", validator,
		"count", count,
	)

	// N invalid Undelegate messages
	undelegateFee := bi.Ether(0.1) // Fee for undelegation
	txOpts, err = backend.BindOpts(ctx, delegator)
	if err != nil {
		return err
	}
	txOpts.Value = bi.MulRaw(undelegateFee, count)

	if err := awaitBalance(ctx, backend, delegator, txOpts.Value); err != nil {
		return err
	}

	tx, err = proxy.UndelegateN(txOpts, tutil.RandomAddress(), undelegateFee, delegateAmount, bi.N(count))
	if err != nil {
		return errors.Wrap(err, "undelegate")
	}
	rec, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	log.Info(ctx, "DoS staking undelegation submitted",
		"height", rec.BlockNumber,
		"delegator", delegator,
		"count", count,
	)

	return nil
}

func awaitBalance(ctx context.Context, backend *ethbackend.Backend, delegator common.Address, require *big.Int) error {
	// Await balance
	backoff := expbackoff.New(ctx)
	for {
		bal, err := backend.BalanceAt(ctx, delegator, nil)
		if err != nil {
			return err
		} else if bi.LT(bal, require) {
			log.Info(ctx, "Waiting for DoS delegator to be funded", "balance", bi.ToEtherF64(bal), "require", bi.ToEtherF64(require), "delegator", delegator.Hex())
			backoff()

			continue
		}

		return nil
	}
}

// deployStakingProxy deploys a proxy smart contract that simply batches arbitrary calls to Staking.sol.
func deployStakingProxy(
	ctx context.Context,
	backend *ethbackend.Backend,
	deployer common.Address,
) (common.Address, error) {
	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "binding opts")
	}

	address, tx, _, err := bindings.DeployStakingProxy(txOpts, backend, common.HexToAddress(predeploys.Staking))
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deployment")
	}
	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return common.Address{}, errors.Wrap(err, "mining")
	}

	return address, nil
}
