//nolint:revive,testifylint // Using assert here to log error afterwards.
package tutil

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/log"

	"github.com/stretchr/testify/assert"
)

// RequireNoError asserts that err is nil. It also logs the error to show the stacktrace.
func RequireNoError(tb testing.TB, err error) {
	tb.Helper()

	if !assert.NoErrorf(tb, err, "See log line for error details") {
		log.Error(context.Background(), "Unexpected error", err)
		tb.FailNow()
	}
}
