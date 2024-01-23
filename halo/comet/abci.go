package comet

import (
	"context"
	"math/big"

	halopb "github.com/omni-network/omni/halo/halopb/v1"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	"google.golang.org/protobuf/proto"
)

const (
	// version of the Halo application wrt cometBFT.
	appVersion = 0
)

// Info returns information about the application state.
// V0 in-memory chain always starts from scratch, at height 0.
func (a *App) Info(ctx context.Context, req *abci.RequestInfo) (*abci.ResponseInfo, error) {
	resp := &abci.ResponseInfo{
		Data:             "", // CometBFT does not use this field.
		Version:          req.AbciVersion,
		AppVersion:       appVersion,
		LastBlockHeight:  int64(a.state.Height()),
		LastBlockAppHash: a.state.Hash(), // AppHash overwritten by InitChain if LastBlockHeight==0.
	}

	log.Info(ctx, "Starting consensus chain",
		"last_height", resp.LastBlockHeight,
		log.Hex7("app_hash", resp.LastBlockAppHash),
	)

	return resp, nil
}

// InitChain initializes the blockchain.
func (a *App) InitChain(ctx context.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	if req.InitialHeight > 1 {
		return nil, errors.New("initial height must not 1")
	}

	if len(req.AppStateBytes) > 0 {
		if err := a.state.Import(0, req.AppStateBytes); err != nil {
			return nil, errors.Wrap(err, "import state")
		}
	}

	if err := a.state.InitValidators(req.Validators); err != nil {
		return nil, errors.Wrap(err, "set validators")
	}

	resp := &abci.ResponseInitChain{
		AppHash: a.state.Hash(),
		// Return nils below to indicate no-update.
		ConsensusParams: nil,
		Validators:      nil,
	}

	log.Info(ctx, "Initializing brand new consensus chain",
		"init_height", req.InitialHeight,
		"validators", len(req.Validators),
		log.Hex7("genesis_app_hash", resp.AppHash),
	)

	return resp, nil
}

// PrepareProposal returns a proposal for the next block.
// Note returning an error results in a panic cometbft and CONSENSUS_FAILURE log.
func (a *App) PrepareProposal(ctx context.Context, req *abci.RequestPrepareProposal) (
	*abci.ResponsePrepareProposal, error,
) {
	if len(req.Txs) > 0 {
		return nil, errors.New("unexpected transactions in proposal")
	}

	latestEHeight, err := a.ethCl.BlockNumber(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "latest execution block number")
	}

	latestEBlock, err := a.ethCl.BlockByNumber(ctx, big.NewInt(int64(latestEHeight)))
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

	forkchoiceResp, err := a.ethCl.ForkchoiceUpdatedV2(ctx, forkchoiceState, &payloadAttrs)
	if err != nil {
		return nil, err
	} else if forkchoiceResp.PayloadStatus.Status != engine.VALID {
		return nil, errors.New("status not valid")
	}

	payloadResp, err := a.ethCl.GetPayloadV2(ctx, *forkchoiceResp.PayloadID)
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

	log.Info(ctx, "Proposing new block",
		"height", req.Height,
		log.Hex7("execution_block_hash", payloadResp.ExecutionPayload.BlockHash[:]),
	)

	return &abci.ResponsePrepareProposal{Txs: [][]byte{tx}}, nil
}

// ProcessProposal validates a proposal.
func (a *App) ProcessProposal(ctx context.Context, req *abci.RequestProcessProposal) (
	*abci.ResponseProcessProposal, error,
) {
	cpayload, err := payloadFromTXs(req.Txs)
	if err != nil {
		return nil, err
	}

	// Push it back to the execution client (mark it as possible new head).
	status, err := a.ethCl.NewPayloadV2(ctx, cpayload.EPayload)
	if err != nil {
		return nil, err
	} else if status.Status != engine.VALID {
		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
	}

	// Mark all local attestations as "proposed", i.e., included in latest proposed block.
	localHeaders := headersByPubKey(cpayload.Aggregates, a.attestSvc.LocalPubKey())
	if err := a.attestSvc.SetProposed(localHeaders); err != nil {
		return nil, err
	}

	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}

