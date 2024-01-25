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
    /// @dev Test that an XSubmission with a single XMsg succeeds
    ///      Check that the correct XReceipt's are emitter, and stream offset is incremented.
    function test_xsubmit_xmsgSingle_succeeds() public {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](1);

        xmsgs[0] = _inbound_increment(0);

        XTypes.Submission memory submission = _xsub(xmsgs);

        uint256 count = counter.count();
        uint64 sourceChainId = xmsgs[0].sourceChainId;
        uint64 offset = portal.inXStreamOffset(sourceChainId);

        vm.prank(relayer);
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
            relayer,
            true // success
        );
    }

    /// @dev Test that an XSubmission with a batch of XMsgs succeeds.
    ///      Check that the correct XReceipt's are emitter, and stream offset is incremented.
    function test_xsubmit_xmsgBatch_succeeds() public {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](4);

        xmsgs[0] = _inbound_increment(0);
        xmsgs[1] = _inbound_increment(1);
        xmsgs[2] = _inbound_increment(2);
        xmsgs[3] = _inbound_increment(3);

        XTypes.Submission memory submission = _xsub(xmsgs);

        uint256 count = counter.count();
        uint64 sourceChainId = xmsgs[0].sourceChainId;
        uint64 offset = portal.inXStreamOffset(sourceChainId);

        vm.prank(relayer);
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
                relayer,
                true // success
            );
        }
    }

    /// @dev Test that an XSubmission with a batch of XMsgs, in which one reverts, succeeds.
    ///      Check that the correct XReceipt's are emitter, and stream offset is incremented.
    function test_xsubmit_xmsgBatchWithRevert_succeeds() public {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](4);

        xmsgs[0] = _inbound_increment(0);
        xmsgs[1] = _inbound_increment(1);
        xmsgs[2] = _inbound_revert(2);
        xmsgs[3] = _inbound_increment(3);

        XTypes.Submission memory submission = _xsub(xmsgs);

        uint256 count = counter.count();
        uint64 sourceChainId = xmsgs[0].sourceChainId;
        uint64 offset = portal.inXStreamOffset(sourceChainId);

        vm.prank(relayer);
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
            relayer,
            true // success
        );

        _assertReceiptEmitted(
            logs[1],
            sourceChainId,
            offset + 1,
            relayer,
            true // success
        );

        // this one fails
        _assertReceiptEmitted(
            logs[2],
            sourceChainId,
            offset + 2,
            relayer,
            false // failure
        );

        _assertReceiptEmitted(
            logs[3],
            sourceChainId,
            offset + 3,
            relayer,
            true // success
        );
    }

    /// @dev Test that an XSubmission with a batch of XMsgs with an XMsg behind the current offset reverts
    function test_xsubmit_xmsgBatchOneBehindOffset_reverts() public {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](4);

        xmsgs[0] = _inbound_increment(0);
        xmsgs[1] = _inbound_increment(1);
        xmsgs[2] = _inbound_increment(2);
        xmsgs[3] = _inbound_increment(2); // intentionally behind offset

        XTypes.Submission memory submission = _xsub(xmsgs);

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.xsubmit(submission);
    }

    /// @dev Test that an XSubmission with a batch of XMsgs with an XMsg ahead the current offset reverts
    function test_xsubmit_xmsgBatchOneAheadOffset_reverts() public {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](4);

        xmsgs[0] = _inbound_increment(0);
        xmsgs[1] = _inbound_increment(1);
        xmsgs[2] = _inbound_increment(2);
        xmsgs[3] = _inbound_increment(4); // intentionally ahead offset

        XTypes.Submission memory submission = _xsub(xmsgs);

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.xsubmit(submission);
    }

    /// @dev Test that an XSubmission with a batch of XMsgs in which one has the wrong destChainId reverts
    function test_xsubmit_xmsgBatchWrongChainId_reverts() public {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](4);

        xmsgs[0] = _inbound_increment(0);
        xmsgs[1] = _inbound_increment(1);
        xmsgs[2] = _inbound_increment(2);
        xmsgs[3] = _inbound_increment(3);

        xmsgs[1].destChainId = xmsgs[0].destChainId + 1; // intentionally wrong chainId

        XTypes.Submission memory submission = _xsub(xmsgs);

        vm.expectRevert("OmniPortal: wrong destChainId");
        portal.xsubmit(submission);
    }
}
