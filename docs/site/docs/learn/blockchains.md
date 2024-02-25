---
sidebar_position: 2
id: blockchains
---

# Modular Blockchains

The concept of modular blockchains is a mental model for building blockchains that has grown in popularity over the past few years in the crypto industry.

## Monolithic Blockchains

Historically, blockchains were built in a monolithic fashion. Most Layer 1 blockchains were built as an "end to end" service – taking responsibility for all services that a blockchain must provide. Layer 1 solutions like Bitcoin and Ethereum were built in this fashion.

As we learned more, new Layer 1s were developed that optimized their stacks to make different trade-offs in terms of security, scalability, and speed. Many of these new Layer 1s adopted the monolithic approach – optimizing the full chain stack to provide the optimal desired experience. Examples of these newer monolithic Layer 1s include [Solana](https://solana.com/), [Aptos](https://aptosfoundation.org/), and [Sui](https://sui.io/).

## Modular Blockchains

However, as time has passed the latest innovations have grown to more closely reflect the architecture of microservices in the traditional software engineering world. Blockchain developers realized that we could decouple the various components that make up a blockchain and optimize them independently.

This approach was first proposed in Ethereum's rollup-centric roadmap, and more recently has been popularized by projects such as [Arbitrum](https://arbitrum.io/), [Celestia](https://celestia.org/) and now Omni.

### Components

Today, our understanding of blockchains has led us to disambiguate the stack into three primary components, broadly speaking:

- **Data Availability**: ensuring that a blockchain's transaction data is available to network participants at block proposal time (read further [here](https://ethereum.org/en/developers/docs/data-availability/)).
- **Execution**: running a blockchain's state transition function (STF) on the input data (transactions from the Data Availability module) and computing its output.
- **Consensus**: mechanism for nodes to agree on the computed output of the execution layer (and optionally finalizing the chain's state).

## Ethereum and Rollups

As the Ethereum ecosystem developed, we learned that scaling transaction throughput and decreasing costs would be extraordinarily difficult on Layer 1 Ethereum while maintaining its high security. Over time, various proposals have attempted to solve this problem: state channels, side chains, and plasma to name a few. However, each of these options came with unfortunate trade-offs that typically sacrificed security to achieve these goals. But their development was not in vain! They led to the growth of rollups, the currently agreed upon optimal solution for scaling Ethereum.

Hence, the [rollup-centric roadmap](https://ethereum-magicians.org/t/a-rollup-centric-ethereum-roadmap/4698) was proposed.

Today, Ethereum Layer 1 handles Data Availability, Execution, and Consensus. But over time, Ethereum will primarily become a Data Availability layer. Rollups will handle execution off-chain, and utilize either Ethereum's smart contract layer, or their own sovereign settlement layer, for consensus. This allows rollups (Layer 2s) to borrow Ethereum security while scaling throughput and decreasing cost. Upgrades like [EIP-4844](https://www.eip4844.com/) will even further improve Ethereum as a Data Availability layer by providing the service at lower cost.

You can view the growth of rollup ecosystems at [L2Beat.com](https://l2beat.com/), a great resource for understanding the security trade-offs, current state of development, and adoption levels for rollups.
