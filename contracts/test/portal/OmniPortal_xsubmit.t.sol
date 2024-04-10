// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { OmniPortal } from "src/protocol/OmniPortal.sol";

import { Base } from "./common/Base.sol";
import { Counter } from "./common/Counter.sol";
import { Vm } from "forge-std/Vm.sol";
import { console } from "forge-std/console.sol";

/**
 * @title OmniPortal_xsubmit_Test
 * @dev Tests of OmniPortal.xsubmit
 */
contract OmniPortal_xsubmit_Test is Base {
    function test_xsubmit_xblock1_succeeds() public {
        _testSubmitXBlock("xblock1", genesisValSetId, portal, counter);
    }

    function test_xsubmit_xblock2_succeeds() public {
        // need to submit xblock1 first, to set the streamOffset
        XTypes.Submission memory xsub1 = readXSubmission("xblock1", portal.chainId());
        portal.xsubmit(xsub1);

        _testSubmitXBlock("xblock2", genesisValSetId, portal, counter);
    }

    function test_xsubmit_xblock1_chainB_succeeds() public {
        _testSubmitXBlock("xblock1", genesisValSetId, chainBPortal, chainBCounter);
    }

    function test_xsubmit_xblock2_chainB_succeeds() public {
        // need to submit xblock1 first, to set the streamOffset
        XTypes.Submission memory xsub1 = readXSubmission("xblock1", chainBId);
        chainBPortal.xsubmit(xsub1);

        _testSubmitXBlock("xblock2", genesisValSetId, chainBPortal, chainBCounter);
    }

    function test_xsubmit_noXmsgs_reverts() public {
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId());
        xsub.msgs = new XTypes.Msg[](0);

        vm.expectRevert("OmniPortal: no xmsgs");
        portal.xsubmit(xsub);
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
        XTypes.SigTuple[] memory sigs = new XTypes.SigTuple[](2);
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

    function test_xsubmit_addValidatorSet_succeeds() public {
        XTypes.Submission memory xsub = readXSubmission("addValSet2", broadcastChainId);
        _testXSubmit(portal, xsub);

        // test that validatorSet[2] is set correctly
        uint64 valSet2Id = 2;
        XTypes.Validator[] storage valSet2 = validatorSet[valSet2Id];
        uint64 totalPower;

        for (uint256 i = 0; i < valSet2.length; i++) {
            totalPower += valSet2[i].power;
            assertEq(portal.validatorSet(valSet2Id, valSet2[i].addr), valSet2[i].power);
        }

        assertEq(portal.validatorSetTotalPower(valSet2Id), totalPower);

        // test that we can submit a block with the new validatorSet
        _testSubmitXBlock("xblock1", valSet2Id, portal, counter);
    }

    /// @dev test that an xsubmission from a source chain can still use the last valSetId, if an
    ///      xsubmission with the new valSetId has not been submitted for that source chain
    function test_xsubmit_notNewValSet_succeeds() public {
        // add new validator set
        XTypes.Submission memory xsub = readXSubmission("addValSet2", broadcastChainId);
        _testXSubmit(portal, xsub);

        // test that we can submit a block with the genesisValSetId
        _testSubmitXBlock("xblock1", genesisValSetId, portal, counter);
    }

    /// @dev test that an xsubmission from a source chain cannot use an old valSetId, if an
    ///      xsubmission with a newer valSetId has been submitted for that source chain
    function test_xsubmit_oldValSet_reverts() public {
        // add new validator set
        XTypes.Submission memory xsub = readXSubmission("addValSet2", broadcastChainId);
        _testXSubmit(portal, xsub);

        // submit a block with the valSetId 2
        _testSubmitXBlock("xblock1", 2, portal, counter);

        // test that we cannot submit a block with the genesisValSetId
        xsub = readXSubmission("xblock1", portal.chainId(), genesisValSetId);

        vm.expectRevert("OmniPortal: old val set");
        portal.xsubmit(xsub);
    }

    // test that an xsubmission with an unknown valSetId reverts
    function test_xsubmit_uknownValSetId_reverts() public {
        // generate an xsubmission for val set 2, without submitting the val set
        XTypes.Submission memory xsub = readXSubmission("xblock1", portal.chainId(), 2);

        vm.expectRevert("OmniPortal: unknown val set");
        portal.xsubmit(xsub);
    }

    /// @dev helper to test that an xsubmission makes the appropriate calls (to counter_), and emits
    ///      the correct receipts.
    function _testSubmitXBlock(string memory name, uint64 valSetId, OmniPortal portal_, Counter counter_) internal {
        XTypes.Submission memory xsub = readXSubmission(name, portal_.chainId(), valSetId);

        uint256 expectedCount = numIncrements(xsub.msgs) + counter_.count();
        uint64 sourceChainId = xsub.blockHeader.sourceChainId;

        _testXSubmit(portal_, xsub);

        assertEq(counter_.count(), expectedCount);
        assertEq(counter_.countByChainId(sourceChainId), expectedCount);
    }

    /// @dev Helper to test and submit an XSubmission. It verifies gas usage matches
    ///      portal.estimageGas(...), verifies all xcalls are made, and verifies stream
    ///      offsets and block heights are updated.
    function _testXSubmit(OmniPortal portal_, XTypes.Submission memory xsub) internal {
        uint64 sourceChainId = xsub.blockHeader.sourceChainId;
        uint64 expectedOffset = xsub.msgs[xsub.msgs.length - 1].streamOffset;

        // relayer will use this call to estimate gas for an xsubmit
        // this gas limit must be >= the gas used by the xsubmit
        uint256 gasLimit = portal_.estimateGas(xsub);

        bool isCChain = sourceChainId == portal_.omniCChainID();

        // cchain xcalls are executed at the portal, and not detected by vm.expectCall()
        if (!isCChain) {
            expectCalls(xsub.msgs);
            vm.recordLogs();
        }

        uint256 gasUsed = gasleft();

        vm.prank(relayer);
        portal_.xsubmit(xsub);

        gasUsed = gasUsed - gasleft();

        // log gas usage
        console.log("gasUsed: ", gasUsed);
        console.log("gasLimit: ", gasLimit);

        require(gasUsed < gasLimit, "gasLimit exceeded");
        console.log("gasLimit - gasUsed: ", gasLimit - gasUsed);

        assertEq(portal_.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(portal_.inXStreamBlockHeight(sourceChainId), xsub.blockHeader.blockHeight);

        if (!isCChain) {
            assertReceipts(vm.getRecordedLogs(), xsub.msgs);
        }
    }
}
