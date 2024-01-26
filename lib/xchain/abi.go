package xchain

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	typUint64  = "uint64"
	typBytes32 = "bytes32"
	typBytes   = "bytes"
	typAddress = "address"
)

//nolint:gochecknoglobals // Static ABI types
var (
	headerABI = mustABITuple([]abi.ArgumentMarshaling{
		{Name: "SourceChainID", Type: typUint64},
		{Name: "BlockHeight", Type: typUint64},
		{Name: "BlockHash", Type: typBytes32},
	})

	msgABI = mustABITuple([]abi.ArgumentMarshaling{
		{Name: "SourceChainID", Type: typUint64},
		{Name: "DestChainID", Type: typUint64},
		{Name: "StreamOffset", Type: typUint64},
		{Name: "SourceMsgSender", Type: typAddress},
		{Name: "DestAddress", Type: typAddress},
		{Name: "Data", Type: typBytes},
		{Name: "DestGasLimit", Type: typUint64},
	})
)

// encodeMsg ABI encodes a cross chain message into a byte slice.
func encodeMsg(msg Msg) ([]byte, error) {
	resp, err := msgABI.Pack(msg)
	if err != nil {
		return nil, errors.Wrap(err, "pack xchain msg")
	}

	return resp, nil
}

// encodeHeader ABI encodes a cross chain block header into a byte slice.
func encodeHeader(header BlockHeader) ([]byte, error) {
	resp, err := headerABI.Pack(header)
	if err != nil {
		return nil, errors.Wrap(err, "pack xchain header")
	}

	return resp, nil
}

// mustABITuple returns an ABI tuple typ with the provided components.
// It panics on error.
func mustABITuple(components []abi.ArgumentMarshaling) abi.Arguments {
	typ, err := abi.NewType("tuple", "", components)
	if err != nil {
		panic(err)
	}

	return abi.Arguments{abi.Argument{Type: typ}}
}
