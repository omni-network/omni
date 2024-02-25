package keeper

import (
	"context"
	"math/big"
	"testing"
	"time"

	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	attesttypes "github.com/omni-network/omni/halo/attest/types"
	etypes "github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/errors"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"

	eengine "github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

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
	atypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	btypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

// todo(lazar): make it table tests, ok for now
func TestKeeper_PrepareProposal(t *testing.T) {
	t.Parallel()

	// Test case 1: Test when there are no transactions in the proposal
	t.Run("NoTransactions", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		keeper := NewKeeper(cdc, storeService, &MockEngineAPI{}, txConfig, MockAddressProvider{})

		req := &abci.RequestPrepareProposal{
			Txs:    nil,        // Set to nil to simulate no transactions
			Height: 1,          // Set height to 1 for this test case
			Time:   time.Now(), // Set time to current time or mock a time
		}

		resp, err := keeper.PrepareProposal(ctx, req)

		// Assert that the response is as expected
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Empty(t, resp.Txs) // Expecting no transactions in the response
	})

	// Test case 2: Test when there are transactions in the proposal
	t.Run("WithTransactions", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		keeper := NewKeeper(cdc, storeService, &MockEngineAPI{}, txConfig, MockAddressProvider{})

		req := &abci.RequestPrepareProposal{
			Txs:    [][]byte{[]byte("test1")}, // Set to some transactions to simulate transactions in the proposal
			Height: 2,                         // Set height to 2 for this test case
			Time:   time.Now(),                // Set time to current time or mock a time
		}

		resp, err := keeper.PrepareProposal(ctx, req)

		// Assert that the response is as expected
		require.Error(t, err) // Expecting an error
		require.Nil(t, resp)
	})

	// Test case 3: Test when the block number is successfully fetched
	t.Run("Block number err", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		mockEngine := MockEngineAPI{
			BlockNumberFunc: func(ctx context.Context) (uint64, error) {
				return 0, errors.New("mocked error")
			},
		}

		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, MockAddressProvider{})
		height := int64(2)

		req := &abci.RequestPrepareProposal{
			Txs:    nil, // Set to nil to simulate no transactions
			Height: height,
			Time:   time.Now(), // Set time to current time or mock a time
		}

		resp, err := keeper.PrepareProposal(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	// Test case 4: Test when the block number is successfully fetched
	t.Run("Block by number err", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		mockEngine := MockEngineAPI{
			BlockNumberFunc: func(ctx context.Context) (uint64, error) {
				return 0, nil
			},
			BlockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
				return nil, errors.New("mocked error")
			},
		}

		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, MockAddressProvider{})
		height := int64(2)

		req := &abci.RequestPrepareProposal{
			Txs:    nil, // Set to nil to simulate no transactions
			Height: height,
			Time:   time.Now(), // Set time to current time or mock a time
		}

		resp, err := keeper.PrepareProposal(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	// Test case 4: Test when the block number is successfully fetched
	t.Run("Block by number err", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		mockEngine := MockEngineAPI{
			BlockNumberFunc: func(ctx context.Context) (uint64, error) {
				return 0, nil
			},
			BlockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
				return nil, errors.New("mocked error")
			},
		}

		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, MockAddressProvider{})
		height := int64(2)

		req := &abci.RequestPrepareProposal{
			Txs:    nil, // Set to nil to simulate no transactions
			Height: height,
			Time:   time.Now(), // Set time to current time or mock a time
		}

		resp, err := keeper.PrepareProposal(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	// Test case 5: Test when the forkchoice update errs
	t.Run("forkchoiceUpdateV2  err", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		mockEngine := MockEngineAPI{
			BlockNumberFunc: func(ctx context.Context) (uint64, error) {
				return 0, nil
			},
			BlockByNumberFunc: func(ctx context.Context, number *big.Int) (*types.Block, error) {
				fuzzer := engine.NewFuzzer(0)
				var block *types.Block
				fuzzer.Fuzz(&block)
				return block, nil
			},
			ForkchoiceUpdatedV2Func: func(ctx context.Context, update eengine.ForkchoiceStateV1,
				payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
				return eengine.ForkChoiceResponse{}, errors.New("mocked error")
			},
		}

		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, MockAddressProvider{})
		height := int64(2)

		req := &abci.RequestPrepareProposal{
			Txs:    nil, // Set to nil to simulate no transactions
			Height: height,
			Time:   time.Now(), // Set time to current time or mock a time
		}

		resp, err := keeper.PrepareProposal(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("use mock impl", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := setupCtxStore(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		me, err := engine.NewMock()
		require.NoError(t, err)
		mockEngine := MockEngineAPI{
			Mock: me,
		}

		keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig, MockAddressProvider{})
		height := int64(2)
		keeper.mutablePayload.UpdatedAt = time.Now()

		req := &abci.RequestPrepareProposal{
			Txs:    nil,
			Height: height,
			Time:   time.Now(),
		}

		resp, err := keeper.PrepareProposal(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
	})

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

var _ engine.API = (*MockEngineAPI)(nil)
var _ etypes.AddressProvider = (*MockAddressProvider)(nil)

type MockEngineAPI struct {
	*engine.Mock            // avoid repeating the implementation but also allow for custom implementations of mocks
	BlockNumberFunc         func(ctx context.Context) (uint64, error)
	BlockByNumberFunc       func(ctx context.Context, number *big.Int) (*types.Block, error)
	ForkchoiceUpdatedV2Func func(ctx context.Context, update eengine.ForkchoiceStateV1,
		payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error)
}
type MockAddressProvider struct{}

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
