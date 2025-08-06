package loadgen

import (
	"context"
	"math/rand/v2"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
)

const (
	StakingMethodDelegate   = 0
	StakingMethodUndelegate = 1
)

func maybeDosForever(ctx context.Context, network netconf.ID, backend *ethbackend.Backend, delegator, validator common.Address, period time.Duration) {
	if network != netconf.Staging {
		return
	}

	log.Info(ctx, "Starting periodic staking dos", "delegator", delegator.Hex(), "validator", validator.Hex(), "period", period)

	nextPeriod := func() time.Duration {
		jitter := time.Duration(float64(period) * rand.Float64() * loadgenJitter) //nolint:gosec // Weak random ok for load tests.
		return period + jitter
	}

	timer := time.NewTimer(nextPeriod())
	defer timer.Stop()

	count := 100

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := dosOnce(ctx, backend, delegator, validator, count); err != nil {
				log.Warn(ctx, "Failed to submit DoS (will retry)", err)
			}
			timer.Reset(nextPeriod())
			count += 100
		}
	}
}

func dosOnce(ctx context.Context, backend *ethbackend.Backend, delegator, validator common.Address, count int) error {
	backoff := expbackoff.New(ctx)
	for {
		ethBalance, err := backend.EtherBalanceAt(ctx, delegator)
		if err != nil {
			return err
		} else if ethBalance < float64(count)*1.1 {
			log.Info(ctx, "Waiting for delegator to be funded", "balance", ethBalance, "delegator", delegator.Hex())
			backoff()

			continue
		}

		break // Continue funding below
	}

	proxyAddr, err := deployStakingProxy(ctx, backend, delegator)
	if err != nil {
		return errors.Wrap(err, "deploying staking proxy")
	}

	proxy, err := bindings.NewStakingProxy(proxyAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new staking proxy")
	}

	txOpts, err := backend.BindOpts(ctx, delegator)
	if err != nil {
		return err
	}
	txOpts.Value = bi.Ether(count) // 1 ETH (in wei)

	var calls []bindings.StakingProxyCall
	for i := 0; i < count; i++ {
		// Invalid undelegate message
		invalid := bi.N(1000)
		calls = append(calls, bindings.StakingProxyCall{
			Method:    StakingMethodUndelegate,
			Value:     bi.Ether(0.1), // Undelegate fee
			Validator: tutil.RandomAddress(),
			Amount:    invalid,
		},
			bindings.StakingProxyCall{
				Method:    StakingMethodDelegate,
				Value:     bi.Ether(1),
				Validator: validator,
			})
	}

	tx, err := proxy.Proxy(txOpts, calls)
	if err != nil {
		return errors.Wrap(err, "deposit")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	log.Info(ctx, "DoS staking event submitted",
		"height", rec.BlockNumber,
		"delegator", delegator,
		"validator", validator,
		"count", count,
	)

	return nil
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
