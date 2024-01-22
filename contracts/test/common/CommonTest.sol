// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Test } from "forge-std/Test.sol";
import { TestPortal } from "test/common/TestPortal.sol";
import { Counter } from "test/common/Counter.sol";
import { Reverter } from "test/common/Reverter.sol";
import { Events } from "test/common/Events.sol";
import { XChain } from "src/libraries/XChain.sol";

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

    TestPortal portal;
    Counter counter;
    Reverter reverter;

    function setUp() public {
        deployer = makeAddr("deployer");
        xcaller = makeAddr("xcaller");
        relayer = makeAddr("relayer");

        vm.startPrank(deployer);
        portal = new TestPortal();
        counter = new Counter();
        reverter = new Reverter();
        vm.stopPrank();
    }

    /**
     * Test XMsgs
     */

    /// @dev Get XMsg fields for an outbound Counter.increment() xcall
    function _outbound_increment() internal view returns (XChain.Msg memory) {
        return XChain.Msg({
            sourceChainId: portal.chainId(),
            destChainId: otherChainId,
            streamOffset: portal.outXStreamOffset(otherChainId),
            sender: address(counter),
            to: address(counter),
            data: abi.encodeWithSignature("increment()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Get XMsg fields for an inbound Counter.increment() xmsg
    function _inbound_increment() internal view returns (XChain.Msg memory) {
        return XChain.Msg({
            sourceChainId: otherChainId,
            destChainId: portal.chainId(),
            streamOffset: portal.inXStreamOffset(otherChainId),
            sender: address(counter),
            to: address(counter),
            data: abi.encodeWithSignature("increment()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Get XMsg fields for an inbound Reverter.revert() xmsg
    function _inbound_revert() internal view returns (XChain.Msg memory) {
        return XChain.Msg({
            sourceChainId: otherChainId,
            destChainId: portal.chainId(),
            streamOffset: portal.inXStreamOffset(otherChainId),
            sender: address(reverter),
            to: address(reverter),
            data: abi.encodeWithSignature("forceRevert()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }
}
