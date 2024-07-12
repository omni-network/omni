package admin

import (
	"context"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

type PausePortalConfig struct {
	Chain string // Pause a specific chain
}

func DefaultPausePortalConfig() PausePortalConfig {
	return PausePortalConfig{
		Chain: "",
	}
}

func (cfg PausePortalConfig) Validate() error {
	if cfg.Chain == "" {
		return errors.New("chain is required")
	}

	return nil
}

// PausePortal pauses the portal contracts on a network. Only single chain is supported.
func PausePortal(ctx context.Context, def app.Definition, cfg PausePortalConfig) error {
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

	if err := pausePortalForge(ctx, c.rpc, s.admin, c.PortalAddress); err != nil {
		return errors.Wrap(err, "pause portal")
	}

	return nil
}

func pausePortalForge(ctx context.Context, rpc string, sender common.Address, portal common.Address) error {
	calldata, err := adminABI.Pack("pausePortal", portal)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, "PausePortal", rpc, calldata, sender)
	if err != nil {
		return errors.Wrap(err, "pause portal forge")
	}

	log.Info(ctx, "PausePortal forge script", "out", out)

	return nil
}
