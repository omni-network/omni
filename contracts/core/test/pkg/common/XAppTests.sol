// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { XAppTesterBase, IsXCallProxy } from "./XAppTester.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { AddressUtils } from "src/libraries/AddressUtils.sol";
import { XApp } from "src/pkg/XApp.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title XAppTests
 * @dev Shared tests for XApp_Test and XAppUpgradeable_Test
 */
abstract contract XAppTests is Test {
    using AddressUtils for address;

    MockPortal portal;
    XAppTesterBase xapp;
    uint8 defaultConfLevel;

    function setUp() public virtual {
        portal = new MockPortal();
        defaultConfLevel = ConfLevel.Latest;
    }

    function test_xcall() public {
        uint64 destChainId = 1;
        bytes memory data = abi.encodeWithSignature("test()");
        bytes32 to = makeAddr("to").toBytes32();
        uint64 gasLimit = 100_000;

        // requires fees
        vm.expectRevert("XApp: insufficient funds");
        xapp.doXCall(destChainId, to, data, gasLimit);

        // allows user to pay
        uint256 fee = portal.feeFor(destChainId, data, gasLimit);
        address user = makeAddr("user");
        vm.deal(user, fee);
        vm.prank(user);
        _expectXCall(destChainId, defaultConfLevel, to, data, gasLimit, fee);
        xapp.doXCall{ value: fee }(destChainId, to, data, gasLimit);
        assertEq(address(xapp).balance, 0);

        // can pay fees from its balance
        vm.deal(address(xapp), fee);
        _expectXCall(destChainId, defaultConfLevel, to, data, gasLimit, fee);
        xapp.doXCall(destChainId, to, data, gasLimit);
        assertEq(address(xapp).balance, 0);

        // custom conf
        require(ConfLevel.Finalized != defaultConfLevel);
        vm.deal(address(xapp), fee); // needs to pay fees again
        _expectXCall(destChainId, ConfLevel.Finalized, to, data, gasLimit, fee);
        xapp.doXCall(destChainId, ConfLevel.Finalized, to, data, gasLimit);
    }

    function test_xrecv() public {
        uint64 sourceChainId = 1;
        bytes32 sender = makeAddr("sender").toBytes32();

        // xrecv sets xmsg
        xapp.expectXMsg(sourceChainId, sender);
        portal.mockXCall(
            sourceChainId, sender, address(xapp).toBytes32(), abi.encodeWithSignature("checkXRecv()"), 100_000
        );

        // xrecv sets zero values when not an xcall
        xapp.checkXRecv();
    }

    function test_isXCall() public {
        uint64 sourceChainId = 1;
        bytes32 sender = makeAddr("sender").toBytes32();

        // isXCall is true when xmsg.sourceChainId is set
        xapp.expectIsXCall(true);
        portal.mockXCall(
            sourceChainId, sender, address(xapp).toBytes32(), abi.encodeWithSignature("checkIsXCall()"), 100_000
        );

        // isXCall is false when xmsg.sourceChainId is not set
        xapp.expectIsXCall(false);
        xapp.checkIsXCall();

        // isXCall is false when msg.sender is not the portal
        IsXCallProxy proxy = new IsXCallProxy(address(xapp));
        xapp.expectIsXCall(false);
        portal.mockXCall(
            sourceChainId, sender, address(proxy).toBytes32(), abi.encodeWithSignature("checkIsXCall()"), 100_000
        );
    }

    function test_setDefaultConfLevel() public {
        uint8 conf = ConfLevel.Finalized;
        require(XApp(address(xapp)).defaultConfLevel() != conf);
        xapp.setDefaultConfLevel(conf);
        assertEq(XApp(address(xapp)).defaultConfLevel(), conf);
    }

    function test_setOmniPortal() public {
        address newPortal = makeAddr("newPortal");
        xapp.setOmniPortal(newPortal);
        assertEq(address(XApp(address(xapp)).omni()), newPortal);
    }

    function test_omniChainId() public view {
        assertEq(xapp.getOmniChainId(), portal.omniChainId());
    }

    /// @dev Helper function to expect an xcall to the portal
    function _expectXCall(uint64 destChainId, uint8 conf, bytes32 to, bytes memory data, uint64 gasLimit, uint256 fee)
        internal
    {
        vm.expectCall(address(portal), fee, abi.encodeCall(IOmniPortal.xcall, (destChainId, conf, to, data, gasLimit)));
    }
}
