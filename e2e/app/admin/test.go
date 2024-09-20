package admin

import (
	"context"
	"math/rand"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Test tests all admin commands against an ephemeral network.
func Test(ctx context.Context, def app.Definition) error {
	if !def.Testnet.Network.IsEphemeral() {
		return errors.New("only ephemeral networks")
	}

	log.Info(ctx, "Running contract admin tests.")

	network := app.NetworkFromDef(def)

	if err := testUpgradePortal(ctx, def, network); err != nil {
		return err
	}

	if err := testPauseUnpause(ctx, def, network); err != nil {
		return err
	}

	if err := testPauseUnpauseXCall(ctx, def, network); err != nil {
		return err
	}

	if err := testPauseUnpauseXCallTo(ctx, def, network); err != nil {
		return err
	}

	log.Info(ctx, "Done.")

	return nil
}

// testPauseUnpausePortal tests PausePortal and UnpausePortal commands.
func testPauseUnpause(ctx context.Context, def app.Definition, network netconf.Network) error {
	chain := randChain(network)

	err := forOne(ctx, def, chain, PausePortal, checkPortalPaused(true))
	if err != nil {
		return errors.Wrap(err, "pause portal")
	}

	err = forOne(ctx, def, chain, UnpausePortal, checkPortalPaused(false))
	if err != nil {
		return errors.Wrap(err, "unpause portal")
	}

	err = forAll(ctx, def, network, PausePortal, checkPortalPaused(true))
	if err != nil {
		return errors.Wrap(err, "pause all portals")
	}

	err = forAll(ctx, def, network, UnpausePortal, checkPortalPaused(false))
	if err != nil {
		return errors.Wrap(err, "unpause all portals")
	}

	return nil
}

// testPauseUnpauseXCall tests PauseXCall and UnpauseXCall commands (without XCallConfig.To).
func testPauseUnpauseXCall(ctx context.Context, def app.Definition, network netconf.Network) error {
	pauseXCall := func(ctx context.Context, def app.Definition, config Config) error {
		return PauseXCall(ctx, def, config, XCallConfig{})
	}

	unpauseXCall := func(ctx context.Context, def app.Definition, config Config) error {
		return UnpauseXCall(ctx, def, config, XCallConfig{})
	}

	chain := randChain(network)

	err := forOne(ctx, def, chain, pauseXCall, checkXCallPaused(true))
	if err != nil {
		return errors.Wrap(err, "pause xcall")
	}

	err = forOne(ctx, def, chain, unpauseXCall, checkXCallPaused(false))
	if err != nil {
		return errors.Wrap(err, "unpause xcall")
	}

	err = forAll(ctx, def, network, pauseXCall, checkXCallPaused(true))
	if err != nil {
		return errors.Wrap(err, "pause all xcalls")
	}

	err = forAll(ctx, def, network, unpauseXCall, checkXCallPaused(false))
	if err != nil {
		return errors.Wrap(err, "unpause all xcalls")
	}

	return nil
}

// testPauseUnpauseXCallTo tests PauseXCall and UnpauseXCall commands (with XCallConfig.To).
func testPauseUnpauseXCallTo(ctx context.Context, def app.Definition, network netconf.Network) error {
	to := randChain(network)

	pauseXCallTo := func(ctx context.Context, def app.Definition, config Config) error {
		return PauseXCall(ctx, def, config, XCallConfig{To: to.Name})
	}

	unpauseXCallTo := func(ctx context.Context, def app.Definition, config Config) error {
		return UnpauseXCall(ctx, def, config, XCallConfig{To: to.Name})
	}

	chain := randChain(network)

	err := forOne(ctx, def, chain, pauseXCallTo, checkXCallToPaused(to, true))
	if err != nil {
		return errors.Wrap(err, "pause xcall to")
	}

	err = forOne(ctx, def, chain, unpauseXCallTo, checkXCallToPaused(to, false))
	if err != nil {
		return errors.Wrap(err, "unpause xcall to")
	}

	err = forAll(ctx, def, network, pauseXCallTo, checkXCallToPaused(to, true))
	if err != nil {
		return errors.Wrap(err, "pause all xcalls to")
	}

	err = forAll(ctx, def, network, unpauseXCallTo, checkXCallToPaused(to, false))
	if err != nil {
		return errors.Wrap(err, "unpause all xcalls to")
	}

	return nil
}

// testUpgradePortal tests UpgradePortal command.
func testUpgradePortal(ctx context.Context, def app.Definition, network netconf.Network) error {
	// no upgrade check - just make sure it doesn't error
	noCheck := func(context.Context, app.Definition, netconf.Chain) error { return nil }

	err := forOne(ctx, def, randChain(network), UpgradePortal, noCheck)
	if err != nil {
		return errors.Wrap(err, "upgrade portal")
	}

	err = forAll(ctx, def, network, UpgradePortal, noCheck)
	if err != nil {
		return errors.Wrap(err, "upgrade all portals")
	}

	return nil
}

// forOne runs an action & check configured for a single chain (Config{Chain: "name"}).
func forOne(
	ctx context.Context,
	def app.Definition,
	chain netconf.Chain,
	action func(context.Context, app.Definition, Config) error,
	check func(context.Context, app.Definition, netconf.Chain) error,
) error {
	if err := action(ctx, def, Config{Chain: chain.Name}); err != nil {
		return errors.Wrap(err, "act", "chain", chain.Name)
	}

	if err := check(ctx, def, chain); err != nil {
		return errors.Wrap(err, "check", "chain", chain.Name)
	}

	return nil
}

// forAll runs an action & check configured for all chains (Config{Chain: ""}).
func forAll(
	ctx context.Context,
	def app.Definition,
	network netconf.Network,
	action func(context.Context, app.Definition, Config) error,
	check func(context.Context, app.Definition, netconf.Chain) error,
) error {
	if err := action(ctx, def, Config{}); err != nil {
		return errors.Wrap(err, "act")
	}

	for _, chain := range network.EVMChains() {
		if err := check(ctx, def, chain); err != nil {
			return errors.Wrap(err, "check", "chain", chain.Name)
		}
	}

	return nil
}

func randChain(network netconf.Network) netconf.Chain {
	chains := network.EVMChains()
	//nolint:gosec // no need for secure randomneness
	return chains[rand.Intn(len(chains))]
}

func checkPortalPaused(expected bool) func(context.Context, app.Definition, netconf.Chain) error {
	return func(ctx context.Context, def app.Definition, chain netconf.Chain) error {
		backend, err := def.Backends().Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend")
		}

		portal, err := bindings.NewOmniPortal(chain.PortalAddress, backend)
		if err != nil {
			return errors.Wrap(err, "new portal")
		}

		paused, err := portal.IsPaused1(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.Wrap(err, "get paused")
		}

		if paused != expected {
			return errors.New("check paused", "chain", chain.Name, "paused", paused, "expected", expected)
		}

		return nil
	}
}

