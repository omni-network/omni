package app

import (
	"bytes"
	"path/filepath"
	"text/template"

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

	DefaultHomeDir                 = "./halo" // Defaults to "halo" in current directory
	defaultAppStatePersistInterval = 1        // Persist app state every block. Set to 0 to disable persistence.
	defaultSnapshotInterval        = 1000     // Roughly once an hour (given 3s blocks)
	defaultMinRetainBlocks         = 0        // Retain all blocks

	defaultPruningOption = pruningtypes.PruningOptionNothing // Prune nothing
	defaultDBBackend     = db.MemDBBackend
)

// DefaultHaloConfig returns the default halo config.
func DefaultHaloConfig() HaloConfig {
	return HaloConfig{
		HomeDir:                 DefaultHomeDir,
		EngineJWTFile:           "", // No default
		AppStatePersistInterval: defaultAppStatePersistInterval,
		SnapshotInterval:        defaultSnapshotInterval,

		// Halo2
		BackendType:     defaultDBBackend,
		MinRetainBlocks: defaultMinRetainBlocks,
		PruningOption:   defaultPruningOption,
	}
}

// HaloConfig defines all halo specific config.
type HaloConfig struct {
	HomeDir                 string
	EngineJWTFile           string
	AppStatePersistInterval uint64
	SnapshotInterval        uint64 // See cosmossdk.io/store/snapshots/types/options.go

	// Halo2
	BackendType     db.BackendType
	MinRetainBlocks uint64
	PruningOption   string // See cosmossdk.io/store/pruning/types/options.go
}

// ConfigFile returns the default path to the toml halo config file.
func (c HaloConfig) ConfigFile() string {
	return filepath.Join(c.HomeDir, configDir, configFile)
}

func (c HaloConfig) NetworkFile() string {
	return filepath.Join(c.HomeDir, configDir, networkFile)
}

func (c HaloConfig) DataDir() string {
	return filepath.Join(c.HomeDir, dataDir)
}

func (c HaloConfig) AttestStateFile() string {
	return filepath.Join(c.DataDir(), attestStateFile)
}

func (c HaloConfig) AppStateDir() string {
	return c.DataDir() // Maybe add a subdirectory for app state?
}

func (c HaloConfig) SnapshotDir() string {
	return filepath.Join(c.DataDir(), snapshotDataDir)
}

//go:embed config.toml.tmpl
var tomlTemplate []byte

// WriteConfigTOML writes the toml halo config to disk.
func WriteConfigTOML(cfg HaloConfig, logCfg log.Config) error {
	var buffer bytes.Buffer

	t, err := template.New("").Parse(string(tomlTemplate))
	if err != nil {
		return errors.Wrap(err, "parse template")
	}

	s := struct {
		HaloConfig
		Log log.Config
	}{
		HaloConfig: cfg,
		Log:        logCfg,
	}

	if err := t.Execute(&buffer, s); err != nil {
		return errors.Wrap(err, "execute template")
	}

	if err := cmtos.WriteFile(cfg.ConfigFile(), buffer.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}
