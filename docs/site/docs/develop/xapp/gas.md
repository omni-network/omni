---
sidebar_position: 4
---

# Gas Limits

Contract calls consume gas. Omni enforces a gas limit when executing an `xcall`. You can specify a gas limit when making an `xcall`.

```solidity
xcall(
   destChainId,  // destination chain id
   to,           // contract address on the destination chain
   data,         // abi encoded calldata, ex abi.encodeWithSignature("foo()")
   gasLimit      // gas limit to enforce on destination execution
)
```

If you do not specify a gas limit, Omni enforces a default gas limit, currently set to `200_000`. You can read the default gas limit via the portal contract.

```solidity
/// @notice Default xmsg execution gas limit, enforced on destination chain
omni.XMSG_DEFAULT_GAS_LIMIT()
```

If you do not provide sufficient gas for an `xcall`, its execution will revert. It's important to set appropriate gas limits for each `xcall`. Determine gas limits with proper unit testing.