// ExtendVote extends a vote with application-injected data (vote extensions).
func (a *App) ExtendVote(ctx context.Context, _ *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	atts := a.attestSvc.GetAvailable()

	attBytes, err := encode(atts)
	if err != nil {
		return nil, err
	}

	log.Info(ctx, "Attesting to rollup blocks", "attestations", len(atts))

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
func (a *App) FinalizeBlock(ctx context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	cpayload, err := payloadFromTXs(req.Txs)
	if err != nil {
		return nil, err
	}

	fcs := engine.ForkchoiceStateV1{
		HeadBlockHash:      cpayload.EPayload.BlockHash,
		SafeBlockHash:      cpayload.EPayload.BlockHash,
		FinalizedBlockHash: cpayload.EPayload.BlockHash,
	}

	forchainResp, err := a.ethCl.ForkchoiceUpdatedV2(ctx, fcs, nil)
	if err != nil {
		return nil, err
	} else if forchainResp.PayloadStatus.Status != engine.VALID {
		return nil, errors.New("status not valid")
	}

	a.state.AddAttestations(cpayload.Aggregates)

	// Mark all local attestations "committed", i.e., included in this committed block.
	localHeaders := headersByPubKey(cpayload.Aggregates, a.attestSvc.LocalPubKey())
	if err := a.attestSvc.SetCommitted(localHeaders); err != nil {
		return nil, err
	}

	appHash, err := a.state.Finalize()
	if err != nil {
		return nil, err
	}

	log.Info(ctx, "Finalized new consensus block",
		"height", req.Height,
		log.Hex7("app_hash", appHash[:]),
		"attestations", len(cpayload.Aggregates),
		log.Hex7("execution_block_hash", cpayload.EPayload.BlockHash[:]),
	)

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
func (a *App) Commit(context.Context, *abci.RequestCommit) (*abci.ResponseCommit, error) {
	height, err := a.state.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "commit state")
	}

	if a.snapshotInterval > 0 && height%a.snapshotInterval == 0 {
		_, err := a.snapshots.Create(a.state)
		if err != nil {
			return nil, errors.Wrap(err, "create snapshot")
		}

		err = a.snapshots.Prune()
		if err != nil {
			return nil, errors.Wrap(err, "prune snapshots")
		}
	}

	return &abci.ResponseCommit{
		RetainHeight: 0, // Retain all blocks.
	}, nil
}

// ListSnapshots lists all the available snapshots.
func (a *App) ListSnapshots(context.Context, *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	var resp abci.ResponseListSnapshots
	for _, snapshot := range a.snapshots.List() {
		snapshot := snapshot // Pin.
		resp.Snapshots = append(resp.Snapshots, &snapshot)
	}

	return &resp, nil
}

// OfferSnapshot sends a snapshot offer.
func (a *App) OfferSnapshot(_ context.Context, req *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	a.restore.Lock()
	defer a.restore.Unlock()

	if a.restore.Snapshot != nil {
		return nil, errors.New("snapshot already offered")
	}

	a.restore.Snapshot = req.Snapshot

	return &abci.ResponseOfferSnapshot{
		Result: abci.ResponseOfferSnapshot_ACCEPT,
	}, nil
}

// ApplySnapshotChunk applies a chunk of snapshot.
func (a *App) ApplySnapshotChunk(_ context.Context, req *abci.RequestApplySnapshotChunk) (
	*abci.ResponseApplySnapshotChunk, error,
) {
	a.restore.Lock()
	defer a.restore.Unlock()

	if a.restore.Snapshot == nil {
		return nil, errors.New("no snapshot offered")
	}

	a.restore.Chunks = append(a.restore.Chunks, req.Chunk)

	if len(a.restore.Chunks) < int(a.restore.Snapshot.Chunks) {
		return &abci.ResponseApplySnapshotChunk{Result: abci.ResponseApplySnapshotChunk_ACCEPT}, nil
	}

	bz := make([]byte, 0, a.restore.Snapshot.Chunks*snapshotChunkSize)
	for _, chunk := range a.restore.Chunks {
		bz = append(bz, chunk...)
	}

	err := a.state.Import(a.restore.Snapshot.Height, bz)
	if err != nil {
		return nil, errors.Wrap(err, "import state")
	}

	a.restore.Snapshot = nil
	a.restore.Chunks = nil

	return &abci.ResponseApplySnapshotChunk{Result: abci.ResponseApplySnapshotChunk_ACCEPT}, nil
}

// LoadSnapshotChunk returns a chunk of snapshot.
func (a *App) LoadSnapshotChunk(_ context.Context, req *abci.RequestLoadSnapshotChunk) (
	*abci.ResponseLoadSnapshotChunk, error,
) {
	chunk, err := a.snapshots.LoadChunk(req.Height, req.Format, req.Chunk)
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
func (a *App) Query(_ context.Context, query *abci.RequestQuery) (*abci.ResponseQuery, error) {
	if query == nil || len(query.Data) == 0 {
		return nil, errors.New("empty query")
	} else if query.Path != halopb.HaloService_ApprovedFrom_FullMethodName {
		return nil, errors.New("unknown query path")
	}

	// Unmarshal the request.
	req := new(halopb.ApprovedFromRequest)
	if err := proto.Unmarshal(query.Data, req); err != nil {
		return nil, errors.Wrap(err, "unmarshal approved from request")
	}

	// Query the state.
	aggs := a.state.ApprovedFrom(req.GetChainId(), req.GetFromHeight())

	// Construct the response.
	resp := &halopb.ApprovedFromResponse{
		Aggregates: halopb.AggregatesToProto(aggs),
	}

	// Marshal the response.
	bz, err := proto.Marshal(resp)
	if err != nil {
		return nil, errors.Wrap(err, "marshal approved from response")
	}

	// Return the response.
	return &abci.ResponseQuery{
		Code:   abci.CodeTypeOK,
		Key:    query.Data,
		Value:  bz,
		Height: int64(a.state.Height()),
	}, nil
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
