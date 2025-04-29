package solvernet

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// hlChains define solvernet chains secured by hyperlane.
var hlChains = map[netconf.ID][]uint64{
	// TODO(zodomo): add hyperlane chains here
}

// HLChains returns the list of hyperlane-secured chains for a given solvernet network.
func HLChains(network netconf.ID) []netconf.Chain {
	var resp []netconf.Chain
	for _, chainID := range hlChains[network] {
		chain, ok := evmchain.MetadataByID(chainID)
		if !ok {
			panic(errors.New("unknown chain", "chain_id", chainID))
		}
		resp = append(resp, netconf.Chain{
			ID:          chain.ChainID,
			Name:        chain.Name,
			BlockPeriod: chain.BlockPeriod,
		})
	}

	return resp
}

// IsHLChain returns true if the solvernet chain is secured by hyperlane.
func IsHLChain(chainID uint64) bool {
	for _, ids := range hlChains {
		for _, id := range ids {
			if id == chainID {
				return true
			}
		}
	}

	return false
}

// FilterByBackends returns an HL chain selector that excludes chains not in backends.
// Useful when needing to deploy contracts to configured backends.
func FilterByBackends(backends ethbackend.Backends) func(netconf.ID, netconf.Chain) bool {
	return func(_ netconf.ID, chain netconf.Chain) bool {
		_, err := backends.Backend(chain.ID)
		return err == nil
	}
}

// FilterByContracts returns an HL chain selector that excludes chains without inbox contracts deployed.
// Note this also excludes chains without endpoints, or with any other error fetching inbox DeployedAt.
func FilterByContracts(ctx context.Context, endpoints xchain.RPCEndpoints) func(netconf.ID, netconf.Chain) bool {
	return func(network netconf.ID, chain netconf.Chain) bool {
		endpoint, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return false
		}

		ethCl, err := ethclient.DialContext(ctx, chain.Name, endpoint)
		if err != nil {
			return false
		}

		addrs, err := contracts.GetAddresses(ctx, network)
		if err != nil {
			return false
		}

		contract, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, ethCl)
		if err != nil {
			return false
		}

		_, err = contract.DeployedAt(&bind.CallOpts{Context: ctx})

		return err == nil
	}
}

// AddHLNetwork returns a copy of the network with hyperlane-secured chains added.
// Optional selector functions can be provided to filter the chains.
func AddHLNetwork(ctx context.Context, network netconf.Network, selectors ...func(netconf.ID, netconf.Chain) bool) netconf.Network {
	chains := HLChains(network.ID)

	// Filter out chains using provided selectors
	var included, excluded []string
	for _, selector := range selectors {
		var selected []netconf.Chain
		for _, chain := range chains {
			if selector(network.ID, chain) {
				selected = append(selected, chain)
			} else {
				excluded = append(excluded, chain.Name)
			}
		}
		chains = selected
	}

	for _, chain := range chains {
		included = append(included, chain.Name)
	}

	log.Debug(ctx, "Adding hyperlane chains", "included", included, "excluded", excluded)

	return network.AddChains(chains...)
}
