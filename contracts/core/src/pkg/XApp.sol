// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { XAppBase } from "./XAppBase.sol";

/**
 * @title XApp
 * @notice Base contract for Omni cross-chain applications
 */
abstract contract XApp is XAppBase {
    constructor(address portal, address exchange, uint8 defaultConf) {
        _setOmniPortal(portal);
        _setOmniGasEx(exchange);
        _setDefaultConfLevel(defaultConf);
    }
}
