package cursor

import (
	"context"
	"slices"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	db "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

const (
	srcChainID  = uint64(1)
	destChainID = uint64(2)
)

var (
	streamID1 = xchain.StreamID{
		SourceChainID: srcChainID,
		DestChainID:   destChainID,
		ShardID:       xchain.ShardFinalized0,
	}

	streamID2 = xchain.StreamID{
		SourceChainID: srcChainID,
		DestChainID:   destChainID,
		ShardID:       xchain.ShardLatest0,
	}
)

var network = netconf.Network{Chains: []netconf.Chain{
	{ID: srcChainID, Name: "source", Shards: []xchain.ShardID{streamID1.ShardID, streamID2.ShardID}},
	{ID: destChainID, Name: "dest"},
}}

// TestStore tests how cursors are being confirmed.
func TestStore(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	finalized := make(map[xchain.StreamID]uint64)
	getSubmitCursorFunc := func(ctx context.Context, ref xchain.Ref, stream xchain.StreamID) (xchain.SubmitCursor, bool, error) {
		if ref != xchain.FinalizedRef {
			return xchain.SubmitCursor{}, false, errors.New("unexpected ref")
		}

		msgOffset, ok := finalized[stream]

		return xchain.SubmitCursor{
			StreamID:  xchain.StreamID{},
			MsgOffset: msgOffset,
		}, ok, nil
	}

	store, err := New(db.NewMemDB(), getSubmitCursorFunc, network)
	require.NoError(t, err)

	saveCursor := func(stream xchain.StreamID, attOffset uint64, msgOffset uint64) {
		var msgs []xchain.Msg
		if msgOffset != 0 {
			msgs = append(msgs, xchain.Msg{MsgID: xchain.MsgID{StreamID: stream, StreamOffset: msgOffset}})
		}
		err := store.Insert(ctx, stream.ChainVersion(), stream.DestChainID, attOffset, map[xchain.StreamID][]xchain.Msg{
			stream: msgs,
		})
		require.NoError(t, err)

		// Subsequent insert should be no-op
		err = store.Insert(ctx, stream.ChainVersion(), stream.DestChainID, attOffset, map[xchain.StreamID][]xchain.Msg{})
		require.NoError(t, err)
	}

	assert := func(t *testing.T, stream xchain.StreamID, confirmed, latest uint64) {
		t.Helper()

		confirmedOffsets, err := store.WorkerOffsets(ctx, destChainID)
		require.NoError(t, err)
		require.Equal(t, confirmed, confirmedOffsets[stream.ChainVersion()])

		lastOffsets, err := store.LatestOffsets(ctx, destChainID)
		require.NoError(t, err)
		require.Equal(t, latest, lastOffsets[stream.ChainVersion()])
	}

	assertCount := func(t *testing.T, stream1Count, stream2Count uint64) {
		t.Helper()

		counts, err := store.CountCursors(ctx)
		require.NoError(t, err)
		require.Equal(t, stream1Count, counts[streamID1.ChainVersion()])
		require.Equal(t, stream2Count, counts[streamID2.ChainVersion()])
	}

	// Add empty cursor, ensure only latest updated
	saveCursor(streamID1, 11, 0)
	assert(t, streamID1, 0, 11)
	assertCount(t, 1, 0)

	// Confirm streamID1 attOffset=11
	require.NoError(t, store.confirmOnce(ctx))
	assert(t, streamID1, 11, 11)
	assertCount(t, 1, 0)

	// Add streamID2 attOffset=21, msgOffset=201
	saveCursor(streamID2, 21, 201)
	assert(t, streamID1, 11, 11) // Unchanged
	require.NoError(t, store.confirmOnce(ctx))
	assert(t, streamID2, 0, 21) // Still unconfirmed
	assertCount(t, 1, 1)

	// Confirm streamID2 attOffset=21
	finalized[streamID2] = 202
	require.NoError(t, store.confirmOnce(ctx))
	assert(t, streamID1, 11, 11)
	assert(t, streamID2, 21, 21)
	assertCount(t, 1, 1)

	// Save new cursors for each stream, assert only latest updated
	saveCursor(streamID1, 12, 102)
	saveCursor(streamID2, 22, 202)
	assert(t, streamID1, 11, 12)
	assert(t, streamID2, 21, 22)
	assertCount(t, 2, 2)

	// Confirm streamID2 attOffset=22
	require.NoError(t, store.confirmOnce(ctx))
	assert(t, streamID1, 11, 12) // Unconfirmed
	assert(t, streamID2, 22, 22)

	// Add empty cursor on stream1, finalize previous cursor, both should be confirmed
	saveCursor(streamID1, 13, 0)
	finalized[streamID1] = 102
	require.NoError(t, store.confirmOnce(ctx))
	assert(t, streamID1, 13, 13)
	assertCount(t, 3, 2)

	require.NoError(t, store.trimOnce(ctx))
	assert(t, streamID1, 13, 13) // Unchanged
	assert(t, streamID2, 22, 22) // Unchanged
	assertCount(t, 1, 1)
}

func (s *Store) LatestOffsets(
	ctx context.Context,
	destChain uint64,
) (map[xchain.ChainVersion]uint64, error) {
	all, err := listAll(ctx, s.db)
	if err != nil {
		return nil, err
	}

	// Collect the highest cursor for each streamer.
	resp := make(map[xchain.ChainVersion]uint64)
	for s, cursors := range splitByStreamer(all) {
		if s.DstChainID != destChain {
			continue
		}

		slices.Reverse(cursors)

		for _, c := range cursors {
			resp[s.ChainVersion()] = c.GetAttestOffset()
			break
		}
	}

	return resp, nil
}

func (s *Store) CountCursors(
	ctx context.Context,
) (map[xchain.ChainVersion]uint64, error) {
	all, err := listAll(ctx, s.db)
	if err != nil {
		return nil, err
	}

	// Collect the highest cursor for each streamer.
	resp := make(map[xchain.ChainVersion]uint64)
	for s, cursors := range splitByStreamer(all) {
		resp[s.ChainVersion()] = uint64(len(cursors))
	}

	return resp, nil
}
