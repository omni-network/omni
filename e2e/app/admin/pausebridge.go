package admin

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

// pauseBridge pauses a bridge contract (native or L1) on a chain.
func pauseBridge(ctx context.Context, s shared, c chain, addr common.Address, actionID [32]byte, actionLabel string) error {
	log.Info(ctx, "Pausing bridge...", "chain", c.Name, "addr", addr, "action", actionLabel)

	calldata, err := adminABI.Pack("pauseBridge", s.owner, addr, actionID)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.owner)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Bridge paused ✅", "chain", c.Name, "action", actionLabel, "addr", c.PortalAddress, "out", out)

	return nil
}

// unpauseBridge unpauses a bridge contract (native or L1) on a chain.
func unpauseBridge(ctx context.Context, s shared, c chain, addr common.Address, actionID [32]byte, actionLabel string) error {
	log.Info(ctx, "Unpausing bridge...", "chain", c.Name, "addr", addr, "action", actionLabel)

	calldata, err := adminABI.Pack("unpauseBridge", s.owner, addr, actionID)
	if err != nil {
		return errors.Wrap(err, "pack calldata", "chain", c.Name)
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.owner)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out, "chain", c.Name)
	}

	log.Info(ctx, "Bridge unpaused ✅", "chain", c.Name, "action", actionLabel, "addr", c.PortalAddress, "out", out)

	return nil
}
