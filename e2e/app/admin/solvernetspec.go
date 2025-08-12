package admin

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// solverNetSpec defines SolverNet spec per network.
// To update SolverNet spec, update this map.
// Then run `ensure-solvernet-spec` to apply the changes.
var solverNetSpec = map[netconf.ID]NetworkSolverNetSpec{
	netconf.Devnet: {Global: DefaultSolverNetSpec()},
	netconf.Staging: {Global: SolverNetSpec{
		PauseAll:   false,
		PauseOpen:  true,
		PauseClose: false,
	}},
	netconf.Omega: {Global: SolverNetSpec{
		PauseAll:   false,
		PauseOpen:  true,
		PauseClose: false,
	}},
	netconf.Mainnet: {Global: SolverNetSpec{
		PauseAll:   false,
		PauseOpen:  true,
		PauseClose: false,
	}},
}

// SolverNetSpec is the specification for the SolverNetInbox contract.
type SolverNetSpec struct {
	// PauseAll indicates that open and close actions on the SolverNetInbox should be paused.
	PauseAll bool

	// PauseOpen indicates that opening new orders should be paused.
	PauseOpen bool

	// PauseClose indicates that closing orders should be paused.
	PauseClose bool
}

// SolverNetDirectives defines updates required on a SolverNetInbox to match a spec.
type SolverNetDirectives struct {
	// PauseAll indicates that open and close actions on the SolverNetInbox should be paused.
	PauseAll bool

	// UnpauseAll indicates that open and close actions on the SolverNetInbox should be unpaused.
	UnpauseAll bool

	// PauseOpen indicates that opening new orders should be paused.
	PauseOpen bool

	// UnpauseOpen indicates that opening new orders should be unpaused.
	UnpauseOpen bool

	// PauseClose indicates that closing orders should be paused.
	PauseClose bool

	// UnpauseClose indicates that closing orders should be unpaused.
	UnpauseClose bool
}

// NetworkSolverNetSpec defines the SolverNet spec for a network.
type NetworkSolverNetSpec struct {
	// Global is the SolverNet spec maintained on all chains.
	Global SolverNetSpec

	// ChainOverrides overrides the global spec for specific chains. Must specify full spec, not just diff.
	ChainOverrides map[uint64]*SolverNetSpec
}

// DefaultSolverNetSpec returns a default SolverNet spec with nothing paused.
func DefaultSolverNetSpec() SolverNetSpec {
	return SolverNetSpec{
		PauseAll:   false,
		PauseOpen:  false,
		PauseClose: false,
	}
}

// EnsureSolverNetSpec ensures that live SolverNetInbox contracts are configured as per the local spec.
func EnsureSolverNetSpec(ctx context.Context, def app.Definition, cfg Config, localSpecOverride *SolverNetSpec) error {
	return setup(def, cfg).run(ctx, func(ctx context.Context, s shared, c chain) error {
		addrs, err := contracts.GetAddresses(ctx, s.testnet.Network)
		if err != nil {
			return errors.Wrap(err, "get addresses", "chain", c.Name)
		}

		local, err := localSolverNetSpec(s.testnet.Network, c.ChainID)
		if err != nil {
			return errors.Wrap(err, "get local solvernet spec", "chain", c.Name)
		}

		if localSpecOverride != nil {
			local = *localSpecOverride
		}

		ethCl, err := ethclient.DialContext(ctx, c.Name, c.RPCEndpoint)
		if err != nil {
			return errors.Wrap(err, "dial eth client", "chain", c.Name)
		}

		live, err := liveSolverNetSpec(ctx, addrs.SolverNetInbox, ethCl)
		if err != nil {
			return errors.Wrap(err, "get live solvernet spec", "chain", c.Name)
		}

		directives, err := makeSolverNetDirectives(local, live)
		if err != nil {
			return errors.Wrap(err, "make directives", "chain", c.Name)
		}

		return runSolverNetDirectives(ctx, s, c, addrs.SolverNetInbox, directives)
	})
}

// makeSolverNetDirectives returns the directives required to bring a SolverNet inbox in line with the local spec.
func makeSolverNetDirectives(local, live SolverNetSpec) (SolverNetDirectives, error) {
	if err := local.Verify(); err != nil {
		return SolverNetDirectives{}, errors.Wrap(err, "verify local spec")
	}

	if err := live.Verify(); err != nil {
		return SolverNetDirectives{}, errors.Wrap(err, "verify live spec")
	}

	// If local.PauseAll is true but live.PauseAll is false, we need to pause all
	pauseAll := local.PauseAll && !live.PauseAll
	// If local.PauseAll is false but live.PauseAll is true, we need to unpause all
	unpauseAll := !local.PauseAll && live.PauseAll

	// If we're not dealing with pause/unpause all, check individual pause states
	var pauseOpen, unpauseOpen, pauseClose, unpauseClose bool

	if !local.PauseAll && !live.PauseAll {
		// If local.PauseOpen is true but live.PauseOpen is false, we need to pause open
		pauseOpen = local.PauseOpen && !live.PauseOpen
		// If local.PauseOpen is false but live.PauseOpen is true, we need to unpause open
		unpauseOpen = !local.PauseOpen && live.PauseOpen
		// If local.PauseClose is true but live.PauseClose is false, we need to pause close
		pauseClose = local.PauseClose && !live.PauseClose
		// If local.PauseClose is false but live.PauseClose is true, we need to unpause close
		unpauseClose = !local.PauseClose && live.PauseClose
	}

	return SolverNetDirectives{
		PauseAll:     pauseAll,
		UnpauseAll:   unpauseAll,
		PauseOpen:    pauseOpen,
		UnpauseOpen:  unpauseOpen,
		PauseClose:   pauseClose,
		UnpauseClose: unpauseClose,
	}, nil
}

