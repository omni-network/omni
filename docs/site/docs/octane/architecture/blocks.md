---
sidebar_position: 3
---

# Block Building Flow

Putting it all together, the combined flow between the **EngineAPI** and **ABCI 2.0** does the following:

## EVM Client

- Receives EVM transactions from users and maintains the EVM mempool
- Builds EVM blocks from transactions in the mempool
- Sends EVM block payloads to the consensus client via the EngineAPI
- Stores EVM state

## Consensus Module

- Receives EVM payloads (blocks) and commits them to consensus module state as a single transaction
- Stores strictly consensus module state (it does not store any EVM state besides the block metadata, which is passed via the EngineAPI)
- Can read logs from the EVM via the EngineAPI, like the Beacon Chain Deposit Contract.

Here is a visual overview of the flow:

<figure>
  <img src="/img/consensus.png" alt="Consensus in octane" />
  <figcaption>*The consensus process in `octane`*</figcaption>
</figure>
