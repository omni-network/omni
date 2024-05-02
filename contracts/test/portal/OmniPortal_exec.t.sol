// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { OmniPortal } from "src/protocol/OmniPortal.sol";
import { Base } from "./common/Base.sol";
import { TestXTypes } from "./common/TestXTypes.sol";
import { Reverter } from "./common/Reverter.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title OmniPortal_exec_Test
 * @dev Test of OmniPortal._exec, an internal function made public for testing
 */
contract OmniPortal_exec_Test is Base {
    /// @dev Test that exec of a valid XMsg succeeds, and emits the correct XReceipt
    function test_exec_xmsg_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);

        uint256 count = counter.count();
        uint256 countForChain = counter.countByChainId(xmsg.sourceChainId);

        vm.prank(relayer);
        vm.expectCall(xmsg.to, xmsg.data);
        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);

        assertEq(counter.count(), count + 1);
        assertEq(counter.countByChainId(xmsg.sourceChainId), countForChain + 1);
        assertEq(portal.inXStreamOffset(xmsg.sourceChainId), xmsg.streamOffset);
        assertReceipt(vm.getRecordedLogs()[0], xmsg);
    }

    /// @dev Test that exec of an XMsg that reverts succeeds, and emits the correct XReceipt
    function test_exec_xmsgRevert_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_revert(1);

        vm.prank(relayer);
        vm.expectCall(xmsg.to, xmsg.data);
        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);

        assertEq(portal.inXStreamOffset(xmsg.sourceChainId), xmsg.streamOffset);
        assertReceipt(vm.getRecordedLogs()[0], xmsg);
    }

    /// @dev Test that exec of an XMsg with the wrong destChainId reverts
    function test_exec_wrongChainId_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);

        uint64 destChainId = xmsg.destChainId;
        xmsg.destChainId = destChainId + 1; // intentionally wrong chainId

        vm.expectRevert("OmniPortal: wrong destChainId");
        vm.chainId(destChainId);
        portal.exec(xmsg);
    }

    /// @dev Test that exec of an XMsg ahead of the current offset reverts
    function test_exec_aheadOffset_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);

        xmsg.streamOffset = xmsg.streamOffset + 1; // intentionally ahead of offset

        vm.expectRevert("OmniPortal: wrong streamOffset");
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);
    }

    /// @dev Test that exec of an XMsg behind the current offset reverts
    function test_exec_behindOffset_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);

        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg); // execute, to increment offset

        vm.expectRevert("OmniPortal: wrong streamOffset");
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);
    }

    /// @dev Test that when an XMsg execution reverts, the correct error bytes are included in the receipt
    function test_exec_errorBytes() public {
        //
        // Reverter.forceRevert() - empty error
        //
        XTypes.Msg memory xmsg = _reverter_xmsg({
            sourceChainId: chainAId,
            destChainId: thisChainId,
            offset: 1,
            data: abi.encodeWithSelector(Reverter.forceRevert.selector)
        });

        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);
        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);
        TestXTypes.Receipt memory receipt = parseReceipt(logs[0]);
        assertEq(receipt.error.length, 0);

        //
        // Reverter.forceRevertWithReason("my reason")
        //
        xmsg = _reverter_xmsg({
            sourceChainId: chainAId,
            destChainId: thisChainId,
            offset: 2,
            data: abi.encodeWithSelector(Reverter.forceRevertWithReason.selector, "my reason")
        });

        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);
        logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);
        receipt = parseReceipt(logs[0]);
        assertEq(receipt.error, abi.encodeWithSignature("Error(string)", "my reason"));

        //
        // Reverter.panicUnderflow() - Panic(0x11)
        //
        xmsg = _reverter_xmsg({
            sourceChainId: chainAId,
            destChainId: thisChainId,
            offset: 3,
            data: abi.encodeWithSelector(Reverter.panicUnderflow.selector)
        });

        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);
        logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);
        receipt = parseReceipt(logs[0]);
        assertEq(receipt.error, abi.encodeWithSignature("Panic(uint256)", 0x11));

        //
        // Reverter.panicDivisionByZero() - Panic(0x12)
        //
        xmsg = _reverter_xmsg({
            sourceChainId: chainAId,
            destChainId: thisChainId,
            offset: 4,
            data: abi.encodeWithSelector(Reverter.panicDivisionByZero.selector)
        });

        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);
        logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);
        receipt = parseReceipt(logs[0]);
        assertEq(receipt.error, abi.encodeWithSignature("Panic(uint256)", 0x12));

        //
        // Reverter.revertWithReason(<really long error>)
        //
        xmsg = _reverter_xmsg({
            sourceChainId: chainAId,
            destChainId: thisChainId,
            offset: 5,
            data: abi.encodeWithSelector(Reverter.forceRevertWithReason.selector, _repeat("really long", 1000))
        });

        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xmsg);
        logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);
        receipt = parseReceipt(logs[0]);
        assertEq(receipt.error, portal.XRECEIPT_ERROR_EXCEEDS_MAX_BYTES());

        // assert constant is encoded as expected
        assertEq(receipt.error, abi.encodeWithSignature("OmniError(uint256)", 0x1));
    }

    /// @dev Test that sys exec that reverts forwards the revert
    function test_execSys_forwardsRevert() public {
        // We use OmniPortal.addValidatorSet, as that is a valid system call

        XTypes.Validator[] memory emptyValSet; // doesn't matter
        uint64 valSetId = 3; // doesn't matter
        bytes memory data = abi.encodeWithSelector(OmniPortal.addValidatorSet.selector, valSetId, emptyValSet);

        // this will revert with "only cchain" because _xmsg.sourceChainId won't be set
        vm.expectRevert("OmniPortal: only cchain");
        portal.execSys(data);
    }

    /// @dev Helper to repeat a string a number of times
    function _repeat(string memory s, uint256 n) internal pure returns (string memory) {
        string memory result = "";
        for (uint256 i = 0; i < n; i++) {
            result = string.concat(result, s);
        }
        return result;
    }
}
