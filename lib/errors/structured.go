package errors

import (
	stderrors "errors" //nolint:revive // This package wraps std errors.
)

// structured is the implementation of a structured error.
type structured struct {
	err   error
	attrs []any
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
	if !stderrors.As(err, &other) {
		return false
	}

	return stderrors.Is(s.err, other.err)
}
