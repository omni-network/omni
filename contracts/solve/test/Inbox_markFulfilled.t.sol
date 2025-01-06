// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { SolveInbox, ISolveInbox } from "src/SolveInbox.sol";
import { Solve } from "src/Solve.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { InboxBase } from "./InboxBase.sol";

/**
 * @title SolveInbox_markFulfilled_Test
 * @notice Test suite for SolveInbox.markFulfilled(...)
 */
contract SolveInbox_markFulfilled_Test is InboxBase {
    function test_markFulfilled_reverts() public {
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);

        // make request
        vm.deal(user, 1 ether);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // must be accepted
        vm.expectRevert(SolveInbox.NotAccepted.selector);
        inbox.markFulfilled(id, callHash(id, call));
        vm.prank(solver);
        inbox.accept(id);

        // must be xcall from outbox
        vm.expectRevert(SolveInbox.NotOutbox.selector);
        portal.mockXCall({
            sourceChainId: call.chainId,
            sender: address(1234), // not outbox
            data: abi.encodeCall(inbox.markFulfilled, (id, callHash(id, call))),
            to: address(inbox)
        });

        // must be xcall from call.chainId
        vm.expectRevert(SolveInbox.WrongSourceChain.selector);
        portal.mockXCall({
            sourceChainId: 1234, // not call.chainId
            sender: address(outbox),
            data: abi.encodeCall(inbox.markFulfilled, (id, callHash(id, call))),
            to: address(inbox)
        });

        // must have correct call hash
        vm.expectRevert(SolveInbox.WrongCallHash.selector);
        portal.mockXCall({
            sourceChainId: call.chainId,
            sender: address(outbox),
            data: abi.encodeCall(inbox.markFulfilled, (id, bytes32(uint256(1234)))), // not correct call hash
            to: address(inbox)
        });
    }

    function test_markFulfilled_success() public {
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);

        // make request
        vm.deal(user, 1 ether);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // accept
        vm.prank(solver);
        inbox.accept(id);

        // mark fulfilled
        vm.warp(block.timestamp + 1 hours);
        vm.expectEmit(address(inbox));
        emit ISolveInbox.Fulfilled(id, callHash(id, call), solver);
        portal.mockXCall({
            sourceChainId: call.chainId,
            sender: address(outbox),
            data: abi.encodeCall(inbox.markFulfilled, (id, callHash(id, call))),
            to: address(inbox)
        });

        Solve.StatusUpdate[] memory history = inbox.getUpdateHistory(id);

        Solve.Request memory req = inbox.getRequest(id);
        assertEq(uint8(req.status), uint8(Solve.Status.Fulfilled), "req.status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Fulfilled).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Fulfilled)"
        );
        assertEq(history.length, 3, "history.length");
        assertEq(uint8(history[2].status), uint8(Solve.Status.Fulfilled), "history[2].status");
        assertEq(history[2].timestamp, block.timestamp, "history[2].timestamp");
    }
}
