package app

import (
	"context"
	"crypto/ecdsa"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/solver/job"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func streamSVMForever(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	cl *rpc.Client,
	cursors *cursors,
	db *job.DB,
	asyncWork asyncWorkFunc,
) {
	backoff := expbackoff.New(ctx)
	for {
		from, ok, err := cursors.GetTxSig(ctx, chainVer)
		if !ok || err != nil {
			log.Warn(ctx, "Failed reading cursor (will retry)", err)
			backoff()

			continue
		}

		callback := newSVMStreamCallback(cl, chainVer, cursors, db, asyncWork)

		req := svmutil.StreamReq{
			AfterSig:   from,
			Account:    anchorinbox.ProgramID,
			Commitment: rpc.CommitmentConfirmed,
		}

		log.Debug(ctx, "SVM streaming events", "after_sig", from)

		err = svmutil.Stream(ctx, cl, req, callback)
		if ctx.Err() != nil {
			return
		}

		// If the stream fails, we backoff and retry.
		log.Warn(ctx, "SVM stream failed (will retry)", err)
		backoff()
	}
}

// svmProcDeps returns the SVM-specific processor dependencies based on the provided global procdeps.
// Specifically, it replaces the functions that interact with the SVM chain.
func svmProcDeps(
	cl *rpc.Client,
	outboxAddr common.Address,
	solverEVM *ecdsa.PrivateKey,
	deps procDeps,
) procDeps {
	solver := svmutil.MapEVMKey(solverEVM)
	evmFill := deps.Fill

	deps.GetOrder = adaptSVMGetOrder(cl, outboxAddr)
	deps.Reject = adaptSVMReject(cl, solver)
	deps.Claim = adaptSVMClaim(cl, solver)
	deps.Fill = func(ctx context.Context, order Order) error {
		if err := evmFill(ctx, order); err != nil {
			return err
		}

		// TODO(corver): Move this to dedicated outbox event processor.
		return markFilledSVMOrder(ctx, cl, solver, solver.PublicKey(), order.ID)
	}

	return deps
}

// newSVMStreamCallback returns stream handler for anchor inbox events.
// It starts async jobs for each valid event.
func newSVMStreamCallback(
	cl *rpc.Client,
	chainVer xchain.ChainVersion,
	cursors *cursors,
	jobDB *job.DB,
	asyncWork asyncWorkFunc,
) svmutil.StreamCallback {
	return func(ctx context.Context, sig *rpc.TransactionSignature) error {
		if sig.Err != nil {
			log.Debug(ctx, "SVM: Ignoring failed tx", "tx_err", sig.Err, "tx", sig)
			return nil // Skip failed transactions
		}

		txResp, err := svmutil.AwaitConfirmedTransaction(ctx, cl, sig.Signature)
		if err != nil {
			return errors.Wrap(err, "get tx")
		}

		logMsgs, ok, err := svmutil.FilterDataLogs(txResp.Meta.LogMessages, anchorinbox.ProgramID)
		if err != nil {
			return errors.Wrap(err, "filter logs")
		} else if !ok {
			log.Warn(ctx, "SVM: potentially skipping truncated logs [BUG]", nil, "tx", sig)
		}

		events, err := anchorinbox.DecodeLogEvents(logMsgs)
		if err != nil {
			return errors.Wrap(err, "decode events")
		}

		for i, event := range events {
			if event.Name != anchorinbox.EventNameUpdated {
				return errors.New("unexpected event [BUG]", "event", event.Name)
			}

			data, ok := event.Data.(*anchorinbox.EventUpdatedEventData)
			if !ok {
				return errors.New("unexpected event data [BUG]", "event", event.Name)
			}

			indexU64, err := umath.ToUint64(i)
			if err != nil {
				return err
			}

			statusU64, err := umath.ToUint64(data.Status)
			if err != nil {
				return err
			}

			j, err := jobDB.Insert(
				ctx,
				chainVer.ID,
				txResp.Slot,
				sig.Signature.String(),
				indexU64,
				data.OrderId[:],
				statusU64,
			)
			if err != nil {
				return err
			}

			if err := asyncWork(ctx, j); err != nil {
				return errors.Wrap(err, "async work")
			}
		}

		log.Debug(ctx, "SVM: Processed events", "sig", sig.Signature, "slot", sig.Slot, "events", len(events))

		if err := cursors.SetTxSig(ctx, chainVer, sig.Signature); err != nil {
			return errors.Wrap(err, "update cursor")
		}

		return nil
	}
}

// adaptSVMGetOrder adapts the svmGetOrder function to the procDeps interface.
func adaptSVMGetOrder(cl *rpc.Client, outboxAddr common.Address) func(context.Context, uint64, OrderID) (Order, bool, error) {
	return func(ctx context.Context, _ uint64, id OrderID) (Order, bool, error) {
		return svmGetOrder(ctx, cl, outboxAddr, id)
	}
}

// adaptSVMReject adapts the rejectSVMOrder function to the procDeps interface.
func adaptSVMReject(cl *rpc.Client, solver solana.PrivateKey) func(ctx context.Context, order Order, reason stypes.RejectReason) error {
	return func(ctx context.Context, order Order, reason stypes.RejectReason) error {
		return rejectSVMOrder(ctx, cl, solver, order.ID, reason)
	}
}

// adaptSVMClaim adapts the claimSVMOrder function to the procDeps interface.
func adaptSVMClaim(cl *rpc.Client, solver solana.PrivateKey) func(ctx context.Context, order Order) error {
	return func(ctx context.Context, order Order) error {
		return claimSVMOrder(ctx, cl, solver, order.ID)
	}
}
