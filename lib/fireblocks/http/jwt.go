package http

import (
	"crypto/sha256"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/golang-jwt/jwt"
)

type JWTOpts struct {
	URI       string
	Nonce     string
	APIKey    string
	SecretKey string
	Body      string
}

func genJWTToken(opts JWTOpts) (string, error) {
	bodyHash := sha256.Sum256([]byte(opts.Body))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uri":      opts.URI,
			"nonce":    opts.Nonce,
			"iat":      time.Now().Unix(),
			"sub":      opts.APIKey,
			"bodyHash": bodyHash,
			// The expiration time on and after which the JWT must not be accepted for processing, in seconds since Epoch. (Must be less than iat+30sec.)
			"exp": time.Now().Add(time.Second * 15).Unix(),
		})

	tokenString, err := token.SignedString(opts.SecretKey)
	if err != nil {
		return "", errors.Wrap(err, "jwt")
	}

	return tokenString, nil
}
