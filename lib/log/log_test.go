package log_test

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

//go:generate go test . -update -clean

// TestSimpleLogs tests the simple logs.
func TestSimpleLogs(t *testing.T) {
	t.Parallel()

	log.AssertLogging(t, func(t *testing.T, ctx context.Context) {
		t.Helper()

		log.Info(ctx, "info message", "with", "args")
		log.Debug(ctx, "debug this code for me please", "number", 1)
		log.Warn(ctx, "watch out!", os.ErrExist)
		log.Error(ctx, "something went wrong", io.EOF, "float", 1.234)

		err := errors.New("first", "1", 1)
		log.Warn(ctx, "err1", err)
		err = errors.Wrap(err, "second", "2", 2)
		log.Error(ctx, "err2", err)

		// Test attributes in context
		ctx1 := log.WithCtx(ctx, "ctx_key1", "ctx_value1")
		log.Debug(ctx1, "ctx debug message", "debug_key1", "debug_value1")
		ctx2 := log.WithCtx(ctx1, "ctx_key2", "ctx_value2")
		log.Info(ctx2, "ctx info message", "info_key2", "info_value2")
	})
}
