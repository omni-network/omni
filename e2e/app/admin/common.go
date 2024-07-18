package admin

import (
	"context"
	"os/exec"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	fbproxy "github.com/omni-network/omni/e2e/fbproxy/app"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/forkjoin"
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

// runner is an interface for running Admin forge scripts.
type runner interface {
	run(ctx context.Context, input []byte, senders ...common.Address) (string, error)
}

// forgeRunner is a runner that runs forge scripts against an rpc.
type forgeRunner struct {
	rpc string
}

var _ runner = forgeRunner{}

// run runs an Admin forge script against an rpc, returning the output.
func (r forgeRunner) run(ctx context.Context, input []byte, senders ...common.Address) (string, error) {
	return runForge(ctx, "Admin", r.rpc, input, senders...)
}

// action is a function that performs an admin action. It returns the output of the action and an error.
type action func(ctx context.Context, s shared, c chain, r runner) (string, error)

// run runs an admin action with some config on an app definition.
func run(ctx context.Context, def app.Definition, cfg PortalAdminConfig, name string, act action) error {
	if err := cfg.Validate(); err != nil {
		return errors.Wrap(err, "validate config")
	}

	s, err := setup(ctx, def)
	if err != nil {
		return errors.Wrap(err, "setup")
	}

	chains := getChains(s.network, cfg.Chain)

	// runForChain runs the action for a single chain.
	runForChain := func(ctx context.Context, chain string) (string, error) {
		c, err := setupChain(ctx, s, chain)
		if err != nil {
			return "", errors.Wrap(err, "setup chain")
		}

		r := forgeRunner{rpc: c.rpc}

		out, err := act(ctx, s, c, r)
		if err != nil {
			return out, errors.Wrap(err, "run action", "action", name, "chain", chain)
		}

		return out, nil
	}

	results, cancel := forkjoin.NewWithInputs(ctx, runForChain, chains)
	defer cancel()

	return report(log.WithCtx(ctx, "action", name), results)
}

// getChains returns the chain names to run an admin action on, returning all chains if chain is "all".
func getChains(network netconf.Network, chain string) []string {
	if chain == chainAll {
		var chains []string
		for _, c := range network.EVMChains() {
			chains = append(chains, c.Name)
		}

		return chains
	}

	return []string{chain}
}

// runResults is a type alias for the results of a forkjoin of admin actions.
// They take a chain name as input, and return cli output.
type runResults = forkjoin.Results[string, string]

// report logs the results of a forkjoin, returning an error if any runs failed.
func report(ctx context.Context, results runResults) error {
	var failed []string
	var success []string
	runs := 0

	for res := range results {
		runs++

		if res.Err != nil {
			log.Error(ctx, "Run  failed", res.Err, "chain", res.Input)
			failed = append(failed, res.Input)
		} else {
			log.Info(ctx, "Run succeeded", "chain", res.Input, "out", res.Output)
			success = append(success, res.Input)
		}
	}

	if len(failed) == runs {
		return errors.New("all runs failed", "chains", failed)
	} else if len(failed) > 0 {
		log.Error(ctx, "Runs failed", errors.New("runs failed"), "chains", failed)
		log.Info(ctx, "Runs succeeded", "chains", success)

		return errors.New("some runs failed", "chains", failed)
	}

	log.Info(ctx, "All runs succeeded", "chains", success)

	return nil
}

// setup returns common resources for all admin operations.
func setup(ctx context.Context, def app.Definition) (shared, error) {
	netID := def.Testnet.Network

	admin, ok := eoa.Address(netID, eoa.RoleAdmin)
	if !ok {
		return shared{}, errors.New("admin eoas not found", "network", netID)
	}

	deployer, ok := eoa.Address(netID, eoa.RoleDeployer)
	if !ok {
		return shared{}, errors.New("deployer eoas not found", "network", netID)
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

	return execCmd(ctx, dir, "forge", "script", script,
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
