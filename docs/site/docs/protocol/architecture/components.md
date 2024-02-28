---
sidebar_position: 3
---

# Validators

## Components & Communication

Omni nodes are configured with a new modular framework based on [Ethereum’s Engine API](https://github.com/ethereum/execution-apis/tree/4b225e0d273e92982b2c539d63eaaa756c5285a4/src/engine) to combine `halo` with the EVM execution client.

<figure>
  <img src="/img/validator.png" alt="Validator" />
  <figcaption>*Overview of the Components Run by Validators and Interactivity*</figcaption>
</figure>

Using the Engine API, Omni nodes pair existing high performance Ethereum execution clients with a new consensus client, referred to as `halo`, that implements CometBFT consensus.The Engine API allows clients to be substituted or upgraded without breaking the system. This allows the protocol to maintain flexibility on Ethereum and Cosmos technology while promoting client diversity within its execution layer and consensus layer. We consider this new network framework to be a public good that future projects may leverage for their own network designs.

## Halo

It implements the server side of the ABCI++ interface and drives the Omni Execution Layer via the Engine API. CometBFT validators attest to source chain blocks containing cross chain messages using CometBFT Vote Extensions.

Omni's consensus layer is established through the Halo consensus client which supports a Delegated Proof of Stake (DPoS) consensus mechanism implemented by the [CometBFT consensus engine](https://docs.cometbft.com/v0.38/) (formerly known as Tendermint). CometBFT stands out as the optimal consensus engine for Omni nodes for three reasons:

1. CometBFT achieves immediate transaction finalization upon block inclusion, eliminating the need for additional confirmations. This feature allows Omni validators to agree on a single global state for all connected networks, updated every second.
2. CometBFT provides support for DPoS consensus mechanisms, allowing re-staked ETH to be delegated directly to individual Omni validators as part of the network’s security model.
3. CometBFT is one of the most robust and widely tested PoS consensus models and is already used in many production blockchain networks securing billions of dollars.

By default, Omni validators ensure the integrity of rollup transactions by awaiting their posting and finalization on the Ethereum L1 before attesting to them. This proactive approach mitigates the risk of reorganizations on both Ethereum L1 and the rollup, enhancing the overall security and reliability of the system.

### Consensus Process

Omni Consensus clients process the consensus blocks and maintain consensus state that tracks the “status” of each `XBlock` (by the validator set) from `Pending` to `Approved` (including an `AggregateAttestation`). Then only the “latest approved” `XBlock` for each source chain needs to be maintained; earlier `XBlock`s can be trimmed from the state. Validators in the current validator set must attest to all subsequent (after the “last approved”) `XBlock`.

#### Validator Set Changes

When the validator set changes, all `XBlock`s marked as `Pending` need to be updated by:

- Updating the associated validator set ID to the current.
- Deleting all attestations by validators not in the current set.
- Updating the weights of each remaining attestation according to the new validator set.

Validators that already attest to the `Pending` marked `XBlock` during the previous validator set, do not need to re-attest. Only the new set validators must attest (ie. to all `XBlock`s after the latest approved).

### ABCI++

Cosmos’ [ABCI++](https://docs.cometbft.com/v0.37/spec/abci/) is a supplementary tool that allows `halo` clients to adapt the CometBFT consensus engine to the node’s architecture. An ABCI++ adapter is wrapped around the CometBFT consensus engine to convert messages from the Engine API into a format that can be used in CometBFT consensus. These messages are inserted into CometBFT blocks as single transactions – this makes Omni consensus lightweight and enables Omni’s sub-second finality time.

During consensus, validators also use ABCI++ to attest to the state of external Rollup VMs. Omni validators run the state transition function, $f(i, S_n)$, for each Rollup VM and compute the output, $S_{n+1}$.

### Parallelized Consensus & CometBFT

Omni introduces Parallelized Consensus, a consensus framework that allows validators to run consensus for the Omni EVM in parallel with consensus for cross-network messages without compromising on performance.

Omni’s Parallelized Consensus contains two sub-processes:

- validating state changes within the Omni EVM and
- attesting to `XBlock` hashes originating from external rollup VMs.

The aggregate Parallelized Consensus process is visualized below.

<figure>
  <img src="/img/parallel-consensus.png" alt="Parallel Consensus" />
  <figcaption>*Parallelized Consensus: Validating Omni EVM State Changes and Rollup `XBlock` Hash Attestation*</figcaption>
</figure>

1. Every `halo` client runs a node for each rollup VM to check for `XMsg` events emitted from Portal contracts.
2. For every rollup VM block that contains `XMsg`s, `halo` clients build `XBlock`s that contain the corresponding `XMsg`s.
3. Once the calldata for a rollup VM block has been posted and finalized on Ethereum L1, Omni validators use ABCI++ vote extensions to attest to the hash of the corresponding `XBlock`. These attestations are appended to the current consensus layer block.

## EngineAPI

The Engine API allows the Omni execution layer (Omni EVM) to mirror the functionality and design of Ethereum L1’s execution layer (EVM). Users send transactions to the Omni EVM mempool and execution clients share those transactions through the execution layer’s peer-to-peer (P2P) node network. For each block of transactions, a node’s execution client computes the state transition function for the Omni EVM and shares the output with its `halo` client through the Engine API. Since the transaction mempool resides on the execution layer, Omni can scale activity without congesting the CometBFT mempool used by validators on the consensus layer. Previously, using the CometBFT mempool to handle transaction requests caused the network to become overloaded and resulted in liveness disruptions. Similar challenges have been observed in other projects that adopted a comparable approach.

## Execution Consensus

1. When it is a validator's turn to propose a block, its `halo` client requests the latest Omni EVM block from its execution client using the Engine API.
2. The execution client builds a block from the transactions in its mempool and returns the block header to the `halo` client through the Engine API.
3. The `halo` client packages the new block proposal as a single CometBFT transaction and includes it in the consensus layer block.
4. The block is proposed to the rest of the validator network through the consensus layer’s P2P network.
5. Non-proposing validators use the Engine API and their execution clients to run the state transition function on the proposed block header to verify the block’s validity.

## Omni EVM

Post-merge ethereum decoupled the execution layer from the consensus layer introducing a more modular approach to building blockchains. This modular approach allows the EVM to scale (somewhat) independently from consensus, by simply adopting the latest performant execution client like `erigon` or `geth`.

The Omni consensus layer needs smart contracts to manage native staking and delegated re-staking from ETH L1. The Omni EVM is a natural fit as fees would be much lower and syncing with the consensus layer is already built-in. Providing an EVM purposely built for cross-chain dapps that has both low fees and short block times allows for a simple adoption path and hub-and-spoke mental model to onboard projects into Omni Protocol.

The Omni execution layer also benefits from using existing Ethereum execution clients without specialized modifications. This approach eliminates the risk of introducing new bugs and increases release velocity for features strictly related to cross-rollup interoperability. Furthermore, Omni nodes can seamlessly adopt upgrades from any EVM client, ensuring ongoing compatibility with the Omni consensus layer. For example, the Omni EVM natively supports dynamic transaction fees and partial fee burning through its support for [EIP-1559](https://eips.ethereum.org/EIPS/eip-1559). In contrast, frameworks like Ethermint have faced delays spanning multiple years due to challenges in adapting EVM upgrades.
