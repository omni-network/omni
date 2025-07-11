package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/app/geth"
	"github.com/omni-network/omni/e2e/app/static"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/e2e/vmcompose"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/halo/genutil"
	evmgenutil "github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	monapp "github.com/omni-network/omni/monitor/app"
	relayapp "github.com/omni-network/omni/relayer/app"
	solverapp "github.com/omni-network/omni/solver/app"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/pex"
	"github.com/cometbft/cometbft/privval"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	_ "embed" // Embed requires blank import
)

const (
	AppAddressTCP  = "tcp://127.0.0.1:30000"
	AppAddressUNIX = "unix:///var/run/app.sock"

	PrivvalKeyFile   = "config/priv_validator_key.json"
	PrivvalStateFile = "data/priv_validator_state.json"
)

// Setup sets up the testnet configuration.
//
//nolint:gocyclo // Multiple steps
func Setup(ctx context.Context, def Definition, depCfg DeployConfig) error {
	log.Info(ctx, "Setup testnet", "dir", def.Testnet.Dir)

	if err := CleanupDir(ctx, def.Testnet.Dir); err != nil {
		return err
	}

	if err := os.MkdirAll(def.Testnet.Dir, os.ModePerm); err != nil {
		return errors.Wrap(err, "mkdir")
	}

	// set global feature flags, as they might inform setup / tests
	feature.SetGlobals(def.Manifest.FeatureFlags)

	// Setup geth execution genesis
	gethGenesis, err := evmgenutil.MakeGenesis(def.Manifest.Network)
	if err != nil {
		return errors.Wrap(err, "make genesis")
	}
	gethGenesisBz, err := evmgenutil.MarshallBackwardsCompatible(gethGenesis)
	if err != nil {
		return errors.Wrap(err, "marshal genesis")
	}
	if err := geth.WriteAllConfig(def.Testnet, gethGenesis); err != nil {
		return err
	}

	// Setup halo consensus genesis
	var vals []crypto.PubKey
	var valPrivKeys []crypto.PrivKey
	for val := range def.Testnet.Validators {
		vals = append(vals, val.PrivvalKey.PubKey())
		valPrivKeys = append(valPrivKeys, val.PrivvalKey)
	}

	cosmosGenesis, err := genutil.MakeGenesis(
		ctx,
		def.Manifest.Network,
		time.Now(),
		gethGenesis.ToBlock().Hash(),
		def.Manifest.EphemeralGenesis,
		vals...)
	if err != nil {
		return errors.Wrap(err, "make genesis")
	}
	cmtGenesis, err := cosmosGenesis.ToGenesisDoc()
	if err != nil {
		return errors.Wrap(err, "convert genesis")
	}

	logCfg := logConfig(def)
	if err := writeMonitorConfig(ctx, def, logCfg, valPrivKeys); err != nil {
		return err
	}

	if err := writeRelayerConfig(ctx, def, logCfg); err != nil {
		return err
	}

	if err := writeSolverConfig(ctx, def, logCfg); err != nil {
		return err
	}

	if err := writeAnvilState(def.Testnet); err != nil {
		return err
	}

	for _, node := range def.Testnet.Nodes {
		nodeDir := filepath.Join(def.Testnet.Dir, node.Name)

		dirs := []string{
			filepath.Join(nodeDir, "config"),
			filepath.Join(nodeDir, "data"),
		}
		for _, dir := range dirs {
			err := os.MkdirAll(dir, 0o755)
			if err != nil {
				return errors.Wrap(err, "make dir")
			}
		}

		cfg, err := MakeConfig(def.Testnet, node, nodeDir)
		if err != nil {
			return err
		}
		if err := halocmd.WriteCometConfig(filepath.Join(nodeDir, "config", "config.toml"), cfg); err != nil {
			return err
		}

		if err := writeHaloAddressBook(def.Testnet.Network, filepath.Join(nodeDir, "config", "addrbook.json"), node); err != nil {
			return err
		}

		endpoints := internalEndpoints(def, node.Name)
		omniEVM := omniEVMByPrefix(def.Testnet, node.Name)

		if err := writeHaloConfig(
			def,
			nodeDir,
			logCfg,
			depCfg.testConfig,
			node,
			omniEVM.InstanceName,
			endpoints,
		); err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(nodeDir, "config", "jwtsecret"), []byte(omniEVM.JWTSecret), 0o600); err != nil {
			return errors.Wrap(err, "write jwtsecret")
		}

		err = cmtGenesis.SaveAs(filepath.Join(nodeDir, "config", "genesis.json"))
		if err != nil {
			return errors.Wrap(err, "write genesis")
		}
		if err := os.WriteFile(filepath.Join(nodeDir, "config", "execution_genesis.json"), gethGenesisBz, 0o644); err != nil {
			return errors.Wrap(err, "write execution genesis")
		}

		err = (&p2p.NodeKey{PrivKey: node.NodeKey}).SaveAs(filepath.Join(nodeDir, "config", "node_key.json"))
		if err != nil {
			return errors.Wrap(err, "write node key")
		}

		(privval.NewFilePV(node.PrivvalKey,
			filepath.Join(nodeDir, PrivvalKeyFile),
			filepath.Join(nodeDir, PrivvalStateFile),
		)).Save()

		// Initialize the node's data directory (with noop logger since it is noisy).
		initCfg := halocmd.InitConfig{
			HomeDir: nodeDir,
			Network: def.Testnet.Network,
			HaloCfgFunc: func(cfg *halocfg.Config) {
				cfg.RPCEndpoints = endpoints
			},
			Force: true,
		}
		if err := halocmd.InitFiles(log.WithNoopLogger(ctx), initCfg); err != nil {
			return errors.Wrap(err, "init files")
		}
	}

	// Write an external network.json and endpoints.json in base testnet dir.
	// This allows for easy connecting or querying of the network
	endpoints := ExternalEndpoints(def)
	if endpointBytes, err := json.MarshalIndent(endpoints, "", " "); err != nil {
		return errors.Wrap(err, "marshal endpoints")
	} else if err := os.WriteFile(filepath.Join(def.Testnet.Dir, "endpoints.json"), endpointBytes, 0o644); err != nil {
		return errors.Wrap(err, "write endpoints")
	}

	if def.Testnet.Prometheus {
		if err := agent.WriteConfig(ctx, def.Testnet, def.Cfg.AgentSecrets); err != nil {
			return errors.Wrap(err, "write prom config")
		}
	}

	if err := svmSetup(def.Testnet); err != nil {
		return err
	}

	if err := def.Infra.Setup(); err != nil {
		return errors.Wrap(err, "setup provider")
	}

	return nil
}

