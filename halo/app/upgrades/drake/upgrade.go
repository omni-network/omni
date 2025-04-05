// Package drake defines the third Omni consensus chain upgrade, named after
// Sir Francis Drake, an English explorer and privateer best known for making the
// second circumnavigation of the world in a single expedition between 1577 and 1580.
//
// It includes:
// - Unstaking of funds,
// - Processing of rewards withdrawals,
// - Setting the unbonding time to 0.
package drake

import (
	"context"
	"encoding/json"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const UpgradeName = "3_drake"

func StoreUpgrades(_ context.Context) *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{},
	}
}

var (
	// StakingParams are genesis params with UnbondingTime set to 0.
	StakingParams = stypes.Params{
		UnbondingTime:     0 * time.Second,
		MaxValidators:     30,
		MaxEntries:        7,
		HistoricalEntries: 10_000,
		BondDenom:         "stake",
		MinCommissionRate: math.LegacyNewDec(0),
	}
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	staking skeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		log.Info(ctx, "Running 3_drake upgrade handler")

		if err := staking.SetParams(ctx, StakingParams); err != nil {
			return nil, errors.Wrap(err, "set staking params")
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func GenesisState(codec.JSONCodec) (map[string]json.RawMessage, error) {
	return map[string]json.RawMessage{}, nil
}
