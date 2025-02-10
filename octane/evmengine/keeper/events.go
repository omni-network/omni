package keeper

import (
	"context"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// evmEvents returns all EVM log events from the provided block hash in deterministic order.
func (k *Keeper) evmEvents(ctx context.Context, blockHash common.Hash) ([]types.EVMEvent, error) {
	return FetchProcEvents(ctx, k.engineCl, blockHash, k.eventProcs...)
}

// FetchProcEvents fetches all EVM events for the given processors from the given block hash.
func FetchProcEvents(ctx context.Context, cl ethclient.EngineClient, blockHash common.Hash, procs ...types.EvmEventProcessor) ([]types.EVMEvent, error) {
	var all []ethtypes.Log
	for _, proc := range procs {
		// Fetching evm events over the network is unreliable, retry forever.
		err := retryForever(ctx, func(ctx context.Context) (bool, error) {
			addresses, topics := proc.FilterParams()
			logs, err := cl.FilterLogs(ctx, ethereum.FilterQuery{
				BlockHash: &blockHash,
				Addresses: addresses,
				Topics:    topics,
			})
			if err != nil {
				log.Warn(ctx, "Failed fetching evm events (will retry)", err, "proc", proc.Name())
				return false, nil // Retry
			}

			all = append(all, logs...)

			return true, nil // Done
		})
		if err != nil {
			return nil, err
		}
	}

	// Sort by Index
	sort.Slice(all, func(i, j int) bool {
		return all[i].Index < all[j].Index
	})

	// Verify and convert logs
	var resp []types.EVMEvent
	for _, l := range all {
		topics := make([][]byte, 0, len(l.Topics))
		for _, t := range l.Topics {
			topics = append(topics, t.Bytes())
		}

		event := types.EVMEvent{
			Address: l.Address.Bytes(),
			Topics:  topics,
			Data:    l.Data,
		}

		if err := event.Verify(); err != nil {
			return nil, errors.Wrap(err, "verify evm events")
		}

		resp = append(resp, event)
	}

	return resp, nil
}
