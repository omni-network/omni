// Package log provides a wrapper around the slog package (might change implementation later).
// It provides an opinionated interface for logging structured data always with context.
package log

import (
	"context"
)

// Debug logs the message and arguments at default level.
func Debug(ctx context.Context, msg string, args ...any) {
	getLogger(ctx).DebugContext(ctx, msg, args...)
}

// Info logs the message and arguments at info level.
func Info(ctx context.Context, msg string, args ...any) {
	getLogger(ctx).InfoContext(ctx, msg, args...)
}

// Warn logs the message and error and arguments at warning level.
// If err is nil, it will not be logged.
func Warn(ctx context.Context, msg string, err error, args ...any) {
	if err != nil {
		args = append(args, "err", err)
		args = append(args, errAttrs(err)...)
	}
	getLogger(ctx).WarnContext(ctx, msg, args...)
}

// Error logs the message and error and arguments at error level.
// If err is nil, it will not be logged.
func Error(ctx context.Context, msg string, err error, args ...any) {
	if err != nil {
		args = append(args, "err", err)
		args = append(args, errAttrs(err)...)
	}
	getLogger(ctx).ErrorContext(ctx, msg, args...)
}

// errFields is similar to z.Err and returns the structured error fields and
// stack trace but without the error message. It avoids duplication of the error message
// since it is used as the main log message in Error above.
func errAttrs(err error) []any {
	type structErr interface {
		Attrs() []any
	}

	// Using cast instead of errors.As since no other wrapping library
	// is used and this avoids exporting the structured error type.
	serr, ok := err.(structErr) //nolint:errorlint // See comment above
	if !ok {
		return nil
	}

	return serr.Attrs()
}
