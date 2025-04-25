package app

import (
	"context"
	"log/slog"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/tutil"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/stretchr/testify/require"
)

const (
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
		getStatus    solvernet.OrderStatus
		rejectReason stypes.RejectReason
		expect       string
		err          string
	}{
		{
			name:         "reject",
			event:        solvernet.TopicOpened,
			getStatus:    solvernet.StatusPending,
			rejectReason: 1,
			expect:       reject,
		},
		{
			name:      "fulfill",
			event:     solvernet.TopicOpened,
			getStatus: solvernet.StatusPending,
			expect:    fill,
		},
		{
			name:      "claim",
			event:     solvernet.TopicFilled,
			getStatus: solvernet.StatusFilled,
			expect:    claim,
		},
		{
			name:      "ignore rejected",
			event:     solvernet.TopicRejected,
			getStatus: solvernet.StatusRejected,
			expect:    ignored,
		},
		{
			name:      "ignore reverted",
			event:     solvernet.TopicClosed,
			getStatus: solvernet.StatusClosed,
			expect:    ignored,
		},
		{
			name:      "ignore claimed",
			event:     solvernet.TopicClaimed,
			getStatus: solvernet.StatusClaimed,
			expect:    ignored,
		},
		{
			name:      "invalid transition",
			event:     solvernet.TopicFilled,
			getStatus: solvernet.StatusPending,
			err:       "invalid order transition",
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
				ParseID: func(log types.Log) (OrderID, error) {
					return OrderID(log.Topics[1]), nil // Return second topic as order ID
				},
				GetOrder: func(ctx context.Context, chainID uint64, id OrderID) (Order, bool, error) {
					return Order{
						ID:     id,
						Status: test.getStatus,
						pendingData: PendingData{
							FillInstruction: bindings.IERC7683FillInstruction{
								DestinationSettler: [32]byte{},
								DestinationChainId: 0,
								OriginData:         []byte{},
							},
							MaxSpent: []bindings.IERC7683Output{},
						},
					}, true, nil
				},
				DidFill: func(ctx context.Context, order Order) (bool, error) {
					return false, nil
				},
				ShouldReject: func(ctx context.Context, order Order) (stypes.RejectReason, bool, error) {
					return test.rejectReason, test.rejectReason != stypes.RejectNone, nil
				},
				Reject: func(ctx context.Context, order Order, reason stypes.RejectReason) error {
					actual = reject
					require.Equal(t, test.getStatus, order.Status)
					require.Equal(t, test.rejectReason, reason)
					require.EqualValues(t, orderID, order.ID)

					return nil
				},
				Fill: func(ctx context.Context, order Order) error {
					actual = fill
					require.Equal(t, test.getStatus, order.Status)
					require.EqualValues(t, orderID, order.ID)

					return nil
				},
				Claim: func(ctx context.Context, order Order) error {
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
				ChainName:         func(uint64) string { return "" },
				TargetName:        func(PendingData) string { return "" },
				DebugPendingOrder: func(context.Context, Order, types.Log) {},
				InstrumentAge:     func(context.Context, uint64, uint64, Order) slog.Attr { return slog.Attr{} },
			}

			proc := newEventProcFunc(deps, chainID)

			err := proc(t.Context(), types.Log{Topics: []common.Hash{test.event, orderID}})
			if test.err != "" {
				require.ErrorContains(t, err, test.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expect, actual)
		})
	}
}
