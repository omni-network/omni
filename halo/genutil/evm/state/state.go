package state

import (
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/solc"

	"github.com/ethereum/go-ethereum/common"
)

type StorageValues map[string]any

type StorageSlot struct {
	Key   common.Hash
	Value common.Hash
}

// EncodeStorage encodes the given storage values according to the given storage layout.
// If it does not find a slot for a label, it returns an error.
// If it does not support encoding a value, it returns an error.
func EncodeStorage(layout solc.StorageLayout, values StorageValues) ([]StorageSlot, error) {
	var slots []StorageSlot
	for label, value := range values {
		slot, ok := solc.SlotOf(layout, label)
		if !ok {
			return nil, errors.New("label not found", "label", label)
		}

		s, err := encodeStorage(slot, value)
		if err != nil {
			return nil, err
		}

		slots = append(slots, s)
	}

	return slots, nil
}

func encodeStorage(slot uint, value any) (StorageSlot, error) {
	key := encodeSlot(slot)
	v, err := encodeValue(value)
	if err != nil {
		return StorageSlot{}, err
	}

	return StorageSlot{Key: key, Value: v}, nil
}

func encodeSlot(slot uint) common.Hash {
	s := new(big.Int).SetUint64(uint64(slot))

	return common.BigToHash(s)
}

func encodeValue(value any) (common.Hash, error) {
	switch v := value.(type) {
	case []byte:
		return common.BytesToHash(v), nil
	case *big.Int:
		return common.BigToHash(v), nil
	case common.Address:
		return common.HexToHash(v.Hex()), nil
	default:
		return common.Hash{}, errors.New("unsupported type")
	}
}
