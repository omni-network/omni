package cmd

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func networkFromDef(ctx context.Context, def app.Definition) (netconf.Network, error) {
	endpoints := app.ExternalEndpoints(def)
	networkID := def.Testnet.Network

	portalReg, err := makePortalRegistry(networkID, endpoints)
	if err != nil {
		return netconf.Network{}, errors.Wrap(err, "portal registry")
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, networkID, portalReg, endpoints.Keys())
	if err != nil {
		return netconf.Network{}, errors.Wrap(err, "await onchain")
	}

	return network, nil
}

func makePortalRegistry(network netconf.ID, endpoints xchain.RPCEndpoints) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	rpc, err := endpoints.ByNameOrID(meta.Name, meta.ChainID)
	if err != nil {
		return nil, err
	}

	ethCl, err := ethclient.Dial(meta.Name, rpc)
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
