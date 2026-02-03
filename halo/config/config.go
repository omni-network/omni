package config

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	evmredenomsubmit "github.com/omni-network/omni/halo/evmredenom/submit"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"

	cmtos "github.com/cometbft/cometbft/libs/os"

	pruningtypes "cosmossdk.io/store/pruning/types"
	db "github.com/cosmos/cosmos-db"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"

	_ "embed"
)

const (
	configFile           = "halo.toml"
	dataDir              = "data"
	configDir            = "config"
	snapshotDataDir      = "snapshots"
	voterStateFile       = "voter_state.json"
	executionGenesisFile = "execution_genesis.json"

	DefaultHomeDir            = "./halo" // Defaults to "halo" in current directory
	defaultSnapshotInterval   = 100      // Can't be too large, must overlap with geth snapshotcache.
	defaultSnapshotKeepRecent = 2
	defaultMinRetainBlocks    = 1 // Prune all blocks by default, Cosmos will still respect other needs like snapshots

	defaultPruningOption      = pruningtypes.PruningOptionDefault // Note that Halo interprets this to be PruningEverything
	defaultDBBackend          = db.GoLevelDBBackend
	defaultEVMBuildDelay      = time.Millisecond * 600 // 100ms longer than geth's --miner.recommit=500ms.
	defaultEVMBuildOptimistic = true

	defaultAPIEnable   = true                 // Halo runs in docker, so enabled via port mapping
	defaultAPIAddress  = "tcp://0.0.0.0:1317" // Halo runs inside docker
	defaultGRPCEnable  = true                 // Halo runs in docker, so enabled via port mapping
	defaultGRPCAddress = "0.0.0.0:9090"       // Halo runs inside docker
)

// DefaultConfig returns the default halo config.
func DefaultConfig() Config {
	return Config{
		HomeDir:            DefaultHomeDir,
		Network:            "", // No default
		EngineEndpoint:     "", // No default
		EngineJWTFile:      "", // No default
		SnapshotInterval:   defaultSnapshotInterval,
		SnapshotKeepRecent: defaultSnapshotKeepRecent,
		BackendType:        string(defaultDBBackend),
		MinRetainBlocks:    defaultMinRetainBlocks,
		PruningOption:      defaultPruningOption,
		EVMBuildDelay:      defaultEVMBuildDelay,
		EVMBuildOptimistic: defaultEVMBuildOptimistic,
		Tracer:             tracer.DefaultConfig(),
		SDKAPI:             RPCConfig{Enable: defaultAPIEnable, Address: defaultAPIAddress},
		SDKGRPC:            RPCConfig{Enable: defaultGRPCEnable, Address: defaultGRPCAddress},
		FeatureFlags:       feature.Flags{}, // Zero enabled flags by default (note not nil).
	}
}

// Config defines all halo specific config.
type Config struct {
	HomeDir            string
	Network            netconf.ID
	EngineJWTFile      string
	EngineEndpoint     string
	RPCEndpoints       xchain.RPCEndpoints
	SnapshotInterval   uint64 // See cosmossdk.io/store/snapshots/types/options.go
	SnapshotKeepRecent uint32 // See cosmossdk.io/store/snapshots/types/options.go
	BackendType        string // See cosmos-db/db.go
	MinRetainBlocks    uint64
	PruningOption      string // See cosmossdk.io/store/pruning/types/options.go
	EVMBuildDelay      time.Duration
	EVMBuildOptimistic bool
	Tracer             tracer.Config
	UnsafeSkipUpgrades []int
	SDKAPI             RPCConfig `mapstructure:"api"`
	SDKGRPC            RPCConfig `mapstructure:"grpc"`
	FeatureFlags       feature.Flags
	EVMProxyListen     string                  // The address to listen for evm proxy requests on
	EVMProxyTarget     string                  // The target address to proxy evm requests to
	EVMRedenomSubmit   evmredenomsubmit.Config // EVM redenomination submit config
}

// HaltHeight returns the consensus halt height for the given network.
// Returns 0 if no halt height is configured for the network.
func HaltHeight(network netconf.ID) uint64 {
	switch network {
	case netconf.Staging:
		return 1190000
	default:
		return 0
	}
}

// RPCConfig is an abridged version of CosmosSDK srvconfig.API/GRPCConfig.
type RPCConfig struct {
	Enable  bool
	Address string
}

// ConfigFile returns the default path to the toml halo config file.
func (c Config) ConfigFile() string {
	return filepath.Join(c.HomeDir, configDir, configFile)
}

// SDKRPCConfig returns the SDK config with only RPC fields populated.
func (c Config) SDKRPCConfig() srvconfig.Config {
	cfg := srvconfig.DefaultConfig()

	api := cfg.API
	api.Enable = c.SDKAPI.Enable
	api.Address = c.SDKAPI.Address
	if c.Network.IsEphemeral() {
		// Enable CORS by default on ephemeral networks.
		// TODO(corver): Expose all config options rather.
		api.EnableUnsafeCORS = true
	}

	grpc := cfg.GRPC
	grpc.Enable = c.SDKGRPC.Enable
	grpc.Address = c.SDKGRPC.Address

	grpcweb := cfg.GRPCWeb

	return srvconfig.Config{
		API:     api,
		GRPC:    grpc,
		GRPCWeb: grpcweb,
	}
}

func (c Config) DataDir() string {
	return filepath.Join(c.HomeDir, dataDir)
}

func (c Config) VoterStateFile() string {
	return filepath.Join(c.DataDir(), voterStateFile)
}

func (c Config) AppStateDir() string {
	return c.DataDir() // Maybe add a subdirectory for app state?
}

func (c Config) SnapshotDir() string {
	return filepath.Join(c.DataDir(), snapshotDataDir)
}

func (c Config) ExecutionGenesisFile() string {
	return filepath.Join(c.HomeDir, configDir, executionGenesisFile)
}

func (c Config) Verify() error {
	if c.EngineEndpoint == "" {
		return errors.New("flag --engine-endpoint is empty")
	} else if c.EngineJWTFile == "" {
		return errors.New("flag --engine-jwt-file is empty")
	} else if c.Network == "" {
		return errors.New("flag --network is empty")
	} else if err := c.Network.Verify(); err != nil {
		return err
	}

	return nil
}

//go:embed config.toml.tmpl
var tomlTemplate []byte

// WriteConfigTOML writes the toml halo config to disk.
func WriteConfigTOML(cfg Config, logCfg log.Config) error {
	var buffer bytes.Buffer

	t, err := template.New("").
		Funcs(template.FuncMap{"FmtIntSlice": fmtSlice[int]}).
		Parse(string(tomlTemplate))
	if err != nil {
		return errors.Wrap(err, "parse template")
	}

	s := struct {
		Config
		Log     log.Config
		Version string
	}{
		Config:  cfg,
		Log:     logCfg,
		Version: buildinfo.Version(),
	}

	if err := t.Execute(&buffer, s); err != nil {
		return errors.Wrap(err, "execute template")
	}

	if err := cmtos.WriteFile(cfg.ConfigFile(), buffer.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}

func fmtSlice[T any](slice []T) string {
	var sb strings.Builder
	for i, v := range slice {
		if i > 0 {
			_, _ = sb.WriteString(",")
		}
		_, _ = sb.WriteString(fmt.Sprint(v))
	}

	return "[" + sb.String() + "]"
}
