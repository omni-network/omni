package keeper

import (
	"bytes"
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

// sortIndex sorts the logs by index.
func sortIndex(all []ethtypes.Log) {
	sort.Slice(all, func(i, j int) bool {
		return all[i].Index < all[j].Index
	})
}

// sortLegacy sorts the logs by address, topics, and data.
func sortLegacy(all []ethtypes.Log) {
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

// evmEvents returns all EVM log events from the provided block hash in deterministic order.
// It uses the new index sorter.
func (k *Keeper) evmEvents(ctx context.Context, blockHash common.Hash) ([]types.EVMEvent, error) {
	return fetchProcEvents(ctx, k.engineCl, blockHash, sortIndex, k.eventProcs...)
}

// legacyEVMEvents returns all EVM log events from the provided block hash in deterministic order.
// It uses the legacy address, topics, and data sorter.
func (k *Keeper) legacyEVMEvents(ctx context.Context, blockHash common.Hash) ([]types.EVMEvent, error) {
	return fetchProcEvents(ctx, k.engineCl, blockHash, sortLegacy, k.eventProcs...)
}

func FetchProcEvents(ctx context.Context, cl ethclient.EngineClient, blockHash common.Hash, procs ...types.EvmEventProcessor) ([]types.EVMEvent, error) {
	return fetchProcEvents(ctx, cl, blockHash, sortIndex, procs...)
}

// fetchProcEvents fetches all EVM events for the given processors from the given block hash.
func fetchProcEvents(ctx context.Context, cl ethclient.EngineClient, blockHash common.Hash, sortFunc func([]ethtypes.Log), procs ...types.EvmEventProcessor) ([]types.EVMEvent, error) {
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

	sortFunc(all)

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
