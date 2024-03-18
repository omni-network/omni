package ethclient

import (
	"context"
	"crypto/sha256"
	"math/big"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	fuzz "github.com/google/gofuzz"
)

// engineMock mocks the Engine API for testing purposes.
type engineMock struct {
	Client
	fuzzer *fuzz.Fuzzer

	mu       sync.Mutex
	head     *types.Block
	payloads map[engine.PayloadID]engine.ExecutableData
}

// NewEngineMock returns a new mock engine API client.
//
// Note only some methods are implemented, it will panic if you call an unimplemented method.
func NewEngineMock() (EngineClient, error) {
	var (
		// Deterministic genesis block
		height     uint64 // 0
		parentHash common.Hash
		timestamp  = time.Now().Truncate(time.Hour * 24).Unix() // TODO(corver): Improve this.

		// Deterministic fuzzer
		fuzzer = NewFuzzer(timestamp)
	)

	genesisPayload, err := makePayload(fuzzer, height, uint64(timestamp), parentHash, common.Address{}, parentHash)
	if err != nil {
		return nil, errors.Wrap(err, "make next payload")
	}
	genesisBlock, err := engine.ExecutableDataToBlock(genesisPayload, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "executable data to block")
	}

	return &engineMock{
		fuzzer:   fuzzer,
		head:     genesisBlock,
		payloads: make(map[engine.PayloadID]engine.ExecutableData),
	}, nil
}

func (*engineMock) FilterLogs(_ context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	if len(q.Addresses) > 0 {
		return nil, nil // We can't mock contract specific logs
	}

	// If no addresses are provided, we return two random logs
	f := fuzz.NewWithSeed(int64(q.BlockHash[0])).NilChance(0).NumElements(1, 4)
	var resp1, resp2 types.Log
	f.Fuzz(&resp1)
	f.Fuzz(&resp2)

	return []types.Log{resp1, resp2}, nil
}

func (m *engineMock) BlockNumber(_ context.Context) (uint64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.head.NumberU64(), nil
}

func (m *engineMock) BlockByNumber(_ context.Context, number *big.Int) (*types.Block, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if number.Cmp(m.head.Number()) != 0 {
		return nil, errors.New("block not found") // Only support latest block
	}

	return m.head, nil
}

func (m *engineMock) NewPayloadV2(ctx context.Context, params engine.ExecutableData) (engine.PayloadStatusV1, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, err := payloadID(params)
	if err != nil {
		return engine.PayloadStatusV1{}, err
	}

	m.payloads[id] = params

	log.Debug(ctx, "Engine mock received new payload from proposer",
		"height", params.Number,
		log.Hex7("hash", params.BlockHash.Bytes()),
	)

	return engine.PayloadStatusV1{
		Status: engine.VALID,
	}, nil
}

func (m *engineMock) ForkchoiceUpdatedV2(ctx context.Context, update engine.ForkchoiceStateV1,
	attrs *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	resp := engine.ForkChoiceResponse{
		PayloadStatus: engine.PayloadStatusV1{
			Status: engine.VALID,
		},
	}

	// Maybe update head
	//nolint: nestif // this is a mock it's fine
	if m.head.Hash() != update.HeadBlockHash {
		var found bool
		for _, payload := range m.payloads {
			block, err := engine.ExecutableDataToBlock(payload, nil, nil)
			if err != nil {
				return engine.ForkChoiceResponse{}, errors.Wrap(err, "executable data to block")
			}

			if block.Hash() != update.HeadBlockHash {
				continue
			}

			if err := verifyChild(m.head, block); err != nil {
				return engine.ForkChoiceResponse{}, err
			}

			m.head = block
			found = true

			id, err := payloadID(payload)
			if err != nil {
				return engine.ForkChoiceResponse{}, err
			}
			resp.PayloadID = &id

			break
		}
		if !found {
			return engine.ForkChoiceResponse{}, errors.New("forkchoice block not found",
				log.Hex7("forkchoice", m.head.Hash().Bytes()))
		}
	}

	// If we have payload attributes, make a new payload
	if attrs != nil {
		payload, err := makePayload(m.fuzzer, m.head.NumberU64()+1,
			attrs.Timestamp, update.HeadBlockHash, attrs.SuggestedFeeRecipient, attrs.Random)
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

	log.Debug(ctx, "Engine mock forkchoice updated",
		"height", m.head.NumberU64(),
		log.Hex7("forkchoice", update.HeadBlockHash.Bytes()),
	)

	return resp, nil
}

func (m *engineMock) GetPayloadV2(_ context.Context, payloadID engine.PayloadID) (*engine.ExecutionPayloadEnvelope, error) {
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

func (*engineMock) NewPayloadV3(context.Context, engine.ExecutableData, []common.Hash, *common.Hash,
) (engine.PayloadStatusV1, error) {
	panic("implement me")
}

func (*engineMock) ForkchoiceUpdatedV3(context.Context, engine.ForkchoiceStateV1, *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	panic("implement me")
}

func (*engineMock) GetPayloadV3(context.Context, engine.PayloadID) (*engine.ExecutionPayloadEnvelope, error) {
	panic("implement me")
}

// makePayload returns a new fuzzed payload using head as parent if provided.
func makePayload(fuzzer *fuzz.Fuzzer, height uint64, timestamp uint64, parentHash common.Hash,
	feeRecipient common.Address, randao common.Hash) (engine.ExecutableData, error) {
	// Build a new header
	var header types.Header
	fuzzer.Fuzz(&header)
	header.Number = big.NewInt(int64(height))
	header.Time = timestamp
	header.ParentHash = parentHash
	header.MixDigest = randao      // this corresponds to Random field in PayloadAttributes
	header.Coinbase = feeRecipient // this corresponds to SuggestedFeeRecipient field in PayloadAttributes

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

// verifyChild returns an error if child is not a valid child of parent.
func verifyChild(parent *types.Block, child *types.Block) error {
	if parent.NumberU64()+1 != child.NumberU64() {
		return errors.New("forkchoice height not following head",
			"head", parent.NumberU64(),
			"forkchoice", child.NumberU64(),
		)
	}

	if parent.Hash() != child.ParentHash() {
		return errors.New("forkchoice parent hash not head",
			log.Hex7("head", parent.Hash().Bytes()),
			log.Hex7("forkchoice", child.Hash().Bytes()),
		)
	}

	return nil
}
