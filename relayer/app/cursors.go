package relayer

import (
	"context"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// initialXBlockOffset defines the first xBlockOffset for all chains, it starts at 1, not 0.
const initialXBlockOffset = 1

// getSubmittedCursors returns the last submitted cursor for each source chain on the destination chain.
// It also returns the offsets indexed by streamID for each stream.
func getSubmittedCursors(ctx context.Context, network netconf.Network, dstChainID uint64, xClient xchain.Provider,
) ([]xchain.StreamCursor, map[xchain.StreamID]uint64, error) {
	var cursors []xchain.StreamCursor                  //nolint:prealloc // Not worth it.
	initialOffsets := make(map[xchain.StreamID]uint64) // Initial submitted offsets for each stream.
	for _, stream := range network.StreamsTo(dstChainID) {
		cursor, ok, err := xClient.GetSubmittedCursor(ctx, stream)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to get submitted cursors", "src_chain", stream.SourceChainID)
		} else if !ok {
			continue
		}

		initialOffsets[cursor.StreamID] = cursor.MsgOffset
		cursors = append(cursors, cursor)
	}

	return cursors, initialOffsets, nil
}

// filterMsgs filters messages based on offsets for a specific stream.
// It takes a slice of messages, offsets indexed by stream ID, and the target stream ID,
// and returns a filtered slice containing only messages with offsets greater than the specified offset.
func filterMsgs(msgs []xchain.Msg, offsets map[xchain.StreamID]uint64, streamID xchain.StreamID) []xchain.Msg {
	offset, ok := offsets[streamID]
	if !ok {
		return msgs // No offset, so no filtering.
	}

	res := make([]xchain.Msg, 0, len(msgs)) // Res might have over-capacity, but that's fine, we only filter on startup.
	for _, msg := range msgs {
		if msg.StreamOffset <= offset {
			// filter msgs lower than offset
			continue
		}
		res = append(res, msg)
	}

	return res
}

// fromOffsets calculates the starting offsets for all streams (to the destination chain).
// It takes submitted stream cursors, destination and source chains, and the current state, and returns
// a map where keys are source chain IDs and values are the starting offsets for streaming.
func fromOffsets(
	cursors []xchain.StreamCursor, // All actual on-chain submit cursors
	streams []xchain.StreamID, // All expected streams
	state *State, // On-disk local state
) (map[xchain.StreamID]uint64, error) {
	res := make(map[xchain.StreamID]uint64)

	// Initialize all streams to start at 1 by default or if local state is present
	for _, stream := range streams {
		res[stream] = initialXBlockOffset

		// If local persisted state is higher, use that instead, skipping a bunch of empty blocks on startup.
		if offset := state.GetOffset(stream.DestChainID, stream.SourceChainID); offset > initialXBlockOffset {
			res[stream] = offset
		}
	}

	// sort cursors by decreasing offset, so we start streaming from minimum offset per source chain
	sort.Slice(cursors, func(i, j int) bool {
		return cursors[i].BlockOffset > cursors[j].BlockOffset
	})

	for _, cursor := range cursors {
		offset, ok := res[cursor.StreamID]
		if !ok {
			return nil, errors.New("unexpected cursor [BUG]")
		}

		if offset >= cursor.BlockOffset {
			continue // Skip if local state is higher than cursor
		}

		res[cursor.StreamID] = cursor.BlockOffset
	}

	return res, nil
}
