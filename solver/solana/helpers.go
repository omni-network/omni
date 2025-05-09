package solana

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var v0 uint64

// AwaitConfirmedTransaction waits for a transaction to be confirmed.
func AwaitConfirmedTransaction(ctx context.Context, cl *rpc.Client, txSig solana.Signature) (*rpc.GetTransactionResult, error) {
	for {
		tx, err := cl.GetTransaction(ctx, txSig, &rpc.GetTransactionOpts{
			Encoding:                       solana.EncodingBase64,
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &v0,
		})
		if errors.Is(err, rpc.ErrNotFound) {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			return nil, errors.Wrap(err, "get confirmed transaction")
		} else if tx.Meta.Err != nil {
			return tx, errors.New("transaction failed", "meta_err", tx.Meta.Err, "signature", txSig)
		}

		return tx, nil
	}
}
