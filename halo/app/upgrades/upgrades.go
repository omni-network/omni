package upgrades

import (
	"context"
	"encoding/json"

	drake3 "github.com/omni-network/omni/halo/app/upgrades/drake"
	earhart4 "github.com/omni-network/omni/halo/app/upgrades/earhart"
	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	uluwatu1 "github.com/omni-network/omni/halo/app/upgrades/uluwatu"
	evmredenomkeeper "github.com/omni-network/omni/halo/evmredenom/keeper"
	evmredenomsubmit "github.com/omni-network/omni/halo/evmredenom/submit"
	evmenginekeeper "github.com/omni-network/omni/octane/evmengine/keeper"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	slkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
)

type App interface {
	GetModuleManager() *module.Manager
	GetModuleConfigurator() module.Configurator
	GetSlashingKeeper() slkeeper.Keeper
	GetMintKeeper() mintkeeper.Keeper
	GetAccountKeeper() authkeeper.AccountKeeper
	GetEVMEngineKeeper() *evmenginekeeper.Keeper
	GetEVMRedenomKeeper() *evmredenomkeeper.Keeper
	GetEVMRedenomSubmitConfig() evmredenomsubmit.Config
}

// Upgrade defines a network upgrade.
type Upgrade struct {
	// Name of the upgrade. Must be <i>_<name>.
	Name string
	// HandlerFunc returns the upgrade handler for the upgrade.
	HandlerFunc func(App) upgradetypes.UpgradeHandler
	// Store returns the store upgrades for the upgrade, ie, which modules are added/renamed/removed.
	Store func(context.Context) *storetypes.StoreUpgrades
	// GenesisState adds the upgrade as part of genesis allowing Upgrades to be skipped for ephemeral chains.
	GenesisState func(codec.JSONCodec) (map[string]json.RawMessage, error)
}

var Upgrades = []Upgrade{
	{
		Name: uluwatu1.UpgradeName,
		HandlerFunc: func(a App) upgradetypes.UpgradeHandler {
			return uluwatu1.CreateUpgradeHandler(
				a.GetModuleManager(),
				a.GetModuleConfigurator(),
				a.GetSlashingKeeper(),
			)
		},
		Store:        uluwatu1.StoreUpgrades,
		GenesisState: uluwatu1.GenesisState,
	},
	{
		Name: magellan2.UpgradeName,
		HandlerFunc: func(a App) upgradetypes.UpgradeHandler {
			return magellan2.CreateUpgradeHandler(
				a.GetModuleManager(),
				a.GetModuleConfigurator(),
				a.GetMintKeeper(),
				a.GetSlashingKeeper(),
				a.GetAccountKeeper(),
			)
		},
		Store:        magellan2.StoreUpgrades,
		GenesisState: magellan2.GenesisState,
	},
	{
		Name: drake3.UpgradeName,
		HandlerFunc: func(a App) upgradetypes.UpgradeHandler {
			return drake3.CreateUpgradeHandler(
				a.GetModuleManager(),
				a.GetModuleConfigurator(),
			)
		},
		Store:        drake3.StoreUpgrades,
		GenesisState: drake3.GenesisState,
	},
	{
		Name: earhart4.UpgradeName,
		HandlerFunc: func(a App) upgradetypes.UpgradeHandler {
			return earhart4.CreateUpgradeHandler(
				a.GetModuleManager(),
				a.GetModuleConfigurator(),
				a.GetEVMEngineKeeper(),
				a.GetEVMRedenomKeeper(),
				a.GetEVMRedenomSubmitConfig(),
			)
		},
		Store:        earhart4.StoreUpgrades,
		GenesisState: earhart4.GenesisState,
	},
}
