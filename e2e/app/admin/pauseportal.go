package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// PausePortal pauses the portal contracts on a network. Only single chain is supported.
func PausePortal(ctx context.Context, def app.Definition, cfg PortalAdminConfig) error {
	return run(ctx, def, cfg, pausePortal)
}

func pausePortal(ctx context.Context, s shared, c chain, r runner) error {
	calldata, err := adminABI.Pack("pausePortal", s.admin, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := r.run(ctx, calldata, s.admin)
	if err != nil {
		return errors.Wrap(err, "run forge")
	}

	log.Info(ctx, "Admin.pausePortal succeeded", "out", out)

	return nil
}
