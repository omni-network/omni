package types

import "github.com/ethereum/go-ethereum/common"

// NewGenesisState creates a new GenesisState instance.
func NewGenesisState(executionBlockHash common.Hash) *GenesisState {
	return &GenesisState{
		ExecutionBlockHash: executionBlockHash.Bytes(),
	}
}
