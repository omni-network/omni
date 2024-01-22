// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { CommonTest } from "test/common/CommonTest.sol";
import { XChain } from "src/libraries/XChain.sol";

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
}
