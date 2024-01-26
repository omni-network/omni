// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Script } from "forge-std/Script.sol";
import { Fixtures } from "test/common/Fixtures.sol";

/**
 * @title WriteTestXBlocks
 * @dev Script tow write test xblock fixtures to disc. See Fixtures.writeXBlocks()
 */
contract WriteTestXBlocks is Script {
    function run() external {
        Fixtures f = new Fixtures();
        f.setUp();
        f.writeXBlocks();
    }
}
