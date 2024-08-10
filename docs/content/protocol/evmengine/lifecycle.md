---
sidebar_position: 4
---

# EVM Lifecycle

A transaction within the Omni EVM undergoes several key phases in its lifecycle, including proposal preparation, payload generation, and consensus-reaching, before being executed and permanently included in the blockchain state. This process not only ensures the integrity and security of transactions but also maintains the scalability and performance of the Omni network.

## Consensus

The following diagram illustrates the consensus process used by CometBFT in together with ABCI++ and EVM Engine.

<figure>
  <img src="/img/consensus.png" alt="Consensus in halo" />
  <figcaption>*The consensus process in `halo`*</figcaption>
</figure>

### Proposal Preparation and Payload Generation

The initial step in the EVM lifecycle involves the preparation of a proposal for the next block. This phase is crucial for organizing transactions into a coherent block structure that can be processed by the consensus mechanism.

In this stage, the `Keeper` module acts by handling unexpected transactions and leveraging the optimistic payload or creating a new one to ensure the smooth progression of the block preparation process.

Below is a stubbed fragment from source for the function declaration of [the preparation of proposals by the `Keeper`](https://github.com/omni-network/omni/blob/1439d8a99f66a3bb3b7d113c63f8f073512c5377/halo/evmengine/keeper/abci.go#L24-L98):

```go
// PrepareProposal function snippet
func (k *Keeper) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
    // Proposal preparation logic...
}
```

### Consensus and Finalization

Once a proposal is prepared, it enters the consensus mechanism, where validators work together to agree on the final state of the block. Upon reaching consensus, the block is finalized and state transitions are applied to the blockchain.

The processing of proposals is handled by the [base cosmos SDK app's `ProcessProposal`](https://github.com/cosmos/cosmos-sdk/blob/5e7aae0db1f5ee4cdcb1b0ff4d0003d09bfd047a/baseapp/abci.go#L471-L561).

Below is a stubbed function declaration [from source](https://github.com/omni-network/omni/blob/82ab876283a73300b091ebae2e9d14d8204a41f2/halo/app/prouter.go#L20-L65) of where this assignment takes place in the `halo` app:

```go
// Makes a new process proposal handler
func makeProcessProposalHandler(app *App) sdk.ProcessProposalHandler {
    // assigns the process proposal handling for baseapp cosmos SDK use
}
```

Finally, proposals are finalized by the cosmos SDK app's `FinalizeBlock`, [see source](https://github.com/cosmos/cosmos-sdk/blob/5e7aae0db1f5ee4cdcb1b0ff4d0003d09bfd047a/baseapp/abci.go#L874-L911).

This phase is critical for the execution of smart contracts and the application of transactions to the blockchain state, ensuring that all operations are accurately reflected in the network's ledger.
