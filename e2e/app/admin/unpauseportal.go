package admin

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

var unpausePortalABI = mustGetABI(bindings.UnpausePortalMetaData)

// UnpausePortal unpauses the portal contracts on a network. Only single chain is supported.
func UnpausePortal(ctx context.Context, def app.Definition, cfg PausePortalConfig) error {
	if err := cfg.Validate(); err != nil {
		return errors.Wrap(err, "validate config")
	}

	s, err := setup(ctx, def)
	if err != nil {
		return errors.Wrap(err, "setup")
	}

	c, err := setupChain(ctx, s, cfg.Chain)
	if err != nil {
		return errors.Wrap(err, "setup chain")
	}

	if err := unpausePortalForge(ctx, c.rpc, s.admin, c.PortalAddress); err != nil {
		return errors.Wrap(err, "pause portal")
	}

	return nil
}

func unpausePortalForge(ctx context.Context, rpc string, sender common.Address, portal common.Address) error {
	calldata, err := unpausePortalABI.Pack("run", portal)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, "UnpausePortal", rpc, calldata, sender)
	if err != nil {
		return errors.Wrap(err, "pause portal forge")
	}

	log.Info(ctx, "UnpausePortal forge script", "out", out)

	return nil
}
