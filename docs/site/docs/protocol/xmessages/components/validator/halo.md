---
sidebar_position: 1
---

# `halo`

`halo` is the client implements the server side of the ABCI++ interface and drives the Omni Execution Layer via the Engine API. CometBFT validators attest to source chain blocks containing cross chain messages using CometBFT Vote Extensions.

Omni's consensus layer is established through the Halo consensus client which supports a Delegated Proof of Stake (DPoS) consensus mechanism implemented by the [CometBFT consensus engine](https://docs.cometbft.com/v0.38/) (formerly known as Tendermint). CometBFT stands out as the optimal consensus engine for Omni nodes for three reasons:

1. CometBFT achieves immediate transaction finalization upon block inclusion, eliminating the need for additional confirmations. This feature allows Omni validators to agree on a single global state for all connected networks, updated every second.
2. CometBFT provides support for DPoS consensus mechanisms, allowing re-staked ETH to be delegated directly to individual Omni validators as part of the network’s security model.
3. CometBFT is one of the most robust and widely tested PoS consensus models and is already used in many production blockchain networks securing billions of dollars.

By default, Omni validators ensure the integrity of rollup transactions by awaiting their posting and finalization on the Ethereum L1 before attesting to them. This proactive approach mitigates the risk of reorganizations on both Ethereum L1 and the rollup, enhancing the overall security and reliability of the system.