func svmSetup(testnet types.Testnet) error {
	if len(testnet.SVMChains) == 0 {
		return nil
	}

	if err := os.MkdirAll(filepath.Join(testnet.Dir, "svm"), 0o755); err != nil {
		return errors.Wrap(err, "mkdir svm dir")
	}

	return nil
}

// writeAnvilState writes the embedded /static/el-anvil-state.json
// to <testnet.Dir>/anvil/state.json for use by all anvil chains.
func writeAnvilState(testnet types.Testnet) error {
	anvilStateFile := filepath.Join(testnet.Dir, "anvil", "state.json")
	if err := os.MkdirAll(filepath.Dir(anvilStateFile), 0o755); err != nil {
		return errors.Wrap(err, "mkdir")
	}
	if err := os.WriteFile(anvilStateFile, static.GetDevnetElAnvilState(), 0o644); err != nil {
		return errors.Wrap(err, "write anvil state")
	}

	return nil
}

// MakeConfig generates a CometBFT config for a node.
func MakeConfig(testnet types.Testnet, node *e2e.Node, nodeDir string) (*config.Config, error) {
	if node.ABCIProtocol != e2e.ProtocolBuiltin {
		return nil, errors.New("only Builtin ABCI is supported")
	}

	cfg := halocmd.DefaultCometConfig(nodeDir)
	cfg.Moniker = node.Name
	cfg.RPC.ListenAddress = "tcp://0.0.0.0:26657"
	cfg.RPC.PprofListenAddress = ":6060"
	cfg.P2P.ExternalAddress = fmt.Sprintf("tcp://%v:26656", advertisedIP(testnet.Network, node.Mode, node.InternalIP, node.ExternalIP))
	cfg.DBBackend = node.Database
	cfg.StateSync.DiscoveryTime = 5 * time.Second
	cfg.BlockSync.Version = node.BlockSyncVersion

	// CometBFT errors if it does not have a privval key set up, regardless of whether
	// it's actually needed (e.g. for remote KMS or non-validators). We set up a dummy
	// key here by default, and use the real key for actual validators that should use
	// the file privval.
	cfg.PrivValidatorListenAddr = ""
	cfg.PrivValidatorKey = PrivvalKeyFile
	cfg.PrivValidatorState = PrivvalStateFile

	if node.PrivvalProtocol != e2e.ProtocolFile {
		return nil, errors.New("only PrivvalKeyFile is supported")
	}

	if testnet.Network == netconf.Staging {
		cfg.LogLevel = "debug" // Debug log levels on staging
	}

	// Try disabling seedmode to fix joining network issues.
	// if node.Mode == types.ModeSeed {
	// cfg.P2P.SeedMode = true
	// }

	if node.StateSync {
		cfg.StateSync.Enable = true
		cfg.StateSync.RPCServers = []string{}
		for _, peer := range node.Testnet.ArchiveNodes() {
			if peer.Name == node.Name {
				continue
			}
			cfg.StateSync.RPCServers = append(cfg.StateSync.RPCServers, peer.AddressRPC())
		}
		if len(cfg.StateSync.RPCServers) < 2 {
			return nil, errors.New("unable to find 2 suitable state sync RPC servers")
		}
	}

	cfg.P2P.Seeds = ""
	for _, seed := range node.Seeds {
		if len(cfg.P2P.Seeds) > 0 {
			cfg.P2P.Seeds += ","
		}
		cfg.P2P.Seeds += advertisedP2PAddr(testnet.Network, seed)
	}

	cfg.P2P.PersistentPeers = ""
	for _, peer := range node.PersistentPeers {
		if len(cfg.P2P.PersistentPeers) > 0 {
			cfg.P2P.PersistentPeers += ","
		}
		cfg.P2P.PersistentPeers += advertisedP2PAddr(testnet.Network, peer)
	}

	cfg.P2P.PrivatePeerIDs = privatePeerIDs(testnet, node)
	if !isPublicNode(testnet.Network, node.Mode) {
		cfg.P2P.AddrBookStrict = false // Strict addresses only supported by public nodes.
	}

	if node.Prometheus {
		cfg.Instrumentation.Prometheus = true
	}

	return &cfg, nil
}

