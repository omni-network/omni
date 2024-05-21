package data

import (
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/graph-gophers/graphql-go/relay"
)

type ChainsProvider struct {
	// chains are stored in a map where the key is the chain ID since the most common operation is to look up a chain by its ID.
	chains map[string]Chain
}

// NewChainsProvider creates a new [ChainsProvider] instance with the chains for the given network (e.g. "devnet", "testnet" or "mainnet").
func NewChainsProvider(network netconf.ID) *ChainsProvider {
	// Define the chains supported by each network
	networks := map[netconf.ID][]struct {
		id      uint64
		name    string
		logoURL string
	}{
		netconf.Devnet: {
			{id: 1651, name: "Omni Ephemeral", logoURL: "https://chainlist.org/unknown-logo.png"},
			{id: 1652, name: "Mock L1 Fast", logoURL: "https://chainlist.org/unknown-logo.png"},
			{id: 1653, name: "Mock L1 Slow", logoURL: "https://chainlist.org/unknown-logo.png"},
			{id: 1654, name: "Mock L2", logoURL: "https://chainlist.org/unknown-logo.png"},
		},
		netconf.Staging: {
			{id: 1651, name: "Omni Ephemeral", logoURL: "https://chainlist.org/unknown-logo.png"},
			{id: 11155420, name: "Op Sepolia", logoURL: "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg"},
			{id: 1653, name: "Mock L1 Slow", logoURL: "https://chainlist.org/unknown-logo.png"},
		},
		netconf.Testnet: {
			{id: 165, name: "Omni Testnet", logoURL: "https://chainlist.org/unknown-logo.png"},
			{id: 17000, name: "Holesky", logoURL: "https://icons.llamao.fi/icons/chains/rsz_ethereum.jpg"},
			{id: 421614, name: "Arb Sepolia", logoURL: "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg"},
			{id: 11155420, name: "Op Sepolia", logoURL: "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg"},
		},
		netconf.Mainnet: {
			{id: 1, name: "Ethereum", logoURL: "https://icons.llamao.fi/icons/chains/rsz_ethereum.jpg"},
			{id: 166, name: "Omni", logoURL: "https://chainlist.org/unknown-logo.png"},
			{id: 42161, name: "Arbitrum", logoURL: "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg"},
			{id: 10, name: "Optimism", logoURL: "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg"},
		},
	}

	list, ok := networks[network]
	if !ok {
		panic("unknown network: " + network)
	}

	chains := make(map[string]Chain, len(list))
	for _, c := range list {
		key := fmt.Sprintf("0x%x", c.id)
		chains[key] = Chain{
			ID:        relay.MarshalID("chain", c.id),
			ChainID:   hexutil.Big(*hexutil.MustDecodeBig(key)),
			DisplayID: Long(c.id),
			Name:      c.name,
			LogoURL:   c.logoURL,
		}
	}

	return &ChainsProvider{
		chains: chains,
	}
}

func (p *ChainsProvider) Chain(id string) (Chain, bool) {
	c, ok := p.chains[id]
	if !ok {
		return Chain{}, false
	}
	return c, true
}

func (p *ChainsProvider) ChainList() []Chain {
	var chains []Chain
	for _, c := range p.chains {
		chains = append(chains, c)
	}

	slices.SortFunc(chains, func(a, b Chain) int {
		if a.DisplayID < b.DisplayID {
			return -1
		} else if a.DisplayID > b.DisplayID {
			return 1
		}

		return 0
	})

	return chains
}
