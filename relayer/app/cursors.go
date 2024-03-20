package relayer

import (
	"context"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// getSubmittedCursors returns the last submitted cursor for each source chain on the destination chain.
// It also returns the offsets indexed by streamID for each stream.
func getSubmittedCursors(ctx context.Context, network netconf.Network, dstChainID uint64, xClient xchain.Provider,
) ([]xchain.StreamCursor, map[xchain.StreamID]uint64, error) {
	var cursors []xchain.StreamCursor                  //nolint:prealloc // Not worth it.
	initialOffsets := make(map[xchain.StreamID]uint64) // Initial submitted offsets for each stream.
	for _, srcChain := range network.Chains {
		if srcChain.ID == dstChainID {
			continue
		}

		cursor, ok, err := xClient.GetSubmittedCursor(ctx, dstChainID, srcChain.ID)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to get submitted cursors", "src_chain", srcChain.Name)
		} else if !ok {
			continue
		}

		initialOffsets[cursor.StreamID] = cursor.Offset
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

// fromHeights calculates the starting heights for streaming from source chains to a destination chain.
// It takes stream cursors, destination and source chains, and the current state, and returns
// a map where keys are source chain IDs and values are the starting heights for streaming.
func fromHeights(cursors []xchain.StreamCursor, destChain netconf.Chain, chains []netconf.Chain,
	state *State) map[uint64]uint64 {
	res := make(map[uint64]uint64)

	for _, chain := range chains {
		if chain.ID == destChain.ID {
			continue
		}
		res[chain.ID] = chain.DeployHeight
	}

	// sort cursors by decreasing SourceBlockHeight, so we start streaming from minimum height per source chain
	sort.Slice(cursors, func(i, j int) bool {
		return cursors[i].SourceBlockHeight > cursors[j].SourceBlockHeight
	})

	for _, cursor := range cursors {
		if cursor.SourceChainID == destChain.ID {
			continue // Sanity check
		}

		res[cursor.SourceChainID] = cursor.SourceBlockHeight

		// If local persisted state is higher, use that instead, skipping a bunch of empty blocks on startup.
		if height := state.GetHeight(destChain.ID, cursor.SourceChainID); height > cursor.SourceBlockHeight {
			res[cursor.SourceChainID] = height
		}
	}

	return res
}
