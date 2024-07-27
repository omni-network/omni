package xchain

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	typUint8   = "uint8"
	typUint64  = "uint64"
	typBytes32 = "bytes32"
	typBytes   = "bytes"
	typAddress = "address"
)

// submissionHeader defines the header leaf of the attestation merkle tree.
// It contains fields from the AttestHeader and BlockHeader.
type submissionHeader struct {
	SourceChainID    uint64
	ConsensusChainID uint64
	ConfLevel        ConfLevel
	AttestOffset     uint64
	BlockHeight      uint64
	BlockHash        common.Hash
}

//nolint:gochecknoglobals // Static ABI types
var (
	submissionHeaderABI = mustABITuple([]abi.ArgumentMarshaling{
		{Name: "SourceChainID", Type: typUint64},
		{Name: "ConsensusChainID", Type: typUint64},
		{Name: "ConfLevel", Type: typUint8},
		{Name: "AttestOffset", Type: typUint64},
		{Name: "BlockHeight", Type: typUint64},
		{Name: "BlockHash", Type: typBytes32},
	})

	msgABI = mustABITuple([]abi.ArgumentMarshaling{
		{Name: "DestChainID", Type: typUint64},
		{Name: "ShardID", Type: typUint64},
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

// encodeSubmissionHeader ABI encodes a attest header and block header into a byte slice.
func encodeSubmissionHeader(attHeader AttestHeader, blockHeader BlockHeader) ([]byte, error) {
	if attHeader.ChainVersion.ID != blockHeader.ChainID {
		return nil, errors.New("chain ID mismatch")
	}

	resp, err := submissionHeaderABI.Pack(submissionHeader{
		SourceChainID:    attHeader.ChainVersion.ID,
		ConsensusChainID: attHeader.ConsensusChainID,
		ConfLevel:        attHeader.ChainVersion.ConfLevel,
		AttestOffset:     attHeader.AttestOffset,
		BlockHeight:      blockHeader.BlockHeight,
		BlockHash:        blockHeader.BlockHash,
	})
	if err != nil {
		return nil, errors.Wrap(err, "pack xchain submissionHeader")
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
