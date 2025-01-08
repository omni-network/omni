package keeper

import (
	"bytes"
	"context"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/feature"
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

	if feature.FlagSimpleEVMEvents.Enabled(ctx) {
		// Sort by Index
		sort.Slice(all, func(i, j int) bool {
			return all[i].Index < all[j].Index
		})
	} else {
		// Sort by Address > Topics > Data
		// This avoids dependency on runtime ordering.
		sort.Slice(all, func(i, j int) bool {
			if cmp := bytes.Compare(all[i].Address.Bytes(), all[j].Address.Bytes()); cmp != 0 {
				return cmp < 0
			}

			topicI := concatHashes(all[i].Topics)
			topicJ := concatHashes(all[j].Topics)
			if cmp := bytes.Compare(topicI, topicJ); cmp != 0 {
				return cmp < 0
			}

			return bytes.Compare(all[i].Data, all[j].Data) < 0
		})
	}

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

func concatHashes(hashes []common.Hash) []byte {
	resp := make([]byte, 0, len(hashes)*common.HashLength)
	for _, hash := range hashes {
		resp = append(resp, hash.Bytes()...)
	}

	return resp
}
