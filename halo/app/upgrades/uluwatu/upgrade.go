// Package uluwatu defines the first omni consensus chain upgrade named after the iconic surf spot in Bali.
// It only includes attest module migration including constructor and logic changes.
// It doesn't include any store migrations.
package uluwatu

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	slkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	sltypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

const UpgradeName = "1_uluwatu"

func StoreUpgrades(context.Context) *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{} // Zero store upgrades
}

var (
	// SlashingParams as per https://github.com/omni-network/omni/issues/2018.
	SlashingParams = sltypes.Params{
		SignedBlocksWindow:      2000,                            // Roughly 40min
		MinSignedPerWindow:      math.LegacyNewDecWithPrec(5, 1), // Unchanged from default 5%
		DowntimeJailDuration:    12 * time.Hour,                  // Mimics post-valsync-buffer network upgrade.
		SlashFractionDoubleSign: math.LegacyNewDec(0),            // 0% since mainnet V1 has trusted operators.
		SlashFractionDowntime:   math.LegacyNewDec(0),            // 0% since mainnet V1 has trusted operators.
	}
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	slashing slkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if err := slashing.SetParams(ctx, SlashingParams); err != nil {
			return nil, errors.Wrap(err, "set slashing params")
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
