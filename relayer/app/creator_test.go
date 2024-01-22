package relayer_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/lib/xchain"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/cometbft/cometbft/crypto"
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
			e.StreamOffset = uint64(randomBetween(1, 100))
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

	ensureNoDuplicates := func(msgs []xchain.Msg) {
		msgSet := make(map[xchain.MsgID]struct{})
		for _, msg := range msgs {
			if _, exists := msgSet[msg.MsgID]; exists {
				// Fail the test if a duplicate message is found
				t.Fatalf("Duplicate message found: %+v", msg)
			}
			msgSet[msg.MsgID] = struct{}{}
		}
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
			submissions, err := relayer.CreateSubmissions(tt.streamUpdate)
			require.NoError(t, err)
			msgCount := 0
			msgs := make([]xchain.Msg, 0, len(tt.streamUpdate.Msgs))
			for _, submission := range submissions {
				require.NotNil(t, submission.AttestationRoot)
				require.Equal(t, submission.AttestationRoot, att.BlockRoot)
				require.NotNil(t, submission.Proof)
				require.NotNil(t, submission.ProofFlags)
				require.NotNil(t, submission.Signatures)
				for _, msg := range submission.Msgs {
					require.Equal(t, msg.DestChainID, submission.DestChainID)
				}
				msgCount += len(submission.Msgs)
				msgs = append(msgs, submission.Msgs...)
			}
			// ensure no msgs were dropped
			require.Len(t, msgs, len(tt.streamUpdate.Msgs))
			// check for duplicates
			ensureNoDuplicates(msgs)
		})
	}
}

func randomBetween(min, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min+1) + min
}
