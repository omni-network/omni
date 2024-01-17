// Package log provides a wrapper around the slog package (might change implementation later).
// It provides an opinionated interface for logging structured data always with context.
package log

import (
	"context"
	"log/slog"

	pkgerrors "github.com/pkg/errors" //nolint:revive // Need this for stacktraces.
)

type attrsKey struct{}

// WithCtx returns a copy of the context with which the logging attributes are associated.
// Usage:
//
//	ctx := log.WithCtx(ctx, "height", 1234)
//	...
//	log.Info(ctx, "Height processed") // Will contain attribute: height=1234
func WithCtx(ctx context.Context, attrs ...any) context.Context {
	return context.WithValue(ctx, attrsKey{}, mergeAttrs(ctx, attrs))
}

// Debug logs the message and attributes at default level.
func Debug(ctx context.Context, msg string, attrs ...any) {
	getLogger(ctx).DebugContext(ctx, msg, mergeAttrs(ctx, attrs)...)
}

// Info logs the message and attributes at info level.
func Info(ctx context.Context, msg string, attrs ...any) {
	getLogger(ctx).InfoContext(ctx, msg, mergeAttrs(ctx, attrs)...)
}

// Warn logs the message and error and attributes at warning level.
// If err is nil, it will not be logged.
func Warn(ctx context.Context, msg string, err error, attrs ...any) {
	if err != nil {
		attrs = append(attrs, "err", err)
		attrs = append(attrs, errAttrs(err)...)
	}

	getLogger(ctx).WarnContext(ctx, msg, mergeAttrs(ctx, attrs)...)
}

// Error logs the message and error and arguments at error level.
// If err is nil, it will not be logged.
func Error(ctx context.Context, msg string, err error, attrs ...any) {
	if err != nil {
		attrs = append(attrs, "err", err)
		attrs = append(attrs, errAttrs(err)...)
	}
	getLogger(ctx).ErrorContext(ctx, msg, mergeAttrs(ctx, attrs)...)
}

// errFields is similar to z.Err and returns the structured error fields and
// stack trace but without the error message. It avoids duplication of the error message
// since it is used as the main log message in Error above.
func errAttrs(err error) []any {
	type structErr interface {
		Attrs() []any
		StackTrace() pkgerrors.StackTrace
	}

	// Using cast instead of errors.As since no other wrapping library
	// is used and this avoids exporting the structured error type.
	serr, ok := err.(structErr) //nolint:errorlint // See comment above
	if !ok {
		return nil
	}

	return append(serr.Attrs(), slog.Any("stacktrace", serr.StackTrace()))
}

// mergeAttrs returns the attributes from the context merged with the provided attributes.
func mergeAttrs(ctx context.Context, attrs []any) []any {
	resp, _ := ctx.Value(attrsKey{}).([]any) //nolint:revive // We know the type.
	resp = append(resp, attrs...)

	verifyAttrs(resp)

	return resp
}
