package types

import (
	"crypto/ecdsa"
	"encoding/hex"
	"net"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// Testnet wraps e2e.Testnet with additional omni-specific fields.
type Testnet struct {
	*e2e.Testnet
	Network        netconf.ID
	OmniEVMs       []OmniEVM
	AnvilChains    []AnvilChain
	PublicChains   []PublicChain
	OnlyMonitor    bool
	Explorer       bool
	ExplorerMockDB bool
	ExplorerDBConn string
}

func (t Testnet) AVSChain() (EVMChain, error) {
	for _, c := range t.AnvilChains {
		if c.Chain.IsAVSTarget {
			return c.Chain, nil
		}
	}

	for _, c := range t.PublicChains {
		if c.Chain.IsAVSTarget {
			return c.Chain, nil
		}
	}

	return EVMChain{}, errors.New("avs target chain found")
}

// BroadcastOmniEVM returns a Omni EVM to use for e2e app tx broadcasts.
// It prefers a validator nodes since we have an issue with mempool+p2p+startup where
// txs get stuck in non-validator mempool immediately after startup if not connected to peers yet.
func (t Testnet) BroadcastOmniEVM() OmniEVM {
	for _, evm := range t.OmniEVMs {
		if strings.Contains(evm.InstanceName, "validator") {
			return evm
		}
	}

	return t.OmniEVMs[0]
}

func (t Testnet) HasOmniEVM() bool {
	return len(t.OmniEVMs) > 0
}

// EVMChain represents a EVM chain in a omni network.
type EVMChain struct {
	Name               string // Chain Nam.
	ID                 uint64 // Chain ID
	IsPublic           bool
	IsAVSTarget        bool
	BlockPeriod        time.Duration
	FinalizationStrat  netconf.FinalizationStrat
	AVSContractAddress common.Address
}

// OmniEVM represents a omni evm instance in a omni network. Similar to e2e.Node for halo instances.
type OmniEVM struct {
	Chain           EVMChain // For netconf (all instances must have the same chain)
	InstanceName    string   // For docker container name
	AdvertisedIP    net.IP   // For setting up NAT on geth bootnode
	ProxyPort       uint32   // For binding
	InternalRPC     string   // For JSON-RPC queries from halo/relayer
	InternalAuthRPC string   // For engine API queries from halo
	ExternalRPC     string   // For JSON-RPC queries from e2e app.
	IsArchive       bool     // Whether this instance is in archive mode
	JWTSecret       string   // JWT secret for authentication

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
	Chain      EVMChain // For netconf
	RPCAddress string   // For JSON-RPC queries from halo/relayer/e2e app.
}
