package app_test

import (
	"sort"
	"testing"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/docker"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -clean -golden

func TestManifestServiceReference(t *testing.T) {
	t.Parallel()
	manifestFiles := []string{
		"../manifests/staging.toml",
		"../manifests/omega.toml",
	}
	ref := make(map[netconf.ID][]string)
	for _, manifestFile := range manifestFiles {
		manifest, err := app.LoadManifest(manifestFile)
		require.NoError(t, err)

		infraData, err := docker.NewInfraData(manifest)
		require.NoError(t, err)

		services := keys(infraData.Instances)
		services = append(services, "relayer", "monitor")
		sort.Strings(services)
		ref[manifest.Network] = services
	}

	tutil.RequireGoldenJSON(t, ref)
}

func keys[K comparable, V any](m map[K]V) []K {
	var resp []K
	for k := range m {
		resp = append(resp, k)
	}

	return resp
}
