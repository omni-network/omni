package cursor

import (
	"context"
	"math"
	"slices"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

type stream struct {
	SrcVersion xchain.ChainVersion
	DestChain  netconf.Chain
	Network    netconf.Network
}

// getAllStreamCursors returns all cursors from the storage.
//
// Cursors are ordered by source chain ID, confirmation level and destination chain ID in ascending order.
func getAllStreamCursors(ctx context.Context, db CursorTable, stream stream) ([]*Cursor, error) {
	iterator, err := db.List(ctx,
		CursorPrimaryKey{}.WithSrcChainIdConfLevelDstChainId(
			stream.SrcVersion.ID,
			uint32(stream.SrcVersion.ConfLevel),
			stream.DestChain.ID,
		))

	if err != nil {
		return nil, errors.Wrap(err, "list cursors")
	}
	defer iterator.Close()

	var cursors []*Cursor
	for iterator.Next() {
		cursor, err := iterator.Value()
		if err != nil {
			return nil, errors.Wrap(err, "cursor value")
		}
		cursors = append(cursors, cursor)
	}

	return cursors, nil
}

// ConfirmStream checks all stream cursors and confirms them according to the rules:
//   - if a cursor is not empty fetch finalized cursor from the network and confirm only if finalized
//     message offset is higher to the cursor last message offset
//   - if a cursor is empty, confirm it if the previous cursor was confirmed
func ConfirmStream(ctx context.Context, db CursorTable, network netconf.Network, xProvider xchain.Provider, stream stream) error {
	cursors, err := getAllStreamCursors(ctx, db, stream)
	if err != nil {
		return err
	}

	// if we only have empty and not confirmed skip as there's nothing to do
	nonEmptyOrConfirmed := slices.ContainsFunc(cursors, func(c *Cursor) bool {
		return !c.GetEmpty() || c.GetConfirmed()
	})
	if !nonEmptyOrConfirmed {
		return nil
	}

	for _, cursor := range cursors {
		// if cursor not empty, confirm it comparing to the network finalized cursor
		if !cursor.GetEmpty() && !cursor.GetConfirmed() {
			confirmed, err := isFinalized(ctx, network, xProvider, cursor)
			if err != nil {
				return err
			} else if !confirmed {
				// we reached latest confirmed
				break
			}
		}

		cursor.Confirmed = true
		if err := db.Update(ctx, cursor); err != nil {
			return errors.Wrap(err, "cursor update")
		}

		confirmedOffset.
			WithLabelValues(network.ChainVersionName(stream.SrcVersion), stream.DestChain.Name).
			Set(float64(cursor.GetAttestOffset()))
	}

	return nil
}

// TrimStream iterates over all stored stream cursors in reverse order and when it finds the first
// confirmed it leaves it but deletes all previous cursors as they are confirmed.
func TrimStream(ctx context.Context, db CursorTable, stream stream) error {
	cursors, err := getAllStreamCursors(ctx, db, stream)
	if err != nil {
		return err
	}

	var confirmed bool
	slices.Reverse(cursors)
	for _, c := range cursors {
		// we iterate until we find first confirmed, skip first then delete all after
		if !confirmed {
			confirmed = c.GetConfirmed()
			continue
		}

		if err := db.Delete(ctx, c); err != nil {
			return errors.Wrap(err, "delete cursor")
		}
	}

	return nil
}

// isFinalized checks if the cursor is finalized on the network.
func isFinalized(ctx context.Context, network netconf.Network, xProvider xchain.Provider, cursor *Cursor) (bool, error) {
	chain, ok := network.Chain(cursor.GetSrcChainId())
	if !ok {
		return false, errors.New("invalid dest chain id [BUG]")
	}

	lowestOffset := uint64(math.MaxUint64)
	for _, shard := range chain.Shards {
		final, ok, err := xProvider.GetSubmittedCursor(ctx, xchain.FinalizedRef, cursor.StreamID(shard))
		if err != nil {
			return false, errors.Wrap(err, "submitted cursor")
		} else if !ok { // no cursors available yet skip
			return false, nil
		}

		if final.AttestOffset < lowestOffset {
			lowestOffset = final.AttestOffset
		}
	}

	return lowestOffset >= cursor.GetAttestOffset(), nil
}
