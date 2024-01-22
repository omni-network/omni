package relayer_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cometbft/cometbft/crypto"
	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/lib/xchain"
	relayer "github.com/omni-network/omni/relayer/app"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestCreatorService_CreateSubmissions(t *testing.T) {
	t.Parallel()

	const (
		SourceChainID  = 1
		DestChainID    = 2
		ValidatorSetID = 1
	)

	privKey := k1.GenPrivKey()

	fuzzer := fuzz.New().NilChance(0).Funcs(
		func(e *xchain.Msg, c fuzz.Continue) {
			e.DestChainID = uint64(randomBetween(1, 5))
			e.SourceChainID = uint64(randomBetween(1, 5))
			e.StreamOffset = uint64(randomBetween(1, 10))
			e.DestAddress = [20]byte(crypto.CRandBytes(20))
			e.SourceMsgSender = [20]byte(crypto.CRandBytes(20))
			e.Data = crypto.CRandBytes(100)
		},
	)

	var block xchain.Block
	fuzzer.NilChance(0).NumElements(1, 64).Fuzz(&block)

	att, err := attest.CreateAttestation(privKey, block)
	require.NoError(t, err)
	require.Equal(t, block.BlockHeader, att.BlockHeader)
	require.Equal(t, privKey.PubKey().Bytes(), att.Signature.ValidatorPubKey[:])

	aggAtt := xchain.AggAttestation{
		BlockHeader:    att.BlockHeader,
		ValidatorSetID: ValidatorSetID,
		BlockRoot:      att.BlockRoot,
		Signatures:     []xchain.SigTuple{att.Signature},
	}

	tests := []struct {
		name         string
		streamUpdate relayer.StreamUpdate
	}{
		{
			name: "ok",
			streamUpdate: relayer.StreamUpdate{
				StreamID: xchain.StreamID{
					SourceChainID: SourceChainID,
					DestChainID:   DestChainID,
				},
				AggAttestation: aggAtt,
				Msgs:           block.Msgs,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := relayer.CreateSubmissions(tt.streamUpdate)
			require.NoError(t, err)
			for _, g := range got {
				require.NotNil(t, g.AttestationRoot)
				require.Equal(t, g.AttestationRoot, att.BlockRoot)
				require.NotNil(t, g.Proof)
				require.NotNil(t, g.ProofFlags)
				require.NotNil(t, g.Signatures)
			}
		})
	}
}

func randomBetween(min, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min+1) + min
}
