package types

import (
	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// Mode defines the halo consensus node mode.
// Nodes are in general full nodes (light nodes are not supported yet).
// In some cases, additional roles are defined: validator, archive, seed.
//
// Note that the execution clients only have two modes: "default" and "archive".
//
// e2e.Mode is extended so ModeArchive can be added transparently.
type Mode = e2e.Mode

const (
	// ModeValidator defines a validator node.
	// It's validator key has staked tokens and it actively participates in consensus and is subject to rewards and penalties.
	// It must always be online, otherwise it will get stashed/jailed.
	// [genesis_validator_set=true,pruning=default,consensus=default,special_p2p=false].
	// Note technically a validator node is also a "full node".
	ModeValidator = e2e.ModeValidator

	// ModeArchive defines an archive node.
	// It stores all historical blocks and state, it doesn't delete anything ever. It will require TBs of disk.
	// [genesis_validator_set=false,pruning=none,consensus=default,special_p2p=false].
	// Note technically an archive node is also a "full node".
	ModeArchive Mode = "archive"

	// ModeSeed defines a seed node. It must have a long-lived p2p pubkey and address (encoded in repo).
	// It acts as notice board for external nodes to learn about the network and connect to publicly available nodes.
	// It crawls the network regularly, making it available to new nodes.
	// [genesis_validator_set=false,pruning=default,consensus=default,special_p2p=true].
	// Note technically a seed node is also a "full node".
	ModeSeed = e2e.ModeSeed

	// ModeFull defines a full node. A full node a normal node without a special role.
	// [genesis_validator_set=false,pruning=default,consensus=default,special_p2p=false].
	ModeFull = e2e.ModeFull

	// ModeLight defines a light node. This isn't used yet.
	// [genesis_validator_set=false,pruning=no_data,consensus=light,special_p2p=false]
	// Only light nodes are not also full nodes.
	ModeLight = e2e.ModeLight
)

// Manifest wraps e2e.Manifest with additional omni-specific fields.
type Manifest struct {
	e2e.Manifest

	Network netconf.ID `toml:"network"`

	// AnvilChains defines the anvil chains to deploy; chain_a, chain_b, etc.
	AnvilChains []string `toml:"anvil_chains"`

	// PublicChains defines the public chains to connect to; arb_sepolia, etc.
	PublicChains []string `toml:"public_chains"`

	// AVSTarget identifies the chain to deploy the AVS contracts to.
	// It must be one of the anvil or public chains.
	AVSTarget string `toml:"avs_target"`

	// MultiOmniEVMs defines whether to deploy one or multiple Omni EVMs.
	MultiOmniEVMs bool `toml:"multi_omni_evms"`

	// OnlyMonitor indicates that the monitor is the only thing that we deploy in this network.
	OnlyMonitor bool `toml:"only_monitor"`

	// PingPongN defines the number of ping pong messages to send. Defaults 3 if 0.
	PingPongN uint64 `toml:"pingpong_n"`

	// Keys contains long-lived private keys (address by type) by node name.
	Keys map[string]map[key.Type]string `toml:"keys"`

	// Explorer defines whether to deploy the explorer.
	Explorer bool `toml:"explorer"`
}

// Seeds returns a map of seed nodes by name.
func (m Manifest) Seeds() map[string]bool {
	resp := make(map[string]bool)
	for name, node := range m.Nodes {
		if Mode(node.Mode) == ModeSeed {
			resp[name] = true
		}
	}

	return resp
}

// OmniEVMs returns a map of omni evm instances names by <IsArchive> to deploy.
// If only a single Omni EVM is to be deployed, the name is "omni_evm".
// Otherwise, the names are "<node>_evm".
func (m Manifest) OmniEVMs() map[string]bool {
	if !m.MultiOmniEVMs {
		return map[string]bool{
			"omni_evm": false,
		}
	}

	resp := make(map[string]bool)
	for name, node := range m.Nodes {
		resp[name+"_evm"] = Mode(node.Mode) == ModeArchive
	}

	return resp
}
