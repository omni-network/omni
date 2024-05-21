// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_pause_Test
 * @dev Test suite for the AVS pausing functionality
 */
contract OmniAVS_admin_Test is Base {
    /// @dev Test that the owner can pause
    function test_pause_byOwner_succeeds() public {
        vm.prank(omniAVSOwner);
        omniAVS.pause();
        assertEq(omniAVS.paused(), true);
    }

    /// @dev Test that the owner can unpause
    function test_unpause_byOwner_succeeds() public {
        vm.startPrank(omniAVSOwner);
        omniAVS.pause();
        assertEq(omniAVS.paused(), true);
        omniAVS.unpause();
        assertEq(omniAVS.paused(), false);
        vm.stopPrank();
    }

    /// @dev Test that when paused, you cannot register an operator
    function test_registerOperator_whenPaused_reverts() public {
        vm.prank(omniAVSOwner);
        omniAVS.pause();

        address operator = _operator(0);
        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;

        vm.expectRevert("Pausable: paused");
        vm.prank(operator);
        omniAVS.registerOperator(_pubkey(operator), emptySig);
    }

    /// @dev Test that when paused, you cannot syncWithOmni
    function test_syncWithOmni_whenPaused_reverts() public {
        vm.prank(omniAVSOwner);
        omniAVS.pause();

        vm.expectRevert("Pausable: paused");
        omniAVS.syncWithOmni();
    }
}
