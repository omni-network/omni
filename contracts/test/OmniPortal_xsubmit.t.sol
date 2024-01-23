// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { CommonTest } from "test/common/CommonTest.sol";
import { XChain } from "src/libraries/XChain.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title OmniPortal_xsubmit_Test
 * @dev Tests of OmniPortal.xsubmit
 */
contract OmniPortal_xsubmit_Test is CommonTest {
    /// @dev Test that an XSubmission with a single XMsg succeeds
    ///      Check that the correct XReceipt's are emitter, and stream offset is incremented.
    function test_xsubmit_xmsgSingle_succeeds() public {
        XChain.Msg[] memory xmsgs = new XChain.Msg[](1);
        xmsgs[0] = _inbound_increment();

        XChain.Submission memory submission = _xsub(xmsgs);

        uint256 count = counter.count();
        uint64 sourceChainId = xmsgs[0].sourceChainId;
        uint64 offset = portal.inXStreamOffset(sourceChainId);

        vm.prank(xrelayer);
        vm.recordLogs();
        portal.xsubmit(submission);

        assertEq(counter.count(), count + 1);
        assertEq(portal.inXStreamOffset(sourceChainId), offset + 1);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        assertEq(logs.length, xmsgs.length);

        _assertReceiptEmitted(
            logs[0],
            sourceChainId,
            offset,
            xrelayer,
            true // success
        );
    }

    /// @dev Test that an XSubmission with a batch of XMsgs succeeds.
    ///      Check that the correct XReceipt's are emitter, and stream offset is incremented.
    function test_xsubmit_xmsgBatch_succeeds() public {
        XChain.Msg[] memory xmsgs = new XChain.Msg[](4);
        xmsgs[0] = _inbound_increment();
        xmsgs[1] = _inbound_increment();
        xmsgs[2] = _inbound_increment();
        xmsgs[3] = _inbound_increment();

        xmsgs[1].streamOffset = xmsgs[0].streamOffset + 1;
        xmsgs[2].streamOffset = xmsgs[1].streamOffset + 1;
        xmsgs[3].streamOffset = xmsgs[2].streamOffset + 1;

        XChain.Submission memory submission = _xsub(xmsgs);

        uint256 count = counter.count();
        uint64 sourceChainId = xmsgs[0].sourceChainId;
        uint64 offset = portal.inXStreamOffset(sourceChainId);

        vm.prank(xrelayer);
        vm.recordLogs();
        portal.xsubmit(submission);

        assertEq(counter.count(), count + 4);
        assertEq(portal.inXStreamOffset(sourceChainId), offset + 4);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        assertEq(logs.length, xmsgs.length);

        for (uint256 i = 0; i < xmsgs.length; i++) {
            _assertReceiptEmitted(
                logs[i],
                sourceChainId,
                offset + uint64(i),
                xrelayer,
                true // success
            );
        }
    }

    /// @dev Test that an XSubmission with a batch of XMsgs, in which one reverts, succeeds.
    ///      Check that the correct XReceipt's are emitter, and stream offset is incremented.
    function test_xsubmit_xmsgBatchWithRevert_succeeds() public {
        XChain.Msg[] memory xmsgs = new XChain.Msg[](4);
        xmsgs[0] = _inbound_increment();
        xmsgs[1] = _inbound_increment();
        xmsgs[2] = _inbound_revert();
        xmsgs[3] = _inbound_increment();

        xmsgs[1].streamOffset = xmsgs[0].streamOffset + 1;
        xmsgs[2].streamOffset = xmsgs[1].streamOffset + 1;
        xmsgs[3].streamOffset = xmsgs[2].streamOffset + 1;

        XChain.Submission memory submission = _xsub(xmsgs);

        uint256 count = counter.count();
        uint64 sourceChainId = xmsgs[0].sourceChainId;
        uint64 offset = portal.inXStreamOffset(sourceChainId);

        vm.prank(xrelayer);
        vm.recordLogs();
        portal.xsubmit(submission);

        assertEq(counter.count(), count + 3); // only 3, because one msg was a revert
        assertEq(portal.inXStreamOffset(sourceChainId), offset + 4);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        assertEq(logs.length, xmsgs.length);

        _assertReceiptEmitted(
            logs[0],
            sourceChainId,
            offset,
            xrelayer,
            true // success
        );

        _assertReceiptEmitted(
            logs[1],
            sourceChainId,
            offset + 1,
            xrelayer,
            true // success
        );

        // this one fails
        _assertReceiptEmitted(
            logs[2],
            sourceChainId,
            offset + 2,
            xrelayer,
            false // failure
        );

        _assertReceiptEmitted(
            logs[3],
            sourceChainId,
            offset + 3,
            xrelayer,
            true // success
        );
    }

    /// @dev Test that an XSubmission with a batch of XMsgs with an XMsg behind the current offset reverts
    function test_xsubmit_xmsgBatchOneBehindOffset_reverts() public {
        XChain.Msg[] memory xmsgs = new XChain.Msg[](4);
        xmsgs[0] = _inbound_increment();
        xmsgs[1] = _inbound_increment();
        xmsgs[2] = _inbound_increment();
        xmsgs[3] = _inbound_increment();

        xmsgs[1].streamOffset = xmsgs[0].streamOffset + 1;
        xmsgs[2].streamOffset = xmsgs[1].streamOffset + 1;
        xmsgs[3].streamOffset = xmsgs[2].streamOffset - 1; // intentionally behind offset

        XChain.Submission memory submission = _xsub(xmsgs);

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.xsubmit(submission);
    }

    /// @dev Test that an XSubmission with a batch of XMsgs with an XMsg ahead the current offset reverts
    function test_xsubmit_xmsgBatchOneAheadOffset_reverts() public {
        XChain.Msg[] memory xmsgs = new XChain.Msg[](4);
        xmsgs[0] = _inbound_increment();
        xmsgs[1] = _inbound_increment();
        xmsgs[2] = _inbound_increment();
        xmsgs[3] = _inbound_increment();

        xmsgs[1].streamOffset = xmsgs[0].streamOffset + 1;
        xmsgs[2].streamOffset = xmsgs[1].streamOffset + 1;
        xmsgs[3].streamOffset = xmsgs[2].streamOffset + 2; // intentionally ahead offset

        XChain.Submission memory submission = _xsub(xmsgs);

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.xsubmit(submission);
    }

    /// @dev Test that an XSubmission with a batch of XMsgs in which one has the wrong destChainId reverts
    function test_xsubmit_xmsgBatchWrongChainId_reverts() public {
        XChain.Msg[] memory xmsgs = new XChain.Msg[](4);
        xmsgs[0] = _inbound_increment();
        xmsgs[1] = _inbound_increment();
        xmsgs[2] = _inbound_increment();
        xmsgs[3] = _inbound_increment();

        xmsgs[1].streamOffset = xmsgs[0].streamOffset + 1;
        xmsgs[2].streamOffset = xmsgs[1].streamOffset + 1;
        xmsgs[3].streamOffset = xmsgs[2].streamOffset + 1;

        xmsgs[1].destChainId = xmsgs[0].destChainId + 1; // intentionally wrong chainId

        XChain.Submission memory submission = _xsub(xmsgs);

        vm.expectRevert("OmniPortal: wrong destChainId");
        portal.xsubmit(submission);
    }
}
