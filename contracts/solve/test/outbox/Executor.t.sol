// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { SolverNetExecutor, ISolverNetExecutor } from "src/SolverNetExecutor.sol";
import { Reverter } from "test/utils/Reverter.sol";
import { ApprovalReverterERC20 } from "test/utils/ApprovalReverterERC20.sol";

contract SolverNet_Outbox_Executor_Test is TestBase {
    using AddrUtils for address;

    SolverNetExecutor internal executor;
    Reverter internal reverter;
    ApprovalReverterERC20 internal approvalReverter;

    function setUp() public override {
        super.setUp();
        executor = SolverNetExecutor(payable(outbox.executor()));
        reverter = new Reverter();
        approvalReverter = new ApprovalReverterERC20();
        approvalReverter.mint(address(executor), 1 ether);
    }

    function test_executor_reverts() public {
        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.approve(address(0), address(0), 0);

        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.execute(address(0), 0, "");

        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.transfer(address(0), address(0), 0);

        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.transferNative(address(0), 0);

        vm.expectRevert(ISolverNetExecutor.CallFailed.selector);
        vm.prank(address(outbox));
        executor.execute(address(reverter), 0, "");
    }

    function test_approve_succeeds() public {
        vm.prank(address(outbox));
        executor.approve(address(token1), user, 1 ether);

        assertEq(token1.allowance(address(executor), user), 1 ether, "allowance should be 1 ether");
    }

    function test_tryRevokeApproval_succeeds_approve_reverts() public {
        vm.startPrank(address(outbox));
        executor.approve(address(approvalReverter), user, 1 ether);
        executor.tryRevokeApproval(address(approvalReverter), user);
        vm.stopPrank();

        assertEq(approvalReverter.allowance(address(executor), user), 1 ether, "allowance should be 1 ether");
    }

    function test_execute_succeeds() public {
        token1.mint(address(executor), 1 ether);

        vm.prank(address(outbox));
        executor.execute(address(token1), 0, abi.encodeCall(IERC20.transfer, (user, 1 ether)));

        assertEq(token1.balanceOf(user), 1 ether, "balance should be 1 ether");
    }

    function test_transfer_succeeds() public {
        token1.mint(address(executor), 1 ether);

        vm.prank(address(outbox));
        executor.transfer(address(token1), user, 1 ether);

        assertEq(token1.balanceOf(user), 1 ether, "balance should be 1 ether");
    }

    function test_transferNative_succeeds() public {
        vm.deal(address(executor), 1 ether);

        vm.prank(address(outbox));
        executor.transferNative(user, 1 ether);

        assertEq(user.balance, 1 ether, "balance should be 1 ether");
    }
}
