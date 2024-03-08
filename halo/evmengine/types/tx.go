package types

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

func (l *EVMLog) Verify() error {
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
