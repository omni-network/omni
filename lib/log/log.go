// Package log provides a wrapper around the slog package (might change implementation later).
// It provides an opinionated interface for logging structured data always with context.
//
// TODO(corver): Add support for custom loggers in the context.
package log

import (
	"context"
	"log/slog"
)

// Debug logs the message and arguments at default level.
func Debug(ctx context.Context, msg string, args ...any) {
	slog.DebugContext(ctx, msg, args...)
}

// Info logs the message and arguments at info level.
func Info(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, args...)
}

// Warn logs the message and arguments at warning level.
func Warn(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, args...)
}

// Error logs the message and error and arguments at error level.
// If err is nil, it will not be logged.
func Error(ctx context.Context, msg string, err error, args ...any) {
	if err != nil {
		args = append(args, "err", err)
	}
	slog.ErrorContext(ctx, msg, args...)
}
