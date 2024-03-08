---
sidebar_position: 2
---

# High-Level Component Overview

<figure>
  <img src="/img/high-level-components.png" alt="Components" />
  <figcaption>*High-Level Overview of the Components of the Network*</figcaption>
</figure>

- The **Omni network** is responsible for operating the Omni EVM and facilitating cross-network messages (`XMsg`). It consists of two internal chains, a consensus layer and an execution layer, similar to post-merge Ethereum.
  - **Execution layer** - is implemented by a standard Ethereum execution client, like `geth`, `erigon`, etc, providing the Omni EVM.
  - **Consensus layer** - is implemented by the Omni Consensus client, `halo`, that uses CometBFT to provide security for cross-chain messaging and for the Omni execution layer.
- **EigenLayer smart contracts** exist on Ethereum L1 and connect the Omni network with its re-staking participants. The Omni network is registered with EigenLayer as an “Actively Validated Service” AVS and Omni validators serve the role of “Operators” of the AVS.
- [**Portal contracts**](../xmessages/components/portal-send.md) implement the on-chain logic of the Omni protocol and serve as the main interface for creating cross-network messages. They are deployed to all supported rollup VMs as well as the Omni EVM on the Omni network. They all have the same address and calls to and from are abstracted with the [solidity interface](https://github.com/omni-network/omni/blob/22bd4460e254eee4ebf79239897ea04ba9b2db43/contracts/src/interfaces/IOmniPortal.sol).
- [**Relayer**](../xmessages/components/relayer.md) responsible for delivering attested cross-network messages from the Omni network to destination rollup VMs. Monitors the Omni Consensus Layer until ⅔ (>66%) of the validator set attested to the “next” block on each source chain, then proceeds to forwarding the respective `XMsg` list included in the block.
