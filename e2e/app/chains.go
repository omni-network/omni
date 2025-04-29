package app

import (
	"context"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// getSolverNetwork returns the solvernet network for the given definition, including HL chains.
func GetSolverNetwork(ctx context.Context, def Definition) (netconf.Network, error) {
	network := networkFromDef(def)
	return solvernet.AddHLNetwork(ctx, network), nil
}

// getSolverEndpoints returns the RPC endpoints for the given solvernet network, including HL chains.
func GetSolverEndpoints(networkID netconf.ID, def Definition) (xchain.RPCEndpoints, error) {
	endpoints := ExternalEndpoints(def)

	// extend endpoints w/ hl chains
	for _, chain := range solvernet.HLChains(networkID) {
		meta, ok := evmchain.MetadataByID(chain.ID)
		if !ok {
			return xchain.RPCEndpoints{}, errors.New("unknown chain", "chain_id", chain.ID)
		}

		rpc, ok := def.Cfg.RPCOverrides[meta.Name]
		if !ok {
			rpc = types.PublicRPCByName(meta.Name)
			if rpc == "" {
				return xchain.RPCEndpoints{}, errors.New("missing rpc override", "chain", meta.Name)
			}
		}

		endpoints[meta.Name] = rpc
	}

	return endpoints, nil
}

// GetSolverNetworkAndBackends returns the solvernet network and backends for the given definition.
func GetSolverNetworkAndBackends(ctx context.Context, def Definition, cmdName string) (netconf.Network, ethbackend.Backends, error) {
	network, err := GetSolverNetwork(ctx, def)
	if err != nil {
		return netconf.Network{}, ethbackend.Backends{}, errors.Wrap(err, "get network")
	}

	endpoints, err := GetSolverEndpoints(network.ID, def)
	if err != nil {
		return netconf.Network{}, ethbackend.Backends{}, errors.Wrap(err, "get endpoints")
	}

	fireCl, err := NewFireblocksClient(def.Cfg, network.ID, cmdName)
	if err != nil {
		return netconf.Network{}, ethbackend.Backends{}, errors.Wrap(err, "fireblocks client")
	}

	backends, err := ethbackend.FireBackendsFromNetwork(ctx, network, endpoints, fireCl)
	if err != nil {
		return netconf.Network{}, ethbackend.Backends{}, errors.Wrap(err, "fire backends")
	}

	return network, backends, nil
}
