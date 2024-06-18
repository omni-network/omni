package app

import (
	"context"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// makeProcessProposalRouter creates a new process proposal router that only routes
// expected messages to expected modules.
func makeProcessProposalRouter(app *App) *baseapp.MsgServiceRouter {
	router := baseapp.NewMsgServiceRouter()
	router.SetInterfaceRegistry(app.interfaceRegistry)
	app.EVMEngKeeper.RegisterProposalService(router) // EVMEngine calls NewPayload on proposals to verify it.
	app.AttestKeeper.RegisterProposalService(router) // Attester marks attestations as proposed.

	return router
}

// makeProcessProposalHandler creates a new process proposal handler.
// It ensures all messages included in a cpayload proposal are valid.
// It also updates some external state.
func makeProcessProposalHandler(router *baseapp.MsgServiceRouter, txConfig client.TxConfig) sdk.ProcessProposalHandler {
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
				return rejectProposal(ctx, errors.New("proposed doesn't include quorum votes extensions"))
			}
		}

		// Ensure only expected messages types are included the expected number of times.
		allowedMsgCounts := map[string]int{
			sdk.MsgTypeURL(&etypes.MsgExecutionPayload{}): 1, // Only a single EVM execution payload is allowed.
			sdk.MsgTypeURL(&atypes.MsgAddVotes{}):         1, // Only a single attest module MsgAddVotes is allowed.
		}

		for _, rawTX := range req.Txs {
			tx, err := txConfig.TxDecoder()(rawTX)
			if err != nil {
				return rejectProposal(ctx, errors.Wrap(err, "decode transaction"))
			}

			for _, msg := range tx.GetMsgs() {
				typeURL := sdk.MsgTypeURL(msg)

				// Ensure the message type is expected and not included too many times.
				if i, ok := allowedMsgCounts[typeURL]; !ok {
					return rejectProposal(ctx, errors.New("unexpected message type", "msg_type", typeURL))
				} else if i <= 0 {
					return rejectProposal(ctx, errors.New("message type included too many times", "msg_type", typeURL))
				}
				allowedMsgCounts[typeURL]--

				handler := router.Handler(msg)
				if handler == nil {
					return rejectProposal(ctx, errors.New("msg handler not found [BUG]", "msg_type", typeURL))
				}

				_, err := handler(ctx, msg)
				if err != nil {
					return rejectProposal(ctx, errors.Wrap(err, "execute message"))
				}
			}
		}

		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}

func rejectProposal(ctx context.Context, err error) (*abci.ResponseProcessProposal, error) {
	log.Error(ctx, "Rejecting process proposal", err)
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
}
