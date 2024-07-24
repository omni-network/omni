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
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	monapp "github.com/omni-network/omni/monitor/app"
	relayapp "github.com/omni-network/omni/relayer/app"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/p2p"
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
func Setup(ctx context.Context, def Definition, depCfg DeployConfig) error {
	log.Info(ctx, "Setup testnet", "dir", def.Testnet.Dir)

	if err := CleanupDir(ctx, def.Testnet.Dir); err != nil {
		return err
	}

	if err := os.MkdirAll(def.Testnet.Dir, os.ModePerm); err != nil {
		return errors.Wrap(err, "mkdir")
	}

	if def.Manifest.OnlyMonitor {
		return SetupOnlyMonitor(ctx, def)
	}

	// Setup geth execution genesis
	gethGenesis, err := evmgenutil.MakeGenesis(def.Manifest.Network)
	if err != nil {
		return errors.Wrap(err, "make genesis")
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
		def.Manifest.Network,
		time.Now(),
		gethGenesis.ToBlock().Hash(),
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

		cfg, err := MakeConfig(def.Testnet.Network, node, nodeDir)
		if err != nil {
			return err
		}
		config.WriteConfigFile(filepath.Join(nodeDir, "config", "config.toml"), cfg) // panics

		endpoints := internalEndpoints(def, node.Name)
		omniEVM := omniEVMByPrefix(def.Testnet, node.Name)

		if err := writeHaloConfig(
			def.Testnet.Network,
			nodeDir,
			def.Cfg,
			logCfg,
			depCfg.testConfig,
			node.Mode,
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
		}
		if err := halocmd.InitFiles(log.WithNoopLogger(ctx), initCfg); err != nil {
			return errors.Wrap(err, "init files")
		}
	}

	// Write an external network.json and endpoints.json in base testnet dir.
	// This allows for easy connecting or querying of the network
	endpoints := externalEndpoints(def)
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

	if err := def.Infra.Setup(); err != nil {
		return errors.Wrap(err, "setup provider")
	}

	return nil
}

