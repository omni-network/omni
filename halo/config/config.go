package config

import (
	"bytes"
	"path/filepath"
	"text/template"
	"time"

	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"

	cmtos "github.com/cometbft/cometbft/libs/os"

	pruningtypes "cosmossdk.io/store/pruning/types"
	db "github.com/cosmos/cosmos-db"

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
	defaultSnapshotInterval   = 1000     // Roughly once an hour (given 3s blocks)
	defaultSnapshotKeepRecent = 2
	defaultMinRetainBlocks    = 1 // Prune all blocks by default, Cosmsos will still respect other needs like snapshots

	defaultPruningOption      = pruningtypes.PruningOptionDefault // Note that Halo interprets this to be PruningEverything
	defaultDBBackend          = db.GoLevelDBBackend
	defaultEVMBuildDelay      = time.Millisecond * 600 // 100ms longer than geth's --miner.recommit=500ms.
	defaultEVMBuildOptimistic = true
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
	SnapshotKeepRecent uint64 // See cosmossdk.io/store/snapshots/types/options.go
	BackendType        string // See cosmos-db/db.go
	MinRetainBlocks    uint64
	PruningOption      string // See cosmossdk.io/store/pruning/types/options.go
	EVMBuildDelay      time.Duration
	EVMBuildOptimistic bool
	Tracer             tracer.Config
}

// ConfigFile returns the default path to the toml halo config file.
func (c Config) ConfigFile() string {
	return filepath.Join(c.HomeDir, configDir, configFile)
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

	t, err := template.New("").Parse(string(tomlTemplate))
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
