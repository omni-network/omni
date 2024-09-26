package app

import (
	"fmt"
	"io"
	"testing"

	"github.com/omni-network/omni/lib/errors"

	"github.com/stretchr/testify/require"
)

//nolint:forbidigo // We use cosmos errors explicitly.
func TestIsErrOldBinary(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "eof error",
			err:  io.EOF,
			want: false,
		},
		{
			name: "wrong version error",
			err:  fmt.Errorf("wrong app version %d, upgrade handler is missing for %s upgrade plan", 99, "test"),
			want: true,
		},
		{
			name: "wrapped wrong version error",
			err: errors.Wrap(
				fmt.Errorf("wrong app version %d, upgrade handler is missing for %s upgrade plan", 98, "wrapped"),
				"wrapper"),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ok := isErrOldBinary(tt.err)
			require.Equal(t, tt.want, ok)
		})
	}
}
