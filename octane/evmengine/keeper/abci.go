package keeper

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// prepareTimeout is the maximum time to prepare a proposal.
// Timeout results in proposing an empty consensus block.
const prepareTimeout = time.Second * 10

// PrepareProposal returns a proposal for the next block.
// Note returning an error results proposing an empty block.
func (k *Keeper) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (
	*abci.ResponsePrepareProposal, error,
) {
	// Only allow 10s to prepare a proposal. Propose empty block otherwise.
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx.Context(), prepareTimeout)
	defer timeoutCancel()
	ctx = ctx.WithContext(timeoutCtx)

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

	appHash, err := cast.EthHash(ctx.BlockHeader().AppHash)
	if err != nil {
		return nil, err
	}

	reqHeight, err := umath.ToUint64(req.Height)
	if err != nil {
		return nil, err
	}

	// Either use the optimistic payload or create a new one.
	payloadID, height, triggeredAt := k.getOptimisticPayload()
	if reqHeight != height { //nolint:nestif // Not an issue
		// Create a new payload (retrying on network errors).
		err := retryForever(ctx, func(ctx context.Context) (bool, error) {
			fcr, err := k.startBuild(ctx, appHash, req.Time, reqHeight)
			if err != nil {
				log.Warn(ctx, "Preparing proposal failed: build new evm payload (will retry)", err)
				return false, nil // Retry
			} else if isSyncing(fcr.PayloadStatus) {
				return false, errors.New("evm unexpectedly syncing") // Abort, don't retry
			} else if invalid, err := isInvalid(fcr.PayloadStatus); invalid {
				return false, errors.Wrap(err, "proposed invalid payload") // Abort, don't retry
			} else if fcr.PayloadID == nil {
				return false, errors.New("missing payload ID [BUG]") // Abort, don't retry
			} /* else isValid(status) */

			payloadID = *fcr.PayloadID

			return true, nil // Done
		})
		if err != nil {
			return nil, err
		}
		triggeredAt = time.Now()
		log.Debug(ctx, "Started non-optimistic payload", "height", req.Height, "payload", payloadID.String())
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
	err = retryForever(ctx, func(ctx context.Context) (bool, error) {
		var err error
		payloadResp, err = k.engineCl.GetPayloadV3(ctx, payloadID)
		if isUnknownPayload(err) {
			return false, err // Abort, don't retry
		} else if err != nil {
			log.Warn(ctx, "Preparing proposal failed: get evm payload (will retry)", err)
			return false, nil // Retry
		}

		return true, nil // Done
	})
	if err != nil {
		return nil, err
	}

	payloadResp = maybeFuzzPayload(ctx, payloadResp)

	// Create execution payload message
	payloadProto, err := types.PayloadToProto(payloadResp.ExecutionPayload)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	// Convert blobs bundle.
	blobCommitments := unwrapHexBytes(payloadResp.BlobsBundle.Commitments)
	if _, err := blobHashes(blobCommitments); err != nil && !feature.FlagFuzzOctane.Enabled(ctx) { // Sanity check blobs are valid.
		return nil, errors.Wrap(err, "invalid blobs [BUG]")
	}

	// First, collect all vote extension msgs from the vote provider.
	prevHeight := umath.SubtractOrZero(reqHeight, 1) // Safe to subtract one since this is only called AFTER genesis.
	voteMsgs, err := k.voteProvider.PrepareVotes(ctx, req.LocalLastCommit, prevHeight)
	if err != nil {
		return nil, errors.Wrap(err, "prepare votes")
	}

	// Then construct the execution payload message.
	payloadMsg := &types.MsgExecutionPayload{
		Authority:             authtypes.NewModuleAddress(types.ModuleName).String(),
		BlobCommitments:       blobCommitments,
		ExecutionPayloadDeneb: payloadProto,
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
	)

	return &abci.ResponsePrepareProposal{Txs: [][]byte{tx}}, nil
}

