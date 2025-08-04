package testupgrade

import (
	"context"

	earhart4 "github.com/omni-network/omni/halo/app/upgrades/earhart"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
)

// PrepFor prepares tests pre-network-upgrade.
func PrepFor(ctx context.Context, omniEVM *ethbackend.Backend, upgradeName string) error {
	if upgradeName == earhart4.UpgradeName {
		return prepForEarhart(ctx, omniEVM)
	}

	return nil
}

// Ensure ensures the state after the network upgrade.
func Ensure(ctx context.Context, omniEVM *ethbackend.Backend) error {
	return ensureEarhart(ctx, omniEVM)
}
