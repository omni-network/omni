package app

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"

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
		for _, rawTX := range req.Txs {
			tx, err := app.txConfig.TxDecoder()(rawTX)
			if err != nil {
				return handleErr(ctx, errors.Wrap(err, "decode transaction"))
			}

			for _, msg := range tx.GetMsgs() {
				handler := router.Handler(msg)
				if handler == nil {
					return handleErr(ctx, errors.Wrap(err, "msg handler not found"))
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
	log.Warn(ctx, "Rejecting failed process proposal", err)
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
}
