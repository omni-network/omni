// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
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

        // only owner
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        portal.setFeeOracle(address(0x123));

        // cannot be zero
        vm.prank(owner);
        vm.expectRevert("OmniPortal: no zero feeOracle");
        portal.setFeeOracle(address(0));
    }

    function test_pauseAll() public {
        // when not paused, can xcall and xsubmit
        assertFalse(portal.isPaused());

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
        address notOwner = address(0x456);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        portal.pause();

        // owner can pause
        vm.prank(owner);
        portal.pause();
        assertTrue(portal.isPaused());

        // when paused, cannot xcall and xsubmit
        vm.expectRevert("OmniPortal: paused");
        vm.chainId(thisChainId);
        portal.xcall(chainAId, conf, to, data, gasLimit);

        vm.expectRevert("OmniPortal: paused");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub1);
    }

    function test_pauseXCall() public {
        assertFalse(portal.isPaused(portal.ActionXCall()));
        assertFalse(portal.isPaused(portal.ActionXSubmit(), chainAId));
        assertFalse(portal.isPaused(portal.ActionXSubmit(), chainBId));

        // xcall params
        uint8 conf = ConfLevel.Finalized;
        address to = address(0x1234);
        bytes memory data = abi.encodeWithSignature("test()");
        uint64 gasLimit = 100_000;

        // can xcall
        vm.chainId(thisChainId);
        portal.xcall{ value: 1 ether }(chainAId, conf, to, data, gasLimit);

        // pause xcall to chain b
        vm.prank(owner);
        portal.pauseXCallTo(chainBId);

        assertFalse(portal.isPaused(portal.ActionXCall()));
        assertFalse(portal.isPaused(portal.ActionXSubmit(), chainAId));
        assertTrue(portal.isPaused(portal.ActionXCall(), chainBId));

        // can xcall to chain a
        vm.chainId(thisChainId);
        portal.xcall{ value: 1 ether }(chainAId, conf, to, data, gasLimit);

        // cannot xcall to chain b
        vm.expectRevert("OmniPortal: paused");
        vm.chainId(thisChainId);
        portal.xcall(chainBId, conf, to, data, gasLimit);

        // unpause xcall to chain b
        vm.prank(owner);
        portal.unpauseXCallTo(chainBId);

        assertFalse(portal.isPaused(portal.ActionXCall()));

        // can xcall to chain b
        vm.chainId(thisChainId);
        portal.xcall{ value: 1 ether }(chainBId, conf, to, data, gasLimit);

        // pause all xcall
        vm.prank(owner);
        portal.pauseXCall();

        assertTrue(portal.isPaused(portal.ActionXCall()));

        // cannot xcall to chain a
        vm.expectRevert("OmniPortal: paused");
        vm.chainId(thisChainId);
        portal.xcall(chainAId, conf, to, data, gasLimit);

        // cannot xcall to chain b
        vm.expectRevert("OmniPortal: paused");
        vm.chainId(thisChainId);
        portal.xcall(chainBId, conf, to, data, gasLimit);

        // unpause all xcall
        vm.prank(owner);
        portal.unpauseXCall();

        assertFalse(portal.isPaused(portal.ActionXCall()));

        // can xcall to chain a
        vm.chainId(thisChainId);
        portal.xcall{ value: 1 ether }(chainAId, conf, to, data, gasLimit);

        // can xcall to chain b
        vm.chainId(thisChainId);
        portal.xcall{ value: 1 ether }(chainBId, conf, to, data, gasLimit);
    }

    function test_pauseXSubmit() public {
        assertFalse(portal.isPaused(portal.ActionXSubmit()));
        assertFalse(portal.isPaused(portal.ActionXSubmit(), chainAId));
        assertFalse(portal.isPaused(portal.ActionXSubmit(), chainBId));

        // can xsubmit
        // we use a stub xsub, so we don't need to provide a real one
        // when not paused, xsubmit should error with "OmniPortal: no xmsgs"
        // when paused, xsubmit should error with "OmniPortal: paused"
        XTypes.Submission memory xsub;
        xsub.blockHeader = XTypes.BlockHeader({
            sourceChainId: chainAId,
            consensusChainId: omniCChainID,
            confLevel: ConfLevel.Finalized,
            offset: 1,
            sourceBlockHeight: 100,
            sourceBlockHash: keccak256("hash")
        });
        vm.expectRevert("OmniPortal: no xmsgs");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // pause xsubmit from chain b
        vm.prank(owner);
        portal.pauseXSubmitFrom(chainBId);

        assertFalse(portal.isPaused(portal.ActionXSubmit()));
        assertFalse(portal.isPaused(portal.ActionXSubmit(), chainAId));
        assertTrue(portal.isPaused(portal.ActionXSubmit(), chainBId));

        // can xsubmit from chain a
        vm.expectRevert("OmniPortal: no xmsgs");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // cannot xsubmit from chain b
        xsub.blockHeader.sourceChainId = chainBId;
        vm.expectRevert("OmniPortal: paused");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // unpause xsubmit from chain b
        vm.prank(owner);
        portal.unpauseXSubmitFrom(chainBId);

        assertFalse(portal.isPaused(portal.ActionXSubmit()));

        // can xsubmit from chain b
        vm.expectRevert("OmniPortal: no xmsgs");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // pause all xsubmit
        vm.prank(owner);
        portal.pauseXSubmit();

        assertTrue(portal.isPaused(portal.ActionXSubmit()));

        // cannot xsubmit from chain a
        xsub.blockHeader.sourceChainId = chainAId;
        vm.expectRevert("OmniPortal: paused");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // cannot xsubmit from chain b
        xsub.blockHeader.sourceChainId = chainBId;
        vm.expectRevert("OmniPortal: paused");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // unpause all xsubmit
        vm.prank(owner);
        portal.unpauseXSubmit();

        assertFalse(portal.isPaused(portal.ActionXSubmit()));

        // can xsubmit from chain a
        xsub.blockHeader.sourceChainId = chainAId;
        vm.expectRevert("OmniPortal: no xmsgs");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // can xsubmit from chain b
        xsub.blockHeader.sourceChainId = chainBId;
        vm.expectRevert("OmniPortal: no xmsgs");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);
    }

    function test_setInXMsgOffset() public {
        uint64 srcChainId = 1;
        uint64 shardId = 1;
        uint64 offset = 3;

        // only owner
        address notOwner = address(0x456);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        portal.setInXMsgOffset(srcChainId, shardId, offset);

        // set offset
        vm.prank(owner);
        portal.setInXMsgOffset(srcChainId, shardId, offset);
        assertEq(portal.inXMsgOffset(srcChainId, shardId), offset);
    }

    function test_setInXBlockOffset() public {
        uint64 srcChainId = 1;
        uint64 shardId = 1;
        uint64 offset = 3;

        // only owner
        address notOwner = address(0x456);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        portal.setInXBlockOffset(srcChainId, shardId, offset);

        // set offset
        vm.prank(owner);
        portal.setInXBlockOffset(srcChainId, shardId, offset);
        assertEq(portal.inXBlockOffset(srcChainId, shardId), offset);
    }

    function test_setXSubValsetCutoff() public {
        // only owner
        address notOwner = address(0x456);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        portal.setXSubValsetCutoff(1);

        // set cutoff
        vm.prank(owner);
        portal.setXSubValsetCutoff(1);

        assertEq(portal.xsubValsetCutoff(), 1);
    }
}
