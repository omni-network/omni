package monitor

import (
	"context"
	"time"

	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/log"
)

func monitorStakingForever(ctx context.Context, cprov cchain.Provider) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_, ok, err := cprov.AppliedPlan(ctx, magellan2.UpgradeName)
			if err != nil {
				log.Warn(ctx, "Failed to get applied upgrade", err)
				continue
			} else if !ok {
				log.Debug(ctx, "Not monitoring inflation, since magellan upgrade not applied yet")
				continue
			}

			effectiveStakingRewards(ctx, cprov)
		}
	}
}

// effectiveStakingRewards instruments effective staking rewards.
func effectiveStakingRewards(ctx context.Context, cprov cchain.Provider) {
	// Collect data during multiple blocks.
	blocks := uint64(30)
	inflation, _, err := queryutil.AvgInflationRate(ctx, cprov, blocks)
	if err != nil {
		log.Warn(ctx, "Failed to get inflation rate (will retry)", err)
		return
	}

	inflationF64, err := inflation.Float64()
	if err != nil {
		log.Warn(ctx, "Failed to convert inflation rate to float64 [BUG]", err)
		return
	}

	inflationAvg.Set(inflationF64)
}
