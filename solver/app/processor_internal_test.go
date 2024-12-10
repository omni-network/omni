package app

import (
	"context"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/stretchr/testify/require"
)

const (
	accept  = "accept"
	reject  = "reject"
	fulfill = "fulfill"
	claim   = "claim"
	ignored = ""
)

func TestEventProcessor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		event        common.Hash
		getStatus    uint8
		rejectReason rejectReason
		expect       string
	}{
		{
			name:         "accept",
			event:        topicRequested,
			getStatus:    statusPending,
			rejectReason: 0,
			expect:       accept,
		},
		{
			name:         "reject",
			event:        topicRequested,
			getStatus:    statusPending,
			rejectReason: 1,
			expect:       reject,
		},
		{
			name:      "fulfill",
			event:     topicAccepted,
			getStatus: statusAccepted,
			expect:    fulfill,
		},
		{
			name:      "claim",
			event:     topicFulfilled,
			getStatus: statusFulfilled,
			expect:    claim,
		},
		{
			name:      "ignore rejected",
			event:     topicRejected,
			getStatus: statusRejected,
			expect:    ignored,
		},
		{
			name:      "ignore reverted",
			event:     topicReverted,
			getStatus: statusReverted,
			expect:    ignored,
		},
		{
			name:      "ignore claimed",
			event:     topicClaimed,
			getStatus: statusClaimed,
			expect:    ignored,
		},
		{
			name:      "ignore mismatch 1",
			event:     topicRequested,
			getStatus: statusAccepted,
			expect:    ignored,
		},
		{
			name:      "ignore mismatch 2",
			event:     topicAccepted,
			getStatus: statusFulfilled,
			expect:    ignored,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			const chainID = 321
			const height = 123
			reqID := tutil.RandomHash()
			actual := ignored

			deps := procDeps{
				ParseID: func(_ uint64, log types.Log) ([32]byte, error) {
					return log.Topics[1], nil // Return second topic as req ID
				},
				GetRequest: func(ctx context.Context, _ uint64, id [32]byte) (bindings.SolveRequest, bool, error) {
					return bindings.SolveRequest{
						Id:     id,
						Status: test.getStatus,
					}, true, nil
				},
				ShouldReject: func(ctx context.Context, _ uint64, req bindings.SolveRequest) (rejectReason, bool, error) {
					return test.rejectReason, test.rejectReason != 0, nil
				},
				Accept: func(ctx context.Context, _ uint64, req bindings.SolveRequest) error {
					actual = accept
					require.Equal(t, test.getStatus, req.Status)
					require.EqualValues(t, reqID, req.Id)

					return nil
				},
				Reject: func(ctx context.Context, _ uint64, req bindings.SolveRequest, reason rejectReason) error {
					actual = reject
					require.Equal(t, test.getStatus, req.Status)
					require.Equal(t, test.rejectReason, reason)
					require.EqualValues(t, reqID, req.Id)

					return nil
				},
				Fulfill: func(ctx context.Context, _ uint64, req bindings.SolveRequest) error {
					actual = fulfill
					require.Equal(t, test.getStatus, req.Status)
					require.EqualValues(t, reqID, req.Id)

					return nil
				},
				Claim: func(ctx context.Context, _ uint64, req bindings.SolveRequest) error {
					actual = claim
					require.Equal(t, test.getStatus, req.Status)
					require.EqualValues(t, reqID, req.Id)

					return nil
				},
				SetCursor: func(ctx context.Context, c uint64, h uint64) error {
					require.EqualValues(t, chainID, c)
					require.EqualValues(t, height, h)

					return nil
				},
				ChainName: func(uint64) string { return "" },
			}

			processor := newEventProcessor(deps, chainID)

			err := processor(context.Background(), height, []types.Log{{Topics: []common.Hash{test.event, reqID}}})
			require.NoError(t, err)
			require.Equal(t, test.expect, actual)
		})
	}
}
