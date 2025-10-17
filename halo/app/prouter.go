package app

import (
	"context"
	"time"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/gogoproto/proto"
)

// processTimeout is the maximum time to process a proposal.
// Timeout results in rejecting the proposal, which could negatively affect liveness.
// But it avoids blocking forever, which also negatively affects liveness.
// This mitigates against malicious proposals that take forever to process (e.g. due to retryForever).
const processTimeout = time.Minute

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
		timeoutCtx, timeoutCancel := context.WithTimeout(ctx.Context(), processTimeout)
		defer timeoutCancel()
		ctx = ctx.WithContext(timeoutCtx)

		if req.Height == 1 {
			if len(req.Txs) > 0 { // First proposal must be empty.
				return rejectProposal(ctx, errors.New("first proposal not empty"))
			}

			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
		} else if len(req.Txs) > 1 {
			return rejectProposal(ctx, errors.New("unexpected transactions in proposal"))
		}

		// Ensure the proposal includes quorum votes.
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

			if err := verifyTX(tx); err != nil {
				return rejectProposal(ctx, errors.Wrap(err, "verify transaction"))
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

// verifyTX ensures a transaction is empty of all signing fields.
//
//nolint:gocyclo // Just lots of checks.
func verifyTX(tx sdk.Tx) error {
	if tx == nil {
		return errors.New("nil tx")
	}

	stx, ok := tx.(signing.Tx)
	if !ok {
		return errors.New("not a signing tx")
	}

	// Convert to actual proto and verify fields explicitly.
	hasProtoTx, ok := tx.(protoTxProvider)
	if !ok {
		return errors.New("tx does not implement protoTxProvider")
	}

	// Verify basic structure
	ptx := hasProtoTx.GetProtoTx()
	if ptx == nil {
		return errors.New("proto tx is nil")
	} else if ptx.AuthInfo == nil {
		return errors.New("proto tx auth info is nil")
	} else if ptx.AuthInfo.Fee == nil {
		return errors.New("proto tx auth info fee is nil")
	} else if ptx.Body == nil {
		return errors.New("proto tx body is nil")
	}

	const addrLen = 20

	// Ensure all signing fields are empty
	if stx.GetGas() != 0 {
		return errors.New("gas not empty")
	}
	if !stx.GetFee().IsZero() {
		return errors.New("fee not empty")
	}
	if len(stx.FeeGranter()) != 0 {
		return errors.New("fee granter not empty")
	}
	if len(stx.GetMemo()) != 0 {
		return errors.New("memo not empty")
	}
	if stx.GetTimeoutHeight() != 0 {
		return errors.New("timeout height not empty")
	}

	msgsLen := len(stx.GetMsgs())

	signers, err := stx.GetSigners()
	if err != nil {
		return errors.Wrap(err, "get signers")
	} else if len(signers) > msgsLen { // Unique signers can't be more than amount of msgs
		return errors.New("unexpected amount of signers", "count", len(signers), "max", msgsLen)
	}
	for _, signer := range signers {
		if len(signer) != addrLen {
			return errors.New("signer invalid", "len", len(signer))
		}
	}
	if len(signers) > 0 {
		// FeePayer panics if no signers
		if len(stx.FeePayer()) != addrLen {
			return errors.New("fee payer invalid")
		}
	}

	pks, err := stx.GetPubKeys()
	if err != nil {
		return errors.Wrap(err, "get pubkeys")
	} else if len(pks) != 0 {
		return errors.New("pks not empty", "count", len(pks))
	}

	sigs, err := stx.GetSignaturesV2()
	if err != nil {
		return errors.Wrap(err, "get sigs")
	} else if len(sigs) != 0 {
		return errors.New("sigs not empty", "count", len(sigs))
	}

	msgs2, err := stx.GetMsgsV2()
	if err != nil {
		return errors.Wrap(err, "get msgs v2")
	} else if len(msgs2) != msgsLen {
		return errors.New("msgs v2 count mismatch")
	}

	if len(ptx.GetSignatures()) != 0 { //nolint:nestif // Multiple checks are needed
		return errors.New("proto tx signatures not empty")
	} else if ptx.AuthInfo.Tip != nil {
		return errors.New("proto tx tip not nil")
	} else if len(ptx.AuthInfo.SignerInfos) > 0 {
		return errors.New("proto tx signer infos not empty")
	} else if !proto.Equal(ptx.AuthInfo.Fee, new(txtypes.Fee)) {
		return errors.New("proto tx fee not zero")
	} else if ptx.Body.TimeoutHeight != 0 {
		return errors.New("proto tx timeout height not empty")
	} else if ptx.Body.Memo != "" {
		return errors.New("proto tx memo not empty")
	} else if len(ptx.Body.ExtensionOptions) > 0 {
		return errors.New("proto tx extension options not empty")
	} else if len(ptx.Body.NonCriticalExtensionOptions) > 0 {
		return errors.New("proto tx non-critical extension options not empty")
	} else if len(ptx.Body.Messages) != msgsLen {
		return errors.New("proto tx messages count mismatch")
	}

	return nil
}

type protoTxProvider interface {
	GetProtoTx() *txtypes.Tx
}

//nolint:unparam // Explicitly return nil error
func rejectProposal(ctx context.Context, err error) (*abci.ResponseProcessProposal, error) {
	log.Error(ctx, "Rejecting process proposal", err)
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
}
