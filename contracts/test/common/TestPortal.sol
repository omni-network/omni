// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { OmniPortal } from "src/protocol/OmniPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { Validators } from "src/libraries/Validators.sol";

/**
 * @title TestPortal
 * @dev A test contract that exposes OmniPortal internal functions, and allows state manipulation.
 */
contract TestPortal is OmniPortal {
    constructor(address owner_, address feeOracle_, uint64 valSetId_, Validators.Validator[] memory validators_)
        OmniPortal(owner_, feeOracle_, valSetId_, validators_)
    { }

    function exec(XTypes.Msg calldata xmsg) external {
        _exec(xmsg);
    }
}
