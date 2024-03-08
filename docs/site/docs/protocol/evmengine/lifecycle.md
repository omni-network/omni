---
sidebar_position: 4
---

# EVM Lifecycle

A transaction within the Omni EVM undergoes several key phases in its lifecycle, including proposal preparation, payload generation, and consensus-reaching, before being executed and permanently included in the blockchain state. This process not only ensures the integrity and security of transactions but also maintains the scalability and performance of the Omni network.

### Proposal Preparation and Payload Generation

The initial step in the EVM lifecycle involves the preparation of a proposal for the next block. This phase is crucial for organizing transactions into a coherent block structure that can be processed by the consensus mechanism.

```go
// PrepareProposal function snippet
func (k *Keeper) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
    // Proposal preparation logic...
}
```

In this stage, the Keeper module plays a pivotal role, handling unexpected transactions and leveraging the optimistic payload or creating a new one to ensure the smooth progression of the block preparation process.

### Consensus and Finalization

Once a proposal is prepared, it enters the consensus mechanism, where validators work together to agree on the final state of the block. Upon reaching consensus, the block is finalized and state transitions are applied to the blockchain.

```go
// ExecutionPayload function snippet
func (s msgServer) ExecutionPayload(ctx context.Context, msg *types.MsgExecutionPayload) (*types.ExecutionPayloadResponse, error) {
    // Execution payload handling logic...
}
```

This phase is critical for the execution of smart contracts and the application of transactions to the blockchain state, ensuring that all operations are accurately reflected in the network's ledger.
