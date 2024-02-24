---
sidebar_position: 2
---

# XApp

The [`XApp`](contracts#xapp) contract is the base contract from which Omni cross-chain applications should inherit. It simplifies sending / receiving cross chain calls.

<br />

## Installation

`XApp` is a contract maintained in Omni's [monorepo](https://github.com/omni-network/omni/tree/main/contracts), alongside the rest of Omni's smart contracts. To use it, install Omni's smart contracts in your project.

#### npm

```bash
npm install @omni-network/contracts
```

```solidity
import { XApp } from "@omni-network/contracts/src/pkg/XApp.sol"
```

#### forge

```bash
forge install omni-network/omni
```

```solidity
import { XApp } from "omni-network/omni/contracts/src/pkg/XApp.sol"
```

Note that Omni's contracts are under active development, and are subject to change.

<br />

## Example


Here's an example of a simple cross chain contract, `XGreeter`. This contract lets you send greetings from one chain to another.

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
    ///      The `xfunc` modifier reads the current xmsg into `xmsg` storage
    function greet(string calldata greeting) external xfunc {
        require(isXCall(), "XGreeter: only xcall")
        emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
    }
}
```

Let's walk through this step by step.

First, inherit from `XApp`.


```solidity
contract XGreeter is XApp {
    constructor(address portal) XApp(portal) { }

    // ...
}
```

<br />

### `xcall`

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

<br />

### `xmsg`

When receiving an `xcall`, you can read its context `omni.xmsg()`.

```solidity
omni.xmsg().sourceChainId // where this xcall came from
omni.xmsg().sender        // who sent it
```

With this context, we can have our `XGreeter` emit events detailing the source chain and sender.


```solidity
function greet(string calldata greeting) external {
    emit Greetings(omni.xmsg().sender, omni.xmsg().sourceChainId, greeting);
}
```

<br />

For convenience, `XApp` defines the `xfunc` modifier. This modifier reads the current xmsg into storage, and deletes after its functions execution.

```solidity
modifier xfunc() {
    xmsg = omni.xmsg();
    _;
    delete xmsg;
}
```

It also visually marks a function as the target of an `xcall`. Though, the `xfunc` modifier is not required to receive an `xcall`. Using this modifier, we can simplify our `XGreeter` a bit further.


```solidity
function greet(string calldata greeting) external xfunc {
    emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
}
```

Note that not every call is an `xcall`. In these cases, `xmsg` will be its zero value.

```solidity
xmsg.sourceChainId // 0
xmsg.sender        // address(0)
```

<br />

### `isXCall`

You can check if the current call is an xcall with `isXCall`.

```solidity
function isXCall() internal view returns (bool) {
    return omni.isXCall() && msg.sender == address(omni);
}
```

Note that not only does `isXCall` check with the portal that the current transaction is an `xcall`, it also confirms the sender is the portal itself. This helps avoid mistaking calls later in an `xcall` stacktrace with the original xcall. Using this helper, we can ensure that `greet()` can only every be called via an `xcall`.


```solidity
function greet(string calldata greeting) external xfunc onlyXCall {
    require(isXCall(), "XGreeter: only xcall")
    emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
}
```


<br />

## Gas Limits


Contract calls consume gas. Omni enforces a default gas limit on each


## Fees

TODO
