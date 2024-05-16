//nolint:gosec // Use weak random for determinism.
package provider

import (
	"context"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

var (
	_ xchain.Provider = (*Mock)(nil)
)

// todo(Lazar): delete this and pass it to ctor so it's not hard coded and hidden here.
const (
	destChainA = 100
	destChainB = 200
)

// Mock is a mock implementation of the xchain.Provider interface.
// It generates deterministic blocks and messages for any chain that is queried.
// Except for omni consensus chain, we use the real cprovider to fetch blocks.
type Mock struct {
	period    time.Duration
	mu        sync.Mutex
	blocks    []xchain.Block
	uniq      map[uint64]bool
	cChainID  uint64
	cProvider cchain.Provider
}

func NewMock(period time.Duration, cChainID uint64, cProvider cchain.Provider) *Mock {
	return &Mock{
		period:    period,
		uniq:      make(map[uint64]bool),
		cChainID:  cChainID,
		cProvider: cProvider,
	}
}

func (m *Mock) StreamAsync(ctx context.Context, chainID uint64, fromHeight uint64, fromOffset uint64, callback xchain.ProviderCallback) error {
	go func() {
		err := m.stream(ctx, chainID, fromHeight, fromOffset, callback, true)
		if err != nil {
			log.Error(ctx, "Unexpected stream error [BUG]", err)
		}
	}()

	return nil
}

func (m *Mock) StreamBlocks(ctx context.Context, chainID uint64, fromHeight uint64, fromOffset uint64, callback xchain.ProviderCallback) error {
	return m.stream(ctx, chainID, fromHeight, fromOffset, callback, false)
}

//nolint:nilerr // Stream function contract states it returns nil on context error.
func (m *Mock) stream(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	fromOffset uint64,
	callback xchain.ProviderCallback,
	retryCallback bool,
) error {
	sOffset := make(streamOffseter).offset

	if fromOffset == 0 {
		fromOffset = initialXOffset
	}
	bOffset := newBlockOffseter(fromOffset)

	// Similarly to real xprovider, we bump fromHeight to netconf.DeployHeight if below,
	// this is only required for consensus chain in the mock.
	if chainID == m.cChainID {
		if fromHeight < 1 {
			fromHeight = 1
		}
	} else {
		// Populate historical blocks for mocked chains so offsets are consistent for heights.
		for i := uint64(0); i < fromHeight; i++ {
			m.addBlock(m.nextBlock(ctx, chainID, i, sOffset, bOffset.getAndInc))
		}
	}

	defer func() {
		log.Debug(ctx, "Mock subscription ended")
	}()
	height := fromHeight

	for ctx.Err() == nil {
		block := m.nextBlock(ctx, chainID, height, sOffset, bOffset.getAndInc)
		m.addBlock(block)

		err := callback(ctx, block)
		if ctx.Err() != nil {
			return nil
		} else if err != nil {
			if !retryCallback {
				return err
			}
			log.Warn(ctx, "Mock callback failed (will retry)", err)

			continue
		}
		height++

		// Backoff before providing next block
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(m.period):
		}
	}

	return nil
}

func (m *Mock) GetBlock(_ context.Context, chainID uint64, height uint64, offset uint64) (xchain.Block, bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, block := range m.blocks {
		if block.BlockOffset == offset && block.BlockHeight == height && block.SourceChainID == chainID {
			return block, true, nil
		}
	}

	return xchain.Block{}, false, nil
}

func (*Mock) GetSubmittedCursor(_ context.Context, destChain uint64, srcChain uint64,
) (xchain.StreamCursor, bool, error) {
	return xchain.StreamCursor{StreamID: xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChain,
	}}, true, nil
}

func (*Mock) GetEmittedCursor(_ context.Context, _ xchain.EmitRef, srcChainID uint64, destChainID uint64,
) (xchain.StreamCursor, bool, error) {
	return xchain.StreamCursor{StreamID: xchain.StreamID{
		SourceChainID: srcChainID,
		DestChainID:   destChainID,
	}}, true, nil
}

func (m *Mock) addBlock(block xchain.Block) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.uniq[block.BlockOffset] {
		return
	}
	m.blocks = append(m.blocks, block)
	m.uniq[block.BlockOffset] = true
}

func (m *Mock) nextBlock(
	ctx context.Context,
	chainID uint64,
	height uint64,
	sOffsetFunc func(xchain.StreamID) uint64,
	bOffsetFunc func() uint64,
) xchain.Block {
	if m.cChainID == chainID {
		// For omni consensus chain, we query the real cprovider for blocks.
		for {
			b, ok, err := m.cProvider.XBlock(ctx, height, false)
			if ctx.Err() != nil {
				return xchain.Block{}
			} else if err != nil {
				panic(err)
			} else if !ok {
				time.Sleep(m.period / 3)
				continue
			}

			return b
		}
	}

	// Use deterministic randomness based on the chainID and height.
	r := rand.New(rand.NewSource(int64(chainID ^ height)))

	// TODO(corver): add xreceipts
	var msgs []xchain.Msg

	newMsgA := func() xchain.Msg {
		return newMsg(r, chainID, destChainA, sOffsetFunc)
	}
	newMsgB := func() xchain.Msg {
		return newMsg(r, chainID, destChainB, sOffsetFunc)
	}

	switch height % 4 {
	case 0:
		// Empty block, no messages or receipts
	case 1:
		msgs = append(msgs, newMsgA()) // Msgs: 1*chainA, 0*chainB
	case 2:
		msgs = append(msgs, newMsgA(), newMsgB()) // Msgs: 2*chainA, 1*chainB
	case 3:
		msgs = append(msgs, newMsgA(), newMsgA(), newMsgB()) // Msgs: 3*chainA, 1*chainB
	}

	b := xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
			BlockHash:     random32(r),
		},
		Msgs:      msgs,
		Receipts:  nil,        // TODO(corver): Add receipts
		Timestamp: time.Now(), // Should this also be deterministic?
	}

	if b.ShouldAttest() {
		b.BlockHeader.BlockOffset = bOffsetFunc()
	}

	return b
}

func (*Mock) StreamAsyncNoOffset(context.Context, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func (*Mock) StreamBlocksNoOffset(context.Context, uint64, uint64, xchain.ProviderCallback) error {
	panic("unexpected")
}

func newMsg(r *rand.Rand, srcChain, destChain uint64, offsetFunc func(xchain.StreamID) uint64) xchain.Msg {
	streamID := xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChain,
	}

	return xchain.Msg{
		MsgID: xchain.MsgID{
			StreamID:     streamID,
			StreamOffset: offsetFunc(streamID),
		},
		SourceMsgSender: random20(r),
		DestAddress:     random20(r),
		DestGasLimit:    r.Uint64(),
	}
}

// random32 returns a deterministic random 32 array on the seed.
func random32(r io.Reader) [32]byte {
	var resp [32]byte
	_, _ = r.Read(resp[:])

	return resp
}

// random20 returns a deterministic random 20 array on the seed.
func random20(r io.Reader) [20]byte {
	var resp [20]byte
	_, _ = r.Read(resp[:])

	return resp
}

type streamOffseter map[xchain.StreamID]uint64

func (o streamOffseter) offset(id xchain.StreamID) uint64 {
	defer func() { o[id]++ }()
	return o[id]
}

type blockOffseter uint64

func newBlockOffseter(from uint64) *blockOffseter {
	return (*blockOffseter)(&from)
}

func (o *blockOffseter) getAndInc() uint64 {
	defer func() { *o++ }()
	return uint64(*o)
}
