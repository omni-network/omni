package app

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// makeProcessProposalHandler creates a new process proposal handler.
// It ensures all messages included in a cpayload proposal are valid.
// It also updates some external state.
func makeProcessProposalHandler(app *App) sdk.ProcessProposalHandler {
	router := baseapp.NewMsgServiceRouter()
	router.SetInterfaceRegistry(app.interfaceRegistry)
	app.EVMEngKeeper.RegisterProposalService(router) // EVMEngine calls NewPayload on proposals to verify it.
	app.AttestKeeper.RegisterProposalService(router) // Attester marks attestations as proposed.

	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		// Ensure the proposal includes quorum vote extensions (unless first block).
		if req.Height > 1 {
			var totalPower, votedPower int64
			for _, vote := range req.ProposedLastCommit.Votes {
				totalPower += vote.Validator.Power
				if vote.BlockIdFlag != cmttypes.BlockIDFlagCommit {
					continue
				}
				votedPower += vote.Validator.Power
			}
			if totalPower*2/3 >= votedPower {
				return handleErr(ctx, errors.New("proposed doesn't include quorum votes exttensions"))
			}
		}

		for _, rawTX := range req.Txs {
			tx, err := app.txConfig.TxDecoder()(rawTX)
			if err != nil {
				return handleErr(ctx, errors.Wrap(err, "decode transaction"))
			}

			for _, msg := range tx.GetMsgs() {
				handler := router.Handler(msg)
				if handler == nil {
					return handleErr(ctx, errors.New("msg handler not found",
						"msg_type", fmt.Sprintf("%T", msg),
					))
				}

				_, err := handler(ctx, msg)
				if err != nil {
					return handleErr(ctx, errors.Wrap(err, "execute message"))
				}
			}
		}

		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}

func handleErr(ctx context.Context, err error) (*abci.ResponseProcessProposal, error) {
	log.Error(ctx, "Rejecting failed process proposal", err)
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
}
