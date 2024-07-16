package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
)

// UnpausePortal unpauses the portal contracts on a network. Only single chain is supported.
func UnpausePortal(ctx context.Context, def app.Definition, cfg PortalAdminConfig) error {
	return run(ctx, def, cfg, "unpausePortal", unpausePortal)
}

func unpausePortal(ctx context.Context, s shared, c chain, r runner) (string, error) {
	calldata, err := adminABI.Pack("unpausePortal", s.admin, c.PortalAddress)
	if err != nil {
		return "", errors.Wrap(err, "pack calldata")
	}

	out, err := r.run(ctx, calldata, s.admin)
	if err != nil {
		return out, errors.Wrap(err, "run forge")
	}

	return out, nil
}
