package keeper

import (
	"crypto/sha256"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	etypes "github.com/ethereum/go-ethereum/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func Test_msgServer_ExecutionPayload(t *testing.T) {
	t.Parallel()
	cdc := getCodec(t)
	txConfig := authtx.NewTxConfig(cdc, nil)

	mockEngine, err := newMockEngineAPI()
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
			Authority: authtypes.NewModuleAddress(types.ModuleName).String(),
			Data:      payloadData,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		gotPayload, err := mockEngine.GetPayloadV2(ctx, payloadID)
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
