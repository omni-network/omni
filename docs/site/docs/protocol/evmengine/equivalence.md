---
sidebar_position: 5
---

# EVM Equivalence

Post-merge ethereum decoupled the execution layer from the consensus layer introducing a more modular approach to building blockchains. This modular approach allows the EVM to scale (somewhat) independently from consensus, by simply adopting the latest performant execution client like `erigon` or `geth`.

The Omni consensus layer needs smart contracts to manage native staking and delegated re-staking from ETH L1. The Omni EVM is a natural fit as fees would be much lower and syncing with the consensus layer is already built-in. Providing an EVM purposely built for cross-chain dapps that has both low fees and short block times allows for a simple adoption path and hub-and-spoke mental model to onboard projects into Omni Protocol.

The Omni execution layer also benefits from using existing Ethereum execution clients without specialized modifications. This approach eliminates the risk of introducing new bugs and increases release velocity for features strictly related to cross-rollup interoperability. Furthermore, Omni nodes can seamlessly adopt upgrades from any EVM client, ensuring ongoing compatibility with the Omni consensus layer. For example, the Omni EVM natively supports dynamic transaction fees and partial fee burning through its support for [EIP-1559](https://eips.ethereum.org/EIPS/eip-1559). In contrast, frameworks like Ethermint have faced delays spanning multiple years due to challenges in adapting EVM upgrades.
