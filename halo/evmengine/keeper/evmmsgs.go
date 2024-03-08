package keeper

import (
	"bytes"
	"context"
	"slices"
	"sort"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

// evmLogs returns all EVM logs from the provided block hash.
func (k *Keeper) evmLogs(ctx context.Context, blockHash common.Hash) ([]*types.EVMLog, error) {
	var logs []*types.EVMLog
	for _, provider := range k.logProviders {
		ll, err := provider.Logs(ctx, blockHash)
		if err != nil {
			return nil, errors.Wrap(err, "prepare msgs")
		}

		logs = append(logs, ll...)
	}

	// Sort by Address > Topics > Data
	// This avoids dependency on runtime ordering.
	sort.Slice(logs, func(i, j int) bool {
		if cmp := bytes.Compare(logs[i].Address, logs[j].Address); cmp != 0 {
			return cmp < 0
		}

		topicI := slices.Concat(logs[i].Topics...)
		topicJ := slices.Concat(logs[j].Topics...)
		if cmp := bytes.Compare(topicI, topicJ); cmp != 0 {
			return cmp < 0
		}

		return bytes.Compare(logs[i].Data, logs[j].Data) < 0
	})

	return logs, nil
}
