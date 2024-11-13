package cursor

import (
	"context"
	"slices"
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
	db              CursorTable
	confirmInterval time.Duration
	xProvider       xchain.Provider
	streams         []stream
	network         netconf.Network
}

func NewCursors(
	db db.DB,
	xProvider xchain.Provider,
	confirmInterval time.Duration,
	network netconf.Network,
) (*Cursors, error) {
	cursorsTable, err := NewCursorsTable(db)
	if err != nil {
		return nil, err
	}

	var streams []stream
	for _, chain := range network.EVMChains() {
		for _, streamID := range network.StreamsTo(chain.ID) {
			streams = append(streams, stream{
				SrcVersion: xchain.ChainVersion{
					ID:        streamID.SourceChainID,
					ConfLevel: streamID.ConfLevel(),
				},
				DestChain: chain,
				Network:   network,
			})
		}
	}

	return &Cursors{
		db:              cursorsTable,
		confirmInterval: confirmInterval,
		xProvider:       xProvider,
		streams:         streams,
		network:         network,
	}, nil
}

// ConfirmedOffset loads confirmed cursor offset for the chain version and destination chain ID.
// If a cursor doesn't exist a 0, false is returned.
func (c *Cursors) ConfirmedOffset(
	ctx context.Context,
	srcVersion xchain.ChainVersion,
	destChain netconf.Chain,
) (uint64, bool, error) {
	cursors, err := getAllStreamCursors(ctx, c.db, stream{srcVersion, destChain, c.network})
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

// Save a cursor to the storage.
// If a cursor already exists it is updated with the provided data.
func (c *Cursors) Save(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	destChainID uint64,
	attestationOffset uint64,
	empty bool,
	chainVerName string,
	destChainName string,
) error {
	cursor := &Cursor{
		SrcChainId:   chainVer.ID,
		ConfLevel:    uint32(chainVer.ConfLevel),
		DstChainId:   destChainID,
		AttestOffset: attestationOffset,
		Confirmed:    false,
		Empty:        empty,
	}

	err := c.db.Save(ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "insert cursor")
	}

	log.Debug(ctx, "New cursor saved",
		"src_chain", chainVer.ID,
		"conf_level", chainVer.ConfLevel,
		"attest_offset", attestationOffset,
	)

	latestOffset.WithLabelValues(chainVerName, destChainName).Set(float64(cursor.GetAttestOffset()))

	return nil
}

// Monitor all existing cursors for the network in the storage and
// confirm them once the submission is finalized.
func (c *Cursors) Monitor(ctx context.Context) {
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	for ctx.Err() == nil {
		if err := c.runOnce(ctx); err != nil {
			log.Error(ctx, "Cursor worker failed, will retry", err)
		}

		if ctx.Err() != nil {
			return
		}

		backoff()
	}
}

func (c *Cursors) runOnce(ctx context.Context) error {
	ticker := time.NewTicker(c.confirmInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context canceled")
		case <-ticker.C:
			for _, s := range c.streams {
				if err := ConfirmStream(ctx, c.db, c.network, c.xProvider, s); err != nil {
					return err
				}

				if err := TrimStream(ctx, c.db, s); err != nil {
					return err
				}
			}
		}
	}
}
