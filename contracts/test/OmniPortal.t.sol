// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { CommonTest } from "test/common/CommonTest.sol";
import { XChain } from "src/libraries/XChain.sol";
import { Vm } from "forge-std/Vm.sol";

contract OmniPortal_Test is CommonTest {
    /// @dev Test that xcall with default gas limit emits XMsg event and increments outXStreamOffset
    function test_xcall_defaultGasLimit_succeeds() public {
        XChain.Msg memory xmsg = _outbound_increment();

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(xmsg.destChainId, xmsg.streamOffset, xcaller, xmsg.to, xmsg.data, xmsg.gasLimit);

        // make xcall
        vm.prank(xcaller);
        portal.xcall(xmsg.destChainId, xmsg.to, xmsg.data);

        // check outXStreamOffset is incremented
        assertEq(portal.outXStreamOffset(xmsg.destChainId), xmsg.streamOffset + 1);
    }

    /// @dev Test that xcall with explicit gas limit emits XMsg event and increments outXStreamOffset
    function test_xcall_explicitGasLimit_succeeds() public {
        XChain.Msg memory xmsg = _outbound_increment();
        xmsg.gasLimit = portal.XMSG_DEFAULT_GAS_LIMIT() + 1;

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(xmsg.destChainId, xmsg.streamOffset, xcaller, xmsg.to, xmsg.data, xmsg.gasLimit);

        // make xcall
        vm.prank(xcaller);
        portal.xcall(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);

        // check outXStreamOffset is incremented
        assertEq(portal.outXStreamOffset(xmsg.destChainId), xmsg.streamOffset + 1);
    }

    /// @dev Test that xcall with too-low gas limit reverts
    function test_xcall_gasLimitTooLow_reverts() public {
        XChain.Msg memory xmsg = _outbound_increment();
        xmsg.gasLimit = portal.XMSG_MIN_GAS_LIMIT() - 1;

        vm.expectRevert("OmniPortal: gasLimit too low");
        portal.xcall(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with too-high gas limit reverts
    function test_xcall_gasLimitTooHigh_reverts() public {
        XChain.Msg memory xmsg = _outbound_increment();
        xmsg.gasLimit = portal.XMSG_MAX_GAS_LIMIT() + 1;

        vm.expectRevert("OmniPortal: gasLimit too high");
        portal.xcall(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that an XSubmission with a single XMsg succeeds
    ///      Check that the correct XReceipt's are emitter, and stream offset is incremented.
    function test_xsubmit_xmsgSingle_succeeds() public {
        XChain.Msg[] memory xmsgs = new XChain.Msg[](1);
        xmsgs[0] = _inbound_increment();

        XChain.Submission memory submission = _xsub(xmsgs);

        uint256 count = counter.count();
        uint64 offset = portal.inXStreamOffset(otherChainId);

        vm.prank(relayer);
        vm.recordLogs();
        portal.xsubmit(submission);

        assertEq(counter.count(), count + 1);
        assertEq(portal.inXStreamOffset(otherChainId), offset + 1);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        assertEq(logs.length, xmsgs.length);

        _assertReceiptEmitted(
            logs[0],
            otherChainId,
            offset,
            relayer,
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
        uint64 offset = portal.inXStreamOffset(otherChainId);

        vm.prank(relayer);
        vm.recordLogs();
        portal.xsubmit(submission);

        assertEq(counter.count(), count + 4);
        assertEq(portal.inXStreamOffset(otherChainId), offset + 4);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        assertEq(logs.length, xmsgs.length);

        for (uint256 i = 0; i < xmsgs.length; i++) {
            _assertReceiptEmitted(
                logs[i],
                otherChainId,
                offset + uint64(i),
                relayer,
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
        uint64 offset = portal.inXStreamOffset(otherChainId);

        vm.prank(relayer);
        vm.recordLogs();
        portal.xsubmit(submission);

        assertEq(counter.count(), count + 3); // only 3, because one msg was a revert
        assertEq(portal.inXStreamOffset(otherChainId), offset + 4);

        Vm.Log[] memory logs = vm.getRecordedLogs();

        assertEq(logs.length, xmsgs.length);

        _assertReceiptEmitted(
            logs[0],
            otherChainId,
            offset,
            relayer,
            true // success
        );

        _assertReceiptEmitted(
            logs[1],
            otherChainId,
            offset + 1,
            relayer,
            true // success
        );

        // this one fails
        _assertReceiptEmitted(
            logs[2],
            otherChainId,
            offset + 2,
            relayer,
            false // success
        );

        _assertReceiptEmitted(
            logs[3],
            otherChainId,
            offset + 3,
            relayer,
            true // success
        );
    }

    /// @dev Test that an XSubmission with a batch of XMsgs that is out of order reverts
    function test_xsubmit_xmsgBatchOutOfOrder_reverts() public {
        XChain.Msg[] memory xmsgs = new XChain.Msg[](4);
        xmsgs[0] = _inbound_increment();
        xmsgs[1] = _inbound_increment();
        xmsgs[2] = _inbound_increment();
        xmsgs[3] = _inbound_increment();

        xmsgs[1].streamOffset = xmsgs[0].streamOffset + 1;
        xmsgs[2].streamOffset = xmsgs[1].streamOffset + 1;
        xmsgs[3].streamOffset = xmsgs[2].streamOffset - 1; // intentionally out of order

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

    /// @dev Assert that the log is an XReceipt event with the correct fields.
    ///      We use this helper rather than vm.expectEmit(), because gasUsed is difficult to predict.
    function _assertReceiptEmitted(
        Vm.Log memory log,
        uint64 sourceChainId,
        uint64 streamOffset,
        address relayer,
        bool success
    ) private {
        _XReceipt memory receipt = _parseReceipt(log);

        assertEq(receipt.sourceChainId, sourceChainId);
        assertEq(receipt.streamOffset, streamOffset);
        assertEq(receipt.relayer, relayer);
        assertEq(receipt.success, success);
    }

    // We define _XReceipt type here for convenience.
    // _ prefixed because event XReceipt is defined in test/common/Events.sol
    // The struct is not used in production, and is therefore not defined in XChain.sol library.
    struct _XReceipt {
        uint64 sourceChainId;
        uint64 streamOffset;
        uint256 gasUsed;
        address relayer;
        bool success;
    }

    /// @dev Parse an XReceipt log
    function _parseReceipt(Vm.Log memory log) internal returns (_XReceipt memory) {
        assertEq(log.emitter, address(portal));
        assertEq(log.topics.length, 3);
        assertEq(log.topics[0], keccak256("XReceipt(uint64,uint64,uint256,address,bool)"));

        (uint256 gasUsed, address relayer, bool success) = abi.decode(log.data, (uint256, address, bool));

        return _XReceipt({
            sourceChainId: uint64(uint256(log.topics[1])),
            streamOffset: uint64(uint256(log.topics[2])),
            gasUsed: gasUsed,
            relayer: relayer,
            success: success
        });
    }

    /// @dev Create an test XSubmission
    function _xsub(XChain.Msg[] memory xmsgs) internal pure returns (XChain.Submission memory) {
        return XChain.Submission({
            attestationRoot: bytes32(0), // TODO: still unchecked
            blockHeader: XChain.BlockHeader(0, 0, 0), // TODO: still unchecked
            msgs: xmsgs,
            proof: new bytes32[](0), // TODO: still unchecked
            proofFlags: new bool[](0), // TODO: still unchecked
            signatures: new XChain.SigTuple[](0) // TODO: still unchecked
         });
    }
}
