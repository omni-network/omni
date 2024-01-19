package engine

import (
	"context"
	"crypto/sha256"
	"math/big"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	fuzz "github.com/google/gofuzz"
)

var _ API = (*Mock)(nil)

// Mock mocks the Engine API for testing purposes.
type Mock struct {
	fuzzer *fuzz.Fuzzer

	mu       sync.Mutex
	head     *types.Block
	payloads map[engine.PayloadID]engine.ExecutableData
}

// NewMock returns a new mock Engine API.
func NewMock() (*Mock, error) {
	fuzzer := NewFuzzer()

	genesisPayload, err := makeNextPayload(nil, fuzzer)
	if err != nil {
		return nil, errors.Wrap(err, "make next payload")
	}
	genesisBlock, err := engine.ExecutableDataToBlock(genesisPayload, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "executable data to block")
	}

	return &Mock{
		fuzzer:   fuzzer,
		head:     genesisBlock,
		payloads: make(map[engine.PayloadID]engine.ExecutableData),
	}, nil
}

func (m *Mock) BlockNumber(_ context.Context) (uint64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.head.NumberU64(), nil
}

func (m *Mock) BlockByNumber(_ context.Context, number *big.Int) (*types.Block, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if number.Cmp(m.head.Number()) != 0 {
		return nil, errors.New("block not found") // Only support latest block
	}

	return m.head, nil
}

func (m *Mock) NewPayloadV2(_ context.Context, params engine.ExecutableData) (engine.PayloadStatusV1, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, err := payloadID(params)
	if err != nil {
		return engine.PayloadStatusV1{}, err
	}

	m.payloads[id] = params

	return engine.PayloadStatusV1{
		Status: engine.VALID,
	}, nil
}

func (m *Mock) ForkchoiceUpdatedV2(_ context.Context, update engine.ForkchoiceStateV1,
	payloadAttributes *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	resp := engine.ForkChoiceResponse{
		PayloadStatus: engine.PayloadStatusV1{
			Status: engine.VALID,
		},
	}

	// If we have payload attributes, make a new payload
	if payloadAttributes != nil {
		payload, err := makeNextPayload(m.head, m.fuzzer)
		if err != nil {
			return engine.ForkChoiceResponse{}, err
		}

		id, err := payloadID(payload)
		if err != nil {
			return engine.ForkChoiceResponse{}, err
		}

		m.payloads[id] = payload

		resp.PayloadID = &id
	}

	if m.head.Hash() == update.HeadBlockHash {
		// Head is already up to date
		return resp, nil
	}

	// Update head
	var found bool
	for _, payload := range m.payloads {
		block, err := engine.ExecutableDataToBlock(payload, nil, nil)
		if err != nil {
			return engine.ForkChoiceResponse{}, errors.Wrap(err, "executable data to block")
		}

		if block.Hash() != update.HeadBlockHash {
			continue
		}

		if !isChild(m.head, block) {
			return engine.ForkChoiceResponse{}, errors.New("unexpected new head block")
		}

		m.head = block
		found = true

		break
	}
	if !found {
		return engine.ForkChoiceResponse{}, errors.New("head block not found")
	}

	return resp, nil
}

func (m *Mock) GetPayloadV2(_ context.Context, payloadID engine.PayloadID) (*engine.ExecutionPayloadEnvelope, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	payload, ok := m.payloads[payloadID]
	if !ok {
		return nil, errors.New("payload not found")
	}

	return &engine.ExecutionPayloadEnvelope{
		ExecutionPayload: &payload,
	}, nil
}

// TODO(corver): Add support for V3

func (*Mock) NewPayloadV3(context.Context, engine.ExecutableData, []common.Hash, *common.Hash,
) (engine.PayloadStatusV1, error) {
	panic("implement me")
}

func (*Mock) ForkchoiceUpdatedV3(context.Context, engine.ForkchoiceStateV1, *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	panic("implement me")
}

func (*Mock) GetPayloadV3(context.Context, engine.PayloadID) (*engine.ExecutionPayloadEnvelope, error) {
	panic("implement me")
}

// makeNextPayload returns a new fuzzed payload using head as parent if provided.
func makeNextPayload(head *types.Block, fuzzer *fuzz.Fuzzer) (engine.ExecutableData, error) {
	// Set some deterministic genesis fields
	height := uint64(1)
	timestamp := uint64(time.Now().Unix())
	var parentHash common.Hash

	// If we have a head, use it as parent
	if head != nil {
		height = head.NumberU64() + 1
		timestamp = head.Time() + 1
		parentHash = head.Hash()
	}

	// Build a new header
	var header types.Header
	fuzzer.Fuzz(&header)
	header.Number = big.NewInt(int64(height))
	header.Time = timestamp
	header.ParentHash = parentHash

	// Convert header to block
	block := types.NewBlock(&header, nil, nil, nil, trie.NewStackTrie(nil))

	// Convert block to payload
	env := engine.BlockToExecutableData(block, big.NewInt(0), nil)
	payload := *env.ExecutionPayload

	// Ensure the block is valid
	_, err := engine.ExecutableDataToBlock(payload, nil, nil)
	if err != nil {
		return engine.ExecutableData{}, errors.Wrap(err, "executable data to block")
	}

	return payload, nil
}

// payloadID returns a deterministic payload id for the given payload.
func payloadID(payload engine.ExecutableData) (engine.PayloadID, error) {
	bz, err := payload.MarshalJSON()
	if err != nil {
		return engine.PayloadID{}, errors.Wrap(err, "marshal payload")
	}

	hash := sha256.Sum256(bz)

	return engine.PayloadID(hash[:8]), nil
}

func isChild(parent *types.Block, child *types.Block) bool {
	return parent.NumberU64() == child.NumberU64()-1 && parent.Hash() == child.ParentHash()
}
