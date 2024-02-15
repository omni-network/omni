package app

import (
	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/core/address"
)

var _ address.Codec = ethAddrCodec{}

// ethAddrCodec defines an address codec for Ethereum addresses.
type ethAddrCodec struct{}

func (ethAddrCodec) StringToBytes(text string) ([]byte, error) {
	return common.HexToAddress(text).Bytes(), nil
}

func (ethAddrCodec) BytesToString(bz []byte) (string, error) {
	return common.BytesToAddress(bz).String(), nil
}
