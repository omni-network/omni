package fireblocks

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const validFor = 29 * time.Second // Must be less than 30sec.

// token generates a JWT token for the Fireblocks API.
func (c Client) token(uri string, reqBody any) (string, error) {
	nonce := uuid.New().String()

	bz, err := json.Marshal(reqBody)
	if err != nil {
		return "", errors.Wrap(err, "marshaling JSON")
	}
	reqHash := sha256.Sum256(bz)

	claims := jwt.MapClaims{
		"uri":      uri,                             // uri - The URI part of the request (e.g., /v1/transactions?foo=bar).
		"nonce":    nonce,                           // nonce - Unique number or string. Each API request needs to have a different nonce.
		"iat":      time.Now().Unix(),               // iat - The time at which the JWT was issued, in seconds since Epoch.
		"exp":      time.Now().Add(validFor).Unix(), // exp - The expiration time on and after which the JWT must not be accepted for processing, in seconds since Epoch. (Must be less than iat+30sec.)
		"sub":      c.apiKey,                        // sub - The API Key.
		"bodyHash": hex.EncodeToString(reqHash[:]),  // bodyHash - Hex-encoded SHA-256 hash of the raw HTTP request body.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", errors.Wrap(err, "signing token")
	}

	return tokenString, nil
}
