package data

import (
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/omni-network/omni/lib/evmchain"
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
		id          uint64
		name        string
		logoURL     string
		addrUrlFmt  string
		blockURLFmt string
		txURLFmt    string
	}{
		netconf.Devnet: {
			{id: uint64(evmchain.IDOmniEphemeral), name: "Omni Ephemeral", logoURL: "https://chainlist.org/unknown-logo.png", addrUrlFmt: "https://omni-ephemeral.dev/address/%s", blockURLFmt: "https://omni-ephemeral.dev/block/%d", txURLFmt: "https://omni-ephemeral.dev/tx/%s"},
			{id: uint64(evmchain.IDMockL1Fast), name: "Mock L1 Fast", logoURL: "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg", addrUrlFmt: "https://mock-l1-fast.dev/address/%s", blockURLFmt: "https://mock-l1-fast.dev/block/%d", txURLFmt: "https://mock-l1-fast.dev/tx/%s"},
			{id: uint64(evmchain.IDMockL2), name: "Mock L2", logoURL: "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg", addrUrlFmt: "https://mock-l2.dev/address/%s", blockURLFmt: "https://mock-l2.dev/block/%d", txURLFmt: "https://mock-l2.dev/tx/%s"},
			{id: uint64(evmchain.IDMockArb), name: "Mock Arb", logoURL: "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg", addrUrlFmt: "https://sepolia-optimism.etherscan.io/address/%s", blockURLFmt: "https://sepolia-optimism.etherscan.io/block/%d", txURLFmt: "https://sepolia-optimism.etherscan.io/tx/%s"},
			{id: uint64(evmchain.IDMockOp), name: "Mock Op", logoURL: "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg", addrUrlFmt: "https://sepolia.arbiscan.io/address/%s", blockURLFmt: "https://sepolia.arbiscan.io/block/%d", txURLFmt: "https://sepolia.arbiscan.io/tx/%s"},
		},
		netconf.Staging: {
			{id: uint64(evmchain.IDOmniEphemeral), name: "Omni Ephemeral", logoURL: "https://chainlist.org/unknown-logo.png", addrUrlFmt: "https://omni-ephemeral.dev/address/%s", blockURLFmt: "https://omni-ephemeral.dev/block/%d", txURLFmt: "https://omni-ephemeral.dev/tx/%s"},
			{id: uint64(evmchain.IDMockL1Slow), name: "Mock L1 Slow", logoURL: "https://chainlist.org/unknown-logo.png", addrUrlFmt: "https://mock-l1-slow.dev/address/%s", blockURLFmt: "https://mock-l1-slow.dev/block/%d", txURLFmt: "https://mock-l1-slow.dev/tx/%s"},
			{id: uint64(evmchain.IDOpSepolia), name: "Op Sepolia", logoURL: "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg", addrUrlFmt: "https://sepolia-optimism.etherscan.io/address/%s", blockURLFmt: "https://sepolia-optimism.etherscan.io/block/%d", txURLFmt: "https://sepolia-optimism.etherscan.io/tx/%s"},
		},
		netconf.Omega: {
			{id: uint64(evmchain.IDOmniOmega), name: "Omni Omega", logoURL: "https://chainlist.org/unknown-logo.png", addrUrlFmt: "", blockURLFmt: "", txURLFmt: ""},
			{id: uint64(evmchain.IDHolesky), name: "Holesky", logoURL: "https://icons.llamao.fi/icons/chains/rsz_ethereum.jpg", addrUrlFmt: "https://holesky.etherscan.io/address/%s", blockURLFmt: "https://holesky.etherscan.io/block/%d", txURLFmt: "https://holesky.etherscan.io/tx/%s"},
			{id: uint64(evmchain.IDArbSepolia), name: "Arb Sepolia", logoURL: "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg", addrUrlFmt: "https://sepolia.arbiscan.io/address/%s", blockURLFmt: "https://sepolia.arbiscan.io/block/%d", txURLFmt: "https://sepolia.arbiscan.io/tx/%s"},
			{id: uint64(evmchain.IDOpSepolia), name: "Op Sepolia", logoURL: "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg", addrUrlFmt: "https://sepolia-optimism.etherscan.io/address/%s", blockURLFmt: "https://sepolia-optimism.etherscan.io/block/%d", txURLFmt: "https://sepolia-optimism.etherscan.io/tx/%s"},
		},
		netconf.Mainnet: {
			{id: 1, name: "Ethereum", logoURL: "https://icons.llamao.fi/icons/chains/rsz_ethereum.jpg", addrUrlFmt: "https://etherscan.io/address/%s", blockURLFmt: "https://etherscan.io/block/%d", txURLFmt: "https://etherscan.io/tx/%s"},
			{id: 166, name: "Omni", logoURL: "https://chainlist.org/unknown-logo.png", addrUrlFmt: "", blockURLFmt: "", txURLFmt: ""},
			{id: 42161, name: "Arbitrum", logoURL: "https://icons.llamao.fi/icons/chains/rsz_arbitrum.jpg", addrUrlFmt: "https://arbiscan.io/address/%s", blockURLFmt: "https://arbiscan.io/block/%d", txURLFmt: "https://arbiscan.io/tx/%s"},
			{id: 10, name: "Optimism", logoURL: "https://icons.llamao.fi/icons/chains/rsz_optimism.jpg", addrUrlFmt: "https://optimistic.etherscan.io/address/%s", blockURLFmt: "https://optimistic.etherscan.io/block/%d", txURLFmt: "https://optimistic.etherscan.io/tx/%s"},
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
			AddrURLFmt:  c.addrUrlFmt,
			BlockURLFmt: c.blockURLFmt,
			TxURLFmt:    c.txURLFmt,
			ID:          relay.MarshalID("chain", c.id),
			ChainID:     hexutil.Big(*hexutil.MustDecodeBig(key)),
			DisplayID:   Long(c.id),
			Name:        c.name,
			LogoURL:     c.logoURL,
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
