package relayer

import (
	"context"
	"sort"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// StartRelayer starts the relayer logic by subscribing to approved aggregate attestations
// from the consensus chain and processing them into submissions for the destination chains.
func StartRelayer(
	ctx context.Context,
	cProvider cchain.Provider,
	network netconf.Network,
	xClient xchain.Provider,
	creator CreateFunc,
	sender SendFunc,
) error {
	// Get the last submitted cursors for each chain.
	cursors, initialOffsets, err := getSubmittedCursors(ctx, network.ChainIDs(), network.ChainIDs(), xClient)
	if err != nil {
		return err
	}

	// callback processes each approved attestation/xblock.
	callback := newCallback(xClient, initialOffsets, creator, sender)

	// Subscribe to attestations for each chain.
	for chainID, fromHeight := range FromHeights(cursors, network.Chains) {
		cProvider.Subscribe(ctx, chainID, fromHeight, callback)
	}

	return nil
}

// getSubmittedCursors returns the last submitted cursor for each chain.
// It also returns the offsets indexed by streamID for each stream.
func getSubmittedCursors(ctx context.Context, srcChains, dstChains []uint64, xClient xchain.Provider,
) ([]xchain.StreamCursor, map[xchain.StreamID]uint64, error) {
	var cursors []xchain.StreamCursor                  // All submitted cursors from all chains.
	initialOffsets := make(map[xchain.StreamID]uint64) // Initial submitted offsets for each stream.
	for _, destChain := range dstChains {
		for _, srcChain := range srcChains {
			if srcChain == destChain {
				continue
			}

			cursor, ok, err := xClient.GetSubmittedCursor(ctx, destChain, srcChain)
			if err != nil {
				return nil, nil, errors.Wrap(err, "failed to get submitted cursors",
					"dest_chain", destChain,
					"src_chain", srcChain,
				)
			} else if !ok {
				continue
			}

			initialOffsets[cursor.StreamID] = cursor.Offset
			cursors = append(cursors, cursor)
		}
	}

	return cursors, initialOffsets, nil
}

// TODO(corver): Add support for empty submissions by passing a map of chainIDs to generate empty submissions for.
func mapByStreamID(msgs []xchain.Msg) map[xchain.StreamID][]xchain.Msg {
	m := make(map[xchain.StreamID][]xchain.Msg)
	for _, msg := range msgs {
		m[msg.StreamID] = append(m[msg.StreamID], msg)
	}

	return m
}

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

func FromHeights(cursors []xchain.StreamCursor, chains []netconf.Chain) map[uint64]uint64 {
	res := make(map[uint64]uint64)

	for _, chain := range chains {
		res[chain.ID] = chain.DeployHeight
	}

	// sort cursors by decreasing SourceBlockHeight, so we start streaming from minimum height per source chain
	sort.Slice(cursors, func(i, j int) bool {
		return cursors[i].SourceBlockHeight > cursors[j].SourceBlockHeight
	})

	for _, cursor := range cursors {
		res[cursor.SourceChainID] = cursor.SourceBlockHeight
	}

	return res
}
