package app

import (
	"context"

	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	uluwatu1 "github.com/omni-network/omni/halo/app/upgrades/uluwatu"
	"github.com/omni-network/omni/lib/errors"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
)

// Upgrade defines a network upgrade.
type Upgrade struct {
	Name        string
	HandlerFunc func(App) upgradetypes.UpgradeHandler
	Store       func(context.Context) *storetypes.StoreUpgrades
}

var upgrades = []Upgrade{
	{
		Name: uluwatu1.UpgradeName,
		HandlerFunc: func(a App) upgradetypes.UpgradeHandler {
			return uluwatu1.CreateUpgradeHandler(
				a.ModuleManager,
				a.Configurator(),
				a.SlashingKeeper,
			)
		},
		Store: uluwatu1.StoreUpgrades,
	},
	{
		Name: magellan2.UpgradeName,
		HandlerFunc: func(a App) upgradetypes.UpgradeHandler {
			return magellan2.CreateUpgradeHandler(
				a.ModuleManager,
				a.Configurator(),
				a.MintKeeper,
				a.AccountKeeper,
			)
		},
		Store: magellan2.StoreUpgrades,
	},
}

// AllUpgrades returns the names of all known upgrades.
func AllUpgrades() []string {
	var resp []string
	for _, u := range upgrades {
		resp = append(resp, u.Name)
	}

	return resp
}

// NextUpgrade returns the next upgrade name after the provided previous upgrade,
// or false if this the latest upgrade (no next), or an error if the name is not a
// valid upgrade.
func NextUpgrade(prev string) (string, bool, error) {
	if prev == "" { // Return the first upgrade
		return upgrades[0].Name, true, nil
	}

	for i, u := range upgrades {
		if u.Name != prev {
			continue
		}

		if i == len(upgrades)-1 {
			return "", false, nil // No next upgrade
		}

		return upgrades[i+1].Name, true, nil
	}

	return "", false, errors.New("prev upgrade not found [BUG]")
}

func (a App) setUpgradeHandlers(ctx context.Context) error {
	for _, u := range upgrades {
		a.UpgradeKeeper.SetUpgradeHandler(u.Name, u.HandlerFunc(a))
	}

	upgradeInfo, err := a.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		return errors.Wrap(err, "read upgrade info from disk")
	} else if upgradeInfo.Name == "" {
		return nil // No upgrade info found
	}

	for _, u := range upgrades {
		if u.Name != upgradeInfo.Name {
			continue
		}

		a.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, u.Store(ctx)))

		return nil
	}

	return errors.New("unknown upgrade info [BUG]", "name", upgradeInfo.Name)
}
