package admin

import (
	"context"
	"math/rand"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
)

// Test tests all admin commands against an ephemeral network.
func Test(ctx context.Context, def app.Definition) error {
	if !def.Testnet.Network.IsEphemeral() {
		return errors.New("only ephemeral networks")
	}

	network := app.NetworkFromDef(def)

	// pause portal on one chain
	chain := randChain(network)
	if err := PausePortal(ctx, def, Config{Chain: chain.Name}); err != nil {
		return errors.Wrap(err, "pause portal", "chain", chain.Name)
	}

	// check if portal is paused
	if err := checkPortalPaused(def, chain, true); err != nil {
		return errors.Wrap(err, "check paused")
	}

	// unpause portal on one chain
	if err := UnpausePortal(ctx, def, Config{Chain: chain.Name}); err != nil {
		return errors.Wrap(err, "unpause portal", "chain", chain.Name)
	}

	// check if portal is unpaused
	if err := checkPortalPaused(def, chain, false); err != nil {
		return errors.Wrap(err, "check paused")
	}

	// upgrade portal
	if err := UpgradePortal(ctx, def, Config{Chain: chain.Name}); err != nil {
		return errors.Wrap(err, "upgrade portal", "chain", chain.Name)
	}

	// pause all portals
	if err := PausePortal(ctx, def, Config{}); err != nil {
		return errors.Wrap(err, "pause all portals")
	}

	// check if all portals are paused
	for _, chain := range network.EVMChains() {
		if err := checkPortalPaused(def, chain, true); err != nil {
			return errors.Wrap(err, "check paused")
		}
	}

	// unpause all portals
	if err := UnpausePortal(ctx, def, Config{}); err != nil {
		return errors.Wrap(err, "unpause all portals")
	}

	// check if all portals are unpaused
	for _, chain := range network.EVMChains() {
		if err := checkPortalPaused(def, chain, false); err != nil {
			return errors.Wrap(err, "check paused")
		}
	}

	// upgrade all portals
	if err := UpgradePortal(ctx, def, Config{}); err != nil {
		return errors.Wrap(err, "upgrade all portals")
	}

	return nil
}

func randChain(network netconf.Network) netconf.Chain {
	chains := network.EVMChains()
	//nolint:gosec // no need for secure randomneness
	return chains[rand.Intn(len(chains))]
}

func checkPortalPaused(def app.Definition, chain netconf.Chain, expected bool) error {
	backend, err := def.Backends().Backend(chain.ID)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	portal, err := bindings.NewOmniPortal(chain.PortalAddress, backend)
	if err != nil {
		return errors.Wrap(err, "new portal")
	}

	paused, err := portal.IsPaused1(nil)
	if err != nil {
		return errors.Wrap(err, "get paused")
	}

	if paused != expected {
		return errors.New("check paused", "chain", chain.Name, "paused", paused, "expected", expected)
	}

	return nil
}
