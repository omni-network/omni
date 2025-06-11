package solvernet

import (
	"context"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gagliardetto/solana-go/rpc"
)

// hlChains define solvernet chains secured by hyperlane.
var hlChains = map[netconf.ID][]uint64{
	// Mainnet
	netconf.Mainnet: {
		// evmchain.IDBSC,
		// evmchain.IDPolygon,
		evmchain.IDHyperEVM,
		evmchain.IDMantle,
		// evmchain.IDBerachain,
		// evmchain.IDPlume,
	},

	// Testnet
	netconf.Omega: {
		// evmchain.IDBSCTestnet,
		// evmchain.IDHyperEVMTestnet,
		// evmchain.IDPolygonAmoy,
		// evmchain.IDPlumeTestnet,
		evmchain.IDSepolia,
	},

	// Staging
	netconf.Staging: {
		evmchain.IDSepolia,
	},

	// Devnet
	netconf.Devnet: {
		evmchain.IDSepolia,
	},
}

var trustedChains = map[netconf.ID][]uint64{
	netconf.Devnet: {
		evmchain.IDSolanaLocal,
	},
	netconf.Staging: {
		// evmchain.IDSolanaDevnet,
	},
	netconf.Omega: {
		// evmchain.IDSolanaDevnet,
	},
	netconf.Mainnet: {
		// evmchain.IDSolana,
	},
}

func allChains(network netconf.ID) []uint64 {
	var resp []uint64
	resp = append(resp, hlChains[network]...)
	resp = append(resp, trustedChains[network]...)

	return resp
}

// Chains returns the list of solver-specific chains for a given solvernet network.
func Chains(network netconf.ID) []netconf.Chain {
	var resp []netconf.Chain
	for _, chainID := range allChains(network) {
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

// isHLRole returns true if the role is a hyperlane-related role.
func isHLRole(role eoa.Role) bool {
	if role != eoa.RoleRelayer && role != eoa.RoleMonitor && role != eoa.RoleTester && role != eoa.RoleXCaller {
		return true
	}

	return false
}

func isTrustedRole(role eoa.Role) bool {
	return map[eoa.Role]bool{
		eoa.RoleCold:     true,
		eoa.RoleHot:      true,
		eoa.RoleDeployer: true,
		eoa.RoleSolver:   true,
		eoa.RoleFlowgen:  true,
	}[role]
}

// IsSolverOnly returns true if the chain is a solver-specific chain (not part of omni core).
func IsSolverOnly(chainID uint64) bool {
	return IsHLOnly(chainID) || IsTrusted(chainID)
}

// IsDisabled returns true if the chain is disabled.
// This configures all routes for this chain to be disabled locally and remotely.
func IsDisabled(_ uint64) bool {
	return false // id == evmchain.IDSepolia
}

// FilterByEndpoints returns an HL chain selector that excludes chains without endpoints.
func FilterByEndpoints(endpoints xchain.RPCEndpoints) func(netconf.ID, netconf.Chain) bool {
	return func(_ netconf.ID, chain netconf.Chain) bool {
		_, err := endpoints.ByNameOrID(chain.Name, chain.ID)
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

		if evmchain.IsSVM(chain.ID) {
			_, ok, _ := anchorinbox.GetInboxState(ctx, rpc.New(endpoint))
			return ok
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

// AddNetwork returns a copy of the network with solver-specific chains added.
// Optional selector functions can be provided to filter the chains.
func AddNetwork(ctx context.Context, network netconf.Network, selectors ...func(netconf.ID, netconf.Chain) bool) netconf.Network {
	chains := Chains(network.ID)

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

	log.Debug(ctx, "Adding solver-specific chains to network", "network", network.ID, "included", included, "excluded", excluded)

	return network.AddChains(chains...)
}
