// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Test } from "forge-std/Test.sol";
import { XAppUpgradeable } from "src/pkg/XAppUpgradeable.sol";
import { XAppUpgradeableTester } from "./common/XAppTester.sol";
import { XAppTests } from "./common/XAppTests.sol";

/**
 * @title ExtraStrorage
 * @dev Contract to test the storage slots used by XAppUpgradeable
 */
contract ExtraStrorage is XAppUpgradeable {
    uint256 public n;

    constructor(uint256 _n) {
        n = _n;
    }
}

/**
 * @title XAppUpgradeable_Test
 * @dev Test suite for XAppUpgradeable
 */
contract XAppUpgradeable_Test is XAppTests {
    function setUp() public override {
        super.setUp();
        xapp = new XAppUpgradeableTester();
        XAppUpgradeableTester(address(xapp)).initialize(address(portal), defaultConfLevel);
    }

    /// @dev Test that XAppUpgradeable uses 50 storage slots (0-49)
    ///      This makes sure we update __gap properly when adding new storage variables
    function test_storageSlots() public {
        uint256 n = 1;
        uint256 expectSlot = 50;
        ExtraStrorage extra = new ExtraStrorage(n);
        assertEq(vm.load(address(extra), bytes32(expectSlot)), bytes32(n));
    }
}
