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

func TestKeeper_PrepareProposal(t *testing.T) {
	t.Parallel()

	// Test case 1: Test when there are no transactions in the proposal
	t.Run("NoTransactions", func(t *testing.T) {
		t.Parallel()
		ctx, storeService := getTestContext(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		keeper := NewKeeper(cdc, storeService, MockEngineAPI{}, txConfig, MockAddressProvider{})

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
		ctx, storeService := getTestContext(t)
		cdc := getCodec()
		txConfig := authtx.NewTxConfig(cdc, nil)

		keeper := NewKeeper(cdc, storeService, MockEngineAPI{}, txConfig, MockAddressProvider{})

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
}

func getTestContext(t *testing.T) (sdk.Context, store.KVStoreService) {
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

type MockEngineAPI struct{}
type MockAddressProvider struct{}

func (m MockAddressProvider) LocalAddress() common.Address {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) BlockNumber(ctx context.Context) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	// TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) NewPayloadV2(ctx context.Context, params eengine.ExecutableData) (eengine.PayloadStatusV1, error) {
	// TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) NewPayloadV3(ctx context.Context, params eengine.ExecutableData, versionedHashes []common.Hash, beaconRoot *common.Hash) (eengine.PayloadStatusV1, error) {
	// TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) ForkchoiceUpdatedV2(ctx context.Context, update eengine.ForkchoiceStateV1, payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) ForkchoiceUpdatedV3(ctx context.Context, update eengine.ForkchoiceStateV1, payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) GetPayloadV2(ctx context.Context, payloadID eengine.PayloadID) (*eengine.ExecutionPayloadEnvelope, error) {
	// TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) GetPayloadV3(ctx context.Context, payloadID eengine.PayloadID) (*eengine.ExecutionPayloadEnvelope, error) {
	// TODO implement me
	panic("implement me")
}
