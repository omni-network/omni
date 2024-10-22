package app

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

func TestSaneMax(t *testing.T) {
	t.Parallel()

	network := netconf.Staging // Ephemeral chains have the highest thresholds.

	for _, role := range eoa.AllRoles() {
		thresh, ok := eoa.GetFundThresholds(tokens.ETH, network, role)
		if ok {
			expect := etherFloat(saneMax(tokens.ETH))
			actual := etherFloat(thresh.TargetBalance())
			require.GreaterOrEqual(t, expect, actual, "ETH %s %s", network, role)
		}

		thresh, ok = eoa.GetFundThresholds(tokens.OMNI, network, role)
		if ok {
			expect := etherFloat(saneMax(tokens.OMNI))
			actual := etherFloat(thresh.TargetBalance())
			require.GreaterOrEqual(t, expect, actual, "OMNI %s %s", network, role)
		}
	}
}

func etherFloat(b *big.Int) float64 {
	resp, _ := new(big.Int).Div(b, big.NewInt(params.Ether)).Float64()

	return resp
}
