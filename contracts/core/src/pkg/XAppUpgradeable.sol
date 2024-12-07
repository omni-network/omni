// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { XAppBase } from "./XAppBase.sol";

/**
 * @title XAppUpgradeable
 * @notice Base contract for Omni cross-chain applications.
 *         Allows for future upgrades to the XAppBase contract
 */
abstract contract XAppUpgradeable is Initializable, XAppBase {
    function __XApp_init(address omni_, uint8 defaultConfLevel_) internal onlyInitializing {
        _setOmniPortal(omni_);
        _setDefaultConfLevel(defaultConfLevel_);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     *
     * NOTE: We use storage gaps, rather than ERC-7201 namespaced storage, to allow for
     * access to `xmsg` storage variable without invoking a function. So we users can
     * use the following syntax: `xmsg.sender`, or `xmsg.sourceChainId`.
     */
    uint256[47] private __gap;
}
