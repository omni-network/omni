//nolint:dupl // similar code is okay
package admin

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// pauseXCall pauses xcalls on a chain.
func pauseXCall(ctx context.Context, s shared, c chain) error {
	log.Info(ctx, "Pausing xcall...", "chain", c.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("pauseXCall", s.manager, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Xcall paused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}

// pauseXCallTo pauses xcalls to a chain, on a chain.
func pauseXCallTo(ctx context.Context, s shared, c chain, toID uint64) error {
	to, ok := s.network.Chain(toID)
	if !ok {
		return errors.New("chain id not in network", "chain", toID)
	}

	log.Info(ctx, "Pausing xcall...", "chain", c.Name, "to", to.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("pauseXCallTo", s.manager, c.PortalAddress, to.ID)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Xcall paused ✅", "chain", c.Name, "to", to.Name, "addr", c.PortalAddress, "out", out)

	return nil
}

// pauseXCall pauses xcalls on a chain.
func unpauseXCall(ctx context.Context, s shared, c chain) error {
	log.Info(ctx, "Unpausing xcall...", "chain", c.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("unpauseXCall", s.manager, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Xcall unpaused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}

// pauseXCallTo pauses xcalls to a chain, on a chain.
func unpauseXCallTo(ctx context.Context, s shared, c chain, toID uint64) error {
	to, ok := s.network.Chain(toID)
	if !ok {
		return errors.New("chain id not in network", "chain", toID)
	}

	log.Info(ctx, "Unpausing xcall...", "chain", c.Name, "to", to.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("unpauseXCallTo", s.manager, c.PortalAddress, to.ID)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Xcall unpaused ✅", "chain", c.Name, "to", to.Name, "addr", c.PortalAddress, "out", out)

	return nil
}
