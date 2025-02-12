package monitor

import (
	"context"
	"time"

	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/log"
)

func monitorInflationForever(ctx context.Context, cprov cchain.Provider) {
	ticker := time.NewTicker(time.Minute)
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

			inflation, _, err := queryutil.AvgInflationRate(ctx, cprov, 3)
			if err != nil {
				log.Warn(ctx, "Failed to get inflation rate (will retry)", err)
				continue
			}

			inflationF64, err := inflation.Float64()
			if err != nil {
				log.Warn(ctx, "Failed to convert inflation rate to float64 [BUG]", err)
				continue
			}

			inflationAvg.Set(inflationF64)
		}
	}
}
