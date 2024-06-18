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

If you do not provide sufficient gas for an `xcall`, its execution will revert. It's important to set appropriate gas limits for each `xcall`. Determine gas limits with testing.
