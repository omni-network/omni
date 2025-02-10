// Package magellan defines the second Omni consensus chain upgrade, named after
// Ferdinand Magellan, the famed explorer who led the first expedition to
// circumnavigate the globe.
//
// It includes the following features:
// - Support user staking delegations (previously only validator self-delegation supported)
// - Inflation staking rewards (11% APR)
// - Pending withdrawals created in queue (not processed to EVM yet).
// - Simplified EVM events processing.
// - Proto encoding of EVM payload.
package magellan

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

const UpgradeName = "2_magellan"

func StoreUpgrades(_ context.Context) *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{minttypes.StoreKey}, // Add the mint module
	}
}

var (
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
	mint mintkeeper.Keeper,
	account authkeeper.AccountKeeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// Initialize mint module genesis (since it is being added in this upgrade)
		minter := minttypes.InitialMinter(targetInflation)
		genState := minttypes.NewGenesisState(minter, mintParams)

		if err := initMintGenesis(ctx, mint, account, genState); err != nil {
			return nil, errors.Wrap(err, "init mint genesis")
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
