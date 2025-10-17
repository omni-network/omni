// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Vm } from "forge-std/Vm.sol";

/**
 * @title InitializableHelperSolady
 * @notice Helper library to read / manipulate Solady Initializable storage
 */
library InitializableHelperSolady {
    Vm internal constant vm = Vm(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);

    // bytes32(~uint256(uint32(bytes4(keccak256("_INITIALIZABLE_SLOT")))))
    bytes32 internal constant INITIALIZABLE_STORAGE =
        0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffbf601132;

    // INITIALIZABLE_STORAGE stores data with the following layout:
    //
    // Bits Layout:
    // - [0]     `initializing`
    // - [1..64] `initializedVersion`
    //
    // This means that a freshly initialized contract will return 2 in the `getInitialized` function.

    /**
     * @notice Returns the initialized version for a given address.
     * @dev Reverts if the contract is currently initializing.
     */
    function getInitialized(address addr) internal view returns (uint64) {
        bytes32 slot = vm.load(addr, INITIALIZABLE_STORAGE);

        // Check if initializing bit is set (bit 0)
        require((uint256(slot) & 1) == 0, "initializing");

        // Extract only the initializedVersion (bits 1-64)
        // by right-shifting by 1 to remove the initializing bit
        return uint64(uint256(slot) >> 1);
    }

    /**
     * @notice Returns true if the address has been initialized.
     */
    function isInitialized(address addr) internal view returns (bool) {
        return getInitialized(addr) == 1;
    }

    /**
     * @notice Returns true if the initializers are disabled for a given address.
     */
    function areInitializersDisabled(address addr) internal view returns (bool) {
        return getInitialized(addr) == type(uint64).max;
    }

    /**
     * @notice Disables the initializers for a given address.
     * @dev Sets _initialized to left shifted max uint64 (0x01FFFFFFFFFFFFFFFE), which disables all initializers.
     */
    function disableInitializers(address addr) internal {
        vm.store(addr, INITIALIZABLE_STORAGE, bytes32(uint256(type(uint64).max) << 1));
    }
}
