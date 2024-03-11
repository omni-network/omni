---
sidebar_position: 3
---

# ABCI++

ABCI++ plays a pivotal role in the Omni blockchain, enabling a novel architecture that combines the Ethereum Virtual Machine (EVM) with the CometBFT consensus mechanism in a scalable manner. This integration is made possible through the use of Ethereum’s Engine API alongside ABCI++, drawing from Ethereum’s Proof of Stake (PoS) architecture to create a modular framework that distinctly separates Omni's execution and consensus layers.

## Overview

Omni's innovative use of ABCI++ facilitates a clear separation between the execution and consensus layers of the blockchain, addressing the bottlenecks that have hindered performance in previous blockchain designs. By leveraging ABCI++, Omni achieves a truly scalable solution, capable of handling increased network activity without compromising consensus speed or efficiency.

### Handling Transaction Requests

Previous approaches relied on the CometBFT mempool to manage transaction requests, leading to network congestion and compromised consensus speed as activity increased. Omni addresses this challenge by utilizing the Engine API, alongside ABCI++, to move the transaction mempool to the execution layer. This strategic move ensures the CometBFT consensus process remains lightweight and efficient.

```go
// Example: Using ABCI++ for processing proposals and managing consensus payloads
func makeProcessProposalHandler(app *App) sdk.ProcessProposalHandler {
    // ProcessProposalHandler implementation...
}
```

ABCI++ is integrated within the Omni framework to process proposals and manage consensus payloads effectively to maintain a seamless consensus mechanism.

### State Translation

One of the major challenges in previous designs was translating the EVM state to a format compatible with CometBFT. Omni overcomes this hurdle by incorporating an ABCI++ wrapper around the CometBFT engine, enabling seamless state translation and ensuring that Omni EVM blocks are efficiently converted into CometBFT transactions.

```go
// Example: ABCI++ wrapper for state translation
func (k *Keeper) PreparePayload(ctx context.Context, height uint64, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
    // Payload preparation logic...
}
```

## Advancements and Flexibility

The architectural advancements introduced by ABCI++ in Omni not only boost scalability and consensus efficiency but also provide remarkable flexibility. Omni supports the interchangeability and upgrading of execution clients without system disruptions, enabling compatibility with various Ethereum execution clients and ensuring ongoing adaptability to future blockchain innovations.
