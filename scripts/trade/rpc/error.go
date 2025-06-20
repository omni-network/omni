package rpc

import (
	"log/slog"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"
)

const label = "status_code"

// StatusAttr returns a slog.Attr for the given status code.
func StatusAttr(statusCode int) slog.Attr {
	return slog.Int(label, statusCode)
}

// getStatusCode extracts the status code from an error if it has one.
func getStatusCode(err error) (int, bool) {
	if err == nil {
		return 0, false
	}

	attr, ok := getAttr(err, label)
	if !ok || attr.Value.Kind() != slog.KindInt64 {
		return 0, false // Attribute not found or invalid
	}

	resp, err := umath.ToInt(attr.Value.Int64())

	return resp, err == nil
}

// errFields is similar to z.Err and returns the structured error fields and
// stack trace but without the error message. It avoids duplication of the error message
// since it is used as the main log message in Error above.
func getAttr(err error, key string) (slog.Attr, bool) {
	type attrErr interface {
		Attrs() []any
	}

	// Go up the cause chain (from the outermost error to the innermost error)
	for {
		if aErr, ok := err.(attrErr); ok {
			for _, anyAttr := range aErr.Attrs() {
				attr, ok := anyAttr.(slog.Attr)
				if !ok {
					continue // Skip non-slog attributes.
				}

				if attr.Key == key {
					return attr, true // Found the attribute.
				}
			}
		}

		if cause := errors.Unwrap(err); cause != nil {
			err = cause
			continue // Continue up the cause chain.
		}

		// Root cause reached, break the loop.

		return slog.Attr{}, false // Attribute not found.
	}
}
