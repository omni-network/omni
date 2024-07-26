// Package netconf provides the configuration of an Omni network, an instance
// of the Omni cross chain protocol.
package netconf

import (
	"fmt"
	"sort"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

// Network defines a deployment of the Omni cross chain protocol.
// It spans an omni chain (both execution and consensus) and a set of
// supported rollup EVMs.
type Network struct {
	ID     ID      `json:"name"`   // ID of the network. e.g. "simnet", "testnet", "staging", "mainnet"
	Chains []Chain `json:"chains"` // Chains that are part of the network
}

// Verify returns an error if the configuration is invalid.
func (n Network) Verify() error {
	if err := n.ID.Verify(); err != nil {
		return err
	}

	for _, chain := range n.Chains {
		if err := chain.Verify(); err != nil {
			return errors.Wrap(err, "chain", "id", chain.ID, "name", chain.Name)
		}
	}

	return nil
}

// EVMChains returns all evm chains in the network. It excludes the omni consensus chain.
func (n Network) EVMChains() []Chain {
	resp := make([]Chain, 0, len(n.Chains))
	for _, chain := range n.Chains {
		if IsOmniConsensus(n.ID, chain.ID) {
			continue
		}

		resp = append(resp, chain)
	}

	return resp
}

// ChainIDs returns the all chain IDs in the network.
// This is a convenience method.
func (n Network) ChainIDs() []uint64 {
	resp := make([]uint64, 0, len(n.Chains))
	for _, chain := range n.Chains {
		resp = append(resp, chain.ID)
	}

	return resp
}

// ChainNamesByIDs returns the all chain IDs and names in the network.
// This is a convenience method.
func (n Network) ChainNamesByIDs() map[uint64]string {
	resp := make(map[uint64]string)
	for _, chain := range n.Chains {
		resp[chain.ID] = chain.Name
	}

	return resp
}

// ChainVersionNames returns the all chain version names in the network.
// This is a convenience method.
func (n Network) ChainVersionNames() map[xchain.ChainVersion]string {
	resp := make(map[xchain.ChainVersion]string)
	for _, chain := range n.Chains {
		for _, chainVersion := range chain.ChainVersions() {
			resp[chainVersion] = n.ChainVersionName(chainVersion)
		}
	}

	return resp
}

// OmniEVMChain returns the Omni execution chain config or false if it does not exist.
func (n Network) OmniEVMChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if IsOmniExecution(n.ID, chain.ID) {
			return chain, true
		}
	}

	return Chain{}, false
}

// OmniConsensusChain returns the Omni consensus chain config or false if it does not exist.
func (n Network) OmniConsensusChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if IsOmniConsensus(n.ID, chain.ID) {
			return chain, true
		}
	}

	return Chain{}, false
}

// EthereumChain returns the ethereum Layer1 chain config or false if it does not exist.
func (n Network) EthereumChain() (Chain, bool) {
	for _, chain := range n.Chains {
		if IsEthereumChain(n.ID, chain.ID) {
			return chain, true
		}
	}

	return Chain{}, false
}

// IsEthereumChain returns true if the chainID is the EthereumChainID for the given network.
func IsEthereumChain(network ID, chainID uint64) bool {
	switch network {
	case Mainnet:
		return chainID == evmchain.IDEthereum
	case Omega:
		return chainID == evmchain.IDHolesky
	default:
		return chainID == evmchain.IDMockL1Fast || chainID == evmchain.IDMockL1Slow
	}
}

// ChainName returns the chain name for the given ID or an empty string if it does not exist.
func (n Network) ChainName(id uint64) string {
	chain, _ := n.Chain(id)
	return chain.Name
}

// ChainVersionName returns the chain version name for the given ID or an empty string if it does not exist.
func (n Network) ChainVersionName(chainVer xchain.ChainVersion) string {
	chain, _ := n.Chain(chainVer.ID)

	return fmt.Sprintf("%s|%s", chain.Name, chainVer.ConfLevel.Label())
}

// StreamName returns the stream name for the given stream ID.
func (n Network) StreamName(stream xchain.StreamID) string {
	srcChain, _ := n.Chain(stream.SourceChainID)
	destChain, _ := n.Chain(stream.DestChainID)

	return fmt.Sprintf("%s|%s|%s", srcChain.Name, stream.ShardID.Label(), destChain.Name)
}

// Chain returns the chain config for the given ID or false if it does not exist.
func (n Network) Chain(id uint64) (Chain, bool) {
	for _, chain := range n.Chains {
		if chain.ID == id {
			return chain, true
		}
	}

	return Chain{}, false
}

