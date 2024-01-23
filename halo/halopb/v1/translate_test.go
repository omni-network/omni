package v1_test

import (
	"testing"

	halopb "github.com/omni-network/omni/halo/halopb/v1"
	"github.com/omni-network/omni/lib/xchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestTranslate(t *testing.T) {
	t.Parallel()
	var aggs []xchain.AggAttestation
	fuzz.New().NilChance(0).NumElements(1, 8).Fuzz(&aggs)

	aggpb := halopb.AggregatesToProto(aggs)
	aggs2, err := halopb.AggregatesFromProto(aggpb)
	require.NoError(t, err)

	require.Equal(t, aggs, aggs2)
}
