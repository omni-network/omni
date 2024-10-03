package monitor

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/log"

	utypes "cosmossdk.io/x/upgrade/types"
)

// monitorUpgradesForever blocks until the context is closed and
// periodically updates the planned upgrade gauge.
func monitorUpgradesForever(ctx context.Context, cprov cchain.Provider) {
	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			plan, ok, err := cprov.CurrentPlannedPlan(ctx)
			if err != nil {
				log.Warn(ctx, "Failed fetching planned upgrade (will retry)", err)
				continue
			} else if !ok {
				plan = utypes.Plan{
					Name:   "none",
					Height: 0,
				}
			}

			plannedUpgradeGauge.Reset()
			plannedUpgradeGauge.WithLabelValues(plan.Name).Set(float64(plan.Height))
		}
	}
}