func SetupOnlyMonitor(ctx context.Context, def Definition) error {
	logCfg := logConfig(def)
	if err := writeMonitorConfig(ctx, def, logCfg, nil); err != nil {
		return err
	}

	if def.Testnet.Prometheus {
		if err := agent.WriteConfig(ctx, def.Testnet, def.Cfg.AgentSecrets); err != nil {
			return errors.Wrap(err, "write prom config")
		}
	}

	if err := def.Infra.Setup(); err != nil {
		return errors.Wrap(err, "setup provider")
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
func MakeConfig(network netconf.ID, node *e2e.Node, nodeDir string) (*config.Config, error) {
	if node.ABCIProtocol != e2e.ProtocolBuiltin {
		return nil, errors.New("only Builtin ABCI is supported")
	}

	cfg := halocmd.DefaultCometConfig(nodeDir)
	cfg.Moniker = node.Name
	cfg.RPC.ListenAddress = "tcp://0.0.0.0:26657"
	cfg.RPC.PprofListenAddress = ":6060"
	cfg.P2P.ExternalAddress = fmt.Sprintf("tcp://%v:26656", advertisedIP(network, node.Mode, node.InternalIP, node.ExternalIP))
	cfg.P2P.AddrBookStrict = false
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

	if node.Mode == types.ModeSeed {
		cfg.P2P.SeedMode = true
		cfg.P2P.PexReactor = true
	}

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
		cfg.P2P.Seeds += advertisedP2PAddr(network, seed)
	}
	cfg.P2P.PersistentPeers = ""
	for _, peer := range node.PersistentPeers {
		if len(cfg.P2P.PersistentPeers) > 0 {
			cfg.P2P.PersistentPeers += ","
		}
		cfg.P2P.PersistentPeers += advertisedP2PAddr(network, peer)
	}

	if node.Prometheus {
		cfg.Instrumentation.Prometheus = true
	}

	return &cfg, nil
}

// advertisedP2PAddr returns the cometBFT network address <ID@IP:port> to advertise for a node.
func advertisedP2PAddr(network netconf.ID, node *e2e.Node) string {
	id := node.NodeKey.PubKey().Address().Bytes()
	ip := advertisedIP(network, node.Mode, node.InternalIP, node.ExternalIP)

	return fmt.Sprintf("%x@%s:26656", id, ip)
}

// advertisedIP returns the IP address to advertise for a node on a network.
func advertisedIP(network netconf.ID, mode types.Mode, internal, external net.IP) net.IP {
	if network.IsEphemeral() {
		// Ephemeral networks do not support external peers connecting to it, so we use the internal IP.
		return internal
	}

	if mode == types.ModeSeed || mode == types.ModeFull {
		// Only seeds and fullnodes allow external peers to connect to them.
		return external
	}
	// Validators and archive nodes are "secured" and only allow internal peers to connect to them.

	return internal
}

// writeHaloConfig generates an halo application config for a node and writes it to disk.
func writeHaloConfig(
	network netconf.ID,
	nodeDir string,
	defCfg DefinitionConfig,
	logCfg log.Config,
	testCfg bool,
	mode e2e.Mode,
	evmInstance string,
	endpoints xchain.RPCEndpoints,
) error {
	cfg := halocfg.DefaultConfig()

	switch mode {
	case e2e.ModeValidator, e2e.ModeFull:
		cfg.PruningOption = "nothing"
		cfg.MinRetainBlocks = 0
	case e2e.ModeSeed, e2e.ModeLight:
		cfg.PruningOption = "everything"
		cfg.MinRetainBlocks = 1
	default:
		cfg.PruningOption = "default"
		cfg.MinRetainBlocks = 0
	}

	cfg.Network = network
	cfg.HomeDir = nodeDir
	cfg.RPCEndpoints = endpoints
	cfg.EngineEndpoint = fmt.Sprintf("http://%s:8551", evmInstance) //nolint:nosprintfhostport // net.JoinHostPort doesn't prefix http.
	cfg.EngineJWTFile = "/halo/config/jwtsecret"                    // Absolute path inside docker container
	cfg.Tracer.Endpoint = defCfg.TracingEndpoint
	cfg.Tracer.Headers = defCfg.TracingHeaders

	if testCfg {
		cfg.SnapshotInterval = 1   // Write snapshots each block in e2e tests
		cfg.SnapshotKeepRecent = 0 // Keep all snapshots in e2e tests
	}

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
		endpoints = externalEndpoints(def)
	}

	// Save private key
	privKey, err := eoa.PrivateKey(ctx, def.Testnet.Network, eoa.RoleRelayer)
	if err != nil {
		return errors.Wrap(err, "get relayer key")
	}
	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, privKeyFile), privKey); err != nil {
		return errors.Wrap(err, "write private key")
	}

	relayCfg := relayapp.DefaultConfig()
	relayCfg.PrivateKey = privKeyFile
	relayCfg.Network = def.Testnet.Network
	relayCfg.HaloURL = def.Testnet.RandomHaloAddr()
	relayCfg.RPCEndpoints = endpoints

	if err := relayapp.WriteConfigTOML(relayCfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write relayer config")
	}

	return nil
}

func writeMonitorConfig(ctx context.Context, def Definition, logCfg log.Config, valPrivKeys []crypto.PrivKey) error {
	confRoot := filepath.Join(def.Testnet.Dir, "monitor")

	const (
		privKeyFile = "privatekey"
		configFile  = "monitor.toml"
	)

	if err := os.MkdirAll(confRoot, 0o755); err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	// Save network config
	endpoints := internalEndpoints(def, "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		endpoints = externalEndpoints(def)
	}

	// Save private key
	privKey, err := eoa.PrivateKey(ctx, def.Testnet.Network, eoa.RoleMonitor)
	if err != nil {
		return errors.Wrap(err, "get relayer key")
	}
	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, privKeyFile), privKey); err != nil {
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

	cfg := monapp.DefaultConfig()
	cfg.PrivateKey = privKeyFile
	cfg.Network = def.Testnet.Network
	cfg.HaloURL = def.Testnet.BroadcastNode().AddressRPC()
	cfg.LoadGen.ValidatorKeysGlob = validatorKeyGlob
	cfg.RPCEndpoints = endpoints

	if err := monapp.WriteConfigTOML(cfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write relayer config")
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
