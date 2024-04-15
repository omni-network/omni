package voter

import (
	"context"
	"testing"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"

	"github.com/stretchr/testify/require"
)

// LoadVoterForT is a helper function to load a voter for testing.
// It sets the backoff period to 1ms.
func LoadVoterForT(t *testing.T, privKey crypto.PrivKey, path string, provider xchain.Provider,
	deps types.VoterDeps, chains map[uint64]string, backoff func(),
) *Voter {
	t.Helper()
	v, err := LoadVoter(privKey, path, provider, deps, chains)
	require.NoError(t, err)

	v.backoffFunc = func(ctx context.Context) func() { return backoff }

	return v
}

// LatestByChain returns the latest vote by chain for testing purposes only.
func (v *Voter) LatestByChain(chainID uint64) (*types.Vote, bool) {
	return v.latestByChain(chainID)
}
