# Omni Octane

## Introduction

Octane is a next-generation modular framework for the EVM. Developers can use octane to run:

- any EVM execution client that implements the EngineAPI, coupled with
- any consensus client that implements ABCI++.

It is designed for modularity, scalability, and performance, and is built to simplify the integration of the EVM into any blockchain application.

Octane powers the Omni blockchain on testnet and mainnet. You can run Octane separately for your own use case, e.g. for your own EVM blockchain with fast finality, as a decentralized sequencer, and more.

## Overview

Octane is the first consensus implementation of the EngineAPI besides Ethereum itself (the Beacon Chain).

Octane connects any EVM execution client to an instant finality consensus mechanism, and can be used to build virtually any application that requires an EVM. It is built for 1) modularity, 2) performance and scalability, 3) EVM equivalence, and 4) future-proof support for highly performant EVM clients and even non-EVM clients that support the EngineAPI. You can read more about our goals for Octane and specific targets below.

Here are some key use cases for Octane:

- **Cosmos Chains**: If you are building a Cosmos chain and aim to integrate EVM functionalities, Octane provides a reliable, performant, and scalable framework using CometBFT consensus.
- **Layer 2 Rollups:** Octane can be employed as a decentralized sequencer that interfaces with any data availability (DA) layer, including but not limited to [Celestia](https://docs.celestia.org/learn/how-celestia-works/data-availability-layer), [Avail](https://docs.availproject.org/docs/introduction-to-avail/avail-da), and [EigenDA](https://docs.eigenlayer.xyz/eigenda/overview).
- **Layer 3 Networks and beyond:** For projects developing an L3+ solution, such as an [Arbitrum Orbit](https://docs.arbitrum.io/launch-orbit-chain/orbit-gentle-introduction) chain or an [OP Stack](https://docs.optimism.io/) L3, Octane can serve as a decentralized sequencer to build blocks for the network.
- **Optimized EVM clients:** with Octane, you can connect a fast-finality gadget (CometBFT) to any EVM client that implements the Engine API. Thus, it is compatible with projects like Reth, Monad, Sei, and more.
- **Non-EVM Execution clients**: the Octane consensus module implements the consensus side of the EngineAPI. Therefore, it can integrate with non-EVM execution layers, as long as they implement the EngineAPI. We believe that the EngineAPI, established by Ethereum’s Proof-of-Stake architecture, can and should be used as the standard for connecting modular chains (the API between the consensus and execution clients).

## The Story of Octane

In the early days, the Omni team wanted to offer a fast consensus mechanism for the EVM to support programming global logic and state into cross-rollup applications. A natural choice was to build a cosmos chain with EVM support. In late ‘22 and early ‘23, we implemented a proof of concept of the Omni blockchain leveraging Ethermint – at the time, the standard for EVM on Cosmos. In fact, we launched with this architecture for our first two testnets.

The team behind Ethermint used tooling that was best in class at the time, but unfortunately its benchmarks were severely limited in the categories we cared about:

- Scalability & Performance
  - 5 second block times
  - ~20-30 TPS when load tested
  - 500ms RPC queries under standard load
  - State store translation between EVM state and Cosmos state, which create a major disk performance bottleneck
  - Consensus logic breaking during periods of high EVM activity, as the Cosmos mempool is not built for EVM scale load
- EVM Compatibility
  - It requires an EVM adapter, a set of logic that wraps geth to be compatible with Cosmos, so EVM upgrades would need to go through several layers of implementation before being included
  - It is partially EVM compatible, but because of this adapter would never be EVM equivalent
  - It only supported one execution client: geth.
- Long Term Maintenance
  - Any time the EVM is upgraded, it requires new adapter modules to be built for the new opcodes
  - infeasibly high disk usage, with Cosmos + EVM state

Because this architecture did not meet our requirements for scalability, performance, EVM equivalence, or long term maintenance, the Omni team went back to the drawing board. How could we solve these problems and what were our goals for the re-architecture?

### Goals

- Separate execution from consensus
  - The Cosmos mempool is not built for EVM scale. Could we separate the EVM layer from the consensus layer?
- EVM equivalence
  - We wanted our implementation to be no different than the code running on Ethereum L1, so that 100% of developer tooling is compatible with Omni.
- Performance
  - Many CometBFT chains support subsecond block times and finality, can a CometBFT chain running an EVM also support this?
- Modularity and Client diversity
  - We wanted our implementation to work not just for a single EVM client, but for any EVM client in the future. If there are fundamental optimizations that happen to the EVM, we want Omni to be able to support them.

After months of research, we found that with the release of the EngineAPI (Ethereum Proof-of-Stake) and ABCI 2.0 (CometBFT release), we could rearchitect Omni with a model that achieves all of these goals. This architecture was not possible before 2024.

_Enter Octane._

## Architecture

You can read more about the architecture in our documentation portal – check out [this page](https://docs.omni.network/protocol/evmengine/dual) for an overview, and [this one](https://docs.omni.network/protocol/evmengine/lifecycle) for details on the EVM payload lifecycle. Continue reading below for a high level overview.

### Overview

Octane is built on a sophisticated modern architecture that includes the following components:

- It implements the server side of the [ABCI++](https://github.com/cometbft/cometbft/tree/main/spec/abci) interface to mesh CometBFT consensus with the Engine API.
- It drives the Execution Layer via the [Engine API](https://github.com/ethereum/execution-apis/blob/main/src/engine/common.md) for the block building engine.

<figure>
  <img src="/img/consensus.png" alt="Consensus in halo" />
  <figcaption>*The consensus process in `halo`*</figcaption>
</figure>

### EVM Client

- Receives EVM transactions and maintains the EVM mempool
- Builds EVM blocks from transactions in the mempool
- Sends EVM block payloads to the consensus client via the EngineAPI
- Stores EVM state

### Consensus Module

- Receives EVM payloads (blocks) and commits them to consensus module state as a single transaction
- Stores strictly consensus module state (it does not store any EVM state besides the block metadata, which is passed via the EngineAPI)
- Can read logs from the EVM via the EngineAPI, like the Beacon Chain Deposit Contract.

This is a modular approach to execution and consensus, using frontier standards – the EngineAPI and ABCI 2.0.

## Purpose and Benefits

- Rather than wrapping EVM transactions as Cosmos transactions (monolithic approach), EVM transactions are handled by the EVM, while consensus chain transactions are handled by the consensus module. This solves the bottleneck of forcing the Cosmos mempool to handle EVM load.
- This modular approach allows the EVM to scale independently from consensus, by simply adopting the latest performant execution client.
- Transaction throughput and gas throughput is limited only by the execution client, which can be highly optimized.
- Supports blocktimes ~1 second without optimizations.
- Is EVM equivalent: the code running the EVM is an unmodified EVM client, so it is exactly the same code that is running for Ethereum L1.
- Supports client diversity, as any EVM client that implements the EngineAPI can be used.
- Consensus chain transactions can be proxied through the EVM (via predeploys), if desired. Omni’s consensus chain implementation uses predeploys for all consensus chain transactions.
- Can support non-EVM execution clients if they implement the EngineAPI.

## Development and Integration

- **Codebase:** The `octane` modules are part of Omni’s broader consensus chain implementation, called Halo. Our [evmengine](https://github.com/omni-network/omni/tree/main/halo/evmengine) and [evmstaking](https://github.com/omni-network/omni/tree/main/halo/evmstaking) modules, will soon be refactored into their own top-level directory as the protocol matures. This change will simplify the integration process for other teams, and will still be used by Halo.
- We are excited about the potential of Octane and the capabilities it brings to the Omni network. As we continue to refine and expand this section, we look forward to empowering developers and partners with powerful tools to build the next generation of blockchain applications.
- We encourage teams to fork, experiment, and play around with Octane for your own use cases.

Visit [our monorepo](https://github.com/omni-network/omni/) or reach out to the team to learn more!

:::info

Octane is currently licensed under GPLv3 with an open interoperability requirement. We are working with teams that would like to use this implementation, and if you would like to integrate it into your own application, you can reach out to the Omni team to discuss further.

:::
