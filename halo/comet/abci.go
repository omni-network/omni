package comet

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/errors"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	// version of the Halo application wrt cometBFT.
	appVersion = 0
)

// Info returns information about the application state.
// V0 in-memory chain always starts from scratch, at height 0.
func (c *App) Info(_ context.Context, req *abci.RequestInfo) (*abci.ResponseInfo, error) {
	return &abci.ResponseInfo{
		Data:             "", // CometBFT does not use this field.
		Version:          req.AbciVersion,
		AppVersion:       appVersion,
		LastBlockHeight:  int64(c.state.Height()),
		LastBlockAppHash: c.state.Hash(), // AppHash overwritten by InitChain if LastBlockHeight==0.
	}, nil
}

// InitChain initializes the blockchain.
func (c *App) InitChain(_ context.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	if req.InitialHeight > 1 {
		return nil, errors.New("initial height must not 1")
	}

	if len(req.AppStateBytes) > 0 {
		if err := c.state.Import(0, req.AppStateBytes); err != nil {
			return nil, errors.Wrap(err, "import state")
		}
	}

	if err := c.state.InitValidators(req.Validators); err != nil {
		return nil, errors.Wrap(err, "set validators")
	}

	return &abci.ResponseInitChain{
		AppHash: c.state.Hash(),
		// Return nils below to indicate no-update.
		ConsensusParams: nil,
		Validators:      nil,
	}, nil
}

// PrepareProposal returns a proposal for the next block.
// Note returning an error results in a panic cometbft and CONSENSUS_FAILURE log.
func (c *App) PrepareProposal(ctx context.Context, req *abci.RequestPrepareProposal) (
	*abci.ResponsePrepareProposal, error,
) {
	if len(req.Txs) > 0 {
		return nil, errors.New("unexpected transactions in proposal")
	}

	latestEHeight, err := c.ethCl.BlockNumber(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "latest execution block number")
	}

	latestEBlock, err := c.ethCl.BlockByNumber(ctx, big.NewInt(int64(latestEHeight)))
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

	forkchoiceResp, err := c.ethCl.ForkchoiceUpdatedV2(ctx, forkchoiceState, &payloadAttrs)
	if err != nil {
		return nil, err
	} else if forkchoiceResp.PayloadStatus.Status != engine.VALID {
		return nil, errors.New("status not valid")
	}

	payloadResp, err := c.ethCl.GetPayloadV2(ctx, *forkchoiceResp.PayloadID)
	if err != nil {
		return nil, err
	}
	// The previous height's vote extensions are provided to the proposer in
	// the requests last local commit. Simply add all vote extensions from the
	// previous height into the CPayload.
	aggs, err := aggregatesFromProposal(req)
	if err != nil {
		return nil, err
	}

	tx, err := encode(cPayload{
		EPayload:   *payloadResp.ExecutionPayload,
		Aggregates: aggs,
	})
	if err != nil {
		return nil, err
	}

	return &abci.ResponsePrepareProposal{Txs: [][]byte{tx}}, nil
}

// ProcessProposal validates a proposal.
func (c *App) ProcessProposal(ctx context.Context, req *abci.RequestProcessProposal) (
	*abci.ResponseProcessProposal, error,
) {
	cpayload, err := payloadFromTXs(req.Txs)
	if err != nil {
		return nil, err
	}

	// Push it back to the execution client (mark it as possible new head).
	status, err := c.ethCl.NewPayloadV2(ctx, cpayload.EPayload)
	if err != nil {
		return nil, err
	} else if status.Status != engine.VALID {
		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
	}

	// Mark all local attestations as "proposed", i.e., included in latest proposed block.
	localHeaders := headersByPubKey(cpayload.Aggregates, c.attestSvc.LocalPubKey())
	c.attestSvc.SetProposed(localHeaders)

	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}

// ExtendVote extends a vote with application-injected data (vote extensions).
func (c *App) ExtendVote(context.Context, *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	attBytes, err := encode(c.attestSvc.GetAvailable())
	if err != nil {
		return nil, err
	}

	return &abci.ResponseExtendVote{
		VoteExtension: attBytes,
	}, nil
}

// VerifyVoteExtension verifies a vote extension.
func (*App) VerifyVoteExtension(context.Context, *abci.RequestVerifyVoteExtension) (
	*abci.ResponseVerifyVoteExtension, error,
) {
	// TODO(corver): Figure out what to verify.
	return &abci.ResponseVerifyVoteExtension{
		Status: abci.ResponseVerifyVoteExtension_ACCEPT,
	}, nil
}

