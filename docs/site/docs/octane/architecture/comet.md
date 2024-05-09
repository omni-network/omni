---
sidebar_position: 2
---

# CometBFT ABCI 2.0

This section outlines the new CometBFT **ABCI** and **ABCI 2.0** APIs and how it allows us to use CometBFT for consensus (the previous version didn’t).

## ABCI 1.0

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

## ABCI 2.0

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
