// Package log provides a wrapper around the slog package (might change implementation later).
// It provides an opinionated interface for logging structured data always with context.
//
// TODO(corver): Add support for custom loggers in the context.
package log

import (
	"context"
	"log/slog"
	"os"

	charm "github.com/charmbracelet/log"
)

var logger = newConsoleLogger() //nolint:gochecknoglobals // Global logger is our approach.

// Debug logs the message and arguments at default level.
func Debug(ctx context.Context, msg string, args ...any) {
	logger.DebugContext(ctx, msg, args...)
}

// Info logs the message and arguments at info level.
func Info(ctx context.Context, msg string, args ...any) {
	logger.InfoContext(ctx, msg, args...)
}

// Warn logs the message and arguments at warning level.
func Warn(ctx context.Context, msg string, args ...any) {
	logger.WarnContext(ctx, msg, args...)
}

// Error logs the message and error and arguments at error level.
// If err is nil, it will not be logged.
func Error(ctx context.Context, msg string, err error, args ...any) {
	if err != nil {
		args = append(args, "err", err)
	}
	logger.ErrorContext(ctx, msg, args...)
}

// newConsoleLogger returns a new console logger for the following opinionated style:
// - Colored log levels (if tty supports it)
// - Timestamps are concise with millisecond precision
// - Timestamps and structured keys are faint
// This is aimed at local-dev and debugging. Production should use json or logfmt.
func newConsoleLogger() *slog.Logger {
	logger := charm.NewWithOptions(os.Stderr, charm.Options{
		TimeFormat:      "06-01-02 15:04:05.000",
		ReportTimestamp: true,
		Level:           charm.DebugLevel,
	})

	styles := charm.DefaultStyles()
	styles.Timestamp = styles.Timestamp.Faint(true)
	logger.SetStyles(styles)

	return slog.New(logger)
}
