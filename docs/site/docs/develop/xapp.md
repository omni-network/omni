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

## Usage

Inherit from `XApp`.


```solidity
/**
 * @title XGreeter
 * @notice A cross chain greeter
 */
contract XGreeter is XApp {
    constructor(address portal) XApp(portal) { }

    // ...
}
```

<br />

### `xcall`

To call a contract on another chain, use `xcall`.


```solidity
/// @dev Greet on another chain
///      `xcall` is a method inherited from `XApp`
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

When receiving an xcall, you can read its context through the `omni.xmsg()`.


```solidity
/// @dev Emitted when someone greets the ether
event Greetings(address indexed from, uint64 indexed fromChainId, string greeting);

/// @dev Greet on this chain
function greet(string calldata greeting) external {
    emit Greetings(omni.xmsg().sender, omni.xmsg().sourceChainId, greeting);
}
```

For convenience, `XApp` defines the `setXMsg` modifier, that reads the current xmsg into the `xmsg` storage variable before execution, and deletes it afterwards.

```solidity
/// @dev Greet on this chain
///      The `setXMsg` modifier reads the current xmsg into `xmsg` storage
function greet(string calldata greeting) external setXMsg {
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
// ...
function greet() external setXMsg {
    if (isXCall()) {
        // if it's an xcall, emit a greeting from the source chain
        emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
    } else {
        // if it's not, emit a greeting from this chain
        emit Greetings(msg.sender, block.chainId, greeting);
    }
}
```

<br />


### `onlyXCall`

If you want your function to only accept `xcalls`, use the `onlyXCall` modifier.


```solidity
// ...
function greet() external setXMsg onlyXCall {
    // now, it's always an xcall
    emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
}
```

<br />

Full `XGreeter` source:

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
    ///      The `setXMsg` modifier reads the current xmsg into `xmsg` storage
    ///      The `onlyXCall` modifier ensures this call is indeed an xcall
    function greet(string calldata greeting) external setXMsg onlyXCall {
        emit Greetings(xmsg.sender, xmsg.sourceChainId, greeting);
    }
}
```

<br />

## Gas Limits

TODO

## Fees

TODO
