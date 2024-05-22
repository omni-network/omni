// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";

/**
 * @title Counter
 * @dev A simple counter for portal testing purposes.
 */
contract Counter {
    IOmniPortal public omni;

    uint256 public count;

    mapping(uint64 => uint256) public countByChainId;

    constructor(IOmniPortal portal) {
        omni = portal;
    }

    function increment() public {
        if (omni.isXCall() && msg.sender == address(omni)) {
            countByChainId[omni.xmsg().sourceChainId]++;
        }

        count++;
    }
}
