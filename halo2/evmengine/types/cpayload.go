package types

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CPayloadProvider abstracts modules that provide consensus payload messages.
// These include msgs from vote extensions and/or any msgs from EVM contracts.
//
// EVMEngine calls this during PreparePayload to collect all msgs to include in
// the single consensus transaction.
type CPayloadProvider interface {
	PreparePayload(ctx context.Context, height uint64, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error)
}
