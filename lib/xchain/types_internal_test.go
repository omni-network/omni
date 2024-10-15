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
	att.Signatures = att.Signatures[:1]

	// mock valid signatures
	_, cometPrivKey1, addr1 := tutil.PrivateKeyFixture(t)
	_, cometPrivKey2, addr2 := tutil.PrivateKeyFixture(t)

	attRoot, err := att.AttestationRoot()
	require.NoError(t, err)

	sig1, err := k1util.Sign(cometPrivKey1, attRoot)
	require.NoError(t, err)
	sig2, err := k1util.Sign(cometPrivKey2, attRoot)
	require.NoError(t, err)

	// valid signatures
	att.Signatures = []SigTuple{
		{ValidatorAddress: addr1, Signature: sig1},
		{ValidatorAddress: addr2, Signature: sig2},
	}
	err = att.Verify()
	require.NoError(t, err)

	// force invalid signature
	att.Signatures[0].Signature = sig2
	err = att.Verify()
	require.Error(t, err)
	require.Equal(t, "invalid attestation signature", err.Error())

	// force duplicate validator
	att.Signatures[0].Signature = sig1
	att.Signatures[1].ValidatorAddress = addr1
	err = att.Verify()
	require.Error(t, err)
	require.Equal(t, "duplicate validator signature", err.Error())

	// force duplicate signature
	att.Signatures[1].ValidatorAddress = addr2
	att.Signatures[1].Signature = sig1
	err = att.Verify()
	require.Error(t, err)
	require.Equal(t, "duplicate attestation signature", err.Error())

	// empty signatures
	att.Signatures = []SigTuple{}
	err = att.Verify()
	require.Error(t, err)
	require.Equal(t, "empty attestation signatures", err.Error())
}
