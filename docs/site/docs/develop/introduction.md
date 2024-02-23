---
sidebar_position: 1
---

# Introduction

You can use the Omni utility interface contract to interact with the Omni network by calling `omni.xcall(...)` to send cross-rollup transactions.

This guide will walk you through an overview of what is required to build a dApp that can run on multiple rollups.

## Usage Guidelines

### Calls to Another Rollup

To send a transaction to another rollup, use:

```solidity
omni.xcall(
   uint64 destChainId,   // desintation chain id
   address to,           // contract address on the dest rollup to execute the call on
   bytes memory data     // calldata for the transaction, abi encoded
)
```

### Calls from Another Rollup

The destination rollup contract will be called by the rollup deployed [Portal Contract](../use/protocol.md#portal-contract) when it's an Onmi cross-rollup call. These `external` functions can then specifically filter for contract calls from the Portal `address` and use the `omni.isXCall()` helper, as:

```solidity
function receive() external {
    if (omni.isXCall() && msg.sender == address(omni)) {  // assert called from Portal Contract
        // perform some action
    }
}
```

## Fees

Fees are paid in ETH to the Portal Contract and calculated in real-time during transactions via the payable `xcall` function. Fees may be estimated off-chain by querying the Portal Contract `feeFor` method.

### Examples

Below are some examples of usage of the Omni utility interface contract for the handling of fees. For all these examples, it is assumed that the contract imports this utility contract and initialises it as:

<details>
<summary>Initialisation Code in Solidity Contract</summary>

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol"; // interface utility contract
import { XTypes } from "src/libraries/XTypes.sol"; // types utility contract

contract MyContract {
    IOmniPortal public omni;

    constructor(IOmniPortal _omni) {
        omni = _omni; // initialise utility contract
    }

    // ... some additional contract logic
}
```

</details>

#### Paying Fee with Contract Funds

Fee payment using funds from the calling contract and estimating fees with the `feeFor` method, looks as follows:

```solidity
  function xincrement(uint64 destChainId) {
    // encode tx data
    bytes memory data = abi.encodeWithSignature("increment()");

    // calculate fee
    uint256 fee = omni.feeFor(destChainId, data);

    // send tx to sibling contract on dest chain
    omni.xcall{value: fee}(destChainId, siblings[destChainId], data);
  }
```

#### Paying Fee with a Custom Gas Limit

Or, with a custom gas limit:

```solidity
 function xincrement(uint64 destChainId) {
    // encode tx data
    bytes memory data = abi.encodeWithSignature("increment()");

    // custom gas limit
    uint64 gasLimit = 100_000;

    // calculate fee
    uint256 fee = omni.feeFor(destChainId, data, gasLimit);

    // send tx to sibling contract on dest chain
    omni.xcall{value: fee}(destChainId, siblings[destChainId], data, gasLimit);
  }
```

The two above examples charge the calling contract. It therefore must maintain an ETH balance to continue sending Omni transactions. Developers may want to pass Omni fees onto their users.

#### Charging End-Users

Charging fees in ETH makes it easy to charge end users for transactions that use Omni. methods must be made `payable`, calculate fees off-chain when building transactions (on a frontend, for example), and send the required ETH in the transaction.

An example of a function allowing end users to pay for the transaction with `payable` declaration:

```solidity
  function xincrement(uint64 destChainId) payable {
    // fee estimated off chain by call to portal, paid with tx
    omni.xcall{value: msg.value}(
      destChainId,
      siblings[destChainId],
      abi.encodeWithSignature("increment()")
    )
  }
```
