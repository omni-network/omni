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
	"github.com/omni-network/omni/lib/contracts"
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
	manager     common.Address
	upgrader    common.Address
	deployer    common.Address
	endpoints   xchain.RPCEndpoints
	network     netconf.Network
	cfg         Config
	fireAPIKey  string
	fireKeyPath string
}

// chain contains chain specific resources, exteding netconf.Chain with an rpc endpoint.
type chain struct {
	netconf.Chain
	rpc string
}

// setup returns common resources for all admin operations.
func setup(def app.Definition, cfg Config) shared {
	netID := def.Testnet.Network
	owner := eoa.MustAddress(netID, eoa.RoleManager)
	upgrader := eoa.MustAddress(netID, eoa.RoleUpgrader)
	deployer := eoa.MustAddress(netID, eoa.RoleDeployer)
	endpoints := app.ExternalEndpoints(def)
	network := app.NetworkFromDef(def)

	// addrs set lazily in setupChain

	return shared{
		manager:     owner,
		upgrader:    upgrader,
		deployer:    deployer,
		endpoints:   endpoints,
		network:     network,
		cfg:         cfg,
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

	// add portal address if not already set
	if c.PortalAddress == (common.Address{}) {
		addrs, err := contracts.GetAddresses(ctx, s.network.ID)
		if err != nil {
			return chain{}, errors.Wrap(err, "get addresses")
		}

		c.PortalAddress = addrs.Portal
	}

	if s.fireAPIKey != "" || s.fireKeyPath != "" {
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
	fn func(context.Context, shared, chain) error,
	options ...runOpt,
) error {
	opts := runOpts{}
	for _, o := range options {
		o(&opts)
	}

	names, err := maybeAll(s.network, s.cfg.Chain, opts.exclude)
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

func (s shared) runForge(ctx context.Context, rpc string, input []byte, senders ...common.Address,
) (string, error) {
	return runForge(ctx, rpc, input, s.cfg.Broadcast, senders...)
}

// runForge runs an Admin forge script against an rpc, returning the ouptut.
// if the senders are known anvil accounts, it will sign with private keys directly.
// otherwise, it will use the unlocked flag.
func runForge(ctx context.Context, rpc string, input []byte, broadcast bool, senders ...common.Address,
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

	args := []string{
		"script", script,
		"--slow",         // wait for each tx to succed before sending the next
		"--rpc-url", rpc, // rpc endpoint, fb proxy for non-devnet
		"--sig", hexutil.Encode(input), // Admin.sol calldata
	}

	if broadcast {
		// if omitted, transactions are not broadcasted
		args = append(args, "--broadcast")
	}

	if len(anvilPks) > 0 {
		// for dev anvil accounts, we sign with privates key directly
		args = append(args, "--private-keys", strings.Join(dedup(anvilPks), ","))
	} else {
		// else, we use --unlocked flag, to send unsigned eth_sendTransaction requests
		// with 5 minute timeout, to allow for fireblocks signing.
		args = append(args, "--timeout", "300", "--unlocked")
	}

	return execCmd(ctx, dir, "forge", args...)
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
