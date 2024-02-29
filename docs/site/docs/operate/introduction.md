---
sidebar_position: 1
---

# Introduction

## What is a node operator in Omni?

Operators run Omni node software and play an active role in validating the Omni network.

In all other blockchains, validators can only stake the native token of that chain. With Omni, validators can stake **\$OMNI**, but they can also restake **\$ETH**, via Eigenlayer. Operators can also receive **\$ETH** delegations from other users.

## Registering as an Operator

To register with Omni, operators must:

1. Register as an operator in Eigenlayer

    This registers your Ethereum public key with the Eigenlayer. You can follow Eigenlayer's instructions [here](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation).

2. Register as an operator in the Omni AVS smart contracts.

    This step tells Eigenlayer that you'd like to be an operator specifically for the Omni AVS. Thus, the **\$ETH** that you, and your delegators restaked, will be used to secure Omni. Omni provides a CLI for this.

3. Run the Omni client software

    This is the component that participates in network validation. Our consensus clients track delegations and stake from our AVS contracts.

    Similar to Ethereum, Omni validators run 2 components: our consensus client, `halo`, and an EVM execution client `geth`, `erigon`, `nethermind`, etc.

## How do I become an operator?

If you'd like to become an Omni operator, please reach out to the team.
