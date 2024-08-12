---
sidebar_position: 2
---

# Component Overview

The Omni protocol is composed of five primary components: **rollup networks**, **the Omni Network**, **EigenLayer restaking contracts**, **Omni Portal contracts**, and **relayers**.

<figure>
  <img src="/img/components.png" alt="Components" />
  <figcaption>*Overview of the Components of the Network*</figcaption>
</figure>

- A **rollup network** is any Ethereum rollup that performs off-chain transaction execution before settling to Ethereum L1.
- The **Omni Network** is a layer 1 blockchain that connects rollup VMs. Similar to Ethereum, Omni nodes are separated into distinct execution and consensus layers. It consists of two internal chains, a consensus layer and an execution layer, similar to post-merge Ethereum.
    - Similar to Ethereum, Omni nodes are separated into distinct execution and consensus environments
        - The **execution layer** is implemented by standard Ethereum execution clients. like  **`geth`**, **`erigon`**, etc, to provide the Omni EVM.
        - The **consensus** layer is implemented by the Omni consensus client, halo, and uses CometBFT for consensus on XMsgs and Omni EVM blocks.
- **EigenLayer restaking contracts** exist on Ethereum L1 and connect Omni with its restaking participants. Omni is registered with EigenLayer as an [Actively Validated Service](https://docs.eigenlayer.xyz/eigenlayer/overview/key-terms) (AVS) and Omni validators serve the [Operator](https://docs.eigenlayer.xyz/eigenlayer/overview/key-terms) for the AVS.
- **[Portal contracts](../xmessages/components/portals.md)** implement the on-chain logic of the Omni protocol and serve as the main interface for creating cross-network messages. They are deployed to all supported rollup VMs as well as the Omni EVM on the Omni network. They all have the same address and calls to and from are abstracted with the [solidity interface](https://github.com/omni-network/omni/blob/22bd4460e254eee4ebf79239897ea04ba9b2db43/contracts/src/interfaces/IOmniPortal.sol).
- **[Relayers](../xmessages/components/relayer.md)** are responsible for delivering attested cross-network messages from the Omni network to destination rollup VMs. Monitors the Omni Consensus Layer until ⅔ (>66%) of the validator set attested to the “next” block on each source chain, then proceeds to forwarding the respective `XMsg` list included in the block.
