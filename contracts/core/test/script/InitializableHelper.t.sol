// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { InitializableHelper } from "script/utils/InitializableHelper.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title InitializableHelper_Test
 * @notice Helper contract to test InitializableHelper contract, used to disableInitalizers in genesis.
 */
contract InitializableHelper_Test is Test {
    // @dev Test that disableInitializers makes the correct storage updates
    function test_disableInitalizers() public {
        Initializer i = new Initializer();
        address test = makeAddr("test");

        InitializableHelper.disableInitializers(test);

        i.disableInitalizers();

        // test that storage updates are the same
        assertEq(
            vm.load(test, InitializableHelper.INITIALIZABLE_STORAGE),
            vm.load(address(i), InitializableHelper.INITIALIZABLE_STORAGE)
        );

        // test InitializableHelper getters are correct

        assertEq(InitializableHelper.getInitialized(test), type(uint64).max);
        assertEq(InitializableHelper.getInitialized(address(i)), type(uint64).max);

        assertTrue(InitializableHelper.areInitializersDisabled(test));
        assertTrue(InitializableHelper.areInitializersDisabled(address(i)));
    }

    function test_getInitialized() public {
        Initializer i = new Initializer();

        i.initialize();

        assertEq(InitializableHelper.getInitialized(address(i)), uint64(1));
        assertTrue(InitializableHelper.isInitialized(address(i)));
        assertFalse(InitializableHelper.areInitializersDisabled(address(i)));
    }
}

// @dev Helper contract that uses actual Initializable logic, used to check against
//      InitializableHelper storage updates
contract Initializer is Initializable {
    function initialize() public initializer {
        // do nothing
    }

    function disableInitalizers() public {
        _disableInitializers();
    }
}
