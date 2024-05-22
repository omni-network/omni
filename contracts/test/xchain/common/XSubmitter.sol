// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title XSubmitter
 * @dev A contract that calls the OmniPortal xsubmit function.
 *      Used to test reentrancy guard.
 */
contract XSubmitter {
    IOmniPortal public omni;

    constructor(IOmniPortal portal) {
        omni = portal;
    }

    function tryXSubmit() public {
        XTypes.Submission memory xsub;
        omni.xsubmit(xsub);
    }
}
