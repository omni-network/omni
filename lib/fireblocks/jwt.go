package fireblocks

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (c Client) GenJWTToken(uri string, body []byte) (string, error) {
	nonce := uuid.New().String()
	now := time.Now().Unix()
	expiration := now + 55

	h := sha256.New()
	_, err := h.Write(body)
	if err != nil {
		return "", errors.Wrap(err, "sha256")
	}
	hashed := h.Sum(nil)

	// uri - The URI part of the request (e.g., /v1/transactions).
	// nonce - Unique number or string. Each API request needs to have a different nonce.
	// iat - The time at which the JWT was issued, in seconds since Epoch.
	// exp - The expiration time on and after which the JWT must not be accepted for processing, in seconds since Epoch. (Must be less than iat+30sec.)
	// sub - The API Key.
	// bodyHash - Hex-encoded SHA-256 hash of the raw HTTP request body.
	claims := jwt.MapClaims{
		"uri":      uri,
		"nonce":    nonce,
		"iat":      now,
		"exp":      expiration,
		"sub":      c.apiKey,
		"bodyHash": hex.EncodeToString(hashed),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", errors.Wrap(err, "jwt")
	}

	return tokenString, nil
}
