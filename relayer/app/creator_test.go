package relayer

import (
	"context"
	"fmt"
	"testing"

	"github.com/omni-network/omni/lib/xchain"
	"github.com/stretchr/testify/assert"
)

func TestCreatorService_CreateSubmissions(t *testing.T) {
	type args struct {
		ctx          context.Context
		streamUpdate StreamUpdate
	}
	tests := []struct {
		name    string
		args    args
		want    []xchain.Submission
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ok",
			args: args{
				ctx: context.TODO(),
				streamUpdate: StreamUpdate{
					StreamID:       xchain.StreamID{},
					AggAttestation: xchain.AggAttestation{},
					Msgs:           nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := CreatorService{}
			got, err := cr.CreateSubmissions(tt.args.ctx, tt.args.streamUpdate)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateSubmissions(%v, %v)", tt.args.ctx, tt.args.streamUpdate)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CreateSubmissions(%v, %v)", tt.args.ctx, tt.args.streamUpdate)
		})
	}
}
