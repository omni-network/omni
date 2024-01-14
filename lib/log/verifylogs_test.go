//go:build verify_logs

package log_test

import (
	"context"
	"github.com/omni-network/omni/lib/log"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVerifyLogs(t *testing.T) {
	t.Parallel()

	badKeys := []string{
		"",
		"Title",
		"CAPITALS",
		"PascalCase",
		"camelCase",
		"spaces in key",
		"hyphens-in-key",
		"special chars:!@#$%^&*()",
	}
	for _, key := range badKeys {
		t.Run(key, func(t *testing.T) {
			require.Panics(t, func() {
				log.Info(context.Background(), "Test bad log key", key, "value")
			})
		})
	}

	require.Panics(t, func() {
		log.Info(context.Background(), "Test bad log key", "missing_value")
	})

	goodKeys := []string{
		"lowercase",
		"numbers123",
		"under_score",
	}

	for _, key := range goodKeys {
		t.Run(key, func(t *testing.T) {
			require.NotPanics(t, func() {
				log.Info(context.Background(), "Test good log key", key, "value")
			})
		})
	}
}
