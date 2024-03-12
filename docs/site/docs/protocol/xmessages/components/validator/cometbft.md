---
sidebar_position: 2
---

# CometBFT

## `halo`

`halo` is the client implements the server side of the ABCI++ interface and drives the Omni Execution Layer via the Engine API. CometBFT validators attest to source chain blocks containing cross chain messages using CometBFT Vote Extensions.

## CometBFT Consensus Engine

Omni's consensus layer is established through the Halo consensus client which supports a Delegated Proof of Stake (DPoS) consensus mechanism implemented by the [CometBFT consensus engine](https://docs.cometbft.com/v0.38/) (formerly known as Tendermint). CometBFT stands out as the optimal consensus engine for Omni nodes for three reasons:

1. CometBFT achieves immediate transaction finalization upon block inclusion, eliminating the need for additional confirmations. This feature allows Omni validators to agree on a single global state for all connected networks, updated every second.
2. CometBFT provides support for DPoS consensus mechanisms, allowing re-staked ETH to be delegated directly to individual Omni validators as part of the network’s security model.
3. CometBFT is one of the most robust and widely tested PoS consensus models and is already used in many production blockchain networks securing billions of dollars.

By default, Omni validators ensure the integrity of rollup transactions by awaiting their posting and finalization on the Ethereum before attesting to them. This proactive approach mitigates the risk of reorganizations on both Ethereum and the rollup, enhancing the overall security and reliability of the system.

### Consensus Process

Omni Consensus clients process the consensus blocks and maintain consensus state that tracks the “status” of each `XBlock` (by the validator set) from `Pending` to `Approved` (including an `AggregateAttestation`). Then only the “latest approved” `XBlock` for each source chain needs to be maintained; earlier `XBlock`s can be trimmed from the state. Validators in the current validator set must attest to all subsequent (after the “last approved”) `XBlock`.

#### Validator Set Changes

When the validator set changes, all `XBlock`s marked as `Pending` need to be updated by:

- Updating the associated validator set ID to the current.
- Deleting all attestations by validators not in the current set.
- Updating the weights of each remaining attestation according to the new validator set.

Validators that already attest to the `Pending` marked `XBlock` during the previous validator set, do not need to re-attest. Only the new set validators must attest (ie. to all `XBlock`s after the latest approved).

## Execution Consensus

1. When it is a validator's turn to propose a block, its `halo` client requests the latest Omni EVM block from its execution client using the Engine API.
2. The execution client builds a block from the transactions in its mempool and returns the block header to the `halo` client through the Engine API.
3. The `halo` client packages the new block proposal as a single CometBFT transaction and includes it in the consensus layer block.
4. The block is proposed to the rest of the validator network through the consensus layer’s P2P network.
5. Non-proposing validators use the Engine API and their execution clients to run the state transition function on the proposed block header to verify the block’s validity.
