package avs

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
)

func startMonitoring(ctx context.Context, avs *bindings.OmniAVS) {
	go monitorForever(ctx, avs, "operators", monitorOperatorsOnce)
	go monitorForever(ctx, avs, "owner", monitorOwnerOnce)
	go monitorForever(ctx, avs, "allowlistEnabled", monitorAllowlistEnabledOnce)
	go monitorForever(ctx, avs, "paused", monitorPausedOnce)
	go monitorForever(ctx, avs, "minOperatorStake", monitorMinStakeOnce)
	go monitorForever(ctx, avs, "maxOperatorCount", monitorMaxOperatorsOnce)
	go monitorForever(ctx, avs, "strategyParams", monitorStrategyParamsOnce)
}

func monitorOperatorsOnce(ctx context.Context, avs *bindings.OmniAVS) error {
	operators, err := avs.Operators(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get operators")
	}

	// Reset all operator labeled metrics since some operators may have been removed.
	operatorStakeGuage.Reset()
	operatorDelegationsGuage.Reset()

	var total float64
	for _, operator := range operators {
		addr := operator.Addr.Hex()
		staked := weiToEth(operator.Staked)
		delegated := weiToEth(operator.Delegated)

		operatorStakeGuage.WithLabelValues(addr).Set(staked)
		operatorDelegationsGuage.WithLabelValues(addr).Set(delegated)

		total += delegated + staked
	}

	numOperatorsGuage.Set(float64(len(operators)))
	totalDelegationsGuage.Set(total)

	return nil
}

func monitorOwnerOnce(ctx context.Context, avs *bindings.OmniAVS) error {
	owner, err := avs.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get owner")
	}

	ownerGauge.WithLabelValues(owner.Hex()).Set(1)

	return nil
}

func monitorAllowlistEnabledOnce(ctx context.Context, avs *bindings.OmniAVS) error {
	enababled, err := avs.AllowlistEnabled(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get allowlist")
	}

	if enababled {
		allowlistEnabledGuage.Set(1)
	} else {
		allowlistEnabledGuage.Set(0)
	}

	return nil
}

func monitorPausedOnce(ctx context.Context, avs *bindings.OmniAVS) error {
	paused, err := avs.Paused(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get paused")
	}

	if paused {
		pausedGuage.Set(1)
	} else {
		pausedGuage.Set(0)
	}

	return nil
}

func monitorMinStakeOnce(ctx context.Context, avs *bindings.OmniAVS) error {
	stake, err := avs.MinOperatorStake(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get min stake")
	}

	minStakeGuage.Set(weiToEth(stake))

	return nil
}

func monitorMaxOperatorsOnce(ctx context.Context, avs *bindings.OmniAVS) error {
	max, err := avs.MaxOperatorCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get max operators")
	}

	maxOperatorsGuage.Set(float64(max))

	return nil
}

func monitorStrategyParamsOnce(ctx context.Context, avs *bindings.OmniAVS) error {
	params, err := avs.StrategyParams(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get strategy params")
	}

	strategyParamsGuage.Reset()
	for _, p := range params {
		start := p.Strategy.Hex()
		multiplier := p.Multiplier.Int64()
		strategyParamsGuage.WithLabelValues(start).Set(float64(multiplier))
	}

	return nil
}

type monitorOnce func(ctx context.Context, avs *bindings.OmniAVS) error

// monitorForever runs the given monitor function every 30 seconds until the context is canceled.
func monitorForever(ctx context.Context, avs *bindings.OmniAVS, name string, f monitorOnce) {
	interval := time.Minute * 1
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := func() error {
				ctx, cancel := context.WithTimeout(ctx, interval)
				defer cancel()

				return f(ctx, avs)
			}()

			if err != nil {
				log.Error(ctx, fmt.Sprintf("Monitoring AVS %s failed (will retry)", name), err)
			}
		}
	}
}

// weiToEth converts a wei amount to an ether amount.
func weiToEth(wei *big.Int) float64 {
	f, _ := wei.Float64()
	return f / params.Ether
}
