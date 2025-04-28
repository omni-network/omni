package cmd

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func networkFromDef(ctx context.Context, def app.Definition) (netconf.Network, error) {
	endpoints := app.ExternalEndpoints(def)
	networkID := def.Testnet.Network

	portalReg, err := makePortalRegistry(ctx, networkID, endpoints)
	if err != nil {
		return netconf.Network{}, errors.Wrap(err, "portal registry")
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, networkID, portalReg, endpoints.Keys())
	if err != nil {
		return netconf.Network{}, errors.Wrap(err, "await onchain")
	}

	return network, nil
}

func makePortalRegistry(ctx context.Context, network netconf.ID, endpoints xchain.RPCEndpoints) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	rpc, err := endpoints.ByNameOrID(meta.Name, meta.ChainID)
	if err != nil {
		return nil, err
	}

	ethCl, err := ethclient.DialContext(ctx, meta.Name, rpc)
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}

// getSolverNetwork returns the solvernet network for the given definition, including HL chains.
func getSolverNetwork(ctx context.Context, def app.Definition) (netconf.Network, error) {
	network, err := networkFromDef(ctx, def)
	if err != nil {
		return netconf.Network{}, errors.Wrap(err, "network")
	}

	return solvernet.AddHLNetwork(network), nil
}

func getSolverEndpoints(networkID netconf.ID, def app.Definition) (xchain.RPCEndpoints, error) {
	endpoints := app.ExternalEndpoints(def)

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
