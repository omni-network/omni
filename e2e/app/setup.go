package app

import (
	"context"
	"fmt"
	"log/slog"
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
	graphqlapp "github.com/omni-network/omni/explorer/graphql/app"
	indexerapp "github.com/omni-network/omni/explorer/indexer/app"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/halo/genutil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
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

	PrivvalAddressTCP  = "tcp://0.0.0.0:27559"
	PrivvalAddressUNIX = "unix:///var/run/privval.sock"
	PrivvalKeyFile     = "config/priv_validator_key.json"
	PrivvalStateFile   = "data/priv_validator_state.json"
	NetworkConfigFile  = "config/network.json"
)

// Setup sets up the testnet configuration.
func Setup(ctx context.Context, def Definition, depCfg DeployConfig) error {
	log.Info(ctx, "Setup testnet", "dir", def.Testnet.Dir)

	if err := os.MkdirAll(def.Testnet.Dir, os.ModePerm); err != nil {
		return errors.Wrap(err, "mkdir")
	}

	if def.Manifest.OnlyMonitor {
		return SetupOnlyMonitor(ctx, def)
	}

	var vals []crypto.PubKey
	var valPrivKeys []crypto.PrivKey
	for val := range def.Testnet.Validators {
		vals = append(vals, val.PrivvalKey.PubKey())
		valPrivKeys = append(valPrivKeys, val.PrivvalKey)
	}

	cosmosGenesis, err := genutil.MakeGenesis(def.Manifest.Network, time.Now(), vals...)
	if err != nil {
		return errors.Wrap(err, "make genesis")
	}
	cmtGenesis, err := cosmosGenesis.ToGenesisDoc()
	if err != nil {
		return errors.Wrap(err, "convert genesis")
	}

	if err := geth.WriteAllConfig(def.Testnet); err != nil {
		return err
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

	if err := writeExplorerIndexerConfig(ctx, def, logCfg); err != nil {
		return err
	}

	if err := writeExplorerGraphqlConfig(def, logCfg); err != nil {
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

		cfg, err := MakeConfig(node, nodeDir)
		if err != nil {
			return err
		}
		config.WriteConfigFile(filepath.Join(nodeDir, "config", "config.toml"), cfg) // panics

		if err := writeHaloConfig(nodeDir, def.Cfg, logCfg, depCfg.testConfig, node.Mode); err != nil {
			return err
		}

		omniEVM := omniEVMByPrefix(def.Testnet, node.Name)
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

		intNetwork := internalNetwork(def, node.Name)

		if err := netconf.Save(ctx, intNetwork, filepath.Join(nodeDir, NetworkConfigFile)); err != nil {
			return errors.Wrap(err, "write network config")
		}

		// Initialize the node's data directory (with noop logger since it is noisy).
		initCfg := halocmd.InitConfig{HomeDir: nodeDir, Network: def.Testnet.Network}
		if err := halocmd.InitFiles(log.WithNoopLogger(ctx), initCfg); err != nil {
			return errors.Wrap(err, "init files")
		}
	}

	// Write an external network.json in base testnet dir.
	// This allows for easy connecting or querying of the network
	extNetwork := externalNetwork(def)
	if err := netconf.Save(ctx, extNetwork, filepath.Join(def.Testnet.Dir, "network.json")); err != nil {
		return errors.Wrap(err, "write network config")
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
//
//nolint:lll // CometBFT super long names :(
func MakeConfig(node *e2e.Node, nodeDir string) (*config.Config, error) {
	cfg := halocmd.DefaultCometConfig(nodeDir)
	cfg.Moniker = node.Name
	cfg.ProxyApp = AppAddressTCP
	cfg.RPC.ListenAddress = "tcp://0.0.0.0:26657"
	cfg.RPC.PprofListenAddress = ":6060"
	cfg.P2P.ExternalAddress = fmt.Sprintf("tcp://%v", node.AddressP2P(false))
	cfg.P2P.AddrBookStrict = false
	cfg.DBBackend = node.Database
	cfg.StateSync.DiscoveryTime = 5 * time.Second
	cfg.BlockSync.Version = node.BlockSyncVersion
	cfg.Mempool.ExperimentalMaxGossipConnectionsToNonPersistentPeers = int(node.Testnet.ExperimentalMaxGossipConnectionsToNonPersistentPeers)
	cfg.Mempool.ExperimentalMaxGossipConnectionsToPersistentPeers = int(node.Testnet.ExperimentalMaxGossipConnectionsToPersistentPeers)

	switch node.ABCIProtocol {
	case e2e.ProtocolUNIX:
		cfg.ProxyApp = AppAddressUNIX
	case e2e.ProtocolTCP:
		cfg.ProxyApp = AppAddressTCP
	case e2e.ProtocolGRPC:
		cfg.ProxyApp = AppAddressTCP
		cfg.ABCI = "grpc"
	case e2e.ProtocolBuiltin, e2e.ProtocolBuiltinConnSync:
		cfg.ProxyApp = ""
		cfg.ABCI = ""
	default:
		return nil, errors.New("unexpected ABCI protocol")
	}

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
		cfg.P2P.Seeds += seed.AddressP2P(true)
	}
	cfg.P2P.PersistentPeers = ""
	for _, peer := range node.PersistentPeers {
		if len(cfg.P2P.PersistentPeers) > 0 {
			cfg.P2P.PersistentPeers += ","
		}
		cfg.P2P.PersistentPeers += peer.AddressP2P(true)
	}

	if node.Prometheus {
		cfg.Instrumentation.Prometheus = true
	}

	return &cfg, nil
}

// writeHaloConfig generates an halo application config for a node and writes it to disk.
func writeHaloConfig(nodeDir string, defCfg DefinitionConfig, logCfg log.Config, testCfg bool, mode e2e.Mode) error {
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

	cfg.HomeDir = nodeDir
	cfg.EngineJWTFile = "/halo/config/jwtsecret" // Absolute path inside docker container
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
		networkFile = "network.json"
		configFile  = "relayer.toml"
	)

	if err := os.MkdirAll(confRoot, 0o755); err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	// Save network config
	network := internalNetwork(def, "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		network = externalNetwork(def)
	}

	if err := netconf.Save(ctx, network, filepath.Join(confRoot, networkFile)); err != nil {
		return errors.Wrap(err, "save network config")
	}

	// Save private key
	privKey, err := eoa.PrivateKey(ctx, def.Testnet.Network, eoa.RoleRelayer)
	if err != nil {
		return errors.Wrap(err, "get relayer key")
	}
	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, privKeyFile), privKey); err != nil {
		return errors.Wrap(err, "write private key")
	}

	ralayCfg := relayapp.DefaultConfig()
	ralayCfg.PrivateKey = privKeyFile
	ralayCfg.NetworkFile = networkFile
	ralayCfg.HaloURL = def.Testnet.RandomHaloAddr()

	if err := relayapp.WriteConfigTOML(ralayCfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write relayer config")
	}

	return nil
}

