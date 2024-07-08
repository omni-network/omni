// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Script } from "forge-std/Script.sol";
import { Fixtures } from "test/xchain/common/Fixtures.sol";

/**
 * @title WriteTestXBlocks
 * @dev Script to write test xblock fixtures to disk. See Fixtures.writeXBlocks()
 */
contract WriteTestXBlocks is Script {
    function run() external {
        Fixtures f = new Fixtures();
        f.setUp();
        f.writeXBlocks();
    }
}
