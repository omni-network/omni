---
sidebar_position: 3
---

import GitHubCodeBlock from '@site/src/components/GitHubCodeBlock/GitHubCodeBlock';

# Example

Here's an example of a simple cross chain contract, `XGreeter`. This contract lets you send greetings from one chain to another.

## `XGreeter` Contract

<GitHubCodeBlock url="https://github.com/omni-network/omni-forge-template/blob/main/src/XGreeter.sol" />

## Walkthrough

Let's walk through this step by step.

First, inherit from `XApp`.


```solidity
contract XGreeter is XApp {
    constructor(address portal) XApp(portal) { }

    // ...
}
```

### Perform a Cross Chain Call

To call a contract on another chain, use `xcall`.

```solidity
function xgreet(uint64 destChainId, address to, string calldata greeting) external {
    xcall(
        // params for xcall
    );
}
```

### Receive a Cross Chain Call

When receiving an `xcall`, you can read its context via `omni.xmsg()`.

```solidity
xmsg.sourceChainId // where this xcall came from
xmsg.sender        // who sent it
```

With this context, we can have our `XGreeter` emit events detailing the source chain and sender.

```solidity
function greet(string calldata greeting) external {
    emit Greetings(omni.xmsg().sender, omni.xmsg().sourceChainId, greeting);
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


```solidity
function greet(string calldata greeting) external xrecv {
    emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
}
```

### Checking for Cross Chain Calls

Note that not every call is an `xcall`. In these cases, `xmsg` will be its zero value.

```solidity
xmsg.sourceChainId // 0
xmsg.sender        // address(0)
```

You can check if the current call is an `xcall` with `isXCall`.

```solidity
function isXCall() internal view returns (bool) {
    return omni.isXCall() && msg.sender == address(omni);
}
```

Note that not only does `isXCall` check with the portal that the current transaction is an `xcall`, it also confirms the sender is the portal itself. This helps avoid mistaking calls later in an `xcall` stacktrace with the original `xcall`. Using this helper, we can ensure that `greet()` can only ever be called via an `xcall`.

```solidity
function greet(string calldata greeting) external xrecv  {
    require(isXCall(), "XGreeter: only xcall");
    emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
}
```
