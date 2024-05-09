---
sidebar_position: 2
---

# Use Cases

**Octane** powers the Omni blockchain on testnet and mainnet. You can run Octane to build virtually any application that requires an EVM e.g. for your own EVM blockchain with fast finality, as a decentralized sequencer, and more.

Here are some potential use cases for Octane:

- **Cosmos SDK Chains**: If you are building a Cosmos SDK chain and aim to support an EVM, Octane provides a reliable, performant, and scalable framework using CometBFT consensus.
- **Optimized EVM clients:** you can connect a fast-finality gadget (CometBFT) to any EVM client that implements the Engine API. Thus, it is compatible with projects like Reth, Monad, and more.
- **Experimental EVM clients**: many teams fork an EVM client and add experimental functionality, such as new precompiles, EIPs that haven't been released on Ethereum mainnet yet, and/or new opcodes (e.g. running AI inference). In most cases, these experimental clients still implement the EngineAPI, and thus are 100% compatible with Octane.
- **Layer 2 Rollups:** Octane can be employed as a decentralized sequencer that interfaces with any data availability (DA) layer, including but not limited to [Celestia](https://docs.celestia.org/learn/how-celestia-works/data-availability-layer), [Avail](https://docs.availproject.org/docs/introduction-to-avail/avail-da), and [EigenDA](https://docs.eigenlayer.xyz/eigenda/overview).
- **Layer 3 Networks and beyond:** For projects developing an L3+ solution, such as an [Arbitrum Orbit](https://docs.arbitrum.io/launch-orbit-chain/orbit-gentle-introduction) chain or an [OP Stack](https://docs.optimism.io/) L3, Octane can serve as a decentralized sequencer to build blocks for the network.
- **Non-EVM Execution clients**: the Octane consensus module implements the consensus side of the EngineAPI. Therefore, it can integrate with non-EVM execution layers, as long as they implement the EngineAPI. We believe that the EngineAPI, established by Ethereum’s Proof-of-Stake architecture, can and should be used as the standard for connecting modular chains (the API between the consensus and execution clients).
