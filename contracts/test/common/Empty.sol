// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/**
 * @title Empty
 * @dev An empy contract, used to deploy transparent proxies
 */
contract Empty {
    function foo() public pure returns (uint256) {
        return 0;
    }
}
