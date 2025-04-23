package manifests

import (
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/BurntSushi/toml"

	_ "embed"
)

var (
	//go:embed devnet0.toml
	devnet0 []byte

	//go:embed omega.toml
	omega []byte

	//go:embed staging.toml
	staging []byte

	//go:embed mainnet.toml
	mainnet []byte
)

// Devnet0Bytes returns the devnet0.toml manifest bytes.
func Devnet0Bytes() []byte {
	return devnet0
}

// Devnet0 returns the devnet0.toml manifest.
func Devnet0() (types.Manifest, error) {
	return unmarshal(devnet0)
}

// Omega returns the omega.toml manifest.
func Omega() (types.Manifest, error) {
	return unmarshal(omega)
}

// Staging returns the staging.toml manifest.
func Staging() (types.Manifest, error) {
	return unmarshal(staging)
}

// Mainnet returns the mainnet.toml manifest.
func Mainnet() (types.Manifest, error) {
	return unmarshal(mainnet)
}

func unmarshal(b []byte) (types.Manifest, error) {
	var manifest types.Manifest
	_, err := toml.Decode(string(b), &manifest)
	if err != nil {
		return types.Manifest{}, errors.Wrap(err, "parse manifest")
	}

	return manifest, nil
}

func Manifest(network netconf.ID) (types.Manifest, error) {
	switch network {
	case netconf.Omega:
		return Omega()
	case netconf.Staging:
		return Staging()
	case netconf.Mainnet:
		return Mainnet()
	default:
		return types.Manifest{}, errors.New("devnet not supported")
	}
}
