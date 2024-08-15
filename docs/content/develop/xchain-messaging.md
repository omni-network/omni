---
sidebar_position: 2
---

# XChain Messaging

You can use Omni to call a contract on another chain. We call this an `xcall`.

```solidity
omni.xcall(
   destChainId,  // destination chain id, e.g. 1 for Ethereum mainnet
   conf,         // confirmation strategy, e.g. 1 for `latest` and 4 for `finalized`
   to,           // contract address on the destination chain
   data,         // abi encoded calldata, ex abi.encodeWithSignature("foo()")
   gasLimit      // gas limit for the call on the destination chain
)
```

When receiving an `xcall`, you can read its context through the current `xmsg` given your contract implements the `XApp` helper contract.

```solidity
xmsg.sourceChainId   // where this xcall came from
xmsg.sender          // who sent it (msg.sender of the source xcall)
```

`omni` is a reference to the [OmniPortal](./contracts.md#omniportal) contract. One portal is deployed on each supported chain. To make interacting with the portal easier, inherit from the [XApp](./xapp/xapp.md) contract.

## Confirmation Levels

Omni's xchain message protocol currently offers 2 confirmation strategies. You can specify your confirmation strategy with each xcall.

You may want to use the same confirmation strategy for every `xcall` in your dapp. If this is the case, you can just hardcode a confirmation strategy.

Or, you may want to adjust the `xcall`'s confirmation strategy based on the transaction's associated "value". For example, you may choose to implement logic that says:

- `if a user is depositing > 1000 USDC, use "finalized"; else use "latest"`

Here are the details of each type:

### `finalized`

Finalized xmsgs are attested to and delivered only after the rollup's transaction data containing this xmsg finalizes on Ethereum Layer 1. This requires 2 beacon chain epochs, which typically takes about 12 minutes. However, this strategy offers strong benefits around delivery guarantees – a delivered message can only be "reorg'd out" if Ethereum itself reorgs, which is highly unlikely and requires 2/3 of Ethereum's validators to be slashed.

Summary:

- exactly once delivery guarantees
- ~12 minute delivery

### `latest`

Latest xmsgs are attested to and delivered as soon as the transaction with the xmsg is included by the L2 sequencer in a block. This provides a much lower latency for message delivery – roughly 5-10s. However, it does come with an associated risk: the xmsg has a higher risk of being reorg'd out if the L2 sequencer misbehaves or fails. This may result in unintended consequences, and you should decide how much you're willing to trust L2 sequencers.

- an xmsg is delivered on a destination chain, but the txn that initiated the xmsg gets reorged out of the source chain
- an xmsg is included in a txn on a source chain, but not delivered on the destination
- an xmsg is included in a txn on a source chain, and delivered multiple times on the destination

Summary

- low latency: ~5-10s
- no delivery guarantees

Note that these risks are possible with all xchain messaging protocols.

## Bridging

Please note that Omni does not bridge native tokens.

The `msg.value` in an `xcall` is for fees, not for bridging native tokens.

Omni's xchain messaging system allows contracts to send and receive function calls across chains. It is a simple and secure way to interact with contracts on other chains.

However, Omni is not a bridge for canonical or native tokens. If you believe you need a bridge, we invite you to consider building a "chain-agnostic", rather than a cross-chain, application. Chat with the team if you're curious about this mental model! We believe the future of smart contract development should not force users to bridge – it should meet users where they are!
