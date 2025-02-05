package app

import (
	"context"

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
				a.MintKeeper,
				a.AccountKeeper,
			)
		},
		Store: uluwatu1.StoreUpgrades,
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

// NextUpgrade returns the next upgrade name after the provided previous upgrade..
func NextUpgrade(prev string) (string, error) {
	if prev == "" { // Return the first upgrade
		return upgrades[0].Name, nil
	}

	for i, u := range upgrades {
		if i == len(upgrades)-1 {
			break
		}

		if u.Name == prev {
			return upgrades[i+1].Name, nil
		}
	}

	return "", errors.New("prev upgrade not found [BUG]")
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
