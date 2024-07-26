// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ReentrancyGuardUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";

import { TestXTypes } from "./common/TestXTypes.sol";
import { Base } from "./common/Base.sol";
import { Counter } from "./common/Counter.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title OmniPortal_xsubmit_Test
 * @dev Tests of OmniPortal.xsubmit
 */
contract OmniPortal_xsubmit_Test is Base {
    function test_xsubmit_xblock1_succeeds() public {
        _testSubmitXBlock({
            name: "xblock1",
            destChainId: thisChainId,
            valSetId: genesisValSetId,
            portal_: portal,
            counter_: counter
        });
    }

    function test_xsubmit_xblock2_succeeds() public {
        // need to submit xblock1 first, to set the offset
        XTypes.Submission memory xsub1 = readXSubmission({ name: "xblock1", destChainId: thisChainId });
        vm.chainId(thisChainId);
        portal.xsubmit(xsub1);

        _testSubmitXBlock({
            name: "xblock2",
            destChainId: thisChainId,
            valSetId: genesisValSetId,
            portal_: portal,
            counter_: counter
        });
    }

    function test_xsubmit_xblock1_chainB_succeeds() public {
        _testSubmitXBlock({
            name: "xblock1",
            destChainId: chainBId,
            valSetId: genesisValSetId,
            portal_: chainBPortal,
            counter_: chainBCounter
        });
    }

    function test_xsubmit_xblock2_chainB_succeeds() public {
        // need to submit xblock1 first, to set the offset
        XTypes.Submission memory xsub1 = readXSubmission({ name: "xblock1", destChainId: chainBId });
        vm.chainId(chainBId);
        chainBPortal.xsubmit(xsub1);

        _testSubmitXBlock({
            name: "xblock2",
            destChainId: chainBId,
            valSetId: genesisValSetId,
            portal_: chainBPortal,
            counter_: chainBCounter
        });
    }

    function test_xsubmit_reentrancy_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "reentrancy", destChainId: thisChainId });

        vm.recordLogs();
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);

        TestXTypes.Receipt memory receipt = parseReceipt(logs[0]);
        assertEq(receipt.success, false);
        assertEq(receipt.error, abi.encodePacked(ReentrancyGuardUpgradeable.ReentrancyGuardReentrantCall.selector));
    }

    function test_xsubmit_noXmsgs_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock1", destChainId: thisChainId });
        xsub.msgs = new XTypes.Msg[](0);

        vm.expectRevert("OmniPortal: no xmsgs");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);
    }

    function test_xsubmit_wrongDestChainId_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock1", destChainId: thisChainId });

        vm.expectRevert("OmniPortal: wrong dest chain");
        vm.chainId(chainBId);
        chainBPortal.xsubmit(xsub);
    }

    function test_xsubmit_wrongConsensusChainId_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock1", destChainId: thisChainId });

        vm.expectRevert("OmniPortal: wrong cchain ID");
        xsub.blockHeader.consensusChainId = chainBId;
        chainBPortal.xsubmit(xsub);
    }

    function test_xsubmit_wrongStreamOffset_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock2", destChainId: thisChainId });

        vm.expectRevert("OmniPortal: wrong offset");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);
    }

    function test_xsubmit_invalidAttestationRoot_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock1", destChainId: thisChainId });

        xsub.attestationRoot = keccak256("invalid");

        // need to resign invalid root, to pass the quorum check
        xsub.signatures = getSignatures(genesisValSetId, xsub.attestationRoot);

        vm.expectRevert("OmniPortal: invalid proof");
        portal.xsubmit(xsub);
    }

    function test_xsubmit_noQuorum_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock1", destChainId: thisChainId });

        // remove last two signatures, to fail the quorum check
        XTypes.SigTuple[] memory sigs = new XTypes.SigTuple[](2);
        sigs[0] = xsub.signatures[0];
        sigs[1] = xsub.signatures[1];

        xsub.signatures = sigs;

        vm.expectRevert("OmniPortal: no quorum");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);
    }

    function test_xsubmit_duplicateValidator_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock1", destChainId: thisChainId });

        // add duplicate validator
        xsub.signatures[1] = xsub.signatures[0];

        vm.expectRevert("Quorum: duplicate validator");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);
    }

    function test_xsubmit_invalidMsgs_reverts() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock1", destChainId: thisChainId });

        // set invalid msg data, so proof fails
        xsub.msgs[0].data = abi.encodeWithSignature("invalid()");

        vm.expectRevert("OmniPortal: invalid proof");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);
    }

    function test_xsubmit_addValidatorSet_succeeds() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "addValSet2", destChainId: broadcastChainId });
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // test that validatorSet[2] is set correctly
        uint64 valSet2Id = 2;
        XTypes.Validator[] storage valSet2 = validatorSet[valSet2Id];
        uint64 totalPower;

        for (uint256 i = 0; i < valSet2.length; i++) {
            totalPower += valSet2[i].power;
            assertEq(portal.valSet(valSet2Id, valSet2[i].addr), valSet2[i].power);
        }

        assertEq(portal.valSetTotalPower(valSet2Id), totalPower);

        // test that we can submit a block with the new validatorSet
        _testSubmitXBlock({
            name: "xblock1",
            destChainId: thisChainId,
            valSetId: valSet2Id,
            portal_: portal,
            counter_: counter
        });
    }

    /// @dev test that an xsubmission from a source chain can still use the last valSetId, if an
    ///      xsubmission with the new valSetId has not been submitted for that source chain
    function test_xsubmit_notNewValSet_succeeds() public {
        // add new validator set
        XTypes.Submission memory xsub = readXSubmission({ name: "addValSet2", destChainId: broadcastChainId });
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);

        // test that we can submit a block with the genesisValSetId
        // _testSubmitXBlock("xblock1", genesisValSetId, portal, counter);
        _testSubmitXBlock({
            name: "xblock1",
            destChainId: thisChainId,
            valSetId: genesisValSetId,
            portal_: portal,
            counter_: counter
        });
    }

    /// @dev test that an xsubmission from a source chain cannot use an old valSetId, if an
    ///      xsubmission with a newer valSetId has been submitted for that source chain
    function test_xsubmit_oldValSet_reverts() public {
        portal.setLatestValSetId(genesisValSetId + portal.xsubValsetCutoff());

        // test that we cannot submit a block with the genesisValSetId
        XTypes.Submission memory xsub =
            readXSubmission({ name: "xblock1", destChainId: thisChainId, valSetId: genesisValSetId });

        vm.expectRevert("OmniPortal: old val set");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);
    }

    function test_xsubmit_unknownValSetId_reverts() public {
        // generate an xsubmission for val set 2, without submitting the val set
        XTypes.Submission memory xsub = readXSubmission({ name: "xblock1", destChainId: thisChainId, valSetId: 2 });

        vm.expectRevert("OmniPortal: unknown val set");
        vm.chainId(thisChainId);
        portal.xsubmit(xsub);
    }

    /// @dev helper to test that an xsubmission makes the appropriate calls (to counter_), and emits
    ///      the correct receipts
    function _testSubmitXBlock(
        string memory name,
        uint64 valSetId,
        uint64 destChainId,
        IOmniPortal portal_,
        Counter counter_
    ) internal {
        XTypes.Submission memory xsub = readXSubmission(name, destChainId, valSetId);

        uint64 sourceChainId = xsub.blockHeader.sourceChainId;
        uint64 shardId = xsub.blockHeader.confLevel; // conf level is shard id
        uint64 expectedOffset = xsub.msgs[xsub.msgs.length - 1].offset;
        uint256 expectedCount = numIncrements(xsub.msgs) + counter_.count();

        vm.recordLogs();
        expectCalls(xsub.msgs);

        vm.prank(relayer);
        vm.chainId(destChainId);
        portal_.xsubmit{ gas: _xsubGasLimit(xsub) }(xsub);

        assertEq(portal_.inXMsgOffset(sourceChainId, shardId), expectedOffset);
        assertEq(portal_.inXBlockOffset(sourceChainId, shardId), xsub.blockHeader.offset);
        assertEq(counter_.count(), expectedCount);
        assertEq(counter_.countByChainId(sourceChainId), expectedCount);
        assertReceipts(vm.getRecordedLogs(), xsub.msgs, sourceChainId);
    }

    /// @dev A simple algo to over-esimate the gas limit for an xsubmission
    ///      We include this in tests because we mirror this estimation in the relayer, and should
    //       therefore confirm the over-estimation is appropriate.
    function _xsubGasLimit(XTypes.Submission memory xsub) internal pure returns (uint256) {
        // start with a base 500k gas limit
        uint256 gasLimit = 500_000;

        XTypes.Msg memory xmsg;

        // add gas limit for each xmsg
        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            xmsg = xsub.msgs[i];

            if (xmsg.gasLimit > 0) {
                // add application defined xmsg gas limit
                gasLimit += xmsg.gasLimit;

                // add additional 100k per xmsg, to cover overhead
                gasLimit += 100_000;
            } else {
                // only system calls can have 0 gas limit
                // for these, we add 1M gas, as a safe over-estimate
                //
                // sys calls currently only come from the consesus chain, with a single xmsg per
                // submission. these xblock submissions do not need to be split, and therefore do
                // not need accurate gas estimation.
                gasLimit += 1_000_000;
            }
        }

        return gasLimit;
    }
}
