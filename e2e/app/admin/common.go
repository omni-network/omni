package admin

import (
	"context"
	"os/exec"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
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

// adminABI is the ABI for the Admin script contract.
var adminABI = mustGetABI(bindings.AdminMetaData)

// omniEVMName is the name of the omni EVM chain.
const omniEVMName = "omni_evm"

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
func setup(def app.Definition) shared {
	netID := def.Testnet.Network
	admin := eoa.MustAddress(netID, eoa.RoleAdmin)
	deployer := eoa.MustAddress(netID, eoa.RoleDeployer)
	endpoints := app.ExternalEndpoints(def)
	network := app.NetworkFromDef(def)

	return shared{
		admin:       admin,
		deployer:    deployer,
		endpoints:   endpoints,
		network:     network,
		fireAPIKey:  def.Cfg.FireAPIKey,
		fireKeyPath: def.Cfg.FireKeyPath,
	}
}

// setupChain returns chain specific resources.
// starts and fbproxy for non-devnet chains.
func setupChain(ctx context.Context, s shared, name string) (chain, error) {
	c, ok := s.network.ChainByName(name)
	if !ok {
		return chain{}, errors.New("chain not found", "chain", name)
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

type runOpts struct {
	// exclude chains by name.
	exclude []string
}

type runOpt func(*runOpts)

func withExclude(names ...string) runOpt {
	return func(o *runOpts) {
		o.exclude = append(o.exclude, names...)
	}
}

// run runs a function for all configured chains (all applicable, if not specified).
func (s shared) run(
	ctx context.Context,
	cfg Config,
	fn func(context.Context, shared, chain) error,
	options ...runOpt,
) error {
	opts := runOpts{}
	for _, o := range options {
		o(&opts)
	}

	names, err := maybeAll(s.network, cfg.Chain, opts.exclude)
	if err != nil {
		return err
	}

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

// maybeAll returns all chains if chain is empty, otherwise returns chain.
func maybeAll(network netconf.Network, chain string, exclude []string) ([]string, error) {
	excluded := make(map[string]bool)
	for _, e := range exclude {
		excluded[e] = true
	}

	if chain == "" {
		var chains []string

		for _, c := range network.EVMChains() {
			if excluded[c.Name] {
				continue
			}

			chains = append(chains, c.Name)
		}

		return chains, nil
	}

	if excluded[chain] {
		return nil, errors.New("unsupported chain", "chain", chain)
	}

	return []string{chain}, nil
}

// runForge runs an Admin forge script against an rpc, returning the ouptut.
// if the senders are known anvil accounts, it will sign with private keys directly.
// otherwise, it will use the unlocked flag.
func runForge(ctx context.Context, rpc string, input []byte, senders ...common.Address,
) (string, error) {
	// name of admin forge script in contracts/core
	const script = "Admin"

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

	// for dev anvil accounts, we sign with privates key directly
	if len(anvilPks) > 0 {
		return execCmd(ctx, dir, "forge", "script", script,
			"--broadcast", "--slow", "--rpc-url", rpc, "--sig", hexutil.Encode(input),
			"--private-keys", strings.Join(dedup(anvilPks), ","),
		)
	}

	return execCmd(ctx, dir, "forge", "script", script, "--timeout", "300", // 5 minute timeout, allow for fb signing
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
