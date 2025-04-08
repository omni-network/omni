package app

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenmeta"

	"github.com/stretchr/testify/require"
)

func TestSaneMax(t *testing.T) {
	t.Parallel()

	network := netconf.Staging // Ephemeral chains have the highest thresholds.

	for _, role := range eoa.AllRoles() {
		thresh, ok := eoa.GetFundThresholds(tokenmeta.ETH, network, role)
		if ok {
			expect := bi.ToEtherF64(saneMax(tokenmeta.ETH))
			actual := bi.ToEtherF64(thresh.TargetBalance())
			require.GreaterOrEqual(t, expect, actual, "ETH %s %s", network, role)
		}

		thresh, ok = eoa.GetFundThresholds(tokenmeta.OMNI, network, role)
		if ok {
			expect := bi.ToEtherF64(saneMax(tokenmeta.OMNI))
			actual := bi.ToEtherF64(thresh.TargetBalance())
			require.GreaterOrEqual(t, expect, actual, "OMNI %s %s", network, role)
		}
	}
}
