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

	omniDevnetChain = netconf.Chain{
		ID:   evmchain.IDOmniDevnet,
		Name: "omni_evm",
	}
	mockL1Chain = netconf.Chain{
		ID:   evmchain.IDMockL1,
		Name: "mock_l1",
	}
	mockL2Chain = netconf.Chain{
		ID:   evmchain.IDMockL2,
		Name: "mock_l2",
	}
	sepoliaChain = netconf.Chain{
		ID:   evmchain.IDSepolia,
		Name: "sepolia",
	}

	devnetNetwork = netconf.Network{
		ID: netconf.Devnet,
		Chains: []netconf.Chain{
			omniDevnetChain,
			mockL1Chain,
			mockL2Chain,
			sepoliaChain,
		},
	}
)

func makeRoutes() []TestRoute {
	var routes []TestRoute

	routes = append(routes, TestRoute{
		name:        "Omni EVM (Core-only)",
		sourceChain: omniDevnetChain,
		network:     devnetNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID: omniDevnetChain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.None,
				},
			},
			{
				ChainID: mockL1Chain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.OmniCore,
				},
			},
			{
				ChainID: mockL2Chain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.OmniCore,
				},
			},
			// Sepolia (Hyperlane) should be skipped
		},
	})

	routes = append(routes, TestRoute{
		name:        "Sepolia (Hyperlane-only)",
		sourceChain: sepoliaChain,
		network:     devnetNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID: omniDevnetChain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.Hyperlane,
				},
			},
			{
				ChainID: mockL1Chain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.Hyperlane,
				},
			},
			{
				ChainID: mockL2Chain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.Hyperlane,
				},
			},
			{
				ChainID: sepoliaChain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.None,
				},
			},
		},
	})

	routes = append(routes, TestRoute{
		name:        "Mock L1 (Both)",
		sourceChain: mockL1Chain,
		network:     devnetNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID: omniDevnetChain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.OmniCore,
				},
			},
			{
				ChainID: mockL1Chain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.None,
				},
			},
			{
				ChainID: mockL2Chain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.OmniCore,
				},
			},
			{
				ChainID: sepoliaChain.ID,
				Outbox:  dummyOutboxAddr,
				InboxConfig: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.Hyperlane,
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

			actualRoutes, err := getRoutes(tc.sourceChain, tc.network, tc.inboxAddr, tc.outboxAddr)
			require.NoError(t, err)
			require.ElementsMatch(t, tc.expectedRoutes, actualRoutes, "Routes mismatch")
		})
	}
}
