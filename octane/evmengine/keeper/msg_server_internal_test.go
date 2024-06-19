package keeper

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/octane/evmengine/types"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func Test_msgServer_ExecutionPayload(t *testing.T) {
	t.Parallel()
	fastBackoffForT()

	cdc := getCodec(t)
	txConfig := authtx.NewTxConfig(cdc, nil)

	mockEngine, err := newMockEngineAPI(2)
	require.NoError(t, err)
	cmtAPI := newMockCometAPI(t, nil)
	// set the header and proposer so we have the correct next proposer
	header := cmtproto.Header{Height: 1, AppHash: tutil.RandomHash().Bytes()}
	header.ProposerAddress = cmtAPI.validatorSet.Validators[0].Address
	nxtAddr, err := k1util.PubKeyToAddress(cmtAPI.validatorSet.Validators[1].PubKey)
	require.NoError(t, err)

	ctx, storeService := setupCtxStore(t, &header)
	ctx = ctx.WithExecMode(sdk.ExecModeFinalize)

	ap := mockAddressProvider{
		address: nxtAddr,
	}
	frp := newRandomFeeRecipientProvider()
	evmLogProc := mockLogProvider{deliverErr: errors.New("test error")}
	keeper, err := NewKeeper(cdc, storeService, &mockEngine, txConfig, ap, frp)
	require.NoError(t, err)
	keeper.SetCometAPI(cmtAPI)
	keeper.AddEventProcessor(evmLogProc)
	populateGenesisHead(ctx, t, keeper)
	msgSrv := NewMsgServerImpl(keeper)

	var payloadData []byte
	var payloadID engine.PayloadID
	var latestHeight uint64
	var block *etypes.Block
	newPayload := func(ctx context.Context) {
		// get latest block to build on top
		latestBlock, err := mockEngine.HeaderByType(ctx, ethclient.HeadLatest)
		require.NoError(t, err)
		latestHeight = latestBlock.Number.Uint64()

		sdkCtx := sdk.UnwrapSDKContext(ctx)
		appHash := common.BytesToHash(sdkCtx.BlockHeader().AppHash)

		var execPayload engine.ExecutableData
		block, execPayload = mockEngine.nextBlock(
			t,
			latestHeight+1,
			uint64(sdkCtx.BlockHeader().Time.Unix()),
			latestBlock.Hash(),
			frp.LocalFeeRecipient(),
			&appHash,
		)

		payloadID, err = ethclient.MockPayloadID(execPayload, &appHash)
		require.NoError(t, err)

		// Create execution payload message
		payloadData, err = json.Marshal(execPayload)
		require.NoError(t, err)
	}

	assertExecutionPayload := func(ctx context.Context) {
		events, err := evmLogProc.Prepare(ctx, block.Hash())
		require.NoError(t, err)

		resp, err := msgSrv.ExecutionPayload(ctx, &types.MsgExecutionPayload{
			Authority:         authtypes.NewModuleAddress(types.ModuleName).String(),
			ExecutionPayload:  payloadData,
			PrevPayloadEvents: events,
		})
		tutil.RequireNoError(t, err)
		require.NotNil(t, resp)

		gotPayload, err := mockEngine.GetPayloadV3(ctx, payloadID)
		require.NoError(t, err)
		// make sure height is increasing in engine, blocks being built
		require.Equal(t, gotPayload.ExecutionPayload.Number, latestHeight+1)
		require.Equal(t, gotPayload.ExecutionPayload.BlockHash, block.Hash())
		require.Equal(t, gotPayload.ExecutionPayload.FeeRecipient, frp.LocalFeeRecipient())
		require.Empty(t, gotPayload.ExecutionPayload.Withdrawals)
	}

	newPayload(ctx)
	assertExecutionPayload(ctx)

	// now lets run optimistic flow
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second))

	newPayload(ctx)
	keeper.SetBuildOptimistic(true)
	assertExecutionPayload(ctx)
}

// populateGenesisHead inserts the mock genesis execution head into the database.
func populateGenesisHead(ctx context.Context, t *testing.T, keeper *Keeper) {
	t.Helper()
	genesisBlock, err := ethclient.MockGenesisBlock()
	require.NoError(t, err)

	require.NoError(t, keeper.InsertGenesisHead(ctx, genesisBlock.Hash().Bytes()))
}

