package cursor

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	db "github.com/cosmos/cosmos-db"
)

// Cursors implements operations on the persisted cursors for network.
type Cursors struct {
	network         netconf.Network
	cursors         CursorTable
	streams         []*StreamCursors
	confirmInterval time.Duration
}

func NewCursors(
	network netconf.Network,
	db db.DB,
	xProvider xchain.Provider,
	confirmInterval time.Duration,
) (*Cursors, error) {
	cursors, err := NewCursorsTable(db)
	if err != nil {
		return nil, err
	}

	return &Cursors{
		network:         network,
		cursors:         cursors,
		streams:         newNetworkCursors(network, cursors, xProvider),
		confirmInterval: confirmInterval,
	}, nil
}

// Confirmed loads all confirmed cursor for a destination chain
//
// Some cursors might not yet exist so they are not included in the result.
func (c *Cursors) Confirmed(ctx context.Context, destChainID uint64) ([]xchain.SubmitCursor, error) {
	var cursors []xchain.SubmitCursor

	for _, stream := range c.streams {
		var streamsDest bool
		for _, s := range c.network.StreamsTo(destChainID) {
			if s == stream.stream {
				streamsDest = true
				break
			}
		}
		if !streamsDest {
			continue
		}

		cursor, ok, err := stream.GetConfirmed(ctx)
		if err != nil {
			return nil, err
		} else if !ok {
			continue
		}

		cursors = append(cursors, *cursor)
	}

	return cursors, nil
}

// Add a cursor to the storage constructed from the stream, attestation and messages
//
// If a cursor already exists it is updated with the provided data.
func (c *Cursors) Add(
	ctx context.Context,
	stream xchain.StreamID,
	attestationOffset uint64,
	messages []xchain.Msg,
) error {
	var last uint64
	if msgLen := len(messages); msgLen > 0 {
		last = messages[msgLen-1].MsgID.StreamOffset
	}

	cursor := &Cursor{
		SrcChainId:     stream.SourceChainID,
		DstChainId:     stream.DestChainID,
		ConfLevel:      uint32(stream.ConfLevel()),
		AttestOffset:   attestationOffset,
		LastXmsgOffset: last,
		Confirmed:      false,
	}

	err := c.cursors.Save(ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "insert cursor")
	}

	log.Info(ctx, "New cursor persisted",
		"stream", c.network.StreamName(stream),
		"attest_offset", attestationOffset,
		"msg_offset", last,
	)

	return nil
}

// Monitor all existing cursors in the storage and confirm them once the submission is finalized.
func (c *Cursors) Monitor(ctx context.Context) {
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	for ctx.Err() == nil {
		if err := c.runOnce(ctx); err != nil {
			log.Error(ctx, "Cursor worker failed, resetting", err)
		}

		if ctx.Err() != nil {
			return
		}

		backoff()
	}
}

func (c *Cursors) runOnce(ctx context.Context) error {
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(c.confirmInterval))

	for {
		for _, stream := range c.streams {
			if err := stream.Confirm(ctx); err != nil {
				return err
			}

			if err := stream.Trim(ctx); err != nil {
				return err
			}
		}

		backoff()
	}
}
