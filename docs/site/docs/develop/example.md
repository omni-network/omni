---
sidebar_position: 3
---

# Example Application

## Example

This is a simple example of a cross-rollup application that uses `omni.xcall(...)`. It tracks a count, that can be incremented from multiple rollups.

### Summary

- Contract `Counter` interfaces with the Omni network through `IOmniPortal`.
- Write methods
    - Function `increment()` increments counter by 2 when called via Omni's cross-chain (`isXCall`) and otherwise increments by 1.
    - Function `xincrement()` Initiates a cross-chain increment, encoding the increment function call and requires payment of a calculated fee to the Omni portal for the transaction.
        - Calls `omni.xcall` to send the transaction to the specified destination chain.
- View methods
    - Function `getCount(uint64 chainId)` returns the count for a specified chain ID.
    - Function `getCount()` returns the count for the current chain ID.

### Solidity Implementation

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";

contract Counter {
    IOmniPortal public omni;

    constructor(IOmniPortal _omni) {
        omni = _omni;
    }

    /// @dev increment this counter
    function increment() external {
        if (omni.isXCall() && msg.sender == address(omni)) {
            // this is a call from another rollup
            count = count + 2;
        }
        count++;
    }

    /// @dev increment a counter on another chain
    function xincrement(uint64 destChainId, address counter) external payable {
        bytes memory data = abi.encodeWithSelector(this.increment.selector);
        uint256 fee = omni.feeFor(destChainId, data);

		// omit if we don't want to charge the user
        require(msg.value >= fee, "Counter: insufficient fee");

        omni.xcall{ value: fee }(destChainId, counter, data);
    }

    /// @dev get the count for a specific chain
    function getCount(uint64 chainId) external view returns (uint256) {
        return countByChainId[chainId];
    }

    /// @dev get the count for the current chain
    function getCount() external view returns (uint256) {
        return countByChainId[omni.chainId()];
    }
}
```
