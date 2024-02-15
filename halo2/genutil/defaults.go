package genutil

import "github.com/cometbft/cometbft/types"

// DefaultConsensusParams returns the default cometBFT consensus params for omni protocol.
func DefaultConsensusParams() *types.ConsensusParams {
	resp := types.DefaultConsensusParams()
	resp.ABCI.VoteExtensionsEnableHeight = 1                             // Enable vote extensions from the start.
	resp.Validator.PubKeyTypes = []string{types.ABCIPubKeyTypeSecp256k1} // Only k1 keys.

	return resp
}
