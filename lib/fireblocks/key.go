package fireblocks

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/omni-network/omni/lib/errors"
)

// LoadKey loads and returns the RSA256 from disk.
func LoadKey(path string) (*rsa.PrivateKey, error) {
	if path == "" {
		return nil, errors.New("fireblocks key path is empty")
	}

	bz, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "load fireblocks key", "path", path)
	}

	p, _ := pem.Decode(bz)
	k, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "parse fireblocks key")
	}

	resp, ok := k.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("invalid fireblocks key type", "type", fmt.Sprintf("%T", resp))
	}

	return resp, nil
}
