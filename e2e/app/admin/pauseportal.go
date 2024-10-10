package admin

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// pausePortal pauses the portal contract on a chain.
func pausePortal(ctx context.Context, s shared, c chain) error {
	log.Info(ctx, "Pausing portal...", "chain", c.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("pausePortal", s.manager, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Portal paused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}

// unpausePortal pauses the portal contract on a chain.
func unpausePortal(ctx context.Context, s shared, c chain) error {
	log.Info(ctx, "Unpausing portal...", "chain", c.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("unpausePortal", s.manager, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Portal unpaused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}
