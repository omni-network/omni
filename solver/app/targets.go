package app

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/solve/devapp"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/solver/types"
)

var (
	targets map[netconf.ID]types.Targets = map[netconf.ID]types.Targets{
		netconf.Devnet: {devapp.GetApp()},
	}
)

func targetFor(network netconf.ID, call bindings.SolveCall) (types.Target, error) {
	t, ok := targets[network]
	if !ok {
		return nil, errors.New("no targets for network", "network", network)
	}

	return t.ForCall(call)
}
