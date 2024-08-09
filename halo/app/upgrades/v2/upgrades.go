package v2

import (
	"context"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/module"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/omni-network/omni/lib/errors"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	skeeper *skeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fmt.Printf("🔥🔥🔥🔥🔥🔥🔥🔥!! UpgradeHandler: plan=%v, fromVM=%v\n", plan.Name, fromVM)
		params, err := skeeper.GetParams(ctx)
		if err != nil {
			return nil, err
		}

		params.MaxValidators *= 2 // Double the max validators

		if err := skeeper.SetParams(ctx, params); err != nil {
			return nil, errors.Wrap(err, "set params")
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
