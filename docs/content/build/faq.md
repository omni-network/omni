---
sidebar_position: 4
---

# FAQ

### Which chains does Omni support?

**Mainnet**: Ethereum, Omni, Arbitrum One, Optimism, Base

**Omega testnet**: Ethereum Holesky, Omni Omega, Arbitrum Sepolia, Optimism Sepolia, Base Sepolia

More chains will be added overtime.

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

### What are Confirmation Levels?

Omni's xchain message protocol currently offers 2 confirmation strategies. You can specify your confirmation strategy with each xcall.

You may choose to use the same confirmation strategy for every `xcall` in your dapp. If this is the case, you can just hardcode a confirmation strategy.

Or, you may want to adjust the `xcall`'s confirmation strategy based on the transaction's associated "value". For example, you may choose to implement logic that says:

- `if a user is depositing > 1000 USDC, use "finalized"; else use "latest"`

Here are the details of each type:

#### `finalized`

Finalized xmsgs are attested to and delivered only after the rollup's transaction data containing this xmsg finalizes on Ethereum Layer 1. This requires 2 beacon chain epochs, which typically takes about 12 minutes. However, this strategy offer strong benefits around delivery guarantees – a delivered message can only be "reorg'd out" if Ethereum itself reorgs, which is highly unlikely and requires 2/3 of Ethereum's validators to be slashed.

Summary:

- exactly once delivery guarantees
- ~12 minute delivery

#### `latest`

Latest xmsgs are attested to and delivered as soon as the transaction with the xmsg is included by the L2 sequencer in a block. This provides a much lower latency for message delivery – roughly 5-10s. However, it does come with an associated risk: the xmsg has a higher risk of being reorg'd out if the L2 sequencer misbehaves or fails. This may result in unintended consequences, and you should decide how much you're willing to trust L2 sequencers.

- an xmsg is delivered on a destination chain, but the txn that initiated the xmsg gets reorged out of the source chain
- an xmsg is included in a txn on a source chain, but not delivered on the destination
- an xmsg is included in a txn on a source chain, and delivered multiple times on the destination

Summary

- low latency: ~5-10s
- no delivery guarantees

Note that these risks are possible with all xchain messaging protocols.

### Is Omni a bridging protocol?

Please note that Omni does not bridge native tokens.

The `msg.value` in an `xcall` is for fees, not for bridging native tokens.

Omni's xchain messaging system allows contracts to send and receive function calls across chains. It is a simple and secure way to interact with contracts on other chains.

However, Omni is not a bridge for canonical or native tokens. If you believe you need a bridge, we invite you to consider building a "chain-agnostic", rather than a cross-chain, application. Chat with the team if you're curious about this mental model! We believe the future of smart contract development should not force users to bridge – it should meet users where they are!

### Does the Omni EVM support pre-EIP-155 transactions?

Our public RPCs do not support pre-EIP-155 transactions. If you'd like to submit a pre-155 transaction, you can run a full node (see [relevant page](../operate/run-full-node.md)), enable [AllowUnprotectedTxs](https://github.com/ethereum/go-ethereum/blob/e67d5f8c441d5c85cfa69317db8d85794645af14/node/config.go#L199), and submit the transaction to your own endpoint.
