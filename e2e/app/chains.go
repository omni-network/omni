package app

import (
	"context"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
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

// AddSolverNetworkAndBackends returns the solvernet network and backends for the given definition.
func AddSolverNetworkAndBackends(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints, defCfg DefinitionConfig, cmdName string) (netconf.Network, ethbackend.Backends, error) {
	log.Debug(ctx, "Adding solver network and backends", "network", network.ID)

	endpoints, err := AddSolverEndpoints(ctx, network.ID, endpoints, defCfg.RPCOverrides)
	if err != nil {
		return netconf.Network{}, ethbackend.Backends{}, errors.Wrap(err, "get endpoints")
	}

	network = solvernet.AddHLNetwork(ctx, network, solvernet.FilterByEndpoints(endpoints))

	var backends ethbackend.Backends
	if network.ID == netconf.Devnet {
		log.Debug(ctx, "Adding devnet backends", "network", network.ID)
		pks := append(eoa.DevPrivateKeys(), eoa.DevPrivateKeys()...)
		backends, err = ethbackend.BackendsFromNetwork(ctx, network, endpoints, pks...)
		if err != nil {
			return netconf.Network{}, ethbackend.Backends{}, errors.Wrap(err, "backends from network")
		}
	} else {
		log.Debug(ctx, "Adding fireblocks backends", "network", network.ID)
		fireCl, err := NewFireblocksClient(defCfg, network.ID, cmdName)
		if err != nil {
			return netconf.Network{}, ethbackend.Backends{}, errors.Wrap(err, "fireblocks client")
		}

		backends, err = ethbackend.FireBackendsFromNetwork(ctx, network, endpoints, fireCl)
		if err != nil {
			return netconf.Network{}, ethbackend.Backends{}, errors.Wrap(err, "fire backends")
		}
	}

	return network, backends, nil
}

// networkFromDef returns the network configuration from the definition.
// Note that this does not panic as it does in definition.go by manually setting portal addresses without deploy height.
func hlNetworkFromDef(ctx context.Context, def Definition) (netconf.Network, error) {
	var chains []netconf.Chain

	addrs, err := contracts.GetAddresses(ctx, def.Testnet.Network)
	if err != nil {
		return netconf.Network{}, errors.Wrap(err, "get addresses")
	}

	newChain := func(chain types.EVMChain) netconf.Chain {
		var portalAddress common.Address
		if !solvernet.IsHLChain(chain.ChainID) {
			portalAddress = addrs.Portal
		}

		return netconf.Chain{
			ID:              chain.ChainID,
			Name:            chain.Name,
			BlockPeriod:     chain.BlockPeriod,
			Shards:          chain.Shards,
			AttestInterval:  chain.AttestInterval(def.Testnet.Network),
			PortalAddress:   portalAddress,
			DeployHeight:    0, // Portal height isn't needed here as we aren't deploying OmniPortal contracts to Hyperlane networks
			HasEmitPortal:   true,
			HasSubmitPortal: true,
		}
	}

	// Add all public chains
	for _, public := range def.Testnet.PublicChains {
		chains = append(chains, newChain(public.Chain()))
	}

	// Connect to a proper omni_evm that isn't unavailable
	omniEVM := def.Testnet.BroadcastOmniEVM()
	chains = append(chains, newChain(omniEVM.Chain))

	// Add omni consensus chain
	chains = append(chains, def.Testnet.Network.Static().OmniConsensusChain())

	// Add all anvil chains
	for _, anvil := range def.Testnet.AnvilChains {
		chains = append(chains, newChain(anvil.Chain))
	}

	return netconf.Network{
		ID:     def.Testnet.Network,
		Chains: chains,
	}, nil
}
