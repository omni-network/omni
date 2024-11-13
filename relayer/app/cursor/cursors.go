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
	cursors         CursorTable
	confirmInterval time.Duration
	xProvider       xchain.Provider
}

func NewCursors(
	db db.DB,
	xProvider xchain.Provider,
	confirmInterval time.Duration,
) (*Cursors, error) {
	cursors, err := NewCursorsTable(db)
	if err != nil {
		return nil, err
	}

	return &Cursors{
		cursors:         cursors,
		confirmInterval: confirmInterval,
		xProvider:       xProvider,
	}, nil
}

// ConfirmedOffset loads confirmed cursor offset for the chain version and destination chain ID.
// If a cursor doesn't exist a 0, false is returned.
func (c *Cursors) ConfirmedOffset(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	destChainID uint64,
) (uint64, bool, error) {
	return newStreamCursors(chainVer.ID, destChainID, chainVer.ConfLevel, c.cursors, c.xProvider).
		GetConfirmed(ctx)
}

// Save a cursor to the storage.
// If a cursor already exists it is updated with the provided data.
func (c *Cursors) Save(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	destChainID uint64,
	attestationOffset uint64,
	empty bool,
) error {
	cursor := &Cursor{
		SrcChainId:   chainVer.ID,
		ConfLevel:    uint32(chainVer.ConfLevel),
		DstChainId:   destChainID,
		AttestOffset: attestationOffset,
		Confirmed:    false,
		Empty:        empty,
	}

	err := c.cursors.Save(ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "insert cursor")
	}

	log.Info(ctx, "New cursor saved",
		"src_chain", chainVer.ID,
		"conf_level", chainVer.ConfLevel,
		"attest_offset", attestationOffset,
	)

	return nil
}

// Monitor all existing cursors for the network in the storage and
// confirm them once the submission is finalized.
func (c *Cursors) Monitor(ctx context.Context, network netconf.Network) {
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	for ctx.Err() == nil {
		if err := c.runOnce(ctx, network); err != nil {
			log.Error(ctx, "Cursor worker failed, resetting", err)
		}

		if ctx.Err() != nil {
			return
		}

		backoff()
	}
}

func (c *Cursors) runOnce(ctx context.Context, network netconf.Network) error {
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(c.confirmInterval))

	var streams []*StreamCursors
	for _, chain := range network.EVMChains() {
		for _, streamID := range network.StreamsTo(chain.ID) {
			streamCursors := newStreamCursors(streamID.SourceChainID, streamID.DestChainID, streamID.ConfLevel(), c.cursors, c.xProvider)
			streams = append(streams, streamCursors)
		}
	}

	for {
		for _, stream := range streams {
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