// privatePeerIDs returns a comma-separated list of private peer IDs that should not be gossiped.
func privatePeerIDs(testnet types.Testnet, self *e2e.Node) string {
	var ids []string
	for _, node := range testnet.Nodes {
		if node.Name == self.Name {
			continue // Skip self
		}
		if isPublicNode(testnet.Network, node.Mode) {
			continue
		}
		ids = append(ids, fmt.Sprintf("%x", node.NodeKey.PubKey().Address().Bytes()))
	}

	return strings.Join(ids, ",")
}

// advertisedP2PAddr returns the cometBFT network address <ID@IP:port> to advertise for a node.
func advertisedP2PAddr(network netconf.ID, node *e2e.Node) string {
	id := node.NodeKey.PubKey().Address().Bytes()
	ip := advertisedIP(network, node.Mode, node.InternalIP, node.ExternalIP)

	return fmt.Sprintf("%x@%s:26656", id, ip)
}

func advertisedIP(network netconf.ID, mode e2e.Mode, internal net.IP, external net.IP) net.IP {
	if isPublicNode(network, mode) {
		return external
	}

	return internal
}

// isPublicNode returns true if the node should be publicly accessible;
// i.e., allow inbound connections from external peers.
func isPublicNode(network netconf.ID, mode types.Mode) bool {
	if network == netconf.Devnet {
		// Devnet does not support external peers connecting to it, so we use the internal IP.
		return false
	}

	if mode == types.ModeSeed || mode == types.ModeFull || mode == types.ModeArchive {
		// Only seeds and fullnodes and archives allow external peers to connect to them.
		return true
	}

	// Validators nodes are "secured" and only allow internal peers to connect to them.

	return false
}

