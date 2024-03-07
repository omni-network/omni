---
sidebar_position: 2
---

# `XBlock`

`XBlock` plays a crucial role in the Omni protocol by encapsulating cross-chain messages (`XMsg`) and ensuring their integrity and verifiability across the network. This section delves into the lifecycle of an `XBlock`, from its creation to its verification and execution on the destination chain.

## Overview of `XBlock` Creation and Flow

### 1. Validator Packaging into `XBlock`

Validators group multiple `XMsg` into an `XBlock`, which includes metadata about the source chain block and the messages it contains. The `XBlock` structure enables efficient verification through merkle-multi-proofs and simplifies the relaying and submission process.

### 2. Attestation by Validators

Omni protocol validators attest to `XBlock` hashes, ensuring over two-thirds consensus is reached. This attestation process is integral for validating the `XBlock` and its contents before relaying.

### 3. Relayer Role in `XBlock` to `Submission`

Relayers collect attested `XBlocks` and decide on the subset of `XMsg` to submit to each destination chain. They generate merkle-multi-proofs for these messages and submit them as part of an EVM transaction.

## Key Properties and Functions of `XBlock`

- **Efficient Verification**: `XBlock` allows for succinct verification of `XMsg` through merkle-multi-proofs, enabling selective message submission and cost management by relayers.
- **Consensus Optimization**: Only blocks containing `XMsg` require Omni Consensus attestations, reducing the workload on validators and optimizing the consensus process.
- **Deterministic Creation**: `XBlock` is deterministically generated from source chain blocks, ensuring a consistent and reproducible process across the network.

## Technical Details

### `XBlock` Structure

```go
type XBlock (
  uint         SourceChainID  // Source chain ID
  uint         BlockHeight    // Height of the source chain block
  bytes32      BlockHash      // Hash of the source chain block
  XMsg[]       Msgs           // XMsg events included in the source block
  XReceipt[]   Receipts       // XReceipt events resulting from `XMsg` execution
)
```

## Storage and Calculation

`XBlock` is not stored on-chain but is deterministically calculated from source blockchain data. This approach ensures that XBlock data is always available and verifiable without requiring additional storage on the Omni network.

## `XStream` and `XMsg` Ordering

`XMsg` are part of an `XStream`, a logical connection between source and destination chains, ensuring messages are uniquely identified and strictly ordered. This structure supports exactly-once delivery guarantees and strict ordering for cross-chain communication.

## Conclusion

`XBlock` is a foundational component of the Omni protocol, enabling secure, efficient, and verifiable cross-chain messaging. Through its structured lifecycle and the collaborative efforts of users, validators, and relayers, `XBlock` facilitates seamless interoperability across diverse blockchain networks.
