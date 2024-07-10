package genutil

import "github.com/cometbft/cometbft/types"

// DefaultConsensusParams returns the default cometBFT consensus params for omni protocol.
func DefaultConsensusParams() *types.ConsensusParams {
	resp := types.DefaultConsensusParams()
	resp.ABCI.VoteExtensionsEnableHeight = 1                             // Enable vote extensions from the start.
	resp.Validator.PubKeyTypes = []string{types.ABCIPubKeyTypeSecp256k1} // Only k1 keys.
	resp.Block.MaxBytes = -1                                             // Disable max block bytes, since we MUST include the whole EVM block, which is limited by max gas per block.

	return resp
}
