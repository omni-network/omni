package app

import (
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	v2 "github.com/omni-network/omni/halo/app/upgrades/v2"
	v3 "github.com/omni-network/omni/halo/app/upgrades/v3"
	"github.com/omni-network/omni/lib/errors"
)

func (a App) setUpgradeHandlers() error {
	a.UpgradeKeeper.SetUpgradeHandler(
		v2.UpgradeName,
		v2.CreateUpgradeHandler(a.ModuleManager, a.Configurator(), a.StakingKeeper),
	)
	a.UpgradeKeeper.SetUpgradeHandler(
		v3.UpgradeName,
		v3.CreateUpgradeHandler(a.ModuleManager, a.Configurator(), a.AttestKeeper),
	)

	upgradeInfo, err := a.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		return errors.Wrap(err, "read upgrade info from disk")
	}

	if upgradeInfo.Name == "" || a.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return nil
	}

	var storeUpgrades *storetypes.StoreUpgrades

	switch upgradeInfo.Name {
	case v2.UpgradeName:
		// No store upgrades
	case v3.UpgradeName:
		storeUpgrades = &storetypes.StoreUpgrades{
			Added:   []string{"foo"},
			Renamed: nil,
			Deleted: nil,
		}
	default:
		return errors.New("unknown upgrade [BUG]")
	}

	if storeUpgrades != nil {
		a.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	}

	return nil
}
