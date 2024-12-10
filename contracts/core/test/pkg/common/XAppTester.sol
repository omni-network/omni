// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { XApp } from "src/pkg/XApp.sol";
import { XAppBase } from "src/pkg/XAppBase.sol";
import { XAppUpgradeable } from "src/pkg/XAppUpgradeable.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title XAppTesterBase
 * @dev Base contract for XAppTester and XAppUpgradeableTester
 */
abstract contract XAppTesterBase is XAppBase {
    /// @dev Call xcall, default conf
    function doXCall(uint64 destChainId, bytes32 to, bytes calldata data, uint64 gasLimit) public payable {
        xcall(destChainId, to, data, gasLimit);
    }

    /// @dev Call xcall, custom conf
    function doXCall(uint64 destChainId, uint8 conf, bytes32 to, bytes calldata data, uint64 gasLimit) public payable {
        xcall(destChainId, conf, to, data, gasLimit);
    }

    XTypes.MsgContext internal _expectXMsg;

    /// @dev Set an expected xmsg, checked in then next checkXRecv
    function expectXMsg(uint64 sourceChainId, bytes32 sender) public {
        _expectXMsg = XTypes.MsgContext(sourceChainId, sender);
    }

    /// @dev Assert that xmsg is as expected
    function checkXRecv() public xrecv {
        require(xmsg.sourceChainId == _expectXMsg.sourceChainId);
        require(xmsg.sender == _expectXMsg.sender);
        delete _expectXMsg;
    }

    bool internal _expectIsXCall;

    /// @dev Set the expected value for isXCall, checked in the next checkIsXCall
    function expectIsXCall(bool expect) public {
        _expectIsXCall = expect;
    }

    /// @dev Assert that isXCall is as expected
    function checkIsXCall() public {
        require(isXCall() == _expectIsXCall);
        _expectIsXCall = false;
    }

    function setDefaultConfLevel(uint8 conf) public {
        _setDefaultConfLevel(conf);
    }

    function setOmniPortal(address portal) public {
        _setOmniPortal(portal);
    }

    function getOmniChainId() public view returns (uint64) {
        return omniChainId();
    }
}

/**
 * @title XAppTester
 * @dev Test helper for XApp
 */
contract XAppTester is XAppTesterBase, XApp {
    constructor(address portal, uint8 defaultConfLevel) XApp(portal, defaultConfLevel) { }
}

/**
 * @title XAppUpgradeableTester
 * @dev Test helper for XAppUpgradeable
 */
contract XAppUpgradeableTester is XAppTesterBase, XAppUpgradeable {
    function initialize(address portal, uint8 defaultConfLevel) public initializer {
        __XApp_init(portal, defaultConfLevel);
    }
}

/**
 * @title IsXCallProxy
 * @dev Simple proxy contract to call checkIsXCall onn IXAppTester
 *      Used to confirm XApp.isXCall is false when msg.sender is not the portal
 */
contract IsXCallProxy {
    XAppTesterBase xapp;

    constructor(address xapp_) {
        xapp = XAppTesterBase(xapp_);
    }

    function checkIsXCall() public {
        xapp.checkIsXCall();
    }
}
