package admin

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

// pauseSolverNetAll pauses or unpauses all functions on the SolverNetInbox.
func pauseSolverNetAll(ctx context.Context, s shared, c chain, addr common.Address, pause bool) error {
	action := getActionText(pause)

	log.Info(ctx, action+" all on SolverNetInbox...", "chain", c.Name, "addr", addr)

	calldata, err := solverNetAdminABI.Pack("pauseSolverNetAll", s.manager, addr, pause)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, solverNetAdminScriptName, solveContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "All functions on SolverNetInbox "+getStatusText(pause)+" ✅", "chain", c.Name, "addr", addr, "out", out)

	return nil
}

// pauseSolverNetOpen pauses or unpauses the open function on the SolverNetInbox.
func pauseSolverNetOpen(ctx context.Context, s shared, c chain, addr common.Address, pause bool) error {
	action := getActionText(pause)

	log.Info(ctx, action+" open on SolverNetInbox...", "chain", c.Name, "addr", addr)

	calldata, err := solverNetAdminABI.Pack("pauseSolverNetOpen", s.manager, addr, pause)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, solverNetAdminScriptName, solveContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Open function on SolverNetInbox "+getStatusText(pause)+" ✅", "chain", c.Name, "addr", addr, "out", out)

	return nil
}

// pauseSolverNetClose pauses or unpauses the close function on the SolverNetInbox.
func pauseSolverNetClose(ctx context.Context, s shared, c chain, addr common.Address, pause bool) error {
	action := getActionText(pause)

	log.Info(ctx, action+" close on SolverNetInbox...", "chain", c.Name, "addr", addr)

	calldata, err := solverNetAdminABI.Pack("pauseSolverNetClose", s.manager, addr, pause)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.RPCEndpoint, solverNetAdminScriptName, solveContracts, calldata, s.manager)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Close function on SolverNetInbox "+getStatusText(pause)+" ✅", "chain", c.Name, "addr", addr, "out", out)

	return nil
}

// getActionText returns "Pausing" or "Unpausing" based on the pause flag.
func getActionText(pause bool) string {
	if pause {
		return "Pausing"
	}

	return "Unpausing"
}

// getStatusText returns "paused" or "unpaused" based on the pause flag.
func getStatusText(pause bool) string {
	if pause {
		return "paused"
	}

	return "unpaused"
}
