---
sidebar_position: 2
---

# Upgrade an Existing dApp

To upgrade an existing dApp to run on multiple rollups little changes are required to your core dApp logic. The main changes are to enable the dApp to interact with the Omni relaying network through the `IOmniPortal` interface.

## Rollup Deployment Organization

### Central State Management

You can upgrade your existing dApp to run on multiple networks by prioritising state to a designated network deployed contract. This contract will be responsible for managing the state of the dApp across multiple rollups and contracts deployed to other networks will be responsible for updating the state of the contract on the main designated network.

This reduces overall complexity as compared to distributing state across all networks and allows for a single source of truth for the dApp.

### Distributed State Management

Alternatively, you can distribute the state of the dApp across multiple rollups. This approach is more complex and requires more careful management of state across multiple networks as changes to the state of the dApp on one network should be propagated to the state of the dApp on other networks.

## Solidity Integration

- Enable the dApp to send cross-rollup transactions using the `omni.xcall` function.
  - Calculate fees pre-emptively for the cross-rollup transaction using the `omni.feeFor` function.
- Enable the dApp to receive cross-rollup transactions using the `omni.isXCall` method and assert the `msg.sender` is the Omni portal.
  - Unpack cross-rollup transactions using the `omni.xmsg` method.

### Sending Cross-Rollup Transactions

```solidity
function sendMethod(uint64 destinationChainId, address toAddress, string memory message) public payable {
    bytes memory data = abi.encodeWithSignature("setMessage(string)", message);
    uint256 fee = omni.feeFor(destinationChainId, data);
    // omit if we don't want to charge the user
    require(msg.value >= fee, "Counter: insufficient fee");
    omni.xcall{ value: fee }(destinationChainId, toAddress, data);
}
```

### Receiving Cross-Rollup Transactions

```solidity
function receiveMethod(string memory message) public {
    require(omni.isXCall(), "Counter: not a cross-rollup transaction");
    require(omni.xmsg().sender == address(omni), "Counter: not from the Omni portal");
    (string memory _message) = abi.decode(omni.xmsg().data, (string));
    setMessage(_message);
}
```
