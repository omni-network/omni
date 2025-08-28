package app

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/stretchr/testify/require"
)

func TestSaneMax(t *testing.T) {
	t.Parallel()

	network := netconf.Staging // Ephemeral chains have the highest thresholds.

	for _, role := range eoa.AllRoles() {
		thresh, ok := eoa.GetFundThresholds(tokens.ETH, network, role)
		if ok {
			expect := bi.ToEtherF64(saneMax(tokens.ETH))
			actual := bi.ToEtherF64(thresh.TargetBalance())
			require.GreaterOrEqual(t, expect, actual, "ETH %s %s", network, role)
		}

		// TODO(zodomo): remove this once we have deprecated OMNI.
		thresh, ok = eoa.GetFundThresholds(tokens.OMNI, network, role)
		if ok {
			expect := bi.ToEtherF64(saneMax(tokens.OMNI))
			actual := bi.ToEtherF64(thresh.TargetBalance())
			require.GreaterOrEqual(t, expect, actual, "OMNI %s %s", network, role)
		}

		thresh, ok = eoa.GetFundThresholds(tokens.NOM, network, role)
		if ok {
			expect := bi.ToEtherF64(saneMax(tokens.NOM))
			actual := bi.ToEtherF64(thresh.TargetBalance())
			require.GreaterOrEqual(t, expect, actual, "NOM %s %s", network, role)
		}
	}
}
