---
sidebar_position: 2
---

# Ethereum's Roadmap

As the Ethereum ecosystem developed, we learned that scaling transaction throughput and decreasing costs would be extraordinarily difficult on Layer 1 Ethereum while maintaining its high security. Over time, various proposals have attempted to solve this problem: state channels, side chains, and plasma to name a few. However, each of these options came with unfortunate trade-offs that typically sacrificed security to achieve these goals. But their development was not in vain! They led to the growth of rollups, the currently agreed upon optimal solution for scaling Ethereum.

## Rollup-Centric Roadmap

Hence, the [rollup-centric roadmap](https://ethereum-magicians.org/t/a-rollup-centric-ethereum-roadmap/4698) was proposed.

Today, Ethereum Layer 1 handles Data Availability, Execution, and Consensus. But over time, Ethereum will primarily become a Data Availability layer. Rollups will handle execution off-chain, and utilize either Ethereum's smart contract layer, or their own sovereign settlement layer, for consensus. This allows rollups (Layer 2s) to borrow Ethereum security while scaling throughput and decreasing cost. Upgrades like [EIP-4844](https://www.eip4844.com/) will even further improve Ethereum as a Data Availability layer by providing the service at lower cost.

You can view the growth of rollup ecosystems at [L2Beat.com](https://l2beat.com/), a great resource for understanding the security trade-offs, current state of development, and adoption levels for rollups.
