//nolint:wrapcheck // The abci.Application is our application, so we don't need to wrap it. Long lines are fine here.
package app

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"
)

type loggingABCIApp struct {
	abci.Application
}

func (l loggingABCIApp) Info(ctx context.Context, info *abci.RequestInfo) (*abci.ResponseInfo, error) {
	log.Debug(ctx, "👾 ABCI call: Info")
	resp, err := l.Application.Info(ctx, info)
	if err != nil {
		log.Error(ctx, "Info failed [BUG]", err)
	}

	return resp, err
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
	resp, err := l.Application.InitChain(ctx, chain)
	if err != nil {
		log.Error(ctx, "InitChain failed [BUG]", err)
	}

	return resp, err
}

func (l loggingABCIApp) PrepareProposal(ctx context.Context, proposal *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	log.Debug(ctx, "👾 ABCI call: PrepareProposal",
		"height", proposal.Height,
		log.Hex7("proposer", proposal.ProposerAddress),
	)
	resp, err := l.Application.PrepareProposal(ctx, proposal)
	if err != nil {
		log.Error(ctx, "PrepareProposal failed [BUG]", err)
	}

	return resp, err
}

func (l loggingABCIApp) ProcessProposal(ctx context.Context, proposal *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	log.Debug(ctx, "👾 ABCI call: ProcessProposal",
		"height", proposal.Height,
		log.Hex7("proposer", proposal.ProposerAddress),
	)
	resp, err := l.Application.ProcessProposal(ctx, proposal)
	if err != nil {
		log.Error(ctx, "ProcessProposal failed [BUG]", err)
	}

	return resp, err
}

func (l loggingABCIApp) FinalizeBlock(ctx context.Context, block *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	resp, err := l.Application.FinalizeBlock(ctx, block)
	if err != nil {
		log.Error(ctx, "Finalize block failed [BUG]", err, "height", block.Height)
		return resp, err
	}

	attrs := []any{
		"val_updates", len(resp.ValidatorUpdates),
		"height", block.Height,
	}
	for i, update := range resp.ValidatorUpdates {
		attrs = append(attrs, log.Hex7(fmt.Sprintf("pubkey_%d", i), update.PubKey.GetSecp256K1()))
		attrs = append(attrs, fmt.Sprintf("power_%d", i), update.Power)
	}
	log.Debug(ctx, "👾 ABCI response: FinalizeBlock", attrs...)

	for i, res := range resp.TxResults {
		if res.Code == 0 {
			continue
		}
		log.Error(ctx, "FinalizeBlock contains unexpected failed transaction [BUG]", nil,
			"info", res.Info, "code", res.Code, "log", res.Log,
			"code_space", res.Codespace, "index", i, "height", block.Height)
	}

	return resp, err
}

func (l loggingABCIApp) ExtendVote(ctx context.Context, vote *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	log.Debug(ctx, "👾 ABCI call: ExtendVote",
		"height", vote.Height,
	)
	resp, err := l.Application.ExtendVote(ctx, vote)
	if err != nil {
		log.Error(ctx, "ExtendVote failed [BUG]", err)
	}

	return resp, err
}

func (l loggingABCIApp) VerifyVoteExtension(ctx context.Context, extension *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
	log.Debug(ctx, "👾 ABCI call: VerifyVoteExtension",
		"height", extension.Height,
	)
	resp, err := l.Application.VerifyVoteExtension(ctx, extension)
	if err != nil {
		log.Error(ctx, "VerifyVoteExtension failed [BUG]", err)
	}

	return resp, err
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
