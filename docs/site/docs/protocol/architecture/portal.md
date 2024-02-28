---
sidebar_position: 5
---

# Portal Contract

Portal Contracts are a set of smart contracts that implement the on-chain logic of the Omni protocol and are deployed to all supported Rollup EVMs as well as the Omni EVM. They provide the main interface to call a cross-chain smart contract which results in a cross-chain message being emitted as a `XMsg`. They also provides a "pay at source" fee mechanism using the source chain’s native token. They track the omni consensus validator set, used to verify submitted cross chain message attestations ([more below](#submission-validation)).

## Relaying Methods

The Portal Contract exposes both an `xcall` method for cross-chain calls to another chain, and an `xsubmit` method for receiving calls from another chain.

- `xcall` is called by the `omni.xcall()` helper and requires the parameters listed above, with the Portal Contract emitting an `XMsg` Event for every call to it.
- `xsubmit` accepts only calls from valid signatures (currently the Relayer) and requires the destination contract `address` and `data` (containing `method` and method params `data`).

See more on performing cross chain transactions in the [developer section](../../develop/introduction.md).

## `Submission` Validation

Portal Contracts keep a cursor for each source chain that:

- Tracks the latest valid `Submission`’s `XBlockHash` that contained valid `XMsgBatch` to the local destination chain.
- The total messages in that batch.
- The index of the last message that was submitted.
- And implicitly, whether the latest `XBlock` is partially or completely submitted.

They validate the `AggregateAttestation` data:

- Ensuring the `SourceChainID` is known
- Ensuring the `ValidateSetID` is known and the validator set if available.
- If the cursor is partial, ensuring the `XBlockHash` matches that of the cursor.

Validate the `XMsgBatch` data:

- Ensuring the `TargetChainID` matches the local chain ID.
- If the cursor is complete, ensure the `ParentBlockHash` matches that of the cursor.
- Ensure the included `WrappedXMsg` indexes start from 0 if the cursor is complete or follow on the cursor index if partial.

Verify the `AggregateAttestation` signatures:

- Verify all validator signatures over the root `XBlockHash`
- Ensure that quorum is reached; more than 66% validators in the set signed.
- Verify a merke-multi-proof against the `XBlockHash` that proves the following fields of the `XBlock`:
  - All fields used in above validator.
  - All included `XMsgBatch` hashes.
