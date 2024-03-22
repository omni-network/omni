package fireblocks

import (
	"crypto/ecdsa"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

// TransactionResponse returns a transaction response for testing purposes.
func TransactionResponseForT(t *testing.T, id string, sig [65]byte, pubkey *ecdsa.PublicKey) transaction {
	t.Helper()

	return transaction{
		ID:     id,
		Status: "COMPLETED",
		SignedMessages: []signedMessage{{
			PublicKey: hex.EncodeToString(crypto.CompressPubkey(pubkey)),
			Signature: signature{
				FullSig: hex.EncodeToString(sig[:64]),
				R:       hex.EncodeToString(sig[:32]),
				S:       hex.EncodeToString(sig[32:64]),
				V:       int(sig[64]),
			},
		}},
	}
}
