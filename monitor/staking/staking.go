package staking

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/params"
)

// maxDelegationsForRewardsEstimation is the max number of random rewards we track across multiple blocks
// to estimate the average effective staking rewards.
const maxDelegationsForRewardsEstimation = 4

func MonitorForever(ctx context.Context, cprov cchain.Provider) {
	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			allDelegations, err := queryutil.AllDelegations(ctx, cprov)
			if err != nil {
				log.Warn(ctx, "Failed to query all delegations", err)
				continue
			}

			instrStakeSizes(allDelegations)

			if err := instrEffRewards(ctx, cprov, allDelegations); err != nil {
				log.Warn(ctx, "Effective rewards intrumentation failed", err)
			}

			timer.Reset(time.Hour)
		}
	}
}

// instrEffRewards instruments effective staking rewards.
func instrEffRewards(ctx context.Context, cprov cchain.Provider, allDelegations []queryutil.DelegationBalance) error {
	delegations := allDelegations
	// Since we have no validator commissions, we can use just a couple of random delegations to estimate rewards.
	// Once we have validator commissions, this code needs to be removed.
	if len(allDelegations) > maxDelegationsForRewardsEstimation {
		delegations = allDelegations[:maxDelegationsForRewardsEstimation]
	}

	// Collect data during multiple blocks.
	const blocks = uint64(30)
	rewards, ok, err := queryutil.AvgRewardsRate(ctx, cprov, delegations, blocks)
	if err != nil {
		return errors.Wrap(err, "avg rewards")
	}

	if !ok {
		return nil
	}

	rewardsF64, err := rewards.Float64()
	if err != nil {
		return errors.Wrap(err, "rewards to float64 conversion")
	}

	rewardsAvg.Set(rewardsF64)

	return nil
}

// instrStakeSizes delegations instruments delegations data.
func instrStakeSizes(allDelegations []queryutil.DelegationBalance) {
	delegatorsTotal := float64(len(allDelegations))
	delegatorsCount.Set(delegatorsTotal)

	totalStake := new(big.Int)
	var stakes []*big.Int
	for _, del := range allDelegations {
		stake := del.Balance
		totalStake = new(big.Int).Add(totalStake, stake)
		stakes = append(stakes, stake)
	}

	avgStakeWei := new(big.Int).Quo(totalStake, big.NewInt(int64(len(allDelegations))))
	stakeAvg.Set(toEtherF64(avgStakeWei))

	l := len(stakes)
	if l == 0 {
		return
	}
	sort.Slice(stakes, func(i, j int) bool {
		return stakes[i].Cmp(stakes[j]) < 0
	})
	medianVal := stakes[l/2+l%2]
	stakeMedian.Set(toEtherF64(medianVal))
}

func toEtherF64(wei *big.Int) float64 {
	f64, _ := new(big.Int).Div(wei, big.NewInt(params.Ether)).Float64()
	return f64
}
