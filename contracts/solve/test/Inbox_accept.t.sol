// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Ownable } from "solady/src/auth/Ownable.sol";
import { SolveInbox } from "src/SolveInbox.sol";
import { Solve } from "src/Solve.sol";
import { InboxBase } from "./InboxBase.sol";

/**
 * @title SolveInbox_accept_Test
 * @notice Test suite for SolveInbox.accept(...)
 */
contract SolveInbox_accept_Test is InboxBase {
    /// @dev Test all revert conditions for SolveInbox.accept(...)
    function test_accept_reverts() public {
        // needs to have solver role
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.accept(bytes32(0));

        // needs open request
        vm.prank(solver);
        vm.expectRevert(SolveInbox.NotPending.selector);
        inbox.accept(bytes32(0));

        // create request to be cancelled
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // cannot accept cancelled request
        vm.prank(user);
        inbox.cancel(id);
        vm.prank(solver);
        vm.expectRevert(SolveInbox.NotPending.selector);
        inbox.accept(id);

        // create request to be rejected
        vm.prank(user);
        id = inbox.request{ value: 1 ether }(call, deposits);

        // cannot accept rejected request
        vm.startPrank(solver);
        inbox.reject({ id: id, reason: 0 });
        vm.expectRevert(SolveInbox.NotPending.selector);
        inbox.accept(id);
        vm.stopPrank();

        // create valid request to advance through later states
        vm.deal(user, 1 ether);
        vm.prank(user);
        id = inbox.request{ value: 1 ether }(call, deposits);

        // once accepted, non-solvers still cannot call accept
        vm.prank(solver);
        inbox.accept(id);
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.accept(id);

        // once accepted, solvers cannot accept again
        vm.prank(solver);
        vm.expectRevert(SolveInbox.NotPending.selector);
        inbox.accept(id);

        // TODO: complete logic to advance through additional states and then test those
    }

    /// @dev Test accepting the first request
    function test_accept_one_request() public {
        // create valid request
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // accept first request
        vm.prank(solver);
        inbox.accept(id);

        Solve.StatusUpdate[] memory history = inbox.getUpdateHistory(id);

        assertEq(inbox.getRequest(id).acceptedBy, solver, "inbox.getRequest(id).acceptedBy");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Accepted), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Accepted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Accepted)"
        );
        assertEq(history.length, 2, "history.length");
        assertEq(uint8(history[1].status), uint8(Solve.Status.Accepted), "history[1].status");
        assertEq(history[1].timestamp, block.timestamp, "history[1].timestamp");
    }

    /// @dev Test accepting two requests
    function test_accept_two_requests() public {
        // create valid requests
        vm.deal(user, 2 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.startPrank(user);
        bytes32 id1 = inbox.request{ value: 1 ether }(call, deposits);
        bytes32 id2 = inbox.request{ value: 1 ether }(call, deposits);
        vm.stopPrank();

        // accept both requests
        vm.startPrank(solver);
        inbox.accept(id1);
        inbox.accept(id2);
        vm.stopPrank();

        Solve.StatusUpdate[] memory history1 = inbox.getUpdateHistory(id1);
        Solve.StatusUpdate[] memory history2 = inbox.getUpdateHistory(id2);

        assertEq(inbox.getRequest(id1).acceptedBy, solver, "inbox.getRequest(id1).acceptedBy");
        assertEq(inbox.getRequest(id2).acceptedBy, solver, "inbox.getRequest(id2).acceptedBy");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Accepted), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Accepted), "inbox.getRequest(id2).status");
        assertEq(
            id2,
            inbox.getLatestRequestByStatus(Solve.Status.Accepted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Accepted)"
        );
        assertEq(history1.length, 2, "history1.length");
        assertEq(history2.length, 2, "history2.length");
        assertEq(uint8(history1[1].status), uint8(Solve.Status.Accepted), "history1[1].status");
        assertEq(uint8(history2[1].status), uint8(Solve.Status.Accepted), "history2[1].status");
        assertEq(history1[1].timestamp, block.timestamp, "history1[1].timestamp");
        assertEq(history2[1].timestamp, block.timestamp, "history2[1].timestamp");
    }

    /// @dev Test accepting requests out of order
    function test_accept_skip_first() public {
        // create valid requests
        vm.deal(user, 2 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.startPrank(user);
        bytes32 id1 = inbox.request{ value: 1 ether }(call, deposits);
        bytes32 id2 = inbox.request{ value: 1 ether }(call, deposits);
        vm.stopPrank();

        // accept second request
        vm.startPrank(solver);
        inbox.accept(id2);
        vm.stopPrank();

        Solve.StatusUpdate[] memory history1 = inbox.getUpdateHistory(id1);
        Solve.StatusUpdate[] memory history2 = inbox.getUpdateHistory(id2);

        assertEq(inbox.getRequest(id1).acceptedBy, address(0), "inbox.getRequest(id1).acceptedBy");
        assertEq(inbox.getRequest(id2).acceptedBy, solver, "inbox.getRequest(id2).acceptedBy");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Pending), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Accepted), "inbox.getRequest(id2).status");
        assertEq(
            id2,
            inbox.getLatestRequestByStatus(Solve.Status.Accepted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Accepted)"
        );
        assertEq(history1.length, 1, "history1.length");
        assertEq(history2.length, 2, "history2.length");
        assertEq(uint8(history1[0].status), uint8(Solve.Status.Pending), "history1[0].status");
        assertEq(uint8(history2[1].status), uint8(Solve.Status.Accepted), "history2[1].status");
        assertEq(history1[0].timestamp, block.timestamp, "history1[0].timestamp");
        assertEq(history2[1].timestamp, block.timestamp, "history2[1].timestamp");
    }
}