func checkXCallPaused(expected bool) func(context.Context, app.Definition, netconf.Chain) error {
	return func(ctx context.Context, def app.Definition, chain netconf.Chain) error {
		backend, err := def.Backends().Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend")
		}

		portal, err := bindings.NewOmniPortal(chain.PortalAddress, backend)
		if err != nil {
			return errors.Wrap(err, "new portal")
		}

		pauseAction, err := portal.ActionXCall(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.Wrap(err, "get paused")
		}

		paused, err := portal.IsPaused(&bind.CallOpts{Context: ctx}, pauseAction)
		if err != nil {
			return errors.Wrap(err, "get paused")
		}

		if paused != expected {
			return errors.New("check paused", "chain", chain.Name, "paused", paused, "expected", expected)
		}

		return nil
	}
}

func checkXCallToPaused(to netconf.Chain, expected bool) func(context.Context, app.Definition, netconf.Chain) error {
	return func(ctx context.Context, def app.Definition, chain netconf.Chain) error {
		backend, err := def.Backends().Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "get backend")
		}

		portal, err := bindings.NewOmniPortal(chain.PortalAddress, backend)
		if err != nil {
			return errors.Wrap(err, "new portal")
		}

		pauseAction, err := portal.ActionXCall(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.Wrap(err, "get paused")
		}

		paused, err := portal.IsPaused0(&bind.CallOpts{Context: ctx}, pauseAction, to.ID)
		if err != nil {
			return errors.Wrap(err, "get paused")
		}

		if paused != expected {
			return errors.New("check paused", "chain", chain.Name, "paused", paused, "expected", expected)
		}

		return nil
	}
}
