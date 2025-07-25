package keeper

import (
	"context"
	"crypto/sha256"
	"math/big"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
)

var (
	retryTimeout  = time.Minute // Just prevent blocking forever
	backoffFuncMu sync.RWMutex
	backoffFunc   = expbackoff.New // backoffFunc aliased for testing.

	// maxBlobsPerBlock is the maximum number of blobs per block.
	// Copied from https://github.com/ethereum/consensus-specs/blob/dev/specs/deneb/beacon-chain.md#execution-1.
	maxBlobsPerBlock = 6
)

// retryForever retries the given function forever until it returns true or an error.
// In order for the function to be retried, it must return false and no error.
//
// Networking (any IO) is non-deterministic and can fail with temporary errors.
// Keeper logic must however be deterministic, retrying forever mitigates this.
func retryForever(ctx context.Context, fn func(ctx context.Context) (bool, error)) error {
	backoffFuncMu.RLock()
	backoff := backoffFunc(ctx)
	backoffFuncMu.RUnlock()
	for {
		innerCtx, cancel := context.WithTimeout(ctx, retryTimeout)
		ok, err := fn(innerCtx)
		cancel()
		if ctx.Err() != nil {
			return errors.Wrap(ctx.Err(), "retry canceled")
		} else if err != nil {
			return err
		} else if !ok {
			backoff()
			continue
		}

		return nil
	}
}

func unwrapHexBytes(in []hexutil.Bytes) [][]byte {
	var out [][]byte
	for _, i := range in {
		out = append(out, i)
	}

	return out
}

// blobHashes returns the blob hashes from provided commitments.
func blobHashes(commitments [][]byte) ([]common.Hash, error) {
	if len(commitments) > maxBlobsPerBlock {
		return nil, errors.New("too many blobs", "max", maxBlobsPerBlock, "actual", len(commitments))
	}

	hasher := sha256.New()
	resp := make([]common.Hash, 0, len(commitments)) // Default to zero len slice, not nil.
	for _, commitment := range commitments {
		commitment48, err := cast.Array48(commitment)
		if err != nil {
			return nil, err
		}

		resp = append(resp, kzg4844.CalcBlobHashV1(hasher, (*kzg4844.Commitment)(&commitment48)))
	}

	return resp, nil
}

// toGwei converts a wei amount to a gwei amount and the wei remainder.
func toGwei(weiAmount *big.Int) (gweiU64 uint64, weiRemU64 uint64, err error) { //nolint:nonamedreturns // Disambiguate identical return types.
	const giga uint64 = 1e9
	gweiAmount := bi.DivRaw(weiAmount, giga)
	weiRem := bi.Sub(weiAmount, bi.MulRaw(gweiAmount, giga))

	// This should work up to 18G ETH
	if !gweiAmount.IsUint64() {
		return 0, 0, errors.New("invalid amount [BUG]")
	}

	return gweiAmount.Uint64(), weiRem.Uint64(), nil
}
