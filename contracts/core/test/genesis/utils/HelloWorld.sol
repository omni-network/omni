// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title HelloWorld
 * @notice A simple contract to return a greeting. Used to test proxy upgrades
 */
contract HelloWorld {
    function hello() external pure returns (string memory) {
        return "Hello, World!";
    }
}
