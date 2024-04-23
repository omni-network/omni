package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	attesttypes "github.com/omni-network/omni/halo/attest/types"
	etypes "github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"

	"github.com/ethereum/go-ethereum"
	eengine "github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	"cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	cosmosstd "github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	atypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	btypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/gogoproto/proto"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

var zeroAddr common.Address

func TestKeeper_PrepareProposal(t *testing.T) {
	t.Parallel()

	// TestRunErrScenarios tests various error scenarios in the PrepareProposal function.
	// It covers cases where different errors are encountered during the preparation of a proposal,
	// such as when no transactions are provided, when errors occur while fetching block information,
	// or when errors occur during fork choice update.
	t.Run("TestRunErrScenarios", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name       string
			mockEngine mockEngineAPI
			req        *abci.RequestPrepareProposal
			wantErr    bool
		}{
			{
				name:       "no transactions",
				mockEngine: mockEngineAPI{},
				req: &abci.RequestPrepareProposal{
					Txs:    nil,        // Set to nil to simulate no transactions
					Height: 1,          // Set height to 1 for this test case
					Time:   time.Now(), // Set time to current time or mock a time
				},
				wantErr: false,
			},
			{
				name:       "with  transactions",
				mockEngine: mockEngineAPI{},
				req: &abci.RequestPrepareProposal{
					Txs:    [][]byte{[]byte("tx1")}, // simulate transactions
					Height: 1,
					Time:   time.Now(),
				},
				wantErr: true,
			},
			{
				name: "block number err",
				mockEngine: mockEngineAPI{
					blockNumberFunc: func(ctx context.Context) (uint64, error) {
						return 0, errors.New("mocked error")
					},
				},
				req: &abci.RequestPrepareProposal{
					Txs:    nil,
					Height: 2,
					Time:   time.Now(),
				},
				wantErr: true,
			},
			{
				name: "block by number err",
				mockEngine: mockEngineAPI{
					blockNumberFunc: func(ctx context.Context) (uint64, error) {
						return 0, nil
					},
					blockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
						return nil, errors.New("mocked error")
					},
				},
				req: &abci.RequestPrepareProposal{
					Txs:    nil,
					Height: 2,
					Time:   time.Now(),
				},
				wantErr: true,
			},
			{
				name: "forkchoiceUpdateV2  err",
				mockEngine: mockEngineAPI{
					blockNumberFunc: func(ctx context.Context) (uint64, error) {
						return 0, nil
					},
					blockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
						fuzzer := ethclient.NewFuzzer(0)
						var block *types.Block
						fuzzer.Fuzz(&block)

						return block, nil
					},
					forkchoiceUpdatedV3Func: func(ctx context.Context, update eengine.ForkchoiceStateV1,
						payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
						return eengine.ForkChoiceResponse{}, errors.New("mocked error")
					},
				},
				req: &abci.RequestPrepareProposal{
					Txs:    nil,
					Height: 2,
					Time:   time.Now(),
				},
				wantErr: true,
			},
			{
				name: "forkchoiceUpdateV2  not valid",
				mockEngine: mockEngineAPI{
					blockNumberFunc: func(ctx context.Context) (uint64, error) {
						return 0, nil
					},
					blockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
						fuzzer := ethclient.NewFuzzer(0)
						var block *types.Block
						fuzzer.Fuzz(&block)

						return block, nil
					},
					forkchoiceUpdatedV3Func: func(ctx context.Context, update eengine.ForkchoiceStateV1,
						payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
						return eengine.ForkChoiceResponse{
							PayloadStatus: eengine.PayloadStatusV1{
								Status:          eengine.INVALID,
								LatestValidHash: nil,
								ValidationError: nil,
							},
							PayloadID: nil,
						}, nil
					},
				},
				req: &abci.RequestPrepareProposal{
					Txs:    nil,
					Height: 2,
					Time:   time.Now(),
				},
				wantErr: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				ctx, storeService := setupCtxStore(t, nil)
				cdc := getCodec(t)
				txConfig := authtx.NewTxConfig(cdc, nil)

				k := NewKeeper(cdc, storeService, &tt.mockEngine, txConfig)
				k.SetAddressProvider(mockAddressProvider{
					address: common.BytesToAddress([]byte("test")),
				})
				_, err := k.PrepareProposal(ctx, tt.req)
				if (err != nil) != tt.wantErr {
					t.Errorf("PrepareProposal() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			})
		}
	})

	t.Run("TestBuildOptimistic", func(t *testing.T) {
		t.Parallel()
		// setup dependencies
		ctx, storeService := setupCtxStore(t, nil)
		cdc := getCodec(t)
		txConfig := authtx.NewTxConfig(cdc, nil)
		mockEngine, err := newMockEngineAPI(0)
		require.NoError(t, err)

		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig)
		ap := mockAddressProvider{
			address: common.BytesToAddress([]byte("test")),
		}
		keeper.SetAddressProvider(ap)
		keeper.SetVoteProvider(mockVEProvider{})
		keeper.AddEventProcessor(mockLogProvider{})

		// get the genesis block to build on top of
		ts := time.Now()
		latestHeight, err := mockEngine.BlockNumber(ctx)
		require.NoError(t, err)
		latestBlock, err := mockEngine.BlockByNumber(ctx, big.NewInt(int64(latestHeight)))
		require.NoError(t, err)

		// build next two blocks and get the PayloadID of the second
		mockEngine.pushPayload(t, ctx, ap.LocalAddress(), latestBlock.Hash(), ts)
		nextBlock, blockPayload := mockEngine.nextBlock(t, latestHeight+1, uint64(ts.Unix()), latestBlock.Hash(), ap.LocalAddress())
		_, err = mockEngine.mock.NewPayloadV3(ctx, blockPayload, nil, nil)
		require.NoError(t, err)
		payloadID := mockEngine.pushPayload(t, ctx, ap.LocalAddress(), nextBlock.Hash(), ts)

		req := &abci.RequestPrepareProposal{
			Txs:    nil,
			Height: int64(2),
			Time:   time.Now(),
		}

		// initialize mutable payload so we trigger the optimistic flow
		keeper.mutablePayload.Height = 2
		keeper.mutablePayload.UpdatedAt = time.Now()
		keeper.mutablePayload.ID = payloadID

		resp, err := keeper.PrepareProposal(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		// decode the txn and get the messages
		tx, err := txConfig.TxDecoder()(resp.Txs[0])
		require.NoError(t, err)

		for _, msg := range tx.GetMsgs() {
			if _, ok := msg.(*etypes.MsgExecutionPayload); ok {
				assertExecutablePayload(t, msg, ts.Unix(), nextBlock.Hash(), ap, uint64(req.Height))
			}
		}
	})

	t.Run("TestBuildNonOptimistic", func(t *testing.T) {
		t.Parallel()
		// setup dependencies
		ctx, storeService := setupCtxStore(t, nil)
		cdc := getCodec(t)
		txConfig := authtx.NewTxConfig(cdc, nil)

		mockEngine, err := newMockEngineAPI(0)
		require.NoError(t, err)

		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig)
		ap := mockAddressProvider{
			address: common.BytesToAddress([]byte("test")),
		}
		keeper.SetAddressProvider(ap)

		keeper.AddEventProcessor(mockLogProvider{})
		keeper.SetVoteProvider(mockVEProvider{})

		// get the genesis block to build on top of
		ts := time.Now()
		latestHeight, err := mockEngine.BlockNumber(ctx)
		require.NoError(t, err)
		latestBlock, err := mockEngine.BlockByNumber(ctx, big.NewInt(int64(latestHeight)))
		require.NoError(t, err)

		// build next two blocks and get the PayloadID of the second
		mockEngine.pushPayload(t, ctx, ap.LocalAddress(), latestBlock.Hash(), ts)
		nextBlock, blockPayload := mockEngine.nextBlock(t, latestHeight+1, uint64(ts.Unix()), latestBlock.Hash(), ap.LocalAddress())
		_, err = mockEngine.NewPayloadV3(ctx, blockPayload, nil, nil)
		require.NoError(t, err)
		mockEngine.pushPayload(t, ctx, ap.LocalAddress(), nextBlock.Hash(), ts)

		keeper.mutablePayload.UpdatedAt = time.Now()

		req := &abci.RequestPrepareProposal{
			Txs:    nil,
			Height: int64(2),
			Time:   time.Now(),
		}

		resp, err := keeper.PrepareProposal(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		// decode the txn and get the messages
		tx, err := txConfig.TxDecoder()(resp.Txs[0])
		require.NoError(t, err)

		var actualDelCount int
		// assert that the message is an executable payload
		for _, msg := range tx.GetMsgs() {
			if _, ok := msg.(*etypes.MsgExecutionPayload); ok {
				assertExecutablePayload(t, msg, req.Time.Unix()+1, nextBlock.Hash(), ap, uint64(req.Height))
			}
			if msgDelegate, ok := msg.(*stypes.MsgDelegate); ok {
				require.Equal(t, msgDelegate.Amount, sdk.NewInt64Coin("stake", 100))
				actualDelCount++
			}
		}
		// make sure all msg.Delegate are present
		require.Equal(t, 1, actualDelCount)
	})
}

