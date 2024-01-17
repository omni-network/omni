package app

import (
	"path/filepath"

	"github.com/omni-network/omni/lib/netconf"
)

const (
	dataDir         = "data"
	configDir       = "config"
	networkFile     = "network.json"
	attestStateFile = "xattestations_state.json"
)

// HaloConfig defines all halo specific config.
type HaloConfig struct {
	HomeDir                 string
	EngineJWTFile           string
	AppStatePersistInterval uint64
	SnapshotInterval        uint64
}

func (c Config) Network() (netconf.Network, error) {
	return netconf.Load(filepath.Join(c.HomeDir, configDir, networkFile))
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
	return c.DataDir() // Maybe add a subdirectory for snapshots?
}
