package lib_test

import (
	"maps"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// allowed external packages with explanations.
var allowed = map[string]bool{
	"github.com/omni-network/omni/contracts/bindings":          true, // Bindings is the contracts API client library
	"github.com/omni-network/omni/contracts/allocs":            true, // Contains static data only
	"github.com/omni-network/omni/e2e/app/eoa":                 true, // EOA is also a library. TODO(corver): Move to lib/eoa
	"github.com/omni-network/omni/e2e/app/key":                 true, // [transitive] key is a library. TODO(corver): Move to lib/key
	"github.com/omni-network/omni/halo/genutil/evm/predeploys": true, // Contains static data only.
	"github.com/omni-network/omni/halo/sdk":                    true, // sdk contains common config required for cosmos clients
	"github.com/omni-network/omni/halo/comet":                  true, // comet is a lib containing cometbft client and types
	"github.com/omni-network/omni/solver/client":               true, // Solver client lib
	"github.com/omni-network/omni/solver/types":                true, // Solver client types lib
	"github.com/omni-network/omni/solver/fundthresh":           true, // Solver fund thresholds (imported by e2e/app/eoa)
	"github.com/omni-network/omni/anchor/anchorinbox":          true, // Anchor inbox program client lib
	"github.com/omni-network/omni/anchor/localnet":             true, // Anchor local static data.
	"github.com/omni-network/omni/halo/app/upgrades/static":    true, // Contains static data only.
	"github.com/omni-network/omni/scripts":                     true, // Contains static data only.
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice":   true, // gasprice.Tier type used by contract/xfeemgr. TODO(corver): Extract to types xfeemgr package.
	"github.com/omni-network/omni/monitor/xfeemngr/ticker":     true, // [Transitive] Would be solved by extracting gasprice.Tier to own types package.
	"github.com/omni-network/omni/halo/evmredenom":             true, // Static conversion functions.

	// `types` packages are required for clients of those APIs.
	"github.com/omni-network/omni/halo/attest/types":      true,
	"github.com/omni-network/omni/halo/portal/types":      true,
	"github.com/omni-network/omni/halo/valsync/types":     true,
	"github.com/omni-network/omni/octane/evmengine/types": true,
	"github.com/omni-network/omni/halo/registry/types":    true,
	"github.com/omni-network/omni/e2e/types":              true,
	"github.com/omni-network/omni/halo/genutil/genserve":  true, // Not called types, but contains clients API types.
}

// TestLibImports ensures that the lib packages does not import non-lib packages.
func TestLibImports(t *testing.T) {
	t.Parallel()

	bz1, err := exec.CommandContext(t.Context(), "go", "list", "-f", "{{.Imports}}", "./...").Output()
	require.NoError(t, err)
	bz2, err := exec.CommandContext(t.Context(), "go", "list", "-f", "{{.Deps}}", "./...").Output()
	require.NoError(t, err)

	found := maps.Clone(allowed)

	for _, line := range strings.Split(string(bz1)+string(bz2), "\n") {
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")

		require.NotContains(t, line, "[")

		for _, pkg := range strings.Fields(line) {
			if !strings.Contains(pkg, "github.com/omni-network/omni") {
				continue
			}

			if strings.Contains(pkg, "github.com/omni-network/omni/lib") {
				continue
			}

			if allowed[pkg] {
				delete(found, pkg)
				continue
			}

			require.Fail(t, "Illegal /lib/ import (note it may be transitive)", pkg)
		}
	}

	require.Empty(t, found, "Imports allowed but not actually imported. Please remove them from the allowed list.")
}
