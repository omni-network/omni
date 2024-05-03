---
sidebar_position: 3
---

# Relayer

The Relayer plays a pivotal role in the Omni protocol as a permissionless entity that bridges cross-chain messages between source and destination chains. It performs critical functions that ensure the smooth and secure transmission of messages across the network.

## Responsibilities

### Attestation Monitoring and XBlock Cache Management

Like validators, Relayers are tasked with monitoring attestations within the Omni Consensus Chain. This involves:

- Maintaining an `XBlock` cache by tracking all source chain blocks, converting them into `XBlock` format, caching these blocks, and enabling internal indexed querying. This ensures the Relayer has timely access to relevant data for cross-chain message processing.
- Keeping a vigilant eye on the Omni Consensus Chain state for attested `XBlocks`, preparing for the submission of cross-chain messages once a consensus is reached.

### Decision Making for Message Submission

A key decision that Relayers face is determining the number of `XMsg`s to submit to each destination chain. This decision directly influences the cost of transactions due to factors like data size, gas limits, and the computational overhead required for portal contract verification and message processing.

### Cross-Network Message Submission

The Relayer waits for the Omni Consensus Layer to confirm that over two-thirds (>66%) of the validator set have attested to the next block for each source chain. Once this quorum is achieved, the Relayer is responsible for submitting the validated cross-chain messages to their respective destination chains, accompanied by the necessary validator signatures and a multi-merkle-proof. Future iterations of the network plan to incentivize this crucial function.

For the actual submission to a destination chain, Relayers generate a merkle-multi-proof for the `XMsg`s that are to be included, based on the `XBlock` attestations root that has reached a quorum. They then craft an EVM transaction containing this data, aiming to ensure its swift inclusion on the destination chain. The transaction structure is as follows:

```go
// Submission is a cross-chain submission of a set of messages and their proofs.
type Submission struct {
	AttestationRoot common.Hash // Attestation merkle root of the cross-chain Block
	ValidatorSetID  uint64      // Validator set that approved the attestation.
	BlockHeader     BlockHeader // BlockHeader identifies the cross-chain Block
	Msgs            []Msg       // Messages to be submitted
	Proof           [][32]byte  // Merkle multi proofs of the messages
	ProofFlags      []bool      // Flags indicating whether the proof is a left or right proof
	Signatures      []SigTuple  // Validator signatures and public keys
	DestChainID     uint64      // Destination chain ID, for internal use only
}
```

This transaction is a critical step in the relaying process, encapsulating the essence of the Relayer's role in maintaining the integrity and efficiency of cross-chain communication within the Omni protocol.
