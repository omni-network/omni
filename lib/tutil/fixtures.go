package tutil

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

// PrivateKeyFixture generates an ethereum ecds private key.
func PrivateKeyFixture(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()
	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(privKey.PublicKey)

	return privKey, addr
}
