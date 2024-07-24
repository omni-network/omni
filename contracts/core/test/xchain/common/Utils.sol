// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Test } from "forge-std/Test.sol";
import { Vm } from "forge-std/Vm.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { Events } from "./Events.sol";
import { TestXTypes } from "./TestXTypes.sol";
import { Fixtures } from "./Fixtures.sol";

/**
 * @title Utils
 * @dev Defines test utilities.
 */
contract Utils is Test, Events, Fixtures {
    /// @dev Parse an XReceipt log
    function parseReceipt(Vm.Log memory log) internal returns (TestXTypes.Receipt memory) {
        assertEq(log.topics.length, 4);
        assertEq(log.topics[0], XReceipt.selector);

        (uint256 gasUsed, address relayer, bool success, bytes memory errorBytes) =
            abi.decode(log.data, (uint256, address, bool, bytes));

        return TestXTypes.Receipt({
            sourceChainId: uint64(uint256(log.topics[1])),
            shardId: uint64(uint256(log.topics[2])),
            offset: uint64(uint256(log.topics[3])),
            gasUsed: gasUsed,
            relayer: relayer,
            success: success,
            error: errorBytes
        });
    }

    /// _dev Assert that the logs are XReceipt events with the correct fields.
    function assertReceipts(Vm.Log[] memory logs, XTypes.Msg[] memory xmsgs, uint64 sourceChainId) internal {
        assertEq(logs.length, xmsgs.length);
        for (uint256 i = 0; i < logs.length; i++) {
            assertReceipt(logs[i], xmsgs[i], sourceChainId);
        }
    }

    /// @dev Assert that the log is an XReceipt event with the correct fields.
    ///      We use this helper rather than vm.expectEmit(), because gasUsed is difficult to predict.
    function assertReceipt(Vm.Log memory log, XTypes.Msg memory xmsg, uint64 sourceChainId) internal {
        TestXTypes.Receipt memory receipt = parseReceipt(log);

        assertEq(receipt.sourceChainId, sourceChainId);
        assertEq(receipt.offset, xmsg.offset);
        assertEq(receipt.relayer, relayer);
        assertEq(
            receipt.success,
            // little hacky, but deriving receipts from messages helps
            // readability and this let's us do that
            xmsg.to == _reverters[xmsg.destChainId] ? false : true
        );

        // error should be empty if success is true
        if (receipt.success) assertEq(receipt.error, "");
    }

    /// @dev vm.expectCall() for multiple XMsgs
    function expectCalls(XTypes.Msg[] memory xmsgs) internal {
        for (uint256 i = 0; i < xmsgs.length; i++) {
            vm.expectCall(xmsgs[i].to, xmsgs[i].data);
        }
    }

    /// @dev The number of Counter.increment() calls in a list of xmsgs
    function numIncrements(XTypes.Msg[] memory xmsgs) internal view returns (uint256) {
        bytes32 incrHash = keccak256(abi.encodeWithSignature("increment()"));
        uint256 count = 0;

        for (uint256 i = 0; i < xmsgs.length; i++) {
            if (xmsgs[i].to == _counters[xmsgs[i].destChainId] && keccak256(xmsgs[i].data) == incrHash) {
                count++;
            }
        }

        return count;
    }
}
