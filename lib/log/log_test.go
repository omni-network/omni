package log_test

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/omni-network/omni/lib/log"
)

// TestSimpleLogs tests the simple logs.
func TestSimpleLogs(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	log.Info(ctx, "info message", "with", "args")
	log.Debug(ctx, "debug this code for me please", "number", 1)
	log.Warn(ctx, "watch out!", "err", os.ErrExist)
	log.Error(ctx, "something went wrong", io.EOF, "float", 1.234)
}
