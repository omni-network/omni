// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Test } from "forge-std/Test.sol";
import { Vm } from "forge-std/Vm.sol";
import { Events } from "./Events.sol";
import { TestXTypes } from "./TestXTypes.sol";

/**
 * @title Utils
 * @dev Defines test utilities.
 */
contract Utils is Test, Events {
    /// @dev Parse an XReceipt log
    function _parseReceipt(Vm.Log memory log) internal returns (TestXTypes.Receipt memory) {
        assertEq(log.topics.length, 3);
        assertEq(log.topics[0], XReceipt.selector);

        (uint256 gasUsed, address relayer, bool success) = abi.decode(log.data, (uint256, address, bool));

        return TestXTypes.Receipt({
            sourceChainId: uint64(uint256(log.topics[1])),
            streamOffset: uint64(uint256(log.topics[2])),
            gasUsed: gasUsed,
            relayer: relayer,
            success: success
        });
    }

    /// @dev Assert that the log is an XReceipt event with the correct fields.
    ///      We use this helper rather than vm.expectEmit(), because gasUsed is difficult to predict.
    function _assertReceiptEmitted(
        Vm.Log memory log,
        uint64 sourceChainId,
        uint64 streamOffset,
        address relayer,
        bool success
    ) internal {
        TestXTypes.Receipt memory receipt = _parseReceipt(log);

        assertEq(receipt.sourceChainId, sourceChainId);
        assertEq(receipt.streamOffset, streamOffset);
        assertEq(receipt.relayer, relayer);
        assertEq(receipt.success, success);
    }
}
