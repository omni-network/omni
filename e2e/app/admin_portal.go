package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

var portalAdminABI = mustGetABI(bindings.PortalAdminMetaData)

type PortalAdminConfig struct {
	// Chain is the name of the chain on which to run the command  (use "all" for all chains)
	Chain string
}

func DefaultPortalAdminConfig() PortalAdminConfig {
	return PortalAdminConfig{Chain: ""}
}

func (cfg PortalAdminConfig) Validate() error {
	if cfg.Chain == "" {
		return errors.New("chain must be set")
	}

	return nil
}

// PausePortal pauses the portal contracts on a network. Only single chain is supported.
func PausePortal(ctx context.Context, def Definition, cfg PortalAdminConfig) error {
	if err := cfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid config")
	}

	s, err := setup(def)
	if err != nil {
		return err
	}

	return forChains(ctx, maybeAll(s.network, cfg.Chain), s, pausePortal)
}

// UnpausePortal unpauses the portal contracts on a network. Only single chain is supported.
func UnpausePortal(ctx context.Context, def Definition, cfg PortalAdminConfig) error {
	if err := cfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid config")
	}

	s, err := setup(def)
	if err != nil {
		return err
	}

	return forChains(ctx, maybeAll(s.network, cfg.Chain), s, unpausePortal)
}

// UpgradePortal upgrades the portal contracts on a network. Only single chain is supported.
func UpgradePortal(ctx context.Context, def Definition, cfg PortalAdminConfig) error {
	if err := cfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid config")
	}

	s, err := setup(def)
	if err != nil {
		return err
	}

	return forChains(ctx, maybeAll(s.network, cfg.Chain), s, upgradePortal)
}

func pausePortal(ctx context.Context, s shared, c chain) error {
	calldata, err := portalAdminABI.Pack("pause", s.admin, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, scriptPortalAdmin, c.rpc, calldata, s.admin)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Paused portal", "chain", c.Name, "address", c.PortalAddress.Hex(), "out", out)

	return nil
}

func unpausePortal(ctx context.Context, s shared, c chain) error {
	calldata, err := portalAdminABI.Pack("unpause", s.admin, c.PortalAddress)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, scriptPortalAdmin, c.rpc, calldata, s.admin)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Unpaused portal", "chain", c.Name, "address", c.PortalAddress.Hex(), "out", out)

	return nil
}

func upgradePortal(ctx context.Context, s shared, c chain) error {
	// TODO: thie is the calldata to execute on the portal contract post upgrade
	// this is empty for now, but should be replaced with calldata if portal re-initialization is required
	initCalldata := []byte{}

	calldata, err := portalAdminABI.Pack("upgrade", s.admin, s.deployer, c.PortalAddress, initCalldata)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, scriptPortalAdmin, c.rpc, calldata, s.admin, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Upgraded portal", "chain", c.Name, "address", c.PortalAddress.Hex(), "out", out)

	return nil
}
