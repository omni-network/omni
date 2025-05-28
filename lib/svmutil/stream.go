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
	FromSlot   *uint64          // Inclusive if not nil
	AfterSig   solana.Signature // Exclusive
	Account    solana.PublicKey
	Commitment rpc.CommitmentType
}

// StreamCallback abstracts the logic that handles a stream of transaction signatures.
// The logic must be idempotent.
type StreamCallback func(ctx context.Context, sig *rpc.TransactionSignature) error

func Stream(ctx context.Context, cl *rpc.Client, req StreamReq, callback StreamCallback) error {
	if req.FromSlot == nil && req.AfterSig.IsZero() {
		return errors.New("neither FromSlot nor AfterSig provided")
	}

	// Max backoff of 3s
	backoffCfg := expbackoff.DefaultConfig
	backoffCfg.MaxDelay = time.Second * 3
	backoff, reset := expbackoff.NewWithReset(ctx, expbackoff.With(backoffCfg))

	var prev *rpc.TransactionSignature

	for ctx.Err() == nil {
		latest, err := cl.GetSlot(ctx, req.Commitment)
		if ctx.Err() != nil {
			return nil //nolint:nilerr // As per API, return nil on context cancel.
		} else if err != nil {
			return WrapRPCError(err, "getSlot")
		}

		sigs, err := allSigsForAccount(ctx, cl, req)
		if ctx.Err() != nil {
			return nil //nolint:nilerr // As per API, return nil on context cancel.
		} else if err != nil {
			return WrapRPCError(err, "getSignaturesForAddress")
		} else if len(sigs) == 0 {
			// No sigs found between from* and latest
			req.FromSlot = &latest
			backoff()

			continue
		}

		for _, sig := range sigs {
			// Sanity checks
			if req.FromSlot != nil && sig.Slot < *req.FromSlot {
				return errors.New("signature slot is less than fromSlot [BUG]")
			} else if req.AfterSig == sig.Signature {
				return errors.New("signature is equal to afterSig [BUG]")
			} else if prev != nil && sig.BlockTime != nil && *prev.BlockTime > *sig.BlockTime {
				return errors.New("block time is less than prev [BUG]")
			} else if prev != nil && sig.Slot < prev.Slot {
				return errors.New("slot is less than prev [BUG]")
			}
			prev = sig

			ctx := log.WithCtx(ctx, "slot", sig.Slot, log.Hex7("sig", sig.Signature[:]))

			err := callback(ctx, sig)
			if ctx.Err() != nil {
				return nil //nolint:nilerr // As per API, return nil on context cancel.
			} else if err != nil {
				return errors.Wrap(err, "callback", "slot", sig.Slot, "signature", sig.Signature)
			}
		}

		// Update cursor (prev always non-nil at this point)
		req.AfterSig = prev.Signature
		req.FromSlot = &prev.Slot
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
			Before:         before,
			MinContextSlot: req.FromSlot,
			Until:          req.AfterSig,
			Commitment:     req.Commitment,
			Limit:          ptr(limit),
		})
		if err != nil {
			return nil, WrapRPCError(err, "getSignaturesForAddress", "page", i)
		}

		// Sigs are descending order (newer to older, so we need to reverse them and prepend to resp)
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
