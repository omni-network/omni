---
sidebar_position: 3
---

# Example

Here's an example of a simple cross chain contract, `XGreeter`. This contract lets you send greetings from one chain to another.

## `XGreeter` Contract

```solidity
/**
 * @title XGreeter
 * @notice A cross chain greeter
 */
contract XGreeter is XApp {
    /// @dev Emitted when someone greets the ether
    event Greetings(address indexed from, uint64 indexed fromChainId, string greeting);

    constructor(address portal) XApp(portal) { }

    /// @dev Greet on another chain
    ///      `xcall` is a method inherited from `XApp`
    function xgreet(uint64 destChainId, address to, string calldata greeting) external {
        xcall(
            destChainId,
            to,
            abi.encodeWithSignature("greet(string)", greeting)
        );
    }

    /// @dev Greet on this chain
    ///      The `xrecv` modifier reads the current xmsg into `xmsg` storage
    function greet(string calldata greeting) external xrecv {
        require(isXCall(), "XGreeter: only xcall");
        emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
    }
}
```

## Walkthrough

Let's walk through this step by step.

### Inheritance and Constructor

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
        destChainId
        to,
        abi.encodeWithSignature("greet(string)", greeting)
    );
}
```

### Decomposing Messages

When receiving an `xcall`, you can read its context via `omni.xmsg()`.

```solidity
omni.xmsg().sourceChainId // where this xcall came from
omni.xmsg().sender        // who sent it
```

### Emit Events

With this context, we can have our `XGreeter` emit events detailing the source chain and sender.

```solidity
function greet(string calldata greeting) external {
    emit Greetings(omni.xmsg().sender, omni.xmsg().sourceChainId, greeting);
}
```
### Receiving Calls

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

Note that not only does `isXCall` check with the portal that the current transaction is an `xcall`, it also confirms the sender is the portal itself. This helps avoid mistaking calls later in an `xcall` stacktrace with the original `xcall`. Using this helper, we can ensure that `greet()` can only every be called via an `xcall`.

```solidity
function greet(string calldata greeting) external xrecv  {
    require(isXCall(), "XGreeter: only xcall");
    emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
}
```
