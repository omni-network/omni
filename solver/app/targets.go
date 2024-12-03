package app

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/solve/devapp"
	"github.com/omni-network/omni/e2e/solve/symbiotic"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/solver/types"
)

var targetsByNetwork = map[netconf.ID][]types.Target{
	netconf.Devnet: {devapp.MustGetApp(netconf.Devnet), symbiotic.MustGetApp(netconf.Devnet)},
}

// getTarget returns the target for the given network and call.
func getTarget(network netconf.ID, call bindings.SolveCall) (types.Target, error) {
	targets, ok := targetsByNetwork[network]
	if !ok {
		return nil, errors.New("no targets for network", "network", network)
	}

	var resp *types.Target
	for _, target := range targets {
		if target.ChainID() == call.DestChainId && target.Address() == call.Target {
			if resp != nil {
				return nil, errors.New("multiple targets found")
			}
			resp = &target
		}
	}

	if resp == nil {
		return nil, errors.New("no target found")
	}

	return *resp, nil
}
