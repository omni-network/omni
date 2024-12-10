// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { MockPortal } from "core/test/utils/MockPortal.sol";
import { Pledge } from "src/pledge/Pledge.sol";
import { AddressUtils } from "core/src/libraries/AddressUtils.sol";
import { Test } from "forge-std/Test.sol";

contract PledgeTest is Test {
    using AddressUtils for address;

    MockPortal portal;
    Pledge pledge;

    address user = makeAddr("user");

    function setUp() public {
        portal = new MockPortal();
        pledge = new Pledge(address(portal));
    }

    /**
     * @notice Tests that the pledge reverts when the caller is not the OmniPortal.
     */
    function test_pledge_reverts() public {
        vm.deal(user, 1 ether);
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(Pledge.NotXCall.selector));
        pledge.pledge_jwkilcxtschdbaaa();
    }

    /**
     * @notice Tests that the pledge emits the Pledged event when successful.
     */
    function test_pledge_success() public {
        vm.expectEmit(address(pledge));
        emit Pledge.Pledged(user.toBytes32(), 100, block.timestamp);
        portal.mockXCall({
            sourceChainId: 100,
            sender: user.toBytes32(),
            data: abi.encodeCall(pledge.pledge_jwkilcxtschdbaaa, ()),
            to: address(pledge).toBytes32()
        });
    }
}
