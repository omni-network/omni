// Package log provides a wrapper around the slog package (might change implementation later).
// It provides an opinionated interface for logging structured data always with context.
package log

import (
	"context"
	"log/slog"
	"runtime"
	"strings"
	"time"

	pkgerrors "github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type attrsKey struct{}
type skipKey struct{}

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
	log(ctx, slog.LevelDebug, msg, mergeAttrs(ctx, attrs)...)
}

// Info logs the message and attributes at info level.
func Info(ctx context.Context, msg string, attrs ...any) {
	log(ctx, slog.LevelInfo, msg, mergeAttrs(ctx, attrs)...)
}

// Warn logs the message and error and attributes at warning level.
// If err is nil, it will not be logged.
func Warn(ctx context.Context, msg string, err error, attrs ...any) {
	if err != nil {
		attrs = append(attrs, slog.String("err", err.Error()))
		attrs = append(attrs, errAttrs(err)...)
	}

	log(ctx, slog.LevelWarn, msg, mergeAttrs(ctx, attrs)...)
}

// Error logs the message and error and arguments at error level.
// If err is nil, it will not be logged.
func Error(ctx context.Context, msg string, err error, attrs ...any) {
	if err != nil {
		attrs = append(attrs, slog.String("err", err.Error()))
		attrs = append(attrs, errAttrs(err)...)
	}
	log(ctx, slog.LevelError, msg, mergeAttrs(ctx, attrs)...)
}

// log is the low-level logging method for methods that take ...any.
// It must always be called directly by an exported logging method
// or function, because it uses a fixed call depth to obtain the pc.
//
// Copied from stdlib since we wrap slog and the source/caller is incorrect otherwise.
// See https://github.com/golang/go/blob/master/src/log/slog/logger.go#L241
func log(ctx context.Context, level slog.Level, msg string, attrs ...any) {
	logTotal.WithLabelValues(strings.ToLower(level.String())).Inc()

	logger := getLogger(ctx)

	if !logger.Enabled(ctx, level) {
		return
	}

	// Default skip [runtime.Callers, this function, this function's caller]
	var skip = 3
	if v, ok := ctx.Value(skipKey{}).(int); ok {
		skip = v
	}

	var pcs [1]uintptr
	runtime.Callers(skip, pcs[:])

	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(attrs...)

	// Build trace event
	traceAttrs := []attribute.KeyValue{attribute.String("msg", msg)}
	r.Attrs(func(attr slog.Attr) bool {
		traceAttrs = append(traceAttrs, attribute.Stringer(attr.Key, attr.Value))
		return true
	})
	trace.SpanFromContext(ctx).AddEvent(
		"log"+level.String(),
		trace.WithAttributes(traceAttrs...),
	)

	_ = logger.Handler().Handle(ctx, r)
}

// errFields is similar to z.Err and returns the structured error fields and
// stack trace but without the error message. It avoids duplication of the error message
// since it is used as the main log message in Error above.
func errAttrs(err error) []any {
	type stackTracer interface {
		StackTrace() pkgerrors.StackTrace
	}

	type omniErr interface {
		Attrs() []any
	}

	var attrs []any
	var stack pkgerrors.StackTrace

	// Go up the cause chain (from the outermost error to the innermost error)
	for {
		// Use the first encountered omniErr's attributes.
		if len(attrs) == 0 {
			if serr, ok := err.(omniErr); ok {
				attrs = serr.Attrs()
			}
		}

		// Use the last encountered stack trace.
		if serr, ok := err.(stackTracer); ok {
			stack = serr.StackTrace()
		}

		if cause := pkgerrors.Unwrap(err); cause != nil {
			err = cause
			continue // Continue up the cause chain.
		}

		// Root cause reached, break the loop.

		if len(stack) > 0 {
			attrs = append(attrs, slog.Any("stacktrace", stack))
		}

		return attrs
	}
}

// mergeAttrs returns the attributes from the context merged with the provided attributes.
func mergeAttrs(ctx context.Context, attrs []any) []any {
	resp, _ := ctx.Value(attrsKey{}).([]any) //nolint:revive // We know the type.
	resp = append(resp, attrs...)

	return resp
}

// WithSkip returns a copy of the context with the skip value set.
// This is used to control the number of stack frames to skip when `source` is calculated.
func WithSkip(ctx context.Context, skip int) context.Context {
	return context.WithValue(ctx, skipKey{}, skip)
}
