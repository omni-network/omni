// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

/**
 * @title Reverter
 * @dev A contract with functions that revert with a reason string, for portal testing purposes.
 */
contract Reverter {
    function revertWithReason(string memory reason) public pure {
        revert(reason);
    }

    function failRequireWithReason(string memory reason) public pure {
        require(false, reason);
    }
}
