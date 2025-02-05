// Package uluwatu defines the first omni consensus chain upgrade named after the iconic surf spot in Bali.
// It only includes attest module migration including constructor and logic changes.
// It doesn't include any store migrations.
package uluwatu

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/feature"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	sltypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

const UpgradeName = "1_uluwatu"

func StoreUpgrades(ctx context.Context) *storetypes.StoreUpgrades {
	var storeUpgrades storetypes.StoreUpgrades // Zero store upgrades

	if feature.FlagEVMStakingModule.Enabled(ctx) {
		storeUpgrades.Added = []string{minttypes.StoreKey} // Unless staking module flag in which case we add the mint module
	}

	return &storeUpgrades
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

	targetInflation  = math.LegacyNewDecWithPrec(115789, 6) // 11.5789% so that delegators earn ~11% after deducting of 5% validator rewards
	blocksPeriodSecs = 2                                    // BlocksPerYear calculated based on 2 second block times
	mintParams       = minttypes.Params{
		MintDenom:           sdk.DefaultBondDenom,
		InflationRateChange: math.LegacyNewDec(0),
		InflationMin:        targetInflation,
		InflationMax:        targetInflation,
		GoalBonded:          math.LegacyNewDecWithPrec(60, 2), // 60%
		BlocksPerYear:       uint64(365 * 24 * 60 * 60 / blocksPeriodSecs),
	}
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	slashing slkeeper.Keeper,
	mint mintkeeper.Keeper,
	account authkeeper.AccountKeeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if err := slashing.SetParams(ctx, SlashingParams); err != nil {
			return nil, errors.Wrap(err, "set slashing params")
		}

		if feature.FlagEVMStakingModule.Enabled(ctx) {
			// Initialize mint module genesis (since it is being added in this upgrade)
			minter := minttypes.InitialMinter(targetInflation)
			genState := minttypes.NewGenesisState(minter, mintParams)

			if err := initMintGenesis(ctx, mint, account, genState); err != nil {
				return nil, errors.Wrap(err, "init mint genesis")
			}
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

// initMintGenesis initializes mint module genesis
// It is a copy of mintkeeper.InitGenesis but with proper error handling and simple context.
func initMintGenesis(ctx context.Context, mint mintkeeper.Keeper, account minttypes.AccountKeeper, genesis *minttypes.GenesisState) error {
	if err := mint.Minter.Set(ctx, genesis.Minter); err != nil {
		return errors.Wrap(err, "set minter")
	}
	if err := mint.Params.Set(ctx, genesis.Params); err != nil {
		return errors.Wrap(err, "set mint params")
	}

	account.GetModuleAccount(ctx, minttypes.ModuleName) // This panics on error :(

	return nil
}
