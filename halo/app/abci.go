//nolint:wrapcheck // The abci.Application is our application, so we don't need to wrap it. Long lines are fine here.
package app

import (
	"context"
	"fmt"
	"time"

	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type postFinalizeCallback func(sdk.Context) error
type multiStoreProvider func() storetypes.CacheMultiStore

type abciWrapper struct {
	abci.Application
	postFinalize       postFinalizeCallback
	multiStoreProvider multiStoreProvider
	maybeChaos         chaosFunc
}

func newABCIWrapper(
	app abci.Application,
	finaliseCallback postFinalizeCallback,
	multiStoreProvider multiStoreProvider,
	chaosTest bool,
) *abciWrapper {
	chaos := noopChaos
	if chaosTest {
		chaos = newChaosSlasher()
	}

	return &abciWrapper{
		Application:        app,
		postFinalize:       finaliseCallback,
		multiStoreProvider: multiStoreProvider,
		maybeChaos:         chaos,
	}
}

func (l abciWrapper) Info(ctx context.Context, info *abci.RequestInfo) (*abci.ResponseInfo, error) {
	log.Debug(ctx, "👾 ABCI call: Info")
	resp, err := l.Application.Info(ctx, info)
	if err != nil {
		log.Error(ctx, "Info failed [BUG]", err)
	}

	return resp, err
}

func (l abciWrapper) Query(ctx context.Context, query *abci.RequestQuery) (*abci.ResponseQuery, error) {
	return l.Application.Query(ctx, query) // No log here since this can be very noisy
}

func (l abciWrapper) CheckTx(ctx context.Context, tx *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	log.Debug(ctx, "👾 ABCI call: CheckTx")
	return l.Application.CheckTx(ctx, tx)
}

func (l abciWrapper) InitChain(ctx context.Context, chain *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	log.Debug(ctx, "👾 ABCI call: InitChain")
	resp, err := l.Application.InitChain(ctx, chain)
	if err != nil {
		log.Error(ctx, "InitChain failed [BUG]", err)
	}

	return resp, err
}

func (l abciWrapper) PrepareProposal(ctx context.Context, proposal *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	ctx = log.WithCtx(ctx, "height", proposal.Height)
	log.Debug(ctx, "👾 ABCI call: PrepareProposal",
		log.Hex7("proposer", proposal.ProposerAddress),
	)
	resp, err := l.Application.PrepareProposal(ctx, proposal)
	if err != nil {
		log.Error(ctx, "PrepareProposal failed [BUG]", err)
	}

	return resp, err
}

func (l abciWrapper) ProcessProposal(ctx context.Context, proposal *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	ctx = log.WithCtx(ctx, "height", proposal.Height)
	log.Debug(ctx, "👾 ABCI call: ProcessProposal",
		log.Hex7("proposer", proposal.ProposerAddress),
	)
	resp, err := l.Application.ProcessProposal(ctx, proposal)
	if err != nil {
		log.Error(ctx, "ProcessProposal failed [BUG]", err)
	}

	return resp, err
}

func (l abciWrapper) FinalizeBlock(ctx context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	l.maybeChaos(ctx, req)

	ctx = log.WithCtx(ctx, "height", req.Height)
	resp, err := l.Application.FinalizeBlock(ctx, req)
	if err != nil {
		log.Error(ctx, "Finalize req failed [BUG]", err)
		return resp, err
	}

	// Call custom `PostFinalize` callback after the block is finalized.
	header := cmtproto.Header{
		Height:             req.Height,
		Time:               req.Time,
		ProposerAddress:    req.ProposerAddress,
		NextValidatorsHash: req.NextValidatorsHash,
		AppHash:            resp.AppHash, // The app hash after the block is finalized.
	}
	sdkCtx := sdk.NewContext(l.multiStoreProvider(), header, false, nil)
	if err := l.postFinalize(sdkCtx); err != nil {
		log.Error(ctx, "PostFinalize callback failed [BUG]", err)
		return resp, err
	}

	attrs := []any{
		"val_updates", len(resp.ValidatorUpdates),
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
			"code_space", res.Codespace, "index", i, "height", req.Height)
	}

	return resp, err
}

func (l abciWrapper) ExtendVote(ctx context.Context, vote *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	ctx = log.WithCtx(ctx, "height", vote.Height)
	log.Debug(ctx, "👾 ABCI call: ExtendVote")
	resp, err := l.Application.ExtendVote(ctx, vote)
	if err != nil {
		log.Error(ctx, "ExtendVote failed [BUG]", err)
	}

	return resp, err
}

func (l abciWrapper) VerifyVoteExtension(ctx context.Context, extension *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
	ctx = log.WithCtx(ctx, "height", extension.Height)
	log.Debug(ctx, "👾 ABCI call: VerifyVoteExtension")
	resp, err := l.Application.VerifyVoteExtension(ctx, extension)
	if err != nil {
		log.Error(ctx, "VerifyVoteExtension failed [BUG]", err)
	}

	return resp, err
}

func (l abciWrapper) Commit(ctx context.Context, commit *abci.RequestCommit) (*abci.ResponseCommit, error) {
	log.Debug(ctx, "👾 ABCI call: Commit")
	return l.Application.Commit(ctx, commit)
}

func (l abciWrapper) ListSnapshots(ctx context.Context, listSnapshots *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	log.Debug(ctx, "👾 ABCI call: ListSnapshots")
	return l.Application.ListSnapshots(ctx, listSnapshots)
}

func (l abciWrapper) OfferSnapshot(ctx context.Context, snapshot *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	log.Debug(ctx, "👾 ABCI call: OfferSnapshot")
	return l.Application.OfferSnapshot(ctx, snapshot)
}

func (l abciWrapper) LoadSnapshotChunk(ctx context.Context, chunk *abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	log.Debug(ctx, "👾 ABCI call: LoadSnapshotChunk")
	return l.Application.LoadSnapshotChunk(ctx, chunk)
}

func (l abciWrapper) ApplySnapshotChunk(ctx context.Context, chunk *abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	log.Debug(ctx, "👾 ABCI call: ApplySnapshotChunk")
	return l.Application.ApplySnapshotChunk(ctx, chunk)
}

// chaos abstracts a source of chaos testing.
type chaosFunc func(ctx context.Context, req *abci.RequestFinalizeBlock)

// noopChaos is a no-op chaos implementation.
// This is used in production to disable chaos testing.
var noopChaos = func(context.Context, *abci.RequestFinalizeBlock) {}

const chaosSlashMinHeight = 10

// newChaosSlasher return a chaosFunc that slashes a validator once.
func newChaosSlasher() chaosFunc {
	var done bool
	return func(ctx context.Context, req *abci.RequestFinalizeBlock) {
		if done || req.Height < chaosSlashMinHeight {
			return
		}

		// Wait for all validators to be online
		var totalPower int64
		var lastVal abci.Validator
		for _, vote := range req.DecidedLastCommit.Votes {
			if vote.BlockIdFlag != cmtproto.BlockIDFlagCommit {
				return
			}

			totalPower += vote.Validator.Power
			lastVal = vote.Validator
		}

		req.Misbehavior = append(req.Misbehavior, abci.Misbehavior{
			Type:             abci.MisbehaviorType_DUPLICATE_VOTE,
			Validator:        lastVal,
			Height:           req.Height - 1,
			Time:             req.Time.Add(-time.Second),
			TotalVotingPower: totalPower,
		})

		done = true
		log.Warn(ctx, "😱 Chaos slashing validator", nil, log.Hex7("val", lastVal.Address), "height", req.Height)
	}
}
