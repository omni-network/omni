package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaybeRedactHexToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "empty",
			path:     "",
			expected: "",
		},
		{
			name:     "slash",
			path:     "/",
			expected: "/",
		},
		{
			name:     "not hex",
			path:     "/some/path/longerthan16characters",
			expected: "/some/path/longerthan16characters",
		},
		{
			name:     "Non-hex token",
			path:     "/some/path/nonhextoken1234567890",
			expected: "/some/path/nonhextoken1234567890",
		},
		{
			name:     "Short hex token",
			path:     "/some/path/1234",
			expected: "/some/path/1234",
		},
		{
			name:     "Hex token",
			path:     "/some/path/1234567890abcdef",
			expected: "/some/path/12..ef",
		},
		{
			name:     "quick node",
			path:     "/1234567890abcdef547740f492062c7f686d208f",
			expected: "/12..8f",
		},
		{
			name:     "infura",
			path:     "/v3/12341234123446ddb446a687c86a05d8",
			expected: "/v3/12..d8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := maybeRedactHexToken(tt.path)
			require.Equal(t, tt.expected, result)
		})
	}
}
