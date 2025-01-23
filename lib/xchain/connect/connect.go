//nolint:revive // unexported option type is fine.
package connect

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/cchain"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"
)

// Connector provides a simple abstraction to connect to the Omni network.
type Connector struct {
	Network    netconf.Network
	XProvider  xchain.Provider
	CProvider  cchain.Provider
	EthClients map[uint64]ethclient.Client
	Backends   ethbackend.Backends
	CmtCl      rpcclient.Client
}

// Backend returns an ethbackend for the given chainID.
func (c Connector) Backend(chainID uint64) (*ethbackend.Backend, error) {
	return c.Backends.Backend(chainID)
}

// Portal returns an OmniPortal contract.
func (c Connector) Portal(ctx context.Context, chainID uint64) (*bindings.OmniPortal, error) {
	backend, err := c.Backends.Backend(chainID)
	if err != nil {
		return nil, err
	}

	addrs, err := contracts.GetAddresses(ctx, c.Network.ID)
	if err != nil {
		return nil, err
	}

	contract, err := bindings.NewOmniPortal(addrs.Portal, backend)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

type options struct {
	Endpoints xchain.RPCEndpoints
	PrivKeys  []*ecdsa.PrivateKey
}

type option func(*options) error

// WithPublicRPCs returns an option using well known public free RPCs for all xchains.
// This is used be default if no other option is provided.
func WithPublicRPCs() option {
	return func(o *options) error {
		for name, rpc := range o.Endpoints {
			if rpc != "" {
				continue
			}
			o.Endpoints[name] = types.PublicRPCByName(name)
		}

		return nil
	}
}

// WithEndpoint returns an option to set the RPC endpoint for the given chain name or chain ID.
func WithEndpoint(chainName string, rpc string) option {
	return func(o *options) error {
		o.Endpoints[chainName] = rpc

		return nil
	}
}

// WithPrivKey returns an option to add the privkey to underlying backends.
func WithPrivKey(privkeys *ecdsa.PrivateKey) option {
	return func(o *options) error {
		o.PrivKeys = append(o.PrivKeys, privkeys)

		return nil
	}
}

// WithInfuraENV returns an option using the provided ENV VAR as infura API key for all xchains.
func WithInfuraENV(keyVar string) option {
	return func(o *options) error {
		infuraNames := map[string]string{
			"ethreum":      "mainnet",
			"holesky":      "holesky",
			"base":         "base-mainnet",
			"base_sepolia": "base-sepolia",
			"arbitrum_one": "arbitrum-mainnet",
			"arb_sepolia":  "arbitrum-sepolia",
			"optimism":     "optimism-mainnet",
			"op_sepolia":   "optimism-sepolia",
		}

		key, ok := os.LookupEnv(keyVar)
		if !ok {
			return errors.New("infura key not found in env", "key", keyVar)
		}

		for name, rpc := range o.Endpoints {
			if infuraNames[name] != "" && rpc == "" {
				o.Endpoints[name] = fmt.Sprintf("https://%s.infura.io/v3/%s", infuraNames[name], key)
			}
		}

		return nil
	}
}

// New returns a populated Connector for the given network.
// It connects to well-known free public RPCs. Use WithInfuraENV or WithEndpoint to override this.
func New(ctx context.Context, netID netconf.ID, opts ...option) (Connector, error) {
	if len(opts) == 0 {
		opts = append(opts, WithPublicRPCs())
	}

	endpoints := make(xchain.RPCEndpoints)

	// Add default omni consensus and execution RPC endpoints.
	omniCons := netID.Static().OmniConsensusChain()
	omniExec, ok := evmchain.MetadataByID(netID.Static().OmniExecutionChainID)
	if !ok {
		return Connector{}, errors.New("omni evm metadata not found")
	}
	endpoints[omniExec.Name] = netID.Static().ExecutionRPC()
	endpoints[omniCons.Name] = netID.Static().ConsensusRPC()

	// Apply any custom endpoint options.
	o := options{Endpoints: endpoints}
	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return Connector{}, err
		}
	}

	portalReg, err := makePortalRegistry(netID, endpoints)
	if err != nil {
		return Connector{}, err
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, netID, portalReg, nil)
	if err != nil {
		return Connector{}, err
	}

	// Add zero endpoints for all detected chains
	for _, chain := range network.Chains {
		if _, ok := endpoints[chain.Name]; ok {
			continue
		}
		endpoints[chain.Name] = ""
	}

	// Apply option again, since we now know the network.
	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return Connector{}, err
		}
	}

	ethClients := make(map[uint64]ethclient.Client)
	for _, chain := range network.Chains {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			rpc = "unknown"
		} else {
			ethCl, err := ethclient.Dial(chain.Name, rpc)
			if err != nil {
				return Connector{}, errors.Wrap(err, "dial eth client")
			}
			ethClients[chain.ID] = ethCl
		}

		log.Info(ctx, "Network chain", "chain", chain.Name, "id", chain.ID, "rpc", libcmd.Redact("", rpc))
	}

	// Connect to the halo cometBFT RPC server.
	cometCl, err := rpchttp.New(endpoints[omniCons.Name], "/websocket")
	if err != nil {
		return Connector{}, errors.Wrap(err, "comet rpc client")
	}

	cprov := cprovider.NewABCI(cometCl, netID)

	xprov := xprovider.New(network, ethClients, cprov)

	backends, err := ethbackend.BackendsFromNetwork(network, o.Endpoints, o.PrivKeys...)
	if err != nil {
		return Connector{}, errors.Wrap(err, "eth backends")
	}

	return Connector{
		Network:    network,
		XProvider:  xprov,
		CProvider:  cprov,
		EthClients: ethClients,
		Backends:   backends,
		CmtCl:      cometCl,
	}, nil
}

func makePortalRegistry(network netconf.ID, endpoints xchain.RPCEndpoints) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	rpc, err := endpoints.ByNameOrID(meta.Name, meta.ChainID)
	if err != nil {
		return nil, err
	}

	ethCl, err := ethclient.Dial(meta.Name, rpc)
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
