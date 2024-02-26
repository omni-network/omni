package keeper

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	attesttypes "github.com/omni-network/omni/halo/attest/types"
	etypes "github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"

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

// todo(lazar): make it table tests, ok for now.
func TestKeeper_PrepareProposal(t *testing.T) {
	t.Parallel()

	t.Run("run err scenarios", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name       string
			mockEngine MockEngineAPI
			req        *abci.RequestPrepareProposal
			wantErr    bool
		}{
			{
				name:       "no transactions",
				mockEngine: MockEngineAPI{},
				req: &abci.RequestPrepareProposal{
					Txs:    nil,        // Set to nil to simulate no transactions
					Height: 1,          // Set height to 1 for this test case
					Time:   time.Now(), // Set time to current time or mock a time
				},
				wantErr: false,
			},
			{
				name: "block number err",
				mockEngine: MockEngineAPI{
					BlockNumberFunc: func(ctx context.Context) (uint64, error) {
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
				mockEngine: MockEngineAPI{
					BlockNumberFunc: func(ctx context.Context) (uint64, error) {
						return 0, nil
					},
					BlockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
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
				mockEngine: MockEngineAPI{
					BlockNumberFunc: func(ctx context.Context) (uint64, error) {
						return 0, nil
					},
					BlockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
						fuzzer := ethclient.NewFuzzer(0)
						var block *types.Block
						fuzzer.Fuzz(&block)

						return block, nil
					},
					ForkchoiceUpdatedV2Func: func(ctx context.Context, update eengine.ForkchoiceStateV1,
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
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				ctx, storeService := setupCtxStore(t)
				cdc := getCodec()
				txConfig := authtx.NewTxConfig(cdc, nil)
				ap := MockAddressProvider{}

				k := NewKeeper(cdc, storeService, &tt.mockEngine, txConfig, ap)
				_, err := k.PrepareProposal(ctx, tt.req)
				if (err != nil) != tt.wantErr {
					t.Errorf("PrepareProposal() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			})
		}
	})

	t.Run("build non optimistic", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		me, err := ethclient.NewEngineMock()
		require.NoError(t, err)
		mockEngine := MockEngineAPI{
			Mock:   me,
			fuzzer: ethclient.NewFuzzer(time.Now().Truncate(time.Hour * 24).Unix()),
		}
		ap := MockAddressProvider{}
		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap)

		ts := time.Now()
		latestHeight, err := mockEngine.BlockNumber(ctx)
		require.NoError(t, err)
		latestBlock, err := mockEngine.BlockByNumber(ctx, big.NewInt(int64(latestHeight)))
		require.NoError(t, err)

		mockEngine.pushPayload(t, ctx, ap, latestBlock.Hash(), ts)
		nextBlock := mockEngine.nextBlock(latestHeight+1, uint64(ts.Unix()), latestBlock.Hash(), ap.LocalAddress())
		payloadID := mockEngine.pushPayload(t, ctx, ap, nextBlock.Hash(), ts)

		keeper.mutablePayload.UpdatedAt = time.Now()
		keeper.mutablePayload.ID = payloadID

		req := &abci.RequestPrepareProposal{
			Txs:    nil,
			Height: int64(2),
			Time:   time.Now(),
		}
		keeper.mutablePayload.Height = 2

		resp, err := keeper.PrepareProposal(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		tx, err := txConfig.TxDecoder()(resp.Txs[0])
		require.NoError(t, err)

		for _, msg := range tx.GetMsgs() {
			if _, ok := msg.(*etypes.MsgExecutionPayload); ok {
				assertExecutablePayload(t, msg, ts, nextBlock.Hash(), ap, uint64(req.Height))
			}
		}
	})

	t.Run("build optimistic", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		me, err := ethclient.NewEngineMock()
		require.NoError(t, err)
		mockEngine := MockEngineAPI{
			Mock:   me,
			fuzzer: ethclient.NewFuzzer(time.Now().Truncate(time.Hour * 24).Unix()),
		}
		ap := MockAddressProvider{}
		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap)

		ts := time.Now()
		latestHeight, err := mockEngine.BlockNumber(ctx)
		require.NoError(t, err)
		latestBlock, err := mockEngine.BlockByNumber(ctx, big.NewInt(int64(latestHeight)))
		require.NoError(t, err)

		mockEngine.pushPayload(t, ctx, ap, latestBlock.Hash(), ts)
		nextBlock := mockEngine.nextBlock(latestHeight+1, uint64(ts.Unix()), latestBlock.Hash(), ap.LocalAddress())
		payloadID := mockEngine.pushPayload(t, ctx, ap, nextBlock.Hash(), ts)

		keeper.mutablePayload.UpdatedAt = time.Now()
		keeper.mutablePayload.ID = payloadID

		req := &abci.RequestPrepareProposal{
			Txs:    nil,
			Height: int64(2),
			Time:   time.Now(),
		}

		resp, err := keeper.PrepareProposal(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		latestHeight, err = mockEngine.BlockNumber(ctx)
		require.NoError(t, err)
		latestBlock, err = mockEngine.BlockByNumber(ctx, big.NewInt(int64(latestHeight)))
		require.NoError(t, err)

		tx, err := txConfig.TxDecoder()(resp.Txs[0])
		require.NoError(t, err)

		for _, msg := range tx.GetMsgs() {
			if _, ok := msg.(*etypes.MsgExecutionPayload); ok {
				assertExecutablePayload(t, msg, ts, latestBlock.Hash(), ap, uint64(req.Height))
			}
		}
	})
}

func assertExecutablePayload(t *testing.T, msg sdk.Msg, ts time.Time, blockHash common.Hash, ap MockAddressProvider, height uint64) {
	t.Helper()
	executionPayload, ok := msg.(*etypes.MsgExecutionPayload)
	require.True(t, ok)
	require.NotNil(t, executionPayload)
	var ep *eengine.ExecutableData
	err := json.Unmarshal(executionPayload.GetData(), &ep)
	require.NoError(t, err)
	require.Equal(t, int64(ep.Timestamp), ts.Unix()+1)
	require.Equal(t, ep.Random, blockHash)
	require.Equal(t, ep.FeeRecipient, ap.LocalAddress())
	require.Empty(t, ep.Withdrawals)
	require.Equal(t, ep.Number, height)
}

func setupCtxStore(t *testing.T) (sdk.Context, store.KVStoreService) {
	t.Helper()
	key := storetypes.NewKVStoreKey("test")
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(t, key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})

	return ctx, storeService
}

func getCodec() *codec.ProtoCodec {
	sdkConfig := sdk.GetConfig()
	reg, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          authcodec.NewBech32Codec(sdkConfig.GetBech32AccountAddrPrefix()),
			ValidatorAddressCodec: authcodec.NewBech32Codec(sdkConfig.GetBech32ValidatorAddrPrefix()),
		},
	})
	if err != nil {
		panic(err)
	}

	cosmosstd.RegisterInterfaces(reg)
	atypes.RegisterInterfaces(reg)
	stypes.RegisterInterfaces(reg)
	btypes.RegisterInterfaces(reg)
	dtypes.RegisterInterfaces(reg)
	etypes.RegisterInterfaces(reg)
	attesttypes.RegisterInterfaces(reg)

	return codec.NewProtoCodec(reg)
}

