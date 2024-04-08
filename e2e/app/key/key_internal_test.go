package key

import (
	"context"
	"os/exec"
	"testing"

	"github.com/omni-network/omni/lib/netconf"

	"github.com/stretchr/testify/require"
)

func DeleteSecretForT(ctx context.Context, t *testing.T, network netconf.ID, name string, typ Type, addr string) {
	t.Helper()
	secret := secretName(network, name, typ, addr)

	out, err := exec.CommandContext(ctx, "gcloud", "secrets", "delete", secret, "--quiet").CombinedOutput()
	require.NoError(t, err, string(out))
}
