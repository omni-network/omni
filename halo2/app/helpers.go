package app

import (
	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/core/address"
)

var _ address.Codec = EthAddrCodec{}

// EthAddrCodec defines an address codec for Ethereum addresses.
type EthAddrCodec struct{}

func (EthAddrCodec) StringToBytes(text string) ([]byte, error) {
	return common.HexToAddress(text).Bytes(), nil
}

func (EthAddrCodec) BytesToString(bz []byte) (string, error) {
	return common.BytesToAddress(bz).String(), nil
}