var _ ethclient.EngineClient = (*MockEngineAPI)(nil)
var _ etypes.AddressProvider = (*MockAddressProvider)(nil)
var _ etypes.CPayloadProvider = (*MockCPayloadProvider)(nil)

type MockEngineAPI struct {
	fuzzer                  *fuzz.Fuzzer
	Mock                    ethclient.EngineClient // avoid repeating the implementation but also allow for custom implementations of mocks
	BlockNumberFunc         func(ctx context.Context) (uint64, error)
	BlockByNumberFunc       func(ctx context.Context, number *big.Int) (*types.Block, error)
	ForkchoiceUpdatedV2Func func(ctx context.Context, update eengine.ForkchoiceStateV1,
		payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error)
}

func (m *MockEngineAPI) ChainID(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) PendingTransactionCount(ctx context.Context) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) HeaderByType(ctx context.Context, typ ethclient.HeadType) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockEngineAPI) Close() {
	//TODO implement me
	panic("implement me")
}

type MockAddressProvider struct{}
type MockCPayloadProvider struct{}

func (m MockCPayloadProvider) PreparePayload(ctx context.Context, height uint64, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
	// TODO implement me
	panic("implement me")
}

func (m MockAddressProvider) LocalAddress() common.Address {
	return common.BytesToAddress([]byte("test"))
}

