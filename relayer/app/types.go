package relayer

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"sort"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"
)

type StreamUpdate struct {
	xchain.StreamID
	MsgTree     xchain.MsgTree
	Attestation xchain.Attestation // Attestation for the xmsgs
	Msgs        []xchain.Msg       // msgs that increment the cursor
}

// CreateFunc is a function that creates one or more submissions from the given stream update.
type CreateFunc func(streamUpdate StreamUpdate) ([]xchain.Submission, error)

// SendFunc sends a submission to the destination chain by invoking "xsubmit" on portal contract.
type SendFunc func(ctx context.Context, submission xchain.Submission) error

// submissionToBinding converts a go xchain submission to a solidity binding submission.
func submissionToBinding(sub xchain.Submission) bindings.XSubmission {
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

// randomHex7 returns a random 7-character hex string.
func randomHex7() string {
	bytes := make([]byte, 4)
	_, _ = rand.Read(bytes)
	hexString := hex.EncodeToString(bytes)

	// Trim the string to 7 characters
	if len(hexString) > 7 {
		hexString = hexString[:7]
	}

	return hexString
}
