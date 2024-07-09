package relayer

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

const (
	subGasBase               uint64 = 500_000
	subGasXmsgOverhead       uint64 = 100_000
	subGasMax                uint64 = 10_000_000 // Many chains have block gas limit of 30M, so we limit ourselves to 1/3 of that.
	subEphemeralConsensusGas uint64 = 5_000_000
	properGasEstimation      uint64 = 0 // Use proper (RPC) gas estimation
)

// gasEstimator is a function that estimates the gas usage by submitting the set of messages.
// Note that the messages MUST be from the same source chain.
// It returns zero if proper (RPC) gas estimation should be used.
type gasEstimator func(destChain uint64, msgs []xchain.Msg) uint64

// newGasEstimator returns a new gas estimator function.
func newGasEstimator(network netconf.ID) gasEstimator {
	consChainID := network.Static().OmniConsensusChainIDUint64()

	// Some destination chains do not need this model and can simply rely on the proper gas estimation.
	skipModel := map[uint64]bool{
		evmchain.IDArbSepolia: true, // Arbitrum has non-standard gas usage, and super-fast blocks, so we skip the model.
	}

	return func(destChain uint64, msgs []xchain.Msg) uint64 {
		if len(msgs) == 0 {
			return subGasBase
		}

		srcChain := msgs[0].SourceChainID

		if skipModel[destChain] { // Note we must provide destination chain explicitly, since msg.DestChainID can be broadcast=0.
			return properGasEstimation
		}

		if srcChain == consChainID {
			// Consensus chain xmsgs do not have a gas limit, so naiveSubmissionGas doesn't work.

			if network.IsEphemeral() {
				// For ephemeral chains, we use a fixed very high value.
				return subEphemeralConsensusGas
			}

			// For protected chains, we rely on proper gas estimation.
			// Proper gas estimation for protected chains is ok since real world consensus chain messages are rare,
			// so the multiple-submissions-per-block-gas-estimation-wrong-offset issue isn't a problem.

			return properGasEstimation
		}

		return naiveSubmissionGas(msgs)
	}
}

// naiveSubmissionGas returns the estimated max gas usage of a submissions using a naive model:
// - <gasBase> + sum(xmsg.DestGasLimit + <gasXmsgOverhead>).
func naiveSubmissionGas(msgs []xchain.Msg) uint64 {
	resp := subGasBase
	for _, msg := range msgs {
		resp += msg.DestGasLimit + subGasXmsgOverhead
	}

	return resp
}
