package keeper

import (
	"bytes"
	"context"
	"slices"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"
)

// evmEvents returns all EVM log events from the provided block hash.
func (k *Keeper) evmEvents(ctx context.Context, blockHash common.Hash) ([]types.EVMEvent, error) {
	var events []types.EVMEvent
	for _, proc := range k.eventProcs {
		// Fetching evm events over the network is unreliable, retry forever.
		err := retryForever(ctx, func(ctx context.Context) (bool, error) {
			ll, err := proc.Prepare(ctx, blockHash)
			if err != nil {
				log.Warn(ctx, "Failed fetching evm events (will retry)", err, "proc", proc.Name())
				return false, nil // Retry
			}

			events = append(events, ll...)

			return true, nil // Done
		})
		if err != nil {
			return nil, err
		}
	}

	// Verify all events
	for _, event := range events {
		if err := event.Verify(); err != nil {
			return nil, errors.Wrap(err, "verify evm events")
		}
	}

	// Sort by Address > Topics > Data
	// This avoids dependency on runtime ordering.
	sort.Slice(events, func(i, j int) bool {
		if cmp := bytes.Compare(events[i].Address, events[j].Address); cmp != 0 {
			return cmp < 0
		}

		// TODO: replace this with sort.CompareFunc in next network upgrade which is more performant but has slightly different results
		topicI := slices.Concat(events[i].Topics...)
		topicJ := slices.Concat(events[j].Topics...)
		if cmp := bytes.Compare(topicI, topicJ); cmp != 0 {
			return cmp < 0
		}

		return bytes.Compare(events[i].Data, events[j].Data) < 0
	})

	return events, nil
}
