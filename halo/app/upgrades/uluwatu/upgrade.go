// Package uluwatu defines the first omni consensus chain upgrade named after the iconic surf spot in Bali.
// It only includes attest module migration including constructor and logic changes.
// It doesn't include any store migrations.
package uluwatu

import (
	"context"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const UpgradeName = "1_uluwatu"

var StoreUpgrades storetypes.StoreUpgrades // Zero store upgrades

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
