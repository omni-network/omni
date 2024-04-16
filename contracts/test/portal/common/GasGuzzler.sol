// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";

/**
 * @title GasGuzzler
 * @dev A contract that consumes all gas passed to it.
 */
contract GasGuzzler {
    /// @dev Internal counter to increment to consume gas
    uint256 internal _x;

    /// @dev Consumes all gas passed to it, and reverts
    function guzzle() public {
        while (true) {
            _x += 1;
        }
    }
}
