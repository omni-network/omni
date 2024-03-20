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

// genJWTToken generates a JWT token for the Fireblocks API.
func (c Client) genJWTToken(uri string, data any) (string, error) {
	nonce := uuid.New().String()
	now := time.Now().Unix()
	expiration := now + 29 // Must be less than iat+30sec. (their example code uses 55 though)

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "marshaling JSON")
	}
	sha := sha256.Sum256(bytes)

	claims := jwt.MapClaims{
		"uri":      "/" + uri,                  // uri - The URI part of the request (e.g., /v1/transactions).
		"nonce":    nonce,                      // nonce - Unique number or string. Each API request needs to have a different nonce.
		"iat":      now,                        // iat - The time at which the JWT was issued, in seconds since Epoch.
		"exp":      expiration,                 // exp - The expiration time on and after which the JWT must not be accepted for processing, in seconds since Epoch. (Must be less than iat+30sec.)
		"sub":      c.apiKey,                   // sub - The API Key.
		"bodyHash": hex.EncodeToString(sha[:]), // bodyHash - Hex-encoded SHA-256 hash of the raw HTTP request body.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", errors.Wrap(err, "signing token")
	}

	return tokenString, nil
}
