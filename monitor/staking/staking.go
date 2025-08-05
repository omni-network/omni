package staking

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
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
			timer.Reset(time.Hour)

			allDelegations, err := queryutil.AllDelegations(ctx, cprov)
			if err != nil {
				log.Warn(ctx, "Failed to query all delegations (will retry)", err)
				continue
			}

			instrStakeSizes(allDelegations)

			if err := instrEffRewards(ctx, cprov, allDelegations); err != nil {
				log.Warn(ctx, "Effective rewards instrumentation failed (will retry)", err)
			}
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
	rewards, err := queryutil.AvgRewardsRate(ctx, cprov, delegations, blocks)
	if err != nil {
		return errors.Wrap(err, "avg rewards")
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

	totalStake := bi.Zero()
	var stakes []*big.Int
	for _, del := range allDelegations {
		stake := del.Balance
		totalStake = bi.Add(totalStake, stake)
		stakes = append(stakes, stake)
	}

	avgStakeWei := bi.DivRaw(totalStake, len(allDelegations))
	stakeAvg.Set(bi.ToEtherF64(avgStakeWei))

	l := len(stakes)
	if l < 2 {
		return
	}
	sort.Slice(stakes, func(i, j int) bool {
		return bi.LT(stakes[i], stakes[j])
	})
	medianVal := stakes[l/2+l%2]
	stakeMedian.Set(bi.ToEtherF64(medianVal))
}
