// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { MockToken } from "test/utils/MockToken.sol";
import { Inbox } from "src/solve/Inbox.sol";
import { Solve } from "src/solve/Solve.sol";
import { Test } from "forge-std/Test.sol";

import { Ownable } from "solady/src/auth/Ownable.sol";

/**
 * @title Inbox_accept_Test
 * @notice Test suite for solver Inbox.accept(...)
 * @dev TODO: add fuzz / invariant tests
 */
contract Inbox_accept_Test is Test {
    Inbox inbox;

    MockToken token1;
    MockToken token2;

    address user = makeAddr("user");
    address solver = makeAddr("solver");

    function setUp() public {
        inbox = new Inbox();
        // Omni and outbox addresses not needed for these tests
        inbox.initialize(address(this), solver, address(0x1234), address(0x5678));
        token1 = new MockToken();
        token2 = new MockToken();
    }

    /// @dev Test all revert conditions for Inbox.accept(...)
    function test_accept_reverts() public {
        // needs to have solver role
        vm.expectRevert(Ownable.Unauthorized.selector);
        inbox.accept(bytes32(0));

        // needs open request
        vm.prank(solver);
        vm.expectRevert(Inbox.RequestStateInvalid.selector);
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
        vm.expectRevert(Inbox.RequestStateInvalid.selector);
        inbox.accept(id);

        // create request to be rejected
        vm.prank(user);
        id = inbox.request{ value: 1 ether }(call, deposits);

        // cannot accept rejected request
        vm.startPrank(solver);
        inbox.reject(id, Solve.RejectReason.None);
        vm.expectRevert(Inbox.RequestStateInvalid.selector);
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
        vm.expectRevert(Inbox.RequestStateInvalid.selector);
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

        assertEq(inbox.getRequest(id).acceptedBy, solver, "inbox.getRequest(id).acceptedBy");
        assertEq(uint8(inbox.getRequest(id).status), uint8(Solve.Status.Accepted), "inbox.getRequest(id).status");
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

        assertEq(inbox.getRequest(id1).acceptedBy, solver, "inbox.getRequest(id1).acceptedBy");
        assertEq(inbox.getRequest(id2).acceptedBy, solver, "inbox.getRequest(id2).acceptedBy");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Accepted), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Accepted), "inbox.getRequest(id2).status");
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

        assertEq(inbox.getRequest(id1).acceptedBy, address(0), "inbox.getRequest(id1).acceptedBy");
        assertEq(inbox.getRequest(id2).acceptedBy, solver, "inbox.getRequest(id2).acceptedBy");
        assertEq(uint8(inbox.getRequest(id1).status), uint8(Solve.Status.Pending), "inbox.getRequest(id1).status");
        assertEq(uint8(inbox.getRequest(id2).status), uint8(Solve.Status.Accepted), "inbox.getRequest(id2).status");
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
