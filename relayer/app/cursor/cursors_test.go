package cursor_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/relayer/app/cursor"

	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

// Test_StreamConfirmation tests how cursors are being confirmed.
func Test_StreamConfirmation(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream := xchain.StreamID{
		SourceChainID: 1,
		DestChainID:   2,
		ShardID:       0,
	}

	c := func(offset uint64, lastMsg uint64, confirmed bool) *cursor.Cursor {
		return buildCursor(offset, lastMsg, confirmed, stream)
	}

	tests := []struct {
		desc       string
		cursors    []*cursor.Cursor     // db cursors
		final      *xchain.SubmitCursor // on-chain submitted cursor result
		confOffset uint64               // offset up to which cursors should be confirmed
	}{{
		desc:       "single non-empty attestation",
		cursors:    []*cursor.Cursor{c(1, 2, false)},
		final:      &xchain.SubmitCursor{StreamID: stream, MsgOffset: 2, AttestOffset: 1},
		confOffset: 1,
	}, {
		desc: "non-empty attestation followed by empty",
		cursors: []*cursor.Cursor{
			c(1, 2, false),
			c(2, 0, false),
			c(3, 0, false),
		},
		final:      &xchain.SubmitCursor{StreamID: stream, MsgOffset: 2, AttestOffset: 1},
		confOffset: 3,
	}, {
		desc: "non-empty attestation followed by empty, followed by non-empty but non-confirmed, followed by empty",
		cursors: []*cursor.Cursor{
			c(1, 2, false),
			c(2, 0, false),
			c(3, 0, false),
			c(4, 4, false),
			c(5, 0, false),
			c(6, 0, false),
		},
		final:      &xchain.SubmitCursor{StreamID: stream, MsgOffset: 2, AttestOffset: 1}, // only first non-empty confirmed
		confOffset: 3,
	}, {
		desc: "empty confirmed followed by empty non-confirmed",
		cursors: []*cursor.Cursor{
			c(1, 0, true),
			c(2, 0, false),
			c(3, 0, false),
		},
		confOffset: 3,
	}, {
		desc: "empty confirmed followed by empty non-confirmed, followed by non-empty",
		cursors: []*cursor.Cursor{
			c(1, 0, true),
			c(2, 0, false),
			c(3, 0, false),
			c(4, 5, false),
		},
		final:      &xchain.SubmitCursor{StreamID: stream, MsgOffset: 3, AttestOffset: 3},
		confOffset: 3,
	}, {
		desc: "all empty and non-confirmed",
		cursors: []*cursor.Cursor{
			c(1, 0, false),
			c(2, 0, false),
			c(3, 0, false),
		},
		confOffset: 0,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			provider := &mockProvider{}
			cursorsDB, err := cursor.NewCursorsTable(db.NewMemDB())
			require.NoError(t, err)
			cursors := cursor.NewStreamCursors(stream, cursorsDB, provider)

			for _, c := range test.cursors {
				require.NoError(t, cursorsDB.Insert(ctx, c))
			}

			var getSubmittedCalled int
			shouldCallGetSubmitted := test.final != nil

			if shouldCallGetSubmitted {
				provider.getSubmittedCursorFunc = func(
					ctx context.Context,
					ref xchain.Ref,
					s xchain.StreamID,
				) (xchain.SubmitCursor, bool, error) {
					require.Equal(t, xchain.FinalizedRef, ref)
					require.Equal(t, stream, s)
					getSubmittedCalled++

					return *test.final, true, nil
				}
			}

			err = cursors.Confirm(ctx)
			require.NoError(t, err)
			// make sure we only call it once to optimize resource use,
			// and only if we do have non-empty attestation
			require.Equal(t, shouldCallGetSubmitted, getSubmittedCalled > 0)

			for _, c := range test.cursors {
				confirmed, err := cursorsDB.Get(ctx, c.GetSrcChainId(), c.GetConfLevel(), c.GetDstChainId(), c.GetAttestOffset())
				require.NoError(t, err)
				// make sure all cursors before confirmed index are indeed confirmed
				shouldBeConfirmed := c.GetAttestOffset() <= test.confOffset
				require.Equal(t, shouldBeConfirmed, confirmed.GetConfirmed(), fmt.Sprintf("offset %d", c.GetAttestOffset()))
			}
		})
	}
}

