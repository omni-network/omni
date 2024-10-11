package xchain

import (
	"testing"

	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/tutil"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestConfLevelFuzzy(t *testing.T) {
	t.Parallel()
	var fuzzies []ConfLevel
	for conf := ConfUnknown; conf < confSentinel; conf++ {
		if conf.Valid() && conf.IsFuzzy() {
			fuzzies = append(fuzzies, conf)
		}
	}

	require.EqualValues(t, fuzzies, FuzzyConfLevels())
}

// TestAttestation_Verify ensures attestation verification works as expected.
func TestAttestation_Verify(t *testing.T) {
	t.Parallel()
	var att Attestation
	fuzz.New().NilChance(0).Fuzz(&att)
	att.AttestHeader.ChainVersion.ID = att.BlockHeader.ChainID // Align headers

	// mock valid signatures
	privKey, addr := tutil.PrivateKeyFixture(t)
	cometPrivKey, err := k1util.StdPrivKeyToComet(privKey)
	require.NoError(t, err)

	attRoot, err := att.AttestationRoot()
	require.NoError(t, err)

	sig, err := k1util.Sign(cometPrivKey, attRoot)
	require.NoError(t, err)
	for i := range att.Signatures {
		att.Signatures[i] = SigTuple{
			ValidatorAddress: addr,
			Signature:        sig,
		}
	}

	ok, err := att.Verify()
	require.True(t, ok)
	require.NoError(t, err)

	// force invalid signature
	att.Signatures[0].Signature[0] = 0
	ok, err = att.Verify()
	require.False(t, ok)
	require.NoError(t, err)
}
