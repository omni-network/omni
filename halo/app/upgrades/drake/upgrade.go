// Package drake defines the third Omni consensus chain upgrade, named after
// Sir Francis Drake, an English explorer best known for making the second
// circumnavigation of the world in a single expedition between 1577 and 1580.
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

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const UpgradeName = "3_drake"

// UnbondingTime is the time of unbonding applied to all undelegated stakes.
const UnbondingTime = 1 * time.Second

// MaxValidators is the maximum of validators we keep across all networks.
const MaxValidators = uint32(30)

func StoreUpgrades(_ context.Context) *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	staking *skeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		log.Info(ctx, "Running 3_drake upgrade handler")

		p, err := staking.GetParams(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get staking params")
		}

		p.UnbondingTime = UnbondingTime

		if err := staking.SetParams(ctx, p); err != nil {
			return nil, errors.Wrap(err, "set staking params")
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

// GenesisState creates a new genesis state. This state will be used on networks
// defining `3_drake` as `ephemeral_genesis` in their manifests.
func GenesisState(cdc codec.JSONCodec) (map[string]json.RawMessage, error) {
	stakingGenesis := stypes.DefaultGenesisState()
	stakingGenesis.Params.UnbondingTime = UnbondingTime
	stakingGenesis.Params.MaxValidators = MaxValidators

	data, err := cdc.MarshalJSON(stakingGenesis)
	if err != nil {
		return nil, errors.Wrap(err, "marshal staking genesis")
	}

	return map[string]json.RawMessage{
		stypes.ModuleName: data,
	}, nil
}
