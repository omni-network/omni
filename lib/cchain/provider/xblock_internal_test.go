package provider

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/test/tutil"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestXBlock(t *testing.T) {
	t.Parallel()
	f := fuzz.NewWithSeed(99).NilChance(0)

	var height uint64
	f.Fuzz(&height)

	valFunc := func(ctx context.Context, h uint64) ([]cchain.Validator, bool, error) {
		require.EqualValues(t, height, h)
		var resp []cchain.Validator
		f.Fuzz(&resp)

		return resp, true, nil
	}
	prov := Provider{valset: valFunc}

	block, ok, err := prov.XBlock(context.Background(), height)
	require.NoError(t, err)
	require.True(t, ok)
	tutil.RequireGoldenJSON(t, block)
}
