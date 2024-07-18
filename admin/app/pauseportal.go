package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
)

// PausePortal pauses the portal contracts on a network. Only single chain is supported.
func PausePortal(ctx context.Context, cfg Config) error {
	return run(ctx, cfg, "pausePortal", pausePortal)
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