// ChainVersionsTo returns the all chain versions submitted to the provided destination chain.
func (n Network) ChainVersionsTo(dstChainID uint64) []xchain.ChainVersion {
	var resp []xchain.ChainVersion
	for _, chain := range n.Chains {
		if chain.ID == dstChainID {
			continue // Skip self
		}

		resp = append(resp, chain.ChainVersions()...)
	}

	return resp
}

// StreamsTo returns the all streams to the provided destination chain.
func (n Network) StreamsTo(dstChainID uint64) []xchain.StreamID {
	if dstChainID == n.ID.Static().OmniConsensusChainIDUint64() {
		return nil // Consensus chain is never a destination chain.
	}

	var resp []xchain.StreamID
	for _, srcChain := range n.Chains {
		if srcChain.ID == dstChainID {
			continue // Skip self
		}

		for _, shardID := range srcChain.Shards {
			resp = append(resp, xchain.StreamID{
				SourceChainID: srcChain.ID,
				DestChainID:   dstChainID,
				ShardID:       shardID,
			})
		}
	}

	return resp
}

// StreamsFrom returns the all streams from the provided source chain.
func (n Network) StreamsFrom(srcChainID uint64) []xchain.StreamID {
	srcChain, ok := n.Chain(srcChainID)
	if !ok {
		return nil
	}

	var resp []xchain.StreamID
	for _, dstChain := range n.EVMChains() { // Only EVM chain can be destinations.
		if srcChainID == dstChain.ID {
			continue // Skip self
		}

		for _, shardID := range srcChain.Shards {
			resp = append(resp, xchain.StreamID{
				SourceChainID: srcChain.ID,
				DestChainID:   dstChain.ID,
				ShardID:       shardID,
			})
		}
	}

	return resp
}

// StreamsBetween returns the all streams between the provided source and destination chain.
func (n Network) StreamsBetween(srcChainID uint64, dstChainID uint64) []xchain.StreamID {
	if srcChainID == dstChainID {
		return nil
	} else if dstChainID == n.ID.Static().OmniConsensusChainIDUint64() {
		return nil // Consensus chain is never a destination chain.
	}

	srcChain, ok := n.Chain(srcChainID)
	if !ok {
		return nil
	}

	var resp []xchain.StreamID

	for _, shardID := range srcChain.Shards {
		resp = append(resp, xchain.StreamID{
			SourceChainID: srcChain.ID,
			DestChainID:   dstChainID,
			ShardID:       shardID,
		})
	}

	return resp
}

// Chain defines the configuration of an execution chain that supports
// the Omni cross chain protocol. This is most supported Rollup EVMs, but
// also the Omni EVM, and the Omni Consensus chain.
type Chain struct {
	ID             uint64           // Chain ID asa per https://chainlist.org
	Name           string           // Chain name as per https://chainlist.org
	PortalAddress  common.Address   // Address of the omni portal contract on the chain
	DeployHeight   uint64           // Height that the portal contracts were deployed
	BlockPeriod    time.Duration    // Block period of the chain
	Shards         []xchain.ShardID // Supported xmsg shards
	AttestInterval uint64           // Attest to every Nth block, even if empty.
}

// ConfLevels returns the uniq set of confirmation levels
// supported by the chain. This is inferred from the supported shards.
func (c Chain) ConfLevels() []xchain.ConfLevel {
	dedup := make(map[xchain.ConfLevel]struct{})

	for _, shard := range c.Shards {
		conf := shard.ConfLevel()
		if _, ok := dedup[conf]; ok {
			continue
		}
		dedup[conf] = struct{}{}
	}

	confs := make([]xchain.ConfLevel, 0, len(dedup))
	for conf := range dedup {
		confs = append(confs, conf)
	}

	// Sort for deterministic ordering.
	sort.Slice(confs, func(i, j int) bool {
		return confs[i] < confs[j]
	})

	return confs
}

// ChainVersions returns the uniq set of chain versions
// supported by the chain. This is inferred from the supported shards.
func (c Chain) ChainVersions() []xchain.ChainVersion {
	var resp []xchain.ChainVersion
	for _, conf := range c.ConfLevels() {
		resp = append(resp, xchain.ChainVersion{
			ID:        c.ID,
			ConfLevel: conf,
		})
	}

	return resp
}

func (c Chain) ShardsUint64() []uint64 {
	var resp []uint64
	for _, shard := range c.Shards {
		resp = append(resp, uint64(shard))
	}

	return resp
}

func (c Chain) Verify() error {
	if c.ID == 0 {
		return errors.New("zero chain ID")
	}

	if c.Name == "" {
		return errors.New("empty chain name")
	}

	if len(c.Shards) == 0 {
		return errors.New("empty shards")
	}

	return nil
}
