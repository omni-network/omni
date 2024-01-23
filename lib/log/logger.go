package log

import (
	"context"
	"io"
	"log/slog"
	"os"
	"testing"

	charm "github.com/charmbracelet/log"
)

var logger = newConsoleLogger() //nolint:gochecknoglobals // Global logger is our approach.

type loggerKey struct{}

// WithLogger returns a copy of the context with which the logger
// is associated replacing the default global logger.
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func getLogger(ctx context.Context) *slog.Logger {
	if l := ctx.Value(loggerKey{}); l != nil {
		return l.(*slog.Logger) //nolint:forcetypeassert,revive // We know the type.
	}

	return logger
}

// newConsoleLogger returns a new console logger for the following opinionated style:
// - Colored log levels (if tty supports it)
// - Timestamps are concise with millisecond precision
// - Timestamps and structured keys are faint
// This is aimed at local-dev and debugging. Production should use json or logfmt.
func newConsoleLogger(opts ...func(*TestOptions)) *slog.Logger {
	o := TestOptions{
		Writer:   os.Stderr,
		StubTime: false,
	}
	for _, opt := range opts {
		opt(&o)
	}

	timeFormat := "06-01-02 15:04:05.000"
	if o.StubTime {
		timeFormat = "00-00-00 00:00:00"
	}

	logger := charm.NewWithOptions(o.Writer, charm.Options{
		TimeFormat:      timeFormat,
		ReportTimestamp: true,
		Level:           charm.DebugLevel,
	})

	styles := charm.DefaultStyles()
	styles.Timestamp = styles.Timestamp.Faint(true)
	logger.SetStyles(styles)

	return slog.New(logger)
}

// TestOptions allow testing loggers.
type TestOptions struct {
	Writer   io.Writer // Write to some buffer
	StubTime bool      // Stub time in tests for deterministic output.
}

// LoggersForT returns a map of loggers for testing.
func LoggersForT(_ *testing.T) map[string]func(...func(*TestOptions)) *slog.Logger {
	return map[string]func(...func(*TestOptions)) *slog.Logger{
		"console": newConsoleLogger,
	}
}

// WithNoopLogger returns a copy of the context with a noop logger which discards all logs.
func WithNoopLogger(ctx context.Context) context.Context {
	return WithLogger(ctx, newConsoleLogger(func(o *TestOptions) {
		o.Writer = io.Discard
	}))
}
