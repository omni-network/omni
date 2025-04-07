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

	"github.com/omni-network/omni/halo/app/upgrades/magellan"
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

func StoreUpgrades(_ context.Context) *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	staking skeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		log.Info(ctx, "Running 3_drake upgrade handler")

		p, err := staking.GetParams(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get staking params")
		}

		p.UnbondingTime = 0

		if err := staking.SetParams(ctx, p); err != nil {
			return nil, errors.Wrap(err, "set staking params")
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

// Staking params overwrites the default params with 30 validators and no unbonding time.
func StakingParams() stypes.Params {
	p := stypes.DefaultParams()
	p.MaxValidators = 30
	p.UnbondingTime = 0

	return p
}

// GenesisState creates a new genesis state. This state will be used on networks
// defining `3_drake` as `ephemeral_genesis` in their manifests.
func GenesisState(cdc codec.JSONCodec) (map[string]json.RawMessage, error) {
	genesis, err := magellan.GenesisState(cdc)
	if err != nil {
		return nil, errors.Wrap(err, "magellan genesis state")
	}

	stakingGenesis := stypes.DefaultGenesisState()
	stakingGenesis.Params = StakingParams()

	data, err := cdc.MarshalJSON(stakingGenesis)
	if err != nil {
		return nil, errors.Wrap(err, "marshal staking genesis")
	}

	genesis[stypes.ModuleName] = data

	return genesis, nil
}
