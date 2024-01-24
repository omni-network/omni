//nolint:gosec // Use weak random for determinism.
package provider

import (
	"context"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

var (
	_ xchain.Provider = (*Mock)(nil)
)

const (
	destChainA = 100
	destChainB = 200
)

// Mock is a mock implementation of the xchain.Provider interface as well as the relayer.XChainClient.
type Mock struct {
	period time.Duration
	mu     sync.Mutex
	blocks []xchain.Block
}

func NewMock(period time.Duration) *Mock {
	return &Mock{
		period: period,
	}
}

func (m *Mock) Subscribe(ctx context.Context, chainID uint64, fromHeight uint64, callback xchain.ProviderCallback,
) error {
	offset := make(offseter).offset

	// Populate historical blocks so offsets are consistent for heights.
	for i := uint64(0); i < fromHeight; i++ {
		m.addBlock(*nextBlock(chainID, i, offset))
	}

	go func() {
		height := fromHeight

		for ctx.Err() == nil {
			block := nextBlock(chainID, height, offset)
			m.addBlock(*block)
			if err := callback(ctx, block); err != nil {
				log.Warn(ctx, "Mock callback failed, will retry", err)
				continue
			}
			height++

			// Backoff before providing next block
			select {
			case <-ctx.Done():
				return
			case <-time.After(m.period):
			}
		}
	}()

	return nil
}

func (m *Mock) GetBlock(_ context.Context, chainID uint64, height uint64) (xchain.Block, bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, block := range m.blocks {
		if block.BlockHeight == height && block.SourceChainID == chainID {
			return block, true, nil
		}
	}

	return xchain.Block{}, false, nil
}

func (*Mock) GetSubmittedCursor(_ context.Context, destChain uint64, srcChain uint64) (xchain.StreamCursor, error) {
	return xchain.StreamCursor{StreamID: xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChain,
	}}, nil
}

func (m *Mock) addBlock(block xchain.Block) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.blocks = append(m.blocks, block)
}

func nextBlock(chainID uint64, height uint64, offsetFunc func(xchain.StreamID) uint64) *xchain.Block {
	// Use deterministic randomness based on the chainID and height.
	r := rand.New(rand.NewSource(int64(chainID ^ height)))

	// TODO(corver): add xreceipts
	var msgs []xchain.Msg

	newMsgA := func() xchain.Msg {
		return newMsg(r, chainID, destChainA, offsetFunc)
	}
	newMsgB := func() xchain.Msg {
		return newMsg(r, chainID, destChainB, offsetFunc)
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

	return &xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
			BlockHash:     random32(r),
		},
		Msgs:      msgs,
		Receipts:  nil,        // TODO(corver): Add receipts
		Timestamp: time.Now(), // Should this also be deterministic?
	}
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

type offseter map[xchain.StreamID]uint64

func (o offseter) offset(id xchain.StreamID) uint64 {
	defer func() { o[id]++ }()
	return o[id]
}
