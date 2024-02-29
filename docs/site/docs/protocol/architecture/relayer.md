---
sidebar_position: 4
---

# Relayer

The Relayer is a permissionless actor that submits cross chain messages to destination chains and monitors the Omni Consensus Layer until ⅔ (>66%) of the validator set attested to the next block on each source chain. It then submits the applicable cross-chain messages to each destination chain providing the quorum validator signatures and a multi-merkle-proof. Will be incentivized in future versions of the network.

The Relayer is also responsible for monitoring attestations. Similar to validators, relayers should maintain an `XBlock` cache. I.e., track all source chain blocks, convert them to `XBlock`s, cache them, and make them available for internal indexed querying. Relayers monitor the Omni Consensus Chain state for attested `XBlock`.

For each destination chain, the relayer has to decide how many `XMsg`s to submit, which defines the [cost](./fees.md) of transactions being submitted to the destination chain. This is primarily defined by the data size and gas limit of the messages and the portal contract verification and processing overhead.

A merkle-multi-proof is generated for the set of identified `XMsg`s that match the quorum `XBlock` attestations root. The relayer submits a EVM transaction to the destination chain, ensuring it gets included on-chain as soon as possible. The transaction contains the following data:

```go
type Submission (
  AggregateAttestation att    // Aggregate attestation containing quorum signatures for a specific validator set.
  XMsg[]               msgs   // Subset of XMsgs in a XBlock (could also be all)
  bytes                Proofs // Merkle-multi-proof for XMsgBatch
)
```
