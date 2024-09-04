package app

import (
	"context"
	"os/exec"
	"strings"

	"github.com/omni-network/omni/e2e/app/eoa"
	fbproxy "github.com/omni-network/omni/e2e/fbproxy/app"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	scriptPortalAdmin  = "PortalAdmin"
	scriptStakingAdmin = "StakingAdmin"

	chainAll     = "all"
	chainOmniEVM = "omni_evm"
)

// shared contains common resources for all admin operations.
type shared struct {
	admin       common.Address
	deployer    common.Address
	endpoints   xchain.RPCEndpoints
	network     netconf.Network
	fireAPIKey  string
	fireKeyPath string
}

// chain contains chain specific resources, exteding netconf.Chain with an rpc endpoint.
type chain struct {
	netconf.Chain
	rpc string
}

// setup returns common resources for all admin operations.
func setup(def Definition) (shared, error) {
	netID := def.Testnet.Network

	admin, ok := eoa.Address(netID, eoa.RoleAdmin)
	if !ok {
		return shared{}, errors.New("admin eoas not found", "network", netID)
	}

	deployer, ok := eoa.Address(netID, eoa.RoleDeployer)
	if !ok {
		return shared{}, errors.New("deployer eoas not found", "network", netID)
	}

	endpoints := externalEndpoints(def)
	network := networkFromDef(def)

	return shared{
		admin:       admin,
		deployer:    deployer,
		endpoints:   endpoints,
		network:     network,
		fireAPIKey:  def.Cfg.FireAPIKey,
		fireKeyPath: def.Cfg.FireKeyPath,
	}, nil
}

// setupChain returns chain specific resources.
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
	if s.network.ID != netconf.Devnet {
		rpc, err = startFBProxy(ctx, s.network.ID, rpc, s.fireAPIKey, s.fireKeyPath)
		if err != nil {
			return chain{}, errors.Wrap(err, "start fb proxy")
		}
	}

	return chain{Chain: c, rpc: rpc}, nil
}

// forChains runs a function for each chain in names.
func forChains(
	ctx context.Context,
	names []string,
	s shared,
	fn func(context.Context, shared, chain) error,
) error {
	for _, name := range names {
		c, err := setupChain(ctx, s, name)
		if err != nil {
			return errors.Wrap(err, "setup chain", "chain", name)
		}

		if err := fn(ctx, s, c); err != nil {
			return errors.Wrap(err, "chain", "chain", name)
		}
	}

	return nil
}

// maybeAll returns all chain names if chain is "all", otherwise returns chain.
func maybeAll(network netconf.Network, chain string) []string {
	if chain == chainAll {
		var names []string

		for _, c := range network.EVMChains() {
			names = append(names, c.Name)
		}

		return names
	}

	return []string{chain}
}

// runForge runs a forge script against an rpc, returning the ouptut.
// if the senders are known anvil accounts, it will sign with private keys directly.
// otherwise, it will use the unlocked flag.
func runForge(ctx context.Context, script string, rpc string, input []byte, senders ...common.Address,
) (string, error) {
	// assumes running from root
	dir := "./contracts/core"

	anvilPks := make([]string, 0, len(senders))
	for _, sender := range senders {
		pk, ok := anvil.PrivateKey(sender)
		if !ok {
			continue
		}

		anvilPks = append(anvilPks, hexutil.EncodeBig(pk.D))
	}

	if len(anvilPks) > 0 && len(anvilPks) != len(senders) {
		return "", errors.New("cannot mix anvil and non-anvil accounts")
	}

	// for dev anvil accounts, we signed with private key directly
	// TODO: consider refactoring fbproxy to "txproxy" that allows for different
	// signing methods (one fireblocks, one providing private keys directly)
	if len(anvilPks) > 0 {
		return execCmd(ctx, dir, "forge", "script", script,
			"--broadcast", "--slow", "--rpc-url", rpc, "--sig", hexutil.Encode(input),
			"--private-keys", strings.Join(dedup(anvilPks), ","),
		)
	}

	return execCmd(ctx, dir, "forge", "script", script, "--timeout", "300", // 5 minute timeout
		"--broadcast", "--slow", "--rpc-url", rpc, "--sig", hexutil.Encode(input), "--unlocked",
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

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

func dedup(strs []string) []string {
	seen := make(map[string]bool)
	var res []string
	for _, s := range strs {
		if seen[s] {
			continue
		}

		seen[s] = true
		res = append(res, s)
	}

	return res
}
