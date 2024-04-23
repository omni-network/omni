package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

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
	header := cmtproto.Header{Height: 1}
	header.ProposerAddress = cmtAPI.validatorSet.Validators[0].Address
	nxtAddr, err := k1util.PubKeyToAddress(cmtAPI.validatorSet.Validators[1].PubKey)
	require.NoError(t, err)

	ctx, storeService := setupCtxStore(t, &header)
	ctx = ctx.WithExecMode(sdk.ExecModeFinalize)

	keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig)
	ap := mockAddressProvider{
		address: nxtAddr,
	}
	keeper.SetAddressProvider(ap)
	keeper.SetCometAPI(&cmtAPI)
	msgSrv := NewMsgServerImpl(keeper)

	var payloadData []byte
	var payloadID engine.PayloadID
	var latestHeight uint64
	var block *etypes.Block
	newPayload := func() {
		// get latest block to build on top
		latestHeight, err = mockEngine.BlockNumber(ctx)
		require.NoError(t, err)
		latestBlock, err := mockEngine.BlockByNumber(ctx, big.NewInt(int64(latestHeight)))
		require.NoError(t, err)

		b, execPayload := mockEngine.nextBlock(t, latestHeight+1, uint64(time.Now().Unix()), latestBlock.Hash(), ap.LocalAddress())
		block = b

		payloadID, err = toPayloadID(execPayload)
		require.NoError(t, err)

		// Create execution payload message
		payloadData, err = json.Marshal(execPayload)
		require.NoError(t, err)
	}

	assertExecutionPayload := func() {
		resp, err := msgSrv.ExecutionPayload(ctx, &types.MsgExecutionPayload{
			Authority:        authtypes.NewModuleAddress(types.ModuleName).String(),
			ExecutionPayload: payloadData,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		gotPayload, err := mockEngine.GetPayloadV3(ctx, payloadID)
		require.NoError(t, err)
		// make sure height is increasing in engine, blocks being built
		require.Equal(t, gotPayload.ExecutionPayload.Number, latestHeight+1)
		require.Equal(t, gotPayload.ExecutionPayload.BlockHash, block.Hash())
		require.Equal(t, gotPayload.ExecutionPayload.FeeRecipient, ap.LocalAddress())
		require.Empty(t, gotPayload.ExecutionPayload.Withdrawals)
	}

	newPayload()
	assertExecutionPayload()

	// not lets run optimistic flow
	newPayload()
	keeper.SetBuildOptimistic(true)
	assertExecutionPayload()
}

// toPayloadID returns a deterministic payload id for the given payload.
func toPayloadID(payload engine.ExecutableData) (engine.PayloadID, error) {
	bz, err := payload.MarshalJSON()
	if err != nil {
		return engine.PayloadID{}, errors.Wrap(err, "marshal payload")
	}

	hash := sha256.Sum256(bz)

	return engine.PayloadID(hash[:8]), nil
}

func Test_pushPayload(t *testing.T) {
	t.Parallel()

	newPayload := func(ctx context.Context, mockEngine mockEngineAPI, address common.Address) ([]byte, engine.PayloadID) {
		// get latest block to build on top
		latestHeight, err := mockEngine.BlockNumber(ctx)
		require.NoError(t, err)
		latestBlock, err := mockEngine.BlockByNumber(ctx, big.NewInt(int64(latestHeight)))
		require.NoError(t, err)

		_, execPayload := mockEngine.nextBlock(t, latestHeight+1, uint64(time.Now().Unix()), latestBlock.Hash(), address)
		payloadID, err := toPayloadID(execPayload)
		require.NoError(t, err)
		// Create execution payload message
		payloadData, err := json.Marshal(execPayload)
		require.NoError(t, err)

		return payloadData, payloadID
	}
	type args struct {
		msg              *types.MsgExecutionPayload
		newPayloadV3Func func(context.Context, engine.ExecutableData, []common.Hash, *common.Hash) (engine.PayloadStatusV1, error)
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantStatus string
	}{
		{
			name: "fail to unmarshal",
			args: args{
				msg: &types.MsgExecutionPayload{ExecutionPayload: []byte("invalid")},
			},
			wantErr:    true,
			wantStatus: "",
		},
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
			ctx := context.Background()
			mockEngine, err := newMockEngineAPI(0)
			require.NoError(t, err)
			mockEngine.newPayloadV3Func = tt.args.newPayloadV3Func
			payload, payloadID := newPayload(ctx, mockEngine, common.Address{})
			if tt.args.msg == nil {
				tt.args.msg = &types.MsgExecutionPayload{
					ExecutionPayload: payload,
				}
			}

			got, status, err := pushPayload(ctx, &mockEngine, tt.args.msg)
			require.Equal(t, tt.wantStatus, status.Status)
			if (err != nil) != tt.wantErr {
				t.Errorf("pushPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if status.Status == engine.VALID {
				want, err := mockEngine.GetPayloadV3(ctx, payloadID)
				require.NoError(t, err)
				if !reflect.DeepEqual(got, *want.ExecutionPayload) {
					t.Errorf("pushPayload() got = %v, want %v", got, want)
				}
			}
		})
	}
}
