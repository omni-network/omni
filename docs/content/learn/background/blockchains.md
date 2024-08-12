---
sidebar_position: 1
---

# Blockchains

The concept of modular blockchains is a mental model for building blockchains that has grown in popularity over the past few years in the crypto industry. Let's discuss the history of the idea, and how it's relevant to Omni.

## Monolithic Blockchains

Historically, blockchains were built in a monolithic fashion. Most Layer 1 blockchains were built as an "end to end" service – taking responsibility for all services that a blockchain must provide. Layer 1 solutions like Bitcoin and Ethereum were built in this fashion.

As the industry evolved, new Layer 1s were developed that optimized their stacks to make different trade-offs in terms of security, scalability, and speed. Many of these new Layer 1s adopted the monolithic approach – optimizing the full chain stack to provide the optimal desired experience. Examples of these newer monolithic Layer 1s include [Solana](https://solana.com/), [Aptos](https://aptosfoundation.org/), and [Sui](https://sui.io/).

## Modular Blockchains

However, as time has passed the latest innovations have grown to more closely reflect the architecture of microservices in the traditional software engineering world. Blockchain developers realized that we could decouple the various components that make up a blockchain and optimize them independently.

This approach was first proposed in Ethereum's rollup-centric roadmap, and more recently has been popularized by projects such as [Arbitrum](https://arbitrum.io/), [Celestia](https://celestia.org/) and now Omni.

### Components

Today, our understanding of blockchains has led us to disambiguate the stack into three primary components, broadly speaking:

- **Data Availability**: ensuring that a blockchain's transaction data is available to network participants at block proposal time (read further [here](https://ethereum.org/en/developers/docs/data-availability/)).
- **Execution**: running a blockchain's state transition function (STF) on the input data (transactions from the Data Availability module) and computing its output.
- **Consensus**: mechanism for nodes to agree on the computed output of the execution layer (and optionally finalizing the chain's state).
