package cursors

import (
	"context"
	"slices"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
)

// todo mind the query limit on db

type Watcher struct {
	cursors  CursorTable
	provider xchain.Provider
}

func NewWatcher(cursors CursorTable, provider xchain.Provider) (*Watcher, error) {
	return &Watcher{
		cursors,
		provider,
	}, nil
}

func (w *Watcher) getStreamCursors(ctx context.Context, stream xchain.StreamID) ([]*Cursor, error) {
	iterator, err := w.cursors.List(
		ctx,
		CursorPrimaryKey{}.WithSrcChainIdConfLevelDstChainId(
			stream.SourceChainID,
			uint32(stream.ConfLevel()),
			stream.DestChainID,
		))

	if err != nil {
		return nil, errors.Wrap(err, "list cursors")
	}
	defer iterator.Close()

	var cursorsDesc []*Cursor
	for iterator.Next() {
		cursor, err := iterator.Value()
		if err != nil {
			return nil, errors.Wrap(err, "cursor value")
		}
		cursorsDesc = append(cursorsDesc, cursor)
	}

	return cursorsDesc, nil
}

// ConfirmStream checks all cursors for the provided stream and confirms them according to the rules:
//   - if a cursor is not empty fetch finalized cursor from the network and confirm only if finalized
//     message offset is higher to the cursor last message offset
//   - if a cursor is empty, confirm it if the previous cursor was confirmed
func (w *Watcher) ConfirmStream(ctx context.Context, stream xchain.StreamID) error {
	cursors, err := w.getStreamCursors(ctx, stream)
	if err != nil {
		return err
	}

	// check if contains non-empty unconfirmed cursor and only then fetch latest on-chain
	// submitted cursor, to optimize network load
	containsNonEmptyUnconfirmed := slices.ContainsFunc(cursors, func(c *Cursor) bool {
		return !c.Empty() && !c.GetConfirmed()
	})

	var final xchain.SubmitCursor
	if containsNonEmptyUnconfirmed {
		f, ok, err := w.provider.GetSubmittedCursor(ctx, xchain.FinalizedRef, stream)
		if err != nil {
			return errors.Wrap(err, "submitted cursor")
		} else if !ok { // no cursors available yet skip
			return nil
		}
		final = f
	}

	var prevConfirmed bool
	for _, cursor := range cursors {
		// if cursor not empty, confirm it comparing to the network finalized cursor
		if !cursor.Empty() {
			confirmed := final.MsgOffset >= cursor.GetLastXmsgOffset()
			cursor.Confirmed = confirmed
			// we reached latest confirmed break until next cycle
			if !confirmed {
				break
			}
		}

		cursor.Confirmed = cursor.GetConfirmed() || prevConfirmed
		if err := w.cursors.Update(ctx, cursor); err != nil {
			return errors.Wrap(err, "cursor update")
		}
		prevConfirmed = cursor.GetConfirmed()
	}

	return nil
}

// TrimStream iterates over all stored cursors in reverse order and when it finds the first
// confirmed it leaves it but deletes all previous cursors as they are confirmed.
func (w *Watcher) TrimStream(ctx context.Context, stream xchain.StreamID) error {
	cursors, err := w.getStreamCursors(ctx, stream)
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

		if err := w.cursors.Delete(ctx, c); err != nil {
			return errors.Wrap(err, "delete cursor")
		}
	}

	return nil
}
