// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Ownable } from "solady/src/auth/Ownable.sol";
import { SolveInbox, ISolveInbox } from "src/SolveInbox.sol";
import { Solve } from "src/Solve.sol";
import { InboxBase } from "./InboxBase.sol";

/**
 * @title SolveInbox_claim_Test
 * @notice Test suite for SolveInbox.claim(...)
 */
contract SolveInbox_claim_Test is InboxBase {
    address claimTo = makeAddr("claim-to");

    function test_claim_reverts() public {
        // no request
        vm.expectRevert(SolveInbox.NotFulfilled.selector);
        inbox.claim(bytes32(0), claimTo);

        // open request
        vm.deal(user, 1 ether);
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        vm.prank(user);
        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);

        // pending (not fulfilled)
        vm.expectRevert(SolveInbox.NotFulfilled.selector);
        inbox.claim(id, claimTo);

        // accept
        vm.prank(solver);
        inbox.accept(id);

        // accepted (not fulfilled)
        vm.expectRevert(SolveInbox.NotFulfilled.selector);
        inbox.claim(id, claimTo);

        // mark fulfilled
        portal.mockXCall({
            sourceChainId: call.destChainId,
            sender: address(outbox),
            data: abi.encodeCall(inbox.markFulfilled, (id, callHash(id, call))),
            to: address(inbox)
        });

        // not acceptedBy
        vm.expectRevert(Ownable.Unauthorized.selector);
        vm.prank(makeAddr("not-solver"));
        inbox.claim(id, claimTo);

        // no claimTo zero
        vm.expectRevert(SolveInbox.InvalidRecipient.selector);
        vm.prank(solver);
        inbox.claim(id, address(0));
    }

    function test_claim_singleNative() public {
        // open, accept, fulfill
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);
        bytes32 id = openAcceptFulfill(call, deposits, 1 ether);

        // claim
        vm.expectEmit(address(inbox));
        emit ISolveInbox.Claimed(id, solver, claimTo, inbox.getRequest(id).deposits);
        vm.prank(solver);
        inbox.claim(id, claimTo);

        // assert claimed
        Solve.Request memory req = inbox.getRequest(id);
        assertEq(uint8(req.status), uint8(Solve.Status.Claimed), "req.status");
        assertEq(claimTo.balance, 1 ether, "claimTo.balance");
    }

    function test_claim_singleToken() public {
        // open, accept, fulfill
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](1);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });
        bytes32 id = openAcceptFulfill(call, deposits, 0);

        // claim
        vm.expectEmit(address(inbox));
        emit ISolveInbox.Claimed(id, solver, claimTo, inbox.getRequest(id).deposits);
        vm.prank(solver);
        inbox.claim(id, claimTo);

        // assert claimed
        Solve.Request memory req = inbox.getRequest(id);
        assertEq(uint8(req.status), uint8(Solve.Status.Claimed), "req.status");
        assertEq(token1.balanceOf(claimTo), 1 ether, "token.balanceOf(claimTo)");
    }

    function test_claim_multiDeposit() public {
        // open, accept, fulfill
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](2);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });
        deposits[1] = Solve.TokenDeposit({ token: address(token2), amount: 2 ether });
        bytes32 id = openAcceptFulfill(call, deposits, 3 ether);

        // claim
        vm.expectEmit(address(inbox));
        emit ISolveInbox.Claimed(id, solver, claimTo, inbox.getRequest(id).deposits);
        vm.prank(solver);
        inbox.claim(id, claimTo);

        // assert claimed
        Solve.Request memory req = inbox.getRequest(id);
        assertEq(uint8(req.status), uint8(Solve.Status.Claimed), "req.status");
        assertEq(claimTo.balance, 3 ether, "claimTo.balance");
        assertEq(token1.balanceOf(claimTo), 1 ether, "token1.balanceOf(claimTo)");
        assertEq(token2.balanceOf(claimTo), 2 ether, "token2.balanceOf(claimTo)");
    }

    /// @dev Open a request, accept it, mark it as fulfilled, and return the request ID.
    function openAcceptFulfill(Solve.Call memory call, Solve.TokenDeposit[] memory tokenDeposits, uint256 nativeDeposit)
        internal
        returns (bytes32 id)
    {
        // open request
        vm.deal(user, nativeDeposit);
        vm.startPrank(user);
        mintAndApprove(tokenDeposits);
        id = inbox.request{ value: nativeDeposit }(call, tokenDeposits);
        vm.stopPrank();

        // accept
        vm.prank(solver);
        inbox.accept(id);

        // mark fulfilled
        portal.mockXCall({
            sourceChainId: call.destChainId,
            sender: address(outbox),
            data: abi.encodeCall(inbox.markFulfilled, (id, callHash(id, call))),
            to: address(inbox)
        });
    }
}
