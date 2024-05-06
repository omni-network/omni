# Omni Octane

## Introduction

Octane is a next-generation modular framework for the EVM. Developers can use octane to run:

- any EVM execution client that implements the [EngineAPI](https://hackmd.io/@danielrachi/engine_api)
- any consensus client that implements [ABCI 2.0](https://github.com/cometbft/cometbft/tree/main/spec/abci)

Octane is the first consensus implementation of the EngineAPI besides Ethereum itself (the Beacon Chain).

It is designed and built to simplify the integration of the EVM into any blockchain application.

## Use Cases

**Octane** powers the Omni blockchain on testnet and mainnet. You can run Octane to build virtually any application that requires an EVM e.g. for your own EVM blockchain with fast finality, as a decentralized sequencer, and more.

Here are some potential use cases for Octane:

- **Cosmos SDK Chains**: If you are building a Cosmos SDK chain and aim to support an EVM, Octane provides a reliable, performant, and scalable framework using CometBFT consensus.
- **Optimized EVM clients:** you can connect a fast-finality gadget (CometBFT) to any EVM client that implements the Engine API. Thus, it is compatible with projects like Reth, Monad, and more.
- **Experimental EVM clients**: many teams fork an EVM client and add experimental functionality, such as new precompiles, EIPs that haven't been released on Ethereum mainnet yet, and/or new opcodes (e.g. running AI inference). In most cases, these experimental clients still implement the EngineAPI, and thus are 100% compatible with Octane.
- **Layer 2 Rollups:** Octane can be employed as a decentralized sequencer that interfaces with any data availability (DA) layer, including but not limited to [Celestia](https://docs.celestia.org/learn/how-celestia-works/data-availability-layer), [Avail](https://docs.availproject.org/docs/introduction-to-avail/avail-da), and [EigenDA](https://docs.eigenlayer.xyz/eigenda/overview).
- **Layer 3 Networks and beyond:** For projects developing an L3+ solution, such as an [Arbitrum Orbit](https://docs.arbitrum.io/launch-orbit-chain/orbit-gentle-introduction) chain or an [OP Stack](https://docs.optimism.io/) L3, Octane can serve as a decentralized sequencer to build blocks for the network.
- **Non-EVM Execution clients**: the Octane consensus module implements the consensus side of the EngineAPI. Therefore, it can integrate with non-EVM execution layers, as long as they implement the EngineAPI. We believe that the EngineAPI, established by Ethereum’s Proof-of-Stake architecture, can and should be used as the standard for connecting modular chains (the API between the consensus and execution clients).

## The Evolution of Octane

Omni was designed to solve the fragmentation problem in Ethereum's L2 ecosystem. One of the critical desired properties required to solve this problem was and is to introduce _global state_. Global state empowers smart contract developers to build unified applications that benefit from network effects across deployments, and is a core value proposition of Omni. To provide this in a developer friendly way, the natural choice was to introduce an EVM.

In late ‘22 and early ‘23, the Omni team implemented a proof of concept of the Omni blockchain that utilized Ethermint – at the time, the standard for EVM on Cosmos. In fact, we launched with this architecture for our first two testnets.

The team behind Ethermint used tooling that was best in class at the time, but unfortunately its benchmarks were severely limited in the categories we cared about:

- **Scalability & Performance**
  - 5 second block times
  - ~25 TPS when load tested
  - 500ms RPC queries under standard load
  - State store translation between EVM state and Cosmos state, which created a major disk performance bottleneck
  - Consensus logic breaking during periods of high EVM activity, as the Cosmos mempool is not built for EVM scale load
- **EVM Compatibility**
  - It requires an EVM adapter, a set of logic that wraps geth to be compatible with Cosmos, so EVM upgrades would need to go through several layers of implementation before being included
  - It is partially EVM compatible, but because of this adapter it would never be EVM equivalent
  - It only supported one execution client: geth.
- **Long-Term Maintenance and Upgradeability**
  - Any time the EVM is upgraded, it required new adapter modules to be built for the new opcodes
  - infeasibly high state bloat, with Cosmos + EVM state
  - only works with geth, so new high performance EVM clients could never be used

Because this architecture did not meet our requirements for scalability, performance, EVM equivalence, or long term maintenance, the Omni team decided not to proceed with this architecture. How could we solve these problems and what were our goals for the re-architecture?

### Goals

- **Decouple execution from consensus**
  - Multiple execution layer implementations are available with new ones being added constantly. With a decoupled architecture, the latest and greatest implementation can always be introduced.
  **- The Cosmos mempool and state store are not built for EVM transactions, we want to separate consensus processing from execution processing #modularity
- **EVM equivalence**
  - We wanted our implementation to be no different than the code running on Ethereum L1, so that 100% of developer tooling is compatible with Omni.
- **Performance**
  - As an async messaging protocol, fast block times are critical, ideally targeting subsecond finality
  - To introduce a messaging protocol that scales for the entire Ethereum ecosystem, Omni's TPS and GPS must not hit fundamental bottlenecks like the previous implementation

After many months of research, proof-of-concepts, and benchmarking, we we able to design an architecture that achieve all of these goals – unlocked by the release of the EngineAPI (Ethereum Proof-of-Stake) and ABCI 2.0 (CometBFT).

This architecture was not possible before 2024.

_Enter Octane._

## Architecture

You can read more about the architecture in our documentation portal – check out [this page](https://docs.omni.network/protocol/evmengine/dual) for an overview, and [this one](https://docs.omni.network/protocol/evmengine/lifecycle) for details on the EVM payload lifecycle. Continue reading below for a high level overview.

### Overview

Octane is built on a sophisticated modern architecture that includes the following components:

#### Ethereum Engine API

Ethereum introduced the [Engine API](https://github.com/ethereum/execution-apis/tree/main/src/engine) to decouple the execution layer from the consensus layer. This is almost like a `consensus.EngineV2` interface, except that this is an HTTP RPC interface and therefore introduces more decoupling.

Here is a [visual guide](https://hackmd.io/@danielrachi/engine_api) and this is the original [design space doc](https://hackmd.io/@n0ble/consensus_api_design_space).

The benefit to using the Engine API for decoupling execution vs consensus is that multiple execution layer implementations are available with new ones being added constantly. This means that the latest and greatest EVM implementation is always available for use (standing on the shoulders of giants).

Noteworthy implementations:

- [geth](https://github.com/ethereum/go-ethereum): the original and most-used go implementation.
- [erigon](https://github.com/ledgerwatch/erigon): GA from [2022](https://erigon.substack.com/p/post-merge-release-of-erigon-dropping); performant go implementation, almost like geth v2.
- [reth](https://github.com/paradigmxyz/reth): soon to be available rust implementation, almost like Erigon v2.

Octane drives its execution layer via the EngineAPI for the block building engine.

#### CometBFT ABCI 2.0

This section outlines the new CometBFT **ABCI** and **ABCI 2.0** APIs and how it allows us to use CometBFT for consensus (the previous version didn’t).

##### ABCI 1.0

The main purpose behind [ABCI](https://docs.tendermint.com/v0.33/app-dev/app-development.html) (**A**pplication **B**lock**C**hain **I**nterface) is to provide an interface between the application logic (application that someone is developing, our blockchain business logic) and the consensus engine. Application logic is responsible for applying state and to validate transactions. The consensus engine is responsible for ensuring that all transactions are replicated in the same order on every machine. Machines in the consensus engine are validators who apply the transaction logic to the application state (e.g. increase/decrease a user's balance).

The consensus engine and the application logic can be viewed as a client/server relationship. CometBFT maintains 3 connections: mempool, consensus connection, and query.

There are a couple of important ABCI methods that require implementation:

- `CheckTx`: verifies whether the transaction is valid from the application's perspective. It checks for the correctness and validity of the transaction without affecting the state of the blockchain
- `BeginBlock`: allows the application to perform any necessary setup or updates at the start of a new block. It may include initializing variables or performing tasks that are specific to the block being processed
- `DeliverTx`: applies the changes of a transaction to the application's state. It is responsible for updating the blockchain state
- `EndBlock`: allows the application to perform any necessary cleanup or calculations at the end of a block

The whole transaction flow is handled by the proposer and the block gets executed and state applied to the app. The proposer has full control over the whole block construction process, over the ordering of transactions, and over what goes into the block.

<figure>
  <img src="/img/abci-flow.png" alt="ABCI Flow" />
  <figcaption>*the standard ABCI 1.0 flow*</figcaption>
</figure>

##### ABCI 2.0

CometBFT has implemented a upgraded API for improved customisation called **ABCI 2.0**. See the [Spec](https://github.com/cometbft/cometbft/tree/main/spec/abci). In contrast, [ABCI 2.0](https://docs.cometbft.com/v0.38/spec/abci/abci++_basic_concepts) gives a lot more control and granularity over the process. It’s much more of a back-and-forth process between the Application and Consensus layers.

It adds the following methods to the original ABCI API:

- `PrepareProposal`: allows the **proposer’s** application to modify the transactions in a block, add, remove, reorder, replace.
- `ProcessProposal`: allows validators to access to proposed blocks for validation and eager processing.
- `FinalizeBlock`: encapsulates existing `BeginBlock, [DeliverTx], EndBlock` methods in a single method.
- `ExtendVote`: allows validators to append arbitrary data to their vote allowing validators to include data in blocks (not just proposers).
- `VerifyVoteExtension`: allows validators to verify other validator vote extensions.

These changes to CometBFT support a range of new use cases. See how this is used in Interplanetary Consensus ([1](https://docs.google.com/document/d/1cFoTdoRuYgxmWJia6K-b5vmEj-4MvyHCNvShZpyconU/edit), [2](https://docs.ipc.space/key-concepts/architecture#abci++)).

This could be applied to the Omni usecase:

- Use Execution Layer as the EVM and expose the **Engine API** for decoupled consensus.
- Build a consensus layer using CometBFT **ABCI 2.0**.
- Do not use CometBFT transactions, nor the CometBFT mempool
- Proposers provide the EVM block (EngineAPI payload) as a single transaction to CometBFT via the `PrepareProposal`

<figure>
  <img src="/img/abci2-flow.png" alt="ABCI 2.0 Flow" />
  <figcaption>*the ABCI 2.0 flow*</figcaption>
</figure>

#### Block Building Flow

Putting it all together, the combined flow between the **EngineAPI** and **ABCI 2.0** does the following:

##### EVM Client

- Receives EVM transactions from users and maintains the EVM mempool
- Builds EVM blocks from transactions in the mempool
- Sends EVM block payloads to the consensus client via the EngineAPI
- Stores EVM state

##### Consensus Module

- Receives EVM payloads (blocks) and commits them to consensus module state as a single transaction
- Stores strictly consensus module state (it does not store any EVM state besides the block metadata, which is passed via the EngineAPI)
- Can read logs from the EVM via the EngineAPI, like the Beacon Chain Deposit Contract.

Here is a visual overview of the flow:

<figure>
  <img src="/img/consensus.png" alt="Consensus in octane" />
  <figcaption>*The consensus process in `octane`*</figcaption>
</figure>

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
