---
sidebar_position: 3
sidebar_label: Example
---

import GitHubCodeBlock from '@site/src/components/GitHubCodeBlock/GitHubCodeBlock';

# Hello World Contracts

Here's an example of a simple cross-chain contract set for setting and getting a string. This contract lets you send greetings from a rollup chain to a global storage contract deployed on Omni. Two main contracts are used in this example:

1. `RollupGreeter` - A contract deployed on a rollup chain that sends greetings to the Omni chain.
2. `GlobalGreeter` - A contract deployed on the Omni chain that stores greetings.

## `RollupGreeter` Contract

<GitHubCodeBlock url="https://github.com/omni-network/hello-world-template/blob/f4002e90960e013be18c0be44d07aadb34264489/src/RollupGreeter.sol" />

### Walkthrough

Let's walk through this step by step.

First, inherit from `XApp`.


```solidity
contract XGreeter is XApp {
    // ...
}
```

You may also specify the confirmation level for the cross-chain message. This is the level of finalisation of the transaction containing the message you want to wait for. In this example, we set it to `Latest`, meaning the message is relayed on block creation at source. You can see the supported confirmation levels [here](https://github.com/omni-network/omni/blob/main/contracts/src/libraries/ConfLevel.sol).

```solidity
constructor(address portal, address _omniChainGreeter) XApp(portal, ConfLevel.Latest) {
    omniChainGreeter = _omniChainGreeter;
}
```

### Perform a Cross Chain Call

To call a contract on another chain, use `xcall`.

```solidity
function greet(string calldata greeting) external payable {
    xcall(
        // params for xcall
    );
}
```

## `GlobalGreeter` Contract

<GitHubCodeBlock url="https://github.com/omni-network/hello-world-template/blob/1d0ba3c882c47284b1b16bc4b02f68e996a1e4a1/src/GlobalGreeter.sol" />

### Walkthrough

Similar to `RollupGreeter`, we inherit from `XApp`.

```solidity
contract GlobalGreeter is XApp {
    // ...
}
```

Similiar to `RollupGreeter`, we can specify the confirmation level for the cross-chain message.

```solidity
constructor(address portal) XApp(portal, ConfLevel.Latest) {
}
```

### Receive a Cross Chain Call

When receiving an `xcall`, you can read its context via `omni.xmsg()`.

```solidity
xmsg.sourceChainId // where this xcall came from
xmsg.sender        // who sent it
```

With this context, we can have our `XGreeter` extract the source chain and sender of the `xcall` to store the greeting in a struct.

```solidity
function greet(string calldata _greeting) external xrecv {
    // Initialize the fee to 0, for local calls
    uint256 fee = 0;
    if (isXCall() && xmsg.sourceChainId != omni.chainId()) {
        // Calculate the fee for the cross-chain call
        fee = feeFor(xmsg.sourceChainId, abi.encodeWithSelector(this.greet.selector, _greeting), DEST_TX_GAS_LIMIT);
    }

    // Create a Greeting struct to store information about the received greeting
    Greeting memory greeting =
        Greeting(xmsg.sourceChainId, block.timestamp, fee, msg.sender, xmsg.sender, _greeting);

    // Update the lastGreet variable with the information about the received greeting
    lastGreet = greeting;
}
```

For convenience, `XApp` defines the `xrecv` modifier. This modifier reads the current xmsg into storage, and deletes after its function's execution.

```solidity
modifier xrecv() {
    xmsg = omni.xmsg();
    _;
    delete xmsg;
}
```

It also visually marks a function as the target of an `xcall`. Though, the `xrecv` modifier is not required to receive an `xcall`. Using this modifier, we can simplify our `XGreeter` a bit further.

### Checking for Cross Chain Calls

Note that not every call is an `xcall`. In these cases, `xmsg` will be its zero value.

```solidity
xmsg.sourceChainId // 0
xmsg.sender        // address(0)
```

You can check if the current call is an `xcall` with `isXCall`.

<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/ad2f5b7dddc245e7f5b6b662d6c1fc44170694ab/contracts/src/xchain/OmniPortal.sol#L200-L202" />

Note that not only does `isXCall` check with the portal that the current transaction is an `xcall`. This helps avoid mistaking calls later in an `xcall` stacktrace with the original `xcall`. Using this helper, we can ensure that `greet()` can only ever be called via an `xcall`.
