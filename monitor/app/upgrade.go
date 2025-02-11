package monitor

import (
	"context"
	"time"

	haloapp "github.com/omni-network/omni/halo/app"
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
			planned, ok, err := cprov.CurrentPlannedPlan(ctx)
			if err != nil {
				log.Warn(ctx, "Failed fetching planned upgrade (will retry)", err)
				continue
			} else if !ok {
				planned = utypes.Plan{
					Name:   "none",
					Height: 0,
				}
			}

			plannedUpgradeGauge.Reset()
			plannedUpgradeGauge.WithLabelValues(planned.Name).Set(float64(planned.Height))

			applied := utypes.Plan{
				Name:   "none",
				Height: 0,
			}
			for _, upgrade := range haloapp.AllUpgrades() {
				p, ok, err := cprov.AppliedPlan(ctx, upgrade)
				if err != nil {
					log.Warn(ctx, "Failed fetching applied upgrade (will retry)", err, "name", upgrade)
					continue
				} else if !ok || p.Height < applied.Height {
					continue
				}

				applied = p // Update last applied
			}

			appliedUpgradeGauge.Reset()
			appliedUpgradeGauge.WithLabelValues(applied.Name).Set(float64(applied.Height))
		}
	}
}
