package relayer

import (
	"reflect"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/stretchr/testify/require"
)

func Test_translateSubmission(t *testing.T) {
	t.Parallel()
	type args struct {
		submission xchain.Submission
	}
	var sub xchain.Submission
	fuzz.NewWithSeed(1).NilChance(0).Fuzz(&sub)
	var xSub bindings.XChainSubmission
	fuzz.NewWithSeed(1).NilChance(0).Fuzz(&xSub)

	tests := []struct {
		name string
		args args
		want bindings.XChainSubmission
	}{
		{
			name: "",
			args: args{
				submission: sub,
			},
			want: xSub,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := translateSubmission(tt.args.submission)
			require.Equal(t, got.Proof, tt.want.Proof)
			require.Equal(t, got.ProofFlags, tt.want.ProofFlags)
			require.Equal(t, got.AttestationRoot, tt.want.AttestationRoot)
			require.True(t, reflect.DeepEqual(got.BlockHeader, tt.want.BlockHeader))
			require.True(t, reflect.DeepEqual(got.Msgs, tt.want.Msgs))
		})
	}
}
