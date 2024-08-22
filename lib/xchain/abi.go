package xchain

import (
	"bytes"
	"sort"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	typUint8   = "uint8"
	typUint64  = "uint64"
	typBytes32 = "bytes32"
	typBytes   = "bytes"
	typAddress = "address"
	typTuple   = "tuple"
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
	omniPortalABI = mustGetABI(bindings.OmniPortalMetaData)

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

// EncodeXSubmit returns the abi encoding of the xsubmit function call.
func EncodeXSubmit(sub bindings.XSubmission) ([]byte, error) {
	bytes, err := omniPortalABI.Pack("xsubmit", sub)
	if err != nil {
		return nil, errors.Wrap(err, "pack xsubmit")
	}

	return bytes, nil
}

// DecodeXSubmit decodes the xsubmit function call data.
func DecodeXSubmit(txCallData []byte) (bindings.XSubmission, error) {
	const method = "xsubmit"
	m, ok := omniPortalABI.Methods[method]
	if !ok {
		return bindings.XSubmission{}, errors.New("missing method")
	}

	trimmed := bytes.TrimPrefix(txCallData, m.ID)
	if bytes.Equal(trimmed, txCallData) {
		return bindings.XSubmission{}, errors.New("tx data not prefixed with xsubmit method ID")
	}

	unpacked, err := m.Inputs.Unpack(trimmed)
	if err != nil {
		return bindings.XSubmission{}, errors.Wrap(err, "unpack submission")
	}

	wrap := struct {
		Sub bindings.XSubmission
	}{}
	if err := m.Inputs.Copy(&wrap, unpacked); err != nil {
		return bindings.XSubmission{}, errors.Wrap(err, "copy submission")
	}

	return wrap.Sub, nil
}

func SubmissionFromBinding(sub bindings.XSubmission, destChainID uint64) Submission {
	sigs := make([]SigTuple, 0, len(sub.Signatures))
	for _, sig := range sub.Signatures {
		sigs = append(sigs, SigTuple{
			ValidatorAddress: sig.ValidatorAddr,
			Signature:        Signature65(sig.Signature),
		})
	}

	msgs := make([]Msg, 0, len(sub.Msgs))
	for _, msg := range sub.Msgs {
		msgs = append(msgs, Msg{
			MsgID: MsgID{
				StreamID: StreamID{
					DestChainID: msg.DestChainId,
					ShardID:     ShardID(msg.ShardId),
				},
				StreamOffset: msg.Offset,
			},
			SourceMsgSender: msg.Sender,
			DestAddress:     msg.To,
			Data:            msg.Data,
			DestGasLimit:    msg.GasLimit,
		})
	}

	return Submission{
		AttestationRoot: sub.AttestationRoot,
		ValidatorSetID:  sub.ValidatorSetId,
		AttHeader: AttestHeader{
			ConsensusChainID: sub.BlockHeader.ConsensusChainId,
			ChainVersion:     NewChainVersion(sub.BlockHeader.SourceChainId, ConfLevel(sub.BlockHeader.ConfLevel)),
			AttestOffset:     sub.BlockHeader.Offset,
		},
		BlockHeader: BlockHeader{
			ChainID:   sub.BlockHeader.SourceChainId,
			BlockHash: sub.BlockHeader.SourceBlockHash,
		},
		Proof:       sub.Proof,
		ProofFlags:  sub.ProofFlags,
		DestChainID: destChainID,
		Signatures:  sigs,
		Msgs:        msgs,
	}
}

// SubmissionToBinding converts a go xchain submission to a solidity binding submission.
func SubmissionToBinding(sub Submission) bindings.XSubmission {
	// Sort the signatures by validator address to ensure deterministic ordering.
	sort.Slice(sub.Signatures, func(i, j int) bool {
		return sub.Signatures[i].ValidatorAddress.Cmp(sub.Signatures[j].ValidatorAddress) < 0
	})

	sigs := make([]bindings.ValidatorSigTuple, 0, len(sub.Signatures))
	for _, sig := range sub.Signatures {
		sigs = append(sigs, bindings.ValidatorSigTuple{
			ValidatorAddr: sig.ValidatorAddress,
			Signature:     sig.Signature[:],
		})
	}

	msgs := make([]bindings.XMsg, 0, len(sub.Msgs))
	for _, msg := range sub.Msgs {
		msgs = append(msgs, bindings.XMsg{
			DestChainId: msg.DestChainID,
			ShardId:     uint64(msg.ShardID),
			Offset:      msg.StreamOffset,
			Sender:      msg.SourceMsgSender,
			To:          msg.DestAddress,
			Data:        msg.Data,
			GasLimit:    msg.DestGasLimit,
		})
	}

	return bindings.XSubmission{
		AttestationRoot: sub.AttestationRoot,
		ValidatorSetId:  sub.ValidatorSetID,
		BlockHeader: bindings.XBlockHeader{
			SourceChainId:     sub.BlockHeader.ChainID,
			ConsensusChainId:  sub.AttHeader.ConsensusChainID,
			SourceBlockHash:   sub.BlockHeader.BlockHash,
			SourceBlockHeight: sub.BlockHeader.BlockHeight,
			ConfLevel:         uint8(sub.AttHeader.ChainVersion.ConfLevel),
			Offset:            sub.AttHeader.AttestOffset,
		},
		Proof:      sub.Proof,
		ProofFlags: sub.ProofFlags,
		Signatures: sigs,
		Msgs:       msgs,
	}
}

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

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
	typ, err := abi.NewType(typTuple, "", components)
	if err != nil {
		panic(err)
	}

	return abi.Arguments{abi.Argument{Type: typ}}
}
