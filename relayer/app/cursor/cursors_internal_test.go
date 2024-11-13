package cursor

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var streamID = xchain.StreamID{
	SourceChainID: 1,
	DestChainID:   2,
	ShardID:       xchain.ShardFinalized0,
}

var network = netconf.Network{Chains: []netconf.Chain{
	{ID: streamID.SourceChainID, Name: "source", Shards: []xchain.ShardID{streamID.ShardID}},
	{ID: streamID.DestChainID, Name: "mock_l1"},
}}

// Test_StreamConfirmation tests how cursors are being confirmed.
func Test_StreamConfirmation(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream := stream{
		SrcVersion: streamID.ChainVersion(),
		DestChain:  netconf.Chain{ID: streamID.DestChainID},
		Network:    network,
	}

	c := func(offset uint64, empty bool, confirmed bool) *Cursor {
		return buildCursor(offset, empty, confirmed, streamID)
	}

	tests := []struct {
		desc       string
		cursors    []*Cursor            // db cursors
		final      *xchain.SubmitCursor // on-chain submitted cursor result
		confOffset uint64               // offset up to which cursors should be confirmed
	}{{
		desc:       "single non-empty attestation",
		cursors:    []*Cursor{c(1, false, false)},
		final:      &xchain.SubmitCursor{StreamID: streamID, MsgOffset: 2, AttestOffset: 1},
		confOffset: 1,
	}, {
		desc: "non-empty attestation followed by empty",
		cursors: []*Cursor{
			c(1, false, false),
			c(2, true, false),
			c(3, true, false),
		},
		final:      &xchain.SubmitCursor{StreamID: streamID, MsgOffset: 2, AttestOffset: 1},
		confOffset: 3,
	}, {
		desc: "non-empty attestation followed by empty, followed by non-empty but non-confirmed, followed by empty",
		cursors: []*Cursor{
			c(1, false, false),
			c(2, true, false),
			c(3, true, false),
			c(4, false, false),
			c(5, true, false),
			c(6, true, false),
		},
		final:      &xchain.SubmitCursor{StreamID: streamID, MsgOffset: 2, AttestOffset: 1}, // only first non-empty confirmed
		confOffset: 3,
	}, {
		desc: "empty confirmed followed by empty non-confirmed",
		cursors: []*Cursor{
			c(1, true, true),
			c(2, true, false),
			c(3, true, false),
		},
		confOffset: 3,
	}, {
		desc: "empty confirmed followed by empty non-confirmed, followed by non-empty",
		cursors: []*Cursor{
			c(1, true, true),
			c(2, true, false),
			c(3, true, false),
			c(4, false, false),
		},
		final:      &xchain.SubmitCursor{StreamID: streamID, MsgOffset: 3, AttestOffset: 3},
		confOffset: 3,
	}, {
		desc: "all empty and non-confirmed",
		cursors: []*Cursor{
			c(1, true, false),
			c(2, true, false),
			c(3, true, false),
		},
		confOffset: 0,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			provider := &mockProvider{}
			cursorsDB, err := NewCursorsTable(db.NewMemDB())
			require.NoError(t, err)

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
					require.Equal(t, streamID, s)
					getSubmittedCalled++

					return *test.final, true, nil
				}
			}

			err = ConfirmStream(ctx, cursorsDB, network, provider, stream)
			require.NoError(t, err)
			// make sure we only call it once to optimize resource use,
			// and only if we do have non-empty attestation
			require.Equal(t, shouldCallGetSubmitted, getSubmittedCalled > 0)

			for _, c := range test.cursors {
				confirmed, err := cursorsDB.Get(ctx, c.GetSrcChainId(), c.GetConfLevel(), c.GetDstChainId(), c.GetAttestOffset())
				require.NoError(t, err)
				// make sure all cursors before confirmed index are indeed confirmed
				shouldBeConfirmed := c.GetAttestOffset() <= test.confOffset
				assert.Equal(t, shouldBeConfirmed, confirmed.GetConfirmed(), "offset", c.GetAttestOffset())
			}
		})
	}
}

// Test_StreamTrimming tests how confirmed cursors are being removed from storage.
func Test_StreamTrimming(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream := stream{
		SrcVersion: streamID.ChainVersion(),
		DestChain:  netconf.Chain{ID: streamID.DestChainID},
		Network:    network,
	}

	c := func(offset uint64, empty bool, confirmed bool) *Cursor {
		return buildCursor(offset, empty, confirmed, streamID)
	}

	tests := []struct {
		desc         string
		cursors      []*Cursor // db cursors
		deleteOffset uint64    // offset up to which the cursors were deleted
	}{{
		desc:         "single non-empty attestation",
		cursors:      []*Cursor{c(1, false, true)},
		deleteOffset: 0,
	}, {
		desc: "multiple confirmed",
		cursors: []*Cursor{
			c(1, false, true),
			c(2, true, true),
			c(3, true, true),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple confirmed, followed by non-confirmed",
		cursors: []*Cursor{
			c(1, false, true),
			c(2, true, true),
			c(3, true, true),
			c(4, true, false),
			c(5, true, false),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple empty non-confirmed, followed by non-empty confirmed",
		cursors: []*Cursor{
			c(1, true, false),
			c(2, true, false),
			c(3, false, true),
			c(5, true, false),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple empty non-confirmed",
		cursors: []*Cursor{
			c(1, true, false),
			c(2, true, false),
			c(3, true, false),
		},
		deleteOffset: 0,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			cursorsDB, err := NewCursorsTable(db.NewMemDB())
			require.NoError(t, err)

			for _, c := range test.cursors {
				require.NoError(t, cursorsDB.Insert(ctx, c))
			}

			err = TrimStream(ctx, cursorsDB, stream)
			require.NoError(t, err)

			for _, c := range test.cursors {
				_, err := cursorsDB.Get(ctx, c.GetSrcChainId(), c.GetConfLevel(), c.GetDstChainId(), c.GetAttestOffset())
				deleted := ormerrors.IsNotFound(err)
				shouldDelete := c.GetAttestOffset() <= test.deleteOffset
				require.Equal(t, shouldDelete, deleted, "offset", c.GetAttestOffset())
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
	empty bool,
	confirmed bool,
	stream xchain.StreamID,
) *Cursor {
	return &Cursor{
		SrcChainId:   stream.SourceChainID,
		DstChainId:   stream.DestChainID,
		ConfLevel:    uint32(stream.ShardID),
		AttestOffset: offset,
		Confirmed:    confirmed,
		Empty:        empty,
	}
}
