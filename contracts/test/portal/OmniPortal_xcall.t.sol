// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { Base } from "./common/Base.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title OmniPortal_xcall_Test
 * @dev Tests of OmniPortal.xcall
 */
contract OmniPortal_xcall_Test is Base {
    /// @dev Test that xcall with default gas limit emits XMsg event and increments outXStreamOffset
    function test_xcall_defaultGasLimit_succeeds() public {
        XTypes.Msg memory xmsg = _outbound_increment();

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(xmsg.destChainId, 1, xcaller, xmsg.to, xmsg.data, portal.XMSG_DEFAULT_GAS_LIMIT());

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data);

        // make xcall
        vm.prank(xcaller);
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data);

        // check outXStreamOffset is incremented
        assertEq(portal.outXStreamOffset(xmsg.destChainId), 1);
    }

    /// @dev Test that xcall with explicit gas limit emits XMsg event and increments outXStreamOffset
    function test_xcall_explicitGasLimit_succeeds() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        xmsg.gasLimit = portal.XMSG_DEFAULT_GAS_LIMIT() + 1;

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(xmsg.destChainId, 1, xcaller, xmsg.to, xmsg.data, xmsg.gasLimit);

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        // make xcall
        vm.prank(xcaller);
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);

        // check outXStreamOffset is incremented
        assertEq(portal.outXStreamOffset(xmsg.destChainId), 1);
    }

    /// @dev Test that xcall with insufficient fee revert
    function test_xcall_insufficientFee_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit) - 1;

        vm.expectRevert("OmniPortal: insufficient fee");
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with too-low gas limit reverts
    function test_xcall_gasLimitTooLow_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        xmsg.gasLimit = portal.XMSG_MIN_GAS_LIMIT() - 1;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: gasLimit too low");
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with too-high gas limit reverts
    function test_xcall_gasLimitTooHigh_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        xmsg.gasLimit = portal.XMSG_MAX_GAS_LIMIT() + 1;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: gasLimit too high");
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with destChainId == portal.chainId reverts
    function test_xcall_sameChain_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        xmsg.destChainId = portal.chainId();

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: no same-chain xcall");
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }
}
