---
slug: /
---

# Overview

## Why does Omni exist?

The most glaring side effect of Ethereum’s pursuit of the rollup-centric roadmap is fragmentation of users and capital across the broader Ethereum ecosystem. This has degraded both the developer and user experiences on Ethereum.

For developers, they remain siloed and unable to access a majority of the users. Due to the complexity of building an application deployed across multiple rollups, a majority of teams either deploy to one rollup or duplicate deployments of their application across multiple rollups.

For users, the experience is painful. They must navigate application frontends, select networks, switch RPC endpoints, and manually bridge assets between networks, all while ensuring they have enough gas to complete transactions on the destination network.

## What is Omni?

Omni is a chain abstraction layer for the Ethereum ecosystem, allowing developers to make their application available to users and liquidity across all rollups, while only deploying their application on a single rollup. Omni is making the Ethereum ecosystem feel like a single chain for users — no manual bridging, switching RPCs, managing gas, or wallet upgrades are needed to interact with applications powered by Omni. To simplify both user and developer experiences, Omni delegates cross-chain complexity to sophisticated third parties through Omni’s solver network.

## How does Omni work?

The foundation of the Omni Network is composed of two core primitives:

1. An interoperability protocol purpose-built for Ethereum’s rollup ecosystem
2. An EVM that supports computation for Omni’s interoperability protocol

On top of this foundation sits the Omni SDK and the Omni solver network. The Omni SDK is the primary interface for applications to interact with the Omni solver network, which carries out cross-chain actions (intents) on behalf of application users on destination chains.

<img src="/img/architecture.jpg" width="500px"/>

Here’s a quick breakdown of how these components work together:

Omni’s solver network has timelock escrow contracts on every supported rollup — these allow users to deposit funds and declare intents for actions they want completed on applications deployed on other rollups.

After a user deposits funds into escrow and declares their intent on the source rollup, a solver provides “just in time liquidity” on the destination rollup by depositing capital into the target application on behalf of the user.

The solver then provides proof of completing the user’s intent through Omni’s interoperability network, allowing the escrow contract on the origin rollup to release the user’s deposited funds to the solver.
