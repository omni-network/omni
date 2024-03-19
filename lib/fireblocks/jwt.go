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

func (c Client) GenJWTToken(uri string, bodyJSON any) (string, error) {
	nonce := uuid.New().String()
	now := time.Now().Unix()
	expiration := now + 55

	bodyBytes, err := json.Marshal(bodyJSON)
	if err != nil {
		return "", errors.Wrap(err, "marshaling JSON")
	}

	h := sha256.New()
	_, err = h.Write(bodyBytes)
	if err != nil {
		return "", errors.Wrap(err, "sha256")
	}
	hashed := h.Sum(nil)

	claims := jwt.MapClaims{
		"uri":      uri,                        // uri - The URI part of the request (e.g., /v1/transactions).
		"nonce":    nonce,                      // nonce - Unique number or string. Each API request needs to have a different nonce.
		"iat":      now,                        // iat - The time at which the JWT was issued, in seconds since Epoch.
		"exp":      expiration,                 // exp - The expiration time on and after which the JWT must not be accepted for processing, in seconds since Epoch. (Must be less than iat+30sec.)
		"sub":      c.apiKey,                   // sub - The API Key.
		"bodyHash": hex.EncodeToString(hashed), // bodyHash - Hex-encoded SHA-256 hash of the raw HTTP request body.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", errors.Wrap(err, "signing token")
	}

	return tokenString, nil
}
