---
sidebar_position: 2
id: example
---

# Example Application

## Example

This is a simple example of a global application that uses Omni to operate across multiple rollups. It tracks a global count, that can be incremented from multiple rollups.

### On Omni

The `GlobalCounter` contract on Omni tracks the current global count across all rollups. It also tracks the amount incremented by each rollup. It can also increment the count on a rollup (which in turn increments the global counter here).

This contract implements callback functions that simply store whether an `incrementOnChain` call was reverted or successful.

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import {OmniScient} from "omni-std/OmniScient.sol";
import {OmniCodec} from 'omni-std/OmniCodec.sol';

contract GlobalCounter is OmniScient {
    uint256 public count;

    mapping(string => uint256) internal countByChain;
    mapping(uint256 => bool) public incrementSuccess;

    event IncrementOnChainSuccess(uint256 nonce);
    event IncrementOnChainReverted(uint256 nonce);

    event Increment(uint256 count);

    constructor() {
        count = 0;
    }

    function increment() public {
        count += 1;
        countByChain[omni.txSourceChain()] += 1;
        emit Increment(count);
    }

    function incrementOnChain(string memory chain, address counter) public {
      omni.sendTx(
        chain,
        counter,
        abi.encodeWithSignature("increment()")
      );
    }

    function getCountFor(string memory chain) public view returns (uint256) {
        return countByChain[chain];
    }

    function onXChainTxSuccess(OmniCodec.Tx memory _xtx, address _sender, bytes memory _returnValue, uint256 _gasSpent) external override onlyOmni {
        incrementSuccess[_xtx.nonce] = true;
        emit IncrementOnChainSuccess(_xtx.nonce);
    }

    function onXChainTxReverted(OmniCodec.Tx memory _xtx, address _sender, uint256 _gasSpent) external override onlyOmni {
        incrementSuccess[_xtx.nonce] = false;
        emit IncrementOnChainReverted(_xtx.nonce);
    }
}
```

### On Rollups

This `LocalCounter` contract is deployed to rollups. The primary method is the `increment()` method, which sends a transaction to Omni to increment the global counter.

It also includes a `syncGlobalCount` function, which uses the `omni.verifyState` method to check the global count state variable in the `GlobalCounter` contract.

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import {IOmniPortal} from "omni-std/interfaces/IOmniPortal.sol";

contract LocalCounter {
    uint256 public count;
    uint256 public globalCount;
    uint256 public globalBlockNumber;
    address public global; // counter address on omni

    mapping(string => uint256) public countByChain;

    IOmniPortal public omni;

    event Increment(uint256 count);

    constructor(IOmniPortal _portal, address _global) {
        count = 0;
        globalCount = 0;
        omni = _portal;
        global = _global;
    }

    function increment() public {
        count += 1;

        if (omni.isXChainTx()) {
            countByChain[omni.txSourceChain()] += 1;
        }

        omni.sendOmniTx(
            global,
            abi.encodeWithSignature("increment()")
        );

        emit Increment(count);
    }

    function incrementOnChain(string memory chain, address counter) public {
        omni.sendXChainTx(
            chain,
            counter,
            abi.encodeWithSignature("increment()")
        );
    }

    function syncGlobalCount(uint64 _blockNumber, bytes calldata _storageProof, uint256 _globalCount) public {
        require(_blockNumber > globalBlockNumber, "LocalCounter: block number must be greater than global block number");

        bytes memory storageSlotKey = abi.encodePacked(hex"02", global, bytes32(uint256(0)));
        bytes memory storageSlotValue = abi.encodePacked(bytes32(_globalCount));

        bool verified = omni.verifyOmniState(_blockNumber, _storageProof, storageSlotKey, storageSlotValue);

        require(verified, "LocalCounter: invalid proof");

        globalCount = _globalCount;
        globalBlockNumber = _blockNumber;
    }
}
```
