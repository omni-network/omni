package flowgen

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"
)

func newSymbioticRunner() RunFunc {
	return func(_ context.Context, spend Spend) error {
		_, ok := spend[tokens.WSTETH]
		if !ok {
			return errors.New("missing deposit", "token", tokens.WSTETH)
		}

		// TODO

		return nil
	}
}
