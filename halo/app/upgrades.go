package app

import (
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"fmt"
	v2 "github.com/omni-network/omni/halo/app/upgrades/v2"
	"github.com/omni-network/omni/lib/errors"
)

func (a App) setUpgradeHandlers() error {
	a.UpgradeKeeper.SetUpgradeHandler(
		v2.UpgradeName,
		v2.CreateUpgradeHandler(a.ModuleManager, a.Configurator(), a.StakingKeeper),
	)

	upgradeInfo, err := a.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		return errors.Wrap(err, "read upgrade info from disk")
	}
	fmt.Printf("🔥!! upgradeInfo=%#v\n", upgradeInfo.String())

	if upgradeInfo.Name == "" || a.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return nil
	}

	var storeUpgrades *storetypes.StoreUpgrades

	switch upgradeInfo.Name {
	case v2.UpgradeName:
		// No store upgrades
	default:
		return errors.New("unknown upgrade [BUG]")
	}

	if storeUpgrades != nil {
		a.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	}

	return nil
}
