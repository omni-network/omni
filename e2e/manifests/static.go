package manifests

import (
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/BurntSushi/toml"

	_ "embed"
)

var (
	//go:embed devnet0.toml
	devnet0 []byte

	//go:embed devnet1.toml
	devnet1 []byte

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

// Devnet1Bytes returns the devnet1.toml manifest bytes.
func Devnet1Bytes() []byte {
	return devnet1
}

// Devnet1 returns the devnet1.toml manifest.
func Devnet1() (types.Manifest, error) {
	return unmarshal(devnet1)
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
	case netconf.Devnet:
		return Devnet1()
	default:
		return types.Manifest{}, errors.New("unknown network", "network", network)
	}
}

func EVMChains(network netconf.ID) ([]evmchain.Metadata, error) {
	manifest, err := Manifest(network)
	if err != nil {
		return nil, err
	}

	return manifest.EVMChains()
}
