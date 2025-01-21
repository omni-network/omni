package types

import (
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

// customProtobufType defines the interface custom gogo proto types must implement
// in order to be used as a "customtype" extension.
//
// ref: https://github.com/cosmos/gogoproto/blob/master/custom_types.md
type customProtobufType interface {
	Marshal() ([]byte, error)
	MarshalTo(data []byte) (n int, err error)
	Unmarshal(data []byte) error
	Size() int

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

var (
	_ customProtobufType = (*Address)(nil)
	_ customProtobufType = (*Hash)(nil)
)

// Address extends common.Address to implement customProtobufType.
type Address common.Address

func (a Address) Marshal() ([]byte, error) {
	return common.Address(a).Bytes(), nil
}

func (a Address) MarshalTo(data []byte) (int, error) {
	copy(data, common.Address(a).Bytes())

	return common.AddressLength, nil
}

func (a *Address) Unmarshal(data []byte) error {
	if len(data) != common.AddressLength {
		return errors.New("invalid address length")
	}

	var b [common.AddressLength]byte
	copy(b[:], data)
	*a = b

	return nil
}

func (Address) Size() int {
	return common.AddressLength
}

func (a Address) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(common.Address(a))
	if err != nil {
		return nil, errors.Wrap(err, "marshal address")
	}

	return b, nil
}

func (a *Address) UnmarshalJSON(data []byte) error {
	if err := (*common.Address)(a).UnmarshalJSON(data); err != nil {
		return errors.Wrap(err, "unmarshal address")
	}

	return nil
}

// Hash extends common.Hash to implement customProtobufType.
type Hash common.Hash

func (h Hash) Marshal() ([]byte, error) {
	return common.Hash(h).Bytes(), nil
}

func (h Hash) MarshalTo(data []byte) (int, error) {
	copy(data, common.Hash(h).Bytes())

	return common.HashLength, nil
}

func (h *Hash) Unmarshal(data []byte) error {
	if len(data) != common.HashLength {
		return errors.New("invalid hash length")
	}

	var b [common.HashLength]byte
	copy(b[:], data)
	*h = b

	return nil
}

func (Hash) Size() int {
	return common.HashLength
}

func (h Hash) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(common.Hash(h))
	if err != nil {
		return nil, errors.Wrap(err, "marshal hash")
	}

	return b, nil
}

func (h *Hash) UnmarshalJSON(data []byte) error {
	if err := (*common.Hash)(h).UnmarshalJSON(data); err != nil {
		return errors.Wrap(err, "unmarshal hash")
	}

	return nil
}
