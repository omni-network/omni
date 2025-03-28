package validator

import (
	"flag"
	"fmt"
	"testing"

	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "enable integration tests")

func TestMonitorOnce(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration tests")
	}

	ctx := t.Context()

	cprov, err := provider.Dial(netconf.Omega)
	require.NoError(t, err)

	err = monitorOnce(ctx, cprov, func(s sample) {
		log.Info(ctx, "Sample validator", "sample", fmt.Sprintf("%#v", s))
	})
	require.NoError(t, err)
}
