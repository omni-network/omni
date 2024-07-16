package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
)

// PausePortal pauses the portal contracts on a network. Only single chain is supported.
func PausePortal(ctx context.Context, def app.Definition, cfg PortalAdminConfig) error {
	return run(ctx, def, cfg, "pausePortal", pausePortal)
}

func pausePortal(ctx context.Context, s shared, c chain, r runner) (string, error) {
	calldata, err := adminABI.Pack("pausePortal", s.admin, c.PortalAddress)
	if err != nil {
		return "", errors.Wrap(err, "pack calldata")
	}

	out, err := r.run(ctx, calldata, s.admin)
	if err != nil {
		return out, errors.Wrap(err, "run forge")
	}

	return out, nil
}
