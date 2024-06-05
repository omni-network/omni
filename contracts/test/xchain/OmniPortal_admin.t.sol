// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Base } from "./common/Base.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

/**
 * @title OmniPortal_admin_Test
 * @dev Test of OmniPortal admin controls
 */
contract OmniPortal_admin_Test is Base {
    function test_setFeeOracle() public {
        address newFeeOracle = address(0x123);

        // owner can set
        vm.prank(owner);
        portal.setFeeOracle(newFeeOracle);
        assertEq(portal.feeOracle(), newFeeOracle);

        assertEq(portal.feeOracle(), newFeeOracle);

        // only owner
        vm.expectRevert("Ownable: caller is not the owner");
        portal.setFeeOracle(address(0x456));

        // cannot be zero
        vm.prank(owner);
        vm.expectRevert("OmniPortal: no zero feeOracle");
        portal.setFeeOracle(address(0));
    }

    function test_pause() public {
        // when not paused, can xcall and xsubmit
        assertFalse(portal.paused());

        // xcall params
        uint8 conf = ConfLevel.Finalized;
        address to = address(0x1234);
        bytes memory data = abi.encodeWithSignature("test()");
        uint64 gasLimit = 100_000;

        // xcall
        vm.chainId(thisChainId);
        portal.xcall{ value: 1 ether }(chainAId, conf, to, data, gasLimit);

        // xsubmit
        XTypes.Submission memory xsub1 = readXSubmission({ name: "xblock1", destChainId: thisChainId });
        vm.chainId(thisChainId);
        portal.xsubmit(xsub1);

        // only owner can pause
        vm.expectRevert("Ownable: caller is not the owner");
        portal.pause();

        // owner can pause
        vm.prank(owner);
        portal.pause();
        assertTrue(portal.paused());

        // when paused, cannot xcall and xsubmit
        vm.expectRevert("Pausable: paused");
        vm.chainId(thisChainId);
        portal.xcall(chainAId, conf, to, data, gasLimit);

        vm.expectRevert("Pausable: paused");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub1);
    }
}
