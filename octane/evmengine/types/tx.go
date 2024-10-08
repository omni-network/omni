package types

import (
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// ToEthLog converts an EVMEvent to an Ethereum Log.
// Note it assumes that Verify has been called before.
func (l *EVMEvent) ToEthLog() (ethtypes.Log, error) {
	if l == nil {
		return ethtypes.Log{}, errors.New("nil log")
	} else if len(l.Topics) == 0 {
		return ethtypes.Log{}, errors.New("empty topics")
	}

	topics := make([]common.Hash, 0, len(l.Topics))
	for _, t := range l.Topics {
		topics = append(topics, common.BytesToHash(t))
	}

	addr, err := cast.EthAddress(l.Address)
	if err != nil {
		return ethtypes.Log{}, err
	}

	return ethtypes.Log{
		Address: addr,
		Topics:  topics,
		Data:    l.Data,
	}, nil
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
