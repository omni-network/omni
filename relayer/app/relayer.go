package relayer

import (
	"context"
	"sort"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
)

func StartRelayer(
	ctx context.Context,
	cProvider cchain.Provider,
	chainIDs []uint64,
	xClient XChainClient,
	creator Creator,
	sender Sender,
) {
	// Get the last submitted cursors for each chain.
	var cursors []xchain.StreamCursor                  // All submitted cursors from all chains.
	initialOffsets := make(map[xchain.StreamID]uint64) // Initial submitted offsets for each stream.
	for _, chainID := range chainIDs {
		submitted, _ := xClient.GetSubmittedCursors(ctx, chainID)
		for _, cursor := range submitted {
			initialOffsets[cursor.StreamID] = cursor.Offset
			cursors = append(cursors, cursor)
		}
	}

	// callback processes each approved attestation/xblock.
	callback := func(ctx context.Context, att xchain.AggAttestation) error {
		// Get the xblock from the source chain.
		block, ok, err := xClient.GetBlock(ctx, att.SourceChainID, att.BlockHeight)

		if err != nil {
			return err
		} else if !ok {
			return errors.New("attestation block not finalized", "attestation", att)
		}

		// Split into streams
		for streamID, msgs := range mapByStreamID(block.Msgs) {
			msgs = filterMsgs(msgs, initialOffsets[streamID]) // Filter out any partially submitted stream updates.
			if len(msgs) == 0 {
				continue
			}

			update := StreamUpdate{
				StreamID:       streamID,
				AggAttestation: att,
				Msgs:           msgs,
			}

			submissions, err := creator.CreateSubmissions(ctx, update)
			if err != nil {
				return err
			}

			for _, subs := range submissions {
				if err := sender.SendTransaction(ctx, subs); err != nil {
					return err
				}
			}
		}

		return nil
	}

	// Subscribe to attestations for each chain.
	for chainID, fromHeight := range FromHeights(cursors, chainIDs) {
		cProvider.Subscribe(ctx, chainID, fromHeight, callback)
	}
}

func mapByStreamID(msgs []xchain.Msg) map[xchain.StreamID][]xchain.Msg {
	m := make(map[xchain.StreamID][]xchain.Msg)
	for _, msg := range msgs {
		m[msg.StreamID] = append(m[msg.StreamID], msg)
	}

	return m
}

func filterMsgs(msgs []xchain.Msg, offset uint64) []xchain.Msg {
	var res []xchain.Msg //nolint:prealloc // we don't know how many msgs will be filtered out
	for _, msg := range msgs {
		if msg.StreamOffset <= offset {
			// filter msgs lower than offset
			continue
		}
		res = append(res, msg)
	}

	return res
}

func FromHeights(cursors []xchain.StreamCursor, chainIDs []uint64) map[uint64]uint64 {
	res := make(map[uint64]uint64)

	for _, chainID := range chainIDs {
		res[chainID] = 0
	}

	// sort cursors by decreasing SourceBlockHeight so we start streaming from minimum height per source chain
	sort.Slice(cursors, func(i, j int) bool {
		return cursors[i].SourceBlockHeight > cursors[j].SourceBlockHeight
	})

	for _, cursor := range cursors {
		res[cursor.SourceChainID] = cursor.SourceBlockHeight
	}

	return res
}
