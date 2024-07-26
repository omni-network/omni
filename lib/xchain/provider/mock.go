//nolint:gosec // Use weak random for determinism.
package provider

import (
	"context"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

var (
	_ xchain.Provider = (*Mock)(nil)
)

// Mock is a mock implementation of the xchain.Provider interface.
// It generates deterministic blocks and messages for any chain that is queried.
// Except for omni consensus chain, we use the real cprovider to fetch blocks.
type Mock struct {
	period     time.Duration
	mu         sync.Mutex
	blocks     map[blockKey]xchain.Block
	cChainID   uint64
	cProvider  cchain.Provider
	destChains [2]uint64
}

func NewMock(period time.Duration, cChainID uint64, cProvider cchain.Provider) (*Mock, error) {
	// Mock provider produces xmsgs to simnet non-omni evm chains.
	var destChains []uint64
	for _, chain := range netconf.SimnetNetwork().EVMChains() {
		if chain.ID != netconf.Simnet.Static().OmniExecutionChainID {
			destChains = append(destChains, chain.ID)
		}
	}
	if len(destChains) != 2 {
		return nil, errors.New("mock provider requires exactly 2 destination chains")
	}

	return &Mock{
		period:     period,
		blocks:     make(map[blockKey]xchain.Block),
		cChainID:   cChainID,
		cProvider:  cProvider,
		destChains: [2]uint64(destChains),
	}, nil
}

func (m *Mock) StreamAsync(ctx context.Context, req xchain.ProviderRequest, callback xchain.ProviderCallback) error {
	go func() {
		err := m.stream(ctx, req, callback, true)
		if err != nil {
			log.Error(ctx, "Unexpected stream error [BUG]", err)
		}
	}()

	return nil
}

func (m *Mock) StreamBlocks(ctx context.Context, req xchain.ProviderRequest, callback xchain.ProviderCallback) error {
	return m.stream(ctx, req, callback, false)
}

func (*Mock) ChainVersionHeight(context.Context, xchain.ChainVersion) (uint64, error) {
	return 0, errors.New("unsupported")
}

func (m *Mock) stream(
	ctx context.Context,
	req xchain.ProviderRequest,
	callback xchain.ProviderCallback,
	retryCallback bool,
) error {
	chainVer := xchain.ChainVersion{ID: req.ChainID, ConfLevel: req.ConfLevel}

	sOffset := make(streamOffseter).offset

	// Similarly as the real xprovider, we bump fromHeight to netconf.DeployHeight if below,
	// this is only required for consensus chain in the mock.
	fromHeight := req.Height
	if req.ChainID == m.cChainID {
		if fromHeight < 1 {
			fromHeight = 1
		}
	} else {
		// Populate historical blocks for mocked chains so offsets are consistent for heights.
		for i := uint64(0); i < fromHeight; i++ {
			m.addBlock(m.nextBlock(ctx, chainVer, i, sOffset), req.ConfLevel)
		}
	}

	defer func() {
		log.Debug(ctx, "Mock subscription ended")
	}()
	height := fromHeight

	for ctx.Err() == nil {
		block := m.nextBlock(ctx, chainVer, height, sOffset)
		m.addBlock(block, req.ConfLevel)

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

func (m *Mock) GetBlock(_ context.Context, req xchain.ProviderRequest) (xchain.Block, bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	block, ok := m.blocks[blockKey{
		ChainID:   req.ChainID,
		Height:    req.Height,
		ConfLevel: req.ConfLevel,
	}]

	return block, ok, nil
}

func (*Mock) GetSubmittedCursor(_ context.Context, stream xchain.StreamID,
) (xchain.SubmitCursor, bool, error) {
	return xchain.SubmitCursor{StreamID: stream}, true, nil
}

func (*Mock) GetEmittedCursor(_ context.Context, _ xchain.EmitRef, stream xchain.StreamID,
) (xchain.EmitCursor, bool, error) {
	return xchain.EmitCursor{StreamID: stream}, true, nil
}

func (m *Mock) parentBlockHash(chainVer xchain.ChainVersion, height uint64) common.Hash {
	m.mu.Lock()
	defer m.mu.Unlock()

	parentHeight, ok := umath.Subtract(height, 1)
	if !ok {
		return common.Hash{} // Height == 0
	}

	key := blockKey{
		ChainID:   chainVer.ID,
		Height:    parentHeight,
		ConfLevel: chainVer.ConfLevel,
	}

	return m.blocks[key].BlockHash
}

func (m *Mock) addBlock(block xchain.Block, confLevel xchain.ConfLevel) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := blockKey{
		ChainID:   block.ChainID,
		Height:    block.BlockHeight,
		ConfLevel: confLevel,
	}
	m.blocks[key] = block
}

func (m *Mock) nextBlock(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	height uint64,
	sOffsetFunc func(xchain.StreamID) uint64,
) xchain.Block {
	if m.cChainID == chainVer.ID {
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
	r := rand.New(rand.NewSource(int64(chainVer.ID ^ height)))

	// TODO(corver): add xreceipts
	var msgs []xchain.Msg

	newMsgA := func() xchain.Msg {
		return newMsg(r, chainVer.ID, m.destChains[0], sOffsetFunc)
	}
	newMsgB := func() xchain.Msg {
		return newMsg(r, chainVer.ID, m.destChains[1], sOffsetFunc)
	}

	switch height % 4 {
	case 0:
		// Empty block, no messages or receipts, no attestation
	case 1:
		msgs = append(msgs, newMsgA()) // Msgs: 1*chainA, 0*chainB
	case 2:
		msgs = append(msgs, newMsgA(), newMsgB()) // Msgs: 2*chainA, 1*chainB
	case 3:
		msgs = append(msgs, newMsgA(), newMsgA(), newMsgB()) // Msgs: 3*chainA, 1*chainB
	}

	return xchain.Block{
		BlockHeader: xchain.BlockHeader{
			ChainID:     chainVer.ID,
			BlockHeight: height,
			BlockHash:   random32(r),
		},
		Msgs:       msgs,
		Receipts:   nil,        // TODO(corver): Add receipts
		Timestamp:  time.Now(), // Should this also be deterministic?
		ParentHash: m.parentBlockHash(chainVer, height),
	}
}

func newMsg(r *rand.Rand, srcChain, destChain uint64, offsetFunc func(xchain.StreamID) uint64) xchain.Msg {
	streamID := xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChain,
		ShardID:       xchain.ShardFinalized0,
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

type blockKey struct {
	ChainID   uint64
	Height    uint64
	ConfLevel xchain.ConfLevel
}
