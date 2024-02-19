package attester_test

import (
	"testing"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestCreateVerifyAttestation(t *testing.T) {
	t.Parallel()

	privKey := k1.GenPrivKey()
	addr, err := k1util.PubKeyToAddress(privKey.PubKey())
	require.NoError(t, err)

	var block xchain.Block
	fuzz.New().NilChance(0).NumElements(1, 64).Fuzz(&block)

	att, err := attest.CreateAttestation(privKey, block)
	require.NoError(t, err)
	require.Equal(t, block.BlockHeader, att.BlockHeader)
	require.Equal(t, addr, att.Signature.ValidatorAddress)

	// Verify the attestation
	err = attest.VerifyAttestation(att)
	require.NoError(t, err)
}
