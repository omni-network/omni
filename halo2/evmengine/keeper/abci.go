package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"
	"runtime/debug"

	"github.com/omni-network/omni/halo2/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// PrepareProposal returns a proposal for the next block.
// Note returning an error results in a panic cometbft and CONSENSUS_FAILURE log.
func (k Keeper) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (
	*abci.ResponsePrepareProposal, error,
) {
	defer func() {
		if r := recover(); r != nil {
			log.Error(ctx, "PrepareProposal panic", nil, "recover", r)
			fmt.Println("panic stacktrace: \n" + string(debug.Stack())) //nolint:forbidigo // Print stacktrace
			panic(r)
		}
	}()
	if len(req.Txs) > 0 {
		return nil, errors.New("unexpected transactions in proposal")
	}

	if req.Height == 1 {
		// Current issue is that InitChain doesn't reset the gas meter.
		// So if the first block contains any transactions, we get a app_hash_mismatch
		// Since the proposal calculates the incorrect gas for the first block after InitChain.
		// The gas meter is reset at the end of the 1st block, so we can then start including txs.

		log.Warn(ctx, "Creating empty initial block due to gas issue", nil)
		return &abci.ResponsePrepareProposal{}, nil
	}

	latestEHeight, err := k.ethCl.BlockNumber(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "latest execution block number")
	}

	latestEBlock, err := k.ethCl.BlockByNumber(ctx, big.NewInt(int64(latestEHeight)))
	if err != nil {
		return nil, errors.Wrap(err, "latest execution block")
	}

	// CometBFT has instant finality, so head/safe/finalized is latest height.
	forkchoiceState := engine.ForkchoiceStateV1{
		HeadBlockHash:      latestEBlock.Hash(),
		SafeBlockHash:      latestEBlock.Hash(),
		FinalizedBlockHash: latestEBlock.Hash(),
	}

	// Use req time as timestamp for the next block.
	// Or use latest execution block timestamp + 1 if is not greater.
	// Since execution blocks must have unique second-granularity timestamps.
	// TODO(corver): Maybe error if timestamp is not greater than latest execution block.
	timestamp := uint64(req.Time.Unix())
	if timestamp <= latestEBlock.Time() {
		timestamp = latestEBlock.Time() + 1
	}

	payloadAttrs := engine.PayloadAttributes{
		Timestamp:             timestamp,
		Random:                latestEBlock.Hash(),                        // TODO(corver): implement proper randao.
		SuggestedFeeRecipient: common.BytesToAddress(req.ProposerAddress), // TODO(corver): Ensure this is correct.
		Withdrawals:           []*etypes.Withdrawal{},                     // Withdrawals not supported yet.
		BeaconRoot:            nil,
	}

	forkchoiceResp, err := k.ethCl.ForkchoiceUpdatedV2(ctx, forkchoiceState, &payloadAttrs)
	if err != nil {
		return nil, errors.Wrap(err, "forkchoice updated")
	} else if forkchoiceResp.PayloadStatus.Status != engine.VALID {
		return nil, errors.New("status not valid")
	}

	payloadResp, err := k.ethCl.GetPayloadV2(ctx, *forkchoiceResp.PayloadID)
	if err != nil {
		return nil, errors.Wrap(err, "get payload")
	}

	// Create execution payload message
	payloadData, err := json.Marshal(payloadResp.ExecutionPayload)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	// Collect all msgs to include in our single consensus transaction.
	msgs := []sdk.Msg{
		// First and most important is the execution payload.
		&types.MsgExecutionPayload{
			Authority: authtypes.NewModuleAddress(types.ModuleName).String(),
			Data:      payloadData,
		},
	}

	// Next, collect all msgs from the CPayload providers.
	// These include msgs from vote extensions and/or any msgs from EVM contracts.
	for _, provider := range k.providers {
		cmsgs, err := provider.PreparePayload(ctx, uint64(req.Height), req.LocalLastCommit)
		if err != nil {
			return nil, errors.Wrap(err, "prepare payload")
		}
		msgs = append(msgs, cmsgs...)
	}

	b := k.txConfig.NewTxBuilder()
	if err := b.SetMsgs(msgs...); err != nil {
		return nil, errors.Wrap(err, "set tx builder msgs")
	}

	// Note this transaction is not signed. We need to ensure bypass verification somehow.
	tx, err := k.txConfig.TxEncoder()(b.GetTx())
	if err != nil {
		return nil, errors.Wrap(err, "encode tx builder")
	}

	log.Info(ctx, "Proposing new block",
		"height", req.Height,
		log.Hex7("execution_block_hash", payloadResp.ExecutionPayload.BlockHash[:]),
		"msgs", len(msgs),
	)

	return &abci.ResponsePrepareProposal{Txs: [][]byte{tx}}, nil
}
