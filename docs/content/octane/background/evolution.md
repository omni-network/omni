---
sidebar_position: 3
---

# Evolution

Omni was designed to solve the fragmentation problem in Ethereum's L2 ecosystem. One of the critical desired properties required to solve this problem was and is to introduce _global state_. Global state empowers smart contract developers to build unified applications that benefit from network effects across deployments, and is a core value proposition of Omni. To provide this in a developer friendly way, the natural choice was to introduce an EVM.

In late ‘22 and early ‘23, the Omni team implemented a proof of concept of the Omni blockchain that utilized Ethermint – at the time, the standard for EVM on Cosmos. In fact, we launched with this architecture for our first two testnets.

The team behind Ethermint used tooling that was best in class at the time, but unfortunately its benchmarks were severely limited in the categories we cared about:

- **Scalability & Performance**
  - 5 second block times
  - ~25 TPS when load tested
  - 500ms RPC queries under standard load
  - State store translation between EVM state and Cosmos state, which created a major disk performance bottleneck
  - Consensus logic breaking during periods of high EVM activity, as the Cosmos mempool is not built for EVM scale load
- **EVM Compatibility**
  - It requires an EVM adapter, a set of logic that wraps geth to be compatible with Cosmos, so EVM upgrades would need to go through several layers of implementation before being included
  - It is partially EVM compatible, but because of this adapter it would never be EVM equivalent
  - It only supported one execution client: geth.
- **Long-Term Maintenance and Upgradeability**
  - Any time the EVM is upgraded, it required new adapter modules to be built for the new opcodes
  - infeasibly high state bloat, with Cosmos + EVM state
  - only works with geth, so new high performance EVM clients could never be used

Because this architecture did not meet our requirements for scalability, performance, EVM equivalence, or long term maintenance, the Omni team decided not to proceed with this architecture. How could we solve these problems and what were our goals for the re-architecture?
