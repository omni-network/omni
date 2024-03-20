package types

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
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

// GcMode represents the garbage collection mode for geth. Either "full" or "archive".
type GcMode string

const (
	GcModeFull    GcMode = "full"    // default mode for geth
	GcModeArchive GcMode = "archive" // archival mode for geth
)

// Testnet wraps e2e.Testnet with additional omni-specific fields.
type Testnet struct {
	*e2e.Testnet
	Network      string
	OmniEVMs     []OmniEVM
	AnvilChains  []AnvilChain
	PublicChains []PublicChain
	OnlyMonitor  bool
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
	InternalIP      net.IP   // For docker container IP
	ExternalIP      net.IP   // For setting up NAT on geth bootnode
	ProxyPort       uint32   // For binding
	InternalRPC     string   // For JSON-RPC queries from halo/relayer
	InternalAuthRPC string   // For engine API queries from halo
	ExternalRPC     string   // For JSON-RPC queries from e2e app.
	GcMode          GcMode   // Geth config for archive or full mode

	// P2P networking
	NodeKey   *ecdsa.PrivateKey // Private key
	Enode     *enode.Node       // Public key
	BootNodes []*enode.Node     // Peer public keys
}

// NodeKeyHex returns the hex-encoded node key. Used for geth's config.
func (o OmniEVM) NodeKeyHex() string {
	return hex.EncodeToString(crypto.FromECDSA(o.NodeKey))
}

// BootNodesStrArr returns a string array of bootnodes for use in geth's config for bootnodes.
func (o OmniEVM) BootNodesStrArr() string {
	var resp []string
	for _, b := range o.BootNodes {
		resp = append(resp, fmt.Sprintf(`"%s"`, b.String()))
	}

	return strings.Join(resp, ",")
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

// GethConfig represents part of the geth configuration that can't be initialized through the command line args.
type GethConfig struct {
	// BootstrapNodes are used to establish connectivity
	// with the rest of the network.
	BootstrapNodes string

	// Static nodes are used as pre-configured connections which are always
	// maintained and re-connected on disconnects.
	StaticNodes string
}
