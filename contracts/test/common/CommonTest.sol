// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Test } from "forge-std/Test.sol";
import { OmniPortal } from "src/OmniPortal.sol";
import { Counter } from "test/common/Counter.sol";
import { Reverter } from "test/common/Reverter.sol";
import { Events } from "test/common/Events.sol";
import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title CommonTest
 * @dev An extension of forge Test that includes Omni specifc setup, utils, and events.
 */
contract CommonTest is Test, Events {
    // some other chain id to use for source or dest of xmsgs
    uint64 constant otherChainId = 2;

    address deployer;
    address xcaller;
    address relayer;

    OmniPortal portal;
    Counter counter;
    Reverter reverter;

    function setUp() public {
        deployer = makeAddr("deployer");
        xcaller = makeAddr("xcaller");
        relayer = makeAddr("relayer");

        vm.startPrank(deployer);
        portal = new OmniPortal();
        counter = new Counter();
        reverter = new Reverter();
        vm.stopPrank();
    }

    /**
     * Test XMsgs
     */

    /// @dev Get XMsg fields for an outbound increment() xcall
    function _outbound_increment() internal view returns (XTypes.Msg memory) {
        return XTypes.Msg({
            sourceChainId: portal.chainId(),
            destChainId: otherChainId,
            streamOffset: portal.outXStreamOffset(otherChainId),
            sender: address(counter),
            to: address(counter),
            data: abi.encodeWithSignature("increment()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Get XMsg fields for an inbound incrementTimes() xmsg
    function _inbound_increment() internal view returns (XTypes.Msg memory) {
        return XTypes.Msg({
            sourceChainId: otherChainId,
            destChainId: portal.chainId(),
            streamOffset: portal.inXStreamOffset(otherChainId),
            sender: address(counter),
            to: address(counter),
            data: abi.encodeWithSignature("increment()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Get XMsg fields for an inbound increment(uint256) xmsg
    function _inbound_incrementTimes(uint256 times) internal view returns (XTypes.Msg memory) {
        return XTypes.Msg({
            sourceChainId: otherChainId,
            destChainId: portal.chainId(),
            streamOffset: portal.inXStreamOffset(otherChainId),
            sender: address(counter),
            to: address(counter),
            data: abi.encodeWithSignature("incrementTimes(uint256)", times),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Get XMsg fields for an inbound revertWithReason(string) xmsg
    function _inbound_revertWithReason(string memory reason) internal view returns (XTypes.Msg memory) {
        return XTypes.Msg({
            sourceChainId: otherChainId,
            destChainId: portal.chainId(),
            streamOffset: portal.inXStreamOffset(otherChainId),
            sender: address(reverter),
            to: address(reverter),
            data: abi.encodeWithSignature("revertWithReason(string)", reason),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Get XMsg fields for an inbound failRequireWithReason(string) xmsg
    function _inbound_failRequireWithReason(string memory reason) internal view returns (XTypes.Msg memory) {
        return XTypes.Msg({
            sourceChainId: otherChainId,
            destChainId: portal.chainId(),
            streamOffset: portal.inXStreamOffset(otherChainId),
            sender: address(reverter),
            to: address(reverter),
            data: abi.encodeWithSignature("failRequireWithReason(string)", reason),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }
}
