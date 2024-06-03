---
sidebar_position: 5
---

import GitHubCodeBlock from '@site/src/components/GitHubCodeBlock/GitHubCodeBlock';

# Fees

Omni charges fees for each `xcall`. These fees are paid synchronously on the source chain, in the chain's native token. For most rollups, this is ETH.

## Fee Calculation

Fees are determined by an `xcall`'s destination chain, calldata, and gas limit. You can calculate the fee for an `xcall` via the `XApp.feeFor(...)` method.

```solidity
uint256 fee = feeFor(
   destChainId,  // destination chain id
   data,         // abi encoded calldata, ex abi.encodeWithSignature("foo()")
   gasLimit      // gas limit to enforce on execution
)
```

Or, using the default gas limit.
```solidity

uint256 fee = feeFor(
   destChainId,  // destination chain id
   data,         // abi encoded calldata, ex abi.encodeWithSignature("foo()")
)
```

## Fee Payment

`XApp` handles calculating and charging fees when making an `xcall`

<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/059303647e07fc3481e379b710922e2b84b1827f/contracts/src/pkg/XApp.sol#L56-L65" />

`xcall(...)` charges fees to your contract by default. To charge users for fees, calculate the fee with `feeFor(...)`, and verify `msg.value` is sufficient.

```solidity
uint256 fee = feeFor(...)
require(msg.value >= fee, "insufficient fee")
```

You can calculate this fee off chain, and require users send sufficient `xcall` fees with each contract call.

## Example

In the case of our [`XGreeter` example](./example.md), the fee may be different for each greeting, because the length of the greeting message is variable. You can calculate this fee offchain by querying the portal directly (via `OmniPortal.feeFor(...)`). Or, you can introduce a view function on your contract that calculates the fee.


```solidity
function xgreetFee(uint64 destChainId,  string calldata greeting) external view {
    feeFor(
        destChainId,
        abi.encodeWithSignature("greet(string)", greeting)
    );
}
```

For this simple example, this view function is not that helpful. But for `xcalls` with calldata that depends on other contract state, a view function like this can be very helpful.
