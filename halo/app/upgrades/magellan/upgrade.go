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
	"encoding/json"

	evmstakingtypes "github.com/omni-network/omni/halo/evmstaking/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	mintmodule "github.com/cosmos/cosmos-sdk/x/mint"
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
	MintParams       = minttypes.Params{
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
		log.Info(ctx, "Running 2_magellan upgrade handler")

		// evmstaking module doesn't require genesis or params.

		// Initialize mint module genesis (since it is being added in this upgrade)
		if err := initMintGenesis(ctx, mint, account, mintGenesisState()); err != nil {
			return nil, errors.Wrap(err, "init mint genesis")
		}

		// Register mint module consensus version in the version map
		// to avoid the SDK from triggering the default InitGenesis function which overrides above genesis state.
		// See https://docs.cosmos.network/main/learn/advanced/upgrade
		fromVM[minttypes.ModuleName] = mintmodule.AppModule{}.ConsensusVersion()

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

func mintGenesisState() *minttypes.GenesisState {
	minter := minttypes.InitialMinter(targetInflation)
	return minttypes.NewGenesisState(minter, MintParams)
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

func GenesisState(cdc codec.JSONCodec) (map[string]json.RawMessage, error) {
	bz, err := cdc.MarshalJSON(mintGenesisState())
	if err != nil {
		return nil, errors.Wrap(err, "marshal slashing genesis")
	}

	return map[string]json.RawMessage{
		minttypes.ModuleName: bz,
	}, nil
}
