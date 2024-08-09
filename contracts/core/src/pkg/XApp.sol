// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { XAppBase } from "./XAppBase.sol";

/**
 * @title XApp
 * @notice Base contract for Omni cross-chain applications
 */
abstract contract XApp is XAppBase {
    constructor(address omni_, uint8 defaultConfLevel_) {
        _setOmniPortal(omni_);
        _setDefaultConfLevel(defaultConfLevel_);
    }
}
