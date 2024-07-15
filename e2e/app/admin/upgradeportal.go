package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// UpgradePortal upgrades the portal contracts on a network. Only single chain is supported.
func UpgradePortal(ctx context.Context, def app.Definition, cfg PortalAdminConfig) error {
	return run(ctx, def, cfg, upgradePortal)
}

func upgradePortal(ctx context.Context, s shared, c chain, r runner) error {
	// TODO: thie is the calldata to execute on the portal contract post upgrade
	// this is empty for now, but should be replaced with calldata if portal re-initialization is required
	initCalldata := []byte{}

	calldata, err := adminABI.Pack("upgradePortal", s.admin, s.deployer, c.PortalAddress, initCalldata)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := r.run(ctx, calldata, s.admin, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge")
	}

	log.Info(ctx, "Admin.upgradePortal succeeded", "out", out)

	return nil
}
