package types_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -clean -golden

func TestAttestIntervals(t *testing.T) {
	t.Parallel()

	type tuple struct {
		BlockPeriod             string `json:"block_period"`
		EphemeralAttestInterval uint64 `json:"ephemeral_attest_interval"`
		ProtectedAttestInterval uint64 `json:"protected_attest_interval"`
	}

	resp := make(map[string]tuple)

	for _, network := range []netconf.ID{netconf.Omega, netconf.Mainnet} {
		chains, err := manifests.EVMChains(network)
		require.NoError(t, err)

		for _, metadata := range chains {
			resp[metadata.Name] = tuple{
				BlockPeriod:             metadata.BlockPeriod.String(),
				EphemeralAttestInterval: types.EVMChain{Metadata: metadata}.AttestInterval(netconf.Staging),
				ProtectedAttestInterval: types.EVMChain{Metadata: metadata}.AttestInterval(netconf.Mainnet),
			}
		}
	}

	tutil.RequireGoldenJSON(t, resp)
}

func TestNextRPCAddress(t *testing.T) {
	t.Parallel()
	c := types.NewPublicChain(types.EVMChain{}, []string{"1 ", " 2", "3"})

	require.Equal(t, "1", c.NextRPCAddress())
	require.Equal(t, "2", c.NextRPCAddress())
	require.Equal(t, "3", c.NextRPCAddress())
	require.Equal(t, "1", c.NextRPCAddress())
}
