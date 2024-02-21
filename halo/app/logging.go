//nolint:wrapcheck // The abci.Application is our application, so we don't need to wrap it. Long lines are fine here.
package app

import (
	"context"

	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"
)

type loggingABCIApp struct {
	abci.Application
}

func (l loggingABCIApp) Info(ctx context.Context, info *abci.RequestInfo) (*abci.ResponseInfo, error) {
	log.Debug(ctx, "👾 ABCI call: Info")
	return l.Application.Info(ctx, info)
}

func (l loggingABCIApp) Query(ctx context.Context, query *abci.RequestQuery) (*abci.ResponseQuery, error) {
	return l.Application.Query(ctx, query) // No log here since this can be very noisy
}

func (l loggingABCIApp) CheckTx(ctx context.Context, tx *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	log.Debug(ctx, "👾 ABCI call: CheckTx")
	return l.Application.CheckTx(ctx, tx)
}

func (l loggingABCIApp) InitChain(ctx context.Context, chain *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	log.Debug(ctx, "👾 ABCI call: InitChain")
	return l.Application.InitChain(ctx, chain)
}

func (l loggingABCIApp) PrepareProposal(ctx context.Context, proposal *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	log.Debug(ctx, "👾 ABCI call: PrepareProposal")
	return l.Application.PrepareProposal(ctx, proposal)
}

func (l loggingABCIApp) ProcessProposal(ctx context.Context, proposal *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	log.Debug(ctx, "👾 ABCI call: ProcessProposal")
	return l.Application.ProcessProposal(ctx, proposal)
}

func (l loggingABCIApp) FinalizeBlock(ctx context.Context, block *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	log.Debug(ctx, "👾 ABCI call: FinalizeBlock")
	resp, err := l.Application.FinalizeBlock(ctx, block)

	for i, res := range resp.TxResults {
		if res.Code == 0 {
			continue
		}
		log.Error(ctx, "Unexpected failed transaction", nil,
			"info", res.Info, "code", res.Code, "log", res.Log,
			"code_space", res.Codespace, "index", i, "height", block.Height)
	}

	return resp, err
}

func (l loggingABCIApp) ExtendVote(ctx context.Context, vote *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	log.Debug(ctx, "👾 ABCI call: ExtendVote")
	return l.Application.ExtendVote(ctx, vote)
}

func (l loggingABCIApp) VerifyVoteExtension(ctx context.Context, extension *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
	log.Debug(ctx, "👾 ABCI call: VerifyVoteExtension")
	return l.Application.VerifyVoteExtension(ctx, extension)
}

func (l loggingABCIApp) Commit(ctx context.Context, commit *abci.RequestCommit) (*abci.ResponseCommit, error) {
	log.Debug(ctx, "👾 ABCI call: Commit")
	return l.Application.Commit(ctx, commit)
}

func (l loggingABCIApp) ListSnapshots(ctx context.Context, listSnapshots *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	log.Debug(ctx, "👾 ABCI call: ListSnapshots")
	return l.Application.ListSnapshots(ctx, listSnapshots)
}

func (l loggingABCIApp) OfferSnapshot(ctx context.Context, snapshot *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	log.Debug(ctx, "👾 ABCI call: OfferSnapshot")
	return l.Application.OfferSnapshot(ctx, snapshot)
}

func (l loggingABCIApp) LoadSnapshotChunk(ctx context.Context, chunk *abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	log.Debug(ctx, "👾 ABCI call: LoadSnapshotChunk")
	return l.Application.LoadSnapshotChunk(ctx, chunk)
}

func (l loggingABCIApp) ApplySnapshotChunk(ctx context.Context, chunk *abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	log.Debug(ctx, "👾 ABCI call: ApplySnapshotChunk")
	return l.Application.ApplySnapshotChunk(ctx, chunk)
}
