package keeper

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/expbackoff"

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

	ctx, storeService := setupCtxStore(t, nil)
	ctx = ctx.WithExecMode(sdk.ExecModeFinalize)

	keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig)
	propSrv := NewProposalServer(keeper)

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

		b, execPayload := mockEngine.nextBlock(t, latestHeight+1, uint64(time.Now().Unix()), latestBlock.Hash(), common.Address{})
		block = b

		payloadID, err = toPayloadID(execPayload)
		require.NoError(t, err)

		// Create execution payload message
		payloadData, err = json.Marshal(execPayload)
		require.NoError(t, err)
	}

	assertExecutionPayload := func() {
		resp, err := propSrv.ExecutionPayload(ctx, &types.MsgExecutionPayload{
			Authority:        authtypes.NewModuleAddress(types.ModuleName).String(),
			ExecutionPayload: payloadData,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		gotPayload, err := mockEngine.GetPayloadV3(ctx, payloadID)
		require.NoError(t, err)
		require.Equal(t, latestHeight+1, gotPayload.ExecutionPayload.Number)
		require.Equal(t, block.Hash(), gotPayload.ExecutionPayload.BlockHash)
		require.Equal(t, common.Address{}, gotPayload.ExecutionPayload.FeeRecipient)
		require.Empty(t, gotPayload.ExecutionPayload.Withdrawals)
	}

	newPayload()
	assertExecutionPayload()
}

func fastBackoffForT() {
	backoffFunc = func(context.Context, ...func(*expbackoff.Config)) func() {
		return func() {}
	}
}
