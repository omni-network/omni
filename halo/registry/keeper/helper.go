package keeper

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

func (p *Portal) Verify() error {
	if p.GetChainId() == 0 {
		return errors.New("zero chain id")
	}

	if len(p.GetAddress()) != common.AddressLength {
		return errors.New("invalid address length")
	}

	if len(p.GetShardIds()) == 0 {
		return errors.New("no shards")
	}

	dupShards := make(map[uint64]bool)
	for _, s := range p.GetShardIds() {
		if dupShards[s] {
			return errors.New("duplicate shard id")
		}
		dupShards[s] = true
	}

	return nil
}
