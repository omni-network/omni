package queryutil

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/halo/app/upgrades/static"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
)

// CurrentUpgrade returns the current applied upgrade.
//
// Note it will return genesis upgrade if unknown upgrades are applied.
// This is due to CosmosSDK not providing an API to actually fetch applied upgrades :(.
func CurrentUpgrade(ctx context.Context, cprov cchain.Provider) (string, error) {
	var latest string
	for _, upgrade := range static.UpgradeNames {
		plan, ok, err := cprov.AppliedPlan(ctx, upgrade)
		if err != nil {
			return "", errors.Wrap(err, "fetching applied plan")
		} else if !ok {
			continue
		} else if upgrade != plan.Name {
			return "", errors.New("unexpected upgrade plan name [BUG]", "expected", upgrade, "actual", plan.Name)
		}

		latest = upgrade
	}

	return latest, nil
}

// NextUpgrade returns the next upgrade to apply, or false if all upgrades has been applied.
func NextUpgrade(ctx context.Context, cprov cchain.Provider) (string, bool, error) {
	current, err := CurrentUpgrade(ctx, cprov)
	if err != nil {
		return "", false, err
	}

	return static.NextUpgrade(current)
}

func IsPostEVMRedenom(ctx context.Context, provider cchain.Provider) (bool, error) {
	current, err := CurrentUpgrade(ctx, provider)
	if err != nil {
		return false, err
	} else if current == "" {
		return false, nil // No upgrades applied, so not post EVM redenom
	}

	var version uint64
	var name string
	if _, err := fmt.Sscanf(current, "%d_%s", &version, &name); err != nil {
		return false, errors.Wrap(err, "not %d_%s upgrade", "upgrade", current)
	}

	// EVM redenom done if 4_earhart.

	return version >= 4, nil
}
