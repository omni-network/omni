// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.30;

import { INominaPortal } from "src/interfaces/INominaPortal.sol";

/**
 * @title Counter
 * @dev A simple counter for portal testing purposes.
 */
contract Counter {
    INominaPortal public portal;

    uint256 public count;

    mapping(uint64 => uint256) public countByChainId;

    constructor(INominaPortal _portal) {
        portal = _portal;
    }

    function increment() public {
        if (portal.isXCall() && msg.sender == address(portal)) {
            countByChainId[portal.xmsg().sourceChainId]++;
        }

        count++;
    }
}
