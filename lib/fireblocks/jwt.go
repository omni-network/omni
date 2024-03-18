package fireblocks

import (
	"crypto/sha256"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/golang-jwt/jwt/v5"
	uuid2 "github.com/google/uuid"
)

func (c Client) genJWTToken(uri string, body []byte) (string, error) {
	bodyHash := sha256.Sum256(body)
	nonce := uuid2.NewString() // this is our nonce
	// uri - The URI part of the request (e.g., /v1/transactions).
	// nonce - Unique number or string. Each API request needs to have a different nonce.
	// iat - The time at which the JWT was issued, in seconds since Epoch.
	// exp - The expiration time on and after which the JWT must not be accepted for processing, in seconds since Epoch. (Must be less than iat+30sec.)
	// sub - The API Key.
	// bodyHash - Hex-encoded SHA-256 hash of the raw HTTP request body.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uri":      uri,
			"nonce":    nonce,
			"iat":      time.Now().Unix(),
			"sub":      c.apiKey,
			"bodyHash": bodyHash,
			"exp":      time.Now().Add(time.Second * 15).Unix(),
		})

	tokenString, err := token.SignedString(c.clientSecret)
	if err != nil {
		return "", errors.Wrap(err, "jwt")
	}

	return tokenString, nil
}
