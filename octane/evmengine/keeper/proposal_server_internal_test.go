package keeper

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func Test_proposalServer_ExecutionPayload(t *testing.T) {
	t.Parallel()
	cdc := getCodec(t)
	txConfig := authtx.NewTxConfig(cdc, nil)

	mockEngine, err := newMockEngineAPI(0)
	require.NoError(t, err)

	sdkCtx, storeService := setupCtxStore(t, nil)
	sdkCtx = sdkCtx.WithExecMode(sdk.ExecModeFinalize)

	frp := newRandomFeeRecipientProvider()
	keeper, err := NewKeeper(cdc, storeService, &mockEngine, txConfig, nil, frp)
	require.NoError(t, err)
	populateGenesisHead(sdkCtx, t, keeper)

	propSrv := NewProposalServer(keeper)

	var payload engine.ExecutableData
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

		block, payload = mockEngine.nextBlock(
			t,
			latestHeight+1,
			uint64(sdkCtx.BlockHeader().Time.Unix()),
			latestBlock.Hash(),
			frp.LocalFeeRecipient(),
			&appHash,
		)

		payloadID, err = ethclient.MockPayloadID(payload, &appHash)
		require.NoError(t, err)
	}

	assertExecutionPayload := func(ctx context.Context) {
		payloadProto, err := types.PayloadToProto(&payload)
		require.NoError(t, err)

		resp, err := propSrv.ExecutionPayload(ctx, &types.MsgExecutionPayload{
			Authority:             authtypes.NewModuleAddress(types.ModuleName).String(),
			ExecutionPayloadDeneb: payloadProto,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		gotPayload, err := mockEngine.GetPayloadV3(ctx, payloadID)
		require.NoError(t, err)
		require.Equal(t, latestHeight+1, gotPayload.ExecutionPayload.Number)
		require.Equal(t, block.Hash(), gotPayload.ExecutionPayload.BlockHash)
		require.Equal(t, frp.LocalFeeRecipient(), gotPayload.ExecutionPayload.FeeRecipient)
		require.Empty(t, gotPayload.ExecutionPayload.Withdrawals)
	}

	newPayload(sdkCtx)
	assertExecutionPayload(sdkCtx)
}

func fastBackoffForT() {
	backoffFuncMu.Lock()
	defer backoffFuncMu.Unlock()
	backoffFunc = func(context.Context, ...func(*expbackoff.Config)) func() {
		return func() {}
	}
}
