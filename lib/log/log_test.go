package log_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"

	pkgerrors "github.com/pkg/errors"
)

//go:generate go test . -golden -clean

// TestSimpleLogs tests the simple logs.
func TestSimpleLogs(t *testing.T) {
	t.Parallel()

	AssertLogging(t, func(t *testing.T, ctx context.Context) {
		t.Helper()
		{ // Test some basic logging
			log.Info(ctx, "info message", "with", "args")
			log.Debug(ctx, "debug this code for me please", "number", 1)
			log.Warn(ctx, "watch out!", os.ErrExist)
			log.Error(ctx, "something went wrong", io.EOF, "float", 1.234)
		}
		{ // Test errors wrapping and logging
			err := errors.New("first", "1", 1)
			log.Warn(ctx, "err1", err)
			err = errors.Wrap(err, "second", "2", 2)
			log.Error(ctx, "err2", err)
		}
		{ // Test attributes in context
			ctx1 := log.WithCtx(ctx, "ctx_key1", "ctx_value1")
			log.Debug(ctx1, "ctx debug message", "debug_key1", "debug_value1")
			ctx2 := log.WithCtx(ctx1, "ctx_key2", "ctx_value2")
			log.Info(ctx2, "ctx info message", "info_key2", "info_value2")
		}
		{ // Test cometBFT wrapping our errors in pkg errors
			err := errors.New("new", "new", "new")
			err = errors.Wrap(err, "omni wrap", "wrap", "wrap")
			err = pkgerrors.Wrap(err, "pkg wrap")
			log.Warn(ctx, "Pkg wrapped error", err)
		}
	})
}

// AssertLogging returns a function that will assert all loggers' output against
// golden test files.
func AssertLogging(t *testing.T, testFunc func(*testing.T, context.Context)) {
	t.Helper()

	loggers := log.LoggersForT()

	for name, initFunc := range loggers {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer
			logger := initFunc(&buf)

			ctx := context.Background()
			ctx = log.WithLogger(ctx, logger)

			testFunc(t, ctx)

			tutil.RequireGoldenBytes(t, buf.Bytes())
		})
	}
}
