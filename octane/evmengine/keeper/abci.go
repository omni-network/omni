package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"runtime/debug"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/types"

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
	} else if req.MaxTxBytes < cmttypes.MaxBlockSizeBytes*9/10 {
		// ConsensusParams.Block.MaxBytes is set to -1, so req.MaxTxBytes should be close to MaxBlockSizeBytes.
		return nil, errors.New("invalid max tx bytes [BUG]", "max_tx_bytes", req.MaxTxBytes)
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
			response, err := k.startBuild(ctx, appHash, req.Time)
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

// PostFinalize is called by our custom ABCI wrapper after a block is finalized.
// It starts an optimistic build if enabled and if we are the next proposer.
//
// This custom ABCI callback is used since we need to trigger optimistic builds
// immediately after FinalizeBlock with the latest app hash
// which isn't available from cosmosSDK otherwise.
func (k *Keeper) PostFinalize(ctx sdk.Context) error {
	if !k.buildOptimistic {
		return nil // Not enabled.
	}

	// Extract context values
	height := ctx.BlockHeight()
	proposer := ctx.BlockHeader().ProposerAddress
	timestamp := ctx.BlockTime()
	appHash := common.BytesToHash(ctx.BlockHeader().AppHash) // This is the app hash after the block is finalized.

	// Maybe start building the next block if we are the next proposer.
	isNext, err := k.isNextProposer(ctx, proposer, height)
	if err != nil {
		return errors.Wrap(err, "next proposer")
	} else if !isNext {
		return nil // Nothing to do if we are not next proposer.
	}

	nextHeight := height + 1
	logAttr := slog.Int64("next_height", nextHeight)
	log.Debug(ctx, "Starting optimistic EVM payload build", logAttr)

	fcr, err := k.startBuild(ctx, appHash, timestamp)
	if err != nil || isUnknown(fcr.PayloadStatus) {
		log.Warn(ctx, "Starting optimistic build failed", err, logAttr)
		return nil
	} else if isSyncing(fcr.PayloadStatus) {
		log.Warn(ctx, "Starting optimistic build failed; evm syncing", nil, logAttr)
		return nil
	} else if invalid, err := isInvalid(fcr.PayloadStatus); invalid {
		log.Error(ctx, "Starting optimistic build failed; invalid payload [BUG]", err, logAttr)
		return nil
	}

	k.setOptimisticPayload(fcr.PayloadID, uint64(nextHeight))

	return nil
}

// startBuild triggers the building of a new execution payload on top of the current execution head.
// It returns the EngineAPI response which contains a status and payload ID.
func (k *Keeper) startBuild(ctx context.Context, appHash common.Hash, timestamp time.Time) (engine.ForkChoiceResponse, error) {
	head, err := k.getExecutionHead(ctx)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "latest execution block")
	}

	// Use provided time as timestamp for the next block.
	// Or use latest execution block timestamp + 1 if is not greater.
	// Since execution blocks must have unique second-granularity timestamps.
	ts := uint64(timestamp.Unix())
	if ts <= head.GetBlockTime() {
		ts = head.GetBlockTime() + 1 // Subsequent blocks must have a higher timestamp.
	}

	// CometBFT has instant finality, so head/safe/finalized is latest height.
	fcs := engine.ForkchoiceStateV1{
		HeadBlockHash:      head.Hash(),
		SafeBlockHash:      head.Hash(),
		FinalizedBlockHash: head.Hash(),
	}

	attrs := &engine.PayloadAttributes{
		Timestamp:             ts,
		Random:                head.Hash(), // We use head block hash as randao.
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
