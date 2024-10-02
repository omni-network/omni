package app

import (
	uluwatu1 "github.com/omni-network/omni/halo/app/upgrades/uluwatu"
	"github.com/omni-network/omni/lib/errors"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
)

func (a App) setUpgradeHandlers() error {
	upgrades := []struct {
		Name    string
		Handler upgradetypes.UpgradeHandler
		Store   storetypes.StoreUpgrades
	}{
		{
			Name:    uluwatu1.UpgradeName,
			Handler: uluwatu1.CreateUpgradeHandler(a.ModuleManager, a.Configurator(), a.SlashingKeeper),
			Store:   uluwatu1.StoreUpgrades,
		},
	}

	for _, u := range upgrades {
		a.UpgradeKeeper.SetUpgradeHandler(u.Name, u.Handler)
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

		a.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &u.Store))

		return nil
	}

	return errors.New("unknown upgrade info [BUG]", "name", upgradeInfo.Name)
}
