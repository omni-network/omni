package cursor

import (
	"context"
	"slices"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
)

// StreamCursors provides operations on cursors that share same source chain,
// destination chain and confirmation level.
//
// Most importantly it provides operations to confirm and trim cursors
// persisted in the storage.
type StreamCursors struct {
	cursors       CursorTable
	provider      xchain.Provider
	sourceChainID uint64
	destChainID   uint64
	confLevel     xchain.ConfLevel
	sourceName    string
	destName      string
}

func newStreamCursors(
	sourceChainID uint64,
	destChainID uint64,
	confLevel xchain.ConfLevel,
	cursors CursorTable,
	provider xchain.Provider,
	sourceName string,
	destName string,
) *StreamCursors {
	return &StreamCursors{
		cursors:       cursors,
		provider:      provider,
		sourceChainID: sourceChainID,
		destChainID:   destChainID,
		confLevel:     confLevel,
		sourceName:    sourceName,
		destName:      destName,
	}
}

// GetConfirmed returns the latest confirmed cursor attestation offset for the stream.
// If no cursors have yet been confirmed or stored nil, false is returned.
func (s *StreamCursors) GetConfirmed(ctx context.Context) (uint64, bool, error) {
	cursors, err := s.getAllCursors(ctx)
	if err != nil {
		return 0, false, err
	}

	slices.Reverse(cursors)
	for _, c := range cursors {
		if c.GetConfirmed() {
			return c.GetAttestOffset(), true, nil
		}
	}

	return 0, false, nil
}

// getAllCursors returns all cursors from the storage.
//
// Cursors are ordered by source chain ID, confirmation level and destination chain ID in ascending order.
func (s *StreamCursors) getAllCursors(ctx context.Context) ([]*Cursor, error) {
	iterator, err := s.cursors.List(ctx,
		CursorPrimaryKey{}.WithSrcChainIdConfLevelDstChainId(
			s.sourceChainID,
			uint32(s.confLevel),
			s.destChainID,
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

// Confirm checks all stream cursors and confirms them according to the rules:
//   - if a cursor is not empty fetch finalized cursor from the network and confirm only if finalized
//     message offset is higher to the cursor last message offset
//   - if a cursor is empty, confirm it if the previous cursor was confirmed
func (s *StreamCursors) Confirm(ctx context.Context) error {
	cursors, err := s.getAllCursors(ctx)
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
			confirmed, err := s.isFinalized(ctx, cursor)
			if err != nil {
				return err
			} else if !confirmed {
				// we reached latest confirmed
				break
			}
		}

		cursor.Confirmed = true
		if err := s.cursors.Update(ctx, cursor); err != nil {
			return errors.Wrap(err, "cursor update")
		}

		confirmedOffset.WithLabelValues(s.sourceName, s.destName).Set(float64(cursor.GetAttestOffset()))
	}

	return nil
}

// Trim iterates over all stored stream cursors in reverse order and when it finds the first
// confirmed it leaves it but deletes all previous cursors as they are confirmed.
func (s *StreamCursors) Trim(ctx context.Context) error {
	cursors, err := s.getAllCursors(ctx)
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

		if err := s.cursors.Delete(ctx, c); err != nil {
			return errors.Wrap(err, "delete cursor")
		}
	}

	return nil
}

// isFinalized checks if the cursor is finalized on the network.
func (s *StreamCursors) isFinalized(ctx context.Context, cursor *Cursor) (bool, error) {
	final, ok, err := s.provider.GetSubmittedCursor(ctx, xchain.FinalizedRef, cursor.StreamID())
	if err != nil {
		return false, errors.Wrap(err, "submitted cursor")
	} else if !ok { // no cursors available yet skip
		return false, nil
	}

	return final.MsgOffset >= cursor.GetAttestOffset(), nil
}
