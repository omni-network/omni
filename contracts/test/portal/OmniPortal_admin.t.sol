// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Base } from "./common/Base.sol";

/**
 * @title OmniPortal_admin_Test
 * @dev Test of OmniPortal admin controls
 */
contract OmniPortal_admin_Test is Base {
    /// @dev Test owner can set fee oracle
    function test_setFeeOracle_succeeds() public {
        address newFeeOracle = address(0x123);

        vm.expectEmit();
        emit FeeOracleChanged(portal.feeOracle(), newFeeOracle);
        vm.prank(owner);
        portal.setFeeOracle(newFeeOracle);

        assertEq(portal.feeOracle(), newFeeOracle);
    }

    /// @dev Test non-owner cannot set fee oracle
    function test_setFeeOracle_nonOwner_reverts() public {
        vm.prank(relayer);
        vm.expectRevert();
        portal.setFeeOracle(address(0x123));
    }

    /// @dev Test that the fee oracle cannot be set to address(0)
    function test_setFeeOracle_zero_reverts() public {
        vm.prank(owner);
        vm.expectRevert("OmniPortal: no zero feeOracle");
        portal.setFeeOracle(address(0));
    }
}
