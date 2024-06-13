package keeper

import (
	"bytes"
	"context"
	"slices"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"
)

// evmEvents returns all EVM log events from the provided block hash.
func (k *Keeper) evmEvents(ctx context.Context, blockHash common.Hash) ([]*types.EVMEvent, error) {
	var events []*types.EVMEvent
	for _, proc := range k.eventProcs {
		ll, err := proc.Prepare(ctx, blockHash)
		if err != nil {
			return nil, errors.Wrap(err, "prepare msgs")
		}
		for _, log := range ll {
			if err := log.Verify(); err != nil {
				return nil, errors.Wrap(err, "verify log")
			}
		}

		events = append(events, ll...)
	}

	// Sort by Address > Topics > Data
	// This avoids dependency on runtime ordering.
	sort.Slice(events, func(i, j int) bool {
		if cmp := bytes.Compare(events[i].Address, events[j].Address); cmp != 0 {
			return cmp < 0
		}

		topicI := slices.Concat(events[i].Topics...)
		topicJ := slices.Concat(events[j].Topics...)
		if cmp := bytes.Compare(topicI, topicJ); cmp != 0 {
			return cmp < 0
		}

		return bytes.Compare(events[i].Data, events[j].Data) < 0
	})

	return events, nil
}
