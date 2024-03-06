---
sidebar_position: 1
---

# Overview

Omni is designed to enhance the Ethereum ecosystem's scalability and interoperability. It integrates smart contracts across the Omni Chain EVM and Ethereum rollups for cross-chain interactions, with security ensured by a dPoS validator set, supported by a native token and re-staked ETH through Eigenlayer. This framework enables efficient cross-network storage and contract calls.

## Components

<figure>
  <img src="/img/high-level-components.png" alt="Components" />
  <figcaption>*High-Level Overview of the Components of the Network*</figcaption>
</figure>

- The **Omni network** is responsible for operating the Omni EVM and facilitating cross-network messages (`XMsg`). It consists of two internal chains, a consensus layer and an execution layer, similar to post-merge Ethereum.
  - **Execution layer** - is implemented by a standard Ethereum execution client, like `geth`, `erigon`, etc, providing the Omni EVM.
  - **Consensus layer** - is implemented by the Omni Consensus client, `halo`, that uses CometBFT to provide security for cross-chain messaging and for the Omni execution layer.
- **EigenLayer smart contracts** exist on Ethereum L1 and connect the Omni network with its re-staking participants. The Omni network is registered with EigenLayer as an “Actively Validated Service” AVS and Omni validators serve the role of “Operators” of the AVS.
- [**Portal contracts**](../xmessages/components/portal.md) implement the on-chain logic of the Omni protocol and serve as the main interface for creating cross-network messages. They are deployed to all supported rollup VMs as well as the Omni EVM on the Omni network. They all have the same address and calls to and from are abstracted with the [solidity interface](https://github.com/omni-network/omni/blob/22bd4460e254eee4ebf79239897ea04ba9b2db43/contracts/src/interfaces/IOmniPortal.sol).
- [**Relayer**](../xmessages/components/relayer.md) responsible for delivering attested cross-network messages from the Omni network to destination rollup VMs. Monitors the Omni Consensus Layer until ⅔ (>66%) of the validator set attested to the “next” block on each source chain, then proceeds to forwarding the respective `XMsg` list included in the block.

## Following a User Cross-Rollup Action

If we were to follow a simple initiating cross-rollup user call from a rollup (in this example Arbitrum) to another rollup (Optimism), the path taken by the information would look as shown below from a high level.

<figure>
  <img src="/img/high-level-arch.svg" alt="High-Level Arch" />
  <figcaption>*Following a user deposit call to an xapp*</figcaption>
</figure>

### Stepwise Walkthrough

Note: we refer to an `xapp` as a smart contract application that exists on multiple chains. In this example, we'll use Arbitrum as the "source chain" and Optimism as the "destination chain".

1. User calls a function on the xapp contract on Arbitrum that intends to interact with a contract on Optimism
2. The source xapp contract calls the `xcall` method on the Omni Portal contract on Arbitrum[^1]
3. The Portal contract emits a `XMsg` Event containing relevant data for the destination chain contract call
4. Validators read the emitted Event, create an `xBlock` & attest to it.
5. The Relayer service reads the attestations and pushes the information from the `XMsg` in the attested `xBlock` to the destination chain by calling the destination Portal contract's `xsubmit` method
6. The Portal Contract on the destination chain performs the a contract call to the specified method in the destination contract as specified by the original call `xcall` in **2**.

To read further on this message traversal see the [`XMsg` Lifecycle](../xmessages/xmsg.md) section.
