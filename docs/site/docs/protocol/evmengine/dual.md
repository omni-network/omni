---
sidebar_position: 1
---

# Dual Client Model

Omni introduces Integrated Consensus, a consensus framework that allows validators to run consensus for the Omni EVM in parallel with consensus for cross-network messages without compromising on performance.

Omni’s Integrated Consensus contains two sub-processes:

- validating state changes within the Omni EVM and
- attesting to `XBlock` hashes originating from external rollup VMs.

## Integrated Consensus

The aggregate Integrated Consensus process is visualized below.

<figure>
  <img src="/img/integrated-consensus.png" alt="Integrated Consensus" />
  <figcaption>*Integrated Consensus: Validating Omni EVM State Changes and Rollup `XBlock` Hash Attestation*</figcaption>
</figure>

1. Every `halo` client runs a node for each rollup VM to check for `XMsg` events emitted from Portal contracts.
2. For every rollup VM block that contains `XMsg`s, `halo` clients build `XBlock`s that contain the corresponding `XMsg`s.
3. Once the calldata for a rollup VM block has been posted and finalized on Ethereum, Omni validators use ABCI++ vote extensions to attest to the hash of the corresponding `XBlock`. These attestations are appended to the current consensus layer block.