func Test_pushPayload(t *testing.T) {
	t.Parallel()

	newPayload := func(ctx context.Context, mockEngine mockEngineAPI, address common.Address) (engine.ExecutableData, engine.PayloadID) {
		// get latest block to build on top
		latestBlock, err := mockEngine.HeaderByType(ctx, ethclient.HeadLatest)
		require.NoError(t, err)
		latestHeight := latestBlock.Number.Uint64()

		sdkCtx := sdk.UnwrapSDKContext(ctx)
		appHash := common.BytesToHash(sdkCtx.BlockHeader().AppHash)

		_, execPayload := mockEngine.nextBlock(t, latestHeight+1, uint64(time.Now().Unix()), latestBlock.Hash(), address, &appHash)
		payloadID, err := ethclient.MockPayloadID(execPayload, &appHash)
		require.NoError(t, err)

		return execPayload, payloadID
	}
	type args struct {
		transformPayload func(*engine.ExecutableData)
		newPayloadV3Func func(context.Context, engine.ExecutableData, []common.Hash, *common.Hash) (engine.PayloadStatusV1, error)
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantStatus string
	}{
		{
			name: "new payload error",
			args: args{
				newPayloadV3Func: func(context.Context, engine.ExecutableData, []common.Hash, *common.Hash) (engine.PayloadStatusV1, error) {
					return engine.PayloadStatusV1{}, errors.New("error")
				},
			},
			wantErr:    true,
			wantStatus: "",
		},
		{
			name: "new payload invalid",
			args: args{
				newPayloadV3Func: func(context.Context, engine.ExecutableData, []common.Hash, *common.Hash) (engine.PayloadStatusV1, error) {
					return engine.PayloadStatusV1{
						Status:          engine.INVALID,
						LatestValidHash: nil,
						ValidationError: nil,
					}, nil
				},
			},
			wantErr:    false,
			wantStatus: engine.INVALID,
		},
		{
			name: "new payload invalid val err",
			args: args{
				newPayloadV3Func: func(context.Context, engine.ExecutableData, []common.Hash, *common.Hash) (engine.PayloadStatusV1, error) {
					return engine.PayloadStatusV1{
						Status:          engine.INVALID,
						LatestValidHash: nil,
						ValidationError: func() *string { s := "error"; return &s }(),
					}, nil
				},
			},
			wantErr:    false,
			wantStatus: engine.INVALID,
		},
		{
			name: "new payload syncing",
			args: args{
				newPayloadV3Func: func(context.Context, engine.ExecutableData, []common.Hash, *common.Hash) (engine.PayloadStatusV1, error) {
					return engine.PayloadStatusV1{
						Status:          engine.SYNCING,
						LatestValidHash: nil,
						ValidationError: nil,
					}, nil
				},
			},
			wantErr:    false,
			wantStatus: engine.SYNCING,
		},
		{
			name: "new payload accepted",
			args: args{
				newPayloadV3Func: func(context.Context, engine.ExecutableData, []common.Hash, *common.Hash) (engine.PayloadStatusV1, error) {
					return engine.PayloadStatusV1{
						Status:          engine.ACCEPTED,
						LatestValidHash: nil,
						ValidationError: nil,
					}, nil
				},
			},
			wantErr:    false,
			wantStatus: engine.ACCEPTED,
		},
		{
			name:       "valid payload",
			args:       args{},
			wantErr:    false,
			wantStatus: engine.VALID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			appHash := tutil.RandomHash()
			ctx := ctxWithAppHash(t, appHash)
			mockEngine, err := newMockEngineAPI(0)
			require.NoError(t, err)
			mockEngine.newPayloadV3Func = tt.args.newPayloadV3Func
			payload, payloadID := newPayload(ctx, mockEngine, common.Address{})
			if tt.args.transformPayload != nil {
				tt.args.transformPayload(&payload)
			}

			status, err := pushPayload(ctx, &mockEngine, payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("pushPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.wantStatus, status.Status)

			if status.Status == engine.VALID {
				want, err := mockEngine.GetPayloadV3(ctx, payloadID)
				require.NoError(t, err)
				if !reflect.DeepEqual(payload, *want.ExecutionPayload) {
					t.Errorf("pushPayload() got = %v, want %v", payload, want)
				}
			}
		})
	}
}
