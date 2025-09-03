package admin

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	fbproxy "github.com/omni-network/omni/e2e/fbproxy/app"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// adminABI is the ABI for the Admin script contract.
var adminABI = mustGetABI(bindings.AdminMetaData)

// solverNetAdminABI is the ABI for the SolverNetAdmin script contract.
var solverNetAdminABI = mustGetABI(bindings.SolverNetAdminMetaData)

// omniEVMName is the name of the omni EVM chain.
const omniEVMName = "omni_evm"

// adminScriptName is the name of the admin script contract.
const adminScriptName = "Admin"

// solverNetAdminScriptName is the name of the SolverNetAdmin script contract.
const solverNetAdminScriptName = "SolverNetAdmin"

// coreContracts is the path to the core contracts.
const coreContracts = "./contracts/core"

// solveContracts is the path to the SolverNet contracts.
const solveContracts = "./contracts/solve"

// shared contains common resources for all admin operations.
type shared struct {
	manager     common.Address
	upgrader    common.Address
	deployer    common.Address
	endpoints   xchain.RPCEndpoints
	testnet     types.Testnet
	cfg         Config
	fireAPIKey  string
	fireKeyPath string
}

// chain contains chain specific resources, extending EVMChain with the PortalAddress and an RPCEndpoint endpoint.
// Netconf.Chain cannot be used since we don't know the DeployHeight.
type chain struct {
	types.EVMChain
	PortalAddress common.Address
	RPCEndpoint   string
}

// setup returns common resources for all admin operations.
func setup(def app.Definition, cfg Config) shared {
	netID := def.Testnet.Network
	manager := eoa.MustAddress(netID, eoa.RoleManager)
	upgrader := eoa.MustAddress(netID, eoa.RoleUpgrader)
	deployer := eoa.MustAddress(netID, eoa.RoleDeployer)
	endpoints := app.ExternalEndpoints(def)

	// addrs set lazily in setupChain

	return shared{
		manager:     manager,
		upgrader:    upgrader,
		deployer:    deployer,
		endpoints:   endpoints,
		testnet:     def.Testnet,
		cfg:         cfg,
		fireAPIKey:  def.Cfg.FireAPIKey,
		fireKeyPath: def.Cfg.FireKeyPath,
	}
}

// setupChain returns chain specific resources.
// starts and fbproxy for non-devnet chains.
func setupChain(ctx context.Context, s shared, name string) (chain, error) {
	c, ok := s.testnet.EVMChainByName(name)
	if !ok {
		return chain{}, errors.New("chain not found", "chain", name)
	}

	rpc, err := s.endpoints.ByNameOrID(c.Name, c.ChainID)
	if err != nil {
		return chain{}, errors.Wrap(err, "rpc endpoint")
	}

	addrs, err := contracts.GetAddresses(ctx, s.testnet.Network)
	if err != nil {
		return chain{}, errors.Wrap(err, "get addresses")
	}

	if s.fireAPIKey != "" || s.fireKeyPath != "" {
		rpc, err = startFBProxy(ctx, s.testnet.Network, rpc, s.fireAPIKey, s.fireKeyPath)
		if err != nil {
			return chain{}, errors.Wrap(err, "start fb proxy")
		}
	}

	return chain{
		EVMChain:      c,
		PortalAddress: addrs.Portal,
		RPCEndpoint:   rpc,
	}, nil
}

// setupChainHL returns chain specific resources for all chains (including Hyperlane chains).
// starts an fbproxy for non-devnet chains.
func setupChainHL(ctx context.Context, s shared, c netconf.Chain) (chain, error) {
	evmchain, err := getEVMChain(s, c)
	if err != nil {
		return chain{}, errors.Wrap(err, "get evm chain")
	}

	rpc, err := s.endpoints.ByNameOrID(evmchain.Name, evmchain.ChainID)
	if err != nil {
		return chain{}, errors.Wrap(err, "unknown chain", "chain_name", c.Name)
	}

	if s.fireAPIKey != "" || s.fireKeyPath != "" {
		rpc, err = startFBProxy(ctx, s.testnet.Network, rpc, s.fireAPIKey, s.fireKeyPath)
		if err != nil {
			return chain{}, errors.Wrap(err, "start fb proxy")
		}
	}

	addrs, err := contracts.GetAddresses(ctx, s.testnet.Network)
	if err != nil {
		return chain{}, errors.Wrap(err, "get addresses")
	}

	chain := chain{
		EVMChain:      evmchain,
		PortalAddress: addrs.Portal,
		RPCEndpoint:   rpc,
	}
	if solvernet.IsHLOnly(evmchain.ChainID) {
		chain.PortalAddress = common.Address{}
	}

	return chain, nil
}

