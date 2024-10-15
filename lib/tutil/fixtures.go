package tutil

import (
	"crypto/ecdsa"
	"testing"

	"github.com/omni-network/omni/lib/k1util"

	cmtcrypto "github.com/cometbft/cometbft/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

// PrivateKeyFixture generates an ethereum ecds private key.
func PrivateKeyFixture(t *testing.T) (*ecdsa.PrivateKey, cmtcrypto.PrivKey, common.Address) {
	t.Helper()
	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(privKey.PublicKey)

	cometPrivKey, err := k1util.StdPrivKeyToComet(privKey)
	require.NoError(t, err)

	return privKey, cometPrivKey, addr
}
