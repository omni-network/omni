package solvernet

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
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
	// Devnet
	netconf.Devnet: {
		evmchain.IDSolanaLocal,
	},
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

func TrustedChains(network netconf.ID) []netconf.Chain {
	var resp []netconf.Chain
	for _, chainID := range trustedChains[network] {
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

// SkipRole returns true if the role should be skipped for the given chain ID.
func SkipRole(chainID uint64, role eoa.Role) bool {
	if IsHLOnly(chainID) {
		return !IsHLRole(role)
	}

	return false // Otherwise don't skip any roles
}

// IsHLRole returns true if the role is a hyperlane-related role.
func IsHLRole(role eoa.Role) bool {
	if role != eoa.RoleRelayer && role != eoa.RoleMonitor && role != eoa.RoleTester && role != eoa.RoleXCaller {
		return true
	}

	return false
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
			log.Debug(ctx, "Endpoint", "error", err.Error())
			return false
		}

		ethCl, err := ethclient.DialContext(ctx, chain.Name, endpoint)
		if err != nil {
			log.Debug(ctx, "Dial", "error", err.Error())
			return false
		}

		addrs, err := contracts.GetAddresses(ctx, network)
		if err != nil {
			log.Debug(ctx, "Addresses", "error", err.Error())
			return false
		}

		contract, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, ethCl)
		if err != nil {
			log.Debug(ctx, "Inbox contract", "error", err.Error())
			return false
		}

		_, err = contract.DeployedAt(&bind.CallOpts{Context: ctx})
		if err != nil {
			log.Debug(ctx, "Deployed at", "error", err.Error())
		}

		return err == nil
	}
}

// AddHLNetwork returns a copy of the network with hyperlane-secured chains added.
// Optional selector functions can be provided to filter the chains.
func AddHLNetwork(ctx context.Context, network netconf.Network, selectors ...func(netconf.ID, netconf.Chain) bool) netconf.Network {
	resp := filter(network.ID, HLChains(network.ID), selectors...)

	log.Debug(ctx, "Adding hyperlane chains to network", "network", network.ID, "included", resp.FmtIncluded(), "excluded", resp.FmtExcluded())

	return network.AddChains(resp.Included...)
}

// AddTrustedNetwork returns a copy of the network with solver-trusted chains added.
// Optional selector functions can be provided to filter the chains.
func AddTrustedNetwork(ctx context.Context, network netconf.Network, selectors ...func(netconf.ID, netconf.Chain) bool) netconf.Network {
	resp := filter(network.ID, TrustedChains(network.ID), selectors...)

	log.Debug(ctx, "Adding solver-trusted chains to network", "network", network.ID, "included", resp.FmtIncluded(), "excluded", resp.FmtExcluded())

	return network.AddChains(resp.Included...)
}

type filterResp struct {
	Included []netconf.Chain
	Excluded []netconf.Chain
}

func (f filterResp) FmtIncluded() []string {
	var resp []string
	for _, chain := range f.Included {
		resp = append(resp, chain.Name)
	}

	return resp
}

func (f filterResp) FmtExcluded() []string {
	var resp []string
	for _, chain := range f.Excluded {
		resp = append(resp, chain.Name)
	}

	return resp
}

// AddTrustedNetwork returns a copy of the network with solver-trusted chains added.
// Optional selector functions can be provided to filter the chains.
func filter(network netconf.ID, chains []netconf.Chain, selectors ...func(netconf.ID, netconf.Chain) bool) filterResp {
	// Filter out chains using provided selectors
	var excluded []netconf.Chain
	for _, selector := range selectors {
		var selected []netconf.Chain
		for _, chain := range chains {
			if selector(network, chain) {
				selected = append(selected, chain)
			} else {
				excluded = append(excluded, chain)
			}
		}
		chains = selected
	}

	return filterResp{
		Included: chains,
		Excluded: excluded,
	}
}
