---
sidebar_position: 1
---

# Build a Native Multi-Rollup xdApp

To enable a dApp to run on multiple rollups, the dApp must be able to interact with the Omni network. This guide will walk you through the process of building a dApp that can run on multiple rollups.

## Summary

- The dApp will be able to interact with the Omni network through the `IOmniPortal` interface.
- The dApp will send cross-rollup transactions using the `omni.xcall` function.
  - The dApp will calculate the fee for the cross-rollup transaction using the `omni.feeFor` function.
- The dApp will receive cross-rollup transactions using the `omni.isXCall` function and asserting the `msg.sender` is the Omni portal.
  - The dApp will unpack cross-rollup transactions using the `omni.xmsg` function.

## Hello World Donation dApp

This is a simple example of a single-rollup application without Omni. It tracks a total donation amount and allows signers to add a message to the donation.

### Solidity Implementation

```solidity
pragma solidity ^0.8.0;

contract HelloWorldDonation {
    uint256 public totalDonation;
    mapping(address => string) public messages;
    mapping(address => uint256) public donations;

    function donate(string memory message) public payable {
        totalDonation += msg.value;
        messages[msg.sender] = message;
        donations[msg.sender] += msg.value;
    }

    function getMessage(address donor) public view returns (string memory) {
        return messages[donor];
    }

    function getDonation(address donor) public view returns (uint256) {
        return donations[donor];
    }

    function getTotalDonation() public view returns (uint256) {
        return totalDonation;
    }
}
```

Sadly, this dApp is not Omni-enabled. It only works on a single rollup. Let's Omni-enable it!

## Omni-Enabled Hello World Donation xdApp

This is a simple example of a cross-rollup application that uses Omni. It tracks a total donation amount and allows signers to add a message to the donation.

### Solidity Implementation

#### Main Contract

```solidity
pragma solidity ^0.8.0;

import "@omni/contracts/IOmniPortal.sol";

contract HelloWorldDonation {
    uint256 public totalDonation;
    mapping(address => string) public messages;
    mapping(address => uint256) public donations;

    IOmniPortal public omni;

    constructor(IOmniPortal _omni) {
        omni = _omni;
    }

    function donate(string memory message) public payable {
        if (omni.isXCall() && msg.sender == address(omni)) {
            totalDonation += msg.value;
            messages[omni.xmsg().sender] = message;
            donations[omni.xmsg().sender] += msg.value;
        } else {
            totalDonation += msg.value;
            messages[msg.sender] = message;
            donations[msg.sender] += msg.value;
        }
    }

    function getMessage(address donor) public view returns (string memory) {
        return messages[donor];
    }

    function getDonation(address donor) public view returns (uint256) {
        return donations[donor];
    }

    function getTotalDonation() public view returns (uint256) {
        return totalDonation;
    }
}
```

#### Additional Rollup Contract

```solidity
pragma solidity ^0.8.0;

import "@omni/contracts/IOmniPortal.sol";

contract HelloWorldDonation {
    uint64 public destinationChainId;
    address public destinationContract;

    IOmniPortal public omni;

    constructor(IOmniPortal _omni, uint64 _destinationChainId, address _destinationContract) {
        omni = _omni;
        destinationChainId = _destinationChainId;
        destinationContract = _destinationContract;
    }

    function donate(string memory message) public payable {
        xdonate(message);
    }

    function xdonate(string memory message) internal payable {
        bytes memory data = abi.encodeWithSelector(this.donate.selector, message);
        uint256 fee = omni.feeFor(destinationChainId, data);
        require(msg.value >= fee, "HelloWorldDonation: insufficient fee");
        omni.xcall{ value: fee }(destinationChainId, destinationContract, data);
    }

    function getMessage(address donor) public view returns (string memory) {
        bytes memory data = abi.encodeWithSignature("getMessage(address)", donor);
        return abi.decode(omni.xcall(destinationChainId, destinationContract, data), (string));
    }

    function getDonation(address donor) public view returns (uint256) {
        bytes memory data = abi.encodeWithSignature("getDonation(address)", donor);
        return abi.decode(omni.xcall(destinationChainId, destinationContract, data), (uint256));
    }

    function getTotalDonation() public view returns (uint256) {
        bytes memory data = abi.encodeWithSignature("getTotalDonation()");
        return abi.decode(omni.xcall(destinationChainId, destinationContract, data), (uint256));
    }
}
```

### Points to Note

#### Deployment Pattern

- The cross-rollup architecture is implemented using two contracts: the main contract and the additional rollup contract.
  - The main contract is the contract that will be deployed on the main rollup or EVM network.
  - The additional rollup contract is the contract that will be deployed on any additional rollup.
    - The additional rollup contract will interact with the main contract using cross-rollup transactions, forwarding state queries to the main contract, and querying the main contract for state updates.

#### Main Contract

- The `donate` function now checks if the transaction is a cross-rollup transaction using `omni.isXCall` and if the `msg.sender` is the Omni portal.
- The `donate` function now uses `omni.xmsg` to unpack the cross-rollup transaction.
- The `getMessage`, `getDonation`, and `getTotalDonation` functions query the state of the main contract directly.

#### Additional Rollup Contract

- The `donate` function now forwards the donation to the main contract using the `xdonate` function.
- The `xdonate` function calculates the fee for the cross-rollup transaction using `omni.feeFor` and forwards the donation to the main contract using `omni.xcall`.
- The `getMessage`, `getDonation`, and `getTotalDonation` functions now query the main contract for state updates using `omni.xcall`.

## Comparison

### Single-Rollup dApp vs. Omni-Enabled xdApp

| | Single Network dApp | Omni-Enabled xdApp |
|---|---|---|
| Interact with other rollups | ❌ | ✅ |
| Aware of donations across networks | ❌ | ✅ |
| Aware of messages across networks | ❌ | ✅ |
| Complexity of Development | 🟢 Low | 🟡 Medium |
| Ease of Deploying to Additional Rollups | 🔴 High | 🟢 Low |