// writeHaloAddressBook pre-populates the halo address book for a node.
// All persisted peers are added. This aids seed nodes that don't seem
// to add persisted peer consistently.
func writeHaloAddressBook(network netconf.ID, path string, node *e2e.Node) error {
	addrBook := pex.NewAddrBook(path, false)
	for _, peer := range node.PersistentPeers {
		addr := advertisedP2PAddr(network, peer)
		netAddr, err := p2p.NewNetAddressString(addr)
		if err != nil {
			return errors.Wrap(err, "parse net address")
		}
		if err := addrBook.AddAddress(netAddr, netAddr); err != nil {
			return errors.Wrap(err, "add address")
		}
	}
	addrBook.Save()

	return nil
}

// writeHaloConfig generates an halo application config for a node and writes it to disk.
func writeHaloConfig(
	def Definition,
	nodeDir string,
	logCfg log.Config,
	testCfg bool,
	node *e2e.Node,
	evmInstance string,
	endpoints xchain.RPCEndpoints,
) error {
	cfg := halocfg.DefaultConfig()

	if node.Mode == types.ModeArchive {
		cfg.PruningOption = "nothing"
		// Setting this to 0 retains all blocks
		cfg.MinRetainBlocks = 0
	}

	cfg.Network = def.Testnet.Network
	cfg.HomeDir = nodeDir
	cfg.RPCEndpoints = endpoints
	cfg.EngineEndpoint = fmt.Sprintf("http://%s:8551", evmInstance) //nolint:nosprintfhostport // net.JoinHostPort doesn't prefix http.
	cfg.EngineJWTFile = "/halo/config/jwtsecret"                    // Absolute path inside docker container
	cfg.Tracer.Endpoint = def.Cfg.TracingEndpoint
	cfg.Tracer.Headers = def.Cfg.TracingHeaders
	cfg.FeatureFlags = filterFeatureFlags(node, def.Manifest.FeatureFlags)
	cfg.SDKGRPC.Address = "0.0.0.0:9999" // VM port 9090 used by grafana-agent, so use 9999 instead.

	if testCfg {
		cfg.SnapshotInterval = 1   // Write snapshots each block in e2e tests
		cfg.SnapshotKeepRecent = 0 // Keep all snapshots in e2e tests
	}

	if def.Testnet.Network == netconf.Staging {
		logCfg.Level = log.LevelDebug // Debug log levels on staging
	}

	// Enable Proxying of EVM requests
	cfg.EVMProxyListen = "0.0.0.0:8545"
	cfg.EVMProxyTarget = fmt.Sprintf("http://%s:8545", evmInstance) //nolint:nosprintfhostport // net.JoinHostPort doesn't prefix http.

	return halocfg.WriteConfigTOML(cfg, logCfg)
}

// updateConfigStateSync updates the state sync config for a node.
func updateConfigStateSync(nodeDir string, height int64, hash []byte) error {
	cfgPath := filepath.Join(nodeDir, "config", "config.toml")

	// FIXME Apparently there's no function to simply load a config file without
	// involving the entire Viper apparatus, so we'll just resort to regexps.
	bz, err := os.ReadFile(cfgPath)
	if err != nil {
		return errors.Wrap(err, "read config")
	}

	before := string(bz)

	bz = regexp.MustCompile(`(?m)^trust_height =.*`).ReplaceAll(bz, []byte(fmt.Sprintf(`trust_height = %v`, height)))
	bz = regexp.MustCompile(`(?m)^trust_hash =.*`).ReplaceAll(bz, []byte(fmt.Sprintf(`trust_hash = "%X"`, hash)))
	bz = regexp.MustCompile(`(?m)^log_level =.*`).ReplaceAll(bz, []byte(`log_level = "info"`)) // Increase log level.

	after := string(bz)
	if before == after {
		return errors.New("no changes to config")
	}

	if err := os.WriteFile(cfgPath, bz, 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}

func writeRelayerConfig(ctx context.Context, def Definition, logCfg log.Config) error {
	confRoot := filepath.Join(def.Testnet.Dir, "relayer")

	const (
		privKeyFile = "privatekey"
		configFile  = "relayer.toml"
	)

	if err := os.MkdirAll(confRoot, 0o755); err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	// Save network config
	endpoints := internalEndpoints(def, "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		endpoints = ExternalEndpoints(def)
	}

	// Save private key
	privKey, err := eoa.PrivateKey(ctx, def.Testnet.Network, eoa.RoleRelayer)
	if err != nil {
		return errors.Wrap(err, "get relayer key")
	}
	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, privKeyFile), privKey); err != nil {
		return errors.Wrap(err, "write private key")
	}

	archiveNode, ok := def.Testnet.ArchiveNode()
	if !ok {
		return errors.New("archive node not found")
	}

	relayCfg := relayapp.DefaultConfig()
	relayCfg.PrivateKey = privKeyFile
	relayCfg.Network = def.Testnet.Network
	relayCfg.HaloCometURL = archiveNode.AddressRPC()
	relayCfg.HaloGRPCURL = haloGRPCAddress(archiveNode)
	relayCfg.RPCEndpoints = endpoints
	relayCfg.CoinGeckoAPIKey = def.Cfg.CoinGeckoAPIKey

	if err := relayapp.WriteConfigTOML(relayCfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write relayer config")
	}

	return nil
}

