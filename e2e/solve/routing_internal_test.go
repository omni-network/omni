package solve

import (
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

type TestRoute struct {
	name           string
	sourceChain    netconf.Chain
	network        netconf.Network
	inboxAddr      common.Address
	outboxAddr     common.Address
	expectedRoutes []Route
}

var (
	dummyInboxAddr  = common.HexToAddress("0x1111111111111111111111111111111111111111")
	dummyOutboxAddr = common.HexToAddress("0x2222222222222222222222222222222222222222")

	// Core Only.
	omniStaging = netconf.Chain{
		ID:   evmchain.IDOmniDevnet,
		Name: "omni_evm",
	}

	// Core + Hyperlane.
	opSepolia = netconf.Chain{
		ID:   evmchain.IDOpSepolia,
		Name: "op_sepolia",
	}

	// Core + Hyperlane.
	arbSepolia = netconf.Chain{
		ID:   evmchain.IDArbSepolia,
		Name: "arb_sepolia",
	}

	// Hyperlane only.
	sepoliaChain = netconf.Chain{
		ID:   evmchain.IDSepolia,
		Name: "sepolia",
	}

	network = netconf.Network{
		ID: netconf.Staging,
		Chains: []netconf.Chain{
			omniStaging,
			opSepolia,
			arbSepolia,
			sepoliaChain,
		},
	}
)

func makeRoutes() []TestRoute {
	var routes []TestRoute

	routes = append(routes, TestRoute{
		name:        "Omni EVM (Core-only)",
		sourceChain: omniStaging,
		network:     network,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID:           omniStaging.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderNone,
				},
			},
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           arbSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
		},
	})

	routes = append(routes, TestRoute{
		name:        "Sepolia (Hyperlane-only)",
		sourceChain: sepoliaChain,
		network:     network,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			// Omni EVM (Core) should be skipped
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
				},
			},
			{
				ChainID:           arbSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
				},
			},
			{
				ChainID:           sepoliaChain.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderNone,
				},
			},
		},
	})

	routes = append(routes, TestRoute{
		name:        "OP Sepolia (Core + Hyperlane)",
		sourceChain: opSepolia,
		network:     network,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID:           omniStaging.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderNone,
				},
			},
			{
				ChainID:           arbSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           sepoliaChain.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
				},
			},
		},
	})

	return routes
}

func TestGetRoutes(t *testing.T) {
	t.Parallel()

	testCases := makeRoutes()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualRoutes := getRoutes(tc.sourceChain, tc.network, tc.inboxAddr, tc.outboxAddr)
			require.ElementsMatch(t, tc.expectedRoutes, actualRoutes, "Routes mismatch")
		})
	}
}
