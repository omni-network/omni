package relayer_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/attest/voter"
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
		SourceChainID = uint64(randomBetween(1, 5))
		DestChainID   = uint64(randomBetween(1, 5))
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
	fuzzer.NilChance(0).NumElements(2, 64).Fuzz(&block)
	require.NotEmpty(t, block.Msgs)

	var attestHeader xchain.AttestHeader
	fuzzer.Fuzz(&attestHeader)
	attestHeader.ChainVersion.ID = block.ChainID // Align headers

	var valSetID uint64
	fuzzer.Fuzz(&valSetID)

	// make all msg offset sequential
	for i := range block.Msgs {
		block.Msgs[i].StreamOffset = uint64(i)
	}

	vote, err := voter.CreateVote(privKey, attestHeader, block)
	require.NoError(t, err)
	require.Equal(t, block.BlockHeader, vote.BlockHeader.ToXChain())
	require.Equal(t, addr, common.Address(vote.Signature.ValidatorAddress))

	tree, err := xchain.NewMsgTree(block.Msgs)
	require.NoError(t, err)

	att := xchain.Attestation{
		AttestHeader:   vote.AttestHeader.ToXChain(),
		BlockHeader:    vote.BlockHeader.ToXChain(),
		ValidatorSetID: valSetID,
		MsgRoot:        [32]byte(vote.MsgRoot),
		Signatures:     []xchain.SigTuple{vote.Signature.ToXChain()},
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
				Attestation: att,
				Msgs:        block.Msgs,
				MsgTree:     tree,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			subs, err := relayer.CreateSubmissions(tt.streamUpdate)
			require.NoError(t, err)
			msgCount := 0
			msgs := make([]xchain.Msg, 0, len(tt.streamUpdate.Msgs))
			for _, sub := range subs {
				require.EqualValues(t, valSetID, sub.ValidatorSetID)
				require.NotNil(t, sub.AttestationRoot)
				attRoot, err := att.AttestationRoot()
				require.NoError(t, err)
				require.Equal(t, sub.AttestationRoot.Bytes(), attRoot[:])
				require.NotEmpty(t, sub.ProofFlags)
				require.NotEmpty(t, sub.Signatures)
				for _, msg := range sub.Msgs {
					require.Equal(t, msg.DestChainID, sub.DestChainID)
				}
				msgCount += len(sub.Msgs)
				msgs = append(msgs, sub.Msgs...)
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
