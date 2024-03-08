package types

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// VoteExtensionProvider abstracts logic that provides consensus payload messages
// from the last commits vote extensions.
//
// EVMEngine calls this during PreparePayload to collect all vote extensions msgs to include in
// the consensus block.
type VoteExtensionProvider interface {
	PrepareVotes(ctx context.Context, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error)
}

// EvmLogProvider abstracts logic that provide EVM logs of the
// previous execution payload (current head) identified by
// the provided block hash.
//
// EVMEngine calls this during PreparePayload to collect all EVM-logs to include in
// the consensus block. It is also called during ProcessPayload to verify the proposed EVM logs.
type EvmLogProvider interface {
	Logs(ctx context.Context, blockHash common.Hash) ([]*EVMLog, error)
	Addresses() []common.Address
	DeliverLog(ctx context.Context, blockHash common.Hash, log *EVMLog) error
}
