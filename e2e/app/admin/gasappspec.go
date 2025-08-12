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

// gasAppSpec defines the gas app spec for each network.
// To update gas app spec, update this map.
// Then run `ensure-gas-app-spec` to apply the changes.
var gasAppSpec = map[netconf.ID]NetworkGasAppSpec{
	netconf.Devnet: {Global: DefaultGasAppSpec()},
	netconf.Staging: {Global: GasAppSpec{
		PauseGasPump:    true,
		PauseGasStation: true,
	}},
	netconf.Omega: {Global: GasAppSpec{
		PauseGasPump:    true,
		PauseGasStation: true,
	}},
	netconf.Mainnet: {Global: GasAppSpec{
		PauseGasPump:    true,
		PauseGasStation: true,
	}},
}

// GasAppSpec is the specification for the gas app.
type GasAppSpec struct {
	// PauseGasPump indicates that the gas pump should be paused.
	PauseGasPump bool

	// PauseGasStation indicates that the gas station should be paused.
	PauseGasStation bool
}

// GasAppDirectives defines updates required on the gas app to match a spec.
type GasAppDirectives struct {
	// PauseGasPump indicates that the gas pump should be paused.
	PauseGasPump bool

	// UnpauseGasPump indicates that the gas pump should be unpaused.
	UnpauseGasPump bool

	// PauseGasStation indicates that the gas station should be paused.
	PauseGasStation bool

	// UnpauseGasStation indicates that the gas station should be unpaused.
	UnpauseGasStation bool
}

// NetworkGasAppSpec defines the gas app spec for a network.
type NetworkGasAppSpec struct {
	// Global is the gas app spec, maintained on all chains.
	Global GasAppSpec

	// ChainOverrides overrides the global spec for specific chains. Must specify full spec, not just diff.
	ChainOverrides map[uint64]*GasAppSpec
}

// DefaultGasAppSpec returns a default gas app spec with nothing paused.
func DefaultGasAppSpec() GasAppSpec {
	return GasAppSpec{
		PauseGasPump:    false,
		PauseGasStation: false,
	}
}

// Verify checks that there are no conflicting specifications.
func (GasAppSpec) Verify() error {
	// Gas pump and gas station can both be paused since they exist on different chains
	// No validation conflicts exist for this spec
	return nil
}

// Verify checks that there are no conflicting directives.
func (d GasAppDirectives) Verify() error {
	if d.PauseGasPump && d.UnpauseGasPump {
		return errors.New("cannot pause and unpause gas pump")
	}

	if d.PauseGasStation && d.UnpauseGasStation {
		return errors.New("cannot pause and unpause gas station")
	}

	return nil
}

// EnsureGasAppSpec ensures that live gas app contracts are configured as per the local spec.
func EnsureGasAppSpec(ctx context.Context, def app.Definition, cfg Config, localSpecOverride *GasAppSpec) error {
	return setup(def, cfg).run(ctx, func(ctx context.Context, s shared, c chain) error {
		local, err := localGasAppSpec(s.testnet.Network, c.ChainID)
		if err != nil {
			return errors.Wrap(err, "get local gas app spec", "chain", c.Name)
		}

		if localSpecOverride != nil {
			local = *localSpecOverride
		}

		ethCl, err := ethclient.DialContext(ctx, c.Name, c.RPCEndpoint)
		if err != nil {
			return errors.Wrap(err, "dial eth client", "chain", c.Name)
		}

		addrs, err := contracts.GetAddresses(ctx, s.testnet.Network)
		if err != nil {
			return errors.Wrap(err, "get addresses", "chain", c.Name)
		}

		live, err := liveGasAppSpec(ctx, c.Name, addrs.GasPump, addrs.GasStation, ethCl)
		if err != nil {
			return errors.Wrap(err, "get live gas app spec", "chain", c.Name)
		}

		directives, err := makeGasAppDirectives(local, live)
		if err != nil {
			return errors.Wrap(err, "make directives", "chain", c.Name)
		}

		return runGasAppDirectives(ctx, s, c, addrs.GasPump, addrs.GasStation, directives)
	})
}

// makeGasAppDirectives returns the directives required to bring a gas app in line with the local spec.
func makeGasAppDirectives(local, live GasAppSpec) (GasAppDirectives, error) {
	if err := local.Verify(); err != nil {
		return GasAppDirectives{}, errors.Wrap(err, "verify local spec")
	}

	if err := live.Verify(); err != nil {
		return GasAppDirectives{}, errors.Wrap(err, "verify live spec")
	}

	pauseGasPump := local.PauseGasPump && !live.PauseGasPump
	unpauseGasPump := !local.PauseGasPump && live.PauseGasPump
	pauseGasStation := local.PauseGasStation && !live.PauseGasStation
	unpauseGasStation := !local.PauseGasStation && live.PauseGasStation

	return GasAppDirectives{
		PauseGasPump:      pauseGasPump,
		UnpauseGasPump:    unpauseGasPump,
		PauseGasStation:   pauseGasStation,
		UnpauseGasStation: unpauseGasStation,
	}, nil
}

