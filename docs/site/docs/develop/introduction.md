---
sidebar_position: 1
---

# Introduction

You can use Omni to call a contract on another chain. We call this an `xcall`.

```solidity
omni.xcall(
   destChainId,  // desintation chain id, e.g. 1 for Ethereum mainnet
   to,           // contract address on the destination chain
   data,         // abi encoded calldata, ex abi.encodeWithSignature("foo()")
   gasLimit      // (optional) gas limit for the call on the destination chain
)
```

When receiving an `xcall`, you can read its context through the current `xmsg` given your contract implements the `XApp` helper contract.

```solidity
xmsg.sourceChainId   // where this xcall came from
xmsg.sender          // who sent it (msg.sender of the source xcall)
```

`omni` is a reference to the [OmniPortal](./contracts.md#omniportal) contract. One portal is deployed on each supported chain. To make interacting with the portal easier, inherit from the [XApp](./xapp/xapp.md) contract.

:::info Only Finalized Cross-Chain Messages are Relayed

`XMsg` can only be relayed to the destination chain once the transaction it belongs to finalizes on Ethereum.

:::
