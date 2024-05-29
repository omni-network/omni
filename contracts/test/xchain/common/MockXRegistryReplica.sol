// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XRegistryBase } from "src/xchain/XRegistryBase.sol";
import { XRegistryNames } from "src/libraries/XRegistryNames.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";

/**
 * @title MockXRegistryReplica
 * @dev A mock xregistry replica that allows setting of portla address
 */
contract MockXRegistryReplica is XRegistryBase {
    function registerPortal(
        address thisChainsPortal, // required to initSourceChain
        uint64 chainId,
        address addr
    ) external {
        _set(chainId, XRegistryNames.OmniPortal, Predeploys.PortalRegistry, addr);
        OmniPortal(thisChainsPortal).initSourceChain(chainId);
    }
}
