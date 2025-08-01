package admin

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

// pauseGasPump pauses the gas pump contract on a chain.
func pauseGasPump(ctx context.Context, s shared, c chain, addr common.Address) error {
	log.Info(ctx, "Pausing gas pump...", "chain", c.Name, "addr", addr)

	calldata, err := adminABI.Pack("pauseGasPump", s.manager, addr)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, adminScriptName, coreContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Gas pump paused ✅", "chain", c.Name, "addr", addr, "out", out)

	return nil
}

// unpauseGasPump unpauses the gas pump contract on a chain.
func unpauseGasPump(ctx context.Context, s shared, c chain, addr common.Address) error {
	log.Info(ctx, "Unpausing gas pump...", "chain", c.Name, "addr", addr)

	calldata, err := adminABI.Pack("unpauseGasPump", s.manager, addr)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, adminScriptName, coreContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Gas pump unpaused ✅", "chain", c.Name, "addr", addr, "out", out)

	return nil
}

// pauseGasStation pauses the gas station contract on a chain.
func pauseGasStation(ctx context.Context, s shared, c chain, addr common.Address) error {
	log.Info(ctx, "Pausing gas station...", "chain", c.Name, "addr", addr)

	calldata, err := adminABI.Pack("pauseGasStation", s.manager, addr)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, adminScriptName, coreContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Gas station paused ✅", "chain", c.Name, "addr", addr, "out", out)

	return nil
}

// unpauseGasStation unpauses the gas station contract on a chain.
func unpauseGasStation(ctx context.Context, s shared, c chain, addr common.Address) error {
	log.Info(ctx, "Unpausing gas station...", "chain", c.Name, "addr", addr)

	calldata, err := adminABI.Pack("unpauseGasStation", s.manager, addr)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, adminScriptName, coreContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Gas station unpaused ✅", "chain", c.Name, "addr", addr, "out", out)

	return nil
}