// maybeFuzzPayload applies a random mutation to the payload for fuzz testing if enabled.
func maybeFuzzPayload(ctx sdk.Context, resp *engine.ExecutionPayloadEnvelope) *engine.ExecutionPayloadEnvelope {
	if !feature.FlagFuzzOctane.Enabled(ctx) {
		return resp
	}

	switch rand.IntN(5) { //nolint:gosec // Weak RNG is fine for fuzz testing.
	case 0:
		log.Warn(ctx, "Fuzzing proposed octane payload: invalid parent hash", nil)
		resp.ExecutionPayload.BlockHash = resp.ExecutionPayload.ParentHash
	case 1:
		log.Warn(ctx, "Fuzzing proposed octane payload: invalid withdrawal", nil)
		resp.ExecutionPayload.Withdrawals = append(resp.ExecutionPayload.Withdrawals, &etypes.Withdrawal{
			Index:     1,
			Validator: 2,
			Amount:    3,
		})
	case 2:
		log.Warn(ctx, "Fuzzing proposed octane payload: nil blob gas", nil)
		resp.ExecutionPayload.ExcessBlobGas = nil
		resp.ExecutionPayload.BlobGasUsed = nil
	case 3:
		log.Warn(ctx, "Fuzzing proposed octane payload: invalid tx", nil)
		resp.ExecutionPayload.Transactions = append(resp.ExecutionPayload.Transactions, []byte("invalid tx"))
	case 4:
		log.Warn(ctx, "Fuzzing proposed octane payload: invalid blobs bundle", nil)
		if resp.BlobsBundle == nil {
			resp.BlobsBundle = &engine.BlobsBundleV1{}
		}
		for len(resp.BlobsBundle.Commitments) < maxBlobsPerBlock {
			var commit [48]byte
			resp.BlobsBundle.Commitments = append(resp.BlobsBundle.Commitments, commit[:])
		}
	}

	return resp
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
	timestamp := ctx.BlockTime()
	appHash, err := cast.EthHash(ctx.BlockHeader().AppHash) // This is the app hash after the block is finalized.
	if err != nil {
		return err
	}

	// Maybe start building the next block if we are the next proposer.
	isNext, err := k.isNextProposer(ctx, height)
	if err != nil {
		// IsNextProposer does non-deterministic cometBFT queries, don't stall node due to errors.
		log.Warn(ctx, "Next proposer failed, skipping optimistic EVM payload build", err)

		return nil
	} else if !isNext {
		return nil // Nothing to do if we are not next proposer.
	}

	nextHeight, err := umath.ToUint64(height + 1)
	if err != nil {
		return err
	}

	logAttr := slog.Uint64("next_height", nextHeight)
	log.Debug(ctx, "Starting optimistic EVM payload build", logAttr)

	// No need to wrap this in retryForever since this is a best-effort optimisation, if it fails, just skip it.
	fcr, err := k.startBuild(ctx, appHash, timestamp, nextHeight)
	if err != nil {
		log.Warn(ctx, "Starting optimistic build failed", err, logAttr)
		return nil
	} else if isSyncing(fcr.PayloadStatus) {
		log.Warn(ctx, "Starting optimistic build failed; evm syncing", nil, logAttr)
		return nil
	} else if invalid, err := isInvalid(fcr.PayloadStatus); invalid {
		log.Error(ctx, "Starting optimistic build failed; invalid payload [BUG]", err, logAttr)
		return nil
	} else if fcr.PayloadID == nil {
		log.Error(ctx, "Starting optimistic build failed; missing payload ID [BUG]", nil, logAttr)
		return nil
	} /* else isValid(status) */

	k.setOptimisticPayload(*fcr.PayloadID, nextHeight)

	return nil
}

// startBuild triggers the building of a new execution payload on top of the current execution head.
// It returns the EngineAPI response which contains a status and payload ID.
func (k *Keeper) startBuild(ctx context.Context, appHash common.Hash, timestamp time.Time, height uint64) (engine.ForkChoiceResponse, error) {
	head, err := k.getExecutionHead(ctx)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "latest execution block")
	}

	// Use provided time as timestamp for the next block.
	// Or use latest execution block timestamp + 1 if is not greater.
	// Since execution blocks must have unique second-granularity timestamps.
	ts, err := umath.ToUint64(timestamp.Unix())
	if err != nil {
		return engine.ForkChoiceResponse{}, err
	}

	if ts <= head.GetBlockTime() {
		ts = head.GetBlockTime() + 1 // Subsequent blocks must have a higher timestamp.
	}

	headHash, err := head.Hash()
	if err != nil {
		return engine.ForkChoiceResponse{}, err
	}

	// CometBFT has instant finality, so head/safe/finalized is latest height.
	fcs := engine.ForkchoiceStateV1{
		HeadBlockHash:      headHash,
		SafeBlockHash:      headHash,
		FinalizedBlockHash: headHash,
	}

	withdrawals, err := k.eligibleWithdrawals(ctx, height)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "eligible withdrawals")
	}

	attrs := &engine.PayloadAttributes{
		Timestamp:             ts,
		Random:                headHash, // We use head block hash as randao.
		SuggestedFeeRecipient: k.feeRecProvider.LocalFeeRecipient(),
		Withdrawals:           withdrawals,
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
