package key

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

// UploadConfig is the configuration for uploading a key.
type UploadConfig struct {
	Network netconf.ID
	Name    string
	Type    Type
}

// UploadNew generates a new key and uploads it to the gcp secret manager.
func UploadNew(ctx context.Context, cfg UploadConfig) (Key, error) {
	if err := cfg.Network.Verify(); err != nil {
		return Key{}, err
	}
	if err := cfg.Type.Verify(); err != nil {
		return Key{}, err
	}

	k := Generate(cfg.Type)

	addr, err := k.Addr()
	if err != nil {
		return Key{}, err
	}

	// TODO(corver): Delete or overwrite existing key with matching labels.

	secret := secretName(cfg.Network, cfg.Name, cfg.Type, addr)
	if err := createGCPSecret(ctx, secret, k.Bytes(), nil); err != nil {
		return Key{}, errors.Wrap(err, "upload key")
	}

	log.Info(ctx, "🔐 Key uploaded: "+secret, "address", addr, "type", cfg.Type, "network", cfg.Network, "secret", cfg.Name)

	return k, nil
}

// Download retrieves a key from the gcp secret manager.
func Download(ctx context.Context, network netconf.ID, name string, typ Type, addr string) (Key, error) {
	bz, err := getGCPSecret(ctx, secretName(network, name, typ, addr))
	if err != nil {
		return Key{}, errors.Wrap(err, "download key", "network", network, "name", name, "type", typ, "addr", addr)
	}

	k, err := FromBytes(typ, bz)
	if err != nil {
		return Key{}, err
	}

	actualAddr, err := k.Addr()
	if err != nil {
		return Key{}, err
	} else if actualAddr != addr {
		return Key{}, errors.New("unexpected key address")
	}

	return k, nil
}

// secretName returns the name of the secret in the gcp secret manager.
func secretName(network netconf.ID, name string, typ Type, addr string) string {
	return network.String() + "-" + name + "-" + typ.String() + "-" + addr
}
