package main

import (
	"github.com/omni-network/omni/lib/errors"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/BurntSushi/toml"
)

// Manifest wraps e2e.Manifest with additional omni-specific fields.
type Manifest struct {
	e2e.Manifest

	Network string `toml:"network"`
}

// LoadManifest loads a manifest from disk.
func LoadManifest(path string) (Manifest, error) {
	manifest := Manifest{}
	_, err := toml.DecodeFile(path, &manifest)
	if err != nil {
		return manifest, errors.Wrap(err, "decode manifest")
	}

	return manifest, nil
}
