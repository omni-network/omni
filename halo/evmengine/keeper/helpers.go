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
