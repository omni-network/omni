package log

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/omni-network/omni/test/tutil"
)

// AssertLogging returns a function that will assert all loggers' output against
// golden test files.
func AssertLogging(t *testing.T, testFunc func(*testing.T, context.Context)) {
	t.Helper()

	loggers := map[string]func(...func(*options)) *slog.Logger{
		"console": newConsoleLogger,
	}

	for name, initFunc := range loggers {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := initFunc(func(config *options) {
				config.Writer = &buf
				config.StubTime = true
			})

			ctx := context.Background()
			ctx = WithLogger(ctx, logger)

			testFunc(t, ctx)

			tutil.RequireGoldenBytes(t, buf.Bytes())
		})
	}
}
