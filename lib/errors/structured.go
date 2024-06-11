package errors

import (
	pkgerrors "github.com/pkg/errors"
)

// structured is the implementation of a structured error.
type structured struct {
	err   error
	attrs []any
}

// StackTrace implements the pkgerrors.StrackTracer interface.
func (s structured) StackTrace() pkgerrors.StackTrace {
	type stackTracer interface {
		StackTrace() pkgerrors.StackTrace
	}

	tracer, ok := s.err.(stackTracer)
	if !ok {
		return nil
	}

	trace := tracer.StackTrace()

	// Skip the first frame as this is always this package (can't skip it via pkgerrors API).
	// Drop the last frame as that is always the runtime package.
	return trace[1 : len(trace)-1]
}

// Error returns the error message and implements the error interface.
func (s structured) Error() string {
	return s.err.Error()
}

// Attrs returns the structured slog attributes.
func (s structured) Attrs() []any {
	return s.attrs
}

// Unwrap returns the underlying error and
// provides compatibility with stdlib errors.
func (s structured) Unwrap() error {
	return s.err
}

// Is returns true if err is equaled to this structured error.
func (s structured) Is(err error) bool {
	var other structured
	if !pkgerrors.As(err, &other) {
		return false
	}

	return pkgerrors.Is(s.err, other.err)
}
