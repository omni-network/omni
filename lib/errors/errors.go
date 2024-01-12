// Package errors provides a consistent interface for using errors.
// It also supports slog structured logging attributes; i.e. structured errors.
// It is also a drop-in replacement for the standard library errors package.
package errors

import (
	stderrors "errors" //nolint:revive // This package wraps std errors.
	"fmt"
)

// New returns an error that formats as the given text and
// contains the structured (slog) attributes.
func New(msg string, attrs ...any) error {
	return structured{
		err:   stderrors.New(msg),
		attrs: attrs,
	}
}

// Wrap returns a new error wrapping the provided with additional
// structured fields.
func Wrap(err error, msg string, attrs ...any) error {
	var inner structured
	if As(err, &inner) {
		attrs = append(attrs, inner.attrs...) // Append inner attributes
	}

	return structured{
		err:   fmt.Errorf("%s: %w", msg, err), //nolint:forbidigo // Wrap error message using stdlib.
		attrs: attrs,
	}
}
