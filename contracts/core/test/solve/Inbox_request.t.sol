// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { MockToken } from "test/utils/MockToken.sol";
import { Inbox } from "src/solve/Inbox.sol";
import { Solve } from "src/solve/Solve.sol";
import { Test } from "forge-std/Test.sol";

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

    modifier prankUser() {
        vm.startPrank(user);
        _;
        vm.stopPrank();
    }

    function setUp() public {
        inbox = new Inbox();
        token1 = new MockToken();
        token2 = new MockToken();
    }

    /// @dev Test all revert conditions for Inbox.request(...)
    function test_request_reverts() public prankUser {
        Solve.Call memory call = Solve.Call({ destChainId: 0, value: 0, target: address(0), data: bytes("") });
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);

        // needs call.target
        vm.expectRevert(Inbox.InvalidCall.selector);
        inbox.request(call, deposits);
        call.target = address(1);

        // needs destChainId
        vm.expectRevert(Inbox.InvalidCall.selector);
        inbox.request(call, deposits);
        call.destChainId = 1;

        // needs data
        vm.expectRevert(Inbox.InvalidCall.selector);
        inbox.request(call, deposits);
        call.data = bytes("data");

        // needs deposits
        vm.expectRevert(Inbox.NoDeposits.selector);
        inbox.request(call, deposits);
        deposits = new Solve.TokenDeposit[](1);

        // needs non-zero amount
        vm.expectRevert(Inbox.InvalidDeposit.selector);
        inbox.request(call, deposits);
        deposits[0].amount = 1 ether;

        // needs non-zero token
        vm.expectRevert(Inbox.InvalidDeposit.selector);
        inbox.request(call, deposits);
        deposits[0].token = address(token1);

        // needs balalnce & allowance. we do not test ERC20 errors here
        vm.expectRevert();
        inbox.request(call, deposits);
        mintAndApprove(deposits);

        // success
        inbox.request(call, deposits);
    }

    /// @dev Test a single token deposit
    function test_request_singleToken() public prankUser {
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](1);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });

        mintAndApprove(deposits);

        bytes32 id = inbox.request(call, deposits);
        assertEq(token1.balanceOf(address(inbox)), deposits[0].amount, "token1.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 0, "token1.balanceOf(user)");

        assertNewRequest({ id: id, from: user, call: call, deposits: deposits, nativeDeposit: 0 });
    }

    /// @dev Test multiple token deposits
    function test_request_multiToken() public prankUser {
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](2);

        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });
        deposits[1] = Solve.TokenDeposit({ token: address(token2), amount: 2 ether });

        mintAndApprove(deposits);

        bytes32 id = inbox.request(call, deposits);
        assertEq(token1.balanceOf(address(inbox)), deposits[0].amount, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), deposits[1].amount, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 0, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 0, "token2.balanceOf(user)");

        assertNewRequest({ id: id, from: user, call: call, deposits: deposits, nativeDeposit: 0 });
    }

    /// @dev Test a single native deposit
    function test_request_singleNative() public prankUser {
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);

        vm.deal(user, 1 ether);

        bytes32 id = inbox.request{ value: 1 ether }(call, deposits);
        assertEq(address(inbox).balance, 1 ether, "inbox.balance");

        assertNewRequest({ id: id, from: user, call: call, deposits: deposits, nativeDeposit: 1 ether });
    }

    /// @dev Test multiple native deposits
    function test_request_nativeMultiToken() public prankUser {
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](2);

        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });
        deposits[1] = Solve.TokenDeposit({ token: address(token2), amount: 2 ether });

        vm.deal(user, 3 ether);
        mintAndApprove(deposits);

        bytes32 id = inbox.request{ value: 3 ether }(call, deposits);

        assertEq(address(inbox).balance, 3 ether, "inbox.balance");
        assertEq(token1.balanceOf(address(inbox)), deposits[0].amount, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), deposits[1].amount, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(user), 0, "token1.balanceOf(user)");
        assertEq(token2.balanceOf(user), 0, "token2.balanceOf(user)");

        assertNewRequest({ id: id, from: user, call: call, deposits: deposits, nativeDeposit: 3 ether });
    }

    /// @dev Test opening two requests
    function test_request_two() public prankUser {
        Solve.Call memory call = randCall();
        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](0);

        vm.deal(user, 3 ether);

        bytes32 id1 = inbox.request{ value: 1 ether }(call, deposits);
        bytes32 id2 = inbox.request{ value: 2 ether }(call, deposits);

        assertEq(address(inbox).balance, 3 ether, "address(inbox).balance");
        assertNewRequest({ id: id1, from: user, call: call, deposits: deposits, nativeDeposit: 1 ether });
        assertNewRequest({ id: id2, from: user, call: call, deposits: deposits, nativeDeposit: 2 ether });
    }

    /// @dev Test that inbox has the correct state after a request
    function assertNewRequest(
        bytes32 id,
        address from,
        Solve.Call memory call,
        Solve.TokenDeposit[] memory deposits,
        uint256 nativeDeposit
    ) internal view {
        Solve.Request memory req = inbox.getRequest(id);

        assertTrue(req.status == Solve.Status.Open, "_assertNewRequest : req.status");

        assertEq(req.id, id, "_assertNewRequest : req.id");
        assertEq(req.from, from, "_assertNewRequest : req.from");
        assertEq(req.updatedAt, block.timestamp, "_assertNewRequest : req.updatedAt"); // assumes no vm.warp()
        assertEq(req.call.target, call.target, "_assertNewRequest : req.call.target");
        assertEq(req.call.destChainId, call.destChainId, "_assertNewRequest : req.call.destChainId");
        assertEq(req.call.value, call.value, "_assertNewRequest : req.call.value");
        assertEq(req.call.data, call.data, "_assertNewRequest : req.call.data");

        uint256 numDeposits = nativeDeposit > 0 ? deposits.length + 1 : deposits.length;
        assertEq(req.deposits.length, numDeposits, "_assertNewRequest : req.deposits.length");

        // if nativeDeposit, should be first
        if (nativeDeposit > 0) {
            assertEq(req.deposits[0].token, address(0), "_assertNewRequest : req.deposits[0].token");
            assertEq(req.deposits[0].amount, nativeDeposit, "_assertNewRequest : req.deposits[0].amount");
            assertEq(req.deposits[0].isNative, true, "_assertNewRequest : req.deposits[0].isNative");
        }

        uint256 start = nativeDeposit > 0 ? 1 : 0;
        for (uint256 i = start; i < numDeposits; i++) {
            assertEq(req.deposits[i].isNative, false, "_assertNewRequest : req.deposits[i].isNative");
            assertEq(req.deposits[i].token, deposits[i - start].token, "_assertNewRequest : req.deposits[i].token");
            assertEq(req.deposits[i].amount, deposits[i - start].amount, "_assertNewRequest : req.deposits[i].amount");
        }
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
