package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// PrepareProposal returns a proposal for the next block.
// Note returning an error results in a panic cometbft and CONSENSUS_FAILURE log.
func (k *Keeper) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (
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

	appHash := common.BytesToHash(ctx.BlockHeader().AppHash)

	// Either use the optimistic payload or create a new one.
	payloadID, height, triggeredAt := k.getOptimisticPayload()
	if uint64(req.Height) != height {
		// Create a new payload (retrying on network errors).
		err := retryForever(ctx, func(ctx context.Context) (bool, error) {
			response, err := submitPayload(
				ctx,
				k.engineCl,
				req.Time,
				k.feeRecProvider.LocalFeeRecipient(),
				appHash,
			)
			if err != nil {
				log.Warn(ctx, "Preparing proposal failed: build new evm payload (will retry)", err)
				return false, nil
			} else if response.PayloadStatus.Status != engine.VALID {
				return false, errors.New("status not valid")
			}

			payloadID = response.PayloadID

			return true, nil
		})
		if err != nil {
			return nil, err
		}
		triggeredAt = time.Now()
	} else {
		log.Debug(ctx, "Using optimistic payload", "height", height, "payload", payloadID.String())
	}

	// Wait the minimum build_delay for the payload to be available.
	waitTo := triggeredAt.Add(k.buildDelay)
	select {
	case <-ctx.Done():
		return nil, errors.Wrap(ctx.Err(), "context done")
	case <-time.After(time.Until(waitTo)):
	}

	// Fetch the payload (retrying on network errors).
	var payloadResp *engine.ExecutionPayloadEnvelope
	err := retryForever(ctx, func(ctx context.Context) (bool, error) {
		var err error
		payloadResp, err = k.engineCl.GetPayloadV3(ctx, *payloadID)
		if isUnknownPayload(err) {
			return false, err
		} else if err != nil {
			log.Warn(ctx, "Preparing proposal failed: get evm payload (will retry)", err)
			return false, nil
		}

		return true, nil
	})
	if err != nil {
		return nil, err
	}

	// Create execution payload message
	payloadData, err := json.Marshal(payloadResp.ExecutionPayload)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	// First, collect all vote extension msgs from the vote provider.
	voteMsgs, err := k.voteProvider.PrepareVotes(ctx, req.LocalLastCommit)
	if err != nil {
		return nil, errors.Wrap(err, "prepare votes")
	}

	// Next, collect all prev payload evm event logs.
	evmEvents, err := k.evmEvents(ctx, payloadResp.ExecutionPayload.ParentHash)
	if err != nil {
		return nil, errors.Wrap(err, "prepare evm event logs")
	}

	// Then construct the execution payload message.
	payloadMsg := &types.MsgExecutionPayload{
		Authority:         authtypes.NewModuleAddress(types.ModuleName).String(),
		ExecutionPayload:  payloadData,
		PrevPayloadEvents: evmEvents,
	}

	// Combine all the votes messages and the payload message into a single transaction.
	b := k.txConfig.NewTxBuilder()
	if err := b.SetMsgs(append(voteMsgs, payloadMsg)...); err != nil {
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
		"vote_msgs", len(voteMsgs),
		"evm_events", len(evmEvents),
	)

	return &abci.ResponsePrepareProposal{Txs: [][]byte{tx}}, nil
}