func writeMonitorConfig(ctx context.Context, def Definition, logCfg log.Config, valPrivKeys []crypto.PrivKey) error {
	confRoot := filepath.Join(def.Testnet.Dir, "monitor")

	const (
		privKeyFile = "privatekey"
		networkFile = "network.json"
		configFile  = "monitor.toml"
	)

	if err := os.MkdirAll(confRoot, 0o755); err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	// Save network config
	network := internalNetwork(def, "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		network = externalNetwork(def)
	}

	if err := netconf.Save(ctx, network, filepath.Join(confRoot, networkFile)); err != nil {
		return errors.Wrap(err, "save network config")
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
	cfg.NetworkFile = networkFile
	cfg.LoadGen.ValidatorKeysGlob = validatorKeyGlob

	if err := monapp.WriteConfigTOML(cfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write relayer config")
	}

	return nil
}

func writeExplorerIndexerConfig(ctx context.Context, def Definition, logCfg log.Config) error {
	confRoot := filepath.Join(def.Testnet.Dir, "explorer_indexer")

	const (
		networkFile = "network.json"
		configFile  = "indexer.toml"
	)

	err := os.MkdirAll(confRoot, 0o755)
	if err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	// Save network config
	network := internalNetwork(def, "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		network = externalNetwork(def)
	}

	if err := netconf.Save(ctx, network, filepath.Join(confRoot, networkFile)); err != nil {
		return errors.Wrap(err, "save network config")
	}

	cfg := indexerapp.DefaultConfig()
	cfg.NetworkFile = networkFile
	cfg.ExplorerDBConn = def.Cfg.ExplorerDBConn

	if err := indexerapp.WriteConfigTOML(cfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write indexer config")
	}

	return nil
}

func writeExplorerGraphqlConfig(def Definition, logCfg log.Config) error {
	confRoot := filepath.Join(def.Testnet.Dir, "explorer_graphql")

	const (
		configFile = "graphql.toml"
	)

	err := os.MkdirAll(confRoot, 0o755)
	if err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	cfg := graphqlapp.DefaultConfig()
	cfg.ExplorerDBConn = def.Cfg.ExplorerDBConn

	if err := graphqlapp.WriteConfigTOML(cfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write graphql config")
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
