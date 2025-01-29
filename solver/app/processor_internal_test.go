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
	fill    = "fill"
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
			event:        topicOpened,
			getStatus:    statusPending,
			rejectReason: 0,
			expect:       accept,
		},
		{
			name:         "reject",
			event:        topicOpened,
			getStatus:    statusPending,
			rejectReason: 1,
			expect:       reject,
		},
		{
			name:      "fulfill",
			event:     topicAccepted,
			getStatus: statusAccepted,
			expect:    fill,
		},
		{
			name:      "claim",
			event:     topicFilled,
			getStatus: statusFilled,
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
			event:     topicOpened,
			getStatus: statusAccepted,
			expect:    ignored,
		},
		{
			name:      "ignore mismatch 2",
			event:     topicAccepted,
			getStatus: statusFilled,
			expect:    ignored,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			const chainID = 321
			const height = 123
			orderID := tutil.RandomHash()
			actual := ignored

			deps := procDeps{
				ParseID: func(_ uint64, log types.Log) (OrderID, error) {
					return OrderID(log.Topics[1]), nil // Return second topic as order ID
				},
				GetOrder: func(ctx context.Context, chainID uint64, id OrderID) (Order, bool, error) {
					return Order{
						ID:     id,
						Status: test.getStatus,
						FillInstruction: bindings.IERC7683FillInstruction{
							DestinationSettler: [32]byte{},
							DestinationChainId: 0,
							OriginData:         []byte{},
						},
						MaxSpent: []bindings.IERC7683Output{},
					}, true, nil
				},
				ShouldReject: func(ctx context.Context, _ uint64, order Order) (rejectReason, bool, error) {
					return test.rejectReason, test.rejectReason != 0, nil
				},
				Accept: func(ctx context.Context, _ uint64, order Order) error {
					actual = accept
					require.Equal(t, test.getStatus, order.Status)
					require.EqualValues(t, orderID, order.ID)

					return nil
				},
				Reject: func(ctx context.Context, _ uint64, order Order, reason rejectReason) error {
					actual = reject
					require.Equal(t, test.getStatus, order.Status)
					require.Equal(t, test.rejectReason, reason)
					require.EqualValues(t, orderID, order.ID)

					return nil
				},
				Fill: func(ctx context.Context, _ uint64, order Order) error {
					actual = fill
					require.Equal(t, test.getStatus, order.Status)
					require.EqualValues(t, orderID, order.ID)

					return nil
				},
				Claim: func(ctx context.Context, _ uint64, order Order) error {
					actual = claim
					require.Equal(t, test.getStatus, order.Status)
					require.EqualValues(t, orderID, order.ID)

					return nil
				},
				SetCursor: func(ctx context.Context, c uint64, h uint64) error {
					require.EqualValues(t, chainID, c)
					require.EqualValues(t, height, h)

					return nil
				},
				ChainName:  func(uint64) string { return "" },
				TargetName: func(Order) string { return "" },
			}

			processor := newEventProcessor(deps, chainID)

			err := processor(context.Background(), height, []types.Log{{Topics: []common.Hash{test.event, orderID}}})
			require.NoError(t, err)
			require.Equal(t, test.expect, actual)
		})
	}
}
