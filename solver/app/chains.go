package app

import (
	"context"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// AddSolverEndpoints returns the RPC endpoints for the given solvernet network, including HL chains.
func AddSolverEndpoints(ctx context.Context, networkID netconf.ID, endpoints xchain.RPCEndpoints, rpcOverrides map[string]string) (xchain.RPCEndpoints, error) {
	log.Debug(ctx, "Adding solver endpoints", "network", networkID, "endpoints", endpoints, "rpc_overrides", rpcOverrides)

	// extend endpoints w/ hl chains
	for _, chain := range solvernet.HLChains(networkID) {
		meta, ok := evmchain.MetadataByID(chain.ID)
		if !ok {
			return xchain.RPCEndpoints{}, errors.New("unknown chain", "chain_id", chain.ID)
		}

		rpc, ok := rpcOverrides[meta.Name]
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
