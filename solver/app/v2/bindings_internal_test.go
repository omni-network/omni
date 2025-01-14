package appv2

import (
	"math/big"
	"testing"

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

	var data FillOriginData
	f.Fuzz(&data)

	packed, err := inputsFillOriginData.Pack(data)
	require.NoError(t, err)

	parsed, err := parseFillOriginData(packed)
	require.NoError(t, err)

	require.Equal(t, data, parsed)
}
