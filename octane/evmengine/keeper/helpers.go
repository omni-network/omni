package keeper

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
)

// backoffFunc aliased for testing.
var (
	retryTimeout  = time.Minute // Just prevent blocking forever
	backoffFuncMu sync.RWMutex
	backoffFunc   = expbackoff.New
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
