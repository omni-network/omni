---
sidebar_position: 4
---

# FAQ

### Which chains does Omni support?

Mainnet v1 launch: Ethereum, Omni, Arbitrum, Optimism, Base

Omega Testnet: Ethereum Holesky, Omni Omega, Arbitrum Sepolia, Optimism Sepolia, Base Sepolia

### How fast is a message delivered?

Rollup to rollup messages are roughly 5-10s for "latest" messages and 12 minutes for "finalized" messages.

### How should I use finalized vs. latest?

It depends on the use case, but we generally recommend using `finalized` for high value transactions where exactly once delivery semantics is critical. `latest` can be used for smaller magnitude transactions where delivery guarantees can be loosened.

### How much does an xcall cost?

The cost of an `xcall` is a function of:

- destination chain: since an xcall is actually a function execution on a remote chain, we must consider the gas prices on that chain.
- calldata size: impacts gas fee on the destination
- gas limit: max gas you're willing to spend on the destination

### If xcall is successfully sent on the source chain but the tx on the destination fails, what will happen?

An xcall will only revert on the destination if there is logic in your smart contract on the destination that reverts. Generally, we recommend not reverting when receiving an xcall, rather reverting on the source before sending the xcall if possible.

### What happens if a rollup reorgs?

Reorgs have no impact on a `finalized` stream, as these streams wait for the rollup to be finalized on Ethereum before validating an XMsg.

### Why doesn’t Omni wait for the 7-day challenge period for rollups?

Rollups are actually finalized once transaction calldata is posted to L1. The 7 Day period is only from the perspective of L1 – because it takes smart contracts on L1 7 Days to confirm the state of a rollup. But anyone can compute what the current state of the rollup is before the 7 Day window by running an node for that rollup. Omni's validators each run L2 nodes for integrated rollups, and attest to the state of the rollup once txn calldata is posted.
