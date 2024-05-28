package voter

import (
	"context"
	"testing"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"

	"github.com/stretchr/testify/require"
)

// LoadVoterForT is a helper function to load a voter for testing.
// It sets the backoff period to 1ms.
func LoadVoterForT(t *testing.T, privKey crypto.PrivKey, path string, provider xchain.Provider,
	deps types.VoterDeps, network netconf.Network, backoff func(),
) *Voter {
	t.Helper()
	v, err := LoadVoter(privKey, path, provider, deps, network)
	require.NoError(t, err)

	v.backoffFunc = func(ctx context.Context) func() { return backoff }

	return v
}

// LatestByChain returns the latest vote by chain for testing purposes only.
func (v *Voter) LatestByChain(chainVer xchain.ChainVersion) (*types.Vote, bool) {
	return v.latestByChain(chainVer)
}
