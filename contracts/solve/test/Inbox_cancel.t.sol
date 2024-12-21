// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Ownable } from "solady/src/auth/Ownable.sol";
import { SolveInbox } from "src/SolveInbox.sol";
import { Solve } from "src/Solve.sol";
import { InboxBase } from "./InboxBase.sol";

/**
 * @title SolveInbox_cancel_Test
 * @notice Test suite for SolveInbox.cancel(...)
 */
contract SolveInbox_cancel_Test is InboxBase {
    function test_cancel_reverts() public {
        // cannot cancel non-existent request
        vm.expectRevert(SolveInbox.NotPendingOrRejected.selector);
        inbox.cancel(bytes32(0));

        // create request to be cancelled
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // cannot cancel cancelled request
        vm.startPrank(user);
        inbox.cancel(id);
        vm.expectRevert(SolveInbox.NotPendingOrRejected.selector);
        inbox.cancel(id);
        vm.stopPrank();

        // create request to be rejected
        vm.prank(user);
        id = inbox.request{ value: 1 ether }(call, deposits);

        // cannot double cancel rejected request
        vm.prank(solver);
        inbox.reject({ id: id, reason: 0 });
        vm.startPrank(user);
        inbox.cancel(id);
        vm.expectRevert(SolveInbox.NotPendingOrRejected.selector);
        inbox.cancel(id);
        vm.stopPrank();

        // create valid request to advance through later states
        vm.prank(user);
        id = inbox.request{ value: 1 ether }(call, deposits);

        // cannot cancel request not initiated by sender
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.cancel(id);

        // cannot cancel accepted request
        vm.prank(solver);
        inbox.accept(id);
        vm.expectRevert(SolveInbox.NotPendingOrRejected.selector);
        inbox.cancel(id);

        // TODO: complete logic to advance through additional states and then test those
    }

    function test_cancel_one_request() public {
        // create valid request
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // cancel request
        vm.prank(user);
        inbox.cancel(id);

        Solve.StatusUpdate[] memory history = inbox.getUpdateHistory(id);

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Reverted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Reverted)"
        );
        assertEq(history.length, 2, "history.length");
        assertEq(uint8(history[1].status), uint8(Solve.Status.Reverted), "history[1].status");
        assertEq(history[1].timestamp, block.timestamp, "history[1].timestamp");
    }

    function test_cancel_two_requests() public {
        // create valid requests
        vm.deal(user, 2 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.startPrank(user);
        bytes32 id1 = inbox.request{ value: 1 ether }(call, deposits);
        bytes32 id2 = inbox.request{ value: 1 ether }(call, deposits);
        vm.stopPrank();

        // cancel both requests
        vm.startPrank(user);
        inbox.cancel(id1);
        inbox.cancel(id2);
        vm.stopPrank();

        Solve.StatusUpdate[] memory history1 = inbox.getUpdateHistory(id1);
        Solve.StatusUpdate[] memory history2 = inbox.getUpdateHistory(id2);

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 2 ether, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id2).status");
        assertEq(
            id2,
            inbox.getLatestRequestByStatus(Solve.Status.Reverted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Reverted)"
        );
        assertEq(history1.length, 2, "history1.length");
        assertEq(history2.length, 2, "history2.length");
        assertEq(uint8(history1[1].status), uint8(Solve.Status.Reverted), "history1[1].status");
        assertEq(uint8(history2[1].status), uint8(Solve.Status.Reverted), "history2[1].status");
        assertEq(history1[1].timestamp, block.timestamp, "history1[1].timestamp");
        assertEq(history2[1].timestamp, block.timestamp, "history2[1].timestamp");
    }

    function test_cancel_oldest_request() public {
        // create valid requests
        vm.deal(user, 2 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.startPrank(user);
        bytes32 id1 = inbox.request{ value: 1 ether }(call, deposits);
        bytes32 id2 = inbox.request{ value: 1 ether }(call, deposits);
        vm.stopPrank();

        // cancel oldest request
        vm.prank(user);
        inbox.cancel(id1);

        Solve.StatusUpdate[] memory history1 = inbox.getUpdateHistory(id1);
        Solve.StatusUpdate[] memory history2 = inbox.getUpdateHistory(id2);

        assertEq(address(inbox).balance, 1 ether, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Pending), "inbox.getRequest(id2).status");
        assertEq(
            id1,
            inbox.getLatestRequestByStatus(Solve.Status.Reverted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Reverted)"
        );
        assertEq(
            id2,
            inbox.getLatestRequestByStatus(Solve.Status.Pending).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Pending)"
        );
        assertEq(history1.length, 2, "history1.length");
        assertEq(history2.length, 1, "history2.length");
        assertEq(uint8(history1[1].status), uint8(Solve.Status.Reverted), "history1[1].status");
        assertEq(uint8(history2[0].status), uint8(Solve.Status.Pending), "history2[0].status");
        assertEq(history1[1].timestamp, block.timestamp, "history1[1].timestamp");
        assertEq(history2[0].timestamp, block.timestamp, "history2[0].timestamp");
    }

    function test_cancel_singleToken() public {
        // create valid request
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](1);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });

        vm.startPrank(user);
        mintAndApprove(deposits);
        bytes32 id = inbox.request(call, deposits);

        // cancel request
        inbox.cancel(id);
        vm.stopPrank();

        Solve.StatusUpdate[] memory history = inbox.getUpdateHistory(id);

        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 1 ether, "token1.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Reverted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Reverted)"
        );
        assertEq(history.length, 2, "history.length");
        assertEq(uint8(history[1].status), uint8(Solve.Status.Reverted), "history[1].status");
        assertEq(history[1].timestamp, block.timestamp, "history[1].timestamp");
    }

    function test_cancel_multiToken() public {
        // create valid request
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](2);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });
        deposits[1] = Solve.TokenDeposit({ token: address(token2), amount: 1 ether });

        vm.startPrank(user);
        mintAndApprove(deposits);
        bytes32 id = inbox.request(call, deposits);

        // cancel request
        inbox.cancel(id);
        vm.stopPrank();

        Solve.StatusUpdate[] memory history = inbox.getUpdateHistory(id);

        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), 0, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 1 ether, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 1 ether, "token2.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Reverted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Reverted)"
        );
        assertEq(history.length, 2, "history.length");
        assertEq(uint8(history[1].status), uint8(Solve.Status.Reverted), "history[1].status");
        assertEq(history[1].timestamp, block.timestamp, "history[1].timestamp");
    }

    function test_cancel_nativeMultiToken() public {
        // create valid request
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](2);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });
        deposits[1] = Solve.TokenDeposit({ token: address(token2), amount: 1 ether });

        vm.startPrank(user);
        mintAndApprove(deposits);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // cancel request
        inbox.cancel(id);
        vm.stopPrank();

        Solve.StatusUpdate[] memory history = inbox.getUpdateHistory(id);

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), 0, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 1 ether, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 1 ether, "token2.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Reverted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Reverted)"
        );
        assertEq(history.length, 2, "history.length");
        assertEq(uint8(history[1].status), uint8(Solve.Status.Reverted), "history[1].status");
        assertEq(history[1].timestamp, block.timestamp, "history[1].timestamp");
    }

    function test_cancel_rejected_nativeToken_request() public {
        // create valid request
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // reject request
        vm.prank(solver);
        inbox.reject({ id: id, reason: 0 });

        // cancel rejected request
        vm.prank(user);
        inbox.cancel(id);

        Solve.StatusUpdate[] memory history = inbox.getUpdateHistory(id);

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Rejected).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Rejected)"
        );
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Reverted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Reverted)"
        );
        assertEq(history.length, 3, "history.length");
        assertEq(uint8(history[2].status), uint8(Solve.Status.Reverted), "history[2].status");
        assertEq(history[2].timestamp, block.timestamp, "history[2].timestamp");
    }

    function test_cancel_rejected_nativeMultiToken_request() public {
        // create valid request
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](2);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });
        deposits[1] = Solve.TokenDeposit({ token: address(token2), amount: 1 ether });
        vm.startPrank(user);
        mintAndApprove(deposits);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);
        vm.stopPrank();

        // reject request
        vm.prank(solver);
        inbox.reject({ id: id, reason: 0 });

        // cancel rejected request
        vm.prank(user);
        inbox.cancel(id);

        Solve.StatusUpdate[] memory history = inbox.getUpdateHistory(id);

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), 0, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 1 ether, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 1 ether, "token2.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Reverted), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Rejected).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Rejected)"
        );
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Reverted).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Reverted)"
        );
        assertEq(history.length, 3, "history.length");
        assertEq(uint8(history[2].status), uint8(Solve.Status.Reverted), "history[2].status");
        assertEq(history[2].timestamp, block.timestamp, "history[2].timestamp");
    }
}
