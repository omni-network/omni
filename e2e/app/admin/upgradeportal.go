package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// UpgradePortal upgrades the portal contracts on a network.
func UpgradePortal(ctx context.Context, def app.Definition, cfg Config) error {
	return setup(def).run(ctx, cfg, upgradePortal)
}

func upgradePortal(ctx context.Context, s shared, c chain) error {
	// TODO: thie is the calldata to execute on the portal contract post upgrade
	// this is empty for now, but should be replaced with calldata if portal re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradePortal", s.admin, s.deployer, c.PortalAddress, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, scriptAdmin, c.rpc, calldata, s.admin, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Portal upgraded ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}
