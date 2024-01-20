package app

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/omni-network/omni/lib/errors"

	cmtos "github.com/cometbft/cometbft/libs/os"

	_ "embed"
)

const (
	configFile      = "halo.toml"
	dataDir         = "data"
	configDir       = "config"
	snapshotDataDir = "snapshots"
	networkFile     = "network.json"
	attestStateFile = "xattestations_state.json"

	// defaults.

	DefaultHomeDir                 = "./halo" // Defaults to "halo" in current directory
	defaultAppStatePersistInterval = 1        // Persist app state every block. Set to 0 to disable persistence.
	defaultSnapshotInterval        = 1000     // Roughly once an hour (given 3s blocks)
)

// DefaultHaloConfig returns the default halo config.
func DefaultHaloConfig() HaloConfig {
	return HaloConfig{
		HomeDir:                 DefaultHomeDir,
		EngineJWTFile:           "", // No default
		AppStatePersistInterval: defaultAppStatePersistInterval,
		SnapshotInterval:        defaultSnapshotInterval,
	}
}

// HaloConfig defines all halo specific config.
type HaloConfig struct {
	HomeDir                 string
	EngineJWTFile           string
	AppStatePersistInterval uint64
	SnapshotInterval        uint64
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
func WriteConfigTOML(cfg HaloConfig) error {
	var buffer bytes.Buffer

	t, err := template.New("").Parse(string(tomlTemplate))
	if err != nil {
		return errors.Wrap(err, "parse template")
	}

	if err := t.Execute(&buffer, cfg); err != nil {
		panic(err)
	}

	if err := cmtos.WriteFile(cfg.ConfigFile(), buffer.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}
