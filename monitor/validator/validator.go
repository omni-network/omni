package validator

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MonitorForever(ctx context.Context, cprov cchain.Provider) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorOnce(ctx, cprov, sampleValidator)
			if err != nil {
				log.Warn(ctx, "Monitoring validator failed (will retry)", err)
			}
		}
	}
}

func monitorOnce(ctx context.Context, cprov cchain.Provider, sampleFunc func(sample)) error {
	signings, err := cprov.SDKSigningInfos(ctx)
	if err != nil {
		return err
	}

	infos, err := infosByConsAddr(signings)
	if err != nil {
		return err
	}

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

		rewards, ok, err := cprov.SDKRewards(ctx, opAddr)
		if err != nil {
			return err
		} else if !ok {
			return errors.New("missing rewards for validator")
		}

		consAddrEth, err := val.ConsensusEthAddr()
		if err != nil {
			return err
		}

		consAddrCmt, err := val.ConsensusCmtAddr()
		if err != nil {
			return err
		}

		info, ok := infos[consAddrCmt.String()]
		if !ok {
			return errors.New("missing signing info")
		}

		sampleFunc(sample{
			ConsensusEthAddr: consAddrEth,
			ConsensusCmtAddr: consAddrCmt,
			OperatorEthAddr:  opAddr,
			Power:            val.PotentialConsensusPower(sdk.DefaultPowerReduction),
			Jailed:           val.IsJailed(),
			Bonded:           val.IsBonded(),
			Tombstoned:       info.Tombstoned,
			Uptime:           info.Uptime,
			Rewards:          rewards,
		})
	}

	return nil
}

func infosByConsAddr(signings []cchain.SDKSigningInfo) (map[string]cchain.SDKSigningInfo, error) {
	resp := make(map[string]cchain.SDKSigningInfo)
	for _, signing := range signings {
		addr, err := signing.ConsensusCmtAddr()
		if err != nil {
			return nil, err
		}
		resp[addr.String()] = signing
	}

	return resp, nil
}

func boolToFloat(b bool) float64 {
	if !b {
		return 0
	}

	return 1
}
