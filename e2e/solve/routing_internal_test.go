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
	omniOmega = netconf.Chain{
		ID:   evmchain.IDOmniOmega,
		Name: "omni_evm",
	}

	omniStaging = netconf.Chain{
		ID:   evmchain.IDOmniDevnet,
		Name: "omni_evm",
	}

	// Hyperlane only.
	sepoliaChain = netconf.Chain{
		ID:   evmchain.IDSepolia,
		Name: "sepolia",
	}

	// Core + Hyperlane.
	holesky = netconf.Chain{
		ID:   evmchain.IDHolesky,
		Name: "holesky",
	}

	baseSepolia = netconf.Chain{
		ID:   evmchain.IDBaseSepolia,
		Name: "base_sepolia",
	}

	arbSepolia = netconf.Chain{
		ID:   evmchain.IDArbSepolia,
		Name: "arb_sepolia",
	}

	opSepolia = netconf.Chain{
		ID:   evmchain.IDOpSepolia,
		Name: "op_sepolia",
	}

	// Networks.
	stagingNetwork = netconf.Network{
		ID: netconf.Staging,
		Chains: []netconf.Chain{
			omniStaging,
			holesky,
			baseSepolia,
			arbSepolia,
			sepoliaChain,
			opSepolia,
		},
	}

	omegaNetwork = netconf.Network{
		ID: netconf.Omega,
		Chains: []netconf.Chain{
			omniOmega,
			holesky,
			baseSepolia,
			arbSepolia,
			sepoliaChain,
			opSepolia,
		},
	}
)

func makeRoutes() []TestRoute {
	var routes []TestRoute

	// --- Staging Network Test Cases ---

	// Source: Omni Staging (Core-only)
	routes = append(routes, TestRoute{
		name:        "Omni Staging (Core-only) to Staging Network",
		sourceChain: omniStaging,
		network:     stagingNetwork,
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
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			// No route to sepoliaChain (HL-only) from omniStaging (Core-only)
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
		},
	})

	// Source: Holesky (Core+HL)
	routes = append(routes, TestRoute{
		name:        "Holesky (Core+HL) to Staging Network",
		sourceChain: holesky,
		network:     stagingNetwork,
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
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderNone,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			{
				ChainID:           sepoliaChain.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
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
		},
	})

	// Source: Base Sepolia (Core+HL)
	routes = append(routes, TestRoute{
		name:        "Base Sepolia (Core+HL) to Staging Network",
		sourceChain: baseSepolia,
		network:     stagingNetwork,
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
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
		},
	})

	// Source: Arbitrum Sepolia (Core+HL)
	routes = append(routes, TestRoute{
		name:        "Arbitrum Sepolia (Core+HL) to Staging Network",
		sourceChain: arbSepolia,
		network:     stagingNetwork,
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
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
					Provider: solvernet.ProviderNone,
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
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
		},
	})

	// Source: Sepolia (HL-only)
	routes = append(routes, TestRoute{
		name:        "Sepolia (HL-only) to Staging Network",
		sourceChain: sepoliaChain,
		network:     stagingNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			// No route to omniStaging (Core-only) from sepoliaChain (HL-only)
			{
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
				},
			},
		},
	})

	// Source: OP Sepolia (Core+HL)
	routes = append(routes, TestRoute{
		name:        "OP Sepolia (Core+HL) to Staging Network",
		sourceChain: opSepolia,
		network:     stagingNetwork,
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
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			{
				ChainID:           sepoliaChain.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
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
		},
	})

	// --- Omega Network Test Cases ---

	// Source: Omni Omega (Core-only)
	routes = append(routes, TestRoute{
		name:        "Omni Omega (Core-only) to Omega Network",
		sourceChain: omniOmega,
		network:     omegaNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID:           omniOmega.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderNone,
				},
			},
			{
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			// No route to sepoliaChain (HL-only) from omniOmega (Core-only)
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
		},
	})

	// Source: Holesky (Core+HL) on Omega Network
	routes = append(routes, TestRoute{
		name:        "Holesky (Core+HL) to Omega Network",
		sourceChain: holesky,
		network:     omegaNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID:           omniOmega.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderNone,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			{
				ChainID:           sepoliaChain.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
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
		},
	})

	// Source: Base Sepolia (Core+HL) on Omega Network
	routes = append(routes, TestRoute{
		name:        "Base Sepolia (Core+HL) to Omega Network",
		sourceChain: baseSepolia,
		network:     omegaNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID:           omniOmega.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
		},
	})

	// Source: Arbitrum Sepolia (Core+HL) on Omega Network
	routes = append(routes, TestRoute{
		name:        "Arbitrum Sepolia (Core+HL) to Omega Network",
		sourceChain: arbSepolia,
		network:     omegaNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID:           omniOmega.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
					Provider: solvernet.ProviderNone,
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
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
		},
	})

	// Source: Sepolia (HL-only) on Omega Network
	routes = append(routes, TestRoute{
		name:        "Sepolia (HL-only) to Omega Network",
		sourceChain: sepoliaChain,
		network:     omegaNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			// No route to omniOmega (Core-only) from sepoliaChain (HL-only)
			{
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			{
				ChainID:           opSepolia.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
				},
			},
		},
	})

	// Source: OP Sepolia (Core+HL) on Omega Network
	routes = append(routes, TestRoute{
		name:        "OP Sepolia (Core+HL) to Omega Network",
		sourceChain: opSepolia,
		network:     omegaNetwork,
		inboxAddr:   dummyInboxAddr,
		outboxAddr:  dummyOutboxAddr,
		expectedRoutes: []Route{
			{
				ChainID:           omniOmega.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           holesky.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderCore,
				},
			},
			{
				ChainID:           baseSepolia.ID,
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
			{
				ChainID:           sepoliaChain.ID,
				OutboxAddrOnInbox: dummyOutboxAddr,
				InboxConfigOnOutbox: bindings.ISolverNetOutboxInboxConfig{
					Inbox:    dummyInboxAddr,
					Provider: solvernet.ProviderHL,
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
