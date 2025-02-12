// Package errors provides a consistent interface for using errors.
// It also supports slog structured logging attributes; i.e. structured errors.
// It is also a drop-in replacement for the standard library errors package.
package errors

import (
	stderrors "errors"

	pkgerrors "github.com/pkg/errors"
)

// New returns an error that formats as the given text and
// contains the structured (slog) attributes and stack trace.
func New(msg string, attrs ...any) error {
	return structured{
		err:   pkgerrors.New(msg),
		attrs: attrs,
	}
}

// NewSentinel returns a new error that formats as the given text.
// It doesn't contain a stack trace.
// This can be used to support error checking with proper runtime stack traces.
//
//	// ErrNotFound is a sentinel error, it doesn't have a stack trace.
//	var ErrNotFound = errors.NewSentinel("not found")
//
//	// Foo returns a sentinel error with runtime stack trace of this function
//	// instead of the stack trace of ErrNotFound initialization.
//	func Foo() error {
//	  return errors.Wrap(ErrNotFound, "foo failed")
//	}
//
//	// Usage
//	if errors.Is(Foo(), ErrNotFound) {
//	  // Handle ErrNotFound or log it with proper runtime stack traces.
func NewSentinel(msg string) error {
	return stderrors.New(msg) //nolint:wrapcheck // This is explicitly not wrapped.
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

// Cause calls Unwrap until it finds the root cause of the error.
func Cause(err error) error {
	cause := err
	for {
		next := pkgerrors.Unwrap(cause)
		if next == nil {
			return cause
		}
		cause = next
	}
}
