// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XRegistryBase } from "src/xchain/XRegistryBase.sol";
import { XRegistryNames } from "src/libraries/XRegistryNames.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

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
        uint64[] memory shards = new uint64[](2);
        shards[0] = ConfLevel.Finalized;
        shards[1] = ConfLevel.Latest;

        Deployment memory dep = Deployment({ addr: addr, metadata: abi.encode(shards) });

        _set(chainId, XRegistryNames.OmniPortal, Predeploys.PortalRegistry, dep);

        if (chainId == OmniPortal(thisChainsPortal).chainId()) {
            OmniPortal(thisChainsPortal).setShards(shards);
        }
    }
}