// assertExecutablePayload asserts that the given message is an executable payload with the expected values.
func assertExecutablePayload(t *testing.T, msg sdk.Msg, ts int64, blockHash common.Hash, ap mockAddressProvider, height uint64) {
	t.Helper()
	executionPayload, ok := msg.(*etypes.MsgExecutionPayload)
	require.True(t, ok)
	require.NotNil(t, executionPayload)

	payload := new(eengine.ExecutableData)
	err := json.Unmarshal(executionPayload.GetExecutionPayload(), payload)
	require.NoError(t, err)
	require.Equal(t, int64(payload.Timestamp), ts)
	require.Equal(t, payload.Random, blockHash)
	require.Equal(t, payload.FeeRecipient, ap.LocalAddress())
	require.Empty(t, payload.Withdrawals)
	require.Equal(t, payload.Number, height)

	require.Len(t, executionPayload.PrevPayloadEvents, 1)
	evmLog := executionPayload.PrevPayloadEvents[0]
	require.Equal(t, evmLog.Address, zeroAddr.Bytes())
}

func setupCtxStore(t *testing.T, header *cmtproto.Header) (sdk.Context, store.KVStoreService) {
	t.Helper()
	key := storetypes.NewKVStoreKey("test")
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(t, key, storetypes.NewTransientStoreKey("transient_test"))
	if header == nil {
		header = &cmtproto.Header{Time: cmttime.Now()}
	}
	ctx := testCtx.Ctx.WithBlockHeader(*header)

	return ctx, storeService
}

