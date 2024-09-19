package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// UnpausePortal unpauses the portal contracts on a network.
func UnpausePortal(ctx context.Context, def app.Definition, cfg Config) error {
	return setup(def).run(ctx, cfg, unpausePortal)
}

func unpausePortal(ctx context.Context, s shared, c chain) error {
	calldata, err := adminABI.Pack("unpausePortal", s.admin, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := runForge(ctx, scriptAdmin, c.rpc, calldata, s.admin)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Portal unpaused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}
