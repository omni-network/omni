package cursor

import (
	"context"
	"slices"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	db "github.com/cosmos/cosmos-db"
)

type getSubmitCursorFunc func(ctx context.Context, ref xchain.Ref, stream xchain.StreamID) (xchain.SubmitCursor, bool, error)

// Store provides a persisted attestation streamer cursor store.
type Store struct {
	db                  CursorTable
	getSubmitCursorFunc getSubmitCursorFunc
	network             netconf.Network
}

func New(
	db db.DB,
	getSubmitCursorFunc getSubmitCursorFunc,
	network netconf.Network,
) (*Store, error) {
	cursorsTable, err := newCursorsTable(db)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:                  cursorsTable,
		getSubmitCursorFunc: getSubmitCursorFunc,
		network:             network,
	}, nil
}

// WorkerOffsets returns confirmed offsets for the provided destination chain.
func (s *Store) WorkerOffsets(
	ctx context.Context,
	destChain uint64,
) (map[xchain.ChainVersion]uint64, error) {
	all, err := listAll(ctx, s.db)
	if err != nil {
		return nil, err
	}

	// Collect the highest confirmed cursor for each streamer.
	resp := make(map[xchain.ChainVersion]uint64)
	for s, cursors := range splitByStreamer(all) {
		if s.DstChainID != destChain {
			continue
		}

		slices.Reverse(cursors)

		for _, c := range cursors {
			if c.GetConfirmed() {
				resp[s.ChainVersion()] = c.GetAttestOffset()
				break
			}
		}
	}

	return resp, nil
}

// Save an attest offset for the provided streamer.
// Existing cursors' stream offsets are updated and confirmed is reset to false.
func (s *Store) Save(
	ctx context.Context,
	srcVersion xchain.ChainVersion,
	destChain uint64,
	attestOffset uint64,
	streamMsgs map[xchain.StreamID][]xchain.Msg,
) error {
	// Get highest stream offset for each shard
	offsetsByShard := make(map[uint64]uint64)
	for streamID, msgs := range streamMsgs {
		if len(msgs) == 0 {
			continue
		}
		offsetsByShard[uint64(streamID.ShardID)] = msgs[len(msgs)-1].StreamOffset
	}

	c := &Cursor{
		SrcChainId:           srcVersion.ID,
		ConfLevel:            uint32(srcVersion.ConfLevel),
		DstChainId:           destChain,
		AttestOffset:         attestOffset,
		Confirmed:            false,
		StreamOffsetsByShard: offsetsByShard,
	}

	err := s.db.Save(ctx, c)
	if err != nil {
		return errors.Wrap(err, "save cursor")
	}

	latestOffset.
		WithLabelValues(s.network.ChainVersionName(srcVersion), s.network.ChainName(destChain)).
		Set(float64(attestOffset))

	return nil
}

// StartLoops starts goroutines to periodically confirm and trim streams.
// It returns immediately.
func (s *Store) StartLoops(ctx context.Context) {
	go s.trimForever(ctx)
	go s.confirmForever(ctx)
}

func (s *Store) confirmForever(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := s.confirmOnce(ctx); err != nil {
				log.Error(ctx, "Trimming cursor stored failed (will retry))", err)
			}
		}
	}
}

// confirmOnce marks cursors of each streamer as confirmed based on:
// - if previous cursor is confirmed
// - and if all xmsgs are submitted and finalized.
func (s *Store) confirmOnce(ctx context.Context) error {
	all, err := listAll(ctx, s.db)
	if err != nil {
		return errors.Wrap(err, "listAll all cursors")
	}

	for streamer, cursors := range splitByStreamer(all) {
		for _, c := range cursors {
			if c.GetConfirmed() {
				continue
			}

			var unconfirmed bool
			for shardID, offset := range c.GetStreamOffsetsByShard() {
				stream := xchain.StreamID{SourceChainID: c.GetSrcChainId(), DestChainID: c.GetDstChainId(), ShardID: xchain.ShardID(shardID)}
				submitted, ok, err := s.getSubmitCursorFunc(ctx, xchain.FinalizedRef, stream)
				if err != nil {
					log.Warn(ctx, "Get submit cursor failed while confirming", err, "stream", s.network.StreamName(stream))
					unconfirmed = true

					break
				} else if !ok || submitted.MsgOffset < offset {
					unconfirmed = true // Cursor not yet submitted or finalized.
					break
				}
			}

			if unconfirmed {
				break
			}

			c.Confirmed = true

			confirmedOffset.
				WithLabelValues(s.network.ChainVersionName(streamer.ChainVersion()), s.network.ChainName(c.GetDstChainId())).
				Set(float64(c.GetAttestOffset()))

			if err := s.db.Save(ctx, c); err != nil {
				return errors.Wrap(err, "save cursor")
			}
		}
	}

	return nil
}

func (s *Store) trimForever(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := s.trimOnce(ctx); err != nil {
				log.Error(ctx, "Trimming cursor stored failed (will retry))", err)
			}
		}
	}
}

// trimOnce iterates over all streamer's cursors and deletes all initial confirmed
// cursors, except the last confirmed one before any unconfirmed cursors.
func (s *Store) trimOnce(ctx context.Context) error {
	all, err := listAll(ctx, s.db)
	if err != nil {
		return errors.Wrap(err, "listAll all cursors")
	}

	for _, cursors := range splitByStreamer(all) {
		var prev *Cursor
		for i, c := range cursors {
			if i > 0 && c.GetAttestOffset() <= prev.GetAttestOffset() { // Sanity check
				return errors.New("cursors are not sorted by attest offset [BUG]")
			}

			if !c.GetConfirmed() {
				break
			}
			if i > 0 {
				if err := s.db.Delete(ctx, prev); err != nil {
					return errors.Wrap(err, "delete cursor")
				}
			}
			prev = c
		}
	}

	return nil
}

type streamer struct {
	SrcChainID   uint64
	SrcConfLevel uint32
	DstChainID   uint64
}

func (s streamer) ChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{ID: s.SrcChainID, ConfLevel: xchain.ConfLevel(s.SrcConfLevel)}
}

func splitByStreamer(all []*Cursor) map[streamer][]*Cursor {
	resp := make(map[streamer][]*Cursor)
	for _, c := range all {
		s := streamer{
			SrcChainID:   c.GetSrcChainId(),
			SrcConfLevel: c.GetConfLevel(),
			DstChainID:   c.GetDstChainId(),
		}

		resp[s] = append(resp[s], c)
	}

	return resp
}
