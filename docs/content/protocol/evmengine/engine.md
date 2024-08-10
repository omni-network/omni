---
sidebar_position: 2
---

# Engine API

The Engine API represents a cornerstone in the Omni architecture, enabling seamless integration between the Ethereum Virtual Machine (EVM) and the CometBFT consensus engine. This page delves into the practical aspects of the Engine API as illustrated in the provided source files, highlighting its pivotal role in the Omni ecosystem.

## Overview

Utilizing the Ethereum Engine API, Omni introduces a novel architecture that harmonizes the EVM with CometBFT consensus, achieving scalability and efficiency. This integration is facilitated through the Engine API, which manages the interaction between Omni's execution layer and the consensus mechanism, ensuring that transaction processing and state transitions are handled effectively.

### Key Aspects of Engine API Integration

The Engine API enables the following functionalities within Omni's framework:

- **Transaction Processing:** Moves the transaction mempool to the execution layer, alleviating the CometBFT mempool and enhancing network throughput.
- **State Translation:** Employs an ABCI++ wrapper around the CometBFT engine for seamless state translation between the EVM and the consensus layer.

```go
// Example: Integration of the Engine API for state translation and transaction processing
func (c engineClient) NewPayloadV2(ctx context.Context, params engine.ExecutableData) (engine.PayloadStatusV1, error) {
    // Handling of new payloads and transaction processing...
}
```

This snippet from the `ethclient` package illustrates how the Engine API is utilized to create and manage new payloads, a critical step in preparing transaction blocks for consensus.

### Modular Framework

Inspired by Ethereum's Proof of Stake architecture, Omni adopts a modular approach that distinctly separates its execution and consensus layers. This separation not only mitigates performance bottlenecks but also provides a robust framework for handling transactions and state transitions.

```go
// Example: Modular framework leveraging the Engine API
func (k *Keeper) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
    // Proposal preparation logic demonstrating modular framework...
}
```

In this example, the `PrepareProposal` function (from the `keeper` package) showcases the modular integration of the Engine API, underscoring its significance in proposal preparation and consensus building.

## Benefits of the Engine API

- **Scalability and Efficiency:** By offloading the transaction mempool and facilitating efficient state translation, the Engine API contributes to Omni's scalability and sub-second transaction finality.
- **Flexibility:** Supports the interchangeability and upgrading of execution clients without system disruption, ensuring compatibility with various Ethereum execution clients.
