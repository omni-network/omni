---
sidebar_position: 4
---

# Portal Receive Logic

## `xsubmit` Method

The `xsubmit` method is specifically designed for receiving calls from another chain. It is crucial for the secure and validated processing of cross-chain messages:

- Accepts calls only from valid signatures, primarily from the designated Relayer.
- Requires the destination contract `address` and `data`, which includes the `method` and method parameters (`data`).

## Submission Validation

To ensure the integrity and authenticity of incoming cross-chain messages, the following validation processes are performed:

### Aggregate Attestation Data Validation

Checks include:

- Verification that the `SourceChainID` is recognized and the `ValidatorSetID` is known, ensuring the availability of the corresponding validator set.
- For partial submissions, the `XBlockHash` must match the cursor's record.

### XMsgBatch Data Validation

This step involves:

- Confirming that the `TargetChainID` matches the local chain ID.
- For a new batch, verifying that `WrappedXMsg` indexes start at 0 or follow the last processed message index for ongoing submissions.
- If the cursor indicates a complete submission, the `ParentBlockHash` should match the cursor's record.

### Aggregate Attestation Signatures Verification

This process includes:

- Authenticating the validators' signatures on the `XBlockHash` to ensure a quorum is reached (over 66% of the set's validators).
- Verifying a Merkle-multi-proof against the `XBlockHash` that proves the included `XMsgBatch` hashes and all fields used in the above validation steps.

Through these meticulous validation and verification steps, the system guarantees the secure handling of cross-chain messages, facilitating reliable communication across different chains.
