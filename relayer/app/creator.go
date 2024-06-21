package relayer

import (
	"slices"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

const (
	subGasBase               uint64 = 500_000
	subGasXmsgOverhead       uint64 = 100_000
	subGasMax                uint64 = 10_000_000 // Many chains have block gas limit of 30M, so we limit ourselves to 1/3 of that.
	subEphemeralConsensusGas uint64 = 5_000_000
)

// estimateGas returns the estimated max gas usage of a submissions using a naive model:
// - <gasBase> + sum(xmsg.DestGasLimit + <gasXmsgOverhead>).
func estimateGas(msgs []xchain.Msg) uint64 {
	resp := subGasBase
	for _, msg := range msgs {
		resp += msg.DestGasLimit + subGasXmsgOverhead
	}

	return resp
}

// consensusGasLimit returns the gas limit for consensus chain submissions.
// Consensus chain xmsgs do not have a gas limit, so the above estimateGas doesn't work.
//
// For ephemeral chains, we use a fixed very high value.
// For protected chains, we rely on proper gas estimation.
//
// Proper gas estimation for protected chains is ok since real world consensus chain messages are rare,
// so the multiple-submissions-per-block-gas-estimation-wrong-offset issue isn't a problem.
func consensusGasLimit(network netconf.ID) uint64 {
	if network.IsEphemeral() {
		return subEphemeralConsensusGas
	}

	return 0 // we return 0, which is the signal to use proper gas estimation
}

// CreateSubmissions splits the update into multiple submissions that are each small enough (wrt calldata and gas)
// to be submitted on-chain.
func CreateSubmissions(up StreamUpdate) ([]xchain.Submission, error) {
	// Sanity check on input, should only be for a single stream.
	for i, msg := range up.Msgs {
		if msg.SourceChainID != up.SourceChainID {
			return nil, errors.New("invalid msgs [BUG]")
		} else if i > 0 && msg.StreamOffset != up.Msgs[i-1].StreamOffset+1 {
			return nil, errors.New("msgs not sequential [BUG]")
		}
	}

	att := up.Attestation

	attRoot, err := xchain.AttestationRoot(att.BlockHeader, att.MsgRoot)
	if err != nil {
		return nil, err
	}

	var resp []xchain.Submission //nolint:prealloc // Cannot predetermine size
	for _, msgs := range groupMsgsByCost(up.Msgs) {
		multi, err := up.MsgTree.Proof(msgs)
		if err != nil {
			return nil, err
		}

		resp = append(resp, xchain.Submission{
			AttestationRoot: attRoot,
			ValidatorSetID:  att.ValidatorSetID,
			BlockHeader:     att.BlockHeader,
			Msgs:            msgs,
			Proof:           multi.Proof,
			ProofFlags:      multi.ProofFlags,
			Signatures:      att.Signatures,
			DestChainID:     up.DestChainID,
		})
	}

	return resp, nil
}

// groupMsgsByCost split the messages into groups that are each small enough (wrt calldata and gas)
// to be submitted on-chain.
func groupMsgsByCost(msgs []xchain.Msg) [][]xchain.Msg {
	var resp [][]xchain.Msg

	var current []xchain.Msg
	for _, msg := range msgs {
		// If adding the msg to the current batch will cross max limit, then
		// complete the batch by adding it to the response and starting a new empty current batch.
		if estimateGas(append(slices.Clone(current), msg)) > subGasMax {
			resp = append(resp, current)
			current = nil
		}

		// Add the message to the current batch
		current = append(current, msg)
	}

	// Add current batch to response and return
	return append(resp, current)
}
