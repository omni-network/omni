// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { Base } from "./common/Base.sol";
import { TestXTypes } from "./common/TestXTypes.sol";
import { Reverter } from "./common/Reverter.sol";
import { GasGuzzler } from "./common/GasGuzzler.sol";
import { Vm } from "forge-std/Vm.sol";
import { console } from "forge-std/console.sol";

/**
 * @title OmniPortal_exec_Test
 * @dev Test of OmniPortal._exec, an internal function made public for testing
 */
contract OmniPortal_exec_Test is Base {
    /// @dev Test that exec of a valid XMsg succeeds, and emits the correct XReceipt
    function test_exec_xmsg_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);
        XTypes.BlockHeader memory xheader = _xheader(xmsg);

        uint256 count = counter.count();
        uint256 countForChain = counter.countByChainId(xheader.sourceChainId);

        vm.prank(relayer);
        vm.expectCall(xmsg.to, xmsg.data);
        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xheader, xmsg);

        assertEq(counter.count(), count + 1);
        assertEq(counter.countByChainId(xheader.sourceChainId), countForChain + 1);
        assertEq(portal.inXMsgOffset(xheader.sourceChainId, xmsg.shardId), xmsg.offset);
        assertReceipt(vm.getRecordedLogs()[0], xmsg, xheader.sourceChainId);
    }

    /// @dev Test that exec of an XMsg that reverts succeeds, and emits the correct XReceipt
    function test_exec_xmsgRevert_succeeds() public {
        XTypes.Msg memory xmsg = _inbound_revert(1);
        XTypes.BlockHeader memory xheader = _xheader(xmsg);

        vm.prank(relayer);
        vm.expectCall(xmsg.to, xmsg.data);
        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xheader, xmsg);

        assertEq(portal.inXMsgOffset(xheader.sourceChainId, xmsg.shardId), xmsg.offset);
        assertReceipt(vm.getRecordedLogs()[0], xmsg, xheader.sourceChainId);
    }

    /// @dev Test that exec of an XMsg with the wrong destChainId reverts
    function test_exec_wrongDestChainId_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);
        XTypes.BlockHeader memory xheader = _xheader(xmsg);

        uint64 destChainId = xmsg.destChainId;
        xmsg.destChainId = destChainId + 1; // intentionally wrong chainId

        vm.expectRevert("OmniPortal: wrong dest chain");
        vm.chainId(destChainId);
        portal.exec(xheader, xmsg);
    }

    /// @dev Test that exec of an XMsg ahead of the current offset reverts
    function test_exec_aheadOffset_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);
        XTypes.BlockHeader memory xheader = _xheader(xmsg);

        xmsg.offset = xmsg.offset + 1; // intentionally ahead of offset

        vm.expectRevert("OmniPortal: wrong offset");
        vm.chainId(xmsg.destChainId);
        portal.exec(xheader, xmsg);
    }

    /// @dev Test that exec of an XMsg behind the current offset reverts
    function test_exec_behindOffset_reverts() public {
        XTypes.Msg memory xmsg = _inbound_increment(1);
        XTypes.BlockHeader memory xheader = _xheader(xmsg);

        vm.chainId(xmsg.destChainId);
        portal.exec(xheader, xmsg); // execute, to increment offset

        vm.expectRevert("OmniPortal: wrong offset");
        vm.chainId(xmsg.destChainId);
        portal.exec(xheader, xmsg);
    }

    /// @dev Test that when an XMsg execution reverts, the correct error bytes are included in the receipt
    function test_exec_errorSize() public {
        //
        // Reverter.forceRevert() - empty error
        //
        XTypes.Msg memory xmsg = _reverter_xmsg({
            sourceChainId: chainAId,
            destChainId: thisChainId,
            offset: 1,
            data: abi.encodeWithSelector(Reverter.forceRevert.selector)
        });
        XTypes.BlockHeader memory xheader = _xheader(xmsg);

        vm.recordLogs();
        vm.chainId(xmsg.destChainId);
        portal.exec(xheader, xmsg);
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
        portal.exec(xheader, xmsg);
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
        portal.exec(xheader, xmsg);
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
        portal.exec(xheader, xmsg);
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
        portal.exec(xheader, xmsg);
        logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);
        receipt = parseReceipt(logs[0]);

        // assert error is truncated to max length
        assertEq(receipt.error.length, portal.xreceiptMaxErrorSize());
    }

    /// @dev Test that syscall that reverts forwards the revert
    function test_syscall_forwardsRevert() public {
        // We use OmniPortal.addValidatorSet, as that is a valid system call

        XTypes.Validator[] memory emptyValSet; // doesn't matter
        uint64 valSetId = 3; // doesn't matter
        bytes memory data = abi.encodeWithSelector(OmniPortal.addValidatorSet.selector, valSetId, emptyValSet);

        // this will revert with "only cchain" because _xmsg.sourceChainId won't be set
        vm.expectRevert("OmniPortal: only cchain");
        portal.syscall(data);
    }

    /// @dev Test that a call that has not been given enough gas to execute reverts,
    ///       rather than suceeding with XReceipt(success=false)
    function test_call_notEnoughGas_reverts() public {
        GasGuzzler gasGuzzler = new GasGuzzler();

        address to = address(gasGuzzler);
        // we use a high enough gas limit, such that 1/64th of it is enough to cover the "non-xmsg" gas usage of the
        // exec function
        uint64 gasLimit = 5_000_000;
        bytes memory data = abi.encodeWithSelector(GasGuzzler.guzzle.selector);

        // just xmsg gasLimit is insufficient
        uint256 insufficientGas = gasLimit;
        vm.expectRevert(); // reverts with invalid()
        portal.call{ gas: insufficientGas }(to, gasLimit, data);
    }

    /// @dev Helper to repeat a string a number of times
    function _repeat(string memory s, uint256 n) internal pure returns (string memory) {
        string memory result = "";
        for (uint256 i = 0; i < n; i++) {
            result = string.concat(result, s);
        }
        return result;
    }

    // @dev Helper to create a XBlock header for an xmsg
    function _xheader(XTypes.Msg memory xmsg) internal pure returns (XTypes.BlockHeader memory) {
        return XTypes.BlockHeader({
            sourceChainId: chainAId,
            consensusChainId: omniCChainID,
            confLevel: uint8(xmsg.shardId),
            offset: 1,
            sourceBlockHeight: 100,
            sourceBlockHash: bytes32(0)
        });
    }
}
