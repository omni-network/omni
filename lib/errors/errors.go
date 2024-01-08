// Package errors provides a consistent interface for using errors wrapping our chosen errors library.
// It is a drop-in replacement for the standard library errors package.
//
//nolint:wrapcheck // Wrapping not needed in this package.
package errors

import (
	stderrors "errors" //nolint:revive // This is the only
	"fmt"
)

// New returns a new error with the given message.
func New(msg string) error {
	return stderrors.New(msg)
}

// Wrap returns a new error wrapping the given error with the given message.
func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err) //nolint:forbidigo // This is only place where we can use fmt.Errorf.
}

// As finds the first error in err's tree that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
func As(err error, target any) bool {
	return stderrors.As(err, target)
}
