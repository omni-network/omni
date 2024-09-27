//nolint:forbidigo,govet,staticcheck // We use cosmos errors explicitly.
package app

import (
	"fmt"
	"io"
	"testing"

	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/x/upgrade"
	utypes "cosmossdk.io/x/upgrade/types"
	"github.com/stretchr/testify/require"
)

func TestIsErrOldBinary(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		err     error
		want    bool
		upgrade string
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
			name:    "wrong version error",
			err:     fmt.Errorf("wrong app version %d, upgrade handler is missing for %s upgrade plan", 99, "test"),
			want:    true,
			upgrade: "test",
		},
		{
			name: "wrapped wrong version error",
			err: errors.Wrap(
				fmt.Errorf("wrong app version %d, upgrade handler is missing for %s upgrade plan", 98, "genesis upgrade"),
				"wrapper"),
			want:    true,
			upgrade: "genesis upgrade",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			upgrade, ok := isErrOldBinary(tt.err)
			require.Equal(t, tt.want, ok)
			require.Equal(t, tt.upgrade, upgrade)
		})
	}
}

func TestIsErrUpgradeNeeded(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		err     error
		want    bool
		upgrade string
	}{
		{
			name: "nil error",
			err:  nil,
		},
		{
			name: "eof error",
			err:  io.EOF,
		},
		{
			name:    "wrong version error",
			err:     fmt.Errorf(upgrade.BuildUpgradeNeededMsg(utypes.Plan{Name: "1_uluwatu", Height: 1, Info: "genesis upgrade"})),
			want:    true,
			upgrade: "1_uluwatu",
		},
		{
			name: "wrapped wrong version error",
			err: errors.Wrap(
				fmt.Errorf(upgrade.BuildUpgradeNeededMsg(utypes.Plan{Name: "1_uluwatu", Height: 1, Info: "genesis upgrade"})),
				"wrapper"),
			want:    true,
			upgrade: "1_uluwatu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			upgrade, ok := isErrUpgradeNeeded(tt.err)
			require.Equal(t, tt.want, ok)
			require.Equal(t, tt.upgrade, upgrade)
		})
	}
}
