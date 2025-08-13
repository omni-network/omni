// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

import { TestXTypes } from "test/xchain/common/TestXTypes.sol";
import { Base } from "test/xchain/nomina/common/Base.sol";
import { Counter } from "test/xchain/nomina/common/Counter.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title NominaPortal_adversarial
 * @dev Test cases for adversarial scenarios.
 */
contract NominaPortal_adversarial is Base {
    /// @dev Test than an xcall to the portal address fails
    function test_xcallToPortal__fails() public {
        XTypes.Msg memory xmsg = XTypes.Msg({
            destChainId: thisChainId,
            shardId: uint64(ConfLevel.Finalized),
            offset: 1,
            sender: address(1234),
            to: address(portal),
            data: "", // doesn't matter, should fail before execution
            gasLimit: 100_000
        });

        TestXTypes.Receipt memory receipt = _exec(xmsg);

        assertFalse(receipt.success);
        assertEq(receipt.gasUsed, 0); // not executed
        assertEq(receipt.error, abi.encodeWithSignature("Error(string)", "NominaPortal: no xcall to portal"));
    }

    /// @dev Helpler to call `portal.exec`, setting the chainId and returning the receipt.
    function _exec(XTypes.Msg memory xmsg) internal returns (TestXTypes.Receipt memory) {
        vm.recordLogs();
        vm.chainId(thisChainId);
        portal.exec(_xheader(xmsg), xmsg);

        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 1);

        return parseReceipt(logs[0]);
    }

    // @dev Helper to create a XBlock header for an xmsg
    function _xheader(XTypes.Msg memory xmsg) internal pure returns (XTypes.BlockHeader memory) {
        return XTypes.BlockHeader({
            sourceChainId: chainAId,
            consensusChainId: nominaCChainID,
            confLevel: uint8(xmsg.shardId),
            offset: 1,
            sourceBlockHeight: 100,
            sourceBlockHash: bytes32(0)
        });
    }
}
