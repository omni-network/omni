---
sidebar_position: 2
---

# Introduction

Below you'll find a high-level overview of the Omni Network, its design, and its functionality.

## Components

<figure>
  <img src="/img/high-level-components.png" alt="Components" />
  <figcaption>*High-Level Overview of the Components of the Network*</figcaption>
</figure>

- The **Omni network** is responsible for operating the Omni EVM and facilitating cross-network messages (`XMsg`). It consists of two internal chains, a consensus layer and an execution layer, similar to post-merge Ethereum.
  - **Execution layer** - is implemented by a standard Ethereum execution client, like `geth`, `erigon`, etc, providing the Omni EVM.
  - **Consensus layer** - is implemented by the Omni Consensus client, `halo`, that uses CometBFT to provide security for cross-chain messaging and for the Omni execution layer.
- **EigenLayer smart contracts** exist on Ethereum L1 and connect the Omni network with its re-staking participants. The Omni network is registered with EigenLayer as an “Actively Validated Service” AVS and Omni validators serve the role of “Operators” of the AVS.
- [**Portal contracts**](../architecture/portal.md) implement the on-chain logic of the Omni protocol and serve as the main interface for creating cross-network messages. They are deployed to all supported rollup VMs as well as the Omni EVM on the Omni network. They all have the same address and calls to and from are abstracted with the [solidity interface](https://github.com/omni-network/omni/blob/22bd4460e254eee4ebf79239897ea04ba9b2db43/contracts/src/interfaces/IOmniPortal.sol).
- [**Relayer**](../architecture/relayer.md) responsible for delivering attested cross-network messages from the Omni network to destination rollup VMs. Monitors the Omni Consensus Layer until ⅔ (>66%) of the validator set attested to the “next” block on each source chain, then proceeds to forwarding the respective `XMsg` list included in the block.

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

To read further on this message traversal see the [`XMsg` Lifecycle](../architecture/xmsg.md) section.

## Finality

After verifying the contents of a proposed consensus block and appending `XBlock` hash attestations, `halo` clients use CometBFT consensus to vote on the validity of the block.

If the proposal is determined valid by ⅔ of the active validator set, the block is finalized and becomes the latest block committed to the network. These confirmed blocks contain all transactions from the Omni EVM along with the hashes of attested `XBlock`s, allowing anyone with a validator node for a given rollup VM to reconstruct any `XBlock` and verify its contents.

Omni’s [Parallelized Consensus](../architecture/components.md#parallelized-consensus--cometbft) processes state transitions for both the Omni EVM and external VMs, preventing both complementary sub-processes from interfering with one another prior to block finalization. ABCI++ allows validators to keep `XBlock` attestations separate from the single CometBFT transaction within each block.

## Fee Model

Omni fees start with a basic, Solidity-based fee payment interface and an uncomplicated pricing mechanism, with room for future enhancements. The network only supports **\$ETH** for fee payments presently. In the future, developers will also be able to pay with **\$OMNI** and potentially other tokens if they desire, but **\$ETH** will always be supported.

Fees are paid in **\$ETH** and calculated in real-time during transactions via the payable `xcall` function on the portal contracts, ensuring simplicity for developers and compatibility with existing Ethereum tooling. This setup allows for easy off-chain fee estimations and the possibility for developers to pass the cost on to users, with a straightforward upgrade path to a more dynamic fee structure that can adapt to the network's evolving needs without necessitating changes to developer contracts.

For more information on how fees are handled read the [fees protocol section](../architecture/fees.md)
