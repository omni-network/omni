---
sidebar_position: 1
---

# Components

## Components & Communication

Omni nodes are configured with a new modular framework based on [Ethereum’s Engine API](https://github.com/ethereum/execution-apis/tree/4b225e0d273e92982b2c539d63eaaa756c5285a4/src/engine) to combine `halo` with the EVM execution client.

<figure>
  <img src="/img/validator.png" alt="Validator" />
  <figcaption>*Overview of the Components Run by Validators and Interactivity*</figcaption>
</figure>

Using the Engine API, Omni nodes pair existing high performance Ethereum execution clients with a new consensus client, referred to as `halo`, that implements CometBFT consensus.The Engine API allows clients to be substituted or upgraded without breaking the system. This allows the protocol to maintain flexibility on Ethereum and Cosmos technology while promoting client diversity within its execution layer and consensus layer. We consider this new network framework to be a public good that future projects may leverage for their own network designs.
