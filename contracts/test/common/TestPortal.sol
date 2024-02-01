// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { OmniPortal } from "src/OmniPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title TestPortal
 * @dev A test contract that exposes OmniPortal internal functions, and allows state manipulation.
 */
contract TestPortal is OmniPortal {
    constructor(address owner_, address feeOracle_) OmniPortal(owner_, feeOracle_) { }

    function exec(XTypes.Msg calldata xmsg) external {
        _exec(xmsg);
    }
}