func getCodec(t *testing.T) codec.Codec {
	t.Helper()
	sdkConfig := sdk.GetConfig()
	reg, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          authcodec.NewBech32Codec(sdkConfig.GetBech32AccountAddrPrefix()),
			ValidatorAddressCodec: authcodec.NewBech32Codec(sdkConfig.GetBech32ValidatorAddrPrefix()),
		},
	})
	require.NoError(t, err)

	cosmosstd.RegisterInterfaces(reg)
	atypes.RegisterInterfaces(reg)
	stypes.RegisterInterfaces(reg)
	btypes.RegisterInterfaces(reg)
	dtypes.RegisterInterfaces(reg)
	etypes.RegisterInterfaces(reg)
	attesttypes.RegisterInterfaces(reg)

	return codec.NewProtoCodec(reg)
}

var _ ethclient.EngineClient = (*mockEngineAPI)(nil)
var _ etypes.AddressProvider = (*mockAddressProvider)(nil)
var _ etypes.EvmEventProcessor = (*mockLogProvider)(nil)
var _ etypes.VoteExtensionProvider = (*mockVEProvider)(nil)

type mockEngineAPI struct {
	ethclient.EngineClient
	syncings                <-chan struct{}
	fuzzer                  *fuzz.Fuzzer
	mock                    ethclient.EngineClient // avoid repeating the implementation but also allow for custom implementations of mocks
	blockNumberFunc         func(context.Context) (uint64, error)
	blockByNumberFunc       func(context.Context, *big.Int) (*types.Block, error)
	forkchoiceUpdatedV3Func func(context.Context, eengine.ForkchoiceStateV1, *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error)
	newPayloadV3Func        func(context.Context, eengine.ExecutableData, []common.Hash, *common.Hash) (eengine.PayloadStatusV1, error)
}

// newMockEngineAPI returns a new mock engine API with a fuzzer and a mock engine client.
func newMockEngineAPI(syncings int) (mockEngineAPI, error) {
	me, err := ethclient.NewEngineMock()
	if err != nil {
		return mockEngineAPI{}, err
	}

	syncs := make(chan struct{}, syncings)
	for i := 0; i < syncings; i++ {
		syncs <- struct{}{}
	}

	return mockEngineAPI{
		mock:     me,
		syncings: syncs,
		fuzzer:   ethclient.NewFuzzer(time.Now().Truncate(time.Hour * 24).Unix()),
	}, nil
}

type mockAddressProvider struct {
	address common.Address
}

type mockVEProvider struct{}

func (m mockVEProvider) PrepareVotes(_ context.Context, _ abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
	coin := sdk.NewInt64Coin("stake", 100)
	msg := stypes.NewMsgDelegate("addr", "addr", coin)

	return []sdk.Msg{msg}, nil
}

type mockLogProvider struct{}

func (m mockLogProvider) Prepare(_ context.Context, blockHash common.Hash) ([]*etypes.EVMEvent, error) {
	f := fuzz.NewWithSeed(int64(blockHash[0]))

	var topic common.Hash
	f.Fuzz(&topic)

	return []*etypes.EVMEvent{{
		Address: zeroAddr.Bytes(),
		Topics:  [][]byte{topic[:]},
	}}, nil
}

func (m mockLogProvider) Addresses() []common.Address {
	return []common.Address{zeroAddr}
}

func (m mockLogProvider) Deliver(_ context.Context, _ common.Hash, log *etypes.EVMEvent) error {
	if bytes.Equal(log.Address, zeroAddr.Bytes()) {
		panic("unexpected evm log address")
	}

	return nil
}

