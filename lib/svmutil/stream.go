package svmutil

import (
	"context"
	"slices"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// StreamReq defines which transactions to stream.
type StreamReq struct {
	AfterSig   solana.Signature // Exclusive
	Account    solana.PublicKey
	Commitment rpc.CommitmentType
}

// StreamCallback abstracts the logic that handles a stream of transaction signatures.
// The logic must be idempotent.
type StreamCallback func(ctx context.Context, sig *rpc.TransactionSignature) error

func Stream(ctx context.Context, cl *rpc.Client, req StreamReq, callback StreamCallback) error {
	if req.AfterSig.IsZero() {
		return errors.New("zero AfterSig")
	}

	// Max backoff of 3s
	backoffCfg := expbackoff.DefaultConfig
	backoffCfg.MaxDelay = time.Second * 3
	backoff, reset := expbackoff.NewWithReset(ctx, expbackoff.With(backoffCfg))

	var prev *rpc.TransactionSignature

	for ctx.Err() == nil {
		sigs, err := allSigsForAccount(ctx, cl, req)
		if ctx.Err() != nil {
			return nil //nolint:nilerr // As per API, return nil on context cancel.
		} else if err != nil {
			return WrapRPCError(err, "getSignaturesForAddress")
		} else if len(sigs) == 0 {
			// No sigs found between from* and latest
			backoff()

			continue
		}

		for _, sig := range sigs {
			// Sanity checks
			if req.AfterSig == sig.Signature {
				return errors.New("signature is equal to afterSig [BUG]")
			} else if prev != nil && sig.BlockTime != nil && *prev.BlockTime > *sig.BlockTime {
				return errors.New("block time is less than prev [BUG]", "sig", sig.Signature, "prev", prev.Signature, "time", sig.BlockTime, "prev_time", prev.BlockTime)
			} else if prev != nil && sig.Slot < prev.Slot {
				return errors.New("slot is less than prev [BUG]", "sig", sig.Signature, "prev", prev.Signature, "slot", sig.Slot, "prev_slot", prev.Slot)
			}
			prev = sig

			ctx := log.WithCtx(ctx, "slot", sig.Slot, "sig", sig.Signature.String()[:7])

			err := callback(ctx, sig)
			if ctx.Err() != nil {
				return nil //nolint:nilerr // As per API, return nil on context cancel.
			} else if err != nil {
				return errors.Wrap(err, "callback", "slot", sig.Slot, "signature", sig.Signature)
			}
		}

		// Update cursor (prev always non-nil at this point)
		req.AfterSig = prev.Signature
		reset()
	}

	return nil
}

// allSigsForAccount returns all signatures for the given account, starting from provided req.
func allSigsForAccount(ctx context.Context, cl *rpc.Client, req StreamReq) ([]*rpc.TransactionSignature, error) {
	var before solana.Signature
	var resp []*rpc.TransactionSignature
	for i := 0; ; i++ {
		const limit int = 1000
		sigs, err := cl.GetSignaturesForAddressWithOpts(ctx, req.Account, &rpc.GetSignaturesForAddressOpts{
			Before:     before,
			Until:      req.AfterSig,
			Commitment: req.Commitment,
			Limit:      ptr(limit),
		})
		if err != nil {
			return nil, WrapRPCError(err, "getSignaturesForAddress", "page", i)
		}

		// Sigs are in descending order (newer to older, so we need to reverse them and prepend to resp)
		slices.Reverse(sigs)
		resp = append(sigs, resp...)

		if len(sigs) < limit {
			break
		}

		before = sigs[0].Signature
	}

	return resp, nil
}

func ptr[T any](v T) *T {
	return &v
}
