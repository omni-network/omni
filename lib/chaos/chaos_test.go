package chaos_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/chaos"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

func TestMaybeError(t *testing.T) {
	t.Parallel()
	ctx := chaos.WithErrProbability(context.Background(), netconf.Devnet)

	for i := 0; ; i++ {
		err := chaos.MaybeError(ctx)
		if err == nil {
			continue
		}

		log.Warn(ctx, "Got chaos errors", err, "i", i) // Manually confirm stack trace is correct.

		return
	}
}
