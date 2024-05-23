// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XRegistryBase } from "src/xchain/XRegistryBase.sol";
import { XRegistryNames } from "src/libraries/XRegistryNames.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/**
 * @title MockXRegistryReplica
 * @dev A mock xregistry replica that allows setting of portla address
 */
contract MockXRegistryReplica is XRegistryBase {
    function setPortal(uint64 chainId, address addr) external {
        _set(chainId, XRegistryNames.OmniPortal, Predeploys.PortalRegistry, addr);
    }
}
