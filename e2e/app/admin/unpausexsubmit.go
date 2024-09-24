//nolint:dupl // similar code is okay
package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

// UnpauseXSubmit unpauses xsubmits on a network.
func UnpauseXSubmit(ctx context.Context, def app.Definition, baseCfg Config, xsubCfg XSubmitConfig) error {
	s := setup(def)

	if xsubCfg.From == "" {
		return s.run(ctx, baseCfg, unpauseXSubmit)
	}

	from, ok := s.network.ChainByName(xsubCfg.From)
	if !ok {
		return errors.New("chain not found", "chain", xsubCfg.From)
	}

	return s.run(ctx, baseCfg, unpauseXSubmitFrom(from))
}

func unpauseXSubmitFrom(from netconf.Chain) func(ctx context.Context, s shared, c chain) error {
	return func(ctx context.Context, s shared, c chain) error {
		calldata, err := adminABI.Pack("unpauseXSubmitFrom", s.admin, c.PortalAddress, from.ID)
		if err != nil {
			return errors.Wrap(err, "pack calldata", "chain", c.Name)
		}

		out, err := runForge(ctx, c.rpc, calldata, s.admin)
		if err != nil {
			return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
		}

		log.Info(ctx, "XSubmit unpaused ✅", "chain", c.Name, "from", from.Name, "addr", c.PortalAddress, "out", out)

		return nil
	}
}

func unpauseXSubmit(ctx context.Context, s shared, c chain) error {
	calldata, err := adminABI.Pack("unpauseXSubmit", s.admin, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := runForge(ctx, c.rpc, calldata, s.admin)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "XSubmit unpaused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}