// FinalizeBlock finalizes a block.
func (c *App) FinalizeBlock(ctx context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	cpayload, err := payloadFromTXs(req.Txs)
	if err != nil {
		return nil, err
	}

	fcs := engine.ForkchoiceStateV1{
		HeadBlockHash:      cpayload.EPayload.BlockHash,
		SafeBlockHash:      cpayload.EPayload.BlockHash,
		FinalizedBlockHash: cpayload.EPayload.BlockHash,
	}

	forchainResp, err := c.ethCl.ForkchoiceUpdatedV2(ctx, fcs, nil)
	if err != nil {
		return nil, err
	} else if forchainResp.PayloadStatus.Status != engine.VALID {
		return nil, errors.New("status not valid")
	}

	c.state.AddAttestations(cpayload.Aggregates)

	// Mark all local attestations "committed", i.e., included in this committed block.
	localHeaders := headersByPubKey(cpayload.Aggregates, c.attestSvc.LocalPubKey())
	c.attestSvc.SetCommitted(localHeaders)

	appHash, err := c.state.Finalize()
	if err != nil {
		return nil, err
	}

	return &abci.ResponseFinalizeBlock{
		Events: nil, // Events are going to be deprecated from cometBFT.
		TxResults: []*abci.ExecTxResult{{
			Code: abci.CodeTypeOK, // Single zero/ok result is fine.
		}},
		ValidatorUpdates:      nil, // Validator updates not supported yet.
		ConsensusParamUpdates: nil, // ConsensusParam updates not supported yet.
		AppHash:               appHash[:],
	}, nil
}

// Commit commits the state. It also creates a snapshot sometimes.
func (c *App) Commit(context.Context, *abci.RequestCommit) (*abci.ResponseCommit, error) {
	height, err := c.state.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "commit state")
	}

	if c.snapshotInterval > 0 && height%c.snapshotInterval == 0 {
		_, err := c.snapshots.Create(c.state)
		if err != nil {
			return nil, errors.Wrap(err, "create snapshot")
		}

		err = c.snapshots.Prune()
		if err != nil {
			return nil, errors.Wrap(err, "prune snapshots")
		}
	}

	return &abci.ResponseCommit{
		RetainHeight: 0, // Retain all blocks.
	}, nil
}

// ListSnapshots lists all the available snapshots.
func (c *App) ListSnapshots(context.Context, *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	var resp abci.ResponseListSnapshots
	for _, snapshot := range c.snapshots.List() {
		snapshot := snapshot // Pin.
		resp.Snapshots = append(resp.Snapshots, &snapshot)
	}

	return &resp, nil
}

// OfferSnapshot sends a snapshot offer.
func (c *App) OfferSnapshot(_ context.Context, req *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	c.restore.Lock()
	defer c.restore.Unlock()

	if c.restore.Snapshot != nil {
		return nil, errors.New("snapshot already offered")
	}

	c.restore.Snapshot = req.Snapshot

	return &abci.ResponseOfferSnapshot{
		Result: abci.ResponseOfferSnapshot_ACCEPT,
	}, nil
}

// ApplySnapshotChunk applies a chunk of snapshot.
func (c *App) ApplySnapshotChunk(_ context.Context, req *abci.RequestApplySnapshotChunk) (
	*abci.ResponseApplySnapshotChunk, error,
) {
	c.restore.Lock()
	defer c.restore.Unlock()

	if c.restore.Snapshot == nil {
		return nil, errors.New("no snapshot offered")
	}

	c.restore.Chunks = append(c.restore.Chunks, req.Chunk)

	if len(c.restore.Chunks) < int(c.restore.Snapshot.Chunks) {
		return &abci.ResponseApplySnapshotChunk{Result: abci.ResponseApplySnapshotChunk_ACCEPT}, nil
	}

	bz := make([]byte, 0, c.restore.Snapshot.Chunks*snapshotChunkSize)
	for _, chunk := range c.restore.Chunks {
		bz = append(bz, chunk...)
	}

	err := c.state.Import(c.restore.Snapshot.Height, bz)
	if err != nil {
		return nil, errors.Wrap(err, "import state")
	}

	c.restore.Snapshot = nil
	c.restore.Chunks = nil

	return &abci.ResponseApplySnapshotChunk{Result: abci.ResponseApplySnapshotChunk_ACCEPT}, nil
}

// LoadSnapshotChunk returns a chunk of snapshot.
func (c *App) LoadSnapshotChunk(_ context.Context, req *abci.RequestLoadSnapshotChunk) (
	*abci.ResponseLoadSnapshotChunk, error,
) {
	chunk, err := c.snapshots.LoadChunk(req.Height, req.Format, req.Chunk)
	if err != nil {
		return nil, errors.Wrap(err, "load snapshot chunk")
	}

	return &abci.ResponseLoadSnapshotChunk{
		Chunk: chunk,
	}, nil
}

// TODO(corver): Implement the following logic.

// Flush flushes the write buffer.
func (*App) Flush(context.Context, *abci.RequestFlush) (*abci.ResponseFlush, error) {
	return nil, nil //nolint:nilnil // In-memory state, nothing to flush.
}

// Query queries the application state.
func (*App) Query(context.Context, *abci.RequestQuery) (*abci.ResponseQuery, error) {
	return nil, errors.New("queries not supported yet")
}

// Echo returns back the same message it is sent.
func (*App) Echo(_ context.Context, req *abci.RequestEcho) (*abci.ResponseEcho, error) {
	return &abci.ResponseEcho{
		Message: req.Message,
	}, nil
}

// CheckTx validates a transaction.
func (*App) CheckTx(context.Context, *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	return nil, errors.New("unexpected CheckTx request")
}
