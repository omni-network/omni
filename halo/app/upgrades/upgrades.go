package upgrades

import (
	"context"
	"encoding/json"

	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	uluwatu1 "github.com/omni-network/omni/halo/app/upgrades/uluwatu"
	"github.com/omni-network/omni/lib/errors"

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
}

// AllUpgradeNames returns the names of all known Upgrades.
func AllUpgradeNames() []string {
	var resp []string
	for _, u := range Upgrades {
		resp = append(resp, u.Name)
	}

	return resp
}

func LatestUpgrade() string {
	return Upgrades[len(Upgrades)-1].Name
}

// NextUpgrade returns the next upgrade name after the provided previous upgrade,
// or false if this the latest upgrade (no next), or an error if the name is not a
// valid upgrade.
func NextUpgrade(prev string) (string, bool, error) {
	if prev == "" { // Return the first upgrade
		return Upgrades[0].Name, true, nil
	}

	for i, u := range Upgrades {
		if u.Name != prev {
			continue
		}

		if i == len(Upgrades)-1 {
			return "", false, nil // No next upgrade
		}

		return Upgrades[i+1].Name, true, nil
	}

	return "", false, errors.New("prev upgrade not found [BUG]")
}