func writeSolverConfig(ctx context.Context, def Definition, logCfg log.Config) error {
	confRoot := filepath.Join(def.Testnet.Dir, "solver")

	const (
		privKeyFile = "privatekey"
		configFile  = "solver.toml"
	)

	if err := os.MkdirAll(confRoot, 0o755); err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	// Save network config
	endpoints := internalEndpoints(def, "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		endpoints = ExternalEndpoints(def)
	}

	// Extend endpoints with non-manifest HL chains, passed in via rpc overrides.
	for _, chain := range solvernet.Chains(def.Testnet.Network) {
		rpc, ok := def.Cfg.RPCOverrides[chain.Name]
		if !ok {
			continue
		}

		if _, ok := endpoints[chain.Name]; !ok {
			endpoints[chain.Name] = rpc
		}
	}

	solverPrivKey, err := eoa.PrivateKey(ctx, def.Testnet.Network, eoa.RoleSolver)
	if err != nil {
		return errors.Wrap(err, "get solver key")
	}

	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, privKeyFile), solverPrivKey); err != nil {
		return errors.Wrap(err, "write private key")
	}

	solverCfg := solverapp.DefaultConfig()
	solverCfg.SolverPrivKey = privKeyFile
	solverCfg.Network = def.Testnet.Network
	solverCfg.RPCEndpoints = endpoints
	solverCfg.Tracer.Endpoint = def.Cfg.TracingEndpoint
	solverCfg.Tracer.Headers = def.Cfg.TracingHeaders
	solverCfg.CoinGeckoAPIKey = def.Cfg.CoinGeckoAPIKey

	if err := solverapp.WriteConfigTOML(solverCfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write solver config")
	}

	return nil
}

