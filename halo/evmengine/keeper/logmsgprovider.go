package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var _ types.EvmLogProvider = pocLogProvider{}

//nolint:gochecknoglobals // Will remove with PoC
var zeroAddr common.Address

func NewPocLogProvider(k *Keeper) types.EvmLogProvider {
	return pocLogProvider{k: k}
}

// pocLogProvider is a temporary PoC that queries previous block EVM logs and includes
// the total as cosmosSDK msg.
type pocLogProvider struct {
	k *Keeper
}

func (p pocLogProvider) Logs(ctx context.Context, blockHash common.Hash) ([]*types.EVMLog, error) {
	logs, err := p.k.engineCl.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
	})
	if err != nil {
		return nil, errors.Wrap(err, "filter logs")
	}

	evmLogs := make([]*types.EVMLog, len(logs))
	for i, l := range logs {
		evmLogs[i] = &types.EVMLog{
			Address: zeroAddr.Bytes(), // Stub addresses
			Topics:  topicBytes(l.Topics),
			Data:    l.Data,
		}
	}

	return evmLogs, nil
}

func (pocLogProvider) Addresses() []common.Address {
	return []common.Address{zeroAddr}
}

func (pocLogProvider) DeliverLog(context.Context, common.Hash, *types.EVMLog) error {
	return nil
}

func topicBytes(topics []common.Hash) [][]byte {
	t := make([][]byte, len(topics))
	for i, h := range topics {
		t[i] = h[:]
	}

	return t
}
