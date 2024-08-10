---
sidebar_position: 3
---

# ABCI++

Cosmos’ [ABCI++](https://docs.cometbft.com/v0.37/spec/abci/) is a supplementary tool that allows `halo` clients to adapt the CometBFT consensus engine to the node’s architecture. An ABCI++ adapter is wrapped around the CometBFT consensus engine to convert messages from the Engine API into a format that can be used in CometBFT consensus. These messages are inserted into CometBFT blocks as single transactions – this makes Omni consensus lightweight and enables Omni’s sub-second finality time.

During consensus, validators also use ABCI++ to attest to the state of external Rollup VMs. Omni validators run the state transition function, $f(i, S_n)$, for each Rollup VM and compute the output, $S_{n+1}$.
