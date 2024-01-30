// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Base } from "test/common/Base.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title OmniPortal_xsubmit_Test
 * @dev Tests of OmniPortal.xsubmit
 */
contract OmniPortal_xsubmit_Test is Base {
    function test_xsubmit_xblock1_succeeds() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());

        uint64 sourceChainId = xsub.msgs[0].sourceChainId;
        uint64 expectedOffset = xsub.msgs[xsub.msgs.length - 1].streamOffset;

        vm.prank(relayer);
        vm.recordLogs();
        expectCalls(xsub.msgs);
        portal.xsubmit(xsub);

        assertEq(portal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(portal.inXStreamBlockHeight(sourceChainId), xsub.blockHeader.blockHeight);
        assertReceipts(vm.getRecordedLogs(), xsub.msgs);
    }

    function test_xsubmit_xblock2_succeeds() public {
        // need to submit xblock1 first, to set the streamOffset
        XTypes.Submission memory xsub1 = readXSubmission("xblock1", portal.chainId());
        portal.xsubmit(xsub1);

        XTypes.Submission memory xsub2 = readXSubmission("xblock2", portal.chainId());

        uint64 sourceChainId = xsub2.msgs[0].sourceChainId;
        uint64 expectedOffset = xsub2.msgs[xsub2.msgs.length - 1].streamOffset;

        vm.prank(relayer);
        vm.recordLogs();
        expectCalls(xsub2.msgs);
        portal.xsubmit(xsub2);

        assertEq(portal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(portal.inXStreamBlockHeight(sourceChainId), xsub2.blockHeader.blockHeight);
        assertReceipts(vm.getRecordedLogs(), xsub2.msgs);
    }

    function test_xsubmit_xblock1_chainB_succeeds() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", chainBId);

        uint64 sourceChainId = xsub.msgs[0].sourceChainId;
        uint64 expectedOffset = xsub.msgs[xsub.msgs.length - 1].streamOffset;

        vm.prank(relayer);
        vm.recordLogs();
        expectCalls(xsub.msgs);
        chainBPortal.xsubmit(xsub);

        assertEq(chainBPortal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(chainBPortal.inXStreamBlockHeight(sourceChainId), xsub.blockHeader.blockHeight);
        assertReceipts(vm.getRecordedLogs(), xsub.msgs);
    }

    function test_xsubmit_xblock2_chainB_succeeds() public {
        // need to submit xblock1 first, to set the streamOffset
        XTypes.Submission memory xsub1 = readXSubmission("xblock1", chainBId);
        chainBPortal.xsubmit(xsub1);

        XTypes.Submission memory xsub2 = readXSubmission("xblock2", chainBId);

        uint64 sourceChainId = xsub2.msgs[0].sourceChainId;
        uint64 expectedOffset = xsub2.msgs[xsub2.msgs.length - 1].streamOffset;

        vm.prank(relayer);
        vm.recordLogs();
        expectCalls(xsub2.msgs);
        chainBPortal.xsubmit(xsub2);

        assertEq(chainBPortal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(chainBPortal.inXStreamBlockHeight(sourceChainId), xsub2.blockHeader.blockHeight);
        assertReceipts(vm.getRecordedLogs(), xsub2.msgs);
    }

    function test_xsubmit_wrongChainId_reverts() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());

        vm.expectRevert("OmniPortal: wrong destChainId");
        chainBPortal.xsubmit(xsub);
    }

    function test_xsubmit_wrongStreamOffset_reverts() public {
        XTypes.Submission memory xsub = readXSubmission("xblock2", portal.chainId());

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.xsubmit(xsub);
    }

    function test_xsubmit_invalidAttestationRoot_reverts() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());

        xsub.attestationRoot = keccak256("invalid");

        vm.expectRevert("OmniPortal: invalid proof");
        portal.xsubmit(xsub);
    }

    function test_xsubmit_invalidMsgs_reverts() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());

        xsub.msgs[0].data = abi.encodeWithSignature("invalid()");

        vm.expectRevert("OmniPortal: invalid proof");
        portal.xsubmit(xsub);
    }
}