// localGasAppSpec returns the configured, local gas app spec for a chain.
func localGasAppSpec(network netconf.ID, chainID uint64) (GasAppSpec, error) {
	net, ok := gasAppSpec[network]
	if !ok {
		return GasAppSpec{}, errors.New("no spec for network", "network", network)
	}

	spec := net.Global
	if net.ChainOverrides != nil && net.ChainOverrides[chainID] != nil {
		spec = *net.ChainOverrides[chainID]
	}

	if err := spec.Verify(); err != nil {
		return GasAppSpec{}, errors.Wrap(err, "verify spec", "network", network, "chain", chainID)
	}

	return spec, nil
}

// liveGasAppSpec returns the live, on-chain gas app spec for a chain.
func liveGasAppSpec(ctx context.Context, chainName string, gasPumpAddr, gasStationAddr common.Address, ethCl ethclient.Client) (GasAppSpec, error) {
	callopts := &bind.CallOpts{Context: ctx}

	// If Omni EVM, only check gas station
	if chainName == omniEVMName {
		gasStation, err := bindings.NewOmniGasStation(gasStationAddr, ethCl)
		if err != nil {
			return GasAppSpec{}, errors.Wrap(err, "new gas station contract", "chain", chainName)
		}

		gasStationPaused, err := gasStation.Paused(callopts)
		if err != nil {
			return GasAppSpec{}, errors.Wrap(err, "get gas station paused state", "chain", chainName)
		}

		return GasAppSpec{
			PauseGasPump:    false, // Gas pump doesn't exist on Omni EVM
			PauseGasStation: gasStationPaused,
		}, nil
	}

	// For all other chains, only check gas pump
	gasPump, err := bindings.NewOmniGasPump(gasPumpAddr, ethCl)
	if err != nil {
		return GasAppSpec{}, errors.Wrap(err, "new gas pump contract", "chain", chainName)
	}

	gasPumpPaused, err := gasPump.Paused(callopts)
	if err != nil {
		return GasAppSpec{}, errors.Wrap(err, "get gas pump paused state", "chain", chainName)
	}

	return GasAppSpec{
		PauseGasPump:    gasPumpPaused,
		PauseGasStation: false, // Gas station doesn't exist on non-Omni chains
	}, nil
}

// runGasAppDirectives applies GasAppDirectives on chain.
func runGasAppDirectives(ctx context.Context, s shared, c chain, gasPumpAddr, gasStationAddr common.Address, directives GasAppDirectives) error {
	if err := directives.Verify(); err != nil {
		return errors.Wrap(err, "verify directives", "chain", c.Name)
	}

	if isEmpty(directives) {
		log.Info(ctx, "No directives to apply", "chain", c.Name)
		return nil
	}

	if c.Name == omniEVMName {
		return runGasStationDirectives(ctx, s, c, gasStationAddr, directives)
	}

	return runGasPumpDirectives(ctx, s, c, gasPumpAddr, directives)
}

// runGasStationDirectives handles gas station directives for Omni EVM.
func runGasStationDirectives(ctx context.Context, s shared, c chain, gasStationAddr common.Address, directives GasAppDirectives) error {
	if directives.PauseGasStation {
		if err := pauseGasStation(ctx, s, c, gasStationAddr); err != nil {
			return errors.Wrap(err, "pause gas station", "chain", c.Name)
		}
	}

	if directives.UnpauseGasStation {
		if err := unpauseGasStation(ctx, s, c, gasStationAddr); err != nil {
			return errors.Wrap(err, "unpause gas station", "chain", c.Name)
		}
	}

	return nil
}

// runGasPumpDirectives handles gas pump directives for non-Omni chains.
func runGasPumpDirectives(ctx context.Context, s shared, c chain, gasPumpAddr common.Address, directives GasAppDirectives) error {
	if directives.PauseGasPump {
		if err := pauseGasPump(ctx, s, c, gasPumpAddr); err != nil {
			return errors.Wrap(err, "pause gas pump", "chain", c.Name)
		}
	}

	if directives.UnpauseGasPump {
		if err := unpauseGasPump(ctx, s, c, gasPumpAddr); err != nil {
			return errors.Wrap(err, "unpause gas pump", "chain", c.Name)
		}
	}

	return nil
}
