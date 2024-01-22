// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Test } from "forge-std/Test.sol";
import { OmniPortal } from "src/OmniPortal.sol";
import { Counter } from "test/common/Counter.sol";
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

    OmniPortal portal;
    Counter counter;

    function setUp() public {
        deployer = makeAddr("deployer");
        xcaller = makeAddr("xcaller");

        vm.startPrank(deployer);
        portal = new OmniPortal();
        counter = new Counter();
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
}
