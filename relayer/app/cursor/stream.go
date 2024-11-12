package cursor

import (
	"context"
	"slices"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// StreamCursors provides operations on cursors that share same chain StreamID
// Most importantly it provides operations to confirm and trim cursors
// persisted in the storage.
type StreamCursors struct {
	cursors  CursorTable
	provider xchain.Provider
	stream   xchain.StreamID
}

func NewStreamCursors(stream xchain.StreamID, cursors CursorTable, provider xchain.Provider) *StreamCursors {
	return &StreamCursors{
		cursors:  cursors,
		provider: provider,
		stream:   stream,
	}
}

func newNetworkCursors(network netconf.Network, cursors CursorTable, provider xchain.Provider) []*StreamCursors {
	var streams []*StreamCursors
	for _, chain := range network.EVMChains() {
		for _, stream := range network.StreamsTo(chain.ID) {
			streams = append(streams, NewStreamCursors(stream, cursors, provider))
		}
	}

	return streams
}

// GetConfirmed returns the latest confirmed cursor for the stream.
// If no cursors have yet been confirmed or stored nil, false is returned.
func (s *StreamCursors) GetConfirmed(ctx context.Context) (*xchain.SubmitCursor, bool, error) {
	cursors, err := s.getAllCursors(ctx)
	if err != nil {
		return nil, false, err
	}

	slices.Reverse(cursors)
	for _, c := range cursors {
		if c.GetConfirmed() {
			return c.ToSubmitCursor(), true, nil
		}
	}

	return nil, false, nil
}

// getAllCursors returns all cursors from the storage.
//
// Cursors are ordered by source chain ID, confirmation level and destination chain ID in ascending order.
func (s *StreamCursors) getAllCursors(ctx context.Context) ([]*Cursor, error) {
	iterator, err := s.cursors.List(ctx,
		CursorPrimaryKey{}.WithSrcChainIdConfLevelDstChainId(
			s.stream.SourceChainID,
			uint32(s.stream.ConfLevel()), // todo validate
			s.stream.DestChainID,
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

	var prevConfirmed bool
	for _, cursor := range cursors {
		// if cursor not empty, confirm it comparing to the network finalized cursor
		if !cursor.Empty() && !cursor.GetConfirmed() {
			cursor.Confirmed, err = s.isFinalized(ctx, cursor)
			if err != nil {
				return err
			}
			// we reached latest confirmed break until next cycle
			if !cursor.GetConfirmed() {
				break
			}
		}

		cursor.Confirmed = cursor.GetConfirmed() || prevConfirmed
		if err := s.cursors.Update(ctx, cursor); err != nil {
			return errors.Wrap(err, "cursor update")
		}
		prevConfirmed = cursor.GetConfirmed()
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
	final, ok, err := s.provider.GetSubmittedCursor(ctx, xchain.FinalizedRef, s.stream)
	if err != nil {
		return false, errors.Wrap(err, "submitted cursor")
	} else if !ok { // no cursors available yet skip
		return false, nil
	}

	return final.MsgOffset >= cursor.GetAttestOffset(), nil
}
