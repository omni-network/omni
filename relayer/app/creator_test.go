package relayer_test

import (
	"context"
	"testing"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	fuzz "github.com/google/gofuzz"
	"github.com/omni-network/omni/halo/attest"
	relayer "github.com/omni-network/omni/relayer/app"
	"github.com/stretchr/testify/require"

	"github.com/omni-network/omni/lib/xchain"
)

func TestCreatorService_CreateSubmissions(t *testing.T) {
	type args struct {
		ctx          context.Context
		streamUpdate relayer.StreamUpdate
	}
	const (
		SourceChainID  = 1
		DestChainID    = 2
		ValidatorSetID = 1
	)

	privKey := secp256k1.GenPrivKey()

	var block xchain.Block
	fuzz.New().NilChance(0).NumElements(1, 64).Fuzz(&block)

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
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				ctx: context.TODO(),
				streamUpdate: relayer.StreamUpdate{
					StreamID: xchain.StreamID{
						SourceChainID: SourceChainID,
						DestChainID:   DestChainID,
					},
					AggAttestation: aggAtt,
					Msgs:           block.Msgs,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := relayer.CreatorService{}
			got, err := cr.CreateSubmissions(tt.args.ctx, tt.args.streamUpdate)
			require.NoError(t, err)
			for _, g := range got {
				require.NotNil(t, g.AttestationRoot)
				// all leafs provided, there should be no proof
				require.Nil(t, g.Proof)
				require.Equal(t, len(g.Msgs), len(g.ProofFlags))
				require.Equal(t, len(g.Msgs), len(tt.args.streamUpdate.Msgs))
			}
		})
	}
}