// Finalize is called by our custom ABCI wrapper after a block is finalized.
// It starts an optimistic build if enabled and if we are the next proposer.
//
// This custom ABCI callback is used since we need to trigger optimistic builds
// immediately after FinalizeBlock with the latest app hash
// which isn't available from cosmosSDK otherwise.
func (k *Keeper) Finalize(ctx context.Context, req *abci.RequestFinalizeBlock, resp *abci.ResponseFinalizeBlock) error {
	if !k.buildOptimistic {
		return nil // Not enabled.
	}

	// Maybe start building the next block if we are the next proposer.
	isNext, err := k.isNextProposer(ctx, req.ProposerAddress, req.Height)
	if err != nil {
		return errors.Wrap(err, "next proposer")
	} else if !isNext {
		return nil // Nothing to do if we are not next proposer.
	}

	nextHeight := req.Height + 1
	ctx = log.WithCtx(ctx, "next_height", nextHeight)

	log.Debug(ctx, "Starting optimistic EVM payload build")

	fcr, err := k.startOptimisticBuild(ctx, common.BytesToHash(resp.AppHash), req.Time)
	if err != nil || isUnknown(fcr.PayloadStatus) {
		log.Warn(ctx, "Starting optimistic build failed", err)
		return nil
	} else if isSyncing(fcr.PayloadStatus) {
		log.Warn(ctx, "Starting optimistic build failed; evm syncing", nil)
		return nil
	} else if invalid, err := isInvalid(fcr.PayloadStatus); invalid {
		log.Error(ctx, "Starting optimistic build failed; invalid payload [BUG]", err)
		return nil
	}

	k.setOptimisticPayload(fcr.PayloadID, uint64(nextHeight))

	return nil
}

func (k *Keeper) startOptimisticBuild(ctx context.Context, appHash common.Hash, timestamp time.Time) (engine.ForkChoiceResponse, error) {
	latestEBlock, err := k.engineCl.HeaderByType(ctx, ethclient.HeadLatest)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "get head")
	}

	ts := uint64(timestamp.Unix())
	if ts <= latestEBlock.Time {
		ts = latestEBlock.Time + 1 // Subsequent blocks must have a higher timestamp.
	}

	// CometBFT has instant finality, so head/safe/finalized is latest height.
	fcs := engine.ForkchoiceStateV1{
		HeadBlockHash:      latestEBlock.Hash(),
		SafeBlockHash:      latestEBlock.Hash(),
		FinalizedBlockHash: latestEBlock.Hash(),
	}

	attrs := &engine.PayloadAttributes{
		Timestamp:             ts,
		Random:                latestEBlock.Hash(), // We use head block hash as randao.
		SuggestedFeeRecipient: k.feeRecProvider.LocalFeeRecipient(),
		Withdrawals:           []*etypes.Withdrawal{}, // Withdrawals not supported yet.
		BeaconRoot:            &appHash,
	}

	resp, err := k.engineCl.ForkchoiceUpdatedV3(ctx, fcs, attrs)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "forkchoice update")
	}

	return resp, nil
}

func submitPayload(
	ctx context.Context,
	engineCl ethclient.EngineClient,
	ts time.Time,
	feeRecipient common.Address,
	appHash common.Hash,
) (engine.ForkChoiceResponse, error) {
	latestEBlock, err := engineCl.HeaderByType(ctx, ethclient.HeadLatest)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "latest execution block")
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
	timestamp := uint64(ts.Unix())
	if timestamp <= latestEBlock.Time {
		timestamp = latestEBlock.Time + 1
	}

	payloadAttrs := engine.PayloadAttributes{
		Timestamp:             timestamp,
		Random:                latestEBlock.Hash(), // TODO(corver): implement proper randao.
		SuggestedFeeRecipient: feeRecipient,
		Withdrawals:           []*etypes.Withdrawal{}, // Withdrawals not supported yet.
		BeaconRoot:            &appHash,
	}

	forkchoiceResp, err := engineCl.ForkchoiceUpdatedV3(ctx, forkchoiceState, &payloadAttrs)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "forkchoice updated")
	}

	return forkchoiceResp, nil
}

// isUnknownPayload returns true if the error is due to an unknown payload.
func isUnknownPayload(err error) bool {
	if err == nil {
		return false
	}

	// TODO(corver): Add support for typed errors.
	if strings.Contains(
		strings.ToLower(err.Error()),
		strings.ToLower(engine.UnknownPayload.Error()),
	) {
		return true
	}

	return false
}
