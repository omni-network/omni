// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { Base } from "./common/Base.sol";
import { Vm } from "forge-std/Vm.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

/**
 * @title OmniPortal_xcall_Test
 * @dev Tests of OmniPortal.xcall
 */
contract OmniPortal_xcall_Test is Base {
    /// @dev Test that xcall emits XMsg event and increments outXMsgOffset
    function test_xcall_succeeds() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        uint8 conf = uint8(xmsg.shardId);

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);
        uint64 offset = 1;

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(xmsg.destChainId, uint64(conf), offset, xcaller, xmsg.to, xmsg.data, xmsg.gasLimit, fee);

        // make xcall
        vm.prank(xcaller);
        vm.chainId(thisChainId);
        portal.xcall{ value: fee }(xmsg.destChainId, conf, xmsg.to, xmsg.data, xmsg.gasLimit);

        // check outXMsgOffset is incremented
        assertEq(portal.outXMsgOffset(xmsg.destChainId, xmsg.shardId), 1);
    }

    /// @dev Test that xcall with insufficient fee revert
    function test_xcall_insufficientFee_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        uint8 conf = uint8(xmsg.shardId);

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit) - 1;

        vm.expectRevert("OmniPortal: insufficient fee");
        vm.chainId(thisChainId);
        portal.xcall{ value: fee }(xmsg.destChainId, conf, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with too-low gas limit reverts
    function test_xcall_gasLimitTooLow_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        uint8 conf = uint8(xmsg.shardId);

        xmsg.gasLimit = portal.xmsgMinGasLimit() - 1;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: gasLimit too low");
        vm.chainId(thisChainId);
        portal.xcall{ value: fee }(xmsg.destChainId, conf, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with too-high gas limit reverts
    function test_xcall_gasLimitTooHigh_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        uint8 conf = uint8(xmsg.shardId);

        xmsg.gasLimit = portal.xmsgMaxGasLimit() + 1;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: gasLimit too high");
        vm.chainId(thisChainId);
        portal.xcall{ value: fee }(xmsg.destChainId, conf, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with destChainId == portal.chainId reverts
    function test_xcall_sameChain_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        uint8 conf = uint8(xmsg.shardId);

        xmsg.destChainId = thisChainId;

        uint256 fee = portal.feeFor(xmsg.destChainId, xmsg.data, xmsg.gasLimit);

        vm.expectRevert("OmniPortal: unsupported dest");
        vm.chainId(thisChainId);
        portal.xcall{ value: fee }(xmsg.destChainId, conf, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with unsupported destChainId reverts
    function test_xcall_unsupportedDest_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();
        uint8 conf = uint8(xmsg.shardId);

        xmsg.destChainId = chainAId + chainBId + thisChainId;

        vm.expectRevert("OmniPortal: unsupported dest");
        vm.chainId(thisChainId);
        portal.xcall{ value: 10 ether }(xmsg.destChainId, conf, xmsg.to, xmsg.data, xmsg.gasLimit);
    }

    /// @dev Test that xcall with unsupported conf level reverts
    function test_xcall_unsupportedConf_reverts() public {
        XTypes.Msg memory xmsg = _outbound_increment();

        vm.expectRevert("OmniPortal: unsupported shard");
        vm.chainId(thisChainId);
        portal.xcall{ value: 10 ether }(
            xmsg.destChainId,
            2, // not added in MockXRegistryReplica, Finalized == 1 and Latest == 4
            xmsg.to,
            xmsg.data,
            xmsg.gasLimit
        );
    }
}
