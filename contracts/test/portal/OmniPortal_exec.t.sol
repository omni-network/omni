// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniPortal_exec_Test
 * @dev Test of OmniPortal._exec, an internal function made public for testing
 */
contract OmniPortal_exec_Test is Base {
    /// @dev Test that exec of a valid XMsg succeeds, and emits the correct XReceipt
    function test_exec_xmsg_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);

        uint256 count = counter.count();
        uint256 countForChain = counter.countByChainId(xmsg.sourceChainId);

        vm.prank(relayer);
        vm.expectCall(xmsg.to, xmsg.data);
        vm.recordLogs();
        portal.exec(xmsg);

        assertEq(counter.count(), count + 1);
        assertEq(counter.countByChainId(xmsg.sourceChainId), countForChain + 1);
        assertEq(portal.inXStreamOffset(xmsg.sourceChainId), xmsg.streamOffset);
        assertReceipt(vm.getRecordedLogs()[0], xmsg);
    }

    /// @dev Test that exec of an XMsg that reverts succeeds, and emits the correct XReceipt
    function test_exec_xmsgRevert_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_revert(1);

        vm.prank(relayer);
        vm.expectCall(xmsg.to, xmsg.data);
        vm.recordLogs();
        portal.exec(xmsg);

        assertEq(portal.inXStreamOffset(xmsg.sourceChainId), xmsg.streamOffset);
        assertReceipt(vm.getRecordedLogs()[0], xmsg);
    }

    /// @dev Test that exec of an XMsg with the wrong destChainId reverts
    function test_exec_wrongChainId_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);

        xmsg.destChainId = xmsg.destChainId + 1; // intentionally wrong chainId

        vm.expectRevert("OmniPortal: wrong destChainId");
        portal.exec(xmsg);
    }

    /// @dev Test that exec of an XMsg ahead of the current offset reverts
    function test_exec_aheadOffset_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);

        xmsg.streamOffset = xmsg.streamOffset + 1; // intentionally ahead of offset

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.exec(xmsg);
    }

    /// @dev Test that exec of an XMsg behind the current offset reverts
    function test_exec_behindOffset_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);

        portal.exec(xmsg); // execute, to increment offset

        vm.expectRevert("OmniPortal: wrong streamOffset");
        portal.exec(xmsg);
    }
}
