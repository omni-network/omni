package config

import (
	"bytes"
	"text/template"

	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	cmtos "github.com/cometbft/cometbft/libs/os"

	_ "embed"
)

type Config struct {
	Network   netconf.ID
	RPCListen string // Address to listen for RPC requests, e.g. ":8080" or "localhost:8080"
	DBConn    string // Postgress connection string, e.g. "postgres://user:password@localhost:5432/dbname?sslmode=disable"
}

func DefaultConfig() Config {
	return Config{
		Network:   netconf.Devnet,
		RPCListen: ":8080",
		DBConn:    "postgres://admin:password@localhost:5432/trade?sslmode=disable",
	}
}

//go:embed config.toml.tpl
var tomlTemplate []byte

// WriteConfigTOML writes the toml monitor config to disk.
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
