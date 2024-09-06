package validator

import (
	"context"
	"time"

	"github.com/omni-network/omni/halo/sdk"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func MonitorForever(ctx context.Context, cprov cchain.Provider) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorOnce(ctx, cprov)
			if err != nil {
				log.Warn(ctx, "Monitoring validator failed (will retry)", err)
			}
		}
	}
}

func monitorOnce(ctx context.Context, cprov cchain.Provider) error {
	vals, err := cprov.SDKValidators(ctx)
	if err != nil {
		return errors.Wrap(err, "query validators")
	}

	// Reset existing time-series since validator may be removed
	powerGauge.Reset()
	jailedGauge.Reset()
	bondedGauge.Reset()

	for _, val := range vals {
		opAddr, err := val.OperatorEthAddr()
		if err != nil {
			return err
		}

		consAddrEth, err := val.ConsensusEthAddr()
		if err != nil {
			return err
		}

		consAddrCmt, err := val.ConsensusCmtAddr()
		if err != nil {
			return err
		}

		power := val.ConsensusPower(sdk.DefaultPowerReduction)
		jailed := val.IsJailed()
		bonded := val.IsBonded()

		powerGauge.WithLabelValues(consAddrEth.String(), consAddrCmt.String(), opAddr.String()).Set(float64(power))
		jailedGauge.WithLabelValues(consAddrEth.String()).Set(boolToFloat(jailed))
		bondedGauge.WithLabelValues(consAddrEth.String()).Set(boolToFloat(bonded))
	}

	return nil
}

func boolToFloat(b bool) float64 {
	if !b {
		return 0
	}

	return 1
}
