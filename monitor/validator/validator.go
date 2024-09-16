package validator

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	cosmosk1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
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
	vals, err := cprov.Validators(ctx)
	if err != nil {
		return errors.Wrap(err, "query validators")
	}

	// Reset existing time-series since validator may be removed
	powerGauge.Reset()
	jailedGauge.Reset()
	bondedGauge.Reset()

	for _, val := range vals {
		pk := new(cosmosk1.PubKey)
		err := proto.Unmarshal(val.ConsensusPubkey.Value, pk)
		if err != nil {
			return errors.Wrap(err, "unmarshal consensus pubkey")
		}

		pubkey, err := crypto.DecompressPubkey(pk.Bytes())
		if err != nil {
			return errors.Wrap(err, "decompress pubkey")
		}

		opAddr, err := sdk.ValAddressFromBech32(val.OperatorAddress)
		if err != nil {
			return errors.Wrap(err, "parse operator address")
		}

		opAddrEth := common.BytesToAddress(opAddr)
		consAddrEth := crypto.PubkeyToAddress(*pubkey)
		consAddrCmt := pk.Address()
		power := val.ConsensusPower(sdk.DefaultPowerReduction)
		jailed := val.IsJailed()
		bonded := val.IsBonded()

		powerGauge.WithLabelValues(consAddrEth.String(), consAddrCmt.String(), opAddrEth.String()).Set(float64(power))
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
