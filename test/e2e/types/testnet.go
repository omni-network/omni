package types

import (
	"crypto/ecdsa"
	"net"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/ethereum/go-ethereum/p2p/enode"
)

// Testnet wraps e2e.Testnet with additional omni-specific fields.
type Testnet struct {
	*e2e.Testnet
	Network      string
	OmniEVMs     []OmniEVM
	AnvilChains  []AnvilChain
	PublicChains []PublicChain
}

// EVMChain represents a EVM chain in a omni network.
type EVMChain struct {
	Name     string // Chain Nam.
	ID       uint64 // Chain ID
	IsPublic bool
}

// OmniEVM represetns a omni evm instance in a omni network. Similar to e2e.Node for halo instances.
// TODO(corver): Extend OmniEVM with bootnode peer addresses for p2p networking.
type OmniEVM struct {
	Chain           EVMChain // For netconf (all instances must have the same chain)
	InstanceName    string   // For docker container name
	InternalIP      net.IP   // For docker container IP
	ProxyPort       uint32   // For binding
	InternalRPC     string   // For JSON-RPC queries from halo/relayer
	InternalAuthRPC string   // For engine API queries from halo
	ExternalRPC     string   // For JSON-RPC queries from e2e app.

	// P2P networking
	NodeKey   *ecdsa.PrivateKey // Private key
	Enode     *enode.Node       // Public key
	BootNodes []*enode.Node     // Peer public keys
}

// AnvilChain represents an anvil chain instance in a omni network.
type AnvilChain struct {
	Chain       EVMChain // For netconf
	InternalIP  net.IP   // For docker container IP
	ProxyPort   uint32   // For binding
	InternalRPC string   // For JSON-RPC queries from halo/relayer
	ExternalRPC string   // For JSON-RPC queries from e2e app.
}

// PublicChain represents a public chain in a omni network.
type PublicChain struct {
	Chain      EVMChain // For netconf
	RPCAddress string   // For JSON-RPC queries from halo/relayer/e2e app.
}