// getEVMChain returns the EVMChain for a given netconf chain, with special handling for the Omni EVM.
func getEVMChain(s shared, c netconf.Chain) (types.EVMChain, error) {
	if c.Name == omniEVMName {
		return types.OmniEVMByNetwork(s.testnet.Network), nil
	}

	return types.PublicChainByName(c.Name)
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

// addHLEndpoints adds the Hyperlane endpoints to the shared resources.
func (s shared) addHLEndpoints(ctx context.Context, def app.Definition) (shared, error) {
	endpoints, err := app.AddSolverEndpoints(ctx, s.testnet.Network, s.endpoints, def.Cfg.RPCOverrides)
	if err != nil {
		return shared{}, errors.Wrap(err, "add solver endpoints")
	}

	newS := s
	newS.endpoints = endpoints

	return newS, nil
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

	names, err := maybeAll(s.testnet.EVMChains(), s.cfg.Chain, opts.exclude)
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

// runHL runs a function for all configured chains (including Hyperlane chains).
// NOTE: currently does not include options, network is passed as shared.testnet.EVMChains() doesn't reflect the updated network.
func (s shared) runHL(ctx context.Context, def app.Definition, fn func(context.Context, shared, netconf.Network, chain) error) error {
	_s, err := s.addHLEndpoints(ctx, def)
	if err != nil {
		return errors.Wrap(err, "add hl endpoints")
	}

	network, err := app.SolverNetworkFromDef(ctx, def)
	if err != nil {
		return errors.Wrap(err, "network from def")
	}
	network = solvernet.AddNetwork(ctx, network, solvernet.FilterByContracts(ctx, _s.endpoints))

	for _, _chain := range network.EVMChains() {
		if s.cfg.Chain != "" && s.cfg.Chain != _chain.Name {
			// If set, skip all but the specified chain.
			continue
		}

		c, err := setupChainHL(ctx, _s, _chain)
		if err != nil {
			return errors.Wrap(err, "setup chain hl", "chain", _chain.Name)
		}

		if err := fn(ctx, _s, network, c); err != nil {
			return errors.Wrap(err, "chain", "chain", _chain.Name)
		}
	}

	return nil
}

// maybeAll returns all chains if chain is empty, otherwise returns chain.
func maybeAll(chains []types.EVMChain, chain string, exclude []string) ([]string, error) {
	excluded := make(map[string]bool)
	for _, e := range exclude {
		excluded[e] = true
	}

	if chain == "" {
		var resp []string

		for _, c := range chains {
			if excluded[c.Name] {
				continue
			}

			resp = append(resp, c.Name)
		}

		return resp, nil
	}

	if excluded[chain] {
		return nil, errors.New("unsupported chain", "chain", chain)
	}

	return []string{chain}, nil
}

func (s shared) runForge(ctx context.Context, rpc string, script string, dir string, input []byte, senders ...common.Address,
) (string, error) {
	resume := false
	attempts := 0
	const maxAttempts = 6

	// Retrying lets us resume resumes submitting transactions that failed or timed-out previously.
	// This is a temporarary workaround until foundry lets us increase tx timeout.
	// See https://github.com/foundry-rs/foundry/issues/9303

	for {
		if attempts >= 1 {
			resume = true
			time.Sleep(5 * time.Second) // sleep to allow RPCs to catch up if needed
		}

		out, err := runForgeOnce(ctx, rpc, script, dir, input, s.cfg.Broadcast, resume, senders...)
		if err == nil {
			return out, nil
		}

		attempts++
		if attempts >= maxAttempts {
			return out, errors.Wrap(err, "max attempts reached", "out", out)
		}

		log.Warn(ctx, "Run failed, will retry", err, "attempts", attempts, "remaining", maxAttempts-attempts)
	}
}

// runForgeOnce runs an Admin forge script against an rpc, returning the ouptut.
// if the senders are known anvil accounts, it will sign with private keys directly.
// otherwise, it will use the unlocked flag.
func runForgeOnce(ctx context.Context, rpc string, script string, dir string, input []byte, broadcast, resume bool, senders ...common.Address,
) (string, error) {
	pks := make([]string, 0, len(senders))
	for _, sender := range senders {
		pk, ok := eoa.DevPrivateKey(sender)
		if !ok {
			continue
		}
		pks = append(pks, hexutil.EncodeBig(pk.D))
	}
	if len(pks) > 0 && len(pks) != len(senders) {
		return "", errors.New("cannot mix anvil and non-anvil accounts")
	}

	args := []string{
		"script", script,
		"--slow",                                   // wait for each tx to succed before sending the next
		"--block-gas-limit", "1000000000000000000", // much higher than default to avoid gas limit errors
		"--rpc-url", rpc, // rpc endpoint, fb proxy for non-devnet
		"--sig", hexutil.Encode(input), // Admin.sol calldata
	}

	if broadcast {
		// if omitted, transactions are not broadcasted
		args = append(args, "--broadcast")
	}

	if resume {
		// resumes submitting transactions that failed or timed-out previously.
		args = append(args, "--resume")
	}

	if len(pks) > 0 {
		// for dev anvil accounts, we sign with privates key directly
		args = append(args, "--private-keys", strings.Join(dedup(pks), ","))
	} else {
		// else, we use --unlocked flag, to send unsigned eth_sendTransaction requests
		// with 15 minute timeout, to allow for fireblocks signing.
		args = append(args, "--timeout", "900", "--unlocked")
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

func solverNetInboxInitializer() ([]byte, error) {
	// var inboxABI = mustGetABI(bindings.SolverNetInboxMetaData)
	// TODO: replace if re-initialization is required
	return []byte{}, nil
}

func solverNetOutboxInitializer() ([]byte, error) {
	// var outboxABI = mustGetABI(bindings.SolverNetOutboxMetaData)
	// TODO: replace if re-initialization is required
	return []byte{}, nil
}

func solverNetExecutorInitializer() ([]byte, error) {
	// var executorABI = mustGetABI(bindings.SolverNetExecutorMetaData)
	// TODO: replace if re-initialization is required
	return []byte{}, nil
}
