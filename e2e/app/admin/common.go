package admin

import (
	"context"
	"os/exec"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	fbproxy "github.com/omni-network/omni/e2e/fbproxy/app"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// adminABI is the ABI for the Admin script contract.
var adminABI = mustGetABI(bindings.AdminMetaData)

// shared contains common resources for all admin operations.
type shared struct {
	admin       common.Address
	endpoints   xchain.RPCEndpoints
	network     netconf.Network
	fireAPIKey  string
	fireKeyPath string
}

// setup returns common resources for all admin operations.
func setup(ctx context.Context, def app.Definition) (shared, error) {
	netID := def.Testnet.Network

	admin, ok := eoa.Address(netID, eoa.RoleAdmin)
	if !ok {
		return shared{}, errors.New("admin eoas not found", "network", netID)
	}

	endpoints := makeEndpoints(def)

	portalReg, err := makePortalRegistry(netID, endpoints)
	if err != nil {
		return shared{}, errors.Wrap(err, "portal registry")
	}

	network, err := netconf.AwaitOnChain(ctx, netID, portalReg, nil)
	if err != nil {
		return shared{}, errors.Wrap(err, "await on chain")
	}

	return shared{
		admin:       admin,
		endpoints:   endpoints,
		network:     network,
		fireAPIKey:  def.Cfg.FireAPIKey,
		fireKeyPath: def.Cfg.FireKeyPath,
	}, nil
}

// chain contains chain specific resources, exteding netconf.Chain with an rpc endpoint.
type chain struct {
	netconf.Chain
	rpc string
}

func setupChain(ctx context.Context, s shared, name string) (chain, error) {
	meta, ok := netconf.MetadataByName(s.network.ID, name)
	if !ok {
		return chain{}, errors.New("chain not found", "chain", name)
	}

	c, ok := s.network.Chain(meta.ChainID)
	if !ok {
		return chain{}, errors.New("chain not found", "chain", name, "id", meta.ChainID)
	}

	rpc, err := s.endpoints.ByNameOrID(c.Name, c.ID)
	if err != nil {
		return chain{}, errors.Wrap(err, "rpc endpoint")
	}

	// only use fb proxy for non-devnet - devnet uses anvil accounts
	if s.network.ID == netconf.Devnet {
		rpc, err = startFBProxy(ctx, s.network.ID, rpc, s.fireAPIKey, s.fireKeyPath)
		if err != nil {
			return chain{}, errors.Wrap(err, "start fb proxy")
		}
	}

	return chain{Chain: c, rpc: rpc}, nil
}

func runForge(ctx context.Context, script string, rpc string, input []byte, sender common.Address,
) (string, error) {
	// assumes running from root
	dir := "./contracts/core"

	// for dev anvil accounts, we signed with private key directly
	if pk, ok := anvil.PrivateKey(sender); ok {
		return execCmd(ctx, dir, "forge", "script", script,
			"--broadcast", "--slow", "--rpc-url", rpc, "--sig", hexutil.Encode(input),
			"--private-key", hexutil.EncodeBig(pk.D),
		)
	}

	return execCmd(ctx, dir, "forge", "script", script,
		"--broadcast", "--slow", "--rpc-url", rpc, "--sig", hexutil.Encode(input),
		"--unlocked", "--sender", sender.Hex(),
	)
}

// startFBProxy starts a fireblocks proxy to rpc, returns the listen address. The proxy stops when ctx is done.
func startFBProxy(ctx context.Context, netID netconf.ID, baseRPC string, fireAPIKey string, fireKeyPath string,
) (string, error) {
	cfg := fbproxy.Config{
		ListenAddr:  "0.0.0.0:0", // Random available port.
		BaseRPC:     baseRPC,
		Network:     netID,
		FireAPIKey:  fireAPIKey,
		FireKeyPath: fireKeyPath,
	}

	listenAddr, err := fbproxy.Start(ctx, cfg)
	if err != nil {
		return "", errors.Wrap(err, "start fb proxy")
	}

	return "http://" + listenAddr, nil
}

func execCmd(ctx context.Context, dir string, cmd string, args ...string) (string, error) {
	c := exec.CommandContext(ctx, cmd, args...)
	c.Dir = dir

	out, err := c.CombinedOutput()
	if err != nil {
		return string(out), errors.Wrap(err, "exec", "out", string(out))
	}

	return string(out), nil
}

func makeEndpoints(def app.Definition) xchain.RPCEndpoints {
	endpoints := make(xchain.RPCEndpoints)

	// Add all public chains
	for _, public := range def.Testnet.PublicChains {
		endpoints[public.Chain().Name] = public.NextRPCAddress()
	}

	// Connect to a proper omni_evm that isn't unavailable
	omniEVM := def.Testnet.BroadcastOmniEVM()
	endpoints[omniEVM.Chain.Name] = omniEVM.ExternalRPC

	// Add omni consensus chain
	endpoints[def.Testnet.Network.Static().OmniConsensusChain().Name] = def.Testnet.BroadcastNode().AddressRPC()

	// Add all anvil chains
	for _, anvil := range def.Testnet.AnvilChains {
		endpoints[anvil.Chain.Name] = anvil.ExternalRPC
	}

	return endpoints
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

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}
