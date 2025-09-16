// Package cursor provides a persisted attestation streamer cursor store.
// The goal of this package is to persist and identify
// safe/finalized bootstrap cursors per streamer.
//
// Problems it solves:
// - Relayer source of truth is submitted (and finalized) xmsg offsets on destination portals.
// - An attestation is only "fully relayed" when all xmsgs are submitted AND finalized.
// - Empty attestations are not submitted, so no proof exists that is has been processed.
// - While submissions are not finalized, they can technically reorg, and need to be resubmitted.
//
// Solution:
// - Persist attestations (cursors) immediately once submitted, but don't mark as "confirmed".
// - Monitor finalized submit cursors on destination portals.
// - Mark any cursor lower than on-chain finalized offsets as confirmed.
// - Empty cursors are confirmed once previous non-empty cursor is confirmed.
// - On worker bootstrap, use latest confirmed cursors if higher than on-chain offset.
package cursor

import (
	"context"
	"maps"
	"slices"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
)

// streamer represents a process calling to xchain.Provider.StreamAttestations
// by a worker. The goal of this package is to persist and identify
// safe/finalized bootstrap cursors per streamer.
type streamer struct {
	SrcChainID   uint64
	SrcConfLevel uint32
	DstChainID   uint64
}

func (s streamer) ChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{ID: s.SrcChainID, ConfLevel: xchain.ConfLevel(s.SrcConfLevel)}
}

// submitCursorFunc abstracts xchain.Provider.GetSubmitCursor method for testing purposes.
type submitCursorFunc func(ctx context.Context, ref xchain.Ref, stream xchain.StreamID) (xchain.SubmitCursor, bool, error)

// Store provides a persisted attestation streamer cursor store.
type Store struct {
	mu               sync.RWMutex
	db               CursorTable
	submitCursorFunc submitCursorFunc
	network          netconf.Network
}

// New returns a new cursor store.
func New(
	db db.DB,
	submitCursorFunc submitCursorFunc,
	network netconf.Network,
) (*Store, error) {
	cursorsTable, err := newCursorsTable(db)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:               cursorsTable,
		submitCursorFunc: submitCursorFunc,
		network:          network,
	}, nil
}

// WorkerOffsets returns confirmed offsets for the provided destination chain.
func (s *Store) WorkerOffsets(
	ctx context.Context,
	destChain uint64,
) (map[xchain.ChainVersion]uint64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

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

// Insert cursor for the provided streamer if it doesn't exist, otherwise ignore (keep existing).
func (s *Store) Insert(
	ctx context.Context,
	srcVersion xchain.ChainVersion,
	destChain uint64,
	attestOffset uint64,
	submittedMsgs map[xchain.StreamID][]xchain.Msg,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Get the highest stream offset for each shard
	offsetsByShard := make(map[uint64]uint64)
	for streamID, msgs := range submittedMsgs {
		if len(msgs) == 0 {
			continue
		}
		offsetsByShard[uint64(streamID.ShardID)] = msgs[len(msgs)-1].StreamOffset
	}

	ctx = log.WithCtx(ctx,
		"src_chain_version", s.network.ChainVersionName(srcVersion),
		"dest_chain", s.network.ChainName(destChain),
		"attest_offset", attestOffset,
		"stream_offsets", offsetsByShard,
	)

	c := &Cursor{
		SrcChainId:           srcVersion.ID,
		ConfLevel:            uint32(srcVersion.ConfLevel),
		DstChainId:           destChain,
		AttestOffset:         attestOffset,
		Confirmed:            false,
		StreamOffsetsByShard: offsetsByShard,
	}

	err := s.db.Insert(ctx, c)
	if errors.Is(err, ormerrors.AlreadyExists) {
		// Cursor already exists, verify that offsets are identical
		existing, err := s.db.Get(ctx, srcVersion.ID, uint32(srcVersion.ConfLevel), destChain, attestOffset)
		if err != nil {
			return errors.Wrap(err, "get cursor")
		} else if !maps.Equal(offsetsByShard, existing.GetStreamOffsetsByShard()) { // For now just log an error if this happens.
			log.Error(ctx, "Unexpected existing cursor offset mismatch", nil,
				"existing", existing.GetStreamOffsetsByShard(),
			)
		}

		return nil // Don't update latest metric (since existing may be confirmed).
	} else if err != nil {
		return errors.Wrap(err, "insert cursor")
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

// confirmForever blocks until the context is closed and periodically confirms cursors.
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
	s.mu.Lock()
	defer s.mu.Unlock()

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
				if evmchain.IsDisabled(c.GetSrcChainId()) || evmchain.IsDisabled(c.GetDstChainId()) {
					continue // Treat disabled chains as always confirmed.
				}

				stream := xchain.StreamID{SourceChainID: c.GetSrcChainId(), DestChainID: c.GetDstChainId(), ShardID: xchain.ShardID(shardID)}
				submitted, ok, err := s.submitCursorFunc(ctx, xchain.FinalizedRef, stream)
				if err != nil {
					log.Warn(ctx, "Get submit cursor failed while confirming (will retry)", err, "stream", s.network.StreamName(stream))
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

// trimForever blocks until the context is closed and periodically trims cursors.
func (s *Store) trimForever(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := s.trimOnce(ctx); err != nil {
				log.Warn(ctx, "Trimming cursor store failed (will retry)", err)
			}
		}
	}
}

// trimOnce iterates over all streamer's cursors and deletes all initial confirmed
// cursors, except the last confirmed one before any unconfirmed cursors.
func (s *Store) trimOnce(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

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

// splitByStreamer groups cursors by streamer.
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
