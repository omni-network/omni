//nolint:testifylint // Using assert here to log error afterwards.
package tutil

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RequireNoError asserts that err is nil. It also logs the error to show the stacktrace.
//
// This can be used instead of require.NoError(t, err) to show the stacktrace in case of error.
func RequireNoError(tb testing.TB, err error) {
	tb.Helper()

	if !assert.NoErrorf(tb, err, "See log line for error details") {
		log.Error(tb.Context(), "Unexpected error", err)
		tb.FailNow()
	}
}

// RequireIsPositive asserts that a is positive.
func RequireIsPositive(tb testing.TB, a *big.Int, msgAndArgs ...any) {
	tb.Helper()

	if bi.IsPositive(a) {
		return
	}

	require.Fail(tb, fmt.Sprintf("%s is not positive", a), msgAndArgs...)
}

// RequireIsZero asserts that a is zero.
func RequireIsZero(tb testing.TB, a *big.Int, msgAndArgs ...any) {
	tb.Helper()

	if bi.IsZero(a) {
		return
	}

	require.Fail(tb, fmt.Sprintf("%s is not zero", a), msgAndArgs...)
}

// RequireEQ asserts that a and b are equal.
func RequireEQ(tb testing.TB, a, b *big.Int, msgAndArgs ...any) {
	tb.Helper()

	if bi.EQ(a, b) {
		return
	}

	require.Fail(tb, fmt.Sprintf("%s is not equal to %s", a, b), msgAndArgs...)
}

// RequireGT asserts that a is greater than b.
func RequireGT(tb testing.TB, a, b *big.Int, msgAndArgs ...any) {
	tb.Helper()

	if bi.GT(a, b) {
		return
	}

	require.Fail(tb, fmt.Sprintf("%s is not greater than %s", a, b), msgAndArgs...)
}

// RequireGTE asserts that a is greater than or equal to b.
func RequireGTE(tb testing.TB, a, b *big.Int, msgAndArgs ...any) {
	tb.Helper()

	if bi.GTE(a, b) {
		return
	}

	require.Fail(tb, fmt.Sprintf("%s is not greater than or equal to %s", a, b), msgAndArgs...)
}

// RequireLTE asserts that a is less than or equal to b.
func RequireLTE(tb testing.TB, a, b *big.Int, msgAndArgs ...any) {
	tb.Helper()

	if bi.LTE(a, b) {
		return
	}

	require.Fail(tb, fmt.Sprintf("%s is not less than or equal to %s", a, b), msgAndArgs...)
}
