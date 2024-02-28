---
sidebar_position: 1
---

# Introduction

You can use Omni to call a contract on another chain. We call this an `xcall`.

```solidity
omni.xcall(
   destChainId,  // desintation chain id
   to,           // contract address on the destination chain
   data          // abi encoded calldata, ex abi.encodeWithSignature("foo()")
)
```

When receiving an `xcall`, you can read its context through the current `xmsg`.

```solidity
omni.xmsg().sourceChainId // where this xcall came from
omni.xmsg().sender        // who sent it (msg.sender of the source xcall)
```

`omni` is a reference to the [OmniPortal](./contracts.md#omniportal) contract. One portal is deployed on each supported chain. To make interacting with the portal easier, inherit from the [XApp](./xapp/xapp.md) contract.
