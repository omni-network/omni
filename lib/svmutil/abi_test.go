package svmutil_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/umath"

	"github.com/stretchr/testify/require"
)

func TestU128(t *testing.T) {
	t.Parallel()

	ensure := func(t *testing.T, i *big.Int) {
		t.Helper()

		u, err := svmutil.U128(i)
		require.NoError(t, err)
		require.Equal(t, i.String(), u.String())
		tutil.RequireEQ(t, i, u.BigInt())
	}

	ensure(t, bi.N(0))
	ensure(t, bi.N(1))
	ensure(t, bi.N(2_000_000))
	ensure(t, bi.N(123_456_789_123))
	ensure(t, umath.MaxUint128)

	_, err := svmutil.U128(umath.MaxUint256)
	require.Error(t, err)
}
