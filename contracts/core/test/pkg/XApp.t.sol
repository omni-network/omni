// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Test } from "forge-std/Test.sol";
import { XAppTester } from "./common/XAppTester.sol";
import { XAppTests } from "./common/XAppTests.sol";

/**
 * @title XApp_Test
 * @dev Test suite for XApp
 */
contract XApp_Test is XAppTests {
    function setUp() public override {
        super.setUp();
        xapp = new XAppTester(address(portal), defaultConfLevel);
    }
}
