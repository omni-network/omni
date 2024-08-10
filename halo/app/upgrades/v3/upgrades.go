package v3

import (
	"context"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/module"
	attestkeeper "github.com/omni-network/omni/halo/attest/keeper"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	akeeper *attestkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fmt.Printf("🔥🔥🔥🔥🔥🔥🔥🔥!! UpgradeHandler: plan=%v, fromVM=%v\n", plan.Name, fromVM)
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
