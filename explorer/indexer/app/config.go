package app

import (
	"bytes"
	"html/template"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cmtos "github.com/cometbft/cometbft/libs/os"

	_ "embed"
)

const (
	defaultDBURL       = "postgres://omni:password@db:5432/omni_db"
	defaultNetworkFile = "network.json"
)

type Config struct {
	NetworkFile string
	DBUrl       string
}

func DefaultConfig() Config {
	return Config{
		NetworkFile: defaultNetworkFile,
		DBUrl:       defaultDBURL,
	}
}

//go:embed config.toml.tmpl
var tomlTemplate []byte

// WriteConfigTOML writes the toml halo config to disk.
func WriteConfigTOML(cfg Config, logCfg log.Config, dir string) error {
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

	file := filepath.Join(dir, "indexer.toml")
	if err := cmtos.WriteFile(file, buffer.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}
