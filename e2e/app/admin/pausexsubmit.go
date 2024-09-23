//nolint:dupl // similar code is okay
package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

type XSubmitConfig struct {
	// Name of chain to pause/unpause xsubmits from
	From string
}

// PauseXSubmit pauses xsubmits on a network.
func PauseXSubmit(ctx context.Context, def app.Definition, baseCfg Config, xsubCfg XSubmitConfig) error {
	s := setup(def)

	if xsubCfg.From == "" {
		return s.run(ctx, baseCfg, pauseXSubmit)
	}

	from, ok := s.network.ChainByName(xsubCfg.From)
	if !ok {
		return errors.New("chain not found", "chain", xsubCfg.From)
	}

	return s.run(ctx, baseCfg, pauseXSubmitFrom(from))
}

func pauseXSubmitFrom(from netconf.Chain) func(ctx context.Context, s shared, c chain) error {
	return func(ctx context.Context, s shared, c chain) error {
		calldata, err := adminABI.Pack("pauseXSubmitFrom", s.admin, c.PortalAddress, from.ID)
		if err != nil {
			return errors.Wrap(err, "pack calldata", "chain", c.Name)
		}

		out, err := runForge(ctx, c.rpc, calldata, s.admin)
		if err != nil {
			return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
		}

		log.Info(ctx, "XSubmit paused ✅", "chain", c.Name, "to", from.Name, "addr", c.PortalAddress, "out", out)

		return nil
	}
}

func pauseXSubmit(ctx context.Context, s shared, c chain) error {
	calldata, err := adminABI.Pack("pauseXSubmit", s.admin, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := runForge(ctx, c.rpc, calldata, s.admin)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "XSubmit paused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}
