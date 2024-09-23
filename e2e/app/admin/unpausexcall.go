//nolint:dupl // similar code is okay
package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

// UnpauseXCall unpauses xcalls on a network.
func UnpauseXCall(ctx context.Context, def app.Definition, baseCfg Config, xcallCfg XCallConfig) error {
	s := setup(def)

	if xcallCfg.To == "" {
		return s.run(ctx, baseCfg, unpauseXCall)
	}

	to, ok := s.network.ChainByName(xcallCfg.To)
	if !ok {
		return errors.New("chain not found", "chain", xcallCfg.To)
	}

	return s.run(ctx, baseCfg, unpauseXCallTo(to))
}

func unpauseXCallTo(to netconf.Chain) func(ctx context.Context, s shared, c chain) error {
	return func(ctx context.Context, s shared, c chain) error {
		calldata, err := adminABI.Pack("unpauseXCallTo", s.admin, c.PortalAddress, to.ID)
		if err != nil {
			return errors.Wrap(err, "pack calldata", "chain", c.Name)
		}

		out, err := runForge(ctx, c.rpc, calldata, s.admin)
		if err != nil {
			return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
		}

		log.Info(ctx, "Xcall unpaused ✅", "chain", c.Name, "to", to.Name, "addr", c.PortalAddress, "out", out)

		return nil
	}
}

func unpauseXCall(ctx context.Context, s shared, c chain) error {
	calldata, err := adminABI.Pack("unpauseXCall", s.admin, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := runForge(ctx, c.rpc, calldata, s.admin)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Xcall unpaused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}
