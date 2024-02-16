package keeper

import (
	"math/big"

	"github.com/omni-network/omni/halo2/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PrepareProposal returns a proposal for the next block.
// Note returning an error results in a panic cometbft and CONSENSUS_FAILURE log.
func (k Keeper) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (
	*abci.ResponsePrepareProposal, error,
) {
	defer func() {
		// Recover panics
		if r := recover(); r != nil {
			log.Error(ctx, "PrepareProposal panic", nil, "recover", r)
			panic(r)
		}
	}()
	if len(req.Txs) > 0 {
		return nil, errors.New("unexpected transactions in proposal")
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
	payloadData, err := encode(payloadResp.ExecutionPayload)
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	payloadMsg := types.MsgExecutionPayload{
		Authority: "test",
		Data:      payloadData,
	}

	// The previous height's vote extensions are provided to the proposer in
	// the requests last local commit. Simply add all vote extensions from the
	// previous height into the CPayload.
	_, err = aggregatesFromProposal(req.LocalLastCommit)
	if err != nil {
		return nil, err
	}

	b := k.txConfig.NewTxBuilder()

	// Create "CPayload" transaction.
	// It contains ALL the cosmos messages to be included in this block.
	// - engineevm.MsgExecutionPayload    // The execution payload. Must be included in all blocks.
	// - []attest.MsgAggregateAttestation // The aggregate attestations.
	// The following msgs are extracted from the EVM
	// - []staking.MsgCreateValidator     // New native $OMNI staking validator
	// - []staking.MsgDelegate            // Restaked $ETH delegation
	// - []staking.MsgUndelegate 		  // Unstaked $ETH delegation
	if err := b.SetMsgs(&payloadMsg); err != nil {
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
	)

	return &abci.ResponsePrepareProposal{Txs: [][]byte{tx}}, nil
}

func (k Keeper) ProcessProposal(ctx sdk.Context, req *abci.RequestProcessProposal,
) (*abci.ResponseProcessProposal, error) {
	defer func() {
		// Recover panics
		if r := recover(); r != nil {
			log.Error(ctx, "ProcessProposal panic", nil, "recover", r)
			panic(r)
		}
	}()

	if len(req.Txs) != 1 {
		return nil, errors.New("expected 1 transaction in proposal")
	}

	tx, err := k.txConfig.TxDecoder()(req.Txs[0])
	if err != nil {
		return nil, errors.Wrap(err, "decode tx")
	}

	// Extract the staking messages from the EVM (to compare against proposed).
	for _, msg := range tx.GetMsgs() {
		switch m := msg.(type) {
		case *types.MsgExecutionPayload:

			var payload engine.ExecutableData
			if err := decode(m.Data, &payload); err != nil {
				return nil, errors.Wrap(err, "decode payload")
			}

			// Push it back to the execution client (mark it as possible new head).
			status, err := k.ethCl.NewPayloadV2(ctx, payload)
			if err != nil {
				return nil, errors.Wrap(err, "new payload")
			} else if status.Status != engine.VALID {
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
			}

		// Verify the following messages
		// case attest.MsgAggregateAttestation:
		//    Ensure valid and within window
		//    If local, mark as proposed.
		// case staking.MsgCreateValidator:
		// case staking.MsgDelegate:
		// case staking.MsgUndelegate:
		default:
			return nil, errors.New("unexpected message type")
		}
	}

	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}
