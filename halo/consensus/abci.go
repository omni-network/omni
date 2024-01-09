package consensus

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	abci "github.com/cometbft/cometbft/api/cometbft/abci/v1"
)

const (
	// version of the Halo application wrt cometBFT.
	appVersion = 0

	// startHeight of v0 in-memory chain is always 0.
	startHeight = 0
)

// Info returns information about the application state.
// V0 in-memory chain always starts from scratch, at height 0.
func (*Core) Info(_ context.Context, req *abci.InfoRequest) (*abci.InfoResponse, error) {
	return &abci.InfoResponse{
		Data:             "", // CometBFT does not use this field.
		Version:          req.AbciVersion,
		AppVersion:       appVersion,
		LastBlockHeight:  startHeight,
		LastBlockAppHash: nil, // AppHash overwritten by InitChain if LastBlockHeight==0.
	}, nil
}

// InitChain initializes the blockchain.
func (c *Core) InitChain(_ context.Context, req *abci.InitChainRequest) (*abci.InitChainResponse, error) {
	if err := c.state.InitChainState(req); err != nil {
		return nil, errors.Wrap(err, "failed to set validators")
	}

	appHash, err := c.state.AppHash()
	if err != nil {
		return nil, errors.Wrap(err, "failed to compute app hash")
	}

	return &abci.InitChainResponse{
		AppHash: appHash,
		// Return nils below to indicate no-update.
		ConsensusParams: nil,
		Validators:      nil,
	}, nil
}

// PrepareProposal returns a proposal for the next block.
func (*Core) PrepareProposal(_ context.Context, req *abci.PrepareProposalRequest) (
	*abci.PrepareProposalResponse, error,
) {
	if len(req.Txs) > 0 {
		return nil, errors.New("unexpected transactions in proposal")
	}

	// TODO(corver): Build a EngineAPI block and add it to the cpayload.

	// The previous height's vote extensions are provided to the proposer in
	// the requests last local commit. Simply add all vote extensions from the
	// previous height into the CPayload.
	aggs, err := aggregatesFromProposal(req)
	if err != nil {
		return nil, err
	}

	tx, err := encode(cpayload{
		Aggregates: aggs,
	})
	if err != nil {
		return nil, err
	}

	return &abci.PrepareProposalResponse{Txs: [][]byte{tx}}, nil
}

// ProcessProposal validates a proposal.
func (c *Core) ProcessProposal(_ context.Context, req *abci.ProcessProposalRequest) (
	*abci.ProcessProposalResponse, error,
) {
	payload, err := payloadFromTXs(req.Txs)
	if err != nil {
		return nil, err
	}

	// TODO(corver): Submit the execution cpayload to the EngineAPI.

	// Mark all local attestations as "proposed", i.e., included in latest proposed block.
	localHeaders := headersByPubkey(payload.Aggregates, c.attestSvc.LocalPubKey())
	c.attestSvc.SetProposed(localHeaders)

	return &abci.ProcessProposalResponse{Status: abci.PROCESS_PROPOSAL_STATUS_ACCEPT}, nil
}

// ExtendVote extends a vote with application-injected data (vote extensions).
func (c *Core) ExtendVote(context.Context, *abci.ExtendVoteRequest) (*abci.ExtendVoteResponse, error) {
	attBytes, err := encode(c.attestSvc.GetAvailable())
	if err != nil {
		return nil, err
	}

	return &abci.ExtendVoteResponse{
		VoteExtension: attBytes,
	}, nil
}

// VerifyVoteExtension verifies a vote extension.
func (*Core) VerifyVoteExtension(context.Context, *abci.VerifyVoteExtensionRequest) (
	*abci.VerifyVoteExtensionResponse, error,
) {
	// TODO(corver): Figure out what to verify.
	return &abci.VerifyVoteExtensionResponse{
		Status: abci.VERIFY_VOTE_EXTENSION_STATUS_ACCEPT,
	}, nil
}

// FinalizeBlock finalizes a block.
func (c *Core) FinalizeBlock(_ context.Context, req *abci.FinalizeBlockRequest) (*abci.FinalizeBlockResponse, error) {
	payload, err := payloadFromTXs(req.Txs)
	if err != nil {
		return nil, err
	}

	// TODO(corver): update EngineAPI forkchoice.

	c.state.AddAttestations(payload.Aggregates)

	// Mark all local attestations "committed", i.e., included in this committed block.
	localHeaders := headersByPubkey(payload.Aggregates, c.attestSvc.LocalPubKey())
	c.attestSvc.SetCommitted(localHeaders)

	appHash, err := c.state.AppHash()
	if err != nil {
		return nil, err
	}

	return &abci.FinalizeBlockResponse{
		Events: nil, // Events are going to be deprecated from cometBFT.
		TxResults: []*abci.ExecTxResult{{
			Code: abci.CodeTypeOK, // Single zero/ok result is fine.
		}},
		ValidatorUpdates:      nil, // Validator updates not supported yet.
		ConsensusParamUpdates: nil, // ConsensusParam updates not supported yet.
		AppHash:               appHash,
	}, nil
}

// TODO(corver): Implement the following logic.

// Flush flushes the write buffer.
func (*Core) Flush(context.Context, *abci.FlushRequest) (*abci.FlushResponse, error) {
	return nil, nil //nolint:nilnil // In-memory state, nothing to flush.
}

// Commit commits a block of transactions.
func (*Core) Commit(context.Context, *abci.CommitRequest) (*abci.CommitResponse, error) {
	return &abci.CommitResponse{}, nil // In-memory state, nothing to commit.
}

// Query queries the application state.
func (*Core) Query(context.Context, *abci.QueryRequest) (*abci.QueryResponse, error) {
	return nil, errors.New("queries not supported yet")
}

// ListSnapshots lists all the available snapshots.
func (*Core) ListSnapshots(context.Context, *abci.ListSnapshotsRequest) (*abci.ListSnapshotsResponse, error) {
	return nil, errors.New("snapshots not supported yet")
}

// OfferSnapshot sends a snapshot offer.
func (*Core) OfferSnapshot(context.Context, *abci.OfferSnapshotRequest) (*abci.OfferSnapshotResponse, error) {
	return nil, errors.New("snapshots not supported yet")
}

// LoadSnapshotChunk returns a chunk of snapshot.
func (*Core) LoadSnapshotChunk(context.Context, *abci.LoadSnapshotChunkRequest) (
	*abci.LoadSnapshotChunkResponse, error,
) {
	return nil, errors.New("snapshots not supported yet")
}

// ApplySnapshotChunk applies a chunk of snapshot.
func (*Core) ApplySnapshotChunk(context.Context, *abci.ApplySnapshotChunkRequest) (
	*abci.ApplySnapshotChunkResponse, error,
) {
	return nil, errors.New("snapshots not supported yet")
}

// Echo returns back the same message it is sent.
func (*Core) Echo(_ context.Context, req *abci.EchoRequest) (*abci.EchoResponse, error) {
	return &abci.EchoResponse{
		Message: req.Message,
	}, nil
}

// CheckTx validates a transaction.
func (*Core) CheckTx(context.Context, *abci.CheckTxRequest) (*abci.CheckTxResponse, error) {
	return nil, errors.New("unexpected CheckTx request")
}
