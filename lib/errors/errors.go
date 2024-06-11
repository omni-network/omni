// Package errors provides a consistent interface for using errors.
// It also supports slog structured logging attributes; i.e. structured errors.
// It is also a drop-in replacement for the standard library errors package.
package errors

import (
	pkgerrors "github.com/pkg/errors"
)

// New returns an error that formats as the given text and
// contains the structured (slog) attributes.
func New(msg string, attrs ...any) error {
	return structured{
		err:   pkgerrors.New(msg),
		attrs: attrs,
	}
}

// Wrap returns a new error wrapping the provided with additional
// structured fields.
//
//nolint:wrapcheck,inamedparam // This function does custom wrapping and errors.
func Wrap(err error, msg string, attrs ...any) error {
	if err == nil {
		panic("wrap nil error")
	}

	// Support error types that do their own wrapping.
	if wrapper, ok := err.(interface{ Wrap(string, ...any) error }); ok {
		return wrapper.Wrap(msg, attrs...)
	}

	var inner structured
	if As(err, &inner) {
		attrs = append(attrs, inner.attrs...) // Append inner attributes
	}

	return structured{
		err:   pkgerrors.Wrap(err, msg),
		attrs: attrs,
	}
}