func writeMonitorConfig(ctx context.Context, def Definition, logCfg log.Config, valPrivKeys []crypto.PrivKey) error {
	confRoot := filepath.Join(def.Testnet.Dir, "monitor")

	const (
		privKeyFile        = "privatekey"
		xCallerPrivKeyFile = "xcaller_privatekey"
		flowGenPrivKeyFile = "flowgen_privatekey"
		configFile         = "monitor.toml"
	)

	if err := os.MkdirAll(confRoot, 0o755); err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	// Save network config
	endpoints := internalEndpoints(def, "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		endpoints = ExternalEndpoints(def)
	}

	// xfeemngr may need out-of-network endpoints, passed in via rpc overrides.
	// We pass these only to the xfeemngr, so that other processes don't assume
	// they are in-network.
	xfeemngrEndpoints := make(xchain.RPCEndpoints)

	// First, clone the in-network endpoints.
	for name, rpc := range endpoints {
		xfeemngrEndpoints[name] = rpc
	}

	// Then, add the out-of-network endpoints.
	for name, rpc := range def.Cfg.RPCOverrides {
		if _, ok := xfeemngrEndpoints[name]; !ok {
			xfeemngrEndpoints[name] = rpc
		}
	}

	// Extend endpoints with non-manifest HL chains, passed in via rpc overrides.
	for _, chain := range solvernet.Chains(def.Testnet.Network) {
		rpc, ok := def.Cfg.RPCOverrides[chain.Name]
		if !ok {
			continue
		}

		if _, ok := endpoints[chain.Name]; !ok {
			endpoints[chain.Name] = rpc
		}
	}

	// Save private key
	privKey, err := eoa.PrivateKey(ctx, def.Testnet.Network, eoa.RoleMonitor)
	if err != nil {
		return errors.Wrap(err, "get monitor key")
	}
	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, privKeyFile), privKey); err != nil {
		return errors.Wrap(err, "write private key")
	}

	// save xcaller key
	xCallerPrivKey, err := eoa.PrivateKey(ctx, def.Testnet.Network, eoa.RoleXCaller)
	if err != nil {
		return errors.Wrap(err, "get xcaller key")
	}
	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, xCallerPrivKeyFile), xCallerPrivKey); err != nil {
		return errors.Wrap(err, "write xcaller private key")
	}

	// Save flowgen private key
	flowGenPrivKey, err := eoa.PrivateKey(ctx, def.Testnet.Network, eoa.RoleFlowgen)
	if err != nil {
		return errors.Wrap(err, "get flowgen key")
	}
	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, flowGenPrivKeyFile), flowGenPrivKey); err != nil {
		return errors.Wrap(err, "write private key")
	}

	var validatorKeyGlob string
	for i, privKey := range valPrivKeys {
		validatorKeyGlob = "validator_*"

		pk, err := k1util.StdPrivKeyFromComet(privKey)
		if err != nil {
			return errors.Wrap(err, "convert priv key")
		}

		file := fmt.Sprintf("validator_%d", i)

		// Save private key
		if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, file), pk); err != nil {
			return errors.Wrap(err, "write private key")
		}
	}

	archiveNode, ok := def.Testnet.ArchiveNode()
	if !ok {
		return errors.New("monitor must use archive node, no archive node found")
	}

	cfg := monapp.DefaultConfig()
	cfg.PrivateKey = privKeyFile
	cfg.Network = def.Testnet.Network
	cfg.HaloCometURL = archiveNode.AddressRPC()
	cfg.SolverAddress = def.Testnet.SolverInternalAddr
	cfg.HaloGRPCURL = haloGRPCAddress(archiveNode)
	cfg.RouteScanAPIKey = def.Cfg.RouteScanAPIKey
	cfg.LoadGen.ValidatorKeysGlob = validatorKeyGlob
	cfg.LoadGen.XCallerKey = xCallerPrivKeyFile
	cfg.FlowGenKey = flowGenPrivKeyFile
	cfg.RPCEndpoints = endpoints
	cfg.XFeeMngr.RPCEndpoints = xfeemngrEndpoints
	cfg.XFeeMngr.CoinGeckoAPIKey = def.Cfg.CoinGeckoAPIKey

	if err := monapp.WriteConfigTOML(cfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write monitor config")
	}

	return nil
}

// logConfig returns a default e2e log config.
// Default format is console (local dev), but vmcompose uses logfmt.
func logConfig(def Definition) log.Config {
	format := log.FormatConsole
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		format = log.FormatLogfmt
	}

	return log.Config{
		Format: format,
		Level:  slog.LevelDebug.String(),
		Color:  log.ColorForce,
	}
}

// haloGRPCAddress returns the gRPC address for a halo node
// for access internally within the network.
// Note that VM port 9090 used by grafana-agent, so gRPC bound to 9999 instead.
func haloGRPCAddress(node *e2e.Node) string {
	return fmt.Sprintf("%v:9999", node.InternalIP.String())
}

// filterFeatureFlags removes some manifest provided flags from being provided to the node.
// E.g. FlagFuzzOctane should only be applied to the validator01.
func filterFeatureFlags(node *e2e.Node, flags feature.Flags) feature.Flags {
	var resp feature.Flags
	for _, flag := range flags {
		if flag == string(feature.FlagFuzzOctane) && node.Name != "validator01" {
			continue
		}
		resp = append(resp, flag)
	}

	return resp
}
