package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// UnpausePortal unpauses the portal contracts on a network. Only single chain is supported.
func UnpausePortal(ctx context.Context, def app.Definition, cfg PortalAdminConfig) error {
	return run(ctx, def, cfg, unpausePortal)
}

func unpausePortal(ctx context.Context, s shared, c chain, r runner) error {
	calldata, err := adminABI.Pack("unpausePortal", s.admin, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := r.run(ctx, calldata, s.admin)
	if err != nil {
		return errors.Wrap(err, "run forge")
	}

	log.Info(ctx, "Admin.unpausePortal succeeded", "out", out)

	return nil
}
