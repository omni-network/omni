package relayer_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/attest/attester"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestCreatorService_CreateSubmissions(t *testing.T) {
	t.Parallel()

	var (
		SourceChainID  = uint64(randomBetween(1, 5))
		DestChainID    = uint64(randomBetween(1, 5))
		ValidatorSetID = uint64(1)
	)

	privKey := k1.GenPrivKey()
	addr, err := k1util.PubKeyToAddress(privKey.PubKey())
	require.NoError(t, err)

	fuzzer := fuzz.New().NilChance(0).Funcs(
		func(e *xchain.Msg, c fuzz.Continue) {
			e.DestChainID = DestChainID
			e.SourceChainID = SourceChainID
			e.DestAddress = common.Address(crypto.CRandBytes(20))
			e.SourceMsgSender = common.Address(crypto.CRandBytes(20))
			e.Data = crypto.CRandBytes(100)
		},
	)

	var block xchain.Block
	fuzzer.NilChance(0).NumElements(1, 64).Fuzz(&block)

	// make all msg offset sequential
	for i := range block.Msgs {
		block.Msgs[i].StreamOffset = uint64(i)
	}

	att, err := attester.CreateAttestation(privKey, block)
	require.NoError(t, err)
	require.Equal(t, block.BlockHeader, att.BlockHeader.ToXChain())
	require.Equal(t, addr, common.Address(att.Signature.ValidatorAddress))

	tree, err := xchain.NewBlockTree(block)
	require.NoError(t, err)

	aggAtt := xchain.AggAttestation{
		BlockHeader:    att.BlockHeader.ToXChain(),
		ValidatorSetID: ValidatorSetID,
		BlockRoot:      [32]byte(att.BlockRoot),
		Signatures:     []xchain.SigTuple{att.Signature.ToXChain()},
	}

	ensureNoDuplicates := func(t *testing.T, msgs []xchain.Msg) {
		t.Helper()
		msgSet := make(map[xchain.MsgID]struct{})
		for _, msg := range msgs {
			if _, exists := msgSet[msg.MsgID]; exists {
				// Fail the test if a duplicate message is found
				require.Fail(t, "Duplicate message found", msg)
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
				Tree:           tree,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			submissions, err := relayer.CreateSubmissions(tt.streamUpdate)
			require.NoError(t, err)
			msgCount := 0
			msgs := make([]xchain.Msg, 0, len(tt.streamUpdate.Msgs))
			for _, submission := range submissions {
				require.NotNil(t, submission.AttestationRoot)
				require.Equal(t, submission.AttestationRoot, common.Hash(att.BlockRoot))
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
			ensureNoDuplicates(t, msgs)
		})
	}
}

func randomBetween(min, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min+1) + min
}
