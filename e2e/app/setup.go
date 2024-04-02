package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/app/static"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/e2e/vmcompose"
	graphqlapp "github.com/omni-network/omni/explorer/graphql/app"
	indexerapp "github.com/omni-network/omni/explorer/indexer/app"
	halocmd "github.com/omni-network/omni/halo/cmd"
	halocfg "github.com/omni-network/omni/halo/config"
	"github.com/omni-network/omni/halo/genutil"
	evmgenutil "github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	monapp "github.com/omni-network/omni/monitor/app"
	relayapp "github.com/omni-network/omni/relayer/app"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	cmtos "github.com/cometbft/cometbft/libs/os"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"

	_ "embed"
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
func Setup(ctx context.Context, def Definition, agentSecrets agent.Secrets, testCfg bool, explorerDB string) error {
	log.Info(ctx, "Setup testnet", "dir", def.Testnet.Dir)

	if err := os.MkdirAll(def.Testnet.Dir, os.ModePerm); err != nil {
		return errors.Wrap(err, "mkdir")
	}

	if def.Manifest.OnlyMonitor {
		return SetupOnlyMonitor(ctx, def, agentSecrets)
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

	if err := writeOmniEVMConfig(def.Testnet); err != nil {
		return err
	}

	logCfg := logConfig()
	if err := writeMonitorConfig(def, logCfg, valPrivKeys); err != nil {
		return err
	}

	if err := writeRelayerConfig(def, logCfg); err != nil {
		return err
	}

	if err := writeAnvilState(def.Testnet); err != nil {
		return err
	}

	if err := writeExplorerIndexerConfig(def, logCfg, explorerDB); err != nil {
		return err
	}

	if err := writeExplorerGraphqlConfig(def, logCfg, explorerDB); err != nil {
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

		if err := writeHaloConfig(nodeDir, logCfg, testCfg, node.Mode); err != nil {
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

		intNetwork := internalNetwork(def.Testnet, def.Netman().DeployInfo(), node.Name)

		if err := netconf.Save(intNetwork, filepath.Join(nodeDir, NetworkConfigFile)); err != nil {
			return errors.Wrap(err, "write network config")
		}

		// Initialize the node's data directory (with noop logger since it is noisy).
		initCfg := halocmd.InitConfig{HomeDir: nodeDir, Network: def.Testnet.Network}
		if err := halocmd.InitFiles(log.WithNoopLogger(ctx), initCfg); err != nil {
			return errors.Wrap(err, "init files")
		}
	}

	if def.Testnet.Prometheus {
		if err := agent.WriteConfig(ctx, def.Testnet, agentSecrets); err != nil {
			return errors.Wrap(err, "write prom config")
		}
	}

	if err := def.Infra.Setup(); err != nil {
		return errors.Wrap(err, "setup provider")
	}

	return nil
}

func SetupOnlyMonitor(ctx context.Context, def Definition, agentSecrets agent.Secrets) error {
	logCfg := logConfig()
	if err := writeMonitorConfig(def, logCfg, nil); err != nil {
		return err
	}

	if def.Testnet.Prometheus {
		if err := agent.WriteConfig(ctx, def.Testnet, agentSecrets); err != nil {
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
func writeHaloConfig(nodeDir string, logCfg log.Config, testCfg bool, mode e2e.Mode) error {
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

var (
	//go:embed static/geth-keystore.json
	gethKeystore []byte
)

// writeOmniEVMConfig writes the omni evms (geth) config to <root>/<omni_evm>.
func writeOmniEVMConfig(testnet types.Testnet) error {
	gethGenesis, err := evmgenutil.MakeGenesis(testnet.Network)
	if err != nil {
		return errors.Wrap(err, "make genesis")
	}
	gethGenesisBz, err := json.MarshalIndent(gethGenesis, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal genesis")
	}

	gethConfigFiles := func(evm types.OmniEVM) map[string][]byte {
		return map[string][]byte{
			"genesis.json":      gethGenesisBz,
			"keystore/keystore": gethKeystore, // TODO(corver): Remove this, it isn't used.
			"geth_password.txt": []byte(""),   // Empty password
			"geth/nodekey":      ethcrypto.FromECDSA(evm.NodeKey),
			"geth/jwtsecret":    []byte(evm.JWTSecret),
		}
	}

	for _, evm := range testnet.OmniEVMs {
		for file, data := range gethConfigFiles(evm) {
			path := filepath.Join(testnet.Dir, evm.InstanceName, file)
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				return errors.Wrap(err, "mkdir", "path", path)
			}
			if err := os.WriteFile(path, data, 0o644); err != nil {
				return errors.Wrap(err, "write geth config")
			}
		}

		cfg := GethConfig{
			peers:     evm.Peers,
			IsArchive: evm.IsArchive,
			ChainID:   evm.Chain.ID,
		}
		if err := WriteGethConfigTOML(cfg, filepath.Join(testnet.Dir, evm.InstanceName, "config.toml")); err != nil {
			return errors.Wrap(err, "write geth config")
		}
	}

	return nil
}

func writeRelayerConfig(def Definition, logCfg log.Config) error {
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
	network := internalNetwork(def.Testnet, def.Netman().DeployInfo(), "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		network = externalNetwork(def.Testnet, def.Netman().DeployInfo())
	}

	if err := netconf.Save(network, filepath.Join(confRoot, networkFile)); err != nil {
		return errors.Wrap(err, "save network config")
	}

	// Save private key
	if err := ethcrypto.SaveECDSA(filepath.Join(confRoot, privKeyFile), def.Netman().RelayerKey()); err != nil {
		return errors.Wrap(err, "write private key")
	}

	ralayCfg := relayapp.DefaultConfig()
	ralayCfg.PrivateKey = privKeyFile
	ralayCfg.NetworkFile = networkFile

	ralayCfg.HaloURL = random(def.Testnet.Nodes).AddressRPC()

	if err := relayapp.WriteConfigTOML(ralayCfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write relayer config")
	}

	return nil
}

func writeMonitorConfig(def Definition, logCfg log.Config, valPrivKeys []crypto.PrivKey) error {
	confRoot := filepath.Join(def.Testnet.Dir, "monitor")

	const (
		networkFile = "network.json"
		configFile  = "monitor.toml"
	)

	if err := os.MkdirAll(confRoot, 0o755); err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	// Save network config
	network := internalNetwork(def.Testnet, def.Netman().DeployInfo(), "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		network = externalNetwork(def.Testnet, def.Netman().DeployInfo())
	}

	if err := netconf.Save(network, filepath.Join(confRoot, networkFile)); err != nil {
		return errors.Wrap(err, "save network config")
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
	cfg.NetworkFile = networkFile
	cfg.LoadGen.ValidatorKeysGlob = validatorKeyGlob

	if err := monapp.WriteConfigTOML(cfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write relayer config")
	}

	return nil
}

func writeExplorerIndexerConfig(def Definition, logCfg log.Config, explorerDB string) error {
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
	network := internalNetwork(def.Testnet, def.Netman().DeployInfo(), "")
	if def.Infra.GetInfrastructureData().Provider == vmcompose.ProviderName {
		network = externalNetwork(def.Testnet, def.Netman().DeployInfo())
	}

	if err := netconf.Save(network, filepath.Join(confRoot, networkFile)); err != nil {
		return errors.Wrap(err, "save network config")
	}

	cfg := indexerapp.DefaultConfig()
	cfg.NetworkFile = networkFile
	cfg.DBUrl = explorerDB

	if err := indexerapp.WriteConfigTOML(cfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write indexer config")
	}

	return nil
}

func writeExplorerGraphqlConfig(def Definition, logCfg log.Config, explorerDB string) error {
	confRoot := filepath.Join(def.Testnet.Dir, "explorer_graphql")

	const (
		configFile = "graphql.toml"
	)

	err := os.MkdirAll(confRoot, 0o755)
	if err != nil {
		return errors.Wrap(err, "mkdir", "path", confRoot)
	}

	cfg := graphqlapp.DefaultConfig()
	cfg.DBUrl = explorerDB

	if err := graphqlapp.WriteConfigTOML(cfg, logCfg, filepath.Join(confRoot, configFile)); err != nil {
		return errors.Wrap(err, "write graphql config")
	}

	return nil
}

// logConfig returns a default e2e log config.
func logConfig() log.Config {
	return log.Config{
		Format: log.FormatConsole,
		Level:  slog.LevelDebug.String(),
		Color:  log.ColorForce,
	}
}

// GethConfig defines the geth/config.toml config file template variables.
type GethConfig struct {
	peers     []*enode.Node
	ChainID   uint64
	IsArchive bool
}

// TrustedNodes defines nodes that geth will always connect to. They are excluded from maxPeers calculations.
func (c GethConfig) peerENRs() []string {
	var peers []string
	for _, peer := range c.peers {
		peers = append(peers, peer.String())
	}

	return peers
}

// TrustedNodes defines nodes that geth will always connect to. They are excluded from maxPeers calculations.
func (c GethConfig) TrustedNodes() string {
	return quotedStrArr(c.peerENRs())
}

// BootstrapNodes defines nodes that geth will connect to during bootstrapping to find other nodes in the network.
// TODO(corver): Replace this with network seed nodes.
func (c GethConfig) BootstrapNodes() string {
	return quotedStrArr(c.peerENRs())
}

//go:embed geth.toml.tmpl
var gethTomlTemplate []byte

// WriteGethConfigTOML writes the toml config to disk.
func WriteGethConfigTOML(cfg GethConfig, path string) error {
	var buffer bytes.Buffer

	t, err := template.New("").Parse(string(gethTomlTemplate))
	if err != nil {
		return errors.Wrap(err, "parse template")
	}

	if err := t.Execute(&buffer, cfg); err != nil {
		return errors.Wrap(err, "execute template")
	}

	if err := cmtos.WriteFile(path, buffer.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}

// quotedStrArr returns the string slices as quoted string array.
// e.g. ["a", "b", "c"] -> `["a","b","c"]`.
func quotedStrArr(arr []string) string {
	if len(arr) == 0 {
		return "[]"
	}

	return `["` + strings.Join(arr, `","`) + `"]`
}
