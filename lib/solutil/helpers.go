package solutil

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"

	bin "github.com/gagliardetto/binary"
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

// GetAccountDataInto retrieves account data and decodes it into the provided value.
// It uses commitment level of "confirmed".
func GetAccountDataInto[T any](ctx context.Context, cl *rpc.Client, address solana.PublicKey, val T) (*rpc.GetAccountInfoResult, error) {
	info, err := cl.GetAccountInfoWithOpts(ctx, address, &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
		Encoding:   solana.EncodingBase64,
	})
	if err != nil {
		return nil, errors.Wrap(WrapRPCError(err, "GetAccountDataInto"), "get account info")
	}

	err = bin.NewBorshDecoder(info.Value.Data.GetBinary()).Decode(val)
	if err != nil {
		return nil, errors.Wrap(err, "decode account data")
	}

	return info, nil
}

var chainIDsByHash = map[solana.Hash]uint64{
	solana.MustHashFromBase58("5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d"): evmchain.IDSolana,
	solana.MustHashFromBase58("4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY"): evmchain.IDSolanaTest,
}

func ChainID(ctx context.Context, cl *rpc.Client) (uint64, error) {
	hash, err := cl.GetGenesisHash(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get chain ID")
	}

	if chainID, ok := chainIDsByHash[hash]; ok {
		return chainID, nil
	}

	return evmchain.IDSolanaLocal, nil
}
