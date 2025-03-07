package staking

import (
	"context"
	"sort"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/params"
)

func MonitorForever(ctx context.Context, cprov cchain.Provider) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			allDelegations, err := queryutil.AllDelegations(ctx, cprov)
			if err != nil {
				log.Warn(ctx, "Failed to query all delegations", err)
				continue
			}

			instrStakeSizes(allDelegations)
			instrEffRewards(ctx, cprov, allDelegations)
		}
	}
}

// instrEffRewards instruments effective staking rewards.
func instrEffRewards(ctx context.Context, cprov cchain.Provider, allDelegations []queryutil.DelegationBalance) {
	delegations := allDelegations
	// Since we have no validator commissions, we can use just a couple of random delegations to estimate rewards.
	// Once we have validator commissions, this code needs to be removed.
	if len(allDelegations) > 4 {
		delegations = allDelegations[:4]
	}

	// Collect data during multiple blocks.
	const blocks = uint64(30)
	rewards, ok, err := queryutil.AvgRewardsRate(ctx, cprov, delegations, blocks)
	if err != nil {
		log.Warn(ctx, "Failed to get rewards rate (will retry)", err)
		return
	}

	if !ok {
		return
	}

	rewardsF64, err := rewards.Float64()
	if err != nil {
		log.Warn(ctx, "Failed to convert rewards rate to float64 [BUG]", err)
		return
	}

	rewardsAvg.Set(rewardsF64)
}

// instrStakeSizes delegations instruments delegations data.
func instrStakeSizes(allDelegations []queryutil.DelegationBalance) {
	delegatorsTotal := float64(len(allDelegations))
	delegatorsCount.Set(delegatorsTotal)

	var totalStake float64
	var stakes []float64
	for _, del := range allDelegations {
		stake := float64(del.Balance.Amount.Uint64())
		totalStake += stake
		stakes = append(stakes, stake)
	}

	avgStakeWei := totalStake / delegatorsTotal
	stakeAvg.Set(toGwei(avgStakeWei))

	l := len(stakes)
	if l == 0 {
		return
	}
	sort.Slice(stakes, func(i, j int) bool {
		return stakes[i] < stakes[j]
	})
	medianVal := stakes[l/2+l%2]
	stakeMedian.Set(toGwei(medianVal))
}

func toGwei(wei float64) float64 {
	return wei / params.GWei
}
