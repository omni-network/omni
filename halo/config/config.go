package config

import (
	"bytes"
	"path/filepath"
	"text/template"
	"time"

	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tracer"

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
	voterStateFile  = "voter_state.json"
	keystoreGlob    = "*.ecdsa.key.json" // From https://github.com/Layr-Labs/eigenlayer-cli/blob/master/pkg/operator/keys/create.go#L165

	DefaultHomeDir            = "./halo" // Defaults to "halo" in current directory
	defaultSnapshotInterval   = 1000     // Roughly once an hour (given 3s blocks)
	defaultSnapshotKeepRecent = 2
	defaultMinRetainBlocks    = 0 // Retain all blocks

	defaultPruningOption      = pruningtypes.PruningOptionNothing // Prune nothing
	defaultDBBackend          = db.GoLevelDBBackend
	defaultEVMBuildDelay      = time.Millisecond * 600 // 100ms longer than geth's --miner.recommit=500ms.
	defaultEVMBuildOptimistic = true
)

// DefaultConfig returns the default halo config.
func DefaultConfig() Config {
	return Config{
		HomeDir:            DefaultHomeDir,
		EigenKeyPassword:   "", // No default
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
	EigenKeyPassword   string
	EngineJWTFile      string
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

func (c Config) NetworkFile() string {
	return filepath.Join(c.HomeDir, configDir, networkFile)
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

// KeystoreGlob returns the glob pattern for the eigenlayer-format ethereum keystore.
func (c Config) KeystoreGlob() string {
	return filepath.Join(c.HomeDir, configDir, keystoreGlob)
}

// KeystoreFile returns the path to the eigenlayer-format ethereum keystore file and true if it exists.
// It returns false if the file does not exist.
// It returns an error if multiple files are found.
func (c Config) KeystoreFile() (string, bool, error) {
	return statGlobSingle(c.KeystoreGlob())
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

// statGlobSingle returns the single file path for the given glob pattern and true if it exists.
// It returns false if no matching files are found.
// It returns an error if multiple files are found.
func statGlobSingle(pattern string) (string, bool, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return "", false, errors.Wrap(err, "bad glob pattern", "pattern", pattern)
	}

	if len(matches) == 0 {
		return "", false, nil
	} else if len(matches) > 1 {
		return "", true, errors.New("multiple files found for glob pattern", "pattern", pattern, "matches", matches)
	}

	return matches[0], true, nil
}
