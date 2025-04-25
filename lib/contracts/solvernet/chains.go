package solvernet

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
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

// AddHLNetwork returns a copy of the network with hyperlane-secured chains added.
func AddHLNetwork(network netconf.Network) netconf.Network {
	return network.AddChains(HLChains(network.ID)...)
}
