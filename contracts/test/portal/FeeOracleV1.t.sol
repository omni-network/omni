// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Base } from "./common/Base.sol";

/**
 * @title FeeOracleV1_Test
 * @dev Test of FeeOracleV1
 */
contract FeeOracleV1_Test is Base {
    /// @dev Test owner can set fee
    function test_setFee_succeeds() public {
        uint256 fee = 100;

        vm.expectEmit();
        emit FeeChanged(feeOracle.fee(), fee);
        vm.prank(owner);
        feeOracle.setFee(fee);

        assertEq(feeOracle.fee(), fee);
    }

    /// @dev Test non-owner cannot set fee
    function test_setFee_nonOwner_reverts() public {
        vm.prank(relayer);
        vm.expectRevert();
        feeOracle.setFee(100);
    }

    /// @dev Test that the fee cannot be set to 0
    function test_setFee_zero_reverts() public {
        vm.prank(owner);
        vm.expectRevert();
        feeOracle.setFee(0);
    }

    /// @dev Test that the oracle returns a flat fee
    function test_feeFor_succeeds() public {
        uint256 fee = 100;

        vm.expectEmit();
        emit FeeChanged(feeOracle.fee(), fee);
        vm.prank(owner);
        feeOracle.setFee(fee);

        bytes memory data = abi.encodeWithSignature("test()");
        bytes memory data2 = abi.encodeWithSignature("test2()");

        // test for a few different values, fee is flat
        assertEq(feeOracle.feeFor(1, data, 100), fee);
        assertEq(feeOracle.feeFor(1, data2, 100), fee);
        assertEq(feeOracle.feeFor(2, data, 100), fee);
        assertEq(feeOracle.feeFor(2, data2, 1000), fee);
    }
}
