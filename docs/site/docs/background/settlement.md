---
sidebar_position: 3
id: settlement
---

# Rollup Settlement

Rollups come in two main flavors: optimistic rollups and ZK rollups. They have different "settlement" mechanisms for finalizing the state of their rollup.

Optimistic rollups operate with fault proofs. A privileged actor can submit state updates to the rollup's settlement contract on Ethereum Layer 1. Over a period of ~7 days, anyone can dispute the validity of this state update by submitting a proof that the update is invalid. As long as there is 1 honest actor watching this rollup, invalid state updates will be challenged and thrown out by the end of the 7 day window.

For ZK (zero knowledge) rollups, settlement works differently. Instead of the 7 day challenge window, a "prover" runs a sophisticated computation to generate a validity proof. This proof is submitted along with the state updates and provides a mathematical guarantee that the state update is valid.

Both systems come with trade-offs. The 7-day optimistic window introduces friction for users who wish to withdraw their assets from the rollup back to Layer 1. ZK rollups rely on complex math and heavy computation which will make decentralizing them over time more difficult.

Omni introduces a unified settlement system across both rollup categories and decreases settlement time to establish a faster interoperability standard. You can read on in the Protocol section for more details.
