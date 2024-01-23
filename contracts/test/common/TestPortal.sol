// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { OmniPortal } from "src/OmniPortal.sol";
import { XChain } from "src/libraries/XChain.sol";

/**
 * @title TestPortal
 * @dev A test contract that exposes the OmniPortal's internal functions.
 */
contract TestPortal is OmniPortal {
    function exec(XChain.Msg calldata xmsg) external {
        _exec(xmsg);
    }
}
