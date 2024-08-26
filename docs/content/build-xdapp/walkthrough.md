---
sidebar_position: 2
---

import GitHubCodeBlock from '@site/src/components/GitHubCodeBlock/GitHubCodeBlock';

# XDapp walkthrough

Let's see what's happening under the hood. This contract lets you send greetings from a rollup chain to a global storage contract deployed on Omni. Two main contracts are used in this example:

1. `Greeter` - A contract deployed on a rollup that sends greetings to the Omni EVM.
2. `GreetingBook` - A contract deployed on the Omni EVM that stores greetings from all supported chains.

## `Greeter` Contract

<GitHubCodeBlock url="https://github.com/omni-network/hello-world-template/blob/48ff2f5277b4c144802c1ffa894a03ac071f02fc/src/Greeter.sol" />

Let's walk through this step by step.

First, inherit from `XApp`.


```solidity
contract Greeter is XApp {
    // ...
}
```

You may also specify the confirmation level for the cross-chain message. This is the level of finalisation of the transaction containing the message you want to wait for. In this example, we set it to `Latest`, meaning the message is relayed on block creation at source. You can see the supported confirmation levels [here](https://github.com/omni-network/omni/blob/main/contracts/core/src/libraries/ConfLevel.sol).

```solidity
constructor(address portal, address _greetingBook) XApp(portal, ConfLevel.Latest) {
        greetingBook = _greetingBook;
    }
```

### Perform a Cross Chain Call

To call a contract on another chain, use `xcall`.

```solidity
function greet(string calldata greeting) external payable {
    uint256 fee = xcall(
        // params for xcall
    );
}

```

## `GreetingBook` Contract

<GitHubCodeBlock url="https://github.com/omni-network/hello-world-template/blob/eb02c55bc8ef92c09e7cb6e40420353e41e2841c/src/GreetingBook.sol" />

Similar to `Greeter`, we inherit from `XApp`.

```solidity
contract GreetingBook is XApp {
    // ...
}
```

Similiar to `Greeter`, we can specify the confirmation level for the cross-chain message.

```solidity
constructor(address portal) XApp(portal, ConfLevel.Latest) {
}
```

### Receive a Cross Chain Call

When receiving an `xcall`, you can read its context via `omni.xmsg()`, which is shortened by `XAapp` to `xmsg` for convenience.

```solidity
xmsg.sourceChainId // where this xcall came from
xmsg.sender        // who sent it
```

With this context, we can have our `GreetingBook` extract the source chain and sender of the `xcall` to store the greeting in a struct.

```solidity
function greet(address user, string calldata _greeting) external xrecv {
        // ...

        lastGreet = Greeting(user, _greeting, xmsg.sourceChainId, block.timestamp);
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

You can check if the current call is an `xcall`, and the sender is the portal, with `isXCall()`.

```solidity
function greet(address user, string calldata _greeting) external xrecv {
        require(isXCall(), "GreetingBook: only xcalls");

        // ...
    }
```