func (m *MockEngineAPI) BlockNumber(ctx context.Context) (uint64, error) {
	if m.BlockNumberFunc != nil {
		return m.BlockNumberFunc(ctx)
	}

	return m.Mock.BlockNumber(ctx)
}

func (m *MockEngineAPI) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	if m.BlockByNumberFunc != nil {
		return m.BlockByNumberFunc(ctx, number)
	}

	return m.Mock.BlockByNumber(ctx, number)
}

func (m *MockEngineAPI) NewPayloadV2(ctx context.Context, params eengine.ExecutableData) (eengine.PayloadStatusV1, error) {
	return m.Mock.NewPayloadV2(ctx, params)
}

func (m *MockEngineAPI) NewPayloadV3(ctx context.Context, params eengine.ExecutableData, versionedHashes []common.Hash, beaconRoot *common.Hash) (eengine.PayloadStatusV1, error) {
	return m.Mock.NewPayloadV3(ctx, params, versionedHashes, beaconRoot)
}

func (m *MockEngineAPI) ForkchoiceUpdatedV2(ctx context.Context, update eengine.ForkchoiceStateV1, payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
	if m.ForkchoiceUpdatedV2Func != nil {
		return m.ForkchoiceUpdatedV2Func(ctx, update, payloadAttributes)
	}

	return m.Mock.ForkchoiceUpdatedV2(ctx, update, payloadAttributes)
}

func (m *MockEngineAPI) ForkchoiceUpdatedV3(ctx context.Context, update eengine.ForkchoiceStateV1, payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
	panic("implement me")
}

func (m *MockEngineAPI) GetPayloadV2(ctx context.Context, payloadID eengine.PayloadID) (*eengine.ExecutionPayloadEnvelope, error) {
	return m.Mock.GetPayloadV2(ctx, payloadID)
}

func (m *MockEngineAPI) GetPayloadV3(ctx context.Context, payloadID eengine.PayloadID) (*eengine.ExecutionPayloadEnvelope, error) {
	panic("implement me")
}

func (m *MockEngineAPI) pushPayload(t *testing.T, ctx context.Context, ap etypes.AddressProvider, blockHash common.Hash, ts time.Time) *eengine.PayloadID {
	t.Helper()
	forkchoiceState := eengine.ForkchoiceStateV1{
		HeadBlockHash:      blockHash,
		SafeBlockHash:      blockHash,
		FinalizedBlockHash: blockHash,
	}

	payloadAttrs := eengine.PayloadAttributes{
		Timestamp:             uint64(ts.Unix()),
		Random:                blockHash,
		SuggestedFeeRecipient: ap.LocalAddress(),
		Withdrawals:           []*types.Withdrawal{},
		BeaconRoot:            nil,
	}

	forkchoiceResp, err := m.ForkchoiceUpdatedV2(ctx, forkchoiceState, &payloadAttrs)
	require.NoError(t, err)

	return forkchoiceResp.PayloadID
}

func (m *MockEngineAPI) nextBlock(height uint64, timestamp uint64, parentHash common.Hash, feeRecipient common.Address) *types.Block {
	var header types.Header
	m.fuzzer.Fuzz(&header)
	header.Number = big.NewInt(int64(height))
	header.Time = timestamp
	header.ParentHash = parentHash
	header.Coinbase = feeRecipient
	header.MixDigest = parentHash

	// Convert header to block
	block := types.NewBlock(&header, nil, nil, nil, trie.NewStackTrie(nil))

	return block
}
