---
sidebar_position: 1
id: introduction
---

# Introduction

Omni implements a novel protocol architecture that establishes a new precedent for secure, performant, and globally compatible interoperability across the Ethereum ecosystem. This section breaks down the mechanisms and design choices that underpin the protocol.

## Omni's Dual Role

The Omni chain is engineered to perform two critical functions within the blockchain ecosystem:

### Integrated Consensus

#### Cross-Rollup Message Consensus

Omni serves as a bridge between various rollups, enabling them to communicate seamlessly. This functionality is vital for maintaining coherence and interoperability in the increasingly fragmented blockchain landscape.

#### Omni EVM Operation Consensus

The Omni EVM is a parallel execution environment that operates under the same consensus umbrella as the cross-rollup messages. It empowers developers to build and deploy decentralized applications that can interact with different blockchain networks, all within the Omni ecosystem.

## Security

Previous attempts to secure externally verified interoperability networks relied on native assets for cryptoeconomic security. Under this approach, the scale of the protocol’s security is directly tied to the demand for the protcol’s native asset, introducing reflexive dynamics that result in unstable security guarantees.

Omni improves upon existing approaches by using EigenLayer’s restaking, a novel primitive that extends Ethereum L1’s cryptoeconomic security to external networks by reusing staked **\$ETH** from Ethereum’s consensus layer. As a highly liquid, low volatility asset, restaked **\$ETH** allows Omni to achieve significantly greater stability than its predecessors. Additionally, by deriving security from Ethereum, Omni aligns its security base with the rollups it connects, facilitating a security model that grows in tandem with Ethereum’s modular ecosystem.

Omni reinforces its security with a dual staking model. Under this system, staked **\$OMNI** provides additional security alongside restaked **\$ETH**. The protocol dynamically incentivizes validators to contribute more security from either restaked **\$ETH** or **\$OMNI** depending on market conditions.

### Security Model

The security of the Omni chain is upheld by a dual-staking mechanism:

- **Restaked \$ETH**: Omni leverages the existing security of Ethereum by allowing participants to restake their Ethereum holdings, contributing to the Omni network's overall security.
- **Staked \$OMNI**: Alongside **\$ETH**, Omni's native token, **\$OMNI**, is staked as a commitment to the network's integrity, aligning incentives and enhancing security.

### Simplified Yet Powerful

While Omni's infrastructure is complex and performant, our goal is to present the information in a digestible format. Throughout the Protocol section, you will find detailed yet accessible content that breaks down how Omni achieves its vision of a universally connected and secure blockchain ecosystem.

In the upcoming pages, we will delve deeper into the specifics of Omni's consensus mechanisms, the architecture of the Omni EVM, the role of staked assets, and much more. Whether you are a developer, researcher, or simply a blockchain enthusiast, these pages will equip you with a thorough understanding of the Omni Protocol.
