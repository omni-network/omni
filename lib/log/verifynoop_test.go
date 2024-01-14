//go:build !verify_logs

package log_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/log"

	"github.com/stretchr/testify/require"
)

func TestVerifyNoop(t *testing.T) {
	t.Parallel()
	require.NotPanics(t, func() {
		log.Info(context.Background(), "This should not panic", "BADKEY", "value")
	})
}
