// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { Base } from "./common/Base.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

/**
 * @title OmniPortal_sys_Test
 * @dev Test of OmniPortal sys calls
 */
contract OmniPortal_sys_Test is Base {
    function test_setNetwork() public {
        uint64[] memory chain1Shards = new uint64[](2);
        chain1Shards[0] = 3;
        chain1Shards[1] = 4;

        uint64[] memory chain2Shards = new uint64[](2);
        chain2Shards[0] = 5;
        chain2Shards[1] = 6;

        XTypes.Chain memory chain1 = XTypes.Chain({ chainId: 1, shards: chain1Shards });
        XTypes.Chain memory chain2 = XTypes.Chain({ chainId: 2, shards: chain2Shards });

        XTypes.Chain[] memory network = new XTypes.Chain[](2);
        network[0] = chain1;
        network[1] = chain2;

        // assume chain id 1. this should result in
        //  - chain1 is not a supported dest, but defines supprted shards
        //  - chain2 is a supported dest, but does not define supported shards
        vm.chainId(1);
        portal.setNetworkNoAuth(network);

        XTypes.Chain[] memory read = portal.network();

        assertEq(read.length, 2);
        assertEq(read[0].chainId, chain1.chainId);
        assertEq(read[1].chainId, chain2.chainId);
        assertEq(read[0].shards.length, 2);
        assertEq(read[1].shards.length, 2);
        assertEq(read[0].shards[0], chain1Shards[0]);
        assertEq(read[0].shards[1], chain1Shards[1]);
        assertEq(read[1].shards[0], chain2Shards[0]);
        assertEq(read[1].shards[1], chain2Shards[1]);

        // chain 1 is not a supported dest
        assertFalse(portal.isSupportedDest(chain1.chainId));

        // chain 2 is a supported dest
        assertTrue(portal.isSupportedDest(chain2.chainId));

        // chain 1 shards are supported
        for (uint256 i = 0; i < chain1Shards.length; i++) {
            assertTrue(portal.isSupportedShard(chain1Shards[i]));
        }

        // chain 2 shards are not supported
        // NOTE: this works because chain 2 shards differ from chain 1 shards
        // They could have the same supported shards
        for (uint256 i = 0; i < chain2Shards.length; i++) {
            assertFalse(portal.isSupportedShard(chain2Shards[i]));
        }
    }
}
