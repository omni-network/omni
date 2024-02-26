// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Base } from "test/common/Base.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { Vm } from "forge-std/Vm.sol";
import { Validators } from "src/libraries/Validators.sol";

/**
 * @title OmniPortal_xsubmit_Test
 * @dev Tests of OmniPortal.xsubmit
 */
contract OmniPortal_xsubmit_Test is Base {
    function test_xsubmit_xblock1_succeeds() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());

        uint64 sourceChainId = xsub.blockHeader.sourceChainId;
        uint64 expectedOffset = xsub.msgs[xsub.msgs.length - 1].streamOffset;
        uint256 expectedCount = numIncrements(xsub.msgs);

        vm.prank(relayer);
        vm.recordLogs();
        expectCalls(xsub.msgs);
        portal.xsubmit(xsub);

        assertEq(portal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(portal.inXStreamBlockHeight(sourceChainId), xsub.blockHeader.blockHeight);
        assertEq(counter.count(), expectedCount);
        assertEq(counter.countByChainId(sourceChainId), expectedCount);
        assertReceipts(vm.getRecordedLogs(), xsub.msgs);
    }

    function test_xsubmit_xblock2_succeeds() public {
        // need to submit xblock1 first, to set the streamOffset
        XTypes.Submission memory xsub1 = readXSubmission("xblock1", portal.chainId());
        portal.xsubmit(xsub1);

        XTypes.Submission memory xsub2 = readXSubmission("xblock2", portal.chainId());

        uint64 sourceChainId = xsub2.blockHeader.sourceChainId;
        uint64 expectedOffset = xsub2.msgs[xsub2.msgs.length - 1].streamOffset;
        uint256 expectedCount = numIncrements(xsub1.msgs) + numIncrements(xsub2.msgs);

        vm.prank(relayer);
        vm.recordLogs();
        expectCalls(xsub2.msgs);
        portal.xsubmit(xsub2);

        assertEq(portal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(portal.inXStreamBlockHeight(sourceChainId), xsub2.blockHeader.blockHeight);
        assertEq(counter.count(), expectedCount);
        assertEq(counter.countByChainId(sourceChainId), expectedCount);
        assertReceipts(vm.getRecordedLogs(), xsub2.msgs);
    }

    function test_xsubmit_xblock1_chainB_succeeds() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", chainBId);

        uint64 sourceChainId = xsub.blockHeader.sourceChainId;
        uint64 expectedOffset = xsub.msgs[xsub.msgs.length - 1].streamOffset;
        uint256 expectedCount = numIncrements(xsub.msgs);

        vm.prank(relayer);
        vm.recordLogs();
        expectCalls(xsub.msgs);
        chainBPortal.xsubmit(xsub);

        assertEq(chainBPortal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(chainBPortal.inXStreamBlockHeight(sourceChainId), xsub.blockHeader.blockHeight);
        assertEq(chainBCounter.count(), expectedCount);
        assertEq(chainBCounter.countByChainId(sourceChainId), expectedCount);
        assertReceipts(vm.getRecordedLogs(), xsub.msgs);
    }

    function test_xsubmit_xblock2_chainB_succeeds() public {
        // need to submit xblock1 first, to set the streamOffset
        XTypes.Submission memory xsub1 = readXSubmission("xblock1", chainBId);
        chainBPortal.xsubmit(xsub1);

        XTypes.Submission memory xsub2 = readXSubmission("xblock2", chainBId);

        uint64 sourceChainId = xsub2.blockHeader.sourceChainId;
        uint64 expectedOffset = xsub2.msgs[xsub2.msgs.length - 1].streamOffset;
        uint256 expectedCount = numIncrements(xsub1.msgs) + numIncrements(xsub2.msgs);

        vm.prank(relayer);
        vm.recordLogs();
        expectCalls(xsub2.msgs);
        chainBPortal.xsubmit(xsub2);

        assertEq(chainBPortal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(chainBPortal.inXStreamBlockHeight(sourceChainId), xsub2.blockHeader.blockHeight);
        assertEq(chainBCounter.count(), expectedCount);
        assertEq(chainBCounter.countByChainId(sourceChainId), expectedCount);
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

        // need to resign invalid root, to pass the quorum check
        xsub.signatures = getSignatures(genesisValSetId, xsub.attestationRoot);

        vm.expectRevert("OmniPortal: invalid proof");
        portal.xsubmit(xsub);
    }

    function test_xsubmit_noQuorum_reverts() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());

        // remove last two signatures, to fail the quorum check
        Validators.SigTuple[] memory sigs = new Validators.SigTuple[](2);
        sigs[0] = xsub.signatures[0];
        sigs[1] = xsub.signatures[1];

        xsub.signatures = sigs;

        vm.expectRevert("OmniPortal: no quorum");
        portal.xsubmit(xsub);
    }

    function test_xsubmit_duplicateValidator_reverts() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());

        // add duplicate validator
        xsub.signatures[1] = xsub.signatures[0];

        vm.expectRevert("OmniPortal: duplicate validator");
        portal.xsubmit(xsub);
    }

    function test_xsubmit_invalidMsgs_reverts() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());

        // set invalid msg data, so proof fails
        xsub.msgs[0].data = abi.encodeWithSignature("invalid()");

        vm.expectRevert("OmniPortal: invalid proof");
        portal.xsubmit(xsub);
    }
}
