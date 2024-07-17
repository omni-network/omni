package monitor

import (
	"bytes"
	"text/template"

	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/monitor/loadgen"

	cmtos "github.com/cometbft/cometbft/libs/os"

	_ "embed"
)

type Config struct {
	RPCEndpoints   xchain.RPCEndpoints
	Network        netconf.ID
	MonitoringAddr string
	PrivateKey     string
	HaloURL        string
	LoadGen        loadgen.Config
	DBDir          string
}

func DefaultConfig() Config {
	return Config{
		PrivateKey:     "monitor.key",
		MonitoringAddr: ":26660",
		DBDir:          "./db",
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
