package comet

import (
	crand "crypto/rand"
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestAggregate(t *testing.T) {
	t.Parallel()

	const (
		chainA = 123
		chainB = 456

		height1 = 100
		height2 = 101
	)
	valX := newPubkey()
	valY := newPubkey()
	valZ := newPubkey()

	atts := []xchain.Attestation{
		// Three atts for chainA, height1.
		newAtt(chainA, height1, valX),
		newAtt(chainA, height1, valY),
		newAtt(chainA, height1, valZ),

		// Two atts for chainB, height1.
		newAtt(chainB, height1, valX),
		newAtt(chainB, height1, valY),

		// One att for chainA, height2.
		newAtt(chainA, height2, valX),
	}

	aggs := aggregate(atts)
	require.Len(t, aggs, 3)

	// Assert that the atts are aggregated correctly.
	assertAgg(t, aggs[0], chainA, height1, valX, valY, valZ)
	assertAgg(t, aggs[1], chainB, height1, valX, valY)
	assertAgg(t, aggs[2], chainA, height2, valX)
}

// assertAgg asserts that the agg attestation has the expected values.
func assertAgg(t *testing.T, agg xchain.AggAttestation, chainID uint64, height uint64, pubKeys ...[33]byte) {
	t.Helper()
	require.Equal(t, chainID, agg.BlockHeader.SourceChainID)
	require.Equal(t, height, agg.BlockHeader.BlockHeight)
	require.Len(t, agg.Signatures, len(pubKeys))

	for i, pubKey := range pubKeys {
		require.Equal(t, pubKey, agg.Signatures[i].ValidatorPubKey)
	}
}

// newAtt returns a new attestation with deterministic random values.
func newAtt(chainID uint64, height uint64, valPubKey [33]byte) xchain.Attestation {
	r := rand.New(rand.NewSource(int64(chainID ^ height)))

	randBytes32 := func() [32]byte {
		var b [32]byte
		r.Read(b[:])

		return b
	}

	sig := func() [65]byte {
		var b [65]byte
		copy(b[:], valPubKey[:])

		return b
	}

	return xchain.Attestation{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
			BlockHash:     randBytes32(),
		},
		BlockRoot: randBytes32(),
		Signature: xchain.SigTuple{
			ValidatorPubKey: valPubKey,
			Signature:       sig(),
		},
	}
}

// newPubkey returns a random 33 byte pubkey.
func newPubkey() [33]byte {
	var pubKey [33]byte
	_, _ = crand.Read(pubKey[:])

	return pubKey
}
