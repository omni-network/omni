// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

/**
 * @title Counter
 * @dev A simple counter for portal testing purposes.
 */
contract Counter {
    uint256 public count;

    /// @dev Returns count so we can verify XReceipt returnData
    function increment() public returns (uint256) {
        count++;
        return count;
    }

    /// @dev Lets us increase gas usage for testing
    function incrementTimes(uint256 times) public {
        for (uint256 i = 0; i < times; i++) {
            increment();
        }
    }
}
