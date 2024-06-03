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
) ([]xchain.SubmitCursor, error) {
	var cursors []xchain.SubmitCursor //nolint:prealloc // Not worth it.
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
func filterMsgs(ctx context.Context, streamID xchain.StreamID, valSetID uint64, msgs []xchain.Msg, msgFilter *msgCursorFilter) ([]xchain.Msg, error) {
	backoff := expbackoff.New(ctx)
	res := make([]xchain.Msg, 0, len(msgs)) // Res might have over-capacity, but that's fine, we only filter on startup.
	for i := 0; i < len(msgs); {
		msg := msgs[i]

		check, cursor := msgFilter.Check(streamID, valSetID, msg.StreamOffset)
		if check == checkProcess {
			res = append(res, msg)
		} else if check == checkOldVatSet {
			log.Warn(ctx, "Skipping msg with old valSetID", nil,
				"stream", streamID,
				"offset", msg.StreamOffset,
				"valset", valSetID,
				"cursor_valset", cursor.ValsSetID,
			)
		}
		if check != checkGapOffset {
			i++
			continue // Continue to next message
		}
		// else checkGap

		if !streamID.ConfLevel().IsFuzzy() {
			return nil, errors.New("unexpected gap in finalized msg offsets [BUG]",
				"stream", streamID,
				"offset", msg.StreamOffset,
				"cursor_offset", cursor.MsgOffset,
			)
		}

		// Re-orgs of fuzzy conf levels are expected and can create gaps, block until ConfFinalized fills the gap.
		log.Warn(ctx, "Gap in fuzzy msg offsets, waiting for ConfFinalized", nil,
			"stream", streamID,
			"offset", msg.StreamOffset,
			"cursor_offset", cursor.MsgOffset,
		)
		backoff()
		// Retry the same message again
	}

	return res, nil
}

// fromChainVersionOffsets calculates the starting block offsets for all chain versions (to the destination chain).
func fromChainVersionOffsets(
	destChainID uint64, // Destination chain ID
	cursors []xchain.SubmitCursor, // All actual on-chain submit cursors
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

// msgCursorFilter is a filter that keeps track of the last processed message cursor for each stream.
// It is used to filter out messages that have already been processed or that cannot be submitted otherwise.
//
// More specifically, it ensures that fuzzy msgs are submitted either from fuzzy or finalized attestations, whichever comes first.
// It also ensures that valSetID always increases.
type msgCursorFilter struct {
	mu      sync.Mutex
	cursors map[xchain.StreamID]streamCursor
}

type streamCursor struct {
	MsgOffset uint64
	ValsSetID uint64
}

func newMsgOffsetFilter(cursors []xchain.SubmitCursor) (*msgCursorFilter, error) {
	streamCursors := make(map[xchain.StreamID]streamCursor, len(cursors))
	for _, cursor := range cursors {
		streamCursors[cursor.StreamID] = streamCursor{
			MsgOffset: cursor.MsgOffset,
			ValsSetID: cursor.ValidatorSetID,
		}
	}
	if len(streamCursors) != len(cursors) {
		return nil, errors.New("unexpected duplicate cursors [BUG]")
	}

	return &msgCursorFilter{
		cursors: streamCursors,
	}, nil
}

//go:generate stringer -type=checkResult -trimprefix=check

type checkResult int

const (
	// checkProcess indicates that the message offset is sequential and should be processed.
	checkProcess checkResult = iota
	// checkGapOffset indicates that the message offset is too far ahead and therefore contains a gap.
	checkGapOffset
	// checkIgnoreOffset indicates that the message offset was already processed and should be ignored.
	checkIgnoreOffset
	// checkOldVatSet indicates that the message has a lower validator set ID than the last processed message and must be ignored.
	checkOldVatSet
)

// Check updates the stream state and returns checkProcess if the provided offset is sequential.
//
// Otherwise, it does not update the state.
// It returns checkGap if the next message is too far ahead,
// or checkIgnore if the next message was already processed,
// or checkOldVatSet if the next message has a lower validator set ID than the last processed message.
//
// It also returns the existing stream cursor.
func (f *msgCursorFilter) Check(stream xchain.StreamID, valSetID uint64, msgOffset uint64) (checkResult, streamCursor) {
	f.mu.Lock()
	defer f.mu.Unlock()

	cursor := f.cursors[stream]

	expectOffset := cursor.MsgOffset + 1
	if msgOffset > expectOffset {
		return checkGapOffset, cursor
	} else if msgOffset < expectOffset {
		return checkIgnoreOffset, cursor
	}

	if valSetID < cursor.ValsSetID {
		return checkOldVatSet, cursor
	}

	// Update the cursor
	f.cursors[stream] = streamCursor{
		MsgOffset: msgOffset,
		ValsSetID: valSetID,
	}

	return checkProcess, cursor
}
