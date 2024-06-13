package connect

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/cchain"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"
)

// Connector provider a simple abstraction to connect to the Omni network.
type Connector struct {
	Network    netconf.Network
	XProvider  xchain.Provider
	CProvider  cchain.Provider
	ethClients map[uint64]ethclient.Client
}

// Backend returns an ethbackend for the given chainID.
func (c Connector) Backend(chainID uint64) (*ethbackend.Backend, error) {
	chain, ok := c.Network.Chain(chainID)
	if !ok {
		return nil, errors.New("chain not found")
	}

	cl, ok := c.ethClients[chainID]
	if !ok {
		return nil, errors.New("ethclient not confired for chain")
	}

	return ethbackend.NewBackend(chain.Name, chainID, chain.BlockPeriod, cl)
}

// New returns a populated Connector for the given network.
// By default, it supports connecting to the omni evm and consensus chains since they have well known RPCs.
// To connect to other supported rollups, the RPC endpoints must be manually provided.
func New(ctx context.Context, netID netconf.ID, endpoints xchain.RPCEndpoints) (Connector, error) {
	if endpoints == nil {
		endpoints = make(xchain.RPCEndpoints)
	}

	omniEVMID := netID.Static().OmniExecutionChainID
	omniEVMMetadata, ok := evmchain.MetadataByID(omniEVMID)
	if !ok {
		return Connector{}, errors.New("omni evm metadata not found")
	}
	if _, err := endpoints.ByNameOrID(omniEVMMetadata.Name, omniEVMMetadata.ChainID); err != nil {
		endpoints[omniEVMMetadata.Name] = netID.Static().ExecutionRPC()
	}

	omniCons := netID.Static().OmniConsensusChain()
	if _, err := endpoints.ByNameOrID(omniCons.Name, omniCons.ID); err != nil {
		endpoints[omniCons.Name] = netID.Static().ConsensusRPC()
	}

	portalReg, err := makePortalRegistry(netID, endpoints)
	if err != nil {
		return Connector{}, err
	}

	network, err := netconf.AwaitOnChain(ctx, netID, portalReg, nil)
	if err != nil {
		return Connector{}, err
	}

	ethClients := make(map[uint64]ethclient.Client)
	for _, chain := range network.Chains {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			rpc = "unknown"
		} else {
			ethCl, err := ethclient.Dial(chain.Name, rpc)
			if err != nil {
				return Connector{}, errors.Wrap(err, "dial eth client")
			}
			ethClients[chain.ID] = ethCl
		}

		log.Info(ctx, "Detected supported chain", "chain", chain.Name, "id", chain.ID, "rpc", rpc)
	}

	// Connect to the halo cometBFT RPC server.
	cometCl, err := rpchttp.New(endpoints[omniCons.Name], "/websocket")
	if err != nil {
		return Connector{}, errors.Wrap(err, "comet rpc client")
	}

	cprov := cprovider.NewABCIProvider(cometCl, netID, netconf.ChainVersionNamer(netID))

	xprov := xprovider.New(network, ethClients, cprov)

	return Connector{
		Network:    network,
		XProvider:  xprov,
		CProvider:  cprov,
		ethClients: ethClients,
	}, nil
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
