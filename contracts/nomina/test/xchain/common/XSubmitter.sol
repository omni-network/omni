// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.30;

import { INominaPortal } from "src/interfaces/INominaPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title XSubmitter
 * @dev A contract that calls the NominaPortal xsubmit function.
 *      Used to test reentrancy guard.
 */
contract XSubmitter {
    INominaPortal public portal;

    constructor(INominaPortal _portal) {
        portal = _portal;
    }

    function tryXSubmit() public {
        XTypes.Submission memory xsub;
        portal.xsubmit(xsub);
    }
}
