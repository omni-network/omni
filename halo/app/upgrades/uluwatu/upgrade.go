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
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	sltypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

const UpgradeName = "1_uluwatu"

var StoreUpgrades storetypes.StoreUpgrades // Zero store upgrades

// SlashingParams as per https://github.com/omni-network/omni/issues/2018.
var SlashingParams = sltypes.Params{
	SignedBlocksWindow:      2000,                            // Roughly 40min
	MinSignedPerWindow:      math.LegacyNewDecWithPrec(5, 1), // Unchanged from default 5%
	DowntimeJailDuration:    12 * time.Hour,                  // Mimics post-valsync-buffer network upgrade.
	SlashFractionDoubleSign: math.LegacyNewDec(0),            // 0% since mainnet V1 has trusted operators.
	SlashFractionDowntime:   math.LegacyNewDec(0),            // 0% since mainnet V1 has trusted operators.
}

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

func CreateUpgradeHandler2(
	mm *module.Manager,
	configurator module.Configurator,
	k *mintkeeper.Keeper,
	ak minttypes.AccountKeeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		inflation := math.LegacyNewDecWithPrec(11, 2)
		data := minttypes.GenesisState{
			Minter: minttypes.Minter{
				Inflation:        inflation,
				AnnualProvisions: math.LegacyNewDec(0),
			},
			Params: minttypes.Params{
				MintDenom:           "stake",
				InflationRateChange: math.LegacyNewDec(0),
				InflationMin:        inflation,
				InflationMax:        inflation,
				GoalBonded:          math.LegacyNewDecWithPrec(60, 0),
				BlocksPerYear:       15768000, // roughly, one block every 2 seconds
			}}
		k.InitGenesis(sdk.UnwrapSDKContext(ctx), ak, &data)

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
