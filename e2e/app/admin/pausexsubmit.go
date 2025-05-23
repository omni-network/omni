//nolint:dupl // similar code is okay
package admin

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// pauseXSubmit pauses xsubmits on a chain.
func pauseXSubmit(ctx context.Context, s shared, c chain) error {
	log.Info(ctx, "Pausing xsubmit...", "chain", c.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("pauseXSubmit", s.manager, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, adminScriptName, coreContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "XSubmit paused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}

// pauseXSubmitFrom pauses xsubmits from a chain, on a chain.
func pauseXSubmitFrom(ctx context.Context, s shared, c chain, fromID uint64) error {
	from, ok := s.testnet.EVMChainByID(fromID)
	if !ok {
		return errors.New("chain id not in network", "chain", fromID)
	}

	log.Info(ctx, "Pausing xsubmit...", "chain", c.Name, "from", from.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("pauseXSubmitFrom", s.manager, c.PortalAddress, from.ChainID)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, adminScriptName, coreContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "XSubmit paused ✅", "chain", c.Name, "from", from.Name, "addr", c.PortalAddress, "out", out)

	return nil
}

// pauseXSubmit pauses xsubmits on a chain.
func unpauseXSubmit(ctx context.Context, s shared, c chain) error {
	log.Info(ctx, "Unpausing xsubmit...", "chain", c.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("unpauseXSubmit", s.manager, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, adminScriptName, coreContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "XSubmit unpaused ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}

// pauseXSubmitFrom pauses xsubmits from a chain, on a chain.
func unpauseXSubmitFrom(ctx context.Context, s shared, c chain, fromID uint64) error {
	from, ok := s.testnet.EVMChainByID(fromID)
	if !ok {
		return errors.New("chain id not in network", "chain", fromID)
	}

	log.Info(ctx, "Unpausing xsubmit...", "chain", c.Name, "from", from.Name, "addr", c.PortalAddress)

	calldata, err := adminABI.Pack("unpauseXSubmitFrom", s.manager, c.PortalAddress, from.ChainID)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, adminScriptName, coreContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "XSubmit unpaused ✅", "chain", c.Name, "from", from.Name, "addr", c.PortalAddress, "out", out)

	return nil
}
