---
sidebar_position: 4
---

# Goals

- **Decouple execution from consensus**
  - Multiple execution layer implementations are available with new ones being added constantly. With a decoupled architecture, the latest and greatest implementation can always be introduced.
  **- The Cosmos mempool and state store are not built for EVM transactions, we want to separate consensus processing from execution processing #modularity
- **EVM equivalence**
  - We wanted our implementation to be no different than the code running on Ethereum L1, so that 100% of developer tooling is compatible with Omni.
- **Performance**
  - As an async messaging protocol, fast block times are critical, ideally targeting subsecond finality
  - To introduce a messaging protocol that scales for the entire Ethereum ecosystem, Omni's TPS and GPS must not hit fundamental bottlenecks like the previous implementation

After many months of research, proof-of-concepts, and benchmarking, we we able to design an architecture that achieve all of these goals – unlocked by the release of the EngineAPI (Ethereum Proof-of-Stake) and ABCI 2.0 (CometBFT).

This architecture was not possible before 2024.

_Enter Octane._
