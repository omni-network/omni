package bindings

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// DeployWithRetry calls Deploy up to <retry> times on errors.
func (_Create3 *Create3Transactor) DeployWithRetry(opts *bind.TransactOpts, salt [32]byte, creationCode []byte) (*types.Transaction, error) {
	ctx := opts.Context
	if ctx == nil {
		ctx = context.Background()
	}

	const retry = 3
	for i := 1; ; i++ {
		tx, err := _Create3.Deploy(opts, salt, creationCode)
		if err == nil {
			return tx, nil
		} else if i >= retry {
			return nil, errors.Wrap(err, "factory deploy", "attempt", i)
		}

		log.Warn(ctx, "Failed factory deploy (will retry)", err, "attempt", i)
		time.Sleep(time.Second)
	}
}
