//nolint:gosec // Use weak random for determinism.
package provider

import (
	"context"
	"io"
	"math/rand"
	"time"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

var _ xchain.Provider = (*Mock)(nil)

const (
	destChainA = 100
	destChainB = 200
)

// Mock is a mock implementation of the xchain.Provider interface.
// Its zero value is ready for use. It returns semi-deterministic blocks.
type Mock struct{}

func (Mock) Subscribe(ctx context.Context, chainID uint64, fromHeight uint64, callback xchain.ProviderCallback,
) error {
	go func() {
		height := fromHeight
		for ctx.Err() == nil {
			if err := callback(ctx, nextBlock(chainID, height)); err != nil {
				log.Warn(ctx, "Mock callback failed, will retry", err)
				continue
			}
			height++
		}
	}()

	return nil
}

func nextBlock(chainID uint64, height uint64) *xchain.Block {
	r := rand.New(rand.NewSource(int64(chainID ^ height)))

	var (
		msgs   []xchain.Msg
		offset = make(offseter).offset
	)

	newMsgA := func() xchain.Msg {
		return newMsg(r, chainID, destChainA, offset)
	}
	newMsgB := func() xchain.Msg {
		return newMsg(r, chainID, destChainB, offset)
	}

	switch height % 4 {
	case 0:
		// Empty block, no messages or receipts
	case 1:
		msgs = append(msgs, newMsgA()) // One message to chain A
	case 2:
		msgs = append(msgs, newMsgB()) // One message to chain B
	case 3:
		msgs = append(msgs, newMsgA(), newMsgB()) // Two messages, one to chain A and one to chain B
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
		TxHash:          random32(r),
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
