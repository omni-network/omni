package app

import "testing"

// SetVoteWindowForT overrides the vote window for testing.
// Note this must be called before Start.
func SetVoteWindowForT(_ *testing.T, voteWindow uint64) {
	genesisVoteWindowVar = voteWindow
}
