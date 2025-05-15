// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Vm } from "forge-std/Vm.sol";

/**
 * @title InitializableHelper
 * @notice Helper library to read / manipulate OpenZeppelin v5 Initializable storage
 */
library InitializableHelper {
    Vm internal constant vm = Vm(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D);

    // keccak256(abi.encode(uint256(keccak256("openzeppelin.storage.Initializable")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 internal constant INITIALIZABLE_STORAGE = 0xf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00;

    // INITIALIZABLE_STORAGE stores the following struct:
    //
    // struct InitializableStorage {
    //     /**
    //      * @dev Indicates that the contract has been initialized.
    //      */
    //     uint64 _initialized;
    //     /**
    //      * @dev Indicates that the contract is in the process of being initialized.
    //      */
    //     bool _initializing;
    // }

    /**
     * @notice Returns the Initializable._initialized value for a given address, at slot 0.
     * @dev Reverts in _initializing.
     */
    function getInitialized(address addr) internal view returns (uint64) {
        // _initialized is the first field in the storage layout
        bytes32 slot = vm.load(addr, INITIALIZABLE_STORAGE);

        // if _initializing is false, it's bit will be 0, and will not affect uint conversion
        // if _initializing is true, it's bit will be 1, and will affect uint conversion
        // we therefore require it is 0
        require(uint256(slot) <= uint256(type(uint64).max), "initializing");

        return uint64(uint256(slot));
    }

    /**
     * @notice Returns true if the address has been initialized.
     */
    function isInitialized(address addr) internal view returns (bool) {
        return getInitialized(addr) == uint64(1);
    }

    /**
     * @notice Returns true if the initializers are disabled for a given address.
     */
    function areInitializersDisabled(address addr) internal view returns (bool) {
        return getInitialized(addr) == type(uint64).max;
    }

    /**
     * @notice Disables the initializers for a given address.
     * @dev Sets _initialized to max uint64 (0xFFFFFFFFFFFFFFFF), which disables all initializers.
     */
    function disableInitializers(address addr) internal {
        vm.store(addr, INITIALIZABLE_STORAGE, bytes32(uint256(type(uint64).max)));
    }
}