// Verify checks that there are no conflicting specifications.
func (s SolverNetSpec) Verify() error {
	if s.PauseAll && (s.PauseOpen || s.PauseClose) {
		return errors.New("if pause-all, no other pause flags should be specified")
	}

	return nil
}

// Verify checks that there are no conflicting directives.
func (d SolverNetDirectives) Verify() error {
	if d.PauseAll && d.UnpauseAll {
		return errors.New("cannot pause and unpause all")
	}

	if d.PauseAll && (d.PauseOpen || d.PauseClose) {
		return errors.New("if pause-all, no other pause actions should be specified")
	}

	if d.PauseOpen && d.UnpauseOpen {
		return errors.New("cannot pause and unpause open")
	}

	if d.PauseClose && d.UnpauseClose {
		return errors.New("cannot pause and unpause close")
	}

	return nil
}

// localSolverNetSpec returns the configured, local SolverNet spec for a chain.
func localSolverNetSpec(network netconf.ID, chainID uint64) (SolverNetSpec, error) {
	net, ok := solverNetSpec[network]
	if !ok {
		return SolverNetSpec{}, errors.New("no spec for network", "network", network)
	}

	spec := net.Global
	if net.ChainOverrides != nil && net.ChainOverrides[chainID] != nil {
		spec = *net.ChainOverrides[chainID]
	}

	if err := spec.Verify(); err != nil {
		return SolverNetSpec{}, errors.Wrap(err, "verify spec", "network", network, "chain", chainID)
	}

	return spec, nil
}

// liveSolverNetSpec returns the live, on-chain SolverNet spec for a chain.
func liveSolverNetSpec(ctx context.Context, inboxAddr common.Address, ethCl ethclient.Client) (SolverNetSpec, error) {
	inbox, err := bindings.NewSolverNetInbox(inboxAddr, ethCl)
	if err != nil {
		return SolverNetSpec{}, errors.Wrap(err, "new solvernetinbox contract")
	}

	pauseState, err := inbox.PauseState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return SolverNetSpec{}, errors.Wrap(err, "get pause state")
	}

	// Parse the pause state
	// 0 = no pause, 1 = open paused, 2 = close paused, 3 = all paused
	switch pauseState {
	case 0: // NONE_PAUSED
		return SolverNetSpec{
			PauseAll:   false,
			PauseOpen:  false,
			PauseClose: false,
		}, nil
	case 1: // OPEN_PAUSED
		return SolverNetSpec{
			PauseAll:   false,
			PauseOpen:  true,
			PauseClose: false,
		}, nil
	case 2: // CLOSE_PAUSED
		return SolverNetSpec{
			PauseAll:   false,
			PauseOpen:  false,
			PauseClose: true,
		}, nil
	case 3: // ALL_PAUSED
		return SolverNetSpec{
			PauseAll:   true,
			PauseOpen:  false,
			PauseClose: false,
		}, nil
	default:
		return SolverNetSpec{}, errors.New("unknown pause state", "state", pauseState)
	}
}

// runSolverNetDirectives applies SolverNetDirectives on chain.
func runSolverNetDirectives(ctx context.Context, s shared, c chain, addr common.Address, directives SolverNetDirectives) error {
	if err := directives.Verify(); err != nil {
		return errors.Wrap(err, "verify directives", "chain", c.Name)
	}

	if isEmpty(directives) {
		log.Info(ctx, "No directives to apply", "chain", c.Name)
		return nil
	}

	if directives.PauseAll {
		if err := pauseSolverNetAll(ctx, s, c, addr, true); err != nil {
			return errors.Wrap(err, "pause solvernet all", "chain", c.Name)
		}
	}

	if directives.UnpauseAll {
		if err := pauseSolverNetAll(ctx, s, c, addr, false); err != nil {
			return errors.Wrap(err, "unpause solvernet all", "chain", c.Name)
		}
	}

	if directives.PauseOpen {
		if err := pauseSolverNetOpen(ctx, s, c, addr, true); err != nil {
			return errors.Wrap(err, "pause solvernet open", "chain", c.Name)
		}
	}

	if directives.UnpauseOpen {
		if err := pauseSolverNetOpen(ctx, s, c, addr, false); err != nil {
			return errors.Wrap(err, "unpause solvernet open", "chain", c.Name)
		}
	}

	if directives.PauseClose {
		if err := pauseSolverNetClose(ctx, s, c, addr, true); err != nil {
			return errors.Wrap(err, "pause solvernet close", "chain", c.Name)
		}
	}

	if directives.UnpauseClose {
		if err := pauseSolverNetClose(ctx, s, c, addr, false); err != nil {
			return errors.Wrap(err, "unpause solvernet close", "chain", c.Name)
		}
	}

	return nil
}
