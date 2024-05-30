package relayer

import (
	"context"
	"sort"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// initialXBlockOffset defines the first xBlockOffset for all chains, it starts at 1, not 0.
const initialXBlockOffset = 1

// getSubmittedCursors returns the last submitted cursor for each source chain on the destination chain.
func getSubmittedCursors(ctx context.Context, network netconf.Network, dstChainID uint64, xClient xchain.Provider,
) ([]xchain.StreamCursor, error) {
	var cursors []xchain.StreamCursor //nolint:prealloc // Not worth it.
	for _, stream := range network.StreamsTo(dstChainID) {
		cursor, ok, err := xClient.GetSubmittedCursor(ctx, stream)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get submitted cursors", "src_chain", stream.SourceChainID)
		} else if !ok {
			continue
		}

		cursors = append(cursors, cursor)
	}

	return cursors, nil
}

// filterMsgs filters messages based on offsets for a specific stream.
// It takes a slice of messages, offsets indexed by stream ID, and the target stream ID,
// and returns a filtered slice containing only messages with offsets greater than the specified offset.
func filterMsgs(ctx context.Context, streamID xchain.StreamID, msgs []xchain.Msg, msgFilter *msgOffsetFilter) ([]xchain.Msg, error) {
	backoff := expbackoff.New(ctx)
	res := make([]xchain.Msg, 0, len(msgs)) // Res might have over-capacity, but that's fine, we only filter on startup.
	for i := 0; i < len(msgs); {
		msg := msgs[i]

		check := msgFilter.Check(streamID, msg.StreamOffset)
		if check == checkProcess {
			res = append(res, msg)
		}
		if check != checkGap {
			i++
			continue // Continue to next message
		}
		// else checkGap

		if !streamID.ConfLevel().IsFuzzy() {
			return nil, errors.New("unexpected gap in finalized msg offsets [BUG]", "offset", msg.StreamOffset)
		}

		// Re-orgs of fuzzy conf levels are expected and can create gaps, block until ConfFinalized fills the gap.
		log.Warn(ctx, "Gap in fuzzy msg offsets, waiting for ConfFinalized", nil, "stream", streamID, "offset", msg.StreamOffset)
		backoff()
		// Retry the same message again
	}

	return res, nil
}

// fromChainVersionOffsets calculates the starting block offsets for all chain versions (to the destination chain).
func fromChainVersionOffsets(
	destChainID uint64, // Destination chain ID
	cursors []xchain.StreamCursor, // All actual on-chain submit cursors
	chainVers []xchain.ChainVersion, // All expected chain versions
	state *State, // On-disk local state
) (map[xchain.ChainVersion]uint64, error) {
	res := make(map[xchain.ChainVersion]uint64)

	// Initialize all chain versions to start at 1 by default or if local state is present
	for _, chainVer := range chainVers {
		res[chainVer] = initialXBlockOffset

		// If local persisted state is higher, use that instead, skipping a bunch of empty blocks on startup.
		if offset := state.GetOffset(destChainID, chainVer); offset > initialXBlockOffset {
			res[chainVer] = offset
		}
	}

	// sort cursors by decreasing offset, so we start streaming from minimum offset per source chain
	sort.Slice(cursors, func(i, j int) bool {
		return cursors[i].BlockOffset > cursors[j].BlockOffset
	})

	for _, cursor := range cursors {
		if cursor.DestChainID != destChainID {
			return nil, errors.New("unexpected cursor [BUG]")
		}

		offset, ok := res[cursor.ChainVersion()]
		if !ok {
			return nil, errors.New("unexpected cursor [BUG]")
		}

		if offset >= cursor.BlockOffset {
			continue // Skip if local state is higher than cursor
		}

		res[cursor.ChainVersion()] = cursor.BlockOffset
	}

	return res, nil
}

// msgOffsetFilter is a filter that keeps track of the last processed message offset for each stream.
// It is used to filter out messages that have already been processed.
//
// More specifically, it ensures that fuzzy msgs are submitted either from fuzzy or finalized attestations, whichever comes first.
type msgOffsetFilter struct {
	mu      sync.Mutex
	offsets map[xchain.StreamID]uint64
}

func newMsgOffsetFilter(cursors []xchain.StreamCursor) *msgOffsetFilter {
	offsets := make(map[xchain.StreamID]uint64, len(cursors))
	for _, cursor := range cursors {
		offsets[cursor.StreamID] = cursor.MsgOffset
	}

	return &msgOffsetFilter{
		offsets: offsets,
	}
}

type checkResult int

const (
	// checkProcess indicates that the message offset is sequential and should be processed.
	checkProcess checkResult = iota
	// checkGap indicates that the message offset is too far ahead and therefore contains a gap.
	checkGap
	// checkIgnore indicates that the message offset was already processed and should be ignored.
	checkIgnore
)

// Check updates the stream state and returns checkProcess if the provided offset is sequential.
// Otherwise it does not update the state and returns checkGap if the next message is too far ahead,
// or checkIgnore if the next message was already processed.
func (f *msgOffsetFilter) Check(stream xchain.StreamID, msgOffset uint64) checkResult {
	f.mu.Lock()
	defer f.mu.Unlock()

	expect := f.offsets[stream] + 1
	if msgOffset > expect {
		return checkGap
	} else if msgOffset < expect {
		return checkIgnore
	}

	// Update the offset
	f.offsets[stream] = msgOffset

	return checkProcess
}
