// Package magellan defines the second Omni consensus chain upgrade, named after
// Ferdinand Magellan, the famed explorer who led the first expedition to
// circumnavigate the globe.
//
// It includes:
// - Normal staking delegations (previously only self-delegations)
// - 11% inflation rewards for delegations (previously no rewards)
// - Enqueuing of staking withdrawals to EVM (not processed yet)
// - Protobuf encoding of EVM payload in blocks (improved performance and security)
// - Simplified EVM events processing (improved performance and security)
package magellan

import (
	"context"

	evmstakingtypes "github.com/omni-network/omni/halo/evmstaking/types"
	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

const UpgradeName = "2_magellan"

func StoreUpgrades(_ context.Context) *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{
			minttypes.StoreKey,         // Add the mint module
			evmstakingtypes.ModuleName, // Add the EVM staking module
		},
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
		// evmstaking module doesn't require genesis or params.

		// Initialize mint module genesis (since it is being added in this upgrade)
		minter := minttypes.InitialMinter(targetInflation)
		genState := minttypes.NewGenesisState(minter, mintParams)
		if err := initMintGenesis(ctx, mint, account, genState); err != nil {
			return nil, errors.Wrap(err, "init mint genesis")
		}

		// Add burner permission to distribution module in auth module state
		accI, _ := account.GetModuleAccountAndPermissions(ctx, distrtypes.ModuleName)
		if accI == nil {
			return nil, errors.New("distribution module account not found")
		} else if !accI.HasPermission(authtypes.Burner) {
			acc, ok := accI.(*authtypes.ModuleAccount)
			if !ok {
				return nil, errors.New("distribution module account is not a module account")
			}
			acc.Permissions = append(acc.Permissions, authtypes.Burner)

			account.SetModuleAccount(ctx, acc)
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
