package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var _ types.EvmEventProcessor = pocEventProvider{}

//nolint:gochecknoglobals // Will remove with PoC
var zeroAddr common.Address

func NewPocEventProvider(k *Keeper) types.EvmEventProcessor {
	return pocEventProvider{k: k}
}

// pocEventProvider is a temporary PoC that queries previous block EVM logs and includes
// the total as cosmosSDK msg.
type pocEventProvider struct {
	k *Keeper
}

func (p pocEventProvider) Prepare(ctx context.Context, blockHash common.Hash) ([]*types.EVMEvent, error) {
	logs, err := p.k.engineCl.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
	})
	if err != nil {
		return nil, errors.Wrap(err, "filter logs")
	}

	evmEvents := make([]*types.EVMEvent, 0, len(logs))
	for _, l := range logs {
		evmEvents = append(evmEvents, &types.EVMEvent{
			Address: zeroAddr.Bytes(), // Stub addresses
			Topics:  topicBytes(l.Topics),
			Data:    l.Data,
		})
	}

	return evmEvents, nil
}

func (pocEventProvider) Addresses() []common.Address {
	return []common.Address{zeroAddr}
}

func (pocEventProvider) Deliver(context.Context, common.Hash, *types.EVMEvent) error {
	return nil
}

func topicBytes(topics []common.Hash) [][]byte {
	t := make([][]byte, len(topics))
	for i, h := range topics {
		t[i] = h[:]
	}

	return t
}
