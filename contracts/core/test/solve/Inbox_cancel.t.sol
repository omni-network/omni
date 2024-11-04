// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { MockToken } from "test/utils/MockToken.sol";
import { Inbox } from "src/solve/Inbox.sol";
import { Solve } from "src/solve/Solve.sol";
import { Test } from "forge-std/Test.sol";

import { Ownable } from "solady/src/auth/Ownable.sol";

/**
 * @title Inbox_request_Test
 * @notice Test suite for solver Inbox.request(...)
 * @dev TODO: add fuzz / invariant tests
 */
contract Inbox_request_Test is Test {
    Inbox inbox;

    MockToken token1;
    MockToken token2;

    address user = makeAddr("user");
    address solver = makeAddr("solver");

    function setUp() public {
        inbox = new Inbox();
        inbox.initialize(address(this), solver);
        token1 = new MockToken();
        token2 = new MockToken();
    }

    function test_cancel_reverts() public {
        // cannot cancel non-existent request
        vm.expectRevert(Inbox.RequestNotCancelable.selector);
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
        vm.expectRevert(Inbox.RequestNotCancelable.selector);
        inbox.cancel(id);
        vm.stopPrank();

        // create request to be rejected
        vm.prank(user);
        id = inbox.request{ value: 1 ether }(call, deposits);

        // cannot double cancel rejected request
        vm.prank(solver);
        inbox.reject(id);
        vm.startPrank(user);
        inbox.cancel(id);
        vm.expectRevert(Inbox.RequestNotCancelable.selector);
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
        vm.expectRevert(Inbox.RequestNotCancelable.selector);
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

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id).status");
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

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 2 ether, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id2).status");
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

        assertEq(address(inbox).balance, 1 ether, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Open), "inbox.getRequest(id2).status");
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

        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 1 ether, "token1.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id).status");
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

        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), 0, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 1 ether, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 1 ether, "token2.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id).status");
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

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), 0, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 1 ether, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 1 ether, "token2.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id).status");
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
        inbox.reject(id);

        // cancel rejected request
        vm.prank(user);
        inbox.cancel(id);

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id).status");
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
        inbox.reject(id);

        // cancel rejected request
        vm.prank(user);
        inbox.cancel(id);

        assertEq(address(inbox).balance, 0, "address(inbox).balance");
        assertEq(address(user).balance, 1 ether, "address(user).balance");
        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), 0, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 1 ether, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 1 ether, "token2.balanceOf(user)");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Cancelled), "inbox.getRequest(id).status");
    }

    function randCall() internal returns (Solve.Call memory) {
        uint256 rand = vm.randomUint(1, 1000);
        return Solve.Call({
            destChainId: uint64(rand),
            value: rand * 1 ether,
            target: address(uint160(rand)),
            data: abi.encode("data", rand)
        });
    }

    function mintAndApprove(Solve.TokenDeposit[] memory deposits) internal {
        for (uint256 i = 0; i < deposits.length; i++) {
            MockToken(deposits[i].token).approve(address(inbox), deposits[i].amount);
            MockToken(deposits[i].token).mint(user, deposits[i].amount);
        }
    }
}
