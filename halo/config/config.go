package config

import (
	"bytes"
	"path/filepath"
	"text/template"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cmtos "github.com/cometbft/cometbft/libs/os"

	pruningtypes "cosmossdk.io/store/pruning/types"
	db "github.com/cosmos/cosmos-db"

	_ "embed"
)

const (
	configFile      = "halo.toml"
	dataDir         = "data"
	configDir       = "config"
	snapshotDataDir = "snapshots"
	networkFile     = "network.json"
	attestStateFile = "xattestations_state.json"

	DefaultHomeDir          = "./halo" // Defaults to "halo" in current directory
	defaultSnapshotInterval = 1000     // Roughly once an hour (given 3s blocks)
	defaultMinRetainBlocks  = 0        // Retain all blocks

	defaultPruningOption      = pruningtypes.PruningOptionNothing // Prune nothing
	defaultDBBackend          = db.GoLevelDBBackend
	defaultEVMBuildDelay      = time.Millisecond * 600 // 100ms longer than geth's --miner.recommit=500ms.
	defaultEVMBuildOptimistic = true
)

// DefaultConfig returns the default halo config.
func DefaultConfig() Config {
	return Config{
		HomeDir:            DefaultHomeDir,
		EngineJWTFile:      "", // No default
		SnapshotInterval:   defaultSnapshotInterval,
		BackendType:        string(defaultDBBackend),
		MinRetainBlocks:    defaultMinRetainBlocks,
		PruningOption:      defaultPruningOption,
		EVMBuildDelay:      defaultEVMBuildDelay,
		EVMBuildOptimistic: defaultEVMBuildOptimistic,
	}
}

// Config defines all halo specific config.
type Config struct {
	HomeDir            string
	EngineJWTFile      string
	SnapshotInterval   uint64 // See cosmossdk.io/store/snapshots/types/options.go
	BackendType        string // See cosmos-db/db.go
	MinRetainBlocks    uint64
	PruningOption      string // See cosmossdk.io/store/pruning/types/options.go
	EVMBuildDelay      time.Duration
	EVMBuildOptimistic bool
}

// ConfigFile returns the default path to the toml halo config file.
func (c Config) ConfigFile() string {
	return filepath.Join(c.HomeDir, configDir, configFile)
}

func (c Config) NetworkFile() string {
	return filepath.Join(c.HomeDir, configDir, networkFile)
}

func (c Config) DataDir() string {
	return filepath.Join(c.HomeDir, dataDir)
}

func (c Config) AttestStateFile() string {
	return filepath.Join(c.DataDir(), attestStateFile)
}

func (c Config) AppStateDir() string {
	return c.DataDir() // Maybe add a subdirectory for app state?
}

func (c Config) SnapshotDir() string {
	return filepath.Join(c.DataDir(), snapshotDataDir)
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
		Log log.Config
	}{
		Config: cfg,
		Log:    logCfg,
	}

	if err := t.Execute(&buffer, s); err != nil {
		return errors.Wrap(err, "execute template")
	}

	if err := cmtos.WriteFile(cfg.ConfigFile(), buffer.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}
