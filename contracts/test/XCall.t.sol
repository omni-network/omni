// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { CommonTest } from "test/common/CommonTest.sol";
import { XCall } from "src/libraries/XCall.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { Bytes } from "src/libraries/Bytes.sol";

contract XCall_Test is CommonTest {
    /// @dev Test that an XMsg executres properly
    function test_exec_xmsgSuccess_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_increment();

        uint256 count = counter.count();

        vm.expectCall(xmsg.to, xmsg.data);
        XTypes.Receipt memory receipt = XCall.exec(xmsg, relayer, _default_execopts());

        assertEq(receipt.sourceChainId, xmsg.sourceChainId);
        assertEq(receipt.destChainId, xmsg.destChainId);
        assertEq(receipt.streamOffset, xmsg.streamOffset);
        assertEq(receipt.relayer, relayer);
        assertEq(receipt.success, true);
        assertEq(receipt.returnData, abi.encode(count + 1));
        assertEq(counter.count(), count + 1);
    }

    /// @dev Test that an XMsg that reverts emits the correct XReceipt and increments inXStreamOffset
    function test_exec_xmsgRevert_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_revertWithReason("test");

        vm.expectCall(xmsg.to, xmsg.data);
        XTypes.Receipt memory receipt = XCall.exec(xmsg, relayer, _default_execopts());

        assertEq(receipt.sourceChainId, xmsg.sourceChainId);
        assertEq(receipt.destChainId, xmsg.destChainId);
        assertEq(receipt.streamOffset, xmsg.streamOffset);
        assertEq(receipt.relayer, relayer);
        assertEq(receipt.success, false);
        assertEq(receipt.returnData, abi.encodeWithSignature("Error(string)", "test"));
    }

    /// @dev Test that an XMsg that reverts emits the correct XReceipt and increments inXStreamOffset
    function test_exec_xmsgFailRequire_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_failRequireWithReason("test");

        vm.expectCall(xmsg.to, xmsg.data);
        XTypes.Receipt memory receipt = XCall.exec(xmsg, relayer, _default_execopts());

        assertEq(receipt.sourceChainId, xmsg.sourceChainId);
        assertEq(receipt.destChainId, xmsg.destChainId);
        assertEq(receipt.streamOffset, xmsg.streamOffset);
        assertEq(receipt.relayer, relayer);
        assertEq(receipt.success, false);
        assertEq(receipt.returnData, abi.encodeWithSignature("Error(string)", "test"));
    }

    /// @dev Test trim data is returned if it exceeds maxReturnDataSize
    function test_exec_longReturnData_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_failRequireWithReason("test");

        uint64 maxReturnDataSize = 4;
        bytes memory untrimmedReturnData = abi.encodeWithSignature("Error(string)", "test");

        // assert return data is too long
        assert(untrimmedReturnData.length > maxReturnDataSize);

        vm.expectCall(xmsg.to, xmsg.data);
        XTypes.Receipt memory receipt = XCall.exec(
            xmsg, relayer, XCall.ExecOpts({ maxReturnDataSize: maxReturnDataSize, outOfGasErrorMsg: "out of gas" })
        );

        assertEq(receipt.sourceChainId, xmsg.sourceChainId);
        assertEq(receipt.destChainId, xmsg.destChainId);
        assertEq(receipt.streamOffset, xmsg.streamOffset);
        assertEq(receipt.relayer, relayer);
        assertEq(receipt.success, false);

        // assert return data is trimmed
        assert(receipt.returnData.length == maxReturnDataSize);
        assertEq(receipt.returnData, Bytes.slice(untrimmedReturnData, 0, maxReturnDataSize));
    }

    /// @dev Test that an XMsg without a high enough gas limit fails returns an
    ///      XReceipt with correct error message in returnData
    function test_exec_outOfGas_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_increment();

        // low gas limit, so the call will fail
        xmsg.gasLimit = 10;

        uint256 count = counter.count();

        vm.expectCall(xmsg.to, xmsg.data);
        XTypes.Receipt memory receipt = XCall.exec(xmsg, relayer, _default_execopts());

        assertEq(receipt.sourceChainId, xmsg.sourceChainId);
        assertEq(receipt.destChainId, xmsg.destChainId);
        assertEq(receipt.streamOffset, xmsg.streamOffset);
        assertEq(receipt.relayer, relayer);
        assertEq(receipt.success, false);
        assertEq(receipt.returnData, abi.encodeWithSignature("Error(string)", "out of gas"));
        assertEq(counter.count(), count);
    }

    function _default_execopts() internal pure returns (XCall.ExecOpts memory) {
        return XCall.ExecOpts({ maxReturnDataSize: 256, outOfGasErrorMsg: "out of gas" });
    }
}
