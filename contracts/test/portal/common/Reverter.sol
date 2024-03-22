// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Reverter
 * @dev A contract with functions that revert with a reason string, for portal testing purposes.
 */
contract Reverter {
    function forceRevert() public pure {
        revert();
    }
}
