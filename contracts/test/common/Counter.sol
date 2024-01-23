// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

/**
 * @title Counter
 * @dev A simple counter for portal testing purposes.
 */
contract Counter {
    uint256 public count;

    function increment() public {
        count++;
    }
}