func (m mockAddressProvider) LocalAddress() common.Address {
	return m.address
}

func (m mockEngineAPI) maybeSync() (eengine.PayloadStatusV1, bool) {
	select {
	case <-m.syncings:
		return eengine.PayloadStatusV1{
			Status: eengine.SYNCING,
		}, true
	default:
		return eengine.PayloadStatusV1{}, false
	}
}

func (mockEngineAPI) FilterLogs(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

func (m *mockEngineAPI) BlockNumber(ctx context.Context) (uint64, error) {
	if m.blockNumberFunc != nil {
		return m.blockNumberFunc(ctx)
	}

	return m.mock.BlockNumber(ctx)
}

func (m *mockEngineAPI) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	if m.blockByNumberFunc != nil {
		return m.blockByNumberFunc(ctx, number)
	}

	return m.mock.BlockByNumber(ctx, number)
}

func (m *mockEngineAPI) NewPayloadV2(ctx context.Context, params eengine.ExecutableData) (eengine.PayloadStatusV1, error) {
	return m.mock.NewPayloadV2(ctx, params)
}

//nolint:nonamedreturns // Required for defer
func (m *mockEngineAPI) NewPayloadV3(ctx context.Context, params eengine.ExecutableData, versionedHashes []common.Hash, beaconRoot *common.Hash) (resp eengine.PayloadStatusV1, err error) {
	if status, ok := m.maybeSync(); ok {
		defer func() {
			resp.Status = status.Status
		}()
	}

	if m.newPayloadV3Func != nil {
		return m.newPayloadV3Func(ctx, params, versionedHashes, beaconRoot)
	}

	return m.mock.NewPayloadV3(ctx, params, versionedHashes, beaconRoot)
}

//nolint:nonamedreturns // Required for defer
func (m *mockEngineAPI) ForkchoiceUpdatedV3(ctx context.Context, update eengine.ForkchoiceStateV1, payloadAttributes *eengine.PayloadAttributes) (resp eengine.ForkChoiceResponse, err error) {
	if status, ok := m.maybeSync(); ok {
		defer func() {
			resp.PayloadStatus.Status = status.Status
		}()
	}

	if m.forkchoiceUpdatedV3Func != nil {
		return m.forkchoiceUpdatedV3Func(ctx, update, payloadAttributes)
	}

	return m.mock.ForkchoiceUpdatedV3(ctx, update, payloadAttributes)
}

func (m *mockEngineAPI) GetPayloadV3(ctx context.Context, payloadID eengine.PayloadID) (*eengine.ExecutionPayloadEnvelope, error) {
	return m.mock.GetPayloadV3(ctx, payloadID)
}

// pushPayload - invokes the ForkchoiceUpdatedV2 method on the mock engine and returns the payload ID.
func (m *mockEngineAPI) pushPayload(t *testing.T, ctx context.Context, feeRecipient common.Address, blockHash common.Hash, ts time.Time) *eengine.PayloadID {
	t.Helper()
	state := eengine.ForkchoiceStateV1{
		HeadBlockHash:      blockHash,
		SafeBlockHash:      blockHash,
		FinalizedBlockHash: blockHash,
	}

	var zero common.Hash

	payloadAttrs := eengine.PayloadAttributes{
		Timestamp:             uint64(ts.Unix()),
		Random:                blockHash,
		SuggestedFeeRecipient: feeRecipient,
		Withdrawals:           []*types.Withdrawal{},
		BeaconRoot:            &zero,
	}

	resp, err := m.ForkchoiceUpdatedV3(ctx, state, &payloadAttrs)
	require.NoError(t, err)

	return resp.PayloadID
}

// nextBlock creates a new block with the given height, timestamp, parentHash, and feeRecipient. It also returns the
// payload for the block. It's a utility function for testing.
func (m *mockEngineAPI) nextBlock(t *testing.T, height uint64, timestamp uint64, parentHash common.Hash,
	feeRecipient common.Address) (*types.Block, eengine.ExecutableData) {
	t.Helper()
	var header types.Header
	m.fuzzer.Fuzz(&header)
	header.Number = big.NewInt(int64(height))
	header.Time = timestamp
	header.ParentHash = parentHash
	header.Coinbase = feeRecipient
	header.MixDigest = parentHash

	// Convert header to block
	block := types.NewBlock(&header, nil, nil, nil, trie.NewStackTrie(nil))

	// Convert block to payload
	env := eengine.BlockToExecutableData(block, big.NewInt(0), nil)
	payload := *env.ExecutionPayload

	// Ensure the block is valid
	_, err := eengine.ExecutableDataToBlock(payload, nil, nil)
	require.NoError(t, err)

	return block, payload
}
