package app

import (
	"bytes"
	"html/template"

	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cmtos "github.com/cometbft/cometbft/libs/os"

	_ "embed"
)

const (
	defaultNetworkFile = "network.json"
)

type Config struct {
	NetworkFile    string
	ExplorerDBConn string
	MonitoringAddr string
}

func DefaultConfig() Config {
	return Config{
		NetworkFile:    defaultNetworkFile,
		ExplorerDBConn: "postgres://omni:password@explorer_db:5432/omni_db",
		MonitoringAddr: ":26660",
	}
}

//go:embed config.toml.tmpl
var tomlTemplate []byte

// WriteConfigTOML writes the toml indexer config to disk.
func WriteConfigTOML(cfg Config, logCfg log.Config, path string) error {
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

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, s); err != nil {
		return errors.Wrap(err, "execute template")
	}

	if err := cmtos.WriteFile(path, buffer.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}
