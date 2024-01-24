package relayer

import (
	"bytes"
	"text/template"

	"github.com/omni-network/omni/lib/errors"

	cmtos "github.com/cometbft/cometbft/libs/os"

	_ "embed"
)

type Config struct {
	PrivateKey  string
	HaloURL     string
	NetworkFile string
}

func DefaultConfig() Config {
	return Config{
		PrivateKey:  "relayer.key",
		HaloURL:     "localhost:26657",
		NetworkFile: "network.json",
	}
}

//go:embed config.toml.tmpl
var tomlTemplate []byte

// WriteConfigTOML writes the toml halo config to disk.
func WriteConfigTOML(cfg Config, path string) error {
	var buffer bytes.Buffer

	t, err := template.New("").Parse(string(tomlTemplate))
	if err != nil {
		return errors.Wrap(err, "parse template")
	}

	if err := t.Execute(&buffer, cfg); err != nil {
		panic(err)
	}

	if err := cmtos.WriteFile(path, buffer.Bytes(), 0o644); err != nil {
		return errors.Wrap(err, "write config")
	}

	return nil
}
