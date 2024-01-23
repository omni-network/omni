// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Test } from "forge-std/Test.sol";
import { XChain } from "src/libraries/XChain.sol";
import { TestPortal } from "./TestPortal.sol";
import { Counter } from "./Counter.sol";
import { Reverter } from "./Reverter.sol";
import { Events } from "./Events.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title CommonTest
 * @dev An extension of forge Test that includes Omni specifc setup, fixtures, utils, and events.
 */
contract CommonTest is Test, Events {
    // some other chain id to use for source or dest of xmsgs
    uint64 constant otherChainId = 2;

    address deployer;
    address xcaller;
    address xrelayer;

    TestPortal portal;

    Counter counter;
    Counter otherCounter; // on otherChainId

    Reverter reverter;
    Reverter otherReverter; // on otherChainId

    function setUp() public {
        deployer = makeAddr("deployer");
        xcaller = makeAddr("xcaller");
        xrelayer = makeAddr("relayer");

        vm.startPrank(deployer);
        portal = new TestPortal();
        counter = new Counter();
        otherCounter = new Counter();
        reverter = new Reverter();
        otherReverter = new Reverter();
        vm.stopPrank();
    }

    /**
     * Fixtures.
     */

    /// @dev Get XMsg fields for an outbound Counter.increment() xcall
    function _outbound_increment() internal view returns (XChain.Msg memory) {
        return XChain.Msg({
            sourceChainId: portal.chainId(),
            destChainId: otherChainId,
            streamOffset: portal.outXStreamOffset(otherChainId),
            sender: address(counter),
            to: address(otherCounter),
            data: abi.encodeWithSignature("increment()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Get XMsg fields for an inbound Counter.increment() xmsg
    function _inbound_increment() internal view returns (XChain.Msg memory) {
        return XChain.Msg({
            sourceChainId: otherChainId,
            destChainId: portal.chainId(),
            streamOffset: portal.inXStreamOffset(otherChainId),
            sender: address(otherCounter),
            to: address(counter),
            data: abi.encodeWithSignature("increment()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Get XMsg fields for an inbound Reverter.revert() xmsg
    function _inbound_revert() internal view returns (XChain.Msg memory) {
        return XChain.Msg({
            sourceChainId: otherChainId,
            destChainId: portal.chainId(),
            streamOffset: portal.inXStreamOffset(otherChainId),
            sender: address(otherReverter),
            to: address(reverter),
            data: abi.encodeWithSignature("forceRevert()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Create an test XSubmission
    function _xsub(XChain.Msg[] memory xmsgs) internal pure returns (XChain.Submission memory) {
        return XChain.Submission({
            attestationRoot: bytes32(0), // TODO: still unchecked
            blockHeader: XChain.BlockHeader(0, 0, 0), // TODO: still unchecked
            msgs: xmsgs,
            proof: new bytes32[](0), // TODO: still unchecked
            proofFlags: new bool[](0), // TODO: still unchecked
            signatures: new XChain.SigTuple[](0) // TODO: still unchecked
         });
    }

    /**
     * Utils.
     */

    // We define _XReceipt type here for convenience.
    // _ prefixed because event XReceipt is defined in test/common/Events.sol
    // The struct is not used in production, and is therefore not defined in XChain.sol library.
    struct _XReceipt {
        uint64 sourceChainId;
        uint64 streamOffset;
        uint256 gasUsed;
        address relayer;
        bool success;
    }

    /// @dev Parse an XReceipt log
    function _parseReceipt(Vm.Log memory log) internal returns (_XReceipt memory) {
        assertEq(log.emitter, address(portal));
        assertEq(log.topics.length, 3);
        assertEq(log.topics[0], XReceipt.selector);

        (uint256 gasUsed, address relayer, bool success) = abi.decode(log.data, (uint256, address, bool));

        return _XReceipt({
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
        _XReceipt memory receipt = _parseReceipt(log);

        assertEq(receipt.sourceChainId, sourceChainId);
        assertEq(receipt.streamOffset, streamOffset);
        assertEq(receipt.relayer, relayer);
        assertEq(receipt.success, success);
    }
}
