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

    function forceRevertWithReason(string memory reason) public pure {
        revert(reason);
    }

    function panicUnderflow() public pure {
        uint256 x = 0;
        x -= 1;
    }

    function panicDivisionByZero() public pure {
        uint256 x = 1;
        x /= 0;
    }
}
