package types

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// ToEthLog converts an EVMEvent to an Ethereum Log.
// Note it assumes that Verify has been called before.
func (l *EVMEvent) ToEthLog() ethtypes.Log {
	topics := make([]common.Hash, 0, len(l.Topics))
	for _, t := range l.Topics {
		topics = append(topics, common.BytesToHash(t))
	}

	return ethtypes.Log{
		Address: common.BytesToAddress(l.Address),
		Topics:  topics,
		Data:    l.Data,
	}
}

func (l *EVMEvent) Verify() error {
	if l == nil {
		return errors.New("nil log")
	}

	if l.Address == nil {
		return errors.New("nil address")
	}

	if len(l.Topics) == 0 {
		return errors.New("empty topics")
	}

	if len(l.Address) != len(common.Address{}) {
		return errors.New("invalid address length")
	}

	for _, t := range l.Topics {
		if len(t) != len(common.Hash{}) {
			return errors.New("invalid topic length")
		}
	}

	return nil
}
