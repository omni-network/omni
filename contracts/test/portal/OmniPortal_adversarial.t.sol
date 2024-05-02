// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { OmniPortal } from "src/protocol/OmniPortal.sol";

import { TestXTypes } from "./common/TestXTypes.sol";
import { Base } from "./common/Base.sol";
import { Counter } from "./common/Counter.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title OmniPortal_adversarial
 * @dev Test cases for adversarial scenarios.
 */
contract OmniPortal_adversarial is Base {
    /// @dev Test than an xcall to the portal address that calls portal.xcall(...), fails.
    function test_xcallToPortal_selfXCall_fails() public {
        XTypes.Msg memory xmsg = _selfXCall(
            abi.encodeWithSignature(
                "xcall(uint64,address,bytes)", chainBId, address(5678), abi.encodeWithSignature("test()")
            )
        );

        TestXTypes.Receipt memory receipt = _exec(xmsg);

        assertFalse(receipt.success);
        assertEq(receipt.error, abi.encodeWithSignature("Error(string)", "OmniPortal: portal cannot xcall"));
    }

    /// @dev Test an xcall to the portal address that calls some `onlyOwner` function fails.
    function test_xcallToPortal_adminFunc_fails() public {
        XTypes.Msg memory xmsg = _selfXCall(abi.encodeWithSignature("collectFees(address)", address(5678)));

        TestXTypes.Receipt memory receipt = _exec(xmsg);

        assertFalse(receipt.success);
        assertEq(receipt.error, abi.encodeWithSignature("Error(string)", "Ownable: caller is not the owner"));
    }

    /// @dev Test an xcall to the portal address that calls some internal function fails.
    function test_xcallToPortal_internalFunction_fails() public {
        XTypes.Msg memory xmsg = _selfXCall(abi.encodeWithSignature("_transferOwnership(address)", address(5678)));

        TestXTypes.Receipt memory receipt = _exec(xmsg);

        assertFalse(receipt.success);
        assertEq(receipt.error, hex""); // empty error, because it just a revert (can't find matchin function signature)
    }

    /// @dev Helpler to call `portal.exec`, setting the chainId and returning the receipt.
    function _exec(XTypes.Msg memory xmsg) internal returns (TestXTypes.Receipt memory) {
        vm.recordLogs();
        vm.chainId(thisChainId);
        portal.exec(xmsg);

        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);

        return parseReceipt(logs[0]);
    }

    /// @dev Helper to return an XMsg with `to` set to the portal address.
    function _selfXCall(bytes memory data) internal view returns (XTypes.Msg memory) {
        return XTypes.Msg({
            sourceChainId: chainAId,
            destChainId: thisChainId,
            streamOffset: 1,
            sender: address(1234),
            to: address(portal),
            data: data,
            gasLimit: 100_000
        });
    }
}
