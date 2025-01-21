package solvernet_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestParseFillOriginData(t *testing.T) {
	t.Parallel()

	f := fuzz.New().NilChance(0)

	// big.Ints don't fuzz well, so we provide a custom fuzzer
	f.Funcs(func(bi *big.Int, c fuzz.Continue) {
		var val uint64
		c.Fuzz(&val)
		bi.SetUint64(val)
	})

	var data bindings.ISolverNetFillOriginData
	f.Fuzz(&data)

	packed, err := solvernet.PackFillOriginData(data)
	require.NoError(t, err)

	parsed, err := solvernet.ParseFillOriginData(packed)
	require.NoError(t, err)

	require.Equal(t, data, parsed)
}
