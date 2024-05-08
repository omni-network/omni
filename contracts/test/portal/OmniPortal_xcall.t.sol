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

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data);
        uint64 offset = 1;

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(xmsg.destChainId, offset, xcaller, xmsg.to, xmsg.data, portal.xmsgDefaultGasLimit(), fee);

        // make xcall
        vm.prank(xcaller);
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data);

        // check offsets are incremented
        assertEq(portal.outXStreamOffset(xmsg.destChainId), 1);
        assertEq(portal.xblockOffset(), 1);
    }

    /// @dev Test that xcall with explicit gas limit emits XMsg event and increments outXStreamOffset
    function test_xcall_explicitGasLimit_succeeds() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        xmsg.gasLimit = portal.xmsgDefaultGasLimit() + 1;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);
        uint64 offset = 1;

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(xmsg.destChainId, offset, xcaller, xmsg.to, xmsg.data, xmsg.gasLimit, fee);

        // make xcall
        vm.prank(xcaller);
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);

        // check offsets are incremented
        assertEq(portal.outXStreamOffset(xmsg.destChainId), 1);
        assertEq(portal.xblockOffset(), 1);
    }

    function test_xblockOffset_incremented() public {
        address to = makeAddr("to");
        uint256 blockNum = block.number;

        // make multiple xcalls
        portal.xcall{ value: 1 ether }(chainAId, to, abi.encodeWithSignature("test()"));
        portal.xcall{ value: 1 ether }(chainAId, to, abi.encodeWithSignature("test()"));

        // check xblockOffset is incremented once
        assertEq(portal.xblockOffset(), 1);

        // move to later block
        vm.roll(blockNum + 10);
        blockNum = block.number;

        // make multiple xcalls
        portal.xcall{ value: 1 ether }(chainAId, to, abi.encodeWithSignature("test()"));
        portal.xcall{ value: 1 ether }(chainAId, to, abi.encodeWithSignature("test()"));

        // check xblockOffset is incremented once
        assertEq(portal.xblockOffset(), 2);

        // move to later block
        vm.roll(blockNum + 100);
        blockNum = block.number;

        // make single xcalls
        portal.xcall{ value: 1 ether }(chainAId, to, abi.encodeWithSignature("test()"));

        // check xblockOffset is incremented once
        assertEq(portal.xblockOffset(), 3);

        // move to one later block
        vm.roll(blockNum + 1);
        blockNum = block.number;

        // make multiple xcalls
        portal.xcall{ value: 1 ether }(chainAId, to, abi.encodeWithSignature("test()"));
        portal.xcall{ value: 1 ether }(chainAId, to, abi.encodeWithSignature("test()"));
        portal.xcall{ value: 1 ether }(chainAId, to, abi.encodeWithSignature("test()"));

        // check xblockOffset is incremented once
        assertEq(portal.xblockOffset(), 4);
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
        xmsg.gasLimit = portal.xmsgMinGasLimit() - 1;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: gasLimit too low");
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with too-high gas limit reverts
    function test_xcall_gasLimitTooHigh_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        xmsg.gasLimit = portal.xmsgMaxGasLimit() + 1;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: gasLimit too high");
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with destChainId == portal.chainId reverts
    function test_xcall_sameChain_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        xmsg.destChainId = thisChainId;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: no same-chain xcall");
        vm.chainId(thisChainId);
        portal.xcall{ value: fee }(xmsg.destChainId, xmsg.to, xmsg.data, xmsg.gasLimit);
    }
}
