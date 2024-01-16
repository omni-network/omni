package errors

import (
	pkgerrors "github.com/pkg/errors" //nolint:revive // This package wraps pkg/errors.
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

	trace, ok := s.err.(stackTracer) //nolint:errorlint // Using cast as per pkgerror documentation.
	if !ok {
		return nil
	}

	return trace.StackTrace()[1:] // Skip the first frame since pkgerrors doesn't support custom skipping.
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