// Test_StreamTrimming tests how confirmed cursors are being removed from storage.
func Test_StreamTrimming(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream := xchain.StreamID{
		SourceChainID: 1,
		DestChainID:   2,
		ShardID:       0,
	}

	c := func(offset uint64, lastMsg uint64, confirmed bool) *cursor.Cursor {
		return buildCursor(offset, lastMsg, confirmed, stream)
	}

	tests := []struct {
		desc         string
		cursors      []*cursor.Cursor // db cursors
		deleteOffset uint64           // offset up to which the cursors were deleted
	}{{
		desc:         "single non-empty attestation",
		cursors:      []*cursor.Cursor{c(1, 2, true)},
		deleteOffset: 0,
	}, {
		desc: "multiple confirmed",
		cursors: []*cursor.Cursor{
			c(1, 2, true),
			c(2, 0, true),
			c(3, 0, true),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple confirmed, followed by non-confirmed",
		cursors: []*cursor.Cursor{
			c(1, 2, true),
			c(2, 0, true),
			c(3, 0, true),
			c(4, 0, false),
			c(5, 0, false),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple empty non-confirmed, followed by non-empty confirmed",
		cursors: []*cursor.Cursor{
			c(1, 0, false),
			c(2, 0, false),
			c(3, 2, true),
			c(5, 0, false),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple empty non-confirmed",
		cursors: []*cursor.Cursor{
			c(1, 0, false),
			c(2, 0, false),
			c(3, 0, false),
		},
		deleteOffset: 0,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			provider := &mockProvider{}
			cursorsDB, err := cursor.NewCursorsTable(db.NewMemDB())
			require.NoError(t, err)
			cursors := cursor.NewStreamCursors(stream, cursorsDB, provider)

			for _, c := range test.cursors {
				require.NoError(t, cursorsDB.Insert(ctx, c))
			}

			err = cursors.Trim(ctx)
			require.NoError(t, err)

			for _, c := range test.cursors {
				_, err := cursorsDB.Get(ctx, c.GetSrcChainId(), c.GetConfLevel(), c.GetDstChainId(), c.GetAttestOffset())
				deleted := ormerrors.IsNotFound(err)
				shouldDelete := c.GetAttestOffset() <= test.deleteOffset
				require.Equal(t, shouldDelete, deleted, fmt.Sprintf("offset %d", c.GetAttestOffset()))
			}
		})
	}
}

type mockProvider struct {
	xchain.Provider
	getSubmittedCursorFunc func(ctx context.Context, ref xchain.Ref, stream xchain.StreamID) (xchain.SubmitCursor, bool, error)
}

func (m *mockProvider) GetSubmittedCursor(
	ctx context.Context,
	ref xchain.Ref,
	stream xchain.StreamID,
) (xchain.SubmitCursor, bool, error) {
	if m.getSubmittedCursorFunc == nil {
		panic("not defined getSubmittedCursorFunc")
	}

	return m.getSubmittedCursorFunc(ctx, ref, stream)
}

func buildCursor(
	offset uint64,
	lastMsg uint64,
	confirmed bool,
	stream xchain.StreamID,
) *cursor.Cursor {
	return &cursor.Cursor{
		SrcChainId:     stream.SourceChainID,
		DstChainId:     stream.DestChainID,
		ConfLevel:      uint32(stream.ShardID),
		AttestOffset:   offset,
		Confirmed:      confirmed,
		LastXmsgOffset: lastMsg,
	}
}
