package types

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/rand/v2"
	"net"
	"strings"
	"sync/atomic"
	"time"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// Testnet wraps e2e.Omega with additional omni-specific fields.
type Testnet struct {
	*e2e.Testnet
	Network      netconf.ID
	OmniEVMs     []OmniEVM
	AnvilChains  []AnvilChain
	PublicChains []PublicChain
	OnlyMonitor  bool
	Perturb      map[string][]Perturb
}

// RandomHaloAddr returns a random halo address for cprovider and cometBFT rpc clients.
// It uses the internal IP address of a random node that isn't delayed or a seed.
func (t Testnet) RandomHaloAddr() string {
	var eligible []string
	for _, node := range t.Nodes {
		if node.StartAt != 0 || node.Mode == ModeSeed {
			continue // Skip delayed nodes or seed nodes
		}

		eligible = append(eligible, node.AddressRPC())
	}

	if len(eligible) == 0 {
		return ""
	}

	return eligible[rand.IntN(len(eligible))] //nolint:gosec // Weak random used for deterministic tests.
}

// BroadcastOmniEVM returns a Omni EVM to use for e2e app tx broadcasts.
// It prefers a validator nodes since we have an issue with mempool+p2p+startup where
// txs get stuck in non-validator mempool immediately after startup if not connected to peers yet.
// Also avoid validators that are not started immediately or evms with perturbations.
func (t Testnet) BroadcastOmniEVM() OmniEVM {
	isDelayed := func(evm string) bool {
		for _, node := range t.Nodes {
			if node.StartAt > 0 && strings.Contains(evm, node.Name) {
				return true
			}
		}

		return false
	}

	hasPerturbation := func(evm string) bool {
		for service := range t.Perturb {
			if evm == service {
				return true
			}
		}

		return false
	}

	isValidator := func(evm string) bool {
		return strings.Contains(evm, "validator")
	}

	for _, evm := range t.OmniEVMs {
		instance := evm.InstanceName
		if isValidator(instance) && !isDelayed(instance) && !hasPerturbation(instance) {
			return evm
		}
	}

	return t.OmniEVMs[0]
}

// BroadcastNode returns a halo node to use for RPC queries broadcasts.
// It prefers a validator nodes since we have an issue with mempool+p2p+startup where
// txs get stuck in non-validator mempool immediately after startup if not connected to peers yet.
// Also avoid validators that are not started immediately.
func (t Testnet) BroadcastNode() *e2e.Node {
	for _, node := range t.Nodes {
		if !strings.Contains(node.Name, "validator") {
			continue
		}
		if node.StartAt > 0 {
			continue
		}

		return node
	}

	return t.Nodes[0]
}

// ArchiveNode returns the first node running in ModeArchive.
// Note that this is different from the CometBFT Testnet.ArchiveNodes() method.
func (t Testnet) ArchiveNode() (*e2e.Node, bool) {
	for _, node := range t.Nodes {
		if node.Mode == ModeArchive {
			return node, true
		}
	}

	return nil, false
}

// HasPerturbations returns whether the network has any perturbations.
func (t Testnet) HasPerturbations() bool {
	if len(t.Perturb) > 0 {
		return true
	}

	return t.Testnet.HasPerturbations()
}

func (t Testnet) HasOmniEVM() bool {
	return len(t.OmniEVMs) > 0
}

// EVMChain represents a EVM chain in a omni network.
type EVMChain struct {
	evmchain.Metadata
	Shards   []xchain.ShardID
	IsPublic bool
}

// AttestInterval returns the a constant interval for which attestations are always required, even if empty..
func (c EVMChain) AttestInterval(network netconf.ID) uint64 {
	return intervalFromPeriod(network, c.BlockPeriod)
}

// intervalFromPeriod returns the minimum number of blocks between attestations for a given block period.
func intervalFromPeriod(network netconf.ID, period time.Duration) uint64 {
	target := time.Hour
	if network == netconf.Staging {
		target = time.Minute * 10
	} else if network == netconf.Devnet {
		target = time.Second * 10
	}

	if period == 0 {
		return 0
	}

	return uint64(target / period)
}

func (c EVMChain) ShardsUint64() []uint64 {
	var shards []uint64
	for _, shard := range c.Shards {
		shards = append(shards, uint64(shard))
	}

	return shards
}

// OmniEVM represents a omni evm instance in a omni network. Similar to e2e.Node for halo instances.
type OmniEVM struct {
	Chain        EVMChain // For netconf (all instances must have the same chain)
	InstanceName string   // For docker container name
	AdvertisedIP net.IP   // For setting up NAT on geth bootnode
	ProxyPort    uint32   // For binding
	InternalRPC  string   // For JSON-RPC queries from halo/relayer
	ExternalRPC  string   // For JSON-RPC queries from e2e app.
	IsArchive    bool     // Whether this instance is in archive mode
	JWTSecret    string   // JWT secret for authentication

	// P2P networking
	NodeKey *ecdsa.PrivateKey // Private key
	Enode   *enode.Node       // Public key
	Peers   []*enode.Node     // Peer public keys
}

// NodeKeyHex returns the hex-encoded node key. Used for geth's config.
func (o OmniEVM) NodeKeyHex() string {
	return hex.EncodeToString(crypto.FromECDSA(o.NodeKey))
}

// AnvilChain represents an anvil chain instance in a omni network.
type AnvilChain struct {
	Chain       EVMChain // For netconf
	InternalIP  net.IP   // For docker container IP
	ProxyPort   uint32   // For binding
	InternalRPC string   // For JSON-RPC queries from halo/relayer
	ExternalRPC string   // For JSON-RPC queries from e2e app.
	LoadState   string   // File path to load anvil state from
}

// PublicChain represents a public chain in a omni network.
type PublicChain struct {
	chain        EVMChain      // For netconf
	rpcAddresses []string      // For JSON-RPC queries from halo/relayer/e2e app.
	next         *atomic.Int32 // For round-robin RPC address selection
}

func NewPublicChain(chain EVMChain, rpcAddresses []string) PublicChain {
	return PublicChain{
		chain:        chain,
		rpcAddresses: rpcAddresses,
		next:         new(atomic.Int32),
	}
}

// Chain returns the EVM chain.
func (c PublicChain) Chain() EVMChain {
	return c.chain
}

// NextRPCAddress returns the next RPC address in the list.
func (c PublicChain) NextRPCAddress() string {
	i := c.next.Load()
	defer func() {
		c.next.Store(i + 1)
	}()

	l := len(c.rpcAddresses)
	if l == 0 {
		return ""
	}

	return strings.TrimSpace(c.rpcAddresses[int(i)%l])
}
