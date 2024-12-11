// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Ownable } from "solady/src/auth/Ownable.sol";
import { SolveInbox } from "src/SolveInbox.sol";
import { Solve } from "src/Solve.sol";
import { InboxBase } from "./InboxBase.sol";

/**
 * @title SolveInbox_reject_Test
 * @notice Test suite for SolveInbox.reject(...)
 */
contract SolveInbox_reject_Test is InboxBase {
    function test_reject_reverts() public {
        // cannot reject non-existent request
        vm.prank(solver);
        vm.expectRevert(SolveInbox.NotPending.selector);
        inbox.reject(bytes32(0), Solve.RejectReason.None);

        // needs to have solver role
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.reject(bytes32(0), Solve.RejectReason.None);

        // create request to cancel before rejecting
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.startPrank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);
        inbox.cancel(id);
        vm.stopPrank();

        // cannot reject cancelled request
        vm.prank(solver);
        vm.expectRevert(SolveInbox.NotPending.selector);
        inbox.reject(id, Solve.RejectReason.None);

        // create request to accept before rejecting
        vm.deal(user, 1 ether);
        vm.prank(user);
        id = inbox.request{ value: 1 ether }(call, deposits);

        // cannot reject accepted request
        vm.startPrank(solver);
        inbox.accept(id);
        vm.expectRevert(SolveInbox.NotPending.selector);
        inbox.reject(id, Solve.RejectReason.None);
        vm.stopPrank();
    }

    function test_reject_one_request() public {
        // create valid request
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // reject request
        vm.prank(solver);
        inbox.reject(id, Solve.RejectReason.None);

        assertEq(address(inbox).balance, 1 ether, "address(inbox).balance");
        assertEq(address(user).balance, 0, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Rejected), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Rejected).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Rejected)"
        );
        assertEq(
            uint8(inbox.getRequestUpdateHistory(id)[1].status),
            uint8(Solve.Status.Rejected),
            "inbox.getRequestUpdateHistory(id)[1].status"
        );
        assertEq(inbox.getRequestUpdateHistory(id).length, 2, "inbox.getRequestUpdateHistory(id).length");
    }

    function test_reject_two_requests() public {
        // create valid requests
        vm.deal(user, 2 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.startPrank(user);
        bytes32 id1 = inbox.request{ value: 1 ether }(call, deposits);
        bytes32 id2 = inbox.request{ value: 1 ether }(call, deposits);
        vm.stopPrank();

        // reject both requests
        vm.startPrank(solver);
        inbox.reject(id1, Solve.RejectReason.None);
        inbox.reject(id2, Solve.RejectReason.None);
        vm.stopPrank();

        assertEq(address(inbox).balance, 2 ether, "address(inbox).balance");
        assertEq(address(user).balance, 0, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Rejected), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Rejected), "inbox.getRequest(id2).status");
        assertEq(
            id2,
            inbox.getLatestRequestByStatus(Solve.Status.Rejected).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Rejected)"
        );
        assertEq(
            uint8(inbox.getRequestUpdateHistory(id1)[1].status),
            uint8(Solve.Status.Rejected),
            "inbox.getRequestUpdateHistory(id1)[1].status"
        );
        assertEq(
            uint8(inbox.getRequestUpdateHistory(id2)[1].status),
            uint8(Solve.Status.Rejected),
            "inbox.getRequestUpdateHistory(id2)[1].status"
        );
        assertEq(inbox.getRequestUpdateHistory(id1).length, 2, "inbox.getRequestUpdateHistory(id1).length");
        assertEq(inbox.getRequestUpdateHistory(id2).length, 2, "inbox.getRequestUpdateHistory(id2).length");
    }

    function test_reject_oldest_request() public {
        // create valid requests
        vm.deal(user, 2 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.startPrank(user);
        bytes32 id1 = inbox.request{ value: 1 ether }(call, deposits);
        bytes32 id2 = inbox.request{ value: 1 ether }(call, deposits);
        vm.stopPrank();

        // reject oldest request
        vm.startPrank(solver);
        inbox.reject(id1, Solve.RejectReason.None);
        vm.stopPrank();

        assertEq(address(inbox).balance, 2 ether, "address(inbox).balance");
        assertEq(address(user).balance, 0, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Rejected), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Pending), "inbox.getRequest(id2).status");
        assertEq(
            id1,
            inbox.getLatestRequestByStatus(Solve.Status.Rejected).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Rejected)"
        );
        assertEq(
            id2,
            inbox.getLatestRequestByStatus(Solve.Status.Pending).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Pending)"
        );
        assertEq(
            uint8(inbox.getRequestUpdateHistory(id1)[1].status),
            uint8(Solve.Status.Rejected),
            "inbox.getRequestUpdateHistory(id1)[1].status"
        );
        assertEq(
            uint8(inbox.getRequestUpdateHistory(id2)[0].status),
            uint8(Solve.Status.Pending),
            "inbox.getRequestUpdateHistory(id2)[0].status"
        );
        assertEq(inbox.getRequestUpdateHistory(id1).length, 2, "inbox.getRequestUpdateHistory(id1).length");
        assertEq(inbox.getRequestUpdateHistory(id2).length, 1, "inbox.getRequestUpdateHistory(id2).length");
    }

    function test_reject_nativeMultiToken() public {
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
        inbox.reject(id, Solve.RejectReason.None);

        assertEq(address(inbox).balance, 1 ether, "address(inbox).balance");
        assertEq(address(user).balance, 0, "address(user).balance");
        assertEq(token1.balanceOf(address(inbox)), 1 ether, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), 1 ether, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 0, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 0, "token2.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Rejected), "inbox.getRequest(id).status");
        assertEq(
            id,
            inbox.getLatestRequestByStatus(Solve.Status.Rejected).id,
            "inbox.getLatestRequestByStatus(Solve.Status.Rejected)"
        );
        assertEq(
            uint8(inbox.getRequestUpdateHistory(id)[1].status),
            uint8(Solve.Status.Rejected),
            "inbox.getRequestUpdateHistory(id)[1].status"
        );
        assertEq(inbox.getRequestUpdateHistory(id).length, 2, "inbox.getRequestUpdateHistory(id).length");
    }
}
