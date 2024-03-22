package ethclient

import (
	"bytes"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/golang-jwt/jwt/v5"
)

type jwtRoundTripper struct {
	underlyingTransport http.RoundTripper
	jwtSecret           []byte
}

func newJWTRoundTripper(transport http.RoundTripper, jwtSecret []byte) *jwtRoundTripper {
	return &jwtRoundTripper{
		underlyingTransport: transport,
		jwtSecret:           jwtSecret,
	}
}

func (t *jwtRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(t.jwtSecret)
	if err != nil {
		return nil, errors.Wrap(err, "jwt token string")
	}

	req.Header.Set("Authorization", "Bearer "+tokenString)

	resp, err := t.underlyingTransport.RoundTrip(req)
	if err != nil {
		return nil, errors.Wrap(err, "round trip")
	}

	return resp, nil
}

// LoadJWTHexFile loads a hex encoded JWT secret from the provided file.
func LoadJWTHexFile(file string) ([]byte, error) {
	jwtHex, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "read jwt file")
	}

	jwtHex = bytes.TrimSpace(jwtHex)
	jwtHex = bytes.TrimPrefix(jwtHex, []byte("0x"))

	jwtBytes, err := hex.DecodeString(string(jwtHex))
	if err != nil {
		return nil, errors.Wrap(err, "decode jwt file")
	}

	return jwtBytes, nil
}
