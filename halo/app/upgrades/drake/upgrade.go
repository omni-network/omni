// Package drake defines the third Omni consensus chain upgrade, named after
// Sir Francis Drake, an English explorer best known for making the second
// circumnavigation of the world in a single expedition between 1577 and 1580.
//
// It includes:
// - Unstaking of funds,
// - Processing of rewards withdrawals.
package drake

import (
	"context"
	"encoding/json"

	"github.com/omni-network/omni/lib/log"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const UpgradeName = "3_drake"

func StoreUpgrades(_ context.Context) *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}

func CreateUpgradeHandler(mm *module.Manager, configurator module.Configurator) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		log.Info(ctx, "Running 3_drake upgrade handler")

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func GenesisState(codec.JSONCodec) (map[string]json.RawMessage, error) {
	return nil, nil //nolint:nilnil // map is for reading only
}
