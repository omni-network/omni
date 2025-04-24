package targets

import (
	"flag"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

//go:generate go test . -integration -run=TestIntegration -v

func TestIntegration(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration tests")
	}

	err := refreshOnce(t.Context(), netconf.Mainnet)
	require.NoError(t, err)

	// assert known symbiotic mainnet vault in targets
	target, ok := Get(evmchain.IDEthereum, common.HexToAddress("0xC329400492c6ff2438472D4651Ad17389fCb843a"))
	require.True(t, ok)
	require.Equal(t, "Symbiotic:wstETH", target.Name)
}
